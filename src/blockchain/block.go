package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// each block gonna have a data along with a hash
// this block will attatch to last (previous) block
type Block struct {
	Hash         []byte         // hash of this block
	Transactions []*Transaction // for coin based txn (ledger, doc ets) in this block
	PrevHash     []byte         // last block's hash
	Nonce        int            // to store nonce for validation implemtation
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(txs), prevHash, 0}
	// running pow algo
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// handle the errors
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
