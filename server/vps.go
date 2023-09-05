package server

import "MayMistery/LightChain/utils"

func vpsPublicIP() {
	ini := utils.ReadIni()
	utils.RemotePortForword(ini)
}
