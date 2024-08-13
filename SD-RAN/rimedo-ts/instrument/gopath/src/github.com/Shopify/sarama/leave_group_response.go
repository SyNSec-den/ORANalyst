//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:1
)

type LeaveGroupResponse struct {
	Err KError
}

func (r *LeaveGroupResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:7
	_go_fuzz_dep_.CoverTab[103747]++
												pe.putInt16(int16(r.Err))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:9
	// _ = "end of CoverTab[103747]"
}

func (r *LeaveGroupResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:12
	_go_fuzz_dep_.CoverTab[103748]++
													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:14
		_go_fuzz_dep_.CoverTab[103750]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:15
		// _ = "end of CoverTab[103750]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:16
		_go_fuzz_dep_.CoverTab[103751]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:16
		// _ = "end of CoverTab[103751]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:16
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:16
	// _ = "end of CoverTab[103748]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:16
	_go_fuzz_dep_.CoverTab[103749]++
													r.Err = KError(kerr)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:19
	// _ = "end of CoverTab[103749]"
}

func (r *LeaveGroupResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:22
	_go_fuzz_dep_.CoverTab[103752]++
													return 13
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:23
	// _ = "end of CoverTab[103752]"
}

func (r *LeaveGroupResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:26
	_go_fuzz_dep_.CoverTab[103753]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:27
	// _ = "end of CoverTab[103753]"
}

func (r *LeaveGroupResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:30
	_go_fuzz_dep_.CoverTab[103754]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:31
	// _ = "end of CoverTab[103754]"
}

func (r *LeaveGroupResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:34
	_go_fuzz_dep_.CoverTab[103755]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:35
	// _ = "end of CoverTab[103755]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:36
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/leave_group_response.go:36
var _ = _go_fuzz_dep_.CoverTab
