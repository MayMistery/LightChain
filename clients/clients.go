package clients

import (
	"MayMistery/LightChain/cmd"
	"MayMistery/LightChain/utils"
)

func ExecClient(cfg cmd.Config) {
	if cfg.Dep {
		utils.GenDeployment(cfg)
	}
	ChainNode(cfg)
}
