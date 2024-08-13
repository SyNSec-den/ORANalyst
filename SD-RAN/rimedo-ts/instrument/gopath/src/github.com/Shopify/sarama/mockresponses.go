//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1
)

import (
	"fmt"
	"strings"
	"sync"
)

// TestReporter has methods matching go's testing.T to avoid importing
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:9
// `testing` in the main part of the library.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:11
type TestReporter interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// MockResponse is a response builder interface it defines one method that
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:18
// allows generating a response based on a request body. MockResponses are used
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:18
// to program behavior of MockBroker in tests.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:21
type MockResponse interface {
	For(reqBody versionedDecoder) (res encoderWithHeader)
}

// MockWrapper is a mock response builder that returns a particular concrete
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:25
// response regardless of the actual request passed to the `For` method.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:27
type MockWrapper struct {
	res encoderWithHeader
}

func (mw *MockWrapper) For(reqBody versionedDecoder) (res encoderWithHeader) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:31
	_go_fuzz_dep_.CoverTab[104410]++
												return mw.res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:32
	// _ = "end of CoverTab[104410]"
}

func NewMockWrapper(res encoderWithHeader) *MockWrapper {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:35
	_go_fuzz_dep_.CoverTab[104411]++
												return &MockWrapper{res: res}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:36
	// _ = "end of CoverTab[104411]"
}

// MockSequence is a mock response builder that is created from a sequence of
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:39
// concrete responses. Every time when a `MockBroker` calls its `For` method
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:39
// the next response from the sequence is returned. When the end of the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:39
// sequence is reached the last element from the sequence is returned.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:43
type MockSequence struct {
	responses []MockResponse
}

func NewMockSequence(responses ...interface{}) *MockSequence {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:47
	_go_fuzz_dep_.CoverTab[104412]++
												ms := &MockSequence{}
												ms.responses = make([]MockResponse, len(responses))
												for i, res := range responses {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:50
		_go_fuzz_dep_.CoverTab[104414]++
													switch res := res.(type) {
		case MockResponse:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:52
			_go_fuzz_dep_.CoverTab[104415]++
														ms.responses[i] = res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:53
			// _ = "end of CoverTab[104415]"
		case encoderWithHeader:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:54
			_go_fuzz_dep_.CoverTab[104416]++
														ms.responses[i] = NewMockWrapper(res)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:55
			// _ = "end of CoverTab[104416]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:56
			_go_fuzz_dep_.CoverTab[104417]++
														panic(fmt.Sprintf("Unexpected response type: %T", res))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:57
			// _ = "end of CoverTab[104417]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:58
		// _ = "end of CoverTab[104414]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:59
	// _ = "end of CoverTab[104412]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:59
	_go_fuzz_dep_.CoverTab[104413]++
												return ms
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:60
	// _ = "end of CoverTab[104413]"
}

func (mc *MockSequence) For(reqBody versionedDecoder) (res encoderWithHeader) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:63
	_go_fuzz_dep_.CoverTab[104418]++
												res = mc.responses[0].For(reqBody)
												if len(mc.responses) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:65
		_go_fuzz_dep_.CoverTab[104420]++
													mc.responses = mc.responses[1:]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:66
		// _ = "end of CoverTab[104420]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:67
		_go_fuzz_dep_.CoverTab[104421]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:67
		// _ = "end of CoverTab[104421]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:67
	// _ = "end of CoverTab[104418]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:67
	_go_fuzz_dep_.CoverTab[104419]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:68
	// _ = "end of CoverTab[104419]"
}

type MockListGroupsResponse struct {
	groups	map[string]string
	t	TestReporter
}

func NewMockListGroupsResponse(t TestReporter) *MockListGroupsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:76
	_go_fuzz_dep_.CoverTab[104422]++
												return &MockListGroupsResponse{
		groups:	make(map[string]string),
		t:	t,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:80
	// _ = "end of CoverTab[104422]"
}

func (m *MockListGroupsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:83
	_go_fuzz_dep_.CoverTab[104423]++
												request := reqBody.(*ListGroupsRequest)
												_ = request
												response := &ListGroupsResponse{
		Groups: m.groups,
	}
												return response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:89
	// _ = "end of CoverTab[104423]"
}

func (m *MockListGroupsResponse) AddGroup(groupID, protocolType string) *MockListGroupsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:92
	_go_fuzz_dep_.CoverTab[104424]++
												m.groups[groupID] = protocolType
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:94
	// _ = "end of CoverTab[104424]"
}

type MockDescribeGroupsResponse struct {
	groups	map[string]*GroupDescription
	t	TestReporter
}

func NewMockDescribeGroupsResponse(t TestReporter) *MockDescribeGroupsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:102
	_go_fuzz_dep_.CoverTab[104425]++
												return &MockDescribeGroupsResponse{
		t:	t,
		groups:	make(map[string]*GroupDescription),
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:106
	// _ = "end of CoverTab[104425]"
}

func (m *MockDescribeGroupsResponse) AddGroupDescription(groupID string, description *GroupDescription) *MockDescribeGroupsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:109
	_go_fuzz_dep_.CoverTab[104426]++
												m.groups[groupID] = description
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:111
	// _ = "end of CoverTab[104426]"
}

func (m *MockDescribeGroupsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:114
	_go_fuzz_dep_.CoverTab[104427]++
												request := reqBody.(*DescribeGroupsRequest)

												response := &DescribeGroupsResponse{}
												for _, requestedGroup := range request.Groups {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:118
		_go_fuzz_dep_.CoverTab[104429]++
													if group, ok := m.groups[requestedGroup]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:119
			_go_fuzz_dep_.CoverTab[104430]++
														response.Groups = append(response.Groups, group)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:120
			// _ = "end of CoverTab[104430]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:121
			_go_fuzz_dep_.CoverTab[104431]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:124
			response.Groups = append(response.Groups, &GroupDescription{
				GroupId:	requestedGroup,
				State:		"Dead",
			})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:127
			// _ = "end of CoverTab[104431]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:128
		// _ = "end of CoverTab[104429]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:129
	// _ = "end of CoverTab[104427]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:129
	_go_fuzz_dep_.CoverTab[104428]++

												return response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:131
	// _ = "end of CoverTab[104428]"
}

// MockMetadataResponse is a `MetadataResponse` builder.
type MockMetadataResponse struct {
	controllerID	int32
	leaders		map[string]map[int32]int32
	brokers		map[string]int32
	t		TestReporter
}

func NewMockMetadataResponse(t TestReporter) *MockMetadataResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:142
	_go_fuzz_dep_.CoverTab[104432]++
												return &MockMetadataResponse{
		leaders:	make(map[string]map[int32]int32),
		brokers:	make(map[string]int32),
		t:		t,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:147
	// _ = "end of CoverTab[104432]"
}

func (mmr *MockMetadataResponse) SetLeader(topic string, partition, brokerID int32) *MockMetadataResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:150
	_go_fuzz_dep_.CoverTab[104433]++
												partitions := mmr.leaders[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:152
		_go_fuzz_dep_.CoverTab[104435]++
													partitions = make(map[int32]int32)
													mmr.leaders[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:154
		// _ = "end of CoverTab[104435]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:155
		_go_fuzz_dep_.CoverTab[104436]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:155
		// _ = "end of CoverTab[104436]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:155
	// _ = "end of CoverTab[104433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:155
	_go_fuzz_dep_.CoverTab[104434]++
												partitions[partition] = brokerID
												return mmr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:157
	// _ = "end of CoverTab[104434]"
}

func (mmr *MockMetadataResponse) SetBroker(addr string, brokerID int32) *MockMetadataResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:160
	_go_fuzz_dep_.CoverTab[104437]++
												mmr.brokers[addr] = brokerID
												return mmr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:162
	// _ = "end of CoverTab[104437]"
}

func (mmr *MockMetadataResponse) SetController(brokerID int32) *MockMetadataResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:165
	_go_fuzz_dep_.CoverTab[104438]++
												mmr.controllerID = brokerID
												return mmr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:167
	// _ = "end of CoverTab[104438]"
}

