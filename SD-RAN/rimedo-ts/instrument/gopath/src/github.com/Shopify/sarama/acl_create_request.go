//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:1
)

// CreateAclsRequest is an acl creation request
type CreateAclsRequest struct {
	Version		int16
	AclCreations	[]*AclCreation
}

func (c *CreateAclsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:9
	_go_fuzz_dep_.CoverTab[97127]++
												if err := pe.putArrayLength(len(c.AclCreations)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:10
		_go_fuzz_dep_.CoverTab[97130]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:11
		// _ = "end of CoverTab[97130]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:12
		_go_fuzz_dep_.CoverTab[97131]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:12
		// _ = "end of CoverTab[97131]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:12
	// _ = "end of CoverTab[97127]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:12
	_go_fuzz_dep_.CoverTab[97128]++

												for _, aclCreation := range c.AclCreations {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:14
		_go_fuzz_dep_.CoverTab[97132]++
													if err := aclCreation.encode(pe, c.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:15
			_go_fuzz_dep_.CoverTab[97133]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:16
			// _ = "end of CoverTab[97133]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:17
			_go_fuzz_dep_.CoverTab[97134]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:17
			// _ = "end of CoverTab[97134]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:17
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:17
		// _ = "end of CoverTab[97132]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:18
	// _ = "end of CoverTab[97128]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:18
	_go_fuzz_dep_.CoverTab[97129]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:20
	// _ = "end of CoverTab[97129]"
}

func (c *CreateAclsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:23
	_go_fuzz_dep_.CoverTab[97135]++
												c.Version = version
												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:26
		_go_fuzz_dep_.CoverTab[97138]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:27
		// _ = "end of CoverTab[97138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:28
		_go_fuzz_dep_.CoverTab[97139]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:28
		// _ = "end of CoverTab[97139]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:28
	// _ = "end of CoverTab[97135]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:28
	_go_fuzz_dep_.CoverTab[97136]++

												c.AclCreations = make([]*AclCreation, n)

												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:32
		_go_fuzz_dep_.CoverTab[97140]++
													c.AclCreations[i] = new(AclCreation)
													if err := c.AclCreations[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:34
			_go_fuzz_dep_.CoverTab[97141]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:35
			// _ = "end of CoverTab[97141]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:36
			_go_fuzz_dep_.CoverTab[97142]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:36
			// _ = "end of CoverTab[97142]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:36
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:36
		// _ = "end of CoverTab[97140]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:37
	// _ = "end of CoverTab[97136]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:37
	_go_fuzz_dep_.CoverTab[97137]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:39
	// _ = "end of CoverTab[97137]"
}

func (c *CreateAclsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:42
	_go_fuzz_dep_.CoverTab[97143]++
												return 30
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:43
	// _ = "end of CoverTab[97143]"
}

func (c *CreateAclsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:46
	_go_fuzz_dep_.CoverTab[97144]++
												return c.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:47
	// _ = "end of CoverTab[97144]"
}

func (c *CreateAclsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:50
	_go_fuzz_dep_.CoverTab[97145]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:51
	// _ = "end of CoverTab[97145]"
}

func (c *CreateAclsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:54
	_go_fuzz_dep_.CoverTab[97146]++
												switch c.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:56
		_go_fuzz_dep_.CoverTab[97147]++
													return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:57
		// _ = "end of CoverTab[97147]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:58
		_go_fuzz_dep_.CoverTab[97148]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:59
		// _ = "end of CoverTab[97148]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:60
	// _ = "end of CoverTab[97146]"
}

// AclCreation is a wrapper around Resource and Acl type
type AclCreation struct {
	Resource
	Acl
}

func (a *AclCreation) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:69
	_go_fuzz_dep_.CoverTab[97149]++
												if err := a.Resource.encode(pe, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:70
		_go_fuzz_dep_.CoverTab[97152]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:71
		// _ = "end of CoverTab[97152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:72
		_go_fuzz_dep_.CoverTab[97153]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:72
		// _ = "end of CoverTab[97153]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:72
	// _ = "end of CoverTab[97149]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:72
	_go_fuzz_dep_.CoverTab[97150]++
												if err := a.Acl.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:73
		_go_fuzz_dep_.CoverTab[97154]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:74
		// _ = "end of CoverTab[97154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:75
		_go_fuzz_dep_.CoverTab[97155]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:75
		// _ = "end of CoverTab[97155]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:75
	// _ = "end of CoverTab[97150]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:75
	_go_fuzz_dep_.CoverTab[97151]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:77
	// _ = "end of CoverTab[97151]"
}

func (a *AclCreation) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:80
	_go_fuzz_dep_.CoverTab[97156]++
												if err := a.Resource.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:81
		_go_fuzz_dep_.CoverTab[97159]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:82
		// _ = "end of CoverTab[97159]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:83
		_go_fuzz_dep_.CoverTab[97160]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:83
		// _ = "end of CoverTab[97160]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:83
	// _ = "end of CoverTab[97156]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:83
	_go_fuzz_dep_.CoverTab[97157]++
												if err := a.Acl.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:84
		_go_fuzz_dep_.CoverTab[97161]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:85
		// _ = "end of CoverTab[97161]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:86
		_go_fuzz_dep_.CoverTab[97162]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:86
		// _ = "end of CoverTab[97162]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:86
	// _ = "end of CoverTab[97157]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:86
	_go_fuzz_dep_.CoverTab[97158]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:88
	// _ = "end of CoverTab[97158]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_create_request.go:89
var _ = _go_fuzz_dep_.CoverTab
