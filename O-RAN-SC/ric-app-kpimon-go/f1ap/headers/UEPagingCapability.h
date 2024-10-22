/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_UEPagingCapability_H_
#define	_UEPagingCapability_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum UEPagingCapability__iNACTIVEStatePODetermination {
	UEPagingCapability__iNACTIVEStatePODetermination_supported	= 0
	/*
	 * Enumeration is extensible
	 */
} e_UEPagingCapability__iNACTIVEStatePODetermination;

/* Forward declarations */
struct ProtocolExtensionContainer;

/* UEPagingCapability */
typedef struct UEPagingCapability {
	long	*iNACTIVEStatePODetermination;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extension;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} UEPagingCapability_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_iNACTIVEStatePODetermination_2;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_UEPagingCapability;
extern asn_SEQUENCE_specifics_t asn_SPC_UEPagingCapability_specs_1;
extern asn_TYPE_member_t asn_MBR_UEPagingCapability_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _UEPagingCapability_H_ */
#include "asn_internal.h"
