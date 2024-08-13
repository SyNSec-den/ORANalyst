//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:1
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:14
type AlterClientQuotasRequest struct {
	Entries		[]AlterClientQuotasEntry	// The quota configuration entries to alter.
	ValidateOnly	bool				// Whether the alteration should be validated, but not performed.
}

type AlterClientQuotasEntry struct {
	Entity	[]QuotaEntityComponent	// The quota entity to alter.
	Ops	[]ClientQuotasOp	// An individual quota configuration entry to alter.
}

type ClientQuotasOp struct {
	Key	string	// The quota configuration key.
	Value	float64	// The value to set, otherwise ignored if the value is to be removed.
	Remove	bool	// Whether the quota configuration value should be removed, otherwise set.
}

func (a *AlterClientQuotasRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:30
	_go_fuzz_dep_.CoverTab[97977]++

													if err := pe.putArrayLength(len(a.Entries)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:32
		_go_fuzz_dep_.CoverTab[97980]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:33
		// _ = "end of CoverTab[97980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:34
		_go_fuzz_dep_.CoverTab[97981]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:34
		// _ = "end of CoverTab[97981]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:34
	// _ = "end of CoverTab[97977]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:34
	_go_fuzz_dep_.CoverTab[97978]++
													for _, e := range a.Entries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:35
		_go_fuzz_dep_.CoverTab[97982]++
														if err := e.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:36
			_go_fuzz_dep_.CoverTab[97983]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:37
			// _ = "end of CoverTab[97983]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:38
			_go_fuzz_dep_.CoverTab[97984]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:38
			// _ = "end of CoverTab[97984]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:38
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:38
		// _ = "end of CoverTab[97982]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:39
	// _ = "end of CoverTab[97978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:39
	_go_fuzz_dep_.CoverTab[97979]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:42
	pe.putBool(a.ValidateOnly)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:44
	// _ = "end of CoverTab[97979]"
}

func (a *AlterClientQuotasRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:47
	_go_fuzz_dep_.CoverTab[97985]++

													entryCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:50
		_go_fuzz_dep_.CoverTab[97989]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:51
		// _ = "end of CoverTab[97989]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:52
		_go_fuzz_dep_.CoverTab[97990]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:52
		// _ = "end of CoverTab[97990]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:52
	// _ = "end of CoverTab[97985]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:52
	_go_fuzz_dep_.CoverTab[97986]++
													if entryCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:53
		_go_fuzz_dep_.CoverTab[97991]++
														a.Entries = make([]AlterClientQuotasEntry, entryCount)
														for i := range a.Entries {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:55
			_go_fuzz_dep_.CoverTab[97992]++
															e := AlterClientQuotasEntry{}
															if err = e.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:57
				_go_fuzz_dep_.CoverTab[97994]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:58
				// _ = "end of CoverTab[97994]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:59
				_go_fuzz_dep_.CoverTab[97995]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:59
				// _ = "end of CoverTab[97995]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:59
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:59
			// _ = "end of CoverTab[97992]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:59
			_go_fuzz_dep_.CoverTab[97993]++
															a.Entries[i] = e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:60
			// _ = "end of CoverTab[97993]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:61
		// _ = "end of CoverTab[97991]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:62
		_go_fuzz_dep_.CoverTab[97996]++
														a.Entries = []AlterClientQuotasEntry{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:63
		// _ = "end of CoverTab[97996]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:64
	// _ = "end of CoverTab[97986]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:64
	_go_fuzz_dep_.CoverTab[97987]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:67
	validateOnly, err := pd.getBool()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:68
		_go_fuzz_dep_.CoverTab[97997]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:69
		// _ = "end of CoverTab[97997]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:70
		_go_fuzz_dep_.CoverTab[97998]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:70
		// _ = "end of CoverTab[97998]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:70
	// _ = "end of CoverTab[97987]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:70
	_go_fuzz_dep_.CoverTab[97988]++
													a.ValidateOnly = validateOnly

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:73
	// _ = "end of CoverTab[97988]"
}

func (a *AlterClientQuotasEntry) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:76
	_go_fuzz_dep_.CoverTab[97999]++

													if err := pe.putArrayLength(len(a.Entity)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:78
		_go_fuzz_dep_.CoverTab[98004]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:79
		// _ = "end of CoverTab[98004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:80
		_go_fuzz_dep_.CoverTab[98005]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:80
		// _ = "end of CoverTab[98005]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:80
	// _ = "end of CoverTab[97999]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:80
	_go_fuzz_dep_.CoverTab[98000]++
													for _, component := range a.Entity {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:81
		_go_fuzz_dep_.CoverTab[98006]++
														if err := component.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:82
			_go_fuzz_dep_.CoverTab[98007]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:83
			// _ = "end of CoverTab[98007]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:84
			_go_fuzz_dep_.CoverTab[98008]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:84
			// _ = "end of CoverTab[98008]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:84
		// _ = "end of CoverTab[98006]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:85
	// _ = "end of CoverTab[98000]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:85
	_go_fuzz_dep_.CoverTab[98001]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:88
	if err := pe.putArrayLength(len(a.Ops)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:88
		_go_fuzz_dep_.CoverTab[98009]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:89
		// _ = "end of CoverTab[98009]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:90
		_go_fuzz_dep_.CoverTab[98010]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:90
		// _ = "end of CoverTab[98010]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:90
	// _ = "end of CoverTab[98001]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:90
	_go_fuzz_dep_.CoverTab[98002]++
													for _, o := range a.Ops {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:91
		_go_fuzz_dep_.CoverTab[98011]++
														if err := o.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:92
			_go_fuzz_dep_.CoverTab[98012]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:93
			// _ = "end of CoverTab[98012]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:94
			_go_fuzz_dep_.CoverTab[98013]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:94
			// _ = "end of CoverTab[98013]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:94
		// _ = "end of CoverTab[98011]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:95
	// _ = "end of CoverTab[98002]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:95
	_go_fuzz_dep_.CoverTab[98003]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:97
	// _ = "end of CoverTab[98003]"
}

func (a *AlterClientQuotasEntry) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:100
	_go_fuzz_dep_.CoverTab[98014]++

														componentCount, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:103
		_go_fuzz_dep_.CoverTab[98019]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:104
		// _ = "end of CoverTab[98019]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:105
		_go_fuzz_dep_.CoverTab[98020]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:105
		// _ = "end of CoverTab[98020]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:105
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:105
	// _ = "end of CoverTab[98014]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:105
	_go_fuzz_dep_.CoverTab[98015]++
														if componentCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:106
		_go_fuzz_dep_.CoverTab[98021]++
															a.Entity = make([]QuotaEntityComponent, componentCount)
															for i := 0; i < componentCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:108
			_go_fuzz_dep_.CoverTab[98022]++
																component := QuotaEntityComponent{}
																if err := component.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:110
				_go_fuzz_dep_.CoverTab[98024]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:111
				// _ = "end of CoverTab[98024]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:112
				_go_fuzz_dep_.CoverTab[98025]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:112
				// _ = "end of CoverTab[98025]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:112
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:112
			// _ = "end of CoverTab[98022]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:112
			_go_fuzz_dep_.CoverTab[98023]++
																a.Entity[i] = component
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:113
			// _ = "end of CoverTab[98023]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:114
		// _ = "end of CoverTab[98021]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:115
		_go_fuzz_dep_.CoverTab[98026]++
															a.Entity = []QuotaEntityComponent{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:116
		// _ = "end of CoverTab[98026]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:117
	// _ = "end of CoverTab[98015]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:117
	_go_fuzz_dep_.CoverTab[98016]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:120
	opCount, err := pd.getArrayLength()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:121
		_go_fuzz_dep_.CoverTab[98027]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:122
		// _ = "end of CoverTab[98027]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:123
		_go_fuzz_dep_.CoverTab[98028]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:123
		// _ = "end of CoverTab[98028]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:123
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:123
	// _ = "end of CoverTab[98016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:123
	_go_fuzz_dep_.CoverTab[98017]++
														if opCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:124
		_go_fuzz_dep_.CoverTab[98029]++
															a.Ops = make([]ClientQuotasOp, opCount)
															for i := range a.Ops {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:126
			_go_fuzz_dep_.CoverTab[98030]++
																c := ClientQuotasOp{}
																if err = c.decode(pd, version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:128
				_go_fuzz_dep_.CoverTab[98032]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:129
				// _ = "end of CoverTab[98032]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:130
				_go_fuzz_dep_.CoverTab[98033]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:130
				// _ = "end of CoverTab[98033]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:130
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:130
			// _ = "end of CoverTab[98030]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:130
			_go_fuzz_dep_.CoverTab[98031]++
																a.Ops[i] = c
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:131
			// _ = "end of CoverTab[98031]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:132
		// _ = "end of CoverTab[98029]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:133
		_go_fuzz_dep_.CoverTab[98034]++
															a.Ops = []ClientQuotasOp{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:134
		// _ = "end of CoverTab[98034]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:135
	// _ = "end of CoverTab[98017]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:135
	_go_fuzz_dep_.CoverTab[98018]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:137
	// _ = "end of CoverTab[98018]"
}

func (c *ClientQuotasOp) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:140
	_go_fuzz_dep_.CoverTab[98035]++

														if err := pe.putString(c.Key); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:142
		_go_fuzz_dep_.CoverTab[98037]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:143
		// _ = "end of CoverTab[98037]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:144
		_go_fuzz_dep_.CoverTab[98038]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:144
		// _ = "end of CoverTab[98038]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:144
	// _ = "end of CoverTab[98035]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:144
	_go_fuzz_dep_.CoverTab[98036]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:147
	pe.putFloat64(c.Value)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:150
	pe.putBool(c.Remove)

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:152
	// _ = "end of CoverTab[98036]"
}

func (c *ClientQuotasOp) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:155
	_go_fuzz_dep_.CoverTab[98039]++

														key, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:158
		_go_fuzz_dep_.CoverTab[98043]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:159
		// _ = "end of CoverTab[98043]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:160
		_go_fuzz_dep_.CoverTab[98044]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:160
		// _ = "end of CoverTab[98044]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:160
	// _ = "end of CoverTab[98039]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:160
	_go_fuzz_dep_.CoverTab[98040]++
														c.Key = key

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:164
	value, err := pd.getFloat64()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:165
		_go_fuzz_dep_.CoverTab[98045]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:166
		// _ = "end of CoverTab[98045]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:167
		_go_fuzz_dep_.CoverTab[98046]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:167
		// _ = "end of CoverTab[98046]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:167
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:167
	// _ = "end of CoverTab[98040]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:167
	_go_fuzz_dep_.CoverTab[98041]++
														c.Value = value

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:171
	remove, err := pd.getBool()
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:172
		_go_fuzz_dep_.CoverTab[98047]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:173
		// _ = "end of CoverTab[98047]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:174
		_go_fuzz_dep_.CoverTab[98048]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:174
		// _ = "end of CoverTab[98048]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:174
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:174
	// _ = "end of CoverTab[98041]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:174
	_go_fuzz_dep_.CoverTab[98042]++
														c.Remove = remove

														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:177
	// _ = "end of CoverTab[98042]"
}

func (a *AlterClientQuotasRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:180
	_go_fuzz_dep_.CoverTab[98049]++
														return 49
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:181
	// _ = "end of CoverTab[98049]"
}

func (a *AlterClientQuotasRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:184
	_go_fuzz_dep_.CoverTab[98050]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:185
	// _ = "end of CoverTab[98050]"
}

func (a *AlterClientQuotasRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:188
	_go_fuzz_dep_.CoverTab[98051]++
														return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:189
	// _ = "end of CoverTab[98051]"
}

func (a *AlterClientQuotasRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:192
	_go_fuzz_dep_.CoverTab[98052]++
														return V2_6_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:193
	// _ = "end of CoverTab[98052]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:194
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_client_quotas_request.go:194
var _ = _go_fuzz_dep_.CoverTab
