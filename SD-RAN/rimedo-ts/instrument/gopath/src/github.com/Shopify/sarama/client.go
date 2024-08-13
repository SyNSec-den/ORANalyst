//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1
)

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Client is a generic Kafka client. It manages connections to one or more Kafka brokers.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:10
// You MUST call Close() on a client to avoid leaks, it will not be garbage-collected
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:10
// automatically when it passes out of scope. It is safe to share a client amongst many
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:10
// users, however Kafka will process requests from a single client strictly in serial,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:10
// so it is generally more efficient to use the default one client per producer/consumer.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:15
type Client interface {
	// Config returns the Config struct of the client. This struct should not be
	// altered after it has been created.
	Config() *Config

	// Controller returns the cluster controller broker. It will return a
	// locally cached value if it's available. You can call RefreshController
	// to update the cached value. Requires Kafka 0.10 or higher.
	Controller() (*Broker, error)

	// RefreshController retrieves the cluster controller from fresh metadata
	// and stores it in the local cache. Requires Kafka 0.10 or higher.
	RefreshController() (*Broker, error)

	// Brokers returns the current set of active brokers as retrieved from cluster metadata.
	Brokers() []*Broker

	// Broker returns the active Broker if available for the broker ID.
	Broker(brokerID int32) (*Broker, error)

	// Topics returns the set of available topics as retrieved from cluster metadata.
	Topics() ([]string, error)

	// Partitions returns the sorted list of all partition IDs for the given topic.
	Partitions(topic string) ([]int32, error)

	// WritablePartitions returns the sorted list of all writable partition IDs for
	// the given topic, where "writable" means "having a valid leader accepting
	// writes".
	WritablePartitions(topic string) ([]int32, error)

	// Leader returns the broker object that is the leader of the current
	// topic/partition, as determined by querying the cluster metadata.
	Leader(topic string, partitionID int32) (*Broker, error)

	// Replicas returns the set of all replica IDs for the given partition.
	Replicas(topic string, partitionID int32) ([]int32, error)

	// InSyncReplicas returns the set of all in-sync replica IDs for the given
	// partition. In-sync replicas are replicas which are fully caught up with
	// the partition leader.
	InSyncReplicas(topic string, partitionID int32) ([]int32, error)

	// OfflineReplicas returns the set of all offline replica IDs for the given
	// partition. Offline replicas are replicas which are offline
	OfflineReplicas(topic string, partitionID int32) ([]int32, error)

	// RefreshBrokers takes a list of addresses to be used as seed brokers.
	// Existing broker connections are closed and the updated list of seed brokers
	// will be used for the next metadata fetch.
	RefreshBrokers(addrs []string) error

	// RefreshMetadata takes a list of topics and queries the cluster to refresh the
	// available metadata for those topics. If no topics are provided, it will refresh
	// metadata for all topics.
	RefreshMetadata(topics ...string) error

	// GetOffset queries the cluster to get the most recent available offset at the
	// given time (in milliseconds) on the topic/partition combination.
	// Time should be OffsetOldest for the earliest available offset,
	// OffsetNewest for the offset of the message that will be produced next, or a time.
	GetOffset(topic string, partitionID int32, time int64) (int64, error)

	// Coordinator returns the coordinating broker for a consumer group. It will
	// return a locally cached value if it's available. You can call
	// RefreshCoordinator to update the cached value. This function only works on
	// Kafka 0.8.2 and higher.
	Coordinator(consumerGroup string) (*Broker, error)

	// RefreshCoordinator retrieves the coordinator for a consumer group and stores it
	// in local cache. This function only works on Kafka 0.8.2 and higher.
	RefreshCoordinator(consumerGroup string) error

	// InitProducerID retrieves information required for Idempotent Producer
	InitProducerID() (*InitProducerIDResponse, error)

	// Close shuts down all broker connections managed by this client. It is required
	// to call this function before a client object passes out of scope, as it will
	// otherwise leak memory. You must close any Producers or Consumers using a client
	// before you close the client.
	Close() error

	// Closed returns true if the client has already had Close called on it
	Closed() bool
}

const (
	// OffsetNewest stands for the log head offset, i.e. the offset that will be
	// assigned to the next message that will be produced to the partition. You
	// can send this to a client's GetOffset method to get this offset, or when
	// calling ConsumePartition to start consuming new messages.
	OffsetNewest	int64	= -1
	// OffsetOldest stands for the oldest offset available on the broker for a
	// partition. You can send this to a client's GetOffset method to get this
	// offset, or when calling ConsumePartition to start consuming from the
	// oldest offset that is still available on the broker.
	OffsetOldest	int64	= -2
)

type client struct {
	conf		*Config
	closer, closed	chan none	// for shutting down background metadata updater

	// the broker addresses given to us through the constructor are not guaranteed to be returned in
	// the cluster metadata (I *think* it only returns brokers who are currently leading partitions?)
	// so we store them separately
	seedBrokers	[]*Broker
	deadSeeds	[]*Broker

	controllerID	int32					// cluster controller broker id
	brokers		map[int32]*Broker			// maps broker ids to brokers
	metadata	map[string]map[int32]*PartitionMetadata	// maps topics to partition ids to metadata
	metadataTopics	map[string]none				// topics that need to collect metadata
	coordinators	map[string]int32			// Maps consumer group names to coordinating broker IDs

	// If the number of partitions is large, we can get some churn calling cachedPartitions,
	// so the result is cached.  It is important to update this value whenever metadata is changed
	cachedPartitionsResults	map[string][maxPartitionIndex][]int32

	lock	sync.RWMutex	// protects access to the maps that hold cluster state.
}

