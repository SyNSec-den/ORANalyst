//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:1
)

type PartitionMetadata struct {
	Err		KError
	ID		int32
	Leader		int32
	Replicas	[]int32
	Isr		[]int32
	OfflineReplicas	[]int32
}

func (pm *PartitionMetadata) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:12
	_go_fuzz_dep_.CoverTab[104097]++
												tmp, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:14
		_go_fuzz_dep_.CoverTab[104104]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:15
		// _ = "end of CoverTab[104104]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:16
		_go_fuzz_dep_.CoverTab[104105]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:16
		// _ = "end of CoverTab[104105]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:16
	// _ = "end of CoverTab[104097]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:16
	_go_fuzz_dep_.CoverTab[104098]++
												pm.Err = KError(tmp)

												pm.ID, err = pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:20
		_go_fuzz_dep_.CoverTab[104106]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:21
		// _ = "end of CoverTab[104106]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:22
		_go_fuzz_dep_.CoverTab[104107]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:22
		// _ = "end of CoverTab[104107]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:22
	// _ = "end of CoverTab[104098]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:22
	_go_fuzz_dep_.CoverTab[104099]++

												pm.Leader, err = pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:25
		_go_fuzz_dep_.CoverTab[104108]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:26
		// _ = "end of CoverTab[104108]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:27
		_go_fuzz_dep_.CoverTab[104109]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:27
		// _ = "end of CoverTab[104109]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:27
	// _ = "end of CoverTab[104099]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:27
	_go_fuzz_dep_.CoverTab[104100]++

												pm.Replicas, err = pd.getInt32Array()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:30
		_go_fuzz_dep_.CoverTab[104110]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:31
		// _ = "end of CoverTab[104110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:32
		_go_fuzz_dep_.CoverTab[104111]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:32
		// _ = "end of CoverTab[104111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:32
	// _ = "end of CoverTab[104100]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:32
	_go_fuzz_dep_.CoverTab[104101]++

												pm.Isr, err = pd.getInt32Array()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:35
		_go_fuzz_dep_.CoverTab[104112]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:36
		// _ = "end of CoverTab[104112]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:37
		_go_fuzz_dep_.CoverTab[104113]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:37
		// _ = "end of CoverTab[104113]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:37
	// _ = "end of CoverTab[104101]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:37
	_go_fuzz_dep_.CoverTab[104102]++

												if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:39
		_go_fuzz_dep_.CoverTab[104114]++
													pm.OfflineReplicas, err = pd.getInt32Array()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:41
			_go_fuzz_dep_.CoverTab[104115]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:42
			// _ = "end of CoverTab[104115]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:43
			_go_fuzz_dep_.CoverTab[104116]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:43
			// _ = "end of CoverTab[104116]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:43
		// _ = "end of CoverTab[104114]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:44
		_go_fuzz_dep_.CoverTab[104117]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:44
		// _ = "end of CoverTab[104117]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:44
	// _ = "end of CoverTab[104102]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:44
	_go_fuzz_dep_.CoverTab[104103]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:46
	// _ = "end of CoverTab[104103]"
}

func (pm *PartitionMetadata) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:49
	_go_fuzz_dep_.CoverTab[104118]++
												pe.putInt16(int16(pm.Err))
												pe.putInt32(pm.ID)
												pe.putInt32(pm.Leader)

												err = pe.putInt32Array(pm.Replicas)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:55
		_go_fuzz_dep_.CoverTab[104122]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:56
		// _ = "end of CoverTab[104122]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:57
		_go_fuzz_dep_.CoverTab[104123]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:57
		// _ = "end of CoverTab[104123]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:57
	// _ = "end of CoverTab[104118]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:57
	_go_fuzz_dep_.CoverTab[104119]++

												err = pe.putInt32Array(pm.Isr)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:60
		_go_fuzz_dep_.CoverTab[104124]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:61
		// _ = "end of CoverTab[104124]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:62
		_go_fuzz_dep_.CoverTab[104125]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:62
		// _ = "end of CoverTab[104125]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:62
	// _ = "end of CoverTab[104119]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:62
	_go_fuzz_dep_.CoverTab[104120]++

												if version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:64
		_go_fuzz_dep_.CoverTab[104126]++
													err = pe.putInt32Array(pm.OfflineReplicas)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:66
			_go_fuzz_dep_.CoverTab[104127]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:67
			// _ = "end of CoverTab[104127]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:68
			_go_fuzz_dep_.CoverTab[104128]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:68
			// _ = "end of CoverTab[104128]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:68
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:68
		// _ = "end of CoverTab[104126]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:69
		_go_fuzz_dep_.CoverTab[104129]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:69
		// _ = "end of CoverTab[104129]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:69
	// _ = "end of CoverTab[104120]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:69
	_go_fuzz_dep_.CoverTab[104121]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:71
	// _ = "end of CoverTab[104121]"
}

