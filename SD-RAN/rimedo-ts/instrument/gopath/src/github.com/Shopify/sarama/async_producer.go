//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1
)

import (
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/eapache/go-resiliency/breaker"
	"github.com/eapache/queue"
)

// AsyncProducer publishes Kafka messages using a non-blocking API. It routes messages
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:13
// to the correct broker for the provided topic-partition, refreshing metadata as appropriate,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:13
// and parses responses for errors. You must read from the Errors() channel or the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:13
// producer will deadlock. You must call Close() or AsyncClose() on a producer to avoid
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:13
// leaks and message lost: it will not be garbage-collected automatically when it passes
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:13
// out of scope and buffered messages may not be flushed.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:19
type AsyncProducer interface {

	// AsyncClose triggers a shutdown of the producer. The shutdown has completed
	// when both the Errors and Successes channels have been closed. When calling
	// AsyncClose, you *must* continue to read from those channels in order to
	// drain the results of any messages in flight.
	AsyncClose()

	// Close shuts down the producer and waits for any buffered messages to be
	// flushed. You must call this function before a producer object passes out of
	// scope, as it may otherwise leak memory. You must call this before process
	// shutting down, or you may lose messages. You must call this before calling
	// Close on the underlying client.
	Close() error

	// Input is the input channel for the user to write messages to that they
	// wish to send.
	Input() chan<- *ProducerMessage

	// Successes is the success output channel back to the user when Return.Successes is
	// enabled. If Return.Successes is true, you MUST read from this channel or the
	// Producer will deadlock. It is suggested that you send and read messages
	// together in a single select statement.
	Successes() <-chan *ProducerMessage

	// Errors is the error output channel back to the user. You MUST read from this
	// channel or the Producer will deadlock when the channel is full. Alternatively,
	// you can set Producer.Return.Errors in your config to false, which prevents
	// errors to be returned.
	Errors() <-chan *ProducerError
}

// transactionManager keeps the state necessary to ensure idempotent production
type transactionManager struct {
	producerID	int64
	producerEpoch	int16
	sequenceNumbers	map[string]int32
	mutex		sync.Mutex
}

const (
	noProducerID	= -1
	noProducerEpoch	= -1
)

func (t *transactionManager) getAndIncrementSequenceNumber(topic string, partition int32) (int32, int16) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:64
	_go_fuzz_dep_.CoverTab[98540]++
												key := fmt.Sprintf("%s-%d", topic, partition)
												t.mutex.Lock()
												defer t.mutex.Unlock()
												sequence := t.sequenceNumbers[key]
												t.sequenceNumbers[key] = sequence + 1
												return sequence, t.producerEpoch
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:70
	// _ = "end of CoverTab[98540]"
}

func (t *transactionManager) bumpEpoch() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:73
	_go_fuzz_dep_.CoverTab[98541]++
												t.mutex.Lock()
												defer t.mutex.Unlock()
												t.producerEpoch++
												for k := range t.sequenceNumbers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:77
		_go_fuzz_dep_.CoverTab[98542]++
													t.sequenceNumbers[k] = 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:78
		// _ = "end of CoverTab[98542]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:79
	// _ = "end of CoverTab[98541]"
}

func (t *transactionManager) getProducerID() (int64, int16) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:82
	_go_fuzz_dep_.CoverTab[98543]++
												t.mutex.Lock()
												defer t.mutex.Unlock()
												return t.producerID, t.producerEpoch
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:85
	// _ = "end of CoverTab[98543]"
}

func newTransactionManager(conf *Config, client Client) (*transactionManager, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:88
	_go_fuzz_dep_.CoverTab[98544]++
												txnmgr := &transactionManager{
		producerID:	noProducerID,
		producerEpoch:	noProducerEpoch,
	}

	if conf.Producer.Idempotent {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:94
		_go_fuzz_dep_.CoverTab[98546]++
													initProducerIDResponse, err := client.InitProducerID()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:96
			_go_fuzz_dep_.CoverTab[98548]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:97
			// _ = "end of CoverTab[98548]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:98
			_go_fuzz_dep_.CoverTab[98549]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:98
			// _ = "end of CoverTab[98549]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:98
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:98
		// _ = "end of CoverTab[98546]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:98
		_go_fuzz_dep_.CoverTab[98547]++
													txnmgr.producerID = initProducerIDResponse.ProducerID
													txnmgr.producerEpoch = initProducerIDResponse.ProducerEpoch
													txnmgr.sequenceNumbers = make(map[string]int32)
													txnmgr.mutex = sync.Mutex{}

													Logger.Printf("Obtained a ProducerId: %d and ProducerEpoch: %d\n", txnmgr.producerID, txnmgr.producerEpoch)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:104
		// _ = "end of CoverTab[98547]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:105
		_go_fuzz_dep_.CoverTab[98550]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:105
		// _ = "end of CoverTab[98550]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:105
	// _ = "end of CoverTab[98544]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:105
	_go_fuzz_dep_.CoverTab[98545]++

												return txnmgr, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:107
	// _ = "end of CoverTab[98545]"
}

type asyncProducer struct {
	client	Client
	conf	*Config

	errors				chan *ProducerError
	input, successes, retries	chan *ProducerMessage
	inFlight			sync.WaitGroup

	brokers		map[*Broker]*brokerProducer
	brokerRefs	map[*brokerProducer]int
	brokerLock	sync.Mutex

	txnmgr	*transactionManager
}

// NewAsyncProducer creates a new AsyncProducer using the given broker addresses and configuration.
func NewAsyncProducer(addrs []string, conf *Config) (AsyncProducer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:126
	_go_fuzz_dep_.CoverTab[98551]++
												client, err := NewClient(addrs, conf)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:128
		_go_fuzz_dep_.CoverTab[98553]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:129
		// _ = "end of CoverTab[98553]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:130
		_go_fuzz_dep_.CoverTab[98554]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:130
		// _ = "end of CoverTab[98554]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:130
	// _ = "end of CoverTab[98551]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:130
	_go_fuzz_dep_.CoverTab[98552]++
												return newAsyncProducer(client)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:131
	// _ = "end of CoverTab[98552]"
}

// NewAsyncProducerFromClient creates a new Producer using the given client. It is still
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:134
// necessary to call Close() on the underlying client when shutting down this producer.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:136
func NewAsyncProducerFromClient(client Client) (AsyncProducer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:136
	_go_fuzz_dep_.CoverTab[98555]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:139
	cli := &nopCloserClient{client}
												return newAsyncProducer(cli)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:140
	// _ = "end of CoverTab[98555]"
}

func newAsyncProducer(client Client) (AsyncProducer, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:143
	_go_fuzz_dep_.CoverTab[98556]++

												if client.Closed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:145
		_go_fuzz_dep_.CoverTab[98559]++
													return nil, ErrClosedClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:146
		// _ = "end of CoverTab[98559]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:147
		_go_fuzz_dep_.CoverTab[98560]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:147
		// _ = "end of CoverTab[98560]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:147
	// _ = "end of CoverTab[98556]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:147
	_go_fuzz_dep_.CoverTab[98557]++

												txnmgr, err := newTransactionManager(client.Config(), client)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:150
		_go_fuzz_dep_.CoverTab[98561]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:151
		// _ = "end of CoverTab[98561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:152
		_go_fuzz_dep_.CoverTab[98562]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:152
		// _ = "end of CoverTab[98562]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:152
	// _ = "end of CoverTab[98557]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:152
	_go_fuzz_dep_.CoverTab[98558]++

												p := &asyncProducer{
		client:		client,
		conf:		client.Config(),
		errors:		make(chan *ProducerError),
		input:		make(chan *ProducerMessage),
		successes:	make(chan *ProducerMessage),
		retries:	make(chan *ProducerMessage),
		brokers:	make(map[*Broker]*brokerProducer),
		brokerRefs:	make(map[*brokerProducer]int),
		txnmgr:		txnmgr,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:164
	_curRoutineNum116_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:164
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum116_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
	go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
		_go_fuzz_dep_.CoverTab[98563]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
			_go_fuzz_dep_.CoverTab[98564]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum116_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
			// _ = "end of CoverTab[98564]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
		withRecover(p.dispatcher)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
		// _ = "end of CoverTab[98563]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
	_curRoutineNum117_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:167
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum117_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
		_go_fuzz_dep_.CoverTab[98565]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
			_go_fuzz_dep_.CoverTab[98566]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum117_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
			// _ = "end of CoverTab[98566]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
		withRecover(p.retryHandler)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
		// _ = "end of CoverTab[98565]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:168
	}()

												return p, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:170
	// _ = "end of CoverTab[98558]"
}

type flagSet int8

const (
	syn		flagSet	= 1 << iota	// first message from partitionProducer to brokerProducer
	fin					// final message from partitionProducer to brokerProducer and back
	shutdown				// start the shutdown process
)

// ProducerMessage is the collection of elements passed to the Producer in order to send a message.
type ProducerMessage struct {
	Topic	string	// The Kafka topic for this message.
	// The partitioning key for this message. Pre-existing Encoders include
	// StringEncoder and ByteEncoder.
	Key	Encoder
	// The actual message to store in Kafka. Pre-existing Encoders include
	// StringEncoder and ByteEncoder.
	Value	Encoder

	// The headers are key-value pairs that are transparently passed
	// by Kafka between producers and consumers.
	Headers	[]RecordHeader

	// This field is used to hold arbitrary data you wish to include so it
	// will be available when receiving on the Successes and Errors channels.
	// Sarama completely ignores this field and is only to be used for
	// pass-through data.
	Metadata	interface{}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:203
	// Offset is the offset of the message stored on the broker. This is only
	// guaranteed to be defined if the message was successfully delivered and
	// RequiredAcks is not NoResponse.
	Offset	int64
	// Partition is the partition that the message was sent to. This is only
	// guaranteed to be defined if the message was successfully delivered.
	Partition	int32
	// Timestamp can vary in behavior depending on broker configuration, being
	// in either one of the CreateTime or LogAppendTime modes (default CreateTime),
	// and requiring version at least 0.10.0.
	//
	// When configured to CreateTime, the timestamp is specified by the producer
	// either by explicitly setting this field, or when the message is added
	// to a produce set.
	//
	// When configured to LogAppendTime, the timestamp assigned to the message
	// by the broker. This is only guaranteed to be defined if the message was
	// successfully delivered and RequiredAcks is not NoResponse.
	Timestamp	time.Time

	retries		int
	flags		flagSet
	expectation	chan *ProducerError
	sequenceNumber	int32
	producerEpoch	int16
	hasSequence	bool
}

const producerMessageOverhead = 26	// the metadata overhead of CRC, flags, etc.

func (m *ProducerMessage) byteSize(version int) int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:233
	_go_fuzz_dep_.CoverTab[98567]++
												var size int
												if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:235
		_go_fuzz_dep_.CoverTab[98571]++
													size = maximumRecordOverhead
													for _, h := range m.Headers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:237
			_go_fuzz_dep_.CoverTab[98572]++
														size += len(h.Key) + len(h.Value) + 2*binary.MaxVarintLen32
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:238
			// _ = "end of CoverTab[98572]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:239
		// _ = "end of CoverTab[98571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:240
		_go_fuzz_dep_.CoverTab[98573]++
													size = producerMessageOverhead
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:241
		// _ = "end of CoverTab[98573]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:242
	// _ = "end of CoverTab[98567]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:242
	_go_fuzz_dep_.CoverTab[98568]++
												if m.Key != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:243
		_go_fuzz_dep_.CoverTab[98574]++
													size += m.Key.Length()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:244
		// _ = "end of CoverTab[98574]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:245
		_go_fuzz_dep_.CoverTab[98575]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:245
		// _ = "end of CoverTab[98575]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:245
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:245
	// _ = "end of CoverTab[98568]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:245
	_go_fuzz_dep_.CoverTab[98569]++
												if m.Value != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:246
		_go_fuzz_dep_.CoverTab[98576]++
													size += m.Value.Length()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:247
		// _ = "end of CoverTab[98576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:248
		_go_fuzz_dep_.CoverTab[98577]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:248
		// _ = "end of CoverTab[98577]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:248
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:248
	// _ = "end of CoverTab[98569]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:248
	_go_fuzz_dep_.CoverTab[98570]++
												return size
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:249
	// _ = "end of CoverTab[98570]"
}

