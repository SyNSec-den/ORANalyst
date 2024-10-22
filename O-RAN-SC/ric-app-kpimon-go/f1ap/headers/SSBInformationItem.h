/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_SSBInformationItem_H_
#define	_SSBInformationItem_H_


#include "asn_application.h"

/* Including external dependencies */
#include "SSB-TF-Configuration.h"
#include "NRPCI.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* SSBInformationItem */
typedef struct SSBInformationItem {
	SSB_TF_Configuration_t	 sSB_Configuration;
	NRPCI_t	 pCI_NR;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} SSBInformationItem_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_SSBInformationItem;
extern asn_SEQUENCE_specifics_t asn_SPC_SSBInformationItem_specs_1;
extern asn_TYPE_member_t asn_MBR_SSBInformationItem_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _SSBInformationItem_H_ */
#include "asn_internal.h"
