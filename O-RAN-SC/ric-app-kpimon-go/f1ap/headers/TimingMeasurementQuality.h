/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_TimingMeasurementQuality_H_
#define	_TimingMeasurementQuality_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum TimingMeasurementQuality__resolution {
	TimingMeasurementQuality__resolution_m0dot1	= 0,
	TimingMeasurementQuality__resolution_m1	= 1,
	TimingMeasurementQuality__resolution_m10	= 2,
	TimingMeasurementQuality__resolution_m30	= 3
	/*
	 * Enumeration is extensible
	 */
} e_TimingMeasurementQuality__resolution;

/* Forward declarations */
struct ProtocolExtensionContainer;

/* TimingMeasurementQuality */
typedef struct TimingMeasurementQuality {
	long	 measurementQuality;
	long	 resolution;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TimingMeasurementQuality_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_resolution_3;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_TimingMeasurementQuality;
extern asn_SEQUENCE_specifics_t asn_SPC_TimingMeasurementQuality_specs_1;
extern asn_TYPE_member_t asn_MBR_TimingMeasurementQuality_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _TimingMeasurementQuality_H_ */
#include "asn_internal.h"
