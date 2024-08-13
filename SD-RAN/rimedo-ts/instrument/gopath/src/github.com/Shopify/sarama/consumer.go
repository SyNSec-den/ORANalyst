//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1
)

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rcrowley/go-metrics"
)

// ConsumerMessage encapsulates a Kafka message returned by the consumer.
type ConsumerMessage struct {
	Headers		[]*RecordHeader	// only set if kafka is version 0.11+
	Timestamp	time.Time	// only set if kafka is version 0.10+, inner message timestamp
	BlockTimestamp	time.Time	// only set if kafka is version 0.10+, outer (compressed) block timestamp

	Key, Value	[]byte
	Topic		string
	Partition	int32
	Offset		int64
}

// ConsumerError is what is provided to the user when an error occurs.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:26
// It wraps an error and includes the topic and partition.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:28
type ConsumerError struct {
	Topic		string
	Partition	int32
	Err		error
}

func (ce ConsumerError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:34
	_go_fuzz_dep_.CoverTab[100614]++
											return fmt.Sprintf("kafka: error while consuming %s/%d: %s", ce.Topic, ce.Partition, ce.Err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:35
	// _ = "end of CoverTab[100614]"
}

func (ce ConsumerError) Unwrap() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:38
	_go_fuzz_dep_.CoverTab[100615]++
											return ce.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:39
	// _ = "end of CoverTab[100615]"
}

// ConsumerErrors is a type that wraps a batch of errors and implements the Error interface.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:42
// It can be returned from the PartitionConsumer's Close methods to avoid the need to manually drain errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:42
// when stopping.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:45
type ConsumerErrors []*ConsumerError

func (ce ConsumerErrors) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:47
	_go_fuzz_dep_.CoverTab[100616]++
											return fmt.Sprintf("kafka: %d errors while consuming", len(ce))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:48
	// _ = "end of CoverTab[100616]"
}

// Consumer manages PartitionConsumers which process Kafka messages from brokers. You MUST call Close()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:51
// on a consumer to avoid leaks, it will not be garbage-collected automatically when it passes out of
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:51
// scope.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:54
type Consumer interface {
	// Topics returns the set of available topics as retrieved from the cluster
	// metadata. This method is the same as Client.Topics(), and is provided for
	// convenience.
	Topics() ([]string, error)

	// Partitions returns the sorted list of all partition IDs for the given topic.
	// This method is the same as Client.Partitions(), and is provided for convenience.
	Partitions(topic string) ([]int32, error)

	// ConsumePartition creates a PartitionConsumer on the given topic/partition with
	// the given offset. It will return an error if this Consumer is already consuming
	// on the given topic/partition. Offset can be a literal offset, or OffsetNewest
	// or OffsetOldest
	ConsumePartition(topic string, partition int32, offset int64) (PartitionConsumer, error)

	// HighWaterMarks returns the current high water marks for each topic and partition.
	// Consistency between partitions is not guaranteed since high water marks are updated separately.
	HighWaterMarks() map[string]map[int32]int64

	// Close shuts down the consumer. It must be called after all child
	// PartitionConsumers have already been closed.
	Close() error

	// Pause suspends fetching from the requested partitions. Future calls to the broker will not return any
	// records from these partitions until they have been resumed using Resume()/ResumeAll().
	// Note that this method does not affect partition subscription.
	// In particular, it does not cause a group rebalance when automatic assignment is used.
	Pause(topicPartitions map[string][]int32)

	// Resume resumes specified partitions which have been paused with Pause()/PauseAll().
	// New calls to the broker will return records from these partitions if there are any to be fetched.
	Resume(topicPartitions map[string][]int32)

	// Pause suspends fetching from all partitions. Future calls to the broker will not return any
	// records from these partitions until they have been resumed using Resume()/ResumeAll().
	// Note that this method does not affect partition subscription.
	// In particular, it does not cause a group rebalance when automatic assignment is used.
	PauseAll()

	// Resume resumes all partitions which have been paused with Pause()/PauseAll().
	// New calls to the broker will return records from these partitions if there are any to be fetched.
	ResumeAll()
}

type consumer struct {
	conf		*Config
	children	map[string]map[int32]*partitionConsumer
	brokerConsumers	map[*Broker]*brokerConsumer
	client		Client
	lock		sync.Mutex
}

// NewConsumer creates a new consumer using the given broker addresses and configuration.
func NewConsumer(addrs []string, config *Config) (Consumer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:108
	_go_fuzz_dep_.CoverTab[100617]++
											client, err := NewClient(addrs, config)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:110
		_go_fuzz_dep_.CoverTab[100619]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:111
		// _ = "end of CoverTab[100619]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:112
		_go_fuzz_dep_.CoverTab[100620]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:112
		// _ = "end of CoverTab[100620]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:112
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:112
	// _ = "end of CoverTab[100617]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:112
	_go_fuzz_dep_.CoverTab[100618]++
											return newConsumer(client)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:113
	// _ = "end of CoverTab[100618]"
}

// NewConsumerFromClient creates a new consumer using the given client. It is still
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:116
// necessary to call Close() on the underlying client when shutting down this consumer.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:118
func NewConsumerFromClient(client Client) (Consumer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:118
	_go_fuzz_dep_.CoverTab[100621]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:121
	cli := &nopCloserClient{client}
											return newConsumer(cli)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:122
	// _ = "end of CoverTab[100621]"
}

func newConsumer(client Client) (Consumer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:125
	_go_fuzz_dep_.CoverTab[100622]++

											if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:127
		_go_fuzz_dep_.CoverTab[100624]++
												return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:128
		// _ = "end of CoverTab[100624]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:129
		_go_fuzz_dep_.CoverTab[100625]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:129
		// _ = "end of CoverTab[100625]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:129
	// _ = "end of CoverTab[100622]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:129
	_go_fuzz_dep_.CoverTab[100623]++

											c := &consumer{
		client:			client,
		conf:			client.Config(),
		children:		make(map[string]map[int32]*partitionConsumer),
		brokerConsumers:	make(map[*Broker]*brokerConsumer),
	}

											return c, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:138
	// _ = "end of CoverTab[100623]"
}

func (c *consumer) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:141
	_go_fuzz_dep_.CoverTab[100626]++
											return c.client.Close()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:142
	// _ = "end of CoverTab[100626]"
}

func (c *consumer) Topics() ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:145
	_go_fuzz_dep_.CoverTab[100627]++
											return c.client.Topics()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:146
	// _ = "end of CoverTab[100627]"
}

func (c *consumer) Partitions(topic string) ([]int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:149
	_go_fuzz_dep_.CoverTab[100628]++
											return c.client.Partitions(topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:150
	// _ = "end of CoverTab[100628]"
}

func (c *consumer) ConsumePartition(topic string, partition int32, offset int64) (PartitionConsumer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:153
	_go_fuzz_dep_.CoverTab[100629]++
											child := &partitionConsumer{
		consumer:		c,
		conf:			c.conf,
		topic:			topic,
		partition:		partition,
		messages:		make(chan *ConsumerMessage, c.conf.ChannelBufferSize),
		errors:			make(chan *ConsumerError, c.conf.ChannelBufferSize),
		feeder:			make(chan *FetchResponse, 1),
		preferredReadReplica:	invalidPreferredReplicaID,
		trigger:		make(chan none, 1),
		dying:			make(chan none),
		fetchSize:		c.conf.Consumer.Fetch.Default,
	}

	if err := child.chooseStartingOffset(offset); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:168
		_go_fuzz_dep_.CoverTab[100633]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:169
		// _ = "end of CoverTab[100633]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:170
		_go_fuzz_dep_.CoverTab[100634]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:170
		// _ = "end of CoverTab[100634]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:170
	// _ = "end of CoverTab[100629]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:170
	_go_fuzz_dep_.CoverTab[100630]++

											var leader *Broker
											var err error
											if leader, err = c.client.Leader(child.topic, child.partition); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:174
		_go_fuzz_dep_.CoverTab[100635]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:175
		// _ = "end of CoverTab[100635]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:176
		_go_fuzz_dep_.CoverTab[100636]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:176
		// _ = "end of CoverTab[100636]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:176
	// _ = "end of CoverTab[100630]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:176
	_go_fuzz_dep_.CoverTab[100631]++

											if err := c.addChild(child); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:178
		_go_fuzz_dep_.CoverTab[100637]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:179
		// _ = "end of CoverTab[100637]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:180
		_go_fuzz_dep_.CoverTab[100638]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:180
		// _ = "end of CoverTab[100638]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:180
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:180
	// _ = "end of CoverTab[100631]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:180
	_go_fuzz_dep_.CoverTab[100632]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:180
	_curRoutineNum128_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:180
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum128_)

											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
		_go_fuzz_dep_.CoverTab[100639]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
			_go_fuzz_dep_.CoverTab[100640]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum128_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
			// _ = "end of CoverTab[100640]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
		withRecover(child.dispatcher)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
		// _ = "end of CoverTab[100639]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
	_curRoutineNum129_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:182
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum129_)
											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
		_go_fuzz_dep_.CoverTab[100641]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
			_go_fuzz_dep_.CoverTab[100642]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum129_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
			// _ = "end of CoverTab[100642]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
		withRecover(child.responseFeeder)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
		// _ = "end of CoverTab[100641]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:183
	}()

											child.broker = c.refBrokerConsumer(leader)
											child.broker.input <- child

											return child, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:188
	// _ = "end of CoverTab[100632]"
}

