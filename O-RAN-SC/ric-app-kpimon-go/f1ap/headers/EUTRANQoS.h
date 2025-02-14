/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_EUTRANQoS_H_
#define	_EUTRANQoS_H_


#include "asn_application.h"

/* Including external dependencies */
#include "QCI.h"
#include "AllocationAndRetentionPriority.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct GBR_QosInformation;
struct ProtocolExtensionContainer;

/* EUTRANQoS */
typedef struct EUTRANQoS {
	QCI_t	 qCI;
	AllocationAndRetentionPriority_t	 allocationAndRetentionPriority;
	struct GBR_QosInformation	*gbrQosInformation;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} EUTRANQoS_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_EUTRANQoS;
extern asn_SEQUENCE_specifics_t asn_SPC_EUTRANQoS_specs_1;
extern asn_TYPE_member_t asn_MBR_EUTRANQoS_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _EUTRANQoS_H_ */
#include "asn_internal.h"
