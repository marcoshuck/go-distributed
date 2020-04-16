package election2

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"sort"
)

const Namespace = "/election"

func ElectMaster(z *zk.Conn, currentNode string) error {
	parsedNode := parseNode(currentNode)
	children, _, err := z.Children(Namespace)
	if err != nil {
		return err
	}
	if !sort.SliceIsSorted(children, descendingOrder(children)) {
		sort.SliceStable(children, descendingOrder(children))
	}

	master := children[0]
	fmt.Printf("Current node: [%s]. Master: [%s]\n", parsedNode, master)
	if master == parsedNode {
		fmt.Printf("[%s] I'm the master\n", master)
		return nil
	}
	fmt.Printf("[%s] I'm a slave. The master is: %s\n", parsedNode, master)

	var predecessorIndex int
	for i, c := range children {
		if parsedNode == c {
			predecessorIndex = i - 1
		}
		ok, stat, existEvents, err := z.ExistsW(fmt.Sprintf("%s/%s", Namespace, children[predecessorIndex]))
	}
	return nil
}

func descendingOrder(slice []string) func(i, j int) bool {
	return func(i, j int) bool {
		return slice[i] < slice[j]
	}
}

func parseNode(nodeName string) string {
	var node string
	fmt.Sscanf(nodeName, Namespace + "/%s", &node)
	return node
}

