//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:1
)

import (
	"encoding/binary"
	"fmt"
	"io"
)

type protocolBody interface {
	encoder
	versionedDecoder
	key() int16
	version() int16
	headerVersion() int16
	requiredVersion() KafkaVersion
}

type request struct {
	correlationID	int32
	clientID	string
	body		protocolBody
}

func (r *request) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:24
	_go_fuzz_dep_.CoverTab[106543]++
											pe.push(&lengthField{})
											pe.putInt16(r.body.key())
											pe.putInt16(r.body.version())
											pe.putInt32(r.correlationID)

											if r.body.headerVersion() >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:30
		_go_fuzz_dep_.CoverTab[106547]++
												err := pe.putString(r.clientID)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:32
			_go_fuzz_dep_.CoverTab[106548]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:33
			// _ = "end of CoverTab[106548]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:34
			_go_fuzz_dep_.CoverTab[106549]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:34
			// _ = "end of CoverTab[106549]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:34
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:34
		// _ = "end of CoverTab[106547]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:35
		_go_fuzz_dep_.CoverTab[106550]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:35
		// _ = "end of CoverTab[106550]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:35
	// _ = "end of CoverTab[106543]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:35
	_go_fuzz_dep_.CoverTab[106544]++

											if r.body.headerVersion() >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:37
		_go_fuzz_dep_.CoverTab[106551]++

												pe.putUVarint(0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:39
		// _ = "end of CoverTab[106551]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:40
		_go_fuzz_dep_.CoverTab[106552]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:40
		// _ = "end of CoverTab[106552]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:40
	// _ = "end of CoverTab[106544]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:40
	_go_fuzz_dep_.CoverTab[106545]++

											err := r.body.encode(pe)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:43
		_go_fuzz_dep_.CoverTab[106553]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:44
		// _ = "end of CoverTab[106553]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:45
		_go_fuzz_dep_.CoverTab[106554]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:45
		// _ = "end of CoverTab[106554]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:45
	// _ = "end of CoverTab[106545]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:45
	_go_fuzz_dep_.CoverTab[106546]++

											return pe.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:47
	// _ = "end of CoverTab[106546]"
}

func (r *request) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:50
	_go_fuzz_dep_.CoverTab[106555]++
											key, err := pd.getInt16()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:52
		_go_fuzz_dep_.CoverTab[106562]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:53
		// _ = "end of CoverTab[106562]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:54
		_go_fuzz_dep_.CoverTab[106563]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:54
		// _ = "end of CoverTab[106563]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:54
	// _ = "end of CoverTab[106555]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:54
	_go_fuzz_dep_.CoverTab[106556]++

											version, err := pd.getInt16()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:57
		_go_fuzz_dep_.CoverTab[106564]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:58
		// _ = "end of CoverTab[106564]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:59
		_go_fuzz_dep_.CoverTab[106565]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:59
		// _ = "end of CoverTab[106565]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:59
	// _ = "end of CoverTab[106556]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:59
	_go_fuzz_dep_.CoverTab[106557]++

											r.correlationID, err = pd.getInt32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:62
		_go_fuzz_dep_.CoverTab[106566]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:63
		// _ = "end of CoverTab[106566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:64
		_go_fuzz_dep_.CoverTab[106567]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:64
		// _ = "end of CoverTab[106567]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:64
	// _ = "end of CoverTab[106557]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:64
	_go_fuzz_dep_.CoverTab[106558]++

											r.clientID, err = pd.getString()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:67
		_go_fuzz_dep_.CoverTab[106568]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:68
		// _ = "end of CoverTab[106568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:69
		_go_fuzz_dep_.CoverTab[106569]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:69
		// _ = "end of CoverTab[106569]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:69
	// _ = "end of CoverTab[106558]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:69
	_go_fuzz_dep_.CoverTab[106559]++

											r.body = allocateBody(key, version)
											if r.body == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:72
		_go_fuzz_dep_.CoverTab[106570]++
												return PacketDecodingError{fmt.Sprintf("unknown request key (%d)", key)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:73
		// _ = "end of CoverTab[106570]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:74
		_go_fuzz_dep_.CoverTab[106571]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:74
		// _ = "end of CoverTab[106571]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:74
	// _ = "end of CoverTab[106559]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:74
	_go_fuzz_dep_.CoverTab[106560]++

											if r.body.headerVersion() >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:76
		_go_fuzz_dep_.CoverTab[106572]++

												_, err = pd.getUVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:79
			_go_fuzz_dep_.CoverTab[106573]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:80
			// _ = "end of CoverTab[106573]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:81
			_go_fuzz_dep_.CoverTab[106574]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:81
			// _ = "end of CoverTab[106574]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:81
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:81
		// _ = "end of CoverTab[106572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:82
		_go_fuzz_dep_.CoverTab[106575]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:82
		// _ = "end of CoverTab[106575]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:82
	// _ = "end of CoverTab[106560]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:82
	_go_fuzz_dep_.CoverTab[106561]++

											return r.body.decode(pd, version)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:84
	// _ = "end of CoverTab[106561]"
}

