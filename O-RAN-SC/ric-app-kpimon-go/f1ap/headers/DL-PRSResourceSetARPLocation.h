/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_DL_PRSResourceSetARPLocation_H_
#define	_DL_PRSResourceSetARPLocation_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum DL_PRSResourceSetARPLocation_PR {
	DL_PRSResourceSetARPLocation_PR_NOTHING,	/* No components present */
	DL_PRSResourceSetARPLocation_PR_relativeGeodeticLocation,
	DL_PRSResourceSetARPLocation_PR_relativeCartesianLocation,
	DL_PRSResourceSetARPLocation_PR_choice_Extension
} DL_PRSResourceSetARPLocation_PR;

/* Forward declarations */
struct RelativeGeodeticLocation;
struct RelativeCartesianLocation;
struct ProtocolIE_SingleContainer;

/* DL-PRSResourceSetARPLocation */
typedef struct DL_PRSResourceSetARPLocation {
	DL_PRSResourceSetARPLocation_PR present;
	union DL_PRSResourceSetARPLocation_u {
		struct RelativeGeodeticLocation	*relativeGeodeticLocation;
		struct RelativeCartesianLocation	*relativeCartesianLocation;
		struct ProtocolIE_SingleContainer	*choice_Extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} DL_PRSResourceSetARPLocation_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_DL_PRSResourceSetARPLocation;
extern asn_CHOICE_specifics_t asn_SPC_DL_PRSResourceSetARPLocation_specs_1;
extern asn_TYPE_member_t asn_MBR_DL_PRSResourceSetARPLocation_1[3];
extern asn_per_constraints_t asn_PER_type_DL_PRSResourceSetARPLocation_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _DL_PRSResourceSetARPLocation_H_ */
#include "asn_internal.h"