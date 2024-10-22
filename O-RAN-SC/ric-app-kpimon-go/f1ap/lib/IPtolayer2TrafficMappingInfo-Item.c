/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "IPtolayer2TrafficMappingInfo-Item.h"

#include "ProtocolExtensionContainer.h"
asn_TYPE_member_t asn_MBR_IPtolayer2TrafficMappingInfo_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct IPtolayer2TrafficMappingInfo_Item, mappingInformationIndex),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_MappingInformationIndex,
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
		"mappingInformationIndex"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct IPtolayer2TrafficMappingInfo_Item, iPHeaderInformation),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_IPHeaderInformation,
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
		"iPHeaderInformation"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct IPtolayer2TrafficMappingInfo_Item, bHInfo),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BHInfo,
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
		"bHInfo"
		},
	{ ATF_POINTER, 1, offsetof(struct IPtolayer2TrafficMappingInfo_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P185,
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
static const int asn_MAP_IPtolayer2TrafficMappingInfo_Item_oms_1[] = { 3 };
static const ber_tlv_tag_t asn_DEF_IPtolayer2TrafficMappingInfo_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_IPtolayer2TrafficMappingInfo_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* mappingInformationIndex */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* iPHeaderInformation */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* bHInfo */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_IPtolayer2TrafficMappingInfo_Item_specs_1 = {
	sizeof(struct IPtolayer2TrafficMappingInfo_Item),
	offsetof(struct IPtolayer2TrafficMappingInfo_Item, _asn_ctx),
	asn_MAP_IPtolayer2TrafficMappingInfo_Item_tag2el_1,
	4,	/* Count of tags in the map */
	asn_MAP_IPtolayer2TrafficMappingInfo_Item_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	4,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_IPtolayer2TrafficMappingInfo_Item = {
	"IPtolayer2TrafficMappingInfo-Item",
	"IPtolayer2TrafficMappingInfo-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_IPtolayer2TrafficMappingInfo_Item_tags_1,
	sizeof(asn_DEF_IPtolayer2TrafficMappingInfo_Item_tags_1)
		/sizeof(asn_DEF_IPtolayer2TrafficMappingInfo_Item_tags_1[0]), /* 1 */
	asn_DEF_IPtolayer2TrafficMappingInfo_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_IPtolayer2TrafficMappingInfo_Item_tags_1)
		/sizeof(asn_DEF_IPtolayer2TrafficMappingInfo_Item_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_IPtolayer2TrafficMappingInfo_Item_1,
	4,	/* Elements count */
	&asn_SPC_IPtolayer2TrafficMappingInfo_Item_specs_1	/* Additional specs */
};

