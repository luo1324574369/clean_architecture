package main

import (
	"log"

	"github.com/luo1324574369/clean_architecture/service"
)

func main() {
	if err := service.Server.Run(); err != nil {
		log.Fatalf("server err %v", err)
	}
}
