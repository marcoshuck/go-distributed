package simulations

import "gorm.io/gorm"

type Simulation struct {
	gorm.Model
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Image string `json:"image"`
}
