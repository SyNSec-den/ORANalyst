/*
 * Copyright 2020 AT&T Intellectual Property
 * Copyright 2020 Nokia
 * Copyright 2023 Capgemini
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


#ifndef E2_BUILDRUNNAME_H
#define E2_BUILDRUNNAME_H

#include <3rdparty/oranE2/ProtocolIE-Field.h>
#include "oranE2/ProtocolIE-Container.h"
#include "oranE2/ProtocolIE-Field.h"
#include "oranE2/GlobalE2node-gNB-ID.h"
#include "oranE2/GlobalE2node-en-gNB-ID.h"
#include "oranE2/GlobalE2node-ng-eNB-ID.h"
#include "oranE2/GlobalE2node-eNB-ID.h"

/**    02 F8 29
 * return the size of the string //
 */
static int translatePlmnId(char * plmnId, const unsigned char *data, const char* type) {
    auto mcc1 = (unsigned char)((unsigned char)data[0] & (unsigned char)0x0F);
    auto mcc2 = (unsigned char)(((unsigned char)((unsigned char)data[0] & (unsigned char)0xF0)) >> (unsigned char)4);
    ///auto mcc3 = (unsigned char)((data[1] & (unsigned char)0xF0) >> (unsigned char)4);
    auto mcc3 = (unsigned char)((unsigned char)(data[1] & (unsigned char)0x0F));

    auto mnc1 = (unsigned char)(((unsigned char)((unsigned char)data[1] & (unsigned char)0xF0)) >> (unsigned char)4);
    auto mnc2 =  (unsigned char)((unsigned char)data[2] & (unsigned char)0x0F) ;
    //auto mnc3 = (unsigned char)(((unsigned char)(data[1] & (unsigned char)0x0F) >> (unsigned char)4) );
    auto mnc3 = (unsigned char)(((unsigned char)((unsigned char)data[2] & (unsigned char)0xF0)) >> (unsigned char)4);

    int j = 0;
    if (mnc1 != 15) {
        j = snprintf(plmnId, 20, "%s%1d%1d%1d_%1d%1d%1d", type, mcc1, mcc2, mcc3, mnc1, mnc2, mnc3);
    }
    else {
        j = snprintf(plmnId, 20, "%s%1d%1d%1d_0%1d%1d", type, mcc1, mcc2, mcc3, mnc2, mnc3);
    }

    return j;
}

static int translateBitStringToChar(char *ranName, BIT_STRING_t &data, long cuid, long duid) {
    // dont care of last unused bits
    char buffer[256] {};
    // auto j = snprintf(buffer, 256, "%s_", ranName);
    int j = 0;
/*
 // ran name decimal
    unsigned long bitValue = 0;
    for (auto i = 0; i < (int)data.size; i++) {
        bitValue <<= (unsigned long)8;
        bitValue += data.buf[i];
    }

    j = snprintf(buffer, 256, "%s%ld", ranName, bitValue);

    memcpy(ranName, buffer, j);
*/
   

    unsigned b1 = 0;
    unsigned b2 = 0;
    unsigned tmp_digit=0;
    j = snprintf(buffer, 256, "%s_", ranName);
    memcpy(ranName, buffer, j);
    for (auto i = 0; i < (int)data.size; i++) {
        //
        // we need to shift trailing zeros in the bit string (from asn1c) to leading zeros
        //
        tmp_digit = (0==i ? (uint8_t) 0: (uint8_t) data.buf[i-1])<<(8-data.bits_unused);
        tmp_digit = tmp_digit | ((unsigned) data.buf[i] >> data.bits_unused);

	b1 = tmp_digit & (unsigned)0xF0;
        b1 = b1 >> (unsigned)4;
        j = snprintf(buffer, 256, "%s%1x", ranName, b1);
        memcpy(ranName, buffer, j);
        b2 = tmp_digit & (unsigned)0x0F;
        j = snprintf(buffer, 256, "%s%1x", ranName, b2);
        memcpy(ranName, buffer, j);
    }

    /*Note: Deployment where CU-UP and DU are combined
    (but do not include the CP-CP) and have a single E2 connection
     is not supported. The combination of CU-CP, CU-UP, and DU will be
     treated as a single gNB and expect it to have only the
     global gNB ID in its E2 Setup ID*/

    if ( cuid >= 0 && duid >= 0){
           j= snprintf(buffer, 256, "%s", ranName);
       } else if ( cuid >= 0 ){
           j= snprintf(buffer, 256, "%s_%lx", ranName,cuid);
       }else if ( duid >= 0 ){
           j= snprintf(buffer, 256, "%s_%lx", ranName,duid);
       }else{
           j= snprintf(buffer, 256, "%s", ranName);
       }
       printf("cuupid =%ld\n",cuid);
       printf("duid =%ld\n",duid);
       printf("ranName =%s\n",ranName);
       memcpy(ranName, buffer, j);

    return j;
}


