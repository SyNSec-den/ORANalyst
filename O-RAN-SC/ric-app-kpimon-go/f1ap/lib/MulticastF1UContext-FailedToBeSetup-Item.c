/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "MulticastF1UContext-FailedToBeSetup-Item.h"

#include "Cause.h"
#include "ProtocolExtensionContainer.h"
asn_TYPE_member_t asn_MBR_MulticastF1UContext_FailedToBeSetup_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct MulticastF1UContext_FailedToBeSetup_Item, mRB_ID),
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
	{ ATF_POINTER, 2, offsetof(struct MulticastF1UContext_FailedToBeSetup_Item, cause),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_Cause,
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
		"cause"
		},
	{ ATF_POINTER, 1, offsetof(struct MulticastF1UContext_FailedToBeSetup_Item, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P205,
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
static const int asn_MAP_MulticastF1UContext_FailedToBeSetup_Item_oms_1[] = { 1, 2 };
static const ber_tlv_tag_t asn_DEF_MulticastF1UContext_FailedToBeSetup_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_MulticastF1UContext_FailedToBeSetup_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* mRB-ID */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* cause */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_MulticastF1UContext_FailedToBeSetup_Item_specs_1 = {
	sizeof(struct MulticastF1UContext_FailedToBeSetup_Item),
	offsetof(struct MulticastF1UContext_FailedToBeSetup_Item, _asn_ctx),
	asn_MAP_MulticastF1UContext_FailedToBeSetup_Item_tag2el_1,
	3,	/* Count of tags in the map */
	asn_MAP_MulticastF1UContext_FailedToBeSetup_Item_oms_1,	/* Optional members */
	2, 0,	/* Root/Additions */
	3,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_MulticastF1UContext_FailedToBeSetup_Item = {
	"MulticastF1UContext-FailedToBeSetup-Item",
	"MulticastF1UContext-FailedToBeSetup-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_MulticastF1UContext_FailedToBeSetup_Item_tags_1,
	sizeof(asn_DEF_MulticastF1UContext_FailedToBeSetup_Item_tags_1)
		/sizeof(asn_DEF_MulticastF1UContext_FailedToBeSetup_Item_tags_1[0]), /* 1 */
	asn_DEF_MulticastF1UContext_FailedToBeSetup_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_MulticastF1UContext_FailedToBeSetup_Item_tags_1)
		/sizeof(asn_DEF_MulticastF1UContext_FailedToBeSetup_Item_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_MulticastF1UContext_FailedToBeSetup_Item_1,
	3,	/* Elements count */
	&asn_SPC_MulticastF1UContext_FailedToBeSetup_Item_specs_1	/* Additional specs */
};

