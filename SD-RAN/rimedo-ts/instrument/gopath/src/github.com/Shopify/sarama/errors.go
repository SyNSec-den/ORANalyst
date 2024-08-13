//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:1
)

import (
	"errors"
	"fmt"
)

// ErrOutOfBrokers is the error returned when the client has run out of brokers to talk to because all of them errored
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:8
// or otherwise failed to respond.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:10
var ErrOutOfBrokers = errors.New("kafka: client has run out of available brokers to talk to (Is your cluster reachable?)")

// ErrBrokerNotFound is the error returned when there's no broker found for the requested ID.
var ErrBrokerNotFound = errors.New("kafka: broker for ID is not found")

// ErrClosedClient is the error returned when a method is called on a client that has been closed.
var ErrClosedClient = errors.New("kafka: tried to use a client that was closed")

// ErrIncompleteResponse is the error returned when the server returns a syntactically valid response, but it does
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:18
// not contain the expected information.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:20
var ErrIncompleteResponse = errors.New("kafka: response did not contain all the expected topic/partition blocks")

// ErrInvalidPartition is the error returned when a partitioner returns an invalid partition index
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:22
// (meaning one outside of the range [0...numPartitions-1]).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:24
var ErrInvalidPartition = errors.New("kafka: partitioner returned an invalid partition index")

// ErrAlreadyConnected is the error returned when calling Open() on a Broker that is already connected or connecting.
var ErrAlreadyConnected = errors.New("kafka: broker connection already initiated")

// ErrNotConnected is the error returned when trying to send or call Close() on a Broker that is not connected.
var ErrNotConnected = errors.New("kafka: broker not connected")

// ErrInsufficientData is returned when decoding and the packet is truncated. This can be expected
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:32
// when requesting messages, since as an optimization the server is allowed to return a partial message at the end
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:32
// of the message set.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:35
var ErrInsufficientData = errors.New("kafka: insufficient data to decode packet, more bytes expected")

// ErrShuttingDown is returned when a producer receives a message during shutdown.
var ErrShuttingDown = errors.New("kafka: message received by producer in process of shutting down")

// ErrMessageTooLarge is returned when the next message to consume is larger than the configured Consumer.Fetch.Max
var ErrMessageTooLarge = errors.New("kafka: message is larger than Consumer.Fetch.Max")

// ErrConsumerOffsetNotAdvanced is returned when a partition consumer didn't advance its offset after parsing
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:43
// a RecordBatch.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:45
var ErrConsumerOffsetNotAdvanced = errors.New("kafka: consumer offset was not advanced after a RecordBatch")

// ErrControllerNotAvailable is returned when server didn't give correct controller id. May be kafka server's version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:47
// is lower than 0.10.0.0.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:49
var ErrControllerNotAvailable = errors.New("kafka: controller is not available")

// ErrNoTopicsToUpdateMetadata is returned when Meta.Full is set to false but no specific topics were found to update
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:51
// the metadata.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:53
var ErrNoTopicsToUpdateMetadata = errors.New("kafka: no specific topics to update metadata")

// ErrUnknownScramMechanism is returned when user tries to AlterUserScramCredentials with unknown SCRAM mechanism
var ErrUnknownScramMechanism = errors.New("kafka: unknown SCRAM mechanism provided")

// PacketEncodingError is returned from a failure while encoding a Kafka packet. This can happen, for example,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:58
// if you try to encode a string over 2^15 characters in length, since Kafka's encoding rules do not permit that.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:60
type PacketEncodingError struct {
	Info string
}

func (err PacketEncodingError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:64
	_go_fuzz_dep_.CoverTab[102745]++
											return fmt.Sprintf("kafka: error encoding packet: %s", err.Info)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:65
	// _ = "end of CoverTab[102745]"
}

// PacketDecodingError is returned when there was an error (other than truncated data) decoding the Kafka broker's response.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:68
// This can be a bad CRC or length field, or any other invalid value.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:70
type PacketDecodingError struct {
	Info string
}

func (err PacketDecodingError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:74
	_go_fuzz_dep_.CoverTab[102746]++
											return fmt.Sprintf("kafka: error decoding packet: %s", err.Info)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:75
	// _ = "end of CoverTab[102746]"
}

