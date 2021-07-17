package blockchain

// a chain of blocks
type BlockChain struct {
	Blocks []*Block
}

// each block gonna have a data along with a hash
// this block will attatch to last (previous) block
type Block struct {
	Hash []byte    // hash of this block
	Data []byte    // data (ledger, doc ets) in this block 
	PrevHash []byte // last block's hash 
	Nonce int	// to store nonce for validation implemtation
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// running pow algo
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
