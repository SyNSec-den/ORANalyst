/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_ResponseTime_H_
#define	_ResponseTime_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ResponseTime__timeUnit {
	ResponseTime__timeUnit_second	= 0,
	ResponseTime__timeUnit_ten_seconds	= 1,
	ResponseTime__timeUnit_ten_milliseconds	= 2
	/*
	 * Enumeration is extensible
	 */
} e_ResponseTime__timeUnit;

/* Forward declarations */
struct ProtocolExtensionContainer;

/* ResponseTime */
typedef struct ResponseTime {
	long	 time;
	long	 timeUnit;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ResponseTime_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_timeUnit_3;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_ResponseTime;
extern asn_SEQUENCE_specifics_t asn_SPC_ResponseTime_specs_1;
extern asn_TYPE_member_t asn_MBR_ResponseTime_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _ResponseTime_H_ */
#include "asn_internal.h"