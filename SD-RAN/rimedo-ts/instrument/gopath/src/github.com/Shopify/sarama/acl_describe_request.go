//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:1
)

// DescribeAclsRequest is a secribe acl request type
type DescribeAclsRequest struct {
	Version	int
	AclFilter
}

func (d *DescribeAclsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:9
		_go_fuzz_dep_.CoverTab[97290]++
													d.AclFilter.Version = d.Version
													return d.AclFilter.encode(pe)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:11
	// _ = "end of CoverTab[97290]"
}

func (d *DescribeAclsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:14
	_go_fuzz_dep_.CoverTab[97291]++
													d.Version = int(version)
													d.AclFilter.Version = int(version)
													return d.AclFilter.decode(pd, version)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:17
	// _ = "end of CoverTab[97291]"
}

func (d *DescribeAclsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:20
	_go_fuzz_dep_.CoverTab[97292]++
													return 29
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:21
	// _ = "end of CoverTab[97292]"
}

func (d *DescribeAclsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:24
	_go_fuzz_dep_.CoverTab[97293]++
													return int16(d.Version)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:25
	// _ = "end of CoverTab[97293]"
}

func (d *DescribeAclsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:28
	_go_fuzz_dep_.CoverTab[97294]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:29
	// _ = "end of CoverTab[97294]"
}

func (d *DescribeAclsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:32
	_go_fuzz_dep_.CoverTab[97295]++
													switch d.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:34
		_go_fuzz_dep_.CoverTab[97296]++
														return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:35
		// _ = "end of CoverTab[97296]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:36
		_go_fuzz_dep_.CoverTab[97297]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:37
		// _ = "end of CoverTab[97297]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:38
	// _ = "end of CoverTab[97295]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:39
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_describe_request.go:39
var _ = _go_fuzz_dep_.CoverTab
