#!/bin/sh

# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src":"${GOPATH}/src/github.com/onosproject/onos-kpimon/api"

# samples below
# admin.proto cannot be generated with fast marshaler/unmarshaler because it uses gnmi.ModelData
#protoc -I=$proto_imports --doc_out=docs/api  --doc_opt=markdown,admin.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api/admin,plugins=grpc:. api/admin/v1/*.proto
#protoc -I=$proto_imports --doc_out=docs/api  --doc_opt=markdown,diags.md --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mconfig/admin/admin.proto=github.com/onosproject/onos-e2t/api/admin,import_path=github.com/onosproject/onos-e2t/api/diags,plugins=grpc:. api/diags/*.proto

#protoc -I=$proto_imports  --doc_out=docs/api  --doc_opt=markdown,headers.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api,plugins=grpc:. api/ricapi/e2/headers/v1beta1/headers.proto
#protoc -I=$proto_imports  --doc_out=docs/api  --doc_opt=markdown,ricapi.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api,plugins=grpc:. api/ricapi/e2/v1beta1/ricapi_e2.proto
#protoc -I=$proto_imports  --doc_out=docs/api  --doc_opt=markdown,subscription.md  --gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api/subscription/v1beta1,plugins=grpc:. api/subscription/v1beta1/subscription.proto


#protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_commondatatypes.proto
#protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_constants.proto
#protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_containers.proto
#protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_ies.proto
#protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_pdu_contents.proto
#protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=api --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-e2t/api:. api/e2ap/v1beta1/e2ap_pdu_descriptions.proto

cp -r github.com/onosproject/onos-kpimon/* .
rm -rf github.com
