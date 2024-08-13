//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Package mapstructure exposes functionality to convert one arbitrary
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Go type into another, typically to convert a map[string]interface{}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// into a native Go structure.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// The Go structure can be arbitrarily complex, containing slices,
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// other structs, etc. and the decoder will properly decode nested
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// maps and so on into the proper structures in the native Go struct.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// See the examples to see what the decoder is capable of.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// The simplest function to start with is Decode.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// # Field Tags
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// When decoding to a struct, mapstructure will use the field name by
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// default to perform the mapping. For example, if a struct has a field
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// "Username" then mapstructure will look for a key in the source value
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// of "username" (case insensitive).
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type User struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Username string
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// You can change the behavior of mapstructure by using struct tags.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// The default struct tag that mapstructure looks for is "mapstructure"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// but you can customize it using DecoderConfig.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// # Renaming Fields
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// To rename the key that mapstructure looks for, use the "mapstructure"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// tag and set a value directly. For example, to change the "username" example
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// above to "user":
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type User struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Username string `mapstructure:"user"`
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// # Embedded Structs and Squashing
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Embedded structs are treated as if they're another field with that name.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// By default, the two structs below are equivalent when decoding with
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// mapstructure:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Person struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Name string
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Friend struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Person
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Friend struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Person Person
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// This would require an input that looks like below:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	map[string]interface{}{
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    "person": map[string]interface{}{"name": "alice"},
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// If your "person" value is NOT nested, then you can append ",squash" to
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// your tag value and mapstructure will treat it as if the embedded struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// were part of the struct directly. Example:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Friend struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Person `mapstructure:",squash"`
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Now the following input would be accepted:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	map[string]interface{}{
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    "name": "alice",
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// When decoding from a struct to a map, the squash tag squashes the struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// fields into a single map. Using the example structs from above:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	Friend{Person: Person{Name: "alice"}}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Will be decoded into a map:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	map[string]interface{}{
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    "name": "alice",
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// DecoderConfig has a field that changes the behavior of mapstructure
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// to always squash embedded structs.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// # Remainder Values
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// If there are any unmapped keys in the source value, mapstructure by
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// default will silently ignore them. You can error by setting ErrorUnused
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// in DecoderConfig. If you're using Metadata you can also maintain a slice
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// of the unused keys.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// You can also use the ",remain" suffix on your tag to collect all unused
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// values in a map. The field with this tag MUST be a map type and should
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// probably be a "map[string]interface{}" or "map[interface{}]interface{}".
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// See example below:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Friend struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Name  string
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Other map[string]interface{} `mapstructure:",remain"`
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Given the input below, Other would be populated with the other
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// values that weren't used (everything but "name"):
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	map[string]interface{}{
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    "name":    "bob",
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    "address": "123 Maple St.",
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// # Omit Empty Values
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// When decoding from a struct to any other value, you may use the
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// ",omitempty" suffix on your tag to omit that value if it equates to
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// the zero value. The zero value of all types is specified in the Go
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// specification.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// For example, the zero type of a numeric type is zero ("0"). If the struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// field value is zero and a numeric type, the field is empty, and it won't
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// be encoded into the destination type.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Source {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Age int `mapstructure:",omitempty"`
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// # Unexported fields
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Since unexported (private) struct fields cannot be set outside the package
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// where they are defined, the decoder will simply skip them.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// For this output type definition:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Exported struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    private string // this unexported field will be skipped
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Public string
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// Using this map as input:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	map[string]interface{}{
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    "private": "I will be ignored",
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    "Public":  "I made it through!",
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// The following struct will be decoded:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	type Exported struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    private: "" // field is left with an empty string (zero value)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	    Public: "I made it through!"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// # Other Configuration
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// mapstructure is highly configurable. See the DecoderConfig struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1
// for other features and options that are supported.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:159
package mapstructure

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:159
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:159
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:159
)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:159
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:159
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:159
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// DecodeHookFunc is the callback function that can be used for
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// data transformations. See "DecodeHook" in the DecoderConfig
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// struct.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// The type must be one of DecodeHookFuncType, DecodeHookFuncKind, or
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// DecodeHookFuncValue.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// Values are a superset of Types (Values can return types), and Types are a
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// superset of Kinds (Types can return Kinds) and are generally a richer thing
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// to use, but Kinds are simpler if you only need those.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// The reason DecodeHookFunc is multi-typed is for backwards compatibility:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// we started with Kinds and then realized Types were the better solution,
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// but have a promise to not break backwards compat so we now support
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:171
// both.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:185
type DecodeHookFunc interface{}

// DecodeHookFuncType is a DecodeHookFunc which has complete information about
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:187
// the source and target types.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:189
type DecodeHookFuncType func(reflect.Type, reflect.Type, interface{}) (interface{}, error)

// DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:191
// source and target types.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:193
type DecodeHookFuncKind func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error)

// DecodeHookFuncValue is a DecodeHookFunc which has complete access to both the source and target
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:195
// values.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:197
type DecodeHookFuncValue func(from reflect.Value, to reflect.Value) (interface{}, error)

// DecoderConfig is the configuration that is used to create a new decoder
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:199
// and allows customization of various aspects of decoding.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:201
type DecoderConfig struct {
	// DecodeHook, if set, will be called before any decoding and any
	// type conversion (if WeaklyTypedInput is on). This lets you modify
	// the values before they're set down onto the resulting struct. The
	// DecodeHook is called for every map and value in the input. This means
	// that if a struct has embedded fields with squash tags the decode hook
	// is called only once with all of the input data, not once for each
	// embedded struct.
	//
	// If an error is returned, the entire decode will fail with that error.
	DecodeHook	DecodeHookFunc

	// If ErrorUnused is true, then it is an error for there to exist
	// keys in the original map that were unused in the decoding process
	// (extra keys).
	ErrorUnused	bool

	// ZeroFields, if set to true, will zero fields before writing them.
	// For example, a map will be emptied before decoded values are put in
	// it. If this is false, a map will be merged.
	ZeroFields	bool

	// If WeaklyTypedInput is true, the decoder will make the following
	// "weak" conversions:
	//
	//   - bools to string (true = "1", false = "0")
	//   - numbers to string (base 10)
	//   - bools to int/uint (true = 1, false = 0)
	//   - strings to int/uint (base implied by prefix)
	//   - int to bool (true if value != 0)
	//   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
	//     FALSE, false, False. Anything else is an error)
	//   - empty array = empty map and vice versa
	//   - negative numbers to overflowed uint values (base 10)
	//   - slice of maps to a merged map
	//   - single values are converted to slices if required. Each
	//     element is weakly decoded. For example: "4" can become []int{4}
	//     if the target type is an int slice.
	//
	WeaklyTypedInput	bool

	// Squash will squash embedded structs.  A squash tag may also be
	// added to an individual struct field using a tag.  For example:
	//
	//  type Parent struct {
	//      Child `mapstructure:",squash"`
	//  }
	Squash	bool

	// Metadata is the struct that will contain extra metadata about
	// the decoding. If this is nil, then no metadata will be tracked.
	Metadata	*Metadata

	// Result is a pointer to the struct that will contain the decoded
	// value.
	Result	interface{}

	// The tag name that mapstructure reads for field names. This
	// defaults to "mapstructure"
	TagName	string

	// MatchName is the function used to match the map key to the struct
	// field name or tag. Defaults to `strings.EqualFold`. This can be used
	// to implement case-sensitive tag values, support snake casing, etc.
	MatchName	func(mapKey, fieldName string) bool
}

// A Decoder takes a raw interface value and turns it into structured
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:268
// data, keeping track of rich error information along the way in case
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:268
// anything goes wrong. Unlike the basic top-level Decode method, you can
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:268
// more finely control how the Decoder behaves using the DecoderConfig
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:268
// structure. The top-level Decode method is just a convenience that sets
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:268
// up the most basic Decoder.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:274
type Decoder struct {
	config *DecoderConfig
}

// Metadata contains information about decoding a structure that
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:278
// is tedious or difficult to get otherwise.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:280
type Metadata struct {
	// Keys are the keys of the structure which were successfully decoded
	Keys	[]string

	// Unused is a slice of keys that were found in the raw value but
	// weren't decoded since there was no matching field in the result interface
	Unused	[]string
}

