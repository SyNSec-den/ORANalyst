/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "PRS-Measurement-Info-List-Item.h"

#include "ProtocolExtensionContainer.h"
/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
static int
memb_pointA_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const long *)sptr;
	
	if((value >= 0L && value <= 3279165L)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int
memb_measPRSOffset_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const long *)sptr;
	
	if((value >= 0L && value <= 159L)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

#if !defined(ASN_DISABLE_OER_SUPPORT)
static asn_oer_constraints_t asn_OER_type_measPRSPeriodicity_constr_3 CC_NOTUSED = {
	{ 0, 0 },
	-1};
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
static asn_per_constraints_t asn_PER_type_measPRSPeriodicity_constr_3 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  2,  2,  0,  3 }	/* (0..3,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
#if !defined(ASN_DISABLE_OER_SUPPORT)
static asn_oer_constraints_t asn_OER_type_measurementPRSLength_constr_10 CC_NOTUSED = {
	{ 0, 0 },
	-1};
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
static asn_per_constraints_t asn_PER_type_measurementPRSLength_constr_10 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 3,  3,  0,  7 }	/* (0..7) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
#if !defined(ASN_DISABLE_OER_SUPPORT)
static asn_oer_constraints_t asn_OER_memb_pointA_constr_2 CC_NOTUSED = {
	{ 4, 1 }	/* (0..3279165) */,
	-1};
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
static asn_per_constraints_t asn_PER_memb_pointA_constr_2 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 22, -1,  0,  3279165 }	/* (0..3279165) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
#if !defined(ASN_DISABLE_OER_SUPPORT)
static asn_oer_constraints_t asn_OER_memb_measPRSOffset_constr_9 CC_NOTUSED = {
	{ 0, 0 },
	-1};
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
static asn_per_constraints_t asn_PER_memb_measPRSOffset_constr_9 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  8,  8,  0,  159 }	/* (0..159,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
static const asn_INTEGER_enum_map_t asn_MAP_measPRSPeriodicity_value2enum_3[] = {
	{ 0,	4,	"ms20" },
	{ 1,	4,	"ms40" },
	{ 2,	4,	"ms80" },
	{ 3,	5,	"ms160" }
	/* This list is extensible */
};
static const unsigned int asn_MAP_measPRSPeriodicity_enum2value_3[] = {
	3,	/* ms160(3) */
	0,	/* ms20(0) */
	1,	/* ms40(1) */
	2	/* ms80(2) */
	/* This list is extensible */
};
static const asn_INTEGER_specifics_t asn_SPC_measPRSPeriodicity_specs_3 = {
	asn_MAP_measPRSPeriodicity_value2enum_3,	/* "tag" => N; sorted by tag */
	asn_MAP_measPRSPeriodicity_enum2value_3,	/* N => "tag"; sorted by N */
	4,	/* Number of elements in the maps */
	5,	/* Extensions before this member */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_measPRSPeriodicity_tags_3[] = {
	(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_measPRSPeriodicity_3 = {
	"measPRSPeriodicity",
	"measPRSPeriodicity",
	&asn_OP_NativeEnumerated,
	asn_DEF_measPRSPeriodicity_tags_3,
	sizeof(asn_DEF_measPRSPeriodicity_tags_3)
		/sizeof(asn_DEF_measPRSPeriodicity_tags_3[0]) - 1, /* 1 */
	asn_DEF_measPRSPeriodicity_tags_3,	/* Same as above */
	sizeof(asn_DEF_measPRSPeriodicity_tags_3)
		/sizeof(asn_DEF_measPRSPeriodicity_tags_3[0]), /* 2 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		&asn_OER_type_measPRSPeriodicity_constr_3,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		&asn_PER_type_measPRSPeriodicity_constr_3,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		NativeEnumerated_constraint
	},
	0, 0,	/* Defined elsewhere */
	&asn_SPC_measPRSPeriodicity_specs_3	/* Additional specs */
};

static const asn_INTEGER_enum_map_t asn_MAP_measurementPRSLength_value2enum_10[] = {
	{ 0,	7,	"ms1dot5" },
	{ 1,	3,	"ms3" },
	{ 2,	7,	"ms3dot5" },
	{ 3,	3,	"ms4" },
	{ 4,	7,	"ms5dot5" },
	{ 5,	3,	"ms6" },
	{ 6,	4,	"ms10" },
	{ 7,	4,	"ms20" }
};
static const unsigned int asn_MAP_measurementPRSLength_enum2value_10[] = {
	6,	/* ms10(6) */
	0,	/* ms1dot5(0) */
	7,	/* ms20(7) */
	1,	/* ms3(1) */
	2,	/* ms3dot5(2) */
	3,	/* ms4(3) */
	4,	/* ms5dot5(4) */
	5	/* ms6(5) */
};
static const asn_INTEGER_specifics_t asn_SPC_measurementPRSLength_specs_10 = {
	asn_MAP_measurementPRSLength_value2enum_10,	/* "tag" => N; sorted by tag */
	asn_MAP_measurementPRSLength_enum2value_10,	/* N => "tag"; sorted by N */
	8,	/* Number of elements in the maps */
	0,	/* Enumeration is not extensible */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_measurementPRSLength_tags_10[] = {
	(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_measurementPRSLength_10 = {
	"measurementPRSLength",
	"measurementPRSLength",
	&asn_OP_NativeEnumerated,
	asn_DEF_measurementPRSLength_tags_10,
	sizeof(asn_DEF_measurementPRSLength_tags_10)
		/sizeof(asn_DEF_measurementPRSLength_tags_10[0]) - 1, /* 1 */
	asn_DEF_measurementPRSLength_tags_10,	/* Same as above */
	sizeof(asn_DEF_measurementPRSLength_tags_10)
		/sizeof(asn_DEF_measurementPRSLength_tags_10[0]), /* 2 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		&asn_OER_type_measurementPRSLength_constr_10,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		&asn_PER_type_measurementPRSLength_constr_10,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		NativeEnumerated_constraint
	},
	0, 0,	/* Defined elsewhere */
	&asn_SPC_measurementPRSLength_specs_10	/* Additional specs */
};

asn_TYPE_member_t asn_MBR_PRS_Measurement_Info_List_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct PRS_Measurement_Info_List_Item, pointA),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{
#if !defined(ASN_DISABLE_OER_SUPPORT)
			&asn_OER_memb_pointA_constr_2,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
			&asn_PER_memb_pointA_constr_2,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
			memb_pointA_constraint_1
		},
		0, 0, /* No default value */
		"pointA"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct PRS_Measurement_Info_List_Item, measPRSPeriodicity),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_measPRSPeriodicity_3,
		0,
		{
#if !defined(ASN_DISABLE_OER_SUPPORT)
			0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
			0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
			0
		},
		0, 0, /* No default value */
		"measPRSPeriodicity"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct PRS_Measurement_Info_List_Item, measPRSOffset),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{
#if !defined(ASN_DISABLE_OER_SUPPORT)
			&asn_OER_memb_measPRSOffset_constr_9,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
			&asn_PER_memb_measPRSOffset_constr_9,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
			memb_measPRSOffset_constraint_1
		},
		0, 0, /* No default value */
		"measPRSOffset"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct PRS_Measurement_Info_List_Item, measurementPRSLength),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_measurementPRSLength_10,
		0,
		{
#if !defined(ASN_DISABLE_OER_SUPPORT)
			0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
			0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
			0
		},
		0, 0, /* No default value */
		"measurementPRSLength"
		},
	{ ATF_POINTER, 1, offsetof(struct PRS_Measurement_Info_List_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P290,
		0,
		{
#if !defined(ASN_DISABLE_OER_SUPPORT)
			0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
			0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
			0
		},
		0, 0, /* No default value */
		"iE-Extensions"
		},
};
static const int asn_MAP_PRS_Measurement_Info_List_Item_oms_1[] = { 4 };
static const ber_tlv_tag_t asn_DEF_PRS_Measurement_Info_List_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_PRS_Measurement_Info_List_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* pointA */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* measPRSPeriodicity */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* measPRSOffset */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* measurementPRSLength */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_PRS_Measurement_Info_List_Item_specs_1 = {
	sizeof(struct PRS_Measurement_Info_List_Item),
	offsetof(struct PRS_Measurement_Info_List_Item, _asn_ctx),
	asn_MAP_PRS_Measurement_Info_List_Item_tag2el_1,
	5,	/* Count of tags in the map */
	asn_MAP_PRS_Measurement_Info_List_Item_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	5,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_PRS_Measurement_Info_List_Item = {
	"PRS-Measurement-Info-List-Item",
	"PRS-Measurement-Info-List-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_PRS_Measurement_Info_List_Item_tags_1,
	sizeof(asn_DEF_PRS_Measurement_Info_List_Item_tags_1)
		/sizeof(asn_DEF_PRS_Measurement_Info_List_Item_tags_1[0]), /* 1 */
	asn_DEF_PRS_Measurement_Info_List_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_PRS_Measurement_Info_List_Item_tags_1)
		/sizeof(asn_DEF_PRS_Measurement_Info_List_Item_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_PRS_Measurement_Info_List_Item_1,
	5,	/* Elements count */
	&asn_SPC_PRS_Measurement_Info_List_Item_specs_1	/* Additional specs */
};
