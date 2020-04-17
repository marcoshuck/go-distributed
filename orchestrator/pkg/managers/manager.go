package managers

import (
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/nodes"
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/store"
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/workers"
)

type Manager interface {
}

type manager struct {
	nodes.Node
	store *store.Store
	workers *workers.Workers
}

type Managers []Manager