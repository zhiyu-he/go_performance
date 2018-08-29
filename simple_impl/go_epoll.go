package simple_impl

import (
	"syscall"
	"fmt"
)

const (
	EPOLLET = 1 << 31
	MaxEpollEvents = 32
)


func echo(fd int) {
	defer syscall.Close(fd)
	var buf[32*1024]byte
	for {
		nBytes, e := syscall.Read(fd, buf[:])
		if e != nil {
			break
		}
		if nBytes > 0 {
			fmt.Printf(">>>> %s", buf)
			syscall.Write(fd, buf[:nBytes])
			fmt.Printf("<<< %s", buf)
		}
	}
}


