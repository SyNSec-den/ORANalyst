package control

const MAX_SUBSCRIPTION_ATTEMPTS = 100
type APIResponse struct {
Cellids []string `json:"Cellids"`
PlmnIds  []string `json:"PlmnIds"`
}
type RanFunctions struct {
	RanFunctionId         int
	RanFunctionDefinition string
	RanFunctionRevision   int
	RanFunctionOid        string
}

type GlobalNbId struct {
	PlmnId string
	NbId   string
}
type E2nodeComponentInterfaceTypeE1 struct {
}
type E2nodeComponentInterfaceTypeXn struct {
}
type E2nodeComponentInterfaceTypeF1 struct {
}

type NodeConfigs struct {
	E2nodeComponentInterfaceTypeE1 E2nodeComponentInterfaceTypeE1 `json:e2nodeComponentInterfaceTypeE1",omitempty"`
	E2nodeComponentInterfaceTypeXn E2nodeComponentInterfaceTypeXn `json:e2nodeComponentInterfaceTypeXn",omitempty"`
	E2nodeComponentInterfaceTypeF1 E2nodeComponentInterfaceTypeF1 `json:e2nodeComponentInterfaceTypeF1",omitempty"`
	E2nodeComponentInterfaceType   string
	E2nodeComponentRequestPart     string
	E2nodeComponentResponsePart    string `json:e2nodeComponentResponsePart",omitempty"`
}
type Gnb struct {
	RanFunctions []RanFunctions
	GnbType      string
	NodeConfigs  []NodeConfigs
}

/*
	type E2nodeComponentInterfaceType struct {
		E2nodeComponentInterfaceType string
	}

	type E2nodeComponentRequestPart struct {
		E2nodeComponentRequestPart string
	}
*/

type E2mgrResponse struct {
	RanName                      string
	ConnectionStatus             string
	GlobalNbId                   GlobalNbId
	NodeType                     string
	Gnb                          Gnb
	AssociatedE2tInstanceAddress string `json:associatedE2tInstanceAddress",omitempty"`
	SetupFromNetwork             bool
	StatusUpdateTimeStamp        string
}

type DecodedIndicationMessage struct {
	RequestID             int32
	RequestSequenceNumber int32
	FuncID                int32
	ActionID              int32
	IndSN                 int32
	IndType               int32
	IndHeader             []byte
	IndHeaderLength       int32
	IndMessage            []byte
	IndMessageLength      int32
	CallProcessID         []byte
	CallProcessIDLength   int32
}

type CauseItemType struct {
	CauseType int32
	CauseID   int32
}

type ActionAdmittedListType struct {
	ActionID [16]int32
	Count    int
}

type ActionNotAdmittedListType struct {
	ActionID [16]int32
	Cause    [16]CauseItemType
	Count    int
}

type DecodedSubscriptionResponseMessage struct {
	RequestID             int32
	RequestSequenceNumber int32
	FuncID                int32
	ActionAdmittedList    ActionAdmittedListType
	ActionNotAdmittedList ActionNotAdmittedListType
}

type IntPair64 struct {
	DL int64
	UL int64
}

type OctetString struct {
	Buf  []byte
	Size int
}

type Integer OctetString

type PrintableString OctetString

type ActionDefinition OctetString

type BitString struct {
	Buf        []byte
	Size       int
	BitsUnused int
}

type SubsequentAction struct {
	IsValid              int
	SubsequentActionType int64
	TimeToWait           int64
}

type GNBID BitString

type GlobalgNBIDType struct {
	PlmnID    OctetString
	GnbIDType int
	GnbID     interface{}
}

type GlobalKPMnodegNBIDType struct {
	GlobalgNBID GlobalgNBIDType
	GnbCUUPID   *Integer
	GnbDUID     *Integer
}

type ENGNBID BitString

type GlobalKPMnodeengNBIDType struct {
	PlmnID    OctetString
	GnbIDType int
	GnbID     interface{}
}

type NGENBID_Macro BitString

type NGENBID_ShortMacro BitString

type NGENBID_LongMacro BitString

type GlobalKPMnodengeNBIDType struct {
	PlmnID    OctetString
	EnbIDType int
	EnbID     interface{}
}

type ENBID_Macro BitString

type ENBID_Home BitString

type ENBID_ShortMacro BitString

type ENBID_LongMacro BitString

type GlobalKPMnodeeNBIDType struct {
	PlmnID    OctetString
	EnbIDType int
	EnbID     interface{}
}

