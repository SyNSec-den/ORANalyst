/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_TraceActivation_H_
#define	_TraceActivation_H_


#include "asn_application.h"

/* Including external dependencies */
#include "TraceID.h"
#include "InterfacesToTrace.h"
#include "TraceDepth.h"
#include "TransportLayerAddress.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* TraceActivation */
typedef struct TraceActivation {
	TraceID_t	 traceID;
	InterfacesToTrace_t	 interfacesToTrace;
	TraceDepth_t	 traceDepth;
	TransportLayerAddress_t	 traceCollectionEntityIPAddress;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TraceActivation_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TraceActivation;
extern asn_SEQUENCE_specifics_t asn_SPC_TraceActivation_specs_1;
extern asn_TYPE_member_t asn_MBR_TraceActivation_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _TraceActivation_H_ */
#include "asn_internal.h"
