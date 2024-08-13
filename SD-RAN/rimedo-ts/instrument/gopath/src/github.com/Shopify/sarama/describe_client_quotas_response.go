//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:1
)

import (
	"time"
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:19
type DescribeClientQuotasResponse struct {
	ThrottleTime	time.Duration			// The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ErrorCode	KError				// The error code, or `0` if the quota description succeeded.
	ErrorMsg	*string				// The error message, or `null` if the quota description succeeded.
	Entries		[]DescribeClientQuotasEntry	// A result entry.
}

type DescribeClientQuotasEntry struct {
	Entity	[]QuotaEntityComponent	// The quota entity description.
	Values	map[string]float64	// The quota values for the entity.
}

type QuotaEntityComponent struct {
	EntityType	QuotaEntityType
	MatchType	QuotaMatchType
	Name		string
}

func (d *DescribeClientQuotasResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:37
	_go_fuzz_dep_.CoverTab[102062]++

														pe.putInt32(int32(d.ThrottleTime / time.Millisecond))

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:42
	pe.putInt16(int16(d.ErrorCode))

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:45
	if err := pe.putNullableString(d.ErrorMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:45
		_go_fuzz_dep_.CoverTab[102066]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:46
		// _ = "end of CoverTab[102066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:47
		_go_fuzz_dep_.CoverTab[102067]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:47
		// _ = "end of CoverTab[102067]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:47
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:47
	// _ = "end of CoverTab[102062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:47
	_go_fuzz_dep_.CoverTab[102063]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:50
	if err := pe.putArrayLength(len(d.Entries)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:50
		_go_fuzz_dep_.CoverTab[102068]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:51
		// _ = "end of CoverTab[102068]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:52
		_go_fuzz_dep_.CoverTab[102069]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:52
		// _ = "end of CoverTab[102069]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:52
	// _ = "end of CoverTab[102063]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:52
	_go_fuzz_dep_.CoverTab[102064]++
														for _, e := range d.Entries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:53
		_go_fuzz_dep_.CoverTab[102070]++
															if err := e.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:54
			_go_fuzz_dep_.CoverTab[102071]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:55
			// _ = "end of CoverTab[102071]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:56
			_go_fuzz_dep_.CoverTab[102072]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:56
			// _ = "end of CoverTab[102072]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:56
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:56
		// _ = "end of CoverTab[102070]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:57
	// _ = "end of CoverTab[102064]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:57
	_go_fuzz_dep_.CoverTab[102065]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:59
	// _ = "end of CoverTab[102065]"
}

func (d *DescribeClientQuotasResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:62
	_go_fuzz_dep_.CoverTab[102073]++

														throttleTime, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:65
		_go_fuzz_dep_.CoverTab[102079]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:66
		// _ = "end of CoverTab[102079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:67
		_go_fuzz_dep_.CoverTab[102080]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:67
		// _ = "end of CoverTab[102080]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:67
	// _ = "end of CoverTab[102073]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:67
	_go_fuzz_dep_.CoverTab[102074]++
														d.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:71
	errCode, err := pd.getInt16()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:72
		_go_fuzz_dep_.CoverTab[102081]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:73
		// _ = "end of CoverTab[102081]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:74
		_go_fuzz_dep_.CoverTab[102082]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:74
		// _ = "end of CoverTab[102082]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:74
	// _ = "end of CoverTab[102074]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:74
	_go_fuzz_dep_.CoverTab[102075]++
														d.ErrorCode = KError(errCode)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:78
	errMsg, err := pd.getNullableString()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:79
		_go_fuzz_dep_.CoverTab[102083]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:80
		// _ = "end of CoverTab[102083]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:81
		_go_fuzz_dep_.CoverTab[102084]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:81
		// _ = "end of CoverTab[102084]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:81
	// _ = "end of CoverTab[102075]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:81
	_go_fuzz_dep_.CoverTab[102076]++
														d.ErrorMsg = errMsg

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:85
	entryCount, err := pd.getArrayLength()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:86
		_go_fuzz_dep_.CoverTab[102085]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:87
		// _ = "end of CoverTab[102085]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:88
		_go_fuzz_dep_.CoverTab[102086]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:88
		// _ = "end of CoverTab[102086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:88
	// _ = "end of CoverTab[102076]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:88
	_go_fuzz_dep_.CoverTab[102077]++
														if entryCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:89
		_go_fuzz_dep_.CoverTab[102087]++
															d.Entries = make([]DescribeClientQuotasEntry, entryCount)
															for i := range d.Entries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:91
			_go_fuzz_dep_.CoverTab[102088]++
																e := DescribeClientQuotasEntry{}
																if err = e.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:93
				_go_fuzz_dep_.CoverTab[102090]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:94
				// _ = "end of CoverTab[102090]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:95
				_go_fuzz_dep_.CoverTab[102091]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:95
				// _ = "end of CoverTab[102091]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:95
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:95
			// _ = "end of CoverTab[102088]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:95
			_go_fuzz_dep_.CoverTab[102089]++
																d.Entries[i] = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:96
			// _ = "end of CoverTab[102089]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:97
		// _ = "end of CoverTab[102087]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:98
		_go_fuzz_dep_.CoverTab[102092]++
															d.Entries = []DescribeClientQuotasEntry{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:99
		// _ = "end of CoverTab[102092]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:100
	// _ = "end of CoverTab[102077]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:100
	_go_fuzz_dep_.CoverTab[102078]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:102
	// _ = "end of CoverTab[102078]"
}

func (d *DescribeClientQuotasEntry) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:105
	_go_fuzz_dep_.CoverTab[102093]++

														if err := pe.putArrayLength(len(d.Entity)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:107
		_go_fuzz_dep_.CoverTab[102098]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:108
		// _ = "end of CoverTab[102098]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:109
		_go_fuzz_dep_.CoverTab[102099]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:109
		// _ = "end of CoverTab[102099]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:109
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:109
	// _ = "end of CoverTab[102093]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:109
	_go_fuzz_dep_.CoverTab[102094]++
														for _, e := range d.Entity {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:110
		_go_fuzz_dep_.CoverTab[102100]++
															if err := e.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:111
			_go_fuzz_dep_.CoverTab[102101]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:112
			// _ = "end of CoverTab[102101]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:113
			_go_fuzz_dep_.CoverTab[102102]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:113
			// _ = "end of CoverTab[102102]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:113
		// _ = "end of CoverTab[102100]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:114
	// _ = "end of CoverTab[102094]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:114
	_go_fuzz_dep_.CoverTab[102095]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:117
	if err := pe.putArrayLength(len(d.Values)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:117
		_go_fuzz_dep_.CoverTab[102103]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:118
		// _ = "end of CoverTab[102103]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:119
		_go_fuzz_dep_.CoverTab[102104]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:119
		// _ = "end of CoverTab[102104]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:119
	// _ = "end of CoverTab[102095]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:119
	_go_fuzz_dep_.CoverTab[102096]++
														for key, value := range d.Values {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:120
		_go_fuzz_dep_.CoverTab[102105]++

															if err := pe.putString(key); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:122
			_go_fuzz_dep_.CoverTab[102107]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:123
			// _ = "end of CoverTab[102107]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:124
			_go_fuzz_dep_.CoverTab[102108]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:124
			// _ = "end of CoverTab[102108]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:124
		// _ = "end of CoverTab[102105]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:124
		_go_fuzz_dep_.CoverTab[102106]++

															pe.putFloat64(value)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:126
		// _ = "end of CoverTab[102106]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:127
	// _ = "end of CoverTab[102096]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:127
	_go_fuzz_dep_.CoverTab[102097]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:129
	// _ = "end of CoverTab[102097]"
}

func (d *DescribeClientQuotasEntry) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:132
	_go_fuzz_dep_.CoverTab[102109]++

														componentCount, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:135
		_go_fuzz_dep_.CoverTab[102114]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:136
		// _ = "end of CoverTab[102114]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:137
		_go_fuzz_dep_.CoverTab[102115]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:137
		// _ = "end of CoverTab[102115]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:137
	// _ = "end of CoverTab[102109]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:137
	_go_fuzz_dep_.CoverTab[102110]++
														if componentCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:138
		_go_fuzz_dep_.CoverTab[102116]++
															d.Entity = make([]QuotaEntityComponent, componentCount)
															for i := 0; i < componentCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:140
			_go_fuzz_dep_.CoverTab[102117]++
																component := QuotaEntityComponent{}
																if err := component.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:142
				_go_fuzz_dep_.CoverTab[102119]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:143
				// _ = "end of CoverTab[102119]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:144
				_go_fuzz_dep_.CoverTab[102120]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:144
				// _ = "end of CoverTab[102120]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:144
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:144
			// _ = "end of CoverTab[102117]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:144
			_go_fuzz_dep_.CoverTab[102118]++
																d.Entity[i] = component
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:145
			// _ = "end of CoverTab[102118]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:146
		// _ = "end of CoverTab[102116]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:147
		_go_fuzz_dep_.CoverTab[102121]++
															d.Entity = []QuotaEntityComponent{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:148
		// _ = "end of CoverTab[102121]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:149
	// _ = "end of CoverTab[102110]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:149
	_go_fuzz_dep_.CoverTab[102111]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:152
	valueCount, err := pd.getArrayLength()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:153
		_go_fuzz_dep_.CoverTab[102122]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:154
		// _ = "end of CoverTab[102122]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:155
		_go_fuzz_dep_.CoverTab[102123]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:155
		// _ = "end of CoverTab[102123]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:155
	// _ = "end of CoverTab[102111]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:155
	_go_fuzz_dep_.CoverTab[102112]++
														if valueCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:156
		_go_fuzz_dep_.CoverTab[102124]++
															d.Values = make(map[string]float64, valueCount)
															for i := 0; i < valueCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:158
			_go_fuzz_dep_.CoverTab[102125]++

																key, err := pd.getString()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:161
				_go_fuzz_dep_.CoverTab[102128]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:162
				// _ = "end of CoverTab[102128]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:163
				_go_fuzz_dep_.CoverTab[102129]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:163
				// _ = "end of CoverTab[102129]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:163
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:163
			// _ = "end of CoverTab[102125]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:163
			_go_fuzz_dep_.CoverTab[102126]++

																value, err := pd.getFloat64()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:166
				_go_fuzz_dep_.CoverTab[102130]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:167
				// _ = "end of CoverTab[102130]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:168
				_go_fuzz_dep_.CoverTab[102131]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:168
				// _ = "end of CoverTab[102131]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:168
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:168
			// _ = "end of CoverTab[102126]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:168
			_go_fuzz_dep_.CoverTab[102127]++
																d.Values[key] = value
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:169
			// _ = "end of CoverTab[102127]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:170
		// _ = "end of CoverTab[102124]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:171
		_go_fuzz_dep_.CoverTab[102132]++
															d.Values = map[string]float64{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:172
		// _ = "end of CoverTab[102132]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:173
	// _ = "end of CoverTab[102112]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:173
	_go_fuzz_dep_.CoverTab[102113]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:175
	// _ = "end of CoverTab[102113]"
}

func (c *QuotaEntityComponent) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:178
	_go_fuzz_dep_.CoverTab[102133]++

														if err := pe.putString(string(c.EntityType)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:180
		_go_fuzz_dep_.CoverTab[102136]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:181
		// _ = "end of CoverTab[102136]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:182
		_go_fuzz_dep_.CoverTab[102137]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:182
		// _ = "end of CoverTab[102137]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:182
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:182
	// _ = "end of CoverTab[102133]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:182
	_go_fuzz_dep_.CoverTab[102134]++

														if c.MatchType == QuotaMatchDefault {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:184
		_go_fuzz_dep_.CoverTab[102138]++
															if err := pe.putNullableString(nil); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:185
			_go_fuzz_dep_.CoverTab[102139]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:186
			// _ = "end of CoverTab[102139]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:187
			_go_fuzz_dep_.CoverTab[102140]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:187
			// _ = "end of CoverTab[102140]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:187
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:187
		// _ = "end of CoverTab[102138]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:188
		_go_fuzz_dep_.CoverTab[102141]++
															if err := pe.putString(c.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:189
			_go_fuzz_dep_.CoverTab[102142]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:190
			// _ = "end of CoverTab[102142]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:191
			_go_fuzz_dep_.CoverTab[102143]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:191
			// _ = "end of CoverTab[102143]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:191
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:191
		// _ = "end of CoverTab[102141]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:192
	// _ = "end of CoverTab[102134]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:192
	_go_fuzz_dep_.CoverTab[102135]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:194
	// _ = "end of CoverTab[102135]"
}

func (c *QuotaEntityComponent) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:197
	_go_fuzz_dep_.CoverTab[102144]++

														entityType, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:200
		_go_fuzz_dep_.CoverTab[102148]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:201
		// _ = "end of CoverTab[102148]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:202
		_go_fuzz_dep_.CoverTab[102149]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:202
		// _ = "end of CoverTab[102149]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:202
	// _ = "end of CoverTab[102144]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:202
	_go_fuzz_dep_.CoverTab[102145]++
														c.EntityType = QuotaEntityType(entityType)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:206
	entityName, err := pd.getNullableString()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:207
		_go_fuzz_dep_.CoverTab[102150]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:208
		// _ = "end of CoverTab[102150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:209
		_go_fuzz_dep_.CoverTab[102151]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:209
		// _ = "end of CoverTab[102151]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:209
	// _ = "end of CoverTab[102145]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:209
	_go_fuzz_dep_.CoverTab[102146]++

														if entityName == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:211
		_go_fuzz_dep_.CoverTab[102152]++
															c.MatchType = QuotaMatchDefault
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:212
		// _ = "end of CoverTab[102152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:213
		_go_fuzz_dep_.CoverTab[102153]++
															c.MatchType = QuotaMatchExact
															c.Name = *entityName
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:215
		// _ = "end of CoverTab[102153]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:216
	// _ = "end of CoverTab[102146]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:216
	_go_fuzz_dep_.CoverTab[102147]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:218
	// _ = "end of CoverTab[102147]"
}

func (d *DescribeClientQuotasResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:221
	_go_fuzz_dep_.CoverTab[102154]++
														return 48
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:222
	// _ = "end of CoverTab[102154]"
}

func (d *DescribeClientQuotasResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:225
	_go_fuzz_dep_.CoverTab[102155]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:226
	// _ = "end of CoverTab[102155]"
}

func (d *DescribeClientQuotasResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:229
	_go_fuzz_dep_.CoverTab[102156]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:230
	// _ = "end of CoverTab[102156]"
}

func (d *DescribeClientQuotasResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:233
	_go_fuzz_dep_.CoverTab[102157]++
														return V2_6_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:234
	// _ = "end of CoverTab[102157]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:235
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_client_quotas_response.go:235
var _ = _go_fuzz_dep_.CoverTab