type TopicMetadata struct {
	Err		KError
	Name		string
	IsInternal	bool	// Only valid for Version >= 1
	Partitions	[]*PartitionMetadata
}

func (tm *TopicMetadata) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:81
	_go_fuzz_dep_.CoverTab[104130]++
												tmp, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:83
		_go_fuzz_dep_.CoverTab[104136]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:84
		// _ = "end of CoverTab[104136]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:85
		_go_fuzz_dep_.CoverTab[104137]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:85
		// _ = "end of CoverTab[104137]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:85
	// _ = "end of CoverTab[104130]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:85
	_go_fuzz_dep_.CoverTab[104131]++
												tm.Err = KError(tmp)

												tm.Name, err = pd.getString()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:89
		_go_fuzz_dep_.CoverTab[104138]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:90
		// _ = "end of CoverTab[104138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:91
		_go_fuzz_dep_.CoverTab[104139]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:91
		// _ = "end of CoverTab[104139]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:91
	// _ = "end of CoverTab[104131]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:91
	_go_fuzz_dep_.CoverTab[104132]++

												if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:93
		_go_fuzz_dep_.CoverTab[104140]++
													tm.IsInternal, err = pd.getBool()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:95
			_go_fuzz_dep_.CoverTab[104141]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:96
			// _ = "end of CoverTab[104141]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:97
			_go_fuzz_dep_.CoverTab[104142]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:97
			// _ = "end of CoverTab[104142]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:97
		// _ = "end of CoverTab[104140]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:98
		_go_fuzz_dep_.CoverTab[104143]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:98
		// _ = "end of CoverTab[104143]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:98
	// _ = "end of CoverTab[104132]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:98
	_go_fuzz_dep_.CoverTab[104133]++

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:101
		_go_fuzz_dep_.CoverTab[104144]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:102
		// _ = "end of CoverTab[104144]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:103
		_go_fuzz_dep_.CoverTab[104145]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:103
		// _ = "end of CoverTab[104145]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:103
	// _ = "end of CoverTab[104133]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:103
	_go_fuzz_dep_.CoverTab[104134]++
												tm.Partitions = make([]*PartitionMetadata, n)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:105
		_go_fuzz_dep_.CoverTab[104146]++
													tm.Partitions[i] = new(PartitionMetadata)
													err = tm.Partitions[i].decode(pd, version)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:108
			_go_fuzz_dep_.CoverTab[104147]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:109
			// _ = "end of CoverTab[104147]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:110
			_go_fuzz_dep_.CoverTab[104148]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:110
			// _ = "end of CoverTab[104148]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:110
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:110
		// _ = "end of CoverTab[104146]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:111
	// _ = "end of CoverTab[104134]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:111
	_go_fuzz_dep_.CoverTab[104135]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:113
	// _ = "end of CoverTab[104135]"
}

func (tm *TopicMetadata) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:116
	_go_fuzz_dep_.CoverTab[104149]++
												pe.putInt16(int16(tm.Err))

												err = pe.putString(tm.Name)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:120
		_go_fuzz_dep_.CoverTab[104154]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:121
		// _ = "end of CoverTab[104154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:122
		_go_fuzz_dep_.CoverTab[104155]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:122
		// _ = "end of CoverTab[104155]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:122
	// _ = "end of CoverTab[104149]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:122
	_go_fuzz_dep_.CoverTab[104150]++

												if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:124
		_go_fuzz_dep_.CoverTab[104156]++
													pe.putBool(tm.IsInternal)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:125
		// _ = "end of CoverTab[104156]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:126
		_go_fuzz_dep_.CoverTab[104157]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:126
		// _ = "end of CoverTab[104157]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:126
	// _ = "end of CoverTab[104150]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:126
	_go_fuzz_dep_.CoverTab[104151]++

												err = pe.putArrayLength(len(tm.Partitions))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:129
		_go_fuzz_dep_.CoverTab[104158]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:130
		// _ = "end of CoverTab[104158]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:131
		_go_fuzz_dep_.CoverTab[104159]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:131
		// _ = "end of CoverTab[104159]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:131
	// _ = "end of CoverTab[104151]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:131
	_go_fuzz_dep_.CoverTab[104152]++

												for _, pm := range tm.Partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:133
		_go_fuzz_dep_.CoverTab[104160]++
													err = pm.encode(pe, version)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:135
			_go_fuzz_dep_.CoverTab[104161]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:136
			// _ = "end of CoverTab[104161]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:137
			_go_fuzz_dep_.CoverTab[104162]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:137
			// _ = "end of CoverTab[104162]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:137
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:137
		// _ = "end of CoverTab[104160]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:138
	// _ = "end of CoverTab[104152]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:138
	_go_fuzz_dep_.CoverTab[104153]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:140
	// _ = "end of CoverTab[104153]"
}

