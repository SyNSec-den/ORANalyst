/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "UE-MulticastMRBs-RequiredToBeModified-Item.h"

#include "ProtocolExtensionContainer.h"
/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
#if !defined(ASN_DISABLE_OER_SUPPORT)
static asn_oer_constraints_t asn_OER_type_mrb_type_reconfiguration_constr_3 CC_NOTUSED = {
	{ 0, 0 },
	-1};
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
static asn_per_constraints_t asn_PER_type_mrb_type_reconfiguration_constr_3 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  0,  0,  0,  0 }	/* (0..0,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
#if !defined(ASN_DISABLE_OER_SUPPORT)
static asn_oer_constraints_t asn_OER_type_mrb_reconfigured_RLCtype_constr_6 CC_NOTUSED = {
	{ 0, 0 },
	-1};
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
static asn_per_constraints_t asn_PER_type_mrb_reconfigured_RLCtype_constr_6 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  3,  3,  0,  5 }	/* (0..5,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
static const asn_INTEGER_enum_map_t asn_MAP_mrb_type_reconfiguration_value2enum_3[] = {
	{ 0,	4,	"true" }
	/* This list is extensible */
};
static const unsigned int asn_MAP_mrb_type_reconfiguration_enum2value_3[] = {
	0	/* true(0) */
	/* This list is extensible */
};
static const asn_INTEGER_specifics_t asn_SPC_mrb_type_reconfiguration_specs_3 = {
	asn_MAP_mrb_type_reconfiguration_value2enum_3,	/* "tag" => N; sorted by tag */
	asn_MAP_mrb_type_reconfiguration_enum2value_3,	/* N => "tag"; sorted by N */
	1,	/* Number of elements in the maps */
	2,	/* Extensions before this member */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_mrb_type_reconfiguration_tags_3[] = {
	(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_mrb_type_reconfiguration_3 = {
	"mrb-type-reconfiguration",
	"mrb-type-reconfiguration",
	&asn_OP_NativeEnumerated,
	asn_DEF_mrb_type_reconfiguration_tags_3,
	sizeof(asn_DEF_mrb_type_reconfiguration_tags_3)
		/sizeof(asn_DEF_mrb_type_reconfiguration_tags_3[0]) - 1, /* 1 */
	asn_DEF_mrb_type_reconfiguration_tags_3,	/* Same as above */
	sizeof(asn_DEF_mrb_type_reconfiguration_tags_3)
		/sizeof(asn_DEF_mrb_type_reconfiguration_tags_3[0]), /* 2 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		&asn_OER_type_mrb_type_reconfiguration_constr_3,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		&asn_PER_type_mrb_type_reconfiguration_constr_3,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		NativeEnumerated_constraint
	},
	0, 0,	/* Defined elsewhere */
	&asn_SPC_mrb_type_reconfiguration_specs_3	/* Additional specs */
};

static const asn_INTEGER_enum_map_t asn_MAP_mrb_reconfigured_RLCtype_value2enum_6[] = {
	{ 0,	10,	"rlc-um-ptp" },
	{ 1,	10,	"rlc-am-ptp" },
	{ 2,	13,	"rlc-um-dl-ptm" },
	{ 3,	28,	"two-rlc-um-dl-ptp-and-dl-ptm" },
	{ 4,	33,	"three-rlc-um-dl-ptp-ul-ptp-dl-ptm" },
	{ 5,	24,	"two-rlc-am-ptp-um-dl-ptm" }
	/* This list is extensible */
};
static const unsigned int asn_MAP_mrb_reconfigured_RLCtype_enum2value_6[] = {
	1,	/* rlc-am-ptp(1) */
	2,	/* rlc-um-dl-ptm(2) */
	0,	/* rlc-um-ptp(0) */
	4,	/* three-rlc-um-dl-ptp-ul-ptp-dl-ptm(4) */
	5,	/* two-rlc-am-ptp-um-dl-ptm(5) */
	3	/* two-rlc-um-dl-ptp-and-dl-ptm(3) */
	/* This list is extensible */
};
static const asn_INTEGER_specifics_t asn_SPC_mrb_reconfigured_RLCtype_specs_6 = {
	asn_MAP_mrb_reconfigured_RLCtype_value2enum_6,	/* "tag" => N; sorted by tag */
	asn_MAP_mrb_reconfigured_RLCtype_enum2value_6,	/* N => "tag"; sorted by N */
	6,	/* Number of elements in the maps */
	7,	/* Extensions before this member */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_mrb_reconfigured_RLCtype_tags_6[] = {
	(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_mrb_reconfigured_RLCtype_6 = {
	"mrb-reconfigured-RLCtype",
	"mrb-reconfigured-RLCtype",
	&asn_OP_NativeEnumerated,
	asn_DEF_mrb_reconfigured_RLCtype_tags_6,
	sizeof(asn_DEF_mrb_reconfigured_RLCtype_tags_6)
		/sizeof(asn_DEF_mrb_reconfigured_RLCtype_tags_6[0]) - 1, /* 1 */
	asn_DEF_mrb_reconfigured_RLCtype_tags_6,	/* Same as above */
	sizeof(asn_DEF_mrb_reconfigured_RLCtype_tags_6)
		/sizeof(asn_DEF_mrb_reconfigured_RLCtype_tags_6[0]), /* 2 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		&asn_OER_type_mrb_reconfigured_RLCtype_constr_6,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		&asn_PER_type_mrb_reconfigured_RLCtype_constr_6,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		NativeEnumerated_constraint
	},
	0, 0,	/* Defined elsewhere */
	&asn_SPC_mrb_reconfigured_RLCtype_specs_6	/* Additional specs */
};

asn_TYPE_member_t asn_MBR_UE_MulticastMRBs_RequiredToBeModified_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct UE_MulticastMRBs_RequiredToBeModified_Item, mRB_ID),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_MRB_ID,
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
		"mRB-ID"
		},
	{ ATF_POINTER, 3, offsetof(struct UE_MulticastMRBs_RequiredToBeModified_Item, mrb_type_reconfiguration),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_mrb_type_reconfiguration_3,
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
		"mrb-type-reconfiguration"
		},
	{ ATF_POINTER, 2, offsetof(struct UE_MulticastMRBs_RequiredToBeModified_Item, mrb_reconfigured_RLCtype),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_mrb_reconfigured_RLCtype_6,
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
		"mrb-reconfigured-RLCtype"
		},
	{ ATF_POINTER, 1, offsetof(struct UE_MulticastMRBs_RequiredToBeModified_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P469,
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
static const int asn_MAP_UE_MulticastMRBs_RequiredToBeModified_Item_oms_1[] = { 1, 2, 3 };
static const ber_tlv_tag_t asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_UE_MulticastMRBs_RequiredToBeModified_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* mRB-ID */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* mrb-type-reconfiguration */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* mrb-reconfigured-RLCtype */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_UE_MulticastMRBs_RequiredToBeModified_Item_specs_1 = {
	sizeof(struct UE_MulticastMRBs_RequiredToBeModified_Item),
	offsetof(struct UE_MulticastMRBs_RequiredToBeModified_Item, _asn_ctx),
	asn_MAP_UE_MulticastMRBs_RequiredToBeModified_Item_tag2el_1,
	4,	/* Count of tags in the map */
	asn_MAP_UE_MulticastMRBs_RequiredToBeModified_Item_oms_1,	/* Optional members */
	3, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item = {
	"UE-MulticastMRBs-RequiredToBeModified-Item",
	"UE-MulticastMRBs-RequiredToBeModified-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item_tags_1,
	sizeof(asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item_tags_1)
		/sizeof(asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item_tags_1[0]), /* 1 */
	asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item_tags_1)
		/sizeof(asn_DEF_UE_MulticastMRBs_RequiredToBeModified_Item_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_UE_MulticastMRBs_RequiredToBeModified_Item_1,
	4,	/* Elements count */
	&asn_SPC_UE_MulticastMRBs_RequiredToBeModified_Item_specs_1	/* Additional specs */
};
