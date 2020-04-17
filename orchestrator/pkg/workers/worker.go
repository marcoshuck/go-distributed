package workers

import (
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/managers"
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/nodes"
	"sync"
)

type Worker interface {
	GetManager() managers.Manager
	SetManager(manager managers.Manager)
}

type worker struct {
	nodes.Node
	lockManager sync.Mutex
	Manager managers.Manager
}

type Workers []Worker

func NewWorker(name, data string) Worker {
	var w Worker
	w = &worker{
		Node:    nodes.NewNode(name, data),
	}
	return w
}

func (w *worker) GetManager() managers.Manager {
	return w.Manager
}

func (w *worker) SetManager(manager managers.Manager) {
	w.lockManager.Lock()
	defer w.lockManager.Unlock()
	w.Manager = manager
}