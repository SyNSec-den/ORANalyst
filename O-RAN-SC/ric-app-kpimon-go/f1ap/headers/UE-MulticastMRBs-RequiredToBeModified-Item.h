/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_UE_MulticastMRBs_RequiredToBeModified_Item_H_
#define	_UE_MulticastMRBs_RequiredToBeModified_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "MRB-ID.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum UE_MulticastMRBs_RequiredToBeModified_Item__mrb_type_reconfiguration {
	UE_MulticastMRBs_RequiredToBeModified_Item__mrb_type_reconfiguration_true	= 0
	/*
	 * Enumeration is extensible
	 */
} e_UE_MulticastMRBs_RequiredToBeModified_Item__mrb_type_reconfiguration;
typedef enum UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype {
	UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype_rlc_um_ptp	= 0,
	UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype_rlc_am_ptp	= 1,
	UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype_rlc_um_dl_ptm	= 2,
	UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype_two_rlc_um_dl_ptp_and_dl_ptm	= 3,
	UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype_three_rlc_um_dl_ptp_ul_ptp_dl_ptm	= 4,
	UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype_two_rlc_am_ptp_um_dl_ptm	= 5
	/*
	 * Enumeration is extensible
	 */
} e_UE_MulticastMRBs_RequiredToBeModified_Item__mrb_reconfigured_RLCtype;

/* Forward declarations */
struct ProtocolExtensionContainer;

/* UE-MulticastMRBs-RequiredToBeModified-Item */
typedef struct UE_MulticastMRBs_RequiredToBeModified_Item {
	MRB_ID_t	 mRB_ID;
	long	*mrb_type_reconfiguration;	/* OPTIONAL */
	long	*mrb_reconfigured_RLCtype;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} UE_MulticastMRBs_RequiredToBeModified_Item_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_mrb_type_reconfiguration_3;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_mrb_reconfigured_RLCtype_6;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_UE_MulticastMRBs_RequiredToBeModified_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_UE_MulticastMRBs_RequiredToBeModified_Item_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _UE_MulticastMRBs_RequiredToBeModified_Item_H_ */
#include "asn_internal.h"
