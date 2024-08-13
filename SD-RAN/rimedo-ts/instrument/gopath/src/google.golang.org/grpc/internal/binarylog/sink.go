//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:19
package binarylog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:19
)

import (
	"bufio"
	"encoding/binary"
	"io"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	binlogpb "google.golang.org/grpc/binarylog/grpc_binarylog_v1"
)

var (
	// DefaultSink is the sink where the logs will be written to. It's exported
	// for the binarylog package to update.
	DefaultSink Sink = &noopSink{}	// TODO(blog): change this default (file in /tmp).
)

// Sink writes log entry into the binary log sink.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:38
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:38
// sink is a copy of the exported binarylog.Sink, to avoid circular dependency.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:41
type Sink interface {
	// Write will be called to write the log entry into the sink.
	//
	// It should be thread-safe so it can be called in parallel.
	Write(*binlogpb.GrpcLogEntry) error
	// Close will be called when the Sink is replaced by a new Sink.
	Close() error
}

type noopSink struct{}

func (ns *noopSink) Write(*binlogpb.GrpcLogEntry) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:52
	_go_fuzz_dep_.CoverTab[68835]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:52
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:52
	// _ = "end of CoverTab[68835]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:52
}
func (ns *noopSink) Close() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:53
	_go_fuzz_dep_.CoverTab[68836]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:53
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:53
	// _ = "end of CoverTab[68836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:53
}

// newWriterSink creates a binary log sink with the given writer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:55
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:55
// Write() marshals the proto message and writes it to the given writer. Each
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:55
// message is prefixed with a 4 byte big endian unsigned integer as the length.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:55
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:55
// No buffer is done, Close() doesn't try to close the writer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:61
func newWriterSink(w io.Writer) Sink {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:61
	_go_fuzz_dep_.CoverTab[68837]++
												return &writerSink{out: w}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:62
	// _ = "end of CoverTab[68837]"
}

type writerSink struct {
	out io.Writer
}

func (ws *writerSink) Write(e *binlogpb.GrpcLogEntry) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:69
	_go_fuzz_dep_.CoverTab[68838]++
												b, err := proto.Marshal(e)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:71
		_go_fuzz_dep_.CoverTab[68842]++
													grpclogLogger.Errorf("binary logging: failed to marshal proto message: %v", err)
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:73
		// _ = "end of CoverTab[68842]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:74
		_go_fuzz_dep_.CoverTab[68843]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:74
		// _ = "end of CoverTab[68843]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:74
	// _ = "end of CoverTab[68838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:74
	_go_fuzz_dep_.CoverTab[68839]++
												hdr := make([]byte, 4)
												binary.BigEndian.PutUint32(hdr, uint32(len(b)))
												if _, err := ws.out.Write(hdr); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:77
		_go_fuzz_dep_.CoverTab[68844]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:78
		// _ = "end of CoverTab[68844]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:79
		_go_fuzz_dep_.CoverTab[68845]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:79
		// _ = "end of CoverTab[68845]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:79
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:79
	// _ = "end of CoverTab[68839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:79
	_go_fuzz_dep_.CoverTab[68840]++
												if _, err := ws.out.Write(b); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:80
		_go_fuzz_dep_.CoverTab[68846]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:81
		// _ = "end of CoverTab[68846]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:82
		_go_fuzz_dep_.CoverTab[68847]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:82
		// _ = "end of CoverTab[68847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:82
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:82
	// _ = "end of CoverTab[68840]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:82
	_go_fuzz_dep_.CoverTab[68841]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:83
	// _ = "end of CoverTab[68841]"
}

func (ws *writerSink) Close() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:86
	_go_fuzz_dep_.CoverTab[68848]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:86
	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:86
	// _ = "end of CoverTab[68848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:86
}

type bufferedSink struct {
	mu		sync.Mutex
	closer		io.Closer
	out		Sink		// out is built on buf.
	buf		*bufio.Writer	// buf is kept for flush.
	flusherStarted	bool

	writeTicker	*time.Ticker
	done		chan struct{}
}

func (fs *bufferedSink) Write(e *binlogpb.GrpcLogEntry) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:99
	_go_fuzz_dep_.CoverTab[68849]++
												fs.mu.Lock()
												defer fs.mu.Unlock()
												if !fs.flusherStarted {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:102
		_go_fuzz_dep_.CoverTab[68852]++

														fs.startFlushGoroutine()
														fs.flusherStarted = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:105
		// _ = "end of CoverTab[68852]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:106
		_go_fuzz_dep_.CoverTab[68853]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:106
		// _ = "end of CoverTab[68853]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:106
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:106
	// _ = "end of CoverTab[68849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:106
	_go_fuzz_dep_.CoverTab[68850]++
													if err := fs.out.Write(e); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:107
		_go_fuzz_dep_.CoverTab[68854]++
														return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:108
		// _ = "end of CoverTab[68854]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:109
		_go_fuzz_dep_.CoverTab[68855]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:109
		// _ = "end of CoverTab[68855]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:109
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:109
	// _ = "end of CoverTab[68850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:109
	_go_fuzz_dep_.CoverTab[68851]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:110
	// _ = "end of CoverTab[68851]"
}

