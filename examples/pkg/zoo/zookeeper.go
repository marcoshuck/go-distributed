package zoo

import (
	"fmt"
	"github.com/go-zookeeper/zk"
)

const Address = "localhost:2181"
const Timeout = 3000

func CreatePrefixNode(z *zk.Conn, prefix string) (string, error) {
	znodePrefix := fmt.Sprintf("%s/c_", ElectionNamespace)
	zNodeName, err := z.Create(znodePrefix, []byte{}, 3, zk.WorldACL(zk.PermAll))
	if err != nil {
		return "", err
	}
	return zNodeName, nil
}

func CreateNode(z *zk.Conn, name string) (string, error) {
	zNodeName, err := z.Create(name, []byte{}, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		return "", err
	}
	return zNodeName, nil
}

