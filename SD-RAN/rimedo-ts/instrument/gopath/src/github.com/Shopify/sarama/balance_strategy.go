//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1
)

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	// RangeBalanceStrategyName identifies strategies that use the range partition assignment strategy
	RangeBalanceStrategyName	= "range"

	// RoundRobinBalanceStrategyName identifies strategies that use the round-robin partition assignment strategy
	RoundRobinBalanceStrategyName	= "roundrobin"

	// StickyBalanceStrategyName identifies strategies that use the sticky-partition assignment strategy
	StickyBalanceStrategyName	= "sticky"

	defaultGeneration	= -1
)

// BalanceStrategyPlan is the results of any BalanceStrategy.Plan attempt.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:25
// It contains an allocation of topic/partitions by memberID in the form of
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:25
// a `memberID -> topic -> partitions` map.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:28
type BalanceStrategyPlan map[string]map[string][]int32

// Add assigns a topic with a number partitions to a member.
func (p BalanceStrategyPlan) Add(memberID, topic string, partitions ...int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:31
	_go_fuzz_dep_.CoverTab[98946]++
												if len(partitions) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:32
		_go_fuzz_dep_.CoverTab[98949]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:33
		// _ = "end of CoverTab[98949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:34
		_go_fuzz_dep_.CoverTab[98950]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:34
		// _ = "end of CoverTab[98950]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:34
	// _ = "end of CoverTab[98946]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:34
	_go_fuzz_dep_.CoverTab[98947]++
												if _, ok := p[memberID]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:35
		_go_fuzz_dep_.CoverTab[98951]++
													p[memberID] = make(map[string][]int32, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:36
		// _ = "end of CoverTab[98951]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:37
		_go_fuzz_dep_.CoverTab[98952]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:37
		// _ = "end of CoverTab[98952]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:37
	// _ = "end of CoverTab[98947]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:37
	_go_fuzz_dep_.CoverTab[98948]++
												p[memberID][topic] = append(p[memberID][topic], partitions...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:38
	// _ = "end of CoverTab[98948]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:43
// BalanceStrategy is used to balance topics and partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:43
// across members of a consumer group
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:45
type BalanceStrategy interface {
	// Name uniquely identifies the strategy.
	Name() string

	// Plan accepts a map of `memberID -> metadata` and a map of `topic -> partitions`
	// and returns a distribution plan.
	Plan(members map[string]ConsumerGroupMemberMetadata, topics map[string][]int32) (BalanceStrategyPlan, error)

	// AssignmentData returns the serialized assignment data for the specified
	// memberID
	AssignmentData(memberID string, topics map[string][]int32, generationID int32) ([]byte, error)
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:60
// BalanceStrategyRange is the default and assigns partitions as ranges to consumer group members.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:60
// Example with one topic T with six partitions (0..5) and two members (M1, M2):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:60
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:60
//	M1: {T: [0, 1, 2]}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:60
//	M2: {T: [3, 4, 5]}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:64
var BalanceStrategyRange = &balanceStrategy{
	name:	RangeBalanceStrategyName,
	coreFn: func(plan BalanceStrategyPlan, memberIDs []string, topic string, partitions []int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:66
		_go_fuzz_dep_.CoverTab[98953]++
													step := float64(len(partitions)) / float64(len(memberIDs))

													for i, memberID := range memberIDs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:69
			_go_fuzz_dep_.CoverTab[98954]++
														pos := float64(i)
														min := int(math.Floor(pos*step + 0.5))
														max := int(math.Floor((pos+1)*step + 0.5))
														plan.Add(memberID, topic, partitions[min:max]...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:73
			// _ = "end of CoverTab[98954]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:74
		// _ = "end of CoverTab[98953]"
	},
}

// BalanceStrategySticky assigns partitions to members with an attempt to preserve earlier assignments
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
// while maintain a balanced partition distribution.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
// Example with topic T with six partitions (0..5) and two members (M1, M2):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//	M1: {T: [0, 2, 4]}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//	M2: {T: [1, 3, 5]}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
// On reassignment with an additional consumer, you might get an assignment plan like:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//	M1: {T: [0, 2]}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//	M2: {T: [1, 3]}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:78
//	M3: {T: [4, 5]}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:89
var BalanceStrategySticky = &stickyBalanceStrategy{}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:93
type balanceStrategy struct {
	name	string
	coreFn	func(plan BalanceStrategyPlan, memberIDs []string, topic string, partitions []int32)
}

// Name implements BalanceStrategy.
func (s *balanceStrategy) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:99
	_go_fuzz_dep_.CoverTab[98955]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:99
	return s.name
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:99
	// _ = "end of CoverTab[98955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:99
}

// Plan implements BalanceStrategy.
func (s *balanceStrategy) Plan(members map[string]ConsumerGroupMemberMetadata, topics map[string][]int32) (BalanceStrategyPlan, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:102
	_go_fuzz_dep_.CoverTab[98956]++

												mbt := make(map[string][]string)
												for memberID, meta := range members {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:105
		_go_fuzz_dep_.CoverTab[98960]++
													for _, topic := range meta.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:106
			_go_fuzz_dep_.CoverTab[98961]++
														mbt[topic] = append(mbt[topic], memberID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:107
			// _ = "end of CoverTab[98961]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:108
		// _ = "end of CoverTab[98960]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:109
	// _ = "end of CoverTab[98956]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:109
	_go_fuzz_dep_.CoverTab[98957]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:112
	for topic, memberIDs := range mbt {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:112
		_go_fuzz_dep_.CoverTab[98962]++
													sort.Sort(&balanceStrategySortable{
			topic:		topic,
			memberIDs:	memberIDs,
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:116
		// _ = "end of CoverTab[98962]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:117
	// _ = "end of CoverTab[98957]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:117
	_go_fuzz_dep_.CoverTab[98958]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:120
	plan := make(BalanceStrategyPlan, len(members))
	for topic, memberIDs := range mbt {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:121
		_go_fuzz_dep_.CoverTab[98963]++
													s.coreFn(plan, memberIDs, topic, topics[topic])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:122
		// _ = "end of CoverTab[98963]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:123
	// _ = "end of CoverTab[98958]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:123
	_go_fuzz_dep_.CoverTab[98959]++
												return plan, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:124
	// _ = "end of CoverTab[98959]"
}

// AssignmentData simple strategies do not require any shared assignment data
func (s *balanceStrategy) AssignmentData(memberID string, topics map[string][]int32, generationID int32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:128
	_go_fuzz_dep_.CoverTab[98964]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:129
	// _ = "end of CoverTab[98964]"
}

type balanceStrategySortable struct {
	topic		string
	memberIDs	[]string
}

func (p balanceStrategySortable) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:137
	_go_fuzz_dep_.CoverTab[98965]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:137
	return len(p.memberIDs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:137
	// _ = "end of CoverTab[98965]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:137
}
func (p balanceStrategySortable) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:138
	_go_fuzz_dep_.CoverTab[98966]++
												p.memberIDs[i], p.memberIDs[j] = p.memberIDs[j], p.memberIDs[i]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:139
	// _ = "end of CoverTab[98966]"
}

func (p balanceStrategySortable) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:142
	_go_fuzz_dep_.CoverTab[98967]++
												return balanceStrategyHashValue(p.topic, p.memberIDs[i]) < balanceStrategyHashValue(p.topic, p.memberIDs[j])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:143
	// _ = "end of CoverTab[98967]"
}

func balanceStrategyHashValue(vv ...string) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:146
	_go_fuzz_dep_.CoverTab[98968]++
												h := uint32(2166136261)
												for _, s := range vv {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:148
		_go_fuzz_dep_.CoverTab[98970]++
													for _, c := range s {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:149
			_go_fuzz_dep_.CoverTab[98971]++
														h ^= uint32(c)
														h *= 16777619
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:151
			// _ = "end of CoverTab[98971]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:152
		// _ = "end of CoverTab[98970]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:153
	// _ = "end of CoverTab[98968]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:153
	_go_fuzz_dep_.CoverTab[98969]++
												return h
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:154
	// _ = "end of CoverTab[98969]"
}

type stickyBalanceStrategy struct {
	movements partitionMovements
}

// Name implements BalanceStrategy.
func (s *stickyBalanceStrategy) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:162
	_go_fuzz_dep_.CoverTab[98972]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:162
	return StickyBalanceStrategyName
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:162
	// _ = "end of CoverTab[98972]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:162
}

// Plan implements BalanceStrategy.
func (s *stickyBalanceStrategy) Plan(members map[string]ConsumerGroupMemberMetadata, topics map[string][]int32) (BalanceStrategyPlan, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:165
	_go_fuzz_dep_.CoverTab[98973]++

												s.movements = partitionMovements{
		Movements:			make(map[topicPartitionAssignment]consumerPair),
		PartitionMovementsByTopic:	make(map[string]map[consumerPair]map[topicPartitionAssignment]bool),
	}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:173
	currentAssignment, prevAssignment, err := prepopulateCurrentAssignments(members)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:174
		_go_fuzz_dep_.CoverTab[98982]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:175
		// _ = "end of CoverTab[98982]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:176
		_go_fuzz_dep_.CoverTab[98983]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:176
		// _ = "end of CoverTab[98983]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:176
	// _ = "end of CoverTab[98973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:176
	_go_fuzz_dep_.CoverTab[98974]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:179
	isFreshAssignment := false
	if len(currentAssignment) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:180
		_go_fuzz_dep_.CoverTab[98984]++
													isFreshAssignment = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:181
		// _ = "end of CoverTab[98984]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:182
		_go_fuzz_dep_.CoverTab[98985]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:182
		// _ = "end of CoverTab[98985]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:182
	// _ = "end of CoverTab[98974]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:182
	_go_fuzz_dep_.CoverTab[98975]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:185
	partition2AllPotentialConsumers := make(map[topicPartitionAssignment][]string)
	for topic, partitions := range topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:186
		_go_fuzz_dep_.CoverTab[98986]++
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:187
			_go_fuzz_dep_.CoverTab[98987]++
														partition2AllPotentialConsumers[topicPartitionAssignment{Topic: topic, Partition: partition}] = []string{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:188
			// _ = "end of CoverTab[98987]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:189
		// _ = "end of CoverTab[98986]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:190
	// _ = "end of CoverTab[98975]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:190
	_go_fuzz_dep_.CoverTab[98976]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:194
	consumer2AllPotentialPartitions := make(map[string][]topicPartitionAssignment, len(members))
	for memberID, meta := range members {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:195
		_go_fuzz_dep_.CoverTab[98988]++
													consumer2AllPotentialPartitions[memberID] = make([]topicPartitionAssignment, 0)
													for _, topicSubscription := range meta.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:197
			_go_fuzz_dep_.CoverTab[98990]++

														if _, found := topics[topicSubscription]; found {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:199
				_go_fuzz_dep_.CoverTab[98991]++
															for _, partition := range topics[topicSubscription] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:200
					_go_fuzz_dep_.CoverTab[98992]++
																topicPartition := topicPartitionAssignment{Topic: topicSubscription, Partition: partition}
																consumer2AllPotentialPartitions[memberID] = append(consumer2AllPotentialPartitions[memberID], topicPartition)
																partition2AllPotentialConsumers[topicPartition] = append(partition2AllPotentialConsumers[topicPartition], memberID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:203
					// _ = "end of CoverTab[98992]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:204
				// _ = "end of CoverTab[98991]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:205
				_go_fuzz_dep_.CoverTab[98993]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:205
				// _ = "end of CoverTab[98993]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:205
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:205
			// _ = "end of CoverTab[98990]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:206
		// _ = "end of CoverTab[98988]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:206
		_go_fuzz_dep_.CoverTab[98989]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:209
		if _, exists := currentAssignment[memberID]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:209
			_go_fuzz_dep_.CoverTab[98994]++
														currentAssignment[memberID] = make([]topicPartitionAssignment, 0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:210
			// _ = "end of CoverTab[98994]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:211
			_go_fuzz_dep_.CoverTab[98995]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:211
			// _ = "end of CoverTab[98995]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:211
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:211
		// _ = "end of CoverTab[98989]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:212
	// _ = "end of CoverTab[98976]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:212
	_go_fuzz_dep_.CoverTab[98977]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:215
	currentPartitionConsumers := make(map[topicPartitionAssignment]string, len(currentAssignment))
	unvisitedPartitions := make(map[topicPartitionAssignment]bool, len(partition2AllPotentialConsumers))
	for partition := range partition2AllPotentialConsumers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:217
		_go_fuzz_dep_.CoverTab[98996]++
													unvisitedPartitions[partition] = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:218
		// _ = "end of CoverTab[98996]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:219
	// _ = "end of CoverTab[98977]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:219
	_go_fuzz_dep_.CoverTab[98978]++
												var unassignedPartitions []topicPartitionAssignment
												for memberID, partitions := range currentAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:221
		_go_fuzz_dep_.CoverTab[98997]++
													var keepPartitions []topicPartitionAssignment
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:223
			_go_fuzz_dep_.CoverTab[98999]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:226
			if _, exists := partition2AllPotentialConsumers[partition]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:226
				_go_fuzz_dep_.CoverTab[99002]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:227
				// _ = "end of CoverTab[99002]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:228
				_go_fuzz_dep_.CoverTab[99003]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:228
				// _ = "end of CoverTab[99003]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:228
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:228
			// _ = "end of CoverTab[98999]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:228
			_go_fuzz_dep_.CoverTab[99000]++
														delete(unvisitedPartitions, partition)
														currentPartitionConsumers[partition] = memberID

														if !strsContains(members[memberID].Topics, partition.Topic) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:232
				_go_fuzz_dep_.CoverTab[99004]++
															unassignedPartitions = append(unassignedPartitions, partition)
															continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:234
				// _ = "end of CoverTab[99004]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:235
				_go_fuzz_dep_.CoverTab[99005]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:235
				// _ = "end of CoverTab[99005]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:235
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:235
			// _ = "end of CoverTab[99000]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:235
			_go_fuzz_dep_.CoverTab[99001]++
														keepPartitions = append(keepPartitions, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:236
			// _ = "end of CoverTab[99001]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:237
		// _ = "end of CoverTab[98997]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:237
		_go_fuzz_dep_.CoverTab[98998]++
													currentAssignment[memberID] = keepPartitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:238
		// _ = "end of CoverTab[98998]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:239
	// _ = "end of CoverTab[98978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:239
	_go_fuzz_dep_.CoverTab[98979]++
												for unvisited := range unvisitedPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:240
		_go_fuzz_dep_.CoverTab[99006]++
													unassignedPartitions = append(unassignedPartitions, unvisited)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:241
		// _ = "end of CoverTab[99006]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:242
	// _ = "end of CoverTab[98979]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:242
	_go_fuzz_dep_.CoverTab[98980]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:245
	sortedPartitions := sortPartitions(currentAssignment, prevAssignment, isFreshAssignment, partition2AllPotentialConsumers, consumer2AllPotentialPartitions)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:252
	sortedCurrentSubscriptions := sortMemberIDsByPartitionAssignments(currentAssignment)
												s.balance(currentAssignment, prevAssignment, sortedPartitions, unassignedPartitions, sortedCurrentSubscriptions, consumer2AllPotentialPartitions, partition2AllPotentialConsumers, currentPartitionConsumers)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:256
	plan := make(BalanceStrategyPlan, len(currentAssignment))
	for memberID, assignments := range currentAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:257
		_go_fuzz_dep_.CoverTab[99007]++
													if len(assignments) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:258
			_go_fuzz_dep_.CoverTab[99008]++
														plan[memberID] = make(map[string][]int32)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:259
			// _ = "end of CoverTab[99008]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:260
			_go_fuzz_dep_.CoverTab[99009]++
														for _, assignment := range assignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:261
				_go_fuzz_dep_.CoverTab[99010]++
															plan.Add(memberID, assignment.Topic, assignment.Partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:262
				// _ = "end of CoverTab[99010]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:263
			// _ = "end of CoverTab[99009]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:264
		// _ = "end of CoverTab[99007]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:265
	// _ = "end of CoverTab[98980]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:265
	_go_fuzz_dep_.CoverTab[98981]++
												return plan, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:266
	// _ = "end of CoverTab[98981]"
}

// AssignmentData serializes the set of topics currently assigned to the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:269
// specified member as part of the supplied balance plan
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:271
func (s *stickyBalanceStrategy) AssignmentData(memberID string, topics map[string][]int32, generationID int32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:271
	_go_fuzz_dep_.CoverTab[99011]++
												return encode(&StickyAssignorUserDataV1{
		Topics:		topics,
		Generation:	generationID,
	}, nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:275
	// _ = "end of CoverTab[99011]"
}

func strsContains(s []string, value string) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:278
	_go_fuzz_dep_.CoverTab[99012]++
												for _, entry := range s {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:279
		_go_fuzz_dep_.CoverTab[99014]++
													if entry == value {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:280
			_go_fuzz_dep_.CoverTab[99015]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:281
			// _ = "end of CoverTab[99015]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:282
			_go_fuzz_dep_.CoverTab[99016]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:282
			// _ = "end of CoverTab[99016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:282
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:282
		// _ = "end of CoverTab[99014]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:283
	// _ = "end of CoverTab[99012]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:283
	_go_fuzz_dep_.CoverTab[99013]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:284
	// _ = "end of CoverTab[99013]"
}

// Balance assignments across consumers for maximum fairness and stickiness.
func (s *stickyBalanceStrategy) balance(currentAssignment map[string][]topicPartitionAssignment, prevAssignment map[topicPartitionAssignment]consumerGenerationPair, sortedPartitions []topicPartitionAssignment, unassignedPartitions []topicPartitionAssignment, sortedCurrentSubscriptions []string, consumer2AllPotentialPartitions map[string][]topicPartitionAssignment, partition2AllPotentialConsumers map[topicPartitionAssignment][]string, currentPartitionConsumer map[topicPartitionAssignment]string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:288
	_go_fuzz_dep_.CoverTab[99017]++
												initializing := false
												if len(sortedCurrentSubscriptions) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:290
		_go_fuzz_dep_.CoverTab[99024]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:290
		return len(currentAssignment[sortedCurrentSubscriptions[0]]) == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:290
		// _ = "end of CoverTab[99024]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:290
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:290
		_go_fuzz_dep_.CoverTab[99025]++
													initializing = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:291
		// _ = "end of CoverTab[99025]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:292
		_go_fuzz_dep_.CoverTab[99026]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:292
		// _ = "end of CoverTab[99026]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:292
	// _ = "end of CoverTab[99017]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:292
	_go_fuzz_dep_.CoverTab[99018]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:295
	for _, partition := range unassignedPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:295
		_go_fuzz_dep_.CoverTab[99027]++

													if len(partition2AllPotentialConsumers[partition]) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:297
			_go_fuzz_dep_.CoverTab[99029]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:298
			// _ = "end of CoverTab[99029]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:299
			_go_fuzz_dep_.CoverTab[99030]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:299
			// _ = "end of CoverTab[99030]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:299
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:299
		// _ = "end of CoverTab[99027]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:299
		_go_fuzz_dep_.CoverTab[99028]++
													sortedCurrentSubscriptions = assignPartition(partition, sortedCurrentSubscriptions, currentAssignment, consumer2AllPotentialPartitions, currentPartitionConsumer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:300
		// _ = "end of CoverTab[99028]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:301
	// _ = "end of CoverTab[99018]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:301
	_go_fuzz_dep_.CoverTab[99019]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:304
	for partition := range partition2AllPotentialConsumers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:304
		_go_fuzz_dep_.CoverTab[99031]++
													if !canTopicPartitionParticipateInReassignment(partition, partition2AllPotentialConsumers) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:305
			_go_fuzz_dep_.CoverTab[99032]++
														sortedPartitions = removeTopicPartitionFromMemberAssignments(sortedPartitions, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:306
			// _ = "end of CoverTab[99032]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:307
			_go_fuzz_dep_.CoverTab[99033]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:307
			// _ = "end of CoverTab[99033]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:307
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:307
		// _ = "end of CoverTab[99031]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:308
	// _ = "end of CoverTab[99019]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:308
	_go_fuzz_dep_.CoverTab[99020]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:311
	fixedAssignments := make(map[string][]topicPartitionAssignment)
	for memberID := range consumer2AllPotentialPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:312
		_go_fuzz_dep_.CoverTab[99034]++
													if !canConsumerParticipateInReassignment(memberID, currentAssignment, consumer2AllPotentialPartitions, partition2AllPotentialConsumers) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:313
			_go_fuzz_dep_.CoverTab[99035]++
														fixedAssignments[memberID] = currentAssignment[memberID]
														delete(currentAssignment, memberID)
														sortedCurrentSubscriptions = sortMemberIDsByPartitionAssignments(currentAssignment)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:316
			// _ = "end of CoverTab[99035]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:317
			_go_fuzz_dep_.CoverTab[99036]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:317
			// _ = "end of CoverTab[99036]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:317
		// _ = "end of CoverTab[99034]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:318
	// _ = "end of CoverTab[99020]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:318
	_go_fuzz_dep_.CoverTab[99021]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:321
	preBalanceAssignment := deepCopyAssignment(currentAssignment)
	preBalancePartitionConsumers := make(map[topicPartitionAssignment]string, len(currentPartitionConsumer))
	for k, v := range currentPartitionConsumer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:323
		_go_fuzz_dep_.CoverTab[99037]++
													preBalancePartitionConsumers[k] = v
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:324
		// _ = "end of CoverTab[99037]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:325
	// _ = "end of CoverTab[99021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:325
	_go_fuzz_dep_.CoverTab[99022]++

												reassignmentPerformed := s.performReassignments(sortedPartitions, currentAssignment, prevAssignment, sortedCurrentSubscriptions, consumer2AllPotentialPartitions, partition2AllPotentialConsumers, currentPartitionConsumer)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
	if !initializing && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
		_go_fuzz_dep_.CoverTab[99038]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
		return reassignmentPerformed
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
		// _ = "end of CoverTab[99038]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
		_go_fuzz_dep_.CoverTab[99039]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
		return getBalanceScore(currentAssignment) >= getBalanceScore(preBalanceAssignment)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
		// _ = "end of CoverTab[99039]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:331
		_go_fuzz_dep_.CoverTab[99040]++
													currentAssignment = deepCopyAssignment(preBalanceAssignment)
													currentPartitionConsumer = make(map[topicPartitionAssignment]string, len(preBalancePartitionConsumers))
													for k, v := range preBalancePartitionConsumers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:334
			_go_fuzz_dep_.CoverTab[99041]++
														currentPartitionConsumer[k] = v
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:335
			// _ = "end of CoverTab[99041]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:336
		// _ = "end of CoverTab[99040]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:337
		_go_fuzz_dep_.CoverTab[99042]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:337
		// _ = "end of CoverTab[99042]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:337
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:337
	// _ = "end of CoverTab[99022]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:337
	_go_fuzz_dep_.CoverTab[99023]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:340
	for consumer, assignments := range fixedAssignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:340
		_go_fuzz_dep_.CoverTab[99043]++
													currentAssignment[consumer] = assignments
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:341
		// _ = "end of CoverTab[99043]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:342
	// _ = "end of CoverTab[99023]"
}

// BalanceStrategyRoundRobin assigns partitions to members in alternating order.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:345
// For example, there are two topics (t0, t1) and two consumer (m0, m1), and each topic has three partitions (p0, p1, p2):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:345
// M0: [t0p0, t0p2, t1p1]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:345
// M1: [t0p1, t1p0, t1p2]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:349
var BalanceStrategyRoundRobin = new(roundRobinBalancer)

type roundRobinBalancer struct{}

func (b *roundRobinBalancer) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:353
	_go_fuzz_dep_.CoverTab[99044]++
												return RoundRobinBalanceStrategyName
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:354
	// _ = "end of CoverTab[99044]"
}

func (b *roundRobinBalancer) Plan(memberAndMetadata map[string]ConsumerGroupMemberMetadata, topics map[string][]int32) (BalanceStrategyPlan, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:357
	_go_fuzz_dep_.CoverTab[99045]++
												if len(memberAndMetadata) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:358
		_go_fuzz_dep_.CoverTab[99052]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:358
		return len(topics) == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:358
		// _ = "end of CoverTab[99052]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:358
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:358
		_go_fuzz_dep_.CoverTab[99053]++
													return nil, errors.New("members and topics are not provided")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:359
		// _ = "end of CoverTab[99053]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:360
		_go_fuzz_dep_.CoverTab[99054]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:360
		// _ = "end of CoverTab[99054]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:360
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:360
	// _ = "end of CoverTab[99045]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:360
	_go_fuzz_dep_.CoverTab[99046]++
	// sort partitions
	var topicPartitions []topicAndPartition
	for topic, partitions := range topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:363
		_go_fuzz_dep_.CoverTab[99055]++
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:364
			_go_fuzz_dep_.CoverTab[99056]++
														topicPartitions = append(topicPartitions, topicAndPartition{topic: topic, partition: partition})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:365
			// _ = "end of CoverTab[99056]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:366
		// _ = "end of CoverTab[99055]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:367
	// _ = "end of CoverTab[99046]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:367
	_go_fuzz_dep_.CoverTab[99047]++
												sort.SliceStable(topicPartitions, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:368
		_go_fuzz_dep_.CoverTab[99057]++
													pi := topicPartitions[i]
													pj := topicPartitions[j]
													return pi.comparedValue() < pj.comparedValue()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:371
		// _ = "end of CoverTab[99057]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:372
	// _ = "end of CoverTab[99047]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:372
	_go_fuzz_dep_.CoverTab[99048]++

	// sort members
	var members []memberAndTopic
	for memberID, meta := range memberAndMetadata {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:376
		_go_fuzz_dep_.CoverTab[99058]++
													m := memberAndTopic{
			memberID:	memberID,
			topics:		make(map[string]struct{}),
		}
		for _, t := range meta.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:381
			_go_fuzz_dep_.CoverTab[99060]++
														m.topics[t] = struct{}{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:382
			// _ = "end of CoverTab[99060]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:383
		// _ = "end of CoverTab[99058]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:383
		_go_fuzz_dep_.CoverTab[99059]++
													members = append(members, m)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:384
		// _ = "end of CoverTab[99059]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:385
	// _ = "end of CoverTab[99048]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:385
	_go_fuzz_dep_.CoverTab[99049]++
												sort.SliceStable(members, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:386
		_go_fuzz_dep_.CoverTab[99061]++
													mi := members[i]
													mj := members[j]
													return mi.memberID < mj.memberID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:389
		// _ = "end of CoverTab[99061]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:390
	// _ = "end of CoverTab[99049]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:390
	_go_fuzz_dep_.CoverTab[99050]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:393
	plan := make(BalanceStrategyPlan, len(members))
	i := 0
	n := len(members)
	for _, tp := range topicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:396
		_go_fuzz_dep_.CoverTab[99062]++
													m := members[i%n]
													for !m.hasTopic(tp.topic) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:398
			_go_fuzz_dep_.CoverTab[99064]++
														i++
														m = members[i%n]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:400
			// _ = "end of CoverTab[99064]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:401
		// _ = "end of CoverTab[99062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:401
		_go_fuzz_dep_.CoverTab[99063]++
													plan.Add(m.memberID, tp.topic, tp.partition)
													i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:403
		// _ = "end of CoverTab[99063]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:404
	// _ = "end of CoverTab[99050]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:404
	_go_fuzz_dep_.CoverTab[99051]++
												return plan, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:405
	// _ = "end of CoverTab[99051]"
}

func (b *roundRobinBalancer) AssignmentData(memberID string, topics map[string][]int32, generationID int32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:408
	_go_fuzz_dep_.CoverTab[99065]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:409
	// _ = "end of CoverTab[99065]"
}

type topicAndPartition struct {
	topic		string
	partition	int32
}

func (tp *topicAndPartition) comparedValue() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:417
	_go_fuzz_dep_.CoverTab[99066]++
												return fmt.Sprintf("%s-%d", tp.topic, tp.partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:418
	// _ = "end of CoverTab[99066]"
}

type memberAndTopic struct {
	memberID	string
	topics		map[string]struct{}
}

func (m *memberAndTopic) hasTopic(topic string) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:426
	_go_fuzz_dep_.CoverTab[99067]++
												_, isExist := m.topics[topic]
												return isExist
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:428
	// _ = "end of CoverTab[99067]"
}

// Calculate the balance score of the given assignment, as the sum of assigned partitions size difference of all consumer pairs.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:431
// A perfectly balanced assignment (with all consumers getting the same number of partitions) has a balance score of 0.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:431
// Lower balance score indicates a more balanced assignment.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:434
func getBalanceScore(assignment map[string][]topicPartitionAssignment) int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:434
	_go_fuzz_dep_.CoverTab[99068]++
												consumer2AssignmentSize := make(map[string]int, len(assignment))
												for memberID, partitions := range assignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:436
		_go_fuzz_dep_.CoverTab[99071]++
													consumer2AssignmentSize[memberID] = len(partitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:437
		// _ = "end of CoverTab[99071]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:438
	// _ = "end of CoverTab[99068]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:438
	_go_fuzz_dep_.CoverTab[99069]++

												var score float64
												for memberID, consumerAssignmentSize := range consumer2AssignmentSize {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:441
		_go_fuzz_dep_.CoverTab[99072]++
													delete(consumer2AssignmentSize, memberID)
													for _, otherConsumerAssignmentSize := range consumer2AssignmentSize {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:443
			_go_fuzz_dep_.CoverTab[99073]++
														score += math.Abs(float64(consumerAssignmentSize - otherConsumerAssignmentSize))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:444
			// _ = "end of CoverTab[99073]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:445
		// _ = "end of CoverTab[99072]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:446
	// _ = "end of CoverTab[99069]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:446
	_go_fuzz_dep_.CoverTab[99070]++
												return int(score)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:447
	// _ = "end of CoverTab[99070]"
}

// Determine whether the current assignment plan is balanced.
func isBalanced(currentAssignment map[string][]topicPartitionAssignment, allSubscriptions map[string][]topicPartitionAssignment) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:451
	_go_fuzz_dep_.CoverTab[99074]++
												sortedCurrentSubscriptions := sortMemberIDsByPartitionAssignments(currentAssignment)
												min := len(currentAssignment[sortedCurrentSubscriptions[0]])
												max := len(currentAssignment[sortedCurrentSubscriptions[len(sortedCurrentSubscriptions)-1]])
												if min >= max-1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:455
		_go_fuzz_dep_.CoverTab[99078]++

													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:457
		// _ = "end of CoverTab[99078]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:458
		_go_fuzz_dep_.CoverTab[99079]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:458
		// _ = "end of CoverTab[99079]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:458
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:458
	// _ = "end of CoverTab[99074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:458
	_go_fuzz_dep_.CoverTab[99075]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:461
	allPartitions := make(map[topicPartitionAssignment]string)
	for memberID, partitions := range currentAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:462
		_go_fuzz_dep_.CoverTab[99080]++
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:463
			_go_fuzz_dep_.CoverTab[99081]++
														if _, exists := allPartitions[partition]; exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:464
				_go_fuzz_dep_.CoverTab[99083]++
															Logger.Printf("Topic %s Partition %d is assigned more than one consumer", partition.Topic, partition.Partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:465
				// _ = "end of CoverTab[99083]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:466
				_go_fuzz_dep_.CoverTab[99084]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:466
				// _ = "end of CoverTab[99084]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:466
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:466
			// _ = "end of CoverTab[99081]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:466
			_go_fuzz_dep_.CoverTab[99082]++
														allPartitions[partition] = memberID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:467
			// _ = "end of CoverTab[99082]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:468
		// _ = "end of CoverTab[99080]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:469
	// _ = "end of CoverTab[99075]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:469
	_go_fuzz_dep_.CoverTab[99076]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:473
	for _, memberID := range sortedCurrentSubscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:473
		_go_fuzz_dep_.CoverTab[99085]++
													consumerPartitions := currentAssignment[memberID]
													consumerPartitionCount := len(consumerPartitions)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:478
		if consumerPartitionCount == len(allSubscriptions[memberID]) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:478
			_go_fuzz_dep_.CoverTab[99087]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:479
			// _ = "end of CoverTab[99087]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:480
			_go_fuzz_dep_.CoverTab[99088]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:480
			// _ = "end of CoverTab[99088]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:480
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:480
		// _ = "end of CoverTab[99085]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:480
		_go_fuzz_dep_.CoverTab[99086]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:483
		potentialTopicPartitions := allSubscriptions[memberID]
		for _, partition := range potentialTopicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:484
			_go_fuzz_dep_.CoverTab[99089]++
														if !memberAssignmentsIncludeTopicPartition(currentAssignment[memberID], partition) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:485
				_go_fuzz_dep_.CoverTab[99090]++
															otherConsumer := allPartitions[partition]
															otherConsumerPartitionCount := len(currentAssignment[otherConsumer])
															if consumerPartitionCount < otherConsumerPartitionCount {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:488
					_go_fuzz_dep_.CoverTab[99091]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:489
					// _ = "end of CoverTab[99091]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:490
					_go_fuzz_dep_.CoverTab[99092]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:490
					// _ = "end of CoverTab[99092]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:490
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:490
				// _ = "end of CoverTab[99090]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:491
				_go_fuzz_dep_.CoverTab[99093]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:491
				// _ = "end of CoverTab[99093]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:491
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:491
			// _ = "end of CoverTab[99089]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:492
		// _ = "end of CoverTab[99086]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:493
	// _ = "end of CoverTab[99076]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:493
	_go_fuzz_dep_.CoverTab[99077]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:494
	// _ = "end of CoverTab[99077]"
}

// Reassign all topic partitions that need reassignment until balanced.
func (s *stickyBalanceStrategy) performReassignments(reassignablePartitions []topicPartitionAssignment, currentAssignment map[string][]topicPartitionAssignment, prevAssignment map[topicPartitionAssignment]consumerGenerationPair, sortedCurrentSubscriptions []string, consumer2AllPotentialPartitions map[string][]topicPartitionAssignment, partition2AllPotentialConsumers map[topicPartitionAssignment][]string, currentPartitionConsumer map[topicPartitionAssignment]string) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:498
	_go_fuzz_dep_.CoverTab[99094]++
												reassignmentPerformed := false
												modified := false

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:503
	for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:503
		_go_fuzz_dep_.CoverTab[99095]++
													modified = false

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:507
		for _, partition := range reassignablePartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:507
			_go_fuzz_dep_.CoverTab[99097]++
														if isBalanced(currentAssignment, consumer2AllPotentialPartitions) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:508
				_go_fuzz_dep_.CoverTab[99102]++
															break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:509
				// _ = "end of CoverTab[99102]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:510
				_go_fuzz_dep_.CoverTab[99103]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:510
				// _ = "end of CoverTab[99103]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:510
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:510
			// _ = "end of CoverTab[99097]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:510
			_go_fuzz_dep_.CoverTab[99098]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:513
			if len(partition2AllPotentialConsumers[partition]) <= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:513
				_go_fuzz_dep_.CoverTab[99104]++
															Logger.Printf("Expected more than one potential consumer for partition %s topic %d", partition.Topic, partition.Partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:514
				// _ = "end of CoverTab[99104]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:515
				_go_fuzz_dep_.CoverTab[99105]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:515
				// _ = "end of CoverTab[99105]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:515
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:515
			// _ = "end of CoverTab[99098]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:515
			_go_fuzz_dep_.CoverTab[99099]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:518
			consumer := currentPartitionConsumer[partition]
			if consumer == "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:519
				_go_fuzz_dep_.CoverTab[99106]++
															Logger.Printf("Expected topic %s partition %d to be assigned to a consumer", partition.Topic, partition.Partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:520
				// _ = "end of CoverTab[99106]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:521
				_go_fuzz_dep_.CoverTab[99107]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:521
				// _ = "end of CoverTab[99107]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:521
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:521
			// _ = "end of CoverTab[99099]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:521
			_go_fuzz_dep_.CoverTab[99100]++

														if _, exists := prevAssignment[partition]; exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:523
				_go_fuzz_dep_.CoverTab[99108]++
															if len(currentAssignment[consumer]) > (len(currentAssignment[prevAssignment[partition].MemberID]) + 1) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:524
					_go_fuzz_dep_.CoverTab[99109]++
																sortedCurrentSubscriptions = s.reassignPartition(partition, currentAssignment, sortedCurrentSubscriptions, currentPartitionConsumer, prevAssignment[partition].MemberID)
																reassignmentPerformed = true
																modified = true
																continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:528
					// _ = "end of CoverTab[99109]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:529
					_go_fuzz_dep_.CoverTab[99110]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:529
					// _ = "end of CoverTab[99110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:529
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:529
				// _ = "end of CoverTab[99108]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:530
				_go_fuzz_dep_.CoverTab[99111]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:530
				// _ = "end of CoverTab[99111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:530
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:530
			// _ = "end of CoverTab[99100]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:530
			_go_fuzz_dep_.CoverTab[99101]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:533
			for _, otherConsumer := range partition2AllPotentialConsumers[partition] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:533
				_go_fuzz_dep_.CoverTab[99112]++
															if len(currentAssignment[consumer]) > (len(currentAssignment[otherConsumer]) + 1) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:534
					_go_fuzz_dep_.CoverTab[99113]++
																sortedCurrentSubscriptions = s.reassignPartitionToNewConsumer(partition, currentAssignment, sortedCurrentSubscriptions, currentPartitionConsumer, consumer2AllPotentialPartitions)
																reassignmentPerformed = true
																modified = true
																break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:538
					// _ = "end of CoverTab[99113]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:539
					_go_fuzz_dep_.CoverTab[99114]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:539
					// _ = "end of CoverTab[99114]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:539
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:539
				// _ = "end of CoverTab[99112]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:540
			// _ = "end of CoverTab[99101]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:541
		// _ = "end of CoverTab[99095]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:541
		_go_fuzz_dep_.CoverTab[99096]++
													if !modified {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:542
			_go_fuzz_dep_.CoverTab[99115]++
														return reassignmentPerformed
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:543
			// _ = "end of CoverTab[99115]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:544
			_go_fuzz_dep_.CoverTab[99116]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:544
			// _ = "end of CoverTab[99116]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:544
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:544
		// _ = "end of CoverTab[99096]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:545
	// _ = "end of CoverTab[99094]"
}

// Identify a new consumer for a topic partition and reassign it.
func (s *stickyBalanceStrategy) reassignPartitionToNewConsumer(partition topicPartitionAssignment, currentAssignment map[string][]topicPartitionAssignment, sortedCurrentSubscriptions []string, currentPartitionConsumer map[topicPartitionAssignment]string, consumer2AllPotentialPartitions map[string][]topicPartitionAssignment) []string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:549
	_go_fuzz_dep_.CoverTab[99117]++
												for _, anotherConsumer := range sortedCurrentSubscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:550
		_go_fuzz_dep_.CoverTab[99119]++
													if memberAssignmentsIncludeTopicPartition(consumer2AllPotentialPartitions[anotherConsumer], partition) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:551
			_go_fuzz_dep_.CoverTab[99120]++
														return s.reassignPartition(partition, currentAssignment, sortedCurrentSubscriptions, currentPartitionConsumer, anotherConsumer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:552
			// _ = "end of CoverTab[99120]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:553
			_go_fuzz_dep_.CoverTab[99121]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:553
			// _ = "end of CoverTab[99121]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:553
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:553
		// _ = "end of CoverTab[99119]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:554
	// _ = "end of CoverTab[99117]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:554
	_go_fuzz_dep_.CoverTab[99118]++
												return sortedCurrentSubscriptions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:555
	// _ = "end of CoverTab[99118]"
}

// Reassign a specific partition to a new consumer
func (s *stickyBalanceStrategy) reassignPartition(partition topicPartitionAssignment, currentAssignment map[string][]topicPartitionAssignment, sortedCurrentSubscriptions []string, currentPartitionConsumer map[topicPartitionAssignment]string, newConsumer string) []string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:559
	_go_fuzz_dep_.CoverTab[99122]++
												consumer := currentPartitionConsumer[partition]

												partitionToBeMoved := s.movements.getTheActualPartitionToBeMoved(partition, consumer, newConsumer)
												return s.processPartitionMovement(partitionToBeMoved, newConsumer, currentAssignment, sortedCurrentSubscriptions, currentPartitionConsumer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:563
	// _ = "end of CoverTab[99122]"
}

// Track the movement of a topic partition after assignment
func (s *stickyBalanceStrategy) processPartitionMovement(partition topicPartitionAssignment, newConsumer string, currentAssignment map[string][]topicPartitionAssignment, sortedCurrentSubscriptions []string, currentPartitionConsumer map[topicPartitionAssignment]string) []string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:567
	_go_fuzz_dep_.CoverTab[99123]++
												oldConsumer := currentPartitionConsumer[partition]
												s.movements.movePartition(partition, oldConsumer, newConsumer)

												currentAssignment[oldConsumer] = removeTopicPartitionFromMemberAssignments(currentAssignment[oldConsumer], partition)
												currentAssignment[newConsumer] = append(currentAssignment[newConsumer], partition)
												currentPartitionConsumer[partition] = newConsumer
												return sortMemberIDsByPartitionAssignments(currentAssignment)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:574
	// _ = "end of CoverTab[99123]"
}

// Determine whether a specific consumer should be considered for topic partition assignment.
func canConsumerParticipateInReassignment(memberID string, currentAssignment map[string][]topicPartitionAssignment, consumer2AllPotentialPartitions map[string][]topicPartitionAssignment, partition2AllPotentialConsumers map[topicPartitionAssignment][]string) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:578
	_go_fuzz_dep_.CoverTab[99124]++
												currentPartitions := currentAssignment[memberID]
												currentAssignmentSize := len(currentPartitions)
												maxAssignmentSize := len(consumer2AllPotentialPartitions[memberID])
												if currentAssignmentSize > maxAssignmentSize {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:582
		_go_fuzz_dep_.CoverTab[99128]++
													Logger.Printf("The consumer %s is assigned more partitions than the maximum possible", memberID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:583
		// _ = "end of CoverTab[99128]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:584
		_go_fuzz_dep_.CoverTab[99129]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:584
		// _ = "end of CoverTab[99129]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:584
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:584
	// _ = "end of CoverTab[99124]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:584
	_go_fuzz_dep_.CoverTab[99125]++
												if currentAssignmentSize < maxAssignmentSize {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:585
		_go_fuzz_dep_.CoverTab[99130]++

													return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:587
		// _ = "end of CoverTab[99130]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:588
		_go_fuzz_dep_.CoverTab[99131]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:588
		// _ = "end of CoverTab[99131]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:588
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:588
	// _ = "end of CoverTab[99125]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:588
	_go_fuzz_dep_.CoverTab[99126]++
												for _, partition := range currentPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:589
		_go_fuzz_dep_.CoverTab[99132]++
													if canTopicPartitionParticipateInReassignment(partition, partition2AllPotentialConsumers) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:590
			_go_fuzz_dep_.CoverTab[99133]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:591
			// _ = "end of CoverTab[99133]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:592
			_go_fuzz_dep_.CoverTab[99134]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:592
			// _ = "end of CoverTab[99134]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:592
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:592
		// _ = "end of CoverTab[99132]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:593
	// _ = "end of CoverTab[99126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:593
	_go_fuzz_dep_.CoverTab[99127]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:594
	// _ = "end of CoverTab[99127]"
}

// Only consider reassigning those topic partitions that have two or more potential consumers.
func canTopicPartitionParticipateInReassignment(partition topicPartitionAssignment, partition2AllPotentialConsumers map[topicPartitionAssignment][]string) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:598
	_go_fuzz_dep_.CoverTab[99135]++
												return len(partition2AllPotentialConsumers[partition]) >= 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:599
	// _ = "end of CoverTab[99135]"
}

// The assignment should improve the overall balance of the partition assignments to consumers.
func assignPartition(partition topicPartitionAssignment, sortedCurrentSubscriptions []string, currentAssignment map[string][]topicPartitionAssignment, consumer2AllPotentialPartitions map[string][]topicPartitionAssignment, currentPartitionConsumer map[topicPartitionAssignment]string) []string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:603
	_go_fuzz_dep_.CoverTab[99136]++
												for _, memberID := range sortedCurrentSubscriptions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:604
		_go_fuzz_dep_.CoverTab[99138]++
													if memberAssignmentsIncludeTopicPartition(consumer2AllPotentialPartitions[memberID], partition) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:605
			_go_fuzz_dep_.CoverTab[99139]++
														currentAssignment[memberID] = append(currentAssignment[memberID], partition)
														currentPartitionConsumer[partition] = memberID
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:608
			// _ = "end of CoverTab[99139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:609
			_go_fuzz_dep_.CoverTab[99140]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:609
			// _ = "end of CoverTab[99140]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:609
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:609
		// _ = "end of CoverTab[99138]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:610
	// _ = "end of CoverTab[99136]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:610
	_go_fuzz_dep_.CoverTab[99137]++
												return sortMemberIDsByPartitionAssignments(currentAssignment)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:611
	// _ = "end of CoverTab[99137]"
}

// Deserialize topic partition assignment data to aid with creation of a sticky assignment.
func deserializeTopicPartitionAssignment(userDataBytes []byte) (StickyAssignorUserData, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:615
	_go_fuzz_dep_.CoverTab[99141]++
												userDataV1 := &StickyAssignorUserDataV1{}
												if err := decode(userDataBytes, userDataV1); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:617
		_go_fuzz_dep_.CoverTab[99143]++
													userDataV0 := &StickyAssignorUserDataV0{}
													if err := decode(userDataBytes, userDataV0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:619
			_go_fuzz_dep_.CoverTab[99145]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:620
			// _ = "end of CoverTab[99145]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:621
			_go_fuzz_dep_.CoverTab[99146]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:621
			// _ = "end of CoverTab[99146]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:621
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:621
		// _ = "end of CoverTab[99143]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:621
		_go_fuzz_dep_.CoverTab[99144]++
													return userDataV0, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:622
		// _ = "end of CoverTab[99144]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:623
		_go_fuzz_dep_.CoverTab[99147]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:623
		// _ = "end of CoverTab[99147]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:623
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:623
	// _ = "end of CoverTab[99141]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:623
	_go_fuzz_dep_.CoverTab[99142]++
												return userDataV1, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:624
	// _ = "end of CoverTab[99142]"
}

// filterAssignedPartitions returns a map of consumer group members to their list of previously-assigned topic partitions, limited
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:627
// to those topic partitions currently reported by the Kafka cluster.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:629
func filterAssignedPartitions(currentAssignment map[string][]topicPartitionAssignment, partition2AllPotentialConsumers map[topicPartitionAssignment][]string) map[string][]topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:629
	_go_fuzz_dep_.CoverTab[99148]++
												assignments := deepCopyAssignment(currentAssignment)
												for memberID, partitions := range assignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:631
		_go_fuzz_dep_.CoverTab[99150]++

													i := 0
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:634
			_go_fuzz_dep_.CoverTab[99152]++
														if _, exists := partition2AllPotentialConsumers[partition]; exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:635
				_go_fuzz_dep_.CoverTab[99153]++
															partitions[i] = partition
															i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:637
				// _ = "end of CoverTab[99153]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:638
				_go_fuzz_dep_.CoverTab[99154]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:638
				// _ = "end of CoverTab[99154]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:638
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:638
			// _ = "end of CoverTab[99152]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:639
		// _ = "end of CoverTab[99150]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:639
		_go_fuzz_dep_.CoverTab[99151]++
													assignments[memberID] = partitions[:i]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:640
		// _ = "end of CoverTab[99151]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:641
	// _ = "end of CoverTab[99148]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:641
	_go_fuzz_dep_.CoverTab[99149]++
												return assignments
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:642
	// _ = "end of CoverTab[99149]"
}

func removeTopicPartitionFromMemberAssignments(assignments []topicPartitionAssignment, topic topicPartitionAssignment) []topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:645
	_go_fuzz_dep_.CoverTab[99155]++
												for i, assignment := range assignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:646
		_go_fuzz_dep_.CoverTab[99157]++
													if assignment == topic {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:647
			_go_fuzz_dep_.CoverTab[99158]++
														return append(assignments[:i], assignments[i+1:]...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:648
			// _ = "end of CoverTab[99158]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:649
			_go_fuzz_dep_.CoverTab[99159]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:649
			// _ = "end of CoverTab[99159]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:649
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:649
		// _ = "end of CoverTab[99157]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:650
	// _ = "end of CoverTab[99155]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:650
	_go_fuzz_dep_.CoverTab[99156]++
												return assignments
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:651
	// _ = "end of CoverTab[99156]"
}

func memberAssignmentsIncludeTopicPartition(assignments []topicPartitionAssignment, topic topicPartitionAssignment) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:654
	_go_fuzz_dep_.CoverTab[99160]++
												for _, assignment := range assignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:655
		_go_fuzz_dep_.CoverTab[99162]++
													if assignment == topic {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:656
			_go_fuzz_dep_.CoverTab[99163]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:657
			// _ = "end of CoverTab[99163]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:658
			_go_fuzz_dep_.CoverTab[99164]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:658
			// _ = "end of CoverTab[99164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:658
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:658
		// _ = "end of CoverTab[99162]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:659
	// _ = "end of CoverTab[99160]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:659
	_go_fuzz_dep_.CoverTab[99161]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:660
	// _ = "end of CoverTab[99161]"
}

func sortPartitions(currentAssignment map[string][]topicPartitionAssignment, partitionsWithADifferentPreviousAssignment map[topicPartitionAssignment]consumerGenerationPair, isFreshAssignment bool, partition2AllPotentialConsumers map[topicPartitionAssignment][]string, consumer2AllPotentialPartitions map[string][]topicPartitionAssignment) []topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:663
	_go_fuzz_dep_.CoverTab[99165]++
												unassignedPartitions := make(map[topicPartitionAssignment]bool, len(partition2AllPotentialConsumers))
												for partition := range partition2AllPotentialConsumers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:665
		_go_fuzz_dep_.CoverTab[99168]++
													unassignedPartitions[partition] = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:666
		// _ = "end of CoverTab[99168]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:667
	// _ = "end of CoverTab[99165]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:667
	_go_fuzz_dep_.CoverTab[99166]++

												sortedPartitions := make([]topicPartitionAssignment, 0)
												if !isFreshAssignment && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:670
		_go_fuzz_dep_.CoverTab[99169]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:670
		return areSubscriptionsIdentical(partition2AllPotentialConsumers, consumer2AllPotentialPartitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:670
		// _ = "end of CoverTab[99169]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:670
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:670
		_go_fuzz_dep_.CoverTab[99170]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:674
		assignments := filterAssignedPartitions(currentAssignment, partition2AllPotentialConsumers)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:678
		pq := make(assignmentPriorityQueue, len(assignments))
		i := 0
		for consumerID, consumerAssignments := range assignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:680
			_go_fuzz_dep_.CoverTab[99173]++
														pq[i] = &consumerGroupMember{
				id:		consumerID,
				assignments:	consumerAssignments,
			}
														i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:685
			// _ = "end of CoverTab[99173]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:686
		// _ = "end of CoverTab[99170]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:686
		_go_fuzz_dep_.CoverTab[99171]++
													heap.Init(&pq)

													for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:689
			_go_fuzz_dep_.CoverTab[99174]++

														if pq.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:691
				_go_fuzz_dep_.CoverTab[99177]++
															break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:692
				// _ = "end of CoverTab[99177]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:693
				_go_fuzz_dep_.CoverTab[99178]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:693
				// _ = "end of CoverTab[99178]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:693
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:693
			// _ = "end of CoverTab[99174]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:693
			_go_fuzz_dep_.CoverTab[99175]++
														member := pq[0]

			// partitions that were assigned to a different consumer last time
			var prevPartitionIndex int
			for i, partition := range member.assignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:698
				_go_fuzz_dep_.CoverTab[99179]++
															if _, exists := partitionsWithADifferentPreviousAssignment[partition]; exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:699
					_go_fuzz_dep_.CoverTab[99180]++
																prevPartitionIndex = i
																break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:701
					// _ = "end of CoverTab[99180]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:702
					_go_fuzz_dep_.CoverTab[99181]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:702
					// _ = "end of CoverTab[99181]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:702
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:702
				// _ = "end of CoverTab[99179]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:703
			// _ = "end of CoverTab[99175]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:703
			_go_fuzz_dep_.CoverTab[99176]++

														if len(member.assignments) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:705
				_go_fuzz_dep_.CoverTab[99182]++
															partition := member.assignments[prevPartitionIndex]
															sortedPartitions = append(sortedPartitions, partition)
															delete(unassignedPartitions, partition)
															if prevPartitionIndex == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:709
					_go_fuzz_dep_.CoverTab[99184]++
																member.assignments = member.assignments[1:]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:710
					// _ = "end of CoverTab[99184]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:711
					_go_fuzz_dep_.CoverTab[99185]++
																member.assignments = append(member.assignments[:prevPartitionIndex], member.assignments[prevPartitionIndex+1:]...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:712
					// _ = "end of CoverTab[99185]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:713
				// _ = "end of CoverTab[99182]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:713
				_go_fuzz_dep_.CoverTab[99183]++
															heap.Fix(&pq, 0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:714
				// _ = "end of CoverTab[99183]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:715
				_go_fuzz_dep_.CoverTab[99186]++
															heap.Pop(&pq)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:716
				// _ = "end of CoverTab[99186]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:717
			// _ = "end of CoverTab[99176]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:718
		// _ = "end of CoverTab[99171]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:718
		_go_fuzz_dep_.CoverTab[99172]++

													for partition := range unassignedPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:720
			_go_fuzz_dep_.CoverTab[99187]++
														sortedPartitions = append(sortedPartitions, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:721
			// _ = "end of CoverTab[99187]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:722
		// _ = "end of CoverTab[99172]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:723
		_go_fuzz_dep_.CoverTab[99188]++

													sortedPartitions = sortPartitionsByPotentialConsumerAssignments(partition2AllPotentialConsumers)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:725
		// _ = "end of CoverTab[99188]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:726
	// _ = "end of CoverTab[99166]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:726
	_go_fuzz_dep_.CoverTab[99167]++
												return sortedPartitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:727
	// _ = "end of CoverTab[99167]"
}

func sortMemberIDsByPartitionAssignments(assignments map[string][]topicPartitionAssignment) []string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:730
	_go_fuzz_dep_.CoverTab[99189]++

												sortedMemberIDs := make([]string, 0, len(assignments))
												for memberID := range assignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:733
		_go_fuzz_dep_.CoverTab[99192]++
													sortedMemberIDs = append(sortedMemberIDs, memberID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:734
		// _ = "end of CoverTab[99192]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:735
	// _ = "end of CoverTab[99189]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:735
	_go_fuzz_dep_.CoverTab[99190]++
												sort.SliceStable(sortedMemberIDs, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:736
		_go_fuzz_dep_.CoverTab[99193]++
													ret := len(assignments[sortedMemberIDs[i]]) - len(assignments[sortedMemberIDs[j]])
													if ret == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:738
			_go_fuzz_dep_.CoverTab[99195]++
														return sortedMemberIDs[i] < sortedMemberIDs[j]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:739
			// _ = "end of CoverTab[99195]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:740
			_go_fuzz_dep_.CoverTab[99196]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:740
			// _ = "end of CoverTab[99196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:740
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:740
		// _ = "end of CoverTab[99193]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:740
		_go_fuzz_dep_.CoverTab[99194]++
													return len(assignments[sortedMemberIDs[i]]) < len(assignments[sortedMemberIDs[j]])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:741
		// _ = "end of CoverTab[99194]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:742
	// _ = "end of CoverTab[99190]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:742
	_go_fuzz_dep_.CoverTab[99191]++
												return sortedMemberIDs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:743
	// _ = "end of CoverTab[99191]"
}

func sortPartitionsByPotentialConsumerAssignments(partition2AllPotentialConsumers map[topicPartitionAssignment][]string) []topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:746
	_go_fuzz_dep_.CoverTab[99197]++

												sortedPartionIDs := make([]topicPartitionAssignment, len(partition2AllPotentialConsumers))
												i := 0
												for partition := range partition2AllPotentialConsumers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:750
		_go_fuzz_dep_.CoverTab[99200]++
													sortedPartionIDs[i] = partition
													i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:752
		// _ = "end of CoverTab[99200]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:753
	// _ = "end of CoverTab[99197]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:753
	_go_fuzz_dep_.CoverTab[99198]++
												sort.Slice(sortedPartionIDs, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:754
		_go_fuzz_dep_.CoverTab[99201]++
													if len(partition2AllPotentialConsumers[sortedPartionIDs[i]]) == len(partition2AllPotentialConsumers[sortedPartionIDs[j]]) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:755
			_go_fuzz_dep_.CoverTab[99203]++
														ret := strings.Compare(sortedPartionIDs[i].Topic, sortedPartionIDs[j].Topic)
														if ret == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:757
				_go_fuzz_dep_.CoverTab[99205]++
															return sortedPartionIDs[i].Partition < sortedPartionIDs[j].Partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:758
				// _ = "end of CoverTab[99205]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:759
				_go_fuzz_dep_.CoverTab[99206]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:759
				// _ = "end of CoverTab[99206]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:759
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:759
			// _ = "end of CoverTab[99203]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:759
			_go_fuzz_dep_.CoverTab[99204]++
														return ret < 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:760
			// _ = "end of CoverTab[99204]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:761
			_go_fuzz_dep_.CoverTab[99207]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:761
			// _ = "end of CoverTab[99207]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:761
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:761
		// _ = "end of CoverTab[99201]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:761
		_go_fuzz_dep_.CoverTab[99202]++
													return len(partition2AllPotentialConsumers[sortedPartionIDs[i]]) < len(partition2AllPotentialConsumers[sortedPartionIDs[j]])
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:762
		// _ = "end of CoverTab[99202]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:763
	// _ = "end of CoverTab[99198]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:763
	_go_fuzz_dep_.CoverTab[99199]++
												return sortedPartionIDs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:764
	// _ = "end of CoverTab[99199]"
}

func deepCopyAssignment(assignment map[string][]topicPartitionAssignment) map[string][]topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:767
	_go_fuzz_dep_.CoverTab[99208]++
												m := make(map[string][]topicPartitionAssignment, len(assignment))
												for memberID, subscriptions := range assignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:769
		_go_fuzz_dep_.CoverTab[99210]++
													m[memberID] = append(subscriptions[:0:0], subscriptions...)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:770
		// _ = "end of CoverTab[99210]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:771
	// _ = "end of CoverTab[99208]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:771
	_go_fuzz_dep_.CoverTab[99209]++
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:772
	// _ = "end of CoverTab[99209]"
}

func areSubscriptionsIdentical(partition2AllPotentialConsumers map[topicPartitionAssignment][]string, consumer2AllPotentialPartitions map[string][]topicPartitionAssignment) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:775
	_go_fuzz_dep_.CoverTab[99211]++
												curMembers := make(map[string]int)
												for _, cur := range partition2AllPotentialConsumers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:777
		_go_fuzz_dep_.CoverTab[99214]++
													if len(curMembers) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:778
			_go_fuzz_dep_.CoverTab[99218]++
														for _, curMembersElem := range cur {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:779
				_go_fuzz_dep_.CoverTab[99220]++
															curMembers[curMembersElem]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:780
				// _ = "end of CoverTab[99220]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:781
			// _ = "end of CoverTab[99218]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:781
			_go_fuzz_dep_.CoverTab[99219]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:782
			// _ = "end of CoverTab[99219]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:783
			_go_fuzz_dep_.CoverTab[99221]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:783
			// _ = "end of CoverTab[99221]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:783
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:783
		// _ = "end of CoverTab[99214]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:783
		_go_fuzz_dep_.CoverTab[99215]++

													if len(curMembers) != len(cur) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:785
			_go_fuzz_dep_.CoverTab[99222]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:786
			// _ = "end of CoverTab[99222]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:787
			_go_fuzz_dep_.CoverTab[99223]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:787
			// _ = "end of CoverTab[99223]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:787
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:787
		// _ = "end of CoverTab[99215]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:787
		_go_fuzz_dep_.CoverTab[99216]++

													yMap := make(map[string]int)
													for _, yElem := range cur {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:790
			_go_fuzz_dep_.CoverTab[99224]++
														yMap[yElem]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:791
			// _ = "end of CoverTab[99224]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:792
		// _ = "end of CoverTab[99216]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:792
		_go_fuzz_dep_.CoverTab[99217]++

													for curMembersMapKey, curMembersMapVal := range curMembers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:794
			_go_fuzz_dep_.CoverTab[99225]++
														if yMap[curMembersMapKey] != curMembersMapVal {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:795
				_go_fuzz_dep_.CoverTab[99226]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:796
				// _ = "end of CoverTab[99226]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:797
				_go_fuzz_dep_.CoverTab[99227]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:797
				// _ = "end of CoverTab[99227]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:797
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:797
			// _ = "end of CoverTab[99225]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:798
		// _ = "end of CoverTab[99217]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:799
	// _ = "end of CoverTab[99211]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:799
	_go_fuzz_dep_.CoverTab[99212]++

												curPartitions := make(map[topicPartitionAssignment]int)
												for _, cur := range consumer2AllPotentialPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:802
		_go_fuzz_dep_.CoverTab[99228]++
													if len(curPartitions) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:803
			_go_fuzz_dep_.CoverTab[99232]++
														for _, curPartitionElem := range cur {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:804
				_go_fuzz_dep_.CoverTab[99234]++
															curPartitions[curPartitionElem]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:805
				// _ = "end of CoverTab[99234]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:806
			// _ = "end of CoverTab[99232]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:806
			_go_fuzz_dep_.CoverTab[99233]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:807
			// _ = "end of CoverTab[99233]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:808
			_go_fuzz_dep_.CoverTab[99235]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:808
			// _ = "end of CoverTab[99235]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:808
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:808
		// _ = "end of CoverTab[99228]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:808
		_go_fuzz_dep_.CoverTab[99229]++

													if len(curPartitions) != len(cur) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:810
			_go_fuzz_dep_.CoverTab[99236]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:811
			// _ = "end of CoverTab[99236]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:812
			_go_fuzz_dep_.CoverTab[99237]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:812
			// _ = "end of CoverTab[99237]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:812
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:812
		// _ = "end of CoverTab[99229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:812
		_go_fuzz_dep_.CoverTab[99230]++

													yMap := make(map[topicPartitionAssignment]int)
													for _, yElem := range cur {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:815
			_go_fuzz_dep_.CoverTab[99238]++
														yMap[yElem]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:816
			// _ = "end of CoverTab[99238]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:817
		// _ = "end of CoverTab[99230]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:817
		_go_fuzz_dep_.CoverTab[99231]++

													for curMembersMapKey, curMembersMapVal := range curPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:819
			_go_fuzz_dep_.CoverTab[99239]++
														if yMap[curMembersMapKey] != curMembersMapVal {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:820
				_go_fuzz_dep_.CoverTab[99240]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:821
				// _ = "end of CoverTab[99240]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:822
				_go_fuzz_dep_.CoverTab[99241]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:822
				// _ = "end of CoverTab[99241]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:822
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:822
			// _ = "end of CoverTab[99239]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:823
		// _ = "end of CoverTab[99231]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:824
	// _ = "end of CoverTab[99212]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:824
	_go_fuzz_dep_.CoverTab[99213]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:825
	// _ = "end of CoverTab[99213]"
}

// We need to process subscriptions' user data with each consumer's reported generation in mind
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:828
// higher generations overwrite lower generations in case of a conflict
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:828
// note that a conflict could exist only if user data is for different generations
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:831
func prepopulateCurrentAssignments(members map[string]ConsumerGroupMemberMetadata) (map[string][]topicPartitionAssignment, map[topicPartitionAssignment]consumerGenerationPair, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:831
	_go_fuzz_dep_.CoverTab[99242]++
												currentAssignment := make(map[string][]topicPartitionAssignment)
												prevAssignment := make(map[topicPartitionAssignment]consumerGenerationPair)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:836
	sortedPartitionConsumersByGeneration := make(map[topicPartitionAssignment]map[int]string)
	for memberID, meta := range members {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:837
		_go_fuzz_dep_.CoverTab[99245]++
													consumerUserData, err := deserializeTopicPartitionAssignment(meta.UserData)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:839
			_go_fuzz_dep_.CoverTab[99247]++
														return nil, nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:840
			// _ = "end of CoverTab[99247]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:841
			_go_fuzz_dep_.CoverTab[99248]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:841
			// _ = "end of CoverTab[99248]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:841
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:841
		// _ = "end of CoverTab[99245]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:841
		_go_fuzz_dep_.CoverTab[99246]++
													for _, partition := range consumerUserData.partitions() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:842
			_go_fuzz_dep_.CoverTab[99249]++
														if consumers, exists := sortedPartitionConsumersByGeneration[partition]; exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:843
				_go_fuzz_dep_.CoverTab[99250]++
															if consumerUserData.hasGeneration() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:844
					_go_fuzz_dep_.CoverTab[99251]++
																if _, generationExists := consumers[consumerUserData.generation()]; generationExists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:845
						_go_fuzz_dep_.CoverTab[99252]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:848
						Logger.Printf("Topic %s Partition %d is assigned to multiple consumers following sticky assignment generation %d", partition.Topic, partition.Partition, consumerUserData.generation())
																	continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:849
						// _ = "end of CoverTab[99252]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:850
						_go_fuzz_dep_.CoverTab[99253]++
																	consumers[consumerUserData.generation()] = memberID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:851
						// _ = "end of CoverTab[99253]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:852
					// _ = "end of CoverTab[99251]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:853
					_go_fuzz_dep_.CoverTab[99254]++
																consumers[defaultGeneration] = memberID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:854
					// _ = "end of CoverTab[99254]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:855
				// _ = "end of CoverTab[99250]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:856
				_go_fuzz_dep_.CoverTab[99255]++
															generation := defaultGeneration
															if consumerUserData.hasGeneration() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:858
					_go_fuzz_dep_.CoverTab[99257]++
																generation = consumerUserData.generation()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:859
					// _ = "end of CoverTab[99257]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:860
					_go_fuzz_dep_.CoverTab[99258]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:860
					// _ = "end of CoverTab[99258]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:860
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:860
				// _ = "end of CoverTab[99255]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:860
				_go_fuzz_dep_.CoverTab[99256]++
															sortedPartitionConsumersByGeneration[partition] = map[int]string{generation: memberID}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:861
				// _ = "end of CoverTab[99256]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:862
			// _ = "end of CoverTab[99249]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:863
		// _ = "end of CoverTab[99246]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:864
	// _ = "end of CoverTab[99242]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:864
	_go_fuzz_dep_.CoverTab[99243]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:868
	for partition, consumers := range sortedPartitionConsumersByGeneration {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:868
		_go_fuzz_dep_.CoverTab[99259]++
		// sort consumers by generation in decreasing order
		var generations []int
		for generation := range consumers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:871
			_go_fuzz_dep_.CoverTab[99262]++
														generations = append(generations, generation)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:872
			// _ = "end of CoverTab[99262]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:873
		// _ = "end of CoverTab[99259]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:873
		_go_fuzz_dep_.CoverTab[99260]++
													sort.Sort(sort.Reverse(sort.IntSlice(generations)))

													consumer := consumers[generations[0]]
													if _, exists := currentAssignment[consumer]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:877
			_go_fuzz_dep_.CoverTab[99263]++
														currentAssignment[consumer] = []topicPartitionAssignment{partition}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:878
			// _ = "end of CoverTab[99263]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:879
			_go_fuzz_dep_.CoverTab[99264]++
														currentAssignment[consumer] = append(currentAssignment[consumer], partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:880
			// _ = "end of CoverTab[99264]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:881
		// _ = "end of CoverTab[99260]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:881
		_go_fuzz_dep_.CoverTab[99261]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:884
		if len(generations) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:884
			_go_fuzz_dep_.CoverTab[99265]++
														prevAssignment[partition] = consumerGenerationPair{
				MemberID:	consumers[generations[1]],
				Generation:	generations[1],
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:888
			// _ = "end of CoverTab[99265]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:889
			_go_fuzz_dep_.CoverTab[99266]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:889
			// _ = "end of CoverTab[99266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:889
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:889
		// _ = "end of CoverTab[99261]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:890
	// _ = "end of CoverTab[99243]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:890
	_go_fuzz_dep_.CoverTab[99244]++
												return currentAssignment, prevAssignment, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:891
	// _ = "end of CoverTab[99244]"
}

type consumerGenerationPair struct {
	MemberID	string
	Generation	int
}

// consumerPair represents a pair of Kafka consumer ids involved in a partition reassignment.
type consumerPair struct {
	SrcMemberID	string
	DstMemberID	string
}

// partitionMovements maintains some data structures to simplify lookup of partition movements among consumers.
type partitionMovements struct {
	PartitionMovementsByTopic	map[string]map[consumerPair]map[topicPartitionAssignment]bool
	Movements			map[topicPartitionAssignment]consumerPair
}

func (p *partitionMovements) removeMovementRecordOfPartition(partition topicPartitionAssignment) consumerPair {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:911
	_go_fuzz_dep_.CoverTab[99267]++
												pair := p.Movements[partition]
												delete(p.Movements, partition)

												partitionMovementsForThisTopic := p.PartitionMovementsByTopic[partition.Topic]
												delete(partitionMovementsForThisTopic[pair], partition)
												if len(partitionMovementsForThisTopic[pair]) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:917
		_go_fuzz_dep_.CoverTab[99270]++
													delete(partitionMovementsForThisTopic, pair)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:918
		// _ = "end of CoverTab[99270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:919
		_go_fuzz_dep_.CoverTab[99271]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:919
		// _ = "end of CoverTab[99271]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:919
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:919
	// _ = "end of CoverTab[99267]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:919
	_go_fuzz_dep_.CoverTab[99268]++
												if len(p.PartitionMovementsByTopic[partition.Topic]) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:920
		_go_fuzz_dep_.CoverTab[99272]++
													delete(p.PartitionMovementsByTopic, partition.Topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:921
		// _ = "end of CoverTab[99272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:922
		_go_fuzz_dep_.CoverTab[99273]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:922
		// _ = "end of CoverTab[99273]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:922
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:922
	// _ = "end of CoverTab[99268]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:922
	_go_fuzz_dep_.CoverTab[99269]++
												return pair
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:923
	// _ = "end of CoverTab[99269]"
}

func (p *partitionMovements) addPartitionMovementRecord(partition topicPartitionAssignment, pair consumerPair) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:926
	_go_fuzz_dep_.CoverTab[99274]++
												p.Movements[partition] = pair
												if _, exists := p.PartitionMovementsByTopic[partition.Topic]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:928
		_go_fuzz_dep_.CoverTab[99277]++
													p.PartitionMovementsByTopic[partition.Topic] = make(map[consumerPair]map[topicPartitionAssignment]bool)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:929
		// _ = "end of CoverTab[99277]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:930
		_go_fuzz_dep_.CoverTab[99278]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:930
		// _ = "end of CoverTab[99278]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:930
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:930
	// _ = "end of CoverTab[99274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:930
	_go_fuzz_dep_.CoverTab[99275]++
												partitionMovementsForThisTopic := p.PartitionMovementsByTopic[partition.Topic]
												if _, exists := partitionMovementsForThisTopic[pair]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:932
		_go_fuzz_dep_.CoverTab[99279]++
													partitionMovementsForThisTopic[pair] = make(map[topicPartitionAssignment]bool)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:933
		// _ = "end of CoverTab[99279]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:934
		_go_fuzz_dep_.CoverTab[99280]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:934
		// _ = "end of CoverTab[99280]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:934
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:934
	// _ = "end of CoverTab[99275]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:934
	_go_fuzz_dep_.CoverTab[99276]++
												partitionMovementsForThisTopic[pair][partition] = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:935
	// _ = "end of CoverTab[99276]"
}

func (p *partitionMovements) movePartition(partition topicPartitionAssignment, oldConsumer, newConsumer string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:938
	_go_fuzz_dep_.CoverTab[99281]++
												pair := consumerPair{
		SrcMemberID:	oldConsumer,
		DstMemberID:	newConsumer,
	}
	if _, exists := p.Movements[partition]; exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:943
		_go_fuzz_dep_.CoverTab[99282]++

													existingPair := p.removeMovementRecordOfPartition(partition)
													if existingPair.DstMemberID != oldConsumer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:946
			_go_fuzz_dep_.CoverTab[99284]++
														Logger.Printf("Existing pair DstMemberID %s was not equal to the oldConsumer ID %s", existingPair.DstMemberID, oldConsumer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:947
			// _ = "end of CoverTab[99284]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:948
			_go_fuzz_dep_.CoverTab[99285]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:948
			// _ = "end of CoverTab[99285]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:948
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:948
		// _ = "end of CoverTab[99282]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:948
		_go_fuzz_dep_.CoverTab[99283]++
													if existingPair.SrcMemberID != newConsumer {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:949
			_go_fuzz_dep_.CoverTab[99286]++

														p.addPartitionMovementRecord(partition, consumerPair{
				SrcMemberID:	existingPair.SrcMemberID,
				DstMemberID:	newConsumer,
			})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:954
			// _ = "end of CoverTab[99286]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:955
			_go_fuzz_dep_.CoverTab[99287]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:955
			// _ = "end of CoverTab[99287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:955
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:955
		// _ = "end of CoverTab[99283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:956
		_go_fuzz_dep_.CoverTab[99288]++
													p.addPartitionMovementRecord(partition, pair)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:957
		// _ = "end of CoverTab[99288]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:958
	// _ = "end of CoverTab[99281]"
}

func (p *partitionMovements) getTheActualPartitionToBeMoved(partition topicPartitionAssignment, oldConsumer, newConsumer string) topicPartitionAssignment {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:961
	_go_fuzz_dep_.CoverTab[99289]++
												if _, exists := p.PartitionMovementsByTopic[partition.Topic]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:962
		_go_fuzz_dep_.CoverTab[99294]++
													return partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:963
		// _ = "end of CoverTab[99294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:964
		_go_fuzz_dep_.CoverTab[99295]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:964
		// _ = "end of CoverTab[99295]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:964
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:964
	// _ = "end of CoverTab[99289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:964
	_go_fuzz_dep_.CoverTab[99290]++
												if _, exists := p.Movements[partition]; exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:965
		_go_fuzz_dep_.CoverTab[99296]++

													if oldConsumer != p.Movements[partition].DstMemberID {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:967
			_go_fuzz_dep_.CoverTab[99298]++
														Logger.Printf("Partition movement DstMemberID %s was not equal to the oldConsumer ID %s", p.Movements[partition].DstMemberID, oldConsumer)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:968
			// _ = "end of CoverTab[99298]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:969
			_go_fuzz_dep_.CoverTab[99299]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:969
			// _ = "end of CoverTab[99299]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:969
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:969
		// _ = "end of CoverTab[99296]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:969
		_go_fuzz_dep_.CoverTab[99297]++
													oldConsumer = p.Movements[partition].SrcMemberID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:970
		// _ = "end of CoverTab[99297]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:971
		_go_fuzz_dep_.CoverTab[99300]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:971
		// _ = "end of CoverTab[99300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:971
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:971
	// _ = "end of CoverTab[99290]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:971
	_go_fuzz_dep_.CoverTab[99291]++

												partitionMovementsForThisTopic := p.PartitionMovementsByTopic[partition.Topic]
												reversePair := consumerPair{
		SrcMemberID:	newConsumer,
		DstMemberID:	oldConsumer,
	}
	if _, exists := partitionMovementsForThisTopic[reversePair]; !exists {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:978
		_go_fuzz_dep_.CoverTab[99301]++
													return partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:979
		// _ = "end of CoverTab[99301]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:980
		_go_fuzz_dep_.CoverTab[99302]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:980
		// _ = "end of CoverTab[99302]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:980
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:980
	// _ = "end of CoverTab[99291]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:980
	_go_fuzz_dep_.CoverTab[99292]++
												var reversePairPartition topicPartitionAssignment
												for otherPartition := range partitionMovementsForThisTopic[reversePair] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:982
		_go_fuzz_dep_.CoverTab[99303]++
													reversePairPartition = otherPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:983
		// _ = "end of CoverTab[99303]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:984
	// _ = "end of CoverTab[99292]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:984
	_go_fuzz_dep_.CoverTab[99293]++
												return reversePairPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:985
	// _ = "end of CoverTab[99293]"
}

func (p *partitionMovements) isLinked(src, dst string, pairs []consumerPair, currentPath []string) ([]string, bool) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:988
	_go_fuzz_dep_.CoverTab[99304]++
												if src == dst {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:989
		_go_fuzz_dep_.CoverTab[99309]++
													return currentPath, false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:990
		// _ = "end of CoverTab[99309]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:991
		_go_fuzz_dep_.CoverTab[99310]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:991
		// _ = "end of CoverTab[99310]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:991
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:991
	// _ = "end of CoverTab[99304]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:991
	_go_fuzz_dep_.CoverTab[99305]++
												if len(pairs) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:992
		_go_fuzz_dep_.CoverTab[99311]++
													return currentPath, false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:993
		// _ = "end of CoverTab[99311]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:994
		_go_fuzz_dep_.CoverTab[99312]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:994
		// _ = "end of CoverTab[99312]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:994
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:994
	// _ = "end of CoverTab[99305]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:994
	_go_fuzz_dep_.CoverTab[99306]++
												for _, pair := range pairs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:995
		_go_fuzz_dep_.CoverTab[99313]++
													if src == pair.SrcMemberID && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:996
			_go_fuzz_dep_.CoverTab[99314]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:996
			return dst == pair.DstMemberID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:996
			// _ = "end of CoverTab[99314]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:996
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:996
			_go_fuzz_dep_.CoverTab[99315]++
														currentPath = append(currentPath, src, dst)
														return currentPath, true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:998
			// _ = "end of CoverTab[99315]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:999
			_go_fuzz_dep_.CoverTab[99316]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:999
			// _ = "end of CoverTab[99316]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:999
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:999
		// _ = "end of CoverTab[99313]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1000
	// _ = "end of CoverTab[99306]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1000
	_go_fuzz_dep_.CoverTab[99307]++

												for _, pair := range pairs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1002
		_go_fuzz_dep_.CoverTab[99317]++
													if pair.SrcMemberID == src {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1003
			_go_fuzz_dep_.CoverTab[99318]++

														reducedSet := make([]consumerPair, len(pairs)-1)
														i := 0
														for _, p := range pairs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1007
				_go_fuzz_dep_.CoverTab[99320]++
															if p != pair {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1008
					_go_fuzz_dep_.CoverTab[99321]++
																reducedSet[i] = pair
																i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1010
					// _ = "end of CoverTab[99321]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1011
					_go_fuzz_dep_.CoverTab[99322]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1011
					// _ = "end of CoverTab[99322]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1011
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1011
				// _ = "end of CoverTab[99320]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1012
			// _ = "end of CoverTab[99318]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1012
			_go_fuzz_dep_.CoverTab[99319]++

														currentPath = append(currentPath, pair.SrcMemberID)
														return p.isLinked(pair.DstMemberID, dst, reducedSet, currentPath)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1015
			// _ = "end of CoverTab[99319]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1016
			_go_fuzz_dep_.CoverTab[99323]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1016
			// _ = "end of CoverTab[99323]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1016
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1016
		// _ = "end of CoverTab[99317]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1017
	// _ = "end of CoverTab[99307]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1017
	_go_fuzz_dep_.CoverTab[99308]++
												return currentPath, false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1018
	// _ = "end of CoverTab[99308]"
}

func (p *partitionMovements) in(cycle []string, cycles [][]string) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1021
	_go_fuzz_dep_.CoverTab[99324]++
												superCycle := make([]string, len(cycle)-1)
												for i := 0; i < len(cycle)-1; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1023
		_go_fuzz_dep_.CoverTab[99327]++
													superCycle[i] = cycle[i]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1024
		// _ = "end of CoverTab[99327]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1025
	// _ = "end of CoverTab[99324]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1025
	_go_fuzz_dep_.CoverTab[99325]++
												superCycle = append(superCycle, cycle...)
												for _, foundCycle := range cycles {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1027
		_go_fuzz_dep_.CoverTab[99328]++
													if len(foundCycle) == len(cycle) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1028
			_go_fuzz_dep_.CoverTab[99329]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1028
			return indexOfSubList(superCycle, foundCycle) != -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1028
			// _ = "end of CoverTab[99329]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1028
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1028
			_go_fuzz_dep_.CoverTab[99330]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1029
			// _ = "end of CoverTab[99330]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1030
			_go_fuzz_dep_.CoverTab[99331]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1030
			// _ = "end of CoverTab[99331]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1030
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1030
		// _ = "end of CoverTab[99328]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1031
	// _ = "end of CoverTab[99325]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1031
	_go_fuzz_dep_.CoverTab[99326]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1032
	// _ = "end of CoverTab[99326]"
}

func (p *partitionMovements) hasCycles(pairs []consumerPair) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1035
	_go_fuzz_dep_.CoverTab[99332]++
												cycles := make([][]string, 0)
												for _, pair := range pairs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1037
		_go_fuzz_dep_.CoverTab[99335]++

													reducedPairs := make([]consumerPair, len(pairs)-1)
													i := 0
													for _, p := range pairs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1041
			_go_fuzz_dep_.CoverTab[99337]++
														if p != pair {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1042
				_go_fuzz_dep_.CoverTab[99338]++
															reducedPairs[i] = pair
															i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1044
				// _ = "end of CoverTab[99338]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1045
				_go_fuzz_dep_.CoverTab[99339]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1045
				// _ = "end of CoverTab[99339]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1045
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1045
			// _ = "end of CoverTab[99337]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1046
		// _ = "end of CoverTab[99335]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1046
		_go_fuzz_dep_.CoverTab[99336]++
													if path, linked := p.isLinked(pair.DstMemberID, pair.SrcMemberID, reducedPairs, []string{pair.SrcMemberID}); linked {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1047
			_go_fuzz_dep_.CoverTab[99340]++
														if !p.in(path, cycles) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1048
				_go_fuzz_dep_.CoverTab[99341]++
															cycles = append(cycles, path)
															Logger.Printf("A cycle of length %d was found: %v", len(path)-1, path)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1050
				// _ = "end of CoverTab[99341]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1051
				_go_fuzz_dep_.CoverTab[99342]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1051
				// _ = "end of CoverTab[99342]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1051
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1051
			// _ = "end of CoverTab[99340]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1052
			_go_fuzz_dep_.CoverTab[99343]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1052
			// _ = "end of CoverTab[99343]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1052
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1052
		// _ = "end of CoverTab[99336]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1053
	// _ = "end of CoverTab[99332]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1053
	_go_fuzz_dep_.CoverTab[99333]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1058
	for _, cycle := range cycles {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1058
		_go_fuzz_dep_.CoverTab[99344]++
													if len(cycle) == 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1059
			_go_fuzz_dep_.CoverTab[99345]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1060
			// _ = "end of CoverTab[99345]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1061
			_go_fuzz_dep_.CoverTab[99346]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1061
			// _ = "end of CoverTab[99346]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1061
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1061
		// _ = "end of CoverTab[99344]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1062
	// _ = "end of CoverTab[99333]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1062
	_go_fuzz_dep_.CoverTab[99334]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1063
	// _ = "end of CoverTab[99334]"
}

func (p *partitionMovements) isSticky() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1066
	_go_fuzz_dep_.CoverTab[99347]++
												for topic, movements := range p.PartitionMovementsByTopic {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1067
		_go_fuzz_dep_.CoverTab[99349]++
													movementPairs := make([]consumerPair, len(movements))
													i := 0
													for pair := range movements {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1070
			_go_fuzz_dep_.CoverTab[99351]++
														movementPairs[i] = pair
														i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1072
			// _ = "end of CoverTab[99351]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1073
		// _ = "end of CoverTab[99349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1073
		_go_fuzz_dep_.CoverTab[99350]++
													if p.hasCycles(movementPairs) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1074
			_go_fuzz_dep_.CoverTab[99352]++
														Logger.Printf("Stickiness is violated for topic %s", topic)
														Logger.Printf("Partition movements for this topic occurred among the following consumer pairs: %v", movements)
														return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1077
			// _ = "end of CoverTab[99352]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1078
			_go_fuzz_dep_.CoverTab[99353]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1078
			// _ = "end of CoverTab[99353]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1078
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1078
		// _ = "end of CoverTab[99350]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1079
	// _ = "end of CoverTab[99347]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1079
	_go_fuzz_dep_.CoverTab[99348]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1080
	// _ = "end of CoverTab[99348]"
}

func indexOfSubList(source []string, target []string) int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1083
	_go_fuzz_dep_.CoverTab[99354]++
												targetSize := len(target)
												maxCandidate := len(source) - targetSize
nextCand:
	for candidate := 0; candidate <= maxCandidate; candidate++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1087
		_go_fuzz_dep_.CoverTab[99356]++
													j := candidate
													for i := 0; i < targetSize; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1089
			_go_fuzz_dep_.CoverTab[99358]++
														if target[i] != source[j] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1090
				_go_fuzz_dep_.CoverTab[99360]++

															continue nextCand
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1092
				// _ = "end of CoverTab[99360]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1093
				_go_fuzz_dep_.CoverTab[99361]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1093
				// _ = "end of CoverTab[99361]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1093
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1093
			// _ = "end of CoverTab[99358]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1093
			_go_fuzz_dep_.CoverTab[99359]++
														j++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1094
			// _ = "end of CoverTab[99359]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1095
		// _ = "end of CoverTab[99356]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1095
		_go_fuzz_dep_.CoverTab[99357]++

													return candidate
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1097
		// _ = "end of CoverTab[99357]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1098
	// _ = "end of CoverTab[99354]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1098
	_go_fuzz_dep_.CoverTab[99355]++
												return -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1099
	// _ = "end of CoverTab[99355]"
}

type consumerGroupMember struct {
	id		string
	assignments	[]topicPartitionAssignment
}

// assignmentPriorityQueue is a priority-queue of consumer group members that is sorted
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1107
// in descending order (most assignments to least assignments).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1109
type assignmentPriorityQueue []*consumerGroupMember

func (pq assignmentPriorityQueue) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1111
	_go_fuzz_dep_.CoverTab[99362]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1111
	return len(pq)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1111
	// _ = "end of CoverTab[99362]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1111
}

func (pq assignmentPriorityQueue) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1113
	_go_fuzz_dep_.CoverTab[99363]++

												if len(pq[i].assignments) == len(pq[j].assignments) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1115
		_go_fuzz_dep_.CoverTab[99365]++
													return strings.Compare(pq[i].id, pq[j].id) > 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1116
		// _ = "end of CoverTab[99365]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1117
		_go_fuzz_dep_.CoverTab[99366]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1117
		// _ = "end of CoverTab[99366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1117
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1117
	// _ = "end of CoverTab[99363]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1117
	_go_fuzz_dep_.CoverTab[99364]++
												return len(pq[i].assignments) > len(pq[j].assignments)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1118
	// _ = "end of CoverTab[99364]"
}

func (pq assignmentPriorityQueue) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1121
	_go_fuzz_dep_.CoverTab[99367]++
												pq[i], pq[j] = pq[j], pq[i]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1122
	// _ = "end of CoverTab[99367]"
}

func (pq *assignmentPriorityQueue) Push(x interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1125
	_go_fuzz_dep_.CoverTab[99368]++
												member := x.(*consumerGroupMember)
												*pq = append(*pq, member)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1127
	// _ = "end of CoverTab[99368]"
}

func (pq *assignmentPriorityQueue) Pop() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1130
	_go_fuzz_dep_.CoverTab[99369]++
												old := *pq
												n := len(old)
												member := old[n-1]
												*pq = old[0 : n-1]
												return member
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1135
	// _ = "end of CoverTab[99369]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1136
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/balance_strategy.go:1136
var _ = _go_fuzz_dep_.CoverTab
