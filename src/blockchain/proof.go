// here we implement code for proof of work algo.

package blockchain

import (
	"math"
	"fmt"
	"log"
	"bytes"
	"encoding/binary"
	"math/big"
	"crypto/sha256"
)
// forcing network to work by adding a block to the chain 
// it secure too


// take the data from block

// create counter (0-infinity)

// new hash = data + counter

// check hash to meet a set of requirements (in not, create new hash and keep on repeating)

// Requirements:
// 1. initial few bytes must be 0s (aks in case of bitcoin's hash-cash, it required initial 20-bits to be 0s)

const Difficulty = 18 //12 // in real BC, algos. gradually keep on increasing the Values as more blocks added

type ProofOfWork struct {
	Block *Block
	Target *big.Int // know more aboyt bytes and big-int in go
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// use target to shift left (Lsh) the num. of bytes by clculated places
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	// following is an ~  infinite loop (MaxInt64)
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash := sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			fmt.Println()
			return nonce, hash[:]
		} else {
			nonce = nonce+1;
		}
	}
	
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	// little-endian: LSB stored before MSB
	// big-endian: MSB stored before LSB
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic()
	}

	return buff.Bytes()
}

