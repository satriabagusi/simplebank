/*
Author: Satria Bagus(satria.bagus18@gmail.com)
main.go (c) 2023
Desc: description
Created:  2023-06-29T16:52:10.361Z
Modified: !date!
*/

package cmd

import (
	"log"

	"github.com/satriabagusi/simplebank/config"
	"github.com/satriabagusi/simplebank/pkg/server"
	"golang.org/x/tools/cmd/getgo/server"
)

func main() {
	cfg := config.NewConfig()

	cfg.Load()

	srv := server.NewServer()

	err := srv.Initialize(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("failed to initialize server %v", err)
	}

	err = srv.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
