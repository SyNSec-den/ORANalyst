//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:1
)

type LeaveGroupRequest struct {
	GroupId		string
	MemberId	string
}

func (r *LeaveGroupRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:8
	_go_fuzz_dep_.CoverTab[103729]++
												if err := pe.putString(r.GroupId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:9
		_go_fuzz_dep_.CoverTab[103732]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:10
		// _ = "end of CoverTab[103732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:11
		_go_fuzz_dep_.CoverTab[103733]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:11
		// _ = "end of CoverTab[103733]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:11
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:11
	// _ = "end of CoverTab[103729]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:11
	_go_fuzz_dep_.CoverTab[103730]++
												if err := pe.putString(r.MemberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:12
		_go_fuzz_dep_.CoverTab[103734]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:13
		// _ = "end of CoverTab[103734]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:14
		_go_fuzz_dep_.CoverTab[103735]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:14
		// _ = "end of CoverTab[103735]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:14
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:14
	// _ = "end of CoverTab[103730]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:14
	_go_fuzz_dep_.CoverTab[103731]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:16
	// _ = "end of CoverTab[103731]"
}

func (r *LeaveGroupRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:19
	_go_fuzz_dep_.CoverTab[103736]++
												if r.GroupId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:20
		_go_fuzz_dep_.CoverTab[103739]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:21
		// _ = "end of CoverTab[103739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:22
		_go_fuzz_dep_.CoverTab[103740]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:22
		// _ = "end of CoverTab[103740]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:22
	// _ = "end of CoverTab[103736]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:22
	_go_fuzz_dep_.CoverTab[103737]++
												if r.MemberId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:23
		_go_fuzz_dep_.CoverTab[103741]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:24
		// _ = "end of CoverTab[103741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:25
		_go_fuzz_dep_.CoverTab[103742]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:25
		// _ = "end of CoverTab[103742]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:25
	// _ = "end of CoverTab[103737]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:25
	_go_fuzz_dep_.CoverTab[103738]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:27
	// _ = "end of CoverTab[103738]"
}

func (r *LeaveGroupRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:30
	_go_fuzz_dep_.CoverTab[103743]++
												return 13
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:31
	// _ = "end of CoverTab[103743]"
}

func (r *LeaveGroupRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:34
	_go_fuzz_dep_.CoverTab[103744]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:35
	// _ = "end of CoverTab[103744]"
}

func (r *LeaveGroupRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:38
	_go_fuzz_dep_.CoverTab[103745]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:39
	// _ = "end of CoverTab[103745]"
}

func (r *LeaveGroupRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:42
	_go_fuzz_dep_.CoverTab[103746]++
												return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:43
	// _ = "end of CoverTab[103746]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_request.go:44
var _ = _go_fuzz_dep_.CoverTab