func (m *ProducerMessage) clear() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:252
	_go_fuzz_dep_.CoverTab[98578]++
												m.flags = 0
												m.retries = 0
												m.sequenceNumber = 0
												m.producerEpoch = 0
												m.hasSequence = false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:257
	// _ = "end of CoverTab[98578]"
}

// ProducerError is the type of error generated when the producer fails to deliver a message.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:260
// It contains the original ProducerMessage as well as the actual error value.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:262
type ProducerError struct {
	Msg	*ProducerMessage
	Err	error
}

func (pe ProducerError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:267
	_go_fuzz_dep_.CoverTab[98579]++
												return fmt.Sprintf("kafka: Failed to produce message to topic %s: %s", pe.Msg.Topic, pe.Err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:268
	// _ = "end of CoverTab[98579]"
}

func (pe ProducerError) Unwrap() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:271
	_go_fuzz_dep_.CoverTab[98580]++
												return pe.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:272
	// _ = "end of CoverTab[98580]"
}

// ProducerErrors is a type that wraps a batch of "ProducerError"s and implements the Error interface.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:275
// It can be returned from the Producer's Close method to avoid the need to manually drain the Errors channel
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:275
// when closing a producer.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:278
type ProducerErrors []*ProducerError

func (pe ProducerErrors) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:280
	_go_fuzz_dep_.CoverTab[98581]++
												return fmt.Sprintf("kafka: Failed to deliver %d messages.", len(pe))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:281
	// _ = "end of CoverTab[98581]"
}

func (p *asyncProducer) Errors() <-chan *ProducerError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:284
	_go_fuzz_dep_.CoverTab[98582]++
												return p.errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:285
	// _ = "end of CoverTab[98582]"
}

func (p *asyncProducer) Successes() <-chan *ProducerMessage {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:288
	_go_fuzz_dep_.CoverTab[98583]++
												return p.successes
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:289
	// _ = "end of CoverTab[98583]"
}

func (p *asyncProducer) Input() chan<- *ProducerMessage {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:292
	_go_fuzz_dep_.CoverTab[98584]++
												return p.input
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:293
	// _ = "end of CoverTab[98584]"
}

func (p *asyncProducer) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:296
	_go_fuzz_dep_.CoverTab[98585]++
												p.AsyncClose()

												if p.conf.Producer.Return.Successes {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:299
		_go_fuzz_dep_.CoverTab[98589]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:299
		_curRoutineNum118_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:299
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum118_)
													go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
			_go_fuzz_dep_.CoverTab[98590]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
				_go_fuzz_dep_.CoverTab[98591]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum118_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
				// _ = "end of CoverTab[98591]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
			}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
			withRecover(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:300
				_go_fuzz_dep_.CoverTab[98592]++
															for range p.successes {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:301
					_go_fuzz_dep_.CoverTab[98593]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:301
					// _ = "end of CoverTab[98593]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:302
				// _ = "end of CoverTab[98592]"
			})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:303
			// _ = "end of CoverTab[98590]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:303
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:303
		// _ = "end of CoverTab[98589]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:304
		_go_fuzz_dep_.CoverTab[98594]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:304
		// _ = "end of CoverTab[98594]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:304
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:304
	// _ = "end of CoverTab[98585]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:304
	_go_fuzz_dep_.CoverTab[98586]++

												var errors ProducerErrors
												if p.conf.Producer.Return.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:307
		_go_fuzz_dep_.CoverTab[98595]++
													for event := range p.errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:308
			_go_fuzz_dep_.CoverTab[98596]++
														errors = append(errors, event)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:309
			// _ = "end of CoverTab[98596]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:310
		// _ = "end of CoverTab[98595]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:311
		_go_fuzz_dep_.CoverTab[98597]++
													<-p.errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:312
		// _ = "end of CoverTab[98597]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:313
	// _ = "end of CoverTab[98586]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:313
	_go_fuzz_dep_.CoverTab[98587]++

												if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:315
		_go_fuzz_dep_.CoverTab[98598]++
													return errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:316
		// _ = "end of CoverTab[98598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:317
		_go_fuzz_dep_.CoverTab[98599]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:317
		// _ = "end of CoverTab[98599]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:317
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:317
	// _ = "end of CoverTab[98587]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:317
	_go_fuzz_dep_.CoverTab[98588]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:318
	// _ = "end of CoverTab[98588]"
}

func (p *asyncProducer) AsyncClose() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:321
	_go_fuzz_dep_.CoverTab[98600]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:321
	_curRoutineNum119_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:321
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum119_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
		_go_fuzz_dep_.CoverTab[98601]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
			_go_fuzz_dep_.CoverTab[98602]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum119_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
			// _ = "end of CoverTab[98602]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
		withRecover(p.shutdown)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
		// _ = "end of CoverTab[98601]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:322
	// _ = "end of CoverTab[98600]"
}

