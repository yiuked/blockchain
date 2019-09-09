package libs

import (
	"block-chain/config"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

const nodeVersion = 1
const commandLength = 12

var NodeAddress string                      // 节点地址网络IP地址，如IP地址
var MiningAddress string                    // 矿工钱包地址
var KnownNodes = []string{"localhost:3000"} // 全节点列表
var BlocksInTransit = [][]byte{}            // 需要从其它节点抓取数据的区块地址
var MeePool = make(map[string]Transaction)  // 交易池

type addr struct {
	AddrList []string
}

type block struct {
	AddrFrom string
	Block    []byte
}

type getblocks struct {
	AddrFrom string
}

type getdata struct {
	AddrFrom string
	Type     string
	ID       []byte
}

func commandToBytes(command string) []byte {
	var bytes [commandLength]byte

	for i, c := range command {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func extractCommand(request []byte) []byte {
	return request[:commandLength]
}

func requestBlocks() {
	for _, node := range KnownNodes {
		SendGetBlocks(node)
	}
}

// StartServer starts a node(启动一个节点)
func StartServer(nodeID string, minerAddress string) {
	NodeAddress = fmt.Sprintf("localhost:%s", config.MinerPort)
	MiningAddress = minerAddress
	ln, err := net.Listen(config.MineProto, NodeAddress)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	log.Printf("Start listen mine node %s,net addr:%s\n", nodeID, NodeAddress)

	bc := NewBlockchain(nodeID)

	version := BlockVersion{nodeVersion, bc.GetBestHeight(), NodeAddress}
	log.Println(version)

	// 如果不是全节点，说明是新增节点，新增节点交换版本信息
	if NodeAddress != KnownNodes[0] {
		SendVersion(KnownNodes[0], bc)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn, bc)
	}
}

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func nodeIsKnown(addr string) bool {
	for _, node := range KnownNodes {
		if node == addr {
			return true
		}
	}

	return false
}
