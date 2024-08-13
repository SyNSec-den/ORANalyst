//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:1
)

type DeleteGroupsRequest struct {
	Groups []string
}

func (r *DeleteGroupsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:7
	_go_fuzz_dep_.CoverTab[101697]++
													return pe.putStringArray(r.Groups)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:8
	// _ = "end of CoverTab[101697]"
}

func (r *DeleteGroupsRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:11
	_go_fuzz_dep_.CoverTab[101698]++
													r.Groups, err = pd.getStringArray()
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:13
	// _ = "end of CoverTab[101698]"
}

func (r *DeleteGroupsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:16
	_go_fuzz_dep_.CoverTab[101699]++
													return 42
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:17
	// _ = "end of CoverTab[101699]"
}

func (r *DeleteGroupsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:20
	_go_fuzz_dep_.CoverTab[101700]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:21
	// _ = "end of CoverTab[101700]"
}

func (r *DeleteGroupsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:24
	_go_fuzz_dep_.CoverTab[101701]++
													return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:25
	// _ = "end of CoverTab[101701]"
}

func (r *DeleteGroupsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:28
	_go_fuzz_dep_.CoverTab[101702]++
													return V1_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:29
	// _ = "end of CoverTab[101702]"
}

func (r *DeleteGroupsRequest) AddGroup(group string) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:32
	_go_fuzz_dep_.CoverTab[101703]++
													r.Groups = append(r.Groups, group)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:33
	// _ = "end of CoverTab[101703]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:34
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_request.go:34
var _ = _go_fuzz_dep_.CoverTab
