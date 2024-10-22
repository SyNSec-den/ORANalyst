/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_Child_IAB_Nodes_NA_Resource_List_Item_H_
#define	_Child_IAB_Nodes_NA_Resource_List_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "GNB-CU-UE-F1AP-ID.h"
#include "GNB-DU-UE-F1AP-ID.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct NA_Resource_Configuration_List;
struct ProtocolExtensionContainer;

/* Child-IAB-Nodes-NA-Resource-List-Item */
typedef struct Child_IAB_Nodes_NA_Resource_List_Item {
	GNB_CU_UE_F1AP_ID_t	 gNB_CU_UE_F1AP_ID;
	GNB_DU_UE_F1AP_ID_t	 gNB_DU_UE_F1AP_ID;
	struct NA_Resource_Configuration_List	*nA_Resource_Configuration_List;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Child_IAB_Nodes_NA_Resource_List_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Child_IAB_Nodes_NA_Resource_List_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_Child_IAB_Nodes_NA_Resource_List_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_Child_IAB_Nodes_NA_Resource_List_Item_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _Child_IAB_Nodes_NA_Resource_List_Item_H_ */
#include "asn_internal.h"
