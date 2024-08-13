#ifndef _WRAPPER_F1AP_H_
#define _WRAPPER_F1AP_H_
#include<F1AP-PDU.h>
#include<InitiatingMessage.h>
#include<ProtocolIE-SingleContainer.h>
typedef struct CellIds{
        char ** cellids;
        char **plmn;
        int size;
} CellIds_t;

unsigned char decodeBase64Char(char c);
unsigned char* base64Decode(const char* input, size_t* outputLength);
CellIds_t decodeF1apGetCellIds(const char* encodedString );
#endif /* _WRAPPER_F1AP_H_ */

