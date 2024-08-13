// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Package properties provides functions for reading and writing
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// ISO-8859-1 and UTF-8 encoded .properties files and has
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// support for recursive property expansion.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Java properties files are ISO-8859-1 encoded and use Unicode
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// literals for characters outside the ISO character set. Unicode
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// literals can be used in UTF-8 encoded properties files but
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// aren't necessary.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// To load a single properties file use MustLoadFile():
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	p := properties.MustLoadFile(filename, properties.UTF8)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// To load multiple properties files use MustLoadFiles()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// which loads the files in the given order and merges the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// result. Missing properties files can be ignored if the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// 'ignoreMissing' flag is set to true.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Filenames can contain environment variables which are expanded
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// before loading.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	f1 := "/etc/myapp/myapp.conf"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	f2 := "/home/${USER}/myapp.conf"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	p := MustLoadFiles([]string{f1, f2}, properties.UTF8, true)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// All of the different key/value delimiters ' ', ':' and '=' are
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// supported as well as the comment characters '!' and '#' and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// multi-line values.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	! this is a comment
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# and so is this
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# the following expressions are equal
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key=value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key:value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key = value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key : value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key = val\
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	      ue
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Properties stores all comments preceding a key and provides
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// GetComments() and SetComments() methods to retrieve and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// update them. The convenience functions GetComment() and
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// SetComment() allow access to the last comment. The
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// WriteComment() method writes properties files including
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// the comments and with the keys in the original order.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// This can be used for sanitizing properties files.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Property expansion is recursive and circular references
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// and malformed expressions are not allowed and cause an
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// error. Expansion of environment variables is supported.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# standard property
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key = value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# property expansion: key2 = value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key2 = ${key}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# recursive expansion: key3 = value
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key3 = ${key2}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# circular reference (error)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key = ${key}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# malformed expression (error)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	key = ${ke
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# refers to the users' home dir
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	home = ${HOME}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# local key takes precedence over env var: u = foo
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	USER = foo
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	u = ${USER}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// The default property expansion format is ${key} but can be
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// changed by setting different pre- and postfix values on the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Properties object.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	p := properties.NewProperties()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	p.Prefix = "#["
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	p.Postfix = "]#"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Properties provides convenience functions for getting typed
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// values with default values if the key does not exist or the
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// type conversion failed.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# Returns true if the value is either "1", "on", "yes" or "true"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# Returns false for every other value and the default value if
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# the key does not exist.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	v = p.GetBool("key", false)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# Returns the value if the key exists and the format conversion
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# was successful. Otherwise, the default value is returned.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	v = p.GetInt64("key", 999)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	v = p.GetUint64("key", 999)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	v = p.GetFloat64("key", 123.0)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	v = p.GetString("key", "def")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	v = p.GetDuration("key", 999)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// As an alternative properties may be applied with the standard
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// library's flag implementation at any time.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# Standard configuration
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	v = flag.Int("key", 999, "help message")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	flag.Parse()
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# Merge p into the flag set
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	p.MustFlag(flag.CommandLine)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Properties provides several MustXXX() convenience functions
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// which will terminate the app if an error occurs. The behavior
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// of the failure is configurable and the default is to call
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// log.Fatal(err). To have the MustXXX() functions panic instead
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// of logging the error set a different ErrorHandler before
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// you use the Properties package.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	properties.ErrorHandler = properties.PanicHandler
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	# Will panic instead of logging an error
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	p := properties.MustLoadFile("config.properties")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// You can also provide your own ErrorHandler function. The only requirement
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// is that the error handler function must exit after handling the error.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	  properties.ErrorHandler = func(err error) {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//		     fmt.Println(err)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	      os.Exit(1)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	  }
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	  # Will write to stdout and then exit
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	  p := properties.MustLoadFile("config.properties")
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// Properties can also be loaded into a struct via the `Decode`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// method, e.g.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	type S struct {
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	    A string        `properties:"a,default=foo"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	    D time.Duration `properties:"timeout,default=5s"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	    E time.Time     `properties:"expires,layout=2006-01-02,default=2015-01-01"`
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//	}
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// See `Decode()` method for the full documentation.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// The following documents provide a description of the properties
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// file format.
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// http://en.wikipedia.org/wiki/.properties
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
//
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:5
// http://docs.oracle.com/javase/7/docs/api/java/util/Properties.html#load%28java.io.Reader%29
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
package properties

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
import (
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
)

//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/magiconair/properties@v1.8.5/doc.go:156
var _ = _go_fuzz_dep_.CoverTab