func decodeRequest(r io.Reader) (*request, int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:87
	_go_fuzz_dep_.CoverTab[106576]++
											var (
		bytesRead	int
		lengthBytes	= make([]byte, 4)
	)

	if _, err := io.ReadFull(r, lengthBytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:93
		_go_fuzz_dep_.CoverTab[106581]++
												return nil, bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:94
		// _ = "end of CoverTab[106581]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:95
		_go_fuzz_dep_.CoverTab[106582]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:95
		// _ = "end of CoverTab[106582]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:95
	// _ = "end of CoverTab[106576]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:95
	_go_fuzz_dep_.CoverTab[106577]++

											bytesRead += len(lengthBytes)
											length := int32(binary.BigEndian.Uint32(lengthBytes))

											if length <= 4 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:100
		_go_fuzz_dep_.CoverTab[106583]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:100
		return length > MaxRequestSize
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:100
		// _ = "end of CoverTab[106583]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:100
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:100
		_go_fuzz_dep_.CoverTab[106584]++
												return nil, bytesRead, PacketDecodingError{fmt.Sprintf("message of length %d too large or too small", length)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:101
		// _ = "end of CoverTab[106584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:102
		_go_fuzz_dep_.CoverTab[106585]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:102
		// _ = "end of CoverTab[106585]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:102
	// _ = "end of CoverTab[106577]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:102
	_go_fuzz_dep_.CoverTab[106578]++

											encodedReq := make([]byte, length)
											if _, err := io.ReadFull(r, encodedReq); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:105
		_go_fuzz_dep_.CoverTab[106586]++
												return nil, bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:106
		// _ = "end of CoverTab[106586]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:107
		_go_fuzz_dep_.CoverTab[106587]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:107
		// _ = "end of CoverTab[106587]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:107
	// _ = "end of CoverTab[106578]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:107
	_go_fuzz_dep_.CoverTab[106579]++

											bytesRead += len(encodedReq)

											req := &request{}
											if err := decode(encodedReq, req); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:112
		_go_fuzz_dep_.CoverTab[106588]++
												return nil, bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:113
		// _ = "end of CoverTab[106588]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:114
		_go_fuzz_dep_.CoverTab[106589]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:114
		// _ = "end of CoverTab[106589]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:114
	// _ = "end of CoverTab[106579]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:114
	_go_fuzz_dep_.CoverTab[106580]++

											return req, bytesRead, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:116
	// _ = "end of CoverTab[106580]"
}

func allocateBody(key, version int16) protocolBody {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:119
	_go_fuzz_dep_.CoverTab[106590]++
											switch key {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:121
		_go_fuzz_dep_.CoverTab[106592]++
												return &ProduceRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:122
		// _ = "end of CoverTab[106592]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:123
		_go_fuzz_dep_.CoverTab[106593]++
												return &FetchRequest{Version: version}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:124
		// _ = "end of CoverTab[106593]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:125
		_go_fuzz_dep_.CoverTab[106594]++
												return &OffsetRequest{Version: version}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:126
		// _ = "end of CoverTab[106594]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:127
		_go_fuzz_dep_.CoverTab[106595]++
												return &MetadataRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:128
		// _ = "end of CoverTab[106595]"
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:129
		_go_fuzz_dep_.CoverTab[106596]++
												return &OffsetCommitRequest{Version: version}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:130
		// _ = "end of CoverTab[106596]"
	case 9:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:131
		_go_fuzz_dep_.CoverTab[106597]++
												return &OffsetFetchRequest{Version: version}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:132
		// _ = "end of CoverTab[106597]"
	case 10:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:133
		_go_fuzz_dep_.CoverTab[106598]++
												return &FindCoordinatorRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:134
		// _ = "end of CoverTab[106598]"
	case 11:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:135
		_go_fuzz_dep_.CoverTab[106599]++
												return &JoinGroupRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:136
		// _ = "end of CoverTab[106599]"
	case 12:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:137
		_go_fuzz_dep_.CoverTab[106600]++
												return &HeartbeatRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:138
		// _ = "end of CoverTab[106600]"
	case 13:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:139
		_go_fuzz_dep_.CoverTab[106601]++
												return &LeaveGroupRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:140
		// _ = "end of CoverTab[106601]"
	case 14:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:141
		_go_fuzz_dep_.CoverTab[106602]++
												return &SyncGroupRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:142
		// _ = "end of CoverTab[106602]"
	case 15:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:143
		_go_fuzz_dep_.CoverTab[106603]++
												return &DescribeGroupsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:144
		// _ = "end of CoverTab[106603]"
	case 16:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:145
		_go_fuzz_dep_.CoverTab[106604]++
												return &ListGroupsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:146
		// _ = "end of CoverTab[106604]"
	case 17:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:147
		_go_fuzz_dep_.CoverTab[106605]++
												return &SaslHandshakeRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:148
		// _ = "end of CoverTab[106605]"
	case 18:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:149
		_go_fuzz_dep_.CoverTab[106606]++
												return &ApiVersionsRequest{Version: version}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:150
		// _ = "end of CoverTab[106606]"
	case 19:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:151
		_go_fuzz_dep_.CoverTab[106607]++
												return &CreateTopicsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:152
		// _ = "end of CoverTab[106607]"
	case 20:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:153
		_go_fuzz_dep_.CoverTab[106608]++
												return &DeleteTopicsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:154
		// _ = "end of CoverTab[106608]"
	case 21:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:155
		_go_fuzz_dep_.CoverTab[106609]++
												return &DeleteRecordsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:156
		// _ = "end of CoverTab[106609]"
	case 22:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:157
		_go_fuzz_dep_.CoverTab[106610]++
												return &InitProducerIDRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:158
		// _ = "end of CoverTab[106610]"
	case 24:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:159
		_go_fuzz_dep_.CoverTab[106611]++
												return &AddPartitionsToTxnRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:160
		// _ = "end of CoverTab[106611]"
	case 25:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:161
		_go_fuzz_dep_.CoverTab[106612]++
												return &AddOffsetsToTxnRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:162
		// _ = "end of CoverTab[106612]"
	case 26:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:163
		_go_fuzz_dep_.CoverTab[106613]++
												return &EndTxnRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:164
		// _ = "end of CoverTab[106613]"
	case 28:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:165
		_go_fuzz_dep_.CoverTab[106614]++
												return &TxnOffsetCommitRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:166
		// _ = "end of CoverTab[106614]"
	case 29:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:167
		_go_fuzz_dep_.CoverTab[106615]++
												return &DescribeAclsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:168
		// _ = "end of CoverTab[106615]"
	case 30:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:169
		_go_fuzz_dep_.CoverTab[106616]++
												return &CreateAclsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:170
		// _ = "end of CoverTab[106616]"
	case 31:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:171
		_go_fuzz_dep_.CoverTab[106617]++
												return &DeleteAclsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:172
		// _ = "end of CoverTab[106617]"
	case 32:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:173
		_go_fuzz_dep_.CoverTab[106618]++
												return &DescribeConfigsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:174
		// _ = "end of CoverTab[106618]"
	case 33:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:175
		_go_fuzz_dep_.CoverTab[106619]++
												return &AlterConfigsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:176
		// _ = "end of CoverTab[106619]"
	case 35:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:177
		_go_fuzz_dep_.CoverTab[106620]++
												return &DescribeLogDirsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:178
		// _ = "end of CoverTab[106620]"
	case 36:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:179
		_go_fuzz_dep_.CoverTab[106621]++
												return &SaslAuthenticateRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:180
		// _ = "end of CoverTab[106621]"
	case 37:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:181
		_go_fuzz_dep_.CoverTab[106622]++
												return &CreatePartitionsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:182
		// _ = "end of CoverTab[106622]"
	case 42:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:183
		_go_fuzz_dep_.CoverTab[106623]++
												return &DeleteGroupsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:184
		// _ = "end of CoverTab[106623]"
	case 44:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:185
		_go_fuzz_dep_.CoverTab[106624]++
												return &IncrementalAlterConfigsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:186
		// _ = "end of CoverTab[106624]"
	case 45:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:187
		_go_fuzz_dep_.CoverTab[106625]++
												return &AlterPartitionReassignmentsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:188
		// _ = "end of CoverTab[106625]"
	case 46:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:189
		_go_fuzz_dep_.CoverTab[106626]++
												return &ListPartitionReassignmentsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:190
		// _ = "end of CoverTab[106626]"
	case 47:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:191
		_go_fuzz_dep_.CoverTab[106627]++
												return &DeleteOffsetsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:192
		// _ = "end of CoverTab[106627]"
	case 48:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:193
		_go_fuzz_dep_.CoverTab[106628]++
												return &DescribeClientQuotasRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:194
		// _ = "end of CoverTab[106628]"
	case 49:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:195
		_go_fuzz_dep_.CoverTab[106629]++
												return &AlterClientQuotasRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:196
		// _ = "end of CoverTab[106629]"
	case 50:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:197
		_go_fuzz_dep_.CoverTab[106630]++
												return &DescribeUserScramCredentialsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:198
		// _ = "end of CoverTab[106630]"
	case 51:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:199
		_go_fuzz_dep_.CoverTab[106631]++
												return &AlterUserScramCredentialsRequest{}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:200
		// _ = "end of CoverTab[106631]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:200
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:200
		_go_fuzz_dep_.CoverTab[106632]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:200
		// _ = "end of CoverTab[106632]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:201
	// _ = "end of CoverTab[106590]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:201
	_go_fuzz_dep_.CoverTab[106591]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:202
	// _ = "end of CoverTab[106591]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:203
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/request.go:203
var _ = _go_fuzz_dep_.CoverTab
