/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_TDD_InfoRel16_H_
#define	_TDD_InfoRel16_H_


#include "asn_application.h"

/* Including external dependencies */
#include "TDD-UL-DLConfigCommonNR.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct FreqInfoRel16;
struct ProtocolExtensionContainer;

/* TDD-InfoRel16 */
typedef struct TDD_InfoRel16 {
	struct FreqInfoRel16	*tDD_FreqInfo;	/* OPTIONAL */
	struct FreqInfoRel16	*sUL_FreqInfo;	/* OPTIONAL */
	TDD_UL_DLConfigCommonNR_t	*tDD_UL_DLConfigCommonNR;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TDD_InfoRel16_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TDD_InfoRel16;
extern asn_SEQUENCE_specifics_t asn_SPC_TDD_InfoRel16_specs_1;
extern asn_TYPE_member_t asn_MBR_TDD_InfoRel16_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _TDD_InfoRel16_H_ */
#include "asn_internal.h"