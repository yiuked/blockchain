package service

import (
	"block-chain/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

var blockHeaders []models.BlockHeader

const AppVersion = "0.0.1"

func init() {
	if len(blockHeaders) == 0 {
		service := new(BlockService)
		//TODO  加载本地数据,验证数据真实性，如何本地缓存为空，则创建创世块
		var header models.BlockHeader
		header.Index = 0
		header.Version = AppVersion
		header.Timestamp = time.Now()
		header.PrevBlockHash = fmt.Sprintf("%064d", 0)
		header.MerkleRootHash = fmt.Sprintf("%064d", 0)
		service.CalculateHash(&header)

		blockHeaders = append(blockHeaders, header)
	}
}

type BlockService struct {
}

func (s *BlockService) CalculateHash(header *models.BlockHeader) *models.BlockHeader {
	record := string(header.Index) + header.Timestamp.String() + header.PrevBlockHash + header.MerkleRootHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	header.BlockHash = hex.EncodeToString(hashed)
	header.Nonce += 1
	return header
}

func (s *BlockService) GenerateBlock(preHeader models.BlockHeader, account []models.Account) models.Block {
	service := new(BlockService)

	var header models.BlockHeader
	header.Index = preHeader.Index + 1
	header.Version = AppVersion
	header.Timestamp = time.Now()
	header.PrevBlockHash = preHeader.BlockHash
	header.MerkleRootHash = fmt.Sprintf("%064d", 0)

	service.CalculateHash(&header)

	block := models.Block{
		BlockHeader:  header,
		AccountCount: len(account),
		Accounts:     account,
	}

	return block
}

func (s *BlockService) IsBlockValid(header models.BlockHeader, preHeader models.BlockHeader) bool {
	if preHeader.Index+1 != header.Index {
		return false
	}

	if preHeader.BlockHash != header.PrevBlockHash {
		return false
	}

	return true
}

func (s *BlockService) ReplaceChain(remoteBlocks []models.Block, localBlock []models.Block) []models.Block {
	if len(remoteBlocks) > len(localBlock) {
		return remoteBlocks
	}
	return localBlock
}

func (s *BlockService) BlockChain() []models.BlockHeader {
	return blockHeaders
}
