package main

import (
	"io"
	"net"
	"syscall"

	. "github.com/dvyukov/go-fuzz/go-fuzz-defs"
)

type FD int

func setupCommFile() ([]byte, FD, FD) {
	mem, err := syscall.Mmap(3, 0, CoverSize+MaxInputSize+SonarRegionSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		println("failed to mmap fd = 3 errno =", err.(syscall.Errno))
		println("error:", err.Error())
		syscall.Exit(1)
	}
	return mem, 4, 5
}

func fakeMem() []byte {
	return make([]byte, CoverSize+MaxInputSize+SonarRegionSize)
}

// read reads little-endian-encoded uint8+uint64 from fd.
func read(fd FD) (uint8, uint64) {
	rd := 0
	var buf [9]byte
	for rd != len(buf) {
		n, err := fd.read(buf[rd:])
		if err == syscall.EINTR {
			continue
		}
		if n == 0 {
			syscall.Exit(1)
		}
		if err != nil {
			println("failed to read fd =", fd, "errno =", err.(syscall.Errno))
			syscall.Exit(1)
		}
		rd += n
	}
	return buf[0], deserialize64(buf[1:])
}

// write writes little-endian-encoded vals... to fd.
func write(fd FD, vals ...uint64) {
	var tmp [3 * 8]byte
	buf := tmp[:len(vals)*8]
	for i, v := range vals {
		serialize64(buf[i*8:], v)
	}
	wr := 0
	for wr != len(buf) {
		n, err := fd.write(buf[wr:])
		if err == syscall.EINTR {
			continue
		}
		if err != nil {
			println("failed to read fd =", fd, "errno =", err.(syscall.Errno))
			syscall.Exit(1)
		}
		wr += n
	}
}

func (fd FD) read(buf []byte) (int, error) {
	return syscall.Read(int(fd), buf)
}

func (fd FD) write(buf []byte) (int, error) {
	return syscall.Write(int(fd), buf)
}

func serialize64(buf []byte, v uint64) uint8 {
	_ = buf[7]
	buf[0] = byte(v >> 0)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
	return 8
}

func deserialize64(buf []byte) uint64 {
	_ = buf[7]
	return uint64(buf[0])<<0 |
		uint64(buf[1])<<8 |
		uint64(buf[2])<<16 |
		uint64(buf[3])<<24 |
		uint64(buf[4])<<32 |
		uint64(buf[5])<<40 |
		uint64(buf[6])<<48 |
		uint64(buf[7])<<56
}

func writeAll(c net.Conn, b []byte) (err error) {
	totalBytesWritten := 0
	for totalBytesWritten < len(b) {
		n, err := c.Write(b[totalBytesWritten:])
		if err != nil {
			netError, ok := err.(net.Error)
			if ok && netError.Temporary() {
				// If it's a temporary error, we'll retry.
				continue
			} else {
				// If it's not a temporary error, we return with the error.
				return err
			}
		}
		totalBytesWritten += n
	}
	return nil
}

func readAll(c net.Conn, b []byte) (n int, err error) {
	totalBytesRead := 0
	for totalBytesRead < len(b) {
		n, err = c.Read(b[totalBytesRead:])
		if err != nil {
			netError, ok := err.(net.Error)
			if ok && netError.Temporary() {
				// If it's a temporary error, we'll retry.
				continue
			} else if err != io.EOF {
				// If it's not a temporary error nor an EOF, we return with the error.
				return totalBytesRead, err
			}
		}
		totalBytesRead += n
		// If we've read to the end of the file, break out of the loop.
		if err == io.EOF {
			break
		}
	}
	return totalBytesRead, err
}
