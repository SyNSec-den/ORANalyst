/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "LocationDependentMBSF1UInformation-Item.h"

#include "ProtocolExtensionContainer.h"
asn_TYPE_member_t asn_MBR_LocationDependentMBSF1UInformation_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct LocationDependentMBSF1UInformation_Item, mbsAreaSession_ID),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_MBS_Area_Session_ID,
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
		"mbsAreaSession-ID"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct LocationDependentMBSF1UInformation_Item, mbs_f1u_info_at_CU),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_UPTransportLayerInformation,
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
		"mbs-f1u-info-at-CU"
		},
	{ ATF_POINTER, 1, offsetof(struct LocationDependentMBSF1UInformation_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P192,
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
static const int asn_MAP_LocationDependentMBSF1UInformation_Item_oms_1[] = { 2 };
static const ber_tlv_tag_t asn_DEF_LocationDependentMBSF1UInformation_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_LocationDependentMBSF1UInformation_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* mbsAreaSession-ID */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* mbs-f1u-info-at-CU */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_LocationDependentMBSF1UInformation_Item_specs_1 = {
	sizeof(struct LocationDependentMBSF1UInformation_Item),
	offsetof(struct LocationDependentMBSF1UInformation_Item, _asn_ctx),
	asn_MAP_LocationDependentMBSF1UInformation_Item_tag2el_1,
	3,	/* Count of tags in the map */
	asn_MAP_LocationDependentMBSF1UInformation_Item_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	3,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_LocationDependentMBSF1UInformation_Item = {
	"LocationDependentMBSF1UInformation-Item",
	"LocationDependentMBSF1UInformation-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_LocationDependentMBSF1UInformation_Item_tags_1,
	sizeof(asn_DEF_LocationDependentMBSF1UInformation_Item_tags_1)
		/sizeof(asn_DEF_LocationDependentMBSF1UInformation_Item_tags_1[0]), /* 1 */
	asn_DEF_LocationDependentMBSF1UInformation_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_LocationDependentMBSF1UInformation_Item_tags_1)
		/sizeof(asn_DEF_LocationDependentMBSF1UInformation_Item_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_LocationDependentMBSF1UInformation_Item_1,
	3,	/* Elements count */
	&asn_SPC_LocationDependentMBSF1UInformation_Item_specs_1	/* Additional specs */
};

