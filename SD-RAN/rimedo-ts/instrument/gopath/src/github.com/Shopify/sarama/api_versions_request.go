//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:1
)

const defaultClientSoftwareName = "sarama"

type ApiVersionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version	int16
	// ClientSoftwareName contains the name of the client.
	ClientSoftwareName	string
	// ClientSoftwareVersion contains the version of the client.
	ClientSoftwareVersion	string
}

func (r *ApiVersionsRequest) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:14
	_go_fuzz_dep_.CoverTab[98440]++
													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:15
		_go_fuzz_dep_.CoverTab[98442]++
														if err := pe.putCompactString(r.ClientSoftwareName); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:16
			_go_fuzz_dep_.CoverTab[98445]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:17
			// _ = "end of CoverTab[98445]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:18
			_go_fuzz_dep_.CoverTab[98446]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:18
			// _ = "end of CoverTab[98446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:18
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:18
		// _ = "end of CoverTab[98442]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:18
		_go_fuzz_dep_.CoverTab[98443]++
														if err := pe.putCompactString(r.ClientSoftwareVersion); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:19
			_go_fuzz_dep_.CoverTab[98447]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:20
			// _ = "end of CoverTab[98447]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:21
			_go_fuzz_dep_.CoverTab[98448]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:21
			// _ = "end of CoverTab[98448]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:21
		// _ = "end of CoverTab[98443]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:21
		_go_fuzz_dep_.CoverTab[98444]++
														pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:22
		// _ = "end of CoverTab[98444]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:23
		_go_fuzz_dep_.CoverTab[98449]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:23
		// _ = "end of CoverTab[98449]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:23
	// _ = "end of CoverTab[98440]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:23
	_go_fuzz_dep_.CoverTab[98441]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:25
	// _ = "end of CoverTab[98441]"
}

func (r *ApiVersionsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:28
	_go_fuzz_dep_.CoverTab[98450]++
													r.Version = version
													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:30
		_go_fuzz_dep_.CoverTab[98452]++
														if r.ClientSoftwareName, err = pd.getCompactString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:31
			_go_fuzz_dep_.CoverTab[98455]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:32
			// _ = "end of CoverTab[98455]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:33
			_go_fuzz_dep_.CoverTab[98456]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:33
			// _ = "end of CoverTab[98456]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:33
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:33
		// _ = "end of CoverTab[98452]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:33
		_go_fuzz_dep_.CoverTab[98453]++
														if r.ClientSoftwareVersion, err = pd.getCompactString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:34
			_go_fuzz_dep_.CoverTab[98457]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:35
			// _ = "end of CoverTab[98457]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:36
			_go_fuzz_dep_.CoverTab[98458]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:36
			// _ = "end of CoverTab[98458]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:36
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:36
		// _ = "end of CoverTab[98453]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:36
		_go_fuzz_dep_.CoverTab[98454]++
														if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:37
			_go_fuzz_dep_.CoverTab[98459]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:38
			// _ = "end of CoverTab[98459]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:39
			_go_fuzz_dep_.CoverTab[98460]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:39
			// _ = "end of CoverTab[98460]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:39
		// _ = "end of CoverTab[98454]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:40
		_go_fuzz_dep_.CoverTab[98461]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:40
		// _ = "end of CoverTab[98461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:40
	// _ = "end of CoverTab[98450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:40
	_go_fuzz_dep_.CoverTab[98451]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:42
	// _ = "end of CoverTab[98451]"
}

func (r *ApiVersionsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:45
	_go_fuzz_dep_.CoverTab[98462]++
													return 18
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:46
	// _ = "end of CoverTab[98462]"
}

func (r *ApiVersionsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:49
	_go_fuzz_dep_.CoverTab[98463]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:50
	// _ = "end of CoverTab[98463]"
}

func (r *ApiVersionsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:53
	_go_fuzz_dep_.CoverTab[98464]++
													if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:54
		_go_fuzz_dep_.CoverTab[98466]++
														return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:55
		// _ = "end of CoverTab[98466]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:56
		_go_fuzz_dep_.CoverTab[98467]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:56
		// _ = "end of CoverTab[98467]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:56
	// _ = "end of CoverTab[98464]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:56
	_go_fuzz_dep_.CoverTab[98465]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:57
	// _ = "end of CoverTab[98465]"
}

func (r *ApiVersionsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:60
	_go_fuzz_dep_.CoverTab[98468]++
													switch r.Version {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:62
		_go_fuzz_dep_.CoverTab[98469]++
														return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:63
		// _ = "end of CoverTab[98469]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:64
		_go_fuzz_dep_.CoverTab[98470]++
														return V2_4_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:65
		// _ = "end of CoverTab[98470]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:66
		_go_fuzz_dep_.CoverTab[98471]++
														return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:67
		// _ = "end of CoverTab[98471]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:68
	// _ = "end of CoverTab[98468]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:69
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/api_versions_request.go:69
var _ = _go_fuzz_dep_.CoverTab
