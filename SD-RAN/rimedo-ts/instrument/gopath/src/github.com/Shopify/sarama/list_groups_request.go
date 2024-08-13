//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:1
)

type ListGroupsRequest struct{}

func (r *ListGroupsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:5
	_go_fuzz_dep_.CoverTab[103784]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:6
	// _ = "end of CoverTab[103784]"
}

func (r *ListGroupsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:9
	_go_fuzz_dep_.CoverTab[103785]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:10
	// _ = "end of CoverTab[103785]"
}

func (r *ListGroupsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:13
	_go_fuzz_dep_.CoverTab[103786]++
												return 16
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:14
	// _ = "end of CoverTab[103786]"
}

func (r *ListGroupsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:17
	_go_fuzz_dep_.CoverTab[103787]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:18
	// _ = "end of CoverTab[103787]"
}

func (r *ListGroupsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:21
	_go_fuzz_dep_.CoverTab[103788]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:22
	// _ = "end of CoverTab[103788]"
}

func (r *ListGroupsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:25
	_go_fuzz_dep_.CoverTab[103789]++
												return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:26
	// _ = "end of CoverTab[103789]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:27
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_request.go:27
var _ = _go_fuzz_dep_.CoverTab
