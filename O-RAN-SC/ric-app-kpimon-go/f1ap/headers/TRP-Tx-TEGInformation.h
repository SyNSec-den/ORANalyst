/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_TRP_Tx_TEGInformation_H_
#define	_TRP_Tx_TEGInformation_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "TimingErrorMargin.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* TRP-Tx-TEGInformation */
typedef struct TRP_Tx_TEGInformation {
	long	 tRP_Tx_TEGID;
	TimingErrorMargin_t	 tRP_Tx_TimingErrorMargin;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TRP_Tx_TEGInformation_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TRP_Tx_TEGInformation;
extern asn_SEQUENCE_specifics_t asn_SPC_TRP_Tx_TEGInformation_specs_1;
extern asn_TYPE_member_t asn_MBR_TRP_Tx_TEGInformation_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _TRP_Tx_TEGInformation_H_ */
#include "asn_internal.h"