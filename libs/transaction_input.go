package libs

import (
	"bytes"
)

// TXInput represents a transaction input(交易输出不是指收款，而是指整个交易最开始的输入金额)
// 比如：你拿10元钱去买3元钱的汽水，那么10元钱是输入，而输出则是3元输出给老板，7元输出给你自己。
type TXInput struct {
	Txid      []byte  // 这个Txid是否与transaction中的ID相等呢？
	Vout      int     // 一个输出索引（vout），用于标识来自该交易的哪个UTXO被引用（第一个为零）
	Signature []byte  // 交易开始输入金额需要调用出款人的私钥签名
	PubKey    []byte  // 公钥非Hash，用于验证出款人的签名信息是否正确
}

// UsesKey 检测交易输入的公钥hash是否正确
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

