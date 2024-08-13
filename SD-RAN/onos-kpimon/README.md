<!--
SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# onos-kpimon
The xApplication for ONOS SD-RAN (µONOS Architecture) to monitor KPI

## Overview
The `onos-kpimon` is the xApplication running over ONOS SD-RAN to monitor the KPI.
`onos-kpimon` collects KPIs reported by E2 nodes through the KPM service model version 2.0.
Since ONOS SD-RAN has multiple micro-services running on the Kubernetes platform, `onos-kpimon` should run on the Kubernetes along with the other ONOS SD-RAN micro-services.
In order to deploy `onos-kpimon` on the Kubernetes, a Helm chart is necessary, which is in the `sdran-helm-charts` repository.
Note that this application should be running together with the other SD-RAN micro-services, such as `Atomix`, `onos-operator`, `onos-e2t`, `onos-uenib`, `onos-topo`, and `onos-cli`).
Easily, `sd-ran` umbrella chart can be used to deploy all essential micro-services and `onos-kpimon`.

## Interaction with the other ONOS SD-RAN micro-services
To begin with, `onos-kpimon` makes a subscription with E2 nodes connected to `onos-e2t` through `onos-topo` based ONOS xApplication SDK.
Creating a subscription, `onos-kpimon` sets `report interval` and `granularity period` which are the monitoring interval parameters.
Once the subscription is done successfully, each E2 node starts sending indication messages periodically to report KPIs to `onos-kpimon`.
Then, `onos-kpimon` decodes each indication message that has KPI monitoring reports and store them to both KPIMON local store, or `onos-uenib`.
A user can check the stored monitoring results through `onos-cli` as below.
Also, if Prometheus and Grafana are enabled, the user can see the stored monitoring results through Grafana dashboard or Prometheus web GUI.

## Command Line Interface
Go to the `onos-cli`, and command below:
```bash
$ onos kpimon list metrics
Node ID          Cell Object ID       Cell Global ID            Time    RRC.Conn.Avg    RRC.Conn.Max    RRC.ConnEstabAtt.Sum    RRC.ConnEstabSucc.Sum    RRC.ConnReEstabAtt.HOFail    RRC.ConnReEstabAtt.Other    RRC.ConnReEstabAtt.Sum    RRC.ConnReEstabAtt.reconfigFail
5153            13842601454c001             1454c001      06:23:44.0               0               4                       0                        0                            0                           0                         0                                  0
5153            13842601454c002             1454c002      06:23:44.0               0               1                       0                        0                            0                           0                         0                                  0
5153            13842601454c003             1454c003      06:23:44.0               6               6                       0                        0                            0                           0                         0                                  0
5154            138426014550001             14550001      06:23:44.0               0               5                       0                        0                            0                           0                         0                                  0
5154            138426014550002             14550002      06:23:44.0               4               4                       0                        0                            0                           0                         0                                  0
5154            138426014550003             14550003      06:23:44.0               0               2                       0                        0                            0                           0                         0                                  0
```
