//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:1
)

import "sync"

// SyncProducer publishes Kafka messages, blocking until they have been acknowledged. It routes messages to the correct
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
// broker, refreshing metadata as appropriate, and parses responses for errors. You must call Close() on a producer
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
// to avoid leaks, it may not be garbage-collected automatically when it passes out of scope.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
// The SyncProducer comes with two caveats: it will generally be less efficient than the AsyncProducer, and the actual
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
// durability guarantee provided when a message is acknowledged depend on the configured value of `Producer.RequiredAcks`.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
// There are configurations where a message acknowledged by the SyncProducer can still sometimes be lost.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
// For implementation reasons, the SyncProducer requires `Producer.Return.Errors` and `Producer.Return.Successes` to
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:5
// be set to true in its configuration.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:15
type SyncProducer interface {

	// SendMessage produces a given message, and returns only when it either has
	// succeeded or failed to produce. It will return the partition and the offset
	// of the produced message, or an error if the message failed to produce.
	SendMessage(msg *ProducerMessage) (partition int32, offset int64, err error)

	// SendMessages produces a given set of messages, and returns only when all
	// messages in the set have either succeeded or failed. Note that messages
	// can succeed and fail individually; if some succeed and some fail,
	// SendMessages will return an error.
	SendMessages(msgs []*ProducerMessage) error

	// Close shuts down the producer; you must call this function before a producer
	// object passes out of scope, as it may otherwise leak memory.
	// You must call this before calling Close on the underlying client.
	Close() error
}

type syncProducer struct {
	producer	*asyncProducer
	wg		sync.WaitGroup
}

// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
func NewSyncProducer(addrs []string, config *Config) (SyncProducer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:40
	_go_fuzz_dep_.CoverTab[106843]++
												if config == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:41
		_go_fuzz_dep_.CoverTab[106847]++
													config = NewConfig()
													config.Producer.Return.Successes = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:43
		// _ = "end of CoverTab[106847]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:44
		_go_fuzz_dep_.CoverTab[106848]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:44
		// _ = "end of CoverTab[106848]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:44
	// _ = "end of CoverTab[106843]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:44
	_go_fuzz_dep_.CoverTab[106844]++

												if err := verifyProducerConfig(config); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:46
		_go_fuzz_dep_.CoverTab[106849]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:47
		// _ = "end of CoverTab[106849]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:48
		_go_fuzz_dep_.CoverTab[106850]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:48
		// _ = "end of CoverTab[106850]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:48
	// _ = "end of CoverTab[106844]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:48
	_go_fuzz_dep_.CoverTab[106845]++

												p, err := NewAsyncProducer(addrs, config)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:51
		_go_fuzz_dep_.CoverTab[106851]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:52
		// _ = "end of CoverTab[106851]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:53
		_go_fuzz_dep_.CoverTab[106852]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:53
		// _ = "end of CoverTab[106852]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:53
	// _ = "end of CoverTab[106845]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:53
	_go_fuzz_dep_.CoverTab[106846]++
												return newSyncProducerFromAsyncProducer(p.(*asyncProducer)), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:54
	// _ = "end of CoverTab[106846]"
}

// NewSyncProducerFromClient creates a new SyncProducer using the given client. It is still
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:57
// necessary to call Close() on the underlying client when shutting down this producer.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:59
func NewSyncProducerFromClient(client Client) (SyncProducer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:59
	_go_fuzz_dep_.CoverTab[106853]++
												if err := verifyProducerConfig(client.Config()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:60
		_go_fuzz_dep_.CoverTab[106856]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:61
		// _ = "end of CoverTab[106856]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:62
		_go_fuzz_dep_.CoverTab[106857]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:62
		// _ = "end of CoverTab[106857]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:62
	// _ = "end of CoverTab[106853]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:62
	_go_fuzz_dep_.CoverTab[106854]++

												p, err := NewAsyncProducerFromClient(client)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:65
		_go_fuzz_dep_.CoverTab[106858]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:66
		// _ = "end of CoverTab[106858]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:67
		_go_fuzz_dep_.CoverTab[106859]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:67
		// _ = "end of CoverTab[106859]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:67
	// _ = "end of CoverTab[106854]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:67
	_go_fuzz_dep_.CoverTab[106855]++
												return newSyncProducerFromAsyncProducer(p.(*asyncProducer)), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:68
	// _ = "end of CoverTab[106855]"
}

func newSyncProducerFromAsyncProducer(p *asyncProducer) *syncProducer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:71
	_go_fuzz_dep_.CoverTab[106860]++
												sp := &syncProducer{producer: p}

												sp.wg.Add(2)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:74
	_curRoutineNum146_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:74
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum146_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
		_go_fuzz_dep_.CoverTab[106861]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
			_go_fuzz_dep_.CoverTab[106862]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum146_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
			// _ = "end of CoverTab[106862]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
		withRecover(sp.handleSuccesses)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
		// _ = "end of CoverTab[106861]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
	_curRoutineNum147_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:75
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum147_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
		_go_fuzz_dep_.CoverTab[106863]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
			_go_fuzz_dep_.CoverTab[106864]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum147_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
			// _ = "end of CoverTab[106864]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
		withRecover(sp.handleErrors)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
		// _ = "end of CoverTab[106863]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:76
	}()

												return sp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:78
	// _ = "end of CoverTab[106860]"
}

