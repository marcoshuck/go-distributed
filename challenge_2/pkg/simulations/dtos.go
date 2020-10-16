package simulations

type CreateSimulationInput struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Image string `json:"image"`
}

type CreateSimulationOutput struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Image string `json:"image"`
}
