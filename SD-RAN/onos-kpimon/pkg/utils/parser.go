// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	prototypes "github.com/gogo/protobuf/types"
	kpimonapi "github.com/onosproject/onos-api/go/onos/kpimon"
	measurementStore "github.com/onosproject/onos-kpimon/pkg/store/measurements"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("utils", "parser")

// ParseEntry parses measurement store entry
func ParseEntry(entry *measurementStore.Entry) *kpimonapi.MeasurementItems {
	var err error

	measEntryItems := entry.Value.([]measurementStore.MeasurementItem)
	measItem := &kpimonapi.MeasurementItem{}
	measItems := &kpimonapi.MeasurementItems{}
	for _, entryItem := range measEntryItems {
		measItem.MeasurementRecords = make([]*kpimonapi.MeasurementRecord, 0)
		for _, record := range entryItem.MeasurementRecords {
			var value *prototypes.Any
			switch val := record.MeasurementValue.(type) {
			case int64:
				intValue := &kpimonapi.IntegerValue{Value: val}
				value, err = prototypes.MarshalAny(intValue)
				if err != nil {
					log.Warn(err)
					continue
				}

			case float64:
				realValue := &kpimonapi.RealValue{
					Value: val,
				}
				value, err = prototypes.MarshalAny(realValue)
				if err != nil {
					log.Warn(err)
					continue
				}
			case int32:
				noValue := &kpimonapi.NoValue{
					Value: val,
				}
				value, err = prototypes.MarshalAny(noValue)
				if err != nil {
					log.Warn(err)
					continue
				}

			}

			measRecord := &kpimonapi.MeasurementRecord{
				MeasurementName:  record.MeasurementName,
				Timestamp:        record.Timestamp,
				MeasurementValue: value,
			}
			measItem.MeasurementRecords = append(measItem.MeasurementRecords, measRecord)
		}
		measItems.MeasurementItems = append(measItems.MeasurementItems, measItem)
	}
	return measItems
}
