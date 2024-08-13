//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:1
)

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/rcrowley/go-metrics"
)

// ErrClosedConsumerGroup is the error returned when a method is called on a consumer group that has been closed.
var ErrClosedConsumerGroup = errors.New("kafka: tried to use a consumer group that was closed")

// ConsumerGroup is responsible for dividing up processing of topics and partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:17
// over a collection of processes (the members of the consumer group).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:19
type ConsumerGroup interface {
	// Consume joins a cluster of consumers for a given list of topics and
	// starts a blocking ConsumerGroupSession through the ConsumerGroupHandler.
	//
	// The life-cycle of a session is represented by the following steps:
	//
	// 1. The consumers join the group (as explained in https://kafka.apache.org/documentation/#intro_consumers)
	//    and is assigned their "fair share" of partitions, aka 'claims'.
	// 2. Before processing starts, the handler's Setup() hook is called to notify the user
	//    of the claims and allow any necessary preparation or alteration of state.
	// 3. For each of the assigned claims the handler's ConsumeClaim() function is then called
	//    in a separate goroutine which requires it to be thread-safe. Any state must be carefully protected
	//    from concurrent reads/writes.
	// 4. The session will persist until one of the ConsumeClaim() functions exits. This can be either when the
	//    parent context is canceled or when a server-side rebalance cycle is initiated.
	// 5. Once all the ConsumeClaim() loops have exited, the handler's Cleanup() hook is called
	//    to allow the user to perform any final tasks before a rebalance.
	// 6. Finally, marked offsets are committed one last time before claims are released.
	//
	// Please note, that once a rebalance is triggered, sessions must be completed within
	// Config.Consumer.Group.Rebalance.Timeout. This means that ConsumeClaim() functions must exit
	// as quickly as possible to allow time for Cleanup() and the final offset commit. If the timeout
	// is exceeded, the consumer will be removed from the group by Kafka, which will cause offset
	// commit failures.
	// This method should be called inside an infinite loop, when a
	// server-side rebalance happens, the consumer session will need to be
	// recreated to get the new claims.
	Consume(ctx context.Context, topics []string, handler ConsumerGroupHandler) error

	// Errors returns a read channel of errors that occurred during the consumer life-cycle.
	// By default, errors are logged and not returned over this channel.
	// If you want to implement any custom error handling, set your config's
	// Consumer.Return.Errors setting to true, and read from this channel.
	Errors() <-chan error

	// Close stops the ConsumerGroup and detaches any running sessions. It is required to call
	// this function before the object passes out of scope, as it will otherwise leak memory.
	Close() error

	// Pause suspends fetching from the requested partitions. Future calls to the broker will not return any
	// records from these partitions until they have been resumed using Resume()/ResumeAll().
	// Note that this method does not affect partition subscription.
	// In particular, it does not cause a group rebalance when automatic assignment is used.
	Pause(partitions map[string][]int32)

	// Resume resumes specified partitions which have been paused with Pause()/PauseAll().
	// New calls to the broker will return records from these partitions if there are any to be fetched.
	Resume(partitions map[string][]int32)

	// Pause suspends fetching from all partitions. Future calls to the broker will not return any
	// records from these partitions until they have been resumed using Resume()/ResumeAll().
	// Note that this method does not affect partition subscription.
	// In particular, it does not cause a group rebalance when automatic assignment is used.
	PauseAll()

	// Resume resumes all partitions which have been paused with Pause()/PauseAll().
	// New calls to the broker will return records from these partitions if there are any to be fetched.
	ResumeAll()
}

type consumerGroup struct {
	client	Client

	config		*Config
	consumer	Consumer
	groupID		string
	memberID	string
	errors		chan error

	lock		sync.Mutex
	closed		chan none
	closeOnce	sync.Once

	userData	[]byte
}

// NewConsumerGroup creates a new consumer group the given broker addresses and configuration.
func NewConsumerGroup(addrs []string, groupID string, config *Config) (ConsumerGroup, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:96
	_go_fuzz_dep_.CoverTab[100959]++
												client, err := NewClient(addrs, config)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:98
		_go_fuzz_dep_.CoverTab[100962]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:99
		// _ = "end of CoverTab[100962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:100
		_go_fuzz_dep_.CoverTab[100963]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:100
		// _ = "end of CoverTab[100963]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:100
	// _ = "end of CoverTab[100959]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:100
	_go_fuzz_dep_.CoverTab[100960]++

												c, err := newConsumerGroup(groupID, client)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:103
		_go_fuzz_dep_.CoverTab[100964]++
													_ = client.Close()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:104
		// _ = "end of CoverTab[100964]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:105
		_go_fuzz_dep_.CoverTab[100965]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:105
		// _ = "end of CoverTab[100965]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:105
	// _ = "end of CoverTab[100960]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:105
	_go_fuzz_dep_.CoverTab[100961]++
												return c, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:106
	// _ = "end of CoverTab[100961]"
}

// NewConsumerGroupFromClient creates a new consumer group using the given client. It is still
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:109
// necessary to call Close() on the underlying client when shutting down this consumer.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:109
// PLEASE NOTE: consumer groups can only re-use but not share clients.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:112
func NewConsumerGroupFromClient(groupID string, client Client) (ConsumerGroup, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:112
	_go_fuzz_dep_.CoverTab[100966]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:115
	cli := &nopCloserClient{client}
												return newConsumerGroup(groupID, cli)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:116
	// _ = "end of CoverTab[100966]"
}

