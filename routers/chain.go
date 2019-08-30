package routers

import (
	"block-chain/service"
	"github.com/gin-gonic/gin"
)

type ChainRouter struct {
}

func (s *ChainRouter) Chains(c *gin.Context) {
	blockService := new(service.BlockService)
	blockChain := blockService.BlockChain()

	c.JSON(200, blockChain)
}
