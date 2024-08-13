//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:1
)

type AlterUserScramCredentialsRequest struct {
	Version	int16

	// Deletions represent list of SCRAM credentials to remove
	Deletions	[]AlterUserScramCredentialsDelete

	// Upsertions represent list of SCRAM credentials to update/insert
	Upsertions	[]AlterUserScramCredentialsUpsert
}

type AlterUserScramCredentialsDelete struct {
	Name		string
	Mechanism	ScramMechanismType
}

type AlterUserScramCredentialsUpsert struct {
	Name		string
	Mechanism	ScramMechanismType
	Iterations	int32
	Salt		[]byte
	saltedPassword	[]byte

	// This field is never transmitted over the wire
	// @see: https://tools.ietf.org/html/rfc5802
	Password	[]byte
}

func (r *AlterUserScramCredentialsRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:30
	_go_fuzz_dep_.CoverTab[98339]++
															pe.putCompactArrayLength(len(r.Deletions))
															for _, d := range r.Deletions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:32
		_go_fuzz_dep_.CoverTab[98342]++
																if err := pe.putCompactString(d.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:33
			_go_fuzz_dep_.CoverTab[98344]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:34
			// _ = "end of CoverTab[98344]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:35
			_go_fuzz_dep_.CoverTab[98345]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:35
			// _ = "end of CoverTab[98345]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:35
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:35
		// _ = "end of CoverTab[98342]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:35
		_go_fuzz_dep_.CoverTab[98343]++
																pe.putInt8(int8(d.Mechanism))
																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:37
		// _ = "end of CoverTab[98343]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:38
	// _ = "end of CoverTab[98339]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:38
	_go_fuzz_dep_.CoverTab[98340]++

															pe.putCompactArrayLength(len(r.Upsertions))
															for _, u := range r.Upsertions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:41
		_go_fuzz_dep_.CoverTab[98346]++
																if err := pe.putCompactString(u.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:42
			_go_fuzz_dep_.CoverTab[98351]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:43
			// _ = "end of CoverTab[98351]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:44
			_go_fuzz_dep_.CoverTab[98352]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:44
			// _ = "end of CoverTab[98352]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:44
		// _ = "end of CoverTab[98346]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:44
		_go_fuzz_dep_.CoverTab[98347]++
																pe.putInt8(int8(u.Mechanism))
																pe.putInt32(u.Iterations)

																if err := pe.putCompactBytes(u.Salt); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:48
			_go_fuzz_dep_.CoverTab[98353]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:49
			// _ = "end of CoverTab[98353]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:50
			_go_fuzz_dep_.CoverTab[98354]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:50
			// _ = "end of CoverTab[98354]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:50
		// _ = "end of CoverTab[98347]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:50
		_go_fuzz_dep_.CoverTab[98348]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:53
		formatter := scramFormatter{mechanism: u.Mechanism}
		salted, err := formatter.saltedPassword(u.Password, u.Salt, int(u.Iterations))
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:55
			_go_fuzz_dep_.CoverTab[98355]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:56
			// _ = "end of CoverTab[98355]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:57
			_go_fuzz_dep_.CoverTab[98356]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:57
			// _ = "end of CoverTab[98356]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:57
		// _ = "end of CoverTab[98348]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:57
		_go_fuzz_dep_.CoverTab[98349]++

																if err := pe.putCompactBytes(salted); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:59
			_go_fuzz_dep_.CoverTab[98357]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:60
			// _ = "end of CoverTab[98357]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:61
			_go_fuzz_dep_.CoverTab[98358]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:61
			// _ = "end of CoverTab[98358]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:61
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:61
		// _ = "end of CoverTab[98349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:61
		_go_fuzz_dep_.CoverTab[98350]++
																pe.putEmptyTaggedFieldArray()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:62
		// _ = "end of CoverTab[98350]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:63
	// _ = "end of CoverTab[98340]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:63
	_go_fuzz_dep_.CoverTab[98341]++

															pe.putEmptyTaggedFieldArray()
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:66
	// _ = "end of CoverTab[98341]"
}

func (r *AlterUserScramCredentialsRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:69
	_go_fuzz_dep_.CoverTab[98359]++
															numDeletions, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:71
		_go_fuzz_dep_.CoverTab[98365]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:72
		// _ = "end of CoverTab[98365]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:73
		_go_fuzz_dep_.CoverTab[98366]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:73
		// _ = "end of CoverTab[98366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:73
	// _ = "end of CoverTab[98359]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:73
	_go_fuzz_dep_.CoverTab[98360]++

															r.Deletions = make([]AlterUserScramCredentialsDelete, numDeletions)
															for i := 0; i < numDeletions; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:76
		_go_fuzz_dep_.CoverTab[98367]++
																r.Deletions[i] = AlterUserScramCredentialsDelete{}
																if r.Deletions[i].Name, err = pd.getCompactString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:78
			_go_fuzz_dep_.CoverTab[98370]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:79
			// _ = "end of CoverTab[98370]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:80
			_go_fuzz_dep_.CoverTab[98371]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:80
			// _ = "end of CoverTab[98371]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:80
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:80
		// _ = "end of CoverTab[98367]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:80
		_go_fuzz_dep_.CoverTab[98368]++
																mechanism, err := pd.getInt8()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:82
			_go_fuzz_dep_.CoverTab[98372]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:83
			// _ = "end of CoverTab[98372]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:84
			_go_fuzz_dep_.CoverTab[98373]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:84
			// _ = "end of CoverTab[98373]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:84
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:84
		// _ = "end of CoverTab[98368]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:84
		_go_fuzz_dep_.CoverTab[98369]++
																r.Deletions[i].Mechanism = ScramMechanismType(mechanism)
																if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:86
			_go_fuzz_dep_.CoverTab[98374]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:87
			// _ = "end of CoverTab[98374]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:88
			_go_fuzz_dep_.CoverTab[98375]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:88
			// _ = "end of CoverTab[98375]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:88
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:88
		// _ = "end of CoverTab[98369]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:89
	// _ = "end of CoverTab[98360]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:89
	_go_fuzz_dep_.CoverTab[98361]++

															numUpsertions, err := pd.getCompactArrayLength()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:92
		_go_fuzz_dep_.CoverTab[98376]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:93
		// _ = "end of CoverTab[98376]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:94
		_go_fuzz_dep_.CoverTab[98377]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:94
		// _ = "end of CoverTab[98377]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:94
	// _ = "end of CoverTab[98361]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:94
	_go_fuzz_dep_.CoverTab[98362]++

															r.Upsertions = make([]AlterUserScramCredentialsUpsert, numUpsertions)
															for i := 0; i < numUpsertions; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:97
		_go_fuzz_dep_.CoverTab[98378]++
																r.Upsertions[i] = AlterUserScramCredentialsUpsert{}
																if r.Upsertions[i].Name, err = pd.getCompactString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:99
			_go_fuzz_dep_.CoverTab[98384]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:100
			// _ = "end of CoverTab[98384]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:101
			_go_fuzz_dep_.CoverTab[98385]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:101
			// _ = "end of CoverTab[98385]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:101
		// _ = "end of CoverTab[98378]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:101
		_go_fuzz_dep_.CoverTab[98379]++
																mechanism, err := pd.getInt8()
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:103
			_go_fuzz_dep_.CoverTab[98386]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:104
			// _ = "end of CoverTab[98386]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:105
			_go_fuzz_dep_.CoverTab[98387]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:105
			// _ = "end of CoverTab[98387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:105
		// _ = "end of CoverTab[98379]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:105
		_go_fuzz_dep_.CoverTab[98380]++

																r.Upsertions[i].Mechanism = ScramMechanismType(mechanism)
																if r.Upsertions[i].Iterations, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:108
			_go_fuzz_dep_.CoverTab[98388]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:109
			// _ = "end of CoverTab[98388]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:110
			_go_fuzz_dep_.CoverTab[98389]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:110
			// _ = "end of CoverTab[98389]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:110
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:110
		// _ = "end of CoverTab[98380]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:110
		_go_fuzz_dep_.CoverTab[98381]++
																if r.Upsertions[i].Salt, err = pd.getCompactBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:111
			_go_fuzz_dep_.CoverTab[98390]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:112
			// _ = "end of CoverTab[98390]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:113
			_go_fuzz_dep_.CoverTab[98391]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:113
			// _ = "end of CoverTab[98391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:113
		// _ = "end of CoverTab[98381]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:113
		_go_fuzz_dep_.CoverTab[98382]++
																if r.Upsertions[i].saltedPassword, err = pd.getCompactBytes(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:114
			_go_fuzz_dep_.CoverTab[98392]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:115
			// _ = "end of CoverTab[98392]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:116
			_go_fuzz_dep_.CoverTab[98393]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:116
			// _ = "end of CoverTab[98393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:116
		// _ = "end of CoverTab[98382]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:116
		_go_fuzz_dep_.CoverTab[98383]++
																if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:117
			_go_fuzz_dep_.CoverTab[98394]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:118
			// _ = "end of CoverTab[98394]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:119
			_go_fuzz_dep_.CoverTab[98395]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:119
			// _ = "end of CoverTab[98395]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:119
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:119
		// _ = "end of CoverTab[98383]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:120
	// _ = "end of CoverTab[98362]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:120
	_go_fuzz_dep_.CoverTab[98363]++

															if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:122
		_go_fuzz_dep_.CoverTab[98396]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:123
		// _ = "end of CoverTab[98396]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:124
		_go_fuzz_dep_.CoverTab[98397]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:124
		// _ = "end of CoverTab[98397]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:124
	// _ = "end of CoverTab[98363]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:124
	_go_fuzz_dep_.CoverTab[98364]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:125
	// _ = "end of CoverTab[98364]"
}

func (r *AlterUserScramCredentialsRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:128
	_go_fuzz_dep_.CoverTab[98398]++
															return 51
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:129
	// _ = "end of CoverTab[98398]"
}

func (r *AlterUserScramCredentialsRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:132
	_go_fuzz_dep_.CoverTab[98399]++
															return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:133
	// _ = "end of CoverTab[98399]"
}

func (r *AlterUserScramCredentialsRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:136
	_go_fuzz_dep_.CoverTab[98400]++
															return 2
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:137
	// _ = "end of CoverTab[98400]"
}

func (r *AlterUserScramCredentialsRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:140
	_go_fuzz_dep_.CoverTab[98401]++
															return V2_7_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:141
	// _ = "end of CoverTab[98401]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:142
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/alter_user_scram_credentials_request.go:142
var _ = _go_fuzz_dep_.CoverTab
