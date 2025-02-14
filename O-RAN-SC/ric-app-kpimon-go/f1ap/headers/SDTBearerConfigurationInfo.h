/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_SDTBearerConfigurationInfo_H_
#define	_SDTBearerConfigurationInfo_H_


#include "asn_application.h"

/* Including external dependencies */
#include "SDTBearerConfig-List.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolExtensionContainer;

/* SDTBearerConfigurationInfo */
typedef struct SDTBearerConfigurationInfo {
	SDTBearerConfig_List_t	 sDTBearerConfig_List;
	struct ProtocolExtensionContainer	*iE_Extensions;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} SDTBearerConfigurationInfo_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_SDTBearerConfigurationInfo;
extern asn_SEQUENCE_specifics_t asn_SPC_SDTBearerConfigurationInfo_specs_1;
extern asn_TYPE_member_t asn_MBR_SDTBearerConfigurationInfo_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _SDTBearerConfigurationInfo_H_ */
#include "asn_internal.h"
