/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "RxTxTimingErrorMargin.h"

/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
#if !defined(ASN_DISABLE_OER_SUPPORT)
static asn_oer_constraints_t asn_OER_type_RxTxTimingErrorMargin_constr_1 CC_NOTUSED = {
	{ 0, 0 },
	-1};
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
asn_per_constraints_t asn_PER_type_RxTxTimingErrorMargin_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  4,  4,  0,  15 }	/* (0..15,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
static const asn_INTEGER_enum_map_t asn_MAP_RxTxTimingErrorMargin_value2enum_1[] = {
	{ 0,	7,	"tc0dot5" },
	{ 1,	3,	"tc1" },
	{ 2,	3,	"tc2" },
	{ 3,	3,	"tc4" },
	{ 4,	3,	"tc8" },
	{ 5,	4,	"tc12" },
	{ 6,	4,	"tc16" },
	{ 7,	4,	"tc20" },
	{ 8,	4,	"tc24" },
	{ 9,	4,	"tc32" },
	{ 10,	4,	"tc40" },
	{ 11,	4,	"tc48" },
	{ 12,	4,	"tc64" },
	{ 13,	4,	"tc80" },
	{ 14,	4,	"tc96" },
	{ 15,	5,	"tc128" }
	/* This list is extensible */
};
static const unsigned int asn_MAP_RxTxTimingErrorMargin_enum2value_1[] = {
	0,	/* tc0dot5(0) */
	1,	/* tc1(1) */
	5,	/* tc12(5) */
	15,	/* tc128(15) */
	6,	/* tc16(6) */
	2,	/* tc2(2) */
	7,	/* tc20(7) */
	8,	/* tc24(8) */
	9,	/* tc32(9) */
	3,	/* tc4(3) */
	10,	/* tc40(10) */
	11,	/* tc48(11) */
	12,	/* tc64(12) */
	4,	/* tc8(4) */
	13,	/* tc80(13) */
	14	/* tc96(14) */
	/* This list is extensible */
};
const asn_INTEGER_specifics_t asn_SPC_RxTxTimingErrorMargin_specs_1 = {
	asn_MAP_RxTxTimingErrorMargin_value2enum_1,	/* "tag" => N; sorted by tag */
	asn_MAP_RxTxTimingErrorMargin_enum2value_1,	/* N => "tag"; sorted by N */
	16,	/* Number of elements in the maps */
	17,	/* Extensions before this member */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_RxTxTimingErrorMargin_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
asn_TYPE_descriptor_t asn_DEF_RxTxTimingErrorMargin = {
	"RxTxTimingErrorMargin",
	"RxTxTimingErrorMargin",
	&asn_OP_NativeEnumerated,
	asn_DEF_RxTxTimingErrorMargin_tags_1,
	sizeof(asn_DEF_RxTxTimingErrorMargin_tags_1)
		/sizeof(asn_DEF_RxTxTimingErrorMargin_tags_1[0]), /* 1 */
	asn_DEF_RxTxTimingErrorMargin_tags_1,	/* Same as above */
	sizeof(asn_DEF_RxTxTimingErrorMargin_tags_1)
		/sizeof(asn_DEF_RxTxTimingErrorMargin_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		&asn_OER_type_RxTxTimingErrorMargin_constr_1,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		&asn_PER_type_RxTxTimingErrorMargin_constr_1,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		NativeEnumerated_constraint
	},
	0, 0,	/* Defined elsewhere */
	&asn_SPC_RxTxTimingErrorMargin_specs_1	/* Additional specs */
};
