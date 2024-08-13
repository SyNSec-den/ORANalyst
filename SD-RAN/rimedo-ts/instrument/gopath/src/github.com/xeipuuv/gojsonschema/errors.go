//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:1
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:1
)

import (
	"bytes"
	"sync"
	"text/template"
)

var errorTemplates = errorTemplate{template.New("errors-new"), sync.RWMutex{}}

// template.Template is not thread-safe for writing, so some locking is done
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:11
// sync.RWMutex is used for efficiently locking when new templates are created
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:13
type errorTemplate struct {
	*template.Template
	sync.RWMutex
}

type (

	// FalseError. ErrorDetails: -
	FalseError	struct {
		ResultErrorFields
	}

	// RequiredError indicates that a required field is missing
	// ErrorDetails: property string
	RequiredError	struct {
		ResultErrorFields
	}

	// InvalidTypeError indicates that a field has the incorrect type
	// ErrorDetails: expected, given
	InvalidTypeError	struct {
		ResultErrorFields
	}

	// NumberAnyOfError is produced in case of a failing "anyOf" validation
	// ErrorDetails: -
	NumberAnyOfError	struct {
		ResultErrorFields
	}

	// NumberOneOfError is produced in case of a failing "oneOf" validation
	// ErrorDetails: -
	NumberOneOfError	struct {
		ResultErrorFields
	}

	// NumberAllOfError is produced in case of a failing "allOf" validation
	// ErrorDetails: -
	NumberAllOfError	struct {
		ResultErrorFields
	}

	// NumberNotError is produced if a "not" validation failed
	// ErrorDetails: -
	NumberNotError	struct {
		ResultErrorFields
	}

	// MissingDependencyError is produced in case of a "missing dependency" problem
	// ErrorDetails: dependency
	MissingDependencyError	struct {
		ResultErrorFields
	}

	// InternalError indicates an internal error
	// ErrorDetails: error
	InternalError	struct {
		ResultErrorFields
	}

	// ConstError indicates a const error
	// ErrorDetails: allowed
	ConstError	struct {
		ResultErrorFields
	}

	// EnumError indicates an enum error
	// ErrorDetails: allowed
	EnumError	struct {
		ResultErrorFields
	}

	// ArrayNoAdditionalItemsError is produced if additional items were found, but not allowed
	// ErrorDetails: -
	ArrayNoAdditionalItemsError	struct {
		ResultErrorFields
	}

	// ArrayMinItemsError is produced if an array contains less items than the allowed minimum
	// ErrorDetails: min
	ArrayMinItemsError	struct {
		ResultErrorFields
	}

	// ArrayMaxItemsError is produced if an array contains more items than the allowed maximum
	// ErrorDetails: max
	ArrayMaxItemsError	struct {
		ResultErrorFields
	}

	// ItemsMustBeUniqueError is produced if an array requires unique items, but contains non-unique items
	// ErrorDetails: type, i, j
	ItemsMustBeUniqueError	struct {
		ResultErrorFields
	}

	// ArrayContainsError is produced if an array contains invalid items
	// ErrorDetails:
	ArrayContainsError	struct {
		ResultErrorFields
	}

	// ArrayMinPropertiesError is produced if an object contains less properties than the allowed minimum
	// ErrorDetails: min
	ArrayMinPropertiesError	struct {
		ResultErrorFields
	}

	// ArrayMaxPropertiesError is produced if an object contains more properties than the allowed maximum
	// ErrorDetails: max
	ArrayMaxPropertiesError	struct {
		ResultErrorFields
	}

	// AdditionalPropertyNotAllowedError is produced if an object has additional properties, but not allowed
	// ErrorDetails: property
	AdditionalPropertyNotAllowedError	struct {
		ResultErrorFields
	}

	// InvalidPropertyPatternError is produced if an pattern was found
	// ErrorDetails: property, pattern
	InvalidPropertyPatternError	struct {
		ResultErrorFields
	}

	// InvalidPropertyNameError is produced if an invalid-named property was found
	// ErrorDetails: property
	InvalidPropertyNameError	struct {
		ResultErrorFields
	}

	// StringLengthGTEError is produced if a string is shorter than the minimum required length
	// ErrorDetails: min
	StringLengthGTEError	struct {
		ResultErrorFields
	}

	// StringLengthLTEError is produced if a string is longer than the maximum allowed length
	// ErrorDetails: max
	StringLengthLTEError	struct {
		ResultErrorFields
	}

	// DoesNotMatchPatternError is produced if a string does not match the defined pattern
	// ErrorDetails: pattern
	DoesNotMatchPatternError	struct {
		ResultErrorFields
	}

	// DoesNotMatchFormatError is produced if a string does not match the defined format
	// ErrorDetails: format
	DoesNotMatchFormatError	struct {
		ResultErrorFields
	}

	// MultipleOfError is produced if a number is not a multiple of the defined multipleOf
	// ErrorDetails: multiple
	MultipleOfError	struct {
		ResultErrorFields
	}

	// NumberGTEError is produced if a number is lower than the allowed minimum
	// ErrorDetails: min
	NumberGTEError	struct {
		ResultErrorFields
	}

	// NumberGTError is produced if a number is lower than, or equal to the specified minimum, and exclusiveMinimum is set
	// ErrorDetails: min
	NumberGTError	struct {
		ResultErrorFields
	}

	// NumberLTEError is produced if a number is higher than the allowed maximum
	// ErrorDetails: max
	NumberLTEError	struct {
		ResultErrorFields
	}

	// NumberLTError is produced if a number is higher than, or equal to the specified maximum, and exclusiveMaximum is set
	// ErrorDetails: max
	NumberLTError	struct {
		ResultErrorFields
	}

	// ConditionThenError is produced if a condition's "then" validation is invalid
	// ErrorDetails: -
	ConditionThenError	struct {
		ResultErrorFields
	}

	// ConditionElseError is produced if a condition's "else" condition is invalid
	// ErrorDetails: -
	ConditionElseError	struct {
		ResultErrorFields
	}
)

