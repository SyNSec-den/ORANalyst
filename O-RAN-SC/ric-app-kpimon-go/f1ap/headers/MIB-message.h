/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_MIB_message_H_
#define	_MIB_message_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"

#ifdef __cplusplus
extern "C" {
#endif

/* MIB-message */
typedef OCTET_STRING_t	 MIB_message_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_MIB_message;
asn_struct_free_f MIB_message_free;
asn_struct_print_f MIB_message_print;
asn_constr_check_f MIB_message_constraint;
ber_type_decoder_f MIB_message_decode_ber;
der_type_encoder_f MIB_message_encode_der;
xer_type_decoder_f MIB_message_decode_xer;
xer_type_encoder_f MIB_message_encode_xer;
jer_type_encoder_f MIB_message_encode_jer;
oer_type_decoder_f MIB_message_decode_oer;
oer_type_encoder_f MIB_message_encode_oer;
per_type_decoder_f MIB_message_decode_uper;
per_type_encoder_f MIB_message_encode_uper;
per_type_decoder_f MIB_message_decode_aper;
per_type_encoder_f MIB_message_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _MIB_message_H_ */
#include "asn_internal.h"
