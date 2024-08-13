// Copyright 2015 xeipuuv ( https://github.com/xeipuuv )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author           xeipuuv
// author-github    https://github.com/xeipuuv
// author-mail      xeipuuv@gmail.com
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description      Contains const string and messages.
//
// created          01-01-2015

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:26
package gojsonschema

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:26
)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:26
import (
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:26
)

type (
	// locale is an interface for defining custom error strings
	locale	interface {

		// False returns a format-string for "false" schema validation errors
		False() string

		// Required returns a format-string for "required" schema validation errors
		Required() string

		// InvalidType returns a format-string for "invalid type" schema validation errors
		InvalidType() string

		// NumberAnyOf returns a format-string for "anyOf" schema validation errors
		NumberAnyOf() string

		// NumberOneOf returns a format-string for "oneOf" schema validation errors
		NumberOneOf() string

		// NumberAllOf returns a format-string for "allOf" schema validation errors
		NumberAllOf() string

		// NumberNot returns a format-string to format a NumberNotError
		NumberNot() string

		// MissingDependency returns a format-string for "missing dependency" schema validation errors
		MissingDependency() string

		// Internal returns a format-string for internal errors
		Internal() string

		// Const returns a format-string to format a ConstError
		Const() string

		// Enum returns a format-string to format an EnumError
		Enum() string

		// ArrayNotEnoughItems returns a format-string to format an error for arrays having not enough items to match positional list of schema
		ArrayNotEnoughItems() string

		// ArrayNoAdditionalItems returns a format-string to format an ArrayNoAdditionalItemsError
		ArrayNoAdditionalItems() string

		// ArrayMinItems returns a format-string to format an ArrayMinItemsError
		ArrayMinItems() string

		// ArrayMaxItems returns a format-string to format an ArrayMaxItemsError
		ArrayMaxItems() string

		// Unique returns a format-string  to format an ItemsMustBeUniqueError
		Unique() string

		// ArrayContains returns a format-string to format an ArrayContainsError
		ArrayContains() string

		// ArrayMinProperties returns a format-string to format an ArrayMinPropertiesError
		ArrayMinProperties() string

		// ArrayMaxProperties returns a format-string to format an ArrayMaxPropertiesError
		ArrayMaxProperties() string

		// AdditionalPropertyNotAllowed returns a format-string to format an AdditionalPropertyNotAllowedError
		AdditionalPropertyNotAllowed() string

		// InvalidPropertyPattern returns a format-string to format an InvalidPropertyPatternError
		InvalidPropertyPattern() string

		// InvalidPropertyName returns a format-string to format an InvalidPropertyNameError
		InvalidPropertyName() string

		// StringGTE returns a format-string to format an StringLengthGTEError
		StringGTE() string

		// StringLTE returns a format-string to format an StringLengthLTEError
		StringLTE() string

		// DoesNotMatchPattern returns a format-string to format an DoesNotMatchPatternError
		DoesNotMatchPattern() string

		// DoesNotMatchFormat returns a format-string to format an DoesNotMatchFormatError
		DoesNotMatchFormat() string

		// MultipleOf returns a format-string to format an MultipleOfError
		MultipleOf() string

		// NumberGTE returns a format-string to format an NumberGTEError
		NumberGTE() string

		// NumberGT returns a format-string to format an NumberGTError
		NumberGT() string

		// NumberLTE returns a format-string to format an NumberLTEError
		NumberLTE() string

		// NumberLT returns a format-string to format an NumberLTError
		NumberLT() string

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:127
		// RegexPattern returns a format-string to format a regex-pattern error
													RegexPattern() string

		// GreaterThanZero returns a format-string to format an error where a number must be greater than zero
		GreaterThanZero() string

		// MustBeOfA returns a format-string to format an error where a value is of the wrong type
		MustBeOfA() string

		// MustBeOfAn returns a format-string to format an error where a value is of the wrong type
		MustBeOfAn() string

		// CannotBeUsedWithout returns a format-string to format a "cannot be used without" error
		CannotBeUsedWithout() string

		// CannotBeGT returns a format-string to format an error where a value are greater than allowed
		CannotBeGT() string

		// MustBeOfType returns a format-string to format an error where a value does not match the required type
		MustBeOfType() string

		// MustBeValidRegex returns a format-string to format an error where a regex is invalid
		MustBeValidRegex() string

		// MustBeValidFormat returns a format-string to format an error where a value does not match the expected format
		MustBeValidFormat() string

		// MustBeGTEZero returns a format-string to format an error where a value must be greater or equal than 0
		MustBeGTEZero() string

		// KeyCannotBeGreaterThan returns a format-string to format an error where a key is greater than the maximum  allowed
		KeyCannotBeGreaterThan() string

		// KeyItemsMustBeOfType returns a format-string to format an error where a key is of the wrong type
		KeyItemsMustBeOfType() string

		// KeyItemsMustBeUnique returns a format-string to format an error where keys are not unique
		KeyItemsMustBeUnique() string

		// ReferenceMustBeCanonical returns a format-string to format a "reference must be canonical" error
		ReferenceMustBeCanonical() string

		// NotAValidType returns a format-string to format an invalid type error
		NotAValidType() string

		// Duplicated returns a format-string to format an error where types are duplicated
		Duplicated() string

		// HttpBadStatus returns a format-string for errors when loading a schema using HTTP
		HttpBadStatus() string

		// ParseError returns a format-string for JSON parsing errors
		ParseError() string

		// ConditionThen returns a format-string for ConditionThenError errors
		ConditionThen() string

		// ConditionElse returns a format-string for ConditionElseError errors
		ConditionElse() string

		// ErrorFormat returns a format string for errors
		ErrorFormat() string
	}

	// DefaultLocale is the default locale for this package
	DefaultLocale	struct{}
)

