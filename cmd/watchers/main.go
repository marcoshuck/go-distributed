package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/marcoshuck/go-distributed/pkg/zoo"
)

func main() {
	z, _, err := zoo.CreateConnection()
	if err != nil {
		fmt.Errorf("[MAIN] Error while creating connection for ZooKeeper.\nError: [%v]\n", err)
		panic(err)
	}
	fmt.Printf("[MAIN] ZooKeeper connection created. Session ID: [%d]\n", z.SessionID())
	// nodeName, err := zoo.CreateNode(z, "/test")
	if err != nil {
		fmt.Errorf("[MAIN] Error creating ZooKeeper Node.\nError: [%v]\n", err)
		panic(err)
	}
	parentEvents, childrenEvents := zoo.WatchNode(z, "/test")
	go process(parentEvents)
	go process(childrenEvents)
	run(z, "/test", parentEvents)
	defer z.Close()
}


// run keeps the main thread running until something unexpected with ZooKeeper happens.
func run(z *zk.Conn, node string, events <- chan zk.Event) {
	for {
		_, err := z.Sync(node)
		if err != nil {
			return
		}
	}
}

func process(events <- chan zk.Event) {
	for {
		select {
		case e := <-events:
			switch e.Type {
			case zk.EventNodeCreated:
				fmt.Printf("[EVENT] Node created. Path: %s\n", e.Path)
				continue
			case zk.EventNodeDeleted:
				fmt.Printf("[EVENT] Node deleted. Path: %s\n", e.Path)
				continue
			case zk.EventNodeDataChanged:
				fmt.Printf("[EVENT] Node data has changed. Path: %s\n", e.Path)
				continue
			case zk.EventNodeChildrenChanged:
				fmt.Printf("[EVENT] Child data has changed. Path: %s\n", e.Path)
				continue
			}
		}
	}
}