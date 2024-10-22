/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_ActivationRequestType_H_
#define	_ActivationRequestType_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ActivationRequestType {
	ActivationRequestType_activate	= 0,
	ActivationRequestType_deactivate	= 1
	/*
	 * Enumeration is extensible
	 */
} e_ActivationRequestType;

/* ActivationRequestType */
typedef long	 ActivationRequestType_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_ActivationRequestType_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_ActivationRequestType;
extern const asn_INTEGER_specifics_t asn_SPC_ActivationRequestType_specs_1;
asn_struct_free_f ActivationRequestType_free;
asn_struct_print_f ActivationRequestType_print;
asn_constr_check_f ActivationRequestType_constraint;
ber_type_decoder_f ActivationRequestType_decode_ber;
der_type_encoder_f ActivationRequestType_encode_der;
xer_type_decoder_f ActivationRequestType_decode_xer;
xer_type_encoder_f ActivationRequestType_encode_xer;
jer_type_encoder_f ActivationRequestType_encode_jer;
oer_type_decoder_f ActivationRequestType_decode_oer;
oer_type_encoder_f ActivationRequestType_encode_oer;
per_type_decoder_f ActivationRequestType_decode_uper;
per_type_encoder_f ActivationRequestType_encode_uper;
per_type_decoder_f ActivationRequestType_decode_aper;
per_type_encoder_f ActivationRequestType_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _ActivationRequestType_H_ */
#include "asn_internal.h"
