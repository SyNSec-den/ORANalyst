//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:1
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:10
// A filter to be applied to matching client quotas.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:10
// Components: the components to filter on
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:10
// Strict: whether the filter only includes specified components
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:13
type DescribeClientQuotasRequest struct {
	Components	[]QuotaFilterComponent
	Strict		bool
}

// Describe a component for applying a client quota filter.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:18
// EntityType: the entity type the filter component applies to ("user", "client-id", "ip")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:18
// MatchType: the match type of the filter component (any, exact, default)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:18
// Match: the name that's matched exactly (used when MatchType is QuotaMatchExact)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:22
type QuotaFilterComponent struct {
	EntityType	QuotaEntityType
	MatchType	QuotaMatchType
	Match		string
}

func (d *DescribeClientQuotasRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:28
	_go_fuzz_dep_.CoverTab[102008]++

														if err := pe.putArrayLength(len(d.Components)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:30
		_go_fuzz_dep_.CoverTab[102011]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:31
		// _ = "end of CoverTab[102011]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:32
		_go_fuzz_dep_.CoverTab[102012]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:32
		// _ = "end of CoverTab[102012]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:32
	// _ = "end of CoverTab[102008]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:32
	_go_fuzz_dep_.CoverTab[102009]++
														for _, c := range d.Components {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:33
		_go_fuzz_dep_.CoverTab[102013]++
															if err := c.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:34
			_go_fuzz_dep_.CoverTab[102014]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:35
			// _ = "end of CoverTab[102014]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:36
			_go_fuzz_dep_.CoverTab[102015]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:36
			// _ = "end of CoverTab[102015]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:36
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:36
		// _ = "end of CoverTab[102013]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:37
	// _ = "end of CoverTab[102009]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:37
	_go_fuzz_dep_.CoverTab[102010]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:40
	pe.putBool(d.Strict)

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:42
	// _ = "end of CoverTab[102010]"
}

func (d *DescribeClientQuotasRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:45
	_go_fuzz_dep_.CoverTab[102016]++

														componentCount, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:48
		_go_fuzz_dep_.CoverTab[102020]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:49
		// _ = "end of CoverTab[102020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:50
		_go_fuzz_dep_.CoverTab[102021]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:50
		// _ = "end of CoverTab[102021]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:50
	// _ = "end of CoverTab[102016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:50
	_go_fuzz_dep_.CoverTab[102017]++
														if componentCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:51
		_go_fuzz_dep_.CoverTab[102022]++
															d.Components = make([]QuotaFilterComponent, componentCount)
															for i := range d.Components {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:53
			_go_fuzz_dep_.CoverTab[102023]++
																c := QuotaFilterComponent{}
																if err = c.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:55
				_go_fuzz_dep_.CoverTab[102025]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:56
				// _ = "end of CoverTab[102025]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:57
				_go_fuzz_dep_.CoverTab[102026]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:57
				// _ = "end of CoverTab[102026]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:57
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:57
			// _ = "end of CoverTab[102023]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:57
			_go_fuzz_dep_.CoverTab[102024]++
																d.Components[i] = c
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:58
			// _ = "end of CoverTab[102024]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:59
		// _ = "end of CoverTab[102022]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:60
		_go_fuzz_dep_.CoverTab[102027]++
															d.Components = []QuotaFilterComponent{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:61
		// _ = "end of CoverTab[102027]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:62
	// _ = "end of CoverTab[102017]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:62
	_go_fuzz_dep_.CoverTab[102018]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:65
	strict, err := pd.getBool()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:66
		_go_fuzz_dep_.CoverTab[102028]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:67
		// _ = "end of CoverTab[102028]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:68
		_go_fuzz_dep_.CoverTab[102029]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:68
		// _ = "end of CoverTab[102029]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:68
	// _ = "end of CoverTab[102018]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:68
	_go_fuzz_dep_.CoverTab[102019]++
														d.Strict = strict

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:71
	// _ = "end of CoverTab[102019]"
}

func (d *QuotaFilterComponent) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:74
	_go_fuzz_dep_.CoverTab[102030]++

														if err := pe.putString(string(d.EntityType)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:76
		_go_fuzz_dep_.CoverTab[102033]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:77
		// _ = "end of CoverTab[102033]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:78
		_go_fuzz_dep_.CoverTab[102034]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:78
		// _ = "end of CoverTab[102034]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:78
	// _ = "end of CoverTab[102030]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:78
	_go_fuzz_dep_.CoverTab[102031]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:81
	pe.putInt8(int8(d.MatchType))

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:84
	if d.MatchType == QuotaMatchAny {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:84
		_go_fuzz_dep_.CoverTab[102035]++
															if err := pe.putNullableString(nil); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:85
			_go_fuzz_dep_.CoverTab[102036]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:86
			// _ = "end of CoverTab[102036]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:87
			_go_fuzz_dep_.CoverTab[102037]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:87
			// _ = "end of CoverTab[102037]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:87
		// _ = "end of CoverTab[102035]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:88
		_go_fuzz_dep_.CoverTab[102038]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:88
		if d.MatchType == QuotaMatchDefault {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:88
			_go_fuzz_dep_.CoverTab[102039]++
																if err := pe.putString(""); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:89
				_go_fuzz_dep_.CoverTab[102040]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:90
				// _ = "end of CoverTab[102040]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:91
				_go_fuzz_dep_.CoverTab[102041]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:91
				// _ = "end of CoverTab[102041]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:91
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:91
			// _ = "end of CoverTab[102039]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:92
			_go_fuzz_dep_.CoverTab[102042]++
																if err := pe.putString(d.Match); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:93
				_go_fuzz_dep_.CoverTab[102043]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:94
				// _ = "end of CoverTab[102043]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:95
				_go_fuzz_dep_.CoverTab[102044]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:95
				// _ = "end of CoverTab[102044]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:95
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:95
			// _ = "end of CoverTab[102042]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:96
		// _ = "end of CoverTab[102038]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:96
	// _ = "end of CoverTab[102031]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:96
	_go_fuzz_dep_.CoverTab[102032]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:98
	// _ = "end of CoverTab[102032]"
}

func (d *QuotaFilterComponent) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:101
	_go_fuzz_dep_.CoverTab[102045]++

														entityType, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:104
		_go_fuzz_dep_.CoverTab[102050]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:105
		// _ = "end of CoverTab[102050]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:106
		_go_fuzz_dep_.CoverTab[102051]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:106
		// _ = "end of CoverTab[102051]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:106
	// _ = "end of CoverTab[102045]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:106
	_go_fuzz_dep_.CoverTab[102046]++
														d.EntityType = QuotaEntityType(entityType)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:110
	matchType, err := pd.getInt8()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:111
		_go_fuzz_dep_.CoverTab[102052]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:112
		// _ = "end of CoverTab[102052]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:113
		_go_fuzz_dep_.CoverTab[102053]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:113
		// _ = "end of CoverTab[102053]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:113
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:113
	// _ = "end of CoverTab[102046]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:113
	_go_fuzz_dep_.CoverTab[102047]++
														d.MatchType = QuotaMatchType(matchType)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:117
	match, err := pd.getNullableString()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:118
		_go_fuzz_dep_.CoverTab[102054]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:119
		// _ = "end of CoverTab[102054]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:120
		_go_fuzz_dep_.CoverTab[102055]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:120
		// _ = "end of CoverTab[102055]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:120
	// _ = "end of CoverTab[102047]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:120
	_go_fuzz_dep_.CoverTab[102048]++
														if match != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:121
		_go_fuzz_dep_.CoverTab[102056]++
															d.Match = *match
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:122
		// _ = "end of CoverTab[102056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:123
		_go_fuzz_dep_.CoverTab[102057]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:123
		// _ = "end of CoverTab[102057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:123
	// _ = "end of CoverTab[102048]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:123
	_go_fuzz_dep_.CoverTab[102049]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:124
	// _ = "end of CoverTab[102049]"
}

func (d *DescribeClientQuotasRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:127
	_go_fuzz_dep_.CoverTab[102058]++
														return 48
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:128
	// _ = "end of CoverTab[102058]"
}

func (d *DescribeClientQuotasRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:131
	_go_fuzz_dep_.CoverTab[102059]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:132
	// _ = "end of CoverTab[102059]"
}

func (d *DescribeClientQuotasRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:135
	_go_fuzz_dep_.CoverTab[102060]++
														return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:136
	// _ = "end of CoverTab[102060]"
}

func (d *DescribeClientQuotasRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:139
	_go_fuzz_dep_.CoverTab[102061]++
														return V2_6_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:140
	// _ = "end of CoverTab[102061]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_request.go:141
var _ = _go_fuzz_dep_.CoverTab
