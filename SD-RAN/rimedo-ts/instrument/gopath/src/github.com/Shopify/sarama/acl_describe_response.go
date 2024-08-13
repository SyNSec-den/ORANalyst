//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:1
)

import "time"

// DescribeAclsResponse is a describe acl response type
type DescribeAclsResponse struct {
	Version		int16
	ThrottleTime	time.Duration
	Err		KError
	ErrMsg		*string
	ResourceAcls	[]*ResourceAcls
}

func (d *DescribeAclsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:14
	_go_fuzz_dep_.CoverTab[97298]++
													pe.putInt32(int32(d.ThrottleTime / time.Millisecond))
													pe.putInt16(int16(d.Err))

													if err := pe.putNullableString(d.ErrMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:18
		_go_fuzz_dep_.CoverTab[97302]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:19
		// _ = "end of CoverTab[97302]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:20
		_go_fuzz_dep_.CoverTab[97303]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:20
		// _ = "end of CoverTab[97303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:20
	// _ = "end of CoverTab[97298]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:20
	_go_fuzz_dep_.CoverTab[97299]++

													if err := pe.putArrayLength(len(d.ResourceAcls)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:22
		_go_fuzz_dep_.CoverTab[97304]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:23
		// _ = "end of CoverTab[97304]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:24
		_go_fuzz_dep_.CoverTab[97305]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:24
		// _ = "end of CoverTab[97305]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:24
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:24
	// _ = "end of CoverTab[97299]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:24
	_go_fuzz_dep_.CoverTab[97300]++

													for _, resourceAcl := range d.ResourceAcls {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:26
		_go_fuzz_dep_.CoverTab[97306]++
														if err := resourceAcl.encode(pe, d.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:27
			_go_fuzz_dep_.CoverTab[97307]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:28
			// _ = "end of CoverTab[97307]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:29
			_go_fuzz_dep_.CoverTab[97308]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:29
			// _ = "end of CoverTab[97308]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:29
		// _ = "end of CoverTab[97306]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:30
	// _ = "end of CoverTab[97300]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:30
	_go_fuzz_dep_.CoverTab[97301]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:32
	// _ = "end of CoverTab[97301]"
}

func (d *DescribeAclsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:35
	_go_fuzz_dep_.CoverTab[97309]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:37
		_go_fuzz_dep_.CoverTab[97316]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:38
		// _ = "end of CoverTab[97316]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:39
		_go_fuzz_dep_.CoverTab[97317]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:39
		// _ = "end of CoverTab[97317]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:39
	// _ = "end of CoverTab[97309]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:39
	_go_fuzz_dep_.CoverTab[97310]++
													d.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:43
		_go_fuzz_dep_.CoverTab[97318]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:44
		// _ = "end of CoverTab[97318]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:45
		_go_fuzz_dep_.CoverTab[97319]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:45
		// _ = "end of CoverTab[97319]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:45
	// _ = "end of CoverTab[97310]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:45
	_go_fuzz_dep_.CoverTab[97311]++
													d.Err = KError(kerr)

													errmsg, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:49
		_go_fuzz_dep_.CoverTab[97320]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:50
		// _ = "end of CoverTab[97320]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:51
		_go_fuzz_dep_.CoverTab[97321]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:51
		// _ = "end of CoverTab[97321]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:51
	// _ = "end of CoverTab[97311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:51
	_go_fuzz_dep_.CoverTab[97312]++
													if errmsg != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:52
		_go_fuzz_dep_.CoverTab[97322]++
														d.ErrMsg = &errmsg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:53
		// _ = "end of CoverTab[97322]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:54
		_go_fuzz_dep_.CoverTab[97323]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:54
		// _ = "end of CoverTab[97323]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:54
	// _ = "end of CoverTab[97312]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:54
	_go_fuzz_dep_.CoverTab[97313]++

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:57
		_go_fuzz_dep_.CoverTab[97324]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:58
		// _ = "end of CoverTab[97324]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:59
		_go_fuzz_dep_.CoverTab[97325]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:59
		// _ = "end of CoverTab[97325]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:59
	// _ = "end of CoverTab[97313]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:59
	_go_fuzz_dep_.CoverTab[97314]++
													d.ResourceAcls = make([]*ResourceAcls, n)

													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:62
		_go_fuzz_dep_.CoverTab[97326]++
														d.ResourceAcls[i] = new(ResourceAcls)
														if err := d.ResourceAcls[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:64
			_go_fuzz_dep_.CoverTab[97327]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:65
			// _ = "end of CoverTab[97327]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:66
			_go_fuzz_dep_.CoverTab[97328]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:66
			// _ = "end of CoverTab[97328]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:66
		// _ = "end of CoverTab[97326]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:67
	// _ = "end of CoverTab[97314]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:67
	_go_fuzz_dep_.CoverTab[97315]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:69
	// _ = "end of CoverTab[97315]"
}

func (d *DescribeAclsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:72
	_go_fuzz_dep_.CoverTab[97329]++
													return 29
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:73
	// _ = "end of CoverTab[97329]"
}

func (d *DescribeAclsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:76
	_go_fuzz_dep_.CoverTab[97330]++
													return d.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:77
	// _ = "end of CoverTab[97330]"
}

func (d *DescribeAclsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:80
	_go_fuzz_dep_.CoverTab[97331]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:81
	// _ = "end of CoverTab[97331]"
}

func (d *DescribeAclsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:84
	_go_fuzz_dep_.CoverTab[97332]++
													switch d.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:86
		_go_fuzz_dep_.CoverTab[97333]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:87
		// _ = "end of CoverTab[97333]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:88
		_go_fuzz_dep_.CoverTab[97334]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:89
		// _ = "end of CoverTab[97334]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:90
	// _ = "end of CoverTab[97332]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:91
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_response.go:91
var _ = _go_fuzz_dep_.CoverTab
