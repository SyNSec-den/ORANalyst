/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_Extended_GNB_DU_Name_H_
#define	_Extended_GNB_DU_Name_H_


#include "asn_application.h"

/* Including external dependencies */
#include "GNB-DU-NameVisibleString.h"
#include "GNB-DU-NameUTF8String.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* Extended-GNB-DU-Name */
typedef struct Extended_GNB_DU_Name {
	GNB_DU_NameVisibleString_t	*gNB_DU_NameVisibleString;	/* OPTIONAL */
	GNB_DU_NameUTF8String_t	*gNB_DU_NameUTF8String;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Extended_GNB_DU_Name_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Extended_GNB_DU_Name;
extern asn_SEQUENCE_specifics_t asn_SPC_Extended_GNB_DU_Name_specs_1;
extern asn_TYPE_member_t asn_MBR_Extended_GNB_DU_Name_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _Extended_GNB_DU_Name_H_ */
#include "asn_internal.h"