func (mmr *MockMetadataResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:170
	_go_fuzz_dep_.CoverTab[104439]++
												metadataRequest := reqBody.(*MetadataRequest)
												metadataResponse := &MetadataResponse{
		Version:	metadataRequest.version(),
		ControllerID:	mmr.controllerID,
	}
	for addr, brokerID := range mmr.brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:176
		_go_fuzz_dep_.CoverTab[104444]++
													metadataResponse.AddBroker(addr, brokerID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:177
		// _ = "end of CoverTab[104444]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:178
	// _ = "end of CoverTab[104439]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:178
	_go_fuzz_dep_.CoverTab[104440]++

	// Generate set of replicas
	var replicas []int32
	var offlineReplicas []int32
	for _, brokerID := range mmr.brokers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:183
		_go_fuzz_dep_.CoverTab[104445]++
													replicas = append(replicas, brokerID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:184
		// _ = "end of CoverTab[104445]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:185
	// _ = "end of CoverTab[104440]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:185
	_go_fuzz_dep_.CoverTab[104441]++

												if len(metadataRequest.Topics) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:187
		_go_fuzz_dep_.CoverTab[104446]++
													for topic, partitions := range mmr.leaders {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:188
			_go_fuzz_dep_.CoverTab[104448]++
														for partition, brokerID := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:189
				_go_fuzz_dep_.CoverTab[104449]++
															metadataResponse.AddTopicPartition(topic, partition, brokerID, replicas, replicas, offlineReplicas, ErrNoError)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:190
				// _ = "end of CoverTab[104449]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:191
			// _ = "end of CoverTab[104448]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:192
		// _ = "end of CoverTab[104446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:192
		_go_fuzz_dep_.CoverTab[104447]++
													return metadataResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:193
		// _ = "end of CoverTab[104447]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:194
		_go_fuzz_dep_.CoverTab[104450]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:194
		// _ = "end of CoverTab[104450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:194
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:194
	// _ = "end of CoverTab[104441]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:194
	_go_fuzz_dep_.CoverTab[104442]++
												for _, topic := range metadataRequest.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:195
		_go_fuzz_dep_.CoverTab[104451]++
													for partition, brokerID := range mmr.leaders[topic] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:196
			_go_fuzz_dep_.CoverTab[104452]++
														metadataResponse.AddTopicPartition(topic, partition, brokerID, replicas, replicas, offlineReplicas, ErrNoError)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:197
			// _ = "end of CoverTab[104452]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:198
		// _ = "end of CoverTab[104451]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:199
	// _ = "end of CoverTab[104442]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:199
	_go_fuzz_dep_.CoverTab[104443]++
												return metadataResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:200
	// _ = "end of CoverTab[104443]"
}

// MockOffsetResponse is an `OffsetResponse` builder.
type MockOffsetResponse struct {
	offsets	map[string]map[int32]map[int64]int64
	t	TestReporter
	version	int16
}

func NewMockOffsetResponse(t TestReporter) *MockOffsetResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:210
	_go_fuzz_dep_.CoverTab[104453]++
												return &MockOffsetResponse{
		offsets:	make(map[string]map[int32]map[int64]int64),
		t:		t,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:214
	// _ = "end of CoverTab[104453]"
}

func (mor *MockOffsetResponse) SetVersion(version int16) *MockOffsetResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:217
	_go_fuzz_dep_.CoverTab[104454]++
												mor.version = version
												return mor
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:219
	// _ = "end of CoverTab[104454]"
}

func (mor *MockOffsetResponse) SetOffset(topic string, partition int32, time, offset int64) *MockOffsetResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:222
	_go_fuzz_dep_.CoverTab[104455]++
												partitions := mor.offsets[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:224
		_go_fuzz_dep_.CoverTab[104458]++
													partitions = make(map[int32]map[int64]int64)
													mor.offsets[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:226
		// _ = "end of CoverTab[104458]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:227
		_go_fuzz_dep_.CoverTab[104459]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:227
		// _ = "end of CoverTab[104459]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:227
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:227
	// _ = "end of CoverTab[104455]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:227
	_go_fuzz_dep_.CoverTab[104456]++
												times := partitions[partition]
												if times == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:229
		_go_fuzz_dep_.CoverTab[104460]++
													times = make(map[int64]int64)
													partitions[partition] = times
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:231
		// _ = "end of CoverTab[104460]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:232
		_go_fuzz_dep_.CoverTab[104461]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:232
		// _ = "end of CoverTab[104461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:232
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:232
	// _ = "end of CoverTab[104456]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:232
	_go_fuzz_dep_.CoverTab[104457]++
												times[time] = offset
												return mor
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:234
	// _ = "end of CoverTab[104457]"
}

func (mor *MockOffsetResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:237
	_go_fuzz_dep_.CoverTab[104462]++
												offsetRequest := reqBody.(*OffsetRequest)
												offsetResponse := &OffsetResponse{Version: mor.version}
												for topic, partitions := range offsetRequest.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:240
		_go_fuzz_dep_.CoverTab[104464]++
													for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:241
			_go_fuzz_dep_.CoverTab[104465]++
														offset := mor.getOffset(topic, partition, block.time)
														offsetResponse.AddTopicPartition(topic, partition, offset)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:243
			// _ = "end of CoverTab[104465]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:244
		// _ = "end of CoverTab[104464]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:245
	// _ = "end of CoverTab[104462]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:245
	_go_fuzz_dep_.CoverTab[104463]++
												return offsetResponse
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:246
	// _ = "end of CoverTab[104463]"
}

func (mor *MockOffsetResponse) getOffset(topic string, partition int32, time int64) int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:249
	_go_fuzz_dep_.CoverTab[104466]++
												partitions := mor.offsets[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:251
		_go_fuzz_dep_.CoverTab[104470]++
													mor.t.Errorf("missing topic: %s", topic)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:252
		// _ = "end of CoverTab[104470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:253
		_go_fuzz_dep_.CoverTab[104471]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:253
		// _ = "end of CoverTab[104471]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:253
	// _ = "end of CoverTab[104466]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:253
	_go_fuzz_dep_.CoverTab[104467]++
												times := partitions[partition]
												if times == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:255
		_go_fuzz_dep_.CoverTab[104472]++
													mor.t.Errorf("missing partition: %d", partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:256
		// _ = "end of CoverTab[104472]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:257
		_go_fuzz_dep_.CoverTab[104473]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:257
		// _ = "end of CoverTab[104473]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:257
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:257
	// _ = "end of CoverTab[104467]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:257
	_go_fuzz_dep_.CoverTab[104468]++
												offset, ok := times[time]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:259
		_go_fuzz_dep_.CoverTab[104474]++
													mor.t.Errorf("missing time: %d", time)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:260
		// _ = "end of CoverTab[104474]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:261
		_go_fuzz_dep_.CoverTab[104475]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:261
		// _ = "end of CoverTab[104475]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:261
	// _ = "end of CoverTab[104468]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:261
	_go_fuzz_dep_.CoverTab[104469]++
												return offset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:262
	// _ = "end of CoverTab[104469]"
}

// MockFetchResponse is a `FetchResponse` builder.
type MockFetchResponse struct {
	messages	map[string]map[int32]map[int64]Encoder
	messagesLock	*sync.RWMutex
	highWaterMarks	map[string]map[int32]int64
	t		TestReporter
	batchSize	int
	version		int16
}

func NewMockFetchResponse(t TestReporter, batchSize int) *MockFetchResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:275
	_go_fuzz_dep_.CoverTab[104476]++
												return &MockFetchResponse{
		messages:	make(map[string]map[int32]map[int64]Encoder),
		messagesLock:	&sync.RWMutex{},
		highWaterMarks:	make(map[string]map[int32]int64),
		t:		t,
		batchSize:	batchSize,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:282
	// _ = "end of CoverTab[104476]"
}

func (mfr *MockFetchResponse) SetVersion(version int16) *MockFetchResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:285
	_go_fuzz_dep_.CoverTab[104477]++
												mfr.version = version
												return mfr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:287
	// _ = "end of CoverTab[104477]"
}

func (mfr *MockFetchResponse) SetMessage(topic string, partition int32, offset int64, msg Encoder) *MockFetchResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:290
	_go_fuzz_dep_.CoverTab[104478]++
												mfr.messagesLock.Lock()
												defer mfr.messagesLock.Unlock()
												partitions := mfr.messages[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:294
		_go_fuzz_dep_.CoverTab[104481]++
													partitions = make(map[int32]map[int64]Encoder)
													mfr.messages[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:296
		// _ = "end of CoverTab[104481]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:297
		_go_fuzz_dep_.CoverTab[104482]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:297
		// _ = "end of CoverTab[104482]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:297
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:297
	// _ = "end of CoverTab[104478]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:297
	_go_fuzz_dep_.CoverTab[104479]++
												messages := partitions[partition]
												if messages == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:299
		_go_fuzz_dep_.CoverTab[104483]++
													messages = make(map[int64]Encoder)
													partitions[partition] = messages
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:301
		// _ = "end of CoverTab[104483]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:302
		_go_fuzz_dep_.CoverTab[104484]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:302
		// _ = "end of CoverTab[104484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:302
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:302
	// _ = "end of CoverTab[104479]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:302
	_go_fuzz_dep_.CoverTab[104480]++
												messages[offset] = msg
												return mfr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:304
	// _ = "end of CoverTab[104480]"
}

func (mfr *MockFetchResponse) SetHighWaterMark(topic string, partition int32, offset int64) *MockFetchResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:307
	_go_fuzz_dep_.CoverTab[104485]++
												partitions := mfr.highWaterMarks[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:309
		_go_fuzz_dep_.CoverTab[104487]++
													partitions = make(map[int32]int64)
													mfr.highWaterMarks[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:311
		// _ = "end of CoverTab[104487]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:312
		_go_fuzz_dep_.CoverTab[104488]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:312
		// _ = "end of CoverTab[104488]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:312
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:312
	// _ = "end of CoverTab[104485]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:312
	_go_fuzz_dep_.CoverTab[104486]++
												partitions[partition] = offset
												return mfr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:314
	// _ = "end of CoverTab[104486]"
}

func (mfr *MockFetchResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:317
	_go_fuzz_dep_.CoverTab[104489]++
												fetchRequest := reqBody.(*FetchRequest)
												res := &FetchResponse{
		Version: mfr.version,
	}
	for topic, partitions := range fetchRequest.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:322
		_go_fuzz_dep_.CoverTab[104491]++
													for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:323
			_go_fuzz_dep_.CoverTab[104492]++
														initialOffset := block.fetchOffset
														offset := initialOffset
														maxOffset := initialOffset + int64(mfr.getMessageCount(topic, partition))
														for i := 0; i < mfr.batchSize && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:327
				_go_fuzz_dep_.CoverTab[104495]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:327
				return offset < maxOffset
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:327
				// _ = "end of CoverTab[104495]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:327
			}(); {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:327
				_go_fuzz_dep_.CoverTab[104496]++
															msg := mfr.getMessage(topic, partition, offset)
															if msg != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:329
					_go_fuzz_dep_.CoverTab[104498]++
																res.AddMessage(topic, partition, nil, msg, offset)
																i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:331
					// _ = "end of CoverTab[104498]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:332
					_go_fuzz_dep_.CoverTab[104499]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:332
					// _ = "end of CoverTab[104499]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:332
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:332
				// _ = "end of CoverTab[104496]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:332
				_go_fuzz_dep_.CoverTab[104497]++
															offset++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:333
				// _ = "end of CoverTab[104497]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:334
			// _ = "end of CoverTab[104492]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:334
			_go_fuzz_dep_.CoverTab[104493]++
														fb := res.GetBlock(topic, partition)
														if fb == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:336
				_go_fuzz_dep_.CoverTab[104500]++
															res.AddError(topic, partition, ErrNoError)
															fb = res.GetBlock(topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:338
				// _ = "end of CoverTab[104500]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:339
				_go_fuzz_dep_.CoverTab[104501]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:339
				// _ = "end of CoverTab[104501]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:339
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:339
			// _ = "end of CoverTab[104493]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:339
			_go_fuzz_dep_.CoverTab[104494]++
														fb.HighWaterMarkOffset = mfr.getHighWaterMark(topic, partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:340
			// _ = "end of CoverTab[104494]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:341
		// _ = "end of CoverTab[104491]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:342
	// _ = "end of CoverTab[104489]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:342
	_go_fuzz_dep_.CoverTab[104490]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:343
	// _ = "end of CoverTab[104490]"
}

func (mfr *MockFetchResponse) getMessage(topic string, partition int32, offset int64) Encoder {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:346
	_go_fuzz_dep_.CoverTab[104502]++
												mfr.messagesLock.RLock()
												defer mfr.messagesLock.RUnlock()
												partitions := mfr.messages[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:350
		_go_fuzz_dep_.CoverTab[104505]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:351
		// _ = "end of CoverTab[104505]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:352
		_go_fuzz_dep_.CoverTab[104506]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:352
		// _ = "end of CoverTab[104506]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:352
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:352
	// _ = "end of CoverTab[104502]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:352
	_go_fuzz_dep_.CoverTab[104503]++
												messages := partitions[partition]
												if messages == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:354
		_go_fuzz_dep_.CoverTab[104507]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:355
		// _ = "end of CoverTab[104507]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:356
		_go_fuzz_dep_.CoverTab[104508]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:356
		// _ = "end of CoverTab[104508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:356
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:356
	// _ = "end of CoverTab[104503]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:356
	_go_fuzz_dep_.CoverTab[104504]++
												return messages[offset]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:357
	// _ = "end of CoverTab[104504]"
}

func (mfr *MockFetchResponse) getMessageCount(topic string, partition int32) int {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:360
	_go_fuzz_dep_.CoverTab[104509]++
												mfr.messagesLock.RLock()
												defer mfr.messagesLock.RUnlock()
												partitions := mfr.messages[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:364
		_go_fuzz_dep_.CoverTab[104512]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:365
		// _ = "end of CoverTab[104512]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:366
		_go_fuzz_dep_.CoverTab[104513]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:366
		// _ = "end of CoverTab[104513]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:366
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:366
	// _ = "end of CoverTab[104509]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:366
	_go_fuzz_dep_.CoverTab[104510]++
												messages := partitions[partition]
												if messages == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:368
		_go_fuzz_dep_.CoverTab[104514]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:369
		// _ = "end of CoverTab[104514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:370
		_go_fuzz_dep_.CoverTab[104515]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:370
		// _ = "end of CoverTab[104515]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:370
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:370
	// _ = "end of CoverTab[104510]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:370
	_go_fuzz_dep_.CoverTab[104511]++
												return len(messages)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:371
	// _ = "end of CoverTab[104511]"
}

func (mfr *MockFetchResponse) getHighWaterMark(topic string, partition int32) int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:374
	_go_fuzz_dep_.CoverTab[104516]++
												partitions := mfr.highWaterMarks[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:376
		_go_fuzz_dep_.CoverTab[104518]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:377
		// _ = "end of CoverTab[104518]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:378
		_go_fuzz_dep_.CoverTab[104519]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:378
		// _ = "end of CoverTab[104519]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:378
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:378
	// _ = "end of CoverTab[104516]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:378
	_go_fuzz_dep_.CoverTab[104517]++
												return partitions[partition]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:379
	// _ = "end of CoverTab[104517]"
}

// MockConsumerMetadataResponse is a `ConsumerMetadataResponse` builder.
type MockConsumerMetadataResponse struct {
	coordinators	map[string]interface{}
	t		TestReporter
}

func NewMockConsumerMetadataResponse(t TestReporter) *MockConsumerMetadataResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:388
	_go_fuzz_dep_.CoverTab[104520]++
												return &MockConsumerMetadataResponse{
		coordinators:	make(map[string]interface{}),
		t:		t,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:392
	// _ = "end of CoverTab[104520]"
}

func (mr *MockConsumerMetadataResponse) SetCoordinator(group string, broker *MockBroker) *MockConsumerMetadataResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:395
	_go_fuzz_dep_.CoverTab[104521]++
												mr.coordinators[group] = broker
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:397
	// _ = "end of CoverTab[104521]"
}

func (mr *MockConsumerMetadataResponse) SetError(group string, kerror KError) *MockConsumerMetadataResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:400
	_go_fuzz_dep_.CoverTab[104522]++
												mr.coordinators[group] = kerror
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:402
	// _ = "end of CoverTab[104522]"
}

func (mr *MockConsumerMetadataResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:405
	_go_fuzz_dep_.CoverTab[104523]++
												req := reqBody.(*ConsumerMetadataRequest)
												group := req.ConsumerGroup
												res := &ConsumerMetadataResponse{}
												v := mr.coordinators[group]
												switch v := v.(type) {
	case *MockBroker:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:411
		_go_fuzz_dep_.CoverTab[104525]++
													res.Coordinator = &Broker{id: v.BrokerID(), addr: v.Addr()}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:412
		// _ = "end of CoverTab[104525]"
	case KError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:413
		_go_fuzz_dep_.CoverTab[104526]++
													res.Err = v
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:414
		// _ = "end of CoverTab[104526]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:415
	// _ = "end of CoverTab[104523]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:415
	_go_fuzz_dep_.CoverTab[104524]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:416
	// _ = "end of CoverTab[104524]"
}

// MockFindCoordinatorResponse is a `FindCoordinatorResponse` builder.
type MockFindCoordinatorResponse struct {
	groupCoordinators	map[string]interface{}
	transCoordinators	map[string]interface{}
	t			TestReporter
}

func NewMockFindCoordinatorResponse(t TestReporter) *MockFindCoordinatorResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:426
	_go_fuzz_dep_.CoverTab[104527]++
												return &MockFindCoordinatorResponse{
		groupCoordinators:	make(map[string]interface{}),
		transCoordinators:	make(map[string]interface{}),
		t:			t,
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:431
	// _ = "end of CoverTab[104527]"
}

func (mr *MockFindCoordinatorResponse) SetCoordinator(coordinatorType CoordinatorType, group string, broker *MockBroker) *MockFindCoordinatorResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:434
	_go_fuzz_dep_.CoverTab[104528]++
												switch coordinatorType {
	case CoordinatorGroup:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:436
		_go_fuzz_dep_.CoverTab[104530]++
													mr.groupCoordinators[group] = broker
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:437
		// _ = "end of CoverTab[104530]"
	case CoordinatorTransaction:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:438
		_go_fuzz_dep_.CoverTab[104531]++
													mr.transCoordinators[group] = broker
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:439
		// _ = "end of CoverTab[104531]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:439
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:439
		_go_fuzz_dep_.CoverTab[104532]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:439
		// _ = "end of CoverTab[104532]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:440
	// _ = "end of CoverTab[104528]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:440
	_go_fuzz_dep_.CoverTab[104529]++
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:441
	// _ = "end of CoverTab[104529]"
}

func (mr *MockFindCoordinatorResponse) SetError(coordinatorType CoordinatorType, group string, kerror KError) *MockFindCoordinatorResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:444
	_go_fuzz_dep_.CoverTab[104533]++
												switch coordinatorType {
	case CoordinatorGroup:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:446
		_go_fuzz_dep_.CoverTab[104535]++
													mr.groupCoordinators[group] = kerror
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:447
		// _ = "end of CoverTab[104535]"
	case CoordinatorTransaction:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:448
		_go_fuzz_dep_.CoverTab[104536]++
													mr.transCoordinators[group] = kerror
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:449
		// _ = "end of CoverTab[104536]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:449
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:449
		_go_fuzz_dep_.CoverTab[104537]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:449
		// _ = "end of CoverTab[104537]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:450
	// _ = "end of CoverTab[104533]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:450
	_go_fuzz_dep_.CoverTab[104534]++
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:451
	// _ = "end of CoverTab[104534]"
}

func (mr *MockFindCoordinatorResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:454
	_go_fuzz_dep_.CoverTab[104538]++
												req := reqBody.(*FindCoordinatorRequest)
												res := &FindCoordinatorResponse{}
												var v interface{}
												switch req.CoordinatorType {
	case CoordinatorGroup:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:459
		_go_fuzz_dep_.CoverTab[104541]++
													v = mr.groupCoordinators[req.CoordinatorKey]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:460
		// _ = "end of CoverTab[104541]"
	case CoordinatorTransaction:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:461
		_go_fuzz_dep_.CoverTab[104542]++
													v = mr.transCoordinators[req.CoordinatorKey]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:462
		// _ = "end of CoverTab[104542]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:462
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:462
		_go_fuzz_dep_.CoverTab[104543]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:462
		// _ = "end of CoverTab[104543]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:463
	// _ = "end of CoverTab[104538]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:463
	_go_fuzz_dep_.CoverTab[104539]++
												switch v := v.(type) {
	case *MockBroker:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:465
		_go_fuzz_dep_.CoverTab[104544]++
													res.Coordinator = &Broker{id: v.BrokerID(), addr: v.Addr()}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:466
		// _ = "end of CoverTab[104544]"
	case KError:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:467
		_go_fuzz_dep_.CoverTab[104545]++
													res.Err = v
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:468
		// _ = "end of CoverTab[104545]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:469
	// _ = "end of CoverTab[104539]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:469
	_go_fuzz_dep_.CoverTab[104540]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:470
	// _ = "end of CoverTab[104540]"
}

// MockOffsetCommitResponse is a `OffsetCommitResponse` builder.
type MockOffsetCommitResponse struct {
	errors	map[string]map[string]map[int32]KError
	t	TestReporter
}

func NewMockOffsetCommitResponse(t TestReporter) *MockOffsetCommitResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:479
	_go_fuzz_dep_.CoverTab[104546]++
												return &MockOffsetCommitResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:480
	// _ = "end of CoverTab[104546]"
}

func (mr *MockOffsetCommitResponse) SetError(group, topic string, partition int32, kerror KError) *MockOffsetCommitResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:483
	_go_fuzz_dep_.CoverTab[104547]++
												if mr.errors == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:484
		_go_fuzz_dep_.CoverTab[104551]++
													mr.errors = make(map[string]map[string]map[int32]KError)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:485
		// _ = "end of CoverTab[104551]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:486
		_go_fuzz_dep_.CoverTab[104552]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:486
		// _ = "end of CoverTab[104552]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:486
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:486
	// _ = "end of CoverTab[104547]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:486
	_go_fuzz_dep_.CoverTab[104548]++
												topics := mr.errors[group]
												if topics == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:488
		_go_fuzz_dep_.CoverTab[104553]++
													topics = make(map[string]map[int32]KError)
													mr.errors[group] = topics
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:490
		// _ = "end of CoverTab[104553]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:491
		_go_fuzz_dep_.CoverTab[104554]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:491
		// _ = "end of CoverTab[104554]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:491
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:491
	// _ = "end of CoverTab[104548]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:491
	_go_fuzz_dep_.CoverTab[104549]++
												partitions := topics[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:493
		_go_fuzz_dep_.CoverTab[104555]++
													partitions = make(map[int32]KError)
													topics[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:495
		// _ = "end of CoverTab[104555]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:496
		_go_fuzz_dep_.CoverTab[104556]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:496
		// _ = "end of CoverTab[104556]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:496
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:496
	// _ = "end of CoverTab[104549]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:496
	_go_fuzz_dep_.CoverTab[104550]++
												partitions[partition] = kerror
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:498
	// _ = "end of CoverTab[104550]"
}

func (mr *MockOffsetCommitResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:501
	_go_fuzz_dep_.CoverTab[104557]++
												req := reqBody.(*OffsetCommitRequest)
												group := req.ConsumerGroup
												res := &OffsetCommitResponse{}
												for topic, partitions := range req.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:505
		_go_fuzz_dep_.CoverTab[104559]++
													for partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:506
			_go_fuzz_dep_.CoverTab[104560]++
														res.AddError(topic, partition, mr.getError(group, topic, partition))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:507
			// _ = "end of CoverTab[104560]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:508
		// _ = "end of CoverTab[104559]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:509
	// _ = "end of CoverTab[104557]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:509
	_go_fuzz_dep_.CoverTab[104558]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:510
	// _ = "end of CoverTab[104558]"
}

func (mr *MockOffsetCommitResponse) getError(group, topic string, partition int32) KError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:513
	_go_fuzz_dep_.CoverTab[104561]++
												topics := mr.errors[group]
												if topics == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:515
		_go_fuzz_dep_.CoverTab[104565]++
													return ErrNoError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:516
		// _ = "end of CoverTab[104565]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:517
		_go_fuzz_dep_.CoverTab[104566]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:517
		// _ = "end of CoverTab[104566]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:517
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:517
	// _ = "end of CoverTab[104561]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:517
	_go_fuzz_dep_.CoverTab[104562]++
												partitions := topics[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:519
		_go_fuzz_dep_.CoverTab[104567]++
													return ErrNoError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:520
		// _ = "end of CoverTab[104567]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:521
		_go_fuzz_dep_.CoverTab[104568]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:521
		// _ = "end of CoverTab[104568]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:521
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:521
	// _ = "end of CoverTab[104562]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:521
	_go_fuzz_dep_.CoverTab[104563]++
												kerror, ok := partitions[partition]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:523
		_go_fuzz_dep_.CoverTab[104569]++
													return ErrNoError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:524
		// _ = "end of CoverTab[104569]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:525
		_go_fuzz_dep_.CoverTab[104570]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:525
		// _ = "end of CoverTab[104570]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:525
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:525
	// _ = "end of CoverTab[104563]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:525
	_go_fuzz_dep_.CoverTab[104564]++
												return kerror
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:526
	// _ = "end of CoverTab[104564]"
}

// MockProduceResponse is a `ProduceResponse` builder.
type MockProduceResponse struct {
	version	int16
	errors	map[string]map[int32]KError
	t	TestReporter
}

func NewMockProduceResponse(t TestReporter) *MockProduceResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:536
	_go_fuzz_dep_.CoverTab[104571]++
												return &MockProduceResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:537
	// _ = "end of CoverTab[104571]"
}

func (mr *MockProduceResponse) SetVersion(version int16) *MockProduceResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:540
	_go_fuzz_dep_.CoverTab[104572]++
												mr.version = version
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:542
	// _ = "end of CoverTab[104572]"
}

func (mr *MockProduceResponse) SetError(topic string, partition int32, kerror KError) *MockProduceResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:545
	_go_fuzz_dep_.CoverTab[104573]++
												if mr.errors == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:546
		_go_fuzz_dep_.CoverTab[104576]++
													mr.errors = make(map[string]map[int32]KError)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:547
		// _ = "end of CoverTab[104576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:548
		_go_fuzz_dep_.CoverTab[104577]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:548
		// _ = "end of CoverTab[104577]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:548
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:548
	// _ = "end of CoverTab[104573]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:548
	_go_fuzz_dep_.CoverTab[104574]++
												partitions := mr.errors[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:550
		_go_fuzz_dep_.CoverTab[104578]++
													partitions = make(map[int32]KError)
													mr.errors[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:552
		// _ = "end of CoverTab[104578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:553
		_go_fuzz_dep_.CoverTab[104579]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:553
		// _ = "end of CoverTab[104579]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:553
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:553
	// _ = "end of CoverTab[104574]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:553
	_go_fuzz_dep_.CoverTab[104575]++
												partitions[partition] = kerror
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:555
	// _ = "end of CoverTab[104575]"
}

func (mr *MockProduceResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:558
	_go_fuzz_dep_.CoverTab[104580]++
												req := reqBody.(*ProduceRequest)
												res := &ProduceResponse{
		Version: mr.version,
	}
	for topic, partitions := range req.records {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:563
		_go_fuzz_dep_.CoverTab[104582]++
													for partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:564
			_go_fuzz_dep_.CoverTab[104583]++
														res.AddTopicPartition(topic, partition, mr.getError(topic, partition))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:565
			// _ = "end of CoverTab[104583]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:566
		// _ = "end of CoverTab[104582]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:567
	// _ = "end of CoverTab[104580]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:567
	_go_fuzz_dep_.CoverTab[104581]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:568
	// _ = "end of CoverTab[104581]"
}

func (mr *MockProduceResponse) getError(topic string, partition int32) KError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:571
	_go_fuzz_dep_.CoverTab[104584]++
												partitions := mr.errors[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:573
		_go_fuzz_dep_.CoverTab[104587]++
													return ErrNoError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:574
		// _ = "end of CoverTab[104587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:575
		_go_fuzz_dep_.CoverTab[104588]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:575
		// _ = "end of CoverTab[104588]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:575
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:575
	// _ = "end of CoverTab[104584]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:575
	_go_fuzz_dep_.CoverTab[104585]++
												kerror, ok := partitions[partition]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:577
		_go_fuzz_dep_.CoverTab[104589]++
													return ErrNoError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:578
		// _ = "end of CoverTab[104589]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:579
		_go_fuzz_dep_.CoverTab[104590]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:579
		// _ = "end of CoverTab[104590]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:579
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:579
	// _ = "end of CoverTab[104585]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:579
	_go_fuzz_dep_.CoverTab[104586]++
												return kerror
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:580
	// _ = "end of CoverTab[104586]"
}

// MockOffsetFetchResponse is a `OffsetFetchResponse` builder.
type MockOffsetFetchResponse struct {
	offsets	map[string]map[string]map[int32]*OffsetFetchResponseBlock
	error	KError
	t	TestReporter
}

func NewMockOffsetFetchResponse(t TestReporter) *MockOffsetFetchResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:590
	_go_fuzz_dep_.CoverTab[104591]++
												return &MockOffsetFetchResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:591
	// _ = "end of CoverTab[104591]"
}

func (mr *MockOffsetFetchResponse) SetOffset(group, topic string, partition int32, offset int64, metadata string, kerror KError) *MockOffsetFetchResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:594
	_go_fuzz_dep_.CoverTab[104592]++
												if mr.offsets == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:595
		_go_fuzz_dep_.CoverTab[104596]++
													mr.offsets = make(map[string]map[string]map[int32]*OffsetFetchResponseBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:596
		// _ = "end of CoverTab[104596]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:597
		_go_fuzz_dep_.CoverTab[104597]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:597
		// _ = "end of CoverTab[104597]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:597
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:597
	// _ = "end of CoverTab[104592]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:597
	_go_fuzz_dep_.CoverTab[104593]++
												topics := mr.offsets[group]
												if topics == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:599
		_go_fuzz_dep_.CoverTab[104598]++
													topics = make(map[string]map[int32]*OffsetFetchResponseBlock)
													mr.offsets[group] = topics
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:601
		// _ = "end of CoverTab[104598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:602
		_go_fuzz_dep_.CoverTab[104599]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:602
		// _ = "end of CoverTab[104599]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:602
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:602
	// _ = "end of CoverTab[104593]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:602
	_go_fuzz_dep_.CoverTab[104594]++
												partitions := topics[topic]
												if partitions == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:604
		_go_fuzz_dep_.CoverTab[104600]++
													partitions = make(map[int32]*OffsetFetchResponseBlock)
													topics[topic] = partitions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:606
		// _ = "end of CoverTab[104600]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:607
		_go_fuzz_dep_.CoverTab[104601]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:607
		// _ = "end of CoverTab[104601]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:607
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:607
	// _ = "end of CoverTab[104594]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:607
	_go_fuzz_dep_.CoverTab[104595]++
												partitions[partition] = &OffsetFetchResponseBlock{offset, 0, metadata, kerror}
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:609
	// _ = "end of CoverTab[104595]"
}

func (mr *MockOffsetFetchResponse) SetError(kerror KError) *MockOffsetFetchResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:612
	_go_fuzz_dep_.CoverTab[104602]++
												mr.error = kerror
												return mr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:614
	// _ = "end of CoverTab[104602]"
}

func (mr *MockOffsetFetchResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:617
	_go_fuzz_dep_.CoverTab[104603]++
												req := reqBody.(*OffsetFetchRequest)
												group := req.ConsumerGroup
												res := &OffsetFetchResponse{Version: req.Version}

												for topic, partitions := range mr.offsets[group] {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:622
		_go_fuzz_dep_.CoverTab[104606]++
													for partition, block := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:623
			_go_fuzz_dep_.CoverTab[104607]++
														res.AddBlock(topic, partition, block)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:624
			// _ = "end of CoverTab[104607]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:625
		// _ = "end of CoverTab[104606]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:626
	// _ = "end of CoverTab[104603]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:626
	_go_fuzz_dep_.CoverTab[104604]++

												if res.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:628
		_go_fuzz_dep_.CoverTab[104608]++
													res.Err = mr.error
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:629
		// _ = "end of CoverTab[104608]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:630
		_go_fuzz_dep_.CoverTab[104609]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:630
		// _ = "end of CoverTab[104609]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:630
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:630
	// _ = "end of CoverTab[104604]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:630
	_go_fuzz_dep_.CoverTab[104605]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:631
	// _ = "end of CoverTab[104605]"
}

type MockCreateTopicsResponse struct {
	t TestReporter
}

func NewMockCreateTopicsResponse(t TestReporter) *MockCreateTopicsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:638
	_go_fuzz_dep_.CoverTab[104610]++
												return &MockCreateTopicsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:639
	// _ = "end of CoverTab[104610]"
}

func (mr *MockCreateTopicsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:642
	_go_fuzz_dep_.CoverTab[104611]++
												req := reqBody.(*CreateTopicsRequest)
												res := &CreateTopicsResponse{
		Version: req.Version,
	}
	res.TopicErrors = make(map[string]*TopicError)

	for topic := range req.TopicDetails {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:649
		_go_fuzz_dep_.CoverTab[104613]++
													if res.Version >= 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:650
			_go_fuzz_dep_.CoverTab[104615]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:650
			return strings.HasPrefix(topic, "_")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:650
			// _ = "end of CoverTab[104615]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:650
		}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:650
			_go_fuzz_dep_.CoverTab[104616]++
														msg := "insufficient permissions to create topic with reserved prefix"
														res.TopicErrors[topic] = &TopicError{
				Err:	ErrTopicAuthorizationFailed,
				ErrMsg:	&msg,
			}
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:656
			// _ = "end of CoverTab[104616]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:657
			_go_fuzz_dep_.CoverTab[104617]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:657
			// _ = "end of CoverTab[104617]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:657
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:657
		// _ = "end of CoverTab[104613]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:657
		_go_fuzz_dep_.CoverTab[104614]++
													res.TopicErrors[topic] = &TopicError{Err: ErrNoError}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:658
		// _ = "end of CoverTab[104614]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:659
	// _ = "end of CoverTab[104611]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:659
	_go_fuzz_dep_.CoverTab[104612]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:660
	// _ = "end of CoverTab[104612]"
}

type MockDeleteTopicsResponse struct {
	t TestReporter
}

func NewMockDeleteTopicsResponse(t TestReporter) *MockDeleteTopicsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:667
	_go_fuzz_dep_.CoverTab[104618]++
												return &MockDeleteTopicsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:668
	// _ = "end of CoverTab[104618]"
}

func (mr *MockDeleteTopicsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:671
	_go_fuzz_dep_.CoverTab[104619]++
												req := reqBody.(*DeleteTopicsRequest)
												res := &DeleteTopicsResponse{}
												res.TopicErrorCodes = make(map[string]KError)

												for _, topic := range req.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:676
		_go_fuzz_dep_.CoverTab[104621]++
													res.TopicErrorCodes[topic] = ErrNoError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:677
		// _ = "end of CoverTab[104621]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:678
	// _ = "end of CoverTab[104619]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:678
	_go_fuzz_dep_.CoverTab[104620]++
												res.Version = req.Version
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:680
	// _ = "end of CoverTab[104620]"
}

type MockCreatePartitionsResponse struct {
	t TestReporter
}

func NewMockCreatePartitionsResponse(t TestReporter) *MockCreatePartitionsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:687
	_go_fuzz_dep_.CoverTab[104622]++
												return &MockCreatePartitionsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:688
	// _ = "end of CoverTab[104622]"
}

func (mr *MockCreatePartitionsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:691
	_go_fuzz_dep_.CoverTab[104623]++
												req := reqBody.(*CreatePartitionsRequest)
												res := &CreatePartitionsResponse{}
												res.TopicPartitionErrors = make(map[string]*TopicPartitionError)

												for topic := range req.TopicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:696
		_go_fuzz_dep_.CoverTab[104625]++
													if strings.HasPrefix(topic, "_") {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:697
			_go_fuzz_dep_.CoverTab[104627]++
														msg := "insufficient permissions to create partition on topic with reserved prefix"
														res.TopicPartitionErrors[topic] = &TopicPartitionError{
				Err:	ErrTopicAuthorizationFailed,
				ErrMsg:	&msg,
			}
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:703
			// _ = "end of CoverTab[104627]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:704
			_go_fuzz_dep_.CoverTab[104628]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:704
			// _ = "end of CoverTab[104628]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:704
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:704
		// _ = "end of CoverTab[104625]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:704
		_go_fuzz_dep_.CoverTab[104626]++
													res.TopicPartitionErrors[topic] = &TopicPartitionError{Err: ErrNoError}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:705
		// _ = "end of CoverTab[104626]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:706
	// _ = "end of CoverTab[104623]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:706
	_go_fuzz_dep_.CoverTab[104624]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:707
	// _ = "end of CoverTab[104624]"
}

type MockAlterPartitionReassignmentsResponse struct {
	t TestReporter
}

func NewMockAlterPartitionReassignmentsResponse(t TestReporter) *MockAlterPartitionReassignmentsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:714
	_go_fuzz_dep_.CoverTab[104629]++
												return &MockAlterPartitionReassignmentsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:715
	// _ = "end of CoverTab[104629]"
}

func (mr *MockAlterPartitionReassignmentsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:718
	_go_fuzz_dep_.CoverTab[104630]++
												req := reqBody.(*AlterPartitionReassignmentsRequest)
												_ = req
												res := &AlterPartitionReassignmentsResponse{}
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:722
	// _ = "end of CoverTab[104630]"
}

type MockListPartitionReassignmentsResponse struct {
	t TestReporter
}

func NewMockListPartitionReassignmentsResponse(t TestReporter) *MockListPartitionReassignmentsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:729
	_go_fuzz_dep_.CoverTab[104631]++
												return &MockListPartitionReassignmentsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:730
	// _ = "end of CoverTab[104631]"
}

func (mr *MockListPartitionReassignmentsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:733
	_go_fuzz_dep_.CoverTab[104632]++
												req := reqBody.(*ListPartitionReassignmentsRequest)
												_ = req
												res := &ListPartitionReassignmentsResponse{}

												for topic, partitions := range req.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:738
		_go_fuzz_dep_.CoverTab[104634]++
													for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:739
			_go_fuzz_dep_.CoverTab[104635]++
														res.AddBlock(topic, partition, []int32{0}, []int32{1}, []int32{2})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:740
			// _ = "end of CoverTab[104635]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:741
		// _ = "end of CoverTab[104634]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:742
	// _ = "end of CoverTab[104632]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:742
	_go_fuzz_dep_.CoverTab[104633]++

												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:744
	// _ = "end of CoverTab[104633]"
}

type MockDeleteRecordsResponse struct {
	t TestReporter
}

func NewMockDeleteRecordsResponse(t TestReporter) *MockDeleteRecordsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:751
	_go_fuzz_dep_.CoverTab[104636]++
												return &MockDeleteRecordsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:752
	// _ = "end of CoverTab[104636]"
}

func (mr *MockDeleteRecordsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:755
	_go_fuzz_dep_.CoverTab[104637]++
												req := reqBody.(*DeleteRecordsRequest)
												res := &DeleteRecordsResponse{}
												res.Topics = make(map[string]*DeleteRecordsResponseTopic)

												for topic, deleteRecordRequestTopic := range req.Topics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:760
		_go_fuzz_dep_.CoverTab[104639]++
													partitions := make(map[int32]*DeleteRecordsResponsePartition)
													for partition := range deleteRecordRequestTopic.PartitionOffsets {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:762
			_go_fuzz_dep_.CoverTab[104641]++
														partitions[partition] = &DeleteRecordsResponsePartition{Err: ErrNoError}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:763
			// _ = "end of CoverTab[104641]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:764
		// _ = "end of CoverTab[104639]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:764
		_go_fuzz_dep_.CoverTab[104640]++
													res.Topics[topic] = &DeleteRecordsResponseTopic{Partitions: partitions}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:765
		// _ = "end of CoverTab[104640]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:766
	// _ = "end of CoverTab[104637]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:766
	_go_fuzz_dep_.CoverTab[104638]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:767
	// _ = "end of CoverTab[104638]"
}

type MockDescribeConfigsResponse struct {
	t TestReporter
}

func NewMockDescribeConfigsResponse(t TestReporter) *MockDescribeConfigsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:774
	_go_fuzz_dep_.CoverTab[104642]++
												return &MockDescribeConfigsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:775
	// _ = "end of CoverTab[104642]"
}

func (mr *MockDescribeConfigsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:778
	_go_fuzz_dep_.CoverTab[104643]++
												req := reqBody.(*DescribeConfigsRequest)
												res := &DescribeConfigsResponse{
		Version: req.Version,
	}

	includeSynonyms := req.Version > 0
	includeSource := req.Version > 0

	for _, r := range req.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:787
		_go_fuzz_dep_.CoverTab[104645]++
													var configEntries []*ConfigEntry
													switch r.Type {
		case BrokerResource:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:790
			_go_fuzz_dep_.CoverTab[104646]++
														configEntries = append(configEntries,
				&ConfigEntry{
					Name:		"min.insync.replicas",
					Value:		"2",
					ReadOnly:	false,
					Default:	false,
				},
			)
			res.Resources = append(res.Resources, &ResourceResponse{
				Name:		r.Name,
				Configs:	configEntries,
			})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:802
			// _ = "end of CoverTab[104646]"
		case BrokerLoggerResource:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:803
			_go_fuzz_dep_.CoverTab[104647]++
														configEntries = append(configEntries,
				&ConfigEntry{
					Name:		"kafka.controller.KafkaController",
					Value:		"DEBUG",
					ReadOnly:	false,
					Default:	false,
				},
			)
			res.Resources = append(res.Resources, &ResourceResponse{
				Name:		r.Name,
				Configs:	configEntries,
			})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:815
			// _ = "end of CoverTab[104647]"
		case TopicResource:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:816
			_go_fuzz_dep_.CoverTab[104648]++
														maxMessageBytes := &ConfigEntry{
				Name:		"max.message.bytes",
				Value:		"1000000",
				ReadOnly:	false,
				Default:	!includeSource,
				Sensitive:	false,
			}
			if includeSource {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:824
				_go_fuzz_dep_.CoverTab[104653]++
															maxMessageBytes.Source = SourceDefault
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:825
				// _ = "end of CoverTab[104653]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:826
				_go_fuzz_dep_.CoverTab[104654]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:826
				// _ = "end of CoverTab[104654]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:826
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:826
			// _ = "end of CoverTab[104648]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:826
			_go_fuzz_dep_.CoverTab[104649]++
														if includeSynonyms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:827
				_go_fuzz_dep_.CoverTab[104655]++
															maxMessageBytes.Synonyms = []*ConfigSynonym{
					{
						ConfigName:	"max.message.bytes",
						ConfigValue:	"500000",
					},
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:833
				// _ = "end of CoverTab[104655]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:834
				_go_fuzz_dep_.CoverTab[104656]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:834
				// _ = "end of CoverTab[104656]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:834
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:834
			// _ = "end of CoverTab[104649]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:834
			_go_fuzz_dep_.CoverTab[104650]++
														retentionMs := &ConfigEntry{
				Name:		"retention.ms",
				Value:		"5000",
				ReadOnly:	false,
				Default:	false,
				Sensitive:	false,
			}
			if includeSynonyms {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:842
				_go_fuzz_dep_.CoverTab[104657]++
															retentionMs.Synonyms = []*ConfigSynonym{
					{
						ConfigName:	"log.retention.ms",
						ConfigValue:	"2500",
					},
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:848
				// _ = "end of CoverTab[104657]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:849
				_go_fuzz_dep_.CoverTab[104658]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:849
				// _ = "end of CoverTab[104658]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:849
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:849
			// _ = "end of CoverTab[104650]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:849
			_go_fuzz_dep_.CoverTab[104651]++
														password := &ConfigEntry{
				Name:		"password",
				Value:		"12345",
				ReadOnly:	false,
				Default:	false,
				Sensitive:	true,
			}
			configEntries = append(
				configEntries, maxMessageBytes, retentionMs, password)
			res.Resources = append(res.Resources, &ResourceResponse{
				Name:		r.Name,
				Configs:	configEntries,
			})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:862
			// _ = "end of CoverTab[104651]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:862
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:862
			_go_fuzz_dep_.CoverTab[104652]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:862
			// _ = "end of CoverTab[104652]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:863
		// _ = "end of CoverTab[104645]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:864
	// _ = "end of CoverTab[104643]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:864
	_go_fuzz_dep_.CoverTab[104644]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:865
	// _ = "end of CoverTab[104644]"
}

type MockDescribeConfigsResponseWithErrorCode struct {
	t TestReporter
}

func NewMockDescribeConfigsResponseWithErrorCode(t TestReporter) *MockDescribeConfigsResponseWithErrorCode {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:872
	_go_fuzz_dep_.CoverTab[104659]++
												return &MockDescribeConfigsResponseWithErrorCode{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:873
	// _ = "end of CoverTab[104659]"
}

func (mr *MockDescribeConfigsResponseWithErrorCode) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:876
	_go_fuzz_dep_.CoverTab[104660]++
												req := reqBody.(*DescribeConfigsRequest)
												res := &DescribeConfigsResponse{
		Version: req.Version,
	}

	for _, r := range req.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:882
		_go_fuzz_dep_.CoverTab[104662]++
													res.Resources = append(res.Resources, &ResourceResponse{
			Name:		r.Name,
			Type:		r.Type,
			ErrorCode:	83,
			ErrorMsg:	"",
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:888
		// _ = "end of CoverTab[104662]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:889
	// _ = "end of CoverTab[104660]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:889
	_go_fuzz_dep_.CoverTab[104661]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:890
	// _ = "end of CoverTab[104661]"
}

type MockAlterConfigsResponse struct {
	t TestReporter
}

func NewMockAlterConfigsResponse(t TestReporter) *MockAlterConfigsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:897
	_go_fuzz_dep_.CoverTab[104663]++
												return &MockAlterConfigsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:898
	// _ = "end of CoverTab[104663]"
}

func (mr *MockAlterConfigsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:901
	_go_fuzz_dep_.CoverTab[104664]++
												req := reqBody.(*AlterConfigsRequest)
												res := &AlterConfigsResponse{}

												for _, r := range req.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:905
		_go_fuzz_dep_.CoverTab[104666]++
													res.Resources = append(res.Resources, &AlterConfigsResourceResponse{
			Name:		r.Name,
			Type:		r.Type,
			ErrorMsg:	"",
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:910
		// _ = "end of CoverTab[104666]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:911
	// _ = "end of CoverTab[104664]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:911
	_go_fuzz_dep_.CoverTab[104665]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:912
	// _ = "end of CoverTab[104665]"
}

type MockAlterConfigsResponseWithErrorCode struct {
	t TestReporter
}

func NewMockAlterConfigsResponseWithErrorCode(t TestReporter) *MockAlterConfigsResponseWithErrorCode {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:919
	_go_fuzz_dep_.CoverTab[104667]++
												return &MockAlterConfigsResponseWithErrorCode{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:920
	// _ = "end of CoverTab[104667]"
}

func (mr *MockAlterConfigsResponseWithErrorCode) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:923
	_go_fuzz_dep_.CoverTab[104668]++
												req := reqBody.(*AlterConfigsRequest)
												res := &AlterConfigsResponse{}

												for _, r := range req.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:927
		_go_fuzz_dep_.CoverTab[104670]++
													res.Resources = append(res.Resources, &AlterConfigsResourceResponse{
			Name:		r.Name,
			Type:		r.Type,
			ErrorCode:	83,
			ErrorMsg:	"",
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:933
		// _ = "end of CoverTab[104670]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:934
	// _ = "end of CoverTab[104668]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:934
	_go_fuzz_dep_.CoverTab[104669]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:935
	// _ = "end of CoverTab[104669]"
}

type MockIncrementalAlterConfigsResponse struct {
	t TestReporter
}

func NewMockIncrementalAlterConfigsResponse(t TestReporter) *MockIncrementalAlterConfigsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:942
	_go_fuzz_dep_.CoverTab[104671]++
												return &MockIncrementalAlterConfigsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:943
	// _ = "end of CoverTab[104671]"
}

func (mr *MockIncrementalAlterConfigsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:946
	_go_fuzz_dep_.CoverTab[104672]++
												req := reqBody.(*IncrementalAlterConfigsRequest)
												res := &IncrementalAlterConfigsResponse{}

												for _, r := range req.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:950
		_go_fuzz_dep_.CoverTab[104674]++
													res.Resources = append(res.Resources, &AlterConfigsResourceResponse{
			Name:		r.Name,
			Type:		r.Type,
			ErrorMsg:	"",
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:955
		// _ = "end of CoverTab[104674]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:956
	// _ = "end of CoverTab[104672]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:956
	_go_fuzz_dep_.CoverTab[104673]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:957
	// _ = "end of CoverTab[104673]"
}

type MockIncrementalAlterConfigsResponseWithErrorCode struct {
	t TestReporter
}

func NewMockIncrementalAlterConfigsResponseWithErrorCode(t TestReporter) *MockIncrementalAlterConfigsResponseWithErrorCode {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:964
	_go_fuzz_dep_.CoverTab[104675]++
												return &MockIncrementalAlterConfigsResponseWithErrorCode{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:965
	// _ = "end of CoverTab[104675]"
}

func (mr *MockIncrementalAlterConfigsResponseWithErrorCode) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:968
	_go_fuzz_dep_.CoverTab[104676]++
												req := reqBody.(*IncrementalAlterConfigsRequest)
												res := &IncrementalAlterConfigsResponse{}

												for _, r := range req.Resources {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:972
		_go_fuzz_dep_.CoverTab[104678]++
													res.Resources = append(res.Resources, &AlterConfigsResourceResponse{
			Name:		r.Name,
			Type:		r.Type,
			ErrorCode:	83,
			ErrorMsg:	"",
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:978
		// _ = "end of CoverTab[104678]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:979
	// _ = "end of CoverTab[104676]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:979
	_go_fuzz_dep_.CoverTab[104677]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:980
	// _ = "end of CoverTab[104677]"
}

type MockCreateAclsResponse struct {
	t TestReporter
}

func NewMockCreateAclsResponse(t TestReporter) *MockCreateAclsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:987
	_go_fuzz_dep_.CoverTab[104679]++
												return &MockCreateAclsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:988
	// _ = "end of CoverTab[104679]"
}

func (mr *MockCreateAclsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:991
	_go_fuzz_dep_.CoverTab[104680]++
												req := reqBody.(*CreateAclsRequest)
												res := &CreateAclsResponse{}

												for range req.AclCreations {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:995
		_go_fuzz_dep_.CoverTab[104682]++
													res.AclCreationResponses = append(res.AclCreationResponses, &AclCreationResponse{Err: ErrNoError})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:996
		// _ = "end of CoverTab[104682]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:997
	// _ = "end of CoverTab[104680]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:997
	_go_fuzz_dep_.CoverTab[104681]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:998
	// _ = "end of CoverTab[104681]"
}

type MockListAclsResponse struct {
	t TestReporter
}

func NewMockListAclsResponse(t TestReporter) *MockListAclsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1005
	_go_fuzz_dep_.CoverTab[104683]++
												return &MockListAclsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1006
	// _ = "end of CoverTab[104683]"
}

func (mr *MockListAclsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1009
	_go_fuzz_dep_.CoverTab[104684]++
												req := reqBody.(*DescribeAclsRequest)
												res := &DescribeAclsResponse{}
												res.Err = ErrNoError
												acl := &ResourceAcls{}
												if req.ResourceName != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1014
		_go_fuzz_dep_.CoverTab[104689]++
													acl.Resource.ResourceName = *req.ResourceName
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1015
		// _ = "end of CoverTab[104689]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1016
		_go_fuzz_dep_.CoverTab[104690]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1016
		// _ = "end of CoverTab[104690]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1016
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1016
	// _ = "end of CoverTab[104684]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1016
	_go_fuzz_dep_.CoverTab[104685]++
												acl.Resource.ResourcePatternType = req.ResourcePatternTypeFilter
												acl.Resource.ResourceType = req.ResourceType

												host := "*"
												if req.Host != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1021
		_go_fuzz_dep_.CoverTab[104691]++
													host = *req.Host
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1022
		// _ = "end of CoverTab[104691]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1023
		_go_fuzz_dep_.CoverTab[104692]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1023
		// _ = "end of CoverTab[104692]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1023
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1023
	// _ = "end of CoverTab[104685]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1023
	_go_fuzz_dep_.CoverTab[104686]++

												principal := "User:test"
												if req.Principal != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1026
		_go_fuzz_dep_.CoverTab[104693]++
													principal = *req.Principal
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1027
		// _ = "end of CoverTab[104693]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1028
		_go_fuzz_dep_.CoverTab[104694]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1028
		// _ = "end of CoverTab[104694]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1028
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1028
	// _ = "end of CoverTab[104686]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1028
	_go_fuzz_dep_.CoverTab[104687]++

												permissionType := req.PermissionType
												if permissionType == AclPermissionAny {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1031
		_go_fuzz_dep_.CoverTab[104695]++
													permissionType = AclPermissionAllow
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1032
		// _ = "end of CoverTab[104695]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1033
		_go_fuzz_dep_.CoverTab[104696]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1033
		// _ = "end of CoverTab[104696]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1033
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1033
	// _ = "end of CoverTab[104687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1033
	_go_fuzz_dep_.CoverTab[104688]++

												acl.Acls = append(acl.Acls, &Acl{Operation: req.Operation, PermissionType: permissionType, Host: host, Principal: principal})
												res.ResourceAcls = append(res.ResourceAcls, acl)
												res.Version = int16(req.Version)
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1038
	// _ = "end of CoverTab[104688]"
}

type MockSaslAuthenticateResponse struct {
	t		TestReporter
	kerror		KError
	saslAuthBytes	[]byte
}

func NewMockSaslAuthenticateResponse(t TestReporter) *MockSaslAuthenticateResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1047
	_go_fuzz_dep_.CoverTab[104697]++
												return &MockSaslAuthenticateResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1048
	// _ = "end of CoverTab[104697]"
}

func (msar *MockSaslAuthenticateResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1051
	_go_fuzz_dep_.CoverTab[104698]++
												res := &SaslAuthenticateResponse{}
												res.Err = msar.kerror
												res.SaslAuthBytes = msar.saslAuthBytes
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1055
	// _ = "end of CoverTab[104698]"
}

func (msar *MockSaslAuthenticateResponse) SetError(kerror KError) *MockSaslAuthenticateResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1058
	_go_fuzz_dep_.CoverTab[104699]++
												msar.kerror = kerror
												return msar
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1060
	// _ = "end of CoverTab[104699]"
}

func (msar *MockSaslAuthenticateResponse) SetAuthBytes(saslAuthBytes []byte) *MockSaslAuthenticateResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1063
	_go_fuzz_dep_.CoverTab[104700]++
												msar.saslAuthBytes = saslAuthBytes
												return msar
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1065
	// _ = "end of CoverTab[104700]"
}

type MockDeleteAclsResponse struct {
	t TestReporter
}

type MockSaslHandshakeResponse struct {
	enabledMechanisms	[]string
	kerror			KError
	t			TestReporter
}

func NewMockSaslHandshakeResponse(t TestReporter) *MockSaslHandshakeResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1078
	_go_fuzz_dep_.CoverTab[104701]++
												return &MockSaslHandshakeResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1079
	// _ = "end of CoverTab[104701]"
}

func (mshr *MockSaslHandshakeResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1082
	_go_fuzz_dep_.CoverTab[104702]++
												res := &SaslHandshakeResponse{}
												res.Err = mshr.kerror
												res.EnabledMechanisms = mshr.enabledMechanisms
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1086
	// _ = "end of CoverTab[104702]"
}

func (mshr *MockSaslHandshakeResponse) SetError(kerror KError) *MockSaslHandshakeResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1089
	_go_fuzz_dep_.CoverTab[104703]++
												mshr.kerror = kerror
												return mshr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1091
	// _ = "end of CoverTab[104703]"
}

func (mshr *MockSaslHandshakeResponse) SetEnabledMechanisms(enabledMechanisms []string) *MockSaslHandshakeResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1094
	_go_fuzz_dep_.CoverTab[104704]++
												mshr.enabledMechanisms = enabledMechanisms
												return mshr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1096
	// _ = "end of CoverTab[104704]"
}

func NewMockDeleteAclsResponse(t TestReporter) *MockDeleteAclsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1099
	_go_fuzz_dep_.CoverTab[104705]++
												return &MockDeleteAclsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1100
	// _ = "end of CoverTab[104705]"
}

func (mr *MockDeleteAclsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1103
	_go_fuzz_dep_.CoverTab[104706]++
												req := reqBody.(*DeleteAclsRequest)
												res := &DeleteAclsResponse{}

												for range req.Filters {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1107
		_go_fuzz_dep_.CoverTab[104708]++
													response := &FilterResponse{Err: ErrNoError}
													response.MatchingAcls = append(response.MatchingAcls, &MatchingAcl{Err: ErrNoError})
													res.FilterResponses = append(res.FilterResponses, response)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1110
		// _ = "end of CoverTab[104708]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1111
	// _ = "end of CoverTab[104706]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1111
	_go_fuzz_dep_.CoverTab[104707]++
												res.Version = int16(req.Version)
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1113
	// _ = "end of CoverTab[104707]"
}

type MockDeleteGroupsResponse struct {
	deletedGroups []string
}

func NewMockDeleteGroupsRequest(t TestReporter) *MockDeleteGroupsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1120
	_go_fuzz_dep_.CoverTab[104709]++
												return &MockDeleteGroupsResponse{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1121
	// _ = "end of CoverTab[104709]"
}

func (m *MockDeleteGroupsResponse) SetDeletedGroups(groups []string) *MockDeleteGroupsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1124
	_go_fuzz_dep_.CoverTab[104710]++
												m.deletedGroups = groups
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1126
	// _ = "end of CoverTab[104710]"
}

func (m *MockDeleteGroupsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1129
	_go_fuzz_dep_.CoverTab[104711]++
												resp := &DeleteGroupsResponse{
		GroupErrorCodes: map[string]KError{},
	}
	for _, group := range m.deletedGroups {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1133
		_go_fuzz_dep_.CoverTab[104713]++
													resp.GroupErrorCodes[group] = ErrNoError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1134
		// _ = "end of CoverTab[104713]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1135
	// _ = "end of CoverTab[104711]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1135
	_go_fuzz_dep_.CoverTab[104712]++
												return resp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1136
	// _ = "end of CoverTab[104712]"
}

type MockDeleteOffsetResponse struct {
	errorCode	KError
	topic		string
	partition	int32
	errorPartition	KError
}

func NewMockDeleteOffsetRequest(t TestReporter) *MockDeleteOffsetResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1146
	_go_fuzz_dep_.CoverTab[104714]++
												return &MockDeleteOffsetResponse{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1147
	// _ = "end of CoverTab[104714]"
}

func (m *MockDeleteOffsetResponse) SetDeletedOffset(errorCode KError, topic string, partition int32, errorPartition KError) *MockDeleteOffsetResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1150
	_go_fuzz_dep_.CoverTab[104715]++
												m.errorCode = errorCode
												m.topic = topic
												m.partition = partition
												m.errorPartition = errorPartition
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1155
	// _ = "end of CoverTab[104715]"
}

func (m *MockDeleteOffsetResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1158
	_go_fuzz_dep_.CoverTab[104716]++
												resp := &DeleteOffsetsResponse{
		ErrorCode:	m.errorCode,
		Errors: map[string]map[int32]KError{
			m.topic: {m.partition: m.errorPartition},
		},
	}
												return resp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1165
	// _ = "end of CoverTab[104716]"
}

type MockJoinGroupResponse struct {
	t	TestReporter

	ThrottleTime	int32
	Err		KError
	GenerationId	int32
	GroupProtocol	string
	LeaderId	string
	MemberId	string
	Members		map[string][]byte
}

func NewMockJoinGroupResponse(t TestReporter) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1180
	_go_fuzz_dep_.CoverTab[104717]++
												return &MockJoinGroupResponse{
		t:		t,
		Members:	make(map[string][]byte),
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1184
	// _ = "end of CoverTab[104717]"
}

func (m *MockJoinGroupResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1187
	_go_fuzz_dep_.CoverTab[104718]++
												req := reqBody.(*JoinGroupRequest)
												resp := &JoinGroupResponse{
		Version:	req.Version,
		ThrottleTime:	m.ThrottleTime,
		Err:		m.Err,
		GenerationId:	m.GenerationId,
		GroupProtocol:	m.GroupProtocol,
		LeaderId:	m.LeaderId,
		MemberId:	m.MemberId,
		Members:	m.Members,
	}
												return resp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1199
	// _ = "end of CoverTab[104718]"
}

func (m *MockJoinGroupResponse) SetThrottleTime(t int32) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1202
	_go_fuzz_dep_.CoverTab[104719]++
												m.ThrottleTime = t
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1204
	// _ = "end of CoverTab[104719]"
}

func (m *MockJoinGroupResponse) SetError(kerr KError) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1207
	_go_fuzz_dep_.CoverTab[104720]++
												m.Err = kerr
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1209
	// _ = "end of CoverTab[104720]"
}

func (m *MockJoinGroupResponse) SetGenerationId(id int32) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1212
	_go_fuzz_dep_.CoverTab[104721]++
												m.GenerationId = id
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1214
	// _ = "end of CoverTab[104721]"
}

func (m *MockJoinGroupResponse) SetGroupProtocol(proto string) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1217
	_go_fuzz_dep_.CoverTab[104722]++
												m.GroupProtocol = proto
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1219
	// _ = "end of CoverTab[104722]"
}

func (m *MockJoinGroupResponse) SetLeaderId(id string) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1222
	_go_fuzz_dep_.CoverTab[104723]++
												m.LeaderId = id
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1224
	// _ = "end of CoverTab[104723]"
}

func (m *MockJoinGroupResponse) SetMemberId(id string) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1227
	_go_fuzz_dep_.CoverTab[104724]++
												m.MemberId = id
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1229
	// _ = "end of CoverTab[104724]"
}

func (m *MockJoinGroupResponse) SetMember(id string, meta *ConsumerGroupMemberMetadata) *MockJoinGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1232
	_go_fuzz_dep_.CoverTab[104725]++
												bin, err := encode(meta, nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1234
		_go_fuzz_dep_.CoverTab[104727]++
													panic(fmt.Sprintf("error encoding member metadata: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1235
		// _ = "end of CoverTab[104727]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1236
		_go_fuzz_dep_.CoverTab[104728]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1236
		// _ = "end of CoverTab[104728]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1236
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1236
	// _ = "end of CoverTab[104725]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1236
	_go_fuzz_dep_.CoverTab[104726]++
												m.Members[id] = bin
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1238
	// _ = "end of CoverTab[104726]"
}

type MockLeaveGroupResponse struct {
	t	TestReporter

	Err	KError
}

func NewMockLeaveGroupResponse(t TestReporter) *MockLeaveGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1247
	_go_fuzz_dep_.CoverTab[104729]++
												return &MockLeaveGroupResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1248
	// _ = "end of CoverTab[104729]"
}

func (m *MockLeaveGroupResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1251
	_go_fuzz_dep_.CoverTab[104730]++
												resp := &LeaveGroupResponse{
		Err: m.Err,
	}
												return resp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1255
	// _ = "end of CoverTab[104730]"
}

func (m *MockLeaveGroupResponse) SetError(kerr KError) *MockLeaveGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1258
	_go_fuzz_dep_.CoverTab[104731]++
												m.Err = kerr
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1260
	// _ = "end of CoverTab[104731]"
}

type MockSyncGroupResponse struct {
	t	TestReporter

	Err			KError
	MemberAssignment	[]byte
}

func NewMockSyncGroupResponse(t TestReporter) *MockSyncGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1270
	_go_fuzz_dep_.CoverTab[104732]++
												return &MockSyncGroupResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1271
	// _ = "end of CoverTab[104732]"
}

func (m *MockSyncGroupResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1274
	_go_fuzz_dep_.CoverTab[104733]++
												resp := &SyncGroupResponse{
		Err:			m.Err,
		MemberAssignment:	m.MemberAssignment,
	}
												return resp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1279
	// _ = "end of CoverTab[104733]"
}

func (m *MockSyncGroupResponse) SetError(kerr KError) *MockSyncGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1282
	_go_fuzz_dep_.CoverTab[104734]++
												m.Err = kerr
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1284
	// _ = "end of CoverTab[104734]"
}

func (m *MockSyncGroupResponse) SetMemberAssignment(assignment *ConsumerGroupMemberAssignment) *MockSyncGroupResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1287
	_go_fuzz_dep_.CoverTab[104735]++
												bin, err := encode(assignment, nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1289
		_go_fuzz_dep_.CoverTab[104737]++
													panic(fmt.Sprintf("error encoding member assignment: %v", err))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1290
		// _ = "end of CoverTab[104737]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1291
		_go_fuzz_dep_.CoverTab[104738]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1291
		// _ = "end of CoverTab[104738]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1291
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1291
	// _ = "end of CoverTab[104735]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1291
	_go_fuzz_dep_.CoverTab[104736]++
												m.MemberAssignment = bin
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1293
	// _ = "end of CoverTab[104736]"
}

type MockHeartbeatResponse struct {
	t	TestReporter

	Err	KError
}

func NewMockHeartbeatResponse(t TestReporter) *MockHeartbeatResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1302
	_go_fuzz_dep_.CoverTab[104739]++
												return &MockHeartbeatResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1303
	// _ = "end of CoverTab[104739]"
}

func (m *MockHeartbeatResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1306
	_go_fuzz_dep_.CoverTab[104740]++
												resp := &HeartbeatResponse{}
												return resp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1308
	// _ = "end of CoverTab[104740]"
}

func (m *MockHeartbeatResponse) SetError(kerr KError) *MockHeartbeatResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1311
	_go_fuzz_dep_.CoverTab[104741]++
												m.Err = kerr
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1313
	// _ = "end of CoverTab[104741]"
}

type MockDescribeLogDirsResponse struct {
	t	TestReporter
	logDirs	[]DescribeLogDirsResponseDirMetadata
}

func NewMockDescribeLogDirsResponse(t TestReporter) *MockDescribeLogDirsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1321
	_go_fuzz_dep_.CoverTab[104742]++
												return &MockDescribeLogDirsResponse{t: t}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1322
	// _ = "end of CoverTab[104742]"
}

