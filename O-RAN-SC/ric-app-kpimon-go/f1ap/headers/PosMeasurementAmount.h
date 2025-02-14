/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_PosMeasurementAmount_H_
#define	_PosMeasurementAmount_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum PosMeasurementAmount {
	PosMeasurementAmount_ma0	= 0,
	PosMeasurementAmount_ma1	= 1,
	PosMeasurementAmount_ma2	= 2,
	PosMeasurementAmount_ma4	= 3,
	PosMeasurementAmount_ma8	= 4,
	PosMeasurementAmount_ma16	= 5,
	PosMeasurementAmount_ma32	= 6,
	PosMeasurementAmount_ma64	= 7
} e_PosMeasurementAmount;

/* PosMeasurementAmount */
typedef long	 PosMeasurementAmount_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_PosMeasurementAmount_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_PosMeasurementAmount;
extern const asn_INTEGER_specifics_t asn_SPC_PosMeasurementAmount_specs_1;
asn_struct_free_f PosMeasurementAmount_free;
asn_struct_print_f PosMeasurementAmount_print;
asn_constr_check_f PosMeasurementAmount_constraint;
ber_type_decoder_f PosMeasurementAmount_decode_ber;
der_type_encoder_f PosMeasurementAmount_encode_der;
xer_type_decoder_f PosMeasurementAmount_decode_xer;
xer_type_encoder_f PosMeasurementAmount_encode_xer;
jer_type_encoder_f PosMeasurementAmount_encode_jer;
oer_type_decoder_f PosMeasurementAmount_decode_oer;
oer_type_encoder_f PosMeasurementAmount_encode_oer;
per_type_decoder_f PosMeasurementAmount_decode_uper;
per_type_encoder_f PosMeasurementAmount_encode_uper;
per_type_decoder_f PosMeasurementAmount_decode_aper;
per_type_encoder_f PosMeasurementAmount_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _PosMeasurementAmount_H_ */
#include "asn_internal.h"
