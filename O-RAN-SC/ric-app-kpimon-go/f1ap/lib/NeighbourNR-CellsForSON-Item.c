/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "NeighbourNR-CellsForSON-Item.h"

#include "NR-ModeInfoRel16.h"
#include "SSB-PositionsInBurst.h"
#include "NRPRACHConfig.h"
#include "ProtocolExtensionContainer.h"
asn_TYPE_member_t asn_MBR_NeighbourNR_CellsForSON_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct NeighbourNR_CellsForSON_Item, nRCGI),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NRCGI,
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
		"nRCGI"
		},
	{ ATF_POINTER, 4, offsetof(struct NeighbourNR_CellsForSON_Item, nR_ModeInfoRel16),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_NR_ModeInfoRel16,
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
		"nR-ModeInfoRel16"
		},
	{ ATF_POINTER, 3, offsetof(struct NeighbourNR_CellsForSON_Item, sSB_PositionsInBurst),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_SSB_PositionsInBurst,
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
		"sSB-PositionsInBurst"
		},
	{ ATF_POINTER, 2, offsetof(struct NeighbourNR_CellsForSON_Item, nRPRACHConfig),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NRPRACHConfig,
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
		"nRPRACHConfig"
		},
	{ ATF_POINTER, 1, offsetof(struct NeighbourNR_CellsForSON_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P231,
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
static const int asn_MAP_NeighbourNR_CellsForSON_Item_oms_1[] = { 1, 2, 3, 4 };
static const ber_tlv_tag_t asn_DEF_NeighbourNR_CellsForSON_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_NeighbourNR_CellsForSON_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* nRCGI */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* nR-ModeInfoRel16 */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* sSB-PositionsInBurst */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* nRPRACHConfig */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_NeighbourNR_CellsForSON_Item_specs_1 = {
	sizeof(struct NeighbourNR_CellsForSON_Item),
	offsetof(struct NeighbourNR_CellsForSON_Item, _asn_ctx),
	asn_MAP_NeighbourNR_CellsForSON_Item_tag2el_1,
	5,	/* Count of tags in the map */
	asn_MAP_NeighbourNR_CellsForSON_Item_oms_1,	/* Optional members */
	4, 0,	/* Root/Additions */
	5,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_NeighbourNR_CellsForSON_Item = {
	"NeighbourNR-CellsForSON-Item",
	"NeighbourNR-CellsForSON-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_NeighbourNR_CellsForSON_Item_tags_1,
	sizeof(asn_DEF_NeighbourNR_CellsForSON_Item_tags_1)
		/sizeof(asn_DEF_NeighbourNR_CellsForSON_Item_tags_1[0]), /* 1 */
	asn_DEF_NeighbourNR_CellsForSON_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_NeighbourNR_CellsForSON_Item_tags_1)
		/sizeof(asn_DEF_NeighbourNR_CellsForSON_Item_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_NeighbourNR_CellsForSON_Item_1,
	5,	/* Elements count */
	&asn_SPC_NeighbourNR_CellsForSON_Item_specs_1	/* Additional specs */
};