// NewClient creates a new Client. It connects to one of the given broker addresses
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:137
// and uses that broker to automatically fetch metadata on the rest of the kafka cluster. If metadata cannot
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:137
// be retrieved from any of the given broker addresses, the client is not created.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:140
func NewClient(addrs []string, conf *Config) (Client, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:140
	_go_fuzz_dep_.CoverTab[99946]++
											DebugLogger.Println("Initializing new client")

											if conf == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:143
		_go_fuzz_dep_.CoverTab[99951]++
												conf = NewConfig()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:144
		// _ = "end of CoverTab[99951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:145
		_go_fuzz_dep_.CoverTab[99952]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:145
		// _ = "end of CoverTab[99952]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:145
	// _ = "end of CoverTab[99946]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:145
	_go_fuzz_dep_.CoverTab[99947]++

											if err := conf.Validate(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:147
		_go_fuzz_dep_.CoverTab[99953]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:148
		// _ = "end of CoverTab[99953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:149
		_go_fuzz_dep_.CoverTab[99954]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:149
		// _ = "end of CoverTab[99954]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:149
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:149
	// _ = "end of CoverTab[99947]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:149
	_go_fuzz_dep_.CoverTab[99948]++

											if len(addrs) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:151
		_go_fuzz_dep_.CoverTab[99955]++
												return nil, ConfigurationError("You must provide at least one broker address")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:152
		// _ = "end of CoverTab[99955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:153
		_go_fuzz_dep_.CoverTab[99956]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:153
		// _ = "end of CoverTab[99956]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:153
	// _ = "end of CoverTab[99948]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:153
	_go_fuzz_dep_.CoverTab[99949]++

											client := &client{
		conf:				conf,
		closer:				make(chan none),
		closed:				make(chan none),
		brokers:			make(map[int32]*Broker),
		metadata:			make(map[string]map[int32]*PartitionMetadata),
		metadataTopics:			make(map[string]none),
		cachedPartitionsResults:	make(map[string][maxPartitionIndex][]int32),
		coordinators:			make(map[string]int32),
	}

	client.randomizeSeedBrokers(addrs)

	if conf.Metadata.Full {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:168
		_go_fuzz_dep_.CoverTab[99957]++

												err := client.RefreshMetadata()
												switch err {
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:172
			_go_fuzz_dep_.CoverTab[99958]++
													break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:173
			// _ = "end of CoverTab[99958]"
		case ErrLeaderNotAvailable, ErrReplicaNotAvailable, ErrTopicAuthorizationFailed, ErrClusterAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:174
			_go_fuzz_dep_.CoverTab[99959]++

													Logger.Println(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:176
			// _ = "end of CoverTab[99959]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:177
			_go_fuzz_dep_.CoverTab[99960]++
													close(client.closed)
													_ = client.Close()
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:180
			// _ = "end of CoverTab[99960]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:181
		// _ = "end of CoverTab[99957]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:182
		_go_fuzz_dep_.CoverTab[99961]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:182
		// _ = "end of CoverTab[99961]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:182
	// _ = "end of CoverTab[99949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:182
	_go_fuzz_dep_.CoverTab[99950]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:182
	_curRoutineNum127_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:182
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum127_)
											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
		_go_fuzz_dep_.CoverTab[99962]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
			_go_fuzz_dep_.CoverTab[99963]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum127_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
			// _ = "end of CoverTab[99963]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
		withRecover(client.backgroundMetadataUpdater)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
		// _ = "end of CoverTab[99962]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:183
	}()

											DebugLogger.Println("Successfully initialized new client")

											return client, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:187
	// _ = "end of CoverTab[99950]"
}

func (client *client) Config() *Config {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:190
	_go_fuzz_dep_.CoverTab[99964]++
											return client.conf
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:191
	// _ = "end of CoverTab[99964]"
}

func (client *client) Brokers() []*Broker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:194
	_go_fuzz_dep_.CoverTab[99965]++
											client.lock.RLock()
											defer client.lock.RUnlock()
											brokers := make([]*Broker, 0, len(client.brokers))
											for _, broker := range client.brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:198
		_go_fuzz_dep_.CoverTab[99967]++
												brokers = append(brokers, broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:199
		// _ = "end of CoverTab[99967]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:200
	// _ = "end of CoverTab[99965]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:200
	_go_fuzz_dep_.CoverTab[99966]++
											return brokers
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:201
	// _ = "end of CoverTab[99966]"
}

func (client *client) Broker(brokerID int32) (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:204
	_go_fuzz_dep_.CoverTab[99968]++
											client.lock.RLock()
											defer client.lock.RUnlock()
											broker, ok := client.brokers[brokerID]
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:208
		_go_fuzz_dep_.CoverTab[99970]++
												return nil, ErrBrokerNotFound
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:209
		// _ = "end of CoverTab[99970]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:210
		_go_fuzz_dep_.CoverTab[99971]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:210
		// _ = "end of CoverTab[99971]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:210
	// _ = "end of CoverTab[99968]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:210
	_go_fuzz_dep_.CoverTab[99969]++
											_ = broker.Open(client.conf)
											return broker, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:212
	// _ = "end of CoverTab[99969]"
}

func (client *client) InitProducerID() (*InitProducerIDResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:215
	_go_fuzz_dep_.CoverTab[99972]++
											err := ErrOutOfBrokers
											for broker := client.any(); broker != nil; broker = client.any() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:217
		_go_fuzz_dep_.CoverTab[99974]++
												var response *InitProducerIDResponse
												req := &InitProducerIDRequest{}

												response, err = broker.InitProducerID(req)
												switch err.(type) {
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:223
			_go_fuzz_dep_.CoverTab[99975]++
													return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:224
			// _ = "end of CoverTab[99975]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:225
			_go_fuzz_dep_.CoverTab[99976]++

													Logger.Printf("Client got error from broker %d when issuing InitProducerID : %v\n", broker.ID(), err)
													_ = broker.Close()
													client.deregisterBroker(broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:229
			// _ = "end of CoverTab[99976]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:230
		// _ = "end of CoverTab[99974]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:231
	// _ = "end of CoverTab[99972]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:231
	_go_fuzz_dep_.CoverTab[99973]++

											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:233
	// _ = "end of CoverTab[99973]"
}

func (client *client) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:236
	_go_fuzz_dep_.CoverTab[99977]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:237
		_go_fuzz_dep_.CoverTab[99981]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:240
		Logger.Printf("Close() called on already closed client")
												return ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:241
		// _ = "end of CoverTab[99981]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:242
		_go_fuzz_dep_.CoverTab[99982]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:242
		// _ = "end of CoverTab[99982]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:242
	// _ = "end of CoverTab[99977]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:242
	_go_fuzz_dep_.CoverTab[99978]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:245
	close(client.closer)
	<-client.closed

	client.lock.Lock()
	defer client.lock.Unlock()
	DebugLogger.Println("Closing Client")

	for _, broker := range client.brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:252
		_go_fuzz_dep_.CoverTab[99983]++
												safeAsyncClose(broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:253
		// _ = "end of CoverTab[99983]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:254
	// _ = "end of CoverTab[99978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:254
	_go_fuzz_dep_.CoverTab[99979]++

											for _, broker := range client.seedBrokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:256
		_go_fuzz_dep_.CoverTab[99984]++
												safeAsyncClose(broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:257
		// _ = "end of CoverTab[99984]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:258
	// _ = "end of CoverTab[99979]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:258
	_go_fuzz_dep_.CoverTab[99980]++

											client.brokers = nil
											client.metadata = nil
											client.metadataTopics = nil

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:264
	// _ = "end of CoverTab[99980]"
}

func (client *client) Closed() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:267
	_go_fuzz_dep_.CoverTab[99985]++
											client.lock.RLock()
											defer client.lock.RUnlock()

											return client.brokers == nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:271
	// _ = "end of CoverTab[99985]"
}

func (client *client) Topics() ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:274
	_go_fuzz_dep_.CoverTab[99986]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:275
		_go_fuzz_dep_.CoverTab[99989]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:276
		// _ = "end of CoverTab[99989]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:277
		_go_fuzz_dep_.CoverTab[99990]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:277
		// _ = "end of CoverTab[99990]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:277
	// _ = "end of CoverTab[99986]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:277
	_go_fuzz_dep_.CoverTab[99987]++

											client.lock.RLock()
											defer client.lock.RUnlock()

											ret := make([]string, 0, len(client.metadata))
											for topic := range client.metadata {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:283
		_go_fuzz_dep_.CoverTab[99991]++
												ret = append(ret, topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:284
		// _ = "end of CoverTab[99991]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:285
	// _ = "end of CoverTab[99987]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:285
	_go_fuzz_dep_.CoverTab[99988]++

											return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:287
	// _ = "end of CoverTab[99988]"
}

func (client *client) MetadataTopics() ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:290
	_go_fuzz_dep_.CoverTab[99992]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:291
		_go_fuzz_dep_.CoverTab[99995]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:292
		// _ = "end of CoverTab[99995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:293
		_go_fuzz_dep_.CoverTab[99996]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:293
		// _ = "end of CoverTab[99996]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:293
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:293
	// _ = "end of CoverTab[99992]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:293
	_go_fuzz_dep_.CoverTab[99993]++

											client.lock.RLock()
											defer client.lock.RUnlock()

											ret := make([]string, 0, len(client.metadataTopics))
											for topic := range client.metadataTopics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:299
		_go_fuzz_dep_.CoverTab[99997]++
												ret = append(ret, topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:300
		// _ = "end of CoverTab[99997]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:301
	// _ = "end of CoverTab[99993]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:301
	_go_fuzz_dep_.CoverTab[99994]++

											return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:303
	// _ = "end of CoverTab[99994]"
}

func (client *client) Partitions(topic string) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:306
	_go_fuzz_dep_.CoverTab[99998]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:307
		_go_fuzz_dep_.CoverTab[100002]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:308
		// _ = "end of CoverTab[100002]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:309
		_go_fuzz_dep_.CoverTab[100003]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:309
		// _ = "end of CoverTab[100003]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:309
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:309
	// _ = "end of CoverTab[99998]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:309
	_go_fuzz_dep_.CoverTab[99999]++

											partitions := client.cachedPartitions(topic, allPartitions)

											if len(partitions) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:313
		_go_fuzz_dep_.CoverTab[100004]++
												err := client.RefreshMetadata(topic)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:315
			_go_fuzz_dep_.CoverTab[100006]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:316
			// _ = "end of CoverTab[100006]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:317
			_go_fuzz_dep_.CoverTab[100007]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:317
			// _ = "end of CoverTab[100007]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:317
		// _ = "end of CoverTab[100004]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:317
		_go_fuzz_dep_.CoverTab[100005]++
												partitions = client.cachedPartitions(topic, allPartitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:318
		// _ = "end of CoverTab[100005]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:319
		_go_fuzz_dep_.CoverTab[100008]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:319
		// _ = "end of CoverTab[100008]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:319
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:319
	// _ = "end of CoverTab[99999]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:319
	_go_fuzz_dep_.CoverTab[100000]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:322
	if len(partitions) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:322
		_go_fuzz_dep_.CoverTab[100009]++
												return nil, ErrUnknownTopicOrPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:323
		// _ = "end of CoverTab[100009]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:324
		_go_fuzz_dep_.CoverTab[100010]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:324
		// _ = "end of CoverTab[100010]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:324
	// _ = "end of CoverTab[100000]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:324
	_go_fuzz_dep_.CoverTab[100001]++

											return partitions, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:326
	// _ = "end of CoverTab[100001]"
}

func (client *client) WritablePartitions(topic string) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:329
	_go_fuzz_dep_.CoverTab[100011]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:330
		_go_fuzz_dep_.CoverTab[100015]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:331
		// _ = "end of CoverTab[100015]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:332
		_go_fuzz_dep_.CoverTab[100016]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:332
		// _ = "end of CoverTab[100016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:332
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:332
	// _ = "end of CoverTab[100011]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:332
	_go_fuzz_dep_.CoverTab[100012]++

											partitions := client.cachedPartitions(topic, writablePartitions)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:342
	if len(partitions) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:342
		_go_fuzz_dep_.CoverTab[100017]++
												err := client.RefreshMetadata(topic)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:344
			_go_fuzz_dep_.CoverTab[100019]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:345
			// _ = "end of CoverTab[100019]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:346
			_go_fuzz_dep_.CoverTab[100020]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:346
			// _ = "end of CoverTab[100020]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:346
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:346
		// _ = "end of CoverTab[100017]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:346
		_go_fuzz_dep_.CoverTab[100018]++
												partitions = client.cachedPartitions(topic, writablePartitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:347
		// _ = "end of CoverTab[100018]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:348
		_go_fuzz_dep_.CoverTab[100021]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:348
		// _ = "end of CoverTab[100021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:348
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:348
	// _ = "end of CoverTab[100012]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:348
	_go_fuzz_dep_.CoverTab[100013]++

											if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:350
		_go_fuzz_dep_.CoverTab[100022]++
												return nil, ErrUnknownTopicOrPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:351
		// _ = "end of CoverTab[100022]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:352
		_go_fuzz_dep_.CoverTab[100023]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:352
		// _ = "end of CoverTab[100023]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:352
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:352
	// _ = "end of CoverTab[100013]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:352
	_go_fuzz_dep_.CoverTab[100014]++

											return partitions, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:354
	// _ = "end of CoverTab[100014]"
}

func (client *client) Replicas(topic string, partitionID int32) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:357
	_go_fuzz_dep_.CoverTab[100024]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:358
		_go_fuzz_dep_.CoverTab[100029]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:359
		// _ = "end of CoverTab[100029]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:360
		_go_fuzz_dep_.CoverTab[100030]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:360
		// _ = "end of CoverTab[100030]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:360
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:360
	// _ = "end of CoverTab[100024]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:360
	_go_fuzz_dep_.CoverTab[100025]++

											metadata := client.cachedMetadata(topic, partitionID)

											if metadata == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:364
		_go_fuzz_dep_.CoverTab[100031]++
												err := client.RefreshMetadata(topic)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:366
			_go_fuzz_dep_.CoverTab[100033]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:367
			// _ = "end of CoverTab[100033]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:368
			_go_fuzz_dep_.CoverTab[100034]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:368
			// _ = "end of CoverTab[100034]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:368
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:368
		// _ = "end of CoverTab[100031]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:368
		_go_fuzz_dep_.CoverTab[100032]++
												metadata = client.cachedMetadata(topic, partitionID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:369
		// _ = "end of CoverTab[100032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:370
		_go_fuzz_dep_.CoverTab[100035]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:370
		// _ = "end of CoverTab[100035]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:370
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:370
	// _ = "end of CoverTab[100025]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:370
	_go_fuzz_dep_.CoverTab[100026]++

											if metadata == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:372
		_go_fuzz_dep_.CoverTab[100036]++
												return nil, ErrUnknownTopicOrPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:373
		// _ = "end of CoverTab[100036]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:374
		_go_fuzz_dep_.CoverTab[100037]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:374
		// _ = "end of CoverTab[100037]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:374
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:374
	// _ = "end of CoverTab[100026]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:374
	_go_fuzz_dep_.CoverTab[100027]++

											if metadata.Err == ErrReplicaNotAvailable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:376
		_go_fuzz_dep_.CoverTab[100038]++
												return dupInt32Slice(metadata.Replicas), metadata.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:377
		// _ = "end of CoverTab[100038]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:378
		_go_fuzz_dep_.CoverTab[100039]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:378
		// _ = "end of CoverTab[100039]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:378
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:378
	// _ = "end of CoverTab[100027]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:378
	_go_fuzz_dep_.CoverTab[100028]++
											return dupInt32Slice(metadata.Replicas), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:379
	// _ = "end of CoverTab[100028]"
}

func (client *client) InSyncReplicas(topic string, partitionID int32) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:382
	_go_fuzz_dep_.CoverTab[100040]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:383
		_go_fuzz_dep_.CoverTab[100045]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:384
		// _ = "end of CoverTab[100045]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:385
		_go_fuzz_dep_.CoverTab[100046]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:385
		// _ = "end of CoverTab[100046]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:385
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:385
	// _ = "end of CoverTab[100040]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:385
	_go_fuzz_dep_.CoverTab[100041]++

											metadata := client.cachedMetadata(topic, partitionID)

											if metadata == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:389
		_go_fuzz_dep_.CoverTab[100047]++
												err := client.RefreshMetadata(topic)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:391
			_go_fuzz_dep_.CoverTab[100049]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:392
			// _ = "end of CoverTab[100049]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:393
			_go_fuzz_dep_.CoverTab[100050]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:393
			// _ = "end of CoverTab[100050]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:393
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:393
		// _ = "end of CoverTab[100047]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:393
		_go_fuzz_dep_.CoverTab[100048]++
												metadata = client.cachedMetadata(topic, partitionID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:394
		// _ = "end of CoverTab[100048]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:395
		_go_fuzz_dep_.CoverTab[100051]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:395
		// _ = "end of CoverTab[100051]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:395
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:395
	// _ = "end of CoverTab[100041]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:395
	_go_fuzz_dep_.CoverTab[100042]++

											if metadata == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:397
		_go_fuzz_dep_.CoverTab[100052]++
												return nil, ErrUnknownTopicOrPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:398
		// _ = "end of CoverTab[100052]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:399
		_go_fuzz_dep_.CoverTab[100053]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:399
		// _ = "end of CoverTab[100053]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:399
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:399
	// _ = "end of CoverTab[100042]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:399
	_go_fuzz_dep_.CoverTab[100043]++

											if metadata.Err == ErrReplicaNotAvailable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:401
		_go_fuzz_dep_.CoverTab[100054]++
												return dupInt32Slice(metadata.Isr), metadata.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:402
		// _ = "end of CoverTab[100054]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:403
		_go_fuzz_dep_.CoverTab[100055]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:403
		// _ = "end of CoverTab[100055]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:403
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:403
	// _ = "end of CoverTab[100043]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:403
	_go_fuzz_dep_.CoverTab[100044]++
											return dupInt32Slice(metadata.Isr), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:404
	// _ = "end of CoverTab[100044]"
}

func (client *client) OfflineReplicas(topic string, partitionID int32) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:407
	_go_fuzz_dep_.CoverTab[100056]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:408
		_go_fuzz_dep_.CoverTab[100061]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:409
		// _ = "end of CoverTab[100061]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:410
		_go_fuzz_dep_.CoverTab[100062]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:410
		// _ = "end of CoverTab[100062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:410
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:410
	// _ = "end of CoverTab[100056]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:410
	_go_fuzz_dep_.CoverTab[100057]++

											metadata := client.cachedMetadata(topic, partitionID)

											if metadata == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:414
		_go_fuzz_dep_.CoverTab[100063]++
												err := client.RefreshMetadata(topic)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:416
			_go_fuzz_dep_.CoverTab[100065]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:417
			// _ = "end of CoverTab[100065]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:418
			_go_fuzz_dep_.CoverTab[100066]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:418
			// _ = "end of CoverTab[100066]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:418
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:418
		// _ = "end of CoverTab[100063]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:418
		_go_fuzz_dep_.CoverTab[100064]++
												metadata = client.cachedMetadata(topic, partitionID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:419
		// _ = "end of CoverTab[100064]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:420
		_go_fuzz_dep_.CoverTab[100067]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:420
		// _ = "end of CoverTab[100067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:420
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:420
	// _ = "end of CoverTab[100057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:420
	_go_fuzz_dep_.CoverTab[100058]++

											if metadata == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:422
		_go_fuzz_dep_.CoverTab[100068]++
												return nil, ErrUnknownTopicOrPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:423
		// _ = "end of CoverTab[100068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:424
		_go_fuzz_dep_.CoverTab[100069]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:424
		// _ = "end of CoverTab[100069]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:424
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:424
	// _ = "end of CoverTab[100058]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:424
	_go_fuzz_dep_.CoverTab[100059]++

											if metadata.Err == ErrReplicaNotAvailable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:426
		_go_fuzz_dep_.CoverTab[100070]++
												return dupInt32Slice(metadata.OfflineReplicas), metadata.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:427
		// _ = "end of CoverTab[100070]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:428
		_go_fuzz_dep_.CoverTab[100071]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:428
		// _ = "end of CoverTab[100071]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:428
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:428
	// _ = "end of CoverTab[100059]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:428
	_go_fuzz_dep_.CoverTab[100060]++
											return dupInt32Slice(metadata.OfflineReplicas), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:429
	// _ = "end of CoverTab[100060]"
}

func (client *client) Leader(topic string, partitionID int32) (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:432
	_go_fuzz_dep_.CoverTab[100072]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:433
		_go_fuzz_dep_.CoverTab[100075]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:434
		// _ = "end of CoverTab[100075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:435
		_go_fuzz_dep_.CoverTab[100076]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:435
		// _ = "end of CoverTab[100076]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:435
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:435
	// _ = "end of CoverTab[100072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:435
	_go_fuzz_dep_.CoverTab[100073]++

											leader, err := client.cachedLeader(topic, partitionID)

											if leader == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:439
		_go_fuzz_dep_.CoverTab[100077]++
												err = client.RefreshMetadata(topic)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:441
			_go_fuzz_dep_.CoverTab[100079]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:442
			// _ = "end of CoverTab[100079]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:443
			_go_fuzz_dep_.CoverTab[100080]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:443
			// _ = "end of CoverTab[100080]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:443
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:443
		// _ = "end of CoverTab[100077]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:443
		_go_fuzz_dep_.CoverTab[100078]++
												leader, err = client.cachedLeader(topic, partitionID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:444
		// _ = "end of CoverTab[100078]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:445
		_go_fuzz_dep_.CoverTab[100081]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:445
		// _ = "end of CoverTab[100081]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:445
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:445
	// _ = "end of CoverTab[100073]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:445
	_go_fuzz_dep_.CoverTab[100074]++

											return leader, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:447
	// _ = "end of CoverTab[100074]"
}

func (client *client) RefreshBrokers(addrs []string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:450
	_go_fuzz_dep_.CoverTab[100082]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:451
		_go_fuzz_dep_.CoverTab[100085]++
												return ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:452
		// _ = "end of CoverTab[100085]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:453
		_go_fuzz_dep_.CoverTab[100086]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:453
		// _ = "end of CoverTab[100086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:453
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:453
	// _ = "end of CoverTab[100082]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:453
	_go_fuzz_dep_.CoverTab[100083]++

											client.lock.Lock()
											defer client.lock.Unlock()

											for _, broker := range client.brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:458
		_go_fuzz_dep_.CoverTab[100087]++
												_ = broker.Close()
												delete(client.brokers, broker.ID())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:460
		// _ = "end of CoverTab[100087]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:461
	// _ = "end of CoverTab[100083]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:461
	_go_fuzz_dep_.CoverTab[100084]++

											client.seedBrokers = nil
											client.deadSeeds = nil

											client.randomizeSeedBrokers(addrs)

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:468
	// _ = "end of CoverTab[100084]"
}

func (client *client) RefreshMetadata(topics ...string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:471
	_go_fuzz_dep_.CoverTab[100088]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:472
		_go_fuzz_dep_.CoverTab[100092]++
												return ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:473
		// _ = "end of CoverTab[100092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:474
		_go_fuzz_dep_.CoverTab[100093]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:474
		// _ = "end of CoverTab[100093]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:474
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:474
	// _ = "end of CoverTab[100088]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:474
	_go_fuzz_dep_.CoverTab[100089]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:479
	for _, topic := range topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:479
		_go_fuzz_dep_.CoverTab[100094]++
												if topic == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:480
			_go_fuzz_dep_.CoverTab[100095]++
													return ErrInvalidTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:481
			// _ = "end of CoverTab[100095]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:482
			_go_fuzz_dep_.CoverTab[100096]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:482
			// _ = "end of CoverTab[100096]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:482
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:482
		// _ = "end of CoverTab[100094]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:483
	// _ = "end of CoverTab[100089]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:483
	_go_fuzz_dep_.CoverTab[100090]++

											deadline := time.Time{}
											if client.conf.Metadata.Timeout > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:486
		_go_fuzz_dep_.CoverTab[100097]++
												deadline = time.Now().Add(client.conf.Metadata.Timeout)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:487
		// _ = "end of CoverTab[100097]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:488
		_go_fuzz_dep_.CoverTab[100098]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:488
		// _ = "end of CoverTab[100098]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:488
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:488
	// _ = "end of CoverTab[100090]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:488
	_go_fuzz_dep_.CoverTab[100091]++
											return client.tryRefreshMetadata(topics, client.conf.Metadata.Retry.Max, deadline)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:489
	// _ = "end of CoverTab[100091]"
}

func (client *client) GetOffset(topic string, partitionID int32, time int64) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:492
	_go_fuzz_dep_.CoverTab[100099]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:493
		_go_fuzz_dep_.CoverTab[100102]++
												return -1, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:494
		// _ = "end of CoverTab[100102]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:495
		_go_fuzz_dep_.CoverTab[100103]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:495
		// _ = "end of CoverTab[100103]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:495
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:495
	// _ = "end of CoverTab[100099]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:495
	_go_fuzz_dep_.CoverTab[100100]++

											offset, err := client.getOffset(topic, partitionID, time)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:498
		_go_fuzz_dep_.CoverTab[100104]++
												if err := client.RefreshMetadata(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:499
			_go_fuzz_dep_.CoverTab[100106]++
													return -1, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:500
			// _ = "end of CoverTab[100106]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:501
			_go_fuzz_dep_.CoverTab[100107]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:501
			// _ = "end of CoverTab[100107]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:501
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:501
		// _ = "end of CoverTab[100104]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:501
		_go_fuzz_dep_.CoverTab[100105]++
												return client.getOffset(topic, partitionID, time)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:502
		// _ = "end of CoverTab[100105]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:503
		_go_fuzz_dep_.CoverTab[100108]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:503
		// _ = "end of CoverTab[100108]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:503
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:503
	// _ = "end of CoverTab[100100]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:503
	_go_fuzz_dep_.CoverTab[100101]++

											return offset, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:505
	// _ = "end of CoverTab[100101]"
}

func (client *client) Controller() (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:508
	_go_fuzz_dep_.CoverTab[100109]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:509
		_go_fuzz_dep_.CoverTab[100114]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:510
		// _ = "end of CoverTab[100114]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:511
		_go_fuzz_dep_.CoverTab[100115]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:511
		// _ = "end of CoverTab[100115]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:511
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:511
	// _ = "end of CoverTab[100109]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:511
	_go_fuzz_dep_.CoverTab[100110]++

											if !client.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:513
		_go_fuzz_dep_.CoverTab[100116]++
												return nil, ErrUnsupportedVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:514
		// _ = "end of CoverTab[100116]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:515
		_go_fuzz_dep_.CoverTab[100117]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:515
		// _ = "end of CoverTab[100117]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:515
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:515
	// _ = "end of CoverTab[100110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:515
	_go_fuzz_dep_.CoverTab[100111]++

											controller := client.cachedController()
											if controller == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:518
		_go_fuzz_dep_.CoverTab[100118]++
												if err := client.refreshMetadata(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:519
			_go_fuzz_dep_.CoverTab[100120]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:520
			// _ = "end of CoverTab[100120]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:521
			_go_fuzz_dep_.CoverTab[100121]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:521
			// _ = "end of CoverTab[100121]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:521
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:521
		// _ = "end of CoverTab[100118]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:521
		_go_fuzz_dep_.CoverTab[100119]++
												controller = client.cachedController()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:522
		// _ = "end of CoverTab[100119]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:523
		_go_fuzz_dep_.CoverTab[100122]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:523
		// _ = "end of CoverTab[100122]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:523
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:523
	// _ = "end of CoverTab[100111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:523
	_go_fuzz_dep_.CoverTab[100112]++

											if controller == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:525
		_go_fuzz_dep_.CoverTab[100123]++
												return nil, ErrControllerNotAvailable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:526
		// _ = "end of CoverTab[100123]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:527
		_go_fuzz_dep_.CoverTab[100124]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:527
		// _ = "end of CoverTab[100124]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:527
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:527
	// _ = "end of CoverTab[100112]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:527
	_go_fuzz_dep_.CoverTab[100113]++

											_ = controller.Open(client.conf)
											return controller, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:530
	// _ = "end of CoverTab[100113]"
}

// deregisterController removes the cached controllerID
func (client *client) deregisterController() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:534
	_go_fuzz_dep_.CoverTab[100125]++
											client.lock.Lock()
											defer client.lock.Unlock()
											delete(client.brokers, client.controllerID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:537
	// _ = "end of CoverTab[100125]"
}

// RefreshController retrieves the cluster controller from fresh metadata
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:540
// and stores it in the local cache. Requires Kafka 0.10 or higher.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:542
func (client *client) RefreshController() (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:542
	_go_fuzz_dep_.CoverTab[100126]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:543
		_go_fuzz_dep_.CoverTab[100130]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:544
		// _ = "end of CoverTab[100130]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:545
		_go_fuzz_dep_.CoverTab[100131]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:545
		// _ = "end of CoverTab[100131]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:545
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:545
	// _ = "end of CoverTab[100126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:545
	_go_fuzz_dep_.CoverTab[100127]++

											client.deregisterController()

											if err := client.refreshMetadata(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:549
		_go_fuzz_dep_.CoverTab[100132]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:550
		// _ = "end of CoverTab[100132]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:551
		_go_fuzz_dep_.CoverTab[100133]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:551
		// _ = "end of CoverTab[100133]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:551
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:551
	// _ = "end of CoverTab[100127]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:551
	_go_fuzz_dep_.CoverTab[100128]++

											controller := client.cachedController()
											if controller == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:554
		_go_fuzz_dep_.CoverTab[100134]++
												return nil, ErrControllerNotAvailable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:555
		// _ = "end of CoverTab[100134]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:556
		_go_fuzz_dep_.CoverTab[100135]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:556
		// _ = "end of CoverTab[100135]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:556
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:556
	// _ = "end of CoverTab[100128]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:556
	_go_fuzz_dep_.CoverTab[100129]++

											_ = controller.Open(client.conf)
											return controller, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:559
	// _ = "end of CoverTab[100129]"
}

func (client *client) Coordinator(consumerGroup string) (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:562
	_go_fuzz_dep_.CoverTab[100136]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:563
		_go_fuzz_dep_.CoverTab[100140]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:564
		// _ = "end of CoverTab[100140]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:565
		_go_fuzz_dep_.CoverTab[100141]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:565
		// _ = "end of CoverTab[100141]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:565
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:565
	// _ = "end of CoverTab[100136]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:565
	_go_fuzz_dep_.CoverTab[100137]++

											coordinator := client.cachedCoordinator(consumerGroup)

											if coordinator == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:569
		_go_fuzz_dep_.CoverTab[100142]++
												if err := client.RefreshCoordinator(consumerGroup); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:570
			_go_fuzz_dep_.CoverTab[100144]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:571
			// _ = "end of CoverTab[100144]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:572
			_go_fuzz_dep_.CoverTab[100145]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:572
			// _ = "end of CoverTab[100145]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:572
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:572
		// _ = "end of CoverTab[100142]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:572
		_go_fuzz_dep_.CoverTab[100143]++
												coordinator = client.cachedCoordinator(consumerGroup)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:573
		// _ = "end of CoverTab[100143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:574
		_go_fuzz_dep_.CoverTab[100146]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:574
		// _ = "end of CoverTab[100146]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:574
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:574
	// _ = "end of CoverTab[100137]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:574
	_go_fuzz_dep_.CoverTab[100138]++

											if coordinator == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:576
		_go_fuzz_dep_.CoverTab[100147]++
												return nil, ErrConsumerCoordinatorNotAvailable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:577
		// _ = "end of CoverTab[100147]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:578
		_go_fuzz_dep_.CoverTab[100148]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:578
		// _ = "end of CoverTab[100148]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:578
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:578
	// _ = "end of CoverTab[100138]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:578
	_go_fuzz_dep_.CoverTab[100139]++

											_ = coordinator.Open(client.conf)
											return coordinator, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:581
	// _ = "end of CoverTab[100139]"
}

func (client *client) RefreshCoordinator(consumerGroup string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:584
	_go_fuzz_dep_.CoverTab[100149]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:585
		_go_fuzz_dep_.CoverTab[100152]++
												return ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:586
		// _ = "end of CoverTab[100152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:587
		_go_fuzz_dep_.CoverTab[100153]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:587
		// _ = "end of CoverTab[100153]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:587
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:587
	// _ = "end of CoverTab[100149]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:587
	_go_fuzz_dep_.CoverTab[100150]++

											response, err := client.getConsumerMetadata(consumerGroup, client.conf.Metadata.Retry.Max)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:590
		_go_fuzz_dep_.CoverTab[100154]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:591
		// _ = "end of CoverTab[100154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:592
		_go_fuzz_dep_.CoverTab[100155]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:592
		// _ = "end of CoverTab[100155]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:592
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:592
	// _ = "end of CoverTab[100150]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:592
	_go_fuzz_dep_.CoverTab[100151]++

											client.lock.Lock()
											defer client.lock.Unlock()
											client.registerBroker(response.Coordinator)
											client.coordinators[consumerGroup] = response.Coordinator.ID()
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:598
	// _ = "end of CoverTab[100151]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:603
func (client *client) randomizeSeedBrokers(addrs []string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:603
	_go_fuzz_dep_.CoverTab[100156]++
											random := rand.New(rand.NewSource(time.Now().UnixNano()))
											for _, index := range random.Perm(len(addrs)) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:605
		_go_fuzz_dep_.CoverTab[100157]++
												client.seedBrokers = append(client.seedBrokers, NewBroker(addrs[index]))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:606
		// _ = "end of CoverTab[100157]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:607
	// _ = "end of CoverTab[100156]"
}

func (client *client) updateBroker(brokers []*Broker) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:610
	_go_fuzz_dep_.CoverTab[100158]++
											currentBroker := make(map[int32]*Broker, len(brokers))

											for _, broker := range brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:613
		_go_fuzz_dep_.CoverTab[100160]++
												currentBroker[broker.ID()] = broker
												if client.brokers[broker.ID()] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:615
			_go_fuzz_dep_.CoverTab[100161]++
													client.brokers[broker.ID()] = broker
													DebugLogger.Printf("client/brokers registered new broker #%d at %s", broker.ID(), broker.Addr())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:617
			// _ = "end of CoverTab[100161]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:618
			_go_fuzz_dep_.CoverTab[100162]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:618
			if broker.Addr() != client.brokers[broker.ID()].Addr() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:618
				_go_fuzz_dep_.CoverTab[100163]++
														safeAsyncClose(client.brokers[broker.ID()])
														client.brokers[broker.ID()] = broker
														Logger.Printf("client/brokers replaced registered broker #%d with %s", broker.ID(), broker.Addr())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:621
				// _ = "end of CoverTab[100163]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:622
				_go_fuzz_dep_.CoverTab[100164]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:622
				// _ = "end of CoverTab[100164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:622
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:622
			// _ = "end of CoverTab[100162]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:622
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:622
		// _ = "end of CoverTab[100160]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:623
	// _ = "end of CoverTab[100158]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:623
	_go_fuzz_dep_.CoverTab[100159]++

											for id, broker := range client.brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:625
		_go_fuzz_dep_.CoverTab[100165]++
												if _, exist := currentBroker[id]; !exist {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:626
			_go_fuzz_dep_.CoverTab[100166]++
													safeAsyncClose(broker)
													delete(client.brokers, id)
													Logger.Printf("client/broker remove invalid broker #%d with %s", broker.ID(), broker.Addr())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:629
			// _ = "end of CoverTab[100166]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:630
			_go_fuzz_dep_.CoverTab[100167]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:630
			// _ = "end of CoverTab[100167]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:630
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:630
		// _ = "end of CoverTab[100165]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:631
	// _ = "end of CoverTab[100159]"
}

// registerBroker makes sure a broker received by a Metadata or Coordinator request is registered
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:634
// in the brokers map. It returns the broker that is registered, which may be the provided broker,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:634
// or a previously registered Broker instance. You must hold the write lock before calling this function.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:637
func (client *client) registerBroker(broker *Broker) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:637
	_go_fuzz_dep_.CoverTab[100168]++
											if client.brokers == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:638
		_go_fuzz_dep_.CoverTab[100170]++
												Logger.Printf("cannot register broker #%d at %s, client already closed", broker.ID(), broker.Addr())
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:640
		// _ = "end of CoverTab[100170]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:641
		_go_fuzz_dep_.CoverTab[100171]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:641
		// _ = "end of CoverTab[100171]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:641
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:641
	// _ = "end of CoverTab[100168]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:641
	_go_fuzz_dep_.CoverTab[100169]++

											if client.brokers[broker.ID()] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:643
		_go_fuzz_dep_.CoverTab[100172]++
												client.brokers[broker.ID()] = broker
												DebugLogger.Printf("client/brokers registered new broker #%d at %s", broker.ID(), broker.Addr())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:645
		// _ = "end of CoverTab[100172]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:646
		_go_fuzz_dep_.CoverTab[100173]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:646
		if broker.Addr() != client.brokers[broker.ID()].Addr() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:646
			_go_fuzz_dep_.CoverTab[100174]++
													safeAsyncClose(client.brokers[broker.ID()])
													client.brokers[broker.ID()] = broker
													Logger.Printf("client/brokers replaced registered broker #%d with %s", broker.ID(), broker.Addr())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:649
			// _ = "end of CoverTab[100174]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:650
			_go_fuzz_dep_.CoverTab[100175]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:650
			// _ = "end of CoverTab[100175]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:650
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:650
		// _ = "end of CoverTab[100173]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:650
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:650
	// _ = "end of CoverTab[100169]"
}

// deregisterBroker removes a broker from the seedsBroker list, and if it's
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:653
// not the seedbroker, removes it from brokers map completely.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:655
func (client *client) deregisterBroker(broker *Broker) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:655
	_go_fuzz_dep_.CoverTab[100176]++
											client.lock.Lock()
											defer client.lock.Unlock()

											if len(client.seedBrokers) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:659
		_go_fuzz_dep_.CoverTab[100177]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:659
		return broker == client.seedBrokers[0]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:659
		// _ = "end of CoverTab[100177]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:659
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:659
		_go_fuzz_dep_.CoverTab[100178]++
												client.deadSeeds = append(client.deadSeeds, broker)
												client.seedBrokers = client.seedBrokers[1:]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:661
		// _ = "end of CoverTab[100178]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:662
		_go_fuzz_dep_.CoverTab[100179]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:667
		DebugLogger.Printf("client/brokers deregistered broker #%d at %s", broker.ID(), broker.Addr())
												delete(client.brokers, broker.ID())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:668
		// _ = "end of CoverTab[100179]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:669
	// _ = "end of CoverTab[100176]"
}

func (client *client) resurrectDeadBrokers() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:672
	_go_fuzz_dep_.CoverTab[100180]++
											client.lock.Lock()
											defer client.lock.Unlock()

											Logger.Printf("client/brokers resurrecting %d dead seed brokers", len(client.deadSeeds))
											client.seedBrokers = append(client.seedBrokers, client.deadSeeds...)
											client.deadSeeds = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:678
	// _ = "end of CoverTab[100180]"
}

func (client *client) any() *Broker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:681
	_go_fuzz_dep_.CoverTab[100181]++
											client.lock.RLock()
											defer client.lock.RUnlock()

											if len(client.seedBrokers) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:685
		_go_fuzz_dep_.CoverTab[100184]++
												_ = client.seedBrokers[0].Open(client.conf)
												return client.seedBrokers[0]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:687
		// _ = "end of CoverTab[100184]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:688
		_go_fuzz_dep_.CoverTab[100185]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:688
		// _ = "end of CoverTab[100185]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:688
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:688
	// _ = "end of CoverTab[100181]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:688
	_go_fuzz_dep_.CoverTab[100182]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:691
	for _, broker := range client.brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:691
		_go_fuzz_dep_.CoverTab[100186]++
												_ = broker.Open(client.conf)
												return broker
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:693
		// _ = "end of CoverTab[100186]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:694
	// _ = "end of CoverTab[100182]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:694
	_go_fuzz_dep_.CoverTab[100183]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:696
	// _ = "end of CoverTab[100183]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:701
type partitionType int

const (
	allPartitions	partitionType	= iota
	writablePartitions

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:708
	// Ensure this is the last partition type value
											maxPartitionIndex
)

func (client *client) cachedMetadata(topic string, partitionID int32) *PartitionMetadata {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:712
	_go_fuzz_dep_.CoverTab[100187]++
											client.lock.RLock()
											defer client.lock.RUnlock()

											partitions := client.metadata[topic]
											if partitions != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:717
		_go_fuzz_dep_.CoverTab[100189]++
												return partitions[partitionID]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:718
		// _ = "end of CoverTab[100189]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:719
		_go_fuzz_dep_.CoverTab[100190]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:719
		// _ = "end of CoverTab[100190]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:719
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:719
	// _ = "end of CoverTab[100187]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:719
	_go_fuzz_dep_.CoverTab[100188]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:721
	// _ = "end of CoverTab[100188]"
}

func (client *client) cachedPartitions(topic string, partitionSet partitionType) []int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:724
	_go_fuzz_dep_.CoverTab[100191]++
											client.lock.RLock()
											defer client.lock.RUnlock()

											partitions, exists := client.cachedPartitionsResults[topic]

											if !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:730
		_go_fuzz_dep_.CoverTab[100193]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:731
		// _ = "end of CoverTab[100193]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:732
		_go_fuzz_dep_.CoverTab[100194]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:732
		// _ = "end of CoverTab[100194]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:732
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:732
	// _ = "end of CoverTab[100191]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:732
	_go_fuzz_dep_.CoverTab[100192]++
											return partitions[partitionSet]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:733
	// _ = "end of CoverTab[100192]"
}

func (client *client) setPartitionCache(topic string, partitionSet partitionType) []int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:736
	_go_fuzz_dep_.CoverTab[100195]++
											partitions := client.metadata[topic]

											if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:739
		_go_fuzz_dep_.CoverTab[100198]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:740
		// _ = "end of CoverTab[100198]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:741
		_go_fuzz_dep_.CoverTab[100199]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:741
		// _ = "end of CoverTab[100199]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:741
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:741
	// _ = "end of CoverTab[100195]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:741
	_go_fuzz_dep_.CoverTab[100196]++

											ret := make([]int32, 0, len(partitions))
											for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:744
		_go_fuzz_dep_.CoverTab[100200]++
												if partitionSet == writablePartitions && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:745
			_go_fuzz_dep_.CoverTab[100202]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:745
			return partition.Err == ErrLeaderNotAvailable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:745
			// _ = "end of CoverTab[100202]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:745
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:745
			_go_fuzz_dep_.CoverTab[100203]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:746
			// _ = "end of CoverTab[100203]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:747
			_go_fuzz_dep_.CoverTab[100204]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:747
			// _ = "end of CoverTab[100204]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:747
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:747
		// _ = "end of CoverTab[100200]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:747
		_go_fuzz_dep_.CoverTab[100201]++
												ret = append(ret, partition.ID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:748
		// _ = "end of CoverTab[100201]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:749
	// _ = "end of CoverTab[100196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:749
	_go_fuzz_dep_.CoverTab[100197]++

											sort.Sort(int32Slice(ret))
											return ret
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:752
	// _ = "end of CoverTab[100197]"
}

func (client *client) cachedLeader(topic string, partitionID int32) (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:755
	_go_fuzz_dep_.CoverTab[100205]++
											client.lock.RLock()
											defer client.lock.RUnlock()

											partitions := client.metadata[topic]
											if partitions != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:760
		_go_fuzz_dep_.CoverTab[100207]++
												metadata, ok := partitions[partitionID]
												if ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:762
			_go_fuzz_dep_.CoverTab[100208]++
													if metadata.Err == ErrLeaderNotAvailable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:763
				_go_fuzz_dep_.CoverTab[100211]++
														return nil, ErrLeaderNotAvailable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:764
				// _ = "end of CoverTab[100211]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:765
				_go_fuzz_dep_.CoverTab[100212]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:765
				// _ = "end of CoverTab[100212]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:765
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:765
			// _ = "end of CoverTab[100208]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:765
			_go_fuzz_dep_.CoverTab[100209]++
													b := client.brokers[metadata.Leader]
													if b == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:767
				_go_fuzz_dep_.CoverTab[100213]++
														return nil, ErrLeaderNotAvailable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:768
				// _ = "end of CoverTab[100213]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:769
				_go_fuzz_dep_.CoverTab[100214]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:769
				// _ = "end of CoverTab[100214]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:769
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:769
			// _ = "end of CoverTab[100209]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:769
			_go_fuzz_dep_.CoverTab[100210]++
													_ = b.Open(client.conf)
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:771
			// _ = "end of CoverTab[100210]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:772
			_go_fuzz_dep_.CoverTab[100215]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:772
			// _ = "end of CoverTab[100215]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:772
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:772
		// _ = "end of CoverTab[100207]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:773
		_go_fuzz_dep_.CoverTab[100216]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:773
		// _ = "end of CoverTab[100216]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:773
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:773
	// _ = "end of CoverTab[100205]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:773
	_go_fuzz_dep_.CoverTab[100206]++

											return nil, ErrUnknownTopicOrPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:775
	// _ = "end of CoverTab[100206]"
}

func (client *client) getOffset(topic string, partitionID int32, time int64) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:778
	_go_fuzz_dep_.CoverTab[100217]++
											broker, err := client.Leader(topic, partitionID)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:780
		_go_fuzz_dep_.CoverTab[100224]++
												return -1, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:781
		// _ = "end of CoverTab[100224]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:782
		_go_fuzz_dep_.CoverTab[100225]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:782
		// _ = "end of CoverTab[100225]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:782
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:782
	// _ = "end of CoverTab[100217]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:782
	_go_fuzz_dep_.CoverTab[100218]++

											request := &OffsetRequest{}
											if client.conf.Version.IsAtLeast(V0_10_1_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:785
		_go_fuzz_dep_.CoverTab[100226]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:786
		// _ = "end of CoverTab[100226]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:787
		_go_fuzz_dep_.CoverTab[100227]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:787
		// _ = "end of CoverTab[100227]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:787
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:787
	// _ = "end of CoverTab[100218]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:787
	_go_fuzz_dep_.CoverTab[100219]++
											request.AddBlock(topic, partitionID, time, 1)

											response, err := broker.GetAvailableOffsets(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:791
		_go_fuzz_dep_.CoverTab[100228]++
												_ = broker.Close()
												return -1, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:793
		// _ = "end of CoverTab[100228]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:794
		_go_fuzz_dep_.CoverTab[100229]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:794
		// _ = "end of CoverTab[100229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:794
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:794
	// _ = "end of CoverTab[100219]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:794
	_go_fuzz_dep_.CoverTab[100220]++

											block := response.GetBlock(topic, partitionID)
											if block == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:797
		_go_fuzz_dep_.CoverTab[100230]++
												_ = broker.Close()
												return -1, ErrIncompleteResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:799
		// _ = "end of CoverTab[100230]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:800
		_go_fuzz_dep_.CoverTab[100231]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:800
		// _ = "end of CoverTab[100231]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:800
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:800
	// _ = "end of CoverTab[100220]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:800
	_go_fuzz_dep_.CoverTab[100221]++
											if block.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:801
		_go_fuzz_dep_.CoverTab[100232]++
												return -1, block.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:802
		// _ = "end of CoverTab[100232]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:803
		_go_fuzz_dep_.CoverTab[100233]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:803
		// _ = "end of CoverTab[100233]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:803
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:803
	// _ = "end of CoverTab[100221]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:803
	_go_fuzz_dep_.CoverTab[100222]++
											if len(block.Offsets) != 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:804
		_go_fuzz_dep_.CoverTab[100234]++
												return -1, ErrOffsetOutOfRange
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:805
		// _ = "end of CoverTab[100234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:806
		_go_fuzz_dep_.CoverTab[100235]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:806
		// _ = "end of CoverTab[100235]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:806
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:806
	// _ = "end of CoverTab[100222]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:806
	_go_fuzz_dep_.CoverTab[100223]++

											return block.Offsets[0], nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:808
	// _ = "end of CoverTab[100223]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:813
func (client *client) backgroundMetadataUpdater() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:813
	_go_fuzz_dep_.CoverTab[100236]++
											defer close(client.closed)

											if client.conf.Metadata.RefreshFrequency == time.Duration(0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:816
		_go_fuzz_dep_.CoverTab[100238]++
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:817
		// _ = "end of CoverTab[100238]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:818
		_go_fuzz_dep_.CoverTab[100239]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:818
		// _ = "end of CoverTab[100239]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:818
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:818
	// _ = "end of CoverTab[100236]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:818
	_go_fuzz_dep_.CoverTab[100237]++

											ticker := time.NewTicker(client.conf.Metadata.RefreshFrequency)
											defer ticker.Stop()

											for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:823
		_go_fuzz_dep_.CoverTab[100240]++
												select {
		case <-ticker.C:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:825
			_go_fuzz_dep_.CoverTab[100241]++
													if err := client.refreshMetadata(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:826
				_go_fuzz_dep_.CoverTab[100243]++
														Logger.Println("Client background metadata update:", err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:827
				// _ = "end of CoverTab[100243]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:828
				_go_fuzz_dep_.CoverTab[100244]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:828
				// _ = "end of CoverTab[100244]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:828
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:828
			// _ = "end of CoverTab[100241]"
		case <-client.closer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:829
			_go_fuzz_dep_.CoverTab[100242]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:830
			// _ = "end of CoverTab[100242]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:831
		// _ = "end of CoverTab[100240]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:832
	// _ = "end of CoverTab[100237]"
}

func (client *client) refreshMetadata() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:835
	_go_fuzz_dep_.CoverTab[100245]++
											var topics []string

											if !client.conf.Metadata.Full {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:838
		_go_fuzz_dep_.CoverTab[100248]++
												if specificTopics, err := client.MetadataTopics(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:839
			_go_fuzz_dep_.CoverTab[100249]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:840
			// _ = "end of CoverTab[100249]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:841
			_go_fuzz_dep_.CoverTab[100250]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:841
			if len(specificTopics) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:841
				_go_fuzz_dep_.CoverTab[100251]++
														return ErrNoTopicsToUpdateMetadata
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:842
				// _ = "end of CoverTab[100251]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:843
				_go_fuzz_dep_.CoverTab[100252]++
														topics = specificTopics
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:844
				// _ = "end of CoverTab[100252]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:845
			// _ = "end of CoverTab[100250]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:845
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:845
		// _ = "end of CoverTab[100248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:846
		_go_fuzz_dep_.CoverTab[100253]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:846
		// _ = "end of CoverTab[100253]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:846
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:846
	// _ = "end of CoverTab[100245]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:846
	_go_fuzz_dep_.CoverTab[100246]++

											if err := client.RefreshMetadata(topics...); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:848
		_go_fuzz_dep_.CoverTab[100254]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:849
		// _ = "end of CoverTab[100254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:850
		_go_fuzz_dep_.CoverTab[100255]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:850
		// _ = "end of CoverTab[100255]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:850
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:850
	// _ = "end of CoverTab[100246]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:850
	_go_fuzz_dep_.CoverTab[100247]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:852
	// _ = "end of CoverTab[100247]"
}

func (client *client) tryRefreshMetadata(topics []string, attemptsRemaining int, deadline time.Time) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:855
	_go_fuzz_dep_.CoverTab[100256]++
											pastDeadline := func(backoff time.Duration) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:856
		_go_fuzz_dep_.CoverTab[100261]++
												if !deadline.IsZero() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:857
			_go_fuzz_dep_.CoverTab[100263]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:857
			return time.Now().Add(backoff).After(deadline)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:857
			// _ = "end of CoverTab[100263]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:857
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:857
			_go_fuzz_dep_.CoverTab[100264]++

													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:859
			// _ = "end of CoverTab[100264]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:860
			_go_fuzz_dep_.CoverTab[100265]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:860
			// _ = "end of CoverTab[100265]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:860
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:860
		// _ = "end of CoverTab[100261]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:860
		_go_fuzz_dep_.CoverTab[100262]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:861
		// _ = "end of CoverTab[100262]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:862
	// _ = "end of CoverTab[100256]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:862
	_go_fuzz_dep_.CoverTab[100257]++
											retry := func(err error) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:863
		_go_fuzz_dep_.CoverTab[100266]++
												if attemptsRemaining > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:864
			_go_fuzz_dep_.CoverTab[100268]++
													backoff := client.computeBackoff(attemptsRemaining)
													if pastDeadline(backoff) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:866
				_go_fuzz_dep_.CoverTab[100271]++
														Logger.Println("client/metadata skipping last retries as we would go past the metadata timeout")
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:868
				// _ = "end of CoverTab[100271]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:869
				_go_fuzz_dep_.CoverTab[100272]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:869
				// _ = "end of CoverTab[100272]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:869
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:869
			// _ = "end of CoverTab[100268]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:869
			_go_fuzz_dep_.CoverTab[100269]++
													Logger.Printf("client/metadata retrying after %dms... (%d attempts remaining)\n", backoff/time.Millisecond, attemptsRemaining)
													if backoff > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:871
				_go_fuzz_dep_.CoverTab[100273]++
														time.Sleep(backoff)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:872
				// _ = "end of CoverTab[100273]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:873
				_go_fuzz_dep_.CoverTab[100274]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:873
				// _ = "end of CoverTab[100274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:873
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:873
			// _ = "end of CoverTab[100269]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:873
			_go_fuzz_dep_.CoverTab[100270]++
													return client.tryRefreshMetadata(topics, attemptsRemaining-1, deadline)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:874
			// _ = "end of CoverTab[100270]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:875
			_go_fuzz_dep_.CoverTab[100275]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:875
			// _ = "end of CoverTab[100275]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:875
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:875
		// _ = "end of CoverTab[100266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:875
		_go_fuzz_dep_.CoverTab[100267]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:876
		// _ = "end of CoverTab[100267]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:877
	// _ = "end of CoverTab[100257]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:877
	_go_fuzz_dep_.CoverTab[100258]++

											broker := client.any()
											for ; broker != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:880
		_go_fuzz_dep_.CoverTab[100276]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:880
		return !pastDeadline(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:880
		// _ = "end of CoverTab[100276]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:880
	}(); broker = client.any() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:880
		_go_fuzz_dep_.CoverTab[100277]++
												allowAutoTopicCreation := client.conf.Metadata.AllowAutoTopicCreation
												if len(topics) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:882
			_go_fuzz_dep_.CoverTab[100280]++
													DebugLogger.Printf("client/metadata fetching metadata for %v from broker %s\n", topics, broker.addr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:883
			// _ = "end of CoverTab[100280]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:884
			_go_fuzz_dep_.CoverTab[100281]++
													allowAutoTopicCreation = false
													DebugLogger.Printf("client/metadata fetching metadata for all topics from broker %s\n", broker.addr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:886
			// _ = "end of CoverTab[100281]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:887
		// _ = "end of CoverTab[100277]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:887
		_go_fuzz_dep_.CoverTab[100278]++

												req := &MetadataRequest{Topics: topics, AllowAutoTopicCreation: allowAutoTopicCreation}
												if client.conf.Version.IsAtLeast(V1_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:890
			_go_fuzz_dep_.CoverTab[100282]++
													req.Version = 5
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:891
			// _ = "end of CoverTab[100282]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:892
			_go_fuzz_dep_.CoverTab[100283]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:892
			if client.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:892
				_go_fuzz_dep_.CoverTab[100284]++
														req.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:893
				// _ = "end of CoverTab[100284]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:894
				_go_fuzz_dep_.CoverTab[100285]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:894
				// _ = "end of CoverTab[100285]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:894
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:894
			// _ = "end of CoverTab[100283]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:894
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:894
		// _ = "end of CoverTab[100278]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:894
		_go_fuzz_dep_.CoverTab[100279]++
												response, err := broker.GetMetadata(req)
												switch err := err.(type) {
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:897
			_go_fuzz_dep_.CoverTab[100286]++
													allKnownMetaData := len(topics) == 0

													shouldRetry, err := client.updateMetadata(response, allKnownMetaData)
													if shouldRetry {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:901
				_go_fuzz_dep_.CoverTab[100293]++
														Logger.Println("client/metadata found some partitions to be leaderless")
														return retry(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:903
				// _ = "end of CoverTab[100293]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:904
				_go_fuzz_dep_.CoverTab[100294]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:904
				// _ = "end of CoverTab[100294]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:904
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:904
			// _ = "end of CoverTab[100286]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:904
			_go_fuzz_dep_.CoverTab[100287]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:905
			// _ = "end of CoverTab[100287]"

		case PacketEncodingError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:907
			_go_fuzz_dep_.CoverTab[100288]++

													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:909
			// _ = "end of CoverTab[100288]"

		case KError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:911
			_go_fuzz_dep_.CoverTab[100289]++

													if err == ErrSASLAuthenticationFailed {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:913
				_go_fuzz_dep_.CoverTab[100295]++
														Logger.Println("client/metadata failed SASL authentication")
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:915
				// _ = "end of CoverTab[100295]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:916
				_go_fuzz_dep_.CoverTab[100296]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:916
				// _ = "end of CoverTab[100296]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:916
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:916
			// _ = "end of CoverTab[100289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:916
			_go_fuzz_dep_.CoverTab[100290]++

													if err == ErrTopicAuthorizationFailed {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:918
				_go_fuzz_dep_.CoverTab[100297]++
														Logger.Println("client is not authorized to access this topic. The topics were: ", topics)
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:920
				// _ = "end of CoverTab[100297]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:921
				_go_fuzz_dep_.CoverTab[100298]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:921
				// _ = "end of CoverTab[100298]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:921
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:921
			// _ = "end of CoverTab[100290]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:921
			_go_fuzz_dep_.CoverTab[100291]++

													Logger.Printf("client/metadata got error from broker %d while fetching metadata: %v\n", broker.ID(), err)
													_ = broker.Close()
													client.deregisterBroker(broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:925
			// _ = "end of CoverTab[100291]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:927
			_go_fuzz_dep_.CoverTab[100292]++

													Logger.Printf("client/metadata got error from broker %d while fetching metadata: %v\n", broker.ID(), err)
													_ = broker.Close()
													client.deregisterBroker(broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:931
			// _ = "end of CoverTab[100292]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:932
		// _ = "end of CoverTab[100279]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:933
	// _ = "end of CoverTab[100258]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:933
	_go_fuzz_dep_.CoverTab[100259]++

											if broker != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:935
		_go_fuzz_dep_.CoverTab[100299]++
												Logger.Printf("client/metadata not fetching metadata from broker %s as we would go past the metadata timeout\n", broker.addr)
												return retry(ErrOutOfBrokers)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:937
		// _ = "end of CoverTab[100299]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:938
		_go_fuzz_dep_.CoverTab[100300]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:938
		// _ = "end of CoverTab[100300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:938
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:938
	// _ = "end of CoverTab[100259]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:938
	_go_fuzz_dep_.CoverTab[100260]++

											Logger.Println("client/metadata no available broker to send metadata request to")
											client.resurrectDeadBrokers()
											return retry(ErrOutOfBrokers)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:942
	// _ = "end of CoverTab[100260]"
}

// if no fatal error, returns a list of topics that need retrying due to ErrLeaderNotAvailable
func (client *client) updateMetadata(data *MetadataResponse, allKnownMetaData bool) (retry bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:946
	_go_fuzz_dep_.CoverTab[100301]++
											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:947
		_go_fuzz_dep_.CoverTab[100305]++
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:948
		// _ = "end of CoverTab[100305]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:949
		_go_fuzz_dep_.CoverTab[100306]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:949
		// _ = "end of CoverTab[100306]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:949
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:949
	// _ = "end of CoverTab[100301]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:949
	_go_fuzz_dep_.CoverTab[100302]++

											client.lock.Lock()
											defer client.lock.Unlock()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:959
	client.updateBroker(data.Brokers)

	client.controllerID = data.ControllerID

	if allKnownMetaData {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:963
		_go_fuzz_dep_.CoverTab[100307]++
												client.metadata = make(map[string]map[int32]*PartitionMetadata)
												client.metadataTopics = make(map[string]none)
												client.cachedPartitionsResults = make(map[string][maxPartitionIndex][]int32)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:966
		// _ = "end of CoverTab[100307]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:967
		_go_fuzz_dep_.CoverTab[100308]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:967
		// _ = "end of CoverTab[100308]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:967
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:967
	// _ = "end of CoverTab[100302]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:967
	_go_fuzz_dep_.CoverTab[100303]++
											for _, topic := range data.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:968
		_go_fuzz_dep_.CoverTab[100309]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:972
		if _, exists := client.metadataTopics[topic.Name]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:972
			_go_fuzz_dep_.CoverTab[100313]++
													client.metadataTopics[topic.Name] = none{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:973
			// _ = "end of CoverTab[100313]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:974
			_go_fuzz_dep_.CoverTab[100314]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:974
			// _ = "end of CoverTab[100314]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:974
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:974
		// _ = "end of CoverTab[100309]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:974
		_go_fuzz_dep_.CoverTab[100310]++
												delete(client.metadata, topic.Name)
												delete(client.cachedPartitionsResults, topic.Name)

												switch topic.Err {
		case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:979
			_go_fuzz_dep_.CoverTab[100315]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:979
			// _ = "end of CoverTab[100315]"

		case ErrInvalidTopic, ErrTopicAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:981
			_go_fuzz_dep_.CoverTab[100316]++
													err = topic.Err
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:983
			// _ = "end of CoverTab[100316]"
		case ErrUnknownTopicOrPartition:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:984
			_go_fuzz_dep_.CoverTab[100317]++
													err = topic.Err
													retry = true
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:987
			// _ = "end of CoverTab[100317]"
		case ErrLeaderNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:988
			_go_fuzz_dep_.CoverTab[100318]++
													retry = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:989
			// _ = "end of CoverTab[100318]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:990
			_go_fuzz_dep_.CoverTab[100319]++
													Logger.Printf("Unexpected topic-level metadata error: %s", topic.Err)
													err = topic.Err
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:993
			// _ = "end of CoverTab[100319]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:994
		// _ = "end of CoverTab[100310]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:994
		_go_fuzz_dep_.CoverTab[100311]++

												client.metadata[topic.Name] = make(map[int32]*PartitionMetadata, len(topic.Partitions))
												for _, partition := range topic.Partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:997
			_go_fuzz_dep_.CoverTab[100320]++
													client.metadata[topic.Name][partition.ID] = partition
													if partition.Err == ErrLeaderNotAvailable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:999
				_go_fuzz_dep_.CoverTab[100321]++
														retry = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1000
				// _ = "end of CoverTab[100321]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1001
				_go_fuzz_dep_.CoverTab[100322]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1001
				// _ = "end of CoverTab[100322]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1001
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1001
			// _ = "end of CoverTab[100320]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1002
		// _ = "end of CoverTab[100311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1002
		_go_fuzz_dep_.CoverTab[100312]++

												var partitionCache [maxPartitionIndex][]int32
												partitionCache[allPartitions] = client.setPartitionCache(topic.Name, allPartitions)
												partitionCache[writablePartitions] = client.setPartitionCache(topic.Name, writablePartitions)
												client.cachedPartitionsResults[topic.Name] = partitionCache
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1007
		// _ = "end of CoverTab[100312]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1008
	// _ = "end of CoverTab[100303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1008
	_go_fuzz_dep_.CoverTab[100304]++

											return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1010
	// _ = "end of CoverTab[100304]"
}

func (client *client) cachedCoordinator(consumerGroup string) *Broker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1013
	_go_fuzz_dep_.CoverTab[100323]++
											client.lock.RLock()
											defer client.lock.RUnlock()
											if coordinatorID, ok := client.coordinators[consumerGroup]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1016
		_go_fuzz_dep_.CoverTab[100325]++
												return client.brokers[coordinatorID]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1017
		// _ = "end of CoverTab[100325]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1018
		_go_fuzz_dep_.CoverTab[100326]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1018
		// _ = "end of CoverTab[100326]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1018
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1018
	// _ = "end of CoverTab[100323]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1018
	_go_fuzz_dep_.CoverTab[100324]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1019
	// _ = "end of CoverTab[100324]"
}

func (client *client) cachedController() *Broker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1022
	_go_fuzz_dep_.CoverTab[100327]++
											client.lock.RLock()
											defer client.lock.RUnlock()

											return client.brokers[client.controllerID]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1026
	// _ = "end of CoverTab[100327]"
}

func (client *client) computeBackoff(attemptsRemaining int) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1029
	_go_fuzz_dep_.CoverTab[100328]++
											if client.conf.Metadata.Retry.BackoffFunc != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1030
		_go_fuzz_dep_.CoverTab[100330]++
												maxRetries := client.conf.Metadata.Retry.Max
												retries := maxRetries - attemptsRemaining
												return client.conf.Metadata.Retry.BackoffFunc(retries, maxRetries)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1033
		// _ = "end of CoverTab[100330]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1034
		_go_fuzz_dep_.CoverTab[100331]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1034
		// _ = "end of CoverTab[100331]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1034
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1034
	// _ = "end of CoverTab[100328]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1034
	_go_fuzz_dep_.CoverTab[100329]++
											return client.conf.Metadata.Retry.Backoff
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1035
	// _ = "end of CoverTab[100329]"
}

func (client *client) getConsumerMetadata(consumerGroup string, attemptsRemaining int) (*FindCoordinatorResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1038
	_go_fuzz_dep_.CoverTab[100332]++
											retry := func(err error) (*FindCoordinatorResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1039
		_go_fuzz_dep_.CoverTab[100335]++
												if attemptsRemaining > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1040
			_go_fuzz_dep_.CoverTab[100337]++
													backoff := client.computeBackoff(attemptsRemaining)
													Logger.Printf("client/coordinator retrying after %dms... (%d attempts remaining)\n", backoff/time.Millisecond, attemptsRemaining)
													time.Sleep(backoff)
													return client.getConsumerMetadata(consumerGroup, attemptsRemaining-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1044
			// _ = "end of CoverTab[100337]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1045
			_go_fuzz_dep_.CoverTab[100338]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1045
			// _ = "end of CoverTab[100338]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1045
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1045
		// _ = "end of CoverTab[100335]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1045
		_go_fuzz_dep_.CoverTab[100336]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1046
		// _ = "end of CoverTab[100336]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1047
	// _ = "end of CoverTab[100332]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1047
	_go_fuzz_dep_.CoverTab[100333]++

											for broker := client.any(); broker != nil; broker = client.any() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1049
		_go_fuzz_dep_.CoverTab[100339]++
												DebugLogger.Printf("client/coordinator requesting coordinator for consumergroup %s from %s\n", consumerGroup, broker.Addr())

												request := new(FindCoordinatorRequest)
												request.CoordinatorKey = consumerGroup
												request.CoordinatorType = CoordinatorGroup

												response, err := broker.FindCoordinator(request)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1057
			_go_fuzz_dep_.CoverTab[100341]++
													Logger.Printf("client/coordinator request to broker %s failed: %s\n", broker.Addr(), err)

													switch err.(type) {
			case PacketEncodingError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1061
				_go_fuzz_dep_.CoverTab[100342]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1062
				// _ = "end of CoverTab[100342]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1063
				_go_fuzz_dep_.CoverTab[100343]++
														_ = broker.Close()
														client.deregisterBroker(broker)
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1066
				// _ = "end of CoverTab[100343]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1067
			// _ = "end of CoverTab[100341]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1068
			_go_fuzz_dep_.CoverTab[100344]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1068
			// _ = "end of CoverTab[100344]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1068
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1068
		// _ = "end of CoverTab[100339]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1068
		_go_fuzz_dep_.CoverTab[100340]++

												switch response.Err {
		case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1071
			_go_fuzz_dep_.CoverTab[100345]++
													DebugLogger.Printf("client/coordinator coordinator for consumergroup %s is #%d (%s)\n", consumerGroup, response.Coordinator.ID(), response.Coordinator.Addr())
													return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1073
			// _ = "end of CoverTab[100345]"

		case ErrConsumerCoordinatorNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1075
			_go_fuzz_dep_.CoverTab[100346]++
													Logger.Printf("client/coordinator coordinator for consumer group %s is not available\n", consumerGroup)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1081
			if _, err := client.Leader("__consumer_offsets", 0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1081
				_go_fuzz_dep_.CoverTab[100350]++
														Logger.Printf("client/coordinator the __consumer_offsets topic is not initialized completely yet. Waiting 2 seconds...\n")
														time.Sleep(2 * time.Second)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1083
				// _ = "end of CoverTab[100350]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1084
				_go_fuzz_dep_.CoverTab[100351]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1084
				// _ = "end of CoverTab[100351]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1084
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1084
			// _ = "end of CoverTab[100346]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1084
			_go_fuzz_dep_.CoverTab[100347]++

													return retry(ErrConsumerCoordinatorNotAvailable)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1086
			// _ = "end of CoverTab[100347]"
		case ErrGroupAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1087
			_go_fuzz_dep_.CoverTab[100348]++
													Logger.Printf("client was not authorized to access group %s while attempting to find coordinator", consumerGroup)
													return retry(ErrGroupAuthorizationFailed)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1089
			// _ = "end of CoverTab[100348]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1091
			_go_fuzz_dep_.CoverTab[100349]++
													return nil, response.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1092
			// _ = "end of CoverTab[100349]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1093
		// _ = "end of CoverTab[100340]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1094
	// _ = "end of CoverTab[100333]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1094
	_go_fuzz_dep_.CoverTab[100334]++

											Logger.Println("client/coordinator no available broker to send consumer metadata request to")
											client.resurrectDeadBrokers()
											return retry(ErrOutOfBrokers)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1098
	// _ = "end of CoverTab[100334]"
}

// nopCloserClient embeds an existing Client, but disables
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1101
// the Close method (yet all other methods pass
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1101
// through unchanged). This is for use in larger structs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1101
// where it is undesirable to close the client that was
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1101
// passed in by the caller.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1106
type nopCloserClient struct {
	Client
}

// Close intercepts and purposely does not call the underlying
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1110
// client's Close() method.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1112
func (ncc *nopCloserClient) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1112
	_go_fuzz_dep_.CoverTab[100352]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1113
	// _ = "end of CoverTab[100352]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1114
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/client.go:1114
var _ = _go_fuzz_dep_.CoverTab