// Decode takes an input structure and uses reflection to translate it to
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:289
// the output structure. output must be a pointer to a map or struct.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:291
func Decode(input interface{}, output interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:291
	_go_fuzz_dep_.CoverTab[116281]++
												config := &DecoderConfig{
		Metadata:	nil,
		Result:		output,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:298
		_go_fuzz_dep_.CoverTab[116283]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:299
		// _ = "end of CoverTab[116283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:300
		_go_fuzz_dep_.CoverTab[116284]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:300
		// _ = "end of CoverTab[116284]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:300
	// _ = "end of CoverTab[116281]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:300
	_go_fuzz_dep_.CoverTab[116282]++

												return decoder.Decode(input)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:302
	// _ = "end of CoverTab[116282]"
}

// WeakDecode is the same as Decode but is shorthand to enable
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:305
// WeaklyTypedInput. See DecoderConfig for more info.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:307
func WeakDecode(input, output interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:307
	_go_fuzz_dep_.CoverTab[116285]++
												config := &DecoderConfig{
		Metadata:		nil,
		Result:			output,
		WeaklyTypedInput:	true,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:315
		_go_fuzz_dep_.CoverTab[116287]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:316
		// _ = "end of CoverTab[116287]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:317
		_go_fuzz_dep_.CoverTab[116288]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:317
		// _ = "end of CoverTab[116288]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:317
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:317
	// _ = "end of CoverTab[116285]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:317
	_go_fuzz_dep_.CoverTab[116286]++

												return decoder.Decode(input)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:319
	// _ = "end of CoverTab[116286]"
}

// DecodeMetadata is the same as Decode, but is shorthand to
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:322
// enable metadata collection. See DecoderConfig for more info.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:324
func DecodeMetadata(input interface{}, output interface{}, metadata *Metadata) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:324
	_go_fuzz_dep_.CoverTab[116289]++
												config := &DecoderConfig{
		Metadata:	metadata,
		Result:		output,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:331
		_go_fuzz_dep_.CoverTab[116291]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:332
		// _ = "end of CoverTab[116291]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:333
		_go_fuzz_dep_.CoverTab[116292]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:333
		// _ = "end of CoverTab[116292]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:333
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:333
	// _ = "end of CoverTab[116289]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:333
	_go_fuzz_dep_.CoverTab[116290]++

												return decoder.Decode(input)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:335
	// _ = "end of CoverTab[116290]"
}

// WeakDecodeMetadata is the same as Decode, but is shorthand to
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:338
// enable both WeaklyTypedInput and metadata collection. See
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:338
// DecoderConfig for more info.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:341
func WeakDecodeMetadata(input interface{}, output interface{}, metadata *Metadata) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:341
	_go_fuzz_dep_.CoverTab[116293]++
												config := &DecoderConfig{
		Metadata:		metadata,
		Result:			output,
		WeaklyTypedInput:	true,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:349
		_go_fuzz_dep_.CoverTab[116295]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:350
		// _ = "end of CoverTab[116295]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:351
		_go_fuzz_dep_.CoverTab[116296]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:351
		// _ = "end of CoverTab[116296]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:351
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:351
	// _ = "end of CoverTab[116293]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:351
	_go_fuzz_dep_.CoverTab[116294]++

												return decoder.Decode(input)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:353
	// _ = "end of CoverTab[116294]"
}

// NewDecoder returns a new decoder for the given configuration. Once
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:356
// a decoder has been returned, the same configuration must not be used
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:356
// again.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:359
func NewDecoder(config *DecoderConfig) (*Decoder, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:359
	_go_fuzz_dep_.CoverTab[116297]++
												val := reflect.ValueOf(config.Result)
												if val.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:361
		_go_fuzz_dep_.CoverTab[116303]++
													return nil, errors.New("result must be a pointer")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:362
		// _ = "end of CoverTab[116303]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:363
		_go_fuzz_dep_.CoverTab[116304]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:363
		// _ = "end of CoverTab[116304]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:363
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:363
	// _ = "end of CoverTab[116297]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:363
	_go_fuzz_dep_.CoverTab[116298]++

												val = val.Elem()
												if !val.CanAddr() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:366
		_go_fuzz_dep_.CoverTab[116305]++
													return nil, errors.New("result must be addressable (a pointer)")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:367
		// _ = "end of CoverTab[116305]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:368
		_go_fuzz_dep_.CoverTab[116306]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:368
		// _ = "end of CoverTab[116306]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:368
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:368
	// _ = "end of CoverTab[116298]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:368
	_go_fuzz_dep_.CoverTab[116299]++

												if config.Metadata != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:370
		_go_fuzz_dep_.CoverTab[116307]++
													if config.Metadata.Keys == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:371
			_go_fuzz_dep_.CoverTab[116309]++
														config.Metadata.Keys = make([]string, 0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:372
			// _ = "end of CoverTab[116309]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:373
			_go_fuzz_dep_.CoverTab[116310]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:373
			// _ = "end of CoverTab[116310]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:373
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:373
		// _ = "end of CoverTab[116307]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:373
		_go_fuzz_dep_.CoverTab[116308]++

													if config.Metadata.Unused == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:375
			_go_fuzz_dep_.CoverTab[116311]++
														config.Metadata.Unused = make([]string, 0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:376
			// _ = "end of CoverTab[116311]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:377
			_go_fuzz_dep_.CoverTab[116312]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:377
			// _ = "end of CoverTab[116312]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:377
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:377
		// _ = "end of CoverTab[116308]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:378
		_go_fuzz_dep_.CoverTab[116313]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:378
		// _ = "end of CoverTab[116313]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:378
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:378
	// _ = "end of CoverTab[116299]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:378
	_go_fuzz_dep_.CoverTab[116300]++

												if config.TagName == "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:380
		_go_fuzz_dep_.CoverTab[116314]++
													config.TagName = "mapstructure"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:381
		// _ = "end of CoverTab[116314]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:382
		_go_fuzz_dep_.CoverTab[116315]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:382
		// _ = "end of CoverTab[116315]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:382
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:382
	// _ = "end of CoverTab[116300]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:382
	_go_fuzz_dep_.CoverTab[116301]++

												if config.MatchName == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:384
		_go_fuzz_dep_.CoverTab[116316]++
													config.MatchName = strings.EqualFold
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:385
		// _ = "end of CoverTab[116316]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:386
		_go_fuzz_dep_.CoverTab[116317]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:386
		// _ = "end of CoverTab[116317]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:386
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:386
	// _ = "end of CoverTab[116301]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:386
	_go_fuzz_dep_.CoverTab[116302]++

												result := &Decoder{
		config: config,
	}

												return result, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:392
	// _ = "end of CoverTab[116302]"
}

// Decode decodes the given raw interface to the target pointer specified
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:395
// by the configuration.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:397
func (d *Decoder) Decode(input interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:397
	_go_fuzz_dep_.CoverTab[116318]++
												return d.decode("", input, reflect.ValueOf(d.config.Result).Elem())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:398
	// _ = "end of CoverTab[116318]"
}

// Decodes an unknown data type into a specific reflection value.
func (d *Decoder) decode(name string, input interface{}, outVal reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:402
	_go_fuzz_dep_.CoverTab[116319]++
												var inputVal reflect.Value
												if input != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:404
		_go_fuzz_dep_.CoverTab[116326]++
													inputVal = reflect.ValueOf(input)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:409
		if inputVal.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:409
			_go_fuzz_dep_.CoverTab[116327]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:409
			return inputVal.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:409
			// _ = "end of CoverTab[116327]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:409
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:409
			_go_fuzz_dep_.CoverTab[116328]++
														input = nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:410
			// _ = "end of CoverTab[116328]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:411
			_go_fuzz_dep_.CoverTab[116329]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:411
			// _ = "end of CoverTab[116329]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:411
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:411
		// _ = "end of CoverTab[116326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:412
		_go_fuzz_dep_.CoverTab[116330]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:412
		// _ = "end of CoverTab[116330]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:412
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:412
	// _ = "end of CoverTab[116319]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:412
	_go_fuzz_dep_.CoverTab[116320]++

												if input == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:414
		_go_fuzz_dep_.CoverTab[116331]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:417
		if d.config.ZeroFields {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:417
			_go_fuzz_dep_.CoverTab[116333]++
														outVal.Set(reflect.Zero(outVal.Type()))

														if d.config.Metadata != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:420
				_go_fuzz_dep_.CoverTab[116334]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:420
				return name != ""
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:420
				// _ = "end of CoverTab[116334]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:420
			}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:420
				_go_fuzz_dep_.CoverTab[116335]++
															d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:421
				// _ = "end of CoverTab[116335]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:422
				_go_fuzz_dep_.CoverTab[116336]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:422
				// _ = "end of CoverTab[116336]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:422
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:422
			// _ = "end of CoverTab[116333]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:423
			_go_fuzz_dep_.CoverTab[116337]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:423
			// _ = "end of CoverTab[116337]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:423
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:423
		// _ = "end of CoverTab[116331]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:423
		_go_fuzz_dep_.CoverTab[116332]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:424
		// _ = "end of CoverTab[116332]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:425
		_go_fuzz_dep_.CoverTab[116338]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:425
		// _ = "end of CoverTab[116338]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:425
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:425
	// _ = "end of CoverTab[116320]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:425
	_go_fuzz_dep_.CoverTab[116321]++

												if !inputVal.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:427
		_go_fuzz_dep_.CoverTab[116339]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:430
		outVal.Set(reflect.Zero(outVal.Type()))
		if d.config.Metadata != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:431
			_go_fuzz_dep_.CoverTab[116341]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:431
			return name != ""
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:431
			// _ = "end of CoverTab[116341]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:431
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:431
			_go_fuzz_dep_.CoverTab[116342]++
														d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:432
			// _ = "end of CoverTab[116342]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:433
			_go_fuzz_dep_.CoverTab[116343]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:433
			// _ = "end of CoverTab[116343]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:433
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:433
		// _ = "end of CoverTab[116339]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:433
		_go_fuzz_dep_.CoverTab[116340]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:434
		// _ = "end of CoverTab[116340]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:435
		_go_fuzz_dep_.CoverTab[116344]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:435
		// _ = "end of CoverTab[116344]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:435
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:435
	// _ = "end of CoverTab[116321]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:435
	_go_fuzz_dep_.CoverTab[116322]++

												if d.config.DecodeHook != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:437
		_go_fuzz_dep_.CoverTab[116345]++
		// We have a DecodeHook, so let's pre-process the input.
		var err error
		input, err = DecodeHookExec(d.config.DecodeHook, inputVal, outVal)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:441
			_go_fuzz_dep_.CoverTab[116346]++
														return fmt.Errorf("error decoding '%s': %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:442
			// _ = "end of CoverTab[116346]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:443
			_go_fuzz_dep_.CoverTab[116347]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:443
			// _ = "end of CoverTab[116347]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:443
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:443
		// _ = "end of CoverTab[116345]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:444
		_go_fuzz_dep_.CoverTab[116348]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:444
		// _ = "end of CoverTab[116348]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:444
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:444
	// _ = "end of CoverTab[116322]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:444
	_go_fuzz_dep_.CoverTab[116323]++

												var err error
												outputKind := getKind(outVal)
												addMetaKey := true
												switch outputKind {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:450
		_go_fuzz_dep_.CoverTab[116349]++
													err = d.decodeBool(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:451
		// _ = "end of CoverTab[116349]"
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:452
		_go_fuzz_dep_.CoverTab[116350]++
													err = d.decodeBasic(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:453
		// _ = "end of CoverTab[116350]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:454
		_go_fuzz_dep_.CoverTab[116351]++
													err = d.decodeString(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:455
		// _ = "end of CoverTab[116351]"
	case reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:456
		_go_fuzz_dep_.CoverTab[116352]++
													err = d.decodeInt(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:457
		// _ = "end of CoverTab[116352]"
	case reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:458
		_go_fuzz_dep_.CoverTab[116353]++
													err = d.decodeUint(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:459
		// _ = "end of CoverTab[116353]"
	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:460
		_go_fuzz_dep_.CoverTab[116354]++
													err = d.decodeFloat(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:461
		// _ = "end of CoverTab[116354]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:462
		_go_fuzz_dep_.CoverTab[116355]++
													err = d.decodeStruct(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:463
		// _ = "end of CoverTab[116355]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:464
		_go_fuzz_dep_.CoverTab[116356]++
													err = d.decodeMap(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:465
		// _ = "end of CoverTab[116356]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:466
		_go_fuzz_dep_.CoverTab[116357]++
													addMetaKey, err = d.decodePtr(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:467
		// _ = "end of CoverTab[116357]"
	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:468
		_go_fuzz_dep_.CoverTab[116358]++
													err = d.decodeSlice(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:469
		// _ = "end of CoverTab[116358]"
	case reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:470
		_go_fuzz_dep_.CoverTab[116359]++
													err = d.decodeArray(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:471
		// _ = "end of CoverTab[116359]"
	case reflect.Func:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:472
		_go_fuzz_dep_.CoverTab[116360]++
													err = d.decodeFunc(name, input, outVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:473
		// _ = "end of CoverTab[116360]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:474
		_go_fuzz_dep_.CoverTab[116361]++

													return fmt.Errorf("%s: unsupported type: %s", name, outputKind)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:476
		// _ = "end of CoverTab[116361]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:477
	// _ = "end of CoverTab[116323]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:477
	_go_fuzz_dep_.CoverTab[116324]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
	if addMetaKey && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
		_go_fuzz_dep_.CoverTab[116362]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
		return d.config.Metadata != nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
		// _ = "end of CoverTab[116362]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
		_go_fuzz_dep_.CoverTab[116363]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
		return name != ""
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
		// _ = "end of CoverTab[116363]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:481
		_go_fuzz_dep_.CoverTab[116364]++
													d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:482
		// _ = "end of CoverTab[116364]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:483
		_go_fuzz_dep_.CoverTab[116365]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:483
		// _ = "end of CoverTab[116365]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:483
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:483
	// _ = "end of CoverTab[116324]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:483
	_go_fuzz_dep_.CoverTab[116325]++

												return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:485
	// _ = "end of CoverTab[116325]"
}

// This decodes a basic type (bool, int, string, etc.) and sets the
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:488
// value to "data" of that type.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:490
func (d *Decoder) decodeBasic(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:490
	_go_fuzz_dep_.CoverTab[116366]++
												if val.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:491
		_go_fuzz_dep_.CoverTab[116371]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:491
		return val.Elem().IsValid()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:491
		// _ = "end of CoverTab[116371]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:491
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:491
		_go_fuzz_dep_.CoverTab[116372]++
													elem := val.Elem()

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:497
		copied := false
		if !elem.CanAddr() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:498
			_go_fuzz_dep_.CoverTab[116375]++
														copied = true

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:502
			copy := reflect.New(elem.Type())

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:505
			copy.Elem().Set(elem)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:508
			elem = copy
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:508
			// _ = "end of CoverTab[116375]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:509
			_go_fuzz_dep_.CoverTab[116376]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:509
			// _ = "end of CoverTab[116376]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:509
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:509
		// _ = "end of CoverTab[116372]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:509
		_go_fuzz_dep_.CoverTab[116373]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:513
		if err := d.decode(name, data, elem); err != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:513
			_go_fuzz_dep_.CoverTab[116377]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:513
			return !copied
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:513
			// _ = "end of CoverTab[116377]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:513
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:513
			_go_fuzz_dep_.CoverTab[116378]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:514
			// _ = "end of CoverTab[116378]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:515
			_go_fuzz_dep_.CoverTab[116379]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:515
			// _ = "end of CoverTab[116379]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:515
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:515
		// _ = "end of CoverTab[116373]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:515
		_go_fuzz_dep_.CoverTab[116374]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:518
		val.Set(elem.Elem())
													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:519
		// _ = "end of CoverTab[116374]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:520
		_go_fuzz_dep_.CoverTab[116380]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:520
		// _ = "end of CoverTab[116380]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:520
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:520
	// _ = "end of CoverTab[116366]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:520
	_go_fuzz_dep_.CoverTab[116367]++

												dataVal := reflect.ValueOf(data)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:527
	if dataVal.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:527
		_go_fuzz_dep_.CoverTab[116381]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:527
		return dataVal.Type().Elem() == val.Type()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:527
		// _ = "end of CoverTab[116381]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:527
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:527
		_go_fuzz_dep_.CoverTab[116382]++
													dataVal = reflect.Indirect(dataVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:528
		// _ = "end of CoverTab[116382]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:529
		_go_fuzz_dep_.CoverTab[116383]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:529
		// _ = "end of CoverTab[116383]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:529
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:529
	// _ = "end of CoverTab[116367]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:529
	_go_fuzz_dep_.CoverTab[116368]++

												if !dataVal.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:531
		_go_fuzz_dep_.CoverTab[116384]++
													dataVal = reflect.Zero(val.Type())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:532
		// _ = "end of CoverTab[116384]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:533
		_go_fuzz_dep_.CoverTab[116385]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:533
		// _ = "end of CoverTab[116385]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:533
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:533
	// _ = "end of CoverTab[116368]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:533
	_go_fuzz_dep_.CoverTab[116369]++

												dataValType := dataVal.Type()
												if !dataValType.AssignableTo(val.Type()) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:536
		_go_fuzz_dep_.CoverTab[116386]++
													return fmt.Errorf(
			"'%s' expected type '%s', got '%s'",
			name, val.Type(), dataValType)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:539
		// _ = "end of CoverTab[116386]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:540
		_go_fuzz_dep_.CoverTab[116387]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:540
		// _ = "end of CoverTab[116387]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:540
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:540
	// _ = "end of CoverTab[116369]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:540
	_go_fuzz_dep_.CoverTab[116370]++

												val.Set(dataVal)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:543
	// _ = "end of CoverTab[116370]"
}

func (d *Decoder) decodeString(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:546
	_go_fuzz_dep_.CoverTab[116388]++
												dataVal := reflect.Indirect(reflect.ValueOf(data))
												dataKind := getKind(dataVal)

												converted := true
												switch {
	case dataKind == reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:552
		_go_fuzz_dep_.CoverTab[116391]++
													val.SetString(dataVal.String())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:553
		// _ = "end of CoverTab[116391]"
	case dataKind == reflect.Bool && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:554
		_go_fuzz_dep_.CoverTab[116398]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:554
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:554
		// _ = "end of CoverTab[116398]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:554
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:554
		_go_fuzz_dep_.CoverTab[116392]++
													if dataVal.Bool() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:555
			_go_fuzz_dep_.CoverTab[116399]++
														val.SetString("1")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:556
			// _ = "end of CoverTab[116399]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:557
			_go_fuzz_dep_.CoverTab[116400]++
														val.SetString("0")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:558
			// _ = "end of CoverTab[116400]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:559
		// _ = "end of CoverTab[116392]"
	case dataKind == reflect.Int && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:560
		_go_fuzz_dep_.CoverTab[116401]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:560
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:560
		// _ = "end of CoverTab[116401]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:560
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:560
		_go_fuzz_dep_.CoverTab[116393]++
													val.SetString(strconv.FormatInt(dataVal.Int(), 10))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:561
		// _ = "end of CoverTab[116393]"
	case dataKind == reflect.Uint && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:562
		_go_fuzz_dep_.CoverTab[116402]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:562
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:562
		// _ = "end of CoverTab[116402]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:562
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:562
		_go_fuzz_dep_.CoverTab[116394]++
													val.SetString(strconv.FormatUint(dataVal.Uint(), 10))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:563
		// _ = "end of CoverTab[116394]"
	case dataKind == reflect.Float32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:564
		_go_fuzz_dep_.CoverTab[116403]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:564
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:564
		// _ = "end of CoverTab[116403]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:564
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:564
		_go_fuzz_dep_.CoverTab[116395]++
													val.SetString(strconv.FormatFloat(dataVal.Float(), 'f', -1, 64))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:565
		// _ = "end of CoverTab[116395]"
	case dataKind == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:566
		_go_fuzz_dep_.CoverTab[116404]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:566
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:566
		// _ = "end of CoverTab[116404]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:566
	}(),
		dataKind == reflect.Array && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:567
			_go_fuzz_dep_.CoverTab[116405]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:567
			return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:567
			// _ = "end of CoverTab[116405]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:567
		}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:567
		_go_fuzz_dep_.CoverTab[116396]++
													dataType := dataVal.Type()
													elemKind := dataType.Elem().Kind()
													switch elemKind {
		case reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:571
			_go_fuzz_dep_.CoverTab[116406]++
														var uints []uint8
														if dataKind == reflect.Array {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:573
				_go_fuzz_dep_.CoverTab[116409]++
															uints = make([]uint8, dataVal.Len(), dataVal.Len())
															for i := range uints {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:575
					_go_fuzz_dep_.CoverTab[116410]++
																uints[i] = dataVal.Index(i).Interface().(uint8)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:576
					// _ = "end of CoverTab[116410]"
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:577
				// _ = "end of CoverTab[116409]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:578
				_go_fuzz_dep_.CoverTab[116411]++
															uints = dataVal.Interface().([]uint8)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:579
				// _ = "end of CoverTab[116411]"
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:580
			// _ = "end of CoverTab[116406]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:580
			_go_fuzz_dep_.CoverTab[116407]++
														val.SetString(string(uints))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:581
			// _ = "end of CoverTab[116407]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:582
			_go_fuzz_dep_.CoverTab[116408]++
														converted = false
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:583
			// _ = "end of CoverTab[116408]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:584
		// _ = "end of CoverTab[116396]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:585
		_go_fuzz_dep_.CoverTab[116397]++
													converted = false
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:586
		// _ = "end of CoverTab[116397]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:587
	// _ = "end of CoverTab[116388]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:587
	_go_fuzz_dep_.CoverTab[116389]++

												if !converted {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:589
		_go_fuzz_dep_.CoverTab[116412]++
													return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'",
			name, val.Type(), dataVal.Type(), data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:592
		// _ = "end of CoverTab[116412]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:593
		_go_fuzz_dep_.CoverTab[116413]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:593
		// _ = "end of CoverTab[116413]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:593
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:593
	// _ = "end of CoverTab[116389]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:593
	_go_fuzz_dep_.CoverTab[116390]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:595
	// _ = "end of CoverTab[116390]"
}

func (d *Decoder) decodeInt(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:598
	_go_fuzz_dep_.CoverTab[116414]++
												dataVal := reflect.Indirect(reflect.ValueOf(data))
												dataKind := getKind(dataVal)
												dataType := dataVal.Type()

												switch {
	case dataKind == reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:604
		_go_fuzz_dep_.CoverTab[116416]++
													val.SetInt(dataVal.Int())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:605
		// _ = "end of CoverTab[116416]"
	case dataKind == reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:606
		_go_fuzz_dep_.CoverTab[116417]++
													val.SetInt(int64(dataVal.Uint()))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:607
		// _ = "end of CoverTab[116417]"
	case dataKind == reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:608
		_go_fuzz_dep_.CoverTab[116418]++
													val.SetInt(int64(dataVal.Float()))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:609
		// _ = "end of CoverTab[116418]"
	case dataKind == reflect.Bool && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:610
		_go_fuzz_dep_.CoverTab[116425]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:610
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:610
		// _ = "end of CoverTab[116425]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:610
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:610
		_go_fuzz_dep_.CoverTab[116419]++
													if dataVal.Bool() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:611
			_go_fuzz_dep_.CoverTab[116426]++
														val.SetInt(1)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:612
			// _ = "end of CoverTab[116426]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:613
			_go_fuzz_dep_.CoverTab[116427]++
														val.SetInt(0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:614
			// _ = "end of CoverTab[116427]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:615
		// _ = "end of CoverTab[116419]"
	case dataKind == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:616
		_go_fuzz_dep_.CoverTab[116428]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:616
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:616
		// _ = "end of CoverTab[116428]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:616
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:616
		_go_fuzz_dep_.CoverTab[116420]++
													str := dataVal.String()
													if str == "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:618
			_go_fuzz_dep_.CoverTab[116429]++
														str = "0"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:619
			// _ = "end of CoverTab[116429]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:620
			_go_fuzz_dep_.CoverTab[116430]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:620
			// _ = "end of CoverTab[116430]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:620
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:620
		// _ = "end of CoverTab[116420]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:620
		_go_fuzz_dep_.CoverTab[116421]++

													i, err := strconv.ParseInt(str, 0, val.Type().Bits())
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:623
			_go_fuzz_dep_.CoverTab[116431]++
														val.SetInt(i)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:624
			// _ = "end of CoverTab[116431]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:625
			_go_fuzz_dep_.CoverTab[116432]++
														return fmt.Errorf("cannot parse '%s' as int: %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:626
			// _ = "end of CoverTab[116432]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:627
		// _ = "end of CoverTab[116421]"
	case dataType.PkgPath() == "encoding/json" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:628
		_go_fuzz_dep_.CoverTab[116433]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:628
		return dataType.Name() == "Number"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:628
		// _ = "end of CoverTab[116433]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:628
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:628
		_go_fuzz_dep_.CoverTab[116422]++
													jn := data.(json.Number)
													i, err := jn.Int64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:631
			_go_fuzz_dep_.CoverTab[116434]++
														return fmt.Errorf(
				"error decoding json.Number into %s: %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:633
			// _ = "end of CoverTab[116434]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:634
			_go_fuzz_dep_.CoverTab[116435]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:634
			// _ = "end of CoverTab[116435]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:634
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:634
		// _ = "end of CoverTab[116422]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:634
		_go_fuzz_dep_.CoverTab[116423]++
													val.SetInt(i)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:635
		// _ = "end of CoverTab[116423]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:636
		_go_fuzz_dep_.CoverTab[116424]++
													return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'",
			name, val.Type(), dataVal.Type(), data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:639
		// _ = "end of CoverTab[116424]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:640
	// _ = "end of CoverTab[116414]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:640
	_go_fuzz_dep_.CoverTab[116415]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:642
	// _ = "end of CoverTab[116415]"
}

func (d *Decoder) decodeUint(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:645
	_go_fuzz_dep_.CoverTab[116436]++
												dataVal := reflect.Indirect(reflect.ValueOf(data))
												dataKind := getKind(dataVal)
												dataType := dataVal.Type()

												switch {
	case dataKind == reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:651
		_go_fuzz_dep_.CoverTab[116438]++
													i := dataVal.Int()
													if i < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:653
			_go_fuzz_dep_.CoverTab[116450]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:653
			return !d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:653
			// _ = "end of CoverTab[116450]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:653
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:653
			_go_fuzz_dep_.CoverTab[116451]++
														return fmt.Errorf("cannot parse '%s', %d overflows uint",
				name, i)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:655
			// _ = "end of CoverTab[116451]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:656
			_go_fuzz_dep_.CoverTab[116452]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:656
			// _ = "end of CoverTab[116452]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:656
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:656
		// _ = "end of CoverTab[116438]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:656
		_go_fuzz_dep_.CoverTab[116439]++
													val.SetUint(uint64(i))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:657
		// _ = "end of CoverTab[116439]"
	case dataKind == reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:658
		_go_fuzz_dep_.CoverTab[116440]++
													val.SetUint(dataVal.Uint())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:659
		// _ = "end of CoverTab[116440]"
	case dataKind == reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:660
		_go_fuzz_dep_.CoverTab[116441]++
													f := dataVal.Float()
													if f < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:662
			_go_fuzz_dep_.CoverTab[116453]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:662
			return !d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:662
			// _ = "end of CoverTab[116453]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:662
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:662
			_go_fuzz_dep_.CoverTab[116454]++
														return fmt.Errorf("cannot parse '%s', %f overflows uint",
				name, f)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:664
			// _ = "end of CoverTab[116454]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:665
			_go_fuzz_dep_.CoverTab[116455]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:665
			// _ = "end of CoverTab[116455]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:665
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:665
		// _ = "end of CoverTab[116441]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:665
		_go_fuzz_dep_.CoverTab[116442]++
													val.SetUint(uint64(f))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:666
		// _ = "end of CoverTab[116442]"
	case dataKind == reflect.Bool && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:667
		_go_fuzz_dep_.CoverTab[116456]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:667
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:667
		// _ = "end of CoverTab[116456]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:667
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:667
		_go_fuzz_dep_.CoverTab[116443]++
													if dataVal.Bool() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:668
			_go_fuzz_dep_.CoverTab[116457]++
														val.SetUint(1)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:669
			// _ = "end of CoverTab[116457]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:670
			_go_fuzz_dep_.CoverTab[116458]++
														val.SetUint(0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:671
			// _ = "end of CoverTab[116458]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:672
		// _ = "end of CoverTab[116443]"
	case dataKind == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:673
		_go_fuzz_dep_.CoverTab[116459]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:673
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:673
		// _ = "end of CoverTab[116459]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:673
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:673
		_go_fuzz_dep_.CoverTab[116444]++
													str := dataVal.String()
													if str == "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:675
			_go_fuzz_dep_.CoverTab[116460]++
														str = "0"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:676
			// _ = "end of CoverTab[116460]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:677
			_go_fuzz_dep_.CoverTab[116461]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:677
			// _ = "end of CoverTab[116461]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:677
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:677
		// _ = "end of CoverTab[116444]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:677
		_go_fuzz_dep_.CoverTab[116445]++

													i, err := strconv.ParseUint(str, 0, val.Type().Bits())
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:680
			_go_fuzz_dep_.CoverTab[116462]++
														val.SetUint(i)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:681
			// _ = "end of CoverTab[116462]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:682
			_go_fuzz_dep_.CoverTab[116463]++
														return fmt.Errorf("cannot parse '%s' as uint: %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:683
			// _ = "end of CoverTab[116463]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:684
		// _ = "end of CoverTab[116445]"
	case dataType.PkgPath() == "encoding/json" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:685
		_go_fuzz_dep_.CoverTab[116464]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:685
		return dataType.Name() == "Number"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:685
		// _ = "end of CoverTab[116464]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:685
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:685
		_go_fuzz_dep_.CoverTab[116446]++
													jn := data.(json.Number)
													i, err := jn.Int64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:688
			_go_fuzz_dep_.CoverTab[116465]++
														return fmt.Errorf(
				"error decoding json.Number into %s: %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:690
			// _ = "end of CoverTab[116465]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:691
			_go_fuzz_dep_.CoverTab[116466]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:691
			// _ = "end of CoverTab[116466]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:691
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:691
		// _ = "end of CoverTab[116446]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:691
		_go_fuzz_dep_.CoverTab[116447]++
													if i < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:692
			_go_fuzz_dep_.CoverTab[116467]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:692
			return !d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:692
			// _ = "end of CoverTab[116467]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:692
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:692
			_go_fuzz_dep_.CoverTab[116468]++
														return fmt.Errorf("cannot parse '%s', %d overflows uint",
				name, i)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:694
			// _ = "end of CoverTab[116468]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:695
			_go_fuzz_dep_.CoverTab[116469]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:695
			// _ = "end of CoverTab[116469]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:695
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:695
		// _ = "end of CoverTab[116447]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:695
		_go_fuzz_dep_.CoverTab[116448]++
													val.SetUint(uint64(i))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:696
		// _ = "end of CoverTab[116448]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:697
		_go_fuzz_dep_.CoverTab[116449]++
													return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'",
			name, val.Type(), dataVal.Type(), data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:700
		// _ = "end of CoverTab[116449]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:701
	// _ = "end of CoverTab[116436]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:701
	_go_fuzz_dep_.CoverTab[116437]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:703
	// _ = "end of CoverTab[116437]"
}

func (d *Decoder) decodeBool(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:706
	_go_fuzz_dep_.CoverTab[116470]++
												dataVal := reflect.Indirect(reflect.ValueOf(data))
												dataKind := getKind(dataVal)

												switch {
	case dataKind == reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:711
		_go_fuzz_dep_.CoverTab[116472]++
													val.SetBool(dataVal.Bool())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:712
		// _ = "end of CoverTab[116472]"
	case dataKind == reflect.Int && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:713
		_go_fuzz_dep_.CoverTab[116478]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:713
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:713
		// _ = "end of CoverTab[116478]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:713
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:713
		_go_fuzz_dep_.CoverTab[116473]++
													val.SetBool(dataVal.Int() != 0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:714
		// _ = "end of CoverTab[116473]"
	case dataKind == reflect.Uint && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:715
		_go_fuzz_dep_.CoverTab[116479]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:715
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:715
		// _ = "end of CoverTab[116479]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:715
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:715
		_go_fuzz_dep_.CoverTab[116474]++
													val.SetBool(dataVal.Uint() != 0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:716
		// _ = "end of CoverTab[116474]"
	case dataKind == reflect.Float32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:717
		_go_fuzz_dep_.CoverTab[116480]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:717
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:717
		// _ = "end of CoverTab[116480]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:717
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:717
		_go_fuzz_dep_.CoverTab[116475]++
													val.SetBool(dataVal.Float() != 0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:718
		// _ = "end of CoverTab[116475]"
	case dataKind == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:719
		_go_fuzz_dep_.CoverTab[116481]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:719
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:719
		// _ = "end of CoverTab[116481]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:719
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:719
		_go_fuzz_dep_.CoverTab[116476]++
													b, err := strconv.ParseBool(dataVal.String())
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:721
			_go_fuzz_dep_.CoverTab[116482]++
														val.SetBool(b)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:722
			// _ = "end of CoverTab[116482]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:723
			_go_fuzz_dep_.CoverTab[116483]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:723
			if dataVal.String() == "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:723
				_go_fuzz_dep_.CoverTab[116484]++
															val.SetBool(false)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:724
				// _ = "end of CoverTab[116484]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:725
				_go_fuzz_dep_.CoverTab[116485]++
															return fmt.Errorf("cannot parse '%s' as bool: %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:726
				// _ = "end of CoverTab[116485]"
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:727
			// _ = "end of CoverTab[116483]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:727
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:727
		// _ = "end of CoverTab[116476]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:728
		_go_fuzz_dep_.CoverTab[116477]++
													return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'",
			name, val.Type(), dataVal.Type(), data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:731
		// _ = "end of CoverTab[116477]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:732
	// _ = "end of CoverTab[116470]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:732
	_go_fuzz_dep_.CoverTab[116471]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:734
	// _ = "end of CoverTab[116471]"
}

func (d *Decoder) decodeFloat(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:737
	_go_fuzz_dep_.CoverTab[116486]++
												dataVal := reflect.Indirect(reflect.ValueOf(data))
												dataKind := getKind(dataVal)
												dataType := dataVal.Type()

												switch {
	case dataKind == reflect.Int:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:743
		_go_fuzz_dep_.CoverTab[116488]++
													val.SetFloat(float64(dataVal.Int()))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:744
		// _ = "end of CoverTab[116488]"
	case dataKind == reflect.Uint:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:745
		_go_fuzz_dep_.CoverTab[116489]++
													val.SetFloat(float64(dataVal.Uint()))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:746
		// _ = "end of CoverTab[116489]"
	case dataKind == reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:747
		_go_fuzz_dep_.CoverTab[116490]++
													val.SetFloat(dataVal.Float())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:748
		// _ = "end of CoverTab[116490]"
	case dataKind == reflect.Bool && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:749
		_go_fuzz_dep_.CoverTab[116497]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:749
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:749
		// _ = "end of CoverTab[116497]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:749
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:749
		_go_fuzz_dep_.CoverTab[116491]++
													if dataVal.Bool() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:750
			_go_fuzz_dep_.CoverTab[116498]++
														val.SetFloat(1)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:751
			// _ = "end of CoverTab[116498]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:752
			_go_fuzz_dep_.CoverTab[116499]++
														val.SetFloat(0)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:753
			// _ = "end of CoverTab[116499]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:754
		// _ = "end of CoverTab[116491]"
	case dataKind == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:755
		_go_fuzz_dep_.CoverTab[116500]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:755
		return d.config.WeaklyTypedInput
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:755
		// _ = "end of CoverTab[116500]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:755
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:755
		_go_fuzz_dep_.CoverTab[116492]++
													str := dataVal.String()
													if str == "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:757
			_go_fuzz_dep_.CoverTab[116501]++
														str = "0"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:758
			// _ = "end of CoverTab[116501]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:759
			_go_fuzz_dep_.CoverTab[116502]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:759
			// _ = "end of CoverTab[116502]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:759
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:759
		// _ = "end of CoverTab[116492]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:759
		_go_fuzz_dep_.CoverTab[116493]++

													f, err := strconv.ParseFloat(str, val.Type().Bits())
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:762
			_go_fuzz_dep_.CoverTab[116503]++
														val.SetFloat(f)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:763
			// _ = "end of CoverTab[116503]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:764
			_go_fuzz_dep_.CoverTab[116504]++
														return fmt.Errorf("cannot parse '%s' as float: %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:765
			// _ = "end of CoverTab[116504]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:766
		// _ = "end of CoverTab[116493]"
	case dataType.PkgPath() == "encoding/json" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:767
		_go_fuzz_dep_.CoverTab[116505]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:767
		return dataType.Name() == "Number"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:767
		// _ = "end of CoverTab[116505]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:767
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:767
		_go_fuzz_dep_.CoverTab[116494]++
													jn := data.(json.Number)
													i, err := jn.Float64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:770
			_go_fuzz_dep_.CoverTab[116506]++
														return fmt.Errorf(
				"error decoding json.Number into %s: %s", name, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:772
			// _ = "end of CoverTab[116506]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:773
			_go_fuzz_dep_.CoverTab[116507]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:773
			// _ = "end of CoverTab[116507]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:773
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:773
		// _ = "end of CoverTab[116494]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:773
		_go_fuzz_dep_.CoverTab[116495]++
													val.SetFloat(i)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:774
		// _ = "end of CoverTab[116495]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:775
		_go_fuzz_dep_.CoverTab[116496]++
													return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'",
			name, val.Type(), dataVal.Type(), data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:778
		// _ = "end of CoverTab[116496]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:779
	// _ = "end of CoverTab[116486]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:779
	_go_fuzz_dep_.CoverTab[116487]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:781
	// _ = "end of CoverTab[116487]"
}

func (d *Decoder) decodeMap(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:784
	_go_fuzz_dep_.CoverTab[116508]++
												valType := val.Type()
												valKeyType := valType.Key()
												valElemType := valType.Elem()

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:790
	valMap := val

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:793
	if valMap.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:793
		_go_fuzz_dep_.CoverTab[116510]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:793
		return d.config.ZeroFields
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:793
		// _ = "end of CoverTab[116510]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:793
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:793
		_go_fuzz_dep_.CoverTab[116511]++

													mapType := reflect.MapOf(valKeyType, valElemType)
													valMap = reflect.MakeMap(mapType)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:796
		// _ = "end of CoverTab[116511]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:797
		_go_fuzz_dep_.CoverTab[116512]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:797
		// _ = "end of CoverTab[116512]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:797
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:797
	// _ = "end of CoverTab[116508]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:797
	_go_fuzz_dep_.CoverTab[116509]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:800
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	switch dataVal.Kind() {
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:802
		_go_fuzz_dep_.CoverTab[116513]++
													return d.decodeMapFromMap(name, dataVal, val, valMap)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:803
		// _ = "end of CoverTab[116513]"

	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:805
		_go_fuzz_dep_.CoverTab[116514]++
													return d.decodeMapFromStruct(name, dataVal, val, valMap)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:806
		// _ = "end of CoverTab[116514]"

	case reflect.Array, reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:808
		_go_fuzz_dep_.CoverTab[116515]++
													if d.config.WeaklyTypedInput {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:809
			_go_fuzz_dep_.CoverTab[116518]++
														return d.decodeMapFromSlice(name, dataVal, val, valMap)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:810
			// _ = "end of CoverTab[116518]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:811
			_go_fuzz_dep_.CoverTab[116519]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:811
			// _ = "end of CoverTab[116519]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:811
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:811
		// _ = "end of CoverTab[116515]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:811
		_go_fuzz_dep_.CoverTab[116516]++

													fallthrough
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:813
		// _ = "end of CoverTab[116516]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:815
		_go_fuzz_dep_.CoverTab[116517]++
													return fmt.Errorf("'%s' expected a map, got '%s'", name, dataVal.Kind())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:816
		// _ = "end of CoverTab[116517]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:817
	// _ = "end of CoverTab[116509]"
}

func (d *Decoder) decodeMapFromSlice(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:820
	_go_fuzz_dep_.CoverTab[116520]++

												if dataVal.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:822
		_go_fuzz_dep_.CoverTab[116523]++
													val.Set(valMap)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:824
		// _ = "end of CoverTab[116523]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:825
		_go_fuzz_dep_.CoverTab[116524]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:825
		// _ = "end of CoverTab[116524]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:825
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:825
	// _ = "end of CoverTab[116520]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:825
	_go_fuzz_dep_.CoverTab[116521]++

												for i := 0; i < dataVal.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:827
		_go_fuzz_dep_.CoverTab[116525]++
													err := d.decode(
			name+"["+strconv.Itoa(i)+"]",
			dataVal.Index(i).Interface(), val)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:831
			_go_fuzz_dep_.CoverTab[116526]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:832
			// _ = "end of CoverTab[116526]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:833
			_go_fuzz_dep_.CoverTab[116527]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:833
			// _ = "end of CoverTab[116527]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:833
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:833
		// _ = "end of CoverTab[116525]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:834
	// _ = "end of CoverTab[116521]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:834
	_go_fuzz_dep_.CoverTab[116522]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:836
	// _ = "end of CoverTab[116522]"
}

func (d *Decoder) decodeMapFromMap(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:839
	_go_fuzz_dep_.CoverTab[116528]++
												valType := val.Type()
												valKeyType := valType.Key()
												valElemType := valType.Elem()

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:845
	errors := make([]string, 0)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:848
	if dataVal.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:848
		_go_fuzz_dep_.CoverTab[116532]++
													if dataVal.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:849
			_go_fuzz_dep_.CoverTab[116534]++
														if !val.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:850
				_go_fuzz_dep_.CoverTab[116535]++
															val.Set(dataVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:851
				// _ = "end of CoverTab[116535]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:852
				_go_fuzz_dep_.CoverTab[116536]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:852
				// _ = "end of CoverTab[116536]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:852
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:852
			// _ = "end of CoverTab[116534]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:853
			_go_fuzz_dep_.CoverTab[116537]++

														val.Set(valMap)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:855
			// _ = "end of CoverTab[116537]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:856
		// _ = "end of CoverTab[116532]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:856
		_go_fuzz_dep_.CoverTab[116533]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:858
		// _ = "end of CoverTab[116533]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:859
		_go_fuzz_dep_.CoverTab[116538]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:859
		// _ = "end of CoverTab[116538]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:859
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:859
	// _ = "end of CoverTab[116528]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:859
	_go_fuzz_dep_.CoverTab[116529]++

												for _, k := range dataVal.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:861
		_go_fuzz_dep_.CoverTab[116539]++
													fieldName := name + "[" + k.String() + "]"

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:865
		currentKey := reflect.Indirect(reflect.New(valKeyType))
		if err := d.decode(fieldName, k.Interface(), currentKey); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:866
			_go_fuzz_dep_.CoverTab[116542]++
														errors = appendErrors(errors, err)
														continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:868
			// _ = "end of CoverTab[116542]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:869
			_go_fuzz_dep_.CoverTab[116543]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:869
			// _ = "end of CoverTab[116543]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:869
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:869
		// _ = "end of CoverTab[116539]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:869
		_go_fuzz_dep_.CoverTab[116540]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:872
		v := dataVal.MapIndex(k).Interface()
		currentVal := reflect.Indirect(reflect.New(valElemType))
		if err := d.decode(fieldName, v, currentVal); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:874
			_go_fuzz_dep_.CoverTab[116544]++
														errors = appendErrors(errors, err)
														continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:876
			// _ = "end of CoverTab[116544]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:877
			_go_fuzz_dep_.CoverTab[116545]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:877
			// _ = "end of CoverTab[116545]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:877
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:877
		// _ = "end of CoverTab[116540]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:877
		_go_fuzz_dep_.CoverTab[116541]++

													valMap.SetMapIndex(currentKey, currentVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:879
		// _ = "end of CoverTab[116541]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:880
	// _ = "end of CoverTab[116529]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:880
	_go_fuzz_dep_.CoverTab[116530]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:883
	val.Set(valMap)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:886
	if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:886
		_go_fuzz_dep_.CoverTab[116546]++
													return &Error{errors}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:887
		// _ = "end of CoverTab[116546]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:888
		_go_fuzz_dep_.CoverTab[116547]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:888
		// _ = "end of CoverTab[116547]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:888
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:888
	// _ = "end of CoverTab[116530]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:888
	_go_fuzz_dep_.CoverTab[116531]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:890
	// _ = "end of CoverTab[116531]"
}

func (d *Decoder) decodeMapFromStruct(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:893
	_go_fuzz_dep_.CoverTab[116548]++
												typ := dataVal.Type()
												for i := 0; i < typ.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:895
		_go_fuzz_dep_.CoverTab[116551]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:898
		f := typ.Field(i)
		if f.PkgPath != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:899
			_go_fuzz_dep_.CoverTab[116555]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:900
			// _ = "end of CoverTab[116555]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:901
			_go_fuzz_dep_.CoverTab[116556]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:901
			// _ = "end of CoverTab[116556]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:901
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:901
		// _ = "end of CoverTab[116551]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:901
		_go_fuzz_dep_.CoverTab[116552]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:905
		v := dataVal.Field(i)
		if !v.Type().AssignableTo(valMap.Type().Elem()) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:906
			_go_fuzz_dep_.CoverTab[116557]++
														return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", v.Type(), valMap.Type().Elem())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:907
			// _ = "end of CoverTab[116557]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:908
			_go_fuzz_dep_.CoverTab[116558]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:908
			// _ = "end of CoverTab[116558]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:908
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:908
		// _ = "end of CoverTab[116552]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:908
		_go_fuzz_dep_.CoverTab[116553]++

													tagValue := f.Tag.Get(d.config.TagName)
													keyName := f.Name

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
		squash := d.config.Squash && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
			_go_fuzz_dep_.CoverTab[116559]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
			return v.Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
			// _ = "end of CoverTab[116559]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
			_go_fuzz_dep_.CoverTab[116560]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
			return f.Anonymous
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
			// _ = "end of CoverTab[116560]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:914
		}()

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:917
		if index := strings.Index(tagValue, ","); index != -1 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:917
			_go_fuzz_dep_.CoverTab[116561]++
														if tagValue[:index] == "-" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:918
				_go_fuzz_dep_.CoverTab[116565]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:919
				// _ = "end of CoverTab[116565]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:920
				_go_fuzz_dep_.CoverTab[116566]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:920
				// _ = "end of CoverTab[116566]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:920
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:920
			// _ = "end of CoverTab[116561]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:920
			_go_fuzz_dep_.CoverTab[116562]++

														if strings.Index(tagValue[index+1:], "omitempty") != -1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:922
				_go_fuzz_dep_.CoverTab[116567]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:922
				return isEmptyValue(v)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:922
				// _ = "end of CoverTab[116567]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:922
			}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:922
				_go_fuzz_dep_.CoverTab[116568]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:923
				// _ = "end of CoverTab[116568]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:924
				_go_fuzz_dep_.CoverTab[116569]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:924
				// _ = "end of CoverTab[116569]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:924
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:924
			// _ = "end of CoverTab[116562]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:924
			_go_fuzz_dep_.CoverTab[116563]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:927
			squash = !squash && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:927
				_go_fuzz_dep_.CoverTab[116570]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:927
				return strings.Index(tagValue[index+1:], "squash") != -1
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:927
				// _ = "end of CoverTab[116570]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:927
			}()
			if squash {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:928
				_go_fuzz_dep_.CoverTab[116571]++

															if v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:930
					_go_fuzz_dep_.CoverTab[116573]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:930
					return v.Elem().Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:930
					// _ = "end of CoverTab[116573]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:930
				}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:930
					_go_fuzz_dep_.CoverTab[116574]++
																v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:931
					// _ = "end of CoverTab[116574]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:932
					_go_fuzz_dep_.CoverTab[116575]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:932
					// _ = "end of CoverTab[116575]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:932
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:932
				// _ = "end of CoverTab[116571]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:932
				_go_fuzz_dep_.CoverTab[116572]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:935
				if v.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:935
					_go_fuzz_dep_.CoverTab[116576]++
																return fmt.Errorf("cannot squash non-struct type '%s'", v.Type())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:936
					// _ = "end of CoverTab[116576]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:937
					_go_fuzz_dep_.CoverTab[116577]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:937
					// _ = "end of CoverTab[116577]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:937
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:937
				// _ = "end of CoverTab[116572]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:938
				_go_fuzz_dep_.CoverTab[116578]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:938
				// _ = "end of CoverTab[116578]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:938
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:938
			// _ = "end of CoverTab[116563]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:938
			_go_fuzz_dep_.CoverTab[116564]++
														keyName = tagValue[:index]
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:939
			// _ = "end of CoverTab[116564]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:940
			_go_fuzz_dep_.CoverTab[116579]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:940
			if len(tagValue) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:940
				_go_fuzz_dep_.CoverTab[116580]++
															if tagValue == "-" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:941
					_go_fuzz_dep_.CoverTab[116582]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:942
					// _ = "end of CoverTab[116582]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:943
					_go_fuzz_dep_.CoverTab[116583]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:943
					// _ = "end of CoverTab[116583]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:943
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:943
				// _ = "end of CoverTab[116580]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:943
				_go_fuzz_dep_.CoverTab[116581]++
															keyName = tagValue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:944
				// _ = "end of CoverTab[116581]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:945
				_go_fuzz_dep_.CoverTab[116584]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:945
				// _ = "end of CoverTab[116584]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:945
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:945
			// _ = "end of CoverTab[116579]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:945
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:945
		// _ = "end of CoverTab[116553]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:945
		_go_fuzz_dep_.CoverTab[116554]++

													switch v.Kind() {

		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:949
			_go_fuzz_dep_.CoverTab[116585]++
														x := reflect.New(v.Type())
														x.Elem().Set(v)

														vType := valMap.Type()
														vKeyType := vType.Key()
														vElemType := vType.Elem()
														mType := reflect.MapOf(vKeyType, vElemType)
														vMap := reflect.MakeMap(mType)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:963
			addrVal := reflect.New(vMap.Type())
			reflect.Indirect(addrVal).Set(vMap)

			err := d.decode(keyName, x.Interface(), reflect.Indirect(addrVal))
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:967
				_go_fuzz_dep_.CoverTab[116588]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:968
				// _ = "end of CoverTab[116588]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:969
				_go_fuzz_dep_.CoverTab[116589]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:969
				// _ = "end of CoverTab[116589]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:969
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:969
			// _ = "end of CoverTab[116585]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:969
			_go_fuzz_dep_.CoverTab[116586]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:973
			vMap = reflect.Indirect(addrVal)

			if squash {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:975
				_go_fuzz_dep_.CoverTab[116590]++
															for _, k := range vMap.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:976
					_go_fuzz_dep_.CoverTab[116591]++
																valMap.SetMapIndex(k, vMap.MapIndex(k))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:977
					// _ = "end of CoverTab[116591]"
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:978
				// _ = "end of CoverTab[116590]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:979
				_go_fuzz_dep_.CoverTab[116592]++
															valMap.SetMapIndex(reflect.ValueOf(keyName), vMap)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:980
				// _ = "end of CoverTab[116592]"
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:981
			// _ = "end of CoverTab[116586]"

		default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:983
			_go_fuzz_dep_.CoverTab[116587]++
														valMap.SetMapIndex(reflect.ValueOf(keyName), v)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:984
			// _ = "end of CoverTab[116587]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:985
		// _ = "end of CoverTab[116554]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:986
	// _ = "end of CoverTab[116548]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:986
	_go_fuzz_dep_.CoverTab[116549]++

												if val.CanAddr() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:988
		_go_fuzz_dep_.CoverTab[116593]++
													val.Set(valMap)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:989
		// _ = "end of CoverTab[116593]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:990
		_go_fuzz_dep_.CoverTab[116594]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:990
		// _ = "end of CoverTab[116594]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:990
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:990
	// _ = "end of CoverTab[116549]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:990
	_go_fuzz_dep_.CoverTab[116550]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:992
	// _ = "end of CoverTab[116550]"
}

func (d *Decoder) decodePtr(name string, data interface{}, val reflect.Value) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:995
	_go_fuzz_dep_.CoverTab[116595]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:998
	isNil := data == nil
	if !isNil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:999
		_go_fuzz_dep_.CoverTab[116599]++
													switch v := reflect.Indirect(reflect.ValueOf(data)); v.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1006
			_go_fuzz_dep_.CoverTab[116600]++
															isNil = v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1007
			// _ = "end of CoverTab[116600]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1007
		default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1007
			_go_fuzz_dep_.CoverTab[116601]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1007
			// _ = "end of CoverTab[116601]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1008
		// _ = "end of CoverTab[116599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1009
		_go_fuzz_dep_.CoverTab[116602]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1009
		// _ = "end of CoverTab[116602]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1009
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1009
	// _ = "end of CoverTab[116595]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1009
	_go_fuzz_dep_.CoverTab[116596]++
													if isNil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1010
		_go_fuzz_dep_.CoverTab[116603]++
														if !val.IsNil() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1011
			_go_fuzz_dep_.CoverTab[116605]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1011
			return val.CanSet()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1011
			// _ = "end of CoverTab[116605]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1011
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1011
			_go_fuzz_dep_.CoverTab[116606]++
															nilValue := reflect.New(val.Type()).Elem()
															val.Set(nilValue)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1013
			// _ = "end of CoverTab[116606]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1014
			_go_fuzz_dep_.CoverTab[116607]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1014
			// _ = "end of CoverTab[116607]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1014
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1014
		// _ = "end of CoverTab[116603]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1014
		_go_fuzz_dep_.CoverTab[116604]++

														return true, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1016
		// _ = "end of CoverTab[116604]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1017
		_go_fuzz_dep_.CoverTab[116608]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1017
		// _ = "end of CoverTab[116608]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1017
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1017
	// _ = "end of CoverTab[116596]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1017
	_go_fuzz_dep_.CoverTab[116597]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1021
	valType := val.Type()
	valElemType := valType.Elem()
	if val.CanSet() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1023
		_go_fuzz_dep_.CoverTab[116609]++
														realVal := val
														if realVal.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1025
			_go_fuzz_dep_.CoverTab[116612]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1025
			return d.config.ZeroFields
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1025
			// _ = "end of CoverTab[116612]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1025
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1025
			_go_fuzz_dep_.CoverTab[116613]++
															realVal = reflect.New(valElemType)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1026
			// _ = "end of CoverTab[116613]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1027
			_go_fuzz_dep_.CoverTab[116614]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1027
			// _ = "end of CoverTab[116614]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1027
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1027
		// _ = "end of CoverTab[116609]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1027
		_go_fuzz_dep_.CoverTab[116610]++

														if err := d.decode(name, data, reflect.Indirect(realVal)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1029
			_go_fuzz_dep_.CoverTab[116615]++
															return false, err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1030
			// _ = "end of CoverTab[116615]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1031
			_go_fuzz_dep_.CoverTab[116616]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1031
			// _ = "end of CoverTab[116616]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1031
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1031
		// _ = "end of CoverTab[116610]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1031
		_go_fuzz_dep_.CoverTab[116611]++

														val.Set(realVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1033
		// _ = "end of CoverTab[116611]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1034
		_go_fuzz_dep_.CoverTab[116617]++
														if err := d.decode(name, data, reflect.Indirect(val)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1035
			_go_fuzz_dep_.CoverTab[116618]++
															return false, err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1036
			// _ = "end of CoverTab[116618]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1037
			_go_fuzz_dep_.CoverTab[116619]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1037
			// _ = "end of CoverTab[116619]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1037
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1037
		// _ = "end of CoverTab[116617]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1038
	// _ = "end of CoverTab[116597]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1038
	_go_fuzz_dep_.CoverTab[116598]++
													return false, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1039
	// _ = "end of CoverTab[116598]"
}

func (d *Decoder) decodeFunc(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1042
	_go_fuzz_dep_.CoverTab[116620]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1045
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	if val.Type() != dataVal.Type() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1046
		_go_fuzz_dep_.CoverTab[116622]++
														return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'",
			name, val.Type(), dataVal.Type(), data)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1049
		// _ = "end of CoverTab[116622]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1050
		_go_fuzz_dep_.CoverTab[116623]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1050
		// _ = "end of CoverTab[116623]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1050
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1050
	// _ = "end of CoverTab[116620]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1050
	_go_fuzz_dep_.CoverTab[116621]++
													val.Set(dataVal)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1052
	// _ = "end of CoverTab[116621]"
}

func (d *Decoder) decodeSlice(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1055
	_go_fuzz_dep_.CoverTab[116624]++
													dataVal := reflect.Indirect(reflect.ValueOf(data))
													dataValKind := dataVal.Kind()
													valType := val.Type()
													valElemType := valType.Elem()
													sliceType := reflect.SliceOf(valElemType)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1063
	if dataValKind != reflect.Array && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1063
		_go_fuzz_dep_.CoverTab[116630]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1063
		return dataValKind != reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1063
		// _ = "end of CoverTab[116630]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1063
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1063
		_go_fuzz_dep_.CoverTab[116631]++
														if d.config.WeaklyTypedInput {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1064
			_go_fuzz_dep_.CoverTab[116633]++
															switch {

			case dataValKind == reflect.Slice, dataValKind == reflect.Array:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1067
				_go_fuzz_dep_.CoverTab[116634]++
																break
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1068
				// _ = "end of CoverTab[116634]"

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1071
			case dataValKind == reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1071
				_go_fuzz_dep_.CoverTab[116635]++
																if dataVal.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1072
					_go_fuzz_dep_.CoverTab[116639]++
																	val.Set(reflect.MakeSlice(sliceType, 0, 0))
																	return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1074
					// _ = "end of CoverTab[116639]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1075
					_go_fuzz_dep_.CoverTab[116640]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1075
					// _ = "end of CoverTab[116640]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1075
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1075
				// _ = "end of CoverTab[116635]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1075
				_go_fuzz_dep_.CoverTab[116636]++

																return d.decodeSlice(name, []interface{}{data}, val)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1077
				// _ = "end of CoverTab[116636]"

			case dataValKind == reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1079
				_go_fuzz_dep_.CoverTab[116641]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1079
				return valElemType.Kind() == reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1079
				// _ = "end of CoverTab[116641]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1079
			}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1079
				_go_fuzz_dep_.CoverTab[116637]++
																return d.decodeSlice(name, []byte(dataVal.String()), val)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1080
				// _ = "end of CoverTab[116637]"

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1084
			default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1084
				_go_fuzz_dep_.CoverTab[116638]++

																return d.decodeSlice(name, []interface{}{data}, val)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1086
				// _ = "end of CoverTab[116638]"
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1087
			// _ = "end of CoverTab[116633]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1088
			_go_fuzz_dep_.CoverTab[116642]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1088
			// _ = "end of CoverTab[116642]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1088
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1088
		// _ = "end of CoverTab[116631]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1088
		_go_fuzz_dep_.CoverTab[116632]++

														return fmt.Errorf(
			"'%s': source data must be an array or slice, got %s", name, dataValKind)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1091
		// _ = "end of CoverTab[116632]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1092
		_go_fuzz_dep_.CoverTab[116643]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1092
		// _ = "end of CoverTab[116643]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1092
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1092
	// _ = "end of CoverTab[116624]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1092
	_go_fuzz_dep_.CoverTab[116625]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1095
	if dataVal.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1095
		_go_fuzz_dep_.CoverTab[116644]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1096
		// _ = "end of CoverTab[116644]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1097
		_go_fuzz_dep_.CoverTab[116645]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1097
		// _ = "end of CoverTab[116645]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1097
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1097
	// _ = "end of CoverTab[116625]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1097
	_go_fuzz_dep_.CoverTab[116626]++

													valSlice := val
													if valSlice.IsNil() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1100
		_go_fuzz_dep_.CoverTab[116646]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1100
		return d.config.ZeroFields
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1100
		// _ = "end of CoverTab[116646]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1100
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1100
		_go_fuzz_dep_.CoverTab[116647]++

														valSlice = reflect.MakeSlice(sliceType, dataVal.Len(), dataVal.Len())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1102
		// _ = "end of CoverTab[116647]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1103
		_go_fuzz_dep_.CoverTab[116648]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1103
		// _ = "end of CoverTab[116648]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1103
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1103
	// _ = "end of CoverTab[116626]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1103
	_go_fuzz_dep_.CoverTab[116627]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1106
	errors := make([]string, 0)

	for i := 0; i < dataVal.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1108
		_go_fuzz_dep_.CoverTab[116649]++
														currentData := dataVal.Index(i).Interface()
														for valSlice.Len() <= i {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1110
			_go_fuzz_dep_.CoverTab[116651]++
															valSlice = reflect.Append(valSlice, reflect.Zero(valElemType))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1111
			// _ = "end of CoverTab[116651]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1112
		// _ = "end of CoverTab[116649]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1112
		_go_fuzz_dep_.CoverTab[116650]++
														currentField := valSlice.Index(i)

														fieldName := name + "[" + strconv.Itoa(i) + "]"
														if err := d.decode(fieldName, currentData, currentField); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1116
			_go_fuzz_dep_.CoverTab[116652]++
															errors = appendErrors(errors, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1117
			// _ = "end of CoverTab[116652]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1118
			_go_fuzz_dep_.CoverTab[116653]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1118
			// _ = "end of CoverTab[116653]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1118
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1118
		// _ = "end of CoverTab[116650]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1119
	// _ = "end of CoverTab[116627]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1119
	_go_fuzz_dep_.CoverTab[116628]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1122
	val.Set(valSlice)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1125
	if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1125
		_go_fuzz_dep_.CoverTab[116654]++
														return &Error{errors}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1126
		// _ = "end of CoverTab[116654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1127
		_go_fuzz_dep_.CoverTab[116655]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1127
		// _ = "end of CoverTab[116655]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1127
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1127
	// _ = "end of CoverTab[116628]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1127
	_go_fuzz_dep_.CoverTab[116629]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1129
	// _ = "end of CoverTab[116629]"
}

func (d *Decoder) decodeArray(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1132
	_go_fuzz_dep_.CoverTab[116656]++
													dataVal := reflect.Indirect(reflect.ValueOf(data))
													dataValKind := dataVal.Kind()
													valType := val.Type()
													valElemType := valType.Elem()
													arrayType := reflect.ArrayOf(valType.Len(), valElemType)

													valArray := val

													if valArray.Interface() == reflect.Zero(valArray.Type()).Interface() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1141
		_go_fuzz_dep_.CoverTab[116660]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1141
		return d.config.ZeroFields
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1141
		// _ = "end of CoverTab[116660]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1141
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1141
		_go_fuzz_dep_.CoverTab[116661]++

														if dataValKind != reflect.Array && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1143
			_go_fuzz_dep_.CoverTab[116664]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1143
			return dataValKind != reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1143
			// _ = "end of CoverTab[116664]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1143
		}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1143
			_go_fuzz_dep_.CoverTab[116665]++
															if d.config.WeaklyTypedInput {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1144
				_go_fuzz_dep_.CoverTab[116667]++
																switch {

				case dataValKind == reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1147
					_go_fuzz_dep_.CoverTab[116668]++
																	if dataVal.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1148
						_go_fuzz_dep_.CoverTab[116670]++
																		val.Set(reflect.Zero(arrayType))
																		return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1150
						// _ = "end of CoverTab[116670]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1151
						_go_fuzz_dep_.CoverTab[116671]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1151
						// _ = "end of CoverTab[116671]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1151
					}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1151
					// _ = "end of CoverTab[116668]"

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1155
				default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1155
					_go_fuzz_dep_.CoverTab[116669]++

																	return d.decodeArray(name, []interface{}{data}, val)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1157
					// _ = "end of CoverTab[116669]"
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1158
				// _ = "end of CoverTab[116667]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1159
				_go_fuzz_dep_.CoverTab[116672]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1159
				// _ = "end of CoverTab[116672]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1159
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1159
			// _ = "end of CoverTab[116665]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1159
			_go_fuzz_dep_.CoverTab[116666]++

															return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", name, dataValKind)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1162
			// _ = "end of CoverTab[116666]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1164
			_go_fuzz_dep_.CoverTab[116673]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1164
			// _ = "end of CoverTab[116673]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1164
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1164
		// _ = "end of CoverTab[116661]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1164
		_go_fuzz_dep_.CoverTab[116662]++
														if dataVal.Len() > arrayType.Len() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1165
			_go_fuzz_dep_.CoverTab[116674]++
															return fmt.Errorf(
				"'%s': expected source data to have length less or equal to %d, got %d", name, arrayType.Len(), dataVal.Len())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1167
			// _ = "end of CoverTab[116674]"

		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1169
			_go_fuzz_dep_.CoverTab[116675]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1169
			// _ = "end of CoverTab[116675]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1169
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1169
		// _ = "end of CoverTab[116662]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1169
		_go_fuzz_dep_.CoverTab[116663]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1172
		valArray = reflect.New(arrayType).Elem()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1172
		// _ = "end of CoverTab[116663]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1173
		_go_fuzz_dep_.CoverTab[116676]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1173
		// _ = "end of CoverTab[116676]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1173
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1173
	// _ = "end of CoverTab[116656]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1173
	_go_fuzz_dep_.CoverTab[116657]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1176
	errors := make([]string, 0)

	for i := 0; i < dataVal.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1178
		_go_fuzz_dep_.CoverTab[116677]++
														currentData := dataVal.Index(i).Interface()
														currentField := valArray.Index(i)

														fieldName := name + "[" + strconv.Itoa(i) + "]"
														if err := d.decode(fieldName, currentData, currentField); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1183
			_go_fuzz_dep_.CoverTab[116678]++
															errors = appendErrors(errors, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1184
			// _ = "end of CoverTab[116678]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1185
			_go_fuzz_dep_.CoverTab[116679]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1185
			// _ = "end of CoverTab[116679]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1185
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1185
		// _ = "end of CoverTab[116677]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1186
	// _ = "end of CoverTab[116657]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1186
	_go_fuzz_dep_.CoverTab[116658]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1189
	val.Set(valArray)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1192
	if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1192
		_go_fuzz_dep_.CoverTab[116680]++
														return &Error{errors}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1193
		// _ = "end of CoverTab[116680]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1194
		_go_fuzz_dep_.CoverTab[116681]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1194
		// _ = "end of CoverTab[116681]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1194
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1194
	// _ = "end of CoverTab[116658]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1194
	_go_fuzz_dep_.CoverTab[116659]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1196
	// _ = "end of CoverTab[116659]"
}

func (d *Decoder) decodeStruct(name string, data interface{}, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1199
	_go_fuzz_dep_.CoverTab[116682]++
													dataVal := reflect.Indirect(reflect.ValueOf(data))

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1204
	if dataVal.Type() == val.Type() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1204
		_go_fuzz_dep_.CoverTab[116684]++
														val.Set(dataVal)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1206
		// _ = "end of CoverTab[116684]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1207
		_go_fuzz_dep_.CoverTab[116685]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1207
		// _ = "end of CoverTab[116685]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1207
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1207
	// _ = "end of CoverTab[116682]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1207
	_go_fuzz_dep_.CoverTab[116683]++

													dataValKind := dataVal.Kind()
													switch dataValKind {
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1211
		_go_fuzz_dep_.CoverTab[116686]++
														return d.decodeStructFromMap(name, dataVal, val)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1212
		// _ = "end of CoverTab[116686]"

	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1214
		_go_fuzz_dep_.CoverTab[116687]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1220
		mapType := reflect.TypeOf((map[string]interface{})(nil))
														mval := reflect.MakeMap(mapType)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1227
		addrVal := reflect.New(mval.Type())

		reflect.Indirect(addrVal).Set(mval)
		if err := d.decodeMapFromStruct(name, dataVal, reflect.Indirect(addrVal), mval); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1230
			_go_fuzz_dep_.CoverTab[116690]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1231
			// _ = "end of CoverTab[116690]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1232
			_go_fuzz_dep_.CoverTab[116691]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1232
			// _ = "end of CoverTab[116691]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1232
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1232
		// _ = "end of CoverTab[116687]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1232
		_go_fuzz_dep_.CoverTab[116688]++

														result := d.decodeStructFromMap(name, reflect.Indirect(addrVal), val)
														return result
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1235
		// _ = "end of CoverTab[116688]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1237
		_go_fuzz_dep_.CoverTab[116689]++
														return fmt.Errorf("'%s' expected a map, got '%s'", name, dataVal.Kind())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1238
		// _ = "end of CoverTab[116689]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1239
	// _ = "end of CoverTab[116683]"
}

func (d *Decoder) decodeStructFromMap(name string, dataVal, val reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1242
	_go_fuzz_dep_.CoverTab[116692]++
													dataValType := dataVal.Type()
													if kind := dataValType.Key().Kind(); kind != reflect.String && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1244
		_go_fuzz_dep_.CoverTab[116701]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1244
		return kind != reflect.Interface
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1244
		// _ = "end of CoverTab[116701]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1244
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1244
		_go_fuzz_dep_.CoverTab[116702]++
														return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys",
			name, dataValType.Key().Kind())
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1247
		// _ = "end of CoverTab[116702]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1248
		_go_fuzz_dep_.CoverTab[116703]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1248
		// _ = "end of CoverTab[116703]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1248
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1248
	// _ = "end of CoverTab[116692]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1248
	_go_fuzz_dep_.CoverTab[116693]++

													dataValKeys := make(map[reflect.Value]struct{})
													dataValKeysUnused := make(map[interface{}]struct{})
													for _, dataValKey := range dataVal.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1252
		_go_fuzz_dep_.CoverTab[116704]++
														dataValKeys[dataValKey] = struct{}{}
														dataValKeysUnused[dataValKey.Interface()] = struct{}{}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1254
		// _ = "end of CoverTab[116704]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1255
	// _ = "end of CoverTab[116693]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1255
	_go_fuzz_dep_.CoverTab[116694]++

													errors := make([]string, 0)

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1262
	structs := make([]reflect.Value, 1, 5)
	structs[0] = val

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type field struct {
		field	reflect.StructField
		val	reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var remainField *field

	fields := []field{}
	for len(structs) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1277
		_go_fuzz_dep_.CoverTab[116705]++
														structVal := structs[0]
														structs = structs[1:]

														structType := structVal.Type()

														for i := 0; i < structType.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1283
			_go_fuzz_dep_.CoverTab[116706]++
															fieldType := structType.Field(i)
															fieldVal := structVal.Field(i)
															if fieldVal.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1286
				_go_fuzz_dep_.CoverTab[116710]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1286
				return fieldVal.Elem().Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1286
				// _ = "end of CoverTab[116710]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1286
			}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1286
				_go_fuzz_dep_.CoverTab[116711]++

																fieldVal = fieldVal.Elem()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1288
				// _ = "end of CoverTab[116711]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1289
				_go_fuzz_dep_.CoverTab[116712]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1289
				// _ = "end of CoverTab[116712]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1289
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1289
			// _ = "end of CoverTab[116706]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1289
			_go_fuzz_dep_.CoverTab[116707]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
			squash := d.config.Squash && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
				_go_fuzz_dep_.CoverTab[116713]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
				return fieldVal.Kind() == reflect.Struct
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
				// _ = "end of CoverTab[116713]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
				_go_fuzz_dep_.CoverTab[116714]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
				return fieldType.Anonymous
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
				// _ = "end of CoverTab[116714]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1292
			}()
															remain := false

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1296
			tagParts := strings.Split(fieldType.Tag.Get(d.config.TagName), ",")
			for _, tag := range tagParts[1:] {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1297
				_go_fuzz_dep_.CoverTab[116715]++
																if tag == "squash" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1298
					_go_fuzz_dep_.CoverTab[116717]++
																	squash = true
																	break
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1300
					// _ = "end of CoverTab[116717]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1301
					_go_fuzz_dep_.CoverTab[116718]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1301
					// _ = "end of CoverTab[116718]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1301
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1301
				// _ = "end of CoverTab[116715]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1301
				_go_fuzz_dep_.CoverTab[116716]++

																if tag == "remain" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1303
					_go_fuzz_dep_.CoverTab[116719]++
																	remain = true
																	break
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1305
					// _ = "end of CoverTab[116719]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1306
					_go_fuzz_dep_.CoverTab[116720]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1306
					// _ = "end of CoverTab[116720]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1306
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1306
				// _ = "end of CoverTab[116716]"
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1307
			// _ = "end of CoverTab[116707]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1307
			_go_fuzz_dep_.CoverTab[116708]++

															if squash {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1309
				_go_fuzz_dep_.CoverTab[116721]++
																if fieldVal.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1310
					_go_fuzz_dep_.CoverTab[116723]++
																	errors = appendErrors(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", fieldType.Name, fieldVal.Kind()))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1312
					// _ = "end of CoverTab[116723]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1313
					_go_fuzz_dep_.CoverTab[116724]++
																	structs = append(structs, fieldVal)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1314
					// _ = "end of CoverTab[116724]"
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1315
				// _ = "end of CoverTab[116721]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1315
				_go_fuzz_dep_.CoverTab[116722]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1316
				// _ = "end of CoverTab[116722]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1317
				_go_fuzz_dep_.CoverTab[116725]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1317
				// _ = "end of CoverTab[116725]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1317
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1317
			// _ = "end of CoverTab[116708]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1317
			_go_fuzz_dep_.CoverTab[116709]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1320
			if remain {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1320
				_go_fuzz_dep_.CoverTab[116726]++
																remainField = &field{fieldType, fieldVal}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1321
				// _ = "end of CoverTab[116726]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1322
				_go_fuzz_dep_.CoverTab[116727]++

																fields = append(fields, field{fieldType, fieldVal})
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1324
				// _ = "end of CoverTab[116727]"
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1325
			// _ = "end of CoverTab[116709]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1326
		// _ = "end of CoverTab[116705]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1327
	// _ = "end of CoverTab[116694]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1327
	_go_fuzz_dep_.CoverTab[116695]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1330
	for _, f := range fields {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1330
		_go_fuzz_dep_.CoverTab[116728]++
														field, fieldValue := f.field, f.val
														fieldName := field.Name

														tagValue := field.Tag.Get(d.config.TagName)
														tagValue = strings.SplitN(tagValue, ",", 2)[0]
														if tagValue != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1336
			_go_fuzz_dep_.CoverTab[116734]++
															fieldName = tagValue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1337
			// _ = "end of CoverTab[116734]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1338
			_go_fuzz_dep_.CoverTab[116735]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1338
			// _ = "end of CoverTab[116735]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1338
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1338
		// _ = "end of CoverTab[116728]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1338
		_go_fuzz_dep_.CoverTab[116729]++

														rawMapKey := reflect.ValueOf(fieldName)
														rawMapVal := dataVal.MapIndex(rawMapKey)
														if !rawMapVal.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1342
			_go_fuzz_dep_.CoverTab[116736]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1345
			for dataValKey := range dataValKeys {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1345
				_go_fuzz_dep_.CoverTab[116738]++
																mK, ok := dataValKey.Interface().(string)
																if !ok {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1347
					_go_fuzz_dep_.CoverTab[116740]++

																	continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1349
					// _ = "end of CoverTab[116740]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1350
					_go_fuzz_dep_.CoverTab[116741]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1350
					// _ = "end of CoverTab[116741]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1350
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1350
				// _ = "end of CoverTab[116738]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1350
				_go_fuzz_dep_.CoverTab[116739]++

																if d.config.MatchName(mK, fieldName) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1352
					_go_fuzz_dep_.CoverTab[116742]++
																	rawMapKey = dataValKey
																	rawMapVal = dataVal.MapIndex(dataValKey)
																	break
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1355
					// _ = "end of CoverTab[116742]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1356
					_go_fuzz_dep_.CoverTab[116743]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1356
					// _ = "end of CoverTab[116743]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1356
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1356
				// _ = "end of CoverTab[116739]"
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1357
			// _ = "end of CoverTab[116736]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1357
			_go_fuzz_dep_.CoverTab[116737]++

															if !rawMapVal.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1359
				_go_fuzz_dep_.CoverTab[116744]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1362
				continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1362
				// _ = "end of CoverTab[116744]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1363
				_go_fuzz_dep_.CoverTab[116745]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1363
				// _ = "end of CoverTab[116745]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1363
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1363
			// _ = "end of CoverTab[116737]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1364
			_go_fuzz_dep_.CoverTab[116746]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1364
			// _ = "end of CoverTab[116746]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1364
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1364
		// _ = "end of CoverTab[116729]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1364
		_go_fuzz_dep_.CoverTab[116730]++

														if !fieldValue.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1366
			_go_fuzz_dep_.CoverTab[116747]++

															panic("field is not valid")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1368
			// _ = "end of CoverTab[116747]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1369
			_go_fuzz_dep_.CoverTab[116748]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1369
			// _ = "end of CoverTab[116748]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1369
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1369
		// _ = "end of CoverTab[116730]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1369
		_go_fuzz_dep_.CoverTab[116731]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1373
		if !fieldValue.CanSet() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1373
			_go_fuzz_dep_.CoverTab[116749]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1374
			// _ = "end of CoverTab[116749]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1375
			_go_fuzz_dep_.CoverTab[116750]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1375
			// _ = "end of CoverTab[116750]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1375
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1375
		// _ = "end of CoverTab[116731]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1375
		_go_fuzz_dep_.CoverTab[116732]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1378
		delete(dataValKeysUnused, rawMapKey.Interface())

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1382
		if name != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1382
			_go_fuzz_dep_.CoverTab[116751]++
															fieldName = name + "." + fieldName
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1383
			// _ = "end of CoverTab[116751]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1384
			_go_fuzz_dep_.CoverTab[116752]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1384
			// _ = "end of CoverTab[116752]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1384
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1384
		// _ = "end of CoverTab[116732]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1384
		_go_fuzz_dep_.CoverTab[116733]++

														if err := d.decode(fieldName, rawMapVal.Interface(), fieldValue); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1386
			_go_fuzz_dep_.CoverTab[116753]++
															errors = appendErrors(errors, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1387
			// _ = "end of CoverTab[116753]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1388
			_go_fuzz_dep_.CoverTab[116754]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1388
			// _ = "end of CoverTab[116754]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1388
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1388
		// _ = "end of CoverTab[116733]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1389
	// _ = "end of CoverTab[116695]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1389
	_go_fuzz_dep_.CoverTab[116696]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1393
	if remainField != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1393
		_go_fuzz_dep_.CoverTab[116755]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1393
		return len(dataValKeysUnused) > 0
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1393
		// _ = "end of CoverTab[116755]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1393
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1393
		_go_fuzz_dep_.CoverTab[116756]++

														remain := map[interface{}]interface{}{}
														for key := range dataValKeysUnused {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1396
			_go_fuzz_dep_.CoverTab[116759]++
															remain[key] = dataVal.MapIndex(reflect.ValueOf(key)).Interface()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1397
			// _ = "end of CoverTab[116759]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1398
		// _ = "end of CoverTab[116756]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1398
		_go_fuzz_dep_.CoverTab[116757]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1401
		if err := d.decodeMap(name, remain, remainField.val); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1401
			_go_fuzz_dep_.CoverTab[116760]++
															errors = appendErrors(errors, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1402
			// _ = "end of CoverTab[116760]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1403
			_go_fuzz_dep_.CoverTab[116761]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1403
			// _ = "end of CoverTab[116761]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1403
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1403
		// _ = "end of CoverTab[116757]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1403
		_go_fuzz_dep_.CoverTab[116758]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1407
		dataValKeysUnused = nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1407
		// _ = "end of CoverTab[116758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1408
		_go_fuzz_dep_.CoverTab[116762]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1408
		// _ = "end of CoverTab[116762]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1408
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1408
	// _ = "end of CoverTab[116696]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1408
	_go_fuzz_dep_.CoverTab[116697]++

													if d.config.ErrorUnused && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1410
		_go_fuzz_dep_.CoverTab[116763]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1410
		return len(dataValKeysUnused) > 0
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1410
		// _ = "end of CoverTab[116763]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1410
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1410
		_go_fuzz_dep_.CoverTab[116764]++
														keys := make([]string, 0, len(dataValKeysUnused))
														for rawKey := range dataValKeysUnused {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1412
			_go_fuzz_dep_.CoverTab[116766]++
															keys = append(keys, rawKey.(string))
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1413
			// _ = "end of CoverTab[116766]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1414
		// _ = "end of CoverTab[116764]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1414
		_go_fuzz_dep_.CoverTab[116765]++
														sort.Strings(keys)

														err := fmt.Errorf("'%s' has invalid keys: %s", name, strings.Join(keys, ", "))
														errors = appendErrors(errors, err)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1418
		// _ = "end of CoverTab[116765]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1419
		_go_fuzz_dep_.CoverTab[116767]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1419
		// _ = "end of CoverTab[116767]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1419
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1419
	// _ = "end of CoverTab[116697]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1419
	_go_fuzz_dep_.CoverTab[116698]++

													if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1421
		_go_fuzz_dep_.CoverTab[116768]++
														return &Error{errors}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1422
		// _ = "end of CoverTab[116768]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1423
		_go_fuzz_dep_.CoverTab[116769]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1423
		// _ = "end of CoverTab[116769]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1423
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1423
	// _ = "end of CoverTab[116698]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1423
	_go_fuzz_dep_.CoverTab[116699]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1426
	if d.config.Metadata != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1426
		_go_fuzz_dep_.CoverTab[116770]++
														for rawKey := range dataValKeysUnused {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1427
			_go_fuzz_dep_.CoverTab[116771]++
															key := rawKey.(string)
															if name != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1429
				_go_fuzz_dep_.CoverTab[116773]++
																key = name + "." + key
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1430
				// _ = "end of CoverTab[116773]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1431
				_go_fuzz_dep_.CoverTab[116774]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1431
				// _ = "end of CoverTab[116774]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1431
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1431
			// _ = "end of CoverTab[116771]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1431
			_go_fuzz_dep_.CoverTab[116772]++

															d.config.Metadata.Unused = append(d.config.Metadata.Unused, key)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1433
			// _ = "end of CoverTab[116772]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1434
		// _ = "end of CoverTab[116770]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1435
		_go_fuzz_dep_.CoverTab[116775]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1435
		// _ = "end of CoverTab[116775]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1435
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1435
	// _ = "end of CoverTab[116699]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1435
	_go_fuzz_dep_.CoverTab[116700]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1437
	// _ = "end of CoverTab[116700]"
}

func isEmptyValue(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1440
	_go_fuzz_dep_.CoverTab[116776]++
													switch getKind(v) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1442
		_go_fuzz_dep_.CoverTab[116778]++
														return v.Len() == 0
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1443
		// _ = "end of CoverTab[116778]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1444
		_go_fuzz_dep_.CoverTab[116779]++
														return !v.Bool()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1445
		// _ = "end of CoverTab[116779]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1446
		_go_fuzz_dep_.CoverTab[116780]++
														return v.Int() == 0
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1447
		// _ = "end of CoverTab[116780]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1448
		_go_fuzz_dep_.CoverTab[116781]++
														return v.Uint() == 0
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1449
		// _ = "end of CoverTab[116781]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1450
		_go_fuzz_dep_.CoverTab[116782]++
														return v.Float() == 0
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1451
		// _ = "end of CoverTab[116782]"
	case reflect.Interface, reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1452
		_go_fuzz_dep_.CoverTab[116783]++
														return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1453
		// _ = "end of CoverTab[116783]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1453
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1453
		_go_fuzz_dep_.CoverTab[116784]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1453
		// _ = "end of CoverTab[116784]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1454
	// _ = "end of CoverTab[116776]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1454
	_go_fuzz_dep_.CoverTab[116777]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1455
	// _ = "end of CoverTab[116777]"
}

func getKind(val reflect.Value) reflect.Kind {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1458
	_go_fuzz_dep_.CoverTab[116785]++
													kind := val.Kind()

													switch {
	case kind >= reflect.Int && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1462
		_go_fuzz_dep_.CoverTab[116790]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1462
		return kind <= reflect.Int64
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1462
		// _ = "end of CoverTab[116790]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1462
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1462
		_go_fuzz_dep_.CoverTab[116786]++
														return reflect.Int
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1463
		// _ = "end of CoverTab[116786]"
	case kind >= reflect.Uint && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1464
		_go_fuzz_dep_.CoverTab[116791]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1464
		return kind <= reflect.Uint64
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1464
		// _ = "end of CoverTab[116791]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1464
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1464
		_go_fuzz_dep_.CoverTab[116787]++
														return reflect.Uint
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1465
		// _ = "end of CoverTab[116787]"
	case kind >= reflect.Float32 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1466
		_go_fuzz_dep_.CoverTab[116792]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1466
		return kind <= reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1466
		// _ = "end of CoverTab[116792]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1466
	}():
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1466
		_go_fuzz_dep_.CoverTab[116788]++
														return reflect.Float32
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1467
		// _ = "end of CoverTab[116788]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1468
		_go_fuzz_dep_.CoverTab[116789]++
														return kind
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1469
		// _ = "end of CoverTab[116789]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1470
	// _ = "end of CoverTab[116785]"
}

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1471
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/mapstructure@v1.4.2/mapstructure.go:1471
var _ = _go_fuzz_dep_.CoverTab