type MetadataResponse struct {
	Version		int16
	ThrottleTimeMs	int32
	Brokers		[]*Broker
	ClusterID	*string
	ControllerID	int32
	Topics		[]*TopicMetadata
}

func (r *MetadataResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:152
	_go_fuzz_dep_.CoverTab[104163]++
												r.Version = version

												if version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:155
		_go_fuzz_dep_.CoverTab[104171]++
													r.ThrottleTimeMs, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:157
			_go_fuzz_dep_.CoverTab[104172]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:158
			// _ = "end of CoverTab[104172]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:159
			_go_fuzz_dep_.CoverTab[104173]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:159
			// _ = "end of CoverTab[104173]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:159
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:159
		// _ = "end of CoverTab[104171]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:160
		_go_fuzz_dep_.CoverTab[104174]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:160
		// _ = "end of CoverTab[104174]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:160
	// _ = "end of CoverTab[104163]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:160
	_go_fuzz_dep_.CoverTab[104164]++

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:163
		_go_fuzz_dep_.CoverTab[104175]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:164
		// _ = "end of CoverTab[104175]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:165
		_go_fuzz_dep_.CoverTab[104176]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:165
		// _ = "end of CoverTab[104176]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:165
	// _ = "end of CoverTab[104164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:165
	_go_fuzz_dep_.CoverTab[104165]++

												r.Brokers = make([]*Broker, n)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:168
		_go_fuzz_dep_.CoverTab[104177]++
													r.Brokers[i] = new(Broker)
													err = r.Brokers[i].decode(pd, version)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:171
			_go_fuzz_dep_.CoverTab[104178]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:172
			// _ = "end of CoverTab[104178]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:173
			_go_fuzz_dep_.CoverTab[104179]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:173
			// _ = "end of CoverTab[104179]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:173
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:173
		// _ = "end of CoverTab[104177]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:174
	// _ = "end of CoverTab[104165]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:174
	_go_fuzz_dep_.CoverTab[104166]++

												if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:176
		_go_fuzz_dep_.CoverTab[104180]++
													r.ClusterID, err = pd.getNullableString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:178
			_go_fuzz_dep_.CoverTab[104181]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:179
			// _ = "end of CoverTab[104181]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:180
			_go_fuzz_dep_.CoverTab[104182]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:180
			// _ = "end of CoverTab[104182]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:180
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:180
		// _ = "end of CoverTab[104180]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:181
		_go_fuzz_dep_.CoverTab[104183]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:181
		// _ = "end of CoverTab[104183]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:181
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:181
	// _ = "end of CoverTab[104166]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:181
	_go_fuzz_dep_.CoverTab[104167]++

												if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:183
		_go_fuzz_dep_.CoverTab[104184]++
													r.ControllerID, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:185
			_go_fuzz_dep_.CoverTab[104185]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:186
			// _ = "end of CoverTab[104185]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:187
			_go_fuzz_dep_.CoverTab[104186]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:187
			// _ = "end of CoverTab[104186]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:187
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:187
		// _ = "end of CoverTab[104184]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:188
		_go_fuzz_dep_.CoverTab[104187]++
													r.ControllerID = -1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:189
		// _ = "end of CoverTab[104187]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:190
	// _ = "end of CoverTab[104167]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:190
	_go_fuzz_dep_.CoverTab[104168]++

												n, err = pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:193
		_go_fuzz_dep_.CoverTab[104188]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:194
		// _ = "end of CoverTab[104188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:195
		_go_fuzz_dep_.CoverTab[104189]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:195
		// _ = "end of CoverTab[104189]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:195
	// _ = "end of CoverTab[104168]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:195
	_go_fuzz_dep_.CoverTab[104169]++

												r.Topics = make([]*TopicMetadata, n)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:198
		_go_fuzz_dep_.CoverTab[104190]++
													r.Topics[i] = new(TopicMetadata)
													err = r.Topics[i].decode(pd, version)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:201
			_go_fuzz_dep_.CoverTab[104191]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:202
			// _ = "end of CoverTab[104191]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:203
			_go_fuzz_dep_.CoverTab[104192]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:203
			// _ = "end of CoverTab[104192]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:203
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:203
		// _ = "end of CoverTab[104190]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:204
	// _ = "end of CoverTab[104169]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:204
	_go_fuzz_dep_.CoverTab[104170]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:206
	// _ = "end of CoverTab[104170]"
}

