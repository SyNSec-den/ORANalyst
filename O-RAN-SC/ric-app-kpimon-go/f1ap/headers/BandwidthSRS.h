/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_BandwidthSRS_H_
#define	_BandwidthSRS_H_


#include "asn_application.h"

/* Including external dependencies */
#include "FR1-Bandwidth.h"
#include "FR2-Bandwidth.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum BandwidthSRS_PR {
	BandwidthSRS_PR_NOTHING,	/* No components present */
	BandwidthSRS_PR_fR1,
	BandwidthSRS_PR_fR2,
	BandwidthSRS_PR_choice_extension
} BandwidthSRS_PR;

/* Forward declarations */
struct ProtocolIE_SingleContainer;

/* BandwidthSRS */
typedef struct BandwidthSRS {
	BandwidthSRS_PR present;
	union BandwidthSRS_u {
		FR1_Bandwidth_t	 fR1;
		FR2_Bandwidth_t	 fR2;
		struct ProtocolIE_SingleContainer	*choice_extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} BandwidthSRS_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_BandwidthSRS;
extern asn_CHOICE_specifics_t asn_SPC_BandwidthSRS_specs_1;
extern asn_TYPE_member_t asn_MBR_BandwidthSRS_1[3];
extern asn_per_constraints_t asn_PER_type_BandwidthSRS_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _BandwidthSRS_H_ */
#include "asn_internal.h"