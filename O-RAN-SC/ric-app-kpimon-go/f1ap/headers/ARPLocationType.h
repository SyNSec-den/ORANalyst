/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_ARPLocationType_H_
#define	_ARPLocationType_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ARPLocationType_PR {
	ARPLocationType_PR_NOTHING,	/* No components present */
	ARPLocationType_PR_aRPPositionRelativeGeodetic,
	ARPLocationType_PR_aRPPositionRelativeCartesian,
	ARPLocationType_PR_choice_extension
} ARPLocationType_PR;

/* Forward declarations */
struct RelativeGeodeticLocation;
struct RelativeCartesianLocation;
struct ProtocolIE_SingleContainer;

/* ARPLocationType */
typedef struct ARPLocationType {
	ARPLocationType_PR present;
	union ARPLocationType_u {
		struct RelativeGeodeticLocation	*aRPPositionRelativeGeodetic;
		struct RelativeCartesianLocation	*aRPPositionRelativeCartesian;
		struct ProtocolIE_SingleContainer	*choice_extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ARPLocationType_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ARPLocationType;
extern asn_CHOICE_specifics_t asn_SPC_ARPLocationType_specs_1;
extern asn_TYPE_member_t asn_MBR_ARPLocationType_1[3];
extern asn_per_constraints_t asn_PER_type_ARPLocationType_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _ARPLocationType_H_ */
#include "asn_internal.h"
