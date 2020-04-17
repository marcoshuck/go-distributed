package zoo

import (
	"fmt"
	"github.com/go-zookeeper/zk"
)

func WatchNode(z *zk.Conn, target string) (parent, children <-chan zk.Event) {
	ok, _, err := z.Exists(target)
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil, nil
	}
	data, _, parent, err := z.GetW(target)
	if err != nil {
		panic(err)
	}
	childrenNames, _, children, err := z.ChildrenW(target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[WATCHER] Children: %v | Data: %s\n", childrenNames, string(data))
	return
}
