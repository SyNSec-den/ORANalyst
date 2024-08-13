//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:1
)

import (
	"hash"
	"hash/fnv"
	"math/rand"
	"time"
)

// Partitioner is anything that, given a Kafka message and a number of partitions indexed [0...numPartitions-1],
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:10
// decides to which partition to send the message. RandomPartitioner, RoundRobinPartitioner and HashPartitioner are provided
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:10
// as simple default implementations.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:13
type Partitioner interface {
	// Partition takes a message and partition count and chooses a partition
	Partition(message *ProducerMessage, numPartitions int32) (int32, error)

	// RequiresConsistency indicates to the user of the partitioner whether the
	// mapping of key->partition is consistent or not. Specifically, if a
	// partitioner requires consistency then it must be allowed to choose from all
	// partitions (even ones known to be unavailable), and its choice must be
	// respected by the caller. The obvious example is the HashPartitioner.
	RequiresConsistency() bool
}

// DynamicConsistencyPartitioner can optionally be implemented by Partitioners
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:25
// in order to allow more flexibility than is originally allowed by the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:25
// RequiresConsistency method in the Partitioner interface. This allows
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:25
// partitioners to require consistency sometimes, but not all times. It's useful
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:25
// for, e.g., the HashPartitioner, which does not require consistency if the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:25
// message key is nil.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:31
type DynamicConsistencyPartitioner interface {
	Partitioner

	// MessageRequiresConsistency is similar to Partitioner.RequiresConsistency,
	// but takes in the message being partitioned so that the partitioner can
	// make a per-message determination.
	MessageRequiresConsistency(message *ProducerMessage) bool
}

// PartitionerConstructor is the type for a function capable of constructing new Partitioners.
type PartitionerConstructor func(topic string) Partitioner

type manualPartitioner struct{}

// HashPartitionerOption lets you modify default values of the partitioner
type HashPartitionerOption func(*hashPartitioner)

// WithAbsFirst means that the partitioner handles absolute values
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:48
// in the same way as the reference Java implementation
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:50
func WithAbsFirst() HashPartitionerOption {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:50
	_go_fuzz_dep_.CoverTab[105548]++
											return func(hp *hashPartitioner) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:51
		_go_fuzz_dep_.CoverTab[105549]++
												hp.referenceAbs = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:52
		// _ = "end of CoverTab[105549]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:53
	// _ = "end of CoverTab[105548]"
}

// WithCustomHashFunction lets you specify what hash function to use for the partitioning
func WithCustomHashFunction(hasher func() hash.Hash32) HashPartitionerOption {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:57
	_go_fuzz_dep_.CoverTab[105550]++
											return func(hp *hashPartitioner) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:58
		_go_fuzz_dep_.CoverTab[105551]++
												hp.hasher = hasher()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:59
		// _ = "end of CoverTab[105551]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:60
	// _ = "end of CoverTab[105550]"
}

// WithCustomFallbackPartitioner lets you specify what HashPartitioner should be used in case a Distribution Key is empty
func WithCustomFallbackPartitioner(randomHP Partitioner) HashPartitionerOption {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:64
	_go_fuzz_dep_.CoverTab[105552]++
											return func(hp *hashPartitioner) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:65
		_go_fuzz_dep_.CoverTab[105553]++
												hp.random = randomHP
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:66
		// _ = "end of CoverTab[105553]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:67
	// _ = "end of CoverTab[105552]"
}

// NewManualPartitioner returns a Partitioner which uses the partition manually set in the provided
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:70
// ProducerMessage's Partition field as the partition to produce to.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:72
func NewManualPartitioner(topic string) Partitioner {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:72
	_go_fuzz_dep_.CoverTab[105554]++
											return new(manualPartitioner)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:73
	// _ = "end of CoverTab[105554]"
}

func (p *manualPartitioner) Partition(message *ProducerMessage, numPartitions int32) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:76
	_go_fuzz_dep_.CoverTab[105555]++
											return message.Partition, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:77
	// _ = "end of CoverTab[105555]"
}

func (p *manualPartitioner) RequiresConsistency() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:80
	_go_fuzz_dep_.CoverTab[105556]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:81
	// _ = "end of CoverTab[105556]"
}

