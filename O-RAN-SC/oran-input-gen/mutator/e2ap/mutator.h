#ifndef MUTATOR_H
#define MUTATOR_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "E2AP-PDU.h"
#include "asn_mutate.h"
#include "asn_random_fill.h"
#include "per_encoder.h"
#include "per_decoder.h"

#define E2SM_XER_PRINT(stream,type,pdu)                    \
        xer_fprint((stream == NULL) ? stdout : stream,type,pdu);

size_t get_sizeof_void_ptr();
asn_dec_rval_t decode_msg(void* buffer, size_t size, E2AP_PDU_t** output);
int gen_msg(E2AP_PDU_t** output);
int mutate_msg(E2AP_PDU_t** output);
ssize_t encode_msg(E2AP_PDU_t* input, void* buffer);
asn_dec_rval_t decode_msg_file(const char* filename, E2AP_PDU_t** output);
int mutateMessageTest();
int main(int argc, char* argv[]);
void printBuffer(void* buffer, size_t size);

int mutateMessage(void *data, long size, void** buffer);
void setRandSeed();
int decode_msg_print(void* buffer, size_t size);

#endif /* MUTATOR_H */
