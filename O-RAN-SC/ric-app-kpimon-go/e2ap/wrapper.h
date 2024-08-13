#ifndef	_WRAPPER_E2AP_H_
#define	_WRAPPER_E2AP_H_
#include <stdio.h>
#include <stdlib.h>
#include "RICsubscriptionRequest.h"
#include "RICsubscriptionResponse.h"
#include "RICsubscriptionDeleteRequest.h"
#include "RICsubscriptionDeleteResponse.h"
#include "RICcontrolRequest.h"
#include "RICindication.h"
#include "E2AP-PDU.h"
#include "InitiatingMessage.h"
#include "SuccessfulOutcome.h"
#include "UnsuccessfulOutcome.h"
#include "ProtocolIE-Container.h"
#include "ProtocolIE-Field.h"
#include "RICactionDefinition.h"
#include "RICsubsequentAction.h"
#include "CauseE2node.h"
#include "Cause.h"
#include "CauseMisc.h"
#include "CauseProtocol.h"
#include "CauseRICrequest.h"
#include "CauseRICservice.h"
#include "CauseTransport.h"
#include "RICsubscription-List-withCause.h"
#include "RICsubscription-withCause-Item.h"

//#include "CauseRIC.h"

typedef struct RICindicationMessage {
	long requestorID;
	long requestSequenceNumber;
	long ranfunctionID;
	long actionID;
	long indicationSN;
	long indicationType;
	uint8_t *indicationHeader;
	size_t indicationHeaderSize;
	uint8_t *indicationMessage;
	size_t indicationMessageSize;
	uint8_t *callProcessID;
	size_t callProcessIDSize;
} RICindicationMsg;

typedef struct RICcauseItem {
	int ricCauseType;
	long ricCauseID;
} RICcauseItem;

typedef struct RICactionAdmittedList {
	long ricActionID[16];
	int count;
} RICactionAdmittedList;

typedef struct RICactionNotAdmittedList {
	long ricActionID[16];
	RICcauseItem ricCause[16];
	int count;
} RICactionNotAdmittedList;

typedef struct RICsubscriptionResponseMessage {
	long requestorID;
	long requestSequenceNumber;
	long ranfunctionID;
	RICactionAdmittedList ricActionAdmittedList;
	RICactionNotAdmittedList ricActionNotAdmittedList;
} RICsubscriptionResponseMsg;

typedef struct RICactionDefinition {
	uint8_t *actionDefinition;
	int size;
} RICactionDefinition;

typedef struct RICSubsequentAction {
	int isValid;
	long subsequentActionType;
	long timeToWait;
} RICSubsequentAction;

size_t encode_E2AP_PDU(E2AP_PDU_t* pdu, void* buffer, size_t buf_size);
E2AP_PDU_t* decode_E2AP_PDU(const void* buffer, size_t buf_size);

/* RICsubscriptionRequest */
long e2ap_get_ric_subscription_request_sequence_number(void *buffer, size_t buf_size);
ssize_t  e2ap_set_ric_subscription_request_sequence_number(void *buffer, size_t buf_size, long sequence_number);
ssize_t e2ap_encode_ric_subscription_request_message(void *buffer, size_t buf_size, long ricRequestorID, long ricRequestSequenceNumber, long ranFunctionID, void *eventTriggerDefinition, size_t eventTriggerDefinitionSize, int actionCount, long *actionIds, long* actionTypes, RICactionDefinition *actionDefinitions, RICSubsequentAction *subsequentActionTypes);

/* RICsubscriptionResponse */
long e2ap_get_ric_subscription_response_sequence_number(void *buffer, size_t buf_size);
ssize_t  e2ap_set_ric_subscription_response_sequence_number(void *buffer, size_t buf_size, long sequence_number);
RICsubscriptionResponseMsg* e2ap_decode_ric_subscription_response_message(void *buffer, size_t buf_size);

/* RICsubscriptionFailure */
long e2ap_get_ric_subscription_failure_sequence_number(void *buffer, size_t buf_size);

/* RICsubscriptionDeleteRequest */
long e2ap_get_ric_subscription_delete_request_sequence_number(void *buffer, size_t buf_size);
ssize_t  e2ap_set_ric_subscription_delete_request_sequence_number(void *buffer, size_t buf_size, long sequence_number);
ssize_t e2ap_encode_ric_subscription_delete_request_message(void *buffer, size_t buf_size, long ricRequestorID, long ricRequestSequenceNumber, long ranFunctionID);

/* RICsubscriptionDeleteResponse */
long e2ap_get_ric_subscription_delete_response_sequence_number(void *buffer, size_t buf_size);
ssize_t  e2ap_set_ric_subscription_delete_response_sequence_number(void *buffer, size_t buf_size, long sequence_number);

/* RICsubscriptionDeleteFailure */
long e2ap_get_ric_subscription_delete_failure_sequence_number(void *buffer, size_t buf_size);

/* RICindication */
RICindicationMsg* e2ap_decode_ric_indication_message(void *buffer, size_t buf_size);
RICindicationMsg* e2ap_decode_ric_indication_m();
void e2ap_free_decoded_ric_indication_message(RICindicationMsg* msg);

#endif /* _WRAPPER_E2AP_H_ */
