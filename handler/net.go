package handler

import "net"

type Net struct {
	Handler
	Processable
	Formattable

	// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
	// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
	// (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
	// "unixpacket".
	Network    string
	Address    string
	Persistent bool
	BufferSize int
	conn       net.Conn
	close      chan bool
}

func NewNet(bufferSize int, persistent bool) *Net {
	n := &Net{
		BufferSize: bufferSize,
		Persistent: persistent,
		close:      make(chan bool, 0),
	}
	return n
}

func (n *Net) connect() error {
	if n.conn != nil {
		n.Close()
		n.conn = nil
	}
	conn, err := net.Dial(n.Network, n.Address)
	if err != nil {
		return err
	}
	if tcp, ok := conn.(*net.TCPConn); ok {
		tcp.SetKeepAlive(true)
	}

	n.conn = conn
	return nil
}

func (n *Net) Close() {
	<-n.close
}
