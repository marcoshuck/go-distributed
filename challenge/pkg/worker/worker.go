package worker

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/go-distributed/challenge/pkg/db"
	"net/http"
)

type Input struct {
	User string `json:"user"`
	Value int64 `json:"value"`
}

func Sum(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		var body Input
		ctx.BindJSON(&body)

		result, err := db.Add(DB, ip, body.User, body.Value)

		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func Get(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.Param("user")

		result, err := db.Get(DB, user)

		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}