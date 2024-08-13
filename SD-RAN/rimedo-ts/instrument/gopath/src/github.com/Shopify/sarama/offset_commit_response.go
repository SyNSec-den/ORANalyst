//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:1
)

type OffsetCommitResponse struct {
	Version		int16
	ThrottleTimeMs	int32
	Errors		map[string]map[int32]KError
}

func (r *OffsetCommitResponse) AddError(topic string, partition int32, kerror KError) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:9
	_go_fuzz_dep_.CoverTab[104865]++
													if r.Errors == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:10
		_go_fuzz_dep_.CoverTab[104868]++
														r.Errors = make(map[string]map[int32]KError)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:11
		// _ = "end of CoverTab[104868]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:12
		_go_fuzz_dep_.CoverTab[104869]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:12
		// _ = "end of CoverTab[104869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:12
	// _ = "end of CoverTab[104865]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:12
	_go_fuzz_dep_.CoverTab[104866]++
													partitions := r.Errors[topic]
													if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:14
		_go_fuzz_dep_.CoverTab[104870]++
														partitions = make(map[int32]KError)
														r.Errors[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:16
		// _ = "end of CoverTab[104870]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:17
		_go_fuzz_dep_.CoverTab[104871]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:17
		// _ = "end of CoverTab[104871]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:17
	// _ = "end of CoverTab[104866]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:17
	_go_fuzz_dep_.CoverTab[104867]++
													partitions[partition] = kerror
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:18
	// _ = "end of CoverTab[104867]"
}