// False returns a format-string for "false" schema validation errors
func (l DefaultLocale) False() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:196
	_go_fuzz_dep_.CoverTab[195002]++
												return "False always fails validation"
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:197
	// _ = "end of CoverTab[195002]"
}

// Required returns a format-string for "required" schema validation errors
func (l DefaultLocale) Required() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:201
	_go_fuzz_dep_.CoverTab[195003]++
												return `{{.property}} is required`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:202
	// _ = "end of CoverTab[195003]"
}

// InvalidType returns a format-string for "invalid type" schema validation errors
func (l DefaultLocale) InvalidType() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:206
	_go_fuzz_dep_.CoverTab[195004]++
												return `Invalid type. Expected: {{.expected}}, given: {{.given}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:207
	// _ = "end of CoverTab[195004]"
}

// NumberAnyOf returns a format-string for "anyOf" schema validation errors
func (l DefaultLocale) NumberAnyOf() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:211
	_go_fuzz_dep_.CoverTab[195005]++
												return `Must validate at least one schema (anyOf)`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:212
	// _ = "end of CoverTab[195005]"
}

// NumberOneOf returns a format-string for "oneOf" schema validation errors
func (l DefaultLocale) NumberOneOf() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:216
	_go_fuzz_dep_.CoverTab[195006]++
												return `Must validate one and only one schema (oneOf)`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:217
	// _ = "end of CoverTab[195006]"
}

// NumberAllOf returns a format-string for "allOf" schema validation errors
func (l DefaultLocale) NumberAllOf() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:221
	_go_fuzz_dep_.CoverTab[195007]++
												return `Must validate all the schemas (allOf)`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:222
	// _ = "end of CoverTab[195007]"
}

// NumberNot returns a format-string to format a NumberNotError
func (l DefaultLocale) NumberNot() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:226
	_go_fuzz_dep_.CoverTab[195008]++
												return `Must not validate the schema (not)`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:227
	// _ = "end of CoverTab[195008]"
}

// MissingDependency returns a format-string for "missing dependency" schema validation errors
func (l DefaultLocale) MissingDependency() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:231
	_go_fuzz_dep_.CoverTab[195009]++
												return `Has a dependency on {{.dependency}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:232
	// _ = "end of CoverTab[195009]"
}

// Internal returns a format-string for internal errors
func (l DefaultLocale) Internal() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:236
	_go_fuzz_dep_.CoverTab[195010]++
												return `Internal Error {{.error}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:237
	// _ = "end of CoverTab[195010]"
}

// Const returns a format-string to format a ConstError
func (l DefaultLocale) Const() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:241
	_go_fuzz_dep_.CoverTab[195011]++
												return `{{.field}} does not match: {{.allowed}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:242
	// _ = "end of CoverTab[195011]"
}

