/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "F1AP-IEs"
 * 	found in "F1AP-IEs.asn"
 * 	`asn1c -pdu=auto -fincludes-quoted -fcompound-names -findirect-choice -fno-include-deps -no-gen-example -gen-APER`
 */

#include "GNB-DU-Cell-Resource-Configuration.h"

#include "DUF-Slot-Config-List.h"
#include "HSNASlotConfigList.h"
#include "ProtocolExtensionContainer.h"
asn_TYPE_member_t asn_MBR_GNB_DU_Cell_Resource_Configuration_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct GNB_DU_Cell_Resource_Configuration, subcarrierSpacing),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_SubcarrierSpacing,
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
		"subcarrierSpacing"
		},
	{ ATF_POINTER, 2, offsetof(struct GNB_DU_Cell_Resource_Configuration, dUFTransmissionPeriodicity),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_DUFTransmissionPeriodicity,
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
		"dUFTransmissionPeriodicity"
		},
	{ ATF_POINTER, 1, offsetof(struct GNB_DU_Cell_Resource_Configuration, dUF_Slot_Config_List),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_DUF_Slot_Config_List,
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
		"dUF-Slot-Config-List"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct GNB_DU_Cell_Resource_Configuration, hSNATransmissionPeriodicity),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_HSNATransmissionPeriodicity,
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
		"hSNATransmissionPeriodicity"
		},
	{ ATF_POINTER, 2, offsetof(struct GNB_DU_Cell_Resource_Configuration, hsNSASlotConfigList),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_HSNASlotConfigList,
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
		"hsNSASlotConfigList"
		},
	{ ATF_POINTER, 1, offsetof(struct GNB_DU_Cell_Resource_Configuration, iE_Extensions),
		(ASN_TAG_CLASS_CONTEXT | (5 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_ProtocolExtensionContainer_160P153,
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
static const int asn_MAP_GNB_DU_Cell_Resource_Configuration_oms_1[] = { 1, 2, 4, 5 };
static const ber_tlv_tag_t asn_DEF_GNB_DU_Cell_Resource_Configuration_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_GNB_DU_Cell_Resource_Configuration_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* subcarrierSpacing */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* dUFTransmissionPeriodicity */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* dUF-Slot-Config-List */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* hSNATransmissionPeriodicity */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 }, /* hsNSASlotConfigList */
    { (ASN_TAG_CLASS_CONTEXT | (5 << 2)), 5, 0, 0 } /* iE-Extensions */
};
asn_SEQUENCE_specifics_t asn_SPC_GNB_DU_Cell_Resource_Configuration_specs_1 = {
	sizeof(struct GNB_DU_Cell_Resource_Configuration),
	offsetof(struct GNB_DU_Cell_Resource_Configuration, _asn_ctx),
	asn_MAP_GNB_DU_Cell_Resource_Configuration_tag2el_1,
	6,	/* Count of tags in the map */
	asn_MAP_GNB_DU_Cell_Resource_Configuration_oms_1,	/* Optional members */
	4, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_GNB_DU_Cell_Resource_Configuration = {
	"GNB-DU-Cell-Resource-Configuration",
	"GNB-DU-Cell-Resource-Configuration",
	&asn_OP_SEQUENCE,
	asn_DEF_GNB_DU_Cell_Resource_Configuration_tags_1,
	sizeof(asn_DEF_GNB_DU_Cell_Resource_Configuration_tags_1)
		/sizeof(asn_DEF_GNB_DU_Cell_Resource_Configuration_tags_1[0]), /* 1 */
	asn_DEF_GNB_DU_Cell_Resource_Configuration_tags_1,	/* Same as above */
	sizeof(asn_DEF_GNB_DU_Cell_Resource_Configuration_tags_1)
		/sizeof(asn_DEF_GNB_DU_Cell_Resource_Configuration_tags_1[0]), /* 1 */
	{
#if !defined(ASN_DISABLE_OER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_OER_SUPPORT) */
#if !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT)
		0,
#endif  /* !defined(ASN_DISABLE_UPER_SUPPORT) || !defined(ASN_DISABLE_APER_SUPPORT) */
		SEQUENCE_constraint
	},
	asn_MBR_GNB_DU_Cell_Resource_Configuration_1,
	6,	/* Elements count */
	&asn_SPC_GNB_DU_Cell_Resource_Configuration_specs_1	/* Additional specs */
};

