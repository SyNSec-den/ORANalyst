//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:1
)

// Resource holds information about acl resource type
type Resource struct {
	ResourceType		AclResourceType
	ResourceName		string
	ResourcePatternType	AclResourcePatternType
}

func (r *Resource) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:10
	_go_fuzz_dep_.CoverTab[97062]++
												pe.putInt8(int8(r.ResourceType))

												if err := pe.putString(r.ResourceName); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:13
		_go_fuzz_dep_.CoverTab[97065]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:14
		// _ = "end of CoverTab[97065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:15
		_go_fuzz_dep_.CoverTab[97066]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:15
		// _ = "end of CoverTab[97066]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:15
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:15
	// _ = "end of CoverTab[97062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:15
	_go_fuzz_dep_.CoverTab[97063]++

												if version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:17
		_go_fuzz_dep_.CoverTab[97067]++
													if r.ResourcePatternType == AclPatternUnknown {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:18
			_go_fuzz_dep_.CoverTab[97069]++
														Logger.Print("Cannot encode an unknown resource pattern type, using Literal instead")
														r.ResourcePatternType = AclPatternLiteral
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:20
			// _ = "end of CoverTab[97069]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:21
			_go_fuzz_dep_.CoverTab[97070]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:21
			// _ = "end of CoverTab[97070]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:21
		// _ = "end of CoverTab[97067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:21
		_go_fuzz_dep_.CoverTab[97068]++
													pe.putInt8(int8(r.ResourcePatternType))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:22
		// _ = "end of CoverTab[97068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:23
		_go_fuzz_dep_.CoverTab[97071]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:23
		// _ = "end of CoverTab[97071]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:23
	// _ = "end of CoverTab[97063]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:23
	_go_fuzz_dep_.CoverTab[97064]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:25
	// _ = "end of CoverTab[97064]"
}

func (r *Resource) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:28
	_go_fuzz_dep_.CoverTab[97072]++
												resourceType, err := pd.getInt8()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:30
		_go_fuzz_dep_.CoverTab[97076]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:31
		// _ = "end of CoverTab[97076]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:32
		_go_fuzz_dep_.CoverTab[97077]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:32
		// _ = "end of CoverTab[97077]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:32
	// _ = "end of CoverTab[97072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:32
	_go_fuzz_dep_.CoverTab[97073]++
												r.ResourceType = AclResourceType(resourceType)

												if r.ResourceName, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:35
		_go_fuzz_dep_.CoverTab[97078]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:36
		// _ = "end of CoverTab[97078]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:37
		_go_fuzz_dep_.CoverTab[97079]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:37
		// _ = "end of CoverTab[97079]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:37
	// _ = "end of CoverTab[97073]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:37
	_go_fuzz_dep_.CoverTab[97074]++
												if version == 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:38
		_go_fuzz_dep_.CoverTab[97080]++
													pattern, err := pd.getInt8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:40
			_go_fuzz_dep_.CoverTab[97082]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:41
			// _ = "end of CoverTab[97082]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:42
			_go_fuzz_dep_.CoverTab[97083]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:42
			// _ = "end of CoverTab[97083]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:42
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:42
		// _ = "end of CoverTab[97080]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:42
		_go_fuzz_dep_.CoverTab[97081]++
													r.ResourcePatternType = AclResourcePatternType(pattern)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:43
		// _ = "end of CoverTab[97081]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:44
		_go_fuzz_dep_.CoverTab[97084]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:44
		// _ = "end of CoverTab[97084]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:44
	// _ = "end of CoverTab[97074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:44
	_go_fuzz_dep_.CoverTab[97075]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:46
	// _ = "end of CoverTab[97075]"
}

// Acl holds information about acl type
type Acl struct {
	Principal	string
	Host		string
	Operation	AclOperation
	PermissionType	AclPermissionType
}

func (a *Acl) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:57
	_go_fuzz_dep_.CoverTab[97085]++
												if err := pe.putString(a.Principal); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:58
		_go_fuzz_dep_.CoverTab[97088]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:59
		// _ = "end of CoverTab[97088]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:60
		_go_fuzz_dep_.CoverTab[97089]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:60
		// _ = "end of CoverTab[97089]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:60
	// _ = "end of CoverTab[97085]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:60
	_go_fuzz_dep_.CoverTab[97086]++

												if err := pe.putString(a.Host); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:62
		_go_fuzz_dep_.CoverTab[97090]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:63
		// _ = "end of CoverTab[97090]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:64
		_go_fuzz_dep_.CoverTab[97091]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:64
		// _ = "end of CoverTab[97091]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:64
	// _ = "end of CoverTab[97086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:64
	_go_fuzz_dep_.CoverTab[97087]++

												pe.putInt8(int8(a.Operation))
												pe.putInt8(int8(a.PermissionType))

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:69
	// _ = "end of CoverTab[97087]"
}