// Enum returns a format-string to format an EnumError
func (l DefaultLocale) Enum() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:246
	_go_fuzz_dep_.CoverTab[195012]++
												return `{{.field}} must be one of the following: {{.allowed}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:247
	// _ = "end of CoverTab[195012]"
}

// ArrayNoAdditionalItems returns a format-string to format an ArrayNoAdditionalItemsError
func (l DefaultLocale) ArrayNoAdditionalItems() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:251
	_go_fuzz_dep_.CoverTab[195013]++
												return `No additional items allowed on array`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:252
	// _ = "end of CoverTab[195013]"
}

// ArrayNotEnoughItems returns a format-string to format an error for arrays having not enough items to match positional list of schema
func (l DefaultLocale) ArrayNotEnoughItems() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:256
	_go_fuzz_dep_.CoverTab[195014]++
												return `Not enough items on array to match positional list of schema`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:257
	// _ = "end of CoverTab[195014]"
}

// ArrayMinItems returns a format-string to format an ArrayMinItemsError
func (l DefaultLocale) ArrayMinItems() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:261
	_go_fuzz_dep_.CoverTab[195015]++
												return `Array must have at least {{.min}} items`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:262
	// _ = "end of CoverTab[195015]"
}

// ArrayMaxItems returns a format-string to format an ArrayMaxItemsError
func (l DefaultLocale) ArrayMaxItems() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:266
	_go_fuzz_dep_.CoverTab[195016]++
												return `Array must have at most {{.max}} items`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:267
	// _ = "end of CoverTab[195016]"
}

// Unique returns a format-string  to format an ItemsMustBeUniqueError
func (l DefaultLocale) Unique() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:271
	_go_fuzz_dep_.CoverTab[195017]++
												return `{{.type}} items[{{.i}},{{.j}}] must be unique`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:272
	// _ = "end of CoverTab[195017]"
}

// ArrayContains returns a format-string to format an ArrayContainsError
func (l DefaultLocale) ArrayContains() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:276
	_go_fuzz_dep_.CoverTab[195018]++
												return `At least one of the items must match`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:277
	// _ = "end of CoverTab[195018]"
}

// ArrayMinProperties returns a format-string to format an ArrayMinPropertiesError
func (l DefaultLocale) ArrayMinProperties() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:281
	_go_fuzz_dep_.CoverTab[195019]++
												return `Must have at least {{.min}} properties`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:282
	// _ = "end of CoverTab[195019]"
}

// ArrayMaxProperties returns a format-string to format an ArrayMaxPropertiesError
func (l DefaultLocale) ArrayMaxProperties() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:286
	_go_fuzz_dep_.CoverTab[195020]++
												return `Must have at most {{.max}} properties`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:287
	// _ = "end of CoverTab[195020]"
}

// AdditionalPropertyNotAllowed returns a format-string to format an AdditionalPropertyNotAllowedError
func (l DefaultLocale) AdditionalPropertyNotAllowed() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:291
	_go_fuzz_dep_.CoverTab[195021]++
												return `Additional property {{.property}} is not allowed`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:292
	// _ = "end of CoverTab[195021]"
}

// InvalidPropertyPattern returns a format-string to format an InvalidPropertyPatternError
func (l DefaultLocale) InvalidPropertyPattern() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:296
	_go_fuzz_dep_.CoverTab[195022]++
												return `Property "{{.property}}" does not match pattern {{.pattern}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:297
	// _ = "end of CoverTab[195022]"
}

// InvalidPropertyName returns a format-string to format an InvalidPropertyNameError
func (l DefaultLocale) InvalidPropertyName() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:301
	_go_fuzz_dep_.CoverTab[195023]++
												return `Property name of "{{.property}}" does not match`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:302
	// _ = "end of CoverTab[195023]"
}

// StringGTE returns a format-string to format an StringLengthGTEError
func (l DefaultLocale) StringGTE() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:306
	_go_fuzz_dep_.CoverTab[195024]++
												return `String length must be greater than or equal to {{.min}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:307
	// _ = "end of CoverTab[195024]"
}

// StringLTE returns a format-string to format an StringLengthLTEError
func (l DefaultLocale) StringLTE() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:311
	_go_fuzz_dep_.CoverTab[195025]++
												return `String length must be less than or equal to {{.max}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:312
	// _ = "end of CoverTab[195025]"
}

// DoesNotMatchPattern returns a format-string to format an DoesNotMatchPatternError
func (l DefaultLocale) DoesNotMatchPattern() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:316
	_go_fuzz_dep_.CoverTab[195026]++
												return `Does not match pattern '{{.pattern}}'`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:317
	// _ = "end of CoverTab[195026]"
}

