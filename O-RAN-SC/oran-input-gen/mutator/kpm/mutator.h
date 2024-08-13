#ifndef MUTATOR_H
#define MUTATOR_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "E2SM-KPM-IndicationHeader.h"
#include "E2SM-KPM-IndicationMessage.h"
#include "asn_mutate.h"
#include "asn_random_fill.h"
#include "per_encoder.h"
#include "per_decoder.h"

#define E2SM_XER_PRINT(stream,type,pdu)                    \
        xer_fprint((stream == NULL) ? stdout : stream,type,pdu);

size_t get_sizeof_void_ptr();
asn_dec_rval_t decode_msg(void* buffer, size_t size, E2SM_KPM_IndicationMessage_t** output);
int gen_hdr(E2SM_KPM_IndicationHeader_t** output);
int mutate_hdr(E2SM_KPM_IndicationHeader_t** output);
ssize_t encode_hdr(E2SM_KPM_IndicationHeader_t* input, void* buffer);
asn_dec_rval_t decode_hdr_file(const char* filename, E2SM_KPM_IndicationHeader_t** output);
int mutateHeaderTest();
int gen_msg(E2SM_KPM_IndicationMessage_t** output);
int mutate_msg(E2SM_KPM_IndicationMessage_t** output);
ssize_t encode_msg(E2SM_KPM_IndicationMessage_t* input, void* buffer);
asn_dec_rval_t decode_msg_file(const char* filename, E2SM_KPM_IndicationMessage_t** output);
int mutateMessageTest();
int main(int argc, char* argv[]);
void printBuffer(void* buffer, size_t size);

int mutateMessage(void *data, long size, void** buffer);
int mutateHeader(void *data, long size, void* buffer);
void setRandSeed();
int decode_msg_print(void* buffer, size_t size);

#endif /* MUTATOR_H */