func (r *OffsetCommitResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:21
	_go_fuzz_dep_.CoverTab[104872]++
													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:22
		_go_fuzz_dep_.CoverTab[104876]++
														pe.putInt32(r.ThrottleTimeMs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:23
		// _ = "end of CoverTab[104876]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:24
		_go_fuzz_dep_.CoverTab[104877]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:24
		// _ = "end of CoverTab[104877]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:24
	// _ = "end of CoverTab[104872]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:24
	_go_fuzz_dep_.CoverTab[104873]++
													if err := pe.putArrayLength(len(r.Errors)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:25
		_go_fuzz_dep_.CoverTab[104878]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:26
		// _ = "end of CoverTab[104878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:27
		_go_fuzz_dep_.CoverTab[104879]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:27
		// _ = "end of CoverTab[104879]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:27
	// _ = "end of CoverTab[104873]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:27
	_go_fuzz_dep_.CoverTab[104874]++
													for topic, partitions := range r.Errors {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:28
		_go_fuzz_dep_.CoverTab[104880]++
														if err := pe.putString(topic); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:29
			_go_fuzz_dep_.CoverTab[104883]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:30
			// _ = "end of CoverTab[104883]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:31
			_go_fuzz_dep_.CoverTab[104884]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:31
			// _ = "end of CoverTab[104884]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:31
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:31
		// _ = "end of CoverTab[104880]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:31
		_go_fuzz_dep_.CoverTab[104881]++
														if err := pe.putArrayLength(len(partitions)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:32
			_go_fuzz_dep_.CoverTab[104885]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:33
			// _ = "end of CoverTab[104885]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:34
			_go_fuzz_dep_.CoverTab[104886]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:34
			// _ = "end of CoverTab[104886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:34
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:34
		// _ = "end of CoverTab[104881]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:34
		_go_fuzz_dep_.CoverTab[104882]++
														for partition, kerror := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:35
			_go_fuzz_dep_.CoverTab[104887]++
															pe.putInt32(partition)
															pe.putInt16(int16(kerror))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:37
			// _ = "end of CoverTab[104887]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:38
		// _ = "end of CoverTab[104882]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:39
	// _ = "end of CoverTab[104874]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:39
	_go_fuzz_dep_.CoverTab[104875]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:40
	// _ = "end of CoverTab[104875]"
}

func (r *OffsetCommitResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:43
	_go_fuzz_dep_.CoverTab[104888]++
													r.Version = version

													if version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:46
		_go_fuzz_dep_.CoverTab[104892]++
														r.ThrottleTimeMs, err = pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:48
			_go_fuzz_dep_.CoverTab[104893]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:49
			// _ = "end of CoverTab[104893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:50
			_go_fuzz_dep_.CoverTab[104894]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:50
			// _ = "end of CoverTab[104894]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:50
		// _ = "end of CoverTab[104892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:51
		_go_fuzz_dep_.CoverTab[104895]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:51
		// _ = "end of CoverTab[104895]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:51
	// _ = "end of CoverTab[104888]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:51
	_go_fuzz_dep_.CoverTab[104889]++

													numTopics, err := pd.getArrayLength()
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:54
		_go_fuzz_dep_.CoverTab[104896]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:54
		return numTopics == 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:54
		// _ = "end of CoverTab[104896]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:54
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:54
		_go_fuzz_dep_.CoverTab[104897]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:55
		// _ = "end of CoverTab[104897]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:56
		_go_fuzz_dep_.CoverTab[104898]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:56
		// _ = "end of CoverTab[104898]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:56
	// _ = "end of CoverTab[104889]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:56
	_go_fuzz_dep_.CoverTab[104890]++

													r.Errors = make(map[string]map[int32]KError, numTopics)
													for i := 0; i < numTopics; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:59
		_go_fuzz_dep_.CoverTab[104899]++
														name, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:61
			_go_fuzz_dep_.CoverTab[104902]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:62
			// _ = "end of CoverTab[104902]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:63
			_go_fuzz_dep_.CoverTab[104903]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:63
			// _ = "end of CoverTab[104903]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:63
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:63
		// _ = "end of CoverTab[104899]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:63
		_go_fuzz_dep_.CoverTab[104900]++

														numErrors, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:66
			_go_fuzz_dep_.CoverTab[104904]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:67
			// _ = "end of CoverTab[104904]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:68
			_go_fuzz_dep_.CoverTab[104905]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:68
			// _ = "end of CoverTab[104905]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:68
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:68
		// _ = "end of CoverTab[104900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:68
		_go_fuzz_dep_.CoverTab[104901]++

														r.Errors[name] = make(map[int32]KError, numErrors)

														for j := 0; j < numErrors; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:72
			_go_fuzz_dep_.CoverTab[104906]++
															id, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:74
				_go_fuzz_dep_.CoverTab[104909]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:75
				// _ = "end of CoverTab[104909]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:76
				_go_fuzz_dep_.CoverTab[104910]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:76
				// _ = "end of CoverTab[104910]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:76
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:76
			// _ = "end of CoverTab[104906]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:76
			_go_fuzz_dep_.CoverTab[104907]++

															tmp, err := pd.getInt16()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:79
				_go_fuzz_dep_.CoverTab[104911]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:80
				// _ = "end of CoverTab[104911]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:81
				_go_fuzz_dep_.CoverTab[104912]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:81
				// _ = "end of CoverTab[104912]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:81
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:81
			// _ = "end of CoverTab[104907]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:81
			_go_fuzz_dep_.CoverTab[104908]++
															r.Errors[name][id] = KError(tmp)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:82
			// _ = "end of CoverTab[104908]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:83
		// _ = "end of CoverTab[104901]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:84
	// _ = "end of CoverTab[104890]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:84
	_go_fuzz_dep_.CoverTab[104891]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:86
	// _ = "end of CoverTab[104891]"
}

func (r *OffsetCommitResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:89
	_go_fuzz_dep_.CoverTab[104913]++
													return 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:90
	// _ = "end of CoverTab[104913]"
}

func (r *OffsetCommitResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:93
	_go_fuzz_dep_.CoverTab[104914]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:94
	// _ = "end of CoverTab[104914]"
}

func (r *OffsetCommitResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:97
	_go_fuzz_dep_.CoverTab[104915]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:98
	// _ = "end of CoverTab[104915]"
}

func (r *OffsetCommitResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:101
	_go_fuzz_dep_.CoverTab[104916]++
													switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:103
		_go_fuzz_dep_.CoverTab[104917]++
														return V0_8_2_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:104
		// _ = "end of CoverTab[104917]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:105
		_go_fuzz_dep_.CoverTab[104918]++
														return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:106
		// _ = "end of CoverTab[104918]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:107
		_go_fuzz_dep_.CoverTab[104919]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:108
		// _ = "end of CoverTab[104919]"
	case 4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:109
		_go_fuzz_dep_.CoverTab[104920]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:110
		// _ = "end of CoverTab[104920]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:111
		_go_fuzz_dep_.CoverTab[104921]++
														return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:112
		// _ = "end of CoverTab[104921]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:113
	// _ = "end of CoverTab[104916]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:114
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/offset_commit_response.go:114
var _ = _go_fuzz_dep_.CoverTab
