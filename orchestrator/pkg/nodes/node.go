package nodes

type Node interface {
	Restart()
	Run()
	Pause()
	Stop()
	Kill()
}

type node struct {
	UUID string
	Name string
	Role Role
	Status Status
	Data string
}

type Nodes []Node

func NewNode(name, data string) Node {
	var n Node
	n = &node{
		UUID: "",
		Name: name,
		Data: data,
		Status: StatusCreated,
	}
	return n
}

// Restart requests the Node to be restarted.
func (n *node) Restart() {
	n.Status = StatusRestarting
}

// Run requests the node to run their internal job.
func (n *node) Run() {
	n.Status = StatusRunning
}

// Pause requests the node to pause its internal job.
func (n *node) Pause() {
	n.Status = StatusPaused
}

// Stop requests the node to fully stop but avoid being deleted.
func (n *node) Stop() {
	n.Status = StatusStopped
}

// Kill requests the node to be deleted.
func (n *node) Kill() {
	n.Status = StatusDead
}