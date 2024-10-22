/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_E_CID_MeasuredResults_Value_H_
#define	_E_CID_MeasuredResults_Value_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum E_CID_MeasuredResults_Value_PR {
	E_CID_MeasuredResults_Value_PR_NOTHING,	/* No components present */
	E_CID_MeasuredResults_Value_PR_valueAngleofArrivalNR,
	E_CID_MeasuredResults_Value_PR_choice_extension
} E_CID_MeasuredResults_Value_PR;

/* Forward declarations */
struct UL_AoA;
struct ProtocolIE_SingleContainer;

/* E-CID-MeasuredResults-Value */
typedef struct E_CID_MeasuredResults_Value {
	E_CID_MeasuredResults_Value_PR present;
	union E_CID_MeasuredResults_Value_u {
		struct UL_AoA	*valueAngleofArrivalNR;
		struct ProtocolIE_SingleContainer	*choice_extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E_CID_MeasuredResults_Value_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E_CID_MeasuredResults_Value;
extern asn_CHOICE_specifics_t asn_SPC_E_CID_MeasuredResults_Value_specs_1;
extern asn_TYPE_member_t asn_MBR_E_CID_MeasuredResults_Value_1[2];
extern asn_per_constraints_t asn_PER_type_E_CID_MeasuredResults_Value_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _E_CID_MeasuredResults_Value_H_ */
#include "asn_internal.h"
