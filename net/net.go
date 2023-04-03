package net

import (
	"fmt"
	"net"
	"net/netip"
	"syscall"

	"github.com/pterm/pterm"
	"golang.org/x/exp/slices"
)

const PORT = 38999

var localAddrs = func() []string {
	ad, err := net.InterfaceAddrs()
	pterm.Fatal.PrintOnErrorf("Failed to get our addresses %s", err)
	res := make([]string, len(ad))
	for i, v := range ad {
		res[i] = v.String()
	}
	return res
}()

func CheckIfOur(addr string) bool {
	return slices.Contains(localAddrs, addr)
}

func CreateUDP() (*net.UDPConn, error) {
	laddr := net.UDPAddrFromAddrPort(netip.AddrPortFrom(netip.IPv4Unspecified(), PORT))
	sock, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return nil, err
	}
	sysSock, err := sock.SyscallConn()
	if err != nil {
		return nil, err
	}
	err = sysSock.Control(func(fd uintptr) {
		err := setSockOpt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
		if err != nil {
			fmt.Println("Failed to setsockopt due to " + err.Error())
			return
		}
		err = setSockOpt(fd, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1)
		if err != nil {
			fmt.Println("Failed to setsockopt due to " + err.Error())
			return
		}
	})
	return sock, err
}
