package main

import (
	"github.com/marcoshuck/go-distributed/challenge/pkg/server"
	"os"
)

func main() {
	dbHost := os.Getenv("DATABASE_HOST")
	dsn := server.NewDSN(dbHost)

	s := server.NewServer("", 3000, dsn)
	err := s.Listen()
	if err != nil {
		panic(err)
	}
}
