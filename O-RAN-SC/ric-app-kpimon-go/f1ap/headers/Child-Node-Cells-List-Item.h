/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_Child_Node_Cells_List_Item_H_
#define	_Child_Node_Cells_List_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NRCGI.h"
#include "RACH-Config-Common.h"
#include "RACH-Config-Common-IAB.h"
#include "OCTET_STRING.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct IAB_DU_Cell_Resource_Configuration_Mode_Info;
struct IAB_STC_Info;
struct MultiplexingInfo;
struct ProtocolExtensionContainer;

/* Child-Node-Cells-List-Item */
typedef struct Child_Node_Cells_List_Item {
	NRCGI_t	 nRCGI;
	struct IAB_DU_Cell_Resource_Configuration_Mode_Info	*iAB_DU_Cell_Resource_Configuration_Mode_Info;	/* OPTIONAL */
	struct IAB_STC_Info	*iAB_STC_Info;	/* OPTIONAL */
	RACH_Config_Common_t	*rACH_Config_Common;	/* OPTIONAL */
	RACH_Config_Common_IAB_t	*rACH_Config_Common_IAB;	/* OPTIONAL */
	OCTET_STRING_t	*cSI_RS_Configuration;	/* OPTIONAL */
	OCTET_STRING_t	*sR_Configuration;	/* OPTIONAL */
	OCTET_STRING_t	*pDCCH_ConfigSIB1;	/* OPTIONAL */
	OCTET_STRING_t	*sCS_Common;	/* OPTIONAL */
	struct MultiplexingInfo	*multiplexingInfo;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Child_Node_Cells_List_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Child_Node_Cells_List_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_Child_Node_Cells_List_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_Child_Node_Cells_List_Item_1[11];

#ifdef __cplusplus
}
#endif

#endif	/* _Child_Node_Cells_List_Item_H_ */
#include "asn_internal.h"