const (
	bufFlushDuration = 60 * time.Second
)

func (fs *bufferedSink) startFlushGoroutine() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:117
	_go_fuzz_dep_.CoverTab[68856]++
													fs.writeTicker = time.NewTicker(bufFlushDuration)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:118
	_curRoutineNum52_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:118
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum52_)
													go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:119
		_go_fuzz_dep_.CoverTab[68857]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:119
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:119
			_go_fuzz_dep_.CoverTab[68858]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:119
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum52_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:119
			// _ = "end of CoverTab[68858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:119
		}()
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:120
			_go_fuzz_dep_.CoverTab[68859]++
															select {
			case <-fs.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:122
				_go_fuzz_dep_.CoverTab[68862]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:123
				// _ = "end of CoverTab[68862]"
			case <-fs.writeTicker.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:124
				_go_fuzz_dep_.CoverTab[68863]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:124
				// _ = "end of CoverTab[68863]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:125
			// _ = "end of CoverTab[68859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:125
			_go_fuzz_dep_.CoverTab[68860]++
															fs.mu.Lock()
															if err := fs.buf.Flush(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:127
				_go_fuzz_dep_.CoverTab[68864]++
																grpclogLogger.Warningf("failed to flush to Sink: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:128
				// _ = "end of CoverTab[68864]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:129
				_go_fuzz_dep_.CoverTab[68865]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:129
				// _ = "end of CoverTab[68865]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:129
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:129
			// _ = "end of CoverTab[68860]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:129
			_go_fuzz_dep_.CoverTab[68861]++
															fs.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:130
			// _ = "end of CoverTab[68861]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:131
		// _ = "end of CoverTab[68857]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:132
	// _ = "end of CoverTab[68856]"
}

func (fs *bufferedSink) Close() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:135
	_go_fuzz_dep_.CoverTab[68866]++
													fs.mu.Lock()
													defer fs.mu.Unlock()
													if fs.writeTicker != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:138
		_go_fuzz_dep_.CoverTab[68871]++
														fs.writeTicker.Stop()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:139
		// _ = "end of CoverTab[68871]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:140
		_go_fuzz_dep_.CoverTab[68872]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:140
		// _ = "end of CoverTab[68872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:140
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:140
	// _ = "end of CoverTab[68866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:140
	_go_fuzz_dep_.CoverTab[68867]++
													close(fs.done)
													if err := fs.buf.Flush(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:142
		_go_fuzz_dep_.CoverTab[68873]++
														grpclogLogger.Warningf("failed to flush to Sink: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:143
		// _ = "end of CoverTab[68873]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:144
		_go_fuzz_dep_.CoverTab[68874]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:144
		// _ = "end of CoverTab[68874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:144
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:144
	// _ = "end of CoverTab[68867]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:144
	_go_fuzz_dep_.CoverTab[68868]++
													if err := fs.closer.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:145
		_go_fuzz_dep_.CoverTab[68875]++
														grpclogLogger.Warningf("failed to close the underlying WriterCloser: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:146
		// _ = "end of CoverTab[68875]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:147
		_go_fuzz_dep_.CoverTab[68876]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:147
		// _ = "end of CoverTab[68876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:147
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:147
	// _ = "end of CoverTab[68868]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:147
	_go_fuzz_dep_.CoverTab[68869]++
													if err := fs.out.Close(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:148
		_go_fuzz_dep_.CoverTab[68877]++
														grpclogLogger.Warningf("failed to close the Sink: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:149
		// _ = "end of CoverTab[68877]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:150
		_go_fuzz_dep_.CoverTab[68878]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:150
		// _ = "end of CoverTab[68878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:150
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:150
	// _ = "end of CoverTab[68869]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:150
	_go_fuzz_dep_.CoverTab[68870]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:151
	// _ = "end of CoverTab[68870]"
}

// NewBufferedSink creates a binary log sink with the given WriteCloser.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:154
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:154
// Write() marshals the proto message and writes it to the given writer. Each
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:154
// message is prefixed with a 4 byte big endian unsigned integer as the length.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:154
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:154
// Content is kept in a buffer, and is flushed every 60 seconds.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:154
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:154
// Close closes the WriteCloser.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:162
func NewBufferedSink(o io.WriteCloser) Sink {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:162
	_go_fuzz_dep_.CoverTab[68879]++
													bufW := bufio.NewWriter(o)
													return &bufferedSink{
		closer:	o,
		out:	newWriterSink(bufW),
		buf:	bufW,
		done:	make(chan struct{}),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:169
	// _ = "end of CoverTab[68879]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:170
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/sink.go:170
var _ = _go_fuzz_dep_.CoverTab
