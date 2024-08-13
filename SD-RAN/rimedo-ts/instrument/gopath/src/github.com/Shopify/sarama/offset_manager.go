//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:1
)

import (
	"sync"
	"time"
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:10
// OffsetManager uses Kafka to store and fetch consumed partition offsets.
type OffsetManager interface {
	// ManagePartition creates a PartitionOffsetManager on the given topic/partition.
	// It will return an error if this OffsetManager is already managing the given
	// topic/partition.
	ManagePartition(topic string, partition int32) (PartitionOffsetManager, error)

	// Close stops the OffsetManager from managing offsets. It is required to call
	// this function before an OffsetManager object passes out of scope, as it
	// will otherwise leak memory. You must call this after all the
	// PartitionOffsetManagers are closed.
	Close() error

	// Commit commits the offsets. This method can be used if AutoCommit.Enable is
	// set to false.
	Commit()
}

type offsetManager struct {
	client	Client
	conf	*Config
	group	string
	ticker	*time.Ticker

	memberID	string
	generation	int32

	broker		*Broker
	brokerLock	sync.RWMutex

	poms		map[string]map[int32]*partitionOffsetManager
	pomsLock	sync.RWMutex

	closeOnce	sync.Once
	closing		chan none
	closed		chan none
}

// NewOffsetManagerFromClient creates a new OffsetManager from the given client.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:48
// It is still necessary to call Close() on the underlying client when finished with the partition manager.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:50
func NewOffsetManagerFromClient(group string, client Client) (OffsetManager, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:50
	_go_fuzz_dep_.CoverTab[105194]++
												return newOffsetManagerFromClient(group, "", GroupGenerationUndefined, client)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:51
	// _ = "end of CoverTab[105194]"
}

func newOffsetManagerFromClient(group, memberID string, generation int32, client Client) (*offsetManager, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:54
	_go_fuzz_dep_.CoverTab[105195]++

												if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:56
		_go_fuzz_dep_.CoverTab[105198]++
													return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:57
		// _ = "end of CoverTab[105198]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:58
		_go_fuzz_dep_.CoverTab[105199]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:58
		// _ = "end of CoverTab[105199]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:58
	// _ = "end of CoverTab[105195]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:58
	_go_fuzz_dep_.CoverTab[105196]++

												conf := client.Config()
												om := &offsetManager{
		client:	client,
		conf:	conf,
		group:	group,
		poms:	make(map[string]map[int32]*partitionOffsetManager),

		memberID:	memberID,
		generation:	generation,

		closing:	make(chan none),
		closed:		make(chan none),
	}
	if conf.Consumer.Offsets.AutoCommit.Enable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:73
		_go_fuzz_dep_.CoverTab[105200]++
													om.ticker = time.NewTicker(conf.Consumer.Offsets.AutoCommit.Interval)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:74
		_curRoutineNum145_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:74
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum145_)
													go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
			_go_fuzz_dep_.CoverTab[105201]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
				_go_fuzz_dep_.CoverTab[105202]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum145_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
				// _ = "end of CoverTab[105202]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
			}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
			withRecover(om.mainLoop)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
			// _ = "end of CoverTab[105201]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:75
		// _ = "end of CoverTab[105200]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:76
		_go_fuzz_dep_.CoverTab[105203]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:76
		// _ = "end of CoverTab[105203]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:76
	// _ = "end of CoverTab[105196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:76
	_go_fuzz_dep_.CoverTab[105197]++

												return om, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:78
	// _ = "end of CoverTab[105197]"
}

func (om *offsetManager) ManagePartition(topic string, partition int32) (PartitionOffsetManager, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:81
	_go_fuzz_dep_.CoverTab[105204]++
												pom, err := om.newPartitionOffsetManager(topic, partition)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:83
		_go_fuzz_dep_.CoverTab[105208]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:84
		// _ = "end of CoverTab[105208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:85
		_go_fuzz_dep_.CoverTab[105209]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:85
		// _ = "end of CoverTab[105209]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:85
	// _ = "end of CoverTab[105204]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:85
	_go_fuzz_dep_.CoverTab[105205]++

												om.pomsLock.Lock()
												defer om.pomsLock.Unlock()

												topicManagers := om.poms[topic]
												if topicManagers == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:91
		_go_fuzz_dep_.CoverTab[105210]++
													topicManagers = make(map[int32]*partitionOffsetManager)
													om.poms[topic] = topicManagers
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:93
		// _ = "end of CoverTab[105210]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:94
		_go_fuzz_dep_.CoverTab[105211]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:94
		// _ = "end of CoverTab[105211]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:94
	// _ = "end of CoverTab[105205]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:94
	_go_fuzz_dep_.CoverTab[105206]++

												if topicManagers[partition] != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:96
		_go_fuzz_dep_.CoverTab[105212]++
													return nil, ConfigurationError("That topic/partition is already being managed")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:97
		// _ = "end of CoverTab[105212]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:98
		_go_fuzz_dep_.CoverTab[105213]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:98
		// _ = "end of CoverTab[105213]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:98
	// _ = "end of CoverTab[105206]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:98
	_go_fuzz_dep_.CoverTab[105207]++

												topicManagers[partition] = pom
												return pom, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:101
	// _ = "end of CoverTab[105207]"
}

