/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_RBSetConfiguration_H_
#define	_RBSetConfiguration_H_


#include "asn_application.h"

/* Including external dependencies */
#include "SubcarrierSpacing.h"
#include "RBSetSize.h"
#include "NativeInteger.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* RBSetConfiguration */
typedef struct RBSetConfiguration {
	SubcarrierSpacing_t	 subcarrierSpacing;
	RBSetSize_t	 rBSetSize;
	long	 nUmberRBsets;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RBSetConfiguration_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RBSetConfiguration;
extern asn_SEQUENCE_specifics_t asn_SPC_RBSetConfiguration_specs_1;
extern asn_TYPE_member_t asn_MBR_RBSetConfiguration_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _RBSetConfiguration_H_ */
#include "asn_internal.h"