package models

import "time"

type BlockHeader struct {
	Index          int
	Version        string
	Timestamp      time.Time
	PrevBlockHash  string
	MerkleRootHash string
	BlockHash      string
	Nonce          int
}

type Block struct {
	BlockHeader  BlockHeader
	AccountCount int
	Accounts     []Account
}

type Account struct {
}
