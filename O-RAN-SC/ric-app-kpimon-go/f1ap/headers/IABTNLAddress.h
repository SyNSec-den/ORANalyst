/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_IABTNLAddress_H_
#define	_IABTNLAddress_H_


#include "asn_application.h"

/* Including external dependencies */
#include "BIT_STRING.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum IABTNLAddress_PR {
	IABTNLAddress_PR_NOTHING,	/* No components present */
	IABTNLAddress_PR_iPv4Address,
	IABTNLAddress_PR_iPv6Address,
	IABTNLAddress_PR_iPv6Prefix,
	IABTNLAddress_PR_choice_extension
} IABTNLAddress_PR;

/* Forward declarations */
struct ProtocolIE_SingleContainer;

/* IABTNLAddress */
typedef struct IABTNLAddress {
	IABTNLAddress_PR present;
	union IABTNLAddress_u {
		BIT_STRING_t	 iPv4Address;
		BIT_STRING_t	 iPv6Address;
		BIT_STRING_t	 iPv6Prefix;
		struct ProtocolIE_SingleContainer	*choice_extension;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} IABTNLAddress_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_IABTNLAddress;
extern asn_CHOICE_specifics_t asn_SPC_IABTNLAddress_specs_1;
extern asn_TYPE_member_t asn_MBR_IABTNLAddress_1[4];
extern asn_per_constraints_t asn_PER_type_IABTNLAddress_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _IABTNLAddress_H_ */
#include "asn_internal.h"