// ConfigurationError is the type of error returned from a constructor (e.g. NewClient, or NewConsumer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:78
// when the specified configuration is invalid.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:80
type ConfigurationError string

func (err ConfigurationError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:82
	_go_fuzz_dep_.CoverTab[102747]++
											return "kafka: invalid configuration (" + string(err) + ")"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:83
	// _ = "end of CoverTab[102747]"
}

// KError is the type of error that can be returned directly by the Kafka broker.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:86
// See https://cwiki.apache.org/confluence/display/KAFKA/A+Guide+To+The+Kafka+Protocol#AGuideToTheKafkaProtocol-ErrorCodes
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:88
type KError int16

// MultiError is used to contain multi error.
type MultiError struct {
	Errors *[]error
}

func (mErr MultiError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:95
	_go_fuzz_dep_.CoverTab[102748]++
											errString := ""
											for _, err := range *mErr.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:97
		_go_fuzz_dep_.CoverTab[102750]++
												errString += err.Error() + ","
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:98
		// _ = "end of CoverTab[102750]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:99
	// _ = "end of CoverTab[102748]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:99
	_go_fuzz_dep_.CoverTab[102749]++
											return errString
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:100
	// _ = "end of CoverTab[102749]"
}

func (mErr MultiError) PrettyError() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:103
	_go_fuzz_dep_.CoverTab[102751]++
											errString := ""
											for _, err := range *mErr.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:105
		_go_fuzz_dep_.CoverTab[102753]++
												errString += err.Error() + "\n"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:106
		// _ = "end of CoverTab[102753]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:107
	// _ = "end of CoverTab[102751]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:107
	_go_fuzz_dep_.CoverTab[102752]++
											return errString
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:108
	// _ = "end of CoverTab[102752]"
}

// ErrDeleteRecords is the type of error returned when fail to delete the required records
type ErrDeleteRecords struct {
	MultiError
}

func (err ErrDeleteRecords) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:116
	_go_fuzz_dep_.CoverTab[102754]++
											return "kafka server: failed to delete records " + err.MultiError.Error()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:117
	// _ = "end of CoverTab[102754]"
}

type ErrReassignPartitions struct {
	MultiError
}

func (err ErrReassignPartitions) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:124
	_go_fuzz_dep_.CoverTab[102755]++
											return fmt.Sprintf("failed to reassign partitions for topic: \n%s", err.MultiError.PrettyError())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:125
	// _ = "end of CoverTab[102755]"
}

