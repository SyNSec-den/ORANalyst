/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_PDCMeasurementQuantitiesValue_H_
#define	_PDCMeasurementQuantitiesValue_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum PDCMeasurementQuantitiesValue {
	PDCMeasurementQuantitiesValue_nr_pdc_tadv	= 0,
	PDCMeasurementQuantitiesValue_gNB_rx_tx	= 1
	/*
	 * Enumeration is extensible
	 */
} e_PDCMeasurementQuantitiesValue;

/* PDCMeasurementQuantitiesValue */
typedef long	 PDCMeasurementQuantitiesValue_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_PDCMeasurementQuantitiesValue_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_PDCMeasurementQuantitiesValue;
extern const asn_INTEGER_specifics_t asn_SPC_PDCMeasurementQuantitiesValue_specs_1;
asn_struct_free_f PDCMeasurementQuantitiesValue_free;
asn_struct_print_f PDCMeasurementQuantitiesValue_print;
asn_constr_check_f PDCMeasurementQuantitiesValue_constraint;
ber_type_decoder_f PDCMeasurementQuantitiesValue_decode_ber;
der_type_encoder_f PDCMeasurementQuantitiesValue_encode_der;
xer_type_decoder_f PDCMeasurementQuantitiesValue_decode_xer;
xer_type_encoder_f PDCMeasurementQuantitiesValue_encode_xer;
jer_type_encoder_f PDCMeasurementQuantitiesValue_encode_jer;
oer_type_decoder_f PDCMeasurementQuantitiesValue_decode_oer;
oer_type_encoder_f PDCMeasurementQuantitiesValue_encode_oer;
per_type_decoder_f PDCMeasurementQuantitiesValue_decode_uper;
per_type_encoder_f PDCMeasurementQuantitiesValue_encode_uper;
per_type_decoder_f PDCMeasurementQuantitiesValue_decode_aper;
per_type_encoder_f PDCMeasurementQuantitiesValue_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _PDCMeasurementQuantitiesValue_H_ */
#include "asn_internal.h"
