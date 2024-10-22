/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-COMMON-IEs"
 * 	found in "e2sm-v03.01.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -no-gen-OER -D /tmp/workspace/oransim-gerrit/e2sim/asn1c/`
 */

#ifndef	_CoreCPID_H_
#define	_CoreCPID_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum CoreCPID_PR {
	CoreCPID_PR_NOTHING,	/* No components present */
	CoreCPID_PR_fiveGC,
	CoreCPID_PR_ePC
	/* Extensions may appear below */
	
} CoreCPID_PR;

/* Forward declarations */
struct GUAMI;
struct GUMMEI;

/* CoreCPID */
typedef struct CoreCPID {
	CoreCPID_PR present;
	union CoreCPID_u {
		struct GUAMI	*fiveGC;
		struct GUMMEI	*ePC;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} CoreCPID_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_CoreCPID;

#ifdef __cplusplus
}
#endif

#endif	/* _CoreCPID_H_ */
#include "asn_internal.h"