func (om *offsetManager) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:104
	_go_fuzz_dep_.CoverTab[105214]++
												om.closeOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:105
		_go_fuzz_dep_.CoverTab[105216]++

													close(om.closing)
													if om.conf.Consumer.Offsets.AutoCommit.Enable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:108
			_go_fuzz_dep_.CoverTab[105219]++
														<-om.closed
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:109
			// _ = "end of CoverTab[105219]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:110
			_go_fuzz_dep_.CoverTab[105220]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:110
			// _ = "end of CoverTab[105220]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:110
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:110
		// _ = "end of CoverTab[105216]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:110
		_go_fuzz_dep_.CoverTab[105217]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:113
		om.asyncClosePOMs()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:116
		if om.conf.Consumer.Offsets.AutoCommit.Enable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:116
			_go_fuzz_dep_.CoverTab[105221]++
														for attempt := 0; attempt <= om.conf.Consumer.Offsets.Retry.Max; attempt++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:117
				_go_fuzz_dep_.CoverTab[105222]++
															om.flushToBroker()
															if om.releasePOMs(false) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:119
					_go_fuzz_dep_.CoverTab[105223]++
																break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:120
					// _ = "end of CoverTab[105223]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:121
					_go_fuzz_dep_.CoverTab[105224]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:121
					// _ = "end of CoverTab[105224]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:121
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:121
				// _ = "end of CoverTab[105222]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:122
			// _ = "end of CoverTab[105221]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:123
			_go_fuzz_dep_.CoverTab[105225]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:123
			// _ = "end of CoverTab[105225]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:123
		// _ = "end of CoverTab[105217]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:123
		_go_fuzz_dep_.CoverTab[105218]++

													om.releasePOMs(true)
													om.brokerLock.Lock()
													om.broker = nil
													om.brokerLock.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:128
		// _ = "end of CoverTab[105218]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:129
	// _ = "end of CoverTab[105214]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:129
	_go_fuzz_dep_.CoverTab[105215]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:130
	// _ = "end of CoverTab[105215]"
}

func (om *offsetManager) computeBackoff(retries int) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:133
	_go_fuzz_dep_.CoverTab[105226]++
												if om.conf.Metadata.Retry.BackoffFunc != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:134
		_go_fuzz_dep_.CoverTab[105227]++
													return om.conf.Metadata.Retry.BackoffFunc(retries, om.conf.Metadata.Retry.Max)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:135
		// _ = "end of CoverTab[105227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:136
		_go_fuzz_dep_.CoverTab[105228]++
													return om.conf.Metadata.Retry.Backoff
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:137
		// _ = "end of CoverTab[105228]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:138
	// _ = "end of CoverTab[105226]"
}

