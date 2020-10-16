package simulations

import (
	"log"
	"time"
)

type Repository interface {
	Create(simulation Simulation) (Simulation, error)
}

type repository struct {
	simulations []Simulation
}

func (r *repository) Create(simulation Simulation) (Simulation, error) {
	simulation.ID = uint(len(r.simulations))
	simulation.CreatedAt = time.Now()
	simulation.UpdatedAt = time.Now()
	r.simulations = append(r.simulations, simulation)
	log.Printf("simulation [%d] created\n", simulation.ID)
	return simulation, nil
}

func NewRepository() Repository {
	return &repository{}
}
