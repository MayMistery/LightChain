package main

import (
	"MayMistery/LightChain/cmd"
	"MayMistery/LightChain/server"
	"testing"
)

func TestFileServer(t *testing.T) {
	cfg := cmd.Config{Host: "0.0.0.0", Port: 9000}
	server.FileServer(cfg)
}

func TestPortforward(t *testing.T) {

}