// Numeric error codes returned by the Kafka server.
const (
	ErrNoError				KError	= 0
	ErrUnknown				KError	= -1
	ErrOffsetOutOfRange			KError	= 1
	ErrInvalidMessage			KError	= 2
	ErrUnknownTopicOrPartition		KError	= 3
	ErrInvalidMessageSize			KError	= 4
	ErrLeaderNotAvailable			KError	= 5
	ErrNotLeaderForPartition		KError	= 6
	ErrRequestTimedOut			KError	= 7
	ErrBrokerNotAvailable			KError	= 8
	ErrReplicaNotAvailable			KError	= 9
	ErrMessageSizeTooLarge			KError	= 10
	ErrStaleControllerEpochCode		KError	= 11
	ErrOffsetMetadataTooLarge		KError	= 12
	ErrNetworkException			KError	= 13
	ErrOffsetsLoadInProgress		KError	= 14
	ErrConsumerCoordinatorNotAvailable	KError	= 15
	ErrNotCoordinatorForConsumer		KError	= 16
	ErrInvalidTopic				KError	= 17
	ErrMessageSetSizeTooLarge		KError	= 18
	ErrNotEnoughReplicas			KError	= 19
	ErrNotEnoughReplicasAfterAppend		KError	= 20
	ErrInvalidRequiredAcks			KError	= 21
	ErrIllegalGeneration			KError	= 22
	ErrInconsistentGroupProtocol		KError	= 23
	ErrInvalidGroupId			KError	= 24
	ErrUnknownMemberId			KError	= 25
	ErrInvalidSessionTimeout		KError	= 26
	ErrRebalanceInProgress			KError	= 27
	ErrInvalidCommitOffsetSize		KError	= 28
	ErrTopicAuthorizationFailed		KError	= 29
	ErrGroupAuthorizationFailed		KError	= 30
	ErrClusterAuthorizationFailed		KError	= 31
	ErrInvalidTimestamp			KError	= 32
	ErrUnsupportedSASLMechanism		KError	= 33
	ErrIllegalSASLState			KError	= 34
	ErrUnsupportedVersion			KError	= 35
	ErrTopicAlreadyExists			KError	= 36
	ErrInvalidPartitions			KError	= 37
	ErrInvalidReplicationFactor		KError	= 38
	ErrInvalidReplicaAssignment		KError	= 39
	ErrInvalidConfig			KError	= 40
	ErrNotController			KError	= 41
	ErrInvalidRequest			KError	= 42
	ErrUnsupportedForMessageFormat		KError	= 43
	ErrPolicyViolation			KError	= 44
	ErrOutOfOrderSequenceNumber		KError	= 45
	ErrDuplicateSequenceNumber		KError	= 46
	ErrInvalidProducerEpoch			KError	= 47
	ErrInvalidTxnState			KError	= 48
	ErrInvalidProducerIDMapping		KError	= 49
	ErrInvalidTransactionTimeout		KError	= 50
	ErrConcurrentTransactions		KError	= 51
	ErrTransactionCoordinatorFenced		KError	= 52
	ErrTransactionalIDAuthorizationFailed	KError	= 53
	ErrSecurityDisabled			KError	= 54
	ErrOperationNotAttempted		KError	= 55
	ErrKafkaStorageError			KError	= 56
	ErrLogDirNotFound			KError	= 57
	ErrSASLAuthenticationFailed		KError	= 58
	ErrUnknownProducerID			KError	= 59
	ErrReassignmentInProgress		KError	= 60
	ErrDelegationTokenAuthDisabled		KError	= 61
	ErrDelegationTokenNotFound		KError	= 62
	ErrDelegationTokenOwnerMismatch		KError	= 63
	ErrDelegationTokenRequestNotAllowed	KError	= 64
	ErrDelegationTokenAuthorizationFailed	KError	= 65
	ErrDelegationTokenExpired		KError	= 66
	ErrInvalidPrincipalType			KError	= 67
	ErrNonEmptyGroup			KError	= 68
	ErrGroupIDNotFound			KError	= 69
	ErrFetchSessionIDNotFound		KError	= 70
	ErrInvalidFetchSessionEpoch		KError	= 71
	ErrListenerNotFound			KError	= 72
	ErrTopicDeletionDisabled		KError	= 73
	ErrFencedLeaderEpoch			KError	= 74
	ErrUnknownLeaderEpoch			KError	= 75
	ErrUnsupportedCompressionType		KError	= 76
	ErrStaleBrokerEpoch			KError	= 77
	ErrOffsetNotAvailable			KError	= 78
	ErrMemberIdRequired			KError	= 79
	ErrPreferredLeaderNotAvailable		KError	= 80
	ErrGroupMaxSizeReached			KError	= 81
	ErrFencedInstancedId			KError	= 82
	ErrEligibleLeadersNotAvailable		KError	= 83
	ErrElectionNotNeeded			KError	= 84
	ErrNoReassignmentInProgress		KError	= 85
	ErrGroupSubscribedToTopic		KError	= 86
	ErrInvalidRecord			KError	= 87
	ErrUnstableOffsetCommit			KError	= 88
)

