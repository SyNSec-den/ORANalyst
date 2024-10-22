/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_Choice_TRP_Beam_Antenna_Info_Item_H_
#define	_Choice_TRP_Beam_Antenna_Info_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "TRPID.h"
#include "NULL.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum Choice_TRP_Beam_Antenna_Info_Item_PR {
	Choice_TRP_Beam_Antenna_Info_Item_PR_NOTHING,	/* No components present */
	Choice_TRP_Beam_Antenna_Info_Item_PR_reference,
	Choice_TRP_Beam_Antenna_Info_Item_PR_explicit,
	Choice_TRP_Beam_Antenna_Info_Item_PR_noChange,
	Choice_TRP_Beam_Antenna_Info_Item_PR_choice_extension
} Choice_TRP_Beam_Antenna_Info_Item_PR;

/* Forward declarations */
struct TRP_BeamAntennaExplicitInformation;
struct ProtocolIE_SingleContainer;

/* Choice-TRP-Beam-Antenna-Info-Item */
typedef struct Choice_TRP_Beam_Antenna_Info_Item {
	Choice_TRP_Beam_Antenna_Info_Item_PR present;
	union Choice_TRP_Beam_Antenna_Info_Item_u {
		TRPID_t	 reference;
		struct TRP_BeamAntennaExplicitInformation	*Explicit;
		NULL_t	 noChange;
		struct ProtocolIE_SingleContainer	*choice_extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Choice_TRP_Beam_Antenna_Info_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_Choice_TRP_Beam_Antenna_Info_Item;
extern asn_CHOICE_specifics_t asn_SPC_Choice_TRP_Beam_Antenna_Info_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_Choice_TRP_Beam_Antenna_Info_Item_1[4];
extern asn_per_constraints_t asn_PER_type_Choice_TRP_Beam_Antenna_Info_Item_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _Choice_TRP_Beam_Antenna_Info_Item_H_ */
#include "asn_internal.h"
