package nodes

import (
	"github.com/go-zookeeper/zk"
	"github.com/google/uuid"
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/conn"
	"time"
)

type Node interface {
	Connect(host string, timeout uint)
	Restart()
	Run()
	Stop()
	Kill()
}

type node struct {
	UUID string
	Name string
	Role Role
	Status Status
	Data string
	conn *zk.Conn
	events <- chan zk.Event

}

type Nodes []Node

func NewNode(name, data string) Node {
	var n Node
	uuid := uuid.Must(uuid.NewUUID())
	n = &node{
		UUID: uuid.String(),
		Name: name,
		Data: data,
		Status: StatusCreated,
	}
	return n
}

// Connect connects the node to the given host and set the timeout connection with the given timeout in seconds.
func (n *node) Connect(host string, timeout uint) {
	n.Status = StatusConnecting
	attempts := 10
	var i int
	var err error
	for i = 1; i <= attempts; i++ {
		n.conn, n.events, err = conn.Connect(host, timeout * uint(time.Second.Milliseconds()))
		if err != nil {
			n.Status = StatusErrConnecting
			time.Sleep(time.Duration(i) * time.Second)
			if i == attempts {
				return
			}
			continue
		}
		break
	}
	n.Status = StatusConnected
}

// Restart requests the Node to be restarted.
func (n *node) Restart() {
	n.Status = StatusRestarting
}

// Run requests the node to run their internal job.
func (n *node) Run() {
	n.Status = StatusRunning
}

// Stop requests the node to fully stop but avoid being deleted.
func (n *node) Stop() {
	n.Status = StatusStopping
}

// Kill requests the node to be deleted.
func (n *node) Kill() {
	n.Status = StatusKilling

	conn.Disconnect(n.conn)

	n.Status = StatusDead
}