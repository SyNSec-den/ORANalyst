/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_BCBearerContextF1U_TNLInfo_H_
#define	_BCBearerContextF1U_TNLInfo_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum BCBearerContextF1U_TNLInfo_PR {
	BCBearerContextF1U_TNLInfo_PR_NOTHING,	/* No components present */
	BCBearerContextF1U_TNLInfo_PR_locationindpendent,
	BCBearerContextF1U_TNLInfo_PR_locationdependent,
	BCBearerContextF1U_TNLInfo_PR_choice_extension
} BCBearerContextF1U_TNLInfo_PR;

/* Forward declarations */
struct MBSF1UInformation;
struct LocationDependentMBSF1UInformation;
struct ProtocolIE_SingleContainer;

/* BCBearerContextF1U-TNLInfo */
typedef struct BCBearerContextF1U_TNLInfo {
	BCBearerContextF1U_TNLInfo_PR present;
	union BCBearerContextF1U_TNLInfo_u {
		struct MBSF1UInformation	*locationindpendent;
		struct LocationDependentMBSF1UInformation	*locationdependent;
		struct ProtocolIE_SingleContainer	*choice_extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} BCBearerContextF1U_TNLInfo_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_BCBearerContextF1U_TNLInfo;
extern asn_CHOICE_specifics_t asn_SPC_BCBearerContextF1U_TNLInfo_specs_1;
extern asn_TYPE_member_t asn_MBR_BCBearerContextF1U_TNLInfo_1[3];
extern asn_per_constraints_t asn_PER_type_BCBearerContextF1U_TNLInfo_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _BCBearerContextF1U_TNLInfo_H_ */
#include "asn_internal.h"
