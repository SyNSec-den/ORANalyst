#include <gtest/gtest.h>
#include <gmock/gmock.h>
#include "sctpThread.h"
#include <sys/epoll.h>
#include "E2AP-PDU.h"

using namespace testing;

typedef struct {
    int32_t mtype;                      // message type  ("long" network integer)
    int32_t	plen;                       // payload length (sender data length in payload)
    int32_t rmr_ver;                    // our internal message version number
    unsigned char xid[RMR_MAX_XID];	    // space for user transaction id or somesuch
    unsigned char sid[RMR_MAX_SID];	    // sender ID for return to sender needs
    unsigned char src[RMR_MAX_SRC];	    // name:port of the sender (source)
    unsigned char meid[RMR_MAX_MEID];   // managed element id.
    struct timespec	ts;                 // timestamp ???
    int32_t	flags;                      // HFL_* constants
    int32_t	len0;                       // length of the RMr header data
    int32_t	len1;                       // length of the tracing data
    int32_t	len2;                       // length of data 1 (d1)
    int32_t	len3;                       // length of data 2 (d2)
    int32_t	sub_id;                     // subscription id (-1 invalid)

    unsigned char srcip[RMR_MAX_SRC];   // ip address and port of the source
} uta_mhdr_t;

void init_memories(ReportingMessages_t &message, RmrMessagesBuffer_t &rmrMessageBuffer, sctp_params_t &sctp_ut_params);

TEST(sctp, TEST1) {
    mdclog_level_set(MDCLOG_DEBUG);
    string s;
    s = translateRmrErrorMessages(0);
    EXPECT_THAT(s, HasSubstr("RMR_OK"));
    s = translateRmrErrorMessages(1);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_BADARG"));
    s = translateRmrErrorMessages(2);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_NOENDPT"));
    s = translateRmrErrorMessages(3);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_EMPTY"));
    s = translateRmrErrorMessages(4);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_NOHDR"));
    s = translateRmrErrorMessages(5);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_SENDFAILED"));
    s = translateRmrErrorMessages(6);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_CALLFAILED"));
    s = translateRmrErrorMessages(7);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_NOWHOPEN"));
    s = translateRmrErrorMessages(8);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_WHID"));
    s = translateRmrErrorMessages(9);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_OVERFLOW"));
    s = translateRmrErrorMessages(10);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_RETRY"));
    s = translateRmrErrorMessages(11);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_RCVFAILED"));
    s = translateRmrErrorMessages(12);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_TIMEOUT"));
    s = translateRmrErrorMessages(13);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_UNSET"));
    s = translateRmrErrorMessages(14);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_TRUNC"));
    s = translateRmrErrorMessages(15);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_INITFAILED"));
    s = translateRmrErrorMessages(16);
    EXPECT_THAT(s, HasSubstr("RMR_ERR_NOTSUPP"));
    s = translateRmrErrorMessages(17);
    EXPECT_THAT(s, HasSubstr("UNDOCUMENTED RMR_ERR"));
}

auto *peerInfo = (ConnectedCU_t *)calloc(1, sizeof(ConnectedCU_t));

TEST(sctp, TEST2) {
    struct epoll_event event;
    event.events = EPOLLIN;
    event.data.fd = 0;
    ConnectedCU_t data1;
    ConnectedCU_t *data = &data1;
    event.data.ptr = (void *)data;
    sctp_params_t sctp_ut_params;
    sctp_params_t* sctp = &sctp_ut_params;
    ReportingMessages_t reporting_msg;
    RmrMessagesBuffer_t rmrmessagesbuffer;
    handleEinprogressMessages(event,reporting_msg,rmrmessagesbuffer,sctp);
}

TEST(sctp, TEST3) {
    struct epoll_event event;
    event.events = EPOLLIN;
    event.data.fd = 0;
    ConnectedCU_t data1;
    ConnectedCU_t *data = &data1;
    event.data.ptr = (void *)data;
    sctp_params_t sctp_ut_params;
    sctp_params_t* sctp = &sctp_ut_params;
    Sctp_Map_t m1;
    sctp ->sctpMap =  &m1;
    ReportingMessages_t reporting_msg;
    RmrMessagesBuffer_t rmrmessagesbuffer;
    handlepoll_error(event,reporting_msg,rmrmessagesbuffer,sctp);
}

TEST(sctp, TEST4) {
    ConnectedCU_t cu;
    ConnectedCU_t* connected_cu = &cu;
    Sctp_Map_t m1;
    Sctp_Map_t *m = &m1;
    cleanHashEntry(connected_cu,m);
}

