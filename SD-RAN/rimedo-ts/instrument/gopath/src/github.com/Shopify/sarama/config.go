//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:1
)

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"regexp"
	"time"

	"github.com/rcrowley/go-metrics"
	"golang.org/x/net/proxy"
)

const defaultClientID = "sarama"

var validID = regexp.MustCompile(`\A[A-Za-z0-9._-]+\z`)

// Config is used to pass multiple configuration options to Sarama's constructors.
type Config struct {
	// Admin is the namespace for ClusterAdmin properties used by the administrative Kafka client.
	Admin	struct {
		Retry	struct {
			// The total number of times to retry sending (retriable) admin requests (default 5).
			// Similar to the `retries` setting of the JVM AdminClientConfig.
			Max	int
			// Backoff time between retries of a failed request (default 100ms)
			Backoff	time.Duration
		}
		// The maximum duration the administrative Kafka client will wait for ClusterAdmin operations,
		// including topics, brokers, configurations and ACLs (defaults to 3 seconds).
		Timeout	time.Duration
	}

	// Net is the namespace for network-level properties used by the Broker, and
	// shared by the Client/Producer/Consumer.
	Net	struct {
		// How many outstanding requests a connection is allowed to have before
		// sending on it blocks (default 5).
		// Throughput can improve but message ordering is not guaranteed if Producer.Idempotent is disabled, see:
		// https://kafka.apache.org/protocol#protocol_network
		// https://kafka.apache.org/28/documentation.html#producerconfigs_max.in.flight.requests.per.connection
		MaxOpenRequests	int

		// All three of the below configurations are similar to the
		// `socket.timeout.ms` setting in JVM kafka. All of them default
		// to 30 seconds.
		DialTimeout	time.Duration	// How long to wait for the initial connection.
		ReadTimeout	time.Duration	// How long to wait for a response.
		WriteTimeout	time.Duration	// How long to wait for a transmit.

		TLS	struct {
			// Whether or not to use TLS when connecting to the broker
			// (defaults to false).
			Enable	bool
			// The TLS configuration to use for secure connections if
			// enabled (defaults to nil).
			Config	*tls.Config
		}

		// SASL based authentication with broker. While there are multiple SASL authentication methods
		// the current implementation is limited to plaintext (SASL/PLAIN) authentication
		SASL	struct {
			// Whether or not to use SASL authentication when connecting to the broker
			// (defaults to false).
			Enable	bool
			// SASLMechanism is the name of the enabled SASL mechanism.
			// Possible values: OAUTHBEARER, PLAIN (defaults to PLAIN).
			Mechanism	SASLMechanism
			// Version is the SASL Protocol Version to use
			// Kafka > 1.x should use V1, except on Azure EventHub which use V0
			Version	int16
			// Whether or not to send the Kafka SASL handshake first if enabled
			// (defaults to true). You should only set this to false if you're using
			// a non-Kafka SASL proxy.
			Handshake	bool
			// AuthIdentity is an (optional) authorization identity (authzid) to
			// use for SASL/PLAIN authentication (if different from User) when
			// an authenticated user is permitted to act as the presented
			// alternative user. See RFC4616 for details.
			AuthIdentity	string
			// User is the authentication identity (authcid) to present for
			// SASL/PLAIN or SASL/SCRAM authentication
			User	string
			// Password for SASL/PLAIN authentication
			Password	string
			// authz id used for SASL/SCRAM authentication
			SCRAMAuthzID	string
			// SCRAMClientGeneratorFunc is a generator of a user provided implementation of a SCRAM
			// client used to perform the SCRAM exchange with the server.
			SCRAMClientGeneratorFunc	func() SCRAMClient
			// TokenProvider is a user-defined callback for generating
			// access tokens for SASL/OAUTHBEARER auth. See the
			// AccessTokenProvider interface docs for proper implementation
			// guidelines.
			TokenProvider	AccessTokenProvider

			GSSAPI	GSSAPIConfig
		}

		// KeepAlive specifies the keep-alive period for an active network connection (defaults to 0).
		// If zero or positive, keep-alives are enabled.
		// If negative, keep-alives are disabled.
		KeepAlive	time.Duration

		// LocalAddr is the local address to use when dialing an
		// address. The address must be of a compatible type for the
		// network being dialed.
		// If nil, a local address is automatically chosen.
		LocalAddr	net.Addr

		Proxy	struct {
			// Whether or not to use proxy when connecting to the broker
			// (defaults to false).
			Enable	bool
			// The proxy dialer to use enabled (defaults to nil).
			Dialer	proxy.Dialer
		}
	}

	// Metadata is the namespace for metadata management properties used by the
	// Client, and shared by the Producer/Consumer.
	Metadata	struct {
		Retry	struct {
			// The total number of times to retry a metadata request when the
			// cluster is in the middle of a leader election (default 3).
			Max	int
			// How long to wait for leader election to occur before retrying
			// (default 250ms). Similar to the JVM's `retry.backoff.ms`.
			Backoff	time.Duration
			// Called to compute backoff time dynamically. Useful for implementing
			// more sophisticated backoff strategies. This takes precedence over
			// `Backoff` if set.
			BackoffFunc	func(retries, maxRetries int) time.Duration
		}
		// How frequently to refresh the cluster metadata in the background.
		// Defaults to 10 minutes. Set to 0 to disable. Similar to
		// `topic.metadata.refresh.interval.ms` in the JVM version.
		RefreshFrequency	time.Duration

		// Whether to maintain a full set of metadata for all topics, or just
		// the minimal set that has been necessary so far. The full set is simpler
		// and usually more convenient, but can take up a substantial amount of
		// memory if you have many topics and partitions. Defaults to true.
		Full	bool

		// How long to wait for a successful metadata response.
		// Disabled by default which means a metadata request against an unreachable
		// cluster (all brokers are unreachable or unresponsive) can take up to
		// `Net.[Dial|Read]Timeout * BrokerCount * (Metadata.Retry.Max + 1) + Metadata.Retry.Backoff * Metadata.Retry.Max`
		// to fail.
		Timeout	time.Duration

		// Whether to allow auto-create topics in metadata refresh. If set to true,
		// the broker may auto-create topics that we requested which do not already exist,
		// if it is configured to do so (`auto.create.topics.enable` is true). Defaults to true.
		AllowAutoTopicCreation	bool
	}

	// Producer is the namespace for configuration related to producing messages,
	// used by the Producer.
	Producer	struct {
		// The maximum permitted size of a message (defaults to 1000000). Should be
		// set equal to or smaller than the broker's `message.max.bytes`.
		MaxMessageBytes	int
		// The level of acknowledgement reliability needed from the broker (defaults
		// to WaitForLocal). Equivalent to the `request.required.acks` setting of the
		// JVM producer.
		RequiredAcks	RequiredAcks
		// The maximum duration the broker will wait the receipt of the number of
		// RequiredAcks (defaults to 10 seconds). This is only relevant when
		// RequiredAcks is set to WaitForAll or a number > 1. Only supports
		// millisecond resolution, nanoseconds will be truncated. Equivalent to
		// the JVM producer's `request.timeout.ms` setting.
		Timeout	time.Duration
		// The type of compression to use on messages (defaults to no compression).
		// Similar to `compression.codec` setting of the JVM producer.
		Compression	CompressionCodec
		// The level of compression to use on messages. The meaning depends
		// on the actual compression type used and defaults to default compression
		// level for the codec.
		CompressionLevel	int
		// Generates partitioners for choosing the partition to send messages to
		// (defaults to hashing the message key). Similar to the `partitioner.class`
		// setting for the JVM producer.
		Partitioner	PartitionerConstructor
		// If enabled, the producer will ensure that exactly one copy of each message is
		// written.
		Idempotent	bool

		// Return specifies what channels will be populated. If they are set to true,
		// you must read from the respective channels to prevent deadlock. If,
		// however, this config is used to create a `SyncProducer`, both must be set
		// to true and you shall not read from the channels since the producer does
		// this internally.
		Return	struct {
			// If enabled, successfully delivered messages will be returned on the
			// Successes channel (default disabled).
			Successes	bool

			// If enabled, messages that failed to deliver will be returned on the
			// Errors channel, including error (default enabled).
			Errors	bool
		}

		// The following config options control how often messages are batched up and
		// sent to the broker. By default, messages are sent as fast as possible, and
		// all messages received while the current batch is in-flight are placed
		// into the subsequent batch.
		Flush	struct {
			// The best-effort number of bytes needed to trigger a flush. Use the
			// global sarama.MaxRequestSize to set a hard upper limit.
			Bytes	int
			// The best-effort number of messages needed to trigger a flush. Use
			// `MaxMessages` to set a hard upper limit.
			Messages	int
			// The best-effort frequency of flushes. Equivalent to
			// `queue.buffering.max.ms` setting of JVM producer.
			Frequency	time.Duration
			// The maximum number of messages the producer will send in a single
			// broker request. Defaults to 0 for unlimited. Similar to
			// `queue.buffering.max.messages` in the JVM producer.
			MaxMessages	int
		}

		Retry	struct {
			// The total number of times to retry sending a message (default 3).
			// Similar to the `message.send.max.retries` setting of the JVM producer.
			Max	int
			// How long to wait for the cluster to settle between retries
			// (default 100ms). Similar to the `retry.backoff.ms` setting of the
			// JVM producer.
			Backoff	time.Duration
			// Called to compute backoff time dynamically. Useful for implementing
			// more sophisticated backoff strategies. This takes precedence over
			// `Backoff` if set.
			BackoffFunc	func(retries, maxRetries int) time.Duration
		}

		// Interceptors to be called when the producer dispatcher reads the
		// message for the first time. Interceptors allows to intercept and
		// possible mutate the message before they are published to Kafka
		// cluster. *ProducerMessage modified by the first interceptor's
		// OnSend() is passed to the second interceptor OnSend(), and so on in
		// the interceptor chain.
		Interceptors	[]ProducerInterceptor
	}

	// Consumer is the namespace for configuration related to consuming messages,
	// used by the Consumer.
	Consumer	struct {

		// Group is the namespace for configuring consumer group.
		Group	struct {
			Session	struct {
				// The timeout used to detect consumer failures when using Kafka's group management facility.
				// The consumer sends periodic heartbeats to indicate its liveness to the broker.
				// If no heartbeats are received by the broker before the expiration of this session timeout,
				// then the broker will remove this consumer from the group and initiate a rebalance.
				// Note that the value must be in the allowable range as configured in the broker configuration
				// by `group.min.session.timeout.ms` and `group.max.session.timeout.ms` (default 10s)
				Timeout time.Duration
			}
			Heartbeat	struct {
				// The expected time between heartbeats to the consumer coordinator when using Kafka's group
				// management facilities. Heartbeats are used to ensure that the consumer's session stays active and
				// to facilitate rebalancing when new consumers join or leave the group.
				// The value must be set lower than Consumer.Group.Session.Timeout, but typically should be set no
				// higher than 1/3 of that value.
				// It can be adjusted even lower to control the expected time for normal rebalances (default 3s)
				Interval time.Duration
			}
			Rebalance	struct {
				// Strategy for allocating topic partitions to members (default BalanceStrategyRange)
				Strategy	BalanceStrategy
				// The maximum allowed time for each worker to join the group once a rebalance has begun.
				// This is basically a limit on the amount of time needed for all tasks to flush any pending
				// data and commit offsets. If the timeout is exceeded, then the worker will be removed from
				// the group, which will cause offset commit failures (default 60s).
				Timeout	time.Duration

				Retry	struct {
					// When a new consumer joins a consumer group the set of consumers attempt to "rebalance"
					// the load to assign partitions to each consumer. If the set of consumers changes while
					// this assignment is taking place the rebalance will fail and retry. This setting controls
					// the maximum number of attempts before giving up (default 4).
					Max	int
					// Backoff time between retries during rebalance (default 2s)
					Backoff	time.Duration
				}
			}
			Member	struct {
				// Custom metadata to include when joining the group. The user data for all joined members
				// can be retrieved by sending a DescribeGroupRequest to the broker that is the
				// coordinator for the group.
				UserData []byte
			}
		}

		Retry	struct {
			// How long to wait after a failing to read from a partition before
			// trying again (default 2s).
			Backoff	time.Duration
			// Called to compute backoff time dynamically. Useful for implementing
			// more sophisticated backoff strategies. This takes precedence over
			// `Backoff` if set.
			BackoffFunc	func(retries int) time.Duration
		}

		// Fetch is the namespace for controlling how many bytes are retrieved by any
		// given request.
		Fetch	struct {
			// The minimum number of message bytes to fetch in a request - the broker
			// will wait until at least this many are available. The default is 1,
			// as 0 causes the consumer to spin when no messages are available.
			// Equivalent to the JVM's `fetch.min.bytes`.
			Min	int32
			// The default number of message bytes to fetch from the broker in each
			// request (default 1MB). This should be larger than the majority of
			// your messages, or else the consumer will spend a lot of time
			// negotiating sizes and not actually consuming. Similar to the JVM's
			// `fetch.message.max.bytes`.
			Default	int32
			// The maximum number of message bytes to fetch from the broker in a
			// single request. Messages larger than this will return
			// ErrMessageTooLarge and will not be consumable, so you must be sure
			// this is at least as large as your largest message. Defaults to 0
			// (no limit). Similar to the JVM's `fetch.message.max.bytes`. The
			// global `sarama.MaxResponseSize` still applies.
			Max	int32
		}
		// The maximum amount of time the broker will wait for Consumer.Fetch.Min
		// bytes to become available before it returns fewer than that anyways. The
		// default is 250ms, since 0 causes the consumer to spin when no events are
		// available. 100-500ms is a reasonable range for most cases. Kafka only
		// supports precision up to milliseconds; nanoseconds will be truncated.
		// Equivalent to the JVM's `fetch.wait.max.ms`.
		MaxWaitTime	time.Duration

		// The maximum amount of time the consumer expects a message takes to
		// process for the user. If writing to the Messages channel takes longer
		// than this, that partition will stop fetching more messages until it
		// can proceed again.
		// Note that, since the Messages channel is buffered, the actual grace time is
		// (MaxProcessingTime * ChannelBufferSize). Defaults to 100ms.
		// If a message is not written to the Messages channel between two ticks
		// of the expiryTicker then a timeout is detected.
		// Using a ticker instead of a timer to detect timeouts should typically
		// result in many fewer calls to Timer functions which may result in a
		// significant performance improvement if many messages are being sent
		// and timeouts are infrequent.
		// The disadvantage of using a ticker instead of a timer is that
		// timeouts will be less accurate. That is, the effective timeout could
		// be between `MaxProcessingTime` and `2 * MaxProcessingTime`. For
		// example, if `MaxProcessingTime` is 100ms then a delay of 180ms
		// between two messages being sent may not be recognized as a timeout.
		MaxProcessingTime	time.Duration

		// Return specifies what channels will be populated. If they are set to true,
		// you must read from them to prevent deadlock.
		Return	struct {
			// If enabled, any errors that occurred while consuming are returned on
			// the Errors channel (default disabled).
			Errors bool
		}

		// Offsets specifies configuration for how and when to commit consumed
		// offsets. This currently requires the manual use of an OffsetManager
		// but will eventually be automated.
		Offsets	struct {
			// Deprecated: CommitInterval exists for historical compatibility
			// and should not be used. Please use Consumer.Offsets.AutoCommit
			CommitInterval	time.Duration

			// AutoCommit specifies configuration for commit messages automatically.
			AutoCommit	struct {
				// Whether or not to auto-commit updated offsets back to the broker.
				// (default enabled).
				Enable	bool

				// How frequently to commit updated offsets. Ineffective unless
				// auto-commit is enabled (default 1s)
				Interval	time.Duration
			}

			// The initial offset to use if no offset was previously committed.
			// Should be OffsetNewest or OffsetOldest. Defaults to OffsetNewest.
			Initial	int64

			// The retention duration for committed offsets. If zero, disabled
			// (in which case the `offsets.retention.minutes` option on the
			// broker will be used).  Kafka only supports precision up to
			// milliseconds; nanoseconds will be truncated. Requires Kafka
			// broker version 0.9.0 or later.
			// (default is 0: disabled).
			Retention	time.Duration

			Retry	struct {
				// The total number of times to retry failing commit
				// requests during OffsetManager shutdown (default 3).
				Max int
			}
		}

		// IsolationLevel support 2 mode:
		// 	- use `ReadUncommitted` (default) to consume and return all messages in message channel
		//	- use `ReadCommitted` to hide messages that are part of an aborted transaction
		IsolationLevel	IsolationLevel

		// Interceptors to be called just before the record is sent to the
		// messages channel. Interceptors allows to intercept and possible
		// mutate the message before they are returned to the client.
		// *ConsumerMessage modified by the first interceptor's OnConsume() is
		// passed to the second interceptor OnConsume(), and so on in the
		// interceptor chain.
		Interceptors	[]ConsumerInterceptor
	}

	// A user-provided string sent with every request to the brokers for logging,
	// debugging, and auditing purposes. Defaults to "sarama", but you should
	// probably set it to something specific to your application.
	ClientID	string
	// A rack identifier for this client. This can be any string value which
	// indicates where this client is physically located.
	// It corresponds with the broker config 'broker.rack'
	RackID	string
	// The number of events to buffer in internal and external channels. This
	// permits the producer and consumer to continue processing some messages
	// in the background while user code is working, greatly improving throughput.
	// Defaults to 256.
	ChannelBufferSize	int
	// ApiVersionsRequest determines whether Sarama should send an
	// ApiVersionsRequest message to each broker as part of its initial
	// connection. This defaults to `true` to match the official Java client
	// and most 3rdparty ones.
	ApiVersionsRequest	bool
	// The version of Kafka that Sarama will assume it is running against.
	// Defaults to the oldest supported stable version. Since Kafka provides
	// backwards-compatibility, setting it to a version older than you have
	// will not break anything, although it may prevent you from using the
	// latest features. Setting it to a version greater than you are actually
	// running may lead to random breakage.
	Version	KafkaVersion
	// The registry to define metrics into.
	// Defaults to a local registry.
	// If you want to disable metrics gathering, set "metrics.UseNilMetrics" to "true"
	// prior to starting Sarama.
	// See Examples on how to use the metrics registry
	MetricRegistry	metrics.Registry
}

