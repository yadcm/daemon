package listen

import (
	"net"
	"net/netip"
	"yadcmd/internal/app/config"
)

func Unix(conf *config.ServeAnyUnix) (net.Listener, error) {
	return net.Listen("unix", conf.Socket)
}
func TCP(conf *config.ServeAnyTCP) (net.Listener, error) {
	return net.Listen("tcp", net.TCPAddrFromAddrPort(
		netip.AddrPortFrom(
			netip.AddrFrom4(
				conf.IP,
			),
			conf.Port,
		),
	).String())
}