type NRCGIType struct {
	PlmnID   OctetString
	NRCellID BitString
}

type SliceIDType struct {
	SST OctetString
	SD  *OctetString
}

type GNB_DU_Name PrintableString

type GNB_CU_CP_Name PrintableString

type GNB_CU_UP_Name PrintableString

type IndicationHeaderFormat1 struct {
	GlobalKPMnodeIDType int32
	GlobalKPMnodeID     interface{}
	NRCGI               *NRCGIType
	PlmnID              *OctetString
	SliceID             *SliceIDType
	FiveQI              int64
	Qci                 int64
	UeMessageType       int32
	GnbDUID             *Integer
	GnbNameType         int32
	GnbName             interface{}
	GlobalgNBID         *GlobalgNBIDType
}

type IndicationHeader struct {
	IndHdrType int32
	IndHdr     interface{}
}

type FQIPERSlicesPerPlmnPerCellType struct {
	FiveQI   int64
	PrbUsage IntPair64
}

type SlicePerPlmnPerCellType struct {
	SliceID                         SliceIDType
	FQIPERSlicesPerPlmnPerCells     [64]FQIPERSlicesPerPlmnPerCellType
	FQIPERSlicesPerPlmnPerCellCount int
}

type DUPM5GCContainerType struct {
	SlicePerPlmnPerCells     [1024]SlicePerPlmnPerCellType
	SlicePerPlmnPerCellCount int
}

type DUPMEPCPerQCIReportType struct {
	QCI      int64
	PrbUsage IntPair64
}

type DUPMEPCContainerType struct {
	PerQCIReports     [256]DUPMEPCPerQCIReportType
	PerQCIReportCount int
}

type ServedPlmnPerCellType struct {
	PlmnID  OctetString
	DUPM5GC *DUPM5GCContainerType
	DUPMEPC *DUPMEPCContainerType
}

type CellResourceReportType struct {
	NRCGI                  NRCGIType
	TotalofAvailablePRBs   IntPair64
	ServedPlmnPerCells     [12]ServedPlmnPerCellType
	ServedPlmnPerCellCount int
}

type ODUPFContainerType struct {
	CellResourceReports     [512]CellResourceReportType
	CellResourceReportCount int
}

type CUCPResourceStatusType struct {
	NumberOfActiveUEs int64
}

type OCUCPPFContainerType struct {
	GNBCUCPName        *PrintableString
	CUCPResourceStatus CUCPResourceStatusType
}

type FQIPERSlicesPerPlmnType struct {
	FiveQI      int64
	PDCPBytesDL *Integer
	PDCPBytesUL *Integer
}

type SliceToReportType struct {
	SliceID                  SliceIDType
	FQIPERSlicesPerPlmns     [64]FQIPERSlicesPerPlmnType
	FQIPERSlicesPerPlmnCount int
}

type CUUPPM5GCType struct {
	SliceToReports     [1024]SliceToReportType
	SliceToReportCount int
}

type CUUPPMEPCPerQCIReportType struct {
	QCI         int64
	PDCPBytesDL *Integer
	PDCPBytesUL *Integer
}

type CUUPPMEPCType struct {
	CUUPPMEPCPerQCIReports     [256]CUUPPMEPCPerQCIReportType
	CUUPPMEPCPerQCIReportCount int
}

type CUUPPlmnType struct {
	PlmnID    OctetString
	CUUPPM5GC *CUUPPM5GCType
	CUUPPMEPC *CUUPPMEPCType
}

type CUUPMeasurementContainerType struct {
	CUUPPlmns     [12]CUUPPlmnType
	CUUPPlmnCount int
}

type CUUPPFContainerItemType struct {
	InterfaceType    int64
	OCUUPPMContainer CUUPMeasurementContainerType
}

type OCUUPPFContainerType struct {
	GNBCUUPName              *PrintableString
	CUUPPFContainerItems     [3]CUUPPFContainerItemType
	CUUPPFContainerItemCount int
}

type DUUsageReportUeResourceReportItemType struct {
	CRNTI      Integer
	PRBUsageDL int64
	PRBUsageUL int64
}

type DUUsageReportCellResourceReportItemType struct {
	NRCGI                     NRCGIType
	UeResourceReportItems     [32]DUUsageReportUeResourceReportItemType
	UeResourceReportItemCount int
}

