// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Viper is an application configuration system.
// It believes that applications can be configured a variety of ways
// via flags, ENVIRONMENT variables, configuration files retrieved
// from the file system, or a remote key/value store.

// Each item takes precedence over the item below it:

// overrides
// flag
// env
// config
// key/value store
// default

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:20
package viper

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:20
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:20
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:20
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:20
)

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/magiconair/properties"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/pflag"
	"github.com/subosito/gotenv"
	"gopkg.in/ini.v1"

	"github.com/spf13/viper/internal/encoding"
	"github.com/spf13/viper/internal/encoding/hcl"
	"github.com/spf13/viper/internal/encoding/json"
	"github.com/spf13/viper/internal/encoding/toml"
	"github.com/spf13/viper/internal/encoding/yaml"
)

// ConfigMarshalError happens when failing to marshal the configuration.
type ConfigMarshalError struct {
	err error
}

// Error returns the formatted configuration error.
func (e ConfigMarshalError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:60
	_go_fuzz_dep_.CoverTab[129634]++
										return fmt.Sprintf("While marshaling config: %s", e.err.Error())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:61
	// _ = "end of CoverTab[129634]"
}

var v *Viper

type RemoteResponse struct {
	Value	[]byte
	Error	error
}

var (
	encoderRegistry	= encoding.NewEncoderRegistry()
	decoderRegistry	= encoding.NewDecoderRegistry()
)

func init() {
	v = New()

	{
		codec := yaml.Codec{}

		encoderRegistry.RegisterEncoder("yaml", codec)
		decoderRegistry.RegisterDecoder("yaml", codec)

		encoderRegistry.RegisterEncoder("yml", codec)
		decoderRegistry.RegisterDecoder("yml", codec)
	}

	{
		codec := json.Codec{}

		encoderRegistry.RegisterEncoder("json", codec)
		decoderRegistry.RegisterDecoder("json", codec)
	}

	{
		codec := toml.Codec{}

		encoderRegistry.RegisterEncoder("toml", codec)
		decoderRegistry.RegisterDecoder("toml", codec)
	}

	{
		codec := hcl.Codec{}

		encoderRegistry.RegisterEncoder("hcl", codec)
		decoderRegistry.RegisterDecoder("hcl", codec)

		encoderRegistry.RegisterEncoder("tfvars", codec)
		decoderRegistry.RegisterDecoder("tfvars", codec)
	}
}

type remoteConfigFactory interface {
	Get(rp RemoteProvider) (io.Reader, error)
	Watch(rp RemoteProvider) (io.Reader, error)
	WatchChannel(rp RemoteProvider) (<-chan *RemoteResponse, chan bool)
}

// RemoteConfig is optional, see the remote package
var RemoteConfig remoteConfigFactory

// UnsupportedConfigError denotes encountering an unsupported
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:123
// configuration filetype.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:125
type UnsupportedConfigError string

// Error returns the formatted configuration error.
func (str UnsupportedConfigError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:128
	_go_fuzz_dep_.CoverTab[129635]++
										return fmt.Sprintf("Unsupported Config Type %q", string(str))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:129
	// _ = "end of CoverTab[129635]"
}

// UnsupportedRemoteProviderError denotes encountering an unsupported remote
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:132
// provider. Currently only etcd and Consul are supported.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:134
type UnsupportedRemoteProviderError string

// Error returns the formatted remote provider error.
func (str UnsupportedRemoteProviderError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:137
	_go_fuzz_dep_.CoverTab[129636]++
										return fmt.Sprintf("Unsupported Remote Provider Type %q", string(str))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:138
	// _ = "end of CoverTab[129636]"
}

// RemoteConfigError denotes encountering an error while trying to
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:141
// pull the configuration from the remote provider.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:143
type RemoteConfigError string

// Error returns the formatted remote provider error
func (rce RemoteConfigError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:146
	_go_fuzz_dep_.CoverTab[129637]++
										return fmt.Sprintf("Remote Configurations Error: %s", string(rce))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:147
	// _ = "end of CoverTab[129637]"
}

// ConfigFileNotFoundError denotes failing to find configuration file.
type ConfigFileNotFoundError struct {
	name, locations string
}

// Error returns the formatted configuration error.
func (fnfe ConfigFileNotFoundError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:156
	_go_fuzz_dep_.CoverTab[129638]++
										return fmt.Sprintf("Config File %q Not Found in %q", fnfe.name, fnfe.locations)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:157
	// _ = "end of CoverTab[129638]"
}

// ConfigFileAlreadyExistsError denotes failure to write new configuration file.
type ConfigFileAlreadyExistsError string

// Error returns the formatted error when configuration already exists.
func (faee ConfigFileAlreadyExistsError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:164
	_go_fuzz_dep_.CoverTab[129639]++
										return fmt.Sprintf("Config File %q Already Exists", string(faee))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:165
	// _ = "end of CoverTab[129639]"
}

// A DecoderConfigOption can be passed to viper.Unmarshal to configure
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:168
// mapstructure.DecoderConfig options
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:170
type DecoderConfigOption func(*mapstructure.DecoderConfig)

// DecodeHook returns a DecoderConfigOption which overrides the default
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:172
// DecoderConfig.DecodeHook value, the default is:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:172
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:172
//	 mapstructure.ComposeDecodeHookFunc(
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:172
//			mapstructure.StringToTimeDurationHookFunc(),
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:172
//			mapstructure.StringToSliceHookFunc(","),
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:172
//		)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:179
func DecodeHook(hook mapstructure.DecodeHookFunc) DecoderConfigOption {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:179
	_go_fuzz_dep_.CoverTab[129640]++
										return func(c *mapstructure.DecoderConfig) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:180
		_go_fuzz_dep_.CoverTab[129641]++
											c.DecodeHook = hook
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:181
		// _ = "end of CoverTab[129641]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:182
	// _ = "end of CoverTab[129640]"
}

// Viper is a prioritized configuration registry. It
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// maintains a set of configuration sources, fetches
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// values to populate those, and provides them according
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// to the source's priority.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// The priority of the sources is the following:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// 1. overrides
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// 2. flags
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// 3. env. variables
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// 4. config file
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// 5. key/value store
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// 6. defaults
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// For example, if values from the following sources were loaded:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	Defaults : {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"secret": "",
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"user": "default",
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"endpoint": "https://localhost"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	Config : {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"user": "root"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"secret": "defaultsecret"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	Env : {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"secret": "somesecretkey"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// The resulting config will have the following values:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	{
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"secret": "somesecretkey",
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"user": "root",
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//		"endpoint": "https://localhost"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:185
// Note: Vipers are not safe for concurrent Get() and Set() operations.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:221
type Viper struct {
	// Delimiter that separates a list of keys
	// used to access a nested value in one go
	keyDelim	string

	// A set of paths to look for the config file in
	configPaths	[]string

	// The filesystem to read config from.
	fs	afero.Fs

	// A set of remote providers to search for the configuration
	remoteProviders	[]*defaultRemoteProvider

	// Name of file to look for inside the path
	configName		string
	configFile		string
	configType		string
	configPermissions	os.FileMode
	envPrefix		string

	// Specific commands for ini parsing
	iniLoadOptions	ini.LoadOptions

	automaticEnvApplied	bool
	envKeyReplacer		StringReplacer
	allowEmptyEnv		bool

	config		map[string]interface{}
	override	map[string]interface{}
	defaults	map[string]interface{}
	kvstore		map[string]interface{}
	pflags		map[string]FlagValue
	env		map[string][]string
	aliases		map[string]string
	typeByDefValue	bool

	// Store read properties on the object so that we can write back in order with comments.
	// This will only be used if the configuration read is a properties file.
	properties	*properties.Properties

	onConfigChange	func(fsnotify.Event)
}

// New returns an initialized Viper instance.
func New() *Viper {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:266
	_go_fuzz_dep_.CoverTab[129642]++
										v := new(Viper)
										v.keyDelim = "."
										v.configName = "config"
										v.configPermissions = os.FileMode(0644)
										v.fs = afero.NewOsFs()
										v.config = make(map[string]interface{})
										v.override = make(map[string]interface{})
										v.defaults = make(map[string]interface{})
										v.kvstore = make(map[string]interface{})
										v.pflags = make(map[string]FlagValue)
										v.env = make(map[string][]string)
										v.aliases = make(map[string]string)
										v.typeByDefValue = false

										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:281
	// _ = "end of CoverTab[129642]"
}

// Option configures Viper using the functional options paradigm popularized by Rob Pike and Dave Cheney.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:284
// If you're unfamiliar with this style,
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:284
// see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html and
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:284
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:288
type Option interface {
	apply(v *Viper)
}

type optionFunc func(v *Viper)

func (fn optionFunc) apply(v *Viper) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:294
	_go_fuzz_dep_.CoverTab[129643]++
										fn(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:295
	// _ = "end of CoverTab[129643]"
}

// KeyDelimiter sets the delimiter used for determining key parts.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:298
// By default it's value is ".".
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:300
func KeyDelimiter(d string) Option {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:300
	_go_fuzz_dep_.CoverTab[129644]++
										return optionFunc(func(v *Viper) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:301
		_go_fuzz_dep_.CoverTab[129645]++
											v.keyDelim = d
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:302
		// _ = "end of CoverTab[129645]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:303
	// _ = "end of CoverTab[129644]"
}

// StringReplacer applies a set of replacements to a string.
type StringReplacer interface {
	// Replace returns a copy of s with all replacements performed.
	Replace(s string) string
}

// EnvKeyReplacer sets a replacer used for mapping environment variables to internal keys.
func EnvKeyReplacer(r StringReplacer) Option {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:313
	_go_fuzz_dep_.CoverTab[129646]++
										return optionFunc(func(v *Viper) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:314
		_go_fuzz_dep_.CoverTab[129647]++
											v.envKeyReplacer = r
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:315
		// _ = "end of CoverTab[129647]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:316
	// _ = "end of CoverTab[129646]"
}

// NewWithOptions creates a new Viper instance.
func NewWithOptions(opts ...Option) *Viper {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:320
	_go_fuzz_dep_.CoverTab[129648]++
										v := New()

										for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:323
		_go_fuzz_dep_.CoverTab[129650]++
											opt.apply(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:324
		// _ = "end of CoverTab[129650]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:325
	// _ = "end of CoverTab[129648]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:325
	_go_fuzz_dep_.CoverTab[129649]++

										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:327
	// _ = "end of CoverTab[129649]"
}

// Reset is intended for testing, will reset all to default settings.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:330
// In the public interface for the viper package so applications
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:330
// can use it in their testing as well.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:333
func Reset() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:333
	_go_fuzz_dep_.CoverTab[129651]++
										v = New()
										SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}
										SupportedRemoteProviders = []string{"etcd", "consul", "firestore"}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:336
	// _ = "end of CoverTab[129651]"
}

type defaultRemoteProvider struct {
	provider	string
	endpoint	string
	path		string
	secretKeyring	string
}

func (rp defaultRemoteProvider) Provider() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:346
	_go_fuzz_dep_.CoverTab[129652]++
										return rp.provider
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:347
	// _ = "end of CoverTab[129652]"
}

func (rp defaultRemoteProvider) Endpoint() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:350
	_go_fuzz_dep_.CoverTab[129653]++
										return rp.endpoint
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:351
	// _ = "end of CoverTab[129653]"
}

func (rp defaultRemoteProvider) Path() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:354
	_go_fuzz_dep_.CoverTab[129654]++
										return rp.path
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:355
	// _ = "end of CoverTab[129654]"
}

func (rp defaultRemoteProvider) SecretKeyring() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:358
	_go_fuzz_dep_.CoverTab[129655]++
										return rp.secretKeyring
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:359
	// _ = "end of CoverTab[129655]"
}

// RemoteProvider stores the configuration necessary
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:362
// to connect to a remote key/value store.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:362
// Optional secretKeyring to unencrypt encrypted values
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:362
// can be provided.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:366
type RemoteProvider interface {
	Provider() string
	Endpoint() string
	Path() string
	SecretKeyring() string
}

// SupportedExts are universally supported extensions.
var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}

// SupportedRemoteProviders are universally supported remote providers.
var SupportedRemoteProviders = []string{"etcd", "consul", "firestore"}

func OnConfigChange(run func(in fsnotify.Event)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:379
	_go_fuzz_dep_.CoverTab[129656]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:379
	v.OnConfigChange(run)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:379
	// _ = "end of CoverTab[129656]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:379
}
func (v *Viper) OnConfigChange(run func(in fsnotify.Event)) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:380
	_go_fuzz_dep_.CoverTab[129657]++
										v.onConfigChange = run
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:381
	// _ = "end of CoverTab[129657]"
}

func WatchConfig() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:384
	_go_fuzz_dep_.CoverTab[129658]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:384
	v.WatchConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:384
	// _ = "end of CoverTab[129658]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:384
}

func (v *Viper) WatchConfig() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:386
	_go_fuzz_dep_.CoverTab[129659]++
										initWG := sync.WaitGroup{}
										initWG.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:388
	_curRoutineNum154_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:388
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum154_)
										go func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:389
		_go_fuzz_dep_.CoverTab[129661]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:389
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:389
			_go_fuzz_dep_.CoverTab[129665]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:389
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum154_)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:389
			// _ = "end of CoverTab[129665]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:389
		}()
											watcher, err := newWatcher()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:391
			_go_fuzz_dep_.CoverTab[129666]++
												log.Fatal(err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:392
			// _ = "end of CoverTab[129666]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:393
			_go_fuzz_dep_.CoverTab[129667]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:393
			// _ = "end of CoverTab[129667]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:393
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:393
		// _ = "end of CoverTab[129661]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:393
		_go_fuzz_dep_.CoverTab[129662]++
											defer watcher.Close()

											filename, err := v.getConfigFile()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:397
			_go_fuzz_dep_.CoverTab[129668]++
												log.Printf("error: %v\n", err)
												initWG.Done()
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:400
			// _ = "end of CoverTab[129668]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:401
			_go_fuzz_dep_.CoverTab[129669]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:401
			// _ = "end of CoverTab[129669]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:401
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:401
		// _ = "end of CoverTab[129662]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:401
		_go_fuzz_dep_.CoverTab[129663]++

											configFile := filepath.Clean(filename)
											configDir, _ := filepath.Split(configFile)
											realConfigFile, _ := filepath.EvalSymlinks(filename)

											eventsWG := sync.WaitGroup{}
											eventsWG.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:408
		_curRoutineNum155_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:408
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum155_)
											go func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:409
			_go_fuzz_dep_.CoverTab[129670]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:409
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:409
				_go_fuzz_dep_.CoverTab[129671]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:409
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum155_)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:409
				// _ = "end of CoverTab[129671]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:409
			}()
												for {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:410
				_go_fuzz_dep_.CoverTab[129672]++
													select {
				case event, ok := <-watcher.Events:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:412
					_go_fuzz_dep_.CoverTab[129673]++
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:413
						_go_fuzz_dep_.CoverTab[129677]++
															eventsWG.Done()
															return
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:415
						// _ = "end of CoverTab[129677]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:416
						_go_fuzz_dep_.CoverTab[129678]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:416
						// _ = "end of CoverTab[129678]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:416
					}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:416
					// _ = "end of CoverTab[129673]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:416
					_go_fuzz_dep_.CoverTab[129674]++
														currentConfigFile, _ := filepath.EvalSymlinks(filename)
					// we only care about the config file with the following cases:
					// 1 - if the config file was modified or created
					// 2 - if the real path to the config file changed (eg: k8s ConfigMap replacement)
					const writeOrCreateMask = fsnotify.Write | fsnotify.Create
					if (filepath.Clean(event.Name) == configFile && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:422
						_go_fuzz_dep_.CoverTab[129679]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:422
						return event.Op&writeOrCreateMask != 0
															// _ = "end of CoverTab[129679]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:423
					}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:423
						_go_fuzz_dep_.CoverTab[129680]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:423
						return (currentConfigFile != "" && func() bool {
																_go_fuzz_dep_.CoverTab[129681]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:424
							return currentConfigFile != realConfigFile
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:424
							// _ = "end of CoverTab[129681]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:424
						}())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:424
						// _ = "end of CoverTab[129680]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:424
					}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:424
						_go_fuzz_dep_.CoverTab[129682]++
															realConfigFile = currentConfigFile
															err := v.ReadInConfig()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:427
							_go_fuzz_dep_.CoverTab[129684]++
																log.Printf("error reading config file: %v\n", err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:428
							// _ = "end of CoverTab[129684]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:429
							_go_fuzz_dep_.CoverTab[129685]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:429
							// _ = "end of CoverTab[129685]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:429
						}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:429
						// _ = "end of CoverTab[129682]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:429
						_go_fuzz_dep_.CoverTab[129683]++
															if v.onConfigChange != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:430
							_go_fuzz_dep_.CoverTab[129686]++
																v.onConfigChange(event)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:431
							// _ = "end of CoverTab[129686]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:432
							_go_fuzz_dep_.CoverTab[129687]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:432
							// _ = "end of CoverTab[129687]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:432
						}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:432
						// _ = "end of CoverTab[129683]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:433
						_go_fuzz_dep_.CoverTab[129688]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:433
						if filepath.Clean(event.Name) == configFile && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:433
							_go_fuzz_dep_.CoverTab[129689]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:433
							return event.Op&fsnotify.Remove&fsnotify.Remove != 0
																// _ = "end of CoverTab[129689]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:434
						}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:434
							_go_fuzz_dep_.CoverTab[129690]++
																eventsWG.Done()
																return
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:436
							// _ = "end of CoverTab[129690]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:437
							_go_fuzz_dep_.CoverTab[129691]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:437
							// _ = "end of CoverTab[129691]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:437
						}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:437
						// _ = "end of CoverTab[129688]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:437
					}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:437
					// _ = "end of CoverTab[129674]"

				case err, ok := <-watcher.Errors:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:439
					_go_fuzz_dep_.CoverTab[129675]++
														if ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:440
						_go_fuzz_dep_.CoverTab[129692]++
															log.Printf("watcher error: %v\n", err)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:441
						// _ = "end of CoverTab[129692]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:442
						_go_fuzz_dep_.CoverTab[129693]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:442
						// _ = "end of CoverTab[129693]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:442
					}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:442
					// _ = "end of CoverTab[129675]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:442
					_go_fuzz_dep_.CoverTab[129676]++
														eventsWG.Done()
														return
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:444
					// _ = "end of CoverTab[129676]"
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:445
				// _ = "end of CoverTab[129672]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:446
			// _ = "end of CoverTab[129670]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:447
		// _ = "end of CoverTab[129663]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:447
		_go_fuzz_dep_.CoverTab[129664]++
											watcher.Add(configDir)
											initWG.Done()
											eventsWG.Wait()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:450
		// _ = "end of CoverTab[129664]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:451
	// _ = "end of CoverTab[129659]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:451
	_go_fuzz_dep_.CoverTab[129660]++
										initWG.Wait()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:452
	// _ = "end of CoverTab[129660]"
}