func verifyProducerConfig(config *Config) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:81
	_go_fuzz_dep_.CoverTab[106865]++
												if !config.Producer.Return.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:82
		_go_fuzz_dep_.CoverTab[106868]++
													return ConfigurationError("Producer.Return.Errors must be true to be used in a SyncProducer")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:83
		// _ = "end of CoverTab[106868]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:84
		_go_fuzz_dep_.CoverTab[106869]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:84
		// _ = "end of CoverTab[106869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:84
	// _ = "end of CoverTab[106865]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:84
	_go_fuzz_dep_.CoverTab[106866]++
												if !config.Producer.Return.Successes {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:85
		_go_fuzz_dep_.CoverTab[106870]++
													return ConfigurationError("Producer.Return.Successes must be true to be used in a SyncProducer")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:86
		// _ = "end of CoverTab[106870]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:87
		_go_fuzz_dep_.CoverTab[106871]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:87
		// _ = "end of CoverTab[106871]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:87
	// _ = "end of CoverTab[106866]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:87
	_go_fuzz_dep_.CoverTab[106867]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:88
	// _ = "end of CoverTab[106867]"
}

func (sp *syncProducer) SendMessage(msg *ProducerMessage) (partition int32, offset int64, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:91
	_go_fuzz_dep_.CoverTab[106872]++
												expectation := make(chan *ProducerError, 1)
												msg.expectation = expectation
												sp.producer.Input() <- msg

												if pErr := <-expectation; pErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:96
		_go_fuzz_dep_.CoverTab[106874]++
													return -1, -1, pErr.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:97
		// _ = "end of CoverTab[106874]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:98
		_go_fuzz_dep_.CoverTab[106875]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:98
		// _ = "end of CoverTab[106875]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:98
	// _ = "end of CoverTab[106872]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:98
	_go_fuzz_dep_.CoverTab[106873]++

												return msg.Partition, msg.Offset, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:100
	// _ = "end of CoverTab[106873]"
}

func (sp *syncProducer) SendMessages(msgs []*ProducerMessage) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:103
	_go_fuzz_dep_.CoverTab[106876]++
												expectations := make(chan chan *ProducerError, len(msgs))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:104
	_curRoutineNum148_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:104
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum148_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:105
		_go_fuzz_dep_.CoverTab[106880]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:105
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:105
			_go_fuzz_dep_.CoverTab[106882]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:105
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum148_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:105
			// _ = "end of CoverTab[106882]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:105
		}()
													for _, msg := range msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:106
			_go_fuzz_dep_.CoverTab[106883]++
														expectation := make(chan *ProducerError, 1)
														msg.expectation = expectation
														sp.producer.Input() <- msg
														expectations <- expectation
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:110
			// _ = "end of CoverTab[106883]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:111
		// _ = "end of CoverTab[106880]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:111
		_go_fuzz_dep_.CoverTab[106881]++
													close(expectations)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:112
		// _ = "end of CoverTab[106881]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:113
	// _ = "end of CoverTab[106876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:113
	_go_fuzz_dep_.CoverTab[106877]++

												var errors ProducerErrors
												for expectation := range expectations {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:116
		_go_fuzz_dep_.CoverTab[106884]++
													if pErr := <-expectation; pErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:117
			_go_fuzz_dep_.CoverTab[106885]++
														errors = append(errors, pErr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:118
			// _ = "end of CoverTab[106885]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:119
			_go_fuzz_dep_.CoverTab[106886]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:119
			// _ = "end of CoverTab[106886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:119
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:119
		// _ = "end of CoverTab[106884]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:120
	// _ = "end of CoverTab[106877]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:120
	_go_fuzz_dep_.CoverTab[106878]++

												if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:122
		_go_fuzz_dep_.CoverTab[106887]++
													return errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:123
		// _ = "end of CoverTab[106887]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:124
		_go_fuzz_dep_.CoverTab[106888]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:124
		// _ = "end of CoverTab[106888]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:124
	// _ = "end of CoverTab[106878]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:124
	_go_fuzz_dep_.CoverTab[106879]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:125
	// _ = "end of CoverTab[106879]"
}

func (sp *syncProducer) handleSuccesses() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:128
	_go_fuzz_dep_.CoverTab[106889]++
												defer sp.wg.Done()
												for msg := range sp.producer.Successes() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:130
		_go_fuzz_dep_.CoverTab[106890]++
													expectation := msg.expectation
													expectation <- nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:132
		// _ = "end of CoverTab[106890]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:133
	// _ = "end of CoverTab[106889]"
}

func (sp *syncProducer) handleErrors() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:136
	_go_fuzz_dep_.CoverTab[106891]++
												defer sp.wg.Done()
												for err := range sp.producer.Errors() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:138
		_go_fuzz_dep_.CoverTab[106892]++
													expectation := err.Msg.expectation
													expectation <- err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:140
		// _ = "end of CoverTab[106892]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:141
	// _ = "end of CoverTab[106891]"
}

func (sp *syncProducer) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:144
	_go_fuzz_dep_.CoverTab[106893]++
												sp.producer.AsyncClose()
												sp.wg.Wait()
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:147
	// _ = "end of CoverTab[106893]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:148
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_producer.go:148
var _ = _go_fuzz_dep_.CoverTab
