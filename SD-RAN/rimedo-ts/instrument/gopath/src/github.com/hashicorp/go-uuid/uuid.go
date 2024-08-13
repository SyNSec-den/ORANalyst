//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:1
package uuid

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:1
)

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// GenerateRandomBytes is used to generate random bytes of given size.
func GenerateRandomBytes(size int) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:11
	_go_fuzz_dep_.CoverTab[85438]++
											return GenerateRandomBytesWithReader(size, rand.Reader)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:12
	// _ = "end of CoverTab[85438]"
}

// GenerateRandomBytesWithReader is used to generate random bytes of given size read from a given reader.
func GenerateRandomBytesWithReader(size int, reader io.Reader) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:16
	_go_fuzz_dep_.CoverTab[85439]++
											if reader == nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:17
		_go_fuzz_dep_.CoverTab[85442]++
												return nil, fmt.Errorf("provided reader is nil")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:18
		// _ = "end of CoverTab[85442]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:19
		_go_fuzz_dep_.CoverTab[85443]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:19
		// _ = "end of CoverTab[85443]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:19
	// _ = "end of CoverTab[85439]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:19
	_go_fuzz_dep_.CoverTab[85440]++
											buf := make([]byte, size)
											if _, err := io.ReadFull(reader, buf); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:21
		_go_fuzz_dep_.CoverTab[85444]++
												return nil, fmt.Errorf("failed to read random bytes: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:22
		// _ = "end of CoverTab[85444]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:23
		_go_fuzz_dep_.CoverTab[85445]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:23
		// _ = "end of CoverTab[85445]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:23
	// _ = "end of CoverTab[85440]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:23
	_go_fuzz_dep_.CoverTab[85441]++
											return buf, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:24
	// _ = "end of CoverTab[85441]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:28
const uuidLen = 16

// GenerateUUID is used to generate a random UUID
func GenerateUUID() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:31
	_go_fuzz_dep_.CoverTab[85446]++
											return GenerateUUIDWithReader(rand.Reader)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:32
	// _ = "end of CoverTab[85446]"
}

// GenerateUUIDWithReader is used to generate a random UUID with a given Reader
func GenerateUUIDWithReader(reader io.Reader) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:36
	_go_fuzz_dep_.CoverTab[85447]++
											if reader == nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:37
		_go_fuzz_dep_.CoverTab[85450]++
												return "", fmt.Errorf("provided reader is nil")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:38
		// _ = "end of CoverTab[85450]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:39
		_go_fuzz_dep_.CoverTab[85451]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:39
		// _ = "end of CoverTab[85451]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:39
	// _ = "end of CoverTab[85447]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:39
	_go_fuzz_dep_.CoverTab[85448]++
											buf, err := GenerateRandomBytesWithReader(uuidLen, reader)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:41
		_go_fuzz_dep_.CoverTab[85452]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:42
		// _ = "end of CoverTab[85452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:43
		_go_fuzz_dep_.CoverTab[85453]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:43
		// _ = "end of CoverTab[85453]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:43
	// _ = "end of CoverTab[85448]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:43
	_go_fuzz_dep_.CoverTab[85449]++
											return FormatUUID(buf)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:44
	// _ = "end of CoverTab[85449]"
}

func FormatUUID(buf []byte) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:47
	_go_fuzz_dep_.CoverTab[85454]++
											if buflen := len(buf); buflen != uuidLen {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:48
		_go_fuzz_dep_.CoverTab[85456]++
												return "", fmt.Errorf("wrong length byte slice (%d)", buflen)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:49
		// _ = "end of CoverTab[85456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:50
		_go_fuzz_dep_.CoverTab[85457]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:50
		// _ = "end of CoverTab[85457]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:50
	// _ = "end of CoverTab[85454]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:50
	_go_fuzz_dep_.CoverTab[85455]++

											return fmt.Sprintf("%x-%x-%x-%x-%x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16]), nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:57
	// _ = "end of CoverTab[85455]"
}

func ParseUUID(uuid string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:60
	_go_fuzz_dep_.CoverTab[85458]++
											if len(uuid) != 2*uuidLen+4 {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:61
		_go_fuzz_dep_.CoverTab[85463]++
												return nil, fmt.Errorf("uuid string is wrong length")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:62
		// _ = "end of CoverTab[85463]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:63
		_go_fuzz_dep_.CoverTab[85464]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:63
		// _ = "end of CoverTab[85464]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:63
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:63
	// _ = "end of CoverTab[85458]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:63
	_go_fuzz_dep_.CoverTab[85459]++

											if uuid[8] != '-' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:65
		_go_fuzz_dep_.CoverTab[85465]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:65
		return uuid[13] != '-'
												// _ = "end of CoverTab[85465]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:66
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:66
		_go_fuzz_dep_.CoverTab[85466]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:66
		return uuid[18] != '-'
												// _ = "end of CoverTab[85466]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:67
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:67
		_go_fuzz_dep_.CoverTab[85467]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:67
		return uuid[23] != '-'
												// _ = "end of CoverTab[85467]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:68
	}() {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:68
		_go_fuzz_dep_.CoverTab[85468]++
												return nil, fmt.Errorf("uuid is improperly formatted")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:69
		// _ = "end of CoverTab[85468]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:70
		_go_fuzz_dep_.CoverTab[85469]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:70
		// _ = "end of CoverTab[85469]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:70
	// _ = "end of CoverTab[85459]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:70
	_go_fuzz_dep_.CoverTab[85460]++

											hexStr := uuid[0:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:36]

											ret, err := hex.DecodeString(hexStr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:75
		_go_fuzz_dep_.CoverTab[85470]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:76
		// _ = "end of CoverTab[85470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:77
		_go_fuzz_dep_.CoverTab[85471]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:77
		// _ = "end of CoverTab[85471]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:77
	// _ = "end of CoverTab[85460]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:77
	_go_fuzz_dep_.CoverTab[85461]++
											if len(ret) != uuidLen {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:78
		_go_fuzz_dep_.CoverTab[85472]++
												return nil, fmt.Errorf("decoded hex is the wrong length")
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:79
		// _ = "end of CoverTab[85472]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:80
		_go_fuzz_dep_.CoverTab[85473]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:80
		// _ = "end of CoverTab[85473]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:80
	// _ = "end of CoverTab[85461]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:80
	_go_fuzz_dep_.CoverTab[85462]++

											return ret, nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:82
	// _ = "end of CoverTab[85462]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:83
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/go-uuid@v1.0.2/uuid.go:83
var _ = _go_fuzz_dep_.CoverTab
