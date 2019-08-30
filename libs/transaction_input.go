package libs

import (
	"bytes"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte  // 这个Txid是否与transaction中的ID相等呢？
	Vout      int     // 交易金额
	Signature []byte  //
	PubKey    []byte  // 公钥非Hash
}

// UsesKey 检测接收人的公钥是否正确
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// HashPubKey 将公钥进行Hash化处理
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}
