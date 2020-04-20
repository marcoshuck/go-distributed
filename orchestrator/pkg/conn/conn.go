package conn

import (
	"github.com/go-zookeeper/zk"
	"time"
)

// Connect creates a new Zookeeper connection.
// Receives the address to connect to, and the timeout in milliseconds.
func Connect(address string, timeout uint) (*zk.Conn, <-chan zk.Event, error) {
	conn, ev, err := zk.Connect([]string{address}, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return nil, nil, err
	}
	return conn, ev, nil
}

// Disconnection triggers the closing of the given Zookeeper connection.
func Disconnect(conn *zk.Conn) {
	conn.Close()
}