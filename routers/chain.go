package routers

import (
	"block-chain/common"
	"block-chain/config"
	"block-chain/libs"
	"github.com/gin-gonic/gin"
)

type ChainRouter struct {
}

func (s *ChainRouter) Chains(c *gin.Context) {
	bc := libs.NewBlockchain(config.NodeID)
	bci := bc.Iterator()

	var blocks []*libs.Block
	for {
		block := bci.Next()
		blocks = append(blocks, block)
		//for _, tx := range block.Transactions {
		//	if bytes.Compare(tx.ID, ID) == 0 {
		//		return *tx, nil
		//	}
		//}
		//
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	common.Dispatch(c, "0000", "success", blocks)
}

// 创建创世块
func (s *ChainRouter) NewGenesis(c *gin.Context) {
	// 创建创世块
	bc := libs.CreateBlockchain(config.Gaddress, config.NodeID)
	// 更新
	UTXOSet := libs.UTXOSet{bc}
	UTXOSet.Reindex()

	common.Dispatch(c, "0000", "success", bc)
}
