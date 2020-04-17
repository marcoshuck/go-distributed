package zoo

import (
	"github.com/go-zookeeper/zk"
	"time"
)

func CreateConnection() (*zk.Conn, <-chan zk.Event, error) {
	conn, ev, err := zk.Connect([]string{Address}, Timeout* time.Millisecond)
	if err != nil {
		return nil, nil, err
	}
	return conn, ev, nil
}