TEST(sctp, TEST5) {
    sctp_params_t sctp_ut_params;
    sctp_params_t* sctp = &sctp_ut_params;
    sctp->configFilePath.assign("/opt/e2/RIC-E2-TERMINATION/config");
    sctp->configFileName.assign("config.conf");
    handleConfigChange(sctp);
}

TEST(sctp, TEST6) {
    int epoll_fd = epoll_create1(0);
    ConnectedCU_t cu;
    ConnectedCU_t* peerinfo = &cu;
    Sctp_Map_t m1;
    Sctp_Map_t *m = &m1;
    modifyToEpoll(epoll_fd,peerinfo,2,m, (char*)"enodeb1",2);
}


/* TEST7 Begin: */
void init_memories(ReportingMessages_t &message, RmrMessagesBuffer_t &rmrMessageBuffer, sctp_params_t &sctp_ut_params) {
    message.peerInfo = peerInfo;
    message.peerInfo->sctpParams = &sctp_ut_params;
    snprintf(message.message.enodbName, strlen("Nokia_enb "), "%s", (char*)"Nokia_enb");
    
    sctp_ut_params.myIP = "1.2.3.4";
    sctp_ut_params.rmrPort = 38000;
    snprintf(sctp_ut_params.rmrAddress, sizeof(sctp_ut_params.rmrAddress), "%d", (int) (sctp_ut_params.rmrPort));
    
    rmrMessageBuffer.sendMessage = (rmr_mbuf_t*) malloc(sizeof(rmr_mbuf_t));
    rmrMessageBuffer.sendMessage->header = (uta_mhdr_t*) malloc(sizeof(uta_mhdr_t));
    rmrMessageBuffer.sendMessage->len = strlen("Saying Hello from NOKIA ");
    rmrMessageBuffer.sendMessage->payload = (unsigned char*)strdup("Saying Hello from NOKIA");
}

void delete_memories_initiatingMessage(E2AP_PDU_t *pdu, RmrMessagesBuffer_t &rmrMessageBuffer) {
    
    if( (pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array) &&
        (pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array[0]) ) {
        free(pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array[0]);
        pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array[0] = NULL;
    }
    if(pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array) {
        free(pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array);
        pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array = NULL;
    }
    pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.count = 0;
    pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.size = 0;
    
    if(rmrMessageBuffer.sendMessage && rmrMessageBuffer.sendMessage->header) {
        free(rmrMessageBuffer.sendMessage->header);
        rmrMessageBuffer.sendMessage->header = NULL;
    }
    if(rmrMessageBuffer.sendMessage->payload) {
        free(rmrMessageBuffer.sendMessage->payload);
        rmrMessageBuffer.sendMessage->payload = NULL;
    }
    if(rmrMessageBuffer.sendMessage) {
        free(rmrMessageBuffer.sendMessage);
        rmrMessageBuffer.sendMessage = NULL;
    }
    if(rmrMessageBuffer.rmrCtx) {
        rmr_close(rmrMessageBuffer.rmrCtx);
        rmrMessageBuffer.rmrCtx = NULL;
    }
}

void create_asnInitiatingReq_Procedure_RICserviceUpdate(E2AP_PDU_t *pdu, 
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_RICserviceUpdate */
    pdu->choice.initiatingMessage->procedureCode = ProcedureCode_id_RICserviceUpdate; 
    pdu->choice.initiatingMessage->value.present = InitiatingMessage__value_PR_RICserviceUpdate;
   
    int streamId = 0; 
    asnInitiatingRequest(pdu, sctpMap, message, rmrMessageBuffer,streamId);
    delete_memories_initiatingMessage(pdu, rmrMessageBuffer);
}

