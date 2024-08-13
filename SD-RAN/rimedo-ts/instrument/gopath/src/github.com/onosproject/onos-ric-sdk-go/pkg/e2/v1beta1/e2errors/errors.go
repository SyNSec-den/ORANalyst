// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:5
package e2errors

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:5
)

import (
	"fmt"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"google.golang.org/grpc/status"
)

// E2APType is an E2AP error type
type E2APType int

// Error type constants
const (
	Unknown	E2APType	= iota

	RICUnspecified

	RICRANFunctionIDInvalid

	RICActionNotSupported

	RICExcessiveActions

	RICDuplicateAction

	RICDuplicateEvent

	RICFunctionResourceLimit

	RICRequestIDUnknown

	RICInconsistentActionSubsequentActionSequence

	RICControlMessageInvalid

	RICCallProcessIDInvalid

	RICServiceUnspecified

	RICServiceFunctionNotRequired

	RICServiceExcessiveFunctions

	RICServiceRICResourceLimit

	ProtocolUnspecified

	ProtocolTransferSyntaxError

	ProtocolAbstractSyntaxErrorReject

	ProtocolAbstractSyntaxErrorIgnoreAndNotify

	ProtocolMessageNotCompatibleWithReceiverState

	ProtocolSemanticError

	ProtocolAbstractSyntaxErrorFalselyConstructedMessage

	MiscUnspecified

	MiscControlProcessingOverload

	MiscHardwareFailure

	MiscOMIntervention
)

// TypedError is a typed error
type TypedError struct {
	// E2APType is the E2AP error type
	E2APType	E2APType
	// Message is the error message
	Message	string
}

func (e *TypedError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:83
	_go_fuzz_dep_.CoverTab[196388]++
															return e.Message
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:84
	// _ = "end of CoverTab[196388]"
}

var _ error = &TypedError{}

