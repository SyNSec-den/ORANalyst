/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_CG_SDTindicatorMod_H_
#define	_CG_SDTindicatorMod_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum CG_SDTindicatorMod {
	CG_SDTindicatorMod_true	= 0,
	CG_SDTindicatorMod_false	= 1
	/*
	 * Enumeration is extensible
	 */
} e_CG_SDTindicatorMod;

/* CG-SDTindicatorMod */
typedef long	 CG_SDTindicatorMod_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_CG_SDTindicatorMod_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_CG_SDTindicatorMod;
extern const asn_INTEGER_specifics_t asn_SPC_CG_SDTindicatorMod_specs_1;
asn_struct_free_f CG_SDTindicatorMod_free;
asn_struct_print_f CG_SDTindicatorMod_print;
asn_constr_check_f CG_SDTindicatorMod_constraint;
ber_type_decoder_f CG_SDTindicatorMod_decode_ber;
der_type_encoder_f CG_SDTindicatorMod_encode_der;
xer_type_decoder_f CG_SDTindicatorMod_decode_xer;
xer_type_encoder_f CG_SDTindicatorMod_encode_xer;
jer_type_encoder_f CG_SDTindicatorMod_encode_jer;
oer_type_decoder_f CG_SDTindicatorMod_decode_oer;
oer_type_encoder_f CG_SDTindicatorMod_encode_oer;
per_type_decoder_f CG_SDTindicatorMod_decode_uper;
per_type_encoder_f CG_SDTindicatorMod_encode_uper;
per_type_decoder_f CG_SDTindicatorMod_decode_aper;
per_type_encoder_f CG_SDTindicatorMod_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _CG_SDTindicatorMod_H_ */
#include "asn_internal.h"
