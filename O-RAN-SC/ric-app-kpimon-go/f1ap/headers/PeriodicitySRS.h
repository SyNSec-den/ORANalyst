/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_PeriodicitySRS_H_
#define	_PeriodicitySRS_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum PeriodicitySRS {
	PeriodicitySRS_ms0p125	= 0,
	PeriodicitySRS_ms0p25	= 1,
	PeriodicitySRS_ms0p5	= 2,
	PeriodicitySRS_ms0p625	= 3,
	PeriodicitySRS_ms1	= 4,
	PeriodicitySRS_ms1p25	= 5,
	PeriodicitySRS_ms2	= 6,
	PeriodicitySRS_ms2p5	= 7,
	PeriodicitySRS_ms4	= 8,
	PeriodicitySRS_ms5	= 9,
	PeriodicitySRS_ms8	= 10,
	PeriodicitySRS_ms10	= 11,
	PeriodicitySRS_ms16	= 12,
	PeriodicitySRS_ms20	= 13,
	PeriodicitySRS_ms32	= 14,
	PeriodicitySRS_ms40	= 15,
	PeriodicitySRS_ms64	= 16,
	PeriodicitySRS_ms80	= 17,
	PeriodicitySRS_ms160	= 18,
	PeriodicitySRS_ms320	= 19,
	PeriodicitySRS_ms640	= 20,
	PeriodicitySRS_ms1280	= 21,
	PeriodicitySRS_ms2560	= 22,
	PeriodicitySRS_ms5120	= 23,
	PeriodicitySRS_ms10240	= 24
	/*
	 * Enumeration is extensible
	 */
} e_PeriodicitySRS;

/* PeriodicitySRS */
typedef long	 PeriodicitySRS_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_PeriodicitySRS_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_PeriodicitySRS;
extern const asn_INTEGER_specifics_t asn_SPC_PeriodicitySRS_specs_1;
asn_struct_free_f PeriodicitySRS_free;
asn_struct_print_f PeriodicitySRS_print;
asn_constr_check_f PeriodicitySRS_constraint;
ber_type_decoder_f PeriodicitySRS_decode_ber;
der_type_encoder_f PeriodicitySRS_encode_der;
xer_type_decoder_f PeriodicitySRS_decode_xer;
xer_type_encoder_f PeriodicitySRS_encode_xer;
jer_type_encoder_f PeriodicitySRS_encode_jer;
oer_type_decoder_f PeriodicitySRS_decode_oer;
oer_type_encoder_f PeriodicitySRS_encode_oer;
per_type_decoder_f PeriodicitySRS_decode_uper;
per_type_encoder_f PeriodicitySRS_encode_uper;
per_type_decoder_f PeriodicitySRS_decode_aper;
per_type_encoder_f PeriodicitySRS_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _PeriodicitySRS_H_ */
#include "asn_internal.h"