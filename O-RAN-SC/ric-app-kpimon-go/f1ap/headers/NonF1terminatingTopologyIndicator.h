/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_NonF1terminatingTopologyIndicator_H_
#define	_NonF1terminatingTopologyIndicator_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum NonF1terminatingTopologyIndicator {
	NonF1terminatingTopologyIndicator_true	= 0
	/*
	 * Enumeration is extensible
	 */
} e_NonF1terminatingTopologyIndicator;

/* NonF1terminatingTopologyIndicator */
typedef long	 NonF1terminatingTopologyIndicator_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_NonF1terminatingTopologyIndicator_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_NonF1terminatingTopologyIndicator;
extern const asn_INTEGER_specifics_t asn_SPC_NonF1terminatingTopologyIndicator_specs_1;
asn_struct_free_f NonF1terminatingTopologyIndicator_free;
asn_struct_print_f NonF1terminatingTopologyIndicator_print;
asn_constr_check_f NonF1terminatingTopologyIndicator_constraint;
ber_type_decoder_f NonF1terminatingTopologyIndicator_decode_ber;
der_type_encoder_f NonF1terminatingTopologyIndicator_encode_der;
xer_type_decoder_f NonF1terminatingTopologyIndicator_decode_xer;
xer_type_encoder_f NonF1terminatingTopologyIndicator_encode_xer;
jer_type_encoder_f NonF1terminatingTopologyIndicator_encode_jer;
oer_type_decoder_f NonF1terminatingTopologyIndicator_decode_oer;
oer_type_encoder_f NonF1terminatingTopologyIndicator_encode_oer;
per_type_decoder_f NonF1terminatingTopologyIndicator_decode_uper;
per_type_encoder_f NonF1terminatingTopologyIndicator_encode_uper;
per_type_decoder_f NonF1terminatingTopologyIndicator_decode_aper;
per_type_encoder_f NonF1terminatingTopologyIndicator_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _NonF1terminatingTopologyIndicator_H_ */
#include "asn_internal.h"
