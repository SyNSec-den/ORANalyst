/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_RACH_Config_Common_IAB_H_
#define	_RACH_Config_Common_IAB_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"

#ifdef __cplusplus
extern "C" {
#endif

/* RACH-Config-Common-IAB */
typedef OCTET_STRING_t	 RACH_Config_Common_IAB_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RACH_Config_Common_IAB;
asn_struct_free_f RACH_Config_Common_IAB_free;
asn_struct_print_f RACH_Config_Common_IAB_print;
asn_constr_check_f RACH_Config_Common_IAB_constraint;
ber_type_decoder_f RACH_Config_Common_IAB_decode_ber;
der_type_encoder_f RACH_Config_Common_IAB_encode_der;
xer_type_decoder_f RACH_Config_Common_IAB_decode_xer;
xer_type_encoder_f RACH_Config_Common_IAB_encode_xer;
jer_type_encoder_f RACH_Config_Common_IAB_encode_jer;
oer_type_decoder_f RACH_Config_Common_IAB_decode_oer;
oer_type_encoder_f RACH_Config_Common_IAB_encode_oer;
per_type_decoder_f RACH_Config_Common_IAB_decode_uper;
per_type_encoder_f RACH_Config_Common_IAB_encode_uper;
per_type_decoder_f RACH_Config_Common_IAB_decode_aper;
per_type_encoder_f RACH_Config_Common_IAB_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RACH_Config_Common_IAB_H_ */
#include "asn_internal.h"