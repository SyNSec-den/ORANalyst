/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_QosMonitoringRequest_H_
#define	_QosMonitoringRequest_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum QosMonitoringRequest {
	QosMonitoringRequest_ul	= 0,
	QosMonitoringRequest_dl	= 1,
	QosMonitoringRequest_both	= 2,
	/*
	 * Enumeration is extensible
	 */
	QosMonitoringRequest_stop	= 3
} e_QosMonitoringRequest;

/* QosMonitoringRequest */
typedef long	 QosMonitoringRequest_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_QosMonitoringRequest_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_QosMonitoringRequest;
extern const asn_INTEGER_specifics_t asn_SPC_QosMonitoringRequest_specs_1;
asn_struct_free_f QosMonitoringRequest_free;
asn_struct_print_f QosMonitoringRequest_print;
asn_constr_check_f QosMonitoringRequest_constraint;
ber_type_decoder_f QosMonitoringRequest_decode_ber;
der_type_encoder_f QosMonitoringRequest_encode_der;
xer_type_decoder_f QosMonitoringRequest_decode_xer;
xer_type_encoder_f QosMonitoringRequest_encode_xer;
jer_type_encoder_f QosMonitoringRequest_encode_jer;
oer_type_decoder_f QosMonitoringRequest_decode_oer;
oer_type_encoder_f QosMonitoringRequest_encode_oer;
per_type_decoder_f QosMonitoringRequest_decode_uper;
per_type_encoder_f QosMonitoringRequest_encode_uper;
per_type_decoder_f QosMonitoringRequest_decode_aper;
per_type_encoder_f QosMonitoringRequest_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _QosMonitoringRequest_H_ */
#include "asn_internal.h"