// FromGRPC creates a typed error from a gRPC error
func FromGRPC(err error) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:90
	_go_fuzz_dep_.CoverTab[196389]++
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:91
		_go_fuzz_dep_.CoverTab[196394]++
																return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:92
		// _ = "end of CoverTab[196394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:93
		_go_fuzz_dep_.CoverTab[196395]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:93
		// _ = "end of CoverTab[196395]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:93
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:93
	// _ = "end of CoverTab[196389]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:93
	_go_fuzz_dep_.CoverTab[196390]++

															stat, ok := status.FromError(err)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:96
		_go_fuzz_dep_.CoverTab[196396]++
																return New(Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:97
		// _ = "end of CoverTab[196396]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:98
		_go_fuzz_dep_.CoverTab[196397]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:98
		// _ = "end of CoverTab[196397]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:98
	// _ = "end of CoverTab[196390]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:98
	_go_fuzz_dep_.CoverTab[196391]++
															details := stat.Details()

															if len(details) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:101
		_go_fuzz_dep_.CoverTab[196398]++
																return New(Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:102
		// _ = "end of CoverTab[196398]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:103
		_go_fuzz_dep_.CoverTab[196399]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:103
		// _ = "end of CoverTab[196399]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:103
	// _ = "end of CoverTab[196391]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:103
	_go_fuzz_dep_.CoverTab[196392]++

															switch t := details[0].(type) {
	case *e2api.Error:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:106
		_go_fuzz_dep_.CoverTab[196400]++
																cause := t.GetCause().GetCause()
																switch c := cause.(type) {
		case *e2api.Error_Cause_Ric_:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:109
			_go_fuzz_dep_.CoverTab[196402]++
																	switch c.Ric.Type {
			case e2api.Error_Cause_Ric_UNSPECIFIED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:111
				_go_fuzz_dep_.CoverTab[196407]++
																		return New(RICUnspecified, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:112
				// _ = "end of CoverTab[196407]"
			case e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:113
				_go_fuzz_dep_.CoverTab[196408]++
																		return New(RICRANFunctionIDInvalid, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:114
				// _ = "end of CoverTab[196408]"
			case e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:115
				_go_fuzz_dep_.CoverTab[196409]++
																		return New(RICActionNotSupported, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:116
				// _ = "end of CoverTab[196409]"
			case e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:117
				_go_fuzz_dep_.CoverTab[196410]++
																		return New(RICExcessiveActions, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:118
				// _ = "end of CoverTab[196410]"
			case e2api.Error_Cause_Ric_DUPLICATE_ACTION:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:119
				_go_fuzz_dep_.CoverTab[196411]++
																		return New(RICDuplicateAction, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:120
				// _ = "end of CoverTab[196411]"
			case e2api.Error_Cause_Ric_DUPLICATE_EVENT:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:121
				_go_fuzz_dep_.CoverTab[196412]++
																		return New(RICDuplicateEvent, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:122
				// _ = "end of CoverTab[196412]"
			case e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:123
				_go_fuzz_dep_.CoverTab[196413]++
																		return New(RICFunctionResourceLimit, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:124
				// _ = "end of CoverTab[196413]"
			case e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:125
				_go_fuzz_dep_.CoverTab[196414]++
																		return New(RICRequestIDUnknown, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:126
				// _ = "end of CoverTab[196414]"
			case e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:127
				_go_fuzz_dep_.CoverTab[196415]++
																		return New(RICInconsistentActionSubsequentActionSequence, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:128
				// _ = "end of CoverTab[196415]"
			case e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:129
				_go_fuzz_dep_.CoverTab[196416]++
																		return New(RICControlMessageInvalid, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:130
				// _ = "end of CoverTab[196416]"
			case e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:131
				_go_fuzz_dep_.CoverTab[196417]++
																		return New(RICCallProcessIDInvalid, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:132
				// _ = "end of CoverTab[196417]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:133
				_go_fuzz_dep_.CoverTab[196418]++
																		return New(RICUnspecified, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:134
				// _ = "end of CoverTab[196418]"

			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:136
			// _ = "end of CoverTab[196402]"
		case *e2api.Error_Cause_Protocol_:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:137
			_go_fuzz_dep_.CoverTab[196403]++
																	switch c.Protocol.Type {
			case e2api.Error_Cause_Protocol_UNSPECIFIED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:139
				_go_fuzz_dep_.CoverTab[196419]++
																		return New(ProtocolUnspecified, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:140
				// _ = "end of CoverTab[196419]"
			case e2api.Error_Cause_Protocol_TRANSFER_SYNTAX_ERROR:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:141
				_go_fuzz_dep_.CoverTab[196420]++
																		return New(ProtocolTransferSyntaxError, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:142
				// _ = "end of CoverTab[196420]"
			case e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_REJECT:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:143
				_go_fuzz_dep_.CoverTab[196421]++
																		return New(ProtocolAbstractSyntaxErrorReject, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:144
				// _ = "end of CoverTab[196421]"
			case e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:145
				_go_fuzz_dep_.CoverTab[196422]++
																		return New(ProtocolAbstractSyntaxErrorIgnoreAndNotify, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:146
				// _ = "end of CoverTab[196422]"
			case e2api.Error_Cause_Protocol_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:147
				_go_fuzz_dep_.CoverTab[196423]++
																		return New(ProtocolMessageNotCompatibleWithReceiverState, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:148
				// _ = "end of CoverTab[196423]"
			case e2api.Error_Cause_Protocol_SEMANTIC_ERROR:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:149
				_go_fuzz_dep_.CoverTab[196424]++
																		return New(ProtocolSemanticError, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:150
				// _ = "end of CoverTab[196424]"
			case e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:151
				_go_fuzz_dep_.CoverTab[196425]++
																		return New(ProtocolAbstractSyntaxErrorFalselyConstructedMessage, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:152
				// _ = "end of CoverTab[196425]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:152
			default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:152
				_go_fuzz_dep_.CoverTab[196426]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:152
				// _ = "end of CoverTab[196426]"

			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:154
			// _ = "end of CoverTab[196403]"
		case *e2api.Error_Cause_Misc_:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:155
			_go_fuzz_dep_.CoverTab[196404]++
																	switch c.Misc.Type {
			case e2api.Error_Cause_Misc_UNSPECIFIED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:157
				_go_fuzz_dep_.CoverTab[196427]++
																		return New(MiscUnspecified, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:158
				// _ = "end of CoverTab[196427]"
			case e2api.Error_Cause_Misc_CONTROL_PROCESSING_OVERLOAD:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:159
				_go_fuzz_dep_.CoverTab[196428]++
																		return New(MiscControlProcessingOverload, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:160
				// _ = "end of CoverTab[196428]"
			case e2api.Error_Cause_Misc_OM_INTERVENTION:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:161
				_go_fuzz_dep_.CoverTab[196429]++
																		return New(MiscOMIntervention, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:162
				// _ = "end of CoverTab[196429]"
			case e2api.Error_Cause_Misc_HARDWARE_FAILURE:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:163
				_go_fuzz_dep_.CoverTab[196430]++
																		return New(MiscHardwareFailure, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:164
				// _ = "end of CoverTab[196430]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:164
			default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:164
				_go_fuzz_dep_.CoverTab[196431]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:164
				// _ = "end of CoverTab[196431]"

			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:166
			// _ = "end of CoverTab[196404]"
		case *e2api.Error_Cause_RicService_:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:167
			_go_fuzz_dep_.CoverTab[196405]++
																	switch c.RicService.Type {
			case e2api.Error_Cause_RicService_UNSPECIFIED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:169
				_go_fuzz_dep_.CoverTab[196432]++
																		return New(RICServiceUnspecified, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:170
				// _ = "end of CoverTab[196432]"
			case e2api.Error_Cause_RicService_FUNCTION_NOT_REQUIRED:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:171
				_go_fuzz_dep_.CoverTab[196433]++
																		return New(RICServiceFunctionNotRequired, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:172
				// _ = "end of CoverTab[196433]"
			case e2api.Error_Cause_RicService_EXCESSIVE_FUNCTIONS:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:173
				_go_fuzz_dep_.CoverTab[196434]++
																		return New(RICServiceExcessiveFunctions, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:174
				// _ = "end of CoverTab[196434]"
			case e2api.Error_Cause_RicService_RIC_RESOURCE_LIMIT:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:175
				_go_fuzz_dep_.CoverTab[196435]++
																		return New(RICServiceRICResourceLimit, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:176
				// _ = "end of CoverTab[196435]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:176
			default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:176
				_go_fuzz_dep_.CoverTab[196436]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:176
				// _ = "end of CoverTab[196436]"

			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:178
			// _ = "end of CoverTab[196405]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:180
			_go_fuzz_dep_.CoverTab[196406]++
																	return New(Unknown, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:181
			// _ = "end of CoverTab[196406]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:182
		// _ = "end of CoverTab[196400]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:183
		_go_fuzz_dep_.CoverTab[196401]++
																return New(Unknown, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:184
		// _ = "end of CoverTab[196401]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:185
	// _ = "end of CoverTab[196392]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:185
	_go_fuzz_dep_.CoverTab[196393]++
															return New(Unknown, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:186
	// _ = "end of CoverTab[196393]"
}

// New creates a new typed error
func New(t E2APType, msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:190
	_go_fuzz_dep_.CoverTab[196437]++
															if len(args) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:191
		_go_fuzz_dep_.CoverTab[196439]++
																msg = fmt.Sprintf(msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:192
		// _ = "end of CoverTab[196439]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:193
		_go_fuzz_dep_.CoverTab[196440]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:193
		// _ = "end of CoverTab[196440]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:193
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:193
	// _ = "end of CoverTab[196437]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:193
	_go_fuzz_dep_.CoverTab[196438]++
															return &TypedError{
		E2APType:	t,
		Message:	msg,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:197
	// _ = "end of CoverTab[196438]"
}

// NewUnknown returns a new Unknown error
func NewUnknown(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:201
	_go_fuzz_dep_.CoverTab[196441]++
															return New(Unknown, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:202
	// _ = "end of CoverTab[196441]"
}

// NewRICUnspecified returns a new RIC Unspecified error
func NewRICUnspecified(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:206
	_go_fuzz_dep_.CoverTab[196442]++
															return New(RICUnspecified, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:207
	// _ = "end of CoverTab[196442]"
}

// NewRICRANFunctionIDInvalid returns a new RICRANFunctionIDInvalid error
func NewRICRANFunctionIDInvalid(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:211
	_go_fuzz_dep_.CoverTab[196443]++
															return New(RICRANFunctionIDInvalid, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:212
	// _ = "end of CoverTab[196443]"
}

// NewRICActionNotSupported returns a new RICActionNotSupported error
func NewRICActionNotSupported(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:216
	_go_fuzz_dep_.CoverTab[196444]++
															return New(RICActionNotSupported, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:217
	// _ = "end of CoverTab[196444]"
}

// NewRICExcessiveActions returns a new RICExcessiveActions error
func NewRICExcessiveActions(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:221
	_go_fuzz_dep_.CoverTab[196445]++
															return New(RICExcessiveActions, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:222
	// _ = "end of CoverTab[196445]"
}

// NewRICDuplicateAction returns a new RICDuplicateAction  error
func NewRICDuplicateAction(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:226
	_go_fuzz_dep_.CoverTab[196446]++
															return New(RICDuplicateAction, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:227
	// _ = "end of CoverTab[196446]"
}

// NewRICDuplicateEvent returns a new RICDuplicateEvent error
func NewRICDuplicateEvent(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:231
	_go_fuzz_dep_.CoverTab[196447]++
															return New(RICDuplicateEvent, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:232
	// _ = "end of CoverTab[196447]"
}

// NewRICFunctionResourceLimit returns a new RICFunctionResourceLimit error
func NewRICFunctionResourceLimit(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:236
	_go_fuzz_dep_.CoverTab[196448]++
															return New(RICFunctionResourceLimit, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:237
	// _ = "end of CoverTab[196448]"
}

// NewRICRequestIDUnknown returns a new RICRequestIDUnknown error
func NewRICRequestIDUnknown(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:241
	_go_fuzz_dep_.CoverTab[196449]++
															return New(RICRequestIDUnknown, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:242
	// _ = "end of CoverTab[196449]"
}

// NewRICInconsistentActionSubsequentActionSequence returns a new RICInconsistentActionSubsequentActionSequence error
func NewRICInconsistentActionSubsequentActionSequence(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:246
	_go_fuzz_dep_.CoverTab[196450]++
															return New(RICInconsistentActionSubsequentActionSequence, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:247
	// _ = "end of CoverTab[196450]"
}

// NewRICControlMessageInvalid returns a new RICControlMessageInvalid error
func NewRICControlMessageInvalid(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:251
	_go_fuzz_dep_.CoverTab[196451]++
															return New(RICControlMessageInvalid, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:252
	// _ = "end of CoverTab[196451]"
}

// NewRICCallProcessIDInvalid returns a new RICCallProcessIDInvalid error
func NewRICCallProcessIDInvalid(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:256
	_go_fuzz_dep_.CoverTab[196452]++
															return New(RICCallProcessIDInvalid, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:257
	// _ = "end of CoverTab[196452]"
}

// NewRICServiceUnspecified returns a new RICServiceUnspecified error
func NewRICServiceUnspecified(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:261
	_go_fuzz_dep_.CoverTab[196453]++
															return New(RICServiceUnspecified, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:262
	// _ = "end of CoverTab[196453]"
}

// NewRICServiceFunctionNotRequired returns a new 	RICServiceFunctionNotRequired error
func NewRICServiceFunctionNotRequired(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:266
	_go_fuzz_dep_.CoverTab[196454]++
															return New(RICServiceFunctionNotRequired, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:267
	// _ = "end of CoverTab[196454]"
}

// NewRICServiceExcessiveFunctions returns a new RICServiceExcessiveFunctions error
func NewRICServiceExcessiveFunctions(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:271
	_go_fuzz_dep_.CoverTab[196455]++
															return New(RICServiceExcessiveFunctions, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:272
	// _ = "end of CoverTab[196455]"
}

// NewRICServiceRICResourceLimit returns a new 	RICServiceRICResourceLimit error
func NewRICServiceRICResourceLimit(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:276
	_go_fuzz_dep_.CoverTab[196456]++
															return New(RICServiceRICResourceLimit, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:277
	// _ = "end of CoverTab[196456]"
}

// NewProtocolUnspecified returns a new ProtocolUnspecified error
func NewProtocolUnspecified(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:281
	_go_fuzz_dep_.CoverTab[196457]++
															return New(ProtocolUnspecified, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:282
	// _ = "end of CoverTab[196457]"
}

// NewProtocolTransferSyntaxError returns a new ProtocolTransferSyntaxError error
func NewProtocolTransferSyntaxError(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:286
	_go_fuzz_dep_.CoverTab[196458]++
															return New(ProtocolTransferSyntaxError, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:287
	// _ = "end of CoverTab[196458]"
}

// NewProtocolAbstractSyntaxErrorReject returns a new ProtocolAbstractSyntaxErrorReject error
func NewProtocolAbstractSyntaxErrorReject(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:291
	_go_fuzz_dep_.CoverTab[196459]++
															return New(ProtocolAbstractSyntaxErrorReject, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:292
	// _ = "end of CoverTab[196459]"
}

// NewProtocolAbstractSyntaxErrorIgnoreAndNotify returns a new 	ProtocolAbstractSyntaxErrorIgnoreAndNotify error
func NewProtocolAbstractSyntaxErrorIgnoreAndNotify(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:296
	_go_fuzz_dep_.CoverTab[196460]++
															return New(ProtocolAbstractSyntaxErrorIgnoreAndNotify, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:297
	// _ = "end of CoverTab[196460]"
}

// NewProtocolMessageNotCompatibleWithReceiverState returns a new ProtocolMessageNotCompatibleWithReceiverState error
func NewProtocolMessageNotCompatibleWithReceiverState(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:301
	_go_fuzz_dep_.CoverTab[196461]++
															return New(ProtocolMessageNotCompatibleWithReceiverState, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:302
	// _ = "end of CoverTab[196461]"
}

// NewProtocolSemanticError returns a new ProtocolSemanticError error
func NewProtocolSemanticError(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:306
	_go_fuzz_dep_.CoverTab[196462]++
															return New(ProtocolSemanticError, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:307
	// _ = "end of CoverTab[196462]"
}

// NewProtocolAbstractSyntaxErrorFalselyConstructedMessage returns a new ProtocolAbstractSyntaxErrorFalselyConstructedMessage error
func NewProtocolAbstractSyntaxErrorFalselyConstructedMessage(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:311
	_go_fuzz_dep_.CoverTab[196463]++
															return New(ProtocolAbstractSyntaxErrorFalselyConstructedMessage, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:312
	// _ = "end of CoverTab[196463]"
}

// NewMiscUnspecified returns a new MiscUnspecified error
func NewMiscUnspecified(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:316
	_go_fuzz_dep_.CoverTab[196464]++
															return New(MiscUnspecified, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:317
	// _ = "end of CoverTab[196464]"
}

// NewMiscControlProcessingOverload returns a new MiscControlProcessingOverload error
func NewMiscControlProcessingOverload(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:321
	_go_fuzz_dep_.CoverTab[196465]++
															return New(MiscControlProcessingOverload, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:322
	// _ = "end of CoverTab[196465]"
}

// NewMiscHardwareFailure returns a new MiscHardwareFailure error
func NewMiscHardwareFailure(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:326
	_go_fuzz_dep_.CoverTab[196466]++
															return New(MiscHardwareFailure, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:327
	// _ = "end of CoverTab[196466]"
}

// NewMiscOMIntervention returns a new MiscOMIntervention error
func NewMiscOMIntervention(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:331
	_go_fuzz_dep_.CoverTab[196467]++
															return New(MiscOMIntervention, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:332
	// _ = "end of CoverTab[196467]"
}

// TypeOf returns the type of the given error
func TypeOf(err error) E2APType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:336
	_go_fuzz_dep_.CoverTab[196468]++
															if typed, ok := err.(*TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:337
		_go_fuzz_dep_.CoverTab[196470]++
																return typed.E2APType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:338
		// _ = "end of CoverTab[196470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:339
		_go_fuzz_dep_.CoverTab[196471]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:339
		// _ = "end of CoverTab[196471]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:339
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:339
	// _ = "end of CoverTab[196468]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:339
	_go_fuzz_dep_.CoverTab[196469]++
															return Unknown
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:340
	// _ = "end of CoverTab[196469]"
}

// IsType checks whether the given error is of the given type
func IsType(err error, t E2APType) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:344
	_go_fuzz_dep_.CoverTab[196472]++
															if typed, ok := err.(*TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:345
		_go_fuzz_dep_.CoverTab[196474]++
																return typed.E2APType == t
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:346
		// _ = "end of CoverTab[196474]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:347
		_go_fuzz_dep_.CoverTab[196475]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:347
		// _ = "end of CoverTab[196475]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:347
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:347
	// _ = "end of CoverTab[196472]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:347
	_go_fuzz_dep_.CoverTab[196473]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:348
	// _ = "end of CoverTab[196473]"
}

// IsRICUnspecified checks whether the given error is a RIC Unspecified error
func IsRICUnspecified(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:352
	_go_fuzz_dep_.CoverTab[196476]++
															return IsType(err, RICUnspecified)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:353
	// _ = "end of CoverTab[196476]"
}

// IsRICRANFunctionIDInvalid checks whether the given error is a RICRANFunctionIDInvalid error
func IsRICRANFunctionIDInvalid(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:357
	_go_fuzz_dep_.CoverTab[196477]++
															return IsType(err, RICRANFunctionIDInvalid)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:358
	// _ = "end of CoverTab[196477]"
}

// IsRICActionNotSupported checks whether the given error is a RICActionNotSupported error
func IsRICActionNotSupported(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:362
	_go_fuzz_dep_.CoverTab[196478]++
															return IsType(err, RICActionNotSupported)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:363
	// _ = "end of CoverTab[196478]"
}

// IsRICExcessiveActions checks whether the given error is a RICExcessiveActions error
func IsRICExcessiveActions(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:367
	_go_fuzz_dep_.CoverTab[196479]++
															return IsType(err, RICExcessiveActions)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:368
	// _ = "end of CoverTab[196479]"
}

// IsRICDuplicateAction checks whether the given error is a RICDuplicateAction error
func IsRICDuplicateAction(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:372
	_go_fuzz_dep_.CoverTab[196480]++
															return IsType(err, RICDuplicateAction)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:373
	// _ = "end of CoverTab[196480]"
}

// IsRICDuplicateEvent checks whether the given error is a RICDuplicateEvent error
func IsRICDuplicateEvent(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:377
	_go_fuzz_dep_.CoverTab[196481]++
															return IsType(err, RICDuplicateEvent)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:378
	// _ = "end of CoverTab[196481]"
}

// IsRICFunctionResourceLimit checks whether the given error is a 	RICFunctionResourceLimit error
func IsRICFunctionResourceLimit(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:382
	_go_fuzz_dep_.CoverTab[196482]++
															return IsType(err, RICFunctionResourceLimit)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:383
	// _ = "end of CoverTab[196482]"
}

// IsRICRequestIDUnknown checks whether the given error is a RICRequestIDUnknown error
func IsRICRequestIDUnknown(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:387
	_go_fuzz_dep_.CoverTab[196483]++
															return IsType(err, RICRequestIDUnknown)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:388
	// _ = "end of CoverTab[196483]"
}

// IsRICInconsistentActionSubsequentActionSequence checks whether the given error is a 	RICInconsistentActionSubsequentActionSequence error
func IsRICInconsistentActionSubsequentActionSequence(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:392
	_go_fuzz_dep_.CoverTab[196484]++
															return IsType(err, RICInconsistentActionSubsequentActionSequence)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:393
	// _ = "end of CoverTab[196484]"
}

// IsRICControlMessageInvalid checks whether the given error is a RICControlMessageInvalid error
func IsRICControlMessageInvalid(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:397
	_go_fuzz_dep_.CoverTab[196485]++
															return IsType(err, RICControlMessageInvalid)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:398
	// _ = "end of CoverTab[196485]"
}

// IsRICCallProcessIDInvalid checks whether the given error is a RICCallProcessIDInvalid error
func IsRICCallProcessIDInvalid(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:402
	_go_fuzz_dep_.CoverTab[196486]++
															return IsType(err, RICCallProcessIDInvalid)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:403
	// _ = "end of CoverTab[196486]"
}

// IsRICServiceUnspecified checks whether the given error is a RICServiceUnspecified error
func IsRICServiceUnspecified(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:407
	_go_fuzz_dep_.CoverTab[196487]++
															return IsType(err, RICServiceUnspecified)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:408
	// _ = "end of CoverTab[196487]"
}

// IsRICServiceFunctionNotRequired checks whether the given error is a RICServiceFunctionNotRequired error
func IsRICServiceFunctionNotRequired(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:412
	_go_fuzz_dep_.CoverTab[196488]++
															return IsType(err, RICServiceFunctionNotRequired)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:413
	// _ = "end of CoverTab[196488]"
}

// IsRICServiceExcessiveFunctions checks whether the given error is a RICServiceExcessiveFunctions error
func IsRICServiceExcessiveFunctions(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:417
	_go_fuzz_dep_.CoverTab[196489]++
															return IsType(err, RICServiceExcessiveFunctions)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:418
	// _ = "end of CoverTab[196489]"
}

// IsRICServiceRICResourceLimit checks whether the given error is a RICServiceRICResourceLimit error
func IsRICServiceRICResourceLimit(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:422
	_go_fuzz_dep_.CoverTab[196490]++
															return IsType(err, RICServiceRICResourceLimit)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:423
	// _ = "end of CoverTab[196490]"
}

// IsProtocolUnspecified checks whether the given error is a ProtocolUnspecified error
func IsProtocolUnspecified(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:427
	_go_fuzz_dep_.CoverTab[196491]++
															return IsType(err, ProtocolUnspecified)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:428
	// _ = "end of CoverTab[196491]"
}

// IsProtocolTransferSyntaxError checks whether the given error is a ProtocolTransferSyntaxError error
func IsProtocolTransferSyntaxError(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:432
	_go_fuzz_dep_.CoverTab[196492]++
															return IsType(err, ProtocolTransferSyntaxError)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:433
	// _ = "end of CoverTab[196492]"
}

// IsProtocolAbstractSyntaxErrorReject checks whether the given error is a ProtocolAbstractSyntaxErrorReject error
func IsProtocolAbstractSyntaxErrorReject(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:437
	_go_fuzz_dep_.CoverTab[196493]++
															return IsType(err, ProtocolAbstractSyntaxErrorReject)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:438
	// _ = "end of CoverTab[196493]"
}

// IsProtocolAbstractSyntaxErrorIgnoreAndNotify checks whether the given error is a ProtocolAbstractSyntaxErrorIgnoreAndNotify error
func IsProtocolAbstractSyntaxErrorIgnoreAndNotify(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:442
	_go_fuzz_dep_.CoverTab[196494]++
															return IsType(err, ProtocolAbstractSyntaxErrorIgnoreAndNotify)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:443
	// _ = "end of CoverTab[196494]"
}

// IsProtocolMessageNotCompatibleWithReceiverState checks whether the given error is a ProtocolMessageNotCompatibleWithReceiverState error
func IsProtocolMessageNotCompatibleWithReceiverState(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:447
	_go_fuzz_dep_.CoverTab[196495]++
															return IsType(err, ProtocolMessageNotCompatibleWithReceiverState)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:448
	// _ = "end of CoverTab[196495]"
}

// IsProtocolSemanticError checks whether the given error is a 	ProtocolSemanticError error
func IsProtocolSemanticError(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:452
	_go_fuzz_dep_.CoverTab[196496]++
															return IsType(err, ProtocolSemanticError)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:453
	// _ = "end of CoverTab[196496]"
}

// IsProtocolAbstractSyntaxErrorFalselyConstructedMessage checks whether the given error is a ProtocolAbstractSyntaxErrorFalselyConstructedMessage error
func IsProtocolAbstractSyntaxErrorFalselyConstructedMessage(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:457
	_go_fuzz_dep_.CoverTab[196497]++
															return IsType(err, ProtocolAbstractSyntaxErrorFalselyConstructedMessage)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:458
	// _ = "end of CoverTab[196497]"
}

// IsMiscUnspecified checks whether the given error is a MiscUnspecified error
func IsMiscUnspecified(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:462
	_go_fuzz_dep_.CoverTab[196498]++
															return IsType(err, MiscUnspecified)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:463
	// _ = "end of CoverTab[196498]"
}

// IsMiscControlProcessingOverload checks whether the given error is a 	MiscControlProcessingOverload error
func IsMiscControlProcessingOverload(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:467
	_go_fuzz_dep_.CoverTab[196499]++
															return IsType(err, MiscControlProcessingOverload)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:468
	// _ = "end of CoverTab[196499]"
}

// IsMiscHardwareFailure checks whether the given error is a MiscHardwareFailure error
func IsMiscHardwareFailure(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:472
	_go_fuzz_dep_.CoverTab[196500]++
															return IsType(err, MiscHardwareFailure)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:473
	// _ = "end of CoverTab[196500]"
}

// IsMiscOMIntervention checks whether the given error is a MiscOMIntervention error
func IsMiscOMIntervention(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:477
	_go_fuzz_dep_.CoverTab[196501]++
															return IsType(err, MiscOMIntervention)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:478
	// _ = "end of CoverTab[196501]"
}

// IsE2APError checks if a given error is an E2AP error
func IsE2APError(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:482
	_go_fuzz_dep_.CoverTab[196502]++
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:483
		_go_fuzz_dep_.CoverTab[196506]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:484
		// _ = "end of CoverTab[196506]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:485
		_go_fuzz_dep_.CoverTab[196507]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:485
		// _ = "end of CoverTab[196507]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:485
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:485
	// _ = "end of CoverTab[196502]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:485
	_go_fuzz_dep_.CoverTab[196503]++

															stat, ok := status.FromError(err)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:488
		_go_fuzz_dep_.CoverTab[196508]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:489
		// _ = "end of CoverTab[196508]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:490
		_go_fuzz_dep_.CoverTab[196509]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:490
		// _ = "end of CoverTab[196509]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:490
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:490
	// _ = "end of CoverTab[196503]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:490
	_go_fuzz_dep_.CoverTab[196504]++
															details := stat.Details()

															if len(details) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:493
		_go_fuzz_dep_.CoverTab[196510]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:494
		// _ = "end of CoverTab[196510]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:495
		_go_fuzz_dep_.CoverTab[196511]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:495
		// _ = "end of CoverTab[196511]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:495
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:495
	// _ = "end of CoverTab[196504]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:495
	_go_fuzz_dep_.CoverTab[196505]++

															switch details[0].(type) {
	case *e2api.Error:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:498
		_go_fuzz_dep_.CoverTab[196512]++
																return true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:499
		// _ = "end of CoverTab[196512]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:501
		_go_fuzz_dep_.CoverTab[196513]++
																return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:502
		// _ = "end of CoverTab[196513]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:503
	// _ = "end of CoverTab[196505]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:504
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/e2/v1beta1/e2errors/errors.go:504
var _ = _go_fuzz_dep_.CoverTab
