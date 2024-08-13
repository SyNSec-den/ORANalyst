//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1
)

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// ClusterAdmin is the administrative client for Kafka, which supports managing and inspecting topics,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:12
// brokers, configurations and ACLs. The minimum broker version required is 0.10.0.0.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:12
// Methods with stricter requirements will specify the minimum broker version required.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:12
// You MUST call Close() on a client to avoid leaks
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:16
type ClusterAdmin interface {
	// Creates a new topic. This operation is supported by brokers with version 0.10.1.0 or higher.
	// It may take several seconds after CreateTopic returns success for all the brokers
	// to become aware that the topic has been created. During this time, listTopics
	// may not return information about the new topic.The validateOnly option is supported from version 0.10.2.0.
	CreateTopic(topic string, detail *TopicDetail, validateOnly bool) error

	// List the topics available in the cluster with the default options.
	ListTopics() (map[string]TopicDetail, error)

	// Describe some topics in the cluster.
	DescribeTopics(topics []string) (metadata []*TopicMetadata, err error)

	// Delete a topic. It may take several seconds after the DeleteTopic to returns success
	// and for all the brokers to become aware that the topics are gone.
	// During this time, listTopics  may continue to return information about the deleted topic.
	// If delete.topic.enable is false on the brokers, deleteTopic will mark
	// the topic for deletion, but not actually delete them.
	// This operation is supported by brokers with version 0.10.1.0 or higher.
	DeleteTopic(topic string) error

	// Increase the number of partitions of the topics  according to the corresponding values.
	// If partitions are increased for a topic that has a key, the partition logic or ordering of
	// the messages will be affected. It may take several seconds after this method returns
	// success for all the brokers to become aware that the partitions have been created.
	// During this time, ClusterAdmin#describeTopics may not return information about the
	// new partitions. This operation is supported by brokers with version 1.0.0 or higher.
	CreatePartitions(topic string, count int32, assignment [][]int32, validateOnly bool) error

	// Alter the replica assignment for partitions.
	// This operation is supported by brokers with version 2.4.0.0 or higher.
	AlterPartitionReassignments(topic string, assignment [][]int32) error

	// Provides info on ongoing partitions replica reassignments.
	// This operation is supported by brokers with version 2.4.0.0 or higher.
	ListPartitionReassignments(topics string, partitions []int32) (topicStatus map[string]map[int32]*PartitionReplicaReassignmentsStatus, err error)

	// Delete records whose offset is smaller than the given offset of the corresponding partition.
	// This operation is supported by brokers with version 0.11.0.0 or higher.
	DeleteRecords(topic string, partitionOffsets map[int32]int64) error

	// Get the configuration for the specified resources.
	// The returned configuration includes default values and the Default is true
	// can be used to distinguish them from user supplied values.
	// Config entries where ReadOnly is true cannot be updated.
	// The value of config entries where Sensitive is true is always nil so
	// sensitive information is not disclosed.
	// This operation is supported by brokers with version 0.11.0.0 or higher.
	DescribeConfig(resource ConfigResource) ([]ConfigEntry, error)

	// Update the configuration for the specified resources with the default options.
	// This operation is supported by brokers with version 0.11.0.0 or higher.
	// The resources with their configs (topic is the only resource type with configs
	// that can be updated currently Updates are not transactional so they may succeed
	// for some resources while fail for others. The configs for a particular resource are updated automatically.
	AlterConfig(resourceType ConfigResourceType, name string, entries map[string]*string, validateOnly bool) error

	// IncrementalAlterConfig Incrementally Update the configuration for the specified resources with the default options.
	// This operation is supported by brokers with version 2.3.0.0 or higher.
	// Updates are not transactional so they may succeed for some resources while fail for others.
	// The configs for a particular resource are updated automatically.
	IncrementalAlterConfig(resourceType ConfigResourceType, name string, entries map[string]IncrementalAlterConfigsEntry, validateOnly bool) error

	// Creates access control lists (ACLs) which are bound to specific resources.
	// This operation is not transactional so it may succeed for some ACLs while fail for others.
	// If you attempt to add an ACL that duplicates an existing ACL, no error will be raised, but
	// no changes will be made. This operation is supported by brokers with version 0.11.0.0 or higher.
	CreateACL(resource Resource, acl Acl) error

	// Lists access control lists (ACLs) according to the supplied filter.
	// it may take some time for changes made by createAcls or deleteAcls to be reflected in the output of ListAcls
	// This operation is supported by brokers with version 0.11.0.0 or higher.
	ListAcls(filter AclFilter) ([]ResourceAcls, error)

	// Deletes access control lists (ACLs) according to the supplied filters.
	// This operation is not transactional so it may succeed for some ACLs while fail for others.
	// This operation is supported by brokers with version 0.11.0.0 or higher.
	DeleteACL(filter AclFilter, validateOnly bool) ([]MatchingAcl, error)

	// List the consumer groups available in the cluster.
	ListConsumerGroups() (map[string]string, error)

	// Describe the given consumer groups.
	DescribeConsumerGroups(groups []string) ([]*GroupDescription, error)

	// List the consumer group offsets available in the cluster.
	ListConsumerGroupOffsets(group string, topicPartitions map[string][]int32) (*OffsetFetchResponse, error)

	// Deletes a consumer group offset
	DeleteConsumerGroupOffset(group string, topic string, partition int32) error

	// Delete a consumer group.
	DeleteConsumerGroup(group string) error

	// Get information about the nodes in the cluster
	DescribeCluster() (brokers []*Broker, controllerID int32, err error)

	// Get information about all log directories on the given set of brokers
	DescribeLogDirs(brokers []int32) (map[int32][]DescribeLogDirsResponseDirMetadata, error)

	// Get information about SCRAM users
	DescribeUserScramCredentials(users []string) ([]*DescribeUserScramCredentialsResult, error)

	// Delete SCRAM users
	DeleteUserScramCredentials(delete []AlterUserScramCredentialsDelete) ([]*AlterUserScramCredentialsResult, error)

	// Upsert SCRAM users
	UpsertUserScramCredentials(upsert []AlterUserScramCredentialsUpsert) ([]*AlterUserScramCredentialsResult, error)

	// Get client quota configurations corresponding to the specified filter.
	// This operation is supported by brokers with version 2.6.0.0 or higher.
	DescribeClientQuotas(components []QuotaFilterComponent, strict bool) ([]DescribeClientQuotasEntry, error)

	// Alters client quota configurations with the specified alterations.
	// This operation is supported by brokers with version 2.6.0.0 or higher.
	AlterClientQuotas(entity []QuotaEntityComponent, op ClientQuotasOp, validateOnly bool) error

	// Controller returns the cluster controller broker. It will return a
	// locally cached value if it's available.
	Controller() (*Broker, error)

	// Close shuts down the admin and closes underlying client.
	Close() error
}

type clusterAdmin struct {
	client	Client
	conf	*Config
}

// NewClusterAdmin creates a new ClusterAdmin using the given broker addresses and configuration.
func NewClusterAdmin(addrs []string, conf *Config) (ClusterAdmin, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:147
	_go_fuzz_dep_.CoverTab[97529]++
											client, err := NewClient(addrs, conf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:149
		_go_fuzz_dep_.CoverTab[97531]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:150
		// _ = "end of CoverTab[97531]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:151
		_go_fuzz_dep_.CoverTab[97532]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:151
		// _ = "end of CoverTab[97532]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:151
	// _ = "end of CoverTab[97529]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:151
	_go_fuzz_dep_.CoverTab[97530]++
											return NewClusterAdminFromClient(client)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:152
	// _ = "end of CoverTab[97530]"
}

// NewClusterAdminFromClient creates a new ClusterAdmin using the given client.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:155
// Note that underlying client will also be closed on admin's Close() call.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:157
func NewClusterAdminFromClient(client Client) (ClusterAdmin, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:157
	_go_fuzz_dep_.CoverTab[97533]++

											_, err := client.Controller()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:160
		_go_fuzz_dep_.CoverTab[97535]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:161
		// _ = "end of CoverTab[97535]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:162
		_go_fuzz_dep_.CoverTab[97536]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:162
		// _ = "end of CoverTab[97536]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:162
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:162
	// _ = "end of CoverTab[97533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:162
	_go_fuzz_dep_.CoverTab[97534]++

											ca := &clusterAdmin{
		client:	client,
		conf:	client.Config(),
	}
											return ca, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:168
	// _ = "end of CoverTab[97534]"
}

func (ca *clusterAdmin) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:171
	_go_fuzz_dep_.CoverTab[97537]++
											return ca.client.Close()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:172
	// _ = "end of CoverTab[97537]"
}

func (ca *clusterAdmin) Controller() (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:175
	_go_fuzz_dep_.CoverTab[97538]++
											return ca.client.Controller()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:176
	// _ = "end of CoverTab[97538]"
}

func (ca *clusterAdmin) refreshController() (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:179
	_go_fuzz_dep_.CoverTab[97539]++
											return ca.client.RefreshController()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:180
	// _ = "end of CoverTab[97539]"
}

