package main

import "block-chain/cmd"

func main() {
	go cmd.RunMiner()
	cmd.RunWeb()
}
