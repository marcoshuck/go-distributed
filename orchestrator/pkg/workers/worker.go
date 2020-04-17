package workers

import (
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/managers"
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/nodes"
)

type Worker interface {

}

type worker struct {
	nodes.Node
	Manager managers.Manager
}

type Workers []Worker
