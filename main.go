package main

import (
	"block-chain/libs"
	"block-chain/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	pubKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCh5Nk2GLiyQFMIU+h3OEA4UeFbu3dCH5sjd/sLTxxvwjXq7JLqJbt2rCIdzpAXOi4jL+FRGQnHaxUlHUBZsojnCcHvhrz2knV6rXNogt0emL7f7ZMRo8IsQGV8mlKIC9xLnlOQQdRNUssmrROrCG99wpTRRNZjOmLvkcoXdeuaCQIDAQAB"
	libs.CreateBlockchain(pubKey, "1001")
	//router := gin.Default()
	//Routers(router)
	//log.Panic(router.Run(":8080"))
}

func Routers(router *gin.Engine) {
	chainRouter := routers.ChainRouter{}
	router.GET("/chain", chainRouter.Chains)
}
