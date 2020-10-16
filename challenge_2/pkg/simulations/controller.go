package simulations

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/queue"
	"github.com/marcoshuck/go-distributed/challenge_2/pkg/sender"
	"log"
	"net/http"
)

type Controller interface {
	Create(ctx *gin.Context)
}

type controller struct {
	sender sender.Sender
}

func (c *controller) Create(ctx *gin.Context) {
	var s CreateSimulationInput

	err := ctx.BindJSON(&s)
	if err != nil {
		log.Println("error while parsing json body:", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	b, err := json.Marshal(&s)
	if err != nil {
		log.Println("error while marshaling simulation data:", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = c.sender.Send("", queue.SimulationRequests, b, "text/json")
	if err != nil {
		log.Println("error while sending simulation to the queue of simulation requests:", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, s)
}

func NewController(sender sender.Sender) Controller {
	return &controller{
		sender: sender,
	}
}
