/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_IntendedTDD_DL_ULConfig_H_
#define	_IntendedTDD_DL_ULConfig_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"
#include "Slot-Configuration-List.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum IntendedTDD_DL_ULConfig__nRSCS {
	IntendedTDD_DL_ULConfig__nRSCS_scs15	= 0,
	IntendedTDD_DL_ULConfig__nRSCS_scs30	= 1,
	IntendedTDD_DL_ULConfig__nRSCS_scs60	= 2,
	IntendedTDD_DL_ULConfig__nRSCS_scs120	= 3,
	/*
	 * Enumeration is extensible
	 */
	IntendedTDD_DL_ULConfig__nRSCS_scs480	= 4,
	IntendedTDD_DL_ULConfig__nRSCS_scs960	= 5
} e_IntendedTDD_DL_ULConfig__nRSCS;
typedef enum IntendedTDD_DL_ULConfig__nRCP {
	IntendedTDD_DL_ULConfig__nRCP_normal	= 0,
	IntendedTDD_DL_ULConfig__nRCP_extended	= 1
	/*
	 * Enumeration is extensible
	 */
} e_IntendedTDD_DL_ULConfig__nRCP;
typedef enum IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity {
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms0p5	= 0,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms0p625	= 1,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms1	= 2,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms1p25	= 3,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms2	= 4,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms2p5	= 5,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms3	= 6,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms4	= 7,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms5	= 8,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms10	= 9,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms20	= 10,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms40	= 11,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms60	= 12,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms80	= 13,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms100	= 14,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms120	= 15,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms140	= 16,
	IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity_ms160	= 17
	/*
	 * Enumeration is extensible
	 */
} e_IntendedTDD_DL_ULConfig__nRDLULTxPeriodicity;

/* Forward declarations */
struct ProtocolExtensionContainer;

/* IntendedTDD-DL-ULConfig */
typedef struct IntendedTDD_DL_ULConfig {
	long	 nRSCS;
	long	 nRCP;
	long	 nRDLULTxPeriodicity;
	Slot_Configuration_List_t	 slot_Configuration_List;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} IntendedTDD_DL_ULConfig_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_nRSCS_2;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_nRCP_10;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_nRDLULTxPeriodicity_14;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_IntendedTDD_DL_ULConfig;
extern asn_SEQUENCE_specifics_t asn_SPC_IntendedTDD_DL_ULConfig_specs_1;
extern asn_TYPE_member_t asn_MBR_IntendedTDD_DL_ULConfig_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _IntendedTDD_DL_ULConfig_H_ */
#include "asn_internal.h"
