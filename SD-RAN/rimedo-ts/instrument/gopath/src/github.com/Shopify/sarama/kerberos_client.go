//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:1
)

import (
	krb5client "github.com/jcmturner/gokrb5/v8/client"
	krb5config "github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/types"
)

type KerberosGoKrb5Client struct {
	krb5client.Client
}

func (c *KerberosGoKrb5Client) Domain() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:14
	_go_fuzz_dep_.CoverTab[103716]++
												return c.Credentials.Domain()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:15
	// _ = "end of CoverTab[103716]"
}

func (c *KerberosGoKrb5Client) CName() types.PrincipalName {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:18
	_go_fuzz_dep_.CoverTab[103717]++
												return c.Credentials.CName()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:19
	// _ = "end of CoverTab[103717]"
}

// NewKerberosClient creates kerberos client used to obtain TGT and TGS tokens.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:22
// It uses pure go Kerberos 5 solution (RFC-4121 and RFC-4120).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:22
// uses gokrb5 library underlying which is a pure go kerberos client with some GSS-API capabilities.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:25
func NewKerberosClient(config *GSSAPIConfig) (KerberosClient, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:25
	_go_fuzz_dep_.CoverTab[103718]++
												cfg, err := krb5config.Load(config.KerberosConfigPath)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:27
		_go_fuzz_dep_.CoverTab[103720]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:28
		// _ = "end of CoverTab[103720]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:29
		_go_fuzz_dep_.CoverTab[103721]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:29
		// _ = "end of CoverTab[103721]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:29
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:29
	// _ = "end of CoverTab[103718]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:29
	_go_fuzz_dep_.CoverTab[103719]++
												return createClient(config, cfg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:30
	// _ = "end of CoverTab[103719]"
}

func createClient(config *GSSAPIConfig, cfg *krb5config.Config) (KerberosClient, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:33
	_go_fuzz_dep_.CoverTab[103722]++
												var client *krb5client.Client
												if config.AuthType == KRB5_KEYTAB_AUTH {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:35
		_go_fuzz_dep_.CoverTab[103724]++
													kt, err := keytab.Load(config.KeyTabPath)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:37
			_go_fuzz_dep_.CoverTab[103726]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:38
			// _ = "end of CoverTab[103726]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:39
			_go_fuzz_dep_.CoverTab[103727]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:39
			// _ = "end of CoverTab[103727]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:39
		// _ = "end of CoverTab[103724]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:39
		_go_fuzz_dep_.CoverTab[103725]++
													client = krb5client.NewWithKeytab(config.Username, config.Realm, kt, cfg, krb5client.DisablePAFXFAST(config.DisablePAFXFAST))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:40
		// _ = "end of CoverTab[103725]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:41
		_go_fuzz_dep_.CoverTab[103728]++
													client = krb5client.NewWithPassword(config.Username,
			config.Realm, config.Password, cfg, krb5client.DisablePAFXFAST(config.DisablePAFXFAST))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:43
		// _ = "end of CoverTab[103728]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:44
	// _ = "end of CoverTab[103722]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:44
	_go_fuzz_dep_.CoverTab[103723]++
												return &KerberosGoKrb5Client{*client}, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:45
	// _ = "end of CoverTab[103723]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:46
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/kerberos_client.go:46
var _ = _go_fuzz_dep_.CoverTab
