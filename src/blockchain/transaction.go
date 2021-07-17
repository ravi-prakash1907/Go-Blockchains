// coin based transection
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value int
	// in bitcoins, this key is derived
	// from a coplicated language - 'script'
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	// sig is a script provides data i.e. used in output
	Sig string // currently it'll be common thing
}


func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func CoinBaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf(("Coins to %s", to))
	}

	txin := TxInput{[]byte{}, -1, data}
	// if jack mines this block, hell get 100 coins
	txout := TxOutput{100,to} 

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

// unlock data in i/p & o/p
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

