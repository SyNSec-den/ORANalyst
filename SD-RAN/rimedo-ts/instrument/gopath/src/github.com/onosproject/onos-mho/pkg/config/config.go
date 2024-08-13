// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:5
package config

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:5
)

import (
	"context"

	configurable "github.com/onosproject/onos-ric-sdk-go/pkg/config/registry"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	app "github.com/onosproject/onos-ric-sdk-go/pkg/config/app/default"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
	configutils "github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"
)

var log = logging.GetLogger()

const (
	ReportingPeriodConfigPath	= "reportingPeriod"
	PeriodicConfigPath		= "periodic"
	UponRcvMeasConfigPath		= "uponRcvMeasReport"
	UponChangeRrcStatusConfigPath	= "uponChangeRrcStatus"
	A3OffsetRangeConfigPath		= "A3OffsetRange"
	HysteresisRangeConfigPath	= "HysteresisRange"
	CellIndividualOffsetConfigPath	= "CellIndividualOffset"
	FrequencyOffsetConfigPath	= "FrequencyOffset"
	TimeToTriggerConfigPath		= "TimeToTrigger"
)

// Config xApp configuration interface
type Config interface {
	GetReportingPeriod() (uint64, error)
	GetPeriodic() bool
	GetUponRcvMeas() bool
	GetUponChangeRrcStatus() bool
	GetA3OffsetRange() uint64
	GetHysteresisRange() uint64
	GetCellIndividualOffset() uint64
	GetFrequencyOffset() uint64
	GetTimeToTrigger() uint64
	Watch(context.Context, chan event.Event) error
}

// NewConfig initialize the xApp config
func NewConfig(configPath string) (Config, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:47
	_go_fuzz_dep_.CoverTab[194235]++
													appConfig, err := configurable.RegisterConfigurable(configPath, &configurable.RegisterRequest{})
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:49
		_go_fuzz_dep_.CoverTab[194237]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:50
		// _ = "end of CoverTab[194237]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:51
		_go_fuzz_dep_.CoverTab[194238]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:51
		// _ = "end of CoverTab[194238]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:51
	// _ = "end of CoverTab[194235]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:51
	_go_fuzz_dep_.CoverTab[194236]++

													cfg := &mhoConfig{
		appConfig: appConfig.Config.(*app.Config),
	}
													return cfg, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:56
	// _ = "end of CoverTab[194236]"
}

// mhoConfig application configuration
type mhoConfig struct {
	appConfig *app.Config
}

// Watch watch config changes
func (c *mhoConfig) Watch(ctx context.Context, ch chan event.Event) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:65
	_go_fuzz_dep_.CoverTab[194239]++
													err := c.appConfig.Watch(ctx, ch)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:67
		_go_fuzz_dep_.CoverTab[194241]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:68
		// _ = "end of CoverTab[194241]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:69
		_go_fuzz_dep_.CoverTab[194242]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:69
		// _ = "end of CoverTab[194242]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:69
	// _ = "end of CoverTab[194239]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:69
	_go_fuzz_dep_.CoverTab[194240]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:70
	// _ = "end of CoverTab[194240]"
}