func newConsumerGroup(groupID string, client Client) (ConsumerGroup, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:119
	_go_fuzz_dep_.CoverTab[100967]++
												config := client.Config()
												if !config.Version.IsAtLeast(V0_10_2_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:121
		_go_fuzz_dep_.CoverTab[100970]++
													return nil, ConfigurationError("consumer groups require Version to be >= V0_10_2_0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:122
		// _ = "end of CoverTab[100970]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:123
		_go_fuzz_dep_.CoverTab[100971]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:123
		// _ = "end of CoverTab[100971]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:123
	// _ = "end of CoverTab[100967]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:123
	_go_fuzz_dep_.CoverTab[100968]++

												consumer, err := NewConsumerFromClient(client)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:126
		_go_fuzz_dep_.CoverTab[100972]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:127
		// _ = "end of CoverTab[100972]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:128
		_go_fuzz_dep_.CoverTab[100973]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:128
		// _ = "end of CoverTab[100973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:128
	// _ = "end of CoverTab[100968]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:128
	_go_fuzz_dep_.CoverTab[100969]++

												return &consumerGroup{
		client:		client,
		consumer:	consumer,
		config:		config,
		groupID:	groupID,
		errors:		make(chan error, config.ChannelBufferSize),
		closed:		make(chan none),
		userData:	config.Consumer.Group.Member.UserData,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:138
	// _ = "end of CoverTab[100969]"
}

// Errors implements ConsumerGroup.
func (c *consumerGroup) Errors() <-chan error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:142
	_go_fuzz_dep_.CoverTab[100974]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:142
	return c.errors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:142
	// _ = "end of CoverTab[100974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:142
}

// Close implements ConsumerGroup.
func (c *consumerGroup) Close() (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:145
	_go_fuzz_dep_.CoverTab[100975]++
												c.closeOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:146
		_go_fuzz_dep_.CoverTab[100977]++
													close(c.closed)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:150
		if e := c.leave(); e != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:150
			_go_fuzz_dep_.CoverTab[100981]++
														err = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:151
			// _ = "end of CoverTab[100981]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:152
			_go_fuzz_dep_.CoverTab[100982]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:152
			// _ = "end of CoverTab[100982]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:152
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:152
		// _ = "end of CoverTab[100977]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:152
		_go_fuzz_dep_.CoverTab[100978]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:152
		_curRoutineNum132_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:152
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum132_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:155
		go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:155
			_go_fuzz_dep_.CoverTab[100983]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:155
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:155
				_go_fuzz_dep_.CoverTab[100984]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:155
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum132_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:155
				// _ = "end of CoverTab[100984]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:155
			}()
														close(c.errors)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:156
			// _ = "end of CoverTab[100983]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:157
		// _ = "end of CoverTab[100978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:157
		_go_fuzz_dep_.CoverTab[100979]++
													for e := range c.errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:158
			_go_fuzz_dep_.CoverTab[100985]++
														err = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:159
			// _ = "end of CoverTab[100985]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:160
		// _ = "end of CoverTab[100979]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:160
		_go_fuzz_dep_.CoverTab[100980]++

													if e := c.client.Close(); e != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:162
			_go_fuzz_dep_.CoverTab[100986]++
														err = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:163
			// _ = "end of CoverTab[100986]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:164
			_go_fuzz_dep_.CoverTab[100987]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:164
			// _ = "end of CoverTab[100987]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:164
		// _ = "end of CoverTab[100980]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:165
	// _ = "end of CoverTab[100975]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:165
	_go_fuzz_dep_.CoverTab[100976]++
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:166
	// _ = "end of CoverTab[100976]"
}

// Consume implements ConsumerGroup.
func (c *consumerGroup) Consume(ctx context.Context, topics []string, handler ConsumerGroupHandler) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:170
	_go_fuzz_dep_.CoverTab[100988]++

												select {
	case <-c.closed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:173
		_go_fuzz_dep_.CoverTab[100993]++
													return ErrClosedConsumerGroup
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:174
		// _ = "end of CoverTab[100993]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:175
		_go_fuzz_dep_.CoverTab[100994]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:175
		// _ = "end of CoverTab[100994]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:176
	// _ = "end of CoverTab[100988]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:176
	_go_fuzz_dep_.CoverTab[100989]++

												c.lock.Lock()
												defer c.lock.Unlock()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:182
	if len(topics) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:182
		_go_fuzz_dep_.CoverTab[100995]++
													return fmt.Errorf("no topics provided")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:183
		// _ = "end of CoverTab[100995]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:184
		_go_fuzz_dep_.CoverTab[100996]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:184
		// _ = "end of CoverTab[100996]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:184
	// _ = "end of CoverTab[100989]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:184
	_go_fuzz_dep_.CoverTab[100990]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:187
	if err := c.client.RefreshMetadata(topics...); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:187
		_go_fuzz_dep_.CoverTab[100997]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:188
		// _ = "end of CoverTab[100997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:189
		_go_fuzz_dep_.CoverTab[100998]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:189
		// _ = "end of CoverTab[100998]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:189
	// _ = "end of CoverTab[100990]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:189
	_go_fuzz_dep_.CoverTab[100991]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:192
	sess, err := c.newSession(ctx, topics, handler, c.config.Consumer.Group.Rebalance.Retry.Max)
	if err == ErrClosedClient {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:193
		_go_fuzz_dep_.CoverTab[100999]++
													return ErrClosedConsumerGroup
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:194
		// _ = "end of CoverTab[100999]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:195
		_go_fuzz_dep_.CoverTab[101000]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:195
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:195
			_go_fuzz_dep_.CoverTab[101001]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:196
			// _ = "end of CoverTab[101001]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
			_go_fuzz_dep_.CoverTab[101002]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
			// _ = "end of CoverTab[101002]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
		// _ = "end of CoverTab[101000]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
	// _ = "end of CoverTab[100991]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
	_go_fuzz_dep_.CoverTab[100992]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
	_curRoutineNum133_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:197
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum133_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:202
	go c.loopCheckPartitionNumbers(topics, sess)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:205
	<-sess.ctx.Done()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:208
	return sess.release(true)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:208
	// _ = "end of CoverTab[100992]"
}

// Pause implements ConsumerGroup.
func (c *consumerGroup) Pause(partitions map[string][]int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:212
	_go_fuzz_dep_.CoverTab[101003]++
												c.consumer.Pause(partitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:213
	// _ = "end of CoverTab[101003]"
}

// Resume implements ConsumerGroup.
func (c *consumerGroup) Resume(partitions map[string][]int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:217
	_go_fuzz_dep_.CoverTab[101004]++
												c.consumer.Resume(partitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:218
	// _ = "end of CoverTab[101004]"
}

// PauseAll implements ConsumerGroup.
func (c *consumerGroup) PauseAll() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:222
	_go_fuzz_dep_.CoverTab[101005]++
												c.consumer.PauseAll()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:223
	// _ = "end of CoverTab[101005]"
}

// ResumeAll implements ConsumerGroup.
func (c *consumerGroup) ResumeAll() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:227
	_go_fuzz_dep_.CoverTab[101006]++
												c.consumer.ResumeAll()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:228
	// _ = "end of CoverTab[101006]"
}

func (c *consumerGroup) retryNewSession(ctx context.Context, topics []string, handler ConsumerGroupHandler, retries int, refreshCoordinator bool) (*consumerGroupSession, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:231
	_go_fuzz_dep_.CoverTab[101007]++
												select {
	case <-c.closed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:233
		_go_fuzz_dep_.CoverTab[101010]++
													return nil, ErrClosedConsumerGroup
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:234
		// _ = "end of CoverTab[101010]"
	case <-time.After(c.config.Consumer.Group.Rebalance.Retry.Backoff):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:235
		_go_fuzz_dep_.CoverTab[101011]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:235
		// _ = "end of CoverTab[101011]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:236
	// _ = "end of CoverTab[101007]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:236
	_go_fuzz_dep_.CoverTab[101008]++

												if refreshCoordinator {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:238
		_go_fuzz_dep_.CoverTab[101012]++
													err := c.client.RefreshCoordinator(c.groupID)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:240
			_go_fuzz_dep_.CoverTab[101013]++
														return c.retryNewSession(ctx, topics, handler, retries, true)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:241
			// _ = "end of CoverTab[101013]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:242
			_go_fuzz_dep_.CoverTab[101014]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:242
			// _ = "end of CoverTab[101014]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:242
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:242
		// _ = "end of CoverTab[101012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:243
		_go_fuzz_dep_.CoverTab[101015]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:243
		// _ = "end of CoverTab[101015]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:243
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:243
	// _ = "end of CoverTab[101008]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:243
	_go_fuzz_dep_.CoverTab[101009]++

												return c.newSession(ctx, topics, handler, retries-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:245
	// _ = "end of CoverTab[101009]"
}

func (c *consumerGroup) newSession(ctx context.Context, topics []string, handler ConsumerGroupHandler, retries int) (*consumerGroupSession, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:248
	_go_fuzz_dep_.CoverTab[101016]++
												coordinator, err := c.client.Coordinator(c.groupID)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:250
		_go_fuzz_dep_.CoverTab[101029]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:251
			_go_fuzz_dep_.CoverTab[101031]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:252
			// _ = "end of CoverTab[101031]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:253
			_go_fuzz_dep_.CoverTab[101032]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:253
			// _ = "end of CoverTab[101032]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:253
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:253
		// _ = "end of CoverTab[101029]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:253
		_go_fuzz_dep_.CoverTab[101030]++

													return c.retryNewSession(ctx, topics, handler, retries, true)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:255
		// _ = "end of CoverTab[101030]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:256
		_go_fuzz_dep_.CoverTab[101033]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:256
		// _ = "end of CoverTab[101033]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:256
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:256
	// _ = "end of CoverTab[101016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:256
	_go_fuzz_dep_.CoverTab[101017]++

												var (
		metricRegistry		= c.config.MetricRegistry
		consumerGroupJoinTotal	metrics.Counter
		consumerGroupJoinFailed	metrics.Counter
		consumerGroupSyncTotal	metrics.Counter
		consumerGroupSyncFailed	metrics.Counter
	)

	if metricRegistry != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:266
		_go_fuzz_dep_.CoverTab[101034]++
													consumerGroupJoinTotal = metrics.GetOrRegisterCounter(fmt.Sprintf("consumer-group-join-total-%s", c.groupID), metricRegistry)
													consumerGroupJoinFailed = metrics.GetOrRegisterCounter(fmt.Sprintf("consumer-group-join-failed-%s", c.groupID), metricRegistry)
													consumerGroupSyncTotal = metrics.GetOrRegisterCounter(fmt.Sprintf("consumer-group-sync-total-%s", c.groupID), metricRegistry)
													consumerGroupSyncFailed = metrics.GetOrRegisterCounter(fmt.Sprintf("consumer-group-sync-failed-%s", c.groupID), metricRegistry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:270
		// _ = "end of CoverTab[101034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:271
		_go_fuzz_dep_.CoverTab[101035]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:271
		// _ = "end of CoverTab[101035]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:271
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:271
	// _ = "end of CoverTab[101017]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:271
	_go_fuzz_dep_.CoverTab[101018]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:274
	join, err := c.joinGroupRequest(coordinator, topics)
	if consumerGroupJoinTotal != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:275
		_go_fuzz_dep_.CoverTab[101036]++
													consumerGroupJoinTotal.Inc(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:276
		// _ = "end of CoverTab[101036]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:277
		_go_fuzz_dep_.CoverTab[101037]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:277
		// _ = "end of CoverTab[101037]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:277
	// _ = "end of CoverTab[101018]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:277
	_go_fuzz_dep_.CoverTab[101019]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:278
		_go_fuzz_dep_.CoverTab[101038]++
													_ = coordinator.Close()
													if consumerGroupJoinFailed != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:280
			_go_fuzz_dep_.CoverTab[101040]++
														consumerGroupJoinFailed.Inc(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:281
			// _ = "end of CoverTab[101040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:282
			_go_fuzz_dep_.CoverTab[101041]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:282
			// _ = "end of CoverTab[101041]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:282
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:282
		// _ = "end of CoverTab[101038]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:282
		_go_fuzz_dep_.CoverTab[101039]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:283
		// _ = "end of CoverTab[101039]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:284
		_go_fuzz_dep_.CoverTab[101042]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:284
		// _ = "end of CoverTab[101042]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:284
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:284
	// _ = "end of CoverTab[101019]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:284
	_go_fuzz_dep_.CoverTab[101020]++
												if join.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:285
		_go_fuzz_dep_.CoverTab[101043]++
													if consumerGroupJoinFailed != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:286
			_go_fuzz_dep_.CoverTab[101044]++
														consumerGroupJoinFailed.Inc(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:287
			// _ = "end of CoverTab[101044]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:288
			_go_fuzz_dep_.CoverTab[101045]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:288
			// _ = "end of CoverTab[101045]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:288
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:288
		// _ = "end of CoverTab[101043]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:289
		_go_fuzz_dep_.CoverTab[101046]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:289
		// _ = "end of CoverTab[101046]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:289
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:289
	// _ = "end of CoverTab[101020]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:289
	_go_fuzz_dep_.CoverTab[101021]++
												switch join.Err {
	case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:291
		_go_fuzz_dep_.CoverTab[101047]++
													c.memberID = join.MemberId
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:292
		// _ = "end of CoverTab[101047]"
	case ErrUnknownMemberId, ErrIllegalGeneration:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:293
		_go_fuzz_dep_.CoverTab[101048]++
													c.memberID = ""
													return c.newSession(ctx, topics, handler, retries)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:295
		// _ = "end of CoverTab[101048]"
	case ErrNotCoordinatorForConsumer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:296
		_go_fuzz_dep_.CoverTab[101049]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:297
			_go_fuzz_dep_.CoverTab[101054]++
														return nil, join.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:298
			// _ = "end of CoverTab[101054]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:299
			_go_fuzz_dep_.CoverTab[101055]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:299
			// _ = "end of CoverTab[101055]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:299
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:299
		// _ = "end of CoverTab[101049]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:299
		_go_fuzz_dep_.CoverTab[101050]++

													return c.retryNewSession(ctx, topics, handler, retries, true)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:301
		// _ = "end of CoverTab[101050]"
	case ErrRebalanceInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:302
		_go_fuzz_dep_.CoverTab[101051]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:303
			_go_fuzz_dep_.CoverTab[101056]++
														return nil, join.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:304
			// _ = "end of CoverTab[101056]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:305
			_go_fuzz_dep_.CoverTab[101057]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:305
			// _ = "end of CoverTab[101057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:305
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:305
		// _ = "end of CoverTab[101051]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:305
		_go_fuzz_dep_.CoverTab[101052]++

													return c.retryNewSession(ctx, topics, handler, retries, false)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:307
		// _ = "end of CoverTab[101052]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:308
		_go_fuzz_dep_.CoverTab[101053]++
													return nil, join.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:309
		// _ = "end of CoverTab[101053]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:310
	// _ = "end of CoverTab[101021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:310
	_go_fuzz_dep_.CoverTab[101022]++

	// Prepare distribution plan if we joined as the leader
	var plan BalanceStrategyPlan
	if join.LeaderId == join.MemberId {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:314
		_go_fuzz_dep_.CoverTab[101058]++
													members, err := join.GetMembers()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:316
			_go_fuzz_dep_.CoverTab[101060]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:317
			// _ = "end of CoverTab[101060]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:318
			_go_fuzz_dep_.CoverTab[101061]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:318
			// _ = "end of CoverTab[101061]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:318
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:318
		// _ = "end of CoverTab[101058]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:318
		_go_fuzz_dep_.CoverTab[101059]++

													plan, err = c.balance(members)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:321
			_go_fuzz_dep_.CoverTab[101062]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:322
			// _ = "end of CoverTab[101062]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:323
			_go_fuzz_dep_.CoverTab[101063]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:323
			// _ = "end of CoverTab[101063]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:323
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:323
		// _ = "end of CoverTab[101059]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:324
		_go_fuzz_dep_.CoverTab[101064]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:324
		// _ = "end of CoverTab[101064]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:324
	// _ = "end of CoverTab[101022]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:324
	_go_fuzz_dep_.CoverTab[101023]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:327
	groupRequest, err := c.syncGroupRequest(coordinator, plan, join.GenerationId)
	if consumerGroupSyncTotal != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:328
		_go_fuzz_dep_.CoverTab[101065]++
													consumerGroupSyncTotal.Inc(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:329
		// _ = "end of CoverTab[101065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:330
		_go_fuzz_dep_.CoverTab[101066]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:330
		// _ = "end of CoverTab[101066]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:330
	// _ = "end of CoverTab[101023]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:330
	_go_fuzz_dep_.CoverTab[101024]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:331
		_go_fuzz_dep_.CoverTab[101067]++
													_ = coordinator.Close()
													if consumerGroupSyncFailed != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:333
			_go_fuzz_dep_.CoverTab[101069]++
														consumerGroupSyncFailed.Inc(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:334
			// _ = "end of CoverTab[101069]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:335
			_go_fuzz_dep_.CoverTab[101070]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:335
			// _ = "end of CoverTab[101070]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:335
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:335
		// _ = "end of CoverTab[101067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:335
		_go_fuzz_dep_.CoverTab[101068]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:336
		// _ = "end of CoverTab[101068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:337
		_go_fuzz_dep_.CoverTab[101071]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:337
		// _ = "end of CoverTab[101071]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:337
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:337
	// _ = "end of CoverTab[101024]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:337
	_go_fuzz_dep_.CoverTab[101025]++
												if groupRequest.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:338
		_go_fuzz_dep_.CoverTab[101072]++
													if consumerGroupSyncFailed != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:339
			_go_fuzz_dep_.CoverTab[101073]++
														consumerGroupSyncFailed.Inc(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:340
			// _ = "end of CoverTab[101073]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:341
			_go_fuzz_dep_.CoverTab[101074]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:341
			// _ = "end of CoverTab[101074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:341
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:341
		// _ = "end of CoverTab[101072]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:342
		_go_fuzz_dep_.CoverTab[101075]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:342
		// _ = "end of CoverTab[101075]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:342
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:342
	// _ = "end of CoverTab[101025]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:342
	_go_fuzz_dep_.CoverTab[101026]++

												switch groupRequest.Err {
	case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:345
		_go_fuzz_dep_.CoverTab[101076]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:345
		// _ = "end of CoverTab[101076]"
	case ErrUnknownMemberId, ErrIllegalGeneration:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:346
		_go_fuzz_dep_.CoverTab[101077]++
													c.memberID = ""
													return c.newSession(ctx, topics, handler, retries)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:348
		// _ = "end of CoverTab[101077]"
	case ErrNotCoordinatorForConsumer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:349
		_go_fuzz_dep_.CoverTab[101078]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:350
			_go_fuzz_dep_.CoverTab[101083]++
														return nil, groupRequest.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:351
			// _ = "end of CoverTab[101083]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:352
			_go_fuzz_dep_.CoverTab[101084]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:352
			// _ = "end of CoverTab[101084]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:352
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:352
		// _ = "end of CoverTab[101078]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:352
		_go_fuzz_dep_.CoverTab[101079]++

													return c.retryNewSession(ctx, topics, handler, retries, true)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:354
		// _ = "end of CoverTab[101079]"
	case ErrRebalanceInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:355
		_go_fuzz_dep_.CoverTab[101080]++
													if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:356
			_go_fuzz_dep_.CoverTab[101085]++
														return nil, groupRequest.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:357
			// _ = "end of CoverTab[101085]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:358
			_go_fuzz_dep_.CoverTab[101086]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:358
			// _ = "end of CoverTab[101086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:358
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:358
		// _ = "end of CoverTab[101080]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:358
		_go_fuzz_dep_.CoverTab[101081]++

													return c.retryNewSession(ctx, topics, handler, retries, false)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:360
		// _ = "end of CoverTab[101081]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:361
		_go_fuzz_dep_.CoverTab[101082]++
													return nil, groupRequest.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:362
		// _ = "end of CoverTab[101082]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:363
	// _ = "end of CoverTab[101026]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:363
	_go_fuzz_dep_.CoverTab[101027]++

	// Retrieve and sort claims
	var claims map[string][]int32
	if len(groupRequest.MemberAssignment) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:367
		_go_fuzz_dep_.CoverTab[101087]++
													members, err := groupRequest.GetMemberAssignment()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:369
			_go_fuzz_dep_.CoverTab[101090]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:370
			// _ = "end of CoverTab[101090]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:371
			_go_fuzz_dep_.CoverTab[101091]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:371
			// _ = "end of CoverTab[101091]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:371
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:371
		// _ = "end of CoverTab[101087]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:371
		_go_fuzz_dep_.CoverTab[101088]++
													claims = members.Topics

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:377
		if members.UserData != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:377
			_go_fuzz_dep_.CoverTab[101092]++
														c.userData = members.UserData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:378
			// _ = "end of CoverTab[101092]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:379
			_go_fuzz_dep_.CoverTab[101093]++
														c.userData = c.config.Consumer.Group.Member.UserData
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:380
			// _ = "end of CoverTab[101093]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:381
		// _ = "end of CoverTab[101088]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:381
		_go_fuzz_dep_.CoverTab[101089]++

													for _, partitions := range claims {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:383
			_go_fuzz_dep_.CoverTab[101094]++
														sort.Sort(int32Slice(partitions))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:384
			// _ = "end of CoverTab[101094]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:385
		// _ = "end of CoverTab[101089]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:386
		_go_fuzz_dep_.CoverTab[101095]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:386
		// _ = "end of CoverTab[101095]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:386
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:386
	// _ = "end of CoverTab[101027]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:386
	_go_fuzz_dep_.CoverTab[101028]++

												return newConsumerGroupSession(ctx, c, claims, join.MemberId, join.GenerationId, handler)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:388
	// _ = "end of CoverTab[101028]"
}

func (c *consumerGroup) joinGroupRequest(coordinator *Broker, topics []string) (*JoinGroupResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:391
	_go_fuzz_dep_.CoverTab[101096]++
												req := &JoinGroupRequest{
		GroupId:	c.groupID,
		MemberId:	c.memberID,
		SessionTimeout:	int32(c.config.Consumer.Group.Session.Timeout / time.Millisecond),
		ProtocolType:	"consumer",
	}
	if c.config.Version.IsAtLeast(V0_10_1_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:398
		_go_fuzz_dep_.CoverTab[101099]++
													req.Version = 1
													req.RebalanceTimeout = int32(c.config.Consumer.Group.Rebalance.Timeout / time.Millisecond)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:400
		// _ = "end of CoverTab[101099]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:401
		_go_fuzz_dep_.CoverTab[101100]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:401
		// _ = "end of CoverTab[101100]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:401
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:401
	// _ = "end of CoverTab[101096]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:401
	_go_fuzz_dep_.CoverTab[101097]++

												meta := &ConsumerGroupMemberMetadata{
		Topics:		topics,
		UserData:	c.userData,
	}
	strategy := c.config.Consumer.Group.Rebalance.Strategy
	if err := req.AddGroupProtocolMetadata(strategy.Name(), meta); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:408
		_go_fuzz_dep_.CoverTab[101101]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:409
		// _ = "end of CoverTab[101101]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:410
		_go_fuzz_dep_.CoverTab[101102]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:410
		// _ = "end of CoverTab[101102]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:410
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:410
	// _ = "end of CoverTab[101097]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:410
	_go_fuzz_dep_.CoverTab[101098]++

												return coordinator.JoinGroup(req)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:412
	// _ = "end of CoverTab[101098]"
}

func (c *consumerGroup) syncGroupRequest(coordinator *Broker, plan BalanceStrategyPlan, generationID int32) (*SyncGroupResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:415
	_go_fuzz_dep_.CoverTab[101103]++
												req := &SyncGroupRequest{
		GroupId:	c.groupID,
		MemberId:	c.memberID,
		GenerationId:	generationID,
	}
	strategy := c.config.Consumer.Group.Rebalance.Strategy
	for memberID, topics := range plan {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:422
		_go_fuzz_dep_.CoverTab[101105]++
													assignment := &ConsumerGroupMemberAssignment{Topics: topics}
													userDataBytes, err := strategy.AssignmentData(memberID, topics, generationID)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:425
			_go_fuzz_dep_.CoverTab[101107]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:426
			// _ = "end of CoverTab[101107]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:427
			_go_fuzz_dep_.CoverTab[101108]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:427
			// _ = "end of CoverTab[101108]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:427
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:427
		// _ = "end of CoverTab[101105]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:427
		_go_fuzz_dep_.CoverTab[101106]++
													assignment.UserData = userDataBytes
													if err := req.AddGroupAssignmentMember(memberID, assignment); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:429
			_go_fuzz_dep_.CoverTab[101109]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:430
			// _ = "end of CoverTab[101109]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:431
			_go_fuzz_dep_.CoverTab[101110]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:431
			// _ = "end of CoverTab[101110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:431
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:431
		// _ = "end of CoverTab[101106]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:432
	// _ = "end of CoverTab[101103]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:432
	_go_fuzz_dep_.CoverTab[101104]++
												return coordinator.SyncGroup(req)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:433
	// _ = "end of CoverTab[101104]"
}

func (c *consumerGroup) heartbeatRequest(coordinator *Broker, memberID string, generationID int32) (*HeartbeatResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:436
	_go_fuzz_dep_.CoverTab[101111]++
												req := &HeartbeatRequest{
		GroupId:	c.groupID,
		MemberId:	memberID,
		GenerationId:	generationID,
	}

												return coordinator.Heartbeat(req)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:443
	// _ = "end of CoverTab[101111]"
}

func (c *consumerGroup) balance(members map[string]ConsumerGroupMemberMetadata) (BalanceStrategyPlan, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:446
	_go_fuzz_dep_.CoverTab[101112]++
												topics := make(map[string][]int32)
												for _, meta := range members {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:448
		_go_fuzz_dep_.CoverTab[101115]++
													for _, topic := range meta.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:449
			_go_fuzz_dep_.CoverTab[101116]++
														topics[topic] = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:450
			// _ = "end of CoverTab[101116]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:451
		// _ = "end of CoverTab[101115]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:452
	// _ = "end of CoverTab[101112]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:452
	_go_fuzz_dep_.CoverTab[101113]++

												for topic := range topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:454
		_go_fuzz_dep_.CoverTab[101117]++
													partitions, err := c.client.Partitions(topic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:456
			_go_fuzz_dep_.CoverTab[101119]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:457
			// _ = "end of CoverTab[101119]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:458
			_go_fuzz_dep_.CoverTab[101120]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:458
			// _ = "end of CoverTab[101120]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:458
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:458
		// _ = "end of CoverTab[101117]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:458
		_go_fuzz_dep_.CoverTab[101118]++
													topics[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:459
		// _ = "end of CoverTab[101118]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:460
	// _ = "end of CoverTab[101113]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:460
	_go_fuzz_dep_.CoverTab[101114]++

												strategy := c.config.Consumer.Group.Rebalance.Strategy
												return strategy.Plan(members, topics)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:463
	// _ = "end of CoverTab[101114]"
}

// Leaves the cluster, called by Close.
func (c *consumerGroup) leave() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:467
	_go_fuzz_dep_.CoverTab[101121]++
												c.lock.Lock()
												defer c.lock.Unlock()
												if c.memberID == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:470
		_go_fuzz_dep_.CoverTab[101125]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:471
		// _ = "end of CoverTab[101125]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:472
		_go_fuzz_dep_.CoverTab[101126]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:472
		// _ = "end of CoverTab[101126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:472
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:472
	// _ = "end of CoverTab[101121]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:472
	_go_fuzz_dep_.CoverTab[101122]++

												coordinator, err := c.client.Coordinator(c.groupID)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:475
		_go_fuzz_dep_.CoverTab[101127]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:476
		// _ = "end of CoverTab[101127]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:477
		_go_fuzz_dep_.CoverTab[101128]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:477
		// _ = "end of CoverTab[101128]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:477
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:477
	// _ = "end of CoverTab[101122]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:477
	_go_fuzz_dep_.CoverTab[101123]++

												resp, err := coordinator.LeaveGroup(&LeaveGroupRequest{
		GroupId:	c.groupID,
		MemberId:	c.memberID,
	})
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:483
		_go_fuzz_dep_.CoverTab[101129]++
													_ = coordinator.Close()
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:485
		// _ = "end of CoverTab[101129]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:486
		_go_fuzz_dep_.CoverTab[101130]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:486
		// _ = "end of CoverTab[101130]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:486
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:486
	// _ = "end of CoverTab[101123]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:486
	_go_fuzz_dep_.CoverTab[101124]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:489
	c.memberID = ""

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:492
	switch resp.Err {
	case ErrRebalanceInProgress, ErrUnknownMemberId, ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:493
		_go_fuzz_dep_.CoverTab[101131]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:494
		// _ = "end of CoverTab[101131]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:495
		_go_fuzz_dep_.CoverTab[101132]++
													return resp.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:496
		// _ = "end of CoverTab[101132]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:497
	// _ = "end of CoverTab[101124]"
}

func (c *consumerGroup) handleError(err error, topic string, partition int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:500
	_go_fuzz_dep_.CoverTab[101133]++
												if _, ok := err.(*ConsumerError); !ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
		_go_fuzz_dep_.CoverTab[101137]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
		return topic != ""
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
		// _ = "end of CoverTab[101137]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
		_go_fuzz_dep_.CoverTab[101138]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
		return partition > -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
		// _ = "end of CoverTab[101138]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:501
		_go_fuzz_dep_.CoverTab[101139]++
													err = &ConsumerError{
			Topic:		topic,
			Partition:	partition,
			Err:		err,
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:506
		// _ = "end of CoverTab[101139]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:507
		_go_fuzz_dep_.CoverTab[101140]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:507
		// _ = "end of CoverTab[101140]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:507
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:507
	// _ = "end of CoverTab[101133]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:507
	_go_fuzz_dep_.CoverTab[101134]++

												if !c.config.Consumer.Return.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:509
		_go_fuzz_dep_.CoverTab[101141]++
													Logger.Println(err)
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:511
		// _ = "end of CoverTab[101141]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:512
		_go_fuzz_dep_.CoverTab[101142]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:512
		// _ = "end of CoverTab[101142]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:512
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:512
	// _ = "end of CoverTab[101134]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:512
	_go_fuzz_dep_.CoverTab[101135]++

												select {
	case <-c.closed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:515
		_go_fuzz_dep_.CoverTab[101143]++

													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:517
		// _ = "end of CoverTab[101143]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:518
		_go_fuzz_dep_.CoverTab[101144]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:518
		// _ = "end of CoverTab[101144]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:519
	// _ = "end of CoverTab[101135]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:519
	_go_fuzz_dep_.CoverTab[101136]++

												select {
	case c.errors <- err:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:522
		_go_fuzz_dep_.CoverTab[101145]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:522
		// _ = "end of CoverTab[101145]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:523
		_go_fuzz_dep_.CoverTab[101146]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:523
		// _ = "end of CoverTab[101146]"

	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:525
	// _ = "end of CoverTab[101136]"
}

func (c *consumerGroup) loopCheckPartitionNumbers(topics []string, session *consumerGroupSession) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:528
	_go_fuzz_dep_.CoverTab[101147]++
												pause := time.NewTicker(c.config.Metadata.RefreshFrequency)
												defer session.cancel()
												defer pause.Stop()
												var oldTopicToPartitionNum map[string]int
												var err error
												if oldTopicToPartitionNum, err = c.topicToPartitionNumbers(topics); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:534
		_go_fuzz_dep_.CoverTab[101149]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:535
		// _ = "end of CoverTab[101149]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:536
		_go_fuzz_dep_.CoverTab[101150]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:536
		// _ = "end of CoverTab[101150]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:536
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:536
	// _ = "end of CoverTab[101147]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:536
	_go_fuzz_dep_.CoverTab[101148]++
												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:537
		_go_fuzz_dep_.CoverTab[101151]++
													if newTopicToPartitionNum, err := c.topicToPartitionNumbers(topics); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:538
			_go_fuzz_dep_.CoverTab[101153]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:539
			// _ = "end of CoverTab[101153]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:540
			_go_fuzz_dep_.CoverTab[101154]++
														for topic, num := range oldTopicToPartitionNum {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:541
				_go_fuzz_dep_.CoverTab[101155]++
															if newTopicToPartitionNum[topic] != num {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:542
					_go_fuzz_dep_.CoverTab[101156]++
																return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:543
					// _ = "end of CoverTab[101156]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:544
					_go_fuzz_dep_.CoverTab[101157]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:544
					// _ = "end of CoverTab[101157]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:544
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:544
				// _ = "end of CoverTab[101155]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:545
			// _ = "end of CoverTab[101154]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:546
		// _ = "end of CoverTab[101151]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:546
		_go_fuzz_dep_.CoverTab[101152]++
													select {
		case <-pause.C:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:548
			_go_fuzz_dep_.CoverTab[101158]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:548
			// _ = "end of CoverTab[101158]"
		case <-session.ctx.Done():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:549
			_go_fuzz_dep_.CoverTab[101159]++
														Logger.Printf(
				"consumergroup/%s loop check partition number coroutine will exit, topics %s\n",
				c.groupID, topics)

														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:554
			// _ = "end of CoverTab[101159]"
		case <-c.closed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:555
			_go_fuzz_dep_.CoverTab[101160]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:556
			// _ = "end of CoverTab[101160]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:557
		// _ = "end of CoverTab[101152]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:558
	// _ = "end of CoverTab[101148]"
}

func (c *consumerGroup) topicToPartitionNumbers(topics []string) (map[string]int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:561
	_go_fuzz_dep_.CoverTab[101161]++
												topicToPartitionNum := make(map[string]int, len(topics))
												for _, topic := range topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:563
		_go_fuzz_dep_.CoverTab[101163]++
													if partitionNum, err := c.client.Partitions(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:564
			_go_fuzz_dep_.CoverTab[101164]++
														Logger.Printf(
				"consumergroup/%s topic %s get partition number failed %v\n",
				c.groupID, err)
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:568
			// _ = "end of CoverTab[101164]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:569
			_go_fuzz_dep_.CoverTab[101165]++
														topicToPartitionNum[topic] = len(partitionNum)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:570
			// _ = "end of CoverTab[101165]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:571
		// _ = "end of CoverTab[101163]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:572
	// _ = "end of CoverTab[101161]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:572
	_go_fuzz_dep_.CoverTab[101162]++
												return topicToPartitionNum, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:573
	// _ = "end of CoverTab[101162]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:578
// ConsumerGroupSession represents a consumer group member session.
type ConsumerGroupSession interface {
	// Claims returns information about the claimed partitions by topic.
	Claims() map[string][]int32

	// MemberID returns the cluster member ID.
	MemberID() string

	// GenerationID returns the current generation ID.
	GenerationID() int32

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
	MarkOffset(topic string, partition int32, offset int64, metadata string)

	// Commit the offset to the backend
	//
	// Note: calling Commit performs a blocking synchronous operation.
	Commit()

	// ResetOffset resets to the provided offset, alongside a metadata string that
	// represents the state of the partition consumer at that point in time. Reset
	// acts as a counterpart to MarkOffset, the difference being that it allows to
	// reset an offset to an earlier or smaller value, where MarkOffset only
	// allows incrementing the offset. cf MarkOffset for more details.
	ResetOffset(topic string, partition int32, offset int64, metadata string)

	// MarkMessage marks a message as consumed.
	MarkMessage(msg *ConsumerMessage, metadata string)

	// Context returns the session context.
	Context() context.Context
}

type consumerGroupSession struct {
	parent		*consumerGroup
	memberID	string
	generationID	int32
	handler		ConsumerGroupHandler

	claims	map[string][]int32
	offsets	*offsetManager
	ctx	context.Context
	cancel	func()

	waitGroup	sync.WaitGroup
	releaseOnce	sync.Once
	hbDying, hbDead	chan none
}

func newConsumerGroupSession(ctx context.Context, parent *consumerGroup, claims map[string][]int32, memberID string, generationID int32, handler ConsumerGroupHandler) (*consumerGroupSession, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:639
	_go_fuzz_dep_.CoverTab[101166]++

												offsets, err := newOffsetManagerFromClient(parent.groupID, memberID, generationID, parent.client)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:642
		_go_fuzz_dep_.CoverTab[101171]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:643
		// _ = "end of CoverTab[101171]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:644
		_go_fuzz_dep_.CoverTab[101172]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:644
		// _ = "end of CoverTab[101172]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:644
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:644
	// _ = "end of CoverTab[101166]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:644
	_go_fuzz_dep_.CoverTab[101167]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:647
	ctx, cancel := context.WithCancel(ctx)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:650
	sess := &consumerGroupSession{
		parent:		parent,
		memberID:	memberID,
		generationID:	generationID,
		handler:	handler,
		offsets:	offsets,
		claims:		claims,
		ctx:		ctx,
		cancel:		cancel,
		hbDying:	make(chan none),
		hbDead:		make(chan none),
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:661
	_curRoutineNum134_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:661
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum134_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:664
	go sess.heartbeatLoop()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:667
	for topic, partitions := range claims {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:667
		_go_fuzz_dep_.CoverTab[101173]++
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:668
			_go_fuzz_dep_.CoverTab[101174]++
														pom, err := offsets.ManagePartition(topic, partition)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:670
				_go_fuzz_dep_.CoverTab[101176]++
															_ = sess.release(false)
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:672
				// _ = "end of CoverTab[101176]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:673
				_go_fuzz_dep_.CoverTab[101177]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:673
				// _ = "end of CoverTab[101177]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:673
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:673
			// _ = "end of CoverTab[101174]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:673
			_go_fuzz_dep_.CoverTab[101175]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:673
			_curRoutineNum135_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:673
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum135_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:676
			go func(topic string, partition int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:676
				_go_fuzz_dep_.CoverTab[101178]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:676
				defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:676
					_go_fuzz_dep_.CoverTab[101179]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:676
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum135_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:676
					// _ = "end of CoverTab[101179]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:676
				}()
															for err := range pom.Errors() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:677
					_go_fuzz_dep_.CoverTab[101180]++
																sess.parent.handleError(err, topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:678
					// _ = "end of CoverTab[101180]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:679
				// _ = "end of CoverTab[101178]"
			}(topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:680
			// _ = "end of CoverTab[101175]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:681
		// _ = "end of CoverTab[101173]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:682
	// _ = "end of CoverTab[101167]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:682
	_go_fuzz_dep_.CoverTab[101168]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:685
	if err := handler.Setup(sess); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:685
		_go_fuzz_dep_.CoverTab[101181]++
													_ = sess.release(true)
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:687
		// _ = "end of CoverTab[101181]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:688
		_go_fuzz_dep_.CoverTab[101182]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:688
		// _ = "end of CoverTab[101182]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:688
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:688
	// _ = "end of CoverTab[101168]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:688
	_go_fuzz_dep_.CoverTab[101169]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:691
	for topic, partitions := range claims {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:691
		_go_fuzz_dep_.CoverTab[101183]++
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:692
			_go_fuzz_dep_.CoverTab[101184]++
														sess.waitGroup.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:693
			_curRoutineNum136_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:693
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum136_)

														go func(topic string, partition int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:695
				_go_fuzz_dep_.CoverTab[101185]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:695
				defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:695
					_go_fuzz_dep_.CoverTab[101186]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:695
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum136_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:695
					// _ = "end of CoverTab[101186]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:695
				}()
															defer sess.waitGroup.Done()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:700
				defer sess.cancel()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:703
				sess.consume(topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:703
				// _ = "end of CoverTab[101185]"
			}(topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:704
			// _ = "end of CoverTab[101184]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:705
		// _ = "end of CoverTab[101183]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:706
	// _ = "end of CoverTab[101169]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:706
	_go_fuzz_dep_.CoverTab[101170]++
												return sess, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:707
	// _ = "end of CoverTab[101170]"
}

func (s *consumerGroupSession) Claims() map[string][]int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:710
	_go_fuzz_dep_.CoverTab[101187]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:710
	return s.claims
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:710
	// _ = "end of CoverTab[101187]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:710
}
func (s *consumerGroupSession) MemberID() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:711
	_go_fuzz_dep_.CoverTab[101188]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:711
	return s.memberID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:711
	// _ = "end of CoverTab[101188]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:711
}
func (s *consumerGroupSession) GenerationID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:712
	_go_fuzz_dep_.CoverTab[101189]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:712
	return s.generationID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:712
	// _ = "end of CoverTab[101189]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:712
}

func (s *consumerGroupSession) MarkOffset(topic string, partition int32, offset int64, metadata string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:714
	_go_fuzz_dep_.CoverTab[101190]++
												if pom := s.offsets.findPOM(topic, partition); pom != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:715
		_go_fuzz_dep_.CoverTab[101191]++
													pom.MarkOffset(offset, metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:716
		// _ = "end of CoverTab[101191]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:717
		_go_fuzz_dep_.CoverTab[101192]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:717
		// _ = "end of CoverTab[101192]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:717
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:717
	// _ = "end of CoverTab[101190]"
}

func (s *consumerGroupSession) Commit() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:720
	_go_fuzz_dep_.CoverTab[101193]++
												s.offsets.Commit()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:721
	// _ = "end of CoverTab[101193]"
}

func (s *consumerGroupSession) ResetOffset(topic string, partition int32, offset int64, metadata string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:724
	_go_fuzz_dep_.CoverTab[101194]++
												if pom := s.offsets.findPOM(topic, partition); pom != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:725
		_go_fuzz_dep_.CoverTab[101195]++
													pom.ResetOffset(offset, metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:726
		// _ = "end of CoverTab[101195]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:727
		_go_fuzz_dep_.CoverTab[101196]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:727
		// _ = "end of CoverTab[101196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:727
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:727
	// _ = "end of CoverTab[101194]"
}

func (s *consumerGroupSession) MarkMessage(msg *ConsumerMessage, metadata string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:730
	_go_fuzz_dep_.CoverTab[101197]++
												s.MarkOffset(msg.Topic, msg.Partition, msg.Offset+1, metadata)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:731
	// _ = "end of CoverTab[101197]"
}

func (s *consumerGroupSession) Context() context.Context {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:734
	_go_fuzz_dep_.CoverTab[101198]++
												return s.ctx
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:735
	// _ = "end of CoverTab[101198]"
}

func (s *consumerGroupSession) consume(topic string, partition int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:738
	_go_fuzz_dep_.CoverTab[101199]++

												select {
	case <-s.ctx.Done():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:741
		_go_fuzz_dep_.CoverTab[101206]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:742
		// _ = "end of CoverTab[101206]"
	case <-s.parent.closed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:743
		_go_fuzz_dep_.CoverTab[101207]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:744
		// _ = "end of CoverTab[101207]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:745
		_go_fuzz_dep_.CoverTab[101208]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:745
		// _ = "end of CoverTab[101208]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:746
	// _ = "end of CoverTab[101199]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:746
	_go_fuzz_dep_.CoverTab[101200]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:749
	offset := s.parent.config.Consumer.Offsets.Initial
	if pom := s.offsets.findPOM(topic, partition); pom != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:750
		_go_fuzz_dep_.CoverTab[101209]++
													offset, _ = pom.NextOffset()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:751
		// _ = "end of CoverTab[101209]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:752
		_go_fuzz_dep_.CoverTab[101210]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:752
		// _ = "end of CoverTab[101210]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:752
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:752
	// _ = "end of CoverTab[101200]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:752
	_go_fuzz_dep_.CoverTab[101201]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:755
	claim, err := newConsumerGroupClaim(s, topic, partition, offset)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:756
		_go_fuzz_dep_.CoverTab[101211]++
													s.parent.handleError(err, topic, partition)
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:758
		// _ = "end of CoverTab[101211]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:759
		_go_fuzz_dep_.CoverTab[101212]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:759
		// _ = "end of CoverTab[101212]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:759
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:759
	// _ = "end of CoverTab[101201]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:759
	_go_fuzz_dep_.CoverTab[101202]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:759
	_curRoutineNum137_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:759
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum137_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:762
	go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:762
		_go_fuzz_dep_.CoverTab[101213]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:762
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:762
			_go_fuzz_dep_.CoverTab[101214]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:762
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum137_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:762
			// _ = "end of CoverTab[101214]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:762
		}()
													for err := range claim.Errors() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:763
			_go_fuzz_dep_.CoverTab[101215]++
														s.parent.handleError(err, topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:764
			// _ = "end of CoverTab[101215]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:765
		// _ = "end of CoverTab[101213]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:766
	// _ = "end of CoverTab[101202]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:766
	_go_fuzz_dep_.CoverTab[101203]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:766
	_curRoutineNum138_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:766
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum138_)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:769
	go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:769
		_go_fuzz_dep_.CoverTab[101216]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:769
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:769
			_go_fuzz_dep_.CoverTab[101218]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:769
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum138_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:769
			// _ = "end of CoverTab[101218]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:769
		}()
													select {
		case <-s.ctx.Done():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:771
			_go_fuzz_dep_.CoverTab[101219]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:771
			// _ = "end of CoverTab[101219]"
		case <-s.parent.closed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:772
			_go_fuzz_dep_.CoverTab[101220]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:772
			// _ = "end of CoverTab[101220]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:773
		// _ = "end of CoverTab[101216]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:773
		_go_fuzz_dep_.CoverTab[101217]++
													claim.AsyncClose()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:774
		// _ = "end of CoverTab[101217]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:775
	// _ = "end of CoverTab[101203]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:775
	_go_fuzz_dep_.CoverTab[101204]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:778
	if err := s.handler.ConsumeClaim(s, claim); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:778
		_go_fuzz_dep_.CoverTab[101221]++
													s.parent.handleError(err, topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:779
		// _ = "end of CoverTab[101221]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:780
		_go_fuzz_dep_.CoverTab[101222]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:780
		// _ = "end of CoverTab[101222]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:780
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:780
	// _ = "end of CoverTab[101204]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:780
	_go_fuzz_dep_.CoverTab[101205]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:783
	claim.AsyncClose()
	for _, err := range claim.waitClosed() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:784
		_go_fuzz_dep_.CoverTab[101223]++
													s.parent.handleError(err, topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:785
		// _ = "end of CoverTab[101223]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:786
	// _ = "end of CoverTab[101205]"
}

func (s *consumerGroupSession) release(withCleanup bool) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:789
	_go_fuzz_dep_.CoverTab[101224]++

												s.cancel()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:794
	s.waitGroup.Wait()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:797
	s.releaseOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:797
		_go_fuzz_dep_.CoverTab[101226]++
													if withCleanup {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:798
			_go_fuzz_dep_.CoverTab[101229]++
														if e := s.handler.Cleanup(s); e != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:799
				_go_fuzz_dep_.CoverTab[101230]++
															s.parent.handleError(e, "", -1)
															err = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:801
				// _ = "end of CoverTab[101230]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:802
				_go_fuzz_dep_.CoverTab[101231]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:802
				// _ = "end of CoverTab[101231]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:802
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:802
			// _ = "end of CoverTab[101229]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:803
			_go_fuzz_dep_.CoverTab[101232]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:803
			// _ = "end of CoverTab[101232]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:803
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:803
		// _ = "end of CoverTab[101226]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:803
		_go_fuzz_dep_.CoverTab[101227]++

													if e := s.offsets.Close(); e != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:805
			_go_fuzz_dep_.CoverTab[101233]++
														err = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:806
			// _ = "end of CoverTab[101233]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:807
			_go_fuzz_dep_.CoverTab[101234]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:807
			// _ = "end of CoverTab[101234]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:807
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:807
		// _ = "end of CoverTab[101227]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:807
		_go_fuzz_dep_.CoverTab[101228]++

													close(s.hbDying)
													<-s.hbDead
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:810
		// _ = "end of CoverTab[101228]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:811
	// _ = "end of CoverTab[101224]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:811
	_go_fuzz_dep_.CoverTab[101225]++

												Logger.Printf(
		"consumergroup/session/%s/%d released\n",
		s.MemberID(), s.GenerationID())

												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:817
	// _ = "end of CoverTab[101225]"
}

func (s *consumerGroupSession) heartbeatLoop() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:820
	_go_fuzz_dep_.CoverTab[101235]++
												defer close(s.hbDead)
												defer s.cancel()
												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:823
		_go_fuzz_dep_.CoverTab[101237]++
													Logger.Printf(
			"consumergroup/session/%s/%d heartbeat loop stopped\n",
			s.MemberID(), s.GenerationID())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:826
		// _ = "end of CoverTab[101237]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:827
	// _ = "end of CoverTab[101235]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:827
	_go_fuzz_dep_.CoverTab[101236]++

												pause := time.NewTicker(s.parent.config.Consumer.Group.Heartbeat.Interval)
												defer pause.Stop()

												retryBackoff := time.NewTimer(s.parent.config.Metadata.Retry.Backoff)
												defer retryBackoff.Stop()

												retries := s.parent.config.Metadata.Retry.Max
												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:836
		_go_fuzz_dep_.CoverTab[101238]++
													coordinator, err := s.parent.client.Coordinator(s.parent.groupID)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:838
			_go_fuzz_dep_.CoverTab[101242]++
														if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:839
				_go_fuzz_dep_.CoverTab[101245]++
															s.parent.handleError(err, "", -1)
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:841
				// _ = "end of CoverTab[101245]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:842
				_go_fuzz_dep_.CoverTab[101246]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:842
				// _ = "end of CoverTab[101246]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:842
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:842
			// _ = "end of CoverTab[101242]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:842
			_go_fuzz_dep_.CoverTab[101243]++
														retryBackoff.Reset(s.parent.config.Metadata.Retry.Backoff)
														select {
			case <-s.hbDying:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:845
				_go_fuzz_dep_.CoverTab[101247]++
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:846
				// _ = "end of CoverTab[101247]"
			case <-retryBackoff.C:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:847
				_go_fuzz_dep_.CoverTab[101248]++
															retries--
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:848
				// _ = "end of CoverTab[101248]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:849
			// _ = "end of CoverTab[101243]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:849
			_go_fuzz_dep_.CoverTab[101244]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:850
			// _ = "end of CoverTab[101244]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:851
			_go_fuzz_dep_.CoverTab[101249]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:851
			// _ = "end of CoverTab[101249]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:851
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:851
		// _ = "end of CoverTab[101238]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:851
		_go_fuzz_dep_.CoverTab[101239]++

													resp, err := s.parent.heartbeatRequest(coordinator, s.memberID, s.generationID)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:854
			_go_fuzz_dep_.CoverTab[101250]++
														_ = coordinator.Close()

														if retries <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:857
				_go_fuzz_dep_.CoverTab[101252]++
															s.parent.handleError(err, "", -1)
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:859
				// _ = "end of CoverTab[101252]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:860
				_go_fuzz_dep_.CoverTab[101253]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:860
				// _ = "end of CoverTab[101253]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:860
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:860
			// _ = "end of CoverTab[101250]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:860
			_go_fuzz_dep_.CoverTab[101251]++

														retries--
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:863
			// _ = "end of CoverTab[101251]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:864
			_go_fuzz_dep_.CoverTab[101254]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:864
			// _ = "end of CoverTab[101254]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:864
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:864
		// _ = "end of CoverTab[101239]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:864
		_go_fuzz_dep_.CoverTab[101240]++

													switch resp.Err {
		case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:867
			_go_fuzz_dep_.CoverTab[101255]++
														retries = s.parent.config.Metadata.Retry.Max
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:868
			// _ = "end of CoverTab[101255]"
		case ErrRebalanceInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:869
			_go_fuzz_dep_.CoverTab[101256]++
														retries = s.parent.config.Metadata.Retry.Max
														s.cancel()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:871
			// _ = "end of CoverTab[101256]"
		case ErrUnknownMemberId, ErrIllegalGeneration:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:872
			_go_fuzz_dep_.CoverTab[101257]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:873
			// _ = "end of CoverTab[101257]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:874
			_go_fuzz_dep_.CoverTab[101258]++
														s.parent.handleError(resp.Err, "", -1)
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:876
			// _ = "end of CoverTab[101258]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:877
		// _ = "end of CoverTab[101240]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:877
		_go_fuzz_dep_.CoverTab[101241]++

													select {
		case <-pause.C:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:880
			_go_fuzz_dep_.CoverTab[101259]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:880
			// _ = "end of CoverTab[101259]"
		case <-s.hbDying:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:881
			_go_fuzz_dep_.CoverTab[101260]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:882
			// _ = "end of CoverTab[101260]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:883
		// _ = "end of CoverTab[101241]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:884
	// _ = "end of CoverTab[101236]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:889
// ConsumerGroupHandler instances are used to handle individual topic/partition claims.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:889
// It also provides hooks for your consumer group session life-cycle and allow you to
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:889
// trigger logic before or after the consume loop(s).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:889
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:889
// PLEASE NOTE that handlers are likely be called from several goroutines concurrently,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:889
// ensure that all state is safely protected against race conditions.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:895
type ConsumerGroupHandler interface {
	// Setup is run at the beginning of a new session, before ConsumeClaim.
	Setup(ConsumerGroupSession) error

	// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
	// but before the offsets are committed for the very last time.
	Cleanup(ConsumerGroupSession) error

	// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
	// Once the Messages() channel is closed, the Handler must finish its processing
	// loop and exit.
	ConsumeClaim(ConsumerGroupSession, ConsumerGroupClaim) error
}

// ConsumerGroupClaim processes Kafka messages from a given topic and partition within a consumer group.
type ConsumerGroupClaim interface {
	// Topic returns the consumed topic name.
	Topic() string

	// Partition returns the consumed partition.
	Partition() int32

	// InitialOffset returns the initial offset that was used as a starting point for this claim.
	InitialOffset() int64

	// HighWaterMarkOffset returns the high water mark offset of the partition,
	// i.e. the offset that will be used for the next message that will be produced.
	// You can use this to determine how far behind the processing is.
	HighWaterMarkOffset() int64

	// Messages returns the read channel for the messages that are returned by
	// the broker. The messages channel will be closed when a new rebalance cycle
	// is due. You must finish processing and mark offsets within
	// Config.Consumer.Group.Session.Timeout before the topic/partition is eventually
	// re-assigned to another group member.
	Messages() <-chan *ConsumerMessage
}

type consumerGroupClaim struct {
	topic		string
	partition	int32
	offset		int64
	PartitionConsumer
}

func newConsumerGroupClaim(sess *consumerGroupSession, topic string, partition int32, offset int64) (*consumerGroupClaim, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:940
	_go_fuzz_dep_.CoverTab[101261]++
												pcm, err := sess.parent.consumer.ConsumePartition(topic, partition, offset)
												if err == ErrOffsetOutOfRange {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:942
		_go_fuzz_dep_.CoverTab[101265]++
													offset = sess.parent.config.Consumer.Offsets.Initial
													pcm, err = sess.parent.consumer.ConsumePartition(topic, partition, offset)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:944
		// _ = "end of CoverTab[101265]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:945
		_go_fuzz_dep_.CoverTab[101266]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:945
		// _ = "end of CoverTab[101266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:945
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:945
	// _ = "end of CoverTab[101261]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:945
	_go_fuzz_dep_.CoverTab[101262]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:946
		_go_fuzz_dep_.CoverTab[101267]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:947
		// _ = "end of CoverTab[101267]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:948
		_go_fuzz_dep_.CoverTab[101268]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:948
		// _ = "end of CoverTab[101268]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:948
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:948
	// _ = "end of CoverTab[101262]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:948
	_go_fuzz_dep_.CoverTab[101263]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:948
	_curRoutineNum139_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:948
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum139_)

												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:950
		_go_fuzz_dep_.CoverTab[101269]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:950
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:950
			_go_fuzz_dep_.CoverTab[101270]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:950
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum139_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:950
			// _ = "end of CoverTab[101270]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:950
		}()
													for err := range pcm.Errors() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:951
			_go_fuzz_dep_.CoverTab[101271]++
														sess.parent.handleError(err, topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:952
			// _ = "end of CoverTab[101271]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:953
		// _ = "end of CoverTab[101269]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:954
	// _ = "end of CoverTab[101263]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:954
	_go_fuzz_dep_.CoverTab[101264]++

												return &consumerGroupClaim{
		topic:			topic,
		partition:		partition,
		offset:			offset,
		PartitionConsumer:	pcm,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:961
	// _ = "end of CoverTab[101264]"
}

func (c *consumerGroupClaim) Topic() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:964
	_go_fuzz_dep_.CoverTab[101272]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:964
	return c.topic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:964
	// _ = "end of CoverTab[101272]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:964
}
func (c *consumerGroupClaim) Partition() int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:965
	_go_fuzz_dep_.CoverTab[101273]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:965
	return c.partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:965
	// _ = "end of CoverTab[101273]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:965
}
func (c *consumerGroupClaim) InitialOffset() int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:966
	_go_fuzz_dep_.CoverTab[101274]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:966
	return c.offset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:966
	// _ = "end of CoverTab[101274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:966
}

// Drains messages and errors, ensures the claim is fully closed.
func (c *consumerGroupClaim) waitClosed() (errs ConsumerErrors) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:969
	_go_fuzz_dep_.CoverTab[101275]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:969
	_curRoutineNum140_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:969
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum140_)
												go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:970
		_go_fuzz_dep_.CoverTab[101278]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:970
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:970
			_go_fuzz_dep_.CoverTab[101279]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:970
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum140_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:970
			// _ = "end of CoverTab[101279]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:970
		}()
													for range c.Messages() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:971
			_go_fuzz_dep_.CoverTab[101280]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:971
			// _ = "end of CoverTab[101280]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:972
		// _ = "end of CoverTab[101278]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:973
	// _ = "end of CoverTab[101275]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:973
	_go_fuzz_dep_.CoverTab[101276]++

												for err := range c.Errors() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:975
		_go_fuzz_dep_.CoverTab[101281]++
													errs = append(errs, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:976
		// _ = "end of CoverTab[101281]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:977
	// _ = "end of CoverTab[101276]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:977
	_go_fuzz_dep_.CoverTab[101277]++
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:978
	// _ = "end of CoverTab[101277]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:979
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/consumer_group.go:979
var _ = _go_fuzz_dep_.CoverTab