int buildRanName(char *ranName, E2setupRequestIEs_t *ie) {
    switch (ie->value.choice.GlobalE2node_ID.present) {
        case GlobalE2node_ID_PR_gNB: {
            auto *gnb = ie->value.choice.GlobalE2node_ID.choice.gNB;
            long cuupid = -1;
            long duid = -1;
           
            if (gnb->global_gNB_ID.gnb_id.present == GNB_ID_Choice_PR_gnb_ID) {
                if (gnb->gNB_CU_UP_ID != NULL && gnb->gNB_DU_ID != NULL){
                    printf("\ngNB_CU_UP_ID and gNB_DU_ID is not null\n");
                    asn_INTEGER2long(gnb->gNB_CU_UP_ID,&cuupid);
                    asn_INTEGER2long(gnb->gNB_DU_ID,&duid);
                    translatePlmnId(ranName, (const unsigned char *)gnb->global_gNB_ID.plmn_id.buf, (const char *)"gnb_");
                }else if (gnb->gNB_DU_ID == NULL && gnb->gNB_CU_UP_ID == NULL ){
                    printf("\ngNB_CU_UP_ID and gNB_DU_ID is null\n");
                    translatePlmnId(ranName, (const unsigned char *)gnb->global_gNB_ID.plmn_id.buf, (const char *)"gnb_");
                }else if (gnb->gNB_CU_UP_ID != NULL ){
                    printf("\ngNB_CU_UP_ID is not null\n");
                    asn_INTEGER2long(gnb->gNB_CU_UP_ID,&cuupid);
                    translatePlmnId(ranName, (const unsigned char *)gnb->global_gNB_ID.plmn_id.buf, (const char *)"gnbc_");
                }else if (gnb->gNB_DU_ID != NULL ){
                    printf("\ngNB_DU_ID is not null\n");
                    asn_INTEGER2long(gnb->gNB_DU_ID,&duid);
                    translatePlmnId(ranName, (const unsigned char *)gnb->global_gNB_ID.plmn_id.buf, (const char *)"gnbd_");
                }
                translateBitStringToChar(ranName, gnb->global_gNB_ID.gnb_id.choice.gnb_ID,cuupid,duid);
            }else{
                translatePlmnId(ranName, (const unsigned char *)gnb->global_gNB_ID.plmn_id.buf, (const char *)"gnb_");
            }
            break;
        }
        case GlobalE2node_ID_PR_en_gNB: {
            auto *enGnb = ie->value.choice.GlobalE2node_ID.choice.en_gNB;
            translatePlmnId(ranName,
                            (const unsigned char *)enGnb->global_en_gNB_ID.pLMN_Identity.buf,
                            (const char *)"en_gnb_");
            if (enGnb->global_en_gNB_ID.gNB_ID.present == ENGNB_ID_PR_gNB_ID) {
                translateBitStringToChar(ranName, enGnb->global_en_gNB_ID.gNB_ID.choice.gNB_ID,-1,-1);
            }
            break;
        }
        case GlobalE2node_ID_PR_ng_eNB: {
            auto *ngEnb = ie->value.choice.GlobalE2node_ID.choice.ng_eNB;
            switch (ngEnb->global_ng_eNB_ID.enb_id.present) {
                case ENB_ID_Choice_PR_enb_ID_macro: {
                    translatePlmnId(ranName, (const unsigned char *)ngEnb->global_ng_eNB_ID.plmn_id.buf, (const char *)"ng_enB_macro_");
                    translateBitStringToChar(ranName, ngEnb->global_ng_eNB_ID.enb_id.choice.enb_ID_macro,-1,-1);
                    break;
                }
                case ENB_ID_Choice_PR_enb_ID_shortmacro: {
                    translatePlmnId(ranName, (const unsigned char *)ngEnb->global_ng_eNB_ID.plmn_id.buf, (const char *)"ng_enB_shortmacro_");
                    translateBitStringToChar(ranName, ngEnb->global_ng_eNB_ID.enb_id.choice.enb_ID_shortmacro,-1,-1);
                    break;
                }
                case ENB_ID_Choice_PR_enb_ID_longmacro: {
                    translatePlmnId(ranName, (const unsigned char *)ngEnb->global_ng_eNB_ID.plmn_id.buf, (const char *)"ng_enB_longmacro_");
                    translateBitStringToChar(ranName, ngEnb->global_ng_eNB_ID.enb_id.choice.enb_ID_longmacro,-1,-1);
                    break;
                }
                case ENB_ID_Choice_PR_NOTHING: {
                    break;
                }
                default:
                    break;
            }
	    break;
        }
        case GlobalE2node_ID_PR_eNB: {
            auto *enb = ie->value.choice.GlobalE2node_ID.choice.eNB;
            switch (enb->global_eNB_ID.eNB_ID.present) {
                case ENB_ID_PR_macro_eNB_ID: {
                    translatePlmnId(ranName, (const unsigned char *)enb->global_eNB_ID.pLMN_Identity.buf, (const char *)"enB_macro_");
                    translateBitStringToChar(ranName, enb->global_eNB_ID.eNB_ID.choice.macro_eNB_ID,-1,-1);
                    break;
                }
                case ENB_ID_PR_home_eNB_ID: {
                    translatePlmnId(ranName, (const unsigned char *)enb->global_eNB_ID.pLMN_Identity.buf, (const char *)"enB_home_");
                    translateBitStringToChar(ranName, enb->global_eNB_ID.eNB_ID.choice.home_eNB_ID,-1,-1);
                    break;
                }
                case ENB_ID_PR_short_Macro_eNB_ID: {
                    translatePlmnId(ranName, (const unsigned char *)enb->global_eNB_ID.pLMN_Identity.buf, (const char *)"enB_shortmacro_");
                    translateBitStringToChar(ranName, enb->global_eNB_ID.eNB_ID.choice.short_Macro_eNB_ID,-1,-1);
                    break;
                }
                case ENB_ID_PR_long_Macro_eNB_ID: {
                    translatePlmnId(ranName, (const unsigned char *)enb->global_eNB_ID.pLMN_Identity.buf, (const char *)"enB_longmacro_");
                    translateBitStringToChar(ranName, enb->global_eNB_ID.eNB_ID.choice.long_Macro_eNB_ID,-1,-1);
                    break;
                }
                case ENB_ID_PR_NOTHING: {
                    break;
                }
                default: {
                    break;
                }
            }
	    break;
        }
        case GlobalE2node_ID_PR_NOTHING:
        default:
            return -1;
    }
    return 0;
}


#endif //E2_BUILDRUNNAME_H
