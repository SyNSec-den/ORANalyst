/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_FR1_Bandwidth_H_
#define	_FR1_Bandwidth_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum FR1_Bandwidth {
	FR1_Bandwidth_bw5	= 0,
	FR1_Bandwidth_bw10	= 1,
	FR1_Bandwidth_bw20	= 2,
	FR1_Bandwidth_bw40	= 3,
	FR1_Bandwidth_bw50	= 4,
	FR1_Bandwidth_bw80	= 5,
	FR1_Bandwidth_bw100	= 6
	/*
	 * Enumeration is extensible
	 */
} e_FR1_Bandwidth;

/* FR1-Bandwidth */
typedef long	 FR1_Bandwidth_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_FR1_Bandwidth_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_FR1_Bandwidth;
extern const asn_INTEGER_specifics_t asn_SPC_FR1_Bandwidth_specs_1;
asn_struct_free_f FR1_Bandwidth_free;
asn_struct_print_f FR1_Bandwidth_print;
asn_constr_check_f FR1_Bandwidth_constraint;
ber_type_decoder_f FR1_Bandwidth_decode_ber;
der_type_encoder_f FR1_Bandwidth_encode_der;
xer_type_decoder_f FR1_Bandwidth_decode_xer;
xer_type_encoder_f FR1_Bandwidth_encode_xer;
jer_type_encoder_f FR1_Bandwidth_encode_jer;
oer_type_decoder_f FR1_Bandwidth_decode_oer;
oer_type_encoder_f FR1_Bandwidth_encode_oer;
per_type_decoder_f FR1_Bandwidth_decode_uper;
per_type_encoder_f FR1_Bandwidth_encode_uper;
per_type_decoder_f FR1_Bandwidth_decode_aper;
per_type_encoder_f FR1_Bandwidth_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _FR1_Bandwidth_H_ */
#include "asn_internal.h"