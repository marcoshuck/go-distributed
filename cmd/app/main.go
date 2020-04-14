package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/marcoshuck/go-distributed/pkg/zoo"
)

const ElectionNamespace = "/election"

func main() {
	z, events, err := zoo.CreateConnection()
	if err != nil {
		fmt.Errorf("[MAIN] Error while creating connection for ZooKeeper.\nError: [%v]\n", err)
		panic(err)
	}
	fmt.Printf("[MAIN] ZooKeeper connection created. Session ID: [%d]\n", z.SessionID())
	nodeName, err := zoo.CreateNode(z)
	if err != nil {
		fmt.Errorf("[MAIN] Error creating ZooKeeper Node.\nError: [%v]\n", err)
		panic(err)
	}
	zoo.ElectMaster(z, nodeName)
	run(events)
	defer z.Close()
}

func run(events <- chan zk.Event) {
	for {
		select {
		case e := <- events:
			fmt.Printf("Event received from server: [%s]\n", e.Server)
		}
	}
}