// NewConfig returns a new configuration instance with sane defaults.
func NewConfig() *Config {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:454
	_go_fuzz_dep_.CoverTab[100424]++
											c := &Config{}

											c.Admin.Retry.Max = 5
											c.Admin.Retry.Backoff = 100 * time.Millisecond
											c.Admin.Timeout = 3 * time.Second

											c.Net.MaxOpenRequests = 5
											c.Net.DialTimeout = 30 * time.Second
											c.Net.ReadTimeout = 30 * time.Second
											c.Net.WriteTimeout = 30 * time.Second
											c.Net.SASL.Handshake = true
											c.Net.SASL.Version = SASLHandshakeV0

											c.Metadata.Retry.Max = 3
											c.Metadata.Retry.Backoff = 250 * time.Millisecond
											c.Metadata.RefreshFrequency = 10 * time.Minute
											c.Metadata.Full = true
											c.Metadata.AllowAutoTopicCreation = true

											c.Producer.MaxMessageBytes = 1000000
											c.Producer.RequiredAcks = WaitForLocal
											c.Producer.Timeout = 10 * time.Second
											c.Producer.Partitioner = NewHashPartitioner
											c.Producer.Retry.Max = 3
											c.Producer.Retry.Backoff = 100 * time.Millisecond
											c.Producer.Return.Errors = true
											c.Producer.CompressionLevel = CompressionLevelDefault

											c.Consumer.Fetch.Min = 1
											c.Consumer.Fetch.Default = 1024 * 1024
											c.Consumer.Retry.Backoff = 2 * time.Second
											c.Consumer.MaxWaitTime = 250 * time.Millisecond
											c.Consumer.MaxProcessingTime = 100 * time.Millisecond
											c.Consumer.Return.Errors = false
											c.Consumer.Offsets.AutoCommit.Enable = true
											c.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
											c.Consumer.Offsets.Initial = OffsetNewest
											c.Consumer.Offsets.Retry.Max = 3

											c.Consumer.Group.Session.Timeout = 10 * time.Second
											c.Consumer.Group.Heartbeat.Interval = 3 * time.Second
											c.Consumer.Group.Rebalance.Strategy = BalanceStrategyRange
											c.Consumer.Group.Rebalance.Timeout = 60 * time.Second
											c.Consumer.Group.Rebalance.Retry.Max = 4
											c.Consumer.Group.Rebalance.Retry.Backoff = 2 * time.Second

											c.ClientID = defaultClientID
											c.ChannelBufferSize = 256
											c.ApiVersionsRequest = true
											c.Version = DefaultVersion
											c.MetricRegistry = metrics.NewRegistry()

											return c
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:507
	// _ = "end of CoverTab[100424]"
}

