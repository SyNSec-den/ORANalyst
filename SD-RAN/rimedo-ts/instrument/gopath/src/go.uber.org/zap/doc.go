// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// Package zap provides fast, structured, leveled logging.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// For applications that log in the hot path, reflection-based serialization
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// and string formatting are prohibitively expensive - they're CPU-intensive
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// and make many small allocations. Put differently, using json.Marshal and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// fmt.Fprintf to log tons of interface{} makes your application slow.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// Zap takes a different approach. It includes a reflection-free,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// zero-allocation JSON encoder, and the base Logger strives to avoid
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// serialization overhead and allocations wherever possible. By building the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// high-level SugaredLogger on that foundation, zap lets users choose when
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// they need to count every allocation and when they'd prefer a more familiar,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// loosely typed API.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// # Choosing a Logger
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// In contexts where performance is nice, but not critical, use the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// SugaredLogger. It's 4-10x faster than other structured logging packages and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// supports both structured and printf-style logging. Like log15 and go-kit,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// the SugaredLogger's structured logging APIs are loosely typed and accept a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// variadic number of key-value pairs. (For more advanced use cases, they also
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// accept strongly typed fields - see the SugaredLogger.With documentation for
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// details.)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	sugar := zap.NewExample().Sugar()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	defer sugar.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	sugar.Infow("failed to fetch URL",
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	  "url", "http://example.com",
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	  "attempt", 3,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	  "backoff", time.Second,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	sugar.Infof("failed to fetch URL: %s", "http://example.com")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// By default, loggers are unbuffered. However, since zap's low-level APIs
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// allow buffering, calling Sync before letting your process exit is a good
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// habit.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// In the rare contexts where every microsecond and every allocation matter,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// use the Logger. It's even faster than the SugaredLogger and allocates far
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// less, but it only supports strongly-typed, structured logging.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	logger := zap.NewExample()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	defer logger.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	logger.Info("failed to fetch URL",
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	  zap.String("url", "http://example.com"),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	  zap.Int("attempt", 3),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	  zap.Duration("backoff", time.Second),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// Choosing between the Logger and SugaredLogger doesn't need to be an
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// application-wide decision: converting between the two is simple and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// inexpensive.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	logger := zap.NewExample()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	defer logger.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	sugar := logger.Sugar()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	plain := sugar.Desugar()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// # Configuring Zap
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// The simplest way to build a Logger is to use zap's opinionated presets:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// NewExample, NewProduction, and NewDevelopment. These presets build a logger
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// with a single function call:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	logger, err := zap.NewProduction()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	  log.Fatalf("can't initialize zap logger: %v", err)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//	defer logger.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// Presets are fine for small projects, but larger projects and organizations
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// naturally require a bit more customization. For most users, zap's Config
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// struct strikes the right balance between flexibility and convenience. See
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// the package-level BasicConfiguration example for sample code.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// More unusual configurations (splitting output between files, sending logs
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// to a message queue, etc.) are possible, but require direct use of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// go.uber.org/zap/zapcore. See the package-level AdvancedConfiguration
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// example for sample code.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// # Extending Zap
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// The zap package itself is a relatively thin wrapper around the interfaces
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// in go.uber.org/zap/zapcore. Extending zap to support a new encoding (e.g.,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// BSON), a new log sink (e.g., Kafka), or something more exotic (perhaps an
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// exception aggregation service, like Sentry or Rollbar) typically requires
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// implementing the zapcore.Encoder, zapcore.WriteSyncer, or zapcore.Core
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// interfaces. See the zapcore documentation for details.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// Similarly, package authors can use the high-performance Encoder and Core
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// implementations in the zapcore package to build their own loggers.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// # Frequently Asked Questions
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// An FAQ covering everything from installation errors to design decisions is
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:21
// available at https://github.com/uber-go/zap/blob/master/FAQ.md.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
)

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/doc.go:113
var _ = _go_fuzz_dep_.CoverTab
