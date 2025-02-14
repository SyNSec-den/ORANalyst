/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_Served_Cells_To_Delete_Item_H_
#define	_Served_Cells_To_Delete_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NRCGI.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* Served-Cells-To-Delete-Item */
typedef struct Served_Cells_To_Delete_Item {
	NRCGI_t	 oldNRCGI;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Served_Cells_To_Delete_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Served_Cells_To_Delete_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_Served_Cells_To_Delete_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_Served_Cells_To_Delete_Item_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _Served_Cells_To_Delete_Item_H_ */
#include "asn_internal.h"
