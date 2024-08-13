//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:1
)

type DescribeGroupsRequest struct {
	Groups []string
}

func (r *DescribeGroupsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:7
	_go_fuzz_dep_.CoverTab[102352]++
													return pe.putStringArray(r.Groups)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:8
	// _ = "end of CoverTab[102352]"
}

func (r *DescribeGroupsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:11
	_go_fuzz_dep_.CoverTab[102353]++
													r.Groups, err = pd.getStringArray()
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:13
	// _ = "end of CoverTab[102353]"
}

func (r *DescribeGroupsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:16
	_go_fuzz_dep_.CoverTab[102354]++
													return 15
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:17
	// _ = "end of CoverTab[102354]"
}

func (r *DescribeGroupsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:20
	_go_fuzz_dep_.CoverTab[102355]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:21
	// _ = "end of CoverTab[102355]"
}

func (r *DescribeGroupsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:24
	_go_fuzz_dep_.CoverTab[102356]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:25
	// _ = "end of CoverTab[102356]"
}

func (r *DescribeGroupsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:28
	_go_fuzz_dep_.CoverTab[102357]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:29
	// _ = "end of CoverTab[102357]"
}

func (r *DescribeGroupsRequest) AddGroup(group string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:32
	_go_fuzz_dep_.CoverTab[102358]++
													r.Groups = append(r.Groups, group)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:33
	// _ = "end of CoverTab[102358]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_groups_request.go:34
var _ = _go_fuzz_dep_.CoverTab
