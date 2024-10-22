/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_EUTRA_SubframeAssignment_H_
#define	_EUTRA_SubframeAssignment_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum EUTRA_SubframeAssignment {
	EUTRA_SubframeAssignment_sa0	= 0,
	EUTRA_SubframeAssignment_sa1	= 1,
	EUTRA_SubframeAssignment_sa2	= 2,
	EUTRA_SubframeAssignment_sa3	= 3,
	EUTRA_SubframeAssignment_sa4	= 4,
	EUTRA_SubframeAssignment_sa5	= 5,
	EUTRA_SubframeAssignment_sa6	= 6
	/*
	 * Enumeration is extensible
	 */
} e_EUTRA_SubframeAssignment;

/* EUTRA-SubframeAssignment */
typedef long	 EUTRA_SubframeAssignment_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_EUTRA_SubframeAssignment_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_EUTRA_SubframeAssignment;
extern const asn_INTEGER_specifics_t asn_SPC_EUTRA_SubframeAssignment_specs_1;
asn_struct_free_f EUTRA_SubframeAssignment_free;
asn_struct_print_f EUTRA_SubframeAssignment_print;
asn_constr_check_f EUTRA_SubframeAssignment_constraint;
ber_type_decoder_f EUTRA_SubframeAssignment_decode_ber;
der_type_encoder_f EUTRA_SubframeAssignment_encode_der;
xer_type_decoder_f EUTRA_SubframeAssignment_decode_xer;
xer_type_encoder_f EUTRA_SubframeAssignment_encode_xer;
jer_type_encoder_f EUTRA_SubframeAssignment_encode_jer;
oer_type_decoder_f EUTRA_SubframeAssignment_decode_oer;
oer_type_encoder_f EUTRA_SubframeAssignment_encode_oer;
per_type_decoder_f EUTRA_SubframeAssignment_decode_uper;
per_type_encoder_f EUTRA_SubframeAssignment_encode_uper;
per_type_decoder_f EUTRA_SubframeAssignment_decode_aper;
per_type_encoder_f EUTRA_SubframeAssignment_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _EUTRA_SubframeAssignment_H_ */
#include "asn_internal.h"
