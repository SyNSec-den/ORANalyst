//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:1
)

// DeleteAclsRequest is a delete acl request
type DeleteAclsRequest struct {
	Version	int
	Filters	[]*AclFilter
}

func (d *DeleteAclsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:9
	_go_fuzz_dep_.CoverTab[97197]++
												if err := pe.putArrayLength(len(d.Filters)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:10
		_go_fuzz_dep_.CoverTab[97200]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:11
		// _ = "end of CoverTab[97200]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:12
		_go_fuzz_dep_.CoverTab[97201]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:12
		// _ = "end of CoverTab[97201]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:12
	// _ = "end of CoverTab[97197]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:12
	_go_fuzz_dep_.CoverTab[97198]++

												for _, filter := range d.Filters {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:14
		_go_fuzz_dep_.CoverTab[97202]++
													filter.Version = d.Version
													if err := filter.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:16
			_go_fuzz_dep_.CoverTab[97203]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:17
			// _ = "end of CoverTab[97203]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:18
			_go_fuzz_dep_.CoverTab[97204]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:18
			// _ = "end of CoverTab[97204]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:18
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:18
		// _ = "end of CoverTab[97202]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:19
	// _ = "end of CoverTab[97198]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:19
	_go_fuzz_dep_.CoverTab[97199]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:21
	// _ = "end of CoverTab[97199]"
}

func (d *DeleteAclsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:24
	_go_fuzz_dep_.CoverTab[97205]++
												d.Version = int(version)
												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:27
		_go_fuzz_dep_.CoverTab[97208]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:28
		// _ = "end of CoverTab[97208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:29
		_go_fuzz_dep_.CoverTab[97209]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:29
		// _ = "end of CoverTab[97209]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:29
	// _ = "end of CoverTab[97205]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:29
	_go_fuzz_dep_.CoverTab[97206]++

												d.Filters = make([]*AclFilter, n)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:32
		_go_fuzz_dep_.CoverTab[97210]++
													d.Filters[i] = new(AclFilter)
													d.Filters[i].Version = int(version)
													if err := d.Filters[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:35
			_go_fuzz_dep_.CoverTab[97211]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:36
			// _ = "end of CoverTab[97211]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:37
			_go_fuzz_dep_.CoverTab[97212]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:37
			// _ = "end of CoverTab[97212]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:37
		// _ = "end of CoverTab[97210]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:38
	// _ = "end of CoverTab[97206]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:38
	_go_fuzz_dep_.CoverTab[97207]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:40
	// _ = "end of CoverTab[97207]"
}

func (d *DeleteAclsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:43
	_go_fuzz_dep_.CoverTab[97213]++
												return 31
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:44
	// _ = "end of CoverTab[97213]"
}

func (d *DeleteAclsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:47
	_go_fuzz_dep_.CoverTab[97214]++
												return int16(d.Version)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:48
	// _ = "end of CoverTab[97214]"
}

func (d *DeleteAclsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:51
	_go_fuzz_dep_.CoverTab[97215]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:52
	// _ = "end of CoverTab[97215]"
}

func (d *DeleteAclsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:55
	_go_fuzz_dep_.CoverTab[97216]++
												switch d.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:57
		_go_fuzz_dep_.CoverTab[97217]++
													return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:58
		// _ = "end of CoverTab[97217]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:59
		_go_fuzz_dep_.CoverTab[97218]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:60
		// _ = "end of CoverTab[97218]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:61
	// _ = "end of CoverTab[97216]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_request.go:62
var _ = _go_fuzz_dep_.CoverTab
