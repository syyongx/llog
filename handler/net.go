package handler

import (
	"github.com/syyongx/llog/types"
	"net"
)

// Net handler struct definition
type Net struct {
	Processing

	// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
	// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
	// (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
	// "unixpacket".
	Network    string
	Address    string
	Persistent bool
	BufferSize int
	conn       net.Conn
}

// NewNet new net handler
func NewNet(bufferSize int, persistent bool, level int, bubble bool) *Net {
	n := &Net{
		BufferSize: bufferSize,
		Persistent: persistent,
	}
	n.SetLevel(level)
	n.SetBubble(bubble)
	n.Writer = n.Write
	return n
}

// Write to network.
func (n *Net) Write(record *types.Record) {
	if !n.Persistent {
		if err := n.connect(); err != nil {
			//
		}
		defer n.conn.Close()
	}
	_, err := n.conn.Write(record.Formatted.Bytes())
	if err != nil {
		//
	}
}

// Close connect.
func (n *Net) Close() {
	if n.Persistent {
		n.conn.Close()
	}
}

// connect
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
