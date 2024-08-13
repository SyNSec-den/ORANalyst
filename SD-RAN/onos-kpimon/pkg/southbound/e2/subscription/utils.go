// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package subscription

import (
	"context"
	"sort"

	actionsstore "github.com/onosproject/onos-kpimon/pkg/store/actions"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"google.golang.org/protobuf/proto"
)

// createSubscriptionActions creates subscription actions
func (m *Manager) createSubscriptionActions(ctx context.Context, reportStyle *topoapi.KPMReportStyle, cells []*topoapi.E2Cell, granularity int64) ([]e2api.Action, error) {
	actions := make([]e2api.Action, 0)

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].CellObjectID < cells[j].CellObjectID
	})

	for index, cell := range cells {
		measInfoList := &e2smkpmv2.MeasurementInfoList{
			Value: make([]*e2smkpmv2.MeasurementInfoItem, 0),
		}

		for _, measurement := range reportStyle.Measurements {
			measTypeMeasName, err := pdubuilder.CreateMeasurementTypeMeasName(measurement.GetName())
			if err != nil {
				return nil, err
			}

			meanInfoItem, err := pdubuilder.CreateMeasurementInfoItem(measTypeMeasName)
			if err != nil {
				return nil, err
			}
			measInfoList.Value = append(measInfoList.Value, meanInfoItem)
		}
		subID := int64(index + 1)
		actionDefinition, err := pdubuilder.CreateActionDefinitionFormat1(cell.GetCellObjectID(), measInfoList, granularity, subID)
		if err != nil {
			return nil, err
		}

		key := actionsstore.NewKey(actionsstore.SubscriptionID{
			SubID: subID,
		})
		// TODO clean up this store if we delete subscriptions
		_, err = m.actionStore.Put(ctx, key, actionDefinition)
		if err != nil {
			log.Warn(err)
			return nil, err
		}

		e2smKpmActionDefinition, err := pdubuilder.CreateE2SmKpmActionDefinitionFormat1(reportStyle.Type, actionDefinition)
		if err != nil {
			return nil, err
		}

		e2smKpmActionDefinitionProto, err := proto.Marshal(e2smKpmActionDefinition)
		if err != nil {
			return nil, err
		}

		action := &e2api.Action{
			ID:   int32(index),
			Type: e2api.ActionType_ACTION_TYPE_REPORT,
			SubsequentAction: &e2api.SubsequentAction{
				Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
				TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
			},
			Payload: e2smKpmActionDefinitionProto,
		}

		actions = append(actions, *action)

	}
	return actions, nil
}
