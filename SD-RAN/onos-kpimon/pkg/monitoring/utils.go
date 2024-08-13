// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package monitoring

import (
	"encoding/binary"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
)

func toUnixNano(timeStamp int64) int64 {
	timeStampUnix := time.Unix(timeStamp, 0).UnixNano()
	return timeStampUnix
}

func getMeasurementName(measID string, measurements []*topoapi.KPMMeasurement) string {
	for _, measurement := range measurements {
		if measurement.GetID() == measID {
			return measurement.GetName()
		}
	}
	return ""
}

func getTimeStampFromHeader(header *e2smkpmv2.E2SmKpmIndicationHeaderFormat1) uint64 {
	timeBytes := (*header).GetColletStartTime().Value
	timeInt32 := binary.BigEndian.Uint32(timeBytes)
	return uint64(timeInt32)
}
