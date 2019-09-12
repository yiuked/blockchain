package routers

import (
	"block-chain/common"
	"block-chain/libs"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type WalletRouter struct {
}

// 钱包列表
func (s *WalletRouter) Wallets(c *gin.Context) {
	wallets, _ := libs.NewWallets(os.Getenv("NODE_ID"))
	common.Dispatch(c, "0000", "success", wallets)
}

// 创建钱包
func (s *WalletRouter) NewWallet(c *gin.Context) {
	wallets, _ := libs.NewWallets(os.Getenv("NODE_ID"))
	address := wallets.CreateWallet()
	wallets.SaveToFile(os.Getenv("NODE_ID"))
	common.Dispatch(c, "0000", "success", address)
}

func (s *WalletRouter) Balance(c *gin.Context) {
	address := c.Query("address")
	if !libs.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := libs.NewBlockchain(os.Getenv("NODE_ID"))
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

	bc := libs.NewBlockchain(os.Getenv("NODE_ID"))
	wallets, err := libs.NewWallets(os.Getenv("NODE_ID"))
	if err != nil {
		common.Dispatch(c, "10401", "Wallets not found", nil)
	}

	payerWallet := wallets.GetWallet(from)

	UTXOSet := libs.UTXOSet{Blockchain: bc}
	// 创建一笔交易
	tx := libs.NewUTXOTransaction(&payerWallet, to, amount, &UTXOSet)

	libs.SendTx(libs.KnownNodes[0], tx)

	common.Dispatch(c, "0000", "success", tx)
}

func (s *WalletRouter) UTXOs(c *gin.Context) {
	address := c.Query("address")

	bc := libs.NewBlockchain(os.Getenv("NODE_ID"))
	wallets, err := libs.NewWallets(os.Getenv("NODE_ID"))
	if err != nil {
		common.Dispatch(c, "10401", "Wallets not found", nil)
	}

	wallet := wallets.GetWallet(address)
	UTXOSet := libs.UTXOSet{Blockchain: bc}

	UTXOs := UTXOSet.FindUTXO(libs.HashPubKey(wallet.PublicKey))

	common.Dispatch(c, "0000", "success", UTXOs)
}

func (s *WalletRouter) MeePool(c *gin.Context) {
	common.Dispatch(c, "0000", "success", libs.MeePool)
}
