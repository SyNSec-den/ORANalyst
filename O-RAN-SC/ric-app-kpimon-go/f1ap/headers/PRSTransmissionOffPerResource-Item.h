/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_PRSTransmissionOffPerResource_Item_H_
#define	_PRSTransmissionOffPerResource_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "PRS-Resource-Set-ID.h"
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;
struct PRSTransmissionOffIndicationPerResource_Item;

/* PRSTransmissionOffPerResource-Item */
typedef struct PRSTransmissionOffPerResource_Item {
	PRS_Resource_Set_ID_t	 pRSResourceSetID;
	struct PRSTransmissionOffPerResource_Item__pRSTransmissionOffIndicationPerResourceList {
		A_SEQUENCE_OF(struct PRSTransmissionOffIndicationPerResource_Item) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} pRSTransmissionOffIndicationPerResourceList;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} PRSTransmissionOffPerResource_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_PRSTransmissionOffPerResource_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_PRSTransmissionOffPerResource_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_PRSTransmissionOffPerResource_Item_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _PRSTransmissionOffPerResource_Item_H_ */
#include "asn_internal.h"