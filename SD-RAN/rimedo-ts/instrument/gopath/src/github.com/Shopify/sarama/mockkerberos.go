//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:1
)

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

type KafkaGSSAPIHandler struct {
	client		*MockKerberosClient
	badResponse	bool
	badKeyChecksum	bool
}

func (h *KafkaGSSAPIHandler) MockKafkaGSSAPI(buffer []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:20
	_go_fuzz_dep_.CoverTab[104373]++

												err := h.client.Login()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:23
		_go_fuzz_dep_.CoverTab[104378]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:24
		// _ = "end of CoverTab[104378]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:25
		_go_fuzz_dep_.CoverTab[104379]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:25
		// _ = "end of CoverTab[104379]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:25
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:25
	// _ = "end of CoverTab[104373]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:25
	_go_fuzz_dep_.CoverTab[104374]++
												if h.badResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:26
		_go_fuzz_dep_.CoverTab[104380]++
													return []byte{0x00, 0x00, 0x00, 0x01, 0xAD}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:27
		// _ = "end of CoverTab[104380]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:28
		_go_fuzz_dep_.CoverTab[104381]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:28
		// _ = "end of CoverTab[104381]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:28
	// _ = "end of CoverTab[104374]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:28
	_go_fuzz_dep_.CoverTab[104375]++

												pack := gssapi.WrapToken{
		Flags:		KRB5_USER_AUTH,
		EC:		12,
		RRC:		0,
		SndSeqNum:	3398292281,
		Payload:	[]byte{0x11, 0x00},
	}

	if h.badKeyChecksum {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:38
		_go_fuzz_dep_.CoverTab[104382]++
													pack.CheckSum = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:39
		// _ = "end of CoverTab[104382]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:40
		_go_fuzz_dep_.CoverTab[104383]++
													err = pack.SetCheckSum(h.client.ASRep.DecryptedEncPart.Key, keyusage.GSSAPI_ACCEPTOR_SEAL)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:42
			_go_fuzz_dep_.CoverTab[104384]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:43
			// _ = "end of CoverTab[104384]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:44
			_go_fuzz_dep_.CoverTab[104385]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:44
			// _ = "end of CoverTab[104385]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:44
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:44
		// _ = "end of CoverTab[104383]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:45
	// _ = "end of CoverTab[104375]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:45
	_go_fuzz_dep_.CoverTab[104376]++

												packBytes, err := pack.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:48
		_go_fuzz_dep_.CoverTab[104386]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:49
		// _ = "end of CoverTab[104386]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:50
		_go_fuzz_dep_.CoverTab[104387]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:50
		// _ = "end of CoverTab[104387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:50
	// _ = "end of CoverTab[104376]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:50
	_go_fuzz_dep_.CoverTab[104377]++
												lenBytes := len(packBytes)
												response := make([]byte, lenBytes+4)
												copy(response[4:], packBytes)
												binary.BigEndian.PutUint32(response, uint32(lenBytes))
												return response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:55
	// _ = "end of CoverTab[104377]"
}

type MockKerberosClient struct {
	asRepBytes	string
	ASRep		messages.ASRep
	credentials	*credentials.Credentials
	mockError	error
	errorStage	string
}

func (c *MockKerberosClient) Login() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:66
	_go_fuzz_dep_.CoverTab[104388]++
												if c.errorStage == "login" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:67
		_go_fuzz_dep_.CoverTab[104393]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:67
		return c.mockError != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:67
		// _ = "end of CoverTab[104393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:67
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:67
		_go_fuzz_dep_.CoverTab[104394]++
													return c.mockError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:68
		// _ = "end of CoverTab[104394]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:69
		_go_fuzz_dep_.CoverTab[104395]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:69
		// _ = "end of CoverTab[104395]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:69
	// _ = "end of CoverTab[104388]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:69
	_go_fuzz_dep_.CoverTab[104389]++
												c.asRepBytes = "6b8202e9308202e5a003020105a10302010ba22b30293027a103020113a220041e301c301aa003020112a1131b114" +
		"558414d504c452e434f4d636c69656e74a30d1b0b4558414d504c452e434f4da4133011a003020101a10a30081b06636c69656e7" +
		"4a5820156618201523082014ea003020105a10d1b0b4558414d504c452e434f4da220301ea003020102a11730151b066b7262746" +
		"7741b0b4558414d504c452e434f4da382011430820110a003020112a103020101a28201020481ffdb9891175d106818e61008c51" +
		"d0b3462bca92f3bf9d4cfa82de4c4d7aff9994ec87c573e3a3d54dcb2bb79618c76f2bf4a3d006f90d5bdbd049bc18f48be39203" +
		"549ca02acaf63f292b12404f9b74c34b83687119d8f56552ccc0c50ebee2a53bb114c1b4619bb1d5d31f0f49b4d40a08a9b4c046" +
		"2e1398d0b648be1c0e50c552ad16e1d8d8e74263dd0bf0ec591e4797dfd40a9a1be4ae830d03a306e053fd7586fef84ffc5e4a83" +
		"7c3122bf3e6a40fe87e84019f6283634461b955712b44a5f7386c278bff94ec2c2dc0403247e29c2450e853471ceababf9b8911f" +
		"997f2e3010b046d2c49eb438afb0f4c210821e80d4ffa4c9521eb895dcd68610b3feaa682012c30820128a003020112a282011f0" +
		"482011bce73cbce3f1dd17661c412005f0f2257c756fe8e98ff97e6ec24b7bab66e5fd3a3827aeeae4757af0c6e892948122d8b2" +
		"03c8df48df0ef5d142d0e416d688f11daa0fcd63d96bdd431d02b8e951c664eeff286a2be62383d274a04016d5f0e141da58cb86" +
		"331de64063062f4f885e8e9ce5b181ca2fdc67897c5995e0ae1ae0c171a64493ff7bd91bc6d89cd4fce1e2b3ea0a10e34b0d5eda" +
		"aa38ee727b50c5632ed1d2f2b457908e616178d0d80b72af209fb8ac9dbaa1768fa45931392b36b6d8c12400f8ded2efaa0654d0" +
		"da1db966e8b5aab4706c800f95d559664646041fdb38b411c62fc0fbe0d25083a28562b0e1c8df16e62e9d5626b0addee489835f" +
		"eedb0f26c05baa596b69b17f47920aa64b29dc77cfcc97ba47885"
	apRepBytes, err := hex.DecodeString(c.asRepBytes)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:86
		_go_fuzz_dep_.CoverTab[104396]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:87
		// _ = "end of CoverTab[104396]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:88
		_go_fuzz_dep_.CoverTab[104397]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:88
		// _ = "end of CoverTab[104397]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:88
	// _ = "end of CoverTab[104389]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:88
	_go_fuzz_dep_.CoverTab[104390]++
												err = c.ASRep.Unmarshal(apRepBytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:90
		_go_fuzz_dep_.CoverTab[104398]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:91
		// _ = "end of CoverTab[104398]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:92
		_go_fuzz_dep_.CoverTab[104399]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:92
		// _ = "end of CoverTab[104399]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:92
	// _ = "end of CoverTab[104390]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:92
	_go_fuzz_dep_.CoverTab[104391]++
												c.credentials = credentials.New("client", "EXAMPLE.COM").WithPassword("qwerty")
												_, err = c.ASRep.DecryptEncPart(c.credentials)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:95
		_go_fuzz_dep_.CoverTab[104400]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:96
		// _ = "end of CoverTab[104400]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:97
		_go_fuzz_dep_.CoverTab[104401]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:97
		// _ = "end of CoverTab[104401]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:97
	// _ = "end of CoverTab[104391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:97
	_go_fuzz_dep_.CoverTab[104392]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:98
	// _ = "end of CoverTab[104392]"
}

func (c *MockKerberosClient) GetServiceTicket(spn string) (messages.Ticket, types.EncryptionKey, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:101
	_go_fuzz_dep_.CoverTab[104402]++
												if c.errorStage == "service_ticket" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:102
		_go_fuzz_dep_.CoverTab[104404]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:102
		return c.mockError != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:102
		// _ = "end of CoverTab[104404]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:102
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:102
		_go_fuzz_dep_.CoverTab[104405]++
													return messages.Ticket{}, types.EncryptionKey{}, c.mockError
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:103
		// _ = "end of CoverTab[104405]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:104
		_go_fuzz_dep_.CoverTab[104406]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:104
		// _ = "end of CoverTab[104406]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:104
	// _ = "end of CoverTab[104402]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:104
	_go_fuzz_dep_.CoverTab[104403]++
												return c.ASRep.Ticket, c.ASRep.DecryptedEncPart.Key, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:105
	// _ = "end of CoverTab[104403]"
}

func (c *MockKerberosClient) Domain() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:108
	_go_fuzz_dep_.CoverTab[104407]++
												return "EXAMPLE.COM"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:109
	// _ = "end of CoverTab[104407]"
}

func (c *MockKerberosClient) CName() types.PrincipalName {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:112
	_go_fuzz_dep_.CoverTab[104408]++
												p := types.PrincipalName{
		NameType:	KRB5_USER_AUTH,
		NameString: []string{
			"kafka",
			"kafka",
		},
	}
												return p
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:120
	// _ = "end of CoverTab[104408]"
}

func (c *MockKerberosClient) Destroy() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:123
	_go_fuzz_dep_.CoverTab[104409]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:123
	// _ = "end of CoverTab[104409]"

}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:125
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockkerberos.go:125
var _ = _go_fuzz_dep_.CoverTab