func (m *MockDescribeLogDirsResponse) SetLogDirs(logDirPath string, topicPartitions map[string]int) *MockDescribeLogDirsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1325
	_go_fuzz_dep_.CoverTab[104743]++
												var topics []DescribeLogDirsResponseTopic
												for topic := range topicPartitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1327
		_go_fuzz_dep_.CoverTab[104745]++
													var partitions []DescribeLogDirsResponsePartition
													for i := 0; i < topicPartitions[topic]; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1329
			_go_fuzz_dep_.CoverTab[104747]++
														partitions = append(partitions, DescribeLogDirsResponsePartition{
				PartitionID:	int32(i),
				IsTemporary:	false,
				OffsetLag:	int64(0),
				Size:		int64(1234),
			})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1335
			// _ = "end of CoverTab[104747]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1336
		// _ = "end of CoverTab[104745]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1336
		_go_fuzz_dep_.CoverTab[104746]++
													topics = append(topics, DescribeLogDirsResponseTopic{
			Topic:		topic,
			Partitions:	partitions,
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1340
		// _ = "end of CoverTab[104746]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1341
	// _ = "end of CoverTab[104743]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1341
	_go_fuzz_dep_.CoverTab[104744]++
												logDir := DescribeLogDirsResponseDirMetadata{
		ErrorCode:	ErrNoError,
		Path:		logDirPath,
		Topics:		topics,
	}
												m.logDirs = []DescribeLogDirsResponseDirMetadata{logDir}
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1348
	// _ = "end of CoverTab[104744]"
}

func (m *MockDescribeLogDirsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1351
	_go_fuzz_dep_.CoverTab[104748]++
												resp := &DescribeLogDirsResponse{
		LogDirs: m.logDirs,
	}
												return resp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1355
	// _ = "end of CoverTab[104748]"
}

type MockApiVersionsResponse struct {
	t	TestReporter
	apiKeys	[]ApiVersionsResponseKey
}

func NewMockApiVersionsResponse(t TestReporter) *MockApiVersionsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1363
	_go_fuzz_dep_.CoverTab[104749]++
												return &MockApiVersionsResponse{
		t:	t,
		apiKeys: []ApiVersionsResponseKey{
			{
				ApiKey:		0,
				MinVersion:	5,
				MaxVersion:	8,
			},
			{
				ApiKey:		1,
				MinVersion:	7,
				MaxVersion:	11,
			},
		},
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1378
	// _ = "end of CoverTab[104749]"
}

func (m *MockApiVersionsResponse) SetApiKeys(apiKeys []ApiVersionsResponseKey) *MockApiVersionsResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1381
	_go_fuzz_dep_.CoverTab[104750]++
												m.apiKeys = apiKeys
												return m
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1383
	// _ = "end of CoverTab[104750]"
}

func (m *MockApiVersionsResponse) For(reqBody versionedDecoder) encoderWithHeader {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1386
	_go_fuzz_dep_.CoverTab[104751]++
												req := reqBody.(*ApiVersionsRequest)
												res := &ApiVersionsResponse{
		Version:	req.Version,
		ApiKeys:	m.apiKeys,
	}
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1392
	// _ = "end of CoverTab[104751]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1393
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockresponses.go:1393
var _ = _go_fuzz_dep_.CoverTab
