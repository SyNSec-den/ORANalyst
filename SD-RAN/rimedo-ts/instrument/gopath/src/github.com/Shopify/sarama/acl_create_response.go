//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:1
)

import "time"

// CreateAclsResponse is a an acl response creation type
type CreateAclsResponse struct {
	ThrottleTime		time.Duration
	AclCreationResponses	[]*AclCreationResponse
}

func (c *CreateAclsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:11
	_go_fuzz_dep_.CoverTab[97163]++
												pe.putInt32(int32(c.ThrottleTime / time.Millisecond))

												if err := pe.putArrayLength(len(c.AclCreationResponses)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:14
		_go_fuzz_dep_.CoverTab[97166]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:15
		// _ = "end of CoverTab[97166]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:16
		_go_fuzz_dep_.CoverTab[97167]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:16
		// _ = "end of CoverTab[97167]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:16
	// _ = "end of CoverTab[97163]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:16
	_go_fuzz_dep_.CoverTab[97164]++

												for _, aclCreationResponse := range c.AclCreationResponses {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:18
		_go_fuzz_dep_.CoverTab[97168]++
													if err := aclCreationResponse.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:19
			_go_fuzz_dep_.CoverTab[97169]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:20
			// _ = "end of CoverTab[97169]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:21
			_go_fuzz_dep_.CoverTab[97170]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:21
			// _ = "end of CoverTab[97170]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:21
		// _ = "end of CoverTab[97168]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:22
	// _ = "end of CoverTab[97164]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:22
	_go_fuzz_dep_.CoverTab[97165]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:24
	// _ = "end of CoverTab[97165]"
}

func (c *CreateAclsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:27
	_go_fuzz_dep_.CoverTab[97171]++
												throttleTime, err := pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:29
		_go_fuzz_dep_.CoverTab[97175]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:30
		// _ = "end of CoverTab[97175]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:31
		_go_fuzz_dep_.CoverTab[97176]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:31
		// _ = "end of CoverTab[97176]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:31
	// _ = "end of CoverTab[97171]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:31
	_go_fuzz_dep_.CoverTab[97172]++
												c.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:35
		_go_fuzz_dep_.CoverTab[97177]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:36
		// _ = "end of CoverTab[97177]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:37
		_go_fuzz_dep_.CoverTab[97178]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:37
		// _ = "end of CoverTab[97178]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:37
	// _ = "end of CoverTab[97172]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:37
	_go_fuzz_dep_.CoverTab[97173]++

												c.AclCreationResponses = make([]*AclCreationResponse, n)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:40
		_go_fuzz_dep_.CoverTab[97179]++
													c.AclCreationResponses[i] = new(AclCreationResponse)
													if err := c.AclCreationResponses[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:42
			_go_fuzz_dep_.CoverTab[97180]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:43
			// _ = "end of CoverTab[97180]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:44
			_go_fuzz_dep_.CoverTab[97181]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:44
			// _ = "end of CoverTab[97181]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:44
		// _ = "end of CoverTab[97179]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:45
	// _ = "end of CoverTab[97173]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:45
	_go_fuzz_dep_.CoverTab[97174]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:47
	// _ = "end of CoverTab[97174]"
}

func (c *CreateAclsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:50
	_go_fuzz_dep_.CoverTab[97182]++
												return 30
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:51
	// _ = "end of CoverTab[97182]"
}

func (c *CreateAclsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:54
	_go_fuzz_dep_.CoverTab[97183]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:55
	// _ = "end of CoverTab[97183]"
}

func (c *CreateAclsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:58
	_go_fuzz_dep_.CoverTab[97184]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:59
	// _ = "end of CoverTab[97184]"
}

func (c *CreateAclsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:62
	_go_fuzz_dep_.CoverTab[97185]++
												return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:63
	// _ = "end of CoverTab[97185]"
}

// AclCreationResponse is an acl creation response type
type AclCreationResponse struct {
	Err	KError
	ErrMsg	*string
}

func (a *AclCreationResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:72
	_go_fuzz_dep_.CoverTab[97186]++
												pe.putInt16(int16(a.Err))

												if err := pe.putNullableString(a.ErrMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:75
		_go_fuzz_dep_.CoverTab[97188]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:76
		// _ = "end of CoverTab[97188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:77
		_go_fuzz_dep_.CoverTab[97189]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:77
		// _ = "end of CoverTab[97189]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:77
	// _ = "end of CoverTab[97186]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:77
	_go_fuzz_dep_.CoverTab[97187]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:79
	// _ = "end of CoverTab[97187]"
}

func (a *AclCreationResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:82
	_go_fuzz_dep_.CoverTab[97190]++
												kerr, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:84
		_go_fuzz_dep_.CoverTab[97193]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:85
		// _ = "end of CoverTab[97193]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:86
		_go_fuzz_dep_.CoverTab[97194]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:86
		// _ = "end of CoverTab[97194]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:86
	// _ = "end of CoverTab[97190]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:86
	_go_fuzz_dep_.CoverTab[97191]++
												a.Err = KError(kerr)

												if a.ErrMsg, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:89
		_go_fuzz_dep_.CoverTab[97195]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:90
		// _ = "end of CoverTab[97195]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:91
		_go_fuzz_dep_.CoverTab[97196]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:91
		// _ = "end of CoverTab[97196]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:91
	// _ = "end of CoverTab[97191]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:91
	_go_fuzz_dep_.CoverTab[97192]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:93
	// _ = "end of CoverTab[97192]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_response.go:94
var _ = _go_fuzz_dep_.CoverTab
