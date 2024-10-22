/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_GNB_DU_Cell_Resource_Configuration_H_
#define	_GNB_DU_Cell_Resource_Configuration_H_


#include "asn_application.h"

/* Including external dependencies */
#include "SubcarrierSpacing.h"
#include "DUFTransmissionPeriodicity.h"
#include "HSNATransmissionPeriodicity.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct DUF_Slot_Config_List;
struct HSNASlotConfigList;
struct ProtocolExtensionContainer;

/* GNB-DU-Cell-Resource-Configuration */
typedef struct GNB_DU_Cell_Resource_Configuration {
	SubcarrierSpacing_t	 subcarrierSpacing;
	DUFTransmissionPeriodicity_t	*dUFTransmissionPeriodicity;	/* OPTIONAL */
	struct DUF_Slot_Config_List	*dUF_Slot_Config_List;	/* OPTIONAL */
	HSNATransmissionPeriodicity_t	 hSNATransmissionPeriodicity;
	struct HSNASlotConfigList	*hsNSASlotConfigList;	/* OPTIONAL */
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} GNB_DU_Cell_Resource_Configuration_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_GNB_DU_Cell_Resource_Configuration;
extern asn_SEQUENCE_specifics_t asn_SPC_GNB_DU_Cell_Resource_Configuration_specs_1;
extern asn_TYPE_member_t asn_MBR_GNB_DU_Cell_Resource_Configuration_1[6];

#ifdef __cplusplus
}
#endif

#endif	/* _GNB_DU_Cell_Resource_Configuration_H_ */
#include "asn_internal.h"
