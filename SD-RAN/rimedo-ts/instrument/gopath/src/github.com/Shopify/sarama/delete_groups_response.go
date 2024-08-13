//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:1
)

import (
	"time"
)

type DeleteGroupsResponse struct {
	ThrottleTime	time.Duration
	GroupErrorCodes	map[string]KError
}

func (r *DeleteGroupsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:12
	_go_fuzz_dep_.CoverTab[101704]++
													pe.putInt32(int32(r.ThrottleTime / time.Millisecond))

													if err := pe.putArrayLength(len(r.GroupErrorCodes)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:15
		_go_fuzz_dep_.CoverTab[101707]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:16
		// _ = "end of CoverTab[101707]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:17
		_go_fuzz_dep_.CoverTab[101708]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:17
		// _ = "end of CoverTab[101708]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:17
	// _ = "end of CoverTab[101704]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:17
	_go_fuzz_dep_.CoverTab[101705]++
													for groupID, errorCode := range r.GroupErrorCodes {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:18
		_go_fuzz_dep_.CoverTab[101709]++
														if err := pe.putString(groupID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:19
			_go_fuzz_dep_.CoverTab[101711]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:20
			// _ = "end of CoverTab[101711]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:21
			_go_fuzz_dep_.CoverTab[101712]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:21
			// _ = "end of CoverTab[101712]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:21
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:21
		// _ = "end of CoverTab[101709]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:21
		_go_fuzz_dep_.CoverTab[101710]++
														pe.putInt16(int16(errorCode))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:22
		// _ = "end of CoverTab[101710]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:23
	// _ = "end of CoverTab[101705]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:23
	_go_fuzz_dep_.CoverTab[101706]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:25
	// _ = "end of CoverTab[101706]"
}

func (r *DeleteGroupsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:28
	_go_fuzz_dep_.CoverTab[101713]++
													throttleTime, err := pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:30
		_go_fuzz_dep_.CoverTab[101718]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:31
		// _ = "end of CoverTab[101718]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:32
		_go_fuzz_dep_.CoverTab[101719]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:32
		// _ = "end of CoverTab[101719]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:32
	// _ = "end of CoverTab[101713]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:32
	_go_fuzz_dep_.CoverTab[101714]++
													r.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:36
		_go_fuzz_dep_.CoverTab[101720]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:37
		// _ = "end of CoverTab[101720]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:38
		_go_fuzz_dep_.CoverTab[101721]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:38
		// _ = "end of CoverTab[101721]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:38
	// _ = "end of CoverTab[101714]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:38
	_go_fuzz_dep_.CoverTab[101715]++
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:39
		_go_fuzz_dep_.CoverTab[101722]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:40
		// _ = "end of CoverTab[101722]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:41
		_go_fuzz_dep_.CoverTab[101723]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:41
		// _ = "end of CoverTab[101723]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:41
	// _ = "end of CoverTab[101715]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:41
	_go_fuzz_dep_.CoverTab[101716]++

													r.GroupErrorCodes = make(map[string]KError, n)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:44
		_go_fuzz_dep_.CoverTab[101724]++
														groupID, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:46
			_go_fuzz_dep_.CoverTab[101727]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:47
			// _ = "end of CoverTab[101727]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:48
			_go_fuzz_dep_.CoverTab[101728]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:48
			// _ = "end of CoverTab[101728]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:48
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:48
		// _ = "end of CoverTab[101724]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:48
		_go_fuzz_dep_.CoverTab[101725]++
														errorCode, err := pd.getInt16()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:50
			_go_fuzz_dep_.CoverTab[101729]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:51
			// _ = "end of CoverTab[101729]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:52
			_go_fuzz_dep_.CoverTab[101730]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:52
			// _ = "end of CoverTab[101730]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:52
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:52
		// _ = "end of CoverTab[101725]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:52
		_go_fuzz_dep_.CoverTab[101726]++

														r.GroupErrorCodes[groupID] = KError(errorCode)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:54
		// _ = "end of CoverTab[101726]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:55
	// _ = "end of CoverTab[101716]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:55
	_go_fuzz_dep_.CoverTab[101717]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:57
	// _ = "end of CoverTab[101717]"
}

func (r *DeleteGroupsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:60
	_go_fuzz_dep_.CoverTab[101731]++
													return 42
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:61
	// _ = "end of CoverTab[101731]"
}

func (r *DeleteGroupsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:64
	_go_fuzz_dep_.CoverTab[101732]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:65
	// _ = "end of CoverTab[101732]"
}

func (r *DeleteGroupsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:68
	_go_fuzz_dep_.CoverTab[101733]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:69
	// _ = "end of CoverTab[101733]"
}

func (r *DeleteGroupsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:72
	_go_fuzz_dep_.CoverTab[101734]++
													return V1_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:73
	// _ = "end of CoverTab[101734]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/delete_groups_response.go:74
var _ = _go_fuzz_dep_.CoverTab