func (r *MetadataResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:209
	_go_fuzz_dep_.CoverTab[104193]++
												if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:210
		_go_fuzz_dep_.CoverTab[104201]++
													pe.putInt32(r.ThrottleTimeMs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:211
		// _ = "end of CoverTab[104201]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:212
		_go_fuzz_dep_.CoverTab[104202]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:212
		// _ = "end of CoverTab[104202]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:212
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:212
	// _ = "end of CoverTab[104193]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:212
	_go_fuzz_dep_.CoverTab[104194]++

												err := pe.putArrayLength(len(r.Brokers))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:215
		_go_fuzz_dep_.CoverTab[104203]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:216
		// _ = "end of CoverTab[104203]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:217
		_go_fuzz_dep_.CoverTab[104204]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:217
		// _ = "end of CoverTab[104204]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:217
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:217
	// _ = "end of CoverTab[104194]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:217
	_go_fuzz_dep_.CoverTab[104195]++
												for _, broker := range r.Brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:218
		_go_fuzz_dep_.CoverTab[104205]++
													err = broker.encode(pe, r.Version)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:220
			_go_fuzz_dep_.CoverTab[104206]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:221
			// _ = "end of CoverTab[104206]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:222
			_go_fuzz_dep_.CoverTab[104207]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:222
			// _ = "end of CoverTab[104207]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:222
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:222
		// _ = "end of CoverTab[104205]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:223
	// _ = "end of CoverTab[104195]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:223
	_go_fuzz_dep_.CoverTab[104196]++

												if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:225
		_go_fuzz_dep_.CoverTab[104208]++
													err := pe.putNullableString(r.ClusterID)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:227
			_go_fuzz_dep_.CoverTab[104209]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:228
			// _ = "end of CoverTab[104209]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:229
			_go_fuzz_dep_.CoverTab[104210]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:229
			// _ = "end of CoverTab[104210]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:229
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:229
		// _ = "end of CoverTab[104208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:230
		_go_fuzz_dep_.CoverTab[104211]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:230
		// _ = "end of CoverTab[104211]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:230
	// _ = "end of CoverTab[104196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:230
	_go_fuzz_dep_.CoverTab[104197]++

												if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:232
		_go_fuzz_dep_.CoverTab[104212]++
													pe.putInt32(r.ControllerID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:233
		// _ = "end of CoverTab[104212]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:234
		_go_fuzz_dep_.CoverTab[104213]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:234
		// _ = "end of CoverTab[104213]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:234
	// _ = "end of CoverTab[104197]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:234
	_go_fuzz_dep_.CoverTab[104198]++

												err = pe.putArrayLength(len(r.Topics))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:237
		_go_fuzz_dep_.CoverTab[104214]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:238
		// _ = "end of CoverTab[104214]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:239
		_go_fuzz_dep_.CoverTab[104215]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:239
		// _ = "end of CoverTab[104215]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:239
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:239
	// _ = "end of CoverTab[104198]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:239
	_go_fuzz_dep_.CoverTab[104199]++
												for _, tm := range r.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:240
		_go_fuzz_dep_.CoverTab[104216]++
													err = tm.encode(pe, r.Version)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:242
			_go_fuzz_dep_.CoverTab[104217]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:243
			// _ = "end of CoverTab[104217]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:244
			_go_fuzz_dep_.CoverTab[104218]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:244
			// _ = "end of CoverTab[104218]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:244
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:244
		// _ = "end of CoverTab[104216]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:245
	// _ = "end of CoverTab[104199]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:245
	_go_fuzz_dep_.CoverTab[104200]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:247
	// _ = "end of CoverTab[104200]"
}

func (r *MetadataResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:250
	_go_fuzz_dep_.CoverTab[104219]++
												return 3
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:251
	// _ = "end of CoverTab[104219]"
}

func (r *MetadataResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:254
	_go_fuzz_dep_.CoverTab[104220]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:255
	// _ = "end of CoverTab[104220]"
}

func (r *MetadataResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:258
	_go_fuzz_dep_.CoverTab[104221]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:259
	// _ = "end of CoverTab[104221]"
}

func (r *MetadataResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:262
	_go_fuzz_dep_.CoverTab[104222]++
												switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:264
		_go_fuzz_dep_.CoverTab[104223]++
													return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:265
		// _ = "end of CoverTab[104223]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:266
		_go_fuzz_dep_.CoverTab[104224]++
													return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:267
		// _ = "end of CoverTab[104224]"
	case 3, 4:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:268
		_go_fuzz_dep_.CoverTab[104225]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:269
		// _ = "end of CoverTab[104225]"
	case 5:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:270
		_go_fuzz_dep_.CoverTab[104226]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:271
		// _ = "end of CoverTab[104226]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:272
		_go_fuzz_dep_.CoverTab[104227]++
													return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:273
		// _ = "end of CoverTab[104227]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:274
	// _ = "end of CoverTab[104222]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:279
func (r *MetadataResponse) AddBroker(addr string, id int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:279
	_go_fuzz_dep_.CoverTab[104228]++
												r.Brokers = append(r.Brokers, &Broker{id: id, addr: addr})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:280
	// _ = "end of CoverTab[104228]"
}

func (r *MetadataResponse) AddTopic(topic string, err KError) *TopicMetadata {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:283
	_go_fuzz_dep_.CoverTab[104229]++
												var tmatch *TopicMetadata

												for _, tm := range r.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:286
		_go_fuzz_dep_.CoverTab[104231]++
													if tm.Name == topic {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:287
			_go_fuzz_dep_.CoverTab[104232]++
														tmatch = tm
														goto foundTopic
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:289
			// _ = "end of CoverTab[104232]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:290
			_go_fuzz_dep_.CoverTab[104233]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:290
			// _ = "end of CoverTab[104233]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:290
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:290
		// _ = "end of CoverTab[104231]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:291
	// _ = "end of CoverTab[104229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:291
	_go_fuzz_dep_.CoverTab[104230]++

												tmatch = new(TopicMetadata)
												tmatch.Name = topic
												r.Topics = append(r.Topics, tmatch)

foundTopic:

	tmatch.Err = err
												return tmatch
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:300
	// _ = "end of CoverTab[104230]"
}

func (r *MetadataResponse) AddTopicPartition(topic string, partition, brokerID int32, replicas, isr []int32, offline []int32, err KError) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:303
	_go_fuzz_dep_.CoverTab[104234]++
												tmatch := r.AddTopic(topic, ErrNoError)
												var pmatch *PartitionMetadata

												for _, pm := range tmatch.Partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:307
		_go_fuzz_dep_.CoverTab[104236]++
													if pm.ID == partition {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:308
			_go_fuzz_dep_.CoverTab[104237]++
														pmatch = pm
														goto foundPartition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:310
			// _ = "end of CoverTab[104237]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:311
			_go_fuzz_dep_.CoverTab[104238]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:311
			// _ = "end of CoverTab[104238]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:311
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:311
		// _ = "end of CoverTab[104236]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:312
	// _ = "end of CoverTab[104234]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:312
	_go_fuzz_dep_.CoverTab[104235]++

												pmatch = new(PartitionMetadata)
												pmatch.ID = partition
												tmatch.Partitions = append(tmatch.Partitions, pmatch)

foundPartition:

	pmatch.Leader = brokerID
												pmatch.Replicas = replicas
												pmatch.Isr = isr
												pmatch.OfflineReplicas = offline
												pmatch.Err = err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:324
	// _ = "end of CoverTab[104235]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:325
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metadata_response.go:325
var _ = _go_fuzz_dep_.CoverTab