// DoesNotMatchFormat returns a format-string to format an DoesNotMatchFormatError
func (l DefaultLocale) DoesNotMatchFormat() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:321
	_go_fuzz_dep_.CoverTab[195027]++
												return `Does not match format '{{.format}}'`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:322
	// _ = "end of CoverTab[195027]"
}

// MultipleOf returns a format-string to format an MultipleOfError
func (l DefaultLocale) MultipleOf() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:326
	_go_fuzz_dep_.CoverTab[195028]++
												return `Must be a multiple of {{.multiple}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:327
	// _ = "end of CoverTab[195028]"
}

// NumberGTE returns the format string to format a NumberGTEError
func (l DefaultLocale) NumberGTE() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:331
	_go_fuzz_dep_.CoverTab[195029]++
												return `Must be greater than or equal to {{.min}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:332
	// _ = "end of CoverTab[195029]"
}

// NumberGT returns the format string to format a NumberGTError
func (l DefaultLocale) NumberGT() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:336
	_go_fuzz_dep_.CoverTab[195030]++
												return `Must be greater than {{.min}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:337
	// _ = "end of CoverTab[195030]"
}

// NumberLTE returns the format string to format a NumberLTEError
func (l DefaultLocale) NumberLTE() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:341
	_go_fuzz_dep_.CoverTab[195031]++
												return `Must be less than or equal to {{.max}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:342
	// _ = "end of CoverTab[195031]"
}

// NumberLT returns the format string to format a NumberLTError
func (l DefaultLocale) NumberLT() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:346
	_go_fuzz_dep_.CoverTab[195032]++
												return `Must be less than {{.max}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:347
	// _ = "end of CoverTab[195032]"
}

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:352
// RegexPattern returns a format-string to format a regex-pattern error
func (l DefaultLocale) RegexPattern() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:353
	_go_fuzz_dep_.CoverTab[195033]++
												return `Invalid regex pattern '{{.pattern}}'`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:354
	// _ = "end of CoverTab[195033]"
}

// GreaterThanZero returns a format-string to format an error where a number must be greater than zero
func (l DefaultLocale) GreaterThanZero() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:358
	_go_fuzz_dep_.CoverTab[195034]++
												return `{{.number}} must be strictly greater than 0`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:359
	// _ = "end of CoverTab[195034]"
}

// MustBeOfA returns a format-string to format an error where a value is of the wrong type
func (l DefaultLocale) MustBeOfA() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:363
	_go_fuzz_dep_.CoverTab[195035]++
												return `{{.x}} must be of a {{.y}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:364
	// _ = "end of CoverTab[195035]"
}

// MustBeOfAn returns a format-string to format an error where a value is of the wrong type
func (l DefaultLocale) MustBeOfAn() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:368
	_go_fuzz_dep_.CoverTab[195036]++
												return `{{.x}} must be of an {{.y}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:369
	// _ = "end of CoverTab[195036]"
}

// CannotBeUsedWithout returns a format-string to format a "cannot be used without" error
func (l DefaultLocale) CannotBeUsedWithout() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:373
	_go_fuzz_dep_.CoverTab[195037]++
												return `{{.x}} cannot be used without {{.y}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:374
	// _ = "end of CoverTab[195037]"
}

// CannotBeGT returns a format-string to format an error where a value are greater than allowed
func (l DefaultLocale) CannotBeGT() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:378
	_go_fuzz_dep_.CoverTab[195038]++
												return `{{.x}} cannot be greater than {{.y}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:379
	// _ = "end of CoverTab[195038]"
}

// MustBeOfType returns a format-string to format an error where a value does not match the required type
func (l DefaultLocale) MustBeOfType() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:383
	_go_fuzz_dep_.CoverTab[195039]++
												return `{{.key}} must be of type {{.type}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:384
	// _ = "end of CoverTab[195039]"
}

// MustBeValidRegex returns a format-string to format an error where a regex is invalid
func (l DefaultLocale) MustBeValidRegex() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:388
	_go_fuzz_dep_.CoverTab[195040]++
												return `{{.key}} must be a valid regex`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:389
	// _ = "end of CoverTab[195040]"
}