func (err KError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:222
	_go_fuzz_dep_.CoverTab[102756]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:225
	switch err {
	case ErrNoError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:226
		_go_fuzz_dep_.CoverTab[102758]++
												return "kafka server: Not an error, why are you printing me?"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:227
		// _ = "end of CoverTab[102758]"
	case ErrUnknown:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:228
		_go_fuzz_dep_.CoverTab[102759]++
												return "kafka server: Unexpected (unknown?) server error."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:229
		// _ = "end of CoverTab[102759]"
	case ErrOffsetOutOfRange:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:230
		_go_fuzz_dep_.CoverTab[102760]++
												return "kafka server: The requested offset is outside the range of offsets maintained by the server for the given topic/partition."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:231
		// _ = "end of CoverTab[102760]"
	case ErrInvalidMessage:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:232
		_go_fuzz_dep_.CoverTab[102761]++
												return "kafka server: Message contents does not match its CRC."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:233
		// _ = "end of CoverTab[102761]"
	case ErrUnknownTopicOrPartition:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:234
		_go_fuzz_dep_.CoverTab[102762]++
												return "kafka server: Request was for a topic or partition that does not exist on this broker."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:235
		// _ = "end of CoverTab[102762]"
	case ErrInvalidMessageSize:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:236
		_go_fuzz_dep_.CoverTab[102763]++
												return "kafka server: The message has a negative size."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:237
		// _ = "end of CoverTab[102763]"
	case ErrLeaderNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:238
		_go_fuzz_dep_.CoverTab[102764]++
												return "kafka server: In the middle of a leadership election, there is currently no leader for this partition and hence it is unavailable for writes."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:239
		// _ = "end of CoverTab[102764]"
	case ErrNotLeaderForPartition:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:240
		_go_fuzz_dep_.CoverTab[102765]++
												return "kafka server: Tried to send a message to a replica that is not the leader for some partition. Your metadata is out of date."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:241
		// _ = "end of CoverTab[102765]"
	case ErrRequestTimedOut:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:242
		_go_fuzz_dep_.CoverTab[102766]++
												return "kafka server: Request exceeded the user-specified time limit in the request."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:243
		// _ = "end of CoverTab[102766]"
	case ErrBrokerNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:244
		_go_fuzz_dep_.CoverTab[102767]++
												return "kafka server: Broker not available. Not a client facing error, we should never receive this!!!"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:245
		// _ = "end of CoverTab[102767]"
	case ErrReplicaNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:246
		_go_fuzz_dep_.CoverTab[102768]++
												return "kafka server: Replica information not available, one or more brokers are down."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:247
		// _ = "end of CoverTab[102768]"
	case ErrMessageSizeTooLarge:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:248
		_go_fuzz_dep_.CoverTab[102769]++
												return "kafka server: Message was too large, server rejected it to avoid allocation error."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:249
		// _ = "end of CoverTab[102769]"
	case ErrStaleControllerEpochCode:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:250
		_go_fuzz_dep_.CoverTab[102770]++
												return "kafka server: StaleControllerEpochCode (internal error code for broker-to-broker communication)."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:251
		// _ = "end of CoverTab[102770]"
	case ErrOffsetMetadataTooLarge:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:252
		_go_fuzz_dep_.CoverTab[102771]++
												return "kafka server: Specified a string larger than the configured maximum for offset metadata."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:253
		// _ = "end of CoverTab[102771]"
	case ErrNetworkException:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:254
		_go_fuzz_dep_.CoverTab[102772]++
												return "kafka server: The server disconnected before a response was received."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:255
		// _ = "end of CoverTab[102772]"
	case ErrOffsetsLoadInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:256
		_go_fuzz_dep_.CoverTab[102773]++
												return "kafka server: The broker is still loading offsets after a leader change for that offset's topic partition."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:257
		// _ = "end of CoverTab[102773]"
	case ErrConsumerCoordinatorNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:258
		_go_fuzz_dep_.CoverTab[102774]++
												return "kafka server: Offset's topic has not yet been created."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:259
		// _ = "end of CoverTab[102774]"
	case ErrNotCoordinatorForConsumer:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:260
		_go_fuzz_dep_.CoverTab[102775]++
												return "kafka server: Request was for a consumer group that is not coordinated by this broker."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:261
		// _ = "end of CoverTab[102775]"
	case ErrInvalidTopic:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:262
		_go_fuzz_dep_.CoverTab[102776]++
												return "kafka server: The request attempted to perform an operation on an invalid topic."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:263
		// _ = "end of CoverTab[102776]"
	case ErrMessageSetSizeTooLarge:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:264
		_go_fuzz_dep_.CoverTab[102777]++
												return "kafka server: The request included message batch larger than the configured segment size on the server."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:265
		// _ = "end of CoverTab[102777]"
	case ErrNotEnoughReplicas:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:266
		_go_fuzz_dep_.CoverTab[102778]++
												return "kafka server: Messages are rejected since there are fewer in-sync replicas than required."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:267
		// _ = "end of CoverTab[102778]"
	case ErrNotEnoughReplicasAfterAppend:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:268
		_go_fuzz_dep_.CoverTab[102779]++
												return "kafka server: Messages are written to the log, but to fewer in-sync replicas than required."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:269
		// _ = "end of CoverTab[102779]"
	case ErrInvalidRequiredAcks:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:270
		_go_fuzz_dep_.CoverTab[102780]++
												return "kafka server: The number of required acks is invalid (should be either -1, 0, or 1)."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:271
		// _ = "end of CoverTab[102780]"
	case ErrIllegalGeneration:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:272
		_go_fuzz_dep_.CoverTab[102781]++
												return "kafka server: The provided generation id is not the current generation."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:273
		// _ = "end of CoverTab[102781]"
	case ErrInconsistentGroupProtocol:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:274
		_go_fuzz_dep_.CoverTab[102782]++
												return "kafka server: The provider group protocol type is incompatible with the other members."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:275
		// _ = "end of CoverTab[102782]"
	case ErrInvalidGroupId:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:276
		_go_fuzz_dep_.CoverTab[102783]++
												return "kafka server: The provided group id was empty."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:277
		// _ = "end of CoverTab[102783]"
	case ErrUnknownMemberId:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:278
		_go_fuzz_dep_.CoverTab[102784]++
												return "kafka server: The provided member is not known in the current generation."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:279
		// _ = "end of CoverTab[102784]"
	case ErrInvalidSessionTimeout:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:280
		_go_fuzz_dep_.CoverTab[102785]++
												return "kafka server: The provided session timeout is outside the allowed range."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:281
		// _ = "end of CoverTab[102785]"
	case ErrRebalanceInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:282
		_go_fuzz_dep_.CoverTab[102786]++
												return "kafka server: A rebalance for the group is in progress. Please re-join the group."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:283
		// _ = "end of CoverTab[102786]"
	case ErrInvalidCommitOffsetSize:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:284
		_go_fuzz_dep_.CoverTab[102787]++
												return "kafka server: The provided commit metadata was too large."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:285
		// _ = "end of CoverTab[102787]"
	case ErrTopicAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:286
		_go_fuzz_dep_.CoverTab[102788]++
												return "kafka server: The client is not authorized to access this topic."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:287
		// _ = "end of CoverTab[102788]"
	case ErrGroupAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:288
		_go_fuzz_dep_.CoverTab[102789]++
												return "kafka server: The client is not authorized to access this group."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:289
		// _ = "end of CoverTab[102789]"
	case ErrClusterAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:290
		_go_fuzz_dep_.CoverTab[102790]++
												return "kafka server: The client is not authorized to send this request type."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:291
		// _ = "end of CoverTab[102790]"
	case ErrInvalidTimestamp:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:292
		_go_fuzz_dep_.CoverTab[102791]++
												return "kafka server: The timestamp of the message is out of acceptable range."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:293
		// _ = "end of CoverTab[102791]"
	case ErrUnsupportedSASLMechanism:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:294
		_go_fuzz_dep_.CoverTab[102792]++
												return "kafka server: The broker does not support the requested SASL mechanism."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:295
		// _ = "end of CoverTab[102792]"
	case ErrIllegalSASLState:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:296
		_go_fuzz_dep_.CoverTab[102793]++
												return "kafka server: Request is not valid given the current SASL state."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:297
		// _ = "end of CoverTab[102793]"
	case ErrUnsupportedVersion:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:298
		_go_fuzz_dep_.CoverTab[102794]++
												return "kafka server: The version of API is not supported."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:299
		// _ = "end of CoverTab[102794]"
	case ErrTopicAlreadyExists:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:300
		_go_fuzz_dep_.CoverTab[102795]++
												return "kafka server: Topic with this name already exists."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:301
		// _ = "end of CoverTab[102795]"
	case ErrInvalidPartitions:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:302
		_go_fuzz_dep_.CoverTab[102796]++
												return "kafka server: Number of partitions is invalid."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:303
		// _ = "end of CoverTab[102796]"
	case ErrInvalidReplicationFactor:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:304
		_go_fuzz_dep_.CoverTab[102797]++
												return "kafka server: Replication-factor is invalid."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:305
		// _ = "end of CoverTab[102797]"
	case ErrInvalidReplicaAssignment:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:306
		_go_fuzz_dep_.CoverTab[102798]++
												return "kafka server: Replica assignment is invalid."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:307
		// _ = "end of CoverTab[102798]"
	case ErrInvalidConfig:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:308
		_go_fuzz_dep_.CoverTab[102799]++
												return "kafka server: Configuration is invalid."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:309
		// _ = "end of CoverTab[102799]"
	case ErrNotController:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:310
		_go_fuzz_dep_.CoverTab[102800]++
												return "kafka server: This is not the correct controller for this cluster."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:311
		// _ = "end of CoverTab[102800]"
	case ErrInvalidRequest:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:312
		_go_fuzz_dep_.CoverTab[102801]++
												return "kafka server: This most likely occurs because of a request being malformed by the client library or the message was sent to an incompatible broker. See the broker logs for more details."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:313
		// _ = "end of CoverTab[102801]"
	case ErrUnsupportedForMessageFormat:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:314
		_go_fuzz_dep_.CoverTab[102802]++
												return "kafka server: The requested operation is not supported by the message format version."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:315
		// _ = "end of CoverTab[102802]"
	case ErrPolicyViolation:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:316
		_go_fuzz_dep_.CoverTab[102803]++
												return "kafka server: Request parameters do not satisfy the configured policy."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:317
		// _ = "end of CoverTab[102803]"
	case ErrOutOfOrderSequenceNumber:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:318
		_go_fuzz_dep_.CoverTab[102804]++
												return "kafka server: The broker received an out of order sequence number."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:319
		// _ = "end of CoverTab[102804]"
	case ErrDuplicateSequenceNumber:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:320
		_go_fuzz_dep_.CoverTab[102805]++
												return "kafka server: The broker received a duplicate sequence number."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:321
		// _ = "end of CoverTab[102805]"
	case ErrInvalidProducerEpoch:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:322
		_go_fuzz_dep_.CoverTab[102806]++
												return "kafka server: Producer attempted an operation with an old epoch."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:323
		// _ = "end of CoverTab[102806]"
	case ErrInvalidTxnState:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:324
		_go_fuzz_dep_.CoverTab[102807]++
												return "kafka server: The producer attempted a transactional operation in an invalid state."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:325
		// _ = "end of CoverTab[102807]"
	case ErrInvalidProducerIDMapping:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:326
		_go_fuzz_dep_.CoverTab[102808]++
												return "kafka server: The producer attempted to use a producer id which is not currently assigned to its transactional id."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:327
		// _ = "end of CoverTab[102808]"
	case ErrInvalidTransactionTimeout:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:328
		_go_fuzz_dep_.CoverTab[102809]++
												return "kafka server: The transaction timeout is larger than the maximum value allowed by the broker (as configured by max.transaction.timeout.ms)."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:329
		// _ = "end of CoverTab[102809]"
	case ErrConcurrentTransactions:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:330
		_go_fuzz_dep_.CoverTab[102810]++
												return "kafka server: The producer attempted to update a transaction while another concurrent operation on the same transaction was ongoing."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:331
		// _ = "end of CoverTab[102810]"
	case ErrTransactionCoordinatorFenced:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:332
		_go_fuzz_dep_.CoverTab[102811]++
												return "kafka server: The transaction coordinator sending a WriteTxnMarker is no longer the current coordinator for a given producer."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:333
		// _ = "end of CoverTab[102811]"
	case ErrTransactionalIDAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:334
		_go_fuzz_dep_.CoverTab[102812]++
												return "kafka server: Transactional ID authorization failed."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:335
		// _ = "end of CoverTab[102812]"
	case ErrSecurityDisabled:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:336
		_go_fuzz_dep_.CoverTab[102813]++
												return "kafka server: Security features are disabled."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:337
		// _ = "end of CoverTab[102813]"
	case ErrOperationNotAttempted:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:338
		_go_fuzz_dep_.CoverTab[102814]++
												return "kafka server: The broker did not attempt to execute this operation."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:339
		// _ = "end of CoverTab[102814]"
	case ErrKafkaStorageError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:340
		_go_fuzz_dep_.CoverTab[102815]++
												return "kafka server: Disk error when trying to access log file on the disk."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:341
		// _ = "end of CoverTab[102815]"
	case ErrLogDirNotFound:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:342
		_go_fuzz_dep_.CoverTab[102816]++
												return "kafka server: The specified log directory is not found in the broker config."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:343
		// _ = "end of CoverTab[102816]"
	case ErrSASLAuthenticationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:344
		_go_fuzz_dep_.CoverTab[102817]++
												return "kafka server: SASL Authentication failed."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:345
		// _ = "end of CoverTab[102817]"
	case ErrUnknownProducerID:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:346
		_go_fuzz_dep_.CoverTab[102818]++
												return "kafka server: The broker could not locate the producer metadata associated with the Producer ID."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:347
		// _ = "end of CoverTab[102818]"
	case ErrReassignmentInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:348
		_go_fuzz_dep_.CoverTab[102819]++
												return "kafka server: A partition reassignment is in progress."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:349
		// _ = "end of CoverTab[102819]"
	case ErrDelegationTokenAuthDisabled:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:350
		_go_fuzz_dep_.CoverTab[102820]++
												return "kafka server: Delegation Token feature is not enabled."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:351
		// _ = "end of CoverTab[102820]"
	case ErrDelegationTokenNotFound:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:352
		_go_fuzz_dep_.CoverTab[102821]++
												return "kafka server: Delegation Token is not found on server."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:353
		// _ = "end of CoverTab[102821]"
	case ErrDelegationTokenOwnerMismatch:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:354
		_go_fuzz_dep_.CoverTab[102822]++
												return "kafka server: Specified Principal is not valid Owner/Renewer."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:355
		// _ = "end of CoverTab[102822]"
	case ErrDelegationTokenRequestNotAllowed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:356
		_go_fuzz_dep_.CoverTab[102823]++
												return "kafka server: Delegation Token requests are not allowed on PLAINTEXT/1-way SSL channels and on delegation token authenticated channels."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:357
		// _ = "end of CoverTab[102823]"
	case ErrDelegationTokenAuthorizationFailed:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:358
		_go_fuzz_dep_.CoverTab[102824]++
												return "kafka server: Delegation Token authorization failed."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:359
		// _ = "end of CoverTab[102824]"
	case ErrDelegationTokenExpired:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:360
		_go_fuzz_dep_.CoverTab[102825]++
												return "kafka server: Delegation Token is expired."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:361
		// _ = "end of CoverTab[102825]"
	case ErrInvalidPrincipalType:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:362
		_go_fuzz_dep_.CoverTab[102826]++
												return "kafka server: Supplied principalType is not supported."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:363
		// _ = "end of CoverTab[102826]"
	case ErrNonEmptyGroup:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:364
		_go_fuzz_dep_.CoverTab[102827]++
												return "kafka server: The group is not empty."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:365
		// _ = "end of CoverTab[102827]"
	case ErrGroupIDNotFound:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:366
		_go_fuzz_dep_.CoverTab[102828]++
												return "kafka server: The group id does not exist."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:367
		// _ = "end of CoverTab[102828]"
	case ErrFetchSessionIDNotFound:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:368
		_go_fuzz_dep_.CoverTab[102829]++
												return "kafka server: The fetch session ID was not found."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:369
		// _ = "end of CoverTab[102829]"
	case ErrInvalidFetchSessionEpoch:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:370
		_go_fuzz_dep_.CoverTab[102830]++
												return "kafka server: The fetch session epoch is invalid."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:371
		// _ = "end of CoverTab[102830]"
	case ErrListenerNotFound:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:372
		_go_fuzz_dep_.CoverTab[102831]++
												return "kafka server: There is no listener on the leader broker that matches the listener on which metadata request was processed."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:373
		// _ = "end of CoverTab[102831]"
	case ErrTopicDeletionDisabled:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:374
		_go_fuzz_dep_.CoverTab[102832]++
												return "kafka server: Topic deletion is disabled."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:375
		// _ = "end of CoverTab[102832]"
	case ErrFencedLeaderEpoch:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:376
		_go_fuzz_dep_.CoverTab[102833]++
												return "kafka server: The leader epoch in the request is older than the epoch on the broker."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:377
		// _ = "end of CoverTab[102833]"
	case ErrUnknownLeaderEpoch:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:378
		_go_fuzz_dep_.CoverTab[102834]++
												return "kafka server: The leader epoch in the request is newer than the epoch on the broker."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:379
		// _ = "end of CoverTab[102834]"
	case ErrUnsupportedCompressionType:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:380
		_go_fuzz_dep_.CoverTab[102835]++
												return "kafka server: The requesting client does not support the compression type of given partition."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:381
		// _ = "end of CoverTab[102835]"
	case ErrStaleBrokerEpoch:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:382
		_go_fuzz_dep_.CoverTab[102836]++
												return "kafka server: Broker epoch has changed"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:383
		// _ = "end of CoverTab[102836]"
	case ErrOffsetNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:384
		_go_fuzz_dep_.CoverTab[102837]++
												return "kafka server: The leader high watermark has not caught up from a recent leader election so the offsets cannot be guaranteed to be monotonically increasing"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:385
		// _ = "end of CoverTab[102837]"
	case ErrMemberIdRequired:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:386
		_go_fuzz_dep_.CoverTab[102838]++
												return "kafka server: The group member needs to have a valid member id before actually entering a consumer group"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:387
		// _ = "end of CoverTab[102838]"
	case ErrPreferredLeaderNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:388
		_go_fuzz_dep_.CoverTab[102839]++
												return "kafka server: The preferred leader was not available"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:389
		// _ = "end of CoverTab[102839]"
	case ErrGroupMaxSizeReached:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:390
		_go_fuzz_dep_.CoverTab[102840]++
												return "kafka server: Consumer group The consumer group has reached its max size. already has the configured maximum number of members."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:391
		// _ = "end of CoverTab[102840]"
	case ErrFencedInstancedId:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:392
		_go_fuzz_dep_.CoverTab[102841]++
												return "kafka server: The broker rejected this static consumer since another consumer with the same group.instance.id has registered with a different member.id."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:393
		// _ = "end of CoverTab[102841]"
	case ErrEligibleLeadersNotAvailable:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:394
		_go_fuzz_dep_.CoverTab[102842]++
												return "kafka server: Eligible topic partition leaders are not available."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:395
		// _ = "end of CoverTab[102842]"
	case ErrElectionNotNeeded:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:396
		_go_fuzz_dep_.CoverTab[102843]++
												return "kafka server: Leader election not needed for topic partition."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:397
		// _ = "end of CoverTab[102843]"
	case ErrNoReassignmentInProgress:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:398
		_go_fuzz_dep_.CoverTab[102844]++
												return "kafka server: No partition reassignment is in progress."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:399
		// _ = "end of CoverTab[102844]"
	case ErrGroupSubscribedToTopic:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:400
		_go_fuzz_dep_.CoverTab[102845]++
												return "kafka server: Deleting offsets of a topic is forbidden while the consumer group is actively subscribed to it."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:401
		// _ = "end of CoverTab[102845]"
	case ErrInvalidRecord:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:402
		_go_fuzz_dep_.CoverTab[102846]++
												return "kafka server: This record has failed the validation on broker and hence will be rejected."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:403
		// _ = "end of CoverTab[102846]"
	case ErrUnstableOffsetCommit:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:404
		_go_fuzz_dep_.CoverTab[102847]++
												return "kafka server: There are unstable offsets that need to be cleared."
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:405
		// _ = "end of CoverTab[102847]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:405
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:405
		_go_fuzz_dep_.CoverTab[102848]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:405
		// _ = "end of CoverTab[102848]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:406
	// _ = "end of CoverTab[102756]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:406
	_go_fuzz_dep_.CoverTab[102757]++

											return fmt.Sprintf("Unknown error, how did this happen? Error code = %d", err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:408
	// _ = "end of CoverTab[102757]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:409
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/errors.go:409
var _ = _go_fuzz_dep_.CoverTab
