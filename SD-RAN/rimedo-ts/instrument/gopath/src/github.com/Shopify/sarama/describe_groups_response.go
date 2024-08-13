//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:1
)

type DescribeGroupsResponse struct {
	Groups []*GroupDescription
}

func (r *DescribeGroupsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:7
	_go_fuzz_dep_.CoverTab[102359]++
													if err := pe.putArrayLength(len(r.Groups)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:8
		_go_fuzz_dep_.CoverTab[102362]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:9
		// _ = "end of CoverTab[102362]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:10
		_go_fuzz_dep_.CoverTab[102363]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:10
		// _ = "end of CoverTab[102363]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:10
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:10
	// _ = "end of CoverTab[102359]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:10
	_go_fuzz_dep_.CoverTab[102360]++

													for _, groupDescription := range r.Groups {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:12
		_go_fuzz_dep_.CoverTab[102364]++
														if err := groupDescription.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:13
			_go_fuzz_dep_.CoverTab[102365]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:14
			// _ = "end of CoverTab[102365]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:15
			_go_fuzz_dep_.CoverTab[102366]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:15
			// _ = "end of CoverTab[102366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:15
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:15
		// _ = "end of CoverTab[102364]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:16
	// _ = "end of CoverTab[102360]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:16
	_go_fuzz_dep_.CoverTab[102361]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:18
	// _ = "end of CoverTab[102361]"
}

func (r *DescribeGroupsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:21
	_go_fuzz_dep_.CoverTab[102367]++
													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:23
		_go_fuzz_dep_.CoverTab[102370]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:24
		// _ = "end of CoverTab[102370]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:25
		_go_fuzz_dep_.CoverTab[102371]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:25
		// _ = "end of CoverTab[102371]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:25
	// _ = "end of CoverTab[102367]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:25
	_go_fuzz_dep_.CoverTab[102368]++

													r.Groups = make([]*GroupDescription, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:28
		_go_fuzz_dep_.CoverTab[102372]++
														r.Groups[i] = new(GroupDescription)
														if err := r.Groups[i].decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:30
			_go_fuzz_dep_.CoverTab[102373]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:31
			// _ = "end of CoverTab[102373]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:32
			_go_fuzz_dep_.CoverTab[102374]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:32
			// _ = "end of CoverTab[102374]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:32
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:32
		// _ = "end of CoverTab[102372]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:33
	// _ = "end of CoverTab[102368]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:33
	_go_fuzz_dep_.CoverTab[102369]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:35
	// _ = "end of CoverTab[102369]"
}

func (r *DescribeGroupsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:38
	_go_fuzz_dep_.CoverTab[102375]++
													return 15
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:39
	// _ = "end of CoverTab[102375]"
}

func (r *DescribeGroupsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:42
	_go_fuzz_dep_.CoverTab[102376]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:43
	// _ = "end of CoverTab[102376]"
}

func (r *DescribeGroupsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:46
	_go_fuzz_dep_.CoverTab[102377]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:47
	// _ = "end of CoverTab[102377]"
}

func (r *DescribeGroupsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:50
	_go_fuzz_dep_.CoverTab[102378]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:51
	// _ = "end of CoverTab[102378]"
}

type GroupDescription struct {
	Err		KError
	GroupId		string
	State		string
	ProtocolType	string
	Protocol	string
	Members		map[string]*GroupMemberDescription
}

func (gd *GroupDescription) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:63
	_go_fuzz_dep_.CoverTab[102379]++
													pe.putInt16(int16(gd.Err))

													if err := pe.putString(gd.GroupId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:66
		_go_fuzz_dep_.CoverTab[102386]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:67
		// _ = "end of CoverTab[102386]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:68
		_go_fuzz_dep_.CoverTab[102387]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:68
		// _ = "end of CoverTab[102387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:68
	// _ = "end of CoverTab[102379]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:68
	_go_fuzz_dep_.CoverTab[102380]++
													if err := pe.putString(gd.State); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:69
		_go_fuzz_dep_.CoverTab[102388]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:70
		// _ = "end of CoverTab[102388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:71
		_go_fuzz_dep_.CoverTab[102389]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:71
		// _ = "end of CoverTab[102389]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:71
	// _ = "end of CoverTab[102380]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:71
	_go_fuzz_dep_.CoverTab[102381]++
													if err := pe.putString(gd.ProtocolType); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:72
		_go_fuzz_dep_.CoverTab[102390]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:73
		// _ = "end of CoverTab[102390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:74
		_go_fuzz_dep_.CoverTab[102391]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:74
		// _ = "end of CoverTab[102391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:74
	// _ = "end of CoverTab[102381]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:74
	_go_fuzz_dep_.CoverTab[102382]++
													if err := pe.putString(gd.Protocol); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:75
		_go_fuzz_dep_.CoverTab[102392]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:76
		// _ = "end of CoverTab[102392]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:77
		_go_fuzz_dep_.CoverTab[102393]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:77
		// _ = "end of CoverTab[102393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:77
	// _ = "end of CoverTab[102382]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:77
	_go_fuzz_dep_.CoverTab[102383]++

													if err := pe.putArrayLength(len(gd.Members)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:79
		_go_fuzz_dep_.CoverTab[102394]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:80
		// _ = "end of CoverTab[102394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:81
		_go_fuzz_dep_.CoverTab[102395]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:81
		// _ = "end of CoverTab[102395]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:81
	// _ = "end of CoverTab[102383]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:81
	_go_fuzz_dep_.CoverTab[102384]++

													for memberId, groupMemberDescription := range gd.Members {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:83
		_go_fuzz_dep_.CoverTab[102396]++
														if err := pe.putString(memberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:84
			_go_fuzz_dep_.CoverTab[102398]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:85
			// _ = "end of CoverTab[102398]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:86
			_go_fuzz_dep_.CoverTab[102399]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:86
			// _ = "end of CoverTab[102399]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:86
		// _ = "end of CoverTab[102396]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:86
		_go_fuzz_dep_.CoverTab[102397]++
														if err := groupMemberDescription.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:87
			_go_fuzz_dep_.CoverTab[102400]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:88
			// _ = "end of CoverTab[102400]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:89
			_go_fuzz_dep_.CoverTab[102401]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:89
			// _ = "end of CoverTab[102401]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:89
		// _ = "end of CoverTab[102397]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:90
	// _ = "end of CoverTab[102384]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:90
	_go_fuzz_dep_.CoverTab[102385]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:92
	// _ = "end of CoverTab[102385]"
}

func (gd *GroupDescription) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:95
	_go_fuzz_dep_.CoverTab[102402]++
													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:97
		_go_fuzz_dep_.CoverTab[102411]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:98
		// _ = "end of CoverTab[102411]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:99
		_go_fuzz_dep_.CoverTab[102412]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:99
		// _ = "end of CoverTab[102412]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:99
	// _ = "end of CoverTab[102402]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:99
	_go_fuzz_dep_.CoverTab[102403]++

													gd.Err = KError(kerr)

													if gd.GroupId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:103
		_go_fuzz_dep_.CoverTab[102413]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:104
		// _ = "end of CoverTab[102413]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:105
		_go_fuzz_dep_.CoverTab[102414]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:105
		// _ = "end of CoverTab[102414]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:105
	// _ = "end of CoverTab[102403]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:105
	_go_fuzz_dep_.CoverTab[102404]++
													if gd.State, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:106
		_go_fuzz_dep_.CoverTab[102415]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:107
		// _ = "end of CoverTab[102415]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:108
		_go_fuzz_dep_.CoverTab[102416]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:108
		// _ = "end of CoverTab[102416]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:108
	// _ = "end of CoverTab[102404]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:108
	_go_fuzz_dep_.CoverTab[102405]++
													if gd.ProtocolType, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:109
		_go_fuzz_dep_.CoverTab[102417]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:110
		// _ = "end of CoverTab[102417]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:111
		_go_fuzz_dep_.CoverTab[102418]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:111
		// _ = "end of CoverTab[102418]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:111
	// _ = "end of CoverTab[102405]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:111
	_go_fuzz_dep_.CoverTab[102406]++
													if gd.Protocol, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:112
		_go_fuzz_dep_.CoverTab[102419]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:113
		// _ = "end of CoverTab[102419]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:114
		_go_fuzz_dep_.CoverTab[102420]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:114
		// _ = "end of CoverTab[102420]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:114
	// _ = "end of CoverTab[102406]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:114
	_go_fuzz_dep_.CoverTab[102407]++

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:117
		_go_fuzz_dep_.CoverTab[102421]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:118
		// _ = "end of CoverTab[102421]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:119
		_go_fuzz_dep_.CoverTab[102422]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:119
		// _ = "end of CoverTab[102422]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:119
	// _ = "end of CoverTab[102407]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:119
	_go_fuzz_dep_.CoverTab[102408]++
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:120
		_go_fuzz_dep_.CoverTab[102423]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:121
		// _ = "end of CoverTab[102423]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:122
		_go_fuzz_dep_.CoverTab[102424]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:122
		// _ = "end of CoverTab[102424]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:122
	// _ = "end of CoverTab[102408]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:122
	_go_fuzz_dep_.CoverTab[102409]++

													gd.Members = make(map[string]*GroupMemberDescription)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:125
		_go_fuzz_dep_.CoverTab[102425]++
														memberId, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:127
			_go_fuzz_dep_.CoverTab[102427]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:128
			// _ = "end of CoverTab[102427]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:129
			_go_fuzz_dep_.CoverTab[102428]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:129
			// _ = "end of CoverTab[102428]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:129
		// _ = "end of CoverTab[102425]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:129
		_go_fuzz_dep_.CoverTab[102426]++

														gd.Members[memberId] = new(GroupMemberDescription)
														if err := gd.Members[memberId].decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:132
			_go_fuzz_dep_.CoverTab[102429]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:133
			// _ = "end of CoverTab[102429]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:134
			_go_fuzz_dep_.CoverTab[102430]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:134
			// _ = "end of CoverTab[102430]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:134
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:134
		// _ = "end of CoverTab[102426]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:135
	// _ = "end of CoverTab[102409]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:135
	_go_fuzz_dep_.CoverTab[102410]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:137
	// _ = "end of CoverTab[102410]"
}

type GroupMemberDescription struct {
	ClientId		string
	ClientHost		string
	MemberMetadata		[]byte
	MemberAssignment	[]byte
}

func (gmd *GroupMemberDescription) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:147
	_go_fuzz_dep_.CoverTab[102431]++
													if err := pe.putString(gmd.ClientId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:148
		_go_fuzz_dep_.CoverTab[102436]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:149
		// _ = "end of CoverTab[102436]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:150
		_go_fuzz_dep_.CoverTab[102437]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:150
		// _ = "end of CoverTab[102437]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:150
	// _ = "end of CoverTab[102431]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:150
	_go_fuzz_dep_.CoverTab[102432]++
													if err := pe.putString(gmd.ClientHost); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:151
		_go_fuzz_dep_.CoverTab[102438]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:152
		// _ = "end of CoverTab[102438]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:153
		_go_fuzz_dep_.CoverTab[102439]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:153
		// _ = "end of CoverTab[102439]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:153
	// _ = "end of CoverTab[102432]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:153
	_go_fuzz_dep_.CoverTab[102433]++
													if err := pe.putBytes(gmd.MemberMetadata); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:154
		_go_fuzz_dep_.CoverTab[102440]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:155
		// _ = "end of CoverTab[102440]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:156
		_go_fuzz_dep_.CoverTab[102441]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:156
		// _ = "end of CoverTab[102441]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:156
	// _ = "end of CoverTab[102433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:156
	_go_fuzz_dep_.CoverTab[102434]++
													if err := pe.putBytes(gmd.MemberAssignment); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:157
		_go_fuzz_dep_.CoverTab[102442]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:158
		// _ = "end of CoverTab[102442]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:159
		_go_fuzz_dep_.CoverTab[102443]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:159
		// _ = "end of CoverTab[102443]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:159
	// _ = "end of CoverTab[102434]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:159
	_go_fuzz_dep_.CoverTab[102435]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:161
	// _ = "end of CoverTab[102435]"
}

func (gmd *GroupMemberDescription) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:164
	_go_fuzz_dep_.CoverTab[102444]++
													if gmd.ClientId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:165
		_go_fuzz_dep_.CoverTab[102449]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:166
		// _ = "end of CoverTab[102449]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:167
		_go_fuzz_dep_.CoverTab[102450]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:167
		// _ = "end of CoverTab[102450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:167
	// _ = "end of CoverTab[102444]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:167
	_go_fuzz_dep_.CoverTab[102445]++
													if gmd.ClientHost, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:168
		_go_fuzz_dep_.CoverTab[102451]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:169
		// _ = "end of CoverTab[102451]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:170
		_go_fuzz_dep_.CoverTab[102452]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:170
		// _ = "end of CoverTab[102452]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:170
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:170
	// _ = "end of CoverTab[102445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:170
	_go_fuzz_dep_.CoverTab[102446]++
													if gmd.MemberMetadata, err = pd.getBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:171
		_go_fuzz_dep_.CoverTab[102453]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:172
		// _ = "end of CoverTab[102453]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:173
		_go_fuzz_dep_.CoverTab[102454]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:173
		// _ = "end of CoverTab[102454]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:173
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:173
	// _ = "end of CoverTab[102446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:173
	_go_fuzz_dep_.CoverTab[102447]++
													if gmd.MemberAssignment, err = pd.getBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:174
		_go_fuzz_dep_.CoverTab[102455]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:175
		// _ = "end of CoverTab[102455]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:176
		_go_fuzz_dep_.CoverTab[102456]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:176
		// _ = "end of CoverTab[102456]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:176
	// _ = "end of CoverTab[102447]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:176
	_go_fuzz_dep_.CoverTab[102448]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:178
	// _ = "end of CoverTab[102448]"
}

func (gmd *GroupMemberDescription) GetMemberAssignment() (*ConsumerGroupMemberAssignment, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:181
	_go_fuzz_dep_.CoverTab[102457]++
													if len(gmd.MemberAssignment) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:182
		_go_fuzz_dep_.CoverTab[102459]++
														return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:183
		// _ = "end of CoverTab[102459]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:184
		_go_fuzz_dep_.CoverTab[102460]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:184
		// _ = "end of CoverTab[102460]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:184
	// _ = "end of CoverTab[102457]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:184
	_go_fuzz_dep_.CoverTab[102458]++
													assignment := new(ConsumerGroupMemberAssignment)
													err := decode(gmd.MemberAssignment, assignment)
													return assignment, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:187
	// _ = "end of CoverTab[102458]"
}

func (gmd *GroupMemberDescription) GetMemberMetadata() (*ConsumerGroupMemberMetadata, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:190
	_go_fuzz_dep_.CoverTab[102461]++
													if len(gmd.MemberMetadata) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:191
		_go_fuzz_dep_.CoverTab[102463]++
														return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:192
		// _ = "end of CoverTab[102463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:193
		_go_fuzz_dep_.CoverTab[102464]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:193
		// _ = "end of CoverTab[102464]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:193
	// _ = "end of CoverTab[102461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:193
	_go_fuzz_dep_.CoverTab[102462]++
													metadata := new(ConsumerGroupMemberMetadata)
													err := decode(gmd.MemberMetadata, metadata)
													return metadata, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:196
	// _ = "end of CoverTab[102462]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:197
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_response.go:197
var _ = _go_fuzz_dep_.CoverTab
