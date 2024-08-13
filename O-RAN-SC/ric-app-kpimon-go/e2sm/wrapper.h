#ifndef	_WRAPPER_H_
#define	_WRAPPER_H_

#include <E2SM-KPM-EventTriggerDefinition.h>
#include <E2SM-KPM-EventTriggerDefinition-Format1.h>
#include <E2SM-KPM-ActionDefinition.h>
#include <E2SM-KPM-ActionDefinition-Format1.h>
#include <MeasurementInfoList.h>
#include <GranularityPeriod.h>
#include <MeasurementInfoItem.h>
#include <MeasurementType.h>
#include <MeasurementTypeID.h>
#include <LabelInfoItem.h>
#include <LabelInfoList.h>
#include <MeasurementLabel.h>
#include <CGI.h>
#include <PLMNIdentity.h>
#include <E2SM-KPM-IndicationHeader.h>
#include <E2SM-KPM-IndicationHeader-Format1.h>
//#include <GlobalNGRANNodeID.h>
#include<GlobalRANNodeID.h>
#include <GlobalGNB-ID.h>
#include <GlobalenGNB-ID.h>
#include <GlobalNgENB-ID.h>
#include <GlobalENB-ID.h>
#include <PLMNIdentity.h>
#include <GNB-ID.h>
#include <GNB-CU-UP-ID.h>
#include <GNB-DU-ID.h>
#include <EN-GNB-ID.h>
#include <ENB-ID.h>
#include <NR-CGI.h>
#include <S-NSSAI.h>
#include <E2SM-KPM-IndicationMessage.h>
#include <E2SM-KPM-IndicationMessage-Format1.h>
#include <E2SM-KPM-IndicationMessage-Format2.h>
#include<MeasurementDataItem.h>
#include<MeasurementRecordItem.h>
#include <TimeStamp.h>


#include <stdio.h>
#include "asn_application.h"
#include <E2SM-KPM-EventTriggerDefinition.h>
#include <E2SM-KPM-EventTriggerDefinition-Format1.h>
#include <E2SM-KPM-ActionDefinition.h>
#include <E2SM-KPM-ActionDefinition-Format1.h>
#include <E2SM-KPM-ActionDefinition-Format3.h>
#include <MeasurementInfoList.h>
#include <GranularityPeriod.h>
#include <MeasurementInfoItem.h>
#include <MeasurementType.h>
#include<MeasurementCondList.h>
#include<MeasurementCondItem.h>
#include<MatchingCondItem.h>
#include<E2SM-KPM-RANfunction-Description.h>
#include<RIC-ReportStyle-Item.h>
#include<MeasurementInfo-Action-Item.h>

#include <MeasurementTypeID.h>
#include <LabelInfoItem.h>
#include <LabelInfoList.h>
#include <MeasurementLabel.h>
#include <CGI.h>
#include <PLMNIdentity.h>
#include <E2SM-KPM-IndicationHeader.h>
#include <E2SM-KPM-IndicationHeader-Format1.h>
//#include <GlobalNGRANNodeID.h>
#include <GlobalGNB-ID.h>
#include <GlobalenGNB-ID.h>
#include <GlobalNgENB-ID.h>
#include <GlobalENB-ID.h>
#include <PLMNIdentity.h>
#include <GNB-ID.h>
#include <GNB-CU-UP-ID.h>
#include <GNB-DU-ID.h>
#include <EN-GNB-ID.h>
#include <ENB-ID.h>
#include <NR-CGI.h>
#include <S-NSSAI.h>
#include <E2SM-KPM-IndicationMessage.h>
#include <E2SM-KPM-IndicationMessage-Format1.h>
#include <E2SM-KPM-IndicationMessage-Format2.h>
#include <TimeStamp.h>
#include<per_encoder.h>
#include <errno.h>

/*
typedef struct Plm {
	uint8_t *plms;
	int size;
} Plm;


typedef struct Nrc {
	uint8_t *nRcs;
	int size;
	int bits_unused;
} Nrc;

*/
typedef struct encode_act_Def_result{
	int * array;
	int length;
} encode_act_Def_result_t;

typedef struct ranCellUeKpi{
        char ** ueKpi;
        char ** cellKpi;
	int ueKpiSize;
	int cellKpiSize;
} ranCellUeKpi_t;

ssize_t e2sm_encode_ric_event_trigger_definition(void *buffer, size_t buf_size, size_t event_trigger_count, long *RT_periods);

//ssize_t e2sm_encode_ric_action_definition(void *buffer, size_t buf_size, long action_Def_count, long meas_ID, long ric_style_type, long GR_periods, Plm plm, Nrc nRc);
//

ranCellUeKpi_t buildRanCellUeKpi(const char *hex_values);

void freeMemorydRanCellUeKpi(ranCellUeKpi_t res);

struct encode_act_Def_result encode_action_Definition(const char *hex_values, int determine);

size_t e2sm_encode_ric_action_definition_format1_by_name(unsigned char *buf, size_t *buf_size, const char **id_tmp , size_t measIdcount, long ric_style_type, unsigned long granulPeriod, unsigned char  *p, unsigned char *nR);

size_t e2sm_encode_ric_action_definition_format1_by_id(unsigned char *buf, size_t *buf_size, long *id , size_t measIdcount, long ric_style_type, unsigned long granulPeriod, unsigned char  *p, unsigned char *nR);

size_t e2sm_encode_ric_action_definition_format3_by_name(unsigned char *buf, size_t *buf_size, char **id_tmp , size_t measIdcount, long ric_style_type, unsigned long granulPeriod);

size_t e2sm_encode_ric_action_definition_format3_by_id(unsigned char *buf, size_t *buf_size, long *id , size_t measIdcount, long ric_style_type, unsigned long granulPeriod);

E2SM_KPM_IndicationHeader_t* e2sm_decode_ric_indication_header(void *buffer, size_t buf_size);
void e2sm_free_ric_indication_header(E2SM_KPM_IndicationHeader_t* indHdr);
E2SM_KPM_IndicationMessage_t* e2sm_decode_ric_indication_message(void *buffer, size_t buf_size);
void e2sm_free_ric_indication_message(E2SM_KPM_IndicationMessage_t* indMsg);

#endif /* _WRAPPER_H_ */
