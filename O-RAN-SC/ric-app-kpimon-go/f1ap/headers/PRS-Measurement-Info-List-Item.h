/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_PRS_Measurement_Info_List_Item_H_
#define	_PRS_Measurement_Info_List_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum PRS_Measurement_Info_List_Item__measPRSPeriodicity {
	PRS_Measurement_Info_List_Item__measPRSPeriodicity_ms20	= 0,
	PRS_Measurement_Info_List_Item__measPRSPeriodicity_ms40	= 1,
	PRS_Measurement_Info_List_Item__measPRSPeriodicity_ms80	= 2,
	PRS_Measurement_Info_List_Item__measPRSPeriodicity_ms160	= 3
	/*
	 * Enumeration is extensible
	 */
} e_PRS_Measurement_Info_List_Item__measPRSPeriodicity;
typedef enum PRS_Measurement_Info_List_Item__measurementPRSLength {
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms1dot5	= 0,
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms3	= 1,
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms3dot5	= 2,
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms4	= 3,
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms5dot5	= 4,
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms6	= 5,
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms10	= 6,
	PRS_Measurement_Info_List_Item__measurementPRSLength_ms20	= 7
} e_PRS_Measurement_Info_List_Item__measurementPRSLength;

/* Forward declarations */
struct ProtocolExtensionContainer;

/* PRS-Measurement-Info-List-Item */
typedef struct PRS_Measurement_Info_List_Item {
	long	 pointA;
	long	 measPRSPeriodicity;
	long	 measPRSOffset;
	long	 measurementPRSLength;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} PRS_Measurement_Info_List_Item_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_measPRSPeriodicity_3;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_measurementPRSLength_10;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_PRS_Measurement_Info_List_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_PRS_Measurement_Info_List_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_PRS_Measurement_Info_List_Item_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _PRS_Measurement_Info_List_Item_H_ */
#include "asn_internal.h"