// isErrNoController returns `true` if the given error type unwraps to an
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:183
// `ErrNotController` response from Kafka
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:185
func isErrNoController(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:185
	_go_fuzz_dep_.CoverTab[97540]++
											switch e := err.(type) {
	case *TopicError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:187
		_go_fuzz_dep_.CoverTab[97542]++
												return e.Err == ErrNotController
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:188
		// _ = "end of CoverTab[97542]"
	case *TopicPartitionError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:189
		_go_fuzz_dep_.CoverTab[97543]++
												return e.Err == ErrNotController
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:190
		// _ = "end of CoverTab[97543]"
	case KError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:191
		_go_fuzz_dep_.CoverTab[97544]++
												return e == ErrNotController
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:192
		// _ = "end of CoverTab[97544]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:193
	// _ = "end of CoverTab[97540]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:193
	_go_fuzz_dep_.CoverTab[97541]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:194
	// _ = "end of CoverTab[97541]"
}

// retryOnError will repeatedly call the given (error-returning) func in the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:197
// case that its response is non-nil and retryable (as determined by the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:197
// provided retryable func) up to the maximum number of tries permitted by
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:197
// the admin client configuration
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:201
func (ca *clusterAdmin) retryOnError(retryable func(error) bool, fn func() error) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:201
	_go_fuzz_dep_.CoverTab[97545]++
											var err error
											for attempt := 0; attempt < ca.conf.Admin.Retry.Max; attempt++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:203
		_go_fuzz_dep_.CoverTab[97547]++
												err = fn()
												if err == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:205
			_go_fuzz_dep_.CoverTab[97549]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:205
			return !retryable(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:205
			// _ = "end of CoverTab[97549]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:205
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:205
			_go_fuzz_dep_.CoverTab[97550]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:206
			// _ = "end of CoverTab[97550]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:207
			_go_fuzz_dep_.CoverTab[97551]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:207
			// _ = "end of CoverTab[97551]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:207
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:207
		// _ = "end of CoverTab[97547]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:207
		_go_fuzz_dep_.CoverTab[97548]++
												Logger.Printf(
			"admin/request retrying after %dms... (%d attempts remaining)\n",
			ca.conf.Admin.Retry.Backoff/time.Millisecond, ca.conf.Admin.Retry.Max-attempt)
												time.Sleep(ca.conf.Admin.Retry.Backoff)
												continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:212
		// _ = "end of CoverTab[97548]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:213
	// _ = "end of CoverTab[97545]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:213
	_go_fuzz_dep_.CoverTab[97546]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:214
	// _ = "end of CoverTab[97546]"
}

func (ca *clusterAdmin) CreateTopic(topic string, detail *TopicDetail, validateOnly bool) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:217
	_go_fuzz_dep_.CoverTab[97552]++
											if topic == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:218
		_go_fuzz_dep_.CoverTab[97557]++
												return ErrInvalidTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:219
		// _ = "end of CoverTab[97557]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:220
		_go_fuzz_dep_.CoverTab[97558]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:220
		// _ = "end of CoverTab[97558]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:220
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:220
	// _ = "end of CoverTab[97552]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:220
	_go_fuzz_dep_.CoverTab[97553]++

											if detail == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:222
		_go_fuzz_dep_.CoverTab[97559]++
												return errors.New("you must specify topic details")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:223
		// _ = "end of CoverTab[97559]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:224
		_go_fuzz_dep_.CoverTab[97560]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:224
		// _ = "end of CoverTab[97560]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:224
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:224
	// _ = "end of CoverTab[97553]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:224
	_go_fuzz_dep_.CoverTab[97554]++

											topicDetails := make(map[string]*TopicDetail)
											topicDetails[topic] = detail

											request := &CreateTopicsRequest{
		TopicDetails:	topicDetails,
		ValidateOnly:	validateOnly,
		Timeout:	ca.conf.Admin.Timeout,
	}

	if ca.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:235
		_go_fuzz_dep_.CoverTab[97561]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:236
		// _ = "end of CoverTab[97561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:237
		_go_fuzz_dep_.CoverTab[97562]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:237
		// _ = "end of CoverTab[97562]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:237
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:237
	// _ = "end of CoverTab[97554]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:237
	_go_fuzz_dep_.CoverTab[97555]++
											if ca.conf.Version.IsAtLeast(V1_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:238
		_go_fuzz_dep_.CoverTab[97563]++
												request.Version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:239
		// _ = "end of CoverTab[97563]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:240
		_go_fuzz_dep_.CoverTab[97564]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:240
		// _ = "end of CoverTab[97564]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:240
	// _ = "end of CoverTab[97555]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:240
	_go_fuzz_dep_.CoverTab[97556]++

											return ca.retryOnError(isErrNoController, func() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:242
		_go_fuzz_dep_.CoverTab[97565]++
												b, err := ca.Controller()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:244
			_go_fuzz_dep_.CoverTab[97570]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:245
			// _ = "end of CoverTab[97570]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:246
			_go_fuzz_dep_.CoverTab[97571]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:246
			// _ = "end of CoverTab[97571]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:246
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:246
		// _ = "end of CoverTab[97565]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:246
		_go_fuzz_dep_.CoverTab[97566]++

												rsp, err := b.CreateTopics(request)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:249
			_go_fuzz_dep_.CoverTab[97572]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:250
			// _ = "end of CoverTab[97572]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:251
			_go_fuzz_dep_.CoverTab[97573]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:251
			// _ = "end of CoverTab[97573]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:251
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:251
		// _ = "end of CoverTab[97566]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:251
		_go_fuzz_dep_.CoverTab[97567]++

												topicErr, ok := rsp.TopicErrors[topic]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:254
			_go_fuzz_dep_.CoverTab[97574]++
													return ErrIncompleteResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:255
			// _ = "end of CoverTab[97574]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:256
			_go_fuzz_dep_.CoverTab[97575]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:256
			// _ = "end of CoverTab[97575]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:256
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:256
		// _ = "end of CoverTab[97567]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:256
		_go_fuzz_dep_.CoverTab[97568]++

												if topicErr.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:258
			_go_fuzz_dep_.CoverTab[97576]++
													if topicErr.Err == ErrNotController {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:259
				_go_fuzz_dep_.CoverTab[97578]++
														_, _ = ca.refreshController()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:260
				// _ = "end of CoverTab[97578]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:261
				_go_fuzz_dep_.CoverTab[97579]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:261
				// _ = "end of CoverTab[97579]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:261
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:261
			// _ = "end of CoverTab[97576]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:261
			_go_fuzz_dep_.CoverTab[97577]++
													return topicErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:262
			// _ = "end of CoverTab[97577]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:263
			_go_fuzz_dep_.CoverTab[97580]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:263
			// _ = "end of CoverTab[97580]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:263
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:263
		// _ = "end of CoverTab[97568]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:263
		_go_fuzz_dep_.CoverTab[97569]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:265
		// _ = "end of CoverTab[97569]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:266
	// _ = "end of CoverTab[97556]"
}

func (ca *clusterAdmin) DescribeTopics(topics []string) (metadata []*TopicMetadata, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:269
	_go_fuzz_dep_.CoverTab[97581]++
											controller, err := ca.Controller()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:271
		_go_fuzz_dep_.CoverTab[97585]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:272
		// _ = "end of CoverTab[97585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:273
		_go_fuzz_dep_.CoverTab[97586]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:273
		// _ = "end of CoverTab[97586]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:273
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:273
	// _ = "end of CoverTab[97581]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:273
	_go_fuzz_dep_.CoverTab[97582]++

											request := &MetadataRequest{
		Topics:			topics,
		AllowAutoTopicCreation:	false,
	}

	if ca.conf.Version.IsAtLeast(V1_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:280
		_go_fuzz_dep_.CoverTab[97587]++
												request.Version = 5
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:281
		// _ = "end of CoverTab[97587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:282
		_go_fuzz_dep_.CoverTab[97588]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:282
		if ca.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:282
			_go_fuzz_dep_.CoverTab[97589]++
													request.Version = 4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:283
			// _ = "end of CoverTab[97589]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:284
			_go_fuzz_dep_.CoverTab[97590]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:284
			// _ = "end of CoverTab[97590]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:284
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:284
		// _ = "end of CoverTab[97588]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:284
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:284
	// _ = "end of CoverTab[97582]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:284
	_go_fuzz_dep_.CoverTab[97583]++

											response, err := controller.GetMetadata(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:287
		_go_fuzz_dep_.CoverTab[97591]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:288
		// _ = "end of CoverTab[97591]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:289
		_go_fuzz_dep_.CoverTab[97592]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:289
		// _ = "end of CoverTab[97592]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:289
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:289
	// _ = "end of CoverTab[97583]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:289
	_go_fuzz_dep_.CoverTab[97584]++
											return response.Topics, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:290
	// _ = "end of CoverTab[97584]"
}

func (ca *clusterAdmin) DescribeCluster() (brokers []*Broker, controllerID int32, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:293
	_go_fuzz_dep_.CoverTab[97593]++
											controller, err := ca.Controller()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:295
		_go_fuzz_dep_.CoverTab[97597]++
												return nil, int32(0), err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:296
		// _ = "end of CoverTab[97597]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:297
		_go_fuzz_dep_.CoverTab[97598]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:297
		// _ = "end of CoverTab[97598]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:297
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:297
	// _ = "end of CoverTab[97593]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:297
	_go_fuzz_dep_.CoverTab[97594]++

											request := &MetadataRequest{
		Topics: []string{},
	}

	if ca.conf.Version.IsAtLeast(V0_10_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:303
		_go_fuzz_dep_.CoverTab[97599]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:304
		// _ = "end of CoverTab[97599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:305
		_go_fuzz_dep_.CoverTab[97600]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:305
		// _ = "end of CoverTab[97600]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:305
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:305
	// _ = "end of CoverTab[97594]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:305
	_go_fuzz_dep_.CoverTab[97595]++

											response, err := controller.GetMetadata(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:308
		_go_fuzz_dep_.CoverTab[97601]++
												return nil, int32(0), err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:309
		// _ = "end of CoverTab[97601]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:310
		_go_fuzz_dep_.CoverTab[97602]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:310
		// _ = "end of CoverTab[97602]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:310
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:310
	// _ = "end of CoverTab[97595]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:310
	_go_fuzz_dep_.CoverTab[97596]++

											return response.Brokers, response.ControllerID, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:312
	// _ = "end of CoverTab[97596]"
}

func (ca *clusterAdmin) findBroker(id int32) (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:315
	_go_fuzz_dep_.CoverTab[97603]++
											brokers := ca.client.Brokers()
											for _, b := range brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:317
		_go_fuzz_dep_.CoverTab[97605]++
												if b.ID() == id {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:318
			_go_fuzz_dep_.CoverTab[97606]++
													return b, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:319
			// _ = "end of CoverTab[97606]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:320
			_go_fuzz_dep_.CoverTab[97607]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:320
			// _ = "end of CoverTab[97607]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:320
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:320
		// _ = "end of CoverTab[97605]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:321
	// _ = "end of CoverTab[97603]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:321
	_go_fuzz_dep_.CoverTab[97604]++
											return nil, fmt.Errorf("could not find broker id %d", id)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:322
	// _ = "end of CoverTab[97604]"
}

func (ca *clusterAdmin) findAnyBroker() (*Broker, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:325
	_go_fuzz_dep_.CoverTab[97608]++
											brokers := ca.client.Brokers()
											if len(brokers) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:327
		_go_fuzz_dep_.CoverTab[97610]++
												index := rand.Intn(len(brokers))
												return brokers[index], nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:329
		// _ = "end of CoverTab[97610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:330
		_go_fuzz_dep_.CoverTab[97611]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:330
		// _ = "end of CoverTab[97611]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:330
	// _ = "end of CoverTab[97608]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:330
	_go_fuzz_dep_.CoverTab[97609]++
											return nil, errors.New("no available broker")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:331
	// _ = "end of CoverTab[97609]"
}

func (ca *clusterAdmin) ListTopics() (map[string]TopicDetail, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:334
	_go_fuzz_dep_.CoverTab[97612]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:341
	b, err := ca.findAnyBroker()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:342
		_go_fuzz_dep_.CoverTab[97620]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:343
		// _ = "end of CoverTab[97620]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:344
		_go_fuzz_dep_.CoverTab[97621]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:344
		// _ = "end of CoverTab[97621]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:344
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:344
	// _ = "end of CoverTab[97612]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:344
	_go_fuzz_dep_.CoverTab[97613]++
											_ = b.Open(ca.client.Config())

											metadataReq := &MetadataRequest{}
											metadataResp, err := b.GetMetadata(metadataReq)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:349
		_go_fuzz_dep_.CoverTab[97622]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:350
		// _ = "end of CoverTab[97622]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:351
		_go_fuzz_dep_.CoverTab[97623]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:351
		// _ = "end of CoverTab[97623]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:351
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:351
	// _ = "end of CoverTab[97613]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:351
	_go_fuzz_dep_.CoverTab[97614]++

											topicsDetailsMap := make(map[string]TopicDetail)

											var describeConfigsResources []*ConfigResource

											for _, topic := range metadataResp.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:357
		_go_fuzz_dep_.CoverTab[97624]++
												topicDetails := TopicDetail{
			NumPartitions: int32(len(topic.Partitions)),
		}
		if len(topic.Partitions) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:361
			_go_fuzz_dep_.CoverTab[97626]++
													topicDetails.ReplicaAssignment = map[int32][]int32{}
													for _, partition := range topic.Partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:363
				_go_fuzz_dep_.CoverTab[97628]++
														topicDetails.ReplicaAssignment[partition.ID] = partition.Replicas
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:364
				// _ = "end of CoverTab[97628]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:365
			// _ = "end of CoverTab[97626]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:365
			_go_fuzz_dep_.CoverTab[97627]++
													topicDetails.ReplicationFactor = int16(len(topic.Partitions[0].Replicas))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:366
			// _ = "end of CoverTab[97627]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:367
			_go_fuzz_dep_.CoverTab[97629]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:367
			// _ = "end of CoverTab[97629]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:367
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:367
		// _ = "end of CoverTab[97624]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:367
		_go_fuzz_dep_.CoverTab[97625]++
												topicsDetailsMap[topic.Name] = topicDetails

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:371
		topicResource := ConfigResource{
			Type:	TopicResource,
			Name:	topic.Name,
		}
												describeConfigsResources = append(describeConfigsResources, &topicResource)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:375
		// _ = "end of CoverTab[97625]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:376
	// _ = "end of CoverTab[97614]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:376
	_go_fuzz_dep_.CoverTab[97615]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:379
	describeConfigsReq := &DescribeConfigsRequest{
		Resources: describeConfigsResources,
	}

	if ca.conf.Version.IsAtLeast(V1_1_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:383
		_go_fuzz_dep_.CoverTab[97630]++
												describeConfigsReq.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:384
		// _ = "end of CoverTab[97630]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:385
		_go_fuzz_dep_.CoverTab[97631]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:385
		// _ = "end of CoverTab[97631]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:385
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:385
	// _ = "end of CoverTab[97615]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:385
	_go_fuzz_dep_.CoverTab[97616]++

											if ca.conf.Version.IsAtLeast(V2_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:387
		_go_fuzz_dep_.CoverTab[97632]++
												describeConfigsReq.Version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:388
		// _ = "end of CoverTab[97632]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:389
		_go_fuzz_dep_.CoverTab[97633]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:389
		// _ = "end of CoverTab[97633]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:389
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:389
	// _ = "end of CoverTab[97616]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:389
	_go_fuzz_dep_.CoverTab[97617]++

											describeConfigsResp, err := b.DescribeConfigs(describeConfigsReq)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:392
		_go_fuzz_dep_.CoverTab[97634]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:393
		// _ = "end of CoverTab[97634]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:394
		_go_fuzz_dep_.CoverTab[97635]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:394
		// _ = "end of CoverTab[97635]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:394
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:394
	// _ = "end of CoverTab[97617]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:394
	_go_fuzz_dep_.CoverTab[97618]++

											for _, resource := range describeConfigsResp.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:396
		_go_fuzz_dep_.CoverTab[97636]++
												topicDetails := topicsDetailsMap[resource.Name]
												topicDetails.ConfigEntries = make(map[string]*string)

												for _, entry := range resource.Configs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:400
			_go_fuzz_dep_.CoverTab[97638]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:403
			if entry.Default || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:403
				_go_fuzz_dep_.CoverTab[97640]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:403
				return entry.Sensitive
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:403
				// _ = "end of CoverTab[97640]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:403
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:403
				_go_fuzz_dep_.CoverTab[97641]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:404
				// _ = "end of CoverTab[97641]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:405
				_go_fuzz_dep_.CoverTab[97642]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:405
				// _ = "end of CoverTab[97642]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:405
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:405
			// _ = "end of CoverTab[97638]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:405
			_go_fuzz_dep_.CoverTab[97639]++
													topicDetails.ConfigEntries[entry.Name] = &entry.Value
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:406
			// _ = "end of CoverTab[97639]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:407
		// _ = "end of CoverTab[97636]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:407
		_go_fuzz_dep_.CoverTab[97637]++

												topicsDetailsMap[resource.Name] = topicDetails
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:409
		// _ = "end of CoverTab[97637]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:410
	// _ = "end of CoverTab[97618]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:410
	_go_fuzz_dep_.CoverTab[97619]++

											return topicsDetailsMap, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:412
	// _ = "end of CoverTab[97619]"
}

func (ca *clusterAdmin) DeleteTopic(topic string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:415
	_go_fuzz_dep_.CoverTab[97643]++
											if topic == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:416
		_go_fuzz_dep_.CoverTab[97646]++
												return ErrInvalidTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:417
		// _ = "end of CoverTab[97646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:418
		_go_fuzz_dep_.CoverTab[97647]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:418
		// _ = "end of CoverTab[97647]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:418
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:418
	// _ = "end of CoverTab[97643]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:418
	_go_fuzz_dep_.CoverTab[97644]++

											request := &DeleteTopicsRequest{
		Topics:		[]string{topic},
		Timeout:	ca.conf.Admin.Timeout,
	}

	if ca.conf.Version.IsAtLeast(V0_11_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:425
		_go_fuzz_dep_.CoverTab[97648]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:426
		// _ = "end of CoverTab[97648]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:427
		_go_fuzz_dep_.CoverTab[97649]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:427
		// _ = "end of CoverTab[97649]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:427
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:427
	// _ = "end of CoverTab[97644]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:427
	_go_fuzz_dep_.CoverTab[97645]++

											return ca.retryOnError(isErrNoController, func() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:429
		_go_fuzz_dep_.CoverTab[97650]++
												b, err := ca.Controller()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:431
			_go_fuzz_dep_.CoverTab[97655]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:432
			// _ = "end of CoverTab[97655]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:433
			_go_fuzz_dep_.CoverTab[97656]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:433
			// _ = "end of CoverTab[97656]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:433
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:433
		// _ = "end of CoverTab[97650]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:433
		_go_fuzz_dep_.CoverTab[97651]++

												rsp, err := b.DeleteTopics(request)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:436
			_go_fuzz_dep_.CoverTab[97657]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:437
			// _ = "end of CoverTab[97657]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:438
			_go_fuzz_dep_.CoverTab[97658]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:438
			// _ = "end of CoverTab[97658]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:438
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:438
		// _ = "end of CoverTab[97651]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:438
		_go_fuzz_dep_.CoverTab[97652]++

												topicErr, ok := rsp.TopicErrorCodes[topic]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:441
			_go_fuzz_dep_.CoverTab[97659]++
													return ErrIncompleteResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:442
			// _ = "end of CoverTab[97659]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:443
			_go_fuzz_dep_.CoverTab[97660]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:443
			// _ = "end of CoverTab[97660]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:443
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:443
		// _ = "end of CoverTab[97652]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:443
		_go_fuzz_dep_.CoverTab[97653]++

												if topicErr != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:445
			_go_fuzz_dep_.CoverTab[97661]++
													if topicErr == ErrNotController {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:446
				_go_fuzz_dep_.CoverTab[97663]++
														_, _ = ca.refreshController()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:447
				// _ = "end of CoverTab[97663]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:448
				_go_fuzz_dep_.CoverTab[97664]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:448
				// _ = "end of CoverTab[97664]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:448
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:448
			// _ = "end of CoverTab[97661]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:448
			_go_fuzz_dep_.CoverTab[97662]++
													return topicErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:449
			// _ = "end of CoverTab[97662]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:450
			_go_fuzz_dep_.CoverTab[97665]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:450
			// _ = "end of CoverTab[97665]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:450
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:450
		// _ = "end of CoverTab[97653]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:450
		_go_fuzz_dep_.CoverTab[97654]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:452
		// _ = "end of CoverTab[97654]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:453
	// _ = "end of CoverTab[97645]"
}

func (ca *clusterAdmin) CreatePartitions(topic string, count int32, assignment [][]int32, validateOnly bool) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:456
	_go_fuzz_dep_.CoverTab[97666]++
											if topic == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:457
		_go_fuzz_dep_.CoverTab[97668]++
												return ErrInvalidTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:458
		// _ = "end of CoverTab[97668]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:459
		_go_fuzz_dep_.CoverTab[97669]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:459
		// _ = "end of CoverTab[97669]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:459
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:459
	// _ = "end of CoverTab[97666]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:459
	_go_fuzz_dep_.CoverTab[97667]++

											topicPartitions := make(map[string]*TopicPartition)
											topicPartitions[topic] = &TopicPartition{Count: count, Assignment: assignment}

											request := &CreatePartitionsRequest{
		TopicPartitions:	topicPartitions,
		Timeout:		ca.conf.Admin.Timeout,
		ValidateOnly:		validateOnly,
	}

	return ca.retryOnError(isErrNoController, func() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:470
		_go_fuzz_dep_.CoverTab[97670]++
												b, err := ca.Controller()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:472
			_go_fuzz_dep_.CoverTab[97675]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:473
			// _ = "end of CoverTab[97675]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:474
			_go_fuzz_dep_.CoverTab[97676]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:474
			// _ = "end of CoverTab[97676]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:474
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:474
		// _ = "end of CoverTab[97670]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:474
		_go_fuzz_dep_.CoverTab[97671]++

												rsp, err := b.CreatePartitions(request)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:477
			_go_fuzz_dep_.CoverTab[97677]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:478
			// _ = "end of CoverTab[97677]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:479
			_go_fuzz_dep_.CoverTab[97678]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:479
			// _ = "end of CoverTab[97678]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:479
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:479
		// _ = "end of CoverTab[97671]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:479
		_go_fuzz_dep_.CoverTab[97672]++

												topicErr, ok := rsp.TopicPartitionErrors[topic]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:482
			_go_fuzz_dep_.CoverTab[97679]++
													return ErrIncompleteResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:483
			// _ = "end of CoverTab[97679]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:484
			_go_fuzz_dep_.CoverTab[97680]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:484
			// _ = "end of CoverTab[97680]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:484
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:484
		// _ = "end of CoverTab[97672]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:484
		_go_fuzz_dep_.CoverTab[97673]++

												if topicErr.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:486
			_go_fuzz_dep_.CoverTab[97681]++
													if topicErr.Err == ErrNotController {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:487
				_go_fuzz_dep_.CoverTab[97683]++
														_, _ = ca.refreshController()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:488
				// _ = "end of CoverTab[97683]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:489
				_go_fuzz_dep_.CoverTab[97684]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:489
				// _ = "end of CoverTab[97684]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:489
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:489
			// _ = "end of CoverTab[97681]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:489
			_go_fuzz_dep_.CoverTab[97682]++
													return topicErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:490
			// _ = "end of CoverTab[97682]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:491
			_go_fuzz_dep_.CoverTab[97685]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:491
			// _ = "end of CoverTab[97685]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:491
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:491
		// _ = "end of CoverTab[97673]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:491
		_go_fuzz_dep_.CoverTab[97674]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:493
		// _ = "end of CoverTab[97674]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:494
	// _ = "end of CoverTab[97667]"
}

func (ca *clusterAdmin) AlterPartitionReassignments(topic string, assignment [][]int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:497
	_go_fuzz_dep_.CoverTab[97686]++
											if topic == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:498
		_go_fuzz_dep_.CoverTab[97689]++
												return ErrInvalidTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:499
		// _ = "end of CoverTab[97689]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:500
		_go_fuzz_dep_.CoverTab[97690]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:500
		// _ = "end of CoverTab[97690]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:500
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:500
	// _ = "end of CoverTab[97686]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:500
	_go_fuzz_dep_.CoverTab[97687]++

											request := &AlterPartitionReassignmentsRequest{
		TimeoutMs:	int32(60000),
		Version:	int16(0),
	}

	for i := 0; i < len(assignment); i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:507
		_go_fuzz_dep_.CoverTab[97691]++
												request.AddBlock(topic, int32(i), assignment[i])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:508
		// _ = "end of CoverTab[97691]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:509
	// _ = "end of CoverTab[97687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:509
	_go_fuzz_dep_.CoverTab[97688]++

											return ca.retryOnError(isErrNoController, func() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:511
		_go_fuzz_dep_.CoverTab[97692]++
												b, err := ca.Controller()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:513
			_go_fuzz_dep_.CoverTab[97696]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:514
			// _ = "end of CoverTab[97696]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:515
			_go_fuzz_dep_.CoverTab[97697]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:515
			// _ = "end of CoverTab[97697]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:515
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:515
		// _ = "end of CoverTab[97692]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:515
		_go_fuzz_dep_.CoverTab[97693]++

												errs := make([]error, 0)

												rsp, err := b.AlterPartitionReassignments(request)

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:521
			_go_fuzz_dep_.CoverTab[97698]++
													errs = append(errs, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:522
			// _ = "end of CoverTab[97698]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:523
			_go_fuzz_dep_.CoverTab[97699]++
													if rsp.ErrorCode > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:524
				_go_fuzz_dep_.CoverTab[97701]++
														errs = append(errs, errors.New(rsp.ErrorCode.Error()))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:525
				// _ = "end of CoverTab[97701]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:526
				_go_fuzz_dep_.CoverTab[97702]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:526
				// _ = "end of CoverTab[97702]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:526
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:526
			// _ = "end of CoverTab[97699]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:526
			_go_fuzz_dep_.CoverTab[97700]++

													for topic, topicErrors := range rsp.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:528
				_go_fuzz_dep_.CoverTab[97703]++
														for partition, partitionError := range topicErrors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:529
					_go_fuzz_dep_.CoverTab[97704]++
															if partitionError.errorCode != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:530
						_go_fuzz_dep_.CoverTab[97705]++
																errStr := fmt.Sprintf("[%s-%d]: %s", topic, partition, partitionError.errorCode.Error())
																errs = append(errs, errors.New(errStr))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:532
						// _ = "end of CoverTab[97705]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:533
						_go_fuzz_dep_.CoverTab[97706]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:533
						// _ = "end of CoverTab[97706]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:533
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:533
					// _ = "end of CoverTab[97704]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:534
				// _ = "end of CoverTab[97703]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:535
			// _ = "end of CoverTab[97700]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:536
		// _ = "end of CoverTab[97693]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:536
		_go_fuzz_dep_.CoverTab[97694]++

												if len(errs) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:538
			_go_fuzz_dep_.CoverTab[97707]++
													return ErrReassignPartitions{MultiError{&errs}}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:539
			// _ = "end of CoverTab[97707]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:540
			_go_fuzz_dep_.CoverTab[97708]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:540
			// _ = "end of CoverTab[97708]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:540
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:540
		// _ = "end of CoverTab[97694]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:540
		_go_fuzz_dep_.CoverTab[97695]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:542
		// _ = "end of CoverTab[97695]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:543
	// _ = "end of CoverTab[97688]"
}

func (ca *clusterAdmin) ListPartitionReassignments(topic string, partitions []int32) (topicStatus map[string]map[int32]*PartitionReplicaReassignmentsStatus, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:546
	_go_fuzz_dep_.CoverTab[97709]++
											if topic == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:547
		_go_fuzz_dep_.CoverTab[97712]++
												return nil, ErrInvalidTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:548
		// _ = "end of CoverTab[97712]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:549
		_go_fuzz_dep_.CoverTab[97713]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:549
		// _ = "end of CoverTab[97713]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:549
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:549
	// _ = "end of CoverTab[97709]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:549
	_go_fuzz_dep_.CoverTab[97710]++

											request := &ListPartitionReassignmentsRequest{
		TimeoutMs:	int32(60000),
		Version:	int16(0),
	}

	request.AddBlock(topic, partitions)

	b, err := ca.Controller()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:559
		_go_fuzz_dep_.CoverTab[97714]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:560
		// _ = "end of CoverTab[97714]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:561
		_go_fuzz_dep_.CoverTab[97715]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:561
		// _ = "end of CoverTab[97715]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:561
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:561
	// _ = "end of CoverTab[97710]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:561
	_go_fuzz_dep_.CoverTab[97711]++
											_ = b.Open(ca.client.Config())

											rsp, err := b.ListPartitionReassignments(request)

											if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:566
		_go_fuzz_dep_.CoverTab[97716]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:566
		return rsp != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:566
		// _ = "end of CoverTab[97716]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:566
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:566
		_go_fuzz_dep_.CoverTab[97717]++
												return rsp.TopicStatus, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:567
		// _ = "end of CoverTab[97717]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:568
		_go_fuzz_dep_.CoverTab[97718]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:569
		// _ = "end of CoverTab[97718]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:570
	// _ = "end of CoverTab[97711]"
}

func (ca *clusterAdmin) DeleteRecords(topic string, partitionOffsets map[int32]int64) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:573
	_go_fuzz_dep_.CoverTab[97719]++
											if topic == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:574
		_go_fuzz_dep_.CoverTab[97724]++
												return ErrInvalidTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:575
		// _ = "end of CoverTab[97724]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:576
		_go_fuzz_dep_.CoverTab[97725]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:576
		// _ = "end of CoverTab[97725]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:576
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:576
	// _ = "end of CoverTab[97719]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:576
	_go_fuzz_dep_.CoverTab[97720]++
											partitionPerBroker := make(map[*Broker][]int32)
											for partition := range partitionOffsets {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:578
		_go_fuzz_dep_.CoverTab[97726]++
												broker, err := ca.client.Leader(topic, partition)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:580
			_go_fuzz_dep_.CoverTab[97728]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:581
			// _ = "end of CoverTab[97728]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:582
			_go_fuzz_dep_.CoverTab[97729]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:582
			// _ = "end of CoverTab[97729]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:582
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:582
		// _ = "end of CoverTab[97726]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:582
		_go_fuzz_dep_.CoverTab[97727]++
												partitionPerBroker[broker] = append(partitionPerBroker[broker], partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:583
		// _ = "end of CoverTab[97727]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:584
	// _ = "end of CoverTab[97720]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:584
	_go_fuzz_dep_.CoverTab[97721]++
											errs := make([]error, 0)
											for broker, partitions := range partitionPerBroker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:586
		_go_fuzz_dep_.CoverTab[97730]++
												topics := make(map[string]*DeleteRecordsRequestTopic)
												recordsToDelete := make(map[int32]int64)
												for _, p := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:589
			_go_fuzz_dep_.CoverTab[97732]++
													recordsToDelete[p] = partitionOffsets[p]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:590
			// _ = "end of CoverTab[97732]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:591
		// _ = "end of CoverTab[97730]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:591
		_go_fuzz_dep_.CoverTab[97731]++
												topics[topic] = &DeleteRecordsRequestTopic{PartitionOffsets: recordsToDelete}
												request := &DeleteRecordsRequest{
			Topics:		topics,
			Timeout:	ca.conf.Admin.Timeout,
		}

		rsp, err := broker.DeleteRecords(request)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:599
			_go_fuzz_dep_.CoverTab[97733]++
													errs = append(errs, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:600
			// _ = "end of CoverTab[97733]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:601
			_go_fuzz_dep_.CoverTab[97734]++
													deleteRecordsResponseTopic, ok := rsp.Topics[topic]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:603
				_go_fuzz_dep_.CoverTab[97735]++
														errs = append(errs, ErrIncompleteResponse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:604
				// _ = "end of CoverTab[97735]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:605
				_go_fuzz_dep_.CoverTab[97736]++
														for _, deleteRecordsResponsePartition := range deleteRecordsResponseTopic.Partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:606
					_go_fuzz_dep_.CoverTab[97737]++
															if deleteRecordsResponsePartition.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:607
						_go_fuzz_dep_.CoverTab[97738]++
																errs = append(errs, errors.New(deleteRecordsResponsePartition.Err.Error()))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:608
						// _ = "end of CoverTab[97738]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:609
						_go_fuzz_dep_.CoverTab[97739]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:609
						// _ = "end of CoverTab[97739]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:609
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:609
					// _ = "end of CoverTab[97737]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:610
				// _ = "end of CoverTab[97736]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:611
			// _ = "end of CoverTab[97734]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:612
		// _ = "end of CoverTab[97731]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:613
	// _ = "end of CoverTab[97721]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:613
	_go_fuzz_dep_.CoverTab[97722]++
											if len(errs) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:614
		_go_fuzz_dep_.CoverTab[97740]++
												return ErrDeleteRecords{MultiError{&errs}}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:615
		// _ = "end of CoverTab[97740]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:616
		_go_fuzz_dep_.CoverTab[97741]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:616
		// _ = "end of CoverTab[97741]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:616
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:616
	// _ = "end of CoverTab[97722]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:616
	_go_fuzz_dep_.CoverTab[97723]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:619
	return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:619
	// _ = "end of CoverTab[97723]"
}

// Returns a bool indicating whether the resource request needs to go to a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:622
// specific broker
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:624
func dependsOnSpecificNode(resource ConfigResource) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:624
	_go_fuzz_dep_.CoverTab[97742]++
											return (resource.Type == BrokerResource && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:625
		_go_fuzz_dep_.CoverTab[97743]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:625
		return resource.Name != ""
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:625
		// _ = "end of CoverTab[97743]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:625
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:625
		_go_fuzz_dep_.CoverTab[97744]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:625
		return resource.Type == BrokerLoggerResource
												// _ = "end of CoverTab[97744]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:626
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:626
	// _ = "end of CoverTab[97742]"
}

func (ca *clusterAdmin) DescribeConfig(resource ConfigResource) ([]ConfigEntry, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:629
	_go_fuzz_dep_.CoverTab[97745]++
											var entries []ConfigEntry
											var resources []*ConfigResource
											resources = append(resources, &resource)

											request := &DescribeConfigsRequest{
		Resources: resources,
	}

	if ca.conf.Version.IsAtLeast(V1_1_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:638
		_go_fuzz_dep_.CoverTab[97752]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:639
		// _ = "end of CoverTab[97752]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:640
		_go_fuzz_dep_.CoverTab[97753]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:640
		// _ = "end of CoverTab[97753]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:640
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:640
	// _ = "end of CoverTab[97745]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:640
	_go_fuzz_dep_.CoverTab[97746]++

											if ca.conf.Version.IsAtLeast(V2_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:642
		_go_fuzz_dep_.CoverTab[97754]++
												request.Version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:643
		// _ = "end of CoverTab[97754]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:644
		_go_fuzz_dep_.CoverTab[97755]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:644
		// _ = "end of CoverTab[97755]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:644
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:644
	// _ = "end of CoverTab[97746]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:644
	_go_fuzz_dep_.CoverTab[97747]++

											var (
		b	*Broker
		err	error
	)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:652
	if dependsOnSpecificNode(resource) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:652
		_go_fuzz_dep_.CoverTab[97756]++
												var id int64
												id, err = strconv.ParseInt(resource.Name, 10, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:655
			_go_fuzz_dep_.CoverTab[97758]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:656
			// _ = "end of CoverTab[97758]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:657
			_go_fuzz_dep_.CoverTab[97759]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:657
			// _ = "end of CoverTab[97759]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:657
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:657
		// _ = "end of CoverTab[97756]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:657
		_go_fuzz_dep_.CoverTab[97757]++
												b, err = ca.findBroker(int32(id))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:658
		// _ = "end of CoverTab[97757]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:659
		_go_fuzz_dep_.CoverTab[97760]++
												b, err = ca.findAnyBroker()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:660
		// _ = "end of CoverTab[97760]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:661
	// _ = "end of CoverTab[97747]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:661
	_go_fuzz_dep_.CoverTab[97748]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:662
		_go_fuzz_dep_.CoverTab[97761]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:663
		// _ = "end of CoverTab[97761]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:664
		_go_fuzz_dep_.CoverTab[97762]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:664
		// _ = "end of CoverTab[97762]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:664
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:664
	// _ = "end of CoverTab[97748]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:664
	_go_fuzz_dep_.CoverTab[97749]++

											_ = b.Open(ca.client.Config())
											rsp, err := b.DescribeConfigs(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:668
		_go_fuzz_dep_.CoverTab[97763]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:669
		// _ = "end of CoverTab[97763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:670
		_go_fuzz_dep_.CoverTab[97764]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:670
		// _ = "end of CoverTab[97764]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:670
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:670
	// _ = "end of CoverTab[97749]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:670
	_go_fuzz_dep_.CoverTab[97750]++

											for _, rspResource := range rsp.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:672
		_go_fuzz_dep_.CoverTab[97765]++
												if rspResource.Name == resource.Name {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:673
			_go_fuzz_dep_.CoverTab[97766]++
													if rspResource.ErrorMsg != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:674
				_go_fuzz_dep_.CoverTab[97769]++
														return nil, errors.New(rspResource.ErrorMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:675
				// _ = "end of CoverTab[97769]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:676
				_go_fuzz_dep_.CoverTab[97770]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:676
				// _ = "end of CoverTab[97770]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:676
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:676
			// _ = "end of CoverTab[97766]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:676
			_go_fuzz_dep_.CoverTab[97767]++
													if rspResource.ErrorCode != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:677
				_go_fuzz_dep_.CoverTab[97771]++
														return nil, KError(rspResource.ErrorCode)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:678
				// _ = "end of CoverTab[97771]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:679
				_go_fuzz_dep_.CoverTab[97772]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:679
				// _ = "end of CoverTab[97772]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:679
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:679
			// _ = "end of CoverTab[97767]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:679
			_go_fuzz_dep_.CoverTab[97768]++
													for _, cfgEntry := range rspResource.Configs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:680
				_go_fuzz_dep_.CoverTab[97773]++
														entries = append(entries, *cfgEntry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:681
				// _ = "end of CoverTab[97773]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:682
			// _ = "end of CoverTab[97768]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:683
			_go_fuzz_dep_.CoverTab[97774]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:683
			// _ = "end of CoverTab[97774]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:683
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:683
		// _ = "end of CoverTab[97765]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:684
	// _ = "end of CoverTab[97750]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:684
	_go_fuzz_dep_.CoverTab[97751]++
											return entries, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:685
	// _ = "end of CoverTab[97751]"
}

func (ca *clusterAdmin) AlterConfig(resourceType ConfigResourceType, name string, entries map[string]*string, validateOnly bool) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:688
	_go_fuzz_dep_.CoverTab[97775]++
											var resources []*AlterConfigsResource
											resources = append(resources, &AlterConfigsResource{
		Type:		resourceType,
		Name:		name,
		ConfigEntries:	entries,
	})

	request := &AlterConfigsRequest{
		Resources:	resources,
		ValidateOnly:	validateOnly,
	}

	var (
		b	*Broker
		err	error
	)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:707
	if dependsOnSpecificNode(ConfigResource{Name: name, Type: resourceType}) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:707
		_go_fuzz_dep_.CoverTab[97780]++
												var id int64
												id, err = strconv.ParseInt(name, 10, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:710
			_go_fuzz_dep_.CoverTab[97782]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:711
			// _ = "end of CoverTab[97782]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:712
			_go_fuzz_dep_.CoverTab[97783]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:712
			// _ = "end of CoverTab[97783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:712
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:712
		// _ = "end of CoverTab[97780]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:712
		_go_fuzz_dep_.CoverTab[97781]++
												b, err = ca.findBroker(int32(id))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:713
		// _ = "end of CoverTab[97781]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:714
		_go_fuzz_dep_.CoverTab[97784]++
												b, err = ca.findAnyBroker()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:715
		// _ = "end of CoverTab[97784]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:716
	// _ = "end of CoverTab[97775]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:716
	_go_fuzz_dep_.CoverTab[97776]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:717
		_go_fuzz_dep_.CoverTab[97785]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:718
		// _ = "end of CoverTab[97785]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:719
		_go_fuzz_dep_.CoverTab[97786]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:719
		// _ = "end of CoverTab[97786]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:719
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:719
	// _ = "end of CoverTab[97776]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:719
	_go_fuzz_dep_.CoverTab[97777]++

											_ = b.Open(ca.client.Config())
											rsp, err := b.AlterConfigs(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:723
		_go_fuzz_dep_.CoverTab[97787]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:724
		// _ = "end of CoverTab[97787]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:725
		_go_fuzz_dep_.CoverTab[97788]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:725
		// _ = "end of CoverTab[97788]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:725
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:725
	// _ = "end of CoverTab[97777]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:725
	_go_fuzz_dep_.CoverTab[97778]++

											for _, rspResource := range rsp.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:727
		_go_fuzz_dep_.CoverTab[97789]++
												if rspResource.Name == name {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:728
			_go_fuzz_dep_.CoverTab[97790]++
													if rspResource.ErrorMsg != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:729
				_go_fuzz_dep_.CoverTab[97792]++
														return errors.New(rspResource.ErrorMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:730
				// _ = "end of CoverTab[97792]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:731
				_go_fuzz_dep_.CoverTab[97793]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:731
				// _ = "end of CoverTab[97793]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:731
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:731
			// _ = "end of CoverTab[97790]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:731
			_go_fuzz_dep_.CoverTab[97791]++
													if rspResource.ErrorCode != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:732
				_go_fuzz_dep_.CoverTab[97794]++
														return KError(rspResource.ErrorCode)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:733
				// _ = "end of CoverTab[97794]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:734
				_go_fuzz_dep_.CoverTab[97795]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:734
				// _ = "end of CoverTab[97795]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:734
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:734
			// _ = "end of CoverTab[97791]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:735
			_go_fuzz_dep_.CoverTab[97796]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:735
			// _ = "end of CoverTab[97796]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:735
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:735
		// _ = "end of CoverTab[97789]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:736
	// _ = "end of CoverTab[97778]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:736
	_go_fuzz_dep_.CoverTab[97779]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:737
	// _ = "end of CoverTab[97779]"
}

func (ca *clusterAdmin) IncrementalAlterConfig(resourceType ConfigResourceType, name string, entries map[string]IncrementalAlterConfigsEntry, validateOnly bool) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:740
	_go_fuzz_dep_.CoverTab[97797]++
											var resources []*IncrementalAlterConfigsResource
											resources = append(resources, &IncrementalAlterConfigsResource{
		Type:		resourceType,
		Name:		name,
		ConfigEntries:	entries,
	})

	request := &IncrementalAlterConfigsRequest{
		Resources:	resources,
		ValidateOnly:	validateOnly,
	}

	var (
		b	*Broker
		err	error
	)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:759
	if dependsOnSpecificNode(ConfigResource{Name: name, Type: resourceType}) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:759
		_go_fuzz_dep_.CoverTab[97802]++
												var id int64
												id, err = strconv.ParseInt(name, 10, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:762
			_go_fuzz_dep_.CoverTab[97804]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:763
			// _ = "end of CoverTab[97804]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:764
			_go_fuzz_dep_.CoverTab[97805]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:764
			// _ = "end of CoverTab[97805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:764
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:764
		// _ = "end of CoverTab[97802]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:764
		_go_fuzz_dep_.CoverTab[97803]++
												b, err = ca.findBroker(int32(id))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:765
		// _ = "end of CoverTab[97803]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:766
		_go_fuzz_dep_.CoverTab[97806]++
												b, err = ca.findAnyBroker()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:767
		// _ = "end of CoverTab[97806]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:768
	// _ = "end of CoverTab[97797]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:768
	_go_fuzz_dep_.CoverTab[97798]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:769
		_go_fuzz_dep_.CoverTab[97807]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:770
		// _ = "end of CoverTab[97807]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:771
		_go_fuzz_dep_.CoverTab[97808]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:771
		// _ = "end of CoverTab[97808]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:771
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:771
	// _ = "end of CoverTab[97798]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:771
	_go_fuzz_dep_.CoverTab[97799]++

											_ = b.Open(ca.client.Config())
											rsp, err := b.IncrementalAlterConfigs(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:775
		_go_fuzz_dep_.CoverTab[97809]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:776
		// _ = "end of CoverTab[97809]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:777
		_go_fuzz_dep_.CoverTab[97810]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:777
		// _ = "end of CoverTab[97810]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:777
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:777
	// _ = "end of CoverTab[97799]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:777
	_go_fuzz_dep_.CoverTab[97800]++

											for _, rspResource := range rsp.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:779
		_go_fuzz_dep_.CoverTab[97811]++
												if rspResource.Name == name {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:780
			_go_fuzz_dep_.CoverTab[97812]++
													if rspResource.ErrorMsg != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:781
				_go_fuzz_dep_.CoverTab[97814]++
														return errors.New(rspResource.ErrorMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:782
				// _ = "end of CoverTab[97814]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:783
				_go_fuzz_dep_.CoverTab[97815]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:783
				// _ = "end of CoverTab[97815]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:783
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:783
			// _ = "end of CoverTab[97812]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:783
			_go_fuzz_dep_.CoverTab[97813]++
													if rspResource.ErrorCode != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:784
				_go_fuzz_dep_.CoverTab[97816]++
														return KError(rspResource.ErrorCode)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:785
				// _ = "end of CoverTab[97816]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:786
				_go_fuzz_dep_.CoverTab[97817]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:786
				// _ = "end of CoverTab[97817]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:786
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:786
			// _ = "end of CoverTab[97813]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:787
			_go_fuzz_dep_.CoverTab[97818]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:787
			// _ = "end of CoverTab[97818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:787
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:787
		// _ = "end of CoverTab[97811]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:788
	// _ = "end of CoverTab[97800]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:788
	_go_fuzz_dep_.CoverTab[97801]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:789
	// _ = "end of CoverTab[97801]"
}

func (ca *clusterAdmin) CreateACL(resource Resource, acl Acl) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:792
	_go_fuzz_dep_.CoverTab[97819]++
											var acls []*AclCreation
											acls = append(acls, &AclCreation{resource, acl})
											request := &CreateAclsRequest{AclCreations: acls}

											if ca.conf.Version.IsAtLeast(V2_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:797
		_go_fuzz_dep_.CoverTab[97822]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:798
		// _ = "end of CoverTab[97822]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:799
		_go_fuzz_dep_.CoverTab[97823]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:799
		// _ = "end of CoverTab[97823]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:799
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:799
	// _ = "end of CoverTab[97819]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:799
	_go_fuzz_dep_.CoverTab[97820]++

											b, err := ca.Controller()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:802
		_go_fuzz_dep_.CoverTab[97824]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:803
		// _ = "end of CoverTab[97824]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:804
		_go_fuzz_dep_.CoverTab[97825]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:804
		// _ = "end of CoverTab[97825]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:804
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:804
	// _ = "end of CoverTab[97820]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:804
	_go_fuzz_dep_.CoverTab[97821]++

											_, err = b.CreateAcls(request)
											return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:807
	// _ = "end of CoverTab[97821]"
}

func (ca *clusterAdmin) ListAcls(filter AclFilter) ([]ResourceAcls, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:810
	_go_fuzz_dep_.CoverTab[97826]++
											request := &DescribeAclsRequest{AclFilter: filter}

											if ca.conf.Version.IsAtLeast(V2_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:813
		_go_fuzz_dep_.CoverTab[97831]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:814
		// _ = "end of CoverTab[97831]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:815
		_go_fuzz_dep_.CoverTab[97832]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:815
		// _ = "end of CoverTab[97832]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:815
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:815
	// _ = "end of CoverTab[97826]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:815
	_go_fuzz_dep_.CoverTab[97827]++

											b, err := ca.Controller()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:818
		_go_fuzz_dep_.CoverTab[97833]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:819
		// _ = "end of CoverTab[97833]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:820
		_go_fuzz_dep_.CoverTab[97834]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:820
		// _ = "end of CoverTab[97834]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:820
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:820
	// _ = "end of CoverTab[97827]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:820
	_go_fuzz_dep_.CoverTab[97828]++

											rsp, err := b.DescribeAcls(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:823
		_go_fuzz_dep_.CoverTab[97835]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:824
		// _ = "end of CoverTab[97835]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:825
		_go_fuzz_dep_.CoverTab[97836]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:825
		// _ = "end of CoverTab[97836]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:825
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:825
	// _ = "end of CoverTab[97828]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:825
	_go_fuzz_dep_.CoverTab[97829]++

											var lAcls []ResourceAcls
											for _, rAcl := range rsp.ResourceAcls {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:828
		_go_fuzz_dep_.CoverTab[97837]++
												lAcls = append(lAcls, *rAcl)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:829
		// _ = "end of CoverTab[97837]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:830
	// _ = "end of CoverTab[97829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:830
	_go_fuzz_dep_.CoverTab[97830]++
											return lAcls, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:831
	// _ = "end of CoverTab[97830]"
}

func (ca *clusterAdmin) DeleteACL(filter AclFilter, validateOnly bool) ([]MatchingAcl, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:834
	_go_fuzz_dep_.CoverTab[97838]++
											var filters []*AclFilter
											filters = append(filters, &filter)
											request := &DeleteAclsRequest{Filters: filters}

											if ca.conf.Version.IsAtLeast(V2_0_0_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:839
		_go_fuzz_dep_.CoverTab[97843]++
												request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:840
		// _ = "end of CoverTab[97843]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:841
		_go_fuzz_dep_.CoverTab[97844]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:841
		// _ = "end of CoverTab[97844]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:841
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:841
	// _ = "end of CoverTab[97838]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:841
	_go_fuzz_dep_.CoverTab[97839]++

											b, err := ca.Controller()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:844
		_go_fuzz_dep_.CoverTab[97845]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:845
		// _ = "end of CoverTab[97845]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:846
		_go_fuzz_dep_.CoverTab[97846]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:846
		// _ = "end of CoverTab[97846]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:846
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:846
	// _ = "end of CoverTab[97839]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:846
	_go_fuzz_dep_.CoverTab[97840]++

											rsp, err := b.DeleteAcls(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:849
		_go_fuzz_dep_.CoverTab[97847]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:850
		// _ = "end of CoverTab[97847]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:851
		_go_fuzz_dep_.CoverTab[97848]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:851
		// _ = "end of CoverTab[97848]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:851
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:851
	// _ = "end of CoverTab[97840]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:851
	_go_fuzz_dep_.CoverTab[97841]++

											var mAcls []MatchingAcl
											for _, fr := range rsp.FilterResponses {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:854
		_go_fuzz_dep_.CoverTab[97849]++
												for _, mACL := range fr.MatchingAcls {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:855
			_go_fuzz_dep_.CoverTab[97850]++
													mAcls = append(mAcls, *mACL)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:856
			// _ = "end of CoverTab[97850]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:857
		// _ = "end of CoverTab[97849]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:858
	// _ = "end of CoverTab[97841]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:858
	_go_fuzz_dep_.CoverTab[97842]++
											return mAcls, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:859
	// _ = "end of CoverTab[97842]"
}

func (ca *clusterAdmin) DescribeConsumerGroups(groups []string) (result []*GroupDescription, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:862
	_go_fuzz_dep_.CoverTab[97851]++
											groupsPerBroker := make(map[*Broker][]string)

											for _, group := range groups {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:865
		_go_fuzz_dep_.CoverTab[97854]++
												controller, err := ca.client.Coordinator(group)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:867
			_go_fuzz_dep_.CoverTab[97856]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:868
			// _ = "end of CoverTab[97856]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:869
			_go_fuzz_dep_.CoverTab[97857]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:869
			// _ = "end of CoverTab[97857]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:869
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:869
		// _ = "end of CoverTab[97854]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:869
		_go_fuzz_dep_.CoverTab[97855]++
												groupsPerBroker[controller] = append(groupsPerBroker[controller], group)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:870
		// _ = "end of CoverTab[97855]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:871
	// _ = "end of CoverTab[97851]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:871
	_go_fuzz_dep_.CoverTab[97852]++

											for broker, brokerGroups := range groupsPerBroker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:873
		_go_fuzz_dep_.CoverTab[97858]++
												response, err := broker.DescribeGroups(&DescribeGroupsRequest{
			Groups: brokerGroups,
		})
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:877
			_go_fuzz_dep_.CoverTab[97860]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:878
			// _ = "end of CoverTab[97860]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:879
			_go_fuzz_dep_.CoverTab[97861]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:879
			// _ = "end of CoverTab[97861]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:879
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:879
		// _ = "end of CoverTab[97858]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:879
		_go_fuzz_dep_.CoverTab[97859]++

												result = append(result, response.Groups...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:881
		// _ = "end of CoverTab[97859]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:882
	// _ = "end of CoverTab[97852]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:882
	_go_fuzz_dep_.CoverTab[97853]++
											return result, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:883
	// _ = "end of CoverTab[97853]"
}

func (ca *clusterAdmin) ListConsumerGroups() (allGroups map[string]string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:886
	_go_fuzz_dep_.CoverTab[97862]++
											allGroups = make(map[string]string)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:890
	brokers := ca.client.Brokers()
	groupMaps := make(chan map[string]string, len(brokers))
	errChan := make(chan error, len(brokers))
	wg := sync.WaitGroup{}

	for _, b := range brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:895
		_go_fuzz_dep_.CoverTab[97865]++
												wg.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:896
		_curRoutineNum114_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:896
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum114_)
												go func(b *Broker, conf *Config) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:897
			_go_fuzz_dep_.CoverTab[97866]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:897
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:897
				_go_fuzz_dep_.CoverTab[97869]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:897
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum114_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:897
				// _ = "end of CoverTab[97869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:897
			}()
													defer wg.Done()
													_ = b.Open(conf)

													response, err := b.ListGroups(&ListGroupsRequest{})
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:902
				_go_fuzz_dep_.CoverTab[97870]++
														errChan <- err
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:904
				// _ = "end of CoverTab[97870]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:905
				_go_fuzz_dep_.CoverTab[97871]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:905
				// _ = "end of CoverTab[97871]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:905
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:905
			// _ = "end of CoverTab[97866]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:905
			_go_fuzz_dep_.CoverTab[97867]++

													groups := make(map[string]string)
													for group, typ := range response.Groups {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:908
				_go_fuzz_dep_.CoverTab[97872]++
														groups[group] = typ
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:909
				// _ = "end of CoverTab[97872]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:910
			// _ = "end of CoverTab[97867]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:910
			_go_fuzz_dep_.CoverTab[97868]++

													groupMaps <- groups
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:912
			// _ = "end of CoverTab[97868]"
		}(b, ca.conf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:913
		// _ = "end of CoverTab[97865]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:914
	// _ = "end of CoverTab[97862]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:914
	_go_fuzz_dep_.CoverTab[97863]++

											wg.Wait()
											close(groupMaps)
											close(errChan)

											for groupMap := range groupMaps {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:920
		_go_fuzz_dep_.CoverTab[97873]++
												for group, protocolType := range groupMap {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:921
			_go_fuzz_dep_.CoverTab[97874]++
													allGroups[group] = protocolType
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:922
			// _ = "end of CoverTab[97874]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:923
		// _ = "end of CoverTab[97873]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:924
	// _ = "end of CoverTab[97863]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:924
	_go_fuzz_dep_.CoverTab[97864]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:927
	err = <-errChan
											return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:928
	// _ = "end of CoverTab[97864]"
}

func (ca *clusterAdmin) ListConsumerGroupOffsets(group string, topicPartitions map[string][]int32) (*OffsetFetchResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:931
	_go_fuzz_dep_.CoverTab[97875]++
											coordinator, err := ca.client.Coordinator(group)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:933
		_go_fuzz_dep_.CoverTab[97878]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:934
		// _ = "end of CoverTab[97878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:935
		_go_fuzz_dep_.CoverTab[97879]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:935
		// _ = "end of CoverTab[97879]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:935
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:935
	// _ = "end of CoverTab[97875]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:935
	_go_fuzz_dep_.CoverTab[97876]++

											request := &OffsetFetchRequest{
		ConsumerGroup:	group,
		partitions:	topicPartitions,
	}

	if ca.conf.Version.IsAtLeast(V0_10_2_0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:942
		_go_fuzz_dep_.CoverTab[97880]++
												request.Version = 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:943
		// _ = "end of CoverTab[97880]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:944
		_go_fuzz_dep_.CoverTab[97881]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:944
		if ca.conf.Version.IsAtLeast(V0_8_2_2) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:944
			_go_fuzz_dep_.CoverTab[97882]++
													request.Version = 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:945
			// _ = "end of CoverTab[97882]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:946
			_go_fuzz_dep_.CoverTab[97883]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:946
			// _ = "end of CoverTab[97883]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:946
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:946
		// _ = "end of CoverTab[97881]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:946
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:946
	// _ = "end of CoverTab[97876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:946
	_go_fuzz_dep_.CoverTab[97877]++

											return coordinator.FetchOffset(request)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:948
	// _ = "end of CoverTab[97877]"
}

func (ca *clusterAdmin) DeleteConsumerGroupOffset(group string, topic string, partition int32) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:951
	_go_fuzz_dep_.CoverTab[97884]++
											coordinator, err := ca.client.Coordinator(group)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:953
		_go_fuzz_dep_.CoverTab[97889]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:954
		// _ = "end of CoverTab[97889]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:955
		_go_fuzz_dep_.CoverTab[97890]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:955
		// _ = "end of CoverTab[97890]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:955
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:955
	// _ = "end of CoverTab[97884]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:955
	_go_fuzz_dep_.CoverTab[97885]++

											request := &DeleteOffsetsRequest{
		Group:	group,
		partitions: map[string][]int32{
			topic: {partition},
		},
	}

	resp, err := coordinator.DeleteOffsets(request)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:965
		_go_fuzz_dep_.CoverTab[97891]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:966
		// _ = "end of CoverTab[97891]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:967
		_go_fuzz_dep_.CoverTab[97892]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:967
		// _ = "end of CoverTab[97892]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:967
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:967
	// _ = "end of CoverTab[97885]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:967
	_go_fuzz_dep_.CoverTab[97886]++

											if resp.ErrorCode != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:969
		_go_fuzz_dep_.CoverTab[97893]++
												return resp.ErrorCode
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:970
		// _ = "end of CoverTab[97893]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:971
		_go_fuzz_dep_.CoverTab[97894]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:971
		// _ = "end of CoverTab[97894]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:971
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:971
	// _ = "end of CoverTab[97886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:971
	_go_fuzz_dep_.CoverTab[97887]++

											if resp.Errors[topic][partition] != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:973
		_go_fuzz_dep_.CoverTab[97895]++
												return resp.Errors[topic][partition]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:974
		// _ = "end of CoverTab[97895]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:975
		_go_fuzz_dep_.CoverTab[97896]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:975
		// _ = "end of CoverTab[97896]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:975
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:975
	// _ = "end of CoverTab[97887]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:975
	_go_fuzz_dep_.CoverTab[97888]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:976
	// _ = "end of CoverTab[97888]"
}

func (ca *clusterAdmin) DeleteConsumerGroup(group string) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:979
	_go_fuzz_dep_.CoverTab[97897]++
											coordinator, err := ca.client.Coordinator(group)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:981
		_go_fuzz_dep_.CoverTab[97902]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:982
		// _ = "end of CoverTab[97902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:983
		_go_fuzz_dep_.CoverTab[97903]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:983
		// _ = "end of CoverTab[97903]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:983
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:983
	// _ = "end of CoverTab[97897]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:983
	_go_fuzz_dep_.CoverTab[97898]++

											request := &DeleteGroupsRequest{
		Groups: []string{group},
	}

	resp, err := coordinator.DeleteGroups(request)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:990
		_go_fuzz_dep_.CoverTab[97904]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:991
		// _ = "end of CoverTab[97904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:992
		_go_fuzz_dep_.CoverTab[97905]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:992
		// _ = "end of CoverTab[97905]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:992
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:992
	// _ = "end of CoverTab[97898]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:992
	_go_fuzz_dep_.CoverTab[97899]++

											groupErr, ok := resp.GroupErrorCodes[group]
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:995
		_go_fuzz_dep_.CoverTab[97906]++
												return ErrIncompleteResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:996
		// _ = "end of CoverTab[97906]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:997
		_go_fuzz_dep_.CoverTab[97907]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:997
		// _ = "end of CoverTab[97907]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:997
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:997
	// _ = "end of CoverTab[97899]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:997
	_go_fuzz_dep_.CoverTab[97900]++

											if groupErr != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:999
		_go_fuzz_dep_.CoverTab[97908]++
												return groupErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1000
		// _ = "end of CoverTab[97908]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1001
		_go_fuzz_dep_.CoverTab[97909]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1001
		// _ = "end of CoverTab[97909]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1001
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1001
	// _ = "end of CoverTab[97900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1001
	_go_fuzz_dep_.CoverTab[97901]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1003
	// _ = "end of CoverTab[97901]"
}

func (ca *clusterAdmin) DescribeLogDirs(brokerIds []int32) (allLogDirs map[int32][]DescribeLogDirsResponseDirMetadata, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1006
	_go_fuzz_dep_.CoverTab[97910]++
											allLogDirs = make(map[int32][]DescribeLogDirsResponseDirMetadata)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1010
	logDirsMaps := make(chan map[int32][]DescribeLogDirsResponseDirMetadata, len(brokerIds))
	errChan := make(chan error, len(brokerIds))
	wg := sync.WaitGroup{}

	for _, b := range brokerIds {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1014
		_go_fuzz_dep_.CoverTab[97913]++
												wg.Add(1)
												broker, err := ca.findBroker(b)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1017
			_go_fuzz_dep_.CoverTab[97915]++
													Logger.Printf("Unable to find broker with ID = %v\n", b)
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1019
			// _ = "end of CoverTab[97915]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1020
			_go_fuzz_dep_.CoverTab[97916]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1020
			// _ = "end of CoverTab[97916]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1020
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1020
		// _ = "end of CoverTab[97913]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1020
		_go_fuzz_dep_.CoverTab[97914]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1020
		_curRoutineNum115_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1020
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum115_)
												go func(b *Broker, conf *Config) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1021
			_go_fuzz_dep_.CoverTab[97917]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1021
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1021
				_go_fuzz_dep_.CoverTab[97919]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1021
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum115_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1021
				// _ = "end of CoverTab[97919]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1021
			}()
													defer wg.Done()
													_ = b.Open(conf)

													response, err := b.DescribeLogDirs(&DescribeLogDirsRequest{})
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1026
				_go_fuzz_dep_.CoverTab[97920]++
														errChan <- err
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1028
				// _ = "end of CoverTab[97920]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1029
				_go_fuzz_dep_.CoverTab[97921]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1029
				// _ = "end of CoverTab[97921]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1029
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1029
			// _ = "end of CoverTab[97917]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1029
			_go_fuzz_dep_.CoverTab[97918]++
													logDirs := make(map[int32][]DescribeLogDirsResponseDirMetadata)
													logDirs[b.ID()] = response.LogDirs
													logDirsMaps <- logDirs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1032
			// _ = "end of CoverTab[97918]"
		}(broker, ca.conf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1033
		// _ = "end of CoverTab[97914]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1034
	// _ = "end of CoverTab[97910]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1034
	_go_fuzz_dep_.CoverTab[97911]++

											wg.Wait()
											close(logDirsMaps)
											close(errChan)

											for logDirsMap := range logDirsMaps {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1040
		_go_fuzz_dep_.CoverTab[97922]++
												for id, logDirs := range logDirsMap {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1041
			_go_fuzz_dep_.CoverTab[97923]++
													allLogDirs[id] = logDirs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1042
			// _ = "end of CoverTab[97923]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1043
		// _ = "end of CoverTab[97922]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1044
	// _ = "end of CoverTab[97911]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1044
	_go_fuzz_dep_.CoverTab[97912]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1047
	err = <-errChan
											return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1048
	// _ = "end of CoverTab[97912]"
}

func (ca *clusterAdmin) DescribeUserScramCredentials(users []string) ([]*DescribeUserScramCredentialsResult, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1051
	_go_fuzz_dep_.CoverTab[97924]++
											req := &DescribeUserScramCredentialsRequest{}
											for _, u := range users {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1053
		_go_fuzz_dep_.CoverTab[97928]++
												req.DescribeUsers = append(req.DescribeUsers, DescribeUserScramCredentialsRequestUser{
			Name: u,
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1056
		// _ = "end of CoverTab[97928]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1057
	// _ = "end of CoverTab[97924]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1057
	_go_fuzz_dep_.CoverTab[97925]++

											b, err := ca.Controller()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1060
		_go_fuzz_dep_.CoverTab[97929]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1061
		// _ = "end of CoverTab[97929]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1062
		_go_fuzz_dep_.CoverTab[97930]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1062
		// _ = "end of CoverTab[97930]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1062
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1062
	// _ = "end of CoverTab[97925]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1062
	_go_fuzz_dep_.CoverTab[97926]++

											rsp, err := b.DescribeUserScramCredentials(req)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1065
		_go_fuzz_dep_.CoverTab[97931]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1066
		// _ = "end of CoverTab[97931]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1067
		_go_fuzz_dep_.CoverTab[97932]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1067
		// _ = "end of CoverTab[97932]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1067
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1067
	// _ = "end of CoverTab[97926]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1067
	_go_fuzz_dep_.CoverTab[97927]++

											return rsp.Results, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1069
	// _ = "end of CoverTab[97927]"
}

func (ca *clusterAdmin) UpsertUserScramCredentials(upsert []AlterUserScramCredentialsUpsert) ([]*AlterUserScramCredentialsResult, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1072
	_go_fuzz_dep_.CoverTab[97933]++
											res, err := ca.AlterUserScramCredentials(upsert, nil)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1074
		_go_fuzz_dep_.CoverTab[97935]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1075
		// _ = "end of CoverTab[97935]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1076
		_go_fuzz_dep_.CoverTab[97936]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1076
		// _ = "end of CoverTab[97936]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1076
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1076
	// _ = "end of CoverTab[97933]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1076
	_go_fuzz_dep_.CoverTab[97934]++

											return res, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1078
	// _ = "end of CoverTab[97934]"
}

func (ca *clusterAdmin) DeleteUserScramCredentials(delete []AlterUserScramCredentialsDelete) ([]*AlterUserScramCredentialsResult, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1081
	_go_fuzz_dep_.CoverTab[97937]++
											res, err := ca.AlterUserScramCredentials(nil, delete)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1083
		_go_fuzz_dep_.CoverTab[97939]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1084
		// _ = "end of CoverTab[97939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1085
		_go_fuzz_dep_.CoverTab[97940]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1085
		// _ = "end of CoverTab[97940]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1085
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1085
	// _ = "end of CoverTab[97937]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1085
	_go_fuzz_dep_.CoverTab[97938]++

											return res, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1087
	// _ = "end of CoverTab[97938]"
}

func (ca *clusterAdmin) AlterUserScramCredentials(u []AlterUserScramCredentialsUpsert, d []AlterUserScramCredentialsDelete) ([]*AlterUserScramCredentialsResult, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1090
	_go_fuzz_dep_.CoverTab[97941]++
											req := &AlterUserScramCredentialsRequest{
		Deletions:	d,
		Upsertions:	u,
	}

	b, err := ca.Controller()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1097
		_go_fuzz_dep_.CoverTab[97944]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1098
		// _ = "end of CoverTab[97944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1099
		_go_fuzz_dep_.CoverTab[97945]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1099
		// _ = "end of CoverTab[97945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1099
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1099
	// _ = "end of CoverTab[97941]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1099
	_go_fuzz_dep_.CoverTab[97942]++

											rsp, err := b.AlterUserScramCredentials(req)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1102
		_go_fuzz_dep_.CoverTab[97946]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1103
		// _ = "end of CoverTab[97946]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1104
		_go_fuzz_dep_.CoverTab[97947]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1104
		// _ = "end of CoverTab[97947]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1104
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1104
	// _ = "end of CoverTab[97942]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1104
	_go_fuzz_dep_.CoverTab[97943]++

											return rsp.Results, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1106
	// _ = "end of CoverTab[97943]"
}

// Describe All : use an empty/nil components slice + strict = false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1109
// Contains components: strict = false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1109
// Contains only components: strict = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1112
func (ca *clusterAdmin) DescribeClientQuotas(components []QuotaFilterComponent, strict bool) ([]DescribeClientQuotasEntry, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1112
	_go_fuzz_dep_.CoverTab[97948]++
											request := &DescribeClientQuotasRequest{
		Components:	components,
		Strict:		strict,
	}

	b, err := ca.Controller()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1119
		_go_fuzz_dep_.CoverTab[97953]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1120
		// _ = "end of CoverTab[97953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1121
		_go_fuzz_dep_.CoverTab[97954]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1121
		// _ = "end of CoverTab[97954]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1121
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1121
	// _ = "end of CoverTab[97948]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1121
	_go_fuzz_dep_.CoverTab[97949]++

											rsp, err := b.DescribeClientQuotas(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1124
		_go_fuzz_dep_.CoverTab[97955]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1125
		// _ = "end of CoverTab[97955]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1126
		_go_fuzz_dep_.CoverTab[97956]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1126
		// _ = "end of CoverTab[97956]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1126
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1126
	// _ = "end of CoverTab[97949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1126
	_go_fuzz_dep_.CoverTab[97950]++

											if rsp.ErrorMsg != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1128
		_go_fuzz_dep_.CoverTab[97957]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1128
		return len(*rsp.ErrorMsg) > 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1128
		// _ = "end of CoverTab[97957]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1128
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1128
		_go_fuzz_dep_.CoverTab[97958]++
												return nil, errors.New(*rsp.ErrorMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1129
		// _ = "end of CoverTab[97958]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1130
		_go_fuzz_dep_.CoverTab[97959]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1130
		// _ = "end of CoverTab[97959]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1130
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1130
	// _ = "end of CoverTab[97950]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1130
	_go_fuzz_dep_.CoverTab[97951]++
											if rsp.ErrorCode != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1131
		_go_fuzz_dep_.CoverTab[97960]++
												return nil, rsp.ErrorCode
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1132
		// _ = "end of CoverTab[97960]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1133
		_go_fuzz_dep_.CoverTab[97961]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1133
		// _ = "end of CoverTab[97961]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1133
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1133
	// _ = "end of CoverTab[97951]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1133
	_go_fuzz_dep_.CoverTab[97952]++

											return rsp.Entries, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1135
	// _ = "end of CoverTab[97952]"
}

func (ca *clusterAdmin) AlterClientQuotas(entity []QuotaEntityComponent, op ClientQuotasOp, validateOnly bool) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1138
	_go_fuzz_dep_.CoverTab[97962]++
											entry := AlterClientQuotasEntry{
		Entity:	entity,
		Ops:	[]ClientQuotasOp{op},
	}

	request := &AlterClientQuotasRequest{
		Entries:	[]AlterClientQuotasEntry{entry},
		ValidateOnly:	validateOnly,
	}

	b, err := ca.Controller()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1150
		_go_fuzz_dep_.CoverTab[97966]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1151
		// _ = "end of CoverTab[97966]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1152
		_go_fuzz_dep_.CoverTab[97967]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1152
		// _ = "end of CoverTab[97967]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1152
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1152
	// _ = "end of CoverTab[97962]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1152
	_go_fuzz_dep_.CoverTab[97963]++

											rsp, err := b.AlterClientQuotas(request)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1155
		_go_fuzz_dep_.CoverTab[97968]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1156
		// _ = "end of CoverTab[97968]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1157
		_go_fuzz_dep_.CoverTab[97969]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1157
		// _ = "end of CoverTab[97969]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1157
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1157
	// _ = "end of CoverTab[97963]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1157
	_go_fuzz_dep_.CoverTab[97964]++

											for _, entry := range rsp.Entries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1159
		_go_fuzz_dep_.CoverTab[97970]++
												if entry.ErrorMsg != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1160
			_go_fuzz_dep_.CoverTab[97972]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1160
			return len(*entry.ErrorMsg) > 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1160
			// _ = "end of CoverTab[97972]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1160
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1160
			_go_fuzz_dep_.CoverTab[97973]++
													return errors.New(*entry.ErrorMsg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1161
			// _ = "end of CoverTab[97973]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1162
			_go_fuzz_dep_.CoverTab[97974]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1162
			// _ = "end of CoverTab[97974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1162
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1162
		// _ = "end of CoverTab[97970]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1162
		_go_fuzz_dep_.CoverTab[97971]++
												if entry.ErrorCode != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1163
			_go_fuzz_dep_.CoverTab[97975]++
													return entry.ErrorCode
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1164
			// _ = "end of CoverTab[97975]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1165
			_go_fuzz_dep_.CoverTab[97976]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1165
			// _ = "end of CoverTab[97976]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1165
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1165
		// _ = "end of CoverTab[97971]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1166
	// _ = "end of CoverTab[97964]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1166
	_go_fuzz_dep_.CoverTab[97965]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1168
	// _ = "end of CoverTab[97965]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1169
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/admin.go:1169
var _ = _go_fuzz_dep_.CoverTab