// GetReportingPeriod gets configured reporting period
func (c *mhoConfig) GetReportingPeriod() (uint64, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:74
	_go_fuzz_dep_.CoverTab[194243]++
													interval, err := c.appConfig.Get(ReportingPeriodConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:76
		_go_fuzz_dep_.CoverTab[194246]++
														log.Error(err)
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:78
		// _ = "end of CoverTab[194246]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:79
		_go_fuzz_dep_.CoverTab[194247]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:79
		// _ = "end of CoverTab[194247]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:79
	// _ = "end of CoverTab[194243]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:79
	_go_fuzz_dep_.CoverTab[194244]++
													val, err := configutils.ToUint64(interval.Value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:81
		_go_fuzz_dep_.CoverTab[194248]++
														log.Error(err)
														return 0, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:83
		// _ = "end of CoverTab[194248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:84
		_go_fuzz_dep_.CoverTab[194249]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:84
		// _ = "end of CoverTab[194249]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:84
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:84
	// _ = "end of CoverTab[194244]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:84
	_go_fuzz_dep_.CoverTab[194245]++

													return val, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:86
	// _ = "end of CoverTab[194245]"
}

// GetPeriodic returns true if periodic trigger is enabled
func (c *mhoConfig) GetPeriodic() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:90
	_go_fuzz_dep_.CoverTab[194250]++
													p, err := c.appConfig.Get(PeriodicConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:92
		_go_fuzz_dep_.CoverTab[194252]++
														log.Error(err)
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:94
		// _ = "end of CoverTab[194252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:95
		_go_fuzz_dep_.CoverTab[194253]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:95
		// _ = "end of CoverTab[194253]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:95
	// _ = "end of CoverTab[194250]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:95
	_go_fuzz_dep_.CoverTab[194251]++
													switch p.Value.(type) {
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:97
		_go_fuzz_dep_.CoverTab[194254]++
														return p.Value.(bool)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:98
		// _ = "end of CoverTab[194254]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:99
		_go_fuzz_dep_.CoverTab[194255]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:100
		// _ = "end of CoverTab[194255]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:101
	// _ = "end of CoverTab[194251]"
}

// GetUponRcvMeas returns true if periodic trigger is enabled
func (c *mhoConfig) GetUponRcvMeas() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:105
	_go_fuzz_dep_.CoverTab[194256]++
													p, err := c.appConfig.Get(UponRcvMeasConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:107
		_go_fuzz_dep_.CoverTab[194258]++
														log.Error(err)
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:109
		// _ = "end of CoverTab[194258]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:110
		_go_fuzz_dep_.CoverTab[194259]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:110
		// _ = "end of CoverTab[194259]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:110
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:110
	// _ = "end of CoverTab[194256]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:110
	_go_fuzz_dep_.CoverTab[194257]++
													switch p.Value.(type) {
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:112
		_go_fuzz_dep_.CoverTab[194260]++
														return p.Value.(bool)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:113
		// _ = "end of CoverTab[194260]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:114
		_go_fuzz_dep_.CoverTab[194261]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:115
		// _ = "end of CoverTab[194261]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:116
	// _ = "end of CoverTab[194257]"
}

// GetUponChangeRrcStatus returns true if periodic trigger is enabled
func (c *mhoConfig) GetUponChangeRrcStatus() bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:120
	_go_fuzz_dep_.CoverTab[194262]++
													p, err := c.appConfig.Get(UponChangeRrcStatusConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:122
		_go_fuzz_dep_.CoverTab[194264]++
														log.Error(err)
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:124
		// _ = "end of CoverTab[194264]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:125
		_go_fuzz_dep_.CoverTab[194265]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:125
		// _ = "end of CoverTab[194265]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:125
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:125
	// _ = "end of CoverTab[194262]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:125
	_go_fuzz_dep_.CoverTab[194263]++
													switch p.Value.(type) {
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:127
		_go_fuzz_dep_.CoverTab[194266]++
														return p.Value.(bool)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:128
		// _ = "end of CoverTab[194266]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:129
		_go_fuzz_dep_.CoverTab[194267]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:130
		// _ = "end of CoverTab[194267]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:131
	// _ = "end of CoverTab[194263]"
}

// GetA3OffsetRange gets configured reporting period
func (c *mhoConfig) GetA3OffsetRange() uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:135
	_go_fuzz_dep_.CoverTab[194268]++
													x, err := c.appConfig.Get(A3OffsetRangeConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:137
		_go_fuzz_dep_.CoverTab[194271]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:139
		// _ = "end of CoverTab[194271]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:140
		_go_fuzz_dep_.CoverTab[194272]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:140
		// _ = "end of CoverTab[194272]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:140
	// _ = "end of CoverTab[194268]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:140
	_go_fuzz_dep_.CoverTab[194269]++
													val, err := configutils.ToUint64(x.Value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:142
		_go_fuzz_dep_.CoverTab[194273]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:144
		// _ = "end of CoverTab[194273]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:145
		_go_fuzz_dep_.CoverTab[194274]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:145
		// _ = "end of CoverTab[194274]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:145
	// _ = "end of CoverTab[194269]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:145
	_go_fuzz_dep_.CoverTab[194270]++

													return val
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:147
	// _ = "end of CoverTab[194270]"
}

// GetHysteresisRange gets configured reporting period
func (c *mhoConfig) GetHysteresisRange() uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:151
	_go_fuzz_dep_.CoverTab[194275]++
													x, err := c.appConfig.Get(HysteresisRangeConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:153
		_go_fuzz_dep_.CoverTab[194278]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:155
		// _ = "end of CoverTab[194278]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:156
		_go_fuzz_dep_.CoverTab[194279]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:156
		// _ = "end of CoverTab[194279]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:156
	// _ = "end of CoverTab[194275]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:156
	_go_fuzz_dep_.CoverTab[194276]++
													val, err := configutils.ToUint64(x.Value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:158
		_go_fuzz_dep_.CoverTab[194280]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:160
		// _ = "end of CoverTab[194280]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:161
		_go_fuzz_dep_.CoverTab[194281]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:161
		// _ = "end of CoverTab[194281]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:161
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:161
	// _ = "end of CoverTab[194276]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:161
	_go_fuzz_dep_.CoverTab[194277]++

													return val
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:163
	// _ = "end of CoverTab[194277]"
}

// GetCellIndividualOffset gets configured reporting period
func (c *mhoConfig) GetCellIndividualOffset() uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:167
	_go_fuzz_dep_.CoverTab[194282]++
													x, err := c.appConfig.Get(CellIndividualOffsetConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:169
		_go_fuzz_dep_.CoverTab[194285]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:171
		// _ = "end of CoverTab[194285]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:172
		_go_fuzz_dep_.CoverTab[194286]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:172
		// _ = "end of CoverTab[194286]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:172
	// _ = "end of CoverTab[194282]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:172
	_go_fuzz_dep_.CoverTab[194283]++
													val, err := configutils.ToUint64(x.Value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:174
		_go_fuzz_dep_.CoverTab[194287]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:176
		// _ = "end of CoverTab[194287]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:177
		_go_fuzz_dep_.CoverTab[194288]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:177
		// _ = "end of CoverTab[194288]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:177
	// _ = "end of CoverTab[194283]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:177
	_go_fuzz_dep_.CoverTab[194284]++

													return val
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:179
	// _ = "end of CoverTab[194284]"
}

// GetFrequencyOffset gets configured reporting period
func (c *mhoConfig) GetFrequencyOffset() uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:183
	_go_fuzz_dep_.CoverTab[194289]++
													x, err := c.appConfig.Get(FrequencyOffsetConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:185
		_go_fuzz_dep_.CoverTab[194292]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:187
		// _ = "end of CoverTab[194292]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:188
		_go_fuzz_dep_.CoverTab[194293]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:188
		// _ = "end of CoverTab[194293]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:188
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:188
	// _ = "end of CoverTab[194289]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:188
	_go_fuzz_dep_.CoverTab[194290]++
													val, err := configutils.ToUint64(x.Value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:190
		_go_fuzz_dep_.CoverTab[194294]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:192
		// _ = "end of CoverTab[194294]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:193
		_go_fuzz_dep_.CoverTab[194295]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:193
		// _ = "end of CoverTab[194295]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:193
	// _ = "end of CoverTab[194290]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:193
	_go_fuzz_dep_.CoverTab[194291]++

													return val
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:195
	// _ = "end of CoverTab[194291]"
}

// GetTimeToTrigger gets configured reporting period
func (c *mhoConfig) GetTimeToTrigger() uint64 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:199
	_go_fuzz_dep_.CoverTab[194296]++
													x, err := c.appConfig.Get(TimeToTriggerConfigPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:201
		_go_fuzz_dep_.CoverTab[194299]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:203
		// _ = "end of CoverTab[194299]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:204
		_go_fuzz_dep_.CoverTab[194300]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:204
		// _ = "end of CoverTab[194300]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:204
	// _ = "end of CoverTab[194296]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:204
	_go_fuzz_dep_.CoverTab[194297]++
													val, err := configutils.ToUint64(x.Value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:206
		_go_fuzz_dep_.CoverTab[194301]++
														log.Error(err)
														return 0
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:208
		// _ = "end of CoverTab[194301]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:209
		_go_fuzz_dep_.CoverTab[194302]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:209
		// _ = "end of CoverTab[194302]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:209
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:209
	// _ = "end of CoverTab[194297]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:209
	_go_fuzz_dep_.CoverTab[194298]++

													return val
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:211
	// _ = "end of CoverTab[194298]"
}

var _ Config = &mhoConfig{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:214
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-mho@v0.3.1/pkg/config/config.go:214
var _ = _go_fuzz_dep_.CoverTab
