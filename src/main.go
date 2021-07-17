// creats a dummy blockchains
package main

import (
	"fmt"
	"bytes"
	"crypto/sha256"
)

// a chain of blocks
type BlockChain struct {
	blocks []*Block
}

// each block gonna have a data along with a hash
// this block will attatch to last (previous) block
type Block struct {
	Hash []byte    // hash of this block
	Data []byte    // data (ledger, doc ets) in this block 
	PrevHash []byte // last block's hash 
}

// to get hash value
func (b *Block) DeriveHash() {
	// join bytes' slices
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // actual blockchain hash is far complex
	b.Hash = hash[:]

	// here hashes are derived using this data + "previous hash"
	// any change in previous data will change its hash and 
	// it'll modify all following blocks' hash too
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// print our blockchain
	for _, block := range chain.blocks {
		// fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("This Hash: %x\n\n", block.Hash)
	}

	fmt.Println("Run the code as many time as you want, the hashes will remain same!!")
	fmt.Println("Also remember, changing data from even from one block will ultimatly lead the all following hashes to change!!")
}