// singleton
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:325
// dispatches messages by topic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:327
func (p *asyncProducer) dispatcher() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:327
	_go_fuzz_dep_.CoverTab[98603]++
												handlers := make(map[string]chan<- *ProducerMessage)
												shuttingDown := false

												for msg := range p.input {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:331
		_go_fuzz_dep_.CoverTab[98605]++
													if msg == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:332
			_go_fuzz_dep_.CoverTab[98612]++
														Logger.Println("Something tried to send a nil message, it was ignored.")
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:334
			// _ = "end of CoverTab[98612]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:335
			_go_fuzz_dep_.CoverTab[98613]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:335
			// _ = "end of CoverTab[98613]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:335
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:335
		// _ = "end of CoverTab[98605]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:335
		_go_fuzz_dep_.CoverTab[98606]++

													if msg.flags&shutdown != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:337
			_go_fuzz_dep_.CoverTab[98614]++
														shuttingDown = true
														p.inFlight.Done()
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:340
			// _ = "end of CoverTab[98614]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:341
			_go_fuzz_dep_.CoverTab[98615]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:341
			if msg.retries == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:341
				_go_fuzz_dep_.CoverTab[98616]++
															if shuttingDown {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:342
					_go_fuzz_dep_.CoverTab[98618]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:345
					pErr := &ProducerError{Msg: msg, Err: ErrShuttingDown}
					if p.conf.Producer.Return.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:346
						_go_fuzz_dep_.CoverTab[98620]++
																	p.errors <- pErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:347
						// _ = "end of CoverTab[98620]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:348
						_go_fuzz_dep_.CoverTab[98621]++
																	Logger.Println(pErr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:349
						// _ = "end of CoverTab[98621]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:350
					// _ = "end of CoverTab[98618]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:350
					_go_fuzz_dep_.CoverTab[98619]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:351
					// _ = "end of CoverTab[98619]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:352
					_go_fuzz_dep_.CoverTab[98622]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:352
					// _ = "end of CoverTab[98622]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:352
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:352
				// _ = "end of CoverTab[98616]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:352
				_go_fuzz_dep_.CoverTab[98617]++
															p.inFlight.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:353
				// _ = "end of CoverTab[98617]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:354
				_go_fuzz_dep_.CoverTab[98623]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:354
				// _ = "end of CoverTab[98623]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:354
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:354
			// _ = "end of CoverTab[98615]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:354
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:354
		// _ = "end of CoverTab[98606]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:354
		_go_fuzz_dep_.CoverTab[98607]++

													for _, interceptor := range p.conf.Producer.Interceptors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:356
			_go_fuzz_dep_.CoverTab[98624]++
														msg.safelyApplyInterceptor(interceptor)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:357
			// _ = "end of CoverTab[98624]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:358
		// _ = "end of CoverTab[98607]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:358
		_go_fuzz_dep_.CoverTab[98608]++

													version := 1
													if p.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:361
			_go_fuzz_dep_.CoverTab[98625]++
														version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:362
			// _ = "end of CoverTab[98625]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:363
			_go_fuzz_dep_.CoverTab[98626]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:363
			if msg.Headers != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:363
				_go_fuzz_dep_.CoverTab[98627]++
															p.returnError(msg, ConfigurationError("Producing headers requires Kafka at least v0.11"))
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:365
				// _ = "end of CoverTab[98627]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:366
				_go_fuzz_dep_.CoverTab[98628]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:366
				// _ = "end of CoverTab[98628]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:366
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:366
			// _ = "end of CoverTab[98626]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:366
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:366
		// _ = "end of CoverTab[98608]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:366
		_go_fuzz_dep_.CoverTab[98609]++
													if msg.byteSize(version) > p.conf.Producer.MaxMessageBytes {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:367
			_go_fuzz_dep_.CoverTab[98629]++
														p.returnError(msg, ErrMessageSizeTooLarge)
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:369
			// _ = "end of CoverTab[98629]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:370
			_go_fuzz_dep_.CoverTab[98630]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:370
			// _ = "end of CoverTab[98630]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:370
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:370
		// _ = "end of CoverTab[98609]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:370
		_go_fuzz_dep_.CoverTab[98610]++

													handler := handlers[msg.Topic]
													if handler == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:373
			_go_fuzz_dep_.CoverTab[98631]++
														handler = p.newTopicProducer(msg.Topic)
														handlers[msg.Topic] = handler
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:375
			// _ = "end of CoverTab[98631]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:376
			_go_fuzz_dep_.CoverTab[98632]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:376
			// _ = "end of CoverTab[98632]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:376
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:376
		// _ = "end of CoverTab[98610]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:376
		_go_fuzz_dep_.CoverTab[98611]++

													handler <- msg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:378
		// _ = "end of CoverTab[98611]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:379
	// _ = "end of CoverTab[98603]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:379
	_go_fuzz_dep_.CoverTab[98604]++

												for _, handler := range handlers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:381
		_go_fuzz_dep_.CoverTab[98633]++
													close(handler)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:382
		// _ = "end of CoverTab[98633]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:383
	// _ = "end of CoverTab[98604]"
}

// one per topic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:386
// partitions messages, then dispatches them by partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:388
type topicProducer struct {
	parent	*asyncProducer
	topic	string
	input	<-chan *ProducerMessage

	breaker		*breaker.Breaker
	handlers	map[int32]chan<- *ProducerMessage
	partitioner	Partitioner
}

func (p *asyncProducer) newTopicProducer(topic string) chan<- *ProducerMessage {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:398
	_go_fuzz_dep_.CoverTab[98634]++
												input := make(chan *ProducerMessage, p.conf.ChannelBufferSize)
												tp := &topicProducer{
		parent:		p,
		topic:		topic,
		input:		input,
		breaker:	breaker.New(3, 1, 10*time.Second),
		handlers:	make(map[int32]chan<- *ProducerMessage),
		partitioner:	p.conf.Producer.Partitioner(topic),
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:407
	_curRoutineNum120_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:407
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum120_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
		_go_fuzz_dep_.CoverTab[98635]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
			_go_fuzz_dep_.CoverTab[98636]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum120_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
			// _ = "end of CoverTab[98636]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
		withRecover(tp.dispatch)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
		// _ = "end of CoverTab[98635]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:408
	}()
												return input
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:409
	// _ = "end of CoverTab[98634]"
}

func (tp *topicProducer) dispatch() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:412
	_go_fuzz_dep_.CoverTab[98637]++
												for msg := range tp.input {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:413
		_go_fuzz_dep_.CoverTab[98639]++
													if msg.retries == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:414
			_go_fuzz_dep_.CoverTab[98642]++
														if err := tp.partitionMessage(msg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:415
				_go_fuzz_dep_.CoverTab[98643]++
															tp.parent.returnError(msg, err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:417
				// _ = "end of CoverTab[98643]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:418
				_go_fuzz_dep_.CoverTab[98644]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:418
				// _ = "end of CoverTab[98644]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:418
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:418
			// _ = "end of CoverTab[98642]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:419
			_go_fuzz_dep_.CoverTab[98645]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:419
			// _ = "end of CoverTab[98645]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:419
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:419
		// _ = "end of CoverTab[98639]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:419
		_go_fuzz_dep_.CoverTab[98640]++

													handler := tp.handlers[msg.Partition]
													if handler == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:422
			_go_fuzz_dep_.CoverTab[98646]++
														handler = tp.parent.newPartitionProducer(msg.Topic, msg.Partition)
														tp.handlers[msg.Partition] = handler
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:424
			// _ = "end of CoverTab[98646]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:425
			_go_fuzz_dep_.CoverTab[98647]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:425
			// _ = "end of CoverTab[98647]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:425
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:425
		// _ = "end of CoverTab[98640]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:425
		_go_fuzz_dep_.CoverTab[98641]++

													handler <- msg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:427
		// _ = "end of CoverTab[98641]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:428
	// _ = "end of CoverTab[98637]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:428
	_go_fuzz_dep_.CoverTab[98638]++

												for _, handler := range tp.handlers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:430
		_go_fuzz_dep_.CoverTab[98648]++
													close(handler)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:431
		// _ = "end of CoverTab[98648]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:432
	// _ = "end of CoverTab[98638]"
}

func (tp *topicProducer) partitionMessage(msg *ProducerMessage) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:435
	_go_fuzz_dep_.CoverTab[98649]++
												var partitions []int32

												err := tp.breaker.Run(func() (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:438
		_go_fuzz_dep_.CoverTab[98654]++
													requiresConsistency := false
													if ep, ok := tp.partitioner.(DynamicConsistencyPartitioner); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:440
			_go_fuzz_dep_.CoverTab[98657]++
														requiresConsistency = ep.MessageRequiresConsistency(msg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:441
			// _ = "end of CoverTab[98657]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:442
			_go_fuzz_dep_.CoverTab[98658]++
														requiresConsistency = tp.partitioner.RequiresConsistency()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:443
			// _ = "end of CoverTab[98658]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:444
		// _ = "end of CoverTab[98654]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:444
		_go_fuzz_dep_.CoverTab[98655]++

													if requiresConsistency {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:446
			_go_fuzz_dep_.CoverTab[98659]++
														partitions, err = tp.parent.client.Partitions(msg.Topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:447
			// _ = "end of CoverTab[98659]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:448
			_go_fuzz_dep_.CoverTab[98660]++
														partitions, err = tp.parent.client.WritablePartitions(msg.Topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:449
			// _ = "end of CoverTab[98660]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:450
		// _ = "end of CoverTab[98655]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:450
		_go_fuzz_dep_.CoverTab[98656]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:451
		// _ = "end of CoverTab[98656]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:452
	// _ = "end of CoverTab[98649]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:452
	_go_fuzz_dep_.CoverTab[98650]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:453
		_go_fuzz_dep_.CoverTab[98661]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:454
		// _ = "end of CoverTab[98661]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:455
		_go_fuzz_dep_.CoverTab[98662]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:455
		// _ = "end of CoverTab[98662]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:455
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:455
	// _ = "end of CoverTab[98650]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:455
	_go_fuzz_dep_.CoverTab[98651]++

												numPartitions := int32(len(partitions))

												if numPartitions == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:459
		_go_fuzz_dep_.CoverTab[98663]++
													return ErrLeaderNotAvailable
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:460
		// _ = "end of CoverTab[98663]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:461
		_go_fuzz_dep_.CoverTab[98664]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:461
		// _ = "end of CoverTab[98664]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:461
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:461
	// _ = "end of CoverTab[98651]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:461
	_go_fuzz_dep_.CoverTab[98652]++

												choice, err := tp.partitioner.Partition(msg, numPartitions)

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:465
		_go_fuzz_dep_.CoverTab[98665]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:466
		// _ = "end of CoverTab[98665]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:467
		_go_fuzz_dep_.CoverTab[98666]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:467
		if choice < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:467
			_go_fuzz_dep_.CoverTab[98667]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:467
			return choice >= numPartitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:467
			// _ = "end of CoverTab[98667]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:467
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:467
			_go_fuzz_dep_.CoverTab[98668]++
														return ErrInvalidPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:468
			// _ = "end of CoverTab[98668]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:469
			_go_fuzz_dep_.CoverTab[98669]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:469
			// _ = "end of CoverTab[98669]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:469
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:469
		// _ = "end of CoverTab[98666]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:469
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:469
	// _ = "end of CoverTab[98652]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:469
	_go_fuzz_dep_.CoverTab[98653]++

												msg.Partition = partitions[choice]

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:473
	// _ = "end of CoverTab[98653]"
}

// one per partition per topic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:476
// dispatches messages to the appropriate broker
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:476
// also responsible for maintaining message order during retries
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:479
type partitionProducer struct {
	parent		*asyncProducer
	topic		string
	partition	int32
	input		<-chan *ProducerMessage

	leader		*Broker
	breaker		*breaker.Breaker
	brokerProducer	*brokerProducer

	// highWatermark tracks the "current" retry level, which is the only one where we actually let messages through,
	// all other messages get buffered in retryState[msg.retries].buf to preserve ordering
	// retryState[msg.retries].expectChaser simply tracks whether we've seen a fin message for a given level (and
	// therefore whether our buffer is complete and safe to flush)
	highWatermark	int
	retryState	[]partitionRetryState
}

type partitionRetryState struct {
	buf		[]*ProducerMessage
	expectChaser	bool
}

func (p *asyncProducer) newPartitionProducer(topic string, partition int32) chan<- *ProducerMessage {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:502
	_go_fuzz_dep_.CoverTab[98670]++
												input := make(chan *ProducerMessage, p.conf.ChannelBufferSize)
												pp := &partitionProducer{
		parent:		p,
		topic:		topic,
		partition:	partition,
		input:		input,

		breaker:	breaker.New(3, 1, 10*time.Second),
		retryState:	make([]partitionRetryState, p.conf.Producer.Retry.Max+1),
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:512
	_curRoutineNum121_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:512
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum121_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
		_go_fuzz_dep_.CoverTab[98671]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
			_go_fuzz_dep_.CoverTab[98672]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum121_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
			// _ = "end of CoverTab[98672]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
		withRecover(pp.dispatch)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
		// _ = "end of CoverTab[98671]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:513
	}()
												return input
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:514
	// _ = "end of CoverTab[98670]"
}

func (pp *partitionProducer) backoff(retries int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:517
	_go_fuzz_dep_.CoverTab[98673]++
												var backoff time.Duration
												if pp.parent.conf.Producer.Retry.BackoffFunc != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:519
		_go_fuzz_dep_.CoverTab[98675]++
													maxRetries := pp.parent.conf.Producer.Retry.Max
													backoff = pp.parent.conf.Producer.Retry.BackoffFunc(retries, maxRetries)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:521
		// _ = "end of CoverTab[98675]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:522
		_go_fuzz_dep_.CoverTab[98676]++
													backoff = pp.parent.conf.Producer.Retry.Backoff
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:523
		// _ = "end of CoverTab[98676]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:524
	// _ = "end of CoverTab[98673]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:524
	_go_fuzz_dep_.CoverTab[98674]++
												if backoff > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:525
		_go_fuzz_dep_.CoverTab[98677]++
													time.Sleep(backoff)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:526
		// _ = "end of CoverTab[98677]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:527
		_go_fuzz_dep_.CoverTab[98678]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:527
		// _ = "end of CoverTab[98678]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:527
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:527
	// _ = "end of CoverTab[98674]"
}

func (pp *partitionProducer) dispatch() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:530
	_go_fuzz_dep_.CoverTab[98679]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:533
	pp.leader, _ = pp.parent.client.Leader(pp.topic, pp.partition)
	if pp.leader != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:534
		_go_fuzz_dep_.CoverTab[98682]++
													pp.brokerProducer = pp.parent.getBrokerProducer(pp.leader)
													pp.parent.inFlight.Add(1)
													pp.brokerProducer.input <- &ProducerMessage{Topic: pp.topic, Partition: pp.partition, flags: syn}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:537
		// _ = "end of CoverTab[98682]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:538
		_go_fuzz_dep_.CoverTab[98683]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:538
		// _ = "end of CoverTab[98683]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:538
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:538
	// _ = "end of CoverTab[98679]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:538
	_go_fuzz_dep_.CoverTab[98680]++

												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:540
		_go_fuzz_dep_.CoverTab[98684]++
													if pp.brokerProducer != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:541
			_go_fuzz_dep_.CoverTab[98685]++
														pp.parent.unrefBrokerProducer(pp.leader, pp.brokerProducer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:542
			// _ = "end of CoverTab[98685]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:543
			_go_fuzz_dep_.CoverTab[98686]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:543
			// _ = "end of CoverTab[98686]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:543
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:543
		// _ = "end of CoverTab[98684]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:544
	// _ = "end of CoverTab[98680]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:544
	_go_fuzz_dep_.CoverTab[98681]++

												for msg := range pp.input {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:546
		_go_fuzz_dep_.CoverTab[98687]++
													if pp.brokerProducer != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:547
			_go_fuzz_dep_.CoverTab[98692]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:547
			return pp.brokerProducer.abandoned != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:547
			// _ = "end of CoverTab[98692]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:547
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:547
			_go_fuzz_dep_.CoverTab[98693]++
														select {
			case <-pp.brokerProducer.abandoned:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:549
				_go_fuzz_dep_.CoverTab[98694]++

															Logger.Printf("producer/leader/%s/%d abandoning broker %d\n", pp.topic, pp.partition, pp.leader.ID())
															pp.parent.unrefBrokerProducer(pp.leader, pp.brokerProducer)
															pp.brokerProducer = nil
															time.Sleep(pp.parent.conf.Producer.Retry.Backoff)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:554
				// _ = "end of CoverTab[98694]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:555
				_go_fuzz_dep_.CoverTab[98695]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:555
				// _ = "end of CoverTab[98695]"

			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:557
			// _ = "end of CoverTab[98693]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:558
			_go_fuzz_dep_.CoverTab[98696]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:558
			// _ = "end of CoverTab[98696]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:558
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:558
		// _ = "end of CoverTab[98687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:558
		_go_fuzz_dep_.CoverTab[98688]++

													if msg.retries > pp.highWatermark {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:560
			_go_fuzz_dep_.CoverTab[98697]++

														pp.newHighWatermark(msg.retries)
														pp.backoff(msg.retries)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:563
			// _ = "end of CoverTab[98697]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:564
			_go_fuzz_dep_.CoverTab[98698]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:564
			if pp.highWatermark > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:564
				_go_fuzz_dep_.CoverTab[98699]++

															if msg.retries < pp.highWatermark {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:566
					_go_fuzz_dep_.CoverTab[98700]++

																if msg.flags&fin == fin {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:568
						_go_fuzz_dep_.CoverTab[98702]++
																	pp.retryState[msg.retries].expectChaser = false
																	pp.parent.inFlight.Done()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:570
						// _ = "end of CoverTab[98702]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:571
						_go_fuzz_dep_.CoverTab[98703]++
																	pp.retryState[msg.retries].buf = append(pp.retryState[msg.retries].buf, msg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:572
						// _ = "end of CoverTab[98703]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:573
					// _ = "end of CoverTab[98700]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:573
					_go_fuzz_dep_.CoverTab[98701]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:574
					// _ = "end of CoverTab[98701]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:575
					_go_fuzz_dep_.CoverTab[98704]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:575
					if msg.flags&fin == fin {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:575
						_go_fuzz_dep_.CoverTab[98705]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:578
						pp.retryState[pp.highWatermark].expectChaser = false
																	pp.flushRetryBuffers()
																	pp.parent.inFlight.Done()
																	continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:581
						// _ = "end of CoverTab[98705]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:582
						_go_fuzz_dep_.CoverTab[98706]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:582
						// _ = "end of CoverTab[98706]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:582
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:582
					// _ = "end of CoverTab[98704]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:582
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:582
				// _ = "end of CoverTab[98699]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:583
				_go_fuzz_dep_.CoverTab[98707]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:583
				// _ = "end of CoverTab[98707]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:583
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:583
			// _ = "end of CoverTab[98698]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:583
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:583
		// _ = "end of CoverTab[98688]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:583
		_go_fuzz_dep_.CoverTab[98689]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:588
		if pp.brokerProducer == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:588
			_go_fuzz_dep_.CoverTab[98708]++
														if err := pp.updateLeader(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:589
				_go_fuzz_dep_.CoverTab[98710]++
															pp.parent.returnError(msg, err)
															pp.backoff(msg.retries)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:592
				// _ = "end of CoverTab[98710]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:593
				_go_fuzz_dep_.CoverTab[98711]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:593
				// _ = "end of CoverTab[98711]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:593
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:593
			// _ = "end of CoverTab[98708]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:593
			_go_fuzz_dep_.CoverTab[98709]++
														Logger.Printf("producer/leader/%s/%d selected broker %d\n", pp.topic, pp.partition, pp.leader.ID())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:594
			// _ = "end of CoverTab[98709]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:595
			_go_fuzz_dep_.CoverTab[98712]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:595
			// _ = "end of CoverTab[98712]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:595
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:595
		// _ = "end of CoverTab[98689]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:595
		_go_fuzz_dep_.CoverTab[98690]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
		if pp.parent.conf.Producer.Idempotent && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
			_go_fuzz_dep_.CoverTab[98713]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
			return msg.retries == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
			// _ = "end of CoverTab[98713]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
			_go_fuzz_dep_.CoverTab[98714]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
			return msg.flags == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
			// _ = "end of CoverTab[98714]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:601
			_go_fuzz_dep_.CoverTab[98715]++
														msg.sequenceNumber, msg.producerEpoch = pp.parent.txnmgr.getAndIncrementSequenceNumber(msg.Topic, msg.Partition)
														msg.hasSequence = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:603
			// _ = "end of CoverTab[98715]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:604
			_go_fuzz_dep_.CoverTab[98716]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:604
			// _ = "end of CoverTab[98716]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:604
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:604
		// _ = "end of CoverTab[98690]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:604
		_go_fuzz_dep_.CoverTab[98691]++

													pp.brokerProducer.input <- msg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:606
		// _ = "end of CoverTab[98691]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:607
	// _ = "end of CoverTab[98681]"
}

func (pp *partitionProducer) newHighWatermark(hwm int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:610
	_go_fuzz_dep_.CoverTab[98717]++
												Logger.Printf("producer/leader/%s/%d state change to [retrying-%d]\n", pp.topic, pp.partition, hwm)
												pp.highWatermark = hwm

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:616
	pp.retryState[pp.highWatermark].expectChaser = true
												pp.parent.inFlight.Add(1)
												pp.brokerProducer.input <- &ProducerMessage{Topic: pp.topic, Partition: pp.partition, flags: fin, retries: pp.highWatermark - 1}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:621
	Logger.Printf("producer/leader/%s/%d abandoning broker %d\n", pp.topic, pp.partition, pp.leader.ID())
												pp.parent.unrefBrokerProducer(pp.leader, pp.brokerProducer)
												pp.brokerProducer = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:623
	// _ = "end of CoverTab[98717]"
}

func (pp *partitionProducer) flushRetryBuffers() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:626
	_go_fuzz_dep_.CoverTab[98718]++
												Logger.Printf("producer/leader/%s/%d state change to [flushing-%d]\n", pp.topic, pp.partition, pp.highWatermark)
												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:628
		_go_fuzz_dep_.CoverTab[98719]++
													pp.highWatermark--

													if pp.brokerProducer == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:631
			_go_fuzz_dep_.CoverTab[98722]++
														if err := pp.updateLeader(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:632
				_go_fuzz_dep_.CoverTab[98724]++
															pp.parent.returnErrors(pp.retryState[pp.highWatermark].buf, err)
															goto flushDone
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:634
				// _ = "end of CoverTab[98724]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:635
				_go_fuzz_dep_.CoverTab[98725]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:635
				// _ = "end of CoverTab[98725]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:635
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:635
			// _ = "end of CoverTab[98722]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:635
			_go_fuzz_dep_.CoverTab[98723]++
														Logger.Printf("producer/leader/%s/%d selected broker %d\n", pp.topic, pp.partition, pp.leader.ID())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:636
			// _ = "end of CoverTab[98723]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:637
			_go_fuzz_dep_.CoverTab[98726]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:637
			// _ = "end of CoverTab[98726]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:637
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:637
		// _ = "end of CoverTab[98719]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:637
		_go_fuzz_dep_.CoverTab[98720]++

													for _, msg := range pp.retryState[pp.highWatermark].buf {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:639
			_go_fuzz_dep_.CoverTab[98727]++
														pp.brokerProducer.input <- msg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:640
			// _ = "end of CoverTab[98727]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:641
		// _ = "end of CoverTab[98720]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:641
		_go_fuzz_dep_.CoverTab[98721]++

	flushDone:
		pp.retryState[pp.highWatermark].buf = nil
		if pp.retryState[pp.highWatermark].expectChaser {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:645
			_go_fuzz_dep_.CoverTab[98728]++
														Logger.Printf("producer/leader/%s/%d state change to [retrying-%d]\n", pp.topic, pp.partition, pp.highWatermark)
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:647
			// _ = "end of CoverTab[98728]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:648
			_go_fuzz_dep_.CoverTab[98729]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:648
			if pp.highWatermark == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:648
				_go_fuzz_dep_.CoverTab[98730]++
															Logger.Printf("producer/leader/%s/%d state change to [normal]\n", pp.topic, pp.partition)
															break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:650
				// _ = "end of CoverTab[98730]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:651
				_go_fuzz_dep_.CoverTab[98731]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:651
				// _ = "end of CoverTab[98731]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:651
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:651
			// _ = "end of CoverTab[98729]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:651
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:651
		// _ = "end of CoverTab[98721]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:652
	// _ = "end of CoverTab[98718]"
}

func (pp *partitionProducer) updateLeader() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:655
	_go_fuzz_dep_.CoverTab[98732]++
												return pp.breaker.Run(func() (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:656
		_go_fuzz_dep_.CoverTab[98733]++
													if err = pp.parent.client.RefreshMetadata(pp.topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:657
			_go_fuzz_dep_.CoverTab[98736]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:658
			// _ = "end of CoverTab[98736]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:659
			_go_fuzz_dep_.CoverTab[98737]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:659
			// _ = "end of CoverTab[98737]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:659
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:659
		// _ = "end of CoverTab[98733]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:659
		_go_fuzz_dep_.CoverTab[98734]++

													if pp.leader, err = pp.parent.client.Leader(pp.topic, pp.partition); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:661
			_go_fuzz_dep_.CoverTab[98738]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:662
			// _ = "end of CoverTab[98738]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:663
			_go_fuzz_dep_.CoverTab[98739]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:663
			// _ = "end of CoverTab[98739]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:663
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:663
		// _ = "end of CoverTab[98734]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:663
		_go_fuzz_dep_.CoverTab[98735]++

													pp.brokerProducer = pp.parent.getBrokerProducer(pp.leader)
													pp.parent.inFlight.Add(1)
													pp.brokerProducer.input <- &ProducerMessage{Topic: pp.topic, Partition: pp.partition, flags: syn}

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:669
		// _ = "end of CoverTab[98735]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:670
	// _ = "end of CoverTab[98732]"
}

// one per broker; also constructs an associated flusher
func (p *asyncProducer) newBrokerProducer(broker *Broker) *brokerProducer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:674
	_go_fuzz_dep_.CoverTab[98740]++
												var (
		input		= make(chan *ProducerMessage)
		bridge		= make(chan *produceSet)
		responses	= make(chan *brokerProducerResponse)
	)

	bp := &brokerProducer{
		parent:		p,
		broker:		broker,
		input:		input,
		output:		bridge,
		responses:	responses,
		stopchan:	make(chan struct{}),
		buffer:		newProduceSet(p),
		currentRetries:	make(map[string]map[int32]error),
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:690
	_curRoutineNum122_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:690
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum122_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
		_go_fuzz_dep_.CoverTab[98743]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
			_go_fuzz_dep_.CoverTab[98744]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum122_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
			// _ = "end of CoverTab[98744]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
		withRecover(bp.run)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
		// _ = "end of CoverTab[98743]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
	_curRoutineNum123_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:691
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum123_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
	go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
		_go_fuzz_dep_.CoverTab[98745]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
			_go_fuzz_dep_.CoverTab[98746]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum123_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
			// _ = "end of CoverTab[98746]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
		withRecover(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:694
			_go_fuzz_dep_.CoverTab[98747]++
														for set := range bridge {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:695
				_go_fuzz_dep_.CoverTab[98749]++
															request := set.buildRequest()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:699
				sendResponse := func(set *produceSet) ProduceCallback {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:699
					_go_fuzz_dep_.CoverTab[98752]++
																return func(response *ProduceResponse, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:700
						_go_fuzz_dep_.CoverTab[98753]++
																	responses <- &brokerProducerResponse{
							set:	set,
							err:	err,
							res:	response,
						}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:705
						// _ = "end of CoverTab[98753]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:706
					// _ = "end of CoverTab[98752]"
				}(set)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:707
				// _ = "end of CoverTab[98749]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:707
				_go_fuzz_dep_.CoverTab[98750]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:712
				err := broker.AsyncProduce(request, sendResponse)
				if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:713
					_go_fuzz_dep_.CoverTab[98754]++

																sendResponse(nil, err)
																continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:716
					// _ = "end of CoverTab[98754]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:717
					_go_fuzz_dep_.CoverTab[98755]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:717
					// _ = "end of CoverTab[98755]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:717
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:717
				// _ = "end of CoverTab[98750]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:717
				_go_fuzz_dep_.CoverTab[98751]++

															if p.conf.Producer.RequiredAcks == NoResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:719
					_go_fuzz_dep_.CoverTab[98756]++

																sendResponse(nil, nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:721
					// _ = "end of CoverTab[98756]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:722
					_go_fuzz_dep_.CoverTab[98757]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:722
					// _ = "end of CoverTab[98757]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:722
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:722
				// _ = "end of CoverTab[98751]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:723
			// _ = "end of CoverTab[98747]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:723
			_go_fuzz_dep_.CoverTab[98748]++
														close(responses)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:724
			// _ = "end of CoverTab[98748]"
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:725
		// _ = "end of CoverTab[98745]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:725
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:725
	// _ = "end of CoverTab[98740]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:725
	_go_fuzz_dep_.CoverTab[98741]++

												if p.conf.Producer.Retry.Max <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:727
		_go_fuzz_dep_.CoverTab[98758]++
													bp.abandoned = make(chan struct{})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:728
		// _ = "end of CoverTab[98758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:729
		_go_fuzz_dep_.CoverTab[98759]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:729
		// _ = "end of CoverTab[98759]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:729
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:729
	// _ = "end of CoverTab[98741]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:729
	_go_fuzz_dep_.CoverTab[98742]++

												return bp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:731
	// _ = "end of CoverTab[98742]"
}

type brokerProducerResponse struct {
	set	*produceSet
	err	error
	res	*ProduceResponse
}

// groups messages together into appropriately-sized batches for sending to the broker
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:740
// handles state related to retries etc
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:742
type brokerProducer struct {
	parent	*asyncProducer
	broker	*Broker

	input		chan *ProducerMessage
	output		chan<- *produceSet
	responses	<-chan *brokerProducerResponse
	abandoned	chan struct{}
	stopchan	chan struct{}

	buffer		*produceSet
	timer		<-chan time.Time
	timerFired	bool

	closing		error
	currentRetries	map[string]map[int32]error
}

func (bp *brokerProducer) run() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:760
	_go_fuzz_dep_.CoverTab[98760]++
												var output chan<- *produceSet
												Logger.Printf("producer/broker/%d starting up\n", bp.broker.ID())

												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:764
		_go_fuzz_dep_.CoverTab[98761]++
													select {
		case msg, ok := <-bp.input:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:766
			_go_fuzz_dep_.CoverTab[98763]++
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:767
				_go_fuzz_dep_.CoverTab[98775]++
															Logger.Printf("producer/broker/%d input chan closed\n", bp.broker.ID())
															bp.shutdown()
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:770
				// _ = "end of CoverTab[98775]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:771
				_go_fuzz_dep_.CoverTab[98776]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:771
				// _ = "end of CoverTab[98776]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:771
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:771
			// _ = "end of CoverTab[98763]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:771
			_go_fuzz_dep_.CoverTab[98764]++

														if msg == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:773
				_go_fuzz_dep_.CoverTab[98777]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:774
				// _ = "end of CoverTab[98777]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:775
				_go_fuzz_dep_.CoverTab[98778]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:775
				// _ = "end of CoverTab[98778]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:775
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:775
			// _ = "end of CoverTab[98764]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:775
			_go_fuzz_dep_.CoverTab[98765]++

														if msg.flags&syn == syn {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:777
				_go_fuzz_dep_.CoverTab[98779]++
															Logger.Printf("producer/broker/%d state change to [open] on %s/%d\n",
					bp.broker.ID(), msg.Topic, msg.Partition)
				if bp.currentRetries[msg.Topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:780
					_go_fuzz_dep_.CoverTab[98781]++
																bp.currentRetries[msg.Topic] = make(map[int32]error)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:781
					// _ = "end of CoverTab[98781]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:782
					_go_fuzz_dep_.CoverTab[98782]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:782
					// _ = "end of CoverTab[98782]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:782
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:782
				// _ = "end of CoverTab[98779]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:782
				_go_fuzz_dep_.CoverTab[98780]++
															bp.currentRetries[msg.Topic][msg.Partition] = nil
															bp.parent.inFlight.Done()
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:785
				// _ = "end of CoverTab[98780]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:786
				_go_fuzz_dep_.CoverTab[98783]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:786
				// _ = "end of CoverTab[98783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:786
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:786
			// _ = "end of CoverTab[98765]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:786
			_go_fuzz_dep_.CoverTab[98766]++

														if reason := bp.needsRetry(msg); reason != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:788
				_go_fuzz_dep_.CoverTab[98784]++
															bp.parent.retryMessage(msg, reason)

															if bp.closing == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:791
					_go_fuzz_dep_.CoverTab[98786]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:791
					return msg.flags&fin == fin
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:791
					// _ = "end of CoverTab[98786]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:791
				}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:791
					_go_fuzz_dep_.CoverTab[98787]++

																delete(bp.currentRetries[msg.Topic], msg.Partition)
																Logger.Printf("producer/broker/%d state change to [closed] on %s/%d\n",
						bp.broker.ID(), msg.Topic, msg.Partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:795
					// _ = "end of CoverTab[98787]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:796
					_go_fuzz_dep_.CoverTab[98788]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:796
					// _ = "end of CoverTab[98788]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:796
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:796
				// _ = "end of CoverTab[98784]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:796
				_go_fuzz_dep_.CoverTab[98785]++

															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:798
				// _ = "end of CoverTab[98785]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:799
				_go_fuzz_dep_.CoverTab[98789]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:799
				// _ = "end of CoverTab[98789]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:799
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:799
			// _ = "end of CoverTab[98766]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:799
			_go_fuzz_dep_.CoverTab[98767]++

														if bp.buffer.wouldOverflow(msg) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:801
				_go_fuzz_dep_.CoverTab[98790]++
															Logger.Printf("producer/broker/%d maximum request accumulated, waiting for space\n", bp.broker.ID())
															if err := bp.waitForSpace(msg, false); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:803
					_go_fuzz_dep_.CoverTab[98791]++
																bp.parent.retryMessage(msg, err)
																continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:805
					// _ = "end of CoverTab[98791]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:806
					_go_fuzz_dep_.CoverTab[98792]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:806
					// _ = "end of CoverTab[98792]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:806
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:806
				// _ = "end of CoverTab[98790]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:807
				_go_fuzz_dep_.CoverTab[98793]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:807
				// _ = "end of CoverTab[98793]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:807
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:807
			// _ = "end of CoverTab[98767]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:807
			_go_fuzz_dep_.CoverTab[98768]++

														if bp.parent.txnmgr.producerID != noProducerID && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:809
				_go_fuzz_dep_.CoverTab[98794]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:809
				return bp.buffer.producerEpoch != msg.producerEpoch
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:809
				// _ = "end of CoverTab[98794]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:809
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:809
				_go_fuzz_dep_.CoverTab[98795]++

															Logger.Printf("producer/broker/%d detected epoch rollover, waiting for new buffer\n", bp.broker.ID())
															if err := bp.waitForSpace(msg, true); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:812
					_go_fuzz_dep_.CoverTab[98796]++
																bp.parent.retryMessage(msg, err)
																continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:814
					// _ = "end of CoverTab[98796]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:815
					_go_fuzz_dep_.CoverTab[98797]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:815
					// _ = "end of CoverTab[98797]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:815
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:815
				// _ = "end of CoverTab[98795]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:816
				_go_fuzz_dep_.CoverTab[98798]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:816
				// _ = "end of CoverTab[98798]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:816
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:816
			// _ = "end of CoverTab[98768]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:816
			_go_fuzz_dep_.CoverTab[98769]++
														if err := bp.buffer.add(msg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:817
				_go_fuzz_dep_.CoverTab[98799]++
															bp.parent.returnError(msg, err)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:819
				// _ = "end of CoverTab[98799]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:820
				_go_fuzz_dep_.CoverTab[98800]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:820
				// _ = "end of CoverTab[98800]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:820
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:820
			// _ = "end of CoverTab[98769]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:820
			_go_fuzz_dep_.CoverTab[98770]++

														if bp.parent.conf.Producer.Flush.Frequency > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:822
				_go_fuzz_dep_.CoverTab[98801]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:822
				return bp.timer == nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:822
				// _ = "end of CoverTab[98801]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:822
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:822
				_go_fuzz_dep_.CoverTab[98802]++
															bp.timer = time.After(bp.parent.conf.Producer.Flush.Frequency)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:823
				// _ = "end of CoverTab[98802]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:824
				_go_fuzz_dep_.CoverTab[98803]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:824
				// _ = "end of CoverTab[98803]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:824
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:824
			// _ = "end of CoverTab[98770]"
		case <-bp.timer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:825
			_go_fuzz_dep_.CoverTab[98771]++
														bp.timerFired = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:826
			// _ = "end of CoverTab[98771]"
		case output <- bp.buffer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:827
			_go_fuzz_dep_.CoverTab[98772]++
														bp.rollOver()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:828
			// _ = "end of CoverTab[98772]"
		case response, ok := <-bp.responses:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:829
			_go_fuzz_dep_.CoverTab[98773]++
														if ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:830
				_go_fuzz_dep_.CoverTab[98804]++
															bp.handleResponse(response)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:831
				// _ = "end of CoverTab[98804]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:832
				_go_fuzz_dep_.CoverTab[98805]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:832
				// _ = "end of CoverTab[98805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:832
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:832
			// _ = "end of CoverTab[98773]"
		case <-bp.stopchan:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:833
			_go_fuzz_dep_.CoverTab[98774]++
														Logger.Printf(
				"producer/broker/%d run loop asked to stop\n", bp.broker.ID())
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:836
			// _ = "end of CoverTab[98774]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:837
		// _ = "end of CoverTab[98761]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:837
		_go_fuzz_dep_.CoverTab[98762]++

													if bp.timerFired || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:839
			_go_fuzz_dep_.CoverTab[98806]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:839
			return bp.buffer.readyToFlush()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:839
			// _ = "end of CoverTab[98806]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:839
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:839
			_go_fuzz_dep_.CoverTab[98807]++
														output = bp.output
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:840
			// _ = "end of CoverTab[98807]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:841
			_go_fuzz_dep_.CoverTab[98808]++
														output = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:842
			// _ = "end of CoverTab[98808]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:843
		// _ = "end of CoverTab[98762]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:844
	// _ = "end of CoverTab[98760]"
}

func (bp *brokerProducer) shutdown() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:847
	_go_fuzz_dep_.CoverTab[98809]++
												for !bp.buffer.empty() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:848
		_go_fuzz_dep_.CoverTab[98812]++
													select {
		case response := <-bp.responses:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:850
			_go_fuzz_dep_.CoverTab[98813]++
														bp.handleResponse(response)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:851
			// _ = "end of CoverTab[98813]"
		case bp.output <- bp.buffer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:852
			_go_fuzz_dep_.CoverTab[98814]++
														bp.rollOver()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:853
			// _ = "end of CoverTab[98814]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:854
		// _ = "end of CoverTab[98812]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:855
	// _ = "end of CoverTab[98809]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:855
	_go_fuzz_dep_.CoverTab[98810]++
												close(bp.output)
												for response := range bp.responses {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:857
		_go_fuzz_dep_.CoverTab[98815]++
													bp.handleResponse(response)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:858
		// _ = "end of CoverTab[98815]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:859
	// _ = "end of CoverTab[98810]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:859
	_go_fuzz_dep_.CoverTab[98811]++
												close(bp.stopchan)
												Logger.Printf("producer/broker/%d shut down\n", bp.broker.ID())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:861
	// _ = "end of CoverTab[98811]"
}

func (bp *brokerProducer) needsRetry(msg *ProducerMessage) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:864
	_go_fuzz_dep_.CoverTab[98816]++
												if bp.closing != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:865
		_go_fuzz_dep_.CoverTab[98818]++
													return bp.closing
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:866
		// _ = "end of CoverTab[98818]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:867
		_go_fuzz_dep_.CoverTab[98819]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:867
		// _ = "end of CoverTab[98819]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:867
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:867
	// _ = "end of CoverTab[98816]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:867
	_go_fuzz_dep_.CoverTab[98817]++

												return bp.currentRetries[msg.Topic][msg.Partition]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:869
	// _ = "end of CoverTab[98817]"
}

func (bp *brokerProducer) waitForSpace(msg *ProducerMessage, forceRollover bool) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:872
	_go_fuzz_dep_.CoverTab[98820]++
												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:873
		_go_fuzz_dep_.CoverTab[98821]++
													select {
		case response := <-bp.responses:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:875
			_go_fuzz_dep_.CoverTab[98822]++
														bp.handleResponse(response)

														if reason := bp.needsRetry(msg); reason != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:878
				_go_fuzz_dep_.CoverTab[98824]++
															return reason
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:879
				// _ = "end of CoverTab[98824]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:880
				_go_fuzz_dep_.CoverTab[98825]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:880
				if !bp.buffer.wouldOverflow(msg) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:880
					_go_fuzz_dep_.CoverTab[98826]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:880
					return !forceRollover
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:880
					// _ = "end of CoverTab[98826]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:880
				}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:880
					_go_fuzz_dep_.CoverTab[98827]++
																return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:881
					// _ = "end of CoverTab[98827]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:882
					_go_fuzz_dep_.CoverTab[98828]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:882
					// _ = "end of CoverTab[98828]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:882
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:882
				// _ = "end of CoverTab[98825]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:882
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:882
			// _ = "end of CoverTab[98822]"
		case bp.output <- bp.buffer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:883
			_go_fuzz_dep_.CoverTab[98823]++
														bp.rollOver()
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:885
			// _ = "end of CoverTab[98823]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:886
		// _ = "end of CoverTab[98821]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:887
	// _ = "end of CoverTab[98820]"
}

func (bp *brokerProducer) rollOver() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:890
	_go_fuzz_dep_.CoverTab[98829]++
												bp.timer = nil
												bp.timerFired = false
												bp.buffer = newProduceSet(bp.parent)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:893
	// _ = "end of CoverTab[98829]"
}

func (bp *brokerProducer) handleResponse(response *brokerProducerResponse) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:896
	_go_fuzz_dep_.CoverTab[98830]++
												if response.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:897
		_go_fuzz_dep_.CoverTab[98832]++
													bp.handleError(response.set, response.err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:898
		// _ = "end of CoverTab[98832]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:899
		_go_fuzz_dep_.CoverTab[98833]++
													bp.handleSuccess(response.set, response.res)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:900
		// _ = "end of CoverTab[98833]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:901
	// _ = "end of CoverTab[98830]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:901
	_go_fuzz_dep_.CoverTab[98831]++

												if bp.buffer.empty() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:903
		_go_fuzz_dep_.CoverTab[98834]++
													bp.rollOver()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:904
		// _ = "end of CoverTab[98834]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:905
		_go_fuzz_dep_.CoverTab[98835]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:905
		// _ = "end of CoverTab[98835]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:905
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:905
	// _ = "end of CoverTab[98831]"
}

func (bp *brokerProducer) handleSuccess(sent *produceSet, response *ProduceResponse) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:908
	_go_fuzz_dep_.CoverTab[98836]++
	// we iterate through the blocks in the request set, not the response, so that we notice
	// if the response is missing a block completely
	var retryTopics []string
	sent.eachPartition(func(topic string, partition int32, pSet *partitionSet) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:912
		_go_fuzz_dep_.CoverTab[98838]++
													if response == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:913
			_go_fuzz_dep_.CoverTab[98841]++

														bp.parent.returnSuccesses(pSet.msgs)
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:916
			// _ = "end of CoverTab[98841]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:917
			_go_fuzz_dep_.CoverTab[98842]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:917
			// _ = "end of CoverTab[98842]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:917
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:917
		// _ = "end of CoverTab[98838]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:917
		_go_fuzz_dep_.CoverTab[98839]++

													block := response.GetBlock(topic, partition)
													if block == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:920
			_go_fuzz_dep_.CoverTab[98843]++
														bp.parent.returnErrors(pSet.msgs, ErrIncompleteResponse)
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:922
			// _ = "end of CoverTab[98843]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:923
			_go_fuzz_dep_.CoverTab[98844]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:923
			// _ = "end of CoverTab[98844]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:923
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:923
		// _ = "end of CoverTab[98839]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:923
		_go_fuzz_dep_.CoverTab[98840]++

													switch block.Err {

		case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:927
			_go_fuzz_dep_.CoverTab[98845]++
														if bp.parent.conf.Version.IsAtLeast(V0_10_0_0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:928
				_go_fuzz_dep_.CoverTab[98852]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:928
				return !block.Timestamp.IsZero()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:928
				// _ = "end of CoverTab[98852]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:928
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:928
				_go_fuzz_dep_.CoverTab[98853]++
															for _, msg := range pSet.msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:929
					_go_fuzz_dep_.CoverTab[98854]++
																msg.Timestamp = block.Timestamp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:930
					// _ = "end of CoverTab[98854]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:931
				// _ = "end of CoverTab[98853]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:932
				_go_fuzz_dep_.CoverTab[98855]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:932
				// _ = "end of CoverTab[98855]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:932
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:932
			// _ = "end of CoverTab[98845]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:932
			_go_fuzz_dep_.CoverTab[98846]++
														for i, msg := range pSet.msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:933
				_go_fuzz_dep_.CoverTab[98856]++
															msg.Offset = block.Offset + int64(i)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:934
				// _ = "end of CoverTab[98856]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:935
			// _ = "end of CoverTab[98846]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:935
			_go_fuzz_dep_.CoverTab[98847]++
														bp.parent.returnSuccesses(pSet.msgs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:936
			// _ = "end of CoverTab[98847]"

		case ErrDuplicateSequenceNumber:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:938
			_go_fuzz_dep_.CoverTab[98848]++
														bp.parent.returnSuccesses(pSet.msgs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:939
			// _ = "end of CoverTab[98848]"

		case ErrInvalidMessage, ErrUnknownTopicOrPartition, ErrLeaderNotAvailable, ErrNotLeaderForPartition,
			ErrRequestTimedOut, ErrNotEnoughReplicas, ErrNotEnoughReplicasAfterAppend:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:942
			_go_fuzz_dep_.CoverTab[98849]++
														if bp.parent.conf.Producer.Retry.Max <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:943
				_go_fuzz_dep_.CoverTab[98857]++
															bp.parent.abandonBrokerConnection(bp.broker)
															bp.parent.returnErrors(pSet.msgs, block.Err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:945
				// _ = "end of CoverTab[98857]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:946
				_go_fuzz_dep_.CoverTab[98858]++
															retryTopics = append(retryTopics, topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:947
				// _ = "end of CoverTab[98858]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:948
			// _ = "end of CoverTab[98849]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:950
			_go_fuzz_dep_.CoverTab[98850]++
														if bp.parent.conf.Producer.Retry.Max <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:951
				_go_fuzz_dep_.CoverTab[98859]++
															bp.parent.abandonBrokerConnection(bp.broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:952
				// _ = "end of CoverTab[98859]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:953
				_go_fuzz_dep_.CoverTab[98860]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:953
				// _ = "end of CoverTab[98860]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:953
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:953
			// _ = "end of CoverTab[98850]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:953
			_go_fuzz_dep_.CoverTab[98851]++
														bp.parent.returnErrors(pSet.msgs, block.Err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:954
			// _ = "end of CoverTab[98851]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:955
		// _ = "end of CoverTab[98840]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:956
	// _ = "end of CoverTab[98836]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:956
	_go_fuzz_dep_.CoverTab[98837]++

												if len(retryTopics) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:958
		_go_fuzz_dep_.CoverTab[98861]++
													if bp.parent.conf.Producer.Idempotent {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:959
			_go_fuzz_dep_.CoverTab[98863]++
														err := bp.parent.client.RefreshMetadata(retryTopics...)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:961
				_go_fuzz_dep_.CoverTab[98864]++
															Logger.Printf("Failed refreshing metadata because of %v\n", err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:962
				// _ = "end of CoverTab[98864]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:963
				_go_fuzz_dep_.CoverTab[98865]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:963
				// _ = "end of CoverTab[98865]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:963
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:963
			// _ = "end of CoverTab[98863]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:964
			_go_fuzz_dep_.CoverTab[98866]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:964
			// _ = "end of CoverTab[98866]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:964
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:964
		// _ = "end of CoverTab[98861]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:964
		_go_fuzz_dep_.CoverTab[98862]++

													sent.eachPartition(func(topic string, partition int32, pSet *partitionSet) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:966
			_go_fuzz_dep_.CoverTab[98867]++
														block := response.GetBlock(topic, partition)
														if block == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:968
				_go_fuzz_dep_.CoverTab[98869]++

															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:970
				// _ = "end of CoverTab[98869]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:971
				_go_fuzz_dep_.CoverTab[98870]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:971
				// _ = "end of CoverTab[98870]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:971
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:971
			// _ = "end of CoverTab[98867]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:971
			_go_fuzz_dep_.CoverTab[98868]++

														switch block.Err {
			case ErrInvalidMessage, ErrUnknownTopicOrPartition, ErrLeaderNotAvailable, ErrNotLeaderForPartition,
				ErrRequestTimedOut, ErrNotEnoughReplicas, ErrNotEnoughReplicasAfterAppend:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:975
				_go_fuzz_dep_.CoverTab[98871]++
															Logger.Printf("producer/broker/%d state change to [retrying] on %s/%d because %v\n",
					bp.broker.ID(), topic, partition, block.Err)
				if bp.currentRetries[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:978
					_go_fuzz_dep_.CoverTab[98875]++
																bp.currentRetries[topic] = make(map[int32]error)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:979
					// _ = "end of CoverTab[98875]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:980
					_go_fuzz_dep_.CoverTab[98876]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:980
					// _ = "end of CoverTab[98876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:980
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:980
				// _ = "end of CoverTab[98871]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:980
				_go_fuzz_dep_.CoverTab[98872]++
															bp.currentRetries[topic][partition] = block.Err
															if bp.parent.conf.Producer.Idempotent {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:982
					_go_fuzz_dep_.CoverTab[98877]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:982
					_curRoutineNum124_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:982
					_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum124_)
																go bp.parent.retryBatch(topic, partition, pSet, block.Err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:983
					// _ = "end of CoverTab[98877]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:984
					_go_fuzz_dep_.CoverTab[98878]++
																bp.parent.retryMessages(pSet.msgs, block.Err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:985
					// _ = "end of CoverTab[98878]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:986
				// _ = "end of CoverTab[98872]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:986
				_go_fuzz_dep_.CoverTab[98873]++

															bp.parent.retryMessages(bp.buffer.dropPartition(topic, partition), block.Err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:988
				// _ = "end of CoverTab[98873]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:988
			default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:988
				_go_fuzz_dep_.CoverTab[98874]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:988
				// _ = "end of CoverTab[98874]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:989
			// _ = "end of CoverTab[98868]"
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:990
		// _ = "end of CoverTab[98862]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:991
		_go_fuzz_dep_.CoverTab[98879]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:991
		// _ = "end of CoverTab[98879]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:991
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:991
	// _ = "end of CoverTab[98837]"
}

func (p *asyncProducer) retryBatch(topic string, partition int32, pSet *partitionSet, kerr KError) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:994
	_go_fuzz_dep_.CoverTab[98880]++
												Logger.Printf("Retrying batch for %v-%d because of %s\n", topic, partition, kerr)
												produceSet := newProduceSet(p)
												produceSet.msgs[topic] = make(map[int32]*partitionSet)
												produceSet.msgs[topic][partition] = pSet
												produceSet.bufferBytes += pSet.bufferBytes
												produceSet.bufferCount += len(pSet.msgs)
												for _, msg := range pSet.msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1001
		_go_fuzz_dep_.CoverTab[98883]++
													if msg.retries >= p.conf.Producer.Retry.Max {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1002
			_go_fuzz_dep_.CoverTab[98885]++
														p.returnError(msg, kerr)
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1004
			// _ = "end of CoverTab[98885]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1005
			_go_fuzz_dep_.CoverTab[98886]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1005
			// _ = "end of CoverTab[98886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1005
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1005
		// _ = "end of CoverTab[98883]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1005
		_go_fuzz_dep_.CoverTab[98884]++
													msg.retries++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1006
		// _ = "end of CoverTab[98884]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1007
	// _ = "end of CoverTab[98880]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1007
	_go_fuzz_dep_.CoverTab[98881]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1010
	leader, err := p.client.Leader(topic, partition)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1011
		_go_fuzz_dep_.CoverTab[98887]++
													Logger.Printf("Failed retrying batch for %v-%d because of %v while looking up for new leader\n", topic, partition, err)
													for _, msg := range pSet.msgs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1013
			_go_fuzz_dep_.CoverTab[98889]++
														p.returnError(msg, kerr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1014
			// _ = "end of CoverTab[98889]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1015
		// _ = "end of CoverTab[98887]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1015
		_go_fuzz_dep_.CoverTab[98888]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1016
		// _ = "end of CoverTab[98888]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1017
		_go_fuzz_dep_.CoverTab[98890]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1017
		// _ = "end of CoverTab[98890]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1017
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1017
	// _ = "end of CoverTab[98881]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1017
	_go_fuzz_dep_.CoverTab[98882]++
												bp := p.getBrokerProducer(leader)
												bp.output <- produceSet
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1019
	// _ = "end of CoverTab[98882]"
}

func (bp *brokerProducer) handleError(sent *produceSet, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1022
	_go_fuzz_dep_.CoverTab[98891]++
												switch err.(type) {
	case PacketEncodingError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1024
		_go_fuzz_dep_.CoverTab[98892]++
													sent.eachPartition(func(topic string, partition int32, pSet *partitionSet) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1025
			_go_fuzz_dep_.CoverTab[98896]++
														bp.parent.returnErrors(pSet.msgs, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1026
			// _ = "end of CoverTab[98896]"
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1027
		// _ = "end of CoverTab[98892]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1028
		_go_fuzz_dep_.CoverTab[98893]++
													Logger.Printf("producer/broker/%d state change to [closing] because %s\n", bp.broker.ID(), err)
													bp.parent.abandonBrokerConnection(bp.broker)
													_ = bp.broker.Close()
													bp.closing = err
													sent.eachPartition(func(topic string, partition int32, pSet *partitionSet) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1033
			_go_fuzz_dep_.CoverTab[98897]++
														bp.parent.retryMessages(pSet.msgs, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1034
			// _ = "end of CoverTab[98897]"
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1035
		// _ = "end of CoverTab[98893]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1035
		_go_fuzz_dep_.CoverTab[98894]++
													bp.buffer.eachPartition(func(topic string, partition int32, pSet *partitionSet) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1036
			_go_fuzz_dep_.CoverTab[98898]++
														bp.parent.retryMessages(pSet.msgs, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1037
			// _ = "end of CoverTab[98898]"
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1038
		// _ = "end of CoverTab[98894]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1038
		_go_fuzz_dep_.CoverTab[98895]++
													bp.rollOver()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1039
		// _ = "end of CoverTab[98895]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1040
	// _ = "end of CoverTab[98891]"
}

// singleton
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1043
// effectively a "bridge" between the flushers and the dispatcher in order to avoid deadlock
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1043
// based on https://godoc.org/github.com/eapache/channels#InfiniteChannel
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1046
func (p *asyncProducer) retryHandler() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1046
	_go_fuzz_dep_.CoverTab[98899]++
												var msg *ProducerMessage
												buf := queue.New()

												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1050
		_go_fuzz_dep_.CoverTab[98900]++
													if buf.Length() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1051
			_go_fuzz_dep_.CoverTab[98903]++
														msg = <-p.retries
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1052
			// _ = "end of CoverTab[98903]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1053
			_go_fuzz_dep_.CoverTab[98904]++
														select {
			case msg = <-p.retries:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1055
				_go_fuzz_dep_.CoverTab[98905]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1055
				// _ = "end of CoverTab[98905]"
			case p.input <- buf.Peek().(*ProducerMessage):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1056
				_go_fuzz_dep_.CoverTab[98906]++
															buf.Remove()
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1058
				// _ = "end of CoverTab[98906]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1059
			// _ = "end of CoverTab[98904]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1060
		// _ = "end of CoverTab[98900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1060
		_go_fuzz_dep_.CoverTab[98901]++

													if msg == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1062
			_go_fuzz_dep_.CoverTab[98907]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1063
			// _ = "end of CoverTab[98907]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1064
			_go_fuzz_dep_.CoverTab[98908]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1064
			// _ = "end of CoverTab[98908]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1064
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1064
		// _ = "end of CoverTab[98901]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1064
		_go_fuzz_dep_.CoverTab[98902]++

													buf.Add(msg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1066
		// _ = "end of CoverTab[98902]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1067
	// _ = "end of CoverTab[98899]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1072
func (p *asyncProducer) shutdown() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1072
	_go_fuzz_dep_.CoverTab[98909]++
												Logger.Println("Producer shutting down.")
												p.inFlight.Add(1)
												p.input <- &ProducerMessage{flags: shutdown}

												p.inFlight.Wait()

												err := p.client.Close()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1080
		_go_fuzz_dep_.CoverTab[98911]++
													Logger.Println("producer/shutdown failed to close the embedded client:", err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1081
		// _ = "end of CoverTab[98911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1082
		_go_fuzz_dep_.CoverTab[98912]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1082
		// _ = "end of CoverTab[98912]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1082
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1082
	// _ = "end of CoverTab[98909]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1082
	_go_fuzz_dep_.CoverTab[98910]++

												close(p.input)
												close(p.retries)
												close(p.errors)
												close(p.successes)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1087
	// _ = "end of CoverTab[98910]"
}

func (p *asyncProducer) returnError(msg *ProducerMessage, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1090
	_go_fuzz_dep_.CoverTab[98913]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1093
	if msg.hasSequence {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1093
		_go_fuzz_dep_.CoverTab[98916]++
													Logger.Printf("producer/txnmanager rolling over epoch due to publish failure on %s/%d", msg.Topic, msg.Partition)
													p.txnmgr.bumpEpoch()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1095
		// _ = "end of CoverTab[98916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1096
		_go_fuzz_dep_.CoverTab[98917]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1096
		// _ = "end of CoverTab[98917]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1096
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1096
	// _ = "end of CoverTab[98913]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1096
	_go_fuzz_dep_.CoverTab[98914]++
												msg.clear()
												pErr := &ProducerError{Msg: msg, Err: err}
												if p.conf.Producer.Return.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1099
		_go_fuzz_dep_.CoverTab[98918]++
													p.errors <- pErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1100
		// _ = "end of CoverTab[98918]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1101
		_go_fuzz_dep_.CoverTab[98919]++
													Logger.Println(pErr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1102
		// _ = "end of CoverTab[98919]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1103
	// _ = "end of CoverTab[98914]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1103
	_go_fuzz_dep_.CoverTab[98915]++
												p.inFlight.Done()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1104
	// _ = "end of CoverTab[98915]"
}

func (p *asyncProducer) returnErrors(batch []*ProducerMessage, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1107
	_go_fuzz_dep_.CoverTab[98920]++
												for _, msg := range batch {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1108
		_go_fuzz_dep_.CoverTab[98921]++
													p.returnError(msg, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1109
		// _ = "end of CoverTab[98921]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1110
	// _ = "end of CoverTab[98920]"
}

func (p *asyncProducer) returnSuccesses(batch []*ProducerMessage) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1113
	_go_fuzz_dep_.CoverTab[98922]++
												for _, msg := range batch {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1114
		_go_fuzz_dep_.CoverTab[98923]++
													if p.conf.Producer.Return.Successes {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1115
			_go_fuzz_dep_.CoverTab[98925]++
														msg.clear()
														p.successes <- msg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1117
			// _ = "end of CoverTab[98925]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1118
			_go_fuzz_dep_.CoverTab[98926]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1118
			// _ = "end of CoverTab[98926]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1118
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1118
		// _ = "end of CoverTab[98923]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1118
		_go_fuzz_dep_.CoverTab[98924]++
													p.inFlight.Done()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1119
		// _ = "end of CoverTab[98924]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1120
	// _ = "end of CoverTab[98922]"
}

func (p *asyncProducer) retryMessage(msg *ProducerMessage, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1123
	_go_fuzz_dep_.CoverTab[98927]++
												if msg.retries >= p.conf.Producer.Retry.Max {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1124
		_go_fuzz_dep_.CoverTab[98928]++
													p.returnError(msg, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1125
		// _ = "end of CoverTab[98928]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1126
		_go_fuzz_dep_.CoverTab[98929]++
													msg.retries++
													p.retries <- msg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1128
		// _ = "end of CoverTab[98929]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1129
	// _ = "end of CoverTab[98927]"
}

func (p *asyncProducer) retryMessages(batch []*ProducerMessage, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1132
	_go_fuzz_dep_.CoverTab[98930]++
												for _, msg := range batch {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1133
		_go_fuzz_dep_.CoverTab[98931]++
													p.retryMessage(msg, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1134
		// _ = "end of CoverTab[98931]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1135
	// _ = "end of CoverTab[98930]"
}

func (p *asyncProducer) getBrokerProducer(broker *Broker) *brokerProducer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1138
	_go_fuzz_dep_.CoverTab[98932]++
												p.brokerLock.Lock()
												defer p.brokerLock.Unlock()

												bp := p.brokers[broker]

												if bp == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1144
		_go_fuzz_dep_.CoverTab[98934]++
													bp = p.newBrokerProducer(broker)
													p.brokers[broker] = bp
													p.brokerRefs[bp] = 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1147
		// _ = "end of CoverTab[98934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1148
		_go_fuzz_dep_.CoverTab[98935]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1148
		// _ = "end of CoverTab[98935]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1148
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1148
	// _ = "end of CoverTab[98932]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1148
	_go_fuzz_dep_.CoverTab[98933]++

												p.brokerRefs[bp]++

												return bp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1152
	// _ = "end of CoverTab[98933]"
}

func (p *asyncProducer) unrefBrokerProducer(broker *Broker, bp *brokerProducer) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1155
	_go_fuzz_dep_.CoverTab[98936]++
												p.brokerLock.Lock()
												defer p.brokerLock.Unlock()

												p.brokerRefs[bp]--
												if p.brokerRefs[bp] == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1160
		_go_fuzz_dep_.CoverTab[98937]++
													close(bp.input)
													delete(p.brokerRefs, bp)

													if p.brokers[broker] == bp {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1164
			_go_fuzz_dep_.CoverTab[98938]++
														delete(p.brokers, broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1165
			// _ = "end of CoverTab[98938]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1166
			_go_fuzz_dep_.CoverTab[98939]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1166
			// _ = "end of CoverTab[98939]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1166
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1166
		// _ = "end of CoverTab[98937]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1167
		_go_fuzz_dep_.CoverTab[98940]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1167
		// _ = "end of CoverTab[98940]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1167
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1167
	// _ = "end of CoverTab[98936]"
}

func (p *asyncProducer) abandonBrokerConnection(broker *Broker) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1170
	_go_fuzz_dep_.CoverTab[98941]++
												p.brokerLock.Lock()
												defer p.brokerLock.Unlock()

												bc, ok := p.brokers[broker]
												if ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1175
		_go_fuzz_dep_.CoverTab[98943]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1175
		return bc.abandoned != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1175
		// _ = "end of CoverTab[98943]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1175
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1175
		_go_fuzz_dep_.CoverTab[98944]++
													close(bc.abandoned)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1176
		// _ = "end of CoverTab[98944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1177
		_go_fuzz_dep_.CoverTab[98945]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1177
		// _ = "end of CoverTab[98945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1177
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1177
	// _ = "end of CoverTab[98941]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1177
	_go_fuzz_dep_.CoverTab[98942]++

												delete(p.brokers, broker)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1179
	// _ = "end of CoverTab[98942]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1180
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/async_producer.go:1180
var _ = _go_fuzz_dep_.CoverTab