void create_asnInitiatingReq_Procedure_RICindication(E2AP_PDU_t *pdu, 
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {
    
    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_RICindication */
    pdu->choice.initiatingMessage->procedureCode = ProcedureCode_id_RICindication; 
    pdu->choice.initiatingMessage->value.present = InitiatingMessage__value_PR_RICindication;
    pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array = (RICindication_IEs**) malloc(1 * sizeof(RICindication_IEs*));
    pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.count = 1;
    pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.size = sizeof(RICindication_IEs);
    pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array[0] = (RICindication_IEs*) malloc(sizeof(RICindication_IEs));
    RICindication_IEs_t *ie = pdu->choice.initiatingMessage->value.choice.RICindication.protocolIEs.list.array[0];
    ie->id = ProtocolIE_ID_id_RICrequestID;
    ie->value.present = RICindication_IEs__value_PR_RICrequestID;
    ie->value.choice.RICrequestID.ricRequestorID = 12345;
    ie->value.choice.RICrequestID.ricInstanceID = 1;
    int streamId = 0;
    asnInitiatingRequest(pdu, sctpMap, message, rmrMessageBuffer,streamId);
    delete_memories_initiatingMessage(pdu, rmrMessageBuffer);
}

void create_asnInitiatingReq_Procedure_ErrorIndication(E2AP_PDU_t *pdu, 
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_ErrorIndication */
    pdu->choice.initiatingMessage->procedureCode = ProcedureCode_id_ErrorIndication; 
    pdu->choice.initiatingMessage->value.present = InitiatingMessage__value_PR_ErrorIndication;
    int streamId = 0;
    asnInitiatingRequest(pdu, sctpMap, message, rmrMessageBuffer, streamId);
    delete_memories_initiatingMessage(pdu, rmrMessageBuffer);
}

void create_asnInitiatingReq_Procedure_Reset(E2AP_PDU_t *pdu, 
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_Reset */
    pdu->choice.initiatingMessage->procedureCode = ProcedureCode_id_Reset; 
    pdu->choice.initiatingMessage->value.present = InitiatingMessage__value_PR_ResetRequest;
    int streamId =0;
    asnInitiatingRequest(pdu, sctpMap, message, rmrMessageBuffer,streamId);
    delete_memories_initiatingMessage(pdu, rmrMessageBuffer);
}

TEST(sctp, TEST7) {
    E2AP_PDU_t              pdu;
    Sctp_Map_t              *sctpMap = new Sctp_Map_t();
    ReportingMessages_t     message;
    RmrMessagesBuffer_t     rmrMessageBuffer;
    sctp_params_t           sctp_ut_params;

    pdu.present = E2AP_PDU_PR_initiatingMessage;
    pdu.choice.initiatingMessage = (InitiatingMessage*) malloc(sizeof(InitiatingMessage));
    memset( (void*)pdu.choice.initiatingMessage, 0, sizeof(pdu.choice.initiatingMessage));
    memset( (void*)&message, 0, sizeof(message));
    memset( (void*)&rmrMessageBuffer, 0, sizeof(rmrMessageBuffer));

    snprintf(sctp_ut_params.rmrAddress, strlen("tcp:4560 "), "%s", (char*)"tcp:4560");
    rmrMessageBuffer.rmrCtx = rmr_init(sctp_ut_params.rmrAddress, RECEIVE_XAPP_BUFFER_SIZE, 0x01);

    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_RICserviceUpdate */
    create_asnInitiatingReq_Procedure_RICserviceUpdate(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_RICindication */
    create_asnInitiatingReq_Procedure_RICindication(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_ErrorIndication */
    create_asnInitiatingReq_Procedure_ErrorIndication(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_initiatingMessage and procedure code as: ProcedureCode_id_Reset */
    create_asnInitiatingReq_Procedure_Reset(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* For Procedure's Default case. */
    pdu.choice.initiatingMessage->procedureCode = ((ProcedureCode_t)100);
    int streamId =0;
    asnInitiatingRequest(&pdu, sctpMap, message, rmrMessageBuffer,streamId);
    
    if(pdu.choice.initiatingMessage) {
        free(pdu.choice.initiatingMessage);
        pdu.choice.initiatingMessage = NULL;
    }
    if(sctpMap) {
        delete sctpMap;
        sctpMap = NULL;
    }
}

/* TEST8 Begin: */
void delete_memories_successfulOutcome(E2AP_PDU_t *pdu, RmrMessagesBuffer_t &rmrMessageBuffer) {

    if( (pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array) && 
        (pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array[0]) ) {
        free(pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array[0]);
        pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array[0] = NULL;
    }
    if(pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array) {
        free(pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array);
        pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array = NULL;
    }
    pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.count = 0;
    pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.size = 0;

    if(rmrMessageBuffer.sendMessage && rmrMessageBuffer.sendMessage->header) {
        free(rmrMessageBuffer.sendMessage->header);
        rmrMessageBuffer.sendMessage->header = NULL;
    }
    if(rmrMessageBuffer.sendMessage->payload) {
        free(rmrMessageBuffer.sendMessage->payload);
        rmrMessageBuffer.sendMessage->payload = NULL;
    }
    if(rmrMessageBuffer.sendMessage) {
        free(rmrMessageBuffer.sendMessage);
        rmrMessageBuffer.sendMessage = NULL;
    }
}


void create_asnSuccessfulMsg_Procedure_Reset(E2AP_PDU_t *pdu,
        Sctp_Map_t *sctpMap,
        ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {
    
    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_Reset */
    pdu->choice.successfulOutcome->procedureCode = ProcedureCode_id_Reset; 
    asnSuccessfulMsg(pdu, sctpMap, message, rmrMessageBuffer);
    delete_memories_successfulOutcome(pdu, rmrMessageBuffer);
}

void create_asnSuccessfulMsg_Procedure_RICcontrol(E2AP_PDU_t *pdu,
        Sctp_Map_t *sctpMap,
        ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    
    pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array = (RICcontrolAcknowledge_IEs_t**) malloc(1 * sizeof(RICcontrolAcknowledge_IEs_t*));
    pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.count = 1;
    pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.size = sizeof(RICcontrolAcknowledge_IEs_t);
    pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array[0] = (RICcontrolAcknowledge_IEs_t*) malloc(sizeof(RICcontrolAcknowledge_IEs_t));
    RICcontrolAcknowledge_IEs_t *ie = pdu->choice.successfulOutcome->value.choice.RICcontrolAcknowledge.protocolIEs.list.array[0];

    ie->id = ProtocolIE_ID_id_RICrequestID;
    ie->value.present = RICcontrolAcknowledge_IEs__value_PR_RICrequestID;
    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_RICcontrol */
    pdu->choice.successfulOutcome->procedureCode = ProcedureCode_id_RICcontrol; 
    asnSuccessfulMsg(pdu, sctpMap, message, rmrMessageBuffer);
    delete_memories_successfulOutcome(pdu, rmrMessageBuffer);
}

void create_asnSuccessfulMsg_Procedure_RICsubscription(E2AP_PDU_t *pdu,
        Sctp_Map_t *sctpMap,
        ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_RICsubscription */
    pdu->choice.successfulOutcome->procedureCode = ProcedureCode_id_RICsubscription; 
    asnSuccessfulMsg(pdu, sctpMap, message, rmrMessageBuffer);
    delete_memories_successfulOutcome(pdu, rmrMessageBuffer);
}

void create_asnSuccessfulMsg_Procedure_RICsubscriptionDelete(E2AP_PDU_t *pdu,
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_RICsubscriptionDelete */
    pdu->choice.successfulOutcome->procedureCode = ProcedureCode_id_RICsubscriptionDelete; 
    asnSuccessfulMsg(pdu, sctpMap, message, rmrMessageBuffer);
    delete_memories_successfulOutcome(pdu, rmrMessageBuffer);
}

TEST(sctp, TEST8) {
    E2AP_PDU_t              pdu;
    Sctp_Map_t              *sctpMap = new Sctp_Map_t();
    ReportingMessages_t     message;
    RmrMessagesBuffer_t     rmrMessageBuffer;
    sctp_params_t           sctp_ut_params;

    pdu.present = E2AP_PDU_PR_successfulOutcome;
    pdu.choice.successfulOutcome = (SuccessfulOutcome*) malloc(sizeof(SuccessfulOutcome));
    memset( (void*)pdu.choice.successfulOutcome, 0, sizeof(pdu.choice.successfulOutcome));
    memset( (void*)&message, 0, sizeof(message));
    memset( (void*)&rmrMessageBuffer, 0, sizeof(rmrMessageBuffer));

    snprintf(sctp_ut_params.rmrAddress, strlen("127.0.0.1 "), "%s", (char*)"127.0.0.1");

    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_Reset */
    create_asnSuccessfulMsg_Procedure_Reset(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_RICcontrol */
    create_asnSuccessfulMsg_Procedure_RICcontrol(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_RICsubscription */
    create_asnSuccessfulMsg_Procedure_RICsubscription(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_RICsubscriptionDelete */
    create_asnSuccessfulMsg_Procedure_RICsubscriptionDelete(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* For Procedure's Default case. */
    pdu.choice.successfulOutcome->procedureCode = ((ProcedureCode_t)100); 
    asnSuccessfulMsg(&pdu, sctpMap, message, rmrMessageBuffer);
    
    if(pdu.choice.successfulOutcome) {
        free(pdu.choice.successfulOutcome);
        pdu.choice.successfulOutcome = NULL;
    }
    if(sctpMap) {
        delete sctpMap;
        sctpMap = NULL;
    }
}

/* TEST9 Begin: */
void delete_memories_unsuccessfulOutcome(E2AP_PDU_t *pdu, RmrMessagesBuffer_t &rmrMessageBuffer) {

    if( (pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array) &&
        (pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array[0]) ) {
        free(pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array[0]);
        pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array[0] = NULL;
    }
    if(pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array) {
        free(pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array);
        pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array = NULL;
    }
    pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.count = 0;
    pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.size = 0;

    if(rmrMessageBuffer.sendMessage && rmrMessageBuffer.sendMessage->header) {
        free(rmrMessageBuffer.sendMessage->header);
        rmrMessageBuffer.sendMessage->header = NULL;
    }
    if(rmrMessageBuffer.sendMessage->payload) {
        free(rmrMessageBuffer.sendMessage->payload);
        rmrMessageBuffer.sendMessage->payload = NULL;
    }
    if(rmrMessageBuffer.sendMessage) {
        free(rmrMessageBuffer.sendMessage);
        rmrMessageBuffer.sendMessage = NULL;
    }
}

void create_asnUnSuccsesfulMsg_Procedure_RICcontrol(E2AP_PDU_t *pdu, 
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    
    pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array = (RICcontrolFailure_IEs_t**) malloc(1*sizeof(RICcontrolFailure_IEs_t*));
    pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.count = 1;
    pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.size = sizeof(RICcontrolFailure_IEs_t);
    pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array[0] = (RICcontrolFailure_IEs_t*) malloc(sizeof(RICcontrolFailure_IEs_t));
    RICcontrolFailure_IEs_t *ie = pdu->choice.unsuccessfulOutcome->value.choice.RICcontrolFailure.protocolIEs.list.array[0];

    ie->id = ProtocolIE_ID_id_RICrequestID;
    ie->value.present = RICcontrolFailure_IEs__value_PR_RICrequestID;
    /* Sending E2AP_PDU_PR_unsuccessfulOutcome and procedure code as: ProcedureCode_id_RICcontrol */
    pdu->choice.unsuccessfulOutcome->procedureCode = ProcedureCode_id_RICcontrol;
    asnUnSuccsesfulMsg(pdu, sctpMap, message, rmrMessageBuffer);
    delete_memories_unsuccessfulOutcome(pdu, rmrMessageBuffer);
}

void create_asnUnSuccsesfulMsg_Procedure_RICsubscription(E2AP_PDU_t *pdu, 
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {

    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_unsuccessfulOutcome and procedure code as: ProcedureCode_id_RICsubscription */
    pdu->choice.unsuccessfulOutcome->procedureCode = ProcedureCode_id_RICsubscription; 
    asnUnSuccsesfulMsg(pdu, sctpMap, message, rmrMessageBuffer);
    delete_memories_unsuccessfulOutcome(pdu, rmrMessageBuffer);
}

void create_asnUnSuccsesfulMsg_Procedure_RICsubscriptionDelete(E2AP_PDU_t *pdu, 
        Sctp_Map_t *sctpMap, 
        ReportingMessages_t &message, 
        RmrMessagesBuffer_t &rmrMessageBuffer,
        sctp_params_t &sctp_ut_params) {
    
    init_memories(message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_unsuccessfulOutcome and procedure code as: ProcedureCode_id_RICsubscriptionDelete */
    pdu->choice.unsuccessfulOutcome->procedureCode = ProcedureCode_id_RICsubscriptionDelete; 
    asnUnSuccsesfulMsg(pdu, sctpMap, message, rmrMessageBuffer);
    delete_memories_unsuccessfulOutcome(pdu, rmrMessageBuffer);
}

TEST(sctp, TEST9) {
    E2AP_PDU_t              pdu;
    Sctp_Map_t              *sctpMap = new Sctp_Map_t();
    ReportingMessages_t     message;
    RmrMessagesBuffer_t     rmrMessageBuffer;
    sctp_params_t           sctp_ut_params;

    pdu.present = E2AP_PDU_PR_unsuccessfulOutcome;
    pdu.choice.unsuccessfulOutcome = (UnsuccessfulOutcome*) malloc(sizeof(UnsuccessfulOutcome));
    memset( (void*)pdu.choice.unsuccessfulOutcome, 0, sizeof(pdu.choice.unsuccessfulOutcome));
    memset( (void*)&message, 0, sizeof(message));
    memset( (void*)&rmrMessageBuffer, 0, sizeof(rmrMessageBuffer));

    snprintf(sctp_ut_params.rmrAddress, strlen("127.0.0.1 "), "%s", (char*)"127.0.0.1");

    /* Sending E2AP_PDU_PR_unsuccessfulOutcome and procedure code as: ProcedureCode_id_RICcontrol */
    create_asnUnSuccsesfulMsg_Procedure_RICcontrol(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_unsuccessfulOutcome and procedure code as: ProcedureCode_id_RICsubscription */
    create_asnUnSuccsesfulMsg_Procedure_RICsubscription(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* Sending E2AP_PDU_PR_unsuccessfulOutcome and procedure code as: ProcedureCode_id_RICsubscriptionDelete */
    create_asnUnSuccsesfulMsg_Procedure_RICsubscriptionDelete(&pdu, sctpMap, message, rmrMessageBuffer, sctp_ut_params);
    /* For Procedure's Default case. */
    pdu.choice.unsuccessfulOutcome->procedureCode = ((ProcedureCode_t)100); 
    asnUnSuccsesfulMsg(&pdu, sctpMap, message, rmrMessageBuffer);
    
    if(pdu.choice.unsuccessfulOutcome) {
        free(pdu.choice.unsuccessfulOutcome);
        pdu.choice.unsuccessfulOutcome = NULL;
    }
    if(sctpMap) {
        delete sctpMap;
        sctpMap = NULL;
    }
}


TEST(sctp, TEST10) {
    int epoll_fd = epoll_create1(0);
    ConnectedCU_t cu;
    ConnectedCU_t* peerinfo = &cu;
    Sctp_Map_t m1;
    Sctp_Map_t *m = &m1;
    addToEpoll(epoll_fd, peerinfo, 2, m, (char*)"enodeb1", 0);
}

TEST(sctp, TEST11) {
    sctp_params_t sctpParams;
    int argc = 5;
    char **argv = (char**) calloc(argc, sizeof(char*));
    argv[0] = (char*) malloc(40 * sizeof(char));
    argv[1] = (char*) malloc(40 * sizeof(char));
    argv[2] = (char*) malloc(40 * sizeof(char));
    argv[3] = (char*) malloc(40 * sizeof(char));
    argv[4] = (char*) malloc(40 * sizeof(char));
    snprintf(argv[0], strlen("./e2 "), "%s", (char*)"./e2");
    snprintf(argv[1], strlen("-p "), "%s", (char*)"-p");
    snprintf(argv[2], strlen("/opt/e2/RIC-E2-TERMINATION/config "), "%s", (char*)"/opt/e2/RIC-E2-TERMINATION/config");
    snprintf(argv[3], strlen("-f "), "%s", (char*)"-f");
    snprintf(argv[4], strlen("config.conf "), "%s", (char*)"config.conf");
    
    auto result = parse(argc, argv, sctpParams);
    sctpParams.podName.assign("E2TermAlpha_pod");
    sctpParams.sctpMap = new mapWrapper();
    sctpParams.epoll_fd = epoll_create1(0);
    buildConfiguration(sctpParams);
    // getRmrContext(sctpParams);
    buildInotify(sctpParams);
    buildListeningPort(sctpParams);
    listener(&sctpParams);

    if(sctpParams.sctpMap) {
        delete sctpParams.sctpMap;
        sctpParams.sctpMap = NULL;
    }
    if(argv) {
        free(argv[0]);
        argv[0] = NULL;
        free(argv[1]);
        argv[1] = NULL;
        free(argv[2]);
        argv[2] = NULL;
        free(argv[3]);
        argv[3] = NULL;
        free(argv[4]);
        argv[4] = NULL;
        free(argv);
        argv=NULL;
    }
}

TEST(sctp, TEST12) {
    ReportingMessages_t     reporting_msg;
    Sctp_Map_t              *sctpMap = new Sctp_Map_t();
    sendSctpMsg(peerInfo, reporting_msg, sctpMap);

    if(sctpMap) {
        delete sctpMap;
        sctpMap = NULL;
    }
}

/*TEST13 Begin*/

void inti_buffers_rcv(ReportingMessages_t &message, RmrMessagesBuffer_t &rmrMessageBuffer) {
    message.peerInfo = peerInfo;
    snprintf(message.message.enodbName, strlen("Nokia_enb "), "%s", (char*)"Nokia_enb");

    rmrMessageBuffer.rcvMessage = (rmr_mbuf_t*) malloc(sizeof(rmr_mbuf_t));
    rmrMessageBuffer.rcvMessage->header = (uta_mhdr_t*) malloc(sizeof(uta_mhdr_t));
    rmrMessageBuffer.rcvMessage->len = strlen("Saying Hello from Ramji ");
    rmrMessageBuffer.rcvMessage->payload = (unsigned char*)strdup("Saying Hello from Ramji");

    rmrMessageBuffer.sendMessage = (rmr_mbuf_t*) malloc(sizeof(rmr_mbuf_t));
    rmrMessageBuffer.sendMessage->header = (uta_mhdr_t*) malloc(sizeof(uta_mhdr_t));
    rmrMessageBuffer.sendMessage->len = strlen("Saying Hello from Ramji ");
    rmrMessageBuffer.sendMessage->payload = (unsigned char*)strdup("Saying Hello from Ramji");

}

void delete_memories_rcv(RmrMessagesBuffer_t &rmrMessageBuffer) {

    if(rmrMessageBuffer.rcvMessage && rmrMessageBuffer.rcvMessage->header) {
        free(rmrMessageBuffer.rcvMessage->header);
        rmrMessageBuffer.rcvMessage->header = NULL;
    }
    if(rmrMessageBuffer.rcvMessage) {
        free(rmrMessageBuffer.rcvMessage);
        rmrMessageBuffer.rcvMessage = NULL;
    }
    if(rmrMessageBuffer.sendMessage && rmrMessageBuffer.sendMessage->header) {
        free(rmrMessageBuffer.sendMessage->header);
        rmrMessageBuffer.sendMessage->header = NULL;
    }
    if(rmrMessageBuffer.sendMessage) {
        free(rmrMessageBuffer.sendMessage);
        rmrMessageBuffer.sendMessage = NULL;
    }

}

void create_receiveXappMessages_RIC_ERROR_INDICATION(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_ERROR_INDICATION;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);

}
void create_receiveXappMessages_RIC_E2_SETUP_FAILURE(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_E2_SETUP_FAILURE;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_SUB_REQ(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_SUB_REQ;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_CONTROL_REQ(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_CONTROL_REQ;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_SERVICE_QUERY(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_SERVICE_QUERY;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_SERVICE_UPDATE_ACK(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_SERVICE_UPDATE_ACK;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_SERVICE_UPDATE_FAILURE(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_SERVICE_UPDATE_FAILURE;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}


void create_receiveXappMessages_RIC_E2_RESET_REQ(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_E2_RESET_REQ;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_E2_RESET_RESP(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_E2_RESET_RESP;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_SCTP_CLEAR_ALL(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_SCTP_CLEAR_ALL;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_HEALTH_CHECK_REQ(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = RIC_HEALTH_CHECK_REQ;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_E2_TERM_KEEP_ALIVE_REQ(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = E2_TERM_KEEP_ALIVE_REQ;
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);
}

void create_receiveXappMessages_RIC_SUB_DEL_REQ(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
        inti_buffers_rcv(message, rmrMessageBuffer);
        rmrMessageBuffer.rcvMessage->mtype = RIC_SUB_DEL_REQ;
        receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
        delete_memories_rcv(rmrMessageBuffer);
}        

void create_receiveXappMessages_RIC_E2_SETUP_RESP(Sctp_Map_t *sctpMap, ReportingMessages_t &message,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
        inti_buffers_rcv(message, rmrMessageBuffer);
        rmrMessageBuffer.rcvMessage->mtype = RIC_E2_SETUP_RESP;
        receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
        delete_memories_rcv(rmrMessageBuffer);
}

TEST(sctp, TEST13) {
    Sctp_Map_t *sctpMap = new Sctp_Map_t();
    ReportingMessages_t message;
    RmrMessagesBuffer_t rmrMessageBuffer;

    memset( (void*)&message, 0, sizeof(message));
    memset( (void*)&rmrMessageBuffer, 0, sizeof(rmrMessageBuffer));

    // Sending E2AP_PDU_PR_successfulOutcome and procedure code as: ProcedureCode_id_Reset
    create_receiveXappMessages_RIC_E2_SETUP_FAILURE(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_ERROR_INDICATION(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_SUB_REQ(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_CONTROL_REQ(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_HEALTH_CHECK_REQ(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_E2_TERM_KEEP_ALIVE_REQ(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_SCTP_CLEAR_ALL(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_E2_RESET_RESP(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_E2_RESET_REQ(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_SERVICE_UPDATE_FAILURE(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_SERVICE_UPDATE_ACK(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_SERVICE_QUERY(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_SUB_DEL_REQ(sctpMap, message, rmrMessageBuffer);
    create_receiveXappMessages_RIC_E2_SETUP_RESP(sctpMap, message, rmrMessageBuffer);

    inti_buffers_rcv(message, rmrMessageBuffer);
    rmrMessageBuffer.rcvMessage->mtype = 52345; /*Dummy Integer Value for default case*/
    receiveXappMessages(sctpMap, rmrMessageBuffer, message.message.time);
    delete_memories_rcv(rmrMessageBuffer);

    if(sctpMap) {
        delete sctpMap;
        sctpMap = NULL;
    }
}
/* TEST 14 Begin: */
void receiveDataFromSctp_asnSuccessfulMsg_Procedure_Default( E2AP_PDU_t *pdu,
        RmrMessagesBuffer_t &rmrMessageBuffer) {
    pdu->present = E2AP_PDU_PR_successfulOutcome;
    pdu->choice.successfulOutcome = (SuccessfulOutcome*) malloc(sizeof(SuccessfulOutcome));
    memset( (void*)pdu->choice.successfulOutcome, 0, sizeof(pdu->choice.successfulOutcome));
    pdu->choice.successfulOutcome->procedureCode = ((ProcedureCode_t)100);

    rmrMessageBuffer.sendMessage->tp_buf = pdu;
}

void delete_pdu_memories(E2AP_PDU_t *pdu) {
    if(pdu->choice.successfulOutcome) {
        free(pdu->choice.successfulOutcome);
        pdu->choice.successfulOutcome = NULL;
    }
    if(pdu->choice.unsuccessfulOutcome) {
        free(pdu->choice.unsuccessfulOutcome);
        pdu->choice.unsuccessfulOutcome = NULL;
    }
}

/*void receiveDataFromSctp_asnUnSuccsesfulMsg_Procedure_Default(E2AP_PDU_t  *pdu,
        RmrMessagesBuffer_t     &rmrMessageBuffer) {
    pdu->present = E2AP_PDU_PR_unsuccessfulOutcome;
    pdu->choice.unsuccessfulOutcome = (UnsuccessfulOutcome*) malloc(sizeof(UnsuccessfulOutcome));
    memset( (void*)pdu->choice.unsuccessfulOutcome, 0, sizeof(pdu->choice.unsuccessfulOutcome));
    pdu->choice.unsuccessfulOutcome->procedureCode = ((ProcedureCode_t)100);

    rmrMessageBuffer.sendMessage->tp_buf = pdu;
}*/

TEST(sctp, TEST14) {
    E2AP_PDU_t              pdu;
    struct epoll_event      events;
    Sctp_Map_t              *sctpMap  = new Sctp_Map_t();
    int                     numOfMessages=0;
    RmrMessagesBuffer_t     rmrMessageBuffer;
    struct timespec         ts;
    ConnectedCU_t           *peerInfo = (ConnectedCU_t *) calloc(1, sizeof(ConnectedCU_t));
    events.data.ptr = peerInfo;
    snprintf(peerInfo->enodbName, strlen("Nokia_enb "), "%s", (char*)"Nokia_enb");

    rmrMessageBuffer.sendMessage = (rmr_mbuf_t*) malloc(sizeof(rmr_mbuf_t));
    rmrMessageBuffer.sendMessage->header = (uta_mhdr_t*) malloc(sizeof(uta_mhdr_t));
    rmrMessageBuffer.sendMessage->len = strlen("Saying Hello from NOKIA ");
    rmrMessageBuffer.sendMessage->payload = (unsigned char*)strdup("Saying Hello from NOKIA");

    clock_gettime(CLOCK_MONOTONIC, &ts);

    receiveDataFromSctp_asnSuccessfulMsg_Procedure_Default(&pdu, rmrMessageBuffer);
    receiveDataFromSctp(&events, sctpMap, numOfMessages, rmrMessageBuffer, ts);
    delete_pdu_memories(&pdu);

    if(rmrMessageBuffer.sendMessage->payload) {
        free(rmrMessageBuffer.sendMessage->payload);
        rmrMessageBuffer.sendMessage->payload = NULL;
    }

    /*rmrMessageBuffer.sendMessage->payload = (unsigned char*)strdup("Saying Hello from NOKIA");
    receiveDataFromSctp_asnUnSuccsesfulMsg_Procedure_Default(&pdu, rmrMessageBuffer);
    receiveDataFromSctp(&events, sctpMap, numOfMessages, rmrMessageBuffer, ts);
    delete_pdu_memories(&pdu);

    if(rmrMessageBuffer.sendMessage->payload) {
        free(rmrMessageBuffer.sendMessage->payload);
        rmrMessageBuffer.sendMessage->payload = NULL;
    }*/
    if(rmrMessageBuffer.sendMessage->header) {
        free(rmrMessageBuffer.sendMessage->header);
        rmrMessageBuffer.sendMessage->header = NULL;
    }
    if(rmrMessageBuffer.sendMessage) {
        free(rmrMessageBuffer.sendMessage);
    }
    if(peerInfo) {
        free(peerInfo);
        peerInfo = NULL;
    }
    if(sctpMap) {
        delete sctpMap;
        sctpMap = NULL;
    }
}

int main(int argc, char **argv) {

   testing::InitGoogleTest(&argc, argv);
   return RUN_ALL_TESTS();

}