func (a *Acl) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:72
	_go_fuzz_dep_.CoverTab[97092]++
												if a.Principal, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:73
		_go_fuzz_dep_.CoverTab[97097]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:74
		// _ = "end of CoverTab[97097]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:75
		_go_fuzz_dep_.CoverTab[97098]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:75
		// _ = "end of CoverTab[97098]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:75
	// _ = "end of CoverTab[97092]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:75
	_go_fuzz_dep_.CoverTab[97093]++

												if a.Host, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:77
		_go_fuzz_dep_.CoverTab[97099]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:78
		// _ = "end of CoverTab[97099]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:79
		_go_fuzz_dep_.CoverTab[97100]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:79
		// _ = "end of CoverTab[97100]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:79
	// _ = "end of CoverTab[97093]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:79
	_go_fuzz_dep_.CoverTab[97094]++

												operation, err := pd.getInt8()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:82
		_go_fuzz_dep_.CoverTab[97101]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:83
		// _ = "end of CoverTab[97101]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:84
		_go_fuzz_dep_.CoverTab[97102]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:84
		// _ = "end of CoverTab[97102]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:84
	// _ = "end of CoverTab[97094]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:84
	_go_fuzz_dep_.CoverTab[97095]++
												a.Operation = AclOperation(operation)

												permissionType, err := pd.getInt8()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:88
		_go_fuzz_dep_.CoverTab[97103]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:89
		// _ = "end of CoverTab[97103]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:90
		_go_fuzz_dep_.CoverTab[97104]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:90
		// _ = "end of CoverTab[97104]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:90
	// _ = "end of CoverTab[97095]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:90
	_go_fuzz_dep_.CoverTab[97096]++
												a.PermissionType = AclPermissionType(permissionType)

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:93
	// _ = "end of CoverTab[97096]"
}

// ResourceAcls is an acl resource type
type ResourceAcls struct {
	Resource
	Acls	[]*Acl
}

func (r *ResourceAcls) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:102
	_go_fuzz_dep_.CoverTab[97105]++
												if err := r.Resource.encode(pe, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:103
		_go_fuzz_dep_.CoverTab[97109]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:104
		// _ = "end of CoverTab[97109]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:105
		_go_fuzz_dep_.CoverTab[97110]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:105
		// _ = "end of CoverTab[97110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:105
	// _ = "end of CoverTab[97105]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:105
	_go_fuzz_dep_.CoverTab[97106]++

												if err := pe.putArrayLength(len(r.Acls)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:107
		_go_fuzz_dep_.CoverTab[97111]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:108
		// _ = "end of CoverTab[97111]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:109
		_go_fuzz_dep_.CoverTab[97112]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:109
		// _ = "end of CoverTab[97112]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:109
	// _ = "end of CoverTab[97106]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:109
	_go_fuzz_dep_.CoverTab[97107]++
												for _, acl := range r.Acls {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:110
		_go_fuzz_dep_.CoverTab[97113]++
													if err := acl.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:111
			_go_fuzz_dep_.CoverTab[97114]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:112
			// _ = "end of CoverTab[97114]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:113
			_go_fuzz_dep_.CoverTab[97115]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:113
			// _ = "end of CoverTab[97115]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:113
		// _ = "end of CoverTab[97113]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:114
	// _ = "end of CoverTab[97107]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:114
	_go_fuzz_dep_.CoverTab[97108]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:116
	// _ = "end of CoverTab[97108]"
}

func (r *ResourceAcls) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:119
	_go_fuzz_dep_.CoverTab[97116]++
												if err := r.Resource.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:120
		_go_fuzz_dep_.CoverTab[97120]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:121
		// _ = "end of CoverTab[97120]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:122
		_go_fuzz_dep_.CoverTab[97121]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:122
		// _ = "end of CoverTab[97121]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:122
	// _ = "end of CoverTab[97116]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:122
	_go_fuzz_dep_.CoverTab[97117]++

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:125
		_go_fuzz_dep_.CoverTab[97122]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:126
		// _ = "end of CoverTab[97122]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:127
		_go_fuzz_dep_.CoverTab[97123]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:127
		// _ = "end of CoverTab[97123]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:127
	// _ = "end of CoverTab[97117]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:127
	_go_fuzz_dep_.CoverTab[97118]++

												r.Acls = make([]*Acl, n)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:130
		_go_fuzz_dep_.CoverTab[97124]++
													r.Acls[i] = new(Acl)
													if err := r.Acls[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:132
			_go_fuzz_dep_.CoverTab[97125]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:133
			// _ = "end of CoverTab[97125]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:134
			_go_fuzz_dep_.CoverTab[97126]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:134
			// _ = "end of CoverTab[97126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:134
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:134
		// _ = "end of CoverTab[97124]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:135
	// _ = "end of CoverTab[97118]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:135
	_go_fuzz_dep_.CoverTab[97119]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:137
	// _ = "end of CoverTab[97119]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:138
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_bindings.go:138
var _ = _go_fuzz_dep_.CoverTab
