/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_CompositeAvailableCapacityGroup_H_
#define	_CompositeAvailableCapacityGroup_H_


#include "asn_application.h"

/* Including external dependencies */
#include "CompositeAvailableCapacity.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* CompositeAvailableCapacityGroup */
typedef struct CompositeAvailableCapacityGroup {
	CompositeAvailableCapacity_t	 compositeAvailableCapacityDownlink;
	CompositeAvailableCapacity_t	 compositeAvailableCapacityUplink;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} CompositeAvailableCapacityGroup_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_CompositeAvailableCapacityGroup;
extern asn_SEQUENCE_specifics_t asn_SPC_CompositeAvailableCapacityGroup_specs_1;
extern asn_TYPE_member_t asn_MBR_CompositeAvailableCapacityGroup_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _CompositeAvailableCapacityGroup_H_ */
#include "asn_internal.h"
