// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/kpimon"
	"github.com/onosproject/onos-ric-sdk-go/pkg/utils/creds"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"testing"
)

// KPIMonServiceAddress defines the address and port for connections to the KPIMON service
const KPIMonServiceAddress = "onos-kpimon:5150"

// ConnectKPIMonServiceHost connects to the onos KPIMon service
func ConnectKPIMonServiceHost() (*grpc.ClientConn, error) {
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	return grpc.DialContext(context.Background(), KPIMonServiceAddress, opts...)
}

// GetKPIMonClient returns an SDK subscription client
func GetKPIMonClient(t *testing.T) kpimon.KpimonClient {
	conn, err := ConnectKPIMonServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return kpimon.NewKpimonClient(conn)
}
