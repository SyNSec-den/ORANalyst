/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_PC5RLCChannelRequiredToBeModifiedList_H_
#define	_PC5RLCChannelRequiredToBeModifiedList_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct PC5RLCChannelRequiredToBeModifiedItem;

/* PC5RLCChannelRequiredToBeModifiedList */
typedef struct PC5RLCChannelRequiredToBeModifiedList {
	A_SEQUENCE_OF(struct PC5RLCChannelRequiredToBeModifiedItem) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} PC5RLCChannelRequiredToBeModifiedList_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_PC5RLCChannelRequiredToBeModifiedList;
extern asn_SET_OF_specifics_t asn_SPC_PC5RLCChannelRequiredToBeModifiedList_specs_1;
extern asn_TYPE_member_t asn_MBR_PC5RLCChannelRequiredToBeModifiedList_1[1];
extern asn_per_constraints_t asn_PER_type_PC5RLCChannelRequiredToBeModifiedList_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _PC5RLCChannelRequiredToBeModifiedList_H_ */
#include "asn_internal.h"