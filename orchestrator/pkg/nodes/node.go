package nodes

type Node interface {

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
	}
	return n
}