// newError takes a ResultError type and sets the type, context, description, details, value, and field
func newError(err ResultError, context *JsonContext, value interface{}, locale locale, details ErrorDetails) {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:213
	_go_fuzz_dep_.CoverTab[194776]++
											var t string
											var d string
											switch err.(type) {
	case *FalseError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:217
		_go_fuzz_dep_.CoverTab[194779]++
												t = "false"
												d = locale.False()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:219
		// _ = "end of CoverTab[194779]"
	case *RequiredError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:220
		_go_fuzz_dep_.CoverTab[194780]++
												t = "required"
												d = locale.Required()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:222
		// _ = "end of CoverTab[194780]"
	case *InvalidTypeError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:223
		_go_fuzz_dep_.CoverTab[194781]++
												t = "invalid_type"
												d = locale.InvalidType()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:225
		// _ = "end of CoverTab[194781]"
	case *NumberAnyOfError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:226
		_go_fuzz_dep_.CoverTab[194782]++
												t = "number_any_of"
												d = locale.NumberAnyOf()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:228
		// _ = "end of CoverTab[194782]"
	case *NumberOneOfError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:229
		_go_fuzz_dep_.CoverTab[194783]++
												t = "number_one_of"
												d = locale.NumberOneOf()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:231
		// _ = "end of CoverTab[194783]"
	case *NumberAllOfError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:232
		_go_fuzz_dep_.CoverTab[194784]++
												t = "number_all_of"
												d = locale.NumberAllOf()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:234
		// _ = "end of CoverTab[194784]"
	case *NumberNotError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:235
		_go_fuzz_dep_.CoverTab[194785]++
												t = "number_not"
												d = locale.NumberNot()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:237
		// _ = "end of CoverTab[194785]"
	case *MissingDependencyError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:238
		_go_fuzz_dep_.CoverTab[194786]++
												t = "missing_dependency"
												d = locale.MissingDependency()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:240
		// _ = "end of CoverTab[194786]"
	case *InternalError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:241
		_go_fuzz_dep_.CoverTab[194787]++
												t = "internal"
												d = locale.Internal()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:243
		// _ = "end of CoverTab[194787]"
	case *ConstError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:244
		_go_fuzz_dep_.CoverTab[194788]++
												t = "const"
												d = locale.Const()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:246
		// _ = "end of CoverTab[194788]"
	case *EnumError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:247
		_go_fuzz_dep_.CoverTab[194789]++
												t = "enum"
												d = locale.Enum()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:249
		// _ = "end of CoverTab[194789]"
	case *ArrayNoAdditionalItemsError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:250
		_go_fuzz_dep_.CoverTab[194790]++
												t = "array_no_additional_items"
												d = locale.ArrayNoAdditionalItems()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:252
		// _ = "end of CoverTab[194790]"
	case *ArrayMinItemsError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:253
		_go_fuzz_dep_.CoverTab[194791]++
												t = "array_min_items"
												d = locale.ArrayMinItems()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:255
		// _ = "end of CoverTab[194791]"
	case *ArrayMaxItemsError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:256
		_go_fuzz_dep_.CoverTab[194792]++
												t = "array_max_items"
												d = locale.ArrayMaxItems()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:258
		// _ = "end of CoverTab[194792]"
	case *ItemsMustBeUniqueError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:259
		_go_fuzz_dep_.CoverTab[194793]++
												t = "unique"
												d = locale.Unique()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:261
		// _ = "end of CoverTab[194793]"
	case *ArrayContainsError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:262
		_go_fuzz_dep_.CoverTab[194794]++
												t = "contains"
												d = locale.ArrayContains()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:264
		// _ = "end of CoverTab[194794]"
	case *ArrayMinPropertiesError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:265
		_go_fuzz_dep_.CoverTab[194795]++
												t = "array_min_properties"
												d = locale.ArrayMinProperties()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:267
		// _ = "end of CoverTab[194795]"
	case *ArrayMaxPropertiesError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:268
		_go_fuzz_dep_.CoverTab[194796]++
												t = "array_max_properties"
												d = locale.ArrayMaxProperties()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:270
		// _ = "end of CoverTab[194796]"
	case *AdditionalPropertyNotAllowedError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:271
		_go_fuzz_dep_.CoverTab[194797]++
												t = "additional_property_not_allowed"
												d = locale.AdditionalPropertyNotAllowed()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:273
		// _ = "end of CoverTab[194797]"
	case *InvalidPropertyPatternError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:274
		_go_fuzz_dep_.CoverTab[194798]++
												t = "invalid_property_pattern"
												d = locale.InvalidPropertyPattern()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:276
		// _ = "end of CoverTab[194798]"
	case *InvalidPropertyNameError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:277
		_go_fuzz_dep_.CoverTab[194799]++
												t = "invalid_property_name"
												d = locale.InvalidPropertyName()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:279
		// _ = "end of CoverTab[194799]"
	case *StringLengthGTEError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:280
		_go_fuzz_dep_.CoverTab[194800]++
												t = "string_gte"
												d = locale.StringGTE()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:282
		// _ = "end of CoverTab[194800]"
	case *StringLengthLTEError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:283
		_go_fuzz_dep_.CoverTab[194801]++
												t = "string_lte"
												d = locale.StringLTE()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:285
		// _ = "end of CoverTab[194801]"
	case *DoesNotMatchPatternError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:286
		_go_fuzz_dep_.CoverTab[194802]++
												t = "pattern"
												d = locale.DoesNotMatchPattern()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:288
		// _ = "end of CoverTab[194802]"
	case *DoesNotMatchFormatError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:289
		_go_fuzz_dep_.CoverTab[194803]++
												t = "format"
												d = locale.DoesNotMatchFormat()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:291
		// _ = "end of CoverTab[194803]"
	case *MultipleOfError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:292
		_go_fuzz_dep_.CoverTab[194804]++
												t = "multiple_of"
												d = locale.MultipleOf()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:294
		// _ = "end of CoverTab[194804]"
	case *NumberGTEError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:295
		_go_fuzz_dep_.CoverTab[194805]++
												t = "number_gte"
												d = locale.NumberGTE()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:297
		// _ = "end of CoverTab[194805]"
	case *NumberGTError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:298
		_go_fuzz_dep_.CoverTab[194806]++
												t = "number_gt"
												d = locale.NumberGT()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:300
		// _ = "end of CoverTab[194806]"
	case *NumberLTEError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:301
		_go_fuzz_dep_.CoverTab[194807]++
												t = "number_lte"
												d = locale.NumberLTE()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:303
		// _ = "end of CoverTab[194807]"
	case *NumberLTError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:304
		_go_fuzz_dep_.CoverTab[194808]++
												t = "number_lt"
												d = locale.NumberLT()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:306
		// _ = "end of CoverTab[194808]"
	case *ConditionThenError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:307
		_go_fuzz_dep_.CoverTab[194809]++
												t = "condition_then"
												d = locale.ConditionThen()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:309
		// _ = "end of CoverTab[194809]"
	case *ConditionElseError:
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:310
		_go_fuzz_dep_.CoverTab[194810]++
												t = "condition_else"
												d = locale.ConditionElse()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:312
		// _ = "end of CoverTab[194810]"
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:313
	// _ = "end of CoverTab[194776]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:313
	_go_fuzz_dep_.CoverTab[194777]++

											err.SetType(t)
											err.SetContext(context)
											err.SetValue(value)
											err.SetDetails(details)
											err.SetDescriptionFormat(d)
											details["field"] = err.Field()

											if _, exists := details["context"]; !exists && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:322
		_go_fuzz_dep_.CoverTab[194811]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:322
		return context != nil
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:322
		// _ = "end of CoverTab[194811]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:322
	}() {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:322
		_go_fuzz_dep_.CoverTab[194812]++
												details["context"] = context.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:323
		// _ = "end of CoverTab[194812]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:324
		_go_fuzz_dep_.CoverTab[194813]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:324
		// _ = "end of CoverTab[194813]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:324
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:324
	// _ = "end of CoverTab[194777]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:324
	_go_fuzz_dep_.CoverTab[194778]++

											err.SetDescription(formatErrorDescription(err.DescriptionFormat(), details))
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:326
	// _ = "end of CoverTab[194778]"
}

// formatErrorDescription takes a string in the default text/template
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:329
// format and converts it to a string with replacements. The fields come
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:329
// from the ErrorDetails struct and vary for each type of error.
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:332
func formatErrorDescription(s string, details ErrorDetails) string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:332
	_go_fuzz_dep_.CoverTab[194814]++

											var tpl *template.Template
											var descrAsBuffer bytes.Buffer
											var err error

											errorTemplates.RLock()
											tpl = errorTemplates.Lookup(s)
											errorTemplates.RUnlock()

											if tpl == nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:342
		_go_fuzz_dep_.CoverTab[194817]++
												errorTemplates.Lock()
												tpl = errorTemplates.New(s)

												if ErrorTemplateFuncs != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:346
			_go_fuzz_dep_.CoverTab[194819]++
													tpl.Funcs(ErrorTemplateFuncs)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:347
			// _ = "end of CoverTab[194819]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:348
			_go_fuzz_dep_.CoverTab[194820]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:348
			// _ = "end of CoverTab[194820]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:348
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:348
		// _ = "end of CoverTab[194817]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:348
		_go_fuzz_dep_.CoverTab[194818]++

												tpl, err = tpl.Parse(s)
												errorTemplates.Unlock()

												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:353
			_go_fuzz_dep_.CoverTab[194821]++
													return err.Error()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:354
			// _ = "end of CoverTab[194821]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:355
			_go_fuzz_dep_.CoverTab[194822]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:355
			// _ = "end of CoverTab[194822]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:355
		}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:355
		// _ = "end of CoverTab[194818]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:356
		_go_fuzz_dep_.CoverTab[194823]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:356
		// _ = "end of CoverTab[194823]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:356
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:356
	// _ = "end of CoverTab[194814]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:356
	_go_fuzz_dep_.CoverTab[194815]++

											err = tpl.Execute(&descrAsBuffer, details)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:359
		_go_fuzz_dep_.CoverTab[194824]++
												return err.Error()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:360
		// _ = "end of CoverTab[194824]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:361
		_go_fuzz_dep_.CoverTab[194825]++
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:361
		// _ = "end of CoverTab[194825]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:361
	}
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:361
	// _ = "end of CoverTab[194815]"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:361
	_go_fuzz_dep_.CoverTab[194816]++

											return descrAsBuffer.String()
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:363
	// _ = "end of CoverTab[194816]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:364
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/errors.go:364
var _ = _go_fuzz_dep_.CoverTab
