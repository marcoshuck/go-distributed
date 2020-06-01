package main

import "github.com/marcoshuck/go-distributed/challenge/pkg/server"

func main() {
	dsn := server.NewDSN("db")
	s := server.NewServer("", 3000, dsn)
	err := s.Listen()
	if err != nil {
		panic(err)
	}
}
