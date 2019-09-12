package libs

import (
	"bytes"
	"encoding/gob"
	"log"
)

// TXOutput represents a transaction output(交易输出不是指出款，而是指整个交易最后生产的钱去哪里)
// 比如：你拿10元钱去买3元钱的汽水，那么10元钱是输入，而输出则是3元输出给老板，7元输出给你自己。
type TXOutput struct {
	Value      int       //
	PubKeyHash []byte
}

// Lock signs the output(签署交易输出)
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
// (检测输入的公钥是否可以使用未使用的交易输出)
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput(创建一个交易输出)
// value 交易金额
// address 交易输出的收款人的钱包地址
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

// TXOutputs collects TXOutput
type TXOutputs struct {
	Outputs []TXOutput
}

// Serialize serializes TXOutputs
func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DeserializeOutputs deserializes TXOutputs
func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