// Validate checks a Config instance. It will return a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:510
// ConfigurationError if the specified values don't make sense.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:512
func (c *Config) Validate() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:512
	_go_fuzz_dep_.CoverTab[100425]++

											if !c.Net.TLS.Enable && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:514
		_go_fuzz_dep_.CoverTab[100453]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:514
		return c.Net.TLS.Config != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:514
		// _ = "end of CoverTab[100453]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:514
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:514
		_go_fuzz_dep_.CoverTab[100454]++
												Logger.Println("Net.TLS is disabled but a non-nil configuration was provided.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:515
		// _ = "end of CoverTab[100454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:516
		_go_fuzz_dep_.CoverTab[100455]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:516
		// _ = "end of CoverTab[100455]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:516
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:516
	// _ = "end of CoverTab[100425]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:516
	_go_fuzz_dep_.CoverTab[100426]++
											if !c.Net.SASL.Enable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:517
		_go_fuzz_dep_.CoverTab[100456]++
												if c.Net.SASL.User != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:518
			_go_fuzz_dep_.CoverTab[100458]++
													Logger.Println("Net.SASL is disabled but a non-empty username was provided.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:519
			// _ = "end of CoverTab[100458]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:520
			_go_fuzz_dep_.CoverTab[100459]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:520
			// _ = "end of CoverTab[100459]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:520
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:520
		// _ = "end of CoverTab[100456]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:520
		_go_fuzz_dep_.CoverTab[100457]++
												if c.Net.SASL.Password != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:521
			_go_fuzz_dep_.CoverTab[100460]++
													Logger.Println("Net.SASL is disabled but a non-empty password was provided.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:522
			// _ = "end of CoverTab[100460]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:523
			_go_fuzz_dep_.CoverTab[100461]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:523
			// _ = "end of CoverTab[100461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:523
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:523
		// _ = "end of CoverTab[100457]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:524
		_go_fuzz_dep_.CoverTab[100462]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:524
		// _ = "end of CoverTab[100462]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:524
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:524
	// _ = "end of CoverTab[100426]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:524
	_go_fuzz_dep_.CoverTab[100427]++
											if c.Producer.RequiredAcks > 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:525
		_go_fuzz_dep_.CoverTab[100463]++
												Logger.Println("Producer.RequiredAcks > 1 is deprecated and will raise an exception with kafka >= 0.8.2.0.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:526
		// _ = "end of CoverTab[100463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:527
		_go_fuzz_dep_.CoverTab[100464]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:527
		// _ = "end of CoverTab[100464]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:527
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:527
	// _ = "end of CoverTab[100427]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:527
	_go_fuzz_dep_.CoverTab[100428]++
											if c.Producer.MaxMessageBytes >= int(MaxRequestSize) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:528
		_go_fuzz_dep_.CoverTab[100465]++
												Logger.Println("Producer.MaxMessageBytes must be smaller than MaxRequestSize; it will be ignored.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:529
		// _ = "end of CoverTab[100465]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:530
		_go_fuzz_dep_.CoverTab[100466]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:530
		// _ = "end of CoverTab[100466]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:530
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:530
	// _ = "end of CoverTab[100428]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:530
	_go_fuzz_dep_.CoverTab[100429]++
											if c.Producer.Flush.Bytes >= int(MaxRequestSize) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:531
		_go_fuzz_dep_.CoverTab[100467]++
												Logger.Println("Producer.Flush.Bytes must be smaller than MaxRequestSize; it will be ignored.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:532
		// _ = "end of CoverTab[100467]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:533
		_go_fuzz_dep_.CoverTab[100468]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:533
		// _ = "end of CoverTab[100468]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:533
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:533
	// _ = "end of CoverTab[100429]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:533
	_go_fuzz_dep_.CoverTab[100430]++
											if (c.Producer.Flush.Bytes > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
		_go_fuzz_dep_.CoverTab[100469]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
		return c.Producer.Flush.Messages > 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
		// _ = "end of CoverTab[100469]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
		_go_fuzz_dep_.CoverTab[100470]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
		return c.Producer.Flush.Frequency == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
		// _ = "end of CoverTab[100470]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:534
		_go_fuzz_dep_.CoverTab[100471]++
												Logger.Println("Producer.Flush: Bytes or Messages are set, but Frequency is not; messages may not get flushed.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:535
		// _ = "end of CoverTab[100471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:536
		_go_fuzz_dep_.CoverTab[100472]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:536
		// _ = "end of CoverTab[100472]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:536
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:536
	// _ = "end of CoverTab[100430]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:536
	_go_fuzz_dep_.CoverTab[100431]++
											if c.Producer.Timeout%time.Millisecond != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:537
		_go_fuzz_dep_.CoverTab[100473]++
												Logger.Println("Producer.Timeout only supports millisecond resolution; nanoseconds will be truncated.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:538
		// _ = "end of CoverTab[100473]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:539
		_go_fuzz_dep_.CoverTab[100474]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:539
		// _ = "end of CoverTab[100474]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:539
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:539
	// _ = "end of CoverTab[100431]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:539
	_go_fuzz_dep_.CoverTab[100432]++
											if c.Consumer.MaxWaitTime < 100*time.Millisecond {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:540
		_go_fuzz_dep_.CoverTab[100475]++
												Logger.Println("Consumer.MaxWaitTime is very low, which can cause high CPU and network usage. See documentation for details.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:541
		// _ = "end of CoverTab[100475]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:542
		_go_fuzz_dep_.CoverTab[100476]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:542
		// _ = "end of CoverTab[100476]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:542
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:542
	// _ = "end of CoverTab[100432]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:542
	_go_fuzz_dep_.CoverTab[100433]++
											if c.Consumer.MaxWaitTime%time.Millisecond != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:543
		_go_fuzz_dep_.CoverTab[100477]++
												Logger.Println("Consumer.MaxWaitTime only supports millisecond precision; nanoseconds will be truncated.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:544
		// _ = "end of CoverTab[100477]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:545
		_go_fuzz_dep_.CoverTab[100478]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:545
		// _ = "end of CoverTab[100478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:545
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:545
	// _ = "end of CoverTab[100433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:545
	_go_fuzz_dep_.CoverTab[100434]++
											if c.Consumer.Offsets.Retention%time.Millisecond != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:546
		_go_fuzz_dep_.CoverTab[100479]++
												Logger.Println("Consumer.Offsets.Retention only supports millisecond precision; nanoseconds will be truncated.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:547
		// _ = "end of CoverTab[100479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:548
		_go_fuzz_dep_.CoverTab[100480]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:548
		// _ = "end of CoverTab[100480]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:548
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:548
	// _ = "end of CoverTab[100434]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:548
	_go_fuzz_dep_.CoverTab[100435]++
											if c.Consumer.Group.Session.Timeout%time.Millisecond != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:549
		_go_fuzz_dep_.CoverTab[100481]++
												Logger.Println("Consumer.Group.Session.Timeout only supports millisecond precision; nanoseconds will be truncated.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:550
		// _ = "end of CoverTab[100481]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:551
		_go_fuzz_dep_.CoverTab[100482]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:551
		// _ = "end of CoverTab[100482]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:551
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:551
	// _ = "end of CoverTab[100435]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:551
	_go_fuzz_dep_.CoverTab[100436]++
											if c.Consumer.Group.Heartbeat.Interval%time.Millisecond != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:552
		_go_fuzz_dep_.CoverTab[100483]++
												Logger.Println("Consumer.Group.Heartbeat.Interval only supports millisecond precision; nanoseconds will be truncated.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:553
		// _ = "end of CoverTab[100483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:554
		_go_fuzz_dep_.CoverTab[100484]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:554
		// _ = "end of CoverTab[100484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:554
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:554
	// _ = "end of CoverTab[100436]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:554
	_go_fuzz_dep_.CoverTab[100437]++
											if c.Consumer.Group.Rebalance.Timeout%time.Millisecond != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:555
		_go_fuzz_dep_.CoverTab[100485]++
												Logger.Println("Consumer.Group.Rebalance.Timeout only supports millisecond precision; nanoseconds will be truncated.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:556
		// _ = "end of CoverTab[100485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:557
		_go_fuzz_dep_.CoverTab[100486]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:557
		// _ = "end of CoverTab[100486]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:557
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:557
	// _ = "end of CoverTab[100437]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:557
	_go_fuzz_dep_.CoverTab[100438]++
											if c.ClientID == defaultClientID {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:558
		_go_fuzz_dep_.CoverTab[100487]++
												Logger.Println("ClientID is the default of 'sarama', you should consider setting it to something application-specific.")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:559
		// _ = "end of CoverTab[100487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:560
		_go_fuzz_dep_.CoverTab[100488]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:560
		// _ = "end of CoverTab[100488]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:560
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:560
	// _ = "end of CoverTab[100438]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:560
	_go_fuzz_dep_.CoverTab[100439]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:563
	switch {
	case c.Net.MaxOpenRequests <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:564
		_go_fuzz_dep_.CoverTab[100489]++
												return ConfigurationError("Net.MaxOpenRequests must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:565
		// _ = "end of CoverTab[100489]"
	case c.Net.DialTimeout <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:566
		_go_fuzz_dep_.CoverTab[100490]++
												return ConfigurationError("Net.DialTimeout must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:567
		// _ = "end of CoverTab[100490]"
	case c.Net.ReadTimeout <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:568
		_go_fuzz_dep_.CoverTab[100491]++
												return ConfigurationError("Net.ReadTimeout must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:569
		// _ = "end of CoverTab[100491]"
	case c.Net.WriteTimeout <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:570
		_go_fuzz_dep_.CoverTab[100492]++
												return ConfigurationError("Net.WriteTimeout must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:571
		// _ = "end of CoverTab[100492]"
	case c.Net.SASL.Enable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:572
		_go_fuzz_dep_.CoverTab[100493]++
												if c.Net.SASL.Mechanism == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:573
			_go_fuzz_dep_.CoverTab[100496]++
													c.Net.SASL.Mechanism = SASLTypePlaintext
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:574
			// _ = "end of CoverTab[100496]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:575
			_go_fuzz_dep_.CoverTab[100497]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:575
			// _ = "end of CoverTab[100497]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:575
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:575
		// _ = "end of CoverTab[100493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:575
		_go_fuzz_dep_.CoverTab[100494]++

												switch c.Net.SASL.Mechanism {
		case SASLTypePlaintext:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:578
			_go_fuzz_dep_.CoverTab[100498]++
													if c.Net.SASL.User == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:579
				_go_fuzz_dep_.CoverTab[100510]++
														return ConfigurationError("Net.SASL.User must not be empty when SASL is enabled")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:580
				// _ = "end of CoverTab[100510]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:581
				_go_fuzz_dep_.CoverTab[100511]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:581
				// _ = "end of CoverTab[100511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:581
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:581
			// _ = "end of CoverTab[100498]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:581
			_go_fuzz_dep_.CoverTab[100499]++
													if c.Net.SASL.Password == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:582
				_go_fuzz_dep_.CoverTab[100512]++
														return ConfigurationError("Net.SASL.Password must not be empty when SASL is enabled")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:583
				// _ = "end of CoverTab[100512]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:584
				_go_fuzz_dep_.CoverTab[100513]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:584
				// _ = "end of CoverTab[100513]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:584
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:584
			// _ = "end of CoverTab[100499]"
		case SASLTypeOAuth:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:585
			_go_fuzz_dep_.CoverTab[100500]++
													if c.Net.SASL.TokenProvider == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:586
				_go_fuzz_dep_.CoverTab[100514]++
														return ConfigurationError("An AccessTokenProvider instance must be provided to Net.SASL.TokenProvider")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:587
				// _ = "end of CoverTab[100514]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:588
				_go_fuzz_dep_.CoverTab[100515]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:588
				// _ = "end of CoverTab[100515]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:588
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:588
			// _ = "end of CoverTab[100500]"
		case SASLTypeSCRAMSHA256, SASLTypeSCRAMSHA512:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:589
			_go_fuzz_dep_.CoverTab[100501]++
													if c.Net.SASL.User == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:590
				_go_fuzz_dep_.CoverTab[100516]++
														return ConfigurationError("Net.SASL.User must not be empty when SASL is enabled")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:591
				// _ = "end of CoverTab[100516]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:592
				_go_fuzz_dep_.CoverTab[100517]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:592
				// _ = "end of CoverTab[100517]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:592
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:592
			// _ = "end of CoverTab[100501]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:592
			_go_fuzz_dep_.CoverTab[100502]++
													if c.Net.SASL.Password == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:593
				_go_fuzz_dep_.CoverTab[100518]++
														return ConfigurationError("Net.SASL.Password must not be empty when SASL is enabled")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:594
				// _ = "end of CoverTab[100518]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:595
				_go_fuzz_dep_.CoverTab[100519]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:595
				// _ = "end of CoverTab[100519]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:595
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:595
			// _ = "end of CoverTab[100502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:595
			_go_fuzz_dep_.CoverTab[100503]++
													if c.Net.SASL.SCRAMClientGeneratorFunc == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:596
				_go_fuzz_dep_.CoverTab[100520]++
														return ConfigurationError("A SCRAMClientGeneratorFunc function must be provided to Net.SASL.SCRAMClientGeneratorFunc")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:597
				// _ = "end of CoverTab[100520]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:598
				_go_fuzz_dep_.CoverTab[100521]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:598
				// _ = "end of CoverTab[100521]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:598
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:598
			// _ = "end of CoverTab[100503]"
		case SASLTypeGSSAPI:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:599
			_go_fuzz_dep_.CoverTab[100504]++
													if c.Net.SASL.GSSAPI.ServiceName == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:600
				_go_fuzz_dep_.CoverTab[100522]++
														return ConfigurationError("Net.SASL.GSSAPI.ServiceName must not be empty when GSS-API mechanism is used")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:601
				// _ = "end of CoverTab[100522]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:602
				_go_fuzz_dep_.CoverTab[100523]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:602
				// _ = "end of CoverTab[100523]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:602
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:602
			// _ = "end of CoverTab[100504]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:602
			_go_fuzz_dep_.CoverTab[100505]++

													if c.Net.SASL.GSSAPI.AuthType == KRB5_USER_AUTH {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:604
				_go_fuzz_dep_.CoverTab[100524]++
														if c.Net.SASL.GSSAPI.Password == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:605
					_go_fuzz_dep_.CoverTab[100525]++
															return ConfigurationError("Net.SASL.GSSAPI.Password must not be empty when GSS-API " +
						"mechanism is used and Net.SASL.GSSAPI.AuthType = KRB5_USER_AUTH")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:607
					// _ = "end of CoverTab[100525]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:608
					_go_fuzz_dep_.CoverTab[100526]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:608
					// _ = "end of CoverTab[100526]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:608
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:608
				// _ = "end of CoverTab[100524]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:609
				_go_fuzz_dep_.CoverTab[100527]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:609
				if c.Net.SASL.GSSAPI.AuthType == KRB5_KEYTAB_AUTH {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:609
					_go_fuzz_dep_.CoverTab[100528]++
															if c.Net.SASL.GSSAPI.KeyTabPath == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:610
						_go_fuzz_dep_.CoverTab[100529]++
																return ConfigurationError("Net.SASL.GSSAPI.KeyTabPath must not be empty when GSS-API mechanism is used" +
							" and  Net.SASL.GSSAPI.AuthType = KRB5_KEYTAB_AUTH")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:612
						// _ = "end of CoverTab[100529]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:613
						_go_fuzz_dep_.CoverTab[100530]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:613
						// _ = "end of CoverTab[100530]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:613
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:613
					// _ = "end of CoverTab[100528]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:614
					_go_fuzz_dep_.CoverTab[100531]++
															return ConfigurationError("Net.SASL.GSSAPI.AuthType is invalid. Possible values are KRB5_USER_AUTH and KRB5_KEYTAB_AUTH")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:615
					// _ = "end of CoverTab[100531]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:616
				// _ = "end of CoverTab[100527]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:616
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:616
			// _ = "end of CoverTab[100505]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:616
			_go_fuzz_dep_.CoverTab[100506]++
													if c.Net.SASL.GSSAPI.KerberosConfigPath == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:617
				_go_fuzz_dep_.CoverTab[100532]++
														return ConfigurationError("Net.SASL.GSSAPI.KerberosConfigPath must not be empty when GSS-API mechanism is used")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:618
				// _ = "end of CoverTab[100532]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:619
				_go_fuzz_dep_.CoverTab[100533]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:619
				// _ = "end of CoverTab[100533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:619
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:619
			// _ = "end of CoverTab[100506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:619
			_go_fuzz_dep_.CoverTab[100507]++
													if c.Net.SASL.GSSAPI.Username == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:620
				_go_fuzz_dep_.CoverTab[100534]++
														return ConfigurationError("Net.SASL.GSSAPI.Username must not be empty when GSS-API mechanism is used")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:621
				// _ = "end of CoverTab[100534]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:622
				_go_fuzz_dep_.CoverTab[100535]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:622
				// _ = "end of CoverTab[100535]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:622
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:622
			// _ = "end of CoverTab[100507]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:622
			_go_fuzz_dep_.CoverTab[100508]++
													if c.Net.SASL.GSSAPI.Realm == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:623
				_go_fuzz_dep_.CoverTab[100536]++
														return ConfigurationError("Net.SASL.GSSAPI.Realm must not be empty when GSS-API mechanism is used")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:624
				// _ = "end of CoverTab[100536]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:625
				_go_fuzz_dep_.CoverTab[100537]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:625
				// _ = "end of CoverTab[100537]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:625
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:625
			// _ = "end of CoverTab[100508]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:626
			_go_fuzz_dep_.CoverTab[100509]++
													msg := fmt.Sprintf("The SASL mechanism configuration is invalid. Possible values are `%s`, `%s`, `%s`, `%s` and `%s`",
				SASLTypeOAuth, SASLTypePlaintext, SASLTypeSCRAMSHA256, SASLTypeSCRAMSHA512, SASLTypeGSSAPI)
													return ConfigurationError(msg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:629
			// _ = "end of CoverTab[100509]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:630
		// _ = "end of CoverTab[100494]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:630
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:630
		_go_fuzz_dep_.CoverTab[100495]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:630
		// _ = "end of CoverTab[100495]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:631
	// _ = "end of CoverTab[100439]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:631
	_go_fuzz_dep_.CoverTab[100440]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:634
	switch {
	case c.Admin.Timeout <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:635
		_go_fuzz_dep_.CoverTab[100538]++
												return ConfigurationError("Admin.Timeout must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:636
		// _ = "end of CoverTab[100538]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:636
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:636
		_go_fuzz_dep_.CoverTab[100539]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:636
		// _ = "end of CoverTab[100539]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:637
	// _ = "end of CoverTab[100440]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:637
	_go_fuzz_dep_.CoverTab[100441]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:640
	switch {
	case c.Metadata.Retry.Max < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:641
		_go_fuzz_dep_.CoverTab[100540]++
												return ConfigurationError("Metadata.Retry.Max must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:642
		// _ = "end of CoverTab[100540]"
	case c.Metadata.Retry.Backoff < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:643
		_go_fuzz_dep_.CoverTab[100541]++
												return ConfigurationError("Metadata.Retry.Backoff must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:644
		// _ = "end of CoverTab[100541]"
	case c.Metadata.RefreshFrequency < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:645
		_go_fuzz_dep_.CoverTab[100542]++
												return ConfigurationError("Metadata.RefreshFrequency must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:646
		// _ = "end of CoverTab[100542]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:646
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:646
		_go_fuzz_dep_.CoverTab[100543]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:646
		// _ = "end of CoverTab[100543]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:647
	// _ = "end of CoverTab[100441]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:647
	_go_fuzz_dep_.CoverTab[100442]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:650
	switch {
	case c.Producer.MaxMessageBytes <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:651
		_go_fuzz_dep_.CoverTab[100544]++
												return ConfigurationError("Producer.MaxMessageBytes must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:652
		// _ = "end of CoverTab[100544]"
	case c.Producer.RequiredAcks < -1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:653
		_go_fuzz_dep_.CoverTab[100545]++
												return ConfigurationError("Producer.RequiredAcks must be >= -1")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:654
		// _ = "end of CoverTab[100545]"
	case c.Producer.Timeout <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:655
		_go_fuzz_dep_.CoverTab[100546]++
												return ConfigurationError("Producer.Timeout must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:656
		// _ = "end of CoverTab[100546]"
	case c.Producer.Partitioner == nil:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:657
		_go_fuzz_dep_.CoverTab[100547]++
												return ConfigurationError("Producer.Partitioner must not be nil")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:658
		// _ = "end of CoverTab[100547]"
	case c.Producer.Flush.Bytes < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:659
		_go_fuzz_dep_.CoverTab[100548]++
												return ConfigurationError("Producer.Flush.Bytes must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:660
		// _ = "end of CoverTab[100548]"
	case c.Producer.Flush.Messages < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:661
		_go_fuzz_dep_.CoverTab[100549]++
												return ConfigurationError("Producer.Flush.Messages must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:662
		// _ = "end of CoverTab[100549]"
	case c.Producer.Flush.Frequency < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:663
		_go_fuzz_dep_.CoverTab[100550]++
												return ConfigurationError("Producer.Flush.Frequency must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:664
		// _ = "end of CoverTab[100550]"
	case c.Producer.Flush.MaxMessages < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:665
		_go_fuzz_dep_.CoverTab[100551]++
												return ConfigurationError("Producer.Flush.MaxMessages must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:666
		// _ = "end of CoverTab[100551]"
	case c.Producer.Flush.MaxMessages > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:667
		_go_fuzz_dep_.CoverTab[100556]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:667
		return c.Producer.Flush.MaxMessages < c.Producer.Flush.Messages
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:667
		// _ = "end of CoverTab[100556]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:667
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:667
		_go_fuzz_dep_.CoverTab[100552]++
												return ConfigurationError("Producer.Flush.MaxMessages must be >= Producer.Flush.Messages when set")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:668
		// _ = "end of CoverTab[100552]"
	case c.Producer.Retry.Max < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:669
		_go_fuzz_dep_.CoverTab[100553]++
												return ConfigurationError("Producer.Retry.Max must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:670
		// _ = "end of CoverTab[100553]"
	case c.Producer.Retry.Backoff < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:671
		_go_fuzz_dep_.CoverTab[100554]++
												return ConfigurationError("Producer.Retry.Backoff must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:672
		// _ = "end of CoverTab[100554]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:672
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:672
		_go_fuzz_dep_.CoverTab[100555]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:672
		// _ = "end of CoverTab[100555]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:673
	// _ = "end of CoverTab[100442]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:673
	_go_fuzz_dep_.CoverTab[100443]++

											if c.Producer.Compression == CompressionLZ4 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:675
		_go_fuzz_dep_.CoverTab[100557]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:675
		return !c.Version.IsAtLeast(V0_10_0_0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:675
		// _ = "end of CoverTab[100557]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:675
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:675
		_go_fuzz_dep_.CoverTab[100558]++
												return ConfigurationError("lz4 compression requires Version >= V0_10_0_0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:676
		// _ = "end of CoverTab[100558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:677
		_go_fuzz_dep_.CoverTab[100559]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:677
		// _ = "end of CoverTab[100559]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:677
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:677
	// _ = "end of CoverTab[100443]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:677
	_go_fuzz_dep_.CoverTab[100444]++

											if c.Producer.Compression == CompressionGZIP {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:679
		_go_fuzz_dep_.CoverTab[100560]++
												if c.Producer.CompressionLevel != CompressionLevelDefault {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:680
			_go_fuzz_dep_.CoverTab[100561]++
													if _, err := gzip.NewWriterLevel(io.Discard, c.Producer.CompressionLevel); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:681
				_go_fuzz_dep_.CoverTab[100562]++
														return ConfigurationError(fmt.Sprintf("gzip compression does not work with level %d: %v", c.Producer.CompressionLevel, err))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:682
				// _ = "end of CoverTab[100562]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:683
				_go_fuzz_dep_.CoverTab[100563]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:683
				// _ = "end of CoverTab[100563]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:683
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:683
			// _ = "end of CoverTab[100561]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:684
			_go_fuzz_dep_.CoverTab[100564]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:684
			// _ = "end of CoverTab[100564]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:684
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:684
		// _ = "end of CoverTab[100560]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:685
		_go_fuzz_dep_.CoverTab[100565]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:685
		// _ = "end of CoverTab[100565]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:685
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:685
	// _ = "end of CoverTab[100444]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:685
	_go_fuzz_dep_.CoverTab[100445]++

											if c.Producer.Compression == CompressionZSTD && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:687
		_go_fuzz_dep_.CoverTab[100566]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:687
		return !c.Version.IsAtLeast(V2_1_0_0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:687
		// _ = "end of CoverTab[100566]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:687
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:687
		_go_fuzz_dep_.CoverTab[100567]++
												return ConfigurationError("zstd compression requires Version >= V2_1_0_0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:688
		// _ = "end of CoverTab[100567]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:689
		_go_fuzz_dep_.CoverTab[100568]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:689
		// _ = "end of CoverTab[100568]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:689
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:689
	// _ = "end of CoverTab[100445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:689
	_go_fuzz_dep_.CoverTab[100446]++

											if c.Producer.Idempotent {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:691
		_go_fuzz_dep_.CoverTab[100569]++
												if !c.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:692
			_go_fuzz_dep_.CoverTab[100573]++
													return ConfigurationError("Idempotent producer requires Version >= V0_11_0_0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:693
			// _ = "end of CoverTab[100573]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:694
			_go_fuzz_dep_.CoverTab[100574]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:694
			// _ = "end of CoverTab[100574]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:694
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:694
		// _ = "end of CoverTab[100569]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:694
		_go_fuzz_dep_.CoverTab[100570]++
												if c.Producer.Retry.Max == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:695
			_go_fuzz_dep_.CoverTab[100575]++
													return ConfigurationError("Idempotent producer requires Producer.Retry.Max >= 1")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:696
			// _ = "end of CoverTab[100575]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:697
			_go_fuzz_dep_.CoverTab[100576]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:697
			// _ = "end of CoverTab[100576]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:697
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:697
		// _ = "end of CoverTab[100570]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:697
		_go_fuzz_dep_.CoverTab[100571]++
												if c.Producer.RequiredAcks != WaitForAll {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:698
			_go_fuzz_dep_.CoverTab[100577]++
													return ConfigurationError("Idempotent producer requires Producer.RequiredAcks to be WaitForAll")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:699
			// _ = "end of CoverTab[100577]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:700
			_go_fuzz_dep_.CoverTab[100578]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:700
			// _ = "end of CoverTab[100578]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:700
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:700
		// _ = "end of CoverTab[100571]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:700
		_go_fuzz_dep_.CoverTab[100572]++
												if c.Net.MaxOpenRequests > 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:701
			_go_fuzz_dep_.CoverTab[100579]++
													return ConfigurationError("Idempotent producer requires Net.MaxOpenRequests to be 1")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:702
			// _ = "end of CoverTab[100579]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:703
			_go_fuzz_dep_.CoverTab[100580]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:703
			// _ = "end of CoverTab[100580]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:703
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:703
		// _ = "end of CoverTab[100572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:704
		_go_fuzz_dep_.CoverTab[100581]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:704
		// _ = "end of CoverTab[100581]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:704
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:704
	// _ = "end of CoverTab[100446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:704
	_go_fuzz_dep_.CoverTab[100447]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:707
	switch {
	case c.Consumer.Fetch.Min <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:708
		_go_fuzz_dep_.CoverTab[100582]++
												return ConfigurationError("Consumer.Fetch.Min must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:709
		// _ = "end of CoverTab[100582]"
	case c.Consumer.Fetch.Default <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:710
		_go_fuzz_dep_.CoverTab[100583]++
												return ConfigurationError("Consumer.Fetch.Default must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:711
		// _ = "end of CoverTab[100583]"
	case c.Consumer.Fetch.Max < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:712
		_go_fuzz_dep_.CoverTab[100584]++
												return ConfigurationError("Consumer.Fetch.Max must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:713
		// _ = "end of CoverTab[100584]"
	case c.Consumer.MaxWaitTime < 1*time.Millisecond:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:714
		_go_fuzz_dep_.CoverTab[100585]++
												return ConfigurationError("Consumer.MaxWaitTime must be >= 1ms")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:715
		// _ = "end of CoverTab[100585]"
	case c.Consumer.MaxProcessingTime <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:716
		_go_fuzz_dep_.CoverTab[100586]++
												return ConfigurationError("Consumer.MaxProcessingTime must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:717
		// _ = "end of CoverTab[100586]"
	case c.Consumer.Retry.Backoff < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:718
		_go_fuzz_dep_.CoverTab[100587]++
												return ConfigurationError("Consumer.Retry.Backoff must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:719
		// _ = "end of CoverTab[100587]"
	case c.Consumer.Offsets.AutoCommit.Interval <= 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:720
		_go_fuzz_dep_.CoverTab[100588]++
												return ConfigurationError("Consumer.Offsets.AutoCommit.Interval must be > 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:721
		// _ = "end of CoverTab[100588]"
	case c.Consumer.Offsets.Initial != OffsetOldest && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:722
		_go_fuzz_dep_.CoverTab[100593]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:722
		return c.Consumer.Offsets.Initial != OffsetNewest
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:722
		// _ = "end of CoverTab[100593]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:722
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:722
		_go_fuzz_dep_.CoverTab[100589]++
												return ConfigurationError("Consumer.Offsets.Initial must be OffsetOldest or OffsetNewest")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:723
		// _ = "end of CoverTab[100589]"
	case c.Consumer.Offsets.Retry.Max < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:724
		_go_fuzz_dep_.CoverTab[100590]++
												return ConfigurationError("Consumer.Offsets.Retry.Max must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:725
		// _ = "end of CoverTab[100590]"
	case c.Consumer.IsolationLevel != ReadUncommitted && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:726
		_go_fuzz_dep_.CoverTab[100594]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:726
		return c.Consumer.IsolationLevel != ReadCommitted
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:726
		// _ = "end of CoverTab[100594]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:726
	}():
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:726
		_go_fuzz_dep_.CoverTab[100591]++
												return ConfigurationError("Consumer.IsolationLevel must be ReadUncommitted or ReadCommitted")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:727
		// _ = "end of CoverTab[100591]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:727
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:727
		_go_fuzz_dep_.CoverTab[100592]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:727
		// _ = "end of CoverTab[100592]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:728
	// _ = "end of CoverTab[100447]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:728
	_go_fuzz_dep_.CoverTab[100448]++

											if c.Consumer.Offsets.CommitInterval != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:730
		_go_fuzz_dep_.CoverTab[100595]++
												Logger.Println("Deprecation warning: Consumer.Offsets.CommitInterval exists for historical compatibility" +
			" and should not be used. Please use Consumer.Offsets.AutoCommit, the current value will be ignored")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:732
		// _ = "end of CoverTab[100595]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:733
		_go_fuzz_dep_.CoverTab[100596]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:733
		// _ = "end of CoverTab[100596]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:733
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:733
	// _ = "end of CoverTab[100448]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:733
	_go_fuzz_dep_.CoverTab[100449]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:736
	if c.Consumer.IsolationLevel == ReadCommitted && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:736
		_go_fuzz_dep_.CoverTab[100597]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:736
		return !c.Version.IsAtLeast(V0_11_0_0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:736
		// _ = "end of CoverTab[100597]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:736
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:736
		_go_fuzz_dep_.CoverTab[100598]++
												return ConfigurationError("ReadCommitted requires Version >= V0_11_0_0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:737
		// _ = "end of CoverTab[100598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:738
		_go_fuzz_dep_.CoverTab[100599]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:738
		// _ = "end of CoverTab[100599]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:738
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:738
	// _ = "end of CoverTab[100449]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:738
	_go_fuzz_dep_.CoverTab[100450]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:741
	switch {
	case c.Consumer.Group.Session.Timeout <= 2*time.Millisecond:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:742
		_go_fuzz_dep_.CoverTab[100600]++
												return ConfigurationError("Consumer.Group.Session.Timeout must be >= 2ms")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:743
		// _ = "end of CoverTab[100600]"
	case c.Consumer.Group.Heartbeat.Interval < 1*time.Millisecond:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:744
		_go_fuzz_dep_.CoverTab[100601]++
												return ConfigurationError("Consumer.Group.Heartbeat.Interval must be >= 1ms")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:745
		// _ = "end of CoverTab[100601]"
	case c.Consumer.Group.Heartbeat.Interval >= c.Consumer.Group.Session.Timeout:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:746
		_go_fuzz_dep_.CoverTab[100602]++
												return ConfigurationError("Consumer.Group.Heartbeat.Interval must be < Consumer.Group.Session.Timeout")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:747
		// _ = "end of CoverTab[100602]"
	case c.Consumer.Group.Rebalance.Strategy == nil:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:748
		_go_fuzz_dep_.CoverTab[100603]++
												return ConfigurationError("Consumer.Group.Rebalance.Strategy must not be empty")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:749
		// _ = "end of CoverTab[100603]"
	case c.Consumer.Group.Rebalance.Timeout <= time.Millisecond:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:750
		_go_fuzz_dep_.CoverTab[100604]++
												return ConfigurationError("Consumer.Group.Rebalance.Timeout must be >= 1ms")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:751
		// _ = "end of CoverTab[100604]"
	case c.Consumer.Group.Rebalance.Retry.Max < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:752
		_go_fuzz_dep_.CoverTab[100605]++
												return ConfigurationError("Consumer.Group.Rebalance.Retry.Max must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:753
		// _ = "end of CoverTab[100605]"
	case c.Consumer.Group.Rebalance.Retry.Backoff < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:754
		_go_fuzz_dep_.CoverTab[100606]++
												return ConfigurationError("Consumer.Group.Rebalance.Retry.Backoff must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:755
		// _ = "end of CoverTab[100606]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:755
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:755
		_go_fuzz_dep_.CoverTab[100607]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:755
		// _ = "end of CoverTab[100607]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:756
	// _ = "end of CoverTab[100450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:756
	_go_fuzz_dep_.CoverTab[100451]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:759
	switch {
	case c.ChannelBufferSize < 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:760
		_go_fuzz_dep_.CoverTab[100608]++
												return ConfigurationError("ChannelBufferSize must be >= 0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:761
		// _ = "end of CoverTab[100608]"
	case !validID.MatchString(c.ClientID):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:762
		_go_fuzz_dep_.CoverTab[100609]++
												return ConfigurationError("ClientID is invalid")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:763
		// _ = "end of CoverTab[100609]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:763
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:763
		_go_fuzz_dep_.CoverTab[100610]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:763
		// _ = "end of CoverTab[100610]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:764
	// _ = "end of CoverTab[100451]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:764
	_go_fuzz_dep_.CoverTab[100452]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:766
	// _ = "end of CoverTab[100452]"
}

func (c *Config) getDialer() proxy.Dialer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:769
	_go_fuzz_dep_.CoverTab[100611]++
											if c.Net.Proxy.Enable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:770
		_go_fuzz_dep_.CoverTab[100612]++
												Logger.Printf("using proxy %s", c.Net.Proxy.Dialer)
												return c.Net.Proxy.Dialer
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:772
		// _ = "end of CoverTab[100612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:773
		_go_fuzz_dep_.CoverTab[100613]++
												return &net.Dialer{
			Timeout:	c.Net.DialTimeout,
			KeepAlive:	c.Net.KeepAlive,
			LocalAddr:	c.Net.LocalAddr,
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:778
		// _ = "end of CoverTab[100613]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:779
	// _ = "end of CoverTab[100611]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:780
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config.go:780
var _ = _go_fuzz_dep_.CoverTab
