package zoo

import (
	"github.com/go-zookeeper/zk"
	"time"
)

const ADDRESS = "localhost:2181"
const TIMEOUT = 3000

func Connect() ([]string, *zk.Stat, <-chan zk.Event) {
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := conn.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	return children, stat, ch
}