type randomPartitioner struct {
	generator *rand.Rand
}

// NewRandomPartitioner returns a Partitioner which chooses a random partition each time.
func NewRandomPartitioner(topic string) Partitioner {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:89
	_go_fuzz_dep_.CoverTab[105557]++
											p := new(randomPartitioner)
											p.generator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
											return p
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:92
	// _ = "end of CoverTab[105557]"
}

func (p *randomPartitioner) Partition(message *ProducerMessage, numPartitions int32) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:95
	_go_fuzz_dep_.CoverTab[105558]++
											return int32(p.generator.Intn(int(numPartitions))), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:96
	// _ = "end of CoverTab[105558]"
}

func (p *randomPartitioner) RequiresConsistency() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:99
		_go_fuzz_dep_.CoverTab[105559]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:100
	// _ = "end of CoverTab[105559]"
}

type roundRobinPartitioner struct {
	partition int32
}

// NewRoundRobinPartitioner returns a Partitioner which walks through the available partitions one at a time.
func NewRoundRobinPartitioner(topic string) Partitioner {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:108
	_go_fuzz_dep_.CoverTab[105560]++
												return &roundRobinPartitioner{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:109
	// _ = "end of CoverTab[105560]"
}

func (p *roundRobinPartitioner) Partition(message *ProducerMessage, numPartitions int32) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:112
	_go_fuzz_dep_.CoverTab[105561]++
												if p.partition >= numPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:113
		_go_fuzz_dep_.CoverTab[105563]++
													p.partition = 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:114
		// _ = "end of CoverTab[105563]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:115
		_go_fuzz_dep_.CoverTab[105564]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:115
		// _ = "end of CoverTab[105564]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:115
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:115
	// _ = "end of CoverTab[105561]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:115
	_go_fuzz_dep_.CoverTab[105562]++
												ret := p.partition
												p.partition++
												return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:118
	// _ = "end of CoverTab[105562]"
}

func (p *roundRobinPartitioner) RequiresConsistency() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:121
	_go_fuzz_dep_.CoverTab[105565]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:122
	// _ = "end of CoverTab[105565]"
}

type hashPartitioner struct {
	random		Partitioner
	hasher		hash.Hash32
	referenceAbs	bool
}

// NewCustomHashPartitioner is a wrapper around NewHashPartitioner, allowing the use of custom hasher.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:131
// The argument is a function providing the instance, implementing the hash.Hash32 interface. This is to ensure that
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:131
// each partition dispatcher gets its own hasher, to avoid concurrency issues by sharing an instance.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:134
func NewCustomHashPartitioner(hasher func() hash.Hash32) PartitionerConstructor {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:134
	_go_fuzz_dep_.CoverTab[105566]++
												return func(topic string) Partitioner {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:135
		_go_fuzz_dep_.CoverTab[105567]++
													p := new(hashPartitioner)
													p.random = NewRandomPartitioner(topic)
													p.hasher = hasher()
													p.referenceAbs = false
													return p
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:140
		// _ = "end of CoverTab[105567]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:141
	// _ = "end of CoverTab[105566]"
}

// NewCustomPartitioner creates a default Partitioner but lets you specify the behavior of each component via options
func NewCustomPartitioner(options ...HashPartitionerOption) PartitionerConstructor {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:145
	_go_fuzz_dep_.CoverTab[105568]++
												return func(topic string) Partitioner {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:146
		_go_fuzz_dep_.CoverTab[105569]++
													p := new(hashPartitioner)
													p.random = NewRandomPartitioner(topic)
													p.hasher = fnv.New32a()
													p.referenceAbs = false
													for _, option := range options {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:151
			_go_fuzz_dep_.CoverTab[105571]++
														option(p)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:152
			// _ = "end of CoverTab[105571]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:153
		// _ = "end of CoverTab[105569]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:153
		_go_fuzz_dep_.CoverTab[105570]++
													return p
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:154
		// _ = "end of CoverTab[105570]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:155
	// _ = "end of CoverTab[105568]"
}

// NewHashPartitioner returns a Partitioner which behaves as follows. If the message's key is nil then a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:158
// random partition is chosen. Otherwise the FNV-1a hash of the encoded bytes of the message key is used,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:158
// modulus the number of partitions. This ensures that messages with the same key always end up on the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:158
// same partition.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:162
func NewHashPartitioner(topic string) Partitioner {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:162
	_go_fuzz_dep_.CoverTab[105572]++
												p := new(hashPartitioner)
												p.random = NewRandomPartitioner(topic)
												p.hasher = fnv.New32a()
												p.referenceAbs = false
												return p
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:167
	// _ = "end of CoverTab[105572]"
}

// NewReferenceHashPartitioner is like NewHashPartitioner except that it handles absolute values
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:170
// in the same way as the reference Java implementation. NewHashPartitioner was supposed to do
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:170
// that but it had a mistake and now there are people depending on both behaviors. This will
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:170
// all go away on the next major version bump.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:174
func NewReferenceHashPartitioner(topic string) Partitioner {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:174
	_go_fuzz_dep_.CoverTab[105573]++
												p := new(hashPartitioner)
												p.random = NewRandomPartitioner(topic)
												p.hasher = fnv.New32a()
												p.referenceAbs = true
												return p
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:179
	// _ = "end of CoverTab[105573]"
}

func (p *hashPartitioner) Partition(message *ProducerMessage, numPartitions int32) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:182
	_go_fuzz_dep_.CoverTab[105574]++
												if message.Key == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:183
		_go_fuzz_dep_.CoverTab[105579]++
													return p.random.Partition(message, numPartitions)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:184
		// _ = "end of CoverTab[105579]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:185
		_go_fuzz_dep_.CoverTab[105580]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:185
		// _ = "end of CoverTab[105580]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:185
	// _ = "end of CoverTab[105574]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:185
	_go_fuzz_dep_.CoverTab[105575]++
												bytes, err := message.Key.Encode()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:187
		_go_fuzz_dep_.CoverTab[105581]++
													return -1, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:188
		// _ = "end of CoverTab[105581]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:189
		_go_fuzz_dep_.CoverTab[105582]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:189
		// _ = "end of CoverTab[105582]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:189
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:189
	// _ = "end of CoverTab[105575]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:189
	_go_fuzz_dep_.CoverTab[105576]++
												p.hasher.Reset()
												_, err = p.hasher.Write(bytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:192
		_go_fuzz_dep_.CoverTab[105583]++
													return -1, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:193
		// _ = "end of CoverTab[105583]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:194
		_go_fuzz_dep_.CoverTab[105584]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:194
		// _ = "end of CoverTab[105584]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:194
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:194
	// _ = "end of CoverTab[105576]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:194
	_go_fuzz_dep_.CoverTab[105577]++
												var partition int32

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:200
	if p.referenceAbs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:200
		_go_fuzz_dep_.CoverTab[105585]++
													partition = (int32(p.hasher.Sum32()) & 0x7fffffff) % numPartitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:201
		// _ = "end of CoverTab[105585]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:202
		_go_fuzz_dep_.CoverTab[105586]++
													partition = int32(p.hasher.Sum32()) % numPartitions
													if partition < 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:204
			_go_fuzz_dep_.CoverTab[105587]++
														partition = -partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:205
			// _ = "end of CoverTab[105587]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:206
			_go_fuzz_dep_.CoverTab[105588]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:206
			// _ = "end of CoverTab[105588]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:206
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:206
		// _ = "end of CoverTab[105586]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:207
	// _ = "end of CoverTab[105577]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:207
	_go_fuzz_dep_.CoverTab[105578]++
												return partition, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:208
	// _ = "end of CoverTab[105578]"
}

func (p *hashPartitioner) RequiresConsistency() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:211
	_go_fuzz_dep_.CoverTab[105589]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:212
	// _ = "end of CoverTab[105589]"
}

func (p *hashPartitioner) MessageRequiresConsistency(message *ProducerMessage) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:215
	_go_fuzz_dep_.CoverTab[105590]++
												return message.Key != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:216
	// _ = "end of CoverTab[105590]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:217
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/partitioner.go:217
var _ = _go_fuzz_dep_.CoverTab
