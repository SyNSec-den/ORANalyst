<!--
SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# Honeycomb Topology Generator

The RAN simulator comes with an accompanying utility that generates a RAN topology YAML file that is ready to be loaded by the RAN simulator.

This utility generates a hexagonal grid of RAN towers (E2 Nodes), each with a prescribed number of cells with equal arc of coverage. The following is the command-line usage:

```
Usage:
  honeycomb topo outfile [flags]

Flags:
      --cell-types strings             List of cell size types (default [FEMTO,ENTERPRISE,OUTDOOR_SMALL,MACRO])
      --controller-addresses strings   List of E2T controller addresses or service names (default [onos-e2t])
      --controller-yaml string         if specified, location of yaml file for controller
      --deform-scale float             scale factor for perturbation (default 0.01)
      --earfcn-start uint32            start point for EARFCN generation (default 42)
      --gnbid-start string             GnbID start in hex (default "5152")
  -h, --help                           help for topo
  -a, --latitude float                 Map centre latitude in degrees (default 52.52)
  -g, --longitude float                Map centre longitude in degrees (default 13.405)
      --max-collisions uint            maximum number of collisions (default 8)
  -d, --max-neighbor-distance float    Maximum 'distance' between neighbor cells; see docs (default 3600)
      --max-neighbors int              Maximum number of neighbors a cell will have; -1 no limit (default 5)
      --max-pci uint                   maximum PCI value (default 503)
      --min-pci uint                   minimum PCI value
  -i, --pitch float32                  pitch between cells in degrees (default 0.02)
      --plmnid string                  PlmnID in MCC-MNC format, e.g. CCCNNN or CCCNN (default "315010")
  -s, --sectors-per-tower uint         sectors per tower (default 3)
      --service-models strings         List of service models supported by the nodes (default [kpm/1,rcpre2/3,kpm2/4,mho/5])
      --single-node                    generate a single node for all cells
  -t, --towers uint                    number of towers
      --ue-count uint                  User Equipment count
      --ue-count-per-cell uint         Desired UE count per cell (default 15)
```

Most options have reasonable defaults and only the `--towers` is mandatory. Here is an example of how to run the command to generate topology for a network with PLMNID of `314628` with 10 E2 nodes (towers), each with default number of 3 cells.

```
go run cmd/honeycomb/honeycomb.go topo  --plmnid 314628 --towers 10 pkg/utils/honeycomb/sample.yaml
```

Here is the [output generated by the above command](../pkg/utils/honeycomb/sample.yaml).

Note that for options that are lists, like `--controller-addresses` for example, you can specify them as comma-separated values `--controller-addresses onos-e2t-1,onos-e2t-2` or you can simply repeat the option `--controller-addresses onos-e2t-1 --controller-addresses onos-e2t-2`.

The `--max-neightbor-distance` parameter (specified in meters) works as follows:  if after traveling this distance along a center line of the coverage sector, the endpoint falls within half this distance from another cell's such endpoint, those two cells will be considered neighbors. This is to assure that the two coverage arcs converge sufficiently.

Note that the utility relies on random number generator and therefore its output is not deterministic.