func (c *consumer) HighWaterMarks() map[string]map[int32]int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:191
	_go_fuzz_dep_.CoverTab[100643]++
											c.lock.Lock()
											defer c.lock.Unlock()

											hwms := make(map[string]map[int32]int64)
											for topic, p := range c.children {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:196
		_go_fuzz_dep_.CoverTab[100645]++
												hwm := make(map[int32]int64, len(p))
												for partition, pc := range p {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:198
			_go_fuzz_dep_.CoverTab[100647]++
													hwm[partition] = pc.HighWaterMarkOffset()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:199
			// _ = "end of CoverTab[100647]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:200
		// _ = "end of CoverTab[100645]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:200
		_go_fuzz_dep_.CoverTab[100646]++
												hwms[topic] = hwm
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:201
		// _ = "end of CoverTab[100646]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:202
	// _ = "end of CoverTab[100643]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:202
	_go_fuzz_dep_.CoverTab[100644]++

											return hwms
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:204
	// _ = "end of CoverTab[100644]"
}

func (c *consumer) addChild(child *partitionConsumer) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:207
	_go_fuzz_dep_.CoverTab[100648]++
											c.lock.Lock()
											defer c.lock.Unlock()

											topicChildren := c.children[child.topic]
											if topicChildren == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:212
		_go_fuzz_dep_.CoverTab[100651]++
												topicChildren = make(map[int32]*partitionConsumer)
												c.children[child.topic] = topicChildren
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:214
		// _ = "end of CoverTab[100651]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:215
		_go_fuzz_dep_.CoverTab[100652]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:215
		// _ = "end of CoverTab[100652]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:215
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:215
	// _ = "end of CoverTab[100648]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:215
	_go_fuzz_dep_.CoverTab[100649]++

											if topicChildren[child.partition] != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:217
		_go_fuzz_dep_.CoverTab[100653]++
												return ConfigurationError("That topic/partition is already being consumed")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:218
		// _ = "end of CoverTab[100653]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:219
		_go_fuzz_dep_.CoverTab[100654]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:219
		// _ = "end of CoverTab[100654]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:219
	// _ = "end of CoverTab[100649]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:219
	_go_fuzz_dep_.CoverTab[100650]++

											topicChildren[child.partition] = child
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:222
	// _ = "end of CoverTab[100650]"
}

func (c *consumer) removeChild(child *partitionConsumer) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:225
	_go_fuzz_dep_.CoverTab[100655]++
											c.lock.Lock()
											defer c.lock.Unlock()

											delete(c.children[child.topic], child.partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:229
	// _ = "end of CoverTab[100655]"
}

func (c *consumer) refBrokerConsumer(broker *Broker) *brokerConsumer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:232
	_go_fuzz_dep_.CoverTab[100656]++
											c.lock.Lock()
											defer c.lock.Unlock()

											bc := c.brokerConsumers[broker]
											if bc == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:237
		_go_fuzz_dep_.CoverTab[100658]++
												bc = c.newBrokerConsumer(broker)
												c.brokerConsumers[broker] = bc
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:239
		// _ = "end of CoverTab[100658]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:240
		_go_fuzz_dep_.CoverTab[100659]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:240
		// _ = "end of CoverTab[100659]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:240
	// _ = "end of CoverTab[100656]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:240
	_go_fuzz_dep_.CoverTab[100657]++

											bc.refs++

											return bc
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:244
	// _ = "end of CoverTab[100657]"
}

func (c *consumer) unrefBrokerConsumer(brokerWorker *brokerConsumer) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:247
	_go_fuzz_dep_.CoverTab[100660]++
											c.lock.Lock()
											defer c.lock.Unlock()

											brokerWorker.refs--

											if brokerWorker.refs == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:253
		_go_fuzz_dep_.CoverTab[100661]++
												close(brokerWorker.input)
												if c.brokerConsumers[brokerWorker.broker] == brokerWorker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:255
			_go_fuzz_dep_.CoverTab[100662]++
													delete(c.brokerConsumers, brokerWorker.broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:256
			// _ = "end of CoverTab[100662]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:257
			_go_fuzz_dep_.CoverTab[100663]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:257
			// _ = "end of CoverTab[100663]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:257
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:257
		// _ = "end of CoverTab[100661]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:258
		_go_fuzz_dep_.CoverTab[100664]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:258
		// _ = "end of CoverTab[100664]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:258
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:258
	// _ = "end of CoverTab[100660]"
}

func (c *consumer) abandonBrokerConsumer(brokerWorker *brokerConsumer) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:261
	_go_fuzz_dep_.CoverTab[100665]++
											c.lock.Lock()
											defer c.lock.Unlock()

											delete(c.brokerConsumers, brokerWorker.broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:265
	// _ = "end of CoverTab[100665]"
}

// Pause implements Consumer.
func (c *consumer) Pause(topicPartitions map[string][]int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:269
	_go_fuzz_dep_.CoverTab[100666]++
											c.lock.Lock()
											defer c.lock.Unlock()

											for topic, partitions := range topicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:273
		_go_fuzz_dep_.CoverTab[100667]++
												for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:274
			_go_fuzz_dep_.CoverTab[100668]++
													if topicConsumers, ok := c.children[topic]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:275
				_go_fuzz_dep_.CoverTab[100669]++
														if partitionConsumer, ok := topicConsumers[partition]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:276
					_go_fuzz_dep_.CoverTab[100670]++
															partitionConsumer.Pause()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:277
					// _ = "end of CoverTab[100670]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:278
					_go_fuzz_dep_.CoverTab[100671]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:278
					// _ = "end of CoverTab[100671]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:278
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:278
				// _ = "end of CoverTab[100669]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:279
				_go_fuzz_dep_.CoverTab[100672]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:279
				// _ = "end of CoverTab[100672]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:279
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:279
			// _ = "end of CoverTab[100668]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:280
		// _ = "end of CoverTab[100667]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:281
	// _ = "end of CoverTab[100666]"
}

// Resume implements Consumer.
func (c *consumer) Resume(topicPartitions map[string][]int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:285
	_go_fuzz_dep_.CoverTab[100673]++
											c.lock.Lock()
											defer c.lock.Unlock()

											for topic, partitions := range topicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:289
		_go_fuzz_dep_.CoverTab[100674]++
												for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:290
			_go_fuzz_dep_.CoverTab[100675]++
													if topicConsumers, ok := c.children[topic]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:291
				_go_fuzz_dep_.CoverTab[100676]++
														if partitionConsumer, ok := topicConsumers[partition]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:292
					_go_fuzz_dep_.CoverTab[100677]++
															partitionConsumer.Resume()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:293
					// _ = "end of CoverTab[100677]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:294
					_go_fuzz_dep_.CoverTab[100678]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:294
					// _ = "end of CoverTab[100678]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:294
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:294
				// _ = "end of CoverTab[100676]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:295
				_go_fuzz_dep_.CoverTab[100679]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:295
				// _ = "end of CoverTab[100679]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:295
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:295
			// _ = "end of CoverTab[100675]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:296
		// _ = "end of CoverTab[100674]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:297
	// _ = "end of CoverTab[100673]"
}

// PauseAll implements Consumer.
func (c *consumer) PauseAll() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:301
	_go_fuzz_dep_.CoverTab[100680]++
											c.lock.Lock()
											defer c.lock.Unlock()

											for _, partitions := range c.children {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:305
		_go_fuzz_dep_.CoverTab[100681]++
												for _, partitionConsumer := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:306
			_go_fuzz_dep_.CoverTab[100682]++
													partitionConsumer.Pause()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:307
			// _ = "end of CoverTab[100682]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:308
		// _ = "end of CoverTab[100681]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:309
	// _ = "end of CoverTab[100680]"
}

