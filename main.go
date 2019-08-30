package main

import (
	"block-chain/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	Routers(router)
	log.Panic(router.Run(":8080"))
}

func Routers(router *gin.Engine) {
	chainRouter := routers.ChainRouter{}
	router.GET("/chain", chainRouter.Chains)
}

