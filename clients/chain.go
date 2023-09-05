package clients

import (
	"MayMistery/LightChain/cmd"
	"MayMistery/LightChain/utils"
	"net"
)

func ChainNode(cfg cmd.Config) {
	AddrList := parseChain(cfg.Chain)
	utils.SockClient(AddrList[0]) //TODO for temp, not correct
	utils.SockServer(AddrList[0])
}

func parseChain(chain string) []net.Addr {
	return nil
}