type DUUsageReportType struct {
	CellResourceReportItems     [512]DUUsageReportCellResourceReportItemType
	CellResourceReportItemCount int
}

type CUCPUsageReportUeResourceReportItemType struct {
	CRNTI          Integer
	ServingCellRF  *OctetString
	NeighborCellRF *OctetString
}

type CUCPUsageReportCellResourceReportItemType struct {
	NRCGI                     NRCGIType
	UeResourceReportItems     [32]CUCPUsageReportUeResourceReportItemType
	UeResourceReportItemCount int
}

type CUCPUsageReportType struct {
	CellResourceReportItems     [16384]CUCPUsageReportCellResourceReportItemType
	CellResourceReportItemCount int
}

type CUUPUsageReportUeResourceReportItemType struct {
	CRNTI       Integer
	PDCPBytesDL *Integer
	PDCPBytesUL *Integer
}

type CUUPUsageReportCellResourceReportItemType struct {
	NRCGI                     NRCGIType
	UeResourceReportItems     [32]CUUPUsageReportUeResourceReportItemType
	UeResourceReportItemCount int
}

type CUUPUsageReportType struct {
	CellResourceReportItems     [512]CUUPUsageReportCellResourceReportItemType
	CellResourceReportItemCount int
}

type PFContainerType struct {
	ContainerType int32
	Container     interface{}
}

type RANContainerType struct {
	Timestamp     OctetString
	ContainerType int32
	Container     interface{}
}

type PMContainerType struct {
	PFContainer  *PFContainerType
	RANContainer *RANContainerType
}

type IndicationMessageFormat1 struct {
	PMContainers     [8]PMContainerType
	PMContainerCount int
}

type IndicationMessage struct {
	StyleType  int64
	IndMsgType int32
	IndMsg     interface{}
}

type Timestamp struct {
	TVsec  int64 `json:"tv_sec"`
	TVnsec int64 `json:"tv_nsec"`
}
//VIAVI 1.4.0 E2TEST cell metrics
/*
type CellMetricsEntry struct {
NR_CGI 	interface{}
NR_PCI 	interface{}
GNB_DU_ID 	interface{}
DRB_UEThpDl 	interface{}
DRB_UEThpUl 	interface{}
QosFlow_PdcpPduVolumeUl 	interface{}
QosFlow_PdcpPduVolumeDl 	interface{}
RRC_ConnMean 	interface{}
RRC_ConnMax 	interface{}
RRC_InactiveConnMean 	interface{}
RRC_InactiveConnMax 	interface{}
RRU_PrbAvailDl 	interface{}
RRU_PrbAvailUl 	interface{}
RRU_PrbUsedDl 	interface{}
RRU_PrbUsedUl 	interface{}
TB_TotNbrDlInitial_Qpsk 	interface{}
TB_TotNbrDlInitial_16Qam 	interface{}
TB_TotNbrDlInitial_64Qam 	interface{}
TB_TotNbrDlInitial_256Qam 	interface{}
TB_IntialErrNbrDl 	interface{}
TB_IntialErrNbrDl_Qpsk 	interface{}
TB_IntialErrNbrDl_16Qam 	interface{}
TB_IntialErrNbrDl_64Qam 	interface{}
TB_IntialErrNbrDl_256Qam 	interface{}
TB_TotNbrDl_X 	interface{}
TB_ErrToltalNbrDl_X 	interface{}
TB_ResidualErrNbrDl 	interface{}
TB_TotNbrUlInit 	interface{}
TB_TotNbrUlInit_Qpsk 	interface{}
TB_TotNbrUlInit_16Qam 	interface{}
TB_TotNbrUlInit_64Qam 	interface{}
TB_TotNbrUlInit_256Qam 	interface{}
TB_ErrNbrUlInitial 	interface{}
TB_ErrNbrUlInitial_Qpsk 	interface{}
TB_ErrNbrUlInitial_16Qam 	interface{}
TB_ErrNbrUlInitial_64Qam 	interface{}
TB_ErrNbrUlInitial_256Qam 	interface{}
TB_TotNbrUl_X 	interface{}
TB_ErrToltalNbrUl_X 	interface{}
TB_ResidualErrNbrUl 	interface{}
Viavi_Geo_x 	interface{}
Viavi_Geo_y 	interface{}
Viavi_Geo_z 	interface{}

}*/

