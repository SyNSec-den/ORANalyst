/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_FreqInfoRel16_H_
#define	_FreqInfoRel16_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "FrequencyShift7p5khz.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct NRCarrierList;
struct ProtocolExtensionContainer;

/* FreqInfoRel16 */
typedef struct FreqInfoRel16 {
	long	*nRARFCN;	/* OPTIONAL */
	FrequencyShift7p5khz_t	*frequencyShift7p5khz;	/* OPTIONAL */
	struct NRCarrierList	*carrierList;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} FreqInfoRel16_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_FreqInfoRel16;
extern asn_SEQUENCE_specifics_t asn_SPC_FreqInfoRel16_specs_1;
extern asn_TYPE_member_t asn_MBR_FreqInfoRel16_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _FreqInfoRel16_H_ */
#include "asn_internal.h"