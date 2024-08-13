//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:1
)

import "time"

type AlterUserScramCredentialsResponse struct {
	Version	int16

	ThrottleTime	time.Duration

	Results	[]*AlterUserScramCredentialsResult
}

type AlterUserScramCredentialsResult struct {
	User	string

	ErrorCode	KError
	ErrorMessage	*string
}

func (r *AlterUserScramCredentialsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:20
	_go_fuzz_dep_.CoverTab[98402]++
															pe.putInt32(int32(r.ThrottleTime / time.Millisecond))
															pe.putCompactArrayLength(len(r.Results))

															for _, u := range r.Results {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:24
		_go_fuzz_dep_.CoverTab[98404]++
																if err := pe.putCompactString(u.User); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:25
			_go_fuzz_dep_.CoverTab[98407]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:26
			// _ = "end of CoverTab[98407]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:27
			_go_fuzz_dep_.CoverTab[98408]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:27
			// _ = "end of CoverTab[98408]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:27
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:27
		// _ = "end of CoverTab[98404]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:27
		_go_fuzz_dep_.CoverTab[98405]++
																pe.putInt16(int16(u.ErrorCode))
																if err := pe.putNullableCompactString(u.ErrorMessage); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:29
			_go_fuzz_dep_.CoverTab[98409]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:30
			// _ = "end of CoverTab[98409]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:31
			_go_fuzz_dep_.CoverTab[98410]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:31
			// _ = "end of CoverTab[98410]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:31
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:31
		// _ = "end of CoverTab[98405]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:31
		_go_fuzz_dep_.CoverTab[98406]++
																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:32
		// _ = "end of CoverTab[98406]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:33
	// _ = "end of CoverTab[98402]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:33
	_go_fuzz_dep_.CoverTab[98403]++

															pe.putEmptyTaggedFieldArray()
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:36
	// _ = "end of CoverTab[98403]"
}

func (r *AlterUserScramCredentialsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:39
	_go_fuzz_dep_.CoverTab[98411]++
															throttleTime, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:41
		_go_fuzz_dep_.CoverTab[98416]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:42
		// _ = "end of CoverTab[98416]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:43
		_go_fuzz_dep_.CoverTab[98417]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:43
		// _ = "end of CoverTab[98417]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:43
	// _ = "end of CoverTab[98411]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:43
	_go_fuzz_dep_.CoverTab[98412]++
															r.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

															numResults, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:47
		_go_fuzz_dep_.CoverTab[98418]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:48
		// _ = "end of CoverTab[98418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:49
		_go_fuzz_dep_.CoverTab[98419]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:49
		// _ = "end of CoverTab[98419]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:49
	// _ = "end of CoverTab[98412]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:49
	_go_fuzz_dep_.CoverTab[98413]++

															if numResults > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:51
		_go_fuzz_dep_.CoverTab[98420]++
																r.Results = make([]*AlterUserScramCredentialsResult, numResults)
																for i := 0; i < numResults; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:53
			_go_fuzz_dep_.CoverTab[98421]++
																	r.Results[i] = &AlterUserScramCredentialsResult{}
																	if r.Results[i].User, err = pd.getCompactString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:55
				_go_fuzz_dep_.CoverTab[98425]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:56
				// _ = "end of CoverTab[98425]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:57
				_go_fuzz_dep_.CoverTab[98426]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:57
				// _ = "end of CoverTab[98426]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:57
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:57
			// _ = "end of CoverTab[98421]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:57
			_go_fuzz_dep_.CoverTab[98422]++

																	kerr, err := pd.getInt16()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:60
				_go_fuzz_dep_.CoverTab[98427]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:61
				// _ = "end of CoverTab[98427]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:62
				_go_fuzz_dep_.CoverTab[98428]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:62
				// _ = "end of CoverTab[98428]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:62
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:62
			// _ = "end of CoverTab[98422]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:62
			_go_fuzz_dep_.CoverTab[98423]++

																	r.Results[i].ErrorCode = KError(kerr)
																	if r.Results[i].ErrorMessage, err = pd.getCompactNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:65
				_go_fuzz_dep_.CoverTab[98429]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:66
				// _ = "end of CoverTab[98429]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:67
				_go_fuzz_dep_.CoverTab[98430]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:67
				// _ = "end of CoverTab[98430]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:67
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:67
			// _ = "end of CoverTab[98423]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:67
			_go_fuzz_dep_.CoverTab[98424]++
																	if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:68
				_go_fuzz_dep_.CoverTab[98431]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:69
				// _ = "end of CoverTab[98431]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:70
				_go_fuzz_dep_.CoverTab[98432]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:70
				// _ = "end of CoverTab[98432]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:70
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:70
			// _ = "end of CoverTab[98424]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:71
		// _ = "end of CoverTab[98420]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:72
		_go_fuzz_dep_.CoverTab[98433]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:72
		// _ = "end of CoverTab[98433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:72
	// _ = "end of CoverTab[98413]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:72
	_go_fuzz_dep_.CoverTab[98414]++

															if _, err := pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:74
		_go_fuzz_dep_.CoverTab[98434]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:75
		// _ = "end of CoverTab[98434]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:76
		_go_fuzz_dep_.CoverTab[98435]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:76
		// _ = "end of CoverTab[98435]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:76
	// _ = "end of CoverTab[98414]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:76
	_go_fuzz_dep_.CoverTab[98415]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:77
	// _ = "end of CoverTab[98415]"
}

func (r *AlterUserScramCredentialsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:80
	_go_fuzz_dep_.CoverTab[98436]++
															return 51
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:81
	// _ = "end of CoverTab[98436]"
}

func (r *AlterUserScramCredentialsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:84
	_go_fuzz_dep_.CoverTab[98437]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:85
	// _ = "end of CoverTab[98437]"
}

func (r *AlterUserScramCredentialsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:88
	_go_fuzz_dep_.CoverTab[98438]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:89
	// _ = "end of CoverTab[98438]"
}

func (r *AlterUserScramCredentialsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:92
	_go_fuzz_dep_.CoverTab[98439]++
															return V2_7_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:93
	// _ = "end of CoverTab[98439]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:94
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_response.go:94
var _ = _go_fuzz_dep_.CoverTab
