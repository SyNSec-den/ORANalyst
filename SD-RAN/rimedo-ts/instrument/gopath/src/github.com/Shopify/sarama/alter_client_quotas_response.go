//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:1
)

import (
	"time"
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:16
type AlterClientQuotasResponse struct {
	ThrottleTime	time.Duration				// The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	Entries		[]AlterClientQuotasEntryResponse	// The quota configuration entries altered.
}

type AlterClientQuotasEntryResponse struct {
	ErrorCode	KError			// The error code, or `0` if the quota alteration succeeded.
	ErrorMsg	*string			// The error message, or `null` if the quota alteration succeeded.
	Entity		[]QuotaEntityComponent	// The quota entity altered.
}

func (a *AlterClientQuotasResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:27
	_go_fuzz_dep_.CoverTab[98053]++

														pe.putInt32(int32(a.ThrottleTime / time.Millisecond))

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:32
	if err := pe.putArrayLength(len(a.Entries)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:32
		_go_fuzz_dep_.CoverTab[98056]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:33
		// _ = "end of CoverTab[98056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:34
		_go_fuzz_dep_.CoverTab[98057]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:34
		// _ = "end of CoverTab[98057]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:34
	// _ = "end of CoverTab[98053]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:34
	_go_fuzz_dep_.CoverTab[98054]++
														for _, e := range a.Entries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:35
		_go_fuzz_dep_.CoverTab[98058]++
															if err := e.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:36
			_go_fuzz_dep_.CoverTab[98059]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:37
			// _ = "end of CoverTab[98059]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:38
			_go_fuzz_dep_.CoverTab[98060]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:38
			// _ = "end of CoverTab[98060]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:38
		// _ = "end of CoverTab[98058]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:39
	// _ = "end of CoverTab[98054]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:39
	_go_fuzz_dep_.CoverTab[98055]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:41
	// _ = "end of CoverTab[98055]"
}

func (a *AlterClientQuotasResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:44
	_go_fuzz_dep_.CoverTab[98061]++

														throttleTime, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:47
		_go_fuzz_dep_.CoverTab[98065]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:48
		// _ = "end of CoverTab[98065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:49
		_go_fuzz_dep_.CoverTab[98066]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:49
		// _ = "end of CoverTab[98066]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:49
	// _ = "end of CoverTab[98061]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:49
	_go_fuzz_dep_.CoverTab[98062]++
														a.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:53
	entryCount, err := pd.getArrayLength()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:54
		_go_fuzz_dep_.CoverTab[98067]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:55
		// _ = "end of CoverTab[98067]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:56
		_go_fuzz_dep_.CoverTab[98068]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:56
		// _ = "end of CoverTab[98068]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:56
	// _ = "end of CoverTab[98062]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:56
	_go_fuzz_dep_.CoverTab[98063]++
														if entryCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:57
		_go_fuzz_dep_.CoverTab[98069]++
															a.Entries = make([]AlterClientQuotasEntryResponse, entryCount)
															for i := range a.Entries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:59
			_go_fuzz_dep_.CoverTab[98070]++
																e := AlterClientQuotasEntryResponse{}
																if err = e.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:61
				_go_fuzz_dep_.CoverTab[98072]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:62
				// _ = "end of CoverTab[98072]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:63
				_go_fuzz_dep_.CoverTab[98073]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:63
				// _ = "end of CoverTab[98073]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:63
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:63
			// _ = "end of CoverTab[98070]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:63
			_go_fuzz_dep_.CoverTab[98071]++
																a.Entries[i] = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:64
			// _ = "end of CoverTab[98071]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:65
		// _ = "end of CoverTab[98069]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:66
		_go_fuzz_dep_.CoverTab[98074]++
															a.Entries = []AlterClientQuotasEntryResponse{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:67
		// _ = "end of CoverTab[98074]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:68
	// _ = "end of CoverTab[98063]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:68
	_go_fuzz_dep_.CoverTab[98064]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:70
	// _ = "end of CoverTab[98064]"
}

func (a *AlterClientQuotasEntryResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:73
	_go_fuzz_dep_.CoverTab[98075]++

														pe.putInt16(int16(a.ErrorCode))

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:78
	if err := pe.putNullableString(a.ErrorMsg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:78
		_go_fuzz_dep_.CoverTab[98079]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:79
		// _ = "end of CoverTab[98079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:80
		_go_fuzz_dep_.CoverTab[98080]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:80
		// _ = "end of CoverTab[98080]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:80
	// _ = "end of CoverTab[98075]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:80
	_go_fuzz_dep_.CoverTab[98076]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:83
	if err := pe.putArrayLength(len(a.Entity)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:83
		_go_fuzz_dep_.CoverTab[98081]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:84
		// _ = "end of CoverTab[98081]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:85
		_go_fuzz_dep_.CoverTab[98082]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:85
		// _ = "end of CoverTab[98082]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:85
	// _ = "end of CoverTab[98076]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:85
	_go_fuzz_dep_.CoverTab[98077]++
														for _, component := range a.Entity {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:86
		_go_fuzz_dep_.CoverTab[98083]++
															if err := component.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:87
			_go_fuzz_dep_.CoverTab[98084]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:88
			// _ = "end of CoverTab[98084]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:89
			_go_fuzz_dep_.CoverTab[98085]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:89
			// _ = "end of CoverTab[98085]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:89
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:89
		// _ = "end of CoverTab[98083]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:90
	// _ = "end of CoverTab[98077]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:90
	_go_fuzz_dep_.CoverTab[98078]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:92
	// _ = "end of CoverTab[98078]"
}

func (a *AlterClientQuotasEntryResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:95
	_go_fuzz_dep_.CoverTab[98086]++

														errCode, err := pd.getInt16()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:98
		_go_fuzz_dep_.CoverTab[98091]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:99
		// _ = "end of CoverTab[98091]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:100
		_go_fuzz_dep_.CoverTab[98092]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:100
		// _ = "end of CoverTab[98092]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:100
	// _ = "end of CoverTab[98086]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:100
	_go_fuzz_dep_.CoverTab[98087]++
														a.ErrorCode = KError(errCode)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:104
	errMsg, err := pd.getNullableString()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:105
		_go_fuzz_dep_.CoverTab[98093]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:106
		// _ = "end of CoverTab[98093]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:107
		_go_fuzz_dep_.CoverTab[98094]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:107
		// _ = "end of CoverTab[98094]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:107
	// _ = "end of CoverTab[98087]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:107
	_go_fuzz_dep_.CoverTab[98088]++
														a.ErrorMsg = errMsg

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:111
	componentCount, err := pd.getArrayLength()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:112
		_go_fuzz_dep_.CoverTab[98095]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:113
		// _ = "end of CoverTab[98095]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:114
		_go_fuzz_dep_.CoverTab[98096]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:114
		// _ = "end of CoverTab[98096]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:114
	// _ = "end of CoverTab[98088]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:114
	_go_fuzz_dep_.CoverTab[98089]++
														if componentCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:115
		_go_fuzz_dep_.CoverTab[98097]++
															a.Entity = make([]QuotaEntityComponent, componentCount)
															for i := 0; i < componentCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:117
			_go_fuzz_dep_.CoverTab[98098]++
																component := QuotaEntityComponent{}
																if err := component.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:119
				_go_fuzz_dep_.CoverTab[98100]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:120
				// _ = "end of CoverTab[98100]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:121
				_go_fuzz_dep_.CoverTab[98101]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:121
				// _ = "end of CoverTab[98101]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:121
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:121
			// _ = "end of CoverTab[98098]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:121
			_go_fuzz_dep_.CoverTab[98099]++
																a.Entity[i] = component
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:122
			// _ = "end of CoverTab[98099]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:123
		// _ = "end of CoverTab[98097]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:124
		_go_fuzz_dep_.CoverTab[98102]++
															a.Entity = []QuotaEntityComponent{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:125
		// _ = "end of CoverTab[98102]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:126
	// _ = "end of CoverTab[98089]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:126
	_go_fuzz_dep_.CoverTab[98090]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:128
	// _ = "end of CoverTab[98090]"
}

func (a *AlterClientQuotasResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:131
	_go_fuzz_dep_.CoverTab[98103]++
														return 49
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:132
	// _ = "end of CoverTab[98103]"
}

func (a *AlterClientQuotasResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:135
	_go_fuzz_dep_.CoverTab[98104]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:136
	// _ = "end of CoverTab[98104]"
}

func (a *AlterClientQuotasResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:139
	_go_fuzz_dep_.CoverTab[98105]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:140
	// _ = "end of CoverTab[98105]"
}

func (a *AlterClientQuotasResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:143
	_go_fuzz_dep_.CoverTab[98106]++
														return V2_6_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:144
	// _ = "end of CoverTab[98106]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:145
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_response.go:145
var _ = _go_fuzz_dep_.CoverTab
