/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_FiveG_ProSeAuthorized_H_
#define	_FiveG_ProSeAuthorized_H_


#include "asn_application.h"

/* Including external dependencies */
#include "FiveG-ProSeDirectDiscovery.h"
#include "FiveG-ProSeDirectCommunication.h"
#include "FiveG-ProSeLayer2UEtoNetworkRelay.h"
#include "FiveG-ProSeLayer3UEtoNetworkRelay.h"
#include "FiveG-ProSeLayer2RemoteUE.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* FiveG-ProSeAuthorized */
typedef struct FiveG_ProSeAuthorized {
	FiveG_ProSeDirectDiscovery_t	*fiveG_proSeDirectDiscovery;	/* OPTIONAL */
	FiveG_ProSeDirectCommunication_t	*fiveG_proSeDirectCommunication;	/* OPTIONAL */
	FiveG_ProSeLayer2UEtoNetworkRelay_t	*fiveG_ProSeLayer2UEtoNetworkRelay;	/* OPTIONAL */
	FiveG_ProSeLayer3UEtoNetworkRelay_t	*fiveG_ProSeLayer3UEtoNetworkRelay;	/* OPTIONAL */
	FiveG_ProSeLayer2RemoteUE_t	*fiveG_ProSeLayer2RemoteUE;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} FiveG_ProSeAuthorized_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_FiveG_ProSeAuthorized;
extern asn_SEQUENCE_specifics_t asn_SPC_FiveG_ProSeAuthorized_specs_1;
extern asn_TYPE_member_t asn_MBR_FiveG_ProSeAuthorized_1[6];

#ifdef __cplusplus
}
#endif

#endif	/* _FiveG_ProSeAuthorized_H_ */
#include "asn_internal.h"