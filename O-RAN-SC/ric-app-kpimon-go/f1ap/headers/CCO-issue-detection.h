/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#ifndef	_CCO_issue_detection_H_
#define	_CCO_issue_detection_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum CCO_issue_detection {
	CCO_issue_detection_coverage	= 0,
	CCO_issue_detection_cell_edge_capacity	= 1
	/*
	 * Enumeration is extensible
	 */
} e_CCO_issue_detection;

/* CCO-issue-detection */
typedef long	 CCO_issue_detection_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_CCO_issue_detection_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_CCO_issue_detection;
extern const asn_INTEGER_specifics_t asn_SPC_CCO_issue_detection_specs_1;
asn_struct_free_f CCO_issue_detection_free;
asn_struct_print_f CCO_issue_detection_print;
asn_constr_check_f CCO_issue_detection_constraint;
ber_type_decoder_f CCO_issue_detection_decode_ber;
der_type_encoder_f CCO_issue_detection_encode_der;
xer_type_decoder_f CCO_issue_detection_decode_xer;
xer_type_encoder_f CCO_issue_detection_encode_xer;
jer_type_encoder_f CCO_issue_detection_encode_jer;
oer_type_decoder_f CCO_issue_detection_decode_oer;
oer_type_encoder_f CCO_issue_detection_encode_oer;
per_type_decoder_f CCO_issue_detection_decode_uper;
per_type_encoder_f CCO_issue_detection_encode_uper;
per_type_decoder_f CCO_issue_detection_decode_aper;
per_type_encoder_f CCO_issue_detection_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _CCO_issue_detection_H_ */
#include "asn_internal.h"