// MustBeValidFormat returns a format-string to format an error where a value does not match the expected format
func (l DefaultLocale) MustBeValidFormat() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:393
	_go_fuzz_dep_.CoverTab[195041]++
												return `{{.key}} must be a valid format {{.given}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:394
	// _ = "end of CoverTab[195041]"
}

// MustBeGTEZero returns a format-string to format an error where a value must be greater or equal than 0
func (l DefaultLocale) MustBeGTEZero() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:398
	_go_fuzz_dep_.CoverTab[195042]++
												return `{{.key}} must be greater than or equal to 0`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:399
	// _ = "end of CoverTab[195042]"
}

// KeyCannotBeGreaterThan returns a format-string to format an error where a value is greater than the maximum  allowed
func (l DefaultLocale) KeyCannotBeGreaterThan() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:403
	_go_fuzz_dep_.CoverTab[195043]++
												return `{{.key}} cannot be greater than {{.y}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:404
	// _ = "end of CoverTab[195043]"
}

// KeyItemsMustBeOfType returns a format-string to format an error where a key is of the wrong type
func (l DefaultLocale) KeyItemsMustBeOfType() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:408
	_go_fuzz_dep_.CoverTab[195044]++
												return `{{.key}} items must be {{.type}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:409
	// _ = "end of CoverTab[195044]"
}

// KeyItemsMustBeUnique returns a format-string to format an error where keys are not unique
func (l DefaultLocale) KeyItemsMustBeUnique() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:413
	_go_fuzz_dep_.CoverTab[195045]++
												return `{{.key}} items must be unique`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:414
	// _ = "end of CoverTab[195045]"
}

// ReferenceMustBeCanonical returns a format-string to format a "reference must be canonical" error
func (l DefaultLocale) ReferenceMustBeCanonical() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:418
	_go_fuzz_dep_.CoverTab[195046]++
												return `Reference {{.reference}} must be canonical`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:419
	// _ = "end of CoverTab[195046]"
}

// NotAValidType returns a format-string to format an invalid type error
func (l DefaultLocale) NotAValidType() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:423
	_go_fuzz_dep_.CoverTab[195047]++
												return `has a primitive type that is NOT VALID -- given: {{.given}} Expected valid values are:{{.expected}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:424
	// _ = "end of CoverTab[195047]"
}

// Duplicated returns a format-string to format an error where types are duplicated
func (l DefaultLocale) Duplicated() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:428
	_go_fuzz_dep_.CoverTab[195048]++
												return `{{.type}} type is duplicated`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:429
	// _ = "end of CoverTab[195048]"
}

// HttpBadStatus returns a format-string for errors when loading a schema using HTTP
func (l DefaultLocale) HttpBadStatus() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:433
	_go_fuzz_dep_.CoverTab[195049]++
												return `Could not read schema from HTTP, response status is {{.status}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:434
	// _ = "end of CoverTab[195049]"
}

// ErrorFormat returns a format string for errors
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:437
// Replacement options: field, description, context, value
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:439
func (l DefaultLocale) ErrorFormat() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:439
	_go_fuzz_dep_.CoverTab[195050]++
												return `{{.field}}: {{.description}}`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:440
	// _ = "end of CoverTab[195050]"
}

// ParseError returns a format-string for JSON parsing errors
func (l DefaultLocale) ParseError() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:444
	_go_fuzz_dep_.CoverTab[195051]++
												return `Expected: {{.expected}}, given: Invalid JSON`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:445
	// _ = "end of CoverTab[195051]"
}

// ConditionThen returns a format-string for ConditionThenError errors
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:448
// If/Else
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:450
func (l DefaultLocale) ConditionThen() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:450
	_go_fuzz_dep_.CoverTab[195052]++
												return `Must validate "then" as "if" was valid`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:451
	// _ = "end of CoverTab[195052]"
}

// ConditionElse returns a format-string for ConditionElseError errors
func (l DefaultLocale) ConditionElse() string {
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:455
	_go_fuzz_dep_.CoverTab[195053]++
												return `Must validate "else" as "if" was not valid`
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:456
	// _ = "end of CoverTab[195053]"
}

// constants
const (
	STRING_NUMBER				= "number"
	STRING_ARRAY_OF_STRINGS			= "array of strings"
	STRING_ARRAY_OF_SCHEMAS			= "array of schemas"
	STRING_SCHEMA				= "valid schema"
	STRING_SCHEMA_OR_ARRAY_OF_STRINGS	= "schema or array of strings"
	STRING_PROPERTIES			= "properties"
	STRING_DEPENDENCY			= "dependency"
	STRING_PROPERTY				= "property"
	STRING_UNDEFINED			= "undefined"
	STRING_CONTEXT_ROOT			= "(root)"
	STRING_ROOT_SCHEMA_PROPERTY		= "(root)"
)

//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:472
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/xeipuuv/gojsonschema@v1.2.0/locales.go:472
var _ = _go_fuzz_dep_.CoverTab