// SetConfigFile explicitly defines the path, name and extension of the config file.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:455
// Viper will use this and not check any of the config paths.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:457
func SetConfigFile(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:457
	_go_fuzz_dep_.CoverTab[129694]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:457
	v.SetConfigFile(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:457
	// _ = "end of CoverTab[129694]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:457
}

func (v *Viper) SetConfigFile(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:459
	_go_fuzz_dep_.CoverTab[129695]++
										if in != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:460
		_go_fuzz_dep_.CoverTab[129696]++
											v.configFile = in
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:461
		// _ = "end of CoverTab[129696]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:462
		_go_fuzz_dep_.CoverTab[129697]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:462
		// _ = "end of CoverTab[129697]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:462
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:462
	// _ = "end of CoverTab[129695]"
}

// SetEnvPrefix defines a prefix that ENVIRONMENT variables will use.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:465
// E.g. if your prefix is "spf", the env registry will look for env
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:465
// variables that start with "SPF_".
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:468
func SetEnvPrefix(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:468
	_go_fuzz_dep_.CoverTab[129698]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:468
	v.SetEnvPrefix(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:468
	// _ = "end of CoverTab[129698]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:468
}

func (v *Viper) SetEnvPrefix(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:470
	_go_fuzz_dep_.CoverTab[129699]++
										if in != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:471
		_go_fuzz_dep_.CoverTab[129700]++
											v.envPrefix = in
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:472
		// _ = "end of CoverTab[129700]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:473
		_go_fuzz_dep_.CoverTab[129701]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:473
		// _ = "end of CoverTab[129701]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:473
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:473
	// _ = "end of CoverTab[129699]"
}

func (v *Viper) mergeWithEnvPrefix(in string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:476
	_go_fuzz_dep_.CoverTab[129702]++
										if v.envPrefix != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:477
		_go_fuzz_dep_.CoverTab[129704]++
											return strings.ToUpper(v.envPrefix + "_" + in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:478
		// _ = "end of CoverTab[129704]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:479
		_go_fuzz_dep_.CoverTab[129705]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:479
		// _ = "end of CoverTab[129705]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:479
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:479
	// _ = "end of CoverTab[129702]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:479
	_go_fuzz_dep_.CoverTab[129703]++

										return strings.ToUpper(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:481
	// _ = "end of CoverTab[129703]"
}

// AllowEmptyEnv tells Viper to consider set,
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:484
// but empty environment variables as valid values instead of falling back.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:484
// For backward compatibility reasons this is false by default.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:487
func AllowEmptyEnv(allowEmptyEnv bool) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:487
	_go_fuzz_dep_.CoverTab[129706]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:487
	v.AllowEmptyEnv(allowEmptyEnv)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:487
	// _ = "end of CoverTab[129706]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:487
}

func (v *Viper) AllowEmptyEnv(allowEmptyEnv bool) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:489
	_go_fuzz_dep_.CoverTab[129707]++
										v.allowEmptyEnv = allowEmptyEnv
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:490
	// _ = "end of CoverTab[129707]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:497
// getEnv is a wrapper around os.Getenv which replaces characters in the original
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:497
// key. This allows env vars which have different keys than the config object
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:497
// keys.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:500
func (v *Viper) getEnv(key string) (string, bool) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:500
	_go_fuzz_dep_.CoverTab[129708]++
										if v.envKeyReplacer != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:501
		_go_fuzz_dep_.CoverTab[129710]++
											key = v.envKeyReplacer.Replace(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:502
		// _ = "end of CoverTab[129710]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:503
		_go_fuzz_dep_.CoverTab[129711]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:503
		// _ = "end of CoverTab[129711]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:503
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:503
	// _ = "end of CoverTab[129708]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:503
	_go_fuzz_dep_.CoverTab[129709]++

										val, ok := os.LookupEnv(key)

										return val, ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
		_go_fuzz_dep_.CoverTab[129712]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
		return (v.allowEmptyEnv || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
			_go_fuzz_dep_.CoverTab[129713]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
			return val != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
			// _ = "end of CoverTab[129713]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
		}())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
		// _ = "end of CoverTab[129712]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
	}()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:507
	// _ = "end of CoverTab[129709]"
}

// ConfigFileUsed returns the file used to populate the config registry.
func ConfigFileUsed() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:511
	_go_fuzz_dep_.CoverTab[129714]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:511
	return v.ConfigFileUsed()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:511
	// _ = "end of CoverTab[129714]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:511
}
func (v *Viper) ConfigFileUsed() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:512
	_go_fuzz_dep_.CoverTab[129715]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:512
	return v.configFile
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:512
	// _ = "end of CoverTab[129715]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:512
}

// AddConfigPath adds a path for Viper to search for the config file in.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:514
// Can be called multiple times to define multiple search paths.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:516
func AddConfigPath(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:516
	_go_fuzz_dep_.CoverTab[129716]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:516
	v.AddConfigPath(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:516
	// _ = "end of CoverTab[129716]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:516
}

func (v *Viper) AddConfigPath(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:518
	_go_fuzz_dep_.CoverTab[129717]++
										if in != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:519
		_go_fuzz_dep_.CoverTab[129718]++
											absin := absPathify(in)
											jww.INFO.Println("adding", absin, "to paths to search")
											if !stringInSlice(absin, v.configPaths) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:522
			_go_fuzz_dep_.CoverTab[129719]++
												v.configPaths = append(v.configPaths, absin)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:523
			// _ = "end of CoverTab[129719]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:524
			_go_fuzz_dep_.CoverTab[129720]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:524
			// _ = "end of CoverTab[129720]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:524
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:524
		// _ = "end of CoverTab[129718]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:525
		_go_fuzz_dep_.CoverTab[129721]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:525
		// _ = "end of CoverTab[129721]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:525
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:525
	// _ = "end of CoverTab[129717]"
}

// AddRemoteProvider adds a remote configuration source.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:528
// Remote Providers are searched in the order they are added.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:528
// provider is a string value: "etcd", "consul" or "firestore" are currently supported.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:528
// endpoint is the url.  etcd requires http://ip:port  consul requires ip:port
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:528
// path is the path in the k/v store to retrieve configuration
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:528
// To retrieve a config file called myapp.json from /configs/myapp.json
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:528
// you should set path to /configs and set config name (SetConfigName()) to
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:528
// "myapp"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:536
func AddRemoteProvider(provider, endpoint, path string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:536
	_go_fuzz_dep_.CoverTab[129722]++
										return v.AddRemoteProvider(provider, endpoint, path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:537
	// _ = "end of CoverTab[129722]"
}

func (v *Viper) AddRemoteProvider(provider, endpoint, path string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:540
	_go_fuzz_dep_.CoverTab[129723]++
										if !stringInSlice(provider, SupportedRemoteProviders) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:541
		_go_fuzz_dep_.CoverTab[129726]++
											return UnsupportedRemoteProviderError(provider)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:542
		// _ = "end of CoverTab[129726]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:543
		_go_fuzz_dep_.CoverTab[129727]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:543
		// _ = "end of CoverTab[129727]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:543
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:543
	// _ = "end of CoverTab[129723]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:543
	_go_fuzz_dep_.CoverTab[129724]++
										if provider != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:544
		_go_fuzz_dep_.CoverTab[129728]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:544
		return endpoint != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:544
		// _ = "end of CoverTab[129728]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:544
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:544
		_go_fuzz_dep_.CoverTab[129729]++
											jww.INFO.Printf("adding %s:%s to remote provider list", provider, endpoint)
											rp := &defaultRemoteProvider{
			endpoint:	endpoint,
			provider:	provider,
			path:		path,
		}
		if !v.providerPathExists(rp) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:551
			_go_fuzz_dep_.CoverTab[129730]++
												v.remoteProviders = append(v.remoteProviders, rp)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:552
			// _ = "end of CoverTab[129730]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:553
			_go_fuzz_dep_.CoverTab[129731]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:553
			// _ = "end of CoverTab[129731]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:553
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:553
		// _ = "end of CoverTab[129729]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:554
		_go_fuzz_dep_.CoverTab[129732]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:554
		// _ = "end of CoverTab[129732]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:554
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:554
	// _ = "end of CoverTab[129724]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:554
	_go_fuzz_dep_.CoverTab[129725]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:555
	// _ = "end of CoverTab[129725]"
}

// AddSecureRemoteProvider adds a remote configuration source.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// Secure Remote Providers are searched in the order they are added.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// provider is a string value: "etcd", "consul" or "firestore" are currently supported.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// endpoint is the url.  etcd requires http://ip:port  consul requires ip:port
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// secretkeyring is the filepath to your openpgp secret keyring.  e.g. /etc/secrets/myring.gpg
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// path is the path in the k/v store to retrieve configuration
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// To retrieve a config file called myapp.json from /configs/myapp.json
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// you should set path to /configs and set config name (SetConfigName()) to
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// "myapp"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:558
// Secure Remote Providers are implemented with github.com/bketelsen/crypt
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:568
func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:568
	_go_fuzz_dep_.CoverTab[129733]++
										return v.AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:569
	// _ = "end of CoverTab[129733]"
}

func (v *Viper) AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:572
	_go_fuzz_dep_.CoverTab[129734]++
										if !stringInSlice(provider, SupportedRemoteProviders) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:573
		_go_fuzz_dep_.CoverTab[129737]++
											return UnsupportedRemoteProviderError(provider)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:574
		// _ = "end of CoverTab[129737]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:575
		_go_fuzz_dep_.CoverTab[129738]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:575
		// _ = "end of CoverTab[129738]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:575
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:575
	// _ = "end of CoverTab[129734]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:575
	_go_fuzz_dep_.CoverTab[129735]++
										if provider != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:576
		_go_fuzz_dep_.CoverTab[129739]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:576
		return endpoint != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:576
		// _ = "end of CoverTab[129739]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:576
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:576
		_go_fuzz_dep_.CoverTab[129740]++
											jww.INFO.Printf("adding %s:%s to remote provider list", provider, endpoint)
											rp := &defaultRemoteProvider{
			endpoint:	endpoint,
			provider:	provider,
			path:		path,
			secretKeyring:	secretkeyring,
		}
		if !v.providerPathExists(rp) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:584
			_go_fuzz_dep_.CoverTab[129741]++
												v.remoteProviders = append(v.remoteProviders, rp)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:585
			// _ = "end of CoverTab[129741]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:586
			_go_fuzz_dep_.CoverTab[129742]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:586
			// _ = "end of CoverTab[129742]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:586
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:586
		// _ = "end of CoverTab[129740]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:587
		_go_fuzz_dep_.CoverTab[129743]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:587
		// _ = "end of CoverTab[129743]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:587
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:587
	// _ = "end of CoverTab[129735]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:587
	_go_fuzz_dep_.CoverTab[129736]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:588
	// _ = "end of CoverTab[129736]"
}

func (v *Viper) providerPathExists(p *defaultRemoteProvider) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:591
	_go_fuzz_dep_.CoverTab[129744]++
										for _, y := range v.remoteProviders {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:592
		_go_fuzz_dep_.CoverTab[129746]++
											if reflect.DeepEqual(y, p) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:593
			_go_fuzz_dep_.CoverTab[129747]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:594
			// _ = "end of CoverTab[129747]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:595
			_go_fuzz_dep_.CoverTab[129748]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:595
			// _ = "end of CoverTab[129748]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:595
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:595
		// _ = "end of CoverTab[129746]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:596
	// _ = "end of CoverTab[129744]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:596
	_go_fuzz_dep_.CoverTab[129745]++
										return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:597
	// _ = "end of CoverTab[129745]"
}

