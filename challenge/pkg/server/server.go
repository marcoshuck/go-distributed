package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/marcoshuck/go-distributed/challenge/pkg/db"
	"github.com/marcoshuck/go-distributed/challenge/pkg/worker"
	"net/http"
	"time"
)

type Server struct {
	Address string
	Port uint
	DB *gorm.DB
	*gin.Engine
}

func NewDSN(ip string) string {
	user := "root"
	password := "test"
	dbname := "challenge"
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, ip, dbname)
}

func NewServer(address string, port uint, dsn string) *Server {
	var database *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		database, err = gorm.Open("mysql", dsn)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}
	if err != nil {
		panic("[server.go:38] Error connecting to the database.")
	}
	s := &Server{
		Address: address,
		Port:    port,
		Engine:  gin.New(),
		DB:		 database,
	}
	s.DB.LogMode(true)
	db.Migrate(database)
	s.Handle("GET", "/sum/:user", worker.Get(s.DB))
	s.Handle("POST", "/sum", worker.Sum(s.DB))
	s.Handle("GET", "/healtz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server is up")
	})
	return s
}

func (s *Server) Listen() error {
	return s.Run(fmt.Sprintf("%s:%d", s.Address, s.Port))
}