package libs

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// 需要发送的交易信息
type TX struct {
	AddFrom     string
	Transaction []byte
}

// 需要发送的区块版本信息
type BlockVersion struct {
	Version    int
	BestHeight int
	AddrFrom   string
}
//
type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

func sendAddr(address string) {
	nodes := addr{KnownNodes}
	nodes.AddrList = append(nodes.AddrList, NodeAddress)
	payload := gobEncode(nodes)
	request := append(commandToBytes("addr"), payload...)

	sendData(address, request)
}

func sendBlock(addr string, b *Block) {
	data := block{NodeAddress, b.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("block"), payload...)

	sendData(addr, request)
}

func sendInv(address, kind string, items [][]byte) {
	inventory := inv{NodeAddress, kind, items}
	payload := gobEncode(inventory)
	request := append(commandToBytes("inv"), payload...)

	sendData(address, request)
}

func SendGetBlocks(address string) {
	payload := gobEncode(getblocks{NodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(address, request)
}

func sendGetData(address, kind string, id []byte) {
	payload := gobEncode(getdata{NodeAddress, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	sendData(address, request)
}

func SendTx(addr string, tnx *Transaction) {
	data := TX{NodeAddress, tnx.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("tx"), payload...)

	sendData(addr, request)
}

func SendVersion(addr string, bc *Blockchain) {
	bestHeight := bc.GetBestHeight()
	payload := gobEncode(BlockVersion{nodeVersion, bestHeight, NodeAddress})

	request := append(commandToBytes("version"), payload...)

	sendData(addr, request)
}

func sendData(addr string, data []byte) {
	conn, err := net.Dial(os.Getenv("MINE_PROTO"), addr)
	if err != nil {
		fmt.Printf("%s is not available\n", addr)
		var updatedNodes []string

		for _, node := range KnownNodes {
			if node != addr {
				updatedNodes = append(updatedNodes, node)
			}
		}

		KnownNodes = updatedNodes

		return
	}
	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}
