//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:1
)

type SyncGroupRequest struct {
	GroupId			string
	GenerationId		int32
	MemberId		string
	GroupAssignments	map[string][]byte
}

func (r *SyncGroupRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:10
	_go_fuzz_dep_.CoverTab[106780]++
												if err := pe.putString(r.GroupId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:11
		_go_fuzz_dep_.CoverTab[106785]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:12
		// _ = "end of CoverTab[106785]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:13
		_go_fuzz_dep_.CoverTab[106786]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:13
		// _ = "end of CoverTab[106786]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:13
	// _ = "end of CoverTab[106780]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:13
	_go_fuzz_dep_.CoverTab[106781]++

												pe.putInt32(r.GenerationId)

												if err := pe.putString(r.MemberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:17
		_go_fuzz_dep_.CoverTab[106787]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:18
		// _ = "end of CoverTab[106787]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:19
		_go_fuzz_dep_.CoverTab[106788]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:19
		// _ = "end of CoverTab[106788]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:19
	// _ = "end of CoverTab[106781]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:19
	_go_fuzz_dep_.CoverTab[106782]++

												if err := pe.putArrayLength(len(r.GroupAssignments)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:21
		_go_fuzz_dep_.CoverTab[106789]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:22
		// _ = "end of CoverTab[106789]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:23
		_go_fuzz_dep_.CoverTab[106790]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:23
		// _ = "end of CoverTab[106790]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:23
	// _ = "end of CoverTab[106782]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:23
	_go_fuzz_dep_.CoverTab[106783]++
												for memberId, memberAssignment := range r.GroupAssignments {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:24
		_go_fuzz_dep_.CoverTab[106791]++
													if err := pe.putString(memberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:25
			_go_fuzz_dep_.CoverTab[106793]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:26
			// _ = "end of CoverTab[106793]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:27
			_go_fuzz_dep_.CoverTab[106794]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:27
			// _ = "end of CoverTab[106794]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:27
		// _ = "end of CoverTab[106791]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:27
		_go_fuzz_dep_.CoverTab[106792]++
													if err := pe.putBytes(memberAssignment); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:28
			_go_fuzz_dep_.CoverTab[106795]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:29
			// _ = "end of CoverTab[106795]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:30
			_go_fuzz_dep_.CoverTab[106796]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:30
			// _ = "end of CoverTab[106796]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:30
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:30
		// _ = "end of CoverTab[106792]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:31
	// _ = "end of CoverTab[106783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:31
	_go_fuzz_dep_.CoverTab[106784]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:33
	// _ = "end of CoverTab[106784]"
}

func (r *SyncGroupRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:36
	_go_fuzz_dep_.CoverTab[106797]++
												if r.GroupId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:37
		_go_fuzz_dep_.CoverTab[106804]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:38
		// _ = "end of CoverTab[106804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:39
		_go_fuzz_dep_.CoverTab[106805]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:39
		// _ = "end of CoverTab[106805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:39
	// _ = "end of CoverTab[106797]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:39
	_go_fuzz_dep_.CoverTab[106798]++
												if r.GenerationId, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:40
		_go_fuzz_dep_.CoverTab[106806]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:41
		// _ = "end of CoverTab[106806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:42
		_go_fuzz_dep_.CoverTab[106807]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:42
		// _ = "end of CoverTab[106807]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:42
	// _ = "end of CoverTab[106798]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:42
	_go_fuzz_dep_.CoverTab[106799]++
												if r.MemberId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:43
		_go_fuzz_dep_.CoverTab[106808]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:44
		// _ = "end of CoverTab[106808]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:45
		_go_fuzz_dep_.CoverTab[106809]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:45
		// _ = "end of CoverTab[106809]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:45
	// _ = "end of CoverTab[106799]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:45
	_go_fuzz_dep_.CoverTab[106800]++

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:48
		_go_fuzz_dep_.CoverTab[106810]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:49
		// _ = "end of CoverTab[106810]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:50
		_go_fuzz_dep_.CoverTab[106811]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:50
		// _ = "end of CoverTab[106811]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:50
	// _ = "end of CoverTab[106800]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:50
	_go_fuzz_dep_.CoverTab[106801]++
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:51
		_go_fuzz_dep_.CoverTab[106812]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:52
		// _ = "end of CoverTab[106812]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:53
		_go_fuzz_dep_.CoverTab[106813]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:53
		// _ = "end of CoverTab[106813]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:53
	// _ = "end of CoverTab[106801]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:53
	_go_fuzz_dep_.CoverTab[106802]++

												r.GroupAssignments = make(map[string][]byte)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:56
		_go_fuzz_dep_.CoverTab[106814]++
													memberId, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:58
			_go_fuzz_dep_.CoverTab[106817]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:59
			// _ = "end of CoverTab[106817]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:60
			_go_fuzz_dep_.CoverTab[106818]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:60
			// _ = "end of CoverTab[106818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:60
		// _ = "end of CoverTab[106814]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:60
		_go_fuzz_dep_.CoverTab[106815]++
													memberAssignment, err := pd.getBytes()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:62
			_go_fuzz_dep_.CoverTab[106819]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:63
			// _ = "end of CoverTab[106819]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:64
			_go_fuzz_dep_.CoverTab[106820]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:64
			// _ = "end of CoverTab[106820]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:64
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:64
		// _ = "end of CoverTab[106815]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:64
		_go_fuzz_dep_.CoverTab[106816]++

													r.GroupAssignments[memberId] = memberAssignment
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:66
		// _ = "end of CoverTab[106816]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:67
	// _ = "end of CoverTab[106802]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:67
	_go_fuzz_dep_.CoverTab[106803]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:69
	// _ = "end of CoverTab[106803]"
}

func (r *SyncGroupRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:72
	_go_fuzz_dep_.CoverTab[106821]++
												return 14
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:73
	// _ = "end of CoverTab[106821]"
}

func (r *SyncGroupRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:76
	_go_fuzz_dep_.CoverTab[106822]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:77
	// _ = "end of CoverTab[106822]"
}

func (r *SyncGroupRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:80
	_go_fuzz_dep_.CoverTab[106823]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:81
	// _ = "end of CoverTab[106823]"
}

func (r *SyncGroupRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:84
	_go_fuzz_dep_.CoverTab[106824]++
												return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:85
	// _ = "end of CoverTab[106824]"
}

func (r *SyncGroupRequest) AddGroupAssignment(memberId string, memberAssignment []byte) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:88
	_go_fuzz_dep_.CoverTab[106825]++
												if r.GroupAssignments == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:89
		_go_fuzz_dep_.CoverTab[106827]++
													r.GroupAssignments = make(map[string][]byte)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:90
		// _ = "end of CoverTab[106827]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:91
		_go_fuzz_dep_.CoverTab[106828]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:91
		// _ = "end of CoverTab[106828]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:91
	// _ = "end of CoverTab[106825]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:91
	_go_fuzz_dep_.CoverTab[106826]++

												r.GroupAssignments[memberId] = memberAssignment
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:93
	// _ = "end of CoverTab[106826]"
}

func (r *SyncGroupRequest) AddGroupAssignmentMember(memberId string, memberAssignment *ConsumerGroupMemberAssignment) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:96
	_go_fuzz_dep_.CoverTab[106829]++
												bin, err := encode(memberAssignment, nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:98
		_go_fuzz_dep_.CoverTab[106831]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:99
		// _ = "end of CoverTab[106831]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:100
		_go_fuzz_dep_.CoverTab[106832]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:100
		// _ = "end of CoverTab[106832]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:100
	// _ = "end of CoverTab[106829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:100
	_go_fuzz_dep_.CoverTab[106830]++

												r.AddGroupAssignment(memberId, bin)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:103
	// _ = "end of CoverTab[106830]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:104
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/sync_group_request.go:104
var _ = _go_fuzz_dep_.CoverTab
