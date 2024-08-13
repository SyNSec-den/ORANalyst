// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"fmt"
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/input"
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/onos-kpimon/pkg/manager"
	"github.com/onosproject/onos-kpimon/pkg/store/measurements"
	"github.com/onosproject/onos-test/pkg/onostest"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getCredentials() (string, string, error) {
	kubClient, err := kubernetes.New()
	if err != nil {
		return "", "", err
	}
	secrets, err := kubClient.CoreV1().Secrets().Get(context.Background(), onostest.SecretsName)
	if err != nil {
		return "", "", err
	}
	username := string(secrets.Object.Data["sd-ran-username"])
	password := string(secrets.Object.Data["sd-ran-password"])

	return username, password, nil
}

// CreateSdranRelease creates a helm release for an sd-ran instance
func CreateSdranRelease(c *input.Context) (*helm.HelmRelease, error) {
	username, password, err := getCredentials()
	registry := c.GetArg("registry").String("")

	if err != nil {
		return nil, err
	}

	sdran := helm.Chart("sd-ran", onostest.SdranChartRepo).
		Release("sd-ran").
		SetUsername(username).
		SetPassword(password).
		Set("import.onos-config.enabled", false).
		Set("import.onos-topo.enabled", true).
		Set("import.ran-simulator.enabled", true).
		Set("import.onos-pci.enabled", false).
		Set("import.onos-kpimon.enabled", false).
		Set("global.image.tag", "latest").
		Set("global.image.registry", registry)

	return sdran, nil
}

// WaitForKPMIndicationMessages is the function to wait until all KPM indication messages arrives successfully
func WaitForKPMIndicationMessages(ctx context.Context, t *testing.T, mgr *manager.Manager) error {
	store := mgr.GetMeasurementStore()

	for {
		select {
		case <-ctx.Done():
			if verifyMonResults(t, store) {
				return nil
			}
			return fmt.Errorf("%s", "Test failed - the number of cells, e2nodes, or UEs is not matched")
		case <-time.After(TestInterval):
			if verifyMonResults(t, store) {
				return nil
			}
		}
	}
}

func verifyMonResults(t *testing.T, store measurements.Store) bool {
	verify := true
	ch := make(chan *measurements.Entry)
	go func() {
		err := store.Entries(context.Background(), ch)
		if err != nil {
			t.Log(err)
		}
	}()
	mKey := make(map[string]int64)
	mValueAvg := make(map[string]int64)
	for e := range ch {
		mKey[e.Key.NodeID]++
		for _, item := range e.Value.([]measurements.MeasurementItem) {
			var avgNumUEs int64
			for _, record := range item.MeasurementRecords {
				switch record.MeasurementName {
				case AvgUEsMeasName:
					avgNumUEs = record.MeasurementValue.(int64)
				}
			}
			mValueAvg[fmt.Sprintf("%s:%s", e.Key.NodeID, e.Key.CellIdentity.CellID)] = avgNumUEs
		}
	}

	var numCells int64
	var numE2Nodes int64
	for _, v := range mKey {
		numE2Nodes++
		numCells = numCells + v
	}

	if numCells != TotalNumCells {
		t.Logf("Waiting until the number of cells becomes %d; currently it is %d", TotalNumCells, numCells)
		verify = false
	}

	if numE2Nodes != TotalNumE2Nodes {
		t.Logf("Waiting until the number of e2 nodes becomes %d; currently it is %d", TotalNumE2Nodes, numE2Nodes)
		verify = false
	}

	var numUEs int64
	for _, v := range mValueAvg {
		numUEs = numUEs + v
	}

	if numUEs != TotalNumUEs {
		t.Logf("Waiting until the number of UEs becomes %d; currently it is %d", TotalNumUEs, numUEs)
		verify = false
	}

	return verify
}

// CreateRanSimulatorWithNameOrDie creates a simulator and fails the test if the creation returned an error
func CreateRanSimulatorWithNameOrDie(t *testing.T, c *input.Context, simName string) *helm.HelmRelease {
	sim := CreateRanSimulatorWithName(t, c, simName)
	assert.NotNil(t, sim)
	return sim
}

// CreateRanSimulatorWithName creates a ran simulator
func CreateRanSimulatorWithName(t *testing.T, c *input.Context, name string) *helm.HelmRelease {
	username, password, err := getCredentials()
	assert.NoError(t, err)

	registry := c.GetArg("registry").String("")

	simulator := helm.
		Chart("ran-simulator", onostest.SdranChartRepo).
		Release(name).
		SetUsername(username).
		SetPassword(password).
		Set("image.tag", "latest").
		Set("fullnameOverride", "").
		Set("global.image.registry", registry)
	err = simulator.Install(true)
	assert.NoError(t, err, "could not install device simulator %v", err)

	return simulator
}
