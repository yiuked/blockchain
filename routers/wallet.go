package routers

import (
	"block-chain/common"
	"block-chain/config"
	"block-chain/libs"
	"github.com/gin-gonic/gin"
	"log"
)

type WalletRouter struct {
}

// 钱包列表
func (s *WalletRouter) Wallets(c *gin.Context) {
	wallets, _ := libs.NewWallets(config.NodeID)
	common.Dispatch(c, "0000", "success", wallets)
}

// 创建钱包
func (s *WalletRouter) NewWallet(c *gin.Context) {
	wallets, _ := libs.NewWallets(config.NodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(config.NodeID)
	common.Dispatch(c, "0000", "success", address)
}

func (s *WalletRouter) Balance(c *gin.Context) {
	address := c.Query("address")
	if !libs.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := libs.NewBlockchain(config.NodeID)
	UTXOSet := libs.UTXOSet{bc}
	UTXOSet.Reindex()

	balance := 0
	pubKeyHash := libs.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	common.Dispatch(c, "0000", "success", balance)
}

func (s *WalletRouter) Transfer(c *gin.Context) {
	from := c.PostForm("from")
	to := c.PostForm("to")
	amount := common.StrToInt(c.PostForm("amount"))

	bc := libs.NewBlockchain(config.NodeID)
	wallets, err := libs.NewWallets(config.NodeID)
	if err != nil {
		common.Dispatch(c, "10401", "Wallets not found", nil)
	}

	payerWallet := wallets.GetWallet(from)

	UTXOSet := libs.UTXOSet{bc}
	// 创建一笔交易
	tx := libs.NewUTXOTransaction(&payerWallet, to, amount, &UTXOSet)
	// 创建币基
	cbTx := libs.NewCoinbaseTX(from, "")
	txs := []*libs.Transaction{cbTx, tx}
	// 自己挖矿
	newBlock := bc.MineBlock(txs)

	UTXOSet.Update(newBlock)

	common.Dispatch(c, "0000", "success", newBlock)
}
