/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_IPHeaderInformation_H_
#define	_IPHeaderInformation_H_


#include "asn_application.h"

/* Including external dependencies */
#include "IABTNLAddress.h"
#include "BIT_STRING.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct DSInformationList;
struct ProtocolExtensionContainer;

/* IPHeaderInformation */
typedef struct IPHeaderInformation {
	IABTNLAddress_t	 destinationIABTNLAddress;
	struct DSInformationList	*dsInformationList;	/* OPTIONAL */
	BIT_STRING_t	*iPv6FlowLabel;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} IPHeaderInformation_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_IPHeaderInformation;
extern asn_SEQUENCE_specifics_t asn_SPC_IPHeaderInformation_specs_1;
extern asn_TYPE_member_t asn_MBR_IPHeaderInformation_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _IPHeaderInformation_H_ */
#include "asn_internal.h"
