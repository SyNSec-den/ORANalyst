/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_TRPTEGInformation_H_
#define	_TRPTEGInformation_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum TRPTEGInformation_PR {
	TRPTEGInformation_PR_NOTHING,	/* No components present */
	TRPTEGInformation_PR_rxTx_TEG,
	TRPTEGInformation_PR_rx_TEG,
	TRPTEGInformation_PR_choice_extension
} TRPTEGInformation_PR;

/* Forward declarations */
struct RxTxTEG;
struct RxTEG;
struct ProtocolIE_SingleContainer;

/* TRPTEGInformation */
typedef struct TRPTEGInformation {
	TRPTEGInformation_PR present;
	union TRPTEGInformation_u {
		struct RxTxTEG	*rxTx_TEG;
		struct RxTEG	*rx_TEG;
		struct ProtocolIE_SingleContainer	*choice_extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TRPTEGInformation_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TRPTEGInformation;
extern asn_CHOICE_specifics_t asn_SPC_TRPTEGInformation_specs_1;
extern asn_TYPE_member_t asn_MBR_TRPTEGInformation_1[3];
extern asn_per_constraints_t asn_PER_type_TRPTEGInformation_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _TRPTEGInformation_H_ */
#include "asn_internal.h"