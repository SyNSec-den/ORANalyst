// +build !js

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:3
package viper

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:3
)

import "github.com/fsnotify/fsnotify"

type watcher = fsnotify.Watcher

func newWatcher() (*watcher, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:9
	_go_fuzz_dep_.CoverTab[130344]++
										return fsnotify.NewWatcher()
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:10
	// _ = "end of CoverTab[130344]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:11
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/watch.go:11
var _ = _go_fuzz_dep_.CoverTab