//VIAVI 1.4.0 simulation cell metrics
type CellMetricsEntry struct {
        DRB_UEThpDl                interface{}
        DRB_UEThpUl                interface{}
	PEE_AvgPower		   interface{}
	PEE_Energy		   interface{}
        QosFlow_TotPdcpPduVolumeDl interface{}
        QosFlow_TotPdcpPduVolumeUl interface{}
        RRC_ConnMax                interface{}
        RRC_ConnMean               interface{}
        RRU_PrbAvailDl             interface{}
        RRU_PrbAvailUl             interface{}
        RRU_PrbTotDl               interface{}
        RRU_PrbTotUl               interface{}
        RRU_PrbUsedDl              interface{}
        RRU_PrbUsedUl              interface{}
        Viavi_Geo_x                interface{}
        Viavi_Geo_y                interface{}
        Viavi_Geo_z                interface{}
        Viavi_GnbDuId              interface{}
        Viavi_NrCgi                interface{}
        Viavi_NrPci                interface{}
	Viavi_Radio_antennaType    interface{}
	Viavi_Radio_azimuth	   interface{}
        Viavi_Radio_power          interface{}
}
//VIAVI 1.3 simulation cell metrics
/*
type CellMetricsEntry struct {
	DRB_UEThpDl                interface{}
	DRB_UEThpUl                interface{}
	QosFlow_TotPdcpPduVolumeDl interface{}
	QosFlow_TotPdcpPduVolumeUl interface{}
	RRC_ConnMax                interface{}
	RRC_ConnMean               interface{}
	RRU_PrbAvailDl             interface{}
	RRU_PrbAvailUl             interface{}
	RRU_PrbTotDl               interface{}
	RRU_PrbTotUl               interface{}
	RRU_PrbUsedDl              interface{}
	RRU_PrbUsedUl              interface{}
	Viavi_Geo_x                interface{}
	Viavi_Geo_y                interface{}
	Viavi_Geo_z                interface{}
	Viavi_GnbDuId              interface{}
	Viavi_NrCgi                interface{}
	Viavi_NrPci                interface{}
	Viavi_Radio_power          interface{}
	Viavi_Radio_sectors        interface{}
}
*/
type CellRFType struct {
	RSRP   int `json:"rsrp"`
	RSRQ   int `json:"rsrq"`
	RSSINR int `json:"rssinr"`
}

type NeighborCellRFType struct {
	CellID string     `json:"CID"`
	CellRF CellRFType `json:"CellRF"`
}
//VIAVI 1.4.0 E2TEST UE metrics
/*
type UeMetricsEntry struct {
Viavi_UE_id 	interface{}
Viavi_Cell_id 	interface{}
Viavi_Slice_id 	interface{}
DRB_UEThpDl 	interface{}
DRB_UEThpUl 	interface{}
RRU_PrbUsedDl 	interface{}
RRU_PrbUsedUl 	interface{}
RF_serving_id 	interface{}
RF_serving_RSRP 	interface{}
RF_serving_RSRQ 	interface{}
RF_serving_RSSINR 	interface{}
RF_nb1_id 	interface{}
RF_nb1_RSRP 	interface{}
RF_nb1_RSRQ 	interface{}
RF_nb1_RSSINR 	interface{}
RF_nb2_id 	interface{}
RF_nb2_RSRP 	interface{}
RF_nb2_RSRQ 	interface{}
RF_nb2_RSSINR 	interface{}
RF_nb3_id 	interface{}
RF_nb3_RSRP 	interface{}
RF_nb3_RSRQ 	interface{}
RF_nb3_RSSINR 	interface{}
RF_nb4_id 	interface{}
RF_nb4_RSRP 	interface{}
RF_nb4_RSRQ 	interface{}
RF_nb4_RSSINR 	interface{}
RF_nb5_id 	interface{}
RF_nb5_RSRP 	interface{}
RF_nb5_RSRQ 	interface{}
RF_nb5_RSSINR 	interface{}
RF_nb6_id 	interface{}
RF_nb6_RSRP 	interface{}
RF_nb6_RSRQ 	interface{}
RF_nb6_RSSINR 	interface{}
RF_nb7_id 	interface{}
RF_nb7_RSRP 	interface{}
RF_nb7_RSRQ 	interface{}
RF_nb7_RSSINR 	interface{}
TB_TotNbrUlInit 	interface{}
TB_IntialErrNbrDl 	interface{}
TB_ResidualErrNbrDl 	interface{}
TB_ResidualErrNbrUl 	interface{}
TB_TotNbrDlInitial_Qpsk 	interface{}
TB_TotNbrDlInitial_16Qam 	interface{}
TB_TotNbrDlInitial_64Qam 	interface{}
TB_TotNbrDlInitial_256Qam 	interface{}
TB_IntialErrNbrDl_Qpsk 	interface{}
TB_IntialErrNbrDl_16Qam 	interface{}
TB_IntialErrNbrDl_64Qam 	interface{}
TB_IntialErrNbrDl_256Qam 	interface{}
TB_TotNbrUlInit_Qpsk 	interface{}
TB_TotNbrUlInit_16Qam 	interface{}
TB_TotNbrUlInit_64Qam 	interface{}
TB_TotNbrUlInit_256Qam 	interface{}
TB_ErrNbrUlInitial 	interface{}
TB_ErrNbrUlInitial_Qpsk 	interface{}
TB_ErrNbrUlInitial_16Qam 	interface{}
TB_ErrNbrUlInitial_64Qam 	interface{}
TB_ErrNbrUlInitial_256Qam 	interface{}
Viavi_Geo_x 	interface{}
Viavi_Geo_y 	interface{}
Viavi_Geo_z 	interface{}

}*/

