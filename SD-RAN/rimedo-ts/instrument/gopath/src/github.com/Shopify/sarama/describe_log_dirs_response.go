//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:1
)

import "time"

type DescribeLogDirsResponse struct {
	ThrottleTime	time.Duration

	// Version 0 and 1 are equal
	// The version number is bumped to indicate that on quota violation brokers send out responses before throttling.
	Version	int16

	LogDirs	[]DescribeLogDirsResponseDirMetadata
}

func (r *DescribeLogDirsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:15
	_go_fuzz_dep_.CoverTab[102498]++
													pe.putInt32(int32(r.ThrottleTime / time.Millisecond))

													if err := pe.putArrayLength(len(r.LogDirs)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:18
		_go_fuzz_dep_.CoverTab[102501]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:19
		// _ = "end of CoverTab[102501]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:20
		_go_fuzz_dep_.CoverTab[102502]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:20
		// _ = "end of CoverTab[102502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:20
	// _ = "end of CoverTab[102498]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:20
	_go_fuzz_dep_.CoverTab[102499]++

													for _, dir := range r.LogDirs {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:22
		_go_fuzz_dep_.CoverTab[102503]++
														if err := dir.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:23
			_go_fuzz_dep_.CoverTab[102504]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:24
			// _ = "end of CoverTab[102504]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:25
			_go_fuzz_dep_.CoverTab[102505]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:25
			// _ = "end of CoverTab[102505]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:25
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:25
		// _ = "end of CoverTab[102503]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:26
	// _ = "end of CoverTab[102499]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:26
	_go_fuzz_dep_.CoverTab[102500]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:28
	// _ = "end of CoverTab[102500]"
}

func (r *DescribeLogDirsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:31
	_go_fuzz_dep_.CoverTab[102506]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:33
		_go_fuzz_dep_.CoverTab[102510]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:34
		// _ = "end of CoverTab[102510]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:35
		_go_fuzz_dep_.CoverTab[102511]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:35
		// _ = "end of CoverTab[102511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:35
	// _ = "end of CoverTab[102506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:35
	_go_fuzz_dep_.CoverTab[102507]++
													r.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:39
	n, err := pd.getArrayLength()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:40
		_go_fuzz_dep_.CoverTab[102512]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:41
		// _ = "end of CoverTab[102512]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:42
		_go_fuzz_dep_.CoverTab[102513]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:42
		// _ = "end of CoverTab[102513]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:42
	// _ = "end of CoverTab[102507]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:42
	_go_fuzz_dep_.CoverTab[102508]++

													r.LogDirs = make([]DescribeLogDirsResponseDirMetadata, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:45
		_go_fuzz_dep_.CoverTab[102514]++
														dir := DescribeLogDirsResponseDirMetadata{}
														if err := dir.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:47
			_go_fuzz_dep_.CoverTab[102516]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:48
			// _ = "end of CoverTab[102516]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:49
			_go_fuzz_dep_.CoverTab[102517]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:49
			// _ = "end of CoverTab[102517]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:49
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:49
		// _ = "end of CoverTab[102514]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:49
		_go_fuzz_dep_.CoverTab[102515]++
														r.LogDirs[i] = dir
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:50
		// _ = "end of CoverTab[102515]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:51
	// _ = "end of CoverTab[102508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:51
	_go_fuzz_dep_.CoverTab[102509]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:53
	// _ = "end of CoverTab[102509]"
}

func (r *DescribeLogDirsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:56
	_go_fuzz_dep_.CoverTab[102518]++
													return 35
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:57
	// _ = "end of CoverTab[102518]"
}

func (r *DescribeLogDirsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:60
	_go_fuzz_dep_.CoverTab[102519]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:61
	// _ = "end of CoverTab[102519]"
}

func (r *DescribeLogDirsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:64
	_go_fuzz_dep_.CoverTab[102520]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:65
	// _ = "end of CoverTab[102520]"
}

func (r *DescribeLogDirsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:68
	_go_fuzz_dep_.CoverTab[102521]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:69
	// _ = "end of CoverTab[102521]"
}

type DescribeLogDirsResponseDirMetadata struct {
	ErrorCode	KError

	// The absolute log directory path
	Path	string
	Topics	[]DescribeLogDirsResponseTopic
}

func (r *DescribeLogDirsResponseDirMetadata) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:80
	_go_fuzz_dep_.CoverTab[102522]++
													pe.putInt16(int16(r.ErrorCode))

													if err := pe.putString(r.Path); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:83
		_go_fuzz_dep_.CoverTab[102526]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:84
		// _ = "end of CoverTab[102526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:85
		_go_fuzz_dep_.CoverTab[102527]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:85
		// _ = "end of CoverTab[102527]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:85
	// _ = "end of CoverTab[102522]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:85
	_go_fuzz_dep_.CoverTab[102523]++

													if err := pe.putArrayLength(len(r.Topics)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:87
		_go_fuzz_dep_.CoverTab[102528]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:88
		// _ = "end of CoverTab[102528]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:89
		_go_fuzz_dep_.CoverTab[102529]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:89
		// _ = "end of CoverTab[102529]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:89
	// _ = "end of CoverTab[102523]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:89
	_go_fuzz_dep_.CoverTab[102524]++
													for _, topic := range r.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:90
		_go_fuzz_dep_.CoverTab[102530]++
														if err := topic.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:91
			_go_fuzz_dep_.CoverTab[102531]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:92
			// _ = "end of CoverTab[102531]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:93
			_go_fuzz_dep_.CoverTab[102532]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:93
			// _ = "end of CoverTab[102532]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:93
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:93
		// _ = "end of CoverTab[102530]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:94
	// _ = "end of CoverTab[102524]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:94
	_go_fuzz_dep_.CoverTab[102525]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:96
	// _ = "end of CoverTab[102525]"
}

func (r *DescribeLogDirsResponseDirMetadata) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:99
	_go_fuzz_dep_.CoverTab[102533]++
													errCode, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:101
		_go_fuzz_dep_.CoverTab[102538]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:102
		// _ = "end of CoverTab[102538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:103
		_go_fuzz_dep_.CoverTab[102539]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:103
		// _ = "end of CoverTab[102539]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:103
	// _ = "end of CoverTab[102533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:103
	_go_fuzz_dep_.CoverTab[102534]++
													r.ErrorCode = KError(errCode)

													path, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:107
		_go_fuzz_dep_.CoverTab[102540]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:108
		// _ = "end of CoverTab[102540]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:109
		_go_fuzz_dep_.CoverTab[102541]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:109
		// _ = "end of CoverTab[102541]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:109
	// _ = "end of CoverTab[102534]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:109
	_go_fuzz_dep_.CoverTab[102535]++
													r.Path = path

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:113
	n, err := pd.getArrayLength()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:114
		_go_fuzz_dep_.CoverTab[102542]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:115
		// _ = "end of CoverTab[102542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:116
		_go_fuzz_dep_.CoverTab[102543]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:116
		// _ = "end of CoverTab[102543]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:116
	// _ = "end of CoverTab[102535]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:116
	_go_fuzz_dep_.CoverTab[102536]++

													r.Topics = make([]DescribeLogDirsResponseTopic, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:119
		_go_fuzz_dep_.CoverTab[102544]++
														t := DescribeLogDirsResponseTopic{}

														if err := t.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:122
			_go_fuzz_dep_.CoverTab[102546]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:123
			// _ = "end of CoverTab[102546]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:124
			_go_fuzz_dep_.CoverTab[102547]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:124
			// _ = "end of CoverTab[102547]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:124
		// _ = "end of CoverTab[102544]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:124
		_go_fuzz_dep_.CoverTab[102545]++

														r.Topics[i] = t
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:126
		// _ = "end of CoverTab[102545]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:127
	// _ = "end of CoverTab[102536]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:127
	_go_fuzz_dep_.CoverTab[102537]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:129
	// _ = "end of CoverTab[102537]"
}

// DescribeLogDirsResponseTopic contains a topic's partitions descriptions
type DescribeLogDirsResponseTopic struct {
	Topic		string
	Partitions	[]DescribeLogDirsResponsePartition
}

func (r *DescribeLogDirsResponseTopic) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:138
	_go_fuzz_dep_.CoverTab[102548]++
													if err := pe.putString(r.Topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:139
		_go_fuzz_dep_.CoverTab[102552]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:140
		// _ = "end of CoverTab[102552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:141
		_go_fuzz_dep_.CoverTab[102553]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:141
		// _ = "end of CoverTab[102553]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:141
	// _ = "end of CoverTab[102548]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:141
	_go_fuzz_dep_.CoverTab[102549]++

													if err := pe.putArrayLength(len(r.Partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:143
		_go_fuzz_dep_.CoverTab[102554]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:144
		// _ = "end of CoverTab[102554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:145
		_go_fuzz_dep_.CoverTab[102555]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:145
		// _ = "end of CoverTab[102555]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:145
	// _ = "end of CoverTab[102549]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:145
	_go_fuzz_dep_.CoverTab[102550]++
													for _, partition := range r.Partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:146
		_go_fuzz_dep_.CoverTab[102556]++
														if err := partition.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:147
			_go_fuzz_dep_.CoverTab[102557]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:148
			// _ = "end of CoverTab[102557]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:149
			_go_fuzz_dep_.CoverTab[102558]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:149
			// _ = "end of CoverTab[102558]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:149
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:149
		// _ = "end of CoverTab[102556]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:150
	// _ = "end of CoverTab[102550]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:150
	_go_fuzz_dep_.CoverTab[102551]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:152
	// _ = "end of CoverTab[102551]"
}

func (r *DescribeLogDirsResponseTopic) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:155
	_go_fuzz_dep_.CoverTab[102559]++
													t, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:157
		_go_fuzz_dep_.CoverTab[102563]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:158
		// _ = "end of CoverTab[102563]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:159
		_go_fuzz_dep_.CoverTab[102564]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:159
		// _ = "end of CoverTab[102564]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:159
	// _ = "end of CoverTab[102559]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:159
	_go_fuzz_dep_.CoverTab[102560]++
													r.Topic = t

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:163
		_go_fuzz_dep_.CoverTab[102565]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:164
		// _ = "end of CoverTab[102565]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:165
		_go_fuzz_dep_.CoverTab[102566]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:165
		// _ = "end of CoverTab[102566]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:165
	// _ = "end of CoverTab[102560]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:165
	_go_fuzz_dep_.CoverTab[102561]++
													r.Partitions = make([]DescribeLogDirsResponsePartition, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:167
		_go_fuzz_dep_.CoverTab[102567]++
														p := DescribeLogDirsResponsePartition{}
														if err := p.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:169
			_go_fuzz_dep_.CoverTab[102569]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:170
			// _ = "end of CoverTab[102569]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:171
			_go_fuzz_dep_.CoverTab[102570]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:171
			// _ = "end of CoverTab[102570]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:171
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:171
		// _ = "end of CoverTab[102567]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:171
		_go_fuzz_dep_.CoverTab[102568]++
														r.Partitions[i] = p
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:172
		// _ = "end of CoverTab[102568]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:173
	// _ = "end of CoverTab[102561]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:173
	_go_fuzz_dep_.CoverTab[102562]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:175
	// _ = "end of CoverTab[102562]"
}

// DescribeLogDirsResponsePartition describes a partition's log directory
type DescribeLogDirsResponsePartition struct {
	PartitionID	int32

	// The size of the log segments of the partition in bytes.
	Size	int64

	// The lag of the log's LEO w.r.t. partition's HW (if it is the current log for the partition) or
	// current replica's LEO (if it is the future log for the partition)
	OffsetLag	int64

	// True if this log is created by AlterReplicaLogDirsRequest and will replace the current log of
	// the replica in the future.
	IsTemporary	bool
}

func (r *DescribeLogDirsResponsePartition) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:194
	_go_fuzz_dep_.CoverTab[102571]++
													pe.putInt32(r.PartitionID)
													pe.putInt64(r.Size)
													pe.putInt64(r.OffsetLag)
													pe.putBool(r.IsTemporary)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:200
	// _ = "end of CoverTab[102571]"
}

func (r *DescribeLogDirsResponsePartition) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:203
	_go_fuzz_dep_.CoverTab[102572]++
													pID, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:205
		_go_fuzz_dep_.CoverTab[102577]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:206
		// _ = "end of CoverTab[102577]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:207
		_go_fuzz_dep_.CoverTab[102578]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:207
		// _ = "end of CoverTab[102578]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:207
	// _ = "end of CoverTab[102572]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:207
	_go_fuzz_dep_.CoverTab[102573]++
													r.PartitionID = pID

													size, err := pd.getInt64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:211
		_go_fuzz_dep_.CoverTab[102579]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:212
		// _ = "end of CoverTab[102579]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:213
		_go_fuzz_dep_.CoverTab[102580]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:213
		// _ = "end of CoverTab[102580]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:213
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:213
	// _ = "end of CoverTab[102573]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:213
	_go_fuzz_dep_.CoverTab[102574]++
													r.Size = size

													lag, err := pd.getInt64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:217
		_go_fuzz_dep_.CoverTab[102581]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:218
		// _ = "end of CoverTab[102581]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:219
		_go_fuzz_dep_.CoverTab[102582]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:219
		// _ = "end of CoverTab[102582]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:219
	// _ = "end of CoverTab[102574]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:219
	_go_fuzz_dep_.CoverTab[102575]++
													r.OffsetLag = lag

													isTemp, err := pd.getBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:223
		_go_fuzz_dep_.CoverTab[102583]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:224
		// _ = "end of CoverTab[102583]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:225
		_go_fuzz_dep_.CoverTab[102584]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:225
		// _ = "end of CoverTab[102584]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:225
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:225
	// _ = "end of CoverTab[102575]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:225
	_go_fuzz_dep_.CoverTab[102576]++
													r.IsTemporary = isTemp

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:228
	// _ = "end of CoverTab[102576]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:229
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_log_dirs_response.go:229
var _ = _go_fuzz_dep_.CoverTab
