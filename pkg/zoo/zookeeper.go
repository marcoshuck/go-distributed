package zoo

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"sort"
	"time"
)

const Address = "localhost:2181"
const Timeout = 3000
const ElectionNamespace = "/election"

func CreateConnection() (*zk.Conn, <-chan zk.Event, error) {
	conn, ev, err := zk.Connect([]string{Address}, Timeout* time.Millisecond)
	if err != nil {
		return nil, nil, err
	}
	return conn, ev, nil
}

func CreateNode(z *zk.Conn) (string, error) {
	znodePrefix := fmt.Sprintf("%s/c_", ElectionNamespace)
	zNodeName, err := z.Create(znodePrefix, []byte{}, 3, zk.WorldACL(zk.PermAll))
	if err != nil {
		return "", err
	}
	return zNodeName, nil
}

func ElectMaster(z *zk.Conn, currentNode string) error {
	currentNode = parseNode(currentNode)
	children, _, err := z.Children(ElectionNamespace)
	if err != nil {
		return err
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i] < children[j]
	})

	master := children[0]
	fmt.Printf("Current node: [%s]. Master: [%s]\n", currentNode, master)
	if master == currentNode {
		fmt.Printf("[%s] I'm the master\n", master)
		return nil
	}
	fmt.Printf("[%s] I'm a slave. The master is: %s\n", currentNode, master)
	return nil
}

func parseNode(nodeName string) string {
	var node string
	fmt.Sscanf(nodeName, ElectionNamespace + "/%s", &node)
	return node
}