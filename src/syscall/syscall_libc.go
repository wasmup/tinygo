// +build darwin nintendoswitch wasi

package syscall

import (
	"unsafe"
)

type sliceHeader struct {
	buf *byte
	len uintptr
	cap uintptr
}

func Close(fd int) (err error) {
	if libc_close(fd) < 0 {
		err = getErrno()
	}
	return
}

func Write(fd int, p []byte) (n int, err error) {
	buf, count := splitSlice(p)
	n = libc_write(fd, buf, uint(count))
	if n < 0 {
		err = getErrno()
	}
	return
}

func Read(fd int, p []byte) (n int, err error) {
	buf, count := splitSlice(p)
	n = libc_read(fd, buf, uint(count))
	if n < 0 {
		err = getErrno()
	}
	return
}

func Seek(fd int, offset int64, whence int) (off int64, err error) {
	return 0, ENOSYS // TODO
}

func Open(path string, flag int, mode uint32) (fd int, err error) {
	data := append([]byte(path), 0)
	fd = libc_open(&data[0], flag, mode)
	if fd < 0 {
		err = getErrno()
	}
	return
}

func Mkdir(path string, mode uint32) (err error) {
	return ENOSYS // TODO
}

func Unlink(path string) (err error) {
	return ENOSYS // TODO
}

func Kill(pid int, sig Signal) (err error) {
	return ENOSYS // TODO
}

func Getuid() int  { return int(libc_getuid()) }
func Getgid() int  { return int(libc_getgid()) }
func Geteuid() int { return int(libc_geteuid()) }
func Getegid() int { return int(libc_getegid()) }
func Getpid() int  { return int(libc_getpid()) }
func Getppid() int { return int(libc_getppid()) }

func Getenv(key string) (value string, found bool) {
	data := append([]byte(key), 0)
	raw := libc_getenv(&data[0])
	if raw == nil {
		return "", false
	}

	ptr := uintptr(unsafe.Pointer(raw))
	for size := uintptr(0); ; size++ {
		v := *(*byte)(unsafe.Pointer(ptr))
		if v == 0 {
			src := *(*[]byte)(unsafe.Pointer(&sliceHeader{buf: raw, len: size, cap: size}))
			return string(src), true
		}
		ptr += unsafe.Sizeof(byte(0))
	}
}

func splitSlice(p []byte) (buf *byte, len uintptr) {
	slice := (*sliceHeader)(unsafe.Pointer(&p))
	return slice.buf, slice.len
}

// ssize_t write(int fd, const void *buf, size_t count)
//export write
func libc_write(fd int, buf *byte, count uint) int

// char *getenv(const char *name);
//export getenv
func libc_getenv(name *byte) *byte

// ssize_t read(int fd, void *buf, size_t count);
//export read
func libc_read(fd int, buf *byte, count uint) int

// int open(const char *pathname, int flags, mode_t mode);
//export open
func libc_open(pathname *byte, flags int, mode uint32) int

// int close(int fd)
//export close
func libc_close(fd int) int

// uid_t getuid(void)
//export getuid
func libc_getuid() int32

// gid_t getgid(void)
//export getgid
func libc_getgid() int32

// uid_t geteuid(void)
//export geteuid
func libc_geteuid() int32

// gid_t getegid(void)
//export getegid
func libc_getegid() int32

// gid_t getpid(void)
//export getpid
func libc_getpid() int32

// gid_t getppid(void)
//export getppid
func libc_getppid() int32
