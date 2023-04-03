package net

import (
	"fmt"
	"net"
	"net/netip"
	"time"

	"github.com/x0f5c3/wiz-go/models"
)

const DEFAULTWAITTIME = 5.0
const RegisterMessage = `{"method":"registration","params":{"phoneMac":"AAAAAAAAAAAA","register":false,"phoneIp":"1.2.3.4","id":"1"}}`

type BroadcastProtocol struct {
	Reg           models.BulbRegistry
	BroadcastAddr *net.UDPAddr
	Transport     *net.UDPConn
}

func NewProtocol(addr ...string) (*BroadcastProtocol, error) {
	rAddr := "255.255.255.255"
	if len(addr) > 0 {
		rAddr = addr[0]
	}
	bAddr, err := netip.ParseAddr(rAddr)
	if err != nil {
		return nil, err
	}
	pAddr := net.UDPAddrFromAddrPort(netip.AddrPortFrom(bAddr, PORT))
	transport, err := CreateUDP()
	if err != nil {
		return nil, err
	}
	reg := models.NewRegistry()
	res := new(BroadcastProtocol)
	res.Reg = reg
	res.BroadcastAddr = pAddr
	res.Transport = transport
	return res, nil
}

func (b *BroadcastProtocol) Discover() (string, error) {
	_, err := b.Transport.WriteToUDP([]byte(RegisterMessage), b.BroadcastAddr)
	if err != nil {
		return "", err
	}
	now := time.NewTimer(time.Second * DEFAULTWAITTIME)
	for {
		select {
		case <-now.C:
			return "", fmt.Errorf("exceeded wait time")
		default:
			buf := make([]byte, 1024)
			n, ad, err := b.Transport.ReadFromUDP(buf)
			if err != nil {
				return "", err
			}
			our := CheckIfOur(ad.String())
			if !our {
				msg := string(buf[:n])
				return msg, nil
			}
		}
	}
}