// searchMap recursively searches for a value for path in source map.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:600
// Returns nil if not found.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:600
// Note: This assumes that the path entries and map keys are lower cased.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:603
func (v *Viper) searchMap(source map[string]interface{}, path []string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:603
	_go_fuzz_dep_.CoverTab[129749]++
										if len(path) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:604
		_go_fuzz_dep_.CoverTab[129752]++
											return source
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:605
		// _ = "end of CoverTab[129752]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:606
		_go_fuzz_dep_.CoverTab[129753]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:606
		// _ = "end of CoverTab[129753]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:606
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:606
	// _ = "end of CoverTab[129749]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:606
	_go_fuzz_dep_.CoverTab[129750]++

										next, ok := source[path[0]]
										if ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:609
		_go_fuzz_dep_.CoverTab[129754]++

											if len(path) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:611
			_go_fuzz_dep_.CoverTab[129756]++
												return next
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:612
			// _ = "end of CoverTab[129756]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:613
			_go_fuzz_dep_.CoverTab[129757]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:613
			// _ = "end of CoverTab[129757]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:613
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:613
		// _ = "end of CoverTab[129754]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:613
		_go_fuzz_dep_.CoverTab[129755]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:616
		switch next.(type) {
		case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:617
			_go_fuzz_dep_.CoverTab[129758]++
												return v.searchMap(cast.ToStringMap(next), path[1:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:618
			// _ = "end of CoverTab[129758]"
		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:619
			_go_fuzz_dep_.CoverTab[129759]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:622
			return v.searchMap(next.(map[string]interface{}), path[1:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:622
			// _ = "end of CoverTab[129759]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:623
			_go_fuzz_dep_.CoverTab[129760]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:625
			// _ = "end of CoverTab[129760]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:626
		// _ = "end of CoverTab[129755]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:627
		_go_fuzz_dep_.CoverTab[129761]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:627
		// _ = "end of CoverTab[129761]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:627
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:627
	// _ = "end of CoverTab[129750]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:627
	_go_fuzz_dep_.CoverTab[129751]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:628
	// _ = "end of CoverTab[129751]"
}

// searchIndexableWithPathPrefixes recursively searches for a value for path in source map/slice.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
// While searchMap() considers each path element as a single map key or slice index, this
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
// function searches for, and prioritizes, merged path elements.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
// e.g., if in the source, "foo" is defined with a sub-key "bar", and "foo.bar"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
// is also defined, this latter value is returned for path ["foo", "bar"].
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
// This should be useful only at config level (other maps may not contain dots
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
// in their keys).
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:631
// Note: This assumes that the path entries and map keys are lower cased.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:642
func (v *Viper) searchIndexableWithPathPrefixes(source interface{}, path []string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:642
	_go_fuzz_dep_.CoverTab[129762]++
										if len(path) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:643
		_go_fuzz_dep_.CoverTab[129765]++
											return source
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:644
		// _ = "end of CoverTab[129765]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:645
		_go_fuzz_dep_.CoverTab[129766]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:645
		// _ = "end of CoverTab[129766]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:645
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:645
	// _ = "end of CoverTab[129762]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:645
	_go_fuzz_dep_.CoverTab[129763]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:648
	for i := len(path); i > 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:648
		_go_fuzz_dep_.CoverTab[129767]++
											prefixKey := strings.ToLower(strings.Join(path[0:i], v.keyDelim))

											var val interface{}
											switch sourceIndexable := source.(type) {
		case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:653
			_go_fuzz_dep_.CoverTab[129769]++
												val = v.searchSliceWithPathPrefixes(sourceIndexable, prefixKey, i, path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:654
			// _ = "end of CoverTab[129769]"
		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:655
			_go_fuzz_dep_.CoverTab[129770]++
												val = v.searchMapWithPathPrefixes(sourceIndexable, prefixKey, i, path)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:656
			// _ = "end of CoverTab[129770]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:657
		// _ = "end of CoverTab[129767]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:657
		_go_fuzz_dep_.CoverTab[129768]++
											if val != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:658
			_go_fuzz_dep_.CoverTab[129771]++
												return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:659
			// _ = "end of CoverTab[129771]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:660
			_go_fuzz_dep_.CoverTab[129772]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:660
			// _ = "end of CoverTab[129772]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:660
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:660
		// _ = "end of CoverTab[129768]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:661
	// _ = "end of CoverTab[129763]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:661
	_go_fuzz_dep_.CoverTab[129764]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:664
	return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:664
	// _ = "end of CoverTab[129764]"
}

// searchSliceWithPathPrefixes searches for a value for path in sourceSlice
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:667
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:667
// This function is part of the searchIndexableWithPathPrefixes recurring search and
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:667
// should not be called directly from functions other than searchIndexableWithPathPrefixes.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:671
func (v *Viper) searchSliceWithPathPrefixes(
	sourceSlice []interface{},
	prefixKey string,
	pathIndex int,
	path []string,
) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:676
	_go_fuzz_dep_.CoverTab[129773]++

										index, err := strconv.Atoi(prefixKey)
										if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:679
		_go_fuzz_dep_.CoverTab[129777]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:679
		return len(sourceSlice) <= index
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:679
		// _ = "end of CoverTab[129777]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:679
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:679
		_go_fuzz_dep_.CoverTab[129778]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:680
		// _ = "end of CoverTab[129778]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:681
		_go_fuzz_dep_.CoverTab[129779]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:681
		// _ = "end of CoverTab[129779]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:681
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:681
	// _ = "end of CoverTab[129773]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:681
	_go_fuzz_dep_.CoverTab[129774]++

										next := sourceSlice[index]

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:686
	if pathIndex == len(path) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:686
		_go_fuzz_dep_.CoverTab[129780]++
											return next
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:687
		// _ = "end of CoverTab[129780]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:688
		_go_fuzz_dep_.CoverTab[129781]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:688
		// _ = "end of CoverTab[129781]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:688
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:688
	// _ = "end of CoverTab[129774]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:688
	_go_fuzz_dep_.CoverTab[129775]++

										switch n := next.(type) {
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:691
		_go_fuzz_dep_.CoverTab[129782]++
											return v.searchIndexableWithPathPrefixes(cast.ToStringMap(n), path[pathIndex:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:692
		// _ = "end of CoverTab[129782]"
	case map[string]interface{}, []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:693
		_go_fuzz_dep_.CoverTab[129783]++
											return v.searchIndexableWithPathPrefixes(n, path[pathIndex:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:694
		// _ = "end of CoverTab[129783]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:695
		_go_fuzz_dep_.CoverTab[129784]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:695
		// _ = "end of CoverTab[129784]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:697
	// _ = "end of CoverTab[129775]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:697
	_go_fuzz_dep_.CoverTab[129776]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:700
	return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:700
	// _ = "end of CoverTab[129776]"
}

// searchMapWithPathPrefixes searches for a value for path in sourceMap
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:703
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:703
// This function is part of the searchIndexableWithPathPrefixes recurring search and
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:703
// should not be called directly from functions other than searchIndexableWithPathPrefixes.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:707
func (v *Viper) searchMapWithPathPrefixes(
	sourceMap map[string]interface{},
	prefixKey string,
	pathIndex int,
	path []string,
) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:712
	_go_fuzz_dep_.CoverTab[129785]++
										next, ok := sourceMap[prefixKey]
										if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:714
		_go_fuzz_dep_.CoverTab[129789]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:715
		// _ = "end of CoverTab[129789]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:716
		_go_fuzz_dep_.CoverTab[129790]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:716
		// _ = "end of CoverTab[129790]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:716
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:716
	// _ = "end of CoverTab[129785]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:716
	_go_fuzz_dep_.CoverTab[129786]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:719
	if pathIndex == len(path) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:719
		_go_fuzz_dep_.CoverTab[129791]++
											return next
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:720
		// _ = "end of CoverTab[129791]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:721
		_go_fuzz_dep_.CoverTab[129792]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:721
		// _ = "end of CoverTab[129792]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:721
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:721
	// _ = "end of CoverTab[129786]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:721
	_go_fuzz_dep_.CoverTab[129787]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:724
	switch n := next.(type) {
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:725
		_go_fuzz_dep_.CoverTab[129793]++
											return v.searchIndexableWithPathPrefixes(cast.ToStringMap(n), path[pathIndex:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:726
		// _ = "end of CoverTab[129793]"
	case map[string]interface{}, []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:727
		_go_fuzz_dep_.CoverTab[129794]++
											return v.searchIndexableWithPathPrefixes(n, path[pathIndex:])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:728
		// _ = "end of CoverTab[129794]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:729
		_go_fuzz_dep_.CoverTab[129795]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:729
		// _ = "end of CoverTab[129795]"

	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:731
	// _ = "end of CoverTab[129787]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:731
	_go_fuzz_dep_.CoverTab[129788]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:734
	return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:734
	// _ = "end of CoverTab[129788]"
}

// isPathShadowedInDeepMap makes sure the given path is not shadowed somewhere
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:737
// on its path in the map.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:737
// e.g., if "foo.bar" has a value in the given map, it “shadows”
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:737
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:737
//	"foo.bar.baz" in a lower-priority map
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:741
func (v *Viper) isPathShadowedInDeepMap(path []string, m map[string]interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:741
	_go_fuzz_dep_.CoverTab[129796]++
										var parentVal interface{}
										for i := 1; i < len(path); i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:743
		_go_fuzz_dep_.CoverTab[129798]++
											parentVal = v.searchMap(m, path[0:i])
											if parentVal == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:745
			_go_fuzz_dep_.CoverTab[129800]++

												return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:747
			// _ = "end of CoverTab[129800]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:748
			_go_fuzz_dep_.CoverTab[129801]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:748
			// _ = "end of CoverTab[129801]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:748
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:748
		// _ = "end of CoverTab[129798]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:748
		_go_fuzz_dep_.CoverTab[129799]++
											switch parentVal.(type) {
		case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:750
			_go_fuzz_dep_.CoverTab[129802]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:751
			// _ = "end of CoverTab[129802]"
		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:752
			_go_fuzz_dep_.CoverTab[129803]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:753
			// _ = "end of CoverTab[129803]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:754
			_go_fuzz_dep_.CoverTab[129804]++

												return strings.Join(path[0:i], v.keyDelim)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:756
			// _ = "end of CoverTab[129804]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:757
		// _ = "end of CoverTab[129799]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:758
	// _ = "end of CoverTab[129796]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:758
	_go_fuzz_dep_.CoverTab[129797]++
										return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:759
	// _ = "end of CoverTab[129797]"
}

// isPathShadowedInFlatMap makes sure the given path is not shadowed somewhere
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:762
// in a sub-path of the map.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:762
// e.g., if "foo.bar" has a value in the given map, it “shadows”
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:762
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:762
//	"foo.bar.baz" in a lower-priority map
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:766
func (v *Viper) isPathShadowedInFlatMap(path []string, mi interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:766
	_go_fuzz_dep_.CoverTab[129805]++
	// unify input map
	var m map[string]interface{}
	switch mi.(type) {
	case map[string]string, map[string]FlagValue:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:770
		_go_fuzz_dep_.CoverTab[129808]++
											m = cast.ToStringMap(mi)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:771
		// _ = "end of CoverTab[129808]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:772
		_go_fuzz_dep_.CoverTab[129809]++
											return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:773
		// _ = "end of CoverTab[129809]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:774
	// _ = "end of CoverTab[129805]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:774
	_go_fuzz_dep_.CoverTab[129806]++

	// scan paths
	var parentKey string
	for i := 1; i < len(path); i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:778
		_go_fuzz_dep_.CoverTab[129810]++
											parentKey = strings.Join(path[0:i], v.keyDelim)
											if _, ok := m[parentKey]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:780
			_go_fuzz_dep_.CoverTab[129811]++
												return parentKey
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:781
			// _ = "end of CoverTab[129811]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:782
			_go_fuzz_dep_.CoverTab[129812]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:782
			// _ = "end of CoverTab[129812]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:782
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:782
		// _ = "end of CoverTab[129810]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:783
	// _ = "end of CoverTab[129806]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:783
	_go_fuzz_dep_.CoverTab[129807]++
										return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:784
	// _ = "end of CoverTab[129807]"
}

// isPathShadowedInAutoEnv makes sure the given path is not shadowed somewhere
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:787
// in the environment, when automatic env is on.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:787
// e.g., if "foo.bar" has a value in the environment, it “shadows”
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:787
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:787
//	"foo.bar.baz" in a lower-priority map
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:791
func (v *Viper) isPathShadowedInAutoEnv(path []string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:791
	_go_fuzz_dep_.CoverTab[129813]++
										var parentKey string
										for i := 1; i < len(path); i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:793
		_go_fuzz_dep_.CoverTab[129815]++
											parentKey = strings.Join(path[0:i], v.keyDelim)
											if _, ok := v.getEnv(v.mergeWithEnvPrefix(parentKey)); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:795
			_go_fuzz_dep_.CoverTab[129816]++
												return parentKey
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:796
			// _ = "end of CoverTab[129816]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:797
			_go_fuzz_dep_.CoverTab[129817]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:797
			// _ = "end of CoverTab[129817]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:797
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:797
		// _ = "end of CoverTab[129815]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:798
	// _ = "end of CoverTab[129813]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:798
	_go_fuzz_dep_.CoverTab[129814]++
										return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:799
	// _ = "end of CoverTab[129814]"
}

// SetTypeByDefaultValue enables or disables the inference of a key value's
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
// type when the Get function is used based upon a key's default value as
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
// opposed to the value returned based on the normal fetch logic.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
// For example, if a key has a default value of []string{} and the same key
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
// is set via an environment variable to "a b c", a call to the Get function
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
// would return a string slice for the key if the key's type is inferred by
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
// the default value and the Get function would return:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
//	[]string {"a", "b", "c"}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
// Otherwise the Get function would return:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:802
//	"a b c"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:816
func SetTypeByDefaultValue(enable bool) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:816
	_go_fuzz_dep_.CoverTab[129818]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:816
	v.SetTypeByDefaultValue(enable)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:816
	// _ = "end of CoverTab[129818]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:816
}

func (v *Viper) SetTypeByDefaultValue(enable bool) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:818
	_go_fuzz_dep_.CoverTab[129819]++
										v.typeByDefValue = enable
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:819
	// _ = "end of CoverTab[129819]"
}

// GetViper gets the global Viper instance.
func GetViper() *Viper {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:823
	_go_fuzz_dep_.CoverTab[129820]++
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:824
	// _ = "end of CoverTab[129820]"
}

// Get can retrieve any value given the key to use.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:827
// Get is case-insensitive for a key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:827
// Get has the behavior of returning the value associated with the first
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:827
// place from where it is set. Viper will check in the following order:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:827
// override, flag, env, config file, key/value store, default
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:827
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:827
// Get returns an interface. For a specific value use one of the Get____ methods.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:834
func Get(key string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:834
	_go_fuzz_dep_.CoverTab[129821]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:834
	return v.Get(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:834
	// _ = "end of CoverTab[129821]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:834
}

func (v *Viper) Get(key string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:836
	_go_fuzz_dep_.CoverTab[129822]++
										lcaseKey := strings.ToLower(key)
										val := v.find(lcaseKey, true)
										if val == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:839
		_go_fuzz_dep_.CoverTab[129825]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:840
		// _ = "end of CoverTab[129825]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:841
		_go_fuzz_dep_.CoverTab[129826]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:841
		// _ = "end of CoverTab[129826]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:841
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:841
	// _ = "end of CoverTab[129822]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:841
	_go_fuzz_dep_.CoverTab[129823]++

										if v.typeByDefValue {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:843
		_go_fuzz_dep_.CoverTab[129827]++

											valType := val
											path := strings.Split(lcaseKey, v.keyDelim)
											defVal := v.searchMap(v.defaults, path)
											if defVal != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:848
			_go_fuzz_dep_.CoverTab[129829]++
												valType = defVal
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:849
			// _ = "end of CoverTab[129829]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:850
			_go_fuzz_dep_.CoverTab[129830]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:850
			// _ = "end of CoverTab[129830]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:850
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:850
		// _ = "end of CoverTab[129827]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:850
		_go_fuzz_dep_.CoverTab[129828]++

											switch valType.(type) {
		case bool:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:853
			_go_fuzz_dep_.CoverTab[129831]++
												return cast.ToBool(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:854
			// _ = "end of CoverTab[129831]"
		case string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:855
			_go_fuzz_dep_.CoverTab[129832]++
												return cast.ToString(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:856
			// _ = "end of CoverTab[129832]"
		case int32, int16, int8, int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:857
			_go_fuzz_dep_.CoverTab[129833]++
												return cast.ToInt(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:858
			// _ = "end of CoverTab[129833]"
		case uint:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:859
			_go_fuzz_dep_.CoverTab[129834]++
												return cast.ToUint(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:860
			// _ = "end of CoverTab[129834]"
		case uint32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:861
			_go_fuzz_dep_.CoverTab[129835]++
												return cast.ToUint32(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:862
			// _ = "end of CoverTab[129835]"
		case uint64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:863
			_go_fuzz_dep_.CoverTab[129836]++
												return cast.ToUint64(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:864
			// _ = "end of CoverTab[129836]"
		case int64:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:865
			_go_fuzz_dep_.CoverTab[129837]++
												return cast.ToInt64(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:866
			// _ = "end of CoverTab[129837]"
		case float64, float32:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:867
			_go_fuzz_dep_.CoverTab[129838]++
												return cast.ToFloat64(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:868
			// _ = "end of CoverTab[129838]"
		case time.Time:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:869
			_go_fuzz_dep_.CoverTab[129839]++
												return cast.ToTime(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:870
			// _ = "end of CoverTab[129839]"
		case time.Duration:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:871
			_go_fuzz_dep_.CoverTab[129840]++
												return cast.ToDuration(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:872
			// _ = "end of CoverTab[129840]"
		case []string:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:873
			_go_fuzz_dep_.CoverTab[129841]++
												return cast.ToStringSlice(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:874
			// _ = "end of CoverTab[129841]"
		case []int:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:875
			_go_fuzz_dep_.CoverTab[129842]++
												return cast.ToIntSlice(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:876
			// _ = "end of CoverTab[129842]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:877
		// _ = "end of CoverTab[129828]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:878
		_go_fuzz_dep_.CoverTab[129843]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:878
		// _ = "end of CoverTab[129843]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:878
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:878
	// _ = "end of CoverTab[129823]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:878
	_go_fuzz_dep_.CoverTab[129824]++

										return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:880
	// _ = "end of CoverTab[129824]"
}

// Sub returns new Viper instance representing a sub tree of this instance.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:883
// Sub is case-insensitive for a key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:885
func Sub(key string) *Viper {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:885
	_go_fuzz_dep_.CoverTab[129844]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:885
	return v.Sub(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:885
	// _ = "end of CoverTab[129844]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:885
}

func (v *Viper) Sub(key string) *Viper {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:887
	_go_fuzz_dep_.CoverTab[129845]++
										subv := New()
										data := v.Get(key)
										if data == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:890
		_go_fuzz_dep_.CoverTab[129848]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:891
		// _ = "end of CoverTab[129848]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:892
		_go_fuzz_dep_.CoverTab[129849]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:892
		// _ = "end of CoverTab[129849]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:892
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:892
	// _ = "end of CoverTab[129845]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:892
	_go_fuzz_dep_.CoverTab[129846]++

										if reflect.TypeOf(data).Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:894
		_go_fuzz_dep_.CoverTab[129850]++
											subv.config = cast.ToStringMap(data)
											return subv
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:896
		// _ = "end of CoverTab[129850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:897
		_go_fuzz_dep_.CoverTab[129851]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:897
		// _ = "end of CoverTab[129851]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:897
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:897
	// _ = "end of CoverTab[129846]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:897
	_go_fuzz_dep_.CoverTab[129847]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:898
	// _ = "end of CoverTab[129847]"
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:902
	_go_fuzz_dep_.CoverTab[129852]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:902
	return v.GetString(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:902
	// _ = "end of CoverTab[129852]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:902
}

func (v *Viper) GetString(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:904
	_go_fuzz_dep_.CoverTab[129853]++
										return cast.ToString(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:905
	// _ = "end of CoverTab[129853]"
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:909
	_go_fuzz_dep_.CoverTab[129854]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:909
	return v.GetBool(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:909
	// _ = "end of CoverTab[129854]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:909
}

func (v *Viper) GetBool(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:911
	_go_fuzz_dep_.CoverTab[129855]++
										return cast.ToBool(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:912
	// _ = "end of CoverTab[129855]"
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:916
	_go_fuzz_dep_.CoverTab[129856]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:916
	return v.GetInt(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:916
	// _ = "end of CoverTab[129856]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:916
}

func (v *Viper) GetInt(key string) int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:918
	_go_fuzz_dep_.CoverTab[129857]++
										return cast.ToInt(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:919
	// _ = "end of CoverTab[129857]"
}

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:923
	_go_fuzz_dep_.CoverTab[129858]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:923
	return v.GetInt32(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:923
	// _ = "end of CoverTab[129858]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:923
}

func (v *Viper) GetInt32(key string) int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:925
	_go_fuzz_dep_.CoverTab[129859]++
										return cast.ToInt32(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:926
	// _ = "end of CoverTab[129859]"
}

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:930
	_go_fuzz_dep_.CoverTab[129860]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:930
	return v.GetInt64(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:930
	// _ = "end of CoverTab[129860]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:930
}

func (v *Viper) GetInt64(key string) int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:932
	_go_fuzz_dep_.CoverTab[129861]++
										return cast.ToInt64(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:933
	// _ = "end of CoverTab[129861]"
}

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:937
	_go_fuzz_dep_.CoverTab[129862]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:937
	return v.GetUint(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:937
	// _ = "end of CoverTab[129862]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:937
}

func (v *Viper) GetUint(key string) uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:939
	_go_fuzz_dep_.CoverTab[129863]++
										return cast.ToUint(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:940
	// _ = "end of CoverTab[129863]"
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:944
	_go_fuzz_dep_.CoverTab[129864]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:944
	return v.GetUint32(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:944
	// _ = "end of CoverTab[129864]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:944
}

func (v *Viper) GetUint32(key string) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:946
	_go_fuzz_dep_.CoverTab[129865]++
										return cast.ToUint32(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:947
	// _ = "end of CoverTab[129865]"
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:951
	_go_fuzz_dep_.CoverTab[129866]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:951
	return v.GetUint64(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:951
	// _ = "end of CoverTab[129866]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:951
}

func (v *Viper) GetUint64(key string) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:953
	_go_fuzz_dep_.CoverTab[129867]++
										return cast.ToUint64(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:954
	// _ = "end of CoverTab[129867]"
}

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:958
	_go_fuzz_dep_.CoverTab[129868]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:958
	return v.GetFloat64(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:958
	// _ = "end of CoverTab[129868]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:958
}

func (v *Viper) GetFloat64(key string) float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:960
	_go_fuzz_dep_.CoverTab[129869]++
										return cast.ToFloat64(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:961
	// _ = "end of CoverTab[129869]"
}

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:965
	_go_fuzz_dep_.CoverTab[129870]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:965
	return v.GetTime(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:965
	// _ = "end of CoverTab[129870]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:965
}

func (v *Viper) GetTime(key string) time.Time {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:967
	_go_fuzz_dep_.CoverTab[129871]++
										return cast.ToTime(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:968
	// _ = "end of CoverTab[129871]"
}

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:972
	_go_fuzz_dep_.CoverTab[129872]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:972
	return v.GetDuration(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:972
	// _ = "end of CoverTab[129872]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:972
}

func (v *Viper) GetDuration(key string) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:974
	_go_fuzz_dep_.CoverTab[129873]++
										return cast.ToDuration(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:975
	// _ = "end of CoverTab[129873]"
}

// GetIntSlice returns the value associated with the key as a slice of int values.
func GetIntSlice(key string) []int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:979
	_go_fuzz_dep_.CoverTab[129874]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:979
	return v.GetIntSlice(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:979
	// _ = "end of CoverTab[129874]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:979
}

func (v *Viper) GetIntSlice(key string) []int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:981
	_go_fuzz_dep_.CoverTab[129875]++
										return cast.ToIntSlice(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:982
	// _ = "end of CoverTab[129875]"
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:986
	_go_fuzz_dep_.CoverTab[129876]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:986
	return v.GetStringSlice(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:986
	// _ = "end of CoverTab[129876]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:986
}

func (v *Viper) GetStringSlice(key string) []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:988
	_go_fuzz_dep_.CoverTab[129877]++
										return cast.ToStringSlice(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:989
	// _ = "end of CoverTab[129877]"
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:993
	_go_fuzz_dep_.CoverTab[129878]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:993
	return v.GetStringMap(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:993
	// _ = "end of CoverTab[129878]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:993
}

func (v *Viper) GetStringMap(key string) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:995
	_go_fuzz_dep_.CoverTab[129879]++
										return cast.ToStringMap(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:996
	// _ = "end of CoverTab[129879]"
}

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1000
	_go_fuzz_dep_.CoverTab[129880]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1000
	return v.GetStringMapString(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1000
	// _ = "end of CoverTab[129880]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1000
}

func (v *Viper) GetStringMapString(key string) map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1002
	_go_fuzz_dep_.CoverTab[129881]++
										return cast.ToStringMapString(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1003
	// _ = "end of CoverTab[129881]"
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1007
	_go_fuzz_dep_.CoverTab[129882]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1007
	return v.GetStringMapStringSlice(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1007
	// _ = "end of CoverTab[129882]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1007
}

func (v *Viper) GetStringMapStringSlice(key string) map[string][]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1009
	_go_fuzz_dep_.CoverTab[129883]++
										return cast.ToStringMapStringSlice(v.Get(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1010
	// _ = "end of CoverTab[129883]"
}

// GetSizeInBytes returns the size of the value associated with the given key
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1013
// in bytes.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1015
func GetSizeInBytes(key string) uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1015
	_go_fuzz_dep_.CoverTab[129884]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1015
	return v.GetSizeInBytes(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1015
	// _ = "end of CoverTab[129884]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1015
}

func (v *Viper) GetSizeInBytes(key string) uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1017
	_go_fuzz_dep_.CoverTab[129885]++
										sizeStr := cast.ToString(v.Get(key))
										return parseSizeInBytes(sizeStr)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1019
	// _ = "end of CoverTab[129885]"
}

// UnmarshalKey takes a single key and unmarshals it into a Struct.
func UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1023
	_go_fuzz_dep_.CoverTab[129886]++
										return v.UnmarshalKey(key, rawVal, opts...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1024
	// _ = "end of CoverTab[129886]"
}

func (v *Viper) UnmarshalKey(key string, rawVal interface{}, opts ...DecoderConfigOption) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1027
	_go_fuzz_dep_.CoverTab[129887]++
										return decode(v.Get(key), defaultDecoderConfig(rawVal, opts...))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1028
	// _ = "end of CoverTab[129887]"
}

// Unmarshal unmarshals the config into a Struct. Make sure that the tags
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1031
// on the fields of the structure are properly set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1033
func Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1033
	_go_fuzz_dep_.CoverTab[129888]++
										return v.Unmarshal(rawVal, opts...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1034
	// _ = "end of CoverTab[129888]"
}

func (v *Viper) Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1037
	_go_fuzz_dep_.CoverTab[129889]++
										return decode(v.AllSettings(), defaultDecoderConfig(rawVal, opts...))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1038
	// _ = "end of CoverTab[129889]"
}

// defaultDecoderConfig returns default mapsstructure.DecoderConfig with suppot
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1041
// of time.Duration values & string slices
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1043
func defaultDecoderConfig(output interface{}, opts ...DecoderConfigOption) *mapstructure.DecoderConfig {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1043
	_go_fuzz_dep_.CoverTab[129890]++
										c := &mapstructure.DecoderConfig{
		Metadata:		nil,
		Result:			output,
		WeaklyTypedInput:	true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1053
		_go_fuzz_dep_.CoverTab[129892]++
											opt(c)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1054
		// _ = "end of CoverTab[129892]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1055
	// _ = "end of CoverTab[129890]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1055
	_go_fuzz_dep_.CoverTab[129891]++
										return c
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1056
	// _ = "end of CoverTab[129891]"
}

// A wrapper around mapstructure.Decode that mimics the WeakDecode functionality
func decode(input interface{}, config *mapstructure.DecoderConfig) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1060
	_go_fuzz_dep_.CoverTab[129893]++
										decoder, err := mapstructure.NewDecoder(config)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1062
		_go_fuzz_dep_.CoverTab[129895]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1063
		// _ = "end of CoverTab[129895]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1064
		_go_fuzz_dep_.CoverTab[129896]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1064
		// _ = "end of CoverTab[129896]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1064
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1064
	// _ = "end of CoverTab[129893]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1064
	_go_fuzz_dep_.CoverTab[129894]++
										return decoder.Decode(input)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1065
	// _ = "end of CoverTab[129894]"
}

// UnmarshalExact unmarshals the config into a Struct, erroring if a field is nonexistent
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1068
// in the destination struct.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1070
func UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1070
	_go_fuzz_dep_.CoverTab[129897]++
										return v.UnmarshalExact(rawVal, opts...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1071
	// _ = "end of CoverTab[129897]"
}

func (v *Viper) UnmarshalExact(rawVal interface{}, opts ...DecoderConfigOption) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1074
	_go_fuzz_dep_.CoverTab[129898]++
										config := defaultDecoderConfig(rawVal, opts...)
										config.ErrorUnused = true

										return decode(v.AllSettings(), config)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1078
	// _ = "end of CoverTab[129898]"
}

// BindPFlags binds a full flag set to the configuration, using each flag's long
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1081
// name as the config key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1083
func BindPFlags(flags *pflag.FlagSet) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1083
	_go_fuzz_dep_.CoverTab[129899]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1083
	return v.BindPFlags(flags)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1083
	// _ = "end of CoverTab[129899]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1083
}

func (v *Viper) BindPFlags(flags *pflag.FlagSet) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1085
	_go_fuzz_dep_.CoverTab[129900]++
										return v.BindFlagValues(pflagValueSet{flags})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1086
	// _ = "end of CoverTab[129900]"
}

// BindPFlag binds a specific key to a pflag (as used by cobra).
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1089
// Example (where serverCmd is a Cobra instance):
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1089
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1089
//	serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1089
//	Viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1095
func BindPFlag(key string, flag *pflag.Flag) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1095
	_go_fuzz_dep_.CoverTab[129901]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1095
	return v.BindPFlag(key, flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1095
	// _ = "end of CoverTab[129901]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1095
}

func (v *Viper) BindPFlag(key string, flag *pflag.Flag) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1097
	_go_fuzz_dep_.CoverTab[129902]++
										if flag == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1098
		_go_fuzz_dep_.CoverTab[129904]++
											return fmt.Errorf("flag for %q is nil", key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1099
		// _ = "end of CoverTab[129904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1100
		_go_fuzz_dep_.CoverTab[129905]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1100
		// _ = "end of CoverTab[129905]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1100
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1100
	// _ = "end of CoverTab[129902]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1100
	_go_fuzz_dep_.CoverTab[129903]++
										return v.BindFlagValue(key, pflagValue{flag})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1101
	// _ = "end of CoverTab[129903]"
}

// BindFlagValues binds a full FlagValue set to the configuration, using each flag's long
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1104
// name as the config key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1106
func BindFlagValues(flags FlagValueSet) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1106
	_go_fuzz_dep_.CoverTab[129906]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1106
	return v.BindFlagValues(flags)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1106
	// _ = "end of CoverTab[129906]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1106
}

func (v *Viper) BindFlagValues(flags FlagValueSet) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1108
	_go_fuzz_dep_.CoverTab[129907]++
										flags.VisitAll(func(flag FlagValue) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1109
		_go_fuzz_dep_.CoverTab[129909]++
											if err = v.BindFlagValue(flag.Name(), flag); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1110
			_go_fuzz_dep_.CoverTab[129910]++
												return
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1111
			// _ = "end of CoverTab[129910]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1112
			_go_fuzz_dep_.CoverTab[129911]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1112
			// _ = "end of CoverTab[129911]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1112
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1112
		// _ = "end of CoverTab[129909]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1113
	// _ = "end of CoverTab[129907]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1113
	_go_fuzz_dep_.CoverTab[129908]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1114
	// _ = "end of CoverTab[129908]"
}

// BindFlagValue binds a specific key to a FlagValue.
func BindFlagValue(key string, flag FlagValue) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1118
	_go_fuzz_dep_.CoverTab[129912]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1118
	return v.BindFlagValue(key, flag)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1118
	// _ = "end of CoverTab[129912]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1118
}

func (v *Viper) BindFlagValue(key string, flag FlagValue) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1120
	_go_fuzz_dep_.CoverTab[129913]++
										if flag == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1121
		_go_fuzz_dep_.CoverTab[129915]++
											return fmt.Errorf("flag for %q is nil", key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1122
		// _ = "end of CoverTab[129915]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1123
		_go_fuzz_dep_.CoverTab[129916]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1123
		// _ = "end of CoverTab[129916]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1123
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1123
	// _ = "end of CoverTab[129913]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1123
	_go_fuzz_dep_.CoverTab[129914]++
										v.pflags[strings.ToLower(key)] = flag
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1125
	// _ = "end of CoverTab[129914]"
}

// BindEnv binds a Viper key to a ENV variable.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1128
// ENV variables are case sensitive.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1128
// If only a key is provided, it will use the env key matching the key, uppercased.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1128
// If more arguments are provided, they will represent the env variable names that
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1128
// should bind to this key and will be taken in the specified order.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1128
// EnvPrefix will be used when set when env name is not provided.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1134
func BindEnv(input ...string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1134
	_go_fuzz_dep_.CoverTab[129917]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1134
	return v.BindEnv(input...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1134
	// _ = "end of CoverTab[129917]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1134
}

func (v *Viper) BindEnv(input ...string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1136
	_go_fuzz_dep_.CoverTab[129918]++
										if len(input) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1137
		_go_fuzz_dep_.CoverTab[129921]++
											return fmt.Errorf("missing key to bind to")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1138
		// _ = "end of CoverTab[129921]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1139
		_go_fuzz_dep_.CoverTab[129922]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1139
		// _ = "end of CoverTab[129922]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1139
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1139
	// _ = "end of CoverTab[129918]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1139
	_go_fuzz_dep_.CoverTab[129919]++

										key := strings.ToLower(input[0])

										if len(input) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1143
		_go_fuzz_dep_.CoverTab[129923]++
											v.env[key] = append(v.env[key], v.mergeWithEnvPrefix(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1144
		// _ = "end of CoverTab[129923]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1145
		_go_fuzz_dep_.CoverTab[129924]++
											v.env[key] = append(v.env[key], input[1:]...)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1146
		// _ = "end of CoverTab[129924]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1147
	// _ = "end of CoverTab[129919]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1147
	_go_fuzz_dep_.CoverTab[129920]++

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1149
	// _ = "end of CoverTab[129920]"
}

// Given a key, find the value.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
// Viper will check to see if an alias exists first.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
// Viper will then check in the following order:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
// flag, env, config file, key/value store.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
// Lastly, if no value was found and flagDefault is true, and if the key
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
// corresponds to a flag, the flag's default value is returned.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1152
// Note: this assumes a lower-cased key given.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1161
func (v *Viper) find(lcaseKey string, flagDefault bool) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1161
	_go_fuzz_dep_.CoverTab[129925]++
										var (
		val	interface{}
		exists	bool
		path	= strings.Split(lcaseKey, v.keyDelim)
		nested	= len(path) > 1
	)

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1170
	if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1170
		_go_fuzz_dep_.CoverTab[129941]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1170
		return v.isPathShadowedInDeepMap(path, castMapStringToMapInterface(v.aliases)) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1170
		// _ = "end of CoverTab[129941]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1170
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1170
		_go_fuzz_dep_.CoverTab[129942]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1171
		// _ = "end of CoverTab[129942]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1172
		_go_fuzz_dep_.CoverTab[129943]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1172
		// _ = "end of CoverTab[129943]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1172
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1172
	// _ = "end of CoverTab[129925]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1172
	_go_fuzz_dep_.CoverTab[129926]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1175
	lcaseKey = v.realKey(lcaseKey)
										path = strings.Split(lcaseKey, v.keyDelim)
										nested = len(path) > 1

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1180
	val = v.searchMap(v.override, path)
	if val != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1181
		_go_fuzz_dep_.CoverTab[129944]++
											return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1182
		// _ = "end of CoverTab[129944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1183
		_go_fuzz_dep_.CoverTab[129945]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1183
		// _ = "end of CoverTab[129945]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1183
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1183
	// _ = "end of CoverTab[129926]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1183
	_go_fuzz_dep_.CoverTab[129927]++
										if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1184
		_go_fuzz_dep_.CoverTab[129946]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1184
		return v.isPathShadowedInDeepMap(path, v.override) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1184
		// _ = "end of CoverTab[129946]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1184
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1184
		_go_fuzz_dep_.CoverTab[129947]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1185
		// _ = "end of CoverTab[129947]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1186
		_go_fuzz_dep_.CoverTab[129948]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1186
		// _ = "end of CoverTab[129948]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1186
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1186
	// _ = "end of CoverTab[129927]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1186
	_go_fuzz_dep_.CoverTab[129928]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1189
	flag, exists := v.pflags[lcaseKey]
	if exists && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1190
		_go_fuzz_dep_.CoverTab[129949]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1190
		return flag.HasChanged()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1190
		// _ = "end of CoverTab[129949]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1190
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1190
		_go_fuzz_dep_.CoverTab[129950]++
											switch flag.ValueType() {
		case "int", "int8", "int16", "int32", "int64":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1192
			_go_fuzz_dep_.CoverTab[129951]++
												return cast.ToInt(flag.ValueString())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1193
			// _ = "end of CoverTab[129951]"
		case "bool":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1194
			_go_fuzz_dep_.CoverTab[129952]++
												return cast.ToBool(flag.ValueString())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1195
			// _ = "end of CoverTab[129952]"
		case "stringSlice", "stringArray":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1196
			_go_fuzz_dep_.CoverTab[129953]++
												s := strings.TrimPrefix(flag.ValueString(), "[")
												s = strings.TrimSuffix(s, "]")
												res, _ := readAsCSV(s)
												return res
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1200
			// _ = "end of CoverTab[129953]"
		case "intSlice":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1201
			_go_fuzz_dep_.CoverTab[129954]++
												s := strings.TrimPrefix(flag.ValueString(), "[")
												s = strings.TrimSuffix(s, "]")
												res, _ := readAsCSV(s)
												return cast.ToIntSlice(res)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1205
			// _ = "end of CoverTab[129954]"
		case "stringToString":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1206
			_go_fuzz_dep_.CoverTab[129955]++
												return stringToStringConv(flag.ValueString())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1207
			// _ = "end of CoverTab[129955]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1208
			_go_fuzz_dep_.CoverTab[129956]++
												return flag.ValueString()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1209
			// _ = "end of CoverTab[129956]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1210
		// _ = "end of CoverTab[129950]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1211
		_go_fuzz_dep_.CoverTab[129957]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1211
		// _ = "end of CoverTab[129957]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1211
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1211
	// _ = "end of CoverTab[129928]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1211
	_go_fuzz_dep_.CoverTab[129929]++
										if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1212
		_go_fuzz_dep_.CoverTab[129958]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1212
		return v.isPathShadowedInFlatMap(path, v.pflags) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1212
		// _ = "end of CoverTab[129958]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1212
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1212
		_go_fuzz_dep_.CoverTab[129959]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1213
		// _ = "end of CoverTab[129959]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1214
		_go_fuzz_dep_.CoverTab[129960]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1214
		// _ = "end of CoverTab[129960]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1214
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1214
	// _ = "end of CoverTab[129929]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1214
	_go_fuzz_dep_.CoverTab[129930]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1217
	if v.automaticEnvApplied {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1217
		_go_fuzz_dep_.CoverTab[129961]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1220
		if val, ok := v.getEnv(v.mergeWithEnvPrefix(lcaseKey)); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1220
			_go_fuzz_dep_.CoverTab[129963]++
												return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1221
			// _ = "end of CoverTab[129963]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1222
			_go_fuzz_dep_.CoverTab[129964]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1222
			// _ = "end of CoverTab[129964]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1222
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1222
		// _ = "end of CoverTab[129961]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1222
		_go_fuzz_dep_.CoverTab[129962]++
											if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1223
			_go_fuzz_dep_.CoverTab[129965]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1223
			return v.isPathShadowedInAutoEnv(path) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1223
			// _ = "end of CoverTab[129965]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1223
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1223
			_go_fuzz_dep_.CoverTab[129966]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1224
			// _ = "end of CoverTab[129966]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1225
			_go_fuzz_dep_.CoverTab[129967]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1225
			// _ = "end of CoverTab[129967]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1225
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1225
		// _ = "end of CoverTab[129962]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1226
		_go_fuzz_dep_.CoverTab[129968]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1226
		// _ = "end of CoverTab[129968]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1226
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1226
	// _ = "end of CoverTab[129930]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1226
	_go_fuzz_dep_.CoverTab[129931]++
										envkeys, exists := v.env[lcaseKey]
										if exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1228
		_go_fuzz_dep_.CoverTab[129969]++
											for _, envkey := range envkeys {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1229
			_go_fuzz_dep_.CoverTab[129970]++
												if val, ok := v.getEnv(envkey); ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1230
				_go_fuzz_dep_.CoverTab[129971]++
													return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1231
				// _ = "end of CoverTab[129971]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1232
				_go_fuzz_dep_.CoverTab[129972]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1232
				// _ = "end of CoverTab[129972]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1232
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1232
			// _ = "end of CoverTab[129970]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1233
		// _ = "end of CoverTab[129969]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1234
		_go_fuzz_dep_.CoverTab[129973]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1234
		// _ = "end of CoverTab[129973]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1234
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1234
	// _ = "end of CoverTab[129931]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1234
	_go_fuzz_dep_.CoverTab[129932]++
										if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1235
		_go_fuzz_dep_.CoverTab[129974]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1235
		return v.isPathShadowedInFlatMap(path, v.env) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1235
		// _ = "end of CoverTab[129974]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1235
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1235
		_go_fuzz_dep_.CoverTab[129975]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1236
		// _ = "end of CoverTab[129975]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1237
		_go_fuzz_dep_.CoverTab[129976]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1237
		// _ = "end of CoverTab[129976]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1237
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1237
	// _ = "end of CoverTab[129932]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1237
	_go_fuzz_dep_.CoverTab[129933]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1240
	val = v.searchIndexableWithPathPrefixes(v.config, path)
	if val != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1241
		_go_fuzz_dep_.CoverTab[129977]++
											return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1242
		// _ = "end of CoverTab[129977]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1243
		_go_fuzz_dep_.CoverTab[129978]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1243
		// _ = "end of CoverTab[129978]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1243
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1243
	// _ = "end of CoverTab[129933]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1243
	_go_fuzz_dep_.CoverTab[129934]++
										if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1244
		_go_fuzz_dep_.CoverTab[129979]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1244
		return v.isPathShadowedInDeepMap(path, v.config) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1244
		// _ = "end of CoverTab[129979]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1244
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1244
		_go_fuzz_dep_.CoverTab[129980]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1245
		// _ = "end of CoverTab[129980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1246
		_go_fuzz_dep_.CoverTab[129981]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1246
		// _ = "end of CoverTab[129981]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1246
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1246
	// _ = "end of CoverTab[129934]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1246
	_go_fuzz_dep_.CoverTab[129935]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1249
	val = v.searchMap(v.kvstore, path)
	if val != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1250
		_go_fuzz_dep_.CoverTab[129982]++
											return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1251
		// _ = "end of CoverTab[129982]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1252
		_go_fuzz_dep_.CoverTab[129983]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1252
		// _ = "end of CoverTab[129983]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1252
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1252
	// _ = "end of CoverTab[129935]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1252
	_go_fuzz_dep_.CoverTab[129936]++
										if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1253
		_go_fuzz_dep_.CoverTab[129984]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1253
		return v.isPathShadowedInDeepMap(path, v.kvstore) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1253
		// _ = "end of CoverTab[129984]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1253
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1253
		_go_fuzz_dep_.CoverTab[129985]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1254
		// _ = "end of CoverTab[129985]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1255
		_go_fuzz_dep_.CoverTab[129986]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1255
		// _ = "end of CoverTab[129986]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1255
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1255
	// _ = "end of CoverTab[129936]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1255
	_go_fuzz_dep_.CoverTab[129937]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1258
	val = v.searchMap(v.defaults, path)
	if val != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1259
		_go_fuzz_dep_.CoverTab[129987]++
											return val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1260
		// _ = "end of CoverTab[129987]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1261
		_go_fuzz_dep_.CoverTab[129988]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1261
		// _ = "end of CoverTab[129988]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1261
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1261
	// _ = "end of CoverTab[129937]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1261
	_go_fuzz_dep_.CoverTab[129938]++
										if nested && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1262
		_go_fuzz_dep_.CoverTab[129989]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1262
		return v.isPathShadowedInDeepMap(path, v.defaults) != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1262
		// _ = "end of CoverTab[129989]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1262
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1262
		_go_fuzz_dep_.CoverTab[129990]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1263
		// _ = "end of CoverTab[129990]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1264
		_go_fuzz_dep_.CoverTab[129991]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1264
		// _ = "end of CoverTab[129991]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1264
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1264
	// _ = "end of CoverTab[129938]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1264
	_go_fuzz_dep_.CoverTab[129939]++

										if flagDefault {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1266
		_go_fuzz_dep_.CoverTab[129992]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1269
		if flag, exists := v.pflags[lcaseKey]; exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1269
			_go_fuzz_dep_.CoverTab[129993]++
												switch flag.ValueType() {
			case "int", "int8", "int16", "int32", "int64":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1271
				_go_fuzz_dep_.CoverTab[129994]++
													return cast.ToInt(flag.ValueString())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1272
				// _ = "end of CoverTab[129994]"
			case "bool":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1273
				_go_fuzz_dep_.CoverTab[129995]++
													return cast.ToBool(flag.ValueString())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1274
				// _ = "end of CoverTab[129995]"
			case "stringSlice", "stringArray":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1275
				_go_fuzz_dep_.CoverTab[129996]++
													s := strings.TrimPrefix(flag.ValueString(), "[")
													s = strings.TrimSuffix(s, "]")
													res, _ := readAsCSV(s)
													return res
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1279
				// _ = "end of CoverTab[129996]"
			case "intSlice":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1280
				_go_fuzz_dep_.CoverTab[129997]++
													s := strings.TrimPrefix(flag.ValueString(), "[")
													s = strings.TrimSuffix(s, "]")
													res, _ := readAsCSV(s)
													return cast.ToIntSlice(res)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1284
				// _ = "end of CoverTab[129997]"
			case "stringToString":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1285
				_go_fuzz_dep_.CoverTab[129998]++
													return stringToStringConv(flag.ValueString())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1286
				// _ = "end of CoverTab[129998]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1287
				_go_fuzz_dep_.CoverTab[129999]++
													return flag.ValueString()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1288
				// _ = "end of CoverTab[129999]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1289
			// _ = "end of CoverTab[129993]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1290
			_go_fuzz_dep_.CoverTab[130000]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1290
			// _ = "end of CoverTab[130000]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1290
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1290
		// _ = "end of CoverTab[129992]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1292
		_go_fuzz_dep_.CoverTab[130001]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1292
		// _ = "end of CoverTab[130001]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1292
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1292
	// _ = "end of CoverTab[129939]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1292
	_go_fuzz_dep_.CoverTab[129940]++

										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1294
	// _ = "end of CoverTab[129940]"
}

func readAsCSV(val string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1297
	_go_fuzz_dep_.CoverTab[130002]++
										if val == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1298
		_go_fuzz_dep_.CoverTab[130004]++
											return []string{}, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1299
		// _ = "end of CoverTab[130004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1300
		_go_fuzz_dep_.CoverTab[130005]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1300
		// _ = "end of CoverTab[130005]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1300
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1300
	// _ = "end of CoverTab[130002]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1300
	_go_fuzz_dep_.CoverTab[130003]++
										stringReader := strings.NewReader(val)
										csvReader := csv.NewReader(stringReader)
										return csvReader.Read()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1303
	// _ = "end of CoverTab[130003]"
}

// mostly copied from pflag's implementation of this operation here https://github.com/spf13/pflag/blob/master/string_to_string.go#L79
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1306
// alterations are: errors are swallowed, map[string]interface{} is returned in order to enable cast.ToStringMap
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1308
func stringToStringConv(val string) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1308
	_go_fuzz_dep_.CoverTab[130006]++
										val = strings.Trim(val, "[]")

										if len(val) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1311
		_go_fuzz_dep_.CoverTab[130010]++
											return map[string]interface{}{}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1312
		// _ = "end of CoverTab[130010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1313
		_go_fuzz_dep_.CoverTab[130011]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1313
		// _ = "end of CoverTab[130011]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1313
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1313
	// _ = "end of CoverTab[130006]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1313
	_go_fuzz_dep_.CoverTab[130007]++
										r := csv.NewReader(strings.NewReader(val))
										ss, err := r.Read()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1316
		_go_fuzz_dep_.CoverTab[130012]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1317
		// _ = "end of CoverTab[130012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1318
		_go_fuzz_dep_.CoverTab[130013]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1318
		// _ = "end of CoverTab[130013]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1318
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1318
	// _ = "end of CoverTab[130007]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1318
	_go_fuzz_dep_.CoverTab[130008]++
										out := make(map[string]interface{}, len(ss))
										for _, pair := range ss {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1320
		_go_fuzz_dep_.CoverTab[130014]++
											kv := strings.SplitN(pair, "=", 2)
											if len(kv) != 2 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1322
			_go_fuzz_dep_.CoverTab[130016]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1323
			// _ = "end of CoverTab[130016]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1324
			_go_fuzz_dep_.CoverTab[130017]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1324
			// _ = "end of CoverTab[130017]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1324
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1324
		// _ = "end of CoverTab[130014]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1324
		_go_fuzz_dep_.CoverTab[130015]++
											out[kv[0]] = kv[1]
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1325
		// _ = "end of CoverTab[130015]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1326
	// _ = "end of CoverTab[130008]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1326
	_go_fuzz_dep_.CoverTab[130009]++
										return out
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1327
	// _ = "end of CoverTab[130009]"
}

// IsSet checks to see if the key has been set in any of the data locations.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1330
// IsSet is case-insensitive for a key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1332
func IsSet(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1332
	_go_fuzz_dep_.CoverTab[130018]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1332
	return v.IsSet(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1332
	// _ = "end of CoverTab[130018]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1332
}

func (v *Viper) IsSet(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1334
	_go_fuzz_dep_.CoverTab[130019]++
										lcaseKey := strings.ToLower(key)
										val := v.find(lcaseKey, false)
										return val != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1337
	// _ = "end of CoverTab[130019]"
}

// AutomaticEnv makes Viper check if environment variables match any of the existing keys
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1340
// (config, default or flags). If matching env vars are found, they are loaded into Viper.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1342
func AutomaticEnv() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1342
	_go_fuzz_dep_.CoverTab[130020]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1342
	v.AutomaticEnv()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1342
	// _ = "end of CoverTab[130020]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1342
}

func (v *Viper) AutomaticEnv() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1344
	_go_fuzz_dep_.CoverTab[130021]++
										v.automaticEnvApplied = true
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1345
	// _ = "end of CoverTab[130021]"
}

// SetEnvKeyReplacer sets the strings.Replacer on the viper object
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1348
// Useful for mapping an environmental variable to a key that does
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1348
// not match it.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1351
func SetEnvKeyReplacer(r *strings.Replacer) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1351
	_go_fuzz_dep_.CoverTab[130022]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1351
	v.SetEnvKeyReplacer(r)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1351
	// _ = "end of CoverTab[130022]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1351
}

func (v *Viper) SetEnvKeyReplacer(r *strings.Replacer) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1353
	_go_fuzz_dep_.CoverTab[130023]++
										v.envKeyReplacer = r
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1354
	// _ = "end of CoverTab[130023]"
}

// RegisterAlias creates an alias that provides another accessor for the same key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1357
// This enables one to change a name without breaking the application.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1359
func RegisterAlias(alias string, key string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1359
	_go_fuzz_dep_.CoverTab[130024]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1359
	v.RegisterAlias(alias, key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1359
	// _ = "end of CoverTab[130024]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1359
}

func (v *Viper) RegisterAlias(alias string, key string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1361
	_go_fuzz_dep_.CoverTab[130025]++
										v.registerAlias(alias, strings.ToLower(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1362
	// _ = "end of CoverTab[130025]"
}

func (v *Viper) registerAlias(alias string, key string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1365
	_go_fuzz_dep_.CoverTab[130026]++
										alias = strings.ToLower(alias)
										if alias != key && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1367
		_go_fuzz_dep_.CoverTab[130027]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1367
		return alias != v.realKey(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1367
		// _ = "end of CoverTab[130027]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1367
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1367
		_go_fuzz_dep_.CoverTab[130028]++
											_, exists := v.aliases[alias]

											if !exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1370
			_go_fuzz_dep_.CoverTab[130029]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1374
			if val, ok := v.config[alias]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1374
				_go_fuzz_dep_.CoverTab[130034]++
													delete(v.config, alias)
													v.config[key] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1376
				// _ = "end of CoverTab[130034]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1377
				_go_fuzz_dep_.CoverTab[130035]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1377
				// _ = "end of CoverTab[130035]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1377
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1377
			// _ = "end of CoverTab[130029]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1377
			_go_fuzz_dep_.CoverTab[130030]++
												if val, ok := v.kvstore[alias]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1378
				_go_fuzz_dep_.CoverTab[130036]++
													delete(v.kvstore, alias)
													v.kvstore[key] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1380
				// _ = "end of CoverTab[130036]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1381
				_go_fuzz_dep_.CoverTab[130037]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1381
				// _ = "end of CoverTab[130037]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1381
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1381
			// _ = "end of CoverTab[130030]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1381
			_go_fuzz_dep_.CoverTab[130031]++
												if val, ok := v.defaults[alias]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1382
				_go_fuzz_dep_.CoverTab[130038]++
													delete(v.defaults, alias)
													v.defaults[key] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1384
				// _ = "end of CoverTab[130038]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1385
				_go_fuzz_dep_.CoverTab[130039]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1385
				// _ = "end of CoverTab[130039]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1385
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1385
			// _ = "end of CoverTab[130031]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1385
			_go_fuzz_dep_.CoverTab[130032]++
												if val, ok := v.override[alias]; ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1386
				_go_fuzz_dep_.CoverTab[130040]++
													delete(v.override, alias)
													v.override[key] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1388
				// _ = "end of CoverTab[130040]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1389
				_go_fuzz_dep_.CoverTab[130041]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1389
				// _ = "end of CoverTab[130041]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1389
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1389
			// _ = "end of CoverTab[130032]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1389
			_go_fuzz_dep_.CoverTab[130033]++
												v.aliases[alias] = key
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1390
			// _ = "end of CoverTab[130033]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1391
			_go_fuzz_dep_.CoverTab[130042]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1391
			// _ = "end of CoverTab[130042]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1391
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1391
		// _ = "end of CoverTab[130028]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1392
		_go_fuzz_dep_.CoverTab[130043]++
											jww.WARN.Println("Creating circular reference alias", alias, key, v.realKey(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1393
		// _ = "end of CoverTab[130043]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1394
	// _ = "end of CoverTab[130026]"
}

func (v *Viper) realKey(key string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1397
	_go_fuzz_dep_.CoverTab[130044]++
										newkey, exists := v.aliases[key]
										if exists {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1399
		_go_fuzz_dep_.CoverTab[130046]++
											jww.DEBUG.Println("Alias", key, "to", newkey)
											return v.realKey(newkey)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1401
		// _ = "end of CoverTab[130046]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1402
		_go_fuzz_dep_.CoverTab[130047]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1402
		// _ = "end of CoverTab[130047]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1402
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1402
	// _ = "end of CoverTab[130044]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1402
	_go_fuzz_dep_.CoverTab[130045]++
										return key
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1403
	// _ = "end of CoverTab[130045]"
}

// InConfig checks to see if the given key (or an alias) is in the config file.
func InConfig(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1407
	_go_fuzz_dep_.CoverTab[130048]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1407
	return v.InConfig(key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1407
	// _ = "end of CoverTab[130048]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1407
}

func (v *Viper) InConfig(key string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1409
	_go_fuzz_dep_.CoverTab[130049]++
										lcaseKey := strings.ToLower(key)

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1413
	lcaseKey = v.realKey(lcaseKey)
										path := strings.Split(lcaseKey, v.keyDelim)

										return v.searchIndexableWithPathPrefixes(v.config, path) != nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1416
	// _ = "end of CoverTab[130049]"
}

// SetDefault sets the default value for this key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1419
// SetDefault is case-insensitive for a key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1419
// Default only used when no value is provided by the user via flag, config or ENV.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1422
func SetDefault(key string, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1422
	_go_fuzz_dep_.CoverTab[130050]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1422
	v.SetDefault(key, value)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1422
	// _ = "end of CoverTab[130050]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1422
}

func (v *Viper) SetDefault(key string, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1424
	_go_fuzz_dep_.CoverTab[130051]++

										key = v.realKey(strings.ToLower(key))
										value = toCaseInsensitiveValue(value)

										path := strings.Split(key, v.keyDelim)
										lastKey := strings.ToLower(path[len(path)-1])
										deepestMap := deepSearch(v.defaults, path[0:len(path)-1])

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1434
	deepestMap[lastKey] = value
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1434
	// _ = "end of CoverTab[130051]"
}

// Set sets the value for the key in the override register.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1437
// Set is case-insensitive for a key.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1437
// Will be used instead of values obtained via
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1437
// flags, config file, ENV, default, or key/value store.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1441
func Set(key string, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1441
	_go_fuzz_dep_.CoverTab[130052]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1441
	v.Set(key, value)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1441
	// _ = "end of CoverTab[130052]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1441
}

func (v *Viper) Set(key string, value interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1443
	_go_fuzz_dep_.CoverTab[130053]++

										key = v.realKey(strings.ToLower(key))
										value = toCaseInsensitiveValue(value)

										path := strings.Split(key, v.keyDelim)
										lastKey := strings.ToLower(path[len(path)-1])
										deepestMap := deepSearch(v.override, path[0:len(path)-1])

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1453
	deepestMap[lastKey] = value
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1453
	// _ = "end of CoverTab[130053]"
}

// ReadInConfig will discover and load the configuration file from disk
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1456
// and key/value stores, searching in one of the defined paths.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1458
func ReadInConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1458
	_go_fuzz_dep_.CoverTab[130054]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1458
	return v.ReadInConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1458
	// _ = "end of CoverTab[130054]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1458
}

func (v *Viper) ReadInConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1460
	_go_fuzz_dep_.CoverTab[130055]++
										jww.INFO.Println("Attempting to read in config file")
										filename, err := v.getConfigFile()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1463
		_go_fuzz_dep_.CoverTab[130060]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1464
		// _ = "end of CoverTab[130060]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1465
		_go_fuzz_dep_.CoverTab[130061]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1465
		// _ = "end of CoverTab[130061]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1465
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1465
	// _ = "end of CoverTab[130055]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1465
	_go_fuzz_dep_.CoverTab[130056]++

										if !stringInSlice(v.getConfigType(), SupportedExts) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1467
		_go_fuzz_dep_.CoverTab[130062]++
											return UnsupportedConfigError(v.getConfigType())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1468
		// _ = "end of CoverTab[130062]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1469
		_go_fuzz_dep_.CoverTab[130063]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1469
		// _ = "end of CoverTab[130063]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1469
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1469
	// _ = "end of CoverTab[130056]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1469
	_go_fuzz_dep_.CoverTab[130057]++

										jww.DEBUG.Println("Reading file: ", filename)
										file, err := afero.ReadFile(v.fs, filename)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1473
		_go_fuzz_dep_.CoverTab[130064]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1474
		// _ = "end of CoverTab[130064]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1475
		_go_fuzz_dep_.CoverTab[130065]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1475
		// _ = "end of CoverTab[130065]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1475
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1475
	// _ = "end of CoverTab[130057]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1475
	_go_fuzz_dep_.CoverTab[130058]++

										config := make(map[string]interface{})

										err = v.unmarshalReader(bytes.NewReader(file), config)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1480
		_go_fuzz_dep_.CoverTab[130066]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1481
		// _ = "end of CoverTab[130066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1482
		_go_fuzz_dep_.CoverTab[130067]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1482
		// _ = "end of CoverTab[130067]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1482
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1482
	// _ = "end of CoverTab[130058]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1482
	_go_fuzz_dep_.CoverTab[130059]++

										v.config = config
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1485
	// _ = "end of CoverTab[130059]"
}

// MergeInConfig merges a new configuration with an existing config.
func MergeInConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1489
	_go_fuzz_dep_.CoverTab[130068]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1489
	return v.MergeInConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1489
	// _ = "end of CoverTab[130068]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1489
}

func (v *Viper) MergeInConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1491
	_go_fuzz_dep_.CoverTab[130069]++
										jww.INFO.Println("Attempting to merge in config file")
										filename, err := v.getConfigFile()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1494
		_go_fuzz_dep_.CoverTab[130073]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1495
		// _ = "end of CoverTab[130073]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1496
		_go_fuzz_dep_.CoverTab[130074]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1496
		// _ = "end of CoverTab[130074]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1496
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1496
	// _ = "end of CoverTab[130069]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1496
	_go_fuzz_dep_.CoverTab[130070]++

										if !stringInSlice(v.getConfigType(), SupportedExts) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1498
		_go_fuzz_dep_.CoverTab[130075]++
											return UnsupportedConfigError(v.getConfigType())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1499
		// _ = "end of CoverTab[130075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1500
		_go_fuzz_dep_.CoverTab[130076]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1500
		// _ = "end of CoverTab[130076]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1500
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1500
	// _ = "end of CoverTab[130070]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1500
	_go_fuzz_dep_.CoverTab[130071]++

										file, err := afero.ReadFile(v.fs, filename)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1503
		_go_fuzz_dep_.CoverTab[130077]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1504
		// _ = "end of CoverTab[130077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1505
		_go_fuzz_dep_.CoverTab[130078]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1505
		// _ = "end of CoverTab[130078]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1505
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1505
	// _ = "end of CoverTab[130071]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1505
	_go_fuzz_dep_.CoverTab[130072]++

										return v.MergeConfig(bytes.NewReader(file))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1507
	// _ = "end of CoverTab[130072]"
}

// ReadConfig will read a configuration file, setting existing keys to nil if the
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1510
// key does not exist in the file.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1512
func ReadConfig(in io.Reader) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1512
	_go_fuzz_dep_.CoverTab[130079]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1512
	return v.ReadConfig(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1512
	// _ = "end of CoverTab[130079]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1512
}

func (v *Viper) ReadConfig(in io.Reader) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1514
	_go_fuzz_dep_.CoverTab[130080]++
										v.config = make(map[string]interface{})
										return v.unmarshalReader(in, v.config)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1516
	// _ = "end of CoverTab[130080]"
}

// MergeConfig merges a new configuration with an existing config.
func MergeConfig(in io.Reader) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1520
	_go_fuzz_dep_.CoverTab[130081]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1520
	return v.MergeConfig(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1520
	// _ = "end of CoverTab[130081]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1520
}

func (v *Viper) MergeConfig(in io.Reader) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1522
	_go_fuzz_dep_.CoverTab[130082]++
										cfg := make(map[string]interface{})
										if err := v.unmarshalReader(in, cfg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1524
		_go_fuzz_dep_.CoverTab[130084]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1525
		// _ = "end of CoverTab[130084]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1526
		_go_fuzz_dep_.CoverTab[130085]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1526
		// _ = "end of CoverTab[130085]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1526
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1526
	// _ = "end of CoverTab[130082]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1526
	_go_fuzz_dep_.CoverTab[130083]++
										return v.MergeConfigMap(cfg)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1527
	// _ = "end of CoverTab[130083]"
}

// MergeConfigMap merges the configuration from the map given with an existing config.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1530
// Note that the map given may be modified.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1532
func MergeConfigMap(cfg map[string]interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1532
	_go_fuzz_dep_.CoverTab[130086]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1532
	return v.MergeConfigMap(cfg)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1532
	// _ = "end of CoverTab[130086]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1532
}

func (v *Viper) MergeConfigMap(cfg map[string]interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1534
	_go_fuzz_dep_.CoverTab[130087]++
										if v.config == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1535
		_go_fuzz_dep_.CoverTab[130089]++
											v.config = make(map[string]interface{})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1536
		// _ = "end of CoverTab[130089]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1537
		_go_fuzz_dep_.CoverTab[130090]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1537
		// _ = "end of CoverTab[130090]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1537
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1537
	// _ = "end of CoverTab[130087]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1537
	_go_fuzz_dep_.CoverTab[130088]++
										insensitiviseMap(cfg)
										mergeMaps(cfg, v.config, nil)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1540
	// _ = "end of CoverTab[130088]"
}

// WriteConfig writes the current configuration to a file.
func WriteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1544
	_go_fuzz_dep_.CoverTab[130091]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1544
	return v.WriteConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1544
	// _ = "end of CoverTab[130091]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1544
}

func (v *Viper) WriteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1546
	_go_fuzz_dep_.CoverTab[130092]++
										filename, err := v.getConfigFile()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1548
		_go_fuzz_dep_.CoverTab[130094]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1549
		// _ = "end of CoverTab[130094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1550
		_go_fuzz_dep_.CoverTab[130095]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1550
		// _ = "end of CoverTab[130095]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1550
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1550
	// _ = "end of CoverTab[130092]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1550
	_go_fuzz_dep_.CoverTab[130093]++
										return v.writeConfig(filename, true)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1551
	// _ = "end of CoverTab[130093]"
}

// SafeWriteConfig writes current configuration to file only if the file does not exist.
func SafeWriteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1555
	_go_fuzz_dep_.CoverTab[130096]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1555
	return v.SafeWriteConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1555
	// _ = "end of CoverTab[130096]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1555
}

func (v *Viper) SafeWriteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1557
	_go_fuzz_dep_.CoverTab[130097]++
										if len(v.configPaths) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1558
		_go_fuzz_dep_.CoverTab[130099]++
											return errors.New("missing configuration for 'configPath'")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1559
		// _ = "end of CoverTab[130099]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1560
		_go_fuzz_dep_.CoverTab[130100]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1560
		// _ = "end of CoverTab[130100]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1560
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1560
	// _ = "end of CoverTab[130097]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1560
	_go_fuzz_dep_.CoverTab[130098]++
										return v.SafeWriteConfigAs(filepath.Join(v.configPaths[0], v.configName+"."+v.configType))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1561
	// _ = "end of CoverTab[130098]"
}

// WriteConfigAs writes current configuration to a given filename.
func WriteConfigAs(filename string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1565
	_go_fuzz_dep_.CoverTab[130101]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1565
	return v.WriteConfigAs(filename)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1565
	// _ = "end of CoverTab[130101]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1565
}

func (v *Viper) WriteConfigAs(filename string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1567
	_go_fuzz_dep_.CoverTab[130102]++
										return v.writeConfig(filename, true)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1568
	// _ = "end of CoverTab[130102]"
}

// SafeWriteConfigAs writes current configuration to a given filename if it does not exist.
func SafeWriteConfigAs(filename string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1572
	_go_fuzz_dep_.CoverTab[130103]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1572
	return v.SafeWriteConfigAs(filename)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1572
	// _ = "end of CoverTab[130103]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1572
}

func (v *Viper) SafeWriteConfigAs(filename string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1574
	_go_fuzz_dep_.CoverTab[130104]++
										alreadyExists, err := afero.Exists(v.fs, filename)
										if alreadyExists && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1576
		_go_fuzz_dep_.CoverTab[130106]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1576
		return err == nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1576
		// _ = "end of CoverTab[130106]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1576
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1576
		_go_fuzz_dep_.CoverTab[130107]++
											return ConfigFileAlreadyExistsError(filename)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1577
		// _ = "end of CoverTab[130107]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1578
		_go_fuzz_dep_.CoverTab[130108]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1578
		// _ = "end of CoverTab[130108]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1578
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1578
	// _ = "end of CoverTab[130104]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1578
	_go_fuzz_dep_.CoverTab[130105]++
										return v.writeConfig(filename, false)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1579
	// _ = "end of CoverTab[130105]"
}

func (v *Viper) writeConfig(filename string, force bool) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1582
	_go_fuzz_dep_.CoverTab[130109]++
										jww.INFO.Println("Attempting to write configuration to file.")
										var configType string

										ext := filepath.Ext(filename)
										if ext != "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1587
		_go_fuzz_dep_.CoverTab[130117]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1587
		return ext != filepath.Base(filename)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1587
		// _ = "end of CoverTab[130117]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1587
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1587
		_go_fuzz_dep_.CoverTab[130118]++
											configType = ext[1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1588
		// _ = "end of CoverTab[130118]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1589
		_go_fuzz_dep_.CoverTab[130119]++
											configType = v.configType
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1590
		// _ = "end of CoverTab[130119]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1591
	// _ = "end of CoverTab[130109]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1591
	_go_fuzz_dep_.CoverTab[130110]++
										if configType == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1592
		_go_fuzz_dep_.CoverTab[130120]++
											return fmt.Errorf("config type could not be determined for %s", filename)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1593
		// _ = "end of CoverTab[130120]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1594
		_go_fuzz_dep_.CoverTab[130121]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1594
		// _ = "end of CoverTab[130121]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1594
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1594
	// _ = "end of CoverTab[130110]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1594
	_go_fuzz_dep_.CoverTab[130111]++

										if !stringInSlice(configType, SupportedExts) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1596
		_go_fuzz_dep_.CoverTab[130122]++
											return UnsupportedConfigError(configType)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1597
		// _ = "end of CoverTab[130122]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1598
		_go_fuzz_dep_.CoverTab[130123]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1598
		// _ = "end of CoverTab[130123]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1598
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1598
	// _ = "end of CoverTab[130111]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1598
	_go_fuzz_dep_.CoverTab[130112]++
										if v.config == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1599
		_go_fuzz_dep_.CoverTab[130124]++
											v.config = make(map[string]interface{})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1600
		// _ = "end of CoverTab[130124]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1601
		_go_fuzz_dep_.CoverTab[130125]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1601
		// _ = "end of CoverTab[130125]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1601
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1601
	// _ = "end of CoverTab[130112]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1601
	_go_fuzz_dep_.CoverTab[130113]++
										flags := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
										if !force {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1603
		_go_fuzz_dep_.CoverTab[130126]++
											flags |= os.O_EXCL
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1604
		// _ = "end of CoverTab[130126]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1605
		_go_fuzz_dep_.CoverTab[130127]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1605
		// _ = "end of CoverTab[130127]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1605
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1605
	// _ = "end of CoverTab[130113]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1605
	_go_fuzz_dep_.CoverTab[130114]++
										f, err := v.fs.OpenFile(filename, flags, v.configPermissions)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1607
		_go_fuzz_dep_.CoverTab[130128]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1608
		// _ = "end of CoverTab[130128]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1609
		_go_fuzz_dep_.CoverTab[130129]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1609
		// _ = "end of CoverTab[130129]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1609
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1609
	// _ = "end of CoverTab[130114]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1609
	_go_fuzz_dep_.CoverTab[130115]++
										defer f.Close()

										if err := v.marshalWriter(f, configType); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1612
		_go_fuzz_dep_.CoverTab[130130]++
											return err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1613
		// _ = "end of CoverTab[130130]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1614
		_go_fuzz_dep_.CoverTab[130131]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1614
		// _ = "end of CoverTab[130131]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1614
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1614
	// _ = "end of CoverTab[130115]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1614
	_go_fuzz_dep_.CoverTab[130116]++

										return f.Sync()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1616
	// _ = "end of CoverTab[130116]"
}

// Unmarshal a Reader into a map.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1619
// Should probably be an unexported function.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1621
func unmarshalReader(in io.Reader, c map[string]interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1621
	_go_fuzz_dep_.CoverTab[130132]++
										return v.unmarshalReader(in, c)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1622
	// _ = "end of CoverTab[130132]"
}

func (v *Viper) unmarshalReader(in io.Reader, c map[string]interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1625
	_go_fuzz_dep_.CoverTab[130133]++
										buf := new(bytes.Buffer)
										buf.ReadFrom(in)

										switch format := strings.ToLower(v.getConfigType()); format {
	case "yaml", "yml", "json", "toml", "hcl", "tfvars":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1630
		_go_fuzz_dep_.CoverTab[130135]++
											err := decoderRegistry.Decode(format, buf.Bytes(), &c)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1632
			_go_fuzz_dep_.CoverTab[130143]++
												return ConfigParseError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1633
			// _ = "end of CoverTab[130143]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1634
			_go_fuzz_dep_.CoverTab[130144]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1634
			// _ = "end of CoverTab[130144]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1634
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1634
		// _ = "end of CoverTab[130135]"

	case "dotenv", "env":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1636
		_go_fuzz_dep_.CoverTab[130136]++
											env, err := gotenv.StrictParse(buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1638
			_go_fuzz_dep_.CoverTab[130145]++
												return ConfigParseError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1639
			// _ = "end of CoverTab[130145]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1640
			_go_fuzz_dep_.CoverTab[130146]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1640
			// _ = "end of CoverTab[130146]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1640
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1640
		// _ = "end of CoverTab[130136]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1640
		_go_fuzz_dep_.CoverTab[130137]++
											for k, v := range env {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1641
			_go_fuzz_dep_.CoverTab[130147]++
												c[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1642
			// _ = "end of CoverTab[130147]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1643
		// _ = "end of CoverTab[130137]"

	case "properties", "props", "prop":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1645
		_go_fuzz_dep_.CoverTab[130138]++
											v.properties = properties.NewProperties()
											var err error
											if v.properties, err = properties.Load(buf.Bytes(), properties.UTF8); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1648
			_go_fuzz_dep_.CoverTab[130148]++
												return ConfigParseError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1649
			// _ = "end of CoverTab[130148]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1650
			_go_fuzz_dep_.CoverTab[130149]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1650
			// _ = "end of CoverTab[130149]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1650
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1650
		// _ = "end of CoverTab[130138]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1650
		_go_fuzz_dep_.CoverTab[130139]++
											for _, key := range v.properties.Keys() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1651
			_go_fuzz_dep_.CoverTab[130150]++
												value, _ := v.properties.Get(key)

												path := strings.Split(key, ".")
												lastKey := strings.ToLower(path[len(path)-1])
												deepestMap := deepSearch(c, path[0:len(path)-1])

												deepestMap[lastKey] = value
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1658
			// _ = "end of CoverTab[130150]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1659
		// _ = "end of CoverTab[130139]"

	case "ini":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1661
		_go_fuzz_dep_.CoverTab[130140]++
											cfg := ini.Empty(v.iniLoadOptions)
											err := cfg.Append(buf.Bytes())
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1664
			_go_fuzz_dep_.CoverTab[130151]++
												return ConfigParseError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1665
			// _ = "end of CoverTab[130151]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1666
			_go_fuzz_dep_.CoverTab[130152]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1666
			// _ = "end of CoverTab[130152]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1666
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1666
		// _ = "end of CoverTab[130140]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1666
		_go_fuzz_dep_.CoverTab[130141]++
											sections := cfg.Sections()
											for i := 0; i < len(sections); i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1668
			_go_fuzz_dep_.CoverTab[130153]++
												section := sections[i]
												keys := section.Keys()
												for j := 0; j < len(keys); j++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1671
				_go_fuzz_dep_.CoverTab[130154]++
													key := keys[j]
													value := cfg.Section(section.Name()).Key(key.Name()).String()
													c[section.Name()+"."+key.Name()] = value
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1674
				// _ = "end of CoverTab[130154]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1675
			// _ = "end of CoverTab[130153]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1676
		// _ = "end of CoverTab[130141]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1676
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1676
		_go_fuzz_dep_.CoverTab[130142]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1676
		// _ = "end of CoverTab[130142]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1677
	// _ = "end of CoverTab[130133]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1677
	_go_fuzz_dep_.CoverTab[130134]++

										insensitiviseMap(c)
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1680
	// _ = "end of CoverTab[130134]"
}

// Marshal a map into Writer.
func (v *Viper) marshalWriter(f afero.File, configType string) error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1684
	_go_fuzz_dep_.CoverTab[130155]++
										c := v.AllSettings()
										switch configType {
	case "yaml", "yml", "json", "toml", "hcl", "tfvars":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1687
		_go_fuzz_dep_.CoverTab[130157]++
											b, err := encoderRegistry.Encode(configType, c)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1689
			_go_fuzz_dep_.CoverTab[130167]++
												return ConfigMarshalError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1690
			// _ = "end of CoverTab[130167]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1691
			_go_fuzz_dep_.CoverTab[130168]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1691
			// _ = "end of CoverTab[130168]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1691
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1691
		// _ = "end of CoverTab[130157]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1691
		_go_fuzz_dep_.CoverTab[130158]++

											_, err = f.WriteString(string(b))
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1694
			_go_fuzz_dep_.CoverTab[130169]++
												return ConfigMarshalError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1695
			// _ = "end of CoverTab[130169]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1696
			_go_fuzz_dep_.CoverTab[130170]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1696
			// _ = "end of CoverTab[130170]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1696
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1696
		// _ = "end of CoverTab[130158]"

	case "prop", "props", "properties":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1698
		_go_fuzz_dep_.CoverTab[130159]++
											if v.properties == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1699
			_go_fuzz_dep_.CoverTab[130171]++
												v.properties = properties.NewProperties()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1700
			// _ = "end of CoverTab[130171]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1701
			_go_fuzz_dep_.CoverTab[130172]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1701
			// _ = "end of CoverTab[130172]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1701
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1701
		// _ = "end of CoverTab[130159]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1701
		_go_fuzz_dep_.CoverTab[130160]++
											p := v.properties
											for _, key := range v.AllKeys() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1703
			_go_fuzz_dep_.CoverTab[130173]++
												_, _, err := p.Set(key, v.GetString(key))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1705
				_go_fuzz_dep_.CoverTab[130174]++
													return ConfigMarshalError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1706
				// _ = "end of CoverTab[130174]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1707
				_go_fuzz_dep_.CoverTab[130175]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1707
				// _ = "end of CoverTab[130175]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1707
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1707
			// _ = "end of CoverTab[130173]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1708
		// _ = "end of CoverTab[130160]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1708
		_go_fuzz_dep_.CoverTab[130161]++
											_, err := p.WriteComment(f, "#", properties.UTF8)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1710
			_go_fuzz_dep_.CoverTab[130176]++
												return ConfigMarshalError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1711
			// _ = "end of CoverTab[130176]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1712
			_go_fuzz_dep_.CoverTab[130177]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1712
			// _ = "end of CoverTab[130177]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1712
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1712
		// _ = "end of CoverTab[130161]"

	case "dotenv", "env":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1714
		_go_fuzz_dep_.CoverTab[130162]++
											lines := []string{}
											for _, key := range v.AllKeys() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1716
			_go_fuzz_dep_.CoverTab[130178]++
												envName := strings.ToUpper(strings.Replace(key, ".", "_", -1))
												val := v.Get(key)
												lines = append(lines, fmt.Sprintf("%v=%v", envName, val))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1719
			// _ = "end of CoverTab[130178]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1720
		// _ = "end of CoverTab[130162]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1720
		_go_fuzz_dep_.CoverTab[130163]++
											s := strings.Join(lines, "\n")
											if _, err := f.WriteString(s); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1722
			_go_fuzz_dep_.CoverTab[130179]++
												return ConfigMarshalError{err}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1723
			// _ = "end of CoverTab[130179]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1724
			_go_fuzz_dep_.CoverTab[130180]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1724
			// _ = "end of CoverTab[130180]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1724
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1724
		// _ = "end of CoverTab[130163]"

	case "ini":
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1726
		_go_fuzz_dep_.CoverTab[130164]++
											keys := v.AllKeys()
											cfg := ini.Empty()
											ini.PrettyFormat = false
											for i := 0; i < len(keys); i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1730
			_go_fuzz_dep_.CoverTab[130181]++
												key := keys[i]
												lastSep := strings.LastIndex(key, ".")
												sectionName := key[:(lastSep)]
												keyName := key[(lastSep + 1):]
												if sectionName == "default" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1735
				_go_fuzz_dep_.CoverTab[130183]++
													sectionName = ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1736
				// _ = "end of CoverTab[130183]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1737
				_go_fuzz_dep_.CoverTab[130184]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1737
				// _ = "end of CoverTab[130184]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1737
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1737
			// _ = "end of CoverTab[130181]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1737
			_go_fuzz_dep_.CoverTab[130182]++
												cfg.Section(sectionName).Key(keyName).SetValue(v.GetString(key))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1738
			// _ = "end of CoverTab[130182]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1739
		// _ = "end of CoverTab[130164]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1739
		_go_fuzz_dep_.CoverTab[130165]++
											cfg.WriteTo(f)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1740
		// _ = "end of CoverTab[130165]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1740
	default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1740
		_go_fuzz_dep_.CoverTab[130166]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1740
		// _ = "end of CoverTab[130166]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1741
	// _ = "end of CoverTab[130155]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1741
	_go_fuzz_dep_.CoverTab[130156]++
										return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1742
	// _ = "end of CoverTab[130156]"
}

func keyExists(k string, m map[string]interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1745
	_go_fuzz_dep_.CoverTab[130185]++
										lk := strings.ToLower(k)
										for mk := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1747
		_go_fuzz_dep_.CoverTab[130187]++
											lmk := strings.ToLower(mk)
											if lmk == lk {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1749
			_go_fuzz_dep_.CoverTab[130188]++
												return mk
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1750
			// _ = "end of CoverTab[130188]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1751
			_go_fuzz_dep_.CoverTab[130189]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1751
			// _ = "end of CoverTab[130189]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1751
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1751
		// _ = "end of CoverTab[130187]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1752
	// _ = "end of CoverTab[130185]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1752
	_go_fuzz_dep_.CoverTab[130186]++
										return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1753
	// _ = "end of CoverTab[130186]"
}

func castToMapStringInterface(
	src map[interface{}]interface{}) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1757
	_go_fuzz_dep_.CoverTab[130190]++
										tgt := map[string]interface{}{}
										for k, v := range src {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1759
		_go_fuzz_dep_.CoverTab[130192]++
											tgt[fmt.Sprintf("%v", k)] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1760
		// _ = "end of CoverTab[130192]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1761
	// _ = "end of CoverTab[130190]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1761
	_go_fuzz_dep_.CoverTab[130191]++
										return tgt
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1762
	// _ = "end of CoverTab[130191]"
}

func castMapStringSliceToMapInterface(src map[string][]string) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1765
	_go_fuzz_dep_.CoverTab[130193]++
										tgt := map[string]interface{}{}
										for k, v := range src {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1767
		_go_fuzz_dep_.CoverTab[130195]++
											tgt[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1768
		// _ = "end of CoverTab[130195]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1769
	// _ = "end of CoverTab[130193]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1769
	_go_fuzz_dep_.CoverTab[130194]++
										return tgt
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1770
	// _ = "end of CoverTab[130194]"
}

func castMapStringToMapInterface(src map[string]string) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1773
	_go_fuzz_dep_.CoverTab[130196]++
										tgt := map[string]interface{}{}
										for k, v := range src {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1775
		_go_fuzz_dep_.CoverTab[130198]++
											tgt[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1776
		// _ = "end of CoverTab[130198]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1777
	// _ = "end of CoverTab[130196]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1777
	_go_fuzz_dep_.CoverTab[130197]++
										return tgt
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1778
	// _ = "end of CoverTab[130197]"
}

func castMapFlagToMapInterface(src map[string]FlagValue) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1781
	_go_fuzz_dep_.CoverTab[130199]++
										tgt := map[string]interface{}{}
										for k, v := range src {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1783
		_go_fuzz_dep_.CoverTab[130201]++
											tgt[k] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1784
		// _ = "end of CoverTab[130201]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1785
	// _ = "end of CoverTab[130199]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1785
	_go_fuzz_dep_.CoverTab[130200]++
										return tgt
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1786
	// _ = "end of CoverTab[130200]"
}

// mergeMaps merges two maps. The `itgt` parameter is for handling go-yaml's
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1789
// insistence on parsing nested structures as `map[interface{}]interface{}`
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1789
// instead of using a `string` as the key for nest structures beyond one level
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1789
// deep. Both map types are supported as there is a go-yaml fork that uses
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1789
// `map[string]interface{}` instead.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1794
func mergeMaps(
	src, tgt map[string]interface{}, itgt map[interface{}]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1795
	_go_fuzz_dep_.CoverTab[130202]++
										for sk, sv := range src {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1796
		_go_fuzz_dep_.CoverTab[130203]++
											tk := keyExists(sk, tgt)
											if tk == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1798
			_go_fuzz_dep_.CoverTab[130207]++
												jww.TRACE.Printf("tk=\"\", tgt[%s]=%v", sk, sv)
												tgt[sk] = sv
												if itgt != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1801
				_go_fuzz_dep_.CoverTab[130209]++
													itgt[sk] = sv
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1802
				// _ = "end of CoverTab[130209]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1803
				_go_fuzz_dep_.CoverTab[130210]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1803
				// _ = "end of CoverTab[130210]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1803
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1803
			// _ = "end of CoverTab[130207]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1803
			_go_fuzz_dep_.CoverTab[130208]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1804
			// _ = "end of CoverTab[130208]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1805
			_go_fuzz_dep_.CoverTab[130211]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1805
			// _ = "end of CoverTab[130211]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1805
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1805
		// _ = "end of CoverTab[130203]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1805
		_go_fuzz_dep_.CoverTab[130204]++

											tv, ok := tgt[tk]
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1808
			_go_fuzz_dep_.CoverTab[130212]++
												jww.TRACE.Printf("tgt[%s] != ok, tgt[%s]=%v", tk, sk, sv)
												tgt[sk] = sv
												if itgt != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1811
				_go_fuzz_dep_.CoverTab[130214]++
													itgt[sk] = sv
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1812
				// _ = "end of CoverTab[130214]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1813
				_go_fuzz_dep_.CoverTab[130215]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1813
				// _ = "end of CoverTab[130215]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1813
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1813
			// _ = "end of CoverTab[130212]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1813
			_go_fuzz_dep_.CoverTab[130213]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1814
			// _ = "end of CoverTab[130213]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1815
			_go_fuzz_dep_.CoverTab[130216]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1815
			// _ = "end of CoverTab[130216]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1815
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1815
		// _ = "end of CoverTab[130204]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1815
		_go_fuzz_dep_.CoverTab[130205]++

											svType := reflect.TypeOf(sv)
											tvType := reflect.TypeOf(tv)
											if tvType != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1819
			_go_fuzz_dep_.CoverTab[130217]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1819
			return svType != tvType
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1819
			// _ = "end of CoverTab[130217]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1819
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1819
			_go_fuzz_dep_.CoverTab[130218]++
												jww.ERROR.Printf(
				"svType != tvType; key=%s, st=%v, tt=%v, sv=%v, tv=%v",
				sk, svType, tvType, sv, tv)
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1823
			// _ = "end of CoverTab[130218]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1824
			_go_fuzz_dep_.CoverTab[130219]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1824
			// _ = "end of CoverTab[130219]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1824
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1824
		// _ = "end of CoverTab[130205]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1824
		_go_fuzz_dep_.CoverTab[130206]++

											jww.TRACE.Printf("processing key=%s, st=%v, tt=%v, sv=%v, tv=%v",
			sk, svType, tvType, sv, tv)

		switch ttv := tv.(type) {
		case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1830
			_go_fuzz_dep_.CoverTab[130220]++
												jww.TRACE.Printf("merging maps (must convert)")
												tsv := sv.(map[interface{}]interface{})
												ssv := castToMapStringInterface(tsv)
												stv := castToMapStringInterface(ttv)
												mergeMaps(ssv, stv, ttv)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1835
			// _ = "end of CoverTab[130220]"
		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1836
			_go_fuzz_dep_.CoverTab[130221]++
												jww.TRACE.Printf("merging maps")
												mergeMaps(sv.(map[string]interface{}), ttv, nil)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1838
			// _ = "end of CoverTab[130221]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1839
			_go_fuzz_dep_.CoverTab[130222]++
												jww.TRACE.Printf("setting value")
												tgt[tk] = sv
												if itgt != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1842
				_go_fuzz_dep_.CoverTab[130223]++
													itgt[tk] = sv
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1843
				// _ = "end of CoverTab[130223]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1844
				_go_fuzz_dep_.CoverTab[130224]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1844
				// _ = "end of CoverTab[130224]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1844
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1844
			// _ = "end of CoverTab[130222]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1845
		// _ = "end of CoverTab[130206]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1846
	// _ = "end of CoverTab[130202]"
}

// ReadRemoteConfig attempts to get configuration from a remote source
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1849
// and read it in the remote configuration registry.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1851
func ReadRemoteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1851
	_go_fuzz_dep_.CoverTab[130225]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1851
	return v.ReadRemoteConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1851
	// _ = "end of CoverTab[130225]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1851
}

func (v *Viper) ReadRemoteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1853
	_go_fuzz_dep_.CoverTab[130226]++
										return v.getKeyValueConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1854
	// _ = "end of CoverTab[130226]"
}

func WatchRemoteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1857
	_go_fuzz_dep_.CoverTab[130227]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1857
	return v.WatchRemoteConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1857
	// _ = "end of CoverTab[130227]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1857
}
func (v *Viper) WatchRemoteConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1858
	_go_fuzz_dep_.CoverTab[130228]++
										return v.watchKeyValueConfig()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1859
	// _ = "end of CoverTab[130228]"
}

func (v *Viper) WatchRemoteConfigOnChannel() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1862
	_go_fuzz_dep_.CoverTab[130229]++
										return v.watchKeyValueConfigOnChannel()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1863
	// _ = "end of CoverTab[130229]"
}

// Retrieve the first found remote configuration.
func (v *Viper) getKeyValueConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1867
	_go_fuzz_dep_.CoverTab[130230]++
										if RemoteConfig == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1868
		_go_fuzz_dep_.CoverTab[130233]++
											return RemoteConfigError("Enable the remote features by doing a blank import of the viper/remote package: '_ github.com/spf13/viper/remote'")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1869
		// _ = "end of CoverTab[130233]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1870
		_go_fuzz_dep_.CoverTab[130234]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1870
		// _ = "end of CoverTab[130234]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1870
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1870
	// _ = "end of CoverTab[130230]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1870
	_go_fuzz_dep_.CoverTab[130231]++

										for _, rp := range v.remoteProviders {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1872
		_go_fuzz_dep_.CoverTab[130235]++
											val, err := v.getRemoteConfig(rp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1874
			_go_fuzz_dep_.CoverTab[130237]++
												jww.ERROR.Printf("get remote config: %s", err)

												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1877
			// _ = "end of CoverTab[130237]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1878
			_go_fuzz_dep_.CoverTab[130238]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1878
			// _ = "end of CoverTab[130238]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1878
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1878
		// _ = "end of CoverTab[130235]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1878
		_go_fuzz_dep_.CoverTab[130236]++

											v.kvstore = val

											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1882
		// _ = "end of CoverTab[130236]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1883
	// _ = "end of CoverTab[130231]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1883
	_go_fuzz_dep_.CoverTab[130232]++
										return RemoteConfigError("No Files Found")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1884
	// _ = "end of CoverTab[130232]"
}

func (v *Viper) getRemoteConfig(provider RemoteProvider) (map[string]interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1887
	_go_fuzz_dep_.CoverTab[130239]++
										reader, err := RemoteConfig.Get(provider)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1889
		_go_fuzz_dep_.CoverTab[130241]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1890
		// _ = "end of CoverTab[130241]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1891
		_go_fuzz_dep_.CoverTab[130242]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1891
		// _ = "end of CoverTab[130242]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1891
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1891
	// _ = "end of CoverTab[130239]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1891
	_go_fuzz_dep_.CoverTab[130240]++
										err = v.unmarshalReader(reader, v.kvstore)
										return v.kvstore, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1893
	// _ = "end of CoverTab[130240]"
}

// Retrieve the first found remote configuration.
func (v *Viper) watchKeyValueConfigOnChannel() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1897
	_go_fuzz_dep_.CoverTab[130243]++
										for _, rp := range v.remoteProviders {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1898
		_go_fuzz_dep_.CoverTab[130245]++
											respc, _ := RemoteConfig.WatchChannel(rp)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1899
		_curRoutineNum156_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1899
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum156_)

											go func(rc <-chan *RemoteResponse) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1901
			_go_fuzz_dep_.CoverTab[130247]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1901
			defer func() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1901
				_go_fuzz_dep_.CoverTab[130248]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1901
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum156_)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1901
				// _ = "end of CoverTab[130248]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1901
			}()
												for {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1902
				_go_fuzz_dep_.CoverTab[130249]++
													b := <-rc
													reader := bytes.NewReader(b.Value)
													v.unmarshalReader(reader, v.kvstore)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1905
				// _ = "end of CoverTab[130249]"
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1906
			// _ = "end of CoverTab[130247]"
		}(respc)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1907
		// _ = "end of CoverTab[130245]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1907
		_go_fuzz_dep_.CoverTab[130246]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1908
		// _ = "end of CoverTab[130246]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1909
	// _ = "end of CoverTab[130243]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1909
	_go_fuzz_dep_.CoverTab[130244]++
										return RemoteConfigError("No Files Found")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1910
	// _ = "end of CoverTab[130244]"
}

// Retrieve the first found remote configuration.
func (v *Viper) watchKeyValueConfig() error {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1914
	_go_fuzz_dep_.CoverTab[130250]++
										for _, rp := range v.remoteProviders {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1915
		_go_fuzz_dep_.CoverTab[130252]++
											val, err := v.watchRemoteConfig(rp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1917
			_go_fuzz_dep_.CoverTab[130254]++
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1918
			// _ = "end of CoverTab[130254]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1919
			_go_fuzz_dep_.CoverTab[130255]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1919
			// _ = "end of CoverTab[130255]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1919
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1919
		// _ = "end of CoverTab[130252]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1919
		_go_fuzz_dep_.CoverTab[130253]++
											v.kvstore = val
											return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1921
		// _ = "end of CoverTab[130253]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1922
	// _ = "end of CoverTab[130250]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1922
	_go_fuzz_dep_.CoverTab[130251]++
										return RemoteConfigError("No Files Found")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1923
	// _ = "end of CoverTab[130251]"
}

func (v *Viper) watchRemoteConfig(provider RemoteProvider) (map[string]interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1926
	_go_fuzz_dep_.CoverTab[130256]++
										reader, err := RemoteConfig.Watch(provider)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1928
		_go_fuzz_dep_.CoverTab[130258]++
											return nil, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1929
		// _ = "end of CoverTab[130258]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1930
		_go_fuzz_dep_.CoverTab[130259]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1930
		// _ = "end of CoverTab[130259]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1930
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1930
	// _ = "end of CoverTab[130256]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1930
	_go_fuzz_dep_.CoverTab[130257]++
										err = v.unmarshalReader(reader, v.kvstore)
										return v.kvstore, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1932
	// _ = "end of CoverTab[130257]"
}

// AllKeys returns all keys holding a value, regardless of where they are set.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1935
// Nested keys are returned with a v.keyDelim separator
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1937
func AllKeys() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1937
	_go_fuzz_dep_.CoverTab[130260]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1937
	return v.AllKeys()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1937
	// _ = "end of CoverTab[130260]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1937
}

func (v *Viper) AllKeys() []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1939
	_go_fuzz_dep_.CoverTab[130261]++
										m := map[string]bool{}

										m = v.flattenAndMergeMap(m, castMapStringToMapInterface(v.aliases), "")
										m = v.flattenAndMergeMap(m, v.override, "")
										m = v.mergeFlatMap(m, castMapFlagToMapInterface(v.pflags))
										m = v.mergeFlatMap(m, castMapStringSliceToMapInterface(v.env))
										m = v.flattenAndMergeMap(m, v.config, "")
										m = v.flattenAndMergeMap(m, v.kvstore, "")
										m = v.flattenAndMergeMap(m, v.defaults, "")

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1951
	a := make([]string, 0, len(m))
	for x := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1952
		_go_fuzz_dep_.CoverTab[130263]++
											a = append(a, x)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1953
		// _ = "end of CoverTab[130263]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1954
	// _ = "end of CoverTab[130261]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1954
	_go_fuzz_dep_.CoverTab[130262]++
										return a
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1955
	// _ = "end of CoverTab[130262]"
}

// flattenAndMergeMap recursively flattens the given map into a map[string]bool
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1958
// of key paths (used as a set, easier to manipulate than a []string):
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1958
//   - each path is merged into a single key string, delimited with v.keyDelim
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1958
//   - if a path is shadowed by an earlier value in the initial shadow map,
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1958
//     it is skipped.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1958
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1958
// The resulting set of paths is merged to the given shadow set at the same time.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1964
func (v *Viper) flattenAndMergeMap(shadow map[string]bool, m map[string]interface{}, prefix string) map[string]bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1964
	_go_fuzz_dep_.CoverTab[130264]++
										if shadow != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
		_go_fuzz_dep_.CoverTab[130269]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
		return prefix != ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
		// _ = "end of CoverTab[130269]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
		_go_fuzz_dep_.CoverTab[130270]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
		return shadow[prefix]
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
		// _ = "end of CoverTab[130270]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1965
		_go_fuzz_dep_.CoverTab[130271]++

											return shadow
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1967
		// _ = "end of CoverTab[130271]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1968
		_go_fuzz_dep_.CoverTab[130272]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1968
		// _ = "end of CoverTab[130272]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1968
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1968
	// _ = "end of CoverTab[130264]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1968
	_go_fuzz_dep_.CoverTab[130265]++
										if shadow == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1969
		_go_fuzz_dep_.CoverTab[130273]++
											shadow = make(map[string]bool)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1970
		// _ = "end of CoverTab[130273]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1971
		_go_fuzz_dep_.CoverTab[130274]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1971
		// _ = "end of CoverTab[130274]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1971
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1971
	// _ = "end of CoverTab[130265]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1971
	_go_fuzz_dep_.CoverTab[130266]++

										var m2 map[string]interface{}
										if prefix != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1974
		_go_fuzz_dep_.CoverTab[130275]++
											prefix += v.keyDelim
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1975
		// _ = "end of CoverTab[130275]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1976
		_go_fuzz_dep_.CoverTab[130276]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1976
		// _ = "end of CoverTab[130276]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1976
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1976
	// _ = "end of CoverTab[130266]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1976
	_go_fuzz_dep_.CoverTab[130267]++
										for k, val := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1977
		_go_fuzz_dep_.CoverTab[130277]++
											fullKey := prefix + k
											switch val.(type) {
		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1980
			_go_fuzz_dep_.CoverTab[130279]++
												m2 = val.(map[string]interface{})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1981
			// _ = "end of CoverTab[130279]"
		case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1982
			_go_fuzz_dep_.CoverTab[130280]++
												m2 = cast.ToStringMap(val)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1983
			// _ = "end of CoverTab[130280]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1984
			_go_fuzz_dep_.CoverTab[130281]++

												shadow[strings.ToLower(fullKey)] = true
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1987
			// _ = "end of CoverTab[130281]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1988
		// _ = "end of CoverTab[130277]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1988
		_go_fuzz_dep_.CoverTab[130278]++

											shadow = v.flattenAndMergeMap(shadow, m2, fullKey)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1990
		// _ = "end of CoverTab[130278]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1991
	// _ = "end of CoverTab[130267]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1991
	_go_fuzz_dep_.CoverTab[130268]++
										return shadow
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1992
	// _ = "end of CoverTab[130268]"
}

// mergeFlatMap merges the given maps, excluding values of the second map
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1995
// shadowed by values from the first map.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1997
func (v *Viper) mergeFlatMap(shadow map[string]bool, m map[string]interface{}) map[string]bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:1997
	_go_fuzz_dep_.CoverTab[130282]++

outer:
	for k := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2000
		_go_fuzz_dep_.CoverTab[130284]++
											path := strings.Split(k, v.keyDelim)
		// scan intermediate paths
		var parentKey string
		for i := 1; i < len(path); i++ {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2004
			_go_fuzz_dep_.CoverTab[130286]++
												parentKey = strings.Join(path[0:i], v.keyDelim)
												if shadow[parentKey] {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2006
				_go_fuzz_dep_.CoverTab[130287]++

													continue outer
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2008
				// _ = "end of CoverTab[130287]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2009
				_go_fuzz_dep_.CoverTab[130288]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2009
				// _ = "end of CoverTab[130288]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2009
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2009
			// _ = "end of CoverTab[130286]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2010
		// _ = "end of CoverTab[130284]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2010
		_go_fuzz_dep_.CoverTab[130285]++

											shadow[strings.ToLower(k)] = true
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2012
		// _ = "end of CoverTab[130285]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2013
	// _ = "end of CoverTab[130282]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2013
	_go_fuzz_dep_.CoverTab[130283]++
										return shadow
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2014
	// _ = "end of CoverTab[130283]"
}

// AllSettings merges all settings and returns them as a map[string]interface{}.
func AllSettings() map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2018
	_go_fuzz_dep_.CoverTab[130289]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2018
	return v.AllSettings()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2018
	// _ = "end of CoverTab[130289]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2018
}

func (v *Viper) AllSettings() map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2020
	_go_fuzz_dep_.CoverTab[130290]++
										m := map[string]interface{}{}

										for _, k := range v.AllKeys() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2023
		_go_fuzz_dep_.CoverTab[130292]++
											value := v.Get(k)
											if value == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2025
			_go_fuzz_dep_.CoverTab[130294]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2028
			continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2028
			// _ = "end of CoverTab[130294]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2029
			_go_fuzz_dep_.CoverTab[130295]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2029
			// _ = "end of CoverTab[130295]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2029
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2029
		// _ = "end of CoverTab[130292]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2029
		_go_fuzz_dep_.CoverTab[130293]++
											path := strings.Split(k, v.keyDelim)
											lastKey := strings.ToLower(path[len(path)-1])
											deepestMap := deepSearch(m, path[0:len(path)-1])

											deepestMap[lastKey] = value
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2034
		// _ = "end of CoverTab[130293]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2035
	// _ = "end of CoverTab[130290]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2035
	_go_fuzz_dep_.CoverTab[130291]++
										return m
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2036
	// _ = "end of CoverTab[130291]"
}

// SetFs sets the filesystem to use to read configuration.
func SetFs(fs afero.Fs) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2040
	_go_fuzz_dep_.CoverTab[130296]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2040
	v.SetFs(fs)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2040
	// _ = "end of CoverTab[130296]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2040
}

func (v *Viper) SetFs(fs afero.Fs) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2042
	_go_fuzz_dep_.CoverTab[130297]++
										v.fs = fs
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2043
	// _ = "end of CoverTab[130297]"
}

// SetConfigName sets name for the config file.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2046
// Does not include extension.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2048
func SetConfigName(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2048
	_go_fuzz_dep_.CoverTab[130298]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2048
	v.SetConfigName(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2048
	// _ = "end of CoverTab[130298]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2048
}

func (v *Viper) SetConfigName(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2050
	_go_fuzz_dep_.CoverTab[130299]++
										if in != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2051
		_go_fuzz_dep_.CoverTab[130300]++
											v.configName = in
											v.configFile = ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2053
		// _ = "end of CoverTab[130300]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2054
		_go_fuzz_dep_.CoverTab[130301]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2054
		// _ = "end of CoverTab[130301]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2054
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2054
	// _ = "end of CoverTab[130299]"
}

// SetConfigType sets the type of the configuration returned by the
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2057
// remote source, e.g. "json".
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2059
func SetConfigType(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2059
	_go_fuzz_dep_.CoverTab[130302]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2059
	v.SetConfigType(in)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2059
	// _ = "end of CoverTab[130302]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2059
}

func (v *Viper) SetConfigType(in string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2061
	_go_fuzz_dep_.CoverTab[130303]++
										if in != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2062
		_go_fuzz_dep_.CoverTab[130304]++
											v.configType = in
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2063
		// _ = "end of CoverTab[130304]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2064
		_go_fuzz_dep_.CoverTab[130305]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2064
		// _ = "end of CoverTab[130305]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2064
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2064
	// _ = "end of CoverTab[130303]"
}

// SetConfigPermissions sets the permissions for the config file.
func SetConfigPermissions(perm os.FileMode) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2068
	_go_fuzz_dep_.CoverTab[130306]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2068
	v.SetConfigPermissions(perm)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2068
	// _ = "end of CoverTab[130306]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2068
}

func (v *Viper) SetConfigPermissions(perm os.FileMode) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2070
	_go_fuzz_dep_.CoverTab[130307]++
										v.configPermissions = perm.Perm()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2071
	// _ = "end of CoverTab[130307]"
}

// IniLoadOptions sets the load options for ini parsing.
func IniLoadOptions(in ini.LoadOptions) Option {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2075
	_go_fuzz_dep_.CoverTab[130308]++
										return optionFunc(func(v *Viper) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2076
		_go_fuzz_dep_.CoverTab[130309]++
											v.iniLoadOptions = in
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2077
		// _ = "end of CoverTab[130309]"
	})
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2078
	// _ = "end of CoverTab[130308]"
}

func (v *Viper) getConfigType() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2081
	_go_fuzz_dep_.CoverTab[130310]++
										if v.configType != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2082
		_go_fuzz_dep_.CoverTab[130314]++
											return v.configType
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2083
		// _ = "end of CoverTab[130314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2084
		_go_fuzz_dep_.CoverTab[130315]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2084
		// _ = "end of CoverTab[130315]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2084
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2084
	// _ = "end of CoverTab[130310]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2084
	_go_fuzz_dep_.CoverTab[130311]++

										cf, err := v.getConfigFile()
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2087
		_go_fuzz_dep_.CoverTab[130316]++
											return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2088
		// _ = "end of CoverTab[130316]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2089
		_go_fuzz_dep_.CoverTab[130317]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2089
		// _ = "end of CoverTab[130317]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2089
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2089
	// _ = "end of CoverTab[130311]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2089
	_go_fuzz_dep_.CoverTab[130312]++

										ext := filepath.Ext(cf)

										if len(ext) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2093
		_go_fuzz_dep_.CoverTab[130318]++
											return ext[1:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2094
		// _ = "end of CoverTab[130318]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2095
		_go_fuzz_dep_.CoverTab[130319]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2095
		// _ = "end of CoverTab[130319]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2095
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2095
	// _ = "end of CoverTab[130312]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2095
	_go_fuzz_dep_.CoverTab[130313]++

										return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2097
	// _ = "end of CoverTab[130313]"
}

func (v *Viper) getConfigFile() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2100
	_go_fuzz_dep_.CoverTab[130320]++
										if v.configFile == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2101
		_go_fuzz_dep_.CoverTab[130322]++
											cf, err := v.findConfigFile()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2103
			_go_fuzz_dep_.CoverTab[130324]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2104
			// _ = "end of CoverTab[130324]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2105
			_go_fuzz_dep_.CoverTab[130325]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2105
			// _ = "end of CoverTab[130325]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2105
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2105
		// _ = "end of CoverTab[130322]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2105
		_go_fuzz_dep_.CoverTab[130323]++
											v.configFile = cf
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2106
		// _ = "end of CoverTab[130323]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2107
		_go_fuzz_dep_.CoverTab[130326]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2107
		// _ = "end of CoverTab[130326]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2107
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2107
	// _ = "end of CoverTab[130320]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2107
	_go_fuzz_dep_.CoverTab[130321]++
										return v.configFile, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2108
	// _ = "end of CoverTab[130321]"
}

func (v *Viper) searchInPath(in string) (filename string) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2111
	_go_fuzz_dep_.CoverTab[130327]++
										jww.DEBUG.Println("Searching for config in ", in)
										for _, ext := range SupportedExts {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2113
		_go_fuzz_dep_.CoverTab[130330]++
											jww.DEBUG.Println("Checking for", filepath.Join(in, v.configName+"."+ext))
											if b, _ := exists(v.fs, filepath.Join(in, v.configName+"."+ext)); b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2115
			_go_fuzz_dep_.CoverTab[130331]++
												jww.DEBUG.Println("Found: ", filepath.Join(in, v.configName+"."+ext))
												return filepath.Join(in, v.configName+"."+ext)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2117
			// _ = "end of CoverTab[130331]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2118
			_go_fuzz_dep_.CoverTab[130332]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2118
			// _ = "end of CoverTab[130332]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2118
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2118
		// _ = "end of CoverTab[130330]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2119
	// _ = "end of CoverTab[130327]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2119
	_go_fuzz_dep_.CoverTab[130328]++

										if v.configType != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2121
		_go_fuzz_dep_.CoverTab[130333]++
											if b, _ := exists(v.fs, filepath.Join(in, v.configName)); b {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2122
			_go_fuzz_dep_.CoverTab[130334]++
												return filepath.Join(in, v.configName)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2123
			// _ = "end of CoverTab[130334]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2124
			_go_fuzz_dep_.CoverTab[130335]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2124
			// _ = "end of CoverTab[130335]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2124
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2124
		// _ = "end of CoverTab[130333]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2125
		_go_fuzz_dep_.CoverTab[130336]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2125
		// _ = "end of CoverTab[130336]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2125
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2125
	// _ = "end of CoverTab[130328]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2125
	_go_fuzz_dep_.CoverTab[130329]++

										return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2127
	// _ = "end of CoverTab[130329]"
}

// Search all configPaths for any config file.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2130
// Returns the first path that exists (and is a config file).
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2132
func (v *Viper) findConfigFile() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2132
	_go_fuzz_dep_.CoverTab[130337]++
										jww.INFO.Println("Searching for config in ", v.configPaths)

										for _, cp := range v.configPaths {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2135
		_go_fuzz_dep_.CoverTab[130339]++
											file := v.searchInPath(cp)
											if file != "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2137
			_go_fuzz_dep_.CoverTab[130340]++
												return file, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2138
			// _ = "end of CoverTab[130340]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2139
			_go_fuzz_dep_.CoverTab[130341]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2139
			// _ = "end of CoverTab[130341]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2139
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2139
		// _ = "end of CoverTab[130339]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2140
	// _ = "end of CoverTab[130337]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2140
	_go_fuzz_dep_.CoverTab[130338]++
										return "", ConfigFileNotFoundError{v.configName, fmt.Sprintf("%s", v.configPaths)}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2141
	// _ = "end of CoverTab[130338]"
}

// Debug prints all configuration registries for debugging
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2144
// purposes.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2146
func Debug()	{ _go_fuzz_dep_.CoverTab[130342]++; v.Debug(); // _ = "end of CoverTab[130342]" }

func (v *Viper) Debug() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2148
	_go_fuzz_dep_.CoverTab[130343]++
										fmt.Printf("Aliases:\n%#v\n", v.aliases)
										fmt.Printf("Override:\n%#v\n", v.override)
										fmt.Printf("PFlags:\n%#v\n", v.pflags)
										fmt.Printf("Env:\n%#v\n", v.env)
										fmt.Printf("Key/Value Store:\n%#v\n", v.kvstore)
										fmt.Printf("Config:\n%#v\n", v.config)
										fmt.Printf("Defaults:\n%#v\n", v.defaults)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2155
	// _ = "end of CoverTab[130343]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2156
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/viper.go:2156
var _ = _go_fuzz_dep_.CoverTab
