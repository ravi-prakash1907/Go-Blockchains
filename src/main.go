// creats a dummy blockchains
package main

import (
	"fmt"
	"strconv"

	"github.com/ravi-prakash1907/golang-walkthrough/Golang-Programs/src/blockchains/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// print our blockchain
	for _, block := range chain.Blocks {
		fmt.Printf("\n\nPrevious Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("This Hash: %x\n\n", block.Hash)

		// for pow
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	fmt.Println("Run the code as many time as you want, the hashes will remain same!!")
	fmt.Println("Also remember, changing data from even from one block will ultimatly lead the all following hashes to change!!")
}