func (om *offsetManager) fetchInitialOffset(topic string, partition int32, retries int) (int64, string, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:141
	_go_fuzz_dep_.CoverTab[105229]++
												broker, err := om.coordinator()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:143
		_go_fuzz_dep_.CoverTab[105233]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:144
			_go_fuzz_dep_.CoverTab[105235]++
														return 0, "", err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:145
			// _ = "end of CoverTab[105235]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:146
			_go_fuzz_dep_.CoverTab[105236]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:146
			// _ = "end of CoverTab[105236]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:146
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:146
		// _ = "end of CoverTab[105233]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:146
		_go_fuzz_dep_.CoverTab[105234]++
													return om.fetchInitialOffset(topic, partition, retries-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:147
		// _ = "end of CoverTab[105234]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:148
		_go_fuzz_dep_.CoverTab[105237]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:148
		// _ = "end of CoverTab[105237]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:148
	// _ = "end of CoverTab[105229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:148
	_go_fuzz_dep_.CoverTab[105230]++

												req := new(OffsetFetchRequest)
												req.Version = 1
												req.ConsumerGroup = om.group
												req.AddPartition(topic, partition)

												resp, err := broker.FetchOffset(req)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:156
		_go_fuzz_dep_.CoverTab[105238]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:157
			_go_fuzz_dep_.CoverTab[105240]++
														return 0, "", err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:158
			// _ = "end of CoverTab[105240]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:159
			_go_fuzz_dep_.CoverTab[105241]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:159
			// _ = "end of CoverTab[105241]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:159
		// _ = "end of CoverTab[105238]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:159
		_go_fuzz_dep_.CoverTab[105239]++
													om.releaseCoordinator(broker)
													return om.fetchInitialOffset(topic, partition, retries-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:161
		// _ = "end of CoverTab[105239]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:162
		_go_fuzz_dep_.CoverTab[105242]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:162
		// _ = "end of CoverTab[105242]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:162
	// _ = "end of CoverTab[105230]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:162
	_go_fuzz_dep_.CoverTab[105231]++

												block := resp.GetBlock(topic, partition)
												if block == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:165
		_go_fuzz_dep_.CoverTab[105243]++
													return 0, "", ErrIncompleteResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:166
		// _ = "end of CoverTab[105243]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:167
		_go_fuzz_dep_.CoverTab[105244]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:167
		// _ = "end of CoverTab[105244]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:167
	// _ = "end of CoverTab[105231]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:167
	_go_fuzz_dep_.CoverTab[105232]++

												switch block.Err {
	case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:170
		_go_fuzz_dep_.CoverTab[105245]++
													return block.Offset, block.Metadata, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:171
		// _ = "end of CoverTab[105245]"
	case ErrNotCoordinatorForConsumer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:172
		_go_fuzz_dep_.CoverTab[105246]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:173
			_go_fuzz_dep_.CoverTab[105252]++
														return 0, "", block.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:174
			// _ = "end of CoverTab[105252]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:175
			_go_fuzz_dep_.CoverTab[105253]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:175
			// _ = "end of CoverTab[105253]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:175
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:175
		// _ = "end of CoverTab[105246]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:175
		_go_fuzz_dep_.CoverTab[105247]++
													om.releaseCoordinator(broker)
													return om.fetchInitialOffset(topic, partition, retries-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:177
		// _ = "end of CoverTab[105247]"
	case ErrOffsetsLoadInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:178
		_go_fuzz_dep_.CoverTab[105248]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:179
			_go_fuzz_dep_.CoverTab[105254]++
														return 0, "", block.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:180
			// _ = "end of CoverTab[105254]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:181
			_go_fuzz_dep_.CoverTab[105255]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:181
			// _ = "end of CoverTab[105255]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:181
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:181
		// _ = "end of CoverTab[105248]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:181
		_go_fuzz_dep_.CoverTab[105249]++
													backoff := om.computeBackoff(retries)
													select {
		case <-om.closing:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:184
			_go_fuzz_dep_.CoverTab[105256]++
														return 0, "", block.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:185
			// _ = "end of CoverTab[105256]"
		case <-time.After(backoff):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:186
			_go_fuzz_dep_.CoverTab[105257]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:186
			// _ = "end of CoverTab[105257]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:187
		// _ = "end of CoverTab[105249]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:187
		_go_fuzz_dep_.CoverTab[105250]++
													return om.fetchInitialOffset(topic, partition, retries-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:188
		// _ = "end of CoverTab[105250]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:189
		_go_fuzz_dep_.CoverTab[105251]++
													return 0, "", block.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:190
		// _ = "end of CoverTab[105251]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:191
	// _ = "end of CoverTab[105232]"
}

func (om *offsetManager) coordinator() (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:194
	_go_fuzz_dep_.CoverTab[105258]++
												om.brokerLock.RLock()
												broker := om.broker
												om.brokerLock.RUnlock()

												if broker != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:199
		_go_fuzz_dep_.CoverTab[105263]++
													return broker, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:200
		// _ = "end of CoverTab[105263]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:201
		_go_fuzz_dep_.CoverTab[105264]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:201
		// _ = "end of CoverTab[105264]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:201
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:201
	// _ = "end of CoverTab[105258]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:201
	_go_fuzz_dep_.CoverTab[105259]++

												om.brokerLock.Lock()
												defer om.brokerLock.Unlock()

												if broker := om.broker; broker != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:206
		_go_fuzz_dep_.CoverTab[105265]++
													return broker, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:207
		// _ = "end of CoverTab[105265]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:208
		_go_fuzz_dep_.CoverTab[105266]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:208
		// _ = "end of CoverTab[105266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:208
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:208
	// _ = "end of CoverTab[105259]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:208
	_go_fuzz_dep_.CoverTab[105260]++

												if err := om.client.RefreshCoordinator(om.group); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:210
		_go_fuzz_dep_.CoverTab[105267]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:211
		// _ = "end of CoverTab[105267]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:212
		_go_fuzz_dep_.CoverTab[105268]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:212
		// _ = "end of CoverTab[105268]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:212
	// _ = "end of CoverTab[105260]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:212
	_go_fuzz_dep_.CoverTab[105261]++

												broker, err := om.client.Coordinator(om.group)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:215
		_go_fuzz_dep_.CoverTab[105269]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:216
		// _ = "end of CoverTab[105269]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:217
		_go_fuzz_dep_.CoverTab[105270]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:217
		// _ = "end of CoverTab[105270]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:217
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:217
	// _ = "end of CoverTab[105261]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:217
	_go_fuzz_dep_.CoverTab[105262]++

												om.broker = broker
												return broker, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:220
	// _ = "end of CoverTab[105262]"
}

func (om *offsetManager) releaseCoordinator(b *Broker) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:223
	_go_fuzz_dep_.CoverTab[105271]++
												om.brokerLock.Lock()
												if om.broker == b {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:225
		_go_fuzz_dep_.CoverTab[105273]++
													om.broker = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:226
		// _ = "end of CoverTab[105273]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:227
		_go_fuzz_dep_.CoverTab[105274]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:227
		// _ = "end of CoverTab[105274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:227
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:227
	// _ = "end of CoverTab[105271]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:227
	_go_fuzz_dep_.CoverTab[105272]++
												om.brokerLock.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:228
	// _ = "end of CoverTab[105272]"
}

func (om *offsetManager) mainLoop() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:231
	_go_fuzz_dep_.CoverTab[105275]++
												defer om.ticker.Stop()
												defer close(om.closed)

												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:235
		_go_fuzz_dep_.CoverTab[105276]++
													select {
		case <-om.ticker.C:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:237
			_go_fuzz_dep_.CoverTab[105277]++
														om.Commit()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:238
			// _ = "end of CoverTab[105277]"
		case <-om.closing:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:239
			_go_fuzz_dep_.CoverTab[105278]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:240
			// _ = "end of CoverTab[105278]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:241
		// _ = "end of CoverTab[105276]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:242
	// _ = "end of CoverTab[105275]"
}

func (om *offsetManager) Commit() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:245
	_go_fuzz_dep_.CoverTab[105279]++
												om.flushToBroker()
												om.releasePOMs(false)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:247
	// _ = "end of CoverTab[105279]"
}

func (om *offsetManager) flushToBroker() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:250
	_go_fuzz_dep_.CoverTab[105280]++
												req := om.constructRequest()
												if req == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:252
		_go_fuzz_dep_.CoverTab[105284]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:253
		// _ = "end of CoverTab[105284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:254
		_go_fuzz_dep_.CoverTab[105285]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:254
		// _ = "end of CoverTab[105285]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:254
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:254
	// _ = "end of CoverTab[105280]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:254
	_go_fuzz_dep_.CoverTab[105281]++

												broker, err := om.coordinator()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:257
		_go_fuzz_dep_.CoverTab[105286]++
													om.handleError(err)
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:259
		// _ = "end of CoverTab[105286]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:260
		_go_fuzz_dep_.CoverTab[105287]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:260
		// _ = "end of CoverTab[105287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:260
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:260
	// _ = "end of CoverTab[105281]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:260
	_go_fuzz_dep_.CoverTab[105282]++

												resp, err := broker.CommitOffset(req)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:263
		_go_fuzz_dep_.CoverTab[105288]++
													om.handleError(err)
													om.releaseCoordinator(broker)
													_ = broker.Close()
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:267
		// _ = "end of CoverTab[105288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:268
		_go_fuzz_dep_.CoverTab[105289]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:268
		// _ = "end of CoverTab[105289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:268
	// _ = "end of CoverTab[105282]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:268
	_go_fuzz_dep_.CoverTab[105283]++

												om.handleResponse(broker, req, resp)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:270
	// _ = "end of CoverTab[105283]"
}

func (om *offsetManager) constructRequest() *OffsetCommitRequest {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:273
	_go_fuzz_dep_.CoverTab[105290]++
												var r *OffsetCommitRequest
												var perPartitionTimestamp int64
												if om.conf.Consumer.Offsets.Retention == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:276
		_go_fuzz_dep_.CoverTab[105294]++
													perPartitionTimestamp = ReceiveTime
													r = &OffsetCommitRequest{
			Version:			1,
			ConsumerGroup:			om.group,
			ConsumerID:			om.memberID,
			ConsumerGroupGeneration:	om.generation,
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:283
		// _ = "end of CoverTab[105294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:284
		_go_fuzz_dep_.CoverTab[105295]++
													r = &OffsetCommitRequest{
			Version:			2,
			RetentionTime:			int64(om.conf.Consumer.Offsets.Retention / time.Millisecond),
			ConsumerGroup:			om.group,
			ConsumerID:			om.memberID,
			ConsumerGroupGeneration:	om.generation,
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:291
		// _ = "end of CoverTab[105295]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:292
	// _ = "end of CoverTab[105290]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:292
	_go_fuzz_dep_.CoverTab[105291]++

												om.pomsLock.RLock()
												defer om.pomsLock.RUnlock()

												for _, topicManagers := range om.poms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:297
		_go_fuzz_dep_.CoverTab[105296]++
													for _, pom := range topicManagers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:298
			_go_fuzz_dep_.CoverTab[105297]++
														pom.lock.Lock()
														if pom.dirty {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:300
				_go_fuzz_dep_.CoverTab[105299]++
															r.AddBlock(pom.topic, pom.partition, pom.offset, perPartitionTimestamp, pom.metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:301
				// _ = "end of CoverTab[105299]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:302
				_go_fuzz_dep_.CoverTab[105300]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:302
				// _ = "end of CoverTab[105300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:302
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:302
			// _ = "end of CoverTab[105297]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:302
			_go_fuzz_dep_.CoverTab[105298]++
														pom.lock.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:303
			// _ = "end of CoverTab[105298]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:304
		// _ = "end of CoverTab[105296]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:305
	// _ = "end of CoverTab[105291]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:305
	_go_fuzz_dep_.CoverTab[105292]++

												if len(r.blocks) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:307
		_go_fuzz_dep_.CoverTab[105301]++
													return r
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:308
		// _ = "end of CoverTab[105301]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:309
		_go_fuzz_dep_.CoverTab[105302]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:309
		// _ = "end of CoverTab[105302]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:309
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:309
	// _ = "end of CoverTab[105292]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:309
	_go_fuzz_dep_.CoverTab[105293]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:311
	// _ = "end of CoverTab[105293]"
}

func (om *offsetManager) handleResponse(broker *Broker, req *OffsetCommitRequest, resp *OffsetCommitResponse) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:314
	_go_fuzz_dep_.CoverTab[105303]++
												om.pomsLock.RLock()
												defer om.pomsLock.RUnlock()

												for _, topicManagers := range om.poms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:318
		_go_fuzz_dep_.CoverTab[105304]++
													for _, pom := range topicManagers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:319
			_go_fuzz_dep_.CoverTab[105305]++
														if req.blocks[pom.topic] == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:320
				_go_fuzz_dep_.CoverTab[105309]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:320
				return req.blocks[pom.topic][pom.partition] == nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:320
				// _ = "end of CoverTab[105309]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:320
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:320
				_go_fuzz_dep_.CoverTab[105310]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:321
				// _ = "end of CoverTab[105310]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:322
				_go_fuzz_dep_.CoverTab[105311]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:322
				// _ = "end of CoverTab[105311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:322
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:322
			// _ = "end of CoverTab[105305]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:322
			_go_fuzz_dep_.CoverTab[105306]++

														var err KError
														var ok bool

														if resp.Errors[pom.topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:327
				_go_fuzz_dep_.CoverTab[105312]++
															pom.handleError(ErrIncompleteResponse)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:329
				// _ = "end of CoverTab[105312]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:330
				_go_fuzz_dep_.CoverTab[105313]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:330
				// _ = "end of CoverTab[105313]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:330
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:330
			// _ = "end of CoverTab[105306]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:330
			_go_fuzz_dep_.CoverTab[105307]++
														if err, ok = resp.Errors[pom.topic][pom.partition]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:331
				_go_fuzz_dep_.CoverTab[105314]++
															pom.handleError(ErrIncompleteResponse)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:333
				// _ = "end of CoverTab[105314]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:334
				_go_fuzz_dep_.CoverTab[105315]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:334
				// _ = "end of CoverTab[105315]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:334
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:334
			// _ = "end of CoverTab[105307]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:334
			_go_fuzz_dep_.CoverTab[105308]++

														switch err {
			case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:337
				_go_fuzz_dep_.CoverTab[105316]++
															block := req.blocks[pom.topic][pom.partition]
															pom.updateCommitted(block.offset, block.metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:339
				// _ = "end of CoverTab[105316]"
			case ErrNotLeaderForPartition, ErrLeaderNotAvailable,
				ErrConsumerCoordinatorNotAvailable, ErrNotCoordinatorForConsumer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:341
				_go_fuzz_dep_.CoverTab[105317]++

															om.releaseCoordinator(broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:343
				// _ = "end of CoverTab[105317]"
			case ErrOffsetMetadataTooLarge, ErrInvalidCommitOffsetSize:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:344
				_go_fuzz_dep_.CoverTab[105318]++

															pom.handleError(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:346
				// _ = "end of CoverTab[105318]"
			case ErrOffsetsLoadInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:347
				_go_fuzz_dep_.CoverTab[105319]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:347
				// _ = "end of CoverTab[105319]"

			case ErrUnknownTopicOrPartition:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:349
				_go_fuzz_dep_.CoverTab[105320]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:354
				fallthrough
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:354
				// _ = "end of CoverTab[105320]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:355
				_go_fuzz_dep_.CoverTab[105321]++

															pom.handleError(err)
															om.releaseCoordinator(broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:358
				// _ = "end of CoverTab[105321]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:359
			// _ = "end of CoverTab[105308]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:360
		// _ = "end of CoverTab[105304]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:361
	// _ = "end of CoverTab[105303]"
}

func (om *offsetManager) handleError(err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:364
	_go_fuzz_dep_.CoverTab[105322]++
												om.pomsLock.RLock()
												defer om.pomsLock.RUnlock()

												for _, topicManagers := range om.poms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:368
		_go_fuzz_dep_.CoverTab[105323]++
													for _, pom := range topicManagers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:369
			_go_fuzz_dep_.CoverTab[105324]++
														pom.handleError(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:370
			// _ = "end of CoverTab[105324]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:371
		// _ = "end of CoverTab[105323]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:372
	// _ = "end of CoverTab[105322]"
}

func (om *offsetManager) asyncClosePOMs() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:375
	_go_fuzz_dep_.CoverTab[105325]++
												om.pomsLock.RLock()
												defer om.pomsLock.RUnlock()

												for _, topicManagers := range om.poms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:379
		_go_fuzz_dep_.CoverTab[105326]++
													for _, pom := range topicManagers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:380
			_go_fuzz_dep_.CoverTab[105327]++
														pom.AsyncClose()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:381
			// _ = "end of CoverTab[105327]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:382
		// _ = "end of CoverTab[105326]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:383
	// _ = "end of CoverTab[105325]"
}

// Releases/removes closed POMs once they are clean (or when forced)
func (om *offsetManager) releasePOMs(force bool) (remaining int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:387
	_go_fuzz_dep_.CoverTab[105328]++
												om.pomsLock.Lock()
												defer om.pomsLock.Unlock()

												for topic, topicManagers := range om.poms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:391
		_go_fuzz_dep_.CoverTab[105330]++
													for partition, pom := range topicManagers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:392
			_go_fuzz_dep_.CoverTab[105332]++
														pom.lock.Lock()
														releaseDue := pom.done && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
				_go_fuzz_dep_.CoverTab[105333]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
				return (force || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
					_go_fuzz_dep_.CoverTab[105334]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
					return !pom.dirty
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
					// _ = "end of CoverTab[105334]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
				}())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
				// _ = "end of CoverTab[105333]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:394
			}()
														pom.lock.Unlock()

														if releaseDue {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:397
				_go_fuzz_dep_.CoverTab[105335]++
															pom.release()

															delete(om.poms[topic], partition)
															if len(om.poms[topic]) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:401
					_go_fuzz_dep_.CoverTab[105336]++
																delete(om.poms, topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:402
					// _ = "end of CoverTab[105336]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:403
					_go_fuzz_dep_.CoverTab[105337]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:403
					// _ = "end of CoverTab[105337]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:403
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:403
				// _ = "end of CoverTab[105335]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:404
				_go_fuzz_dep_.CoverTab[105338]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:404
				// _ = "end of CoverTab[105338]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:404
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:404
			// _ = "end of CoverTab[105332]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:405
		// _ = "end of CoverTab[105330]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:405
		_go_fuzz_dep_.CoverTab[105331]++
													remaining += len(om.poms[topic])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:406
		// _ = "end of CoverTab[105331]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:407
	// _ = "end of CoverTab[105328]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:407
	_go_fuzz_dep_.CoverTab[105329]++
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:408
	// _ = "end of CoverTab[105329]"
}

func (om *offsetManager) findPOM(topic string, partition int32) *partitionOffsetManager {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:411
	_go_fuzz_dep_.CoverTab[105339]++
												om.pomsLock.RLock()
												defer om.pomsLock.RUnlock()

												if partitions, ok := om.poms[topic]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:415
		_go_fuzz_dep_.CoverTab[105341]++
													if pom, ok := partitions[partition]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:416
			_go_fuzz_dep_.CoverTab[105342]++
														return pom
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:417
			// _ = "end of CoverTab[105342]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:418
			_go_fuzz_dep_.CoverTab[105343]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:418
			// _ = "end of CoverTab[105343]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:418
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:418
		// _ = "end of CoverTab[105341]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:419
		_go_fuzz_dep_.CoverTab[105344]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:419
		// _ = "end of CoverTab[105344]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:419
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:419
	// _ = "end of CoverTab[105339]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:419
	_go_fuzz_dep_.CoverTab[105340]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:420
	// _ = "end of CoverTab[105340]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:425
// PartitionOffsetManager uses Kafka to store and fetch consumed partition offsets. You MUST call Close()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:425
// on a partition offset manager to avoid leaks, it will not be garbage-collected automatically when it passes
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:425
// out of scope.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:428
type PartitionOffsetManager interface {
	// NextOffset returns the next offset that should be consumed for the managed
	// partition, accompanied by metadata which can be used to reconstruct the state
	// of the partition consumer when it resumes. NextOffset() will return
	// `config.Consumer.Offsets.Initial` and an empty metadata string if no offset
	// was committed for this partition yet.
	NextOffset() (int64, string)

	// MarkOffset marks the provided offset, alongside a metadata string
	// that represents the state of the partition consumer at that point in time. The
	// metadata string can be used by another consumer to restore that state, so it
	// can resume consumption.
	//
	// To follow upstream conventions, you are expected to mark the offset of the
	// next message to read, not the last message read. Thus, when calling `MarkOffset`
	// you should typically add one to the offset of the last consumed message.
	//
	// Note: calling MarkOffset does not necessarily commit the offset to the backend
	// store immediately for efficiency reasons, and it may never be committed if
	// your application crashes. This means that you may end up processing the same
	// message twice, and your processing should ideally be idempotent.
	MarkOffset(offset int64, metadata string)

	// ResetOffset resets to the provided offset, alongside a metadata string that
	// represents the state of the partition consumer at that point in time. Reset
	// acts as a counterpart to MarkOffset, the difference being that it allows to
	// reset an offset to an earlier or smaller value, where MarkOffset only
	// allows incrementing the offset. cf MarkOffset for more details.
	ResetOffset(offset int64, metadata string)

	// Errors returns a read channel of errors that occur during offset management, if
	// enabled. By default, errors are logged and not returned over this channel. If
	// you want to implement any custom error handling, set your config's
	// Consumer.Return.Errors setting to true, and read from this channel.
	Errors() <-chan *ConsumerError

	// AsyncClose initiates a shutdown of the PartitionOffsetManager. This method will
	// return immediately, after which you should wait until the 'errors' channel has
	// been drained and closed. It is required to call this function, or Close before
	// a consumer object passes out of scope, as it will otherwise leak memory. You
	// must call this before calling Close on the underlying client.
	AsyncClose()

	// Close stops the PartitionOffsetManager from managing offsets. It is required to
	// call this function (or AsyncClose) before a PartitionOffsetManager object
	// passes out of scope, as it will otherwise leak memory. You must call this
	// before calling Close on the underlying client.
	Close() error
}

type partitionOffsetManager struct {
	parent		*offsetManager
	topic		string
	partition	int32

	lock		sync.Mutex
	offset		int64
	metadata	string
	dirty		bool
	done		bool

	releaseOnce	sync.Once
	errors		chan *ConsumerError
}

func (om *offsetManager) newPartitionOffsetManager(topic string, partition int32) (*partitionOffsetManager, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:493
	_go_fuzz_dep_.CoverTab[105345]++
												offset, metadata, err := om.fetchInitialOffset(topic, partition, om.conf.Metadata.Retry.Max)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:495
		_go_fuzz_dep_.CoverTab[105347]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:496
		// _ = "end of CoverTab[105347]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:497
		_go_fuzz_dep_.CoverTab[105348]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:497
		// _ = "end of CoverTab[105348]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:497
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:497
	// _ = "end of CoverTab[105345]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:497
	_go_fuzz_dep_.CoverTab[105346]++

												return &partitionOffsetManager{
		parent:		om,
		topic:		topic,
		partition:	partition,
		errors:		make(chan *ConsumerError, om.conf.ChannelBufferSize),
		offset:		offset,
		metadata:	metadata,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:506
	// _ = "end of CoverTab[105346]"
}

func (pom *partitionOffsetManager) Errors() <-chan *ConsumerError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:509
	_go_fuzz_dep_.CoverTab[105349]++
												return pom.errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:510
	// _ = "end of CoverTab[105349]"
}

func (pom *partitionOffsetManager) MarkOffset(offset int64, metadata string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:513
	_go_fuzz_dep_.CoverTab[105350]++
												pom.lock.Lock()
												defer pom.lock.Unlock()

												if offset > pom.offset {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:517
		_go_fuzz_dep_.CoverTab[105351]++
													pom.offset = offset
													pom.metadata = metadata
													pom.dirty = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:520
		// _ = "end of CoverTab[105351]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:521
		_go_fuzz_dep_.CoverTab[105352]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:521
		// _ = "end of CoverTab[105352]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:521
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:521
	// _ = "end of CoverTab[105350]"
}

func (pom *partitionOffsetManager) ResetOffset(offset int64, metadata string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:524
	_go_fuzz_dep_.CoverTab[105353]++
												pom.lock.Lock()
												defer pom.lock.Unlock()

												if offset <= pom.offset {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:528
		_go_fuzz_dep_.CoverTab[105354]++
													pom.offset = offset
													pom.metadata = metadata
													pom.dirty = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:531
		// _ = "end of CoverTab[105354]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:532
		_go_fuzz_dep_.CoverTab[105355]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:532
		// _ = "end of CoverTab[105355]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:532
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:532
	// _ = "end of CoverTab[105353]"
}

func (pom *partitionOffsetManager) updateCommitted(offset int64, metadata string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:535
	_go_fuzz_dep_.CoverTab[105356]++
												pom.lock.Lock()
												defer pom.lock.Unlock()

												if pom.offset == offset && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:539
		_go_fuzz_dep_.CoverTab[105357]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:539
		return pom.metadata == metadata
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:539
		// _ = "end of CoverTab[105357]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:539
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:539
		_go_fuzz_dep_.CoverTab[105358]++
													pom.dirty = false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:540
		// _ = "end of CoverTab[105358]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:541
		_go_fuzz_dep_.CoverTab[105359]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:541
		// _ = "end of CoverTab[105359]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:541
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:541
	// _ = "end of CoverTab[105356]"
}

func (pom *partitionOffsetManager) NextOffset() (int64, string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:544
	_go_fuzz_dep_.CoverTab[105360]++
												pom.lock.Lock()
												defer pom.lock.Unlock()

												if pom.offset >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:548
		_go_fuzz_dep_.CoverTab[105362]++
													return pom.offset, pom.metadata
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:549
		// _ = "end of CoverTab[105362]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:550
		_go_fuzz_dep_.CoverTab[105363]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:550
		// _ = "end of CoverTab[105363]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:550
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:550
	// _ = "end of CoverTab[105360]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:550
	_go_fuzz_dep_.CoverTab[105361]++

												return pom.parent.conf.Consumer.Offsets.Initial, ""
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:552
	// _ = "end of CoverTab[105361]"
}

func (pom *partitionOffsetManager) AsyncClose() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:555
	_go_fuzz_dep_.CoverTab[105364]++
												pom.lock.Lock()
												pom.done = true
												pom.lock.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:558
	// _ = "end of CoverTab[105364]"
}

func (pom *partitionOffsetManager) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:561
	_go_fuzz_dep_.CoverTab[105365]++
												pom.AsyncClose()

												var errors ConsumerErrors
												for err := range pom.errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:565
		_go_fuzz_dep_.CoverTab[105368]++
													errors = append(errors, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:566
		// _ = "end of CoverTab[105368]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:567
	// _ = "end of CoverTab[105365]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:567
	_go_fuzz_dep_.CoverTab[105366]++

												if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:569
		_go_fuzz_dep_.CoverTab[105369]++
													return errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:570
		// _ = "end of CoverTab[105369]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:571
		_go_fuzz_dep_.CoverTab[105370]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:571
		// _ = "end of CoverTab[105370]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:571
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:571
	// _ = "end of CoverTab[105366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:571
	_go_fuzz_dep_.CoverTab[105367]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:572
	// _ = "end of CoverTab[105367]"
}

func (pom *partitionOffsetManager) handleError(err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:575
	_go_fuzz_dep_.CoverTab[105371]++
												cErr := &ConsumerError{
		Topic:		pom.topic,
		Partition:	pom.partition,
		Err:		err,
	}

	if pom.parent.conf.Consumer.Return.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:582
		_go_fuzz_dep_.CoverTab[105372]++
													pom.errors <- cErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:583
		// _ = "end of CoverTab[105372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:584
		_go_fuzz_dep_.CoverTab[105373]++
													Logger.Println(cErr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:585
		// _ = "end of CoverTab[105373]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:586
	// _ = "end of CoverTab[105371]"
}

func (pom *partitionOffsetManager) release() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:589
	_go_fuzz_dep_.CoverTab[105374]++
												pom.releaseOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:590
		_go_fuzz_dep_.CoverTab[105375]++
													close(pom.errors)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:591
		// _ = "end of CoverTab[105375]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:592
	// _ = "end of CoverTab[105374]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:593
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_manager.go:593
var _ = _go_fuzz_dep_.CoverTab
