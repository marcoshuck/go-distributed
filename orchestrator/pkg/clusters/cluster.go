package clusters

import (
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/managers"
	"github.com/marcoshuck/go-distributed/orchestrator/pkg/workers"
)

type Cluster interface {

}

type cluster struct {
	managers managers.Managers
	workers workers.Workers
}

func NewCluster(managers managers.Managers, workers workers.Workers) Cluster {
	var c Cluster
	c = &cluster{
		managers: managers,
		workers:  workers,
	}
	return c
}