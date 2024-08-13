//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:1
)

// DescribeUserScramCredentialsRequest is a request to get list of SCRAM user names
type DescribeUserScramCredentialsRequest struct {
	// Version 0 is currently only supported
	Version	int16

	// If this is an empty array, all users will be queried
	DescribeUsers	[]DescribeUserScramCredentialsRequestUser
}

// DescribeUserScramCredentialsRequestUser is a describe request about specific user name
type DescribeUserScramCredentialsRequestUser struct {
	Name string
}

func (r *DescribeUserScramCredentialsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:17
	_go_fuzz_dep_.CoverTab[102585]++
															pe.putCompactArrayLength(len(r.DescribeUsers))
															for _, d := range r.DescribeUsers {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:19
		_go_fuzz_dep_.CoverTab[102587]++
																if err := pe.putCompactString(d.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:20
			_go_fuzz_dep_.CoverTab[102589]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:21
			// _ = "end of CoverTab[102589]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:22
			_go_fuzz_dep_.CoverTab[102590]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:22
			// _ = "end of CoverTab[102590]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:22
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:22
		// _ = "end of CoverTab[102587]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:22
		_go_fuzz_dep_.CoverTab[102588]++
																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:23
		// _ = "end of CoverTab[102588]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:24
	// _ = "end of CoverTab[102585]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:24
	_go_fuzz_dep_.CoverTab[102586]++

															pe.putEmptyTaggedFieldArray()
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:27
	// _ = "end of CoverTab[102586]"
}

func (r *DescribeUserScramCredentialsRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:30
	_go_fuzz_dep_.CoverTab[102591]++
															n, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:32
		_go_fuzz_dep_.CoverTab[102596]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:33
		// _ = "end of CoverTab[102596]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:34
		_go_fuzz_dep_.CoverTab[102597]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:34
		// _ = "end of CoverTab[102597]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:34
	// _ = "end of CoverTab[102591]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:34
	_go_fuzz_dep_.CoverTab[102592]++
															if n == -1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:35
		_go_fuzz_dep_.CoverTab[102598]++
																n = 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:36
		// _ = "end of CoverTab[102598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:37
		_go_fuzz_dep_.CoverTab[102599]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:37
		// _ = "end of CoverTab[102599]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:37
	// _ = "end of CoverTab[102592]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:37
	_go_fuzz_dep_.CoverTab[102593]++

															r.DescribeUsers = make([]DescribeUserScramCredentialsRequestUser, n)
															for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:40
		_go_fuzz_dep_.CoverTab[102600]++
																r.DescribeUsers[i] = DescribeUserScramCredentialsRequestUser{}
																if r.DescribeUsers[i].Name, err = pd.getCompactString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:42
			_go_fuzz_dep_.CoverTab[102602]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:43
			// _ = "end of CoverTab[102602]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:44
			_go_fuzz_dep_.CoverTab[102603]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:44
			// _ = "end of CoverTab[102603]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:44
		// _ = "end of CoverTab[102600]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:44
		_go_fuzz_dep_.CoverTab[102601]++
																if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:45
			_go_fuzz_dep_.CoverTab[102604]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:46
			// _ = "end of CoverTab[102604]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:47
			_go_fuzz_dep_.CoverTab[102605]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:47
			// _ = "end of CoverTab[102605]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:47
		// _ = "end of CoverTab[102601]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:48
	// _ = "end of CoverTab[102593]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:48
	_go_fuzz_dep_.CoverTab[102594]++

															if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:50
		_go_fuzz_dep_.CoverTab[102606]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:51
		// _ = "end of CoverTab[102606]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:52
		_go_fuzz_dep_.CoverTab[102607]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:52
		// _ = "end of CoverTab[102607]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:52
	// _ = "end of CoverTab[102594]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:52
	_go_fuzz_dep_.CoverTab[102595]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:53
	// _ = "end of CoverTab[102595]"
}

func (r *DescribeUserScramCredentialsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:56
	_go_fuzz_dep_.CoverTab[102608]++
															return 50
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:57
	// _ = "end of CoverTab[102608]"
}

func (r *DescribeUserScramCredentialsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:60
	_go_fuzz_dep_.CoverTab[102609]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:61
	// _ = "end of CoverTab[102609]"
}

func (r *DescribeUserScramCredentialsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:64
	_go_fuzz_dep_.CoverTab[102610]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:65
	// _ = "end of CoverTab[102610]"
}

func (r *DescribeUserScramCredentialsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:68
	_go_fuzz_dep_.CoverTab[102611]++
															return V2_7_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:69
	// _ = "end of CoverTab[102611]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:70
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_request.go:70
var _ = _go_fuzz_dep_.CoverTab
