package handler

import (
	"github.com/syyongx/llog/types"
	"net"
)

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

func NewNet(bufferSize int, persistent bool, level int, bubble bool) *Net {
	network := &Net{
		BufferSize: bufferSize,
		Persistent: persistent,
	}
	network.SetLevel(level)
	network.SetBubble(bubble)
	return network
}

// Handles a record.
func (n *Net) Handle(record *types.Record) bool {
	if !n.IsHandling(record) {
		return false
	}
	if n.processors != nil {
		n.ProcessRecord(record)
	}
	err := n.GetFormatter().Format(record)
	if err != nil {
		return false
	}
	n.Write(record)

	return false == n.GetBubble()
}

// Handles a set of records.
func (n *Net) HandleBatch(records []*types.Record) {
	for _, record := range records {
		n.Handle(record)
	}
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