//VIAVI 1.4.0 simulation UE metrics
type UeMetricsEntry struct {
        DRB_UECqiDl                 interface{}
        DRB_UECqiUl                 interface{}
        DRB_UEThpDl                 interface{}
        DRB_UEThpUl                 interface{}
	QosFlow_TotPdcpPduVolumeDl  interface{}
        RRU_PrbUsedDl               interface{}
        RRU_PrbUsedUl               interface{}
        TB_TotNbrDl                 interface{}
        TB_TotNbrUl                 interface{}
        Viavi_Cell_id               interface{}
        Viavi_Geo_x                 interface{}
        Viavi_Geo_y                 interface{}
        Viavi_Geo_z                 interface{}
	Viavi_QoS_5qi		    interface{}
	Viavi_QoS_CellId	    interface{}
	Viavi_QoS_DrbId		    interface{}
	Viavi_QoS_Gfbr		    interface{}
	Viavi_Slice_id              interface{}
	Viavi_UE_BeamId             interface{}
	Viavi_UE_FrameCnt           interface{}
        Viavi_UE_RsSinr             interface{}
        Viavi_UE_Rsrp               interface{}
        Viavi_UE_Rsrq               interface{}
        Viavi_UE_TcRnti             interface{}
        Viavi_UE_anomalies          interface{}
        Viavi_UE_id                 interface{}
        Viavi_UE_servingDistance    interface{}
        Viavi_UE_speed              interface{}
        Viavi_UE_targetThroughputDl interface{}
        Viavi_UE_targetThroughputUl interface{}
}
//VIAVI 1.3 simulation UE metrics
/*
type UeMetricsEntry struct {
	DRB_UECqiDl                 interface{}
	DRB_UECqiUl                 interface{}
	DRB_UEThpDl                 interface{}
	DRB_UEThpUl                 interface{}
	RRU_PrbUsedDl               interface{}
	RRU_PrbUsedUl               interface{}
	TB_TotNbrDl                 interface{}
	TB_TotNbrUl                 interface{}
	Viavi_Cell_id               interface{}
	Viavi_Geo_x                 interface{}
	Viavi_Geo_y                 interface{}
	Viavi_Geo_z                 interface{}
	Viavi_Pdcp_kb               interface{}
	Viavi_Slice_gbps            interface{}
	Viavi_Slice_id              interface{}
	Viavi_UE_BeamId             interface{}
	Viavi_UE_FrameCnt           interface{}
	Viavi_UE_RsSinr             interface{}
	Viavi_UE_Rsrp               interface{}
	Viavi_UE_Rsrq               interface{}
	Viavi_UE_TcRnti             interface{}
	Viavi_UE_anomalies          interface{}
	Viavi_UE_id                 interface{}
	Viavi_UE_servingDistance    interface{}
	Viavi_UE_speed              interface{}
	Viavi_UE_targetThroughputDl interface{}
	Viavi_UE_targetThroughputUl interface{}
}
*/
type ViaviMessages struct {
	results ViaviMessageBody
}

type ViaviMessageBody struct {
	statement_id int
	series       ViaviMetrics
}

type ViaviMetrics struct {
	name    string
	columns string
	values  string
}
