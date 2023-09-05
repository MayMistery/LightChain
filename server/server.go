package server

import (
	"MayMistery/LightChain/cmd"
	"MayMistery/LightChain/utils"
	"net"
)

func ExecServer(cfg cmd.Config) {
	if cfg.Local {
		vpsPublicIP()
	}
	fileServer(cfg)

	Addr := getAddr(cfg)
	utils.SockServer(Addr)

	utils.GenDeployment(cfg)
}

func getAddr(cfg cmd.Config) net.Addr {
	return nil
}
