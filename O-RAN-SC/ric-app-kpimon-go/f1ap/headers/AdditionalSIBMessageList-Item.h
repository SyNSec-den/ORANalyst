/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_AdditionalSIBMessageList_Item_H_
#define	_AdditionalSIBMessageList_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* AdditionalSIBMessageList-Item */
typedef struct AdditionalSIBMessageList_Item {
	OCTET_STRING_t	 additionalSIB;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} AdditionalSIBMessageList_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_AdditionalSIBMessageList_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_AdditionalSIBMessageList_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_AdditionalSIBMessageList_Item_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _AdditionalSIBMessageList_Item_H_ */
#include "asn_internal.h"