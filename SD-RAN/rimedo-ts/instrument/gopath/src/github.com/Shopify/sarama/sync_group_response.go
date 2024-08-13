//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:1
)

type SyncGroupResponse struct {
	Err			KError
	MemberAssignment	[]byte
}

func (r *SyncGroupResponse) GetMemberAssignment() (*ConsumerGroupMemberAssignment, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:8
	_go_fuzz_dep_.CoverTab[106833]++
												assignment := new(ConsumerGroupMemberAssignment)
												err := decode(r.MemberAssignment, assignment)
												return assignment, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:11
	// _ = "end of CoverTab[106833]"
}

func (r *SyncGroupResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:14
	_go_fuzz_dep_.CoverTab[106834]++
												pe.putInt16(int16(r.Err))
												return pe.putBytes(r.MemberAssignment)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:16
	// _ = "end of CoverTab[106834]"
}

func (r *SyncGroupResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:19
	_go_fuzz_dep_.CoverTab[106835]++
												kerr, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:21
		_go_fuzz_dep_.CoverTab[106837]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:22
		// _ = "end of CoverTab[106837]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:23
		_go_fuzz_dep_.CoverTab[106838]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:23
		// _ = "end of CoverTab[106838]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:23
	// _ = "end of CoverTab[106835]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:23
	_go_fuzz_dep_.CoverTab[106836]++

												r.Err = KError(kerr)

												r.MemberAssignment, err = pd.getBytes()
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:28
	// _ = "end of CoverTab[106836]"
}

func (r *SyncGroupResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:31
	_go_fuzz_dep_.CoverTab[106839]++
												return 14
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:32
	// _ = "end of CoverTab[106839]"
}

func (r *SyncGroupResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:35
	_go_fuzz_dep_.CoverTab[106840]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:36
	// _ = "end of CoverTab[106840]"
}

func (r *SyncGroupResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:39
	_go_fuzz_dep_.CoverTab[106841]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:40
	// _ = "end of CoverTab[106841]"
}

func (r *SyncGroupResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:43
	_go_fuzz_dep_.CoverTab[106842]++
												return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:44
	// _ = "end of CoverTab[106842]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:45
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_response.go:45
var _ = _go_fuzz_dep_.CoverTab
