/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_ExecuteDuplication_H_
#define	_ExecuteDuplication_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ExecuteDuplication {
	ExecuteDuplication_true	= 0
	/*
	 * Enumeration is extensible
	 */
} e_ExecuteDuplication;

/* ExecuteDuplication */
typedef long	 ExecuteDuplication_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_ExecuteDuplication_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_ExecuteDuplication;
extern const asn_INTEGER_specifics_t asn_SPC_ExecuteDuplication_specs_1;
asn_struct_free_f ExecuteDuplication_free;
asn_struct_print_f ExecuteDuplication_print;
asn_constr_check_f ExecuteDuplication_constraint;
ber_type_decoder_f ExecuteDuplication_decode_ber;
der_type_encoder_f ExecuteDuplication_encode_der;
xer_type_decoder_f ExecuteDuplication_decode_xer;
xer_type_encoder_f ExecuteDuplication_encode_xer;
jer_type_encoder_f ExecuteDuplication_encode_jer;
oer_type_decoder_f ExecuteDuplication_decode_oer;
oer_type_encoder_f ExecuteDuplication_encode_oer;
per_type_decoder_f ExecuteDuplication_decode_uper;
per_type_encoder_f ExecuteDuplication_encode_uper;
per_type_decoder_f ExecuteDuplication_decode_aper;
per_type_encoder_f ExecuteDuplication_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _ExecuteDuplication_H_ */
#include "asn_internal.h"