/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_UACPLMN_Item_H_
#define	_UACPLMN_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "PLMN-Identity.h"
#include "UACType-List.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* UACPLMN-Item */
typedef struct UACPLMN_Item {
	PLMN_Identity_t	 pLMNIdentity;
	UACType_List_t	 uACType_List;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} UACPLMN_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_UACPLMN_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_UACPLMN_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_UACPLMN_Item_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _UACPLMN_Item_H_ */
#include "asn_internal.h"
