//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:1
)

// DescribeLogDirsRequest is a describe request to get partitions' log size
type DescribeLogDirsRequest struct {
	// Version 0 and 1 are equal
	// The version number is bumped to indicate that on quota violation brokers send out responses before throttling.
	Version	int16

	// If this is an empty array, all topics will be queried
	DescribeTopics	[]DescribeLogDirsRequestTopic
}

// DescribeLogDirsRequestTopic is a describe request about the log dir of one or more partitions within a Topic
type DescribeLogDirsRequestTopic struct {
	Topic		string
	PartitionIDs	[]int32
}

func (r *DescribeLogDirsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:19
	_go_fuzz_dep_.CoverTab[102465]++
													length := len(r.DescribeTopics)
													if length == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:21
		_go_fuzz_dep_.CoverTab[102469]++

														length = -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:23
		// _ = "end of CoverTab[102469]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:24
		_go_fuzz_dep_.CoverTab[102470]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:24
		// _ = "end of CoverTab[102470]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:24
	// _ = "end of CoverTab[102465]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:24
	_go_fuzz_dep_.CoverTab[102466]++

													if err := pe.putArrayLength(length); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:26
		_go_fuzz_dep_.CoverTab[102471]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:27
		// _ = "end of CoverTab[102471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:28
		_go_fuzz_dep_.CoverTab[102472]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:28
		// _ = "end of CoverTab[102472]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:28
	// _ = "end of CoverTab[102466]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:28
	_go_fuzz_dep_.CoverTab[102467]++

													for _, d := range r.DescribeTopics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:30
		_go_fuzz_dep_.CoverTab[102473]++
														if err := pe.putString(d.Topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:31
			_go_fuzz_dep_.CoverTab[102475]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:32
			// _ = "end of CoverTab[102475]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:33
			_go_fuzz_dep_.CoverTab[102476]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:33
			// _ = "end of CoverTab[102476]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:33
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:33
		// _ = "end of CoverTab[102473]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:33
		_go_fuzz_dep_.CoverTab[102474]++

														if err := pe.putInt32Array(d.PartitionIDs); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:35
			_go_fuzz_dep_.CoverTab[102477]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:36
			// _ = "end of CoverTab[102477]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:37
			_go_fuzz_dep_.CoverTab[102478]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:37
			// _ = "end of CoverTab[102478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:37
		// _ = "end of CoverTab[102474]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:38
	// _ = "end of CoverTab[102467]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:38
	_go_fuzz_dep_.CoverTab[102468]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:40
	// _ = "end of CoverTab[102468]"
}

func (r *DescribeLogDirsRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:43
	_go_fuzz_dep_.CoverTab[102479]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:45
		_go_fuzz_dep_.CoverTab[102483]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:46
		// _ = "end of CoverTab[102483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:47
		_go_fuzz_dep_.CoverTab[102484]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:47
		// _ = "end of CoverTab[102484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:47
	// _ = "end of CoverTab[102479]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:47
	_go_fuzz_dep_.CoverTab[102480]++
													if n == -1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:48
		_go_fuzz_dep_.CoverTab[102485]++
														n = 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:49
		// _ = "end of CoverTab[102485]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:50
		_go_fuzz_dep_.CoverTab[102486]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:50
		// _ = "end of CoverTab[102486]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:50
	// _ = "end of CoverTab[102480]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:50
	_go_fuzz_dep_.CoverTab[102481]++

													topics := make([]DescribeLogDirsRequestTopic, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:53
		_go_fuzz_dep_.CoverTab[102487]++
														topics[i] = DescribeLogDirsRequestTopic{}

														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:57
			_go_fuzz_dep_.CoverTab[102490]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:58
			// _ = "end of CoverTab[102490]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:59
			_go_fuzz_dep_.CoverTab[102491]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:59
			// _ = "end of CoverTab[102491]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:59
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:59
		// _ = "end of CoverTab[102487]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:59
		_go_fuzz_dep_.CoverTab[102488]++
														topics[i].Topic = topic

														pIDs, err := pd.getInt32Array()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:63
			_go_fuzz_dep_.CoverTab[102492]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:64
			// _ = "end of CoverTab[102492]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:65
			_go_fuzz_dep_.CoverTab[102493]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:65
			// _ = "end of CoverTab[102493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:65
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:65
		// _ = "end of CoverTab[102488]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:65
		_go_fuzz_dep_.CoverTab[102489]++
														topics[i].PartitionIDs = pIDs
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:66
		// _ = "end of CoverTab[102489]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:67
	// _ = "end of CoverTab[102481]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:67
	_go_fuzz_dep_.CoverTab[102482]++
													r.DescribeTopics = topics

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:70
	// _ = "end of CoverTab[102482]"
}

func (r *DescribeLogDirsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:73
	_go_fuzz_dep_.CoverTab[102494]++
													return 35
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:74
	// _ = "end of CoverTab[102494]"
}

func (r *DescribeLogDirsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:77
	_go_fuzz_dep_.CoverTab[102495]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:78
	// _ = "end of CoverTab[102495]"
}

func (r *DescribeLogDirsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:81
	_go_fuzz_dep_.CoverTab[102496]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:82
	// _ = "end of CoverTab[102496]"
}

func (r *DescribeLogDirsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:85
	_go_fuzz_dep_.CoverTab[102497]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:86
	// _ = "end of CoverTab[102497]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_request.go:87
var _ = _go_fuzz_dep_.CoverTab
