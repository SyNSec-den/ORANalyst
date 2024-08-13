// Copy from onosproject/onos-mho/pkg/monitoring/monitor.go
// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
// SPDX-FileCopyrightText: 2019-present Rimedo Labs
//
// SPDX-License-Identifier: Apache-2.0
// modified by RIMEDO-Labs team

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:8
package monitoring

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:8
)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:8
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:8
)

import (
	"context"
	"github.com/onosproject/rimedo-ts/pkg/mho"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-mho/pkg/broker"
)

var log = logging.GetLogger("rimedo-ts", "monitoring")

func NewMonitor(streamReader broker.StreamReader, nodeID topoapi.ID, indChan chan *mho.E2NodeIndication, triggerType e2sm_mho.MhoTriggerType) *Monitor {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:23
	_go_fuzz_dep_.CoverTab[196687]++
										return &Monitor{
		streamReader:	streamReader,
		nodeID:		nodeID,
		indChan:	indChan,
		triggerType:	triggerType,
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:29
	// _ = "end of CoverTab[196687]"
}

type Monitor struct {
	streamReader	broker.StreamReader
	nodeID		topoapi.ID
	indChan		chan *mho.E2NodeIndication
	triggerType	e2sm_mho.MhoTriggerType
}

func (m *Monitor) Start(ctx context.Context) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:39
	_go_fuzz_dep_.CoverTab[196688]++
										errCh := make(chan error)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:40
	_curRoutineNum188_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:40
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum188_)
										go func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:41
		_go_fuzz_dep_.CoverTab[196690]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:41
		defer func() {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:41
			_go_fuzz_dep_.CoverTab[196691]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:41
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum188_)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:41
			// _ = "end of CoverTab[196691]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:41
		}()
											for {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:42
			_go_fuzz_dep_.CoverTab[196692]++
												indMsg, err := m.streamReader.Recv(ctx)
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:44
				_go_fuzz_dep_.CoverTab[196694]++
													log.Errorf("Error reading indication stream, chanID:%v, streamID:%v, err:%v", m.streamReader.ChannelID(), m.streamReader.StreamID(), err)
													errCh <- err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:46
				// _ = "end of CoverTab[196694]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:47
				_go_fuzz_dep_.CoverTab[196695]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:47
				// _ = "end of CoverTab[196695]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:47
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:47
			// _ = "end of CoverTab[196692]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:47
			_go_fuzz_dep_.CoverTab[196693]++
												err = m.processIndication(ctx, indMsg, m.nodeID)
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:49
				_go_fuzz_dep_.CoverTab[196696]++
													log.Errorf("Error processing indication, err:%v", err)
													errCh <- err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:51
				// _ = "end of CoverTab[196696]"
			} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:52
				_go_fuzz_dep_.CoverTab[196697]++
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:52
				// _ = "end of CoverTab[196697]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:52
			}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:52
			// _ = "end of CoverTab[196693]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:53
		// _ = "end of CoverTab[196690]"
	}()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:54
	// _ = "end of CoverTab[196688]"
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:54
	_go_fuzz_dep_.CoverTab[196689]++

										select {
	case err := <-errCh:
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:57
		_go_fuzz_dep_.CoverTab[196698]++
											return err
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:58
		// _ = "end of CoverTab[196698]"
	case <-ctx.Done():
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:59
		_go_fuzz_dep_.CoverTab[196699]++
											return ctx.Err()
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:60
		// _ = "end of CoverTab[196699]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:61
	// _ = "end of CoverTab[196689]"
}

func (m *Monitor) processIndication(ctx context.Context, indication e2api.Indication, nodeID topoapi.ID) error {
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:64
	_go_fuzz_dep_.CoverTab[196700]++

										m.indChan <- &mho.E2NodeIndication{
		NodeID:		string(nodeID),
		TriggerType:	m.triggerType,
		IndMsg: e2api.Indication{
			Payload:	indication.Payload,
			Header:		indication.Header,
		},
	}

										return nil
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:75
	// _ = "end of CoverTab[196700]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:76
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/pkg/monitoring/monitor.go:76
var _ = _go_fuzz_dep_.CoverTab
