//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:1
)

import "time"

type ScramMechanismType int8

const (
	SCRAM_MECHANISM_UNKNOWN	ScramMechanismType	= iota	// 0
	SCRAM_MECHANISM_SHA_256					// 1
	SCRAM_MECHANISM_SHA_512					// 2
)

func (s ScramMechanismType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:13
	_go_fuzz_dep_.CoverTab[102612]++
															switch s {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:15
		_go_fuzz_dep_.CoverTab[102613]++
																return SASLTypeSCRAMSHA256
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:16
		// _ = "end of CoverTab[102613]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:17
		_go_fuzz_dep_.CoverTab[102614]++
																return SASLTypeSCRAMSHA512
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:18
		// _ = "end of CoverTab[102614]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:19
		_go_fuzz_dep_.CoverTab[102615]++
																return "Unknown"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:20
		// _ = "end of CoverTab[102615]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:21
	// _ = "end of CoverTab[102612]"
}

type DescribeUserScramCredentialsResponse struct {
	// Version 0 is currently only supported
	Version	int16

	ThrottleTime	time.Duration

	ErrorCode	KError
	ErrorMessage	*string

	Results	[]*DescribeUserScramCredentialsResult
}

type DescribeUserScramCredentialsResult struct {
	User	string

	ErrorCode	KError
	ErrorMessage	*string

	CredentialInfos	[]*UserScramCredentialsResponseInfo
}

type UserScramCredentialsResponseInfo struct {
	Mechanism	ScramMechanismType
	Iterations	int32
}

func (r *DescribeUserScramCredentialsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:50
	_go_fuzz_dep_.CoverTab[102616]++
															pe.putInt32(int32(r.ThrottleTime / time.Millisecond))

															pe.putInt16(int16(r.ErrorCode))
															if err := pe.putNullableCompactString(r.ErrorMessage); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:54
		_go_fuzz_dep_.CoverTab[102619]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:55
		// _ = "end of CoverTab[102619]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:56
		_go_fuzz_dep_.CoverTab[102620]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:56
		// _ = "end of CoverTab[102620]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:56
	// _ = "end of CoverTab[102616]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:56
	_go_fuzz_dep_.CoverTab[102617]++

															pe.putCompactArrayLength(len(r.Results))
															for _, u := range r.Results {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:59
		_go_fuzz_dep_.CoverTab[102621]++
																if err := pe.putCompactString(u.User); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:60
			_go_fuzz_dep_.CoverTab[102625]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:61
			// _ = "end of CoverTab[102625]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:62
			_go_fuzz_dep_.CoverTab[102626]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:62
			// _ = "end of CoverTab[102626]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:62
		// _ = "end of CoverTab[102621]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:62
		_go_fuzz_dep_.CoverTab[102622]++
																pe.putInt16(int16(u.ErrorCode))
																if err := pe.putNullableCompactString(u.ErrorMessage); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:64
			_go_fuzz_dep_.CoverTab[102627]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:65
			// _ = "end of CoverTab[102627]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:66
			_go_fuzz_dep_.CoverTab[102628]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:66
			// _ = "end of CoverTab[102628]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:66
		// _ = "end of CoverTab[102622]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:66
		_go_fuzz_dep_.CoverTab[102623]++

																pe.putCompactArrayLength(len(u.CredentialInfos))
																for _, c := range u.CredentialInfos {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:69
			_go_fuzz_dep_.CoverTab[102629]++
																	pe.putInt8(int8(c.Mechanism))
																	pe.putInt32(c.Iterations)
																	pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:72
			// _ = "end of CoverTab[102629]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:73
		// _ = "end of CoverTab[102623]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:73
		_go_fuzz_dep_.CoverTab[102624]++

																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:75
		// _ = "end of CoverTab[102624]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:76
	// _ = "end of CoverTab[102617]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:76
	_go_fuzz_dep_.CoverTab[102618]++

															pe.putEmptyTaggedFieldArray()
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:79
	// _ = "end of CoverTab[102618]"
}

func (r *DescribeUserScramCredentialsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:82
	_go_fuzz_dep_.CoverTab[102630]++
															throttleTime, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:84
		_go_fuzz_dep_.CoverTab[102637]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:85
		// _ = "end of CoverTab[102637]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:86
		_go_fuzz_dep_.CoverTab[102638]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:86
		// _ = "end of CoverTab[102638]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:86
	// _ = "end of CoverTab[102630]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:86
	_go_fuzz_dep_.CoverTab[102631]++
															r.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

															kerr, err := pd.getInt16()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:90
		_go_fuzz_dep_.CoverTab[102639]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:91
		// _ = "end of CoverTab[102639]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:92
		_go_fuzz_dep_.CoverTab[102640]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:92
		// _ = "end of CoverTab[102640]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:92
	// _ = "end of CoverTab[102631]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:92
	_go_fuzz_dep_.CoverTab[102632]++

															r.ErrorCode = KError(kerr)
															if r.ErrorMessage, err = pd.getCompactNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:95
		_go_fuzz_dep_.CoverTab[102641]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:96
		// _ = "end of CoverTab[102641]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:97
		_go_fuzz_dep_.CoverTab[102642]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:97
		// _ = "end of CoverTab[102642]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:97
	// _ = "end of CoverTab[102632]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:97
	_go_fuzz_dep_.CoverTab[102633]++

															numUsers, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:100
		_go_fuzz_dep_.CoverTab[102643]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:101
		// _ = "end of CoverTab[102643]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:102
		_go_fuzz_dep_.CoverTab[102644]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:102
		// _ = "end of CoverTab[102644]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:102
	// _ = "end of CoverTab[102633]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:102
	_go_fuzz_dep_.CoverTab[102634]++

															if numUsers > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:104
		_go_fuzz_dep_.CoverTab[102645]++
																r.Results = make([]*DescribeUserScramCredentialsResult, numUsers)
																for i := 0; i < numUsers; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:106
			_go_fuzz_dep_.CoverTab[102646]++
																	r.Results[i] = &DescribeUserScramCredentialsResult{}
																	if r.Results[i].User, err = pd.getCompactString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:108
				_go_fuzz_dep_.CoverTab[102652]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:109
				// _ = "end of CoverTab[102652]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:110
				_go_fuzz_dep_.CoverTab[102653]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:110
				// _ = "end of CoverTab[102653]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:110
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:110
			// _ = "end of CoverTab[102646]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:110
			_go_fuzz_dep_.CoverTab[102647]++

																	errorCode, err := pd.getInt16()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:113
				_go_fuzz_dep_.CoverTab[102654]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:114
				// _ = "end of CoverTab[102654]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:115
				_go_fuzz_dep_.CoverTab[102655]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:115
				// _ = "end of CoverTab[102655]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:115
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:115
			// _ = "end of CoverTab[102647]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:115
			_go_fuzz_dep_.CoverTab[102648]++
																	r.Results[i].ErrorCode = KError(errorCode)
																	if r.Results[i].ErrorMessage, err = pd.getCompactNullableString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:117
				_go_fuzz_dep_.CoverTab[102656]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:118
				// _ = "end of CoverTab[102656]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:119
				_go_fuzz_dep_.CoverTab[102657]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:119
				// _ = "end of CoverTab[102657]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:119
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:119
			// _ = "end of CoverTab[102648]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:119
			_go_fuzz_dep_.CoverTab[102649]++

																	numCredentialInfos, err := pd.getCompactArrayLength()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:122
				_go_fuzz_dep_.CoverTab[102658]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:123
				// _ = "end of CoverTab[102658]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:124
				_go_fuzz_dep_.CoverTab[102659]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:124
				// _ = "end of CoverTab[102659]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:124
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:124
			// _ = "end of CoverTab[102649]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:124
			_go_fuzz_dep_.CoverTab[102650]++

																	r.Results[i].CredentialInfos = make([]*UserScramCredentialsResponseInfo, numCredentialInfos)
																	for j := 0; j < numCredentialInfos; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:127
				_go_fuzz_dep_.CoverTab[102660]++
																		r.Results[i].CredentialInfos[j] = &UserScramCredentialsResponseInfo{}
																		scramMechanism, err := pd.getInt8()
																		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:130
					_go_fuzz_dep_.CoverTab[102663]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:131
					// _ = "end of CoverTab[102663]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:132
					_go_fuzz_dep_.CoverTab[102664]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:132
					// _ = "end of CoverTab[102664]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:132
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:132
				// _ = "end of CoverTab[102660]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:132
				_go_fuzz_dep_.CoverTab[102661]++
																		r.Results[i].CredentialInfos[j].Mechanism = ScramMechanismType(scramMechanism)
																		if r.Results[i].CredentialInfos[j].Iterations, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:134
					_go_fuzz_dep_.CoverTab[102665]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:135
					// _ = "end of CoverTab[102665]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:136
					_go_fuzz_dep_.CoverTab[102666]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:136
					// _ = "end of CoverTab[102666]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:136
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:136
				// _ = "end of CoverTab[102661]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:136
				_go_fuzz_dep_.CoverTab[102662]++
																		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:137
					_go_fuzz_dep_.CoverTab[102667]++
																			return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:138
					// _ = "end of CoverTab[102667]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:139
					_go_fuzz_dep_.CoverTab[102668]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:139
					// _ = "end of CoverTab[102668]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:139
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:139
				// _ = "end of CoverTab[102662]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:140
			// _ = "end of CoverTab[102650]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:140
			_go_fuzz_dep_.CoverTab[102651]++

																	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:142
				_go_fuzz_dep_.CoverTab[102669]++
																		return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:143
				// _ = "end of CoverTab[102669]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:144
				_go_fuzz_dep_.CoverTab[102670]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:144
				// _ = "end of CoverTab[102670]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:144
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:144
			// _ = "end of CoverTab[102651]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:145
		// _ = "end of CoverTab[102645]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:146
		_go_fuzz_dep_.CoverTab[102671]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:146
		// _ = "end of CoverTab[102671]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:146
	// _ = "end of CoverTab[102634]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:146
	_go_fuzz_dep_.CoverTab[102635]++

															if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:148
		_go_fuzz_dep_.CoverTab[102672]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:149
		// _ = "end of CoverTab[102672]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:150
		_go_fuzz_dep_.CoverTab[102673]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:150
		// _ = "end of CoverTab[102673]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:150
	// _ = "end of CoverTab[102635]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:150
	_go_fuzz_dep_.CoverTab[102636]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:151
	// _ = "end of CoverTab[102636]"
}

func (r *DescribeUserScramCredentialsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:154
	_go_fuzz_dep_.CoverTab[102674]++
															return 50
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:155
	// _ = "end of CoverTab[102674]"
}

func (r *DescribeUserScramCredentialsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:158
	_go_fuzz_dep_.CoverTab[102675]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:159
	// _ = "end of CoverTab[102675]"
}

func (r *DescribeUserScramCredentialsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:162
	_go_fuzz_dep_.CoverTab[102676]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:163
	// _ = "end of CoverTab[102676]"
}

func (r *DescribeUserScramCredentialsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:166
	_go_fuzz_dep_.CoverTab[102677]++
															return V2_7_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:167
	// _ = "end of CoverTab[102677]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/describe_user_scram_credentials_response.go:168
var _ = _go_fuzz_dep_.CoverTab
