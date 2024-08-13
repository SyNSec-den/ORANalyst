// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/onosproject/helmit/pkg/registry"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/onosproject/onos-kpimon/test/ha"
	"github.com/onosproject/onos-kpimon/test/kpm"
)

func main() {
	registry.RegisterTestSuite("kpm", &kpm.TestSuite{})
	registry.RegisterTestSuite("ha", &ha.TestSuite{})
	test.Main()
}
