//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:1
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:1
)

import (
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

type (
	// FormatChecker is the interface all formatters added to FormatCheckerChain must implement
	FormatChecker	interface {
		// IsFormat checks if input has the correct format and type
		IsFormat(input interface{}) bool
	}

	// FormatCheckerChain holds the formatters
	FormatCheckerChain	struct {
		formatters map[string]FormatChecker
	}

	// EmailFormatChecker verifies email address formats
	EmailFormatChecker	struct{}

	// IPV4FormatChecker verifies IP addresses in the IPv4 format
	IPV4FormatChecker	struct{}

	// IPV6FormatChecker verifies IP addresses in the IPv6 format
	IPV6FormatChecker	struct{}

	// DateTimeFormatChecker verifies date/time formats per RFC3339 5.6
	//
	// Valid formats:
	// 		Partial Time: HH:MM:SS
	//		Full Date: YYYY-MM-DD
	// 		Full Time: HH:MM:SSZ-07:00
	//		Date Time: YYYY-MM-DDTHH:MM:SSZ-0700
	//
	// 	Where
	//		YYYY = 4DIGIT year
	//		MM = 2DIGIT month ; 01-12
	//		DD = 2DIGIT day-month ; 01-28, 01-29, 01-30, 01-31 based on month/year
	//		HH = 2DIGIT hour ; 00-23
	//		MM = 2DIGIT ; 00-59
	//		SS = 2DIGIT ; 00-58, 00-60 based on leap second rules
	//		T = Literal
	//		Z = Literal
	//
	//	Note: Nanoseconds are also suported in all formats
	//
	// http://tools.ietf.org/html/rfc3339#section-5.6
	DateTimeFormatChecker	struct{}

	// DateFormatChecker verifies date formats
	//
	// Valid format:
	//		Full Date: YYYY-MM-DD
	//
	// 	Where
	//		YYYY = 4DIGIT year
	//		MM = 2DIGIT month ; 01-12
	//		DD = 2DIGIT day-month ; 01-28, 01-29, 01-30, 01-31 based on month/year
	DateFormatChecker	struct{}

	// TimeFormatChecker verifies time formats
	//
	// Valid formats:
	// 		Partial Time: HH:MM:SS
	// 		Full Time: HH:MM:SSZ-07:00
	//
	// 	Where
	//		HH = 2DIGIT hour ; 00-23
	//		MM = 2DIGIT ; 00-59
	//		SS = 2DIGIT ; 00-58, 00-60 based on leap second rules
	//		T = Literal
	//		Z = Literal
	TimeFormatChecker	struct{}

	// URIFormatChecker validates a URI with a valid Scheme per RFC3986
	URIFormatChecker	struct{}

	// URIReferenceFormatChecker validates a URI or relative-reference per RFC3986
	URIReferenceFormatChecker	struct{}

	// URITemplateFormatChecker validates a URI template per RFC6570
	URITemplateFormatChecker	struct{}

	// HostnameFormatChecker validates a hostname is in the correct format
	HostnameFormatChecker	struct{}

	// UUIDFormatChecker validates a UUID is in the correct format
	UUIDFormatChecker	struct{}

	// RegexFormatChecker validates a regex is in the correct format
	RegexFormatChecker	struct{}

	// JSONPointerFormatChecker validates a JSON Pointer per RFC6901
	JSONPointerFormatChecker	struct{}

	// RelativeJSONPointerFormatChecker validates a relative JSON Pointer is in the correct format
	RelativeJSONPointerFormatChecker	struct{}
)

var (
	// FormatCheckers holds the valid formatters, and is a public variable
	// so library users can add custom formatters
	FormatCheckers	= FormatCheckerChain{
		formatters: map[string]FormatChecker{
			"date":				DateFormatChecker{},
			"time":				TimeFormatChecker{},
			"date-time":			DateTimeFormatChecker{},
			"hostname":			HostnameFormatChecker{},
			"email":			EmailFormatChecker{},
			"idn-email":			EmailFormatChecker{},
			"ipv4":				IPV4FormatChecker{},
			"ipv6":				IPV6FormatChecker{},
			"uri":				URIFormatChecker{},
			"uri-reference":		URIReferenceFormatChecker{},
			"iri":				URIFormatChecker{},
			"iri-reference":		URIReferenceFormatChecker{},
			"uri-template":			URITemplateFormatChecker{},
			"uuid":				UUIDFormatChecker{},
			"regex":			RegexFormatChecker{},
			"json-pointer":			JSONPointerFormatChecker{},
			"relative-json-pointer":	RelativeJSONPointerFormatChecker{},
		},
	}

	// Regex credit: https://www.socketloop.com/tutorials/golang-validate-hostname
	rxHostname	= regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`)

	// Use a regex to make sure curly brackets are balanced properly after validating it as a AURI
	rxURITemplate	= regexp.MustCompile("^([^{]*({[^}]*})?)*$")

	rxUUID	= regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$")

	rxJSONPointer	= regexp.MustCompile("^(?:/(?:[^~/]|~0|~1)*)*$")

	rxRelJSONPointer	= regexp.MustCompile("^(?:0|[1-9][0-9]*)(?:#|(?:/(?:[^~/]|~0|~1)*)*)$")

	lock	= new(sync.RWMutex)
)

// Add adds a FormatChecker to the FormatCheckerChain
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:147
// The name used will be the value used for the format key in your json schema
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:149
func (c *FormatCheckerChain) Add(name string, f FormatChecker) *FormatCheckerChain {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:149
	_go_fuzz_dep_.CoverTab[194826]++
													lock.Lock()
													c.formatters[name] = f
													lock.Unlock()

													return c
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:154
	// _ = "end of CoverTab[194826]"
}

// Remove deletes a FormatChecker from the FormatCheckerChain (if it exists)
func (c *FormatCheckerChain) Remove(name string) *FormatCheckerChain {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:158
	_go_fuzz_dep_.CoverTab[194827]++
													lock.Lock()
													delete(c.formatters, name)
													lock.Unlock()

													return c
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:163
	// _ = "end of CoverTab[194827]"
}

// Has checks to see if the FormatCheckerChain holds a FormatChecker with the given name
func (c *FormatCheckerChain) Has(name string) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:167
	_go_fuzz_dep_.CoverTab[194828]++
													lock.RLock()
													_, ok := c.formatters[name]
													lock.RUnlock()

													return ok
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:172
	// _ = "end of CoverTab[194828]"
}

// IsFormat will check an input against a FormatChecker with the given name
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:175
// to see if it is the correct format
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:177
func (c *FormatCheckerChain) IsFormat(name string, input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:177
	_go_fuzz_dep_.CoverTab[194829]++
													lock.RLock()
													f, ok := c.formatters[name]
													lock.RUnlock()

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:183
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:183
		_go_fuzz_dep_.CoverTab[194831]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:184
		// _ = "end of CoverTab[194831]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:185
		_go_fuzz_dep_.CoverTab[194832]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:185
		// _ = "end of CoverTab[194832]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:185
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:185
	// _ = "end of CoverTab[194829]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:185
	_go_fuzz_dep_.CoverTab[194830]++

													return f.IsFormat(input)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:187
	// _ = "end of CoverTab[194830]"
}

// IsFormat checks if input is a correctly formatted e-mail address
func (f EmailFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:191
	_go_fuzz_dep_.CoverTab[194833]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:193
		_go_fuzz_dep_.CoverTab[194835]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:194
		// _ = "end of CoverTab[194835]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:195
		_go_fuzz_dep_.CoverTab[194836]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:195
		// _ = "end of CoverTab[194836]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:195
	// _ = "end of CoverTab[194833]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:195
	_go_fuzz_dep_.CoverTab[194834]++

													_, err := mail.ParseAddress(asString)
													return err == nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:198
	// _ = "end of CoverTab[194834]"
}

// IsFormat checks if input is a correctly formatted IPv4-address
func (f IPV4FormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:202
	_go_fuzz_dep_.CoverTab[194837]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:204
		_go_fuzz_dep_.CoverTab[194839]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:205
		// _ = "end of CoverTab[194839]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:206
		_go_fuzz_dep_.CoverTab[194840]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:206
		// _ = "end of CoverTab[194840]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:206
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:206
	// _ = "end of CoverTab[194837]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:206
	_go_fuzz_dep_.CoverTab[194838]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:209
	ip := net.ParseIP(asString)
	return ip != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:210
		_go_fuzz_dep_.CoverTab[194841]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:210
		return strings.Contains(asString, ".")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:210
		// _ = "end of CoverTab[194841]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:210
	}()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:210
	// _ = "end of CoverTab[194838]"
}

// IsFormat checks if input is a correctly formatted IPv6=address
func (f IPV6FormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:214
	_go_fuzz_dep_.CoverTab[194842]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:216
		_go_fuzz_dep_.CoverTab[194844]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:217
		// _ = "end of CoverTab[194844]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:218
		_go_fuzz_dep_.CoverTab[194845]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:218
		// _ = "end of CoverTab[194845]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:218
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:218
	// _ = "end of CoverTab[194842]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:218
	_go_fuzz_dep_.CoverTab[194843]++

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:221
	ip := net.ParseIP(asString)
	return ip != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:222
		_go_fuzz_dep_.CoverTab[194846]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:222
		return strings.Contains(asString, ":")
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:222
		// _ = "end of CoverTab[194846]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:222
	}()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:222
	// _ = "end of CoverTab[194843]"
}

// IsFormat checks if input is a correctly formatted  date/time per RFC3339 5.6
func (f DateTimeFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:226
	_go_fuzz_dep_.CoverTab[194847]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:228
		_go_fuzz_dep_.CoverTab[194850]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:229
		// _ = "end of CoverTab[194850]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:230
		_go_fuzz_dep_.CoverTab[194851]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:230
		// _ = "end of CoverTab[194851]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:230
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:230
	// _ = "end of CoverTab[194847]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:230
	_go_fuzz_dep_.CoverTab[194848]++

													formats := []string{
		"15:04:05",
		"15:04:05Z07:00",
		"2006-01-02",
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, format := range formats {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:240
		_go_fuzz_dep_.CoverTab[194852]++
														if _, err := time.Parse(format, asString); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:241
			_go_fuzz_dep_.CoverTab[194853]++
															return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:242
			// _ = "end of CoverTab[194853]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:243
			_go_fuzz_dep_.CoverTab[194854]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:243
			// _ = "end of CoverTab[194854]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:243
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:243
		// _ = "end of CoverTab[194852]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:244
	// _ = "end of CoverTab[194848]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:244
	_go_fuzz_dep_.CoverTab[194849]++

													return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:246
	// _ = "end of CoverTab[194849]"
}

// IsFormat checks if input is a correctly formatted  date (YYYY-MM-DD)
func (f DateFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:250
	_go_fuzz_dep_.CoverTab[194855]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:252
		_go_fuzz_dep_.CoverTab[194857]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:253
		// _ = "end of CoverTab[194857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:254
		_go_fuzz_dep_.CoverTab[194858]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:254
		// _ = "end of CoverTab[194858]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:254
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:254
	// _ = "end of CoverTab[194855]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:254
	_go_fuzz_dep_.CoverTab[194856]++
													_, err := time.Parse("2006-01-02", asString)
													return err == nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:256
	// _ = "end of CoverTab[194856]"
}

// IsFormat checks if input correctly formatted time (HH:MM:SS or HH:MM:SSZ-07:00)
func (f TimeFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:260
	_go_fuzz_dep_.CoverTab[194859]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:262
		_go_fuzz_dep_.CoverTab[194862]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:263
		// _ = "end of CoverTab[194862]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:264
		_go_fuzz_dep_.CoverTab[194863]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:264
		// _ = "end of CoverTab[194863]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:264
	// _ = "end of CoverTab[194859]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:264
	_go_fuzz_dep_.CoverTab[194860]++

													if _, err := time.Parse("15:04:05Z07:00", asString); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:266
		_go_fuzz_dep_.CoverTab[194864]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:267
		// _ = "end of CoverTab[194864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:268
		_go_fuzz_dep_.CoverTab[194865]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:268
		// _ = "end of CoverTab[194865]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:268
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:268
	// _ = "end of CoverTab[194860]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:268
	_go_fuzz_dep_.CoverTab[194861]++

													_, err := time.Parse("15:04:05", asString)
													return err == nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:271
	// _ = "end of CoverTab[194861]"
}

// IsFormat checks if input is correctly formatted  URI with a valid Scheme per RFC3986
func (f URIFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:275
	_go_fuzz_dep_.CoverTab[194866]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:277
		_go_fuzz_dep_.CoverTab[194869]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:278
		// _ = "end of CoverTab[194869]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:279
		_go_fuzz_dep_.CoverTab[194870]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:279
		// _ = "end of CoverTab[194870]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:279
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:279
	// _ = "end of CoverTab[194866]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:279
	_go_fuzz_dep_.CoverTab[194867]++

													u, err := url.Parse(asString)

													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:283
		_go_fuzz_dep_.CoverTab[194871]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:283
		return u.Scheme == ""
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:283
		// _ = "end of CoverTab[194871]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:283
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:283
		_go_fuzz_dep_.CoverTab[194872]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:284
		// _ = "end of CoverTab[194872]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:285
		_go_fuzz_dep_.CoverTab[194873]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:285
		// _ = "end of CoverTab[194873]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:285
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:285
	// _ = "end of CoverTab[194867]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:285
	_go_fuzz_dep_.CoverTab[194868]++

													return !strings.Contains(asString, `\`)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:287
	// _ = "end of CoverTab[194868]"
}

// IsFormat checks if input is a correctly formatted URI or relative-reference per RFC3986
func (f URIReferenceFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:291
	_go_fuzz_dep_.CoverTab[194874]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:293
		_go_fuzz_dep_.CoverTab[194876]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:294
		// _ = "end of CoverTab[194876]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:295
		_go_fuzz_dep_.CoverTab[194877]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:295
		// _ = "end of CoverTab[194877]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:295
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:295
	// _ = "end of CoverTab[194874]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:295
	_go_fuzz_dep_.CoverTab[194875]++

													_, err := url.Parse(asString)
													return err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:298
		_go_fuzz_dep_.CoverTab[194878]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:298
		return !strings.Contains(asString, `\`)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:298
		// _ = "end of CoverTab[194878]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:298
	}()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:298
	// _ = "end of CoverTab[194875]"
}

// IsFormat checks if input is a correctly formatted URI template per RFC6570
func (f URITemplateFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:302
	_go_fuzz_dep_.CoverTab[194879]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:304
		_go_fuzz_dep_.CoverTab[194882]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:305
		// _ = "end of CoverTab[194882]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:306
		_go_fuzz_dep_.CoverTab[194883]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:306
		// _ = "end of CoverTab[194883]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:306
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:306
	// _ = "end of CoverTab[194879]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:306
	_go_fuzz_dep_.CoverTab[194880]++

													u, err := url.Parse(asString)
													if err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:309
		_go_fuzz_dep_.CoverTab[194884]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:309
		return strings.Contains(asString, `\`)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:309
		// _ = "end of CoverTab[194884]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:309
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:309
		_go_fuzz_dep_.CoverTab[194885]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:310
		// _ = "end of CoverTab[194885]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:311
		_go_fuzz_dep_.CoverTab[194886]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:311
		// _ = "end of CoverTab[194886]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:311
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:311
	// _ = "end of CoverTab[194880]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:311
	_go_fuzz_dep_.CoverTab[194881]++

													return rxURITemplate.MatchString(u.Path)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:313
	// _ = "end of CoverTab[194881]"
}

// IsFormat checks if input is a correctly formatted hostname
func (f HostnameFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:317
	_go_fuzz_dep_.CoverTab[194887]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:319
		_go_fuzz_dep_.CoverTab[194889]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:320
		// _ = "end of CoverTab[194889]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:321
		_go_fuzz_dep_.CoverTab[194890]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:321
		// _ = "end of CoverTab[194890]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:321
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:321
	// _ = "end of CoverTab[194887]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:321
	_go_fuzz_dep_.CoverTab[194888]++

													return rxHostname.MatchString(asString) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:323
		_go_fuzz_dep_.CoverTab[194891]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:323
		return len(asString) < 256
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:323
		// _ = "end of CoverTab[194891]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:323
	}()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:323
	// _ = "end of CoverTab[194888]"
}

// IsFormat checks if input is a correctly formatted UUID
func (f UUIDFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:327
	_go_fuzz_dep_.CoverTab[194892]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:329
		_go_fuzz_dep_.CoverTab[194894]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:330
		// _ = "end of CoverTab[194894]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:331
		_go_fuzz_dep_.CoverTab[194895]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:331
		// _ = "end of CoverTab[194895]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:331
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:331
	// _ = "end of CoverTab[194892]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:331
	_go_fuzz_dep_.CoverTab[194893]++

													return rxUUID.MatchString(asString)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:333
	// _ = "end of CoverTab[194893]"
}

// IsFormat checks if input is a correctly formatted regular expression
func (f RegexFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:337
	_go_fuzz_dep_.CoverTab[194896]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:339
		_go_fuzz_dep_.CoverTab[194899]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:340
		// _ = "end of CoverTab[194899]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:341
		_go_fuzz_dep_.CoverTab[194900]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:341
		// _ = "end of CoverTab[194900]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:341
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:341
	// _ = "end of CoverTab[194896]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:341
	_go_fuzz_dep_.CoverTab[194897]++

													if asString == "" {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:343
		_go_fuzz_dep_.CoverTab[194901]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:344
		// _ = "end of CoverTab[194901]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:345
		_go_fuzz_dep_.CoverTab[194902]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:345
		// _ = "end of CoverTab[194902]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:345
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:345
	// _ = "end of CoverTab[194897]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:345
	_go_fuzz_dep_.CoverTab[194898]++
													_, err := regexp.Compile(asString)
													return err == nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:347
	// _ = "end of CoverTab[194898]"
}

// IsFormat checks if input is a correctly formatted JSON Pointer per RFC6901
func (f JSONPointerFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:351
	_go_fuzz_dep_.CoverTab[194903]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:353
		_go_fuzz_dep_.CoverTab[194905]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:354
		// _ = "end of CoverTab[194905]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:355
		_go_fuzz_dep_.CoverTab[194906]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:355
		// _ = "end of CoverTab[194906]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:355
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:355
	// _ = "end of CoverTab[194903]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:355
	_go_fuzz_dep_.CoverTab[194904]++

													return rxJSONPointer.MatchString(asString)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:357
	// _ = "end of CoverTab[194904]"
}

// IsFormat checks if input is a correctly formatted relative JSON Pointer
func (f RelativeJSONPointerFormatChecker) IsFormat(input interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:361
	_go_fuzz_dep_.CoverTab[194907]++
													asString, ok := input.(string)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:363
		_go_fuzz_dep_.CoverTab[194909]++
														return false
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:364
		// _ = "end of CoverTab[194909]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:365
		_go_fuzz_dep_.CoverTab[194910]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:365
		// _ = "end of CoverTab[194910]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:365
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:365
	// _ = "end of CoverTab[194907]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:365
	_go_fuzz_dep_.CoverTab[194908]++

													return rxRelJSONPointer.MatchString(asString)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:367
	// _ = "end of CoverTab[194908]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:368
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/format_checkers.go:368
var _ = _go_fuzz_dep_.CoverTab
