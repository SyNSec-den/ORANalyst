/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "TRP-PRS-Info-List-Item.h"

#include "NRCGI.h"
#include "ProtocolExtensionContainer.h"
asn_TYPE_member_t asn_MBR_TRP_PRS_Info_List_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct TRP_PRS_Info_List_Item, tRP_ID),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_TRPID,
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
		"tRP-ID"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TRP_PRS_Info_List_Item, nR_PCI),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NRPCI,
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
		"nR-PCI"
		},
	{ ATF_POINTER, 1, offsetof(struct TRP_PRS_Info_List_Item, cGI_NR),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
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
		"cGI-NR"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TRP_PRS_Info_List_Item, pRSConfiguration),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PRSConfiguration,
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
		"pRSConfiguration"
		},
	{ ATF_POINTER, 1, offsetof(struct TRP_PRS_Info_List_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P450,
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
static const int asn_MAP_TRP_PRS_Info_List_Item_oms_1[] = { 2, 4 };
static const ber_tlv_tag_t asn_DEF_TRP_PRS_Info_List_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_TRP_PRS_Info_List_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* tRP-ID */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* nR-PCI */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* cGI-NR */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* pRSConfiguration */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_TRP_PRS_Info_List_Item_specs_1 = {
	sizeof(struct TRP_PRS_Info_List_Item),
	offsetof(struct TRP_PRS_Info_List_Item, _asn_ctx),
	asn_MAP_TRP_PRS_Info_List_Item_tag2el_1,
	5,	/* Count of tags in the map */
	asn_MAP_TRP_PRS_Info_List_Item_oms_1,	/* Optional members */
	2, 0,	/* Root/Additions */
	5,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_TRP_PRS_Info_List_Item = {
	"TRP-PRS-Info-List-Item",
	"TRP-PRS-Info-List-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_TRP_PRS_Info_List_Item_tags_1,
	sizeof(asn_DEF_TRP_PRS_Info_List_Item_tags_1)
		/sizeof(asn_DEF_TRP_PRS_Info_List_Item_tags_1[0]), /* 1 */
	asn_DEF_TRP_PRS_Info_List_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_TRP_PRS_Info_List_Item_tags_1)
		/sizeof(asn_DEF_TRP_PRS_Info_List_Item_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_TRP_PRS_Info_List_Item_1,
	5,	/* Elements count */
	&asn_SPC_TRP_PRS_Info_List_Item_specs_1	/* Additional specs */
};

