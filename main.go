package main

import (
	"MayMistery/LightChain/clients"
	"MayMistery/LightChain/cmd"
	"MayMistery/LightChain/server"
)

func main() {
	var cfg cmd.Config

	cmd.Flag(&cfg)

	if cfg.Server {
		server.ExecServer(cfg)
	} else if cfg.Client {
		clients.ExecClient(cfg)
	} else if cfg.Alive {
		clients.PingScan(cfg)
	} else {
	}
}
