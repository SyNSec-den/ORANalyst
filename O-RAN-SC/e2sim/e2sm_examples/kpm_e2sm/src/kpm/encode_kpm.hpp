/*****************************************************************************
#                                                                            *
# Copyright 2020 AT&T Intellectual Property                                  *
#                                                                            *
# Licensed under the Apache License, Version 2.0 (the "License");            *
# you may not use this file except in compliance with the License.           *
# You may obtain a copy of the License at                                    *
#                                                                            *
#      http://www.apache.org/licenses/LICENSE-2.0                            *
#                                                                            *
# Unless required by applicable law or agreed to in writing, software        *
# distributed under the License is distributed on an "AS IS" BASIS,          *
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   *
# See the License for the specific language governing permissions and        *
# limitations under the License.                                             *
#                                                                            *
******************************************************************************/

#ifndef ENCODE_KPM_HPP
#define ENCODE_KPM_HPP

extern "C" {
  #include "OCTET_STRING.h"
  #include "asn_application.h"
  #include "GranularityPeriod.h"
  #include "TimeStamp.h"
  #include "E2SM-KPM-IndicationHeader-Format1.h"
  #include "E2SM-KPM-IndicationMessage-Format1.h"
  #include "E2SM-KPM-IndicationMessage.h"
  #include "E2SM-KPM-IndicationHeader.h"  
  #include "E2SM-KPM-RANfunction-Description.h"
  #include "RIC-EventTriggerStyle-Item.h"
  #include "RIC-ReportStyle-Item.h"
  #include "MeasurementDataItem.h"
  #include "MeasurementRecordItem.h"
  #include "MeasurementInfo-Action-Item.h"
  #include "MeasurementInfoList.h"
  #include "MeasurementInfoItem.h"
  #include "LabelInfoItem.h"
  #include "S-NSSAI.h"
}

void encode_kpm(E2SM_KPM_IndicationMessage_t* indicationmessage);

void encode_kpm_bak(E2SM_KPM_IndicationMessage_t* indicationmessage);

void encode_kpm_function_description(E2SM_KPM_RANfunction_Description_t* ranfunc_desc);

void encode_kpm_report_rancontainer_du(E2SM_KPM_IndicationMessage_t *indMsg);

void encode_kpm_report_style1(E2SM_KPM_IndicationMessage_t* indicationmessage);

void kpm_report_indication_header_initialized(E2SM_KPM_IndicationHeader_t *ihead, uint8_t *plmnid_buf, uint8_t *sst_buf, uint8_t *sd_buf, long fqival, long qcival, uint8_t *nrcellid_buf, uint8_t *gnbid_buf, int gnbid_unused, uint8_t *cuupid_buf, uint8_t *duid_buf, uint8_t *cuupname_buf);

void ue_meas_kpm_report_indication_message_initialized(E2SM_KPM_IndicationMessage_t* indicationmessage,uint8_t *nrcellid_buf,uint8_t *crnti_buf,const uint8_t *serving_buf, const uint8_t *neighbor_buf);

void cell_meas_kpm_report_indication_message_style_1_initialized(E2SM_KPM_IndicationMessage_t* indicationmessage, long fiveqi, long dl_prb_usage, long ul_prb_usage, uint8_t* nrcellid_buf, long *dl_prbs, long *ul_prbs);



#endif
