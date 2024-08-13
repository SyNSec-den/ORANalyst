//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:1
)

import "time"

// DeleteAclsResponse is a delete acl response
type DeleteAclsResponse struct {
	Version		int16
	ThrottleTime	time.Duration
	FilterResponses	[]*FilterResponse
}

func (d *DeleteAclsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:12
	_go_fuzz_dep_.CoverTab[97219]++
												pe.putInt32(int32(d.ThrottleTime / time.Millisecond))

												if err := pe.putArrayLength(len(d.FilterResponses)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:15
		_go_fuzz_dep_.CoverTab[97222]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:16
		// _ = "end of CoverTab[97222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:17
		_go_fuzz_dep_.CoverTab[97223]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:17
		// _ = "end of CoverTab[97223]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:17
	// _ = "end of CoverTab[97219]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:17
	_go_fuzz_dep_.CoverTab[97220]++

												for _, filterResponse := range d.FilterResponses {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:19
		_go_fuzz_dep_.CoverTab[97224]++
													if err := filterResponse.encode(pe, d.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:20
			_go_fuzz_dep_.CoverTab[97225]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:21
			// _ = "end of CoverTab[97225]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:22
			_go_fuzz_dep_.CoverTab[97226]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:22
			// _ = "end of CoverTab[97226]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:22
		// _ = "end of CoverTab[97224]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:23
	// _ = "end of CoverTab[97220]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:23
	_go_fuzz_dep_.CoverTab[97221]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:25
	// _ = "end of CoverTab[97221]"
}

func (d *DeleteAclsResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:28
	_go_fuzz_dep_.CoverTab[97227]++
												throttleTime, err := pd.getInt32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:30
		_go_fuzz_dep_.CoverTab[97231]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:31
		// _ = "end of CoverTab[97231]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:32
		_go_fuzz_dep_.CoverTab[97232]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:32
		// _ = "end of CoverTab[97232]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:32
	// _ = "end of CoverTab[97227]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:32
	_go_fuzz_dep_.CoverTab[97228]++
												d.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:36
		_go_fuzz_dep_.CoverTab[97233]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:37
		// _ = "end of CoverTab[97233]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:38
		_go_fuzz_dep_.CoverTab[97234]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:38
		// _ = "end of CoverTab[97234]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:38
	// _ = "end of CoverTab[97228]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:38
	_go_fuzz_dep_.CoverTab[97229]++
												d.FilterResponses = make([]*FilterResponse, n)

												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:41
		_go_fuzz_dep_.CoverTab[97235]++
													d.FilterResponses[i] = new(FilterResponse)
													if err := d.FilterResponses[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:43
			_go_fuzz_dep_.CoverTab[97236]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:44
			// _ = "end of CoverTab[97236]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:45
			_go_fuzz_dep_.CoverTab[97237]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:45
			// _ = "end of CoverTab[97237]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:45
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:45
		// _ = "end of CoverTab[97235]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:46
	// _ = "end of CoverTab[97229]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:46
	_go_fuzz_dep_.CoverTab[97230]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:48
	// _ = "end of CoverTab[97230]"
}

func (d *DeleteAclsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:51
	_go_fuzz_dep_.CoverTab[97238]++
												return 31
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:52
	// _ = "end of CoverTab[97238]"
}

func (d *DeleteAclsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:55
	_go_fuzz_dep_.CoverTab[97239]++
												return d.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:56
	// _ = "end of CoverTab[97239]"
}

func (d *DeleteAclsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:59
	_go_fuzz_dep_.CoverTab[97240]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:60
	// _ = "end of CoverTab[97240]"
}

func (d *DeleteAclsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:63
	_go_fuzz_dep_.CoverTab[97241]++
												return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:64
	// _ = "end of CoverTab[97241]"
}

// FilterResponse is a filter response type
type FilterResponse struct {
	Err		KError
	ErrMsg		*string
	MatchingAcls	[]*MatchingAcl
}

func (f *FilterResponse) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:74
	_go_fuzz_dep_.CoverTab[97242]++
												pe.putInt16(int16(f.Err))
												if err := pe.putNullableString(f.ErrMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:76
		_go_fuzz_dep_.CoverTab[97246]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:77
		// _ = "end of CoverTab[97246]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:78
		_go_fuzz_dep_.CoverTab[97247]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:78
		// _ = "end of CoverTab[97247]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:78
	// _ = "end of CoverTab[97242]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:78
	_go_fuzz_dep_.CoverTab[97243]++

												if err := pe.putArrayLength(len(f.MatchingAcls)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:80
		_go_fuzz_dep_.CoverTab[97248]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:81
		// _ = "end of CoverTab[97248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:82
		_go_fuzz_dep_.CoverTab[97249]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:82
		// _ = "end of CoverTab[97249]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:82
	// _ = "end of CoverTab[97243]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:82
	_go_fuzz_dep_.CoverTab[97244]++
												for _, matchingAcl := range f.MatchingAcls {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:83
		_go_fuzz_dep_.CoverTab[97250]++
													if err := matchingAcl.encode(pe, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:84
			_go_fuzz_dep_.CoverTab[97251]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:85
			// _ = "end of CoverTab[97251]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:86
			_go_fuzz_dep_.CoverTab[97252]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:86
			// _ = "end of CoverTab[97252]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:86
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:86
		// _ = "end of CoverTab[97250]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:87
	// _ = "end of CoverTab[97244]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:87
	_go_fuzz_dep_.CoverTab[97245]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:89
	// _ = "end of CoverTab[97245]"
}

func (f *FilterResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:92
	_go_fuzz_dep_.CoverTab[97253]++
												kerr, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:94
		_go_fuzz_dep_.CoverTab[97258]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:95
		// _ = "end of CoverTab[97258]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:96
		_go_fuzz_dep_.CoverTab[97259]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:96
		// _ = "end of CoverTab[97259]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:96
	// _ = "end of CoverTab[97253]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:96
	_go_fuzz_dep_.CoverTab[97254]++
												f.Err = KError(kerr)

												if f.ErrMsg, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:99
			_go_fuzz_dep_.CoverTab[97260]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:100
		// _ = "end of CoverTab[97260]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:101
		_go_fuzz_dep_.CoverTab[97261]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:101
		// _ = "end of CoverTab[97261]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:101
	// _ = "end of CoverTab[97254]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:101
	_go_fuzz_dep_.CoverTab[97255]++

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:104
		_go_fuzz_dep_.CoverTab[97262]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:105
		// _ = "end of CoverTab[97262]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:106
		_go_fuzz_dep_.CoverTab[97263]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:106
		// _ = "end of CoverTab[97263]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:106
	// _ = "end of CoverTab[97255]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:106
	_go_fuzz_dep_.CoverTab[97256]++
													f.MatchingAcls = make([]*MatchingAcl, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:108
		_go_fuzz_dep_.CoverTab[97264]++
														f.MatchingAcls[i] = new(MatchingAcl)
														if err := f.MatchingAcls[i].decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:110
			_go_fuzz_dep_.CoverTab[97265]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:111
			// _ = "end of CoverTab[97265]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:112
			_go_fuzz_dep_.CoverTab[97266]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:112
			// _ = "end of CoverTab[97266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:112
		// _ = "end of CoverTab[97264]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:113
	// _ = "end of CoverTab[97256]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:113
	_go_fuzz_dep_.CoverTab[97257]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:115
	// _ = "end of CoverTab[97257]"
}

// MatchingAcl is a matching acl type
type MatchingAcl struct {
	Err	KError
	ErrMsg	*string
	Resource
	Acl
}

func (m *MatchingAcl) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:126
	_go_fuzz_dep_.CoverTab[97267]++
													pe.putInt16(int16(m.Err))
													if err := pe.putNullableString(m.ErrMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:128
		_go_fuzz_dep_.CoverTab[97271]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:129
		// _ = "end of CoverTab[97271]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:130
		_go_fuzz_dep_.CoverTab[97272]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:130
		// _ = "end of CoverTab[97272]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:130
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:130
	// _ = "end of CoverTab[97267]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:130
	_go_fuzz_dep_.CoverTab[97268]++

													if err := m.Resource.encode(pe, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:132
		_go_fuzz_dep_.CoverTab[97273]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:133
		// _ = "end of CoverTab[97273]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:134
		_go_fuzz_dep_.CoverTab[97274]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:134
		// _ = "end of CoverTab[97274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:134
	// _ = "end of CoverTab[97268]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:134
	_go_fuzz_dep_.CoverTab[97269]++

													if err := m.Acl.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:136
		_go_fuzz_dep_.CoverTab[97275]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:137
		// _ = "end of CoverTab[97275]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:138
		_go_fuzz_dep_.CoverTab[97276]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:138
		// _ = "end of CoverTab[97276]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:138
	// _ = "end of CoverTab[97269]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:138
	_go_fuzz_dep_.CoverTab[97270]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:140
	// _ = "end of CoverTab[97270]"
}

func (m *MatchingAcl) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:143
	_go_fuzz_dep_.CoverTab[97277]++
													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:145
		_go_fuzz_dep_.CoverTab[97282]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:146
		// _ = "end of CoverTab[97282]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:147
		_go_fuzz_dep_.CoverTab[97283]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:147
		// _ = "end of CoverTab[97283]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:147
	// _ = "end of CoverTab[97277]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:147
	_go_fuzz_dep_.CoverTab[97278]++
													m.Err = KError(kerr)

													if m.ErrMsg, err = pd.getNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:150
		_go_fuzz_dep_.CoverTab[97284]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:151
		// _ = "end of CoverTab[97284]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:152
		_go_fuzz_dep_.CoverTab[97285]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:152
		// _ = "end of CoverTab[97285]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:152
	// _ = "end of CoverTab[97278]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:152
	_go_fuzz_dep_.CoverTab[97279]++

													if err := m.Resource.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:154
		_go_fuzz_dep_.CoverTab[97286]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:155
		// _ = "end of CoverTab[97286]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:156
		_go_fuzz_dep_.CoverTab[97287]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:156
		// _ = "end of CoverTab[97287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:156
	// _ = "end of CoverTab[97279]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:156
	_go_fuzz_dep_.CoverTab[97280]++

													if err := m.Acl.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:158
		_go_fuzz_dep_.CoverTab[97288]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:159
		// _ = "end of CoverTab[97288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:160
		_go_fuzz_dep_.CoverTab[97289]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:160
		// _ = "end of CoverTab[97289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:160
	// _ = "end of CoverTab[97280]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:160
	_go_fuzz_dep_.CoverTab[97281]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:162
	// _ = "end of CoverTab[97281]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:163
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_delete_response.go:163
var _ = _go_fuzz_dep_.CoverTab
