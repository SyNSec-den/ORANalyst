/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_SliceAvailableCapacity_H_
#define	_SliceAvailableCapacity_H_


#include "asn_application.h"

/* Including external dependencies */
#include "SliceAvailableCapacityList.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* SliceAvailableCapacity */
typedef struct SliceAvailableCapacity {
	SliceAvailableCapacityList_t	 sliceAvailableCapacityList;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} SliceAvailableCapacity_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_SliceAvailableCapacity;
extern asn_SEQUENCE_specifics_t asn_SPC_SliceAvailableCapacity_specs_1;
extern asn_TYPE_member_t asn_MBR_SliceAvailableCapacity_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _SliceAvailableCapacity_H_ */
#include "asn_internal.h"
