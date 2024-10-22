/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_L839Info_H_
#define	_L839Info_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum L839Info__restrictedSetConfig {
	L839Info__restrictedSetConfig_unrestrictedSet	= 0,
	L839Info__restrictedSetConfig_restrictedSetTypeA	= 1,
	L839Info__restrictedSetConfig_restrictedSetTypeB	= 2
	/*
	 * Enumeration is extensible
	 */
} e_L839Info__restrictedSetConfig;

/* Forward declarations */
struct ProtocolExtensionContainer;

/* L839Info */
typedef struct L839Info {
	long	 rootSequenceIndex;
	long	 restrictedSetConfig;
	struct ProtocolExtensionContainer	*iE_Extension;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} L839Info_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_restrictedSetConfig_3;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_L839Info;
extern asn_SEQUENCE_specifics_t asn_SPC_L839Info_specs_1;
extern asn_TYPE_member_t asn_MBR_L839Info_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _L839Info_H_ */
#include "asn_internal.h"
