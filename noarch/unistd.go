package noarch

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func Isatty(fd int) int {
	_, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	// TODO need test
	if err != nil {
		return 0
	}
	return 1
}

func Pipe(p []int) int {
	err := unix.Pipe(p)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return 0
}

func Read(fd int, p []byte, num int) int {
	n, err := unix.Read(fd, p)
	_ = num
	if err != nil {
		return -1
	}
	return n
}

func Write(fd int, p []byte, num int) int {
	n, err := unix.Write(fd, p)
	_ = num
	if err != nil {
		return -1
	}
	return n
}

type SsizeT int32
