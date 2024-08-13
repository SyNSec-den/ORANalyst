//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:1
)

type AclFilter struct {
	Version				int
	ResourceType			AclResourceType
	ResourceName			*string
	ResourcePatternTypeFilter	AclResourcePatternType
	Principal			*string
	Host				*string
	Operation			AclOperation
	PermissionType			AclPermissionType
}

func (a *AclFilter) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:14
	_go_fuzz_dep_.CoverTab[97335]++
											pe.putInt8(int8(a.ResourceType))
											if err := pe.putNullableString(a.ResourceName); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:16
		_go_fuzz_dep_.CoverTab[97340]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:17
		// _ = "end of CoverTab[97340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:18
		_go_fuzz_dep_.CoverTab[97341]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:18
		// _ = "end of CoverTab[97341]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:18
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:18
	// _ = "end of CoverTab[97335]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:18
	_go_fuzz_dep_.CoverTab[97336]++

											if a.Version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:20
		_go_fuzz_dep_.CoverTab[97342]++
												pe.putInt8(int8(a.ResourcePatternTypeFilter))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:21
		// _ = "end of CoverTab[97342]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:22
		_go_fuzz_dep_.CoverTab[97343]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:22
		// _ = "end of CoverTab[97343]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:22
	// _ = "end of CoverTab[97336]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:22
	_go_fuzz_dep_.CoverTab[97337]++

											if err := pe.putNullableString(a.Principal); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:24
		_go_fuzz_dep_.CoverTab[97344]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:25
		// _ = "end of CoverTab[97344]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:26
		_go_fuzz_dep_.CoverTab[97345]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:26
		// _ = "end of CoverTab[97345]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:26
	// _ = "end of CoverTab[97337]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:26
	_go_fuzz_dep_.CoverTab[97338]++
											if err := pe.putNullableString(a.Host); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:27
		_go_fuzz_dep_.CoverTab[97346]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:28
		// _ = "end of CoverTab[97346]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:29
		_go_fuzz_dep_.CoverTab[97347]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:29
		// _ = "end of CoverTab[97347]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:29
	// _ = "end of CoverTab[97338]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:29
	_go_fuzz_dep_.CoverTab[97339]++
											pe.putInt8(int8(a.Operation))
											pe.putInt8(int8(a.PermissionType))

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:33
	// _ = "end of CoverTab[97339]"
}

func (a *AclFilter) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:36
	_go_fuzz_dep_.CoverTab[97348]++
											resourceType, err := pd.getInt8()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:38
		_go_fuzz_dep_.CoverTab[97356]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:39
		// _ = "end of CoverTab[97356]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:40
		_go_fuzz_dep_.CoverTab[97357]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:40
		// _ = "end of CoverTab[97357]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:40
	// _ = "end of CoverTab[97348]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:40
	_go_fuzz_dep_.CoverTab[97349]++
											a.ResourceType = AclResourceType(resourceType)

											if a.ResourceName, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:43
		_go_fuzz_dep_.CoverTab[97358]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:44
		// _ = "end of CoverTab[97358]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:45
		_go_fuzz_dep_.CoverTab[97359]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:45
		// _ = "end of CoverTab[97359]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:45
	// _ = "end of CoverTab[97349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:45
	_go_fuzz_dep_.CoverTab[97350]++

											if a.Version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:47
		_go_fuzz_dep_.CoverTab[97360]++
												pattern, err := pd.getInt8()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:49
			_go_fuzz_dep_.CoverTab[97362]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:50
			// _ = "end of CoverTab[97362]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:51
			_go_fuzz_dep_.CoverTab[97363]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:51
			// _ = "end of CoverTab[97363]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:51
		// _ = "end of CoverTab[97360]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:51
		_go_fuzz_dep_.CoverTab[97361]++

												a.ResourcePatternTypeFilter = AclResourcePatternType(pattern)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:53
		// _ = "end of CoverTab[97361]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:54
		_go_fuzz_dep_.CoverTab[97364]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:54
		// _ = "end of CoverTab[97364]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:54
	// _ = "end of CoverTab[97350]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:54
	_go_fuzz_dep_.CoverTab[97351]++

											if a.Principal, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:56
		_go_fuzz_dep_.CoverTab[97365]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:57
		// _ = "end of CoverTab[97365]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:58
		_go_fuzz_dep_.CoverTab[97366]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:58
		// _ = "end of CoverTab[97366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:58
	// _ = "end of CoverTab[97351]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:58
	_go_fuzz_dep_.CoverTab[97352]++

											if a.Host, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:60
		_go_fuzz_dep_.CoverTab[97367]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:61
		// _ = "end of CoverTab[97367]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:62
		_go_fuzz_dep_.CoverTab[97368]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:62
		// _ = "end of CoverTab[97368]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:62
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:62
	// _ = "end of CoverTab[97352]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:62
	_go_fuzz_dep_.CoverTab[97353]++

											operation, err := pd.getInt8()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:65
		_go_fuzz_dep_.CoverTab[97369]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:66
		// _ = "end of CoverTab[97369]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:67
		_go_fuzz_dep_.CoverTab[97370]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:67
		// _ = "end of CoverTab[97370]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:67
	// _ = "end of CoverTab[97353]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:67
	_go_fuzz_dep_.CoverTab[97354]++
											a.Operation = AclOperation(operation)

											permissionType, err := pd.getInt8()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:71
		_go_fuzz_dep_.CoverTab[97371]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:72
		// _ = "end of CoverTab[97371]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:73
		_go_fuzz_dep_.CoverTab[97372]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:73
		// _ = "end of CoverTab[97372]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:73
	// _ = "end of CoverTab[97354]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:73
	_go_fuzz_dep_.CoverTab[97355]++
											a.PermissionType = AclPermissionType(permissionType)

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:76
	// _ = "end of CoverTab[97355]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:77
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_filter.go:77
var _ = _go_fuzz_dep_.CoverTab
