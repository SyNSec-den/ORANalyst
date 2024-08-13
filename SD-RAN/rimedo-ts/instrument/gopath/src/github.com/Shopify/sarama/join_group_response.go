//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:1
)

type JoinGroupResponse struct {
	Version		int16
	ThrottleTime	int32
	Err		KError
	GenerationId	int32
	GroupProtocol	string
	LeaderId	string
	MemberId	string
	Members		map[string][]byte
}

func (r *JoinGroupResponse) GetMembers() (map[string]ConsumerGroupMemberMetadata, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:14
	_go_fuzz_dep_.CoverTab[103645]++
												members := make(map[string]ConsumerGroupMemberMetadata, len(r.Members))
												for id, bin := range r.Members {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:16
		_go_fuzz_dep_.CoverTab[103647]++
													meta := new(ConsumerGroupMemberMetadata)
													if err := decode(bin, meta); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:18
			_go_fuzz_dep_.CoverTab[103649]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:19
			// _ = "end of CoverTab[103649]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:20
			_go_fuzz_dep_.CoverTab[103650]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:20
			// _ = "end of CoverTab[103650]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:20
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:20
		// _ = "end of CoverTab[103647]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:20
		_go_fuzz_dep_.CoverTab[103648]++
													members[id] = *meta
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:21
		// _ = "end of CoverTab[103648]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:22
	// _ = "end of CoverTab[103645]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:22
	_go_fuzz_dep_.CoverTab[103646]++
												return members, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:23
	// _ = "end of CoverTab[103646]"
}

func (r *JoinGroupResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:26
	_go_fuzz_dep_.CoverTab[103651]++
												if r.Version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:27
		_go_fuzz_dep_.CoverTab[103658]++
													pe.putInt32(r.ThrottleTime)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:28
		// _ = "end of CoverTab[103658]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:29
		_go_fuzz_dep_.CoverTab[103659]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:29
		// _ = "end of CoverTab[103659]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:29
	// _ = "end of CoverTab[103651]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:29
	_go_fuzz_dep_.CoverTab[103652]++
												pe.putInt16(int16(r.Err))
												pe.putInt32(r.GenerationId)

												if err := pe.putString(r.GroupProtocol); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:33
		_go_fuzz_dep_.CoverTab[103660]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:34
		// _ = "end of CoverTab[103660]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:35
		_go_fuzz_dep_.CoverTab[103661]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:35
		// _ = "end of CoverTab[103661]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:35
	// _ = "end of CoverTab[103652]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:35
	_go_fuzz_dep_.CoverTab[103653]++
												if err := pe.putString(r.LeaderId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:36
		_go_fuzz_dep_.CoverTab[103662]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:37
		// _ = "end of CoverTab[103662]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:38
		_go_fuzz_dep_.CoverTab[103663]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:38
		// _ = "end of CoverTab[103663]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:38
	// _ = "end of CoverTab[103653]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:38
	_go_fuzz_dep_.CoverTab[103654]++
												if err := pe.putString(r.MemberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:39
		_go_fuzz_dep_.CoverTab[103664]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:40
		// _ = "end of CoverTab[103664]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:41
		_go_fuzz_dep_.CoverTab[103665]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:41
		// _ = "end of CoverTab[103665]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:41
	// _ = "end of CoverTab[103654]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:41
	_go_fuzz_dep_.CoverTab[103655]++

												if err := pe.putArrayLength(len(r.Members)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:43
		_go_fuzz_dep_.CoverTab[103666]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:44
		// _ = "end of CoverTab[103666]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:45
		_go_fuzz_dep_.CoverTab[103667]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:45
		// _ = "end of CoverTab[103667]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:45
	// _ = "end of CoverTab[103655]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:45
	_go_fuzz_dep_.CoverTab[103656]++

												for memberId, memberMetadata := range r.Members {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:47
		_go_fuzz_dep_.CoverTab[103668]++
													if err := pe.putString(memberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:48
			_go_fuzz_dep_.CoverTab[103670]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:49
			// _ = "end of CoverTab[103670]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:50
			_go_fuzz_dep_.CoverTab[103671]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:50
			// _ = "end of CoverTab[103671]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:50
		// _ = "end of CoverTab[103668]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:50
		_go_fuzz_dep_.CoverTab[103669]++

													if err := pe.putBytes(memberMetadata); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:52
			_go_fuzz_dep_.CoverTab[103672]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:53
			// _ = "end of CoverTab[103672]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:54
			_go_fuzz_dep_.CoverTab[103673]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:54
			// _ = "end of CoverTab[103673]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:54
		// _ = "end of CoverTab[103669]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:55
	// _ = "end of CoverTab[103656]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:55
	_go_fuzz_dep_.CoverTab[103657]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:57
	// _ = "end of CoverTab[103657]"
}

func (r *JoinGroupResponse) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:60
	_go_fuzz_dep_.CoverTab[103674]++
												r.Version = version

												if version >= 2 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:63
		_go_fuzz_dep_.CoverTab[103684]++
													if r.ThrottleTime, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:64
			_go_fuzz_dep_.CoverTab[103685]++
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:65
			// _ = "end of CoverTab[103685]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:66
			_go_fuzz_dep_.CoverTab[103686]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:66
			// _ = "end of CoverTab[103686]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:66
		// _ = "end of CoverTab[103684]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:67
		_go_fuzz_dep_.CoverTab[103687]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:67
		// _ = "end of CoverTab[103687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:67
	// _ = "end of CoverTab[103674]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:67
	_go_fuzz_dep_.CoverTab[103675]++

												kerr, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:70
		_go_fuzz_dep_.CoverTab[103688]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:71
		// _ = "end of CoverTab[103688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:72
		_go_fuzz_dep_.CoverTab[103689]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:72
		// _ = "end of CoverTab[103689]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:72
	// _ = "end of CoverTab[103675]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:72
	_go_fuzz_dep_.CoverTab[103676]++

												r.Err = KError(kerr)

												if r.GenerationId, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:76
		_go_fuzz_dep_.CoverTab[103690]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:77
		// _ = "end of CoverTab[103690]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:78
		_go_fuzz_dep_.CoverTab[103691]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:78
		// _ = "end of CoverTab[103691]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:78
	// _ = "end of CoverTab[103676]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:78
	_go_fuzz_dep_.CoverTab[103677]++

												if r.GroupProtocol, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:80
		_go_fuzz_dep_.CoverTab[103692]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:81
		// _ = "end of CoverTab[103692]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:82
		_go_fuzz_dep_.CoverTab[103693]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:82
		// _ = "end of CoverTab[103693]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:82
	// _ = "end of CoverTab[103677]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:82
	_go_fuzz_dep_.CoverTab[103678]++

												if r.LeaderId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:84
		_go_fuzz_dep_.CoverTab[103694]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:85
		// _ = "end of CoverTab[103694]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:86
		_go_fuzz_dep_.CoverTab[103695]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:86
		// _ = "end of CoverTab[103695]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:86
	// _ = "end of CoverTab[103678]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:86
	_go_fuzz_dep_.CoverTab[103679]++

												if r.MemberId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:88
		_go_fuzz_dep_.CoverTab[103696]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:89
		// _ = "end of CoverTab[103696]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:90
		_go_fuzz_dep_.CoverTab[103697]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:90
		// _ = "end of CoverTab[103697]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:90
	// _ = "end of CoverTab[103679]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:90
	_go_fuzz_dep_.CoverTab[103680]++

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:93
		_go_fuzz_dep_.CoverTab[103698]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:94
		// _ = "end of CoverTab[103698]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:95
		_go_fuzz_dep_.CoverTab[103699]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:95
		// _ = "end of CoverTab[103699]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:95
	// _ = "end of CoverTab[103680]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:95
	_go_fuzz_dep_.CoverTab[103681]++
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:96
		_go_fuzz_dep_.CoverTab[103700]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:97
		// _ = "end of CoverTab[103700]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:98
		_go_fuzz_dep_.CoverTab[103701]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:98
		// _ = "end of CoverTab[103701]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:98
	// _ = "end of CoverTab[103681]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:98
	_go_fuzz_dep_.CoverTab[103682]++

												r.Members = make(map[string][]byte)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:101
		_go_fuzz_dep_.CoverTab[103702]++
														memberId, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:103
			_go_fuzz_dep_.CoverTab[103705]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:104
			// _ = "end of CoverTab[103705]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:105
			_go_fuzz_dep_.CoverTab[103706]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:105
			// _ = "end of CoverTab[103706]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:105
		// _ = "end of CoverTab[103702]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:105
		_go_fuzz_dep_.CoverTab[103703]++

														memberMetadata, err := pd.getBytes()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:108
			_go_fuzz_dep_.CoverTab[103707]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:109
			// _ = "end of CoverTab[103707]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:110
			_go_fuzz_dep_.CoverTab[103708]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:110
			// _ = "end of CoverTab[103708]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:110
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:110
		// _ = "end of CoverTab[103703]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:110
		_go_fuzz_dep_.CoverTab[103704]++

														r.Members[memberId] = memberMetadata
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:112
		// _ = "end of CoverTab[103704]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:113
	// _ = "end of CoverTab[103682]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:113
	_go_fuzz_dep_.CoverTab[103683]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:115
	// _ = "end of CoverTab[103683]"
}

func (r *JoinGroupResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:118
	_go_fuzz_dep_.CoverTab[103709]++
													return 11
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:119
	// _ = "end of CoverTab[103709]"
}

func (r *JoinGroupResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:122
	_go_fuzz_dep_.CoverTab[103710]++
													return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:123
	// _ = "end of CoverTab[103710]"
}

func (r *JoinGroupResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:126
	_go_fuzz_dep_.CoverTab[103711]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:127
	// _ = "end of CoverTab[103711]"
}

func (r *JoinGroupResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:130
	_go_fuzz_dep_.CoverTab[103712]++
													switch r.Version {
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:132
		_go_fuzz_dep_.CoverTab[103713]++
														return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:133
		// _ = "end of CoverTab[103713]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:134
		_go_fuzz_dep_.CoverTab[103714]++
														return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:135
		// _ = "end of CoverTab[103714]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:136
		_go_fuzz_dep_.CoverTab[103715]++
														return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:137
		// _ = "end of CoverTab[103715]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:138
	// _ = "end of CoverTab[103712]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:139
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_response.go:139
var _ = _go_fuzz_dep_.CoverTab