// ResumeAll implements Consumer.
func (c *consumer) ResumeAll() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:313
	_go_fuzz_dep_.CoverTab[100683]++
											c.lock.Lock()
											defer c.lock.Unlock()

											for _, partitions := range c.children {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:317
		_go_fuzz_dep_.CoverTab[100684]++
												for _, partitionConsumer := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:318
			_go_fuzz_dep_.CoverTab[100685]++
													partitionConsumer.Resume()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:319
			// _ = "end of CoverTab[100685]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:320
		// _ = "end of CoverTab[100684]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:321
	// _ = "end of CoverTab[100683]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// PartitionConsumer processes Kafka messages from a given topic and partition. You MUST call one of Close() or
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// AsyncClose() on a PartitionConsumer to avoid leaks; it will not be garbage-collected automatically when it passes out
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// of scope.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// The simplest way of using a PartitionConsumer is to loop over its Messages channel using a for/range
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// loop. The PartitionConsumer will only stop itself in one case: when the offset being consumed is reported
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// as out of range by the brokers. In this case you should decide what you want to do (try a different offset,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// notify a human, etc) and handle it appropriately. For all other error cases, it will just keep retrying.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// By default, it logs these errors to sarama.Logger; if you want to be notified directly of all errors, set
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// your config's Consumer.Return.Errors to true and read from the Errors channel, using a select statement
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// or a separate goroutine. Check out the Consumer examples to see implementations of these different approaches.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// To terminate such a for/range loop while the loop is executing, call AsyncClose. This will kick off the process of
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// consumer tear-down & return immediately. Continue to loop, servicing the Messages channel until the teardown process
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// AsyncClose initiated closes it (thus terminating the for/range loop). If you've already ceased reading Messages, call
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// Close; this will signal the PartitionConsumer's goroutines to begin shutting down (just like AsyncClose), but will
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:326
// also drain the Messages channel, harvest all errors & return them once cleanup has completed.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:343
type PartitionConsumer interface {
	// AsyncClose initiates a shutdown of the PartitionConsumer. This method will return immediately, after which you
	// should continue to service the 'Messages' and 'Errors' channels until they are empty. It is required to call this
	// function, or Close before a consumer object passes out of scope, as it will otherwise leak memory. You must call
	// this before calling Close on the underlying client.
	AsyncClose()

	// Close stops the PartitionConsumer from fetching messages. It will initiate a shutdown just like AsyncClose, drain
	// the Messages channel, harvest any errors & return them to the caller. Note that if you are continuing to service
	// the Messages channel when this function is called, you will be competing with Close for messages; consider
	// calling AsyncClose, instead. It is required to call this function (or AsyncClose) before a consumer object passes
	// out of scope, as it will otherwise leak memory. You must call this before calling Close on the underlying client.
	Close() error

	// Messages returns the read channel for the messages that are returned by
	// the broker.
	Messages() <-chan *ConsumerMessage

	// Errors returns a read channel of errors that occurred during consuming, if
	// enabled. By default, errors are logged and not returned over this channel.
	// If you want to implement any custom error handling, set your config's
	// Consumer.Return.Errors setting to true, and read from this channel.
	Errors() <-chan *ConsumerError

	// HighWaterMarkOffset returns the high water mark offset of the partition,
	// i.e. the offset that will be used for the next message that will be produced.
	// You can use this to determine how far behind the processing is.
	HighWaterMarkOffset() int64

	// Pause suspends fetching from this partition. Future calls to the broker will not return
	// any records from these partition until it have been resumed using Resume().
	// Note that this method does not affect partition subscription.
	// In particular, it does not cause a group rebalance when automatic assignment is used.
	Pause()

	// Resume resumes this partition which have been paused with Pause().
	// New calls to the broker will return records from these partitions if there are any to be fetched.
	// If the partition was not previously paused, this method is a no-op.
	Resume()

	// IsPaused indicates if this partition consumer is paused or not
	IsPaused() bool
}

type partitionConsumer struct {
	highWaterMarkOffset	int64	// must be at the top of the struct because https://golang.org/pkg/sync/atomic/#pkg-note-BUG

	consumer	*consumer
	conf		*Config
	broker		*brokerConsumer
	messages	chan *ConsumerMessage
	errors		chan *ConsumerError
	feeder		chan *FetchResponse

	preferredReadReplica	int32

	trigger, dying	chan none
	closeOnce	sync.Once
	topic		string
	partition	int32
	responseResult	error
	fetchSize	int32
	offset		int64
	retries		int32

	paused	int32
}

var errTimedOut = errors.New("timed out feeding messages to the user")	// not user-facing

func (child *partitionConsumer) sendError(err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:413
	_go_fuzz_dep_.CoverTab[100686]++
											cErr := &ConsumerError{
		Topic:		child.topic,
		Partition:	child.partition,
		Err:		err,
	}

	if child.conf.Consumer.Return.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:420
		_go_fuzz_dep_.CoverTab[100687]++
												child.errors <- cErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:421
		// _ = "end of CoverTab[100687]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:422
		_go_fuzz_dep_.CoverTab[100688]++
												Logger.Println(cErr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:423
		// _ = "end of CoverTab[100688]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:424
	// _ = "end of CoverTab[100686]"
}

func (child *partitionConsumer) computeBackoff() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:427
	_go_fuzz_dep_.CoverTab[100689]++
											if child.conf.Consumer.Retry.BackoffFunc != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:428
		_go_fuzz_dep_.CoverTab[100691]++
												retries := atomic.AddInt32(&child.retries, 1)
												return child.conf.Consumer.Retry.BackoffFunc(int(retries))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:430
		// _ = "end of CoverTab[100691]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:431
		_go_fuzz_dep_.CoverTab[100692]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:431
		// _ = "end of CoverTab[100692]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:431
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:431
	// _ = "end of CoverTab[100689]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:431
	_go_fuzz_dep_.CoverTab[100690]++
											return child.conf.Consumer.Retry.Backoff
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:432
	// _ = "end of CoverTab[100690]"
}

func (child *partitionConsumer) dispatcher() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:435
	_go_fuzz_dep_.CoverTab[100693]++
											for range child.trigger {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:436
		_go_fuzz_dep_.CoverTab[100696]++
												select {
		case <-child.dying:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:438
			_go_fuzz_dep_.CoverTab[100697]++
													close(child.trigger)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:439
			// _ = "end of CoverTab[100697]"
		case <-time.After(child.computeBackoff()):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:440
			_go_fuzz_dep_.CoverTab[100698]++
													if child.broker != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:441
				_go_fuzz_dep_.CoverTab[100700]++
														child.consumer.unrefBrokerConsumer(child.broker)
														child.broker = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:443
				// _ = "end of CoverTab[100700]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:444
				_go_fuzz_dep_.CoverTab[100701]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:444
				// _ = "end of CoverTab[100701]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:444
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:444
			// _ = "end of CoverTab[100698]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:444
			_go_fuzz_dep_.CoverTab[100699]++

													if err := child.dispatch(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:446
				_go_fuzz_dep_.CoverTab[100702]++
														child.sendError(err)
														child.trigger <- none{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:448
				// _ = "end of CoverTab[100702]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:449
				_go_fuzz_dep_.CoverTab[100703]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:449
				// _ = "end of CoverTab[100703]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:449
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:449
			// _ = "end of CoverTab[100699]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:450
		// _ = "end of CoverTab[100696]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:451
	// _ = "end of CoverTab[100693]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:451
	_go_fuzz_dep_.CoverTab[100694]++

											if child.broker != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:453
		_go_fuzz_dep_.CoverTab[100704]++
												child.consumer.unrefBrokerConsumer(child.broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:454
		// _ = "end of CoverTab[100704]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:455
		_go_fuzz_dep_.CoverTab[100705]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:455
		// _ = "end of CoverTab[100705]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:455
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:455
	// _ = "end of CoverTab[100694]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:455
	_go_fuzz_dep_.CoverTab[100695]++
											child.consumer.removeChild(child)
											close(child.feeder)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:457
	// _ = "end of CoverTab[100695]"
}

func (child *partitionConsumer) preferredBroker() (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:460
	_go_fuzz_dep_.CoverTab[100706]++
											if child.preferredReadReplica >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:461
		_go_fuzz_dep_.CoverTab[100708]++
												broker, err := child.consumer.client.Broker(child.preferredReadReplica)
												if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:463
			_go_fuzz_dep_.CoverTab[100710]++
													return broker, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:464
			// _ = "end of CoverTab[100710]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:465
			_go_fuzz_dep_.CoverTab[100711]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:465
			// _ = "end of CoverTab[100711]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:465
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:465
		// _ = "end of CoverTab[100708]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:465
		_go_fuzz_dep_.CoverTab[100709]++
												Logger.Printf(
			"consumer/%s/%d failed to find active broker for preferred read replica %d - will fallback to leader",
			child.topic, child.partition, child.preferredReadReplica)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:472
		child.preferredReadReplica = invalidPreferredReplicaID
												_ = child.consumer.client.RefreshMetadata(child.topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:473
		// _ = "end of CoverTab[100709]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:474
		_go_fuzz_dep_.CoverTab[100712]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:474
		// _ = "end of CoverTab[100712]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:474
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:474
	// _ = "end of CoverTab[100706]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:474
	_go_fuzz_dep_.CoverTab[100707]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:477
	return child.consumer.client.Leader(child.topic, child.partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:477
	// _ = "end of CoverTab[100707]"
}

func (child *partitionConsumer) dispatch() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:480
	_go_fuzz_dep_.CoverTab[100713]++
											if err := child.consumer.client.RefreshMetadata(child.topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:481
		_go_fuzz_dep_.CoverTab[100716]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:482
		// _ = "end of CoverTab[100716]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:483
		_go_fuzz_dep_.CoverTab[100717]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:483
		// _ = "end of CoverTab[100717]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:483
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:483
	// _ = "end of CoverTab[100713]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:483
	_go_fuzz_dep_.CoverTab[100714]++

											broker, err := child.preferredBroker()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:486
		_go_fuzz_dep_.CoverTab[100718]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:487
		// _ = "end of CoverTab[100718]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:488
		_go_fuzz_dep_.CoverTab[100719]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:488
		// _ = "end of CoverTab[100719]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:488
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:488
	// _ = "end of CoverTab[100714]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:488
	_go_fuzz_dep_.CoverTab[100715]++

											child.broker = child.consumer.refBrokerConsumer(broker)

											child.broker.input <- child

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:494
	// _ = "end of CoverTab[100715]"
}

func (child *partitionConsumer) chooseStartingOffset(offset int64) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:497
	_go_fuzz_dep_.CoverTab[100720]++
											newestOffset, err := child.consumer.client.GetOffset(child.topic, child.partition, OffsetNewest)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:499
		_go_fuzz_dep_.CoverTab[100724]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:500
		// _ = "end of CoverTab[100724]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:501
		_go_fuzz_dep_.CoverTab[100725]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:501
		// _ = "end of CoverTab[100725]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:501
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:501
	// _ = "end of CoverTab[100720]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:501
	_go_fuzz_dep_.CoverTab[100721]++

											child.highWaterMarkOffset = newestOffset

											oldestOffset, err := child.consumer.client.GetOffset(child.topic, child.partition, OffsetOldest)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:506
		_go_fuzz_dep_.CoverTab[100726]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:507
		// _ = "end of CoverTab[100726]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:508
		_go_fuzz_dep_.CoverTab[100727]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:508
		// _ = "end of CoverTab[100727]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:508
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:508
	// _ = "end of CoverTab[100721]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:508
	_go_fuzz_dep_.CoverTab[100722]++

											switch {
	case offset == OffsetNewest:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:511
		_go_fuzz_dep_.CoverTab[100728]++
												child.offset = newestOffset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:512
		// _ = "end of CoverTab[100728]"
	case offset == OffsetOldest:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:513
		_go_fuzz_dep_.CoverTab[100729]++
												child.offset = oldestOffset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:514
		// _ = "end of CoverTab[100729]"
	case offset >= oldestOffset && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:515
		_go_fuzz_dep_.CoverTab[100732]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:515
		return offset <= newestOffset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:515
		// _ = "end of CoverTab[100732]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:515
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:515
		_go_fuzz_dep_.CoverTab[100730]++
												child.offset = offset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:516
		// _ = "end of CoverTab[100730]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:517
		_go_fuzz_dep_.CoverTab[100731]++
												return ErrOffsetOutOfRange
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:518
		// _ = "end of CoverTab[100731]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:519
	// _ = "end of CoverTab[100722]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:519
	_go_fuzz_dep_.CoverTab[100723]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:521
	// _ = "end of CoverTab[100723]"
}

func (child *partitionConsumer) Messages() <-chan *ConsumerMessage {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:524
	_go_fuzz_dep_.CoverTab[100733]++
											return child.messages
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:525
	// _ = "end of CoverTab[100733]"
}

func (child *partitionConsumer) Errors() <-chan *ConsumerError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:528
	_go_fuzz_dep_.CoverTab[100734]++
											return child.errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:529
	// _ = "end of CoverTab[100734]"
}

func (child *partitionConsumer) AsyncClose() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:532
	_go_fuzz_dep_.CoverTab[100735]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:537
	child.closeOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:537
		_go_fuzz_dep_.CoverTab[100736]++
												close(child.dying)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:538
		// _ = "end of CoverTab[100736]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:539
	// _ = "end of CoverTab[100735]"
}

func (child *partitionConsumer) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:542
	_go_fuzz_dep_.CoverTab[100737]++
											child.AsyncClose()

											var consumerErrors ConsumerErrors
											for err := range child.errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:546
		_go_fuzz_dep_.CoverTab[100740]++
												consumerErrors = append(consumerErrors, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:547
		// _ = "end of CoverTab[100740]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:548
	// _ = "end of CoverTab[100737]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:548
	_go_fuzz_dep_.CoverTab[100738]++

											if len(consumerErrors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:550
		_go_fuzz_dep_.CoverTab[100741]++
												return consumerErrors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:551
		// _ = "end of CoverTab[100741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:552
		_go_fuzz_dep_.CoverTab[100742]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:552
		// _ = "end of CoverTab[100742]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:552
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:552
	// _ = "end of CoverTab[100738]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:552
	_go_fuzz_dep_.CoverTab[100739]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:553
	// _ = "end of CoverTab[100739]"
}

func (child *partitionConsumer) HighWaterMarkOffset() int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:556
	_go_fuzz_dep_.CoverTab[100743]++
											return atomic.LoadInt64(&child.highWaterMarkOffset)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:557
	// _ = "end of CoverTab[100743]"
}

func (child *partitionConsumer) responseFeeder() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:560
	_go_fuzz_dep_.CoverTab[100744]++
											var msgs []*ConsumerMessage
											expiryTicker := time.NewTicker(child.conf.Consumer.MaxProcessingTime)
											firstAttempt := true

feederLoop:
	for response := range child.feeder {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:566
		_go_fuzz_dep_.CoverTab[100746]++
												msgs, child.responseResult = child.parseResponse(response)

												if child.responseResult == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:569
			_go_fuzz_dep_.CoverTab[100749]++
													atomic.StoreInt32(&child.retries, 0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:570
			// _ = "end of CoverTab[100749]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:571
			_go_fuzz_dep_.CoverTab[100750]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:571
			// _ = "end of CoverTab[100750]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:571
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:571
		// _ = "end of CoverTab[100746]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:571
		_go_fuzz_dep_.CoverTab[100747]++

												for i, msg := range msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:573
			_go_fuzz_dep_.CoverTab[100751]++
													child.interceptors(msg)
		messageSelect:
			select {
			case <-child.dying:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:577
				_go_fuzz_dep_.CoverTab[100752]++
														child.broker.acks.Done()
														continue feederLoop
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:579
				// _ = "end of CoverTab[100752]"
			case child.messages <- msg:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:580
				_go_fuzz_dep_.CoverTab[100753]++
														firstAttempt = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:581
				// _ = "end of CoverTab[100753]"
			case <-expiryTicker.C:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:582
				_go_fuzz_dep_.CoverTab[100754]++
														if !firstAttempt {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:583
					_go_fuzz_dep_.CoverTab[100755]++
															child.responseResult = errTimedOut
															child.broker.acks.Done()
				remainingLoop:
					for _, msg = range msgs[i:] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:587
						_go_fuzz_dep_.CoverTab[100757]++
																child.interceptors(msg)
																select {
						case child.messages <- msg:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:590
							_go_fuzz_dep_.CoverTab[100758]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:590
							// _ = "end of CoverTab[100758]"
						case <-child.dying:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:591
							_go_fuzz_dep_.CoverTab[100759]++
																	break remainingLoop
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:592
							// _ = "end of CoverTab[100759]"
						}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:593
						// _ = "end of CoverTab[100757]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:594
					// _ = "end of CoverTab[100755]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:594
					_go_fuzz_dep_.CoverTab[100756]++
															child.broker.input <- child
															continue feederLoop
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:596
					// _ = "end of CoverTab[100756]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:597
					_go_fuzz_dep_.CoverTab[100760]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:600
					firstAttempt = false
															goto messageSelect
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:601
					// _ = "end of CoverTab[100760]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:602
				// _ = "end of CoverTab[100754]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:603
			// _ = "end of CoverTab[100751]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:604
		// _ = "end of CoverTab[100747]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:604
		_go_fuzz_dep_.CoverTab[100748]++

												child.broker.acks.Done()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:606
		// _ = "end of CoverTab[100748]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:607
	// _ = "end of CoverTab[100744]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:607
	_go_fuzz_dep_.CoverTab[100745]++

											expiryTicker.Stop()
											close(child.messages)
											close(child.errors)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:611
	// _ = "end of CoverTab[100745]"
}

func (child *partitionConsumer) parseMessages(msgSet *MessageSet) ([]*ConsumerMessage, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:614
	_go_fuzz_dep_.CoverTab[100761]++
											var messages []*ConsumerMessage
											for _, msgBlock := range msgSet.Messages {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:616
		_go_fuzz_dep_.CoverTab[100764]++
												for _, msg := range msgBlock.Messages() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:617
			_go_fuzz_dep_.CoverTab[100765]++
													offset := msg.Offset
													timestamp := msg.Msg.Timestamp
													if msg.Msg.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:620
				_go_fuzz_dep_.CoverTab[100768]++
														baseOffset := msgBlock.Offset - msgBlock.Messages()[len(msgBlock.Messages())-1].Offset
														offset += baseOffset
														if msg.Msg.LogAppendTime {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:623
					_go_fuzz_dep_.CoverTab[100769]++
															timestamp = msgBlock.Msg.Timestamp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:624
					// _ = "end of CoverTab[100769]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:625
					_go_fuzz_dep_.CoverTab[100770]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:625
					// _ = "end of CoverTab[100770]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:625
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:625
				// _ = "end of CoverTab[100768]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:626
				_go_fuzz_dep_.CoverTab[100771]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:626
				// _ = "end of CoverTab[100771]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:626
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:626
			// _ = "end of CoverTab[100765]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:626
			_go_fuzz_dep_.CoverTab[100766]++
													if offset < child.offset {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:627
				_go_fuzz_dep_.CoverTab[100772]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:628
				// _ = "end of CoverTab[100772]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:629
				_go_fuzz_dep_.CoverTab[100773]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:629
				// _ = "end of CoverTab[100773]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:629
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:629
			// _ = "end of CoverTab[100766]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:629
			_go_fuzz_dep_.CoverTab[100767]++
													messages = append(messages, &ConsumerMessage{
				Topic:		child.topic,
				Partition:	child.partition,
				Key:		msg.Msg.Key,
				Value:		msg.Msg.Value,
				Offset:		offset,
				Timestamp:	timestamp,
				BlockTimestamp:	msgBlock.Msg.Timestamp,
			})
													child.offset = offset + 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:639
			// _ = "end of CoverTab[100767]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:640
		// _ = "end of CoverTab[100764]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:641
	// _ = "end of CoverTab[100761]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:641
	_go_fuzz_dep_.CoverTab[100762]++
											if len(messages) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:642
		_go_fuzz_dep_.CoverTab[100774]++
												child.offset++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:643
		// _ = "end of CoverTab[100774]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:644
		_go_fuzz_dep_.CoverTab[100775]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:644
		// _ = "end of CoverTab[100775]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:644
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:644
	// _ = "end of CoverTab[100762]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:644
	_go_fuzz_dep_.CoverTab[100763]++
											return messages, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:645
	// _ = "end of CoverTab[100763]"
}

func (child *partitionConsumer) parseRecords(batch *RecordBatch) ([]*ConsumerMessage, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:648
	_go_fuzz_dep_.CoverTab[100776]++
											messages := make([]*ConsumerMessage, 0, len(batch.Records))

											for _, rec := range batch.Records {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:651
		_go_fuzz_dep_.CoverTab[100779]++
												offset := batch.FirstOffset + rec.OffsetDelta
												if offset < child.offset {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:653
			_go_fuzz_dep_.CoverTab[100782]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:654
			// _ = "end of CoverTab[100782]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:655
			_go_fuzz_dep_.CoverTab[100783]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:655
			// _ = "end of CoverTab[100783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:655
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:655
		// _ = "end of CoverTab[100779]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:655
		_go_fuzz_dep_.CoverTab[100780]++
												timestamp := batch.FirstTimestamp.Add(rec.TimestampDelta)
												if batch.LogAppendTime {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:657
			_go_fuzz_dep_.CoverTab[100784]++
													timestamp = batch.MaxTimestamp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:658
			// _ = "end of CoverTab[100784]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:659
			_go_fuzz_dep_.CoverTab[100785]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:659
			// _ = "end of CoverTab[100785]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:659
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:659
		// _ = "end of CoverTab[100780]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:659
		_go_fuzz_dep_.CoverTab[100781]++
												messages = append(messages, &ConsumerMessage{
			Topic:		child.topic,
			Partition:	child.partition,
			Key:		rec.Key,
			Value:		rec.Value,
			Offset:		offset,
			Timestamp:	timestamp,
			Headers:	rec.Headers,
		})
												child.offset = offset + 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:669
		// _ = "end of CoverTab[100781]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:670
	// _ = "end of CoverTab[100776]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:670
	_go_fuzz_dep_.CoverTab[100777]++
											if len(messages) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:671
		_go_fuzz_dep_.CoverTab[100786]++
												child.offset++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:672
		// _ = "end of CoverTab[100786]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:673
		_go_fuzz_dep_.CoverTab[100787]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:673
		// _ = "end of CoverTab[100787]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:673
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:673
	// _ = "end of CoverTab[100777]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:673
	_go_fuzz_dep_.CoverTab[100778]++
											return messages, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:674
	// _ = "end of CoverTab[100778]"
}

func (child *partitionConsumer) parseResponse(response *FetchResponse) ([]*ConsumerMessage, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:677
	_go_fuzz_dep_.CoverTab[100788]++
											var (
		metricRegistry		= child.conf.MetricRegistry
		consumerBatchSizeMetric	metrics.Histogram
	)

	if metricRegistry != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:683
		_go_fuzz_dep_.CoverTab[100797]++
												consumerBatchSizeMetric = getOrRegisterHistogram("consumer-batch-size", metricRegistry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:684
		// _ = "end of CoverTab[100797]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:685
		_go_fuzz_dep_.CoverTab[100798]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:685
		// _ = "end of CoverTab[100798]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:685
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:685
	// _ = "end of CoverTab[100788]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:685
	_go_fuzz_dep_.CoverTab[100789]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:688
	if response.ThrottleTime != time.Duration(0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:688
		_go_fuzz_dep_.CoverTab[100799]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:688
		return len(response.Blocks) == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:688
		// _ = "end of CoverTab[100799]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:688
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:688
		_go_fuzz_dep_.CoverTab[100800]++
												Logger.Printf(
			"consumer/broker/%d FetchResponse throttled %v\n",
			child.broker.broker.ID(), response.ThrottleTime)
												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:692
		// _ = "end of CoverTab[100800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:693
		_go_fuzz_dep_.CoverTab[100801]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:693
		// _ = "end of CoverTab[100801]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:693
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:693
	// _ = "end of CoverTab[100789]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:693
	_go_fuzz_dep_.CoverTab[100790]++

											block := response.GetBlock(child.topic, child.partition)
											if block == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:696
		_go_fuzz_dep_.CoverTab[100802]++
												return nil, ErrIncompleteResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:697
		// _ = "end of CoverTab[100802]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:698
		_go_fuzz_dep_.CoverTab[100803]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:698
		// _ = "end of CoverTab[100803]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:698
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:698
	// _ = "end of CoverTab[100790]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:698
	_go_fuzz_dep_.CoverTab[100791]++

											if block.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:700
		_go_fuzz_dep_.CoverTab[100804]++
												return nil, block.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:701
		// _ = "end of CoverTab[100804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:702
		_go_fuzz_dep_.CoverTab[100805]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:702
		// _ = "end of CoverTab[100805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:702
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:702
	// _ = "end of CoverTab[100791]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:702
	_go_fuzz_dep_.CoverTab[100792]++

											nRecs, err := block.numRecords()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:705
		_go_fuzz_dep_.CoverTab[100806]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:706
		// _ = "end of CoverTab[100806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:707
		_go_fuzz_dep_.CoverTab[100807]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:707
		// _ = "end of CoverTab[100807]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:707
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:707
	// _ = "end of CoverTab[100792]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:707
	_go_fuzz_dep_.CoverTab[100793]++

											consumerBatchSizeMetric.Update(int64(nRecs))

											if block.PreferredReadReplica != invalidPreferredReplicaID {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:711
		_go_fuzz_dep_.CoverTab[100808]++
												child.preferredReadReplica = block.PreferredReadReplica
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:712
		// _ = "end of CoverTab[100808]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:713
		_go_fuzz_dep_.CoverTab[100809]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:713
		// _ = "end of CoverTab[100809]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:713
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:713
	// _ = "end of CoverTab[100793]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:713
	_go_fuzz_dep_.CoverTab[100794]++

											if nRecs == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:715
		_go_fuzz_dep_.CoverTab[100810]++
												partialTrailingMessage, err := block.isPartial()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:717
			_go_fuzz_dep_.CoverTab[100813]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:718
			// _ = "end of CoverTab[100813]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:719
			_go_fuzz_dep_.CoverTab[100814]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:719
			// _ = "end of CoverTab[100814]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:719
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:719
		// _ = "end of CoverTab[100810]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:719
		_go_fuzz_dep_.CoverTab[100811]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:722
		if partialTrailingMessage {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:722
			_go_fuzz_dep_.CoverTab[100815]++
													if child.conf.Consumer.Fetch.Max > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:723
				_go_fuzz_dep_.CoverTab[100816]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:723
				return child.fetchSize == child.conf.Consumer.Fetch.Max
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:723
				// _ = "end of CoverTab[100816]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:723
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:723
				_go_fuzz_dep_.CoverTab[100817]++

														child.sendError(ErrMessageTooLarge)
														child.offset++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:726
				// _ = "end of CoverTab[100817]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:727
				_go_fuzz_dep_.CoverTab[100818]++
														child.fetchSize *= 2

														if child.fetchSize < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:730
					_go_fuzz_dep_.CoverTab[100820]++
															child.fetchSize = math.MaxInt32
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:731
					// _ = "end of CoverTab[100820]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:732
					_go_fuzz_dep_.CoverTab[100821]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:732
					// _ = "end of CoverTab[100821]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:732
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:732
				// _ = "end of CoverTab[100818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:732
				_go_fuzz_dep_.CoverTab[100819]++
														if child.conf.Consumer.Fetch.Max > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:733
					_go_fuzz_dep_.CoverTab[100822]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:733
					return child.fetchSize > child.conf.Consumer.Fetch.Max
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:733
					// _ = "end of CoverTab[100822]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:733
				}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:733
					_go_fuzz_dep_.CoverTab[100823]++
															child.fetchSize = child.conf.Consumer.Fetch.Max
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:734
					// _ = "end of CoverTab[100823]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:735
					_go_fuzz_dep_.CoverTab[100824]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:735
					// _ = "end of CoverTab[100824]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:735
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:735
				// _ = "end of CoverTab[100819]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:736
			// _ = "end of CoverTab[100815]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:737
			_go_fuzz_dep_.CoverTab[100825]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:737
			if block.LastRecordsBatchOffset != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:737
				_go_fuzz_dep_.CoverTab[100826]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:737
				return *block.LastRecordsBatchOffset < block.HighWaterMarkOffset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:737
				// _ = "end of CoverTab[100826]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:737
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:737
				_go_fuzz_dep_.CoverTab[100827]++

														Logger.Printf("consumer/broker/%d received batch with zero records but high watermark was not reached, topic %s, partition %d, offset %d\n", child.broker.broker.ID(), child.topic, child.partition, *block.LastRecordsBatchOffset)
														child.offset = *block.LastRecordsBatchOffset + 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:740
				// _ = "end of CoverTab[100827]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:741
				_go_fuzz_dep_.CoverTab[100828]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:741
				// _ = "end of CoverTab[100828]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:741
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:741
			// _ = "end of CoverTab[100825]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:741
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:741
		// _ = "end of CoverTab[100811]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:741
		_go_fuzz_dep_.CoverTab[100812]++

												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:743
		// _ = "end of CoverTab[100812]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:744
		_go_fuzz_dep_.CoverTab[100829]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:744
		// _ = "end of CoverTab[100829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:744
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:744
	// _ = "end of CoverTab[100794]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:744
	_go_fuzz_dep_.CoverTab[100795]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:747
	child.fetchSize = child.conf.Consumer.Fetch.Default
											atomic.StoreInt64(&child.highWaterMarkOffset, block.HighWaterMarkOffset)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:753
	abortedProducerIDs := make(map[int64]struct{}, len(block.AbortedTransactions))
	abortedTransactions := block.getAbortedTransactions()

	var messages []*ConsumerMessage
	for _, records := range block.RecordsSet {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:757
		_go_fuzz_dep_.CoverTab[100830]++
												switch records.recordsType {
		case legacyRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:759
			_go_fuzz_dep_.CoverTab[100831]++
													messageSetMessages, err := child.parseMessages(records.MsgSet)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:761
				_go_fuzz_dep_.CoverTab[100840]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:762
				// _ = "end of CoverTab[100840]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:763
				_go_fuzz_dep_.CoverTab[100841]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:763
				// _ = "end of CoverTab[100841]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:763
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:763
			// _ = "end of CoverTab[100831]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:763
			_go_fuzz_dep_.CoverTab[100832]++

													messages = append(messages, messageSetMessages...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:765
			// _ = "end of CoverTab[100832]"
		case defaultRecords:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:766
			_go_fuzz_dep_.CoverTab[100833]++

													for _, txn := range abortedTransactions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:768
				_go_fuzz_dep_.CoverTab[100842]++
														if txn.FirstOffset > records.RecordBatch.LastOffset() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:769
					_go_fuzz_dep_.CoverTab[100844]++
															break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:770
					// _ = "end of CoverTab[100844]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:771
					_go_fuzz_dep_.CoverTab[100845]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:771
					// _ = "end of CoverTab[100845]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:771
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:771
				// _ = "end of CoverTab[100842]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:771
				_go_fuzz_dep_.CoverTab[100843]++
														abortedProducerIDs[txn.ProducerID] = struct{}{}

														abortedTransactions = abortedTransactions[1:]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:774
				// _ = "end of CoverTab[100843]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:775
			// _ = "end of CoverTab[100833]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:775
			_go_fuzz_dep_.CoverTab[100834]++

													recordBatchMessages, err := child.parseRecords(records.RecordBatch)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:778
				_go_fuzz_dep_.CoverTab[100846]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:779
				// _ = "end of CoverTab[100846]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:780
				_go_fuzz_dep_.CoverTab[100847]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:780
				// _ = "end of CoverTab[100847]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:780
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:780
			// _ = "end of CoverTab[100834]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:780
			_go_fuzz_dep_.CoverTab[100835]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:787
			isControl, err := records.isControl()
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:788
				_go_fuzz_dep_.CoverTab[100848]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:792
				if child.conf.Consumer.IsolationLevel == ReadCommitted {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:792
					_go_fuzz_dep_.CoverTab[100850]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:793
					// _ = "end of CoverTab[100850]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:794
					_go_fuzz_dep_.CoverTab[100851]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:794
					// _ = "end of CoverTab[100851]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:794
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:794
				// _ = "end of CoverTab[100848]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:794
				_go_fuzz_dep_.CoverTab[100849]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:795
				// _ = "end of CoverTab[100849]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:796
				_go_fuzz_dep_.CoverTab[100852]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:796
				// _ = "end of CoverTab[100852]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:796
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:796
			// _ = "end of CoverTab[100835]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:796
			_go_fuzz_dep_.CoverTab[100836]++
													if isControl {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:797
				_go_fuzz_dep_.CoverTab[100853]++
														controlRecord, err := records.getControlRecord()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:799
					_go_fuzz_dep_.CoverTab[100856]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:800
					// _ = "end of CoverTab[100856]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:801
					_go_fuzz_dep_.CoverTab[100857]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:801
					// _ = "end of CoverTab[100857]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:801
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:801
				// _ = "end of CoverTab[100853]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:801
				_go_fuzz_dep_.CoverTab[100854]++

														if controlRecord.Type == ControlRecordAbort {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:803
					_go_fuzz_dep_.CoverTab[100858]++
															delete(abortedProducerIDs, records.RecordBatch.ProducerID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:804
					// _ = "end of CoverTab[100858]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:805
					_go_fuzz_dep_.CoverTab[100859]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:805
					// _ = "end of CoverTab[100859]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:805
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:805
				// _ = "end of CoverTab[100854]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:805
				_go_fuzz_dep_.CoverTab[100855]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:806
				// _ = "end of CoverTab[100855]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:807
				_go_fuzz_dep_.CoverTab[100860]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:807
				// _ = "end of CoverTab[100860]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:807
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:807
			// _ = "end of CoverTab[100836]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:807
			_go_fuzz_dep_.CoverTab[100837]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:810
			if child.conf.Consumer.IsolationLevel == ReadCommitted {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:810
				_go_fuzz_dep_.CoverTab[100861]++
														_, isAborted := abortedProducerIDs[records.RecordBatch.ProducerID]
														if records.RecordBatch.IsTransactional && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:812
					_go_fuzz_dep_.CoverTab[100862]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:812
					return isAborted
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:812
					// _ = "end of CoverTab[100862]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:812
				}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:812
					_go_fuzz_dep_.CoverTab[100863]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:813
					// _ = "end of CoverTab[100863]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:814
					_go_fuzz_dep_.CoverTab[100864]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:814
					// _ = "end of CoverTab[100864]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:814
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:814
				// _ = "end of CoverTab[100861]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:815
				_go_fuzz_dep_.CoverTab[100865]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:815
				// _ = "end of CoverTab[100865]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:815
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:815
			// _ = "end of CoverTab[100837]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:815
			_go_fuzz_dep_.CoverTab[100838]++

													messages = append(messages, recordBatchMessages...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:817
			// _ = "end of CoverTab[100838]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:818
			_go_fuzz_dep_.CoverTab[100839]++
													return nil, fmt.Errorf("unknown records type: %v", records.recordsType)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:819
			// _ = "end of CoverTab[100839]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:820
		// _ = "end of CoverTab[100830]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:821
	// _ = "end of CoverTab[100795]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:821
	_go_fuzz_dep_.CoverTab[100796]++

											return messages, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:823
	// _ = "end of CoverTab[100796]"
}

func (child *partitionConsumer) interceptors(msg *ConsumerMessage) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:826
	_go_fuzz_dep_.CoverTab[100866]++
											for _, interceptor := range child.conf.Consumer.Interceptors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:827
		_go_fuzz_dep_.CoverTab[100867]++
												msg.safelyApplyInterceptor(interceptor)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:828
		// _ = "end of CoverTab[100867]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:829
	// _ = "end of CoverTab[100866]"
}

// Pause implements PartitionConsumer.
func (child *partitionConsumer) Pause() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:833
	_go_fuzz_dep_.CoverTab[100868]++
											atomic.StoreInt32(&child.paused, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:834
	// _ = "end of CoverTab[100868]"
}

// Resume implements PartitionConsumer.
func (child *partitionConsumer) Resume() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:838
	_go_fuzz_dep_.CoverTab[100869]++
											atomic.StoreInt32(&child.paused, 0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:839
	// _ = "end of CoverTab[100869]"
}

// IsPaused implements PartitionConsumer.
func (child *partitionConsumer) IsPaused() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:843
	_go_fuzz_dep_.CoverTab[100870]++
											return atomic.LoadInt32(&child.paused) == 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:844
	// _ = "end of CoverTab[100870]"
}

type brokerConsumer struct {
	consumer		*consumer
	broker			*Broker
	input			chan *partitionConsumer
	newSubscriptions	chan []*partitionConsumer
	subscriptions		map[*partitionConsumer]none
	wait			chan none
	acks			sync.WaitGroup
	refs			int
}

func (c *consumer) newBrokerConsumer(broker *Broker) *brokerConsumer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:858
	_go_fuzz_dep_.CoverTab[100871]++
											bc := &brokerConsumer{
		consumer:		c,
		broker:			broker,
		input:			make(chan *partitionConsumer),
		newSubscriptions:	make(chan []*partitionConsumer),
		wait:			make(chan none),
		subscriptions:		make(map[*partitionConsumer]none),
		refs:			0,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:867
	_curRoutineNum130_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:867
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum130_)

											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
		_go_fuzz_dep_.CoverTab[100872]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
			_go_fuzz_dep_.CoverTab[100873]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum130_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
			// _ = "end of CoverTab[100873]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
		withRecover(bc.subscriptionManager)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
		// _ = "end of CoverTab[100872]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
	_curRoutineNum131_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:869
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum131_)
											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
		_go_fuzz_dep_.CoverTab[100874]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
			_go_fuzz_dep_.CoverTab[100875]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum131_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
			// _ = "end of CoverTab[100875]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
		withRecover(bc.subscriptionConsumer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
		// _ = "end of CoverTab[100874]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:870
	}()

											return bc
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:872
	// _ = "end of CoverTab[100871]"
}

// The subscriptionManager constantly accepts new subscriptions on `input` (even when the main subscriptionConsumer
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:875
// goroutine is in the middle of a network request) and batches it up. The main worker goroutine picks
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:875
// up a batch of new subscriptions between every network request by reading from `newSubscriptions`, so we give
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:875
// it nil if no new subscriptions are available. We also write to `wait` only when new subscriptions is available,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:875
// so the main goroutine can block waiting for work if it has none.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:880
func (bc *brokerConsumer) subscriptionManager() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:880
	_go_fuzz_dep_.CoverTab[100876]++
											var buffer []*partitionConsumer

											for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:883
		_go_fuzz_dep_.CoverTab[100879]++
												if len(buffer) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:884
			_go_fuzz_dep_.CoverTab[100880]++
													select {
			case event, ok := <-bc.input:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:886
				_go_fuzz_dep_.CoverTab[100881]++
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:887
					_go_fuzz_dep_.CoverTab[100885]++
															goto done
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:888
					// _ = "end of CoverTab[100885]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:889
					_go_fuzz_dep_.CoverTab[100886]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:889
					// _ = "end of CoverTab[100886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:889
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:889
				// _ = "end of CoverTab[100881]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:889
				_go_fuzz_dep_.CoverTab[100882]++
														buffer = append(buffer, event)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:890
				// _ = "end of CoverTab[100882]"
			case bc.newSubscriptions <- buffer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:891
				_go_fuzz_dep_.CoverTab[100883]++
														buffer = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:892
				// _ = "end of CoverTab[100883]"
			case bc.wait <- none{}:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:893
				_go_fuzz_dep_.CoverTab[100884]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:893
				// _ = "end of CoverTab[100884]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:894
			// _ = "end of CoverTab[100880]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:895
			_go_fuzz_dep_.CoverTab[100887]++
													select {
			case event, ok := <-bc.input:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:897
				_go_fuzz_dep_.CoverTab[100888]++
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:898
					_go_fuzz_dep_.CoverTab[100891]++
															goto done
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:899
					// _ = "end of CoverTab[100891]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:900
					_go_fuzz_dep_.CoverTab[100892]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:900
					// _ = "end of CoverTab[100892]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:900
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:900
				// _ = "end of CoverTab[100888]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:900
				_go_fuzz_dep_.CoverTab[100889]++
														buffer = append(buffer, event)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:901
				// _ = "end of CoverTab[100889]"
			case bc.newSubscriptions <- nil:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:902
				_go_fuzz_dep_.CoverTab[100890]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:902
				// _ = "end of CoverTab[100890]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:903
			// _ = "end of CoverTab[100887]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:904
		// _ = "end of CoverTab[100879]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:905
	// _ = "end of CoverTab[100876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:905
	_go_fuzz_dep_.CoverTab[100877]++

done:
	close(bc.wait)
	if len(buffer) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:909
		_go_fuzz_dep_.CoverTab[100893]++
												bc.newSubscriptions <- buffer
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:910
		// _ = "end of CoverTab[100893]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:911
		_go_fuzz_dep_.CoverTab[100894]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:911
		// _ = "end of CoverTab[100894]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:911
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:911
	// _ = "end of CoverTab[100877]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:911
	_go_fuzz_dep_.CoverTab[100878]++
											close(bc.newSubscriptions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:912
	// _ = "end of CoverTab[100878]"
}

// subscriptionConsumer ensures we will get nil right away if no new subscriptions is available
func (bc *brokerConsumer) subscriptionConsumer() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:916
	_go_fuzz_dep_.CoverTab[100895]++
											<-bc.wait

											for newSubscriptions := range bc.newSubscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:919
		_go_fuzz_dep_.CoverTab[100896]++
												bc.updateSubscriptions(newSubscriptions)

												if len(bc.subscriptions) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:922
			_go_fuzz_dep_.CoverTab[100900]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:925
			<-bc.wait
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:926
			// _ = "end of CoverTab[100900]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:927
			_go_fuzz_dep_.CoverTab[100901]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:927
			// _ = "end of CoverTab[100901]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:927
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:927
		// _ = "end of CoverTab[100896]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:927
		_go_fuzz_dep_.CoverTab[100897]++

												response, err := bc.fetchNewMessages()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:930
			_go_fuzz_dep_.CoverTab[100902]++
													Logger.Printf("consumer/broker/%d disconnecting due to error processing FetchRequest: %s\n", bc.broker.ID(), err)
													bc.abort(err)
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:933
			// _ = "end of CoverTab[100902]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:934
			_go_fuzz_dep_.CoverTab[100903]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:934
			// _ = "end of CoverTab[100903]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:934
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:934
		// _ = "end of CoverTab[100897]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:934
		_go_fuzz_dep_.CoverTab[100898]++

												bc.acks.Add(len(bc.subscriptions))
												for child := range bc.subscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:937
			_go_fuzz_dep_.CoverTab[100904]++
													child.feeder <- response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:938
			// _ = "end of CoverTab[100904]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:939
		// _ = "end of CoverTab[100898]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:939
		_go_fuzz_dep_.CoverTab[100899]++
												bc.acks.Wait()
												bc.handleResponses()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:941
		// _ = "end of CoverTab[100899]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:942
	// _ = "end of CoverTab[100895]"
}

func (bc *brokerConsumer) updateSubscriptions(newSubscriptions []*partitionConsumer) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:945
	_go_fuzz_dep_.CoverTab[100905]++
											for _, child := range newSubscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:946
		_go_fuzz_dep_.CoverTab[100907]++
												bc.subscriptions[child] = none{}
												Logger.Printf("consumer/broker/%d added subscription to %s/%d\n", bc.broker.ID(), child.topic, child.partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:948
		// _ = "end of CoverTab[100907]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:949
	// _ = "end of CoverTab[100905]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:949
	_go_fuzz_dep_.CoverTab[100906]++

											for child := range bc.subscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:951
		_go_fuzz_dep_.CoverTab[100908]++
												select {
		case <-child.dying:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:953
			_go_fuzz_dep_.CoverTab[100909]++
													Logger.Printf("consumer/broker/%d closed dead subscription to %s/%d\n", bc.broker.ID(), child.topic, child.partition)
													close(child.trigger)
													delete(bc.subscriptions, child)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:956
			// _ = "end of CoverTab[100909]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:957
			_go_fuzz_dep_.CoverTab[100910]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:957
			// _ = "end of CoverTab[100910]"

		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:959
		// _ = "end of CoverTab[100908]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:960
	// _ = "end of CoverTab[100906]"
}

// handleResponses handles the response codes left for us by our subscriptions, and abandons ones that have been closed
func (bc *brokerConsumer) handleResponses() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:964
	_go_fuzz_dep_.CoverTab[100911]++
											for child := range bc.subscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:965
		_go_fuzz_dep_.CoverTab[100912]++
												result := child.responseResult
												child.responseResult = nil

												if result == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:969
			_go_fuzz_dep_.CoverTab[100914]++
													if preferredBroker, err := child.preferredBroker(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:970
				_go_fuzz_dep_.CoverTab[100916]++
														if bc.broker.ID() != preferredBroker.ID() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:971
					_go_fuzz_dep_.CoverTab[100917]++

															Logger.Printf(
						"consumer/broker/%d abandoned in favor of preferred replica broker/%d\n",
						bc.broker.ID(), preferredBroker.ID())
															child.trigger <- none{}
															delete(bc.subscriptions, child)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:977
					// _ = "end of CoverTab[100917]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:978
					_go_fuzz_dep_.CoverTab[100918]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:978
					// _ = "end of CoverTab[100918]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:978
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:978
				// _ = "end of CoverTab[100916]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:979
				_go_fuzz_dep_.CoverTab[100919]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:979
				// _ = "end of CoverTab[100919]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:979
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:979
			// _ = "end of CoverTab[100914]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:979
			_go_fuzz_dep_.CoverTab[100915]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:980
			// _ = "end of CoverTab[100915]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:981
			_go_fuzz_dep_.CoverTab[100920]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:981
			// _ = "end of CoverTab[100920]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:981
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:981
		// _ = "end of CoverTab[100912]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:981
		_go_fuzz_dep_.CoverTab[100913]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:984
		child.preferredReadReplica = invalidPreferredReplicaID

		switch result {
		case errTimedOut:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:987
			_go_fuzz_dep_.CoverTab[100921]++
													Logger.Printf("consumer/broker/%d abandoned subscription to %s/%d because consuming was taking too long\n",
				bc.broker.ID(), child.topic, child.partition)
													delete(bc.subscriptions, child)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:990
			// _ = "end of CoverTab[100921]"
		case ErrOffsetOutOfRange:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:991
			_go_fuzz_dep_.CoverTab[100922]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:994
			child.sendError(result)
													Logger.Printf("consumer/%s/%d shutting down because %s\n", child.topic, child.partition, result)
													close(child.trigger)
													delete(bc.subscriptions, child)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:997
			// _ = "end of CoverTab[100922]"
		case ErrUnknownTopicOrPartition, ErrNotLeaderForPartition, ErrLeaderNotAvailable, ErrReplicaNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:998
			_go_fuzz_dep_.CoverTab[100923]++

													Logger.Printf("consumer/broker/%d abandoned subscription to %s/%d because %s\n",
				bc.broker.ID(), child.topic, child.partition, result)
													child.trigger <- none{}
													delete(bc.subscriptions, child)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1003
			// _ = "end of CoverTab[100923]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1004
			_go_fuzz_dep_.CoverTab[100924]++

													child.sendError(result)
													Logger.Printf("consumer/broker/%d abandoned subscription to %s/%d because %s\n",
				bc.broker.ID(), child.topic, child.partition, result)
													child.trigger <- none{}
													delete(bc.subscriptions, child)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1010
			// _ = "end of CoverTab[100924]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1011
		// _ = "end of CoverTab[100913]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1012
	// _ = "end of CoverTab[100911]"
}

func (bc *brokerConsumer) abort(err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1015
	_go_fuzz_dep_.CoverTab[100925]++
											bc.consumer.abandonBrokerConsumer(bc)
											_ = bc.broker.Close()

											for child := range bc.subscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1019
		_go_fuzz_dep_.CoverTab[100927]++
												child.sendError(err)
												child.trigger <- none{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1021
		// _ = "end of CoverTab[100927]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1022
	// _ = "end of CoverTab[100925]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1022
	_go_fuzz_dep_.CoverTab[100926]++

											for newSubscriptions := range bc.newSubscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1024
		_go_fuzz_dep_.CoverTab[100928]++
												if len(newSubscriptions) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1025
			_go_fuzz_dep_.CoverTab[100930]++
													<-bc.wait
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1027
			// _ = "end of CoverTab[100930]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1028
			_go_fuzz_dep_.CoverTab[100931]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1028
			// _ = "end of CoverTab[100931]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1028
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1028
		// _ = "end of CoverTab[100928]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1028
		_go_fuzz_dep_.CoverTab[100929]++
												for _, child := range newSubscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1029
			_go_fuzz_dep_.CoverTab[100932]++
													child.sendError(err)
													child.trigger <- none{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1031
			// _ = "end of CoverTab[100932]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1032
		// _ = "end of CoverTab[100929]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1033
	// _ = "end of CoverTab[100926]"
}

func (bc *brokerConsumer) fetchNewMessages() (*FetchResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1036
	_go_fuzz_dep_.CoverTab[100933]++
											request := &FetchRequest{
		MinBytes:	bc.consumer.conf.Consumer.Fetch.Min,
		MaxWaitTime:	int32(bc.consumer.conf.Consumer.MaxWaitTime / time.Millisecond),
	}
	if bc.consumer.conf.Version.IsAtLeast(V0_9_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1041
		_go_fuzz_dep_.CoverTab[100942]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1042
		// _ = "end of CoverTab[100942]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1043
		_go_fuzz_dep_.CoverTab[100943]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1043
		// _ = "end of CoverTab[100943]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1043
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1043
	// _ = "end of CoverTab[100933]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1043
	_go_fuzz_dep_.CoverTab[100934]++
											if bc.consumer.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1044
		_go_fuzz_dep_.CoverTab[100944]++
												request.Version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1045
		// _ = "end of CoverTab[100944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1046
		_go_fuzz_dep_.CoverTab[100945]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1046
		// _ = "end of CoverTab[100945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1046
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1046
	// _ = "end of CoverTab[100934]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1046
	_go_fuzz_dep_.CoverTab[100935]++
											if bc.consumer.conf.Version.IsAtLeast(V0_10_1_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1047
		_go_fuzz_dep_.CoverTab[100946]++
												request.Version = 3
												request.MaxBytes = MaxResponseSize
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1049
		// _ = "end of CoverTab[100946]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1050
		_go_fuzz_dep_.CoverTab[100947]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1050
		// _ = "end of CoverTab[100947]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1050
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1050
	// _ = "end of CoverTab[100935]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1050
	_go_fuzz_dep_.CoverTab[100936]++
											if bc.consumer.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1051
		_go_fuzz_dep_.CoverTab[100948]++
												request.Version = 4
												request.Isolation = bc.consumer.conf.Consumer.IsolationLevel
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1053
		// _ = "end of CoverTab[100948]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1054
		_go_fuzz_dep_.CoverTab[100949]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1054
		// _ = "end of CoverTab[100949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1054
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1054
	// _ = "end of CoverTab[100936]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1054
	_go_fuzz_dep_.CoverTab[100937]++
											if bc.consumer.conf.Version.IsAtLeast(V1_1_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1055
		_go_fuzz_dep_.CoverTab[100950]++
												request.Version = 7

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1060
		request.SessionID = 0
												request.SessionEpoch = -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1061
		// _ = "end of CoverTab[100950]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1062
		_go_fuzz_dep_.CoverTab[100951]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1062
		// _ = "end of CoverTab[100951]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1062
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1062
	// _ = "end of CoverTab[100937]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1062
	_go_fuzz_dep_.CoverTab[100938]++
											if bc.consumer.conf.Version.IsAtLeast(V2_1_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1063
		_go_fuzz_dep_.CoverTab[100952]++
												request.Version = 10
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1064
		// _ = "end of CoverTab[100952]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1065
		_go_fuzz_dep_.CoverTab[100953]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1065
		// _ = "end of CoverTab[100953]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1065
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1065
	// _ = "end of CoverTab[100938]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1065
	_go_fuzz_dep_.CoverTab[100939]++
											if bc.consumer.conf.Version.IsAtLeast(V2_3_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1066
		_go_fuzz_dep_.CoverTab[100954]++
												request.Version = 11
												request.RackID = bc.consumer.conf.RackID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1068
		// _ = "end of CoverTab[100954]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1069
		_go_fuzz_dep_.CoverTab[100955]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1069
		// _ = "end of CoverTab[100955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1069
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1069
	// _ = "end of CoverTab[100939]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1069
	_go_fuzz_dep_.CoverTab[100940]++

											for child := range bc.subscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1071
		_go_fuzz_dep_.CoverTab[100956]++
												if !child.IsPaused() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1072
			_go_fuzz_dep_.CoverTab[100957]++
													request.AddBlock(child.topic, child.partition, child.offset, child.fetchSize)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1073
			// _ = "end of CoverTab[100957]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1074
			_go_fuzz_dep_.CoverTab[100958]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1074
			// _ = "end of CoverTab[100958]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1074
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1074
		// _ = "end of CoverTab[100956]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1075
	// _ = "end of CoverTab[100940]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1075
	_go_fuzz_dep_.CoverTab[100941]++

											return bc.broker.Fetch(request)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1077
	// _ = "end of CoverTab[100941]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1078
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer.go:1078
var _ = _go_fuzz_dep_.CoverTab
