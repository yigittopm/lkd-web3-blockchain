package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

var (
	Difficulty = 5
)

type Blockchain struct {
	Blockchain []*Block
}

type Block struct {
	Id           int
	Hash         string
	PrevHash     string
	Timestamp    time.Time
	Transactions []Transaction
	Nonce        int
}

type Address string

type Transaction struct {
	From   Address
	To     Address
	Amount int
}

func (b *Block) AddTransaction(tx Transaction) {
	b.Transactions = append(b.Transactions, tx)
}

func NewTransaction(from, to Address, amount int) Transaction {
	return Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
}

func NewGenesisBlock() *Block {
	block := &Block{}
	hash, nonce := GenerateHash(block)

	return &Block{
		Id:           0,
		Hash:         hash,
		PrevHash:     "0000000000000000000000000000000000000000000000000000000000000000",
		Timestamp:    time.Now(),
		Transactions: []Transaction{},
		Nonce:        nonce,
	}
}

func NewBlock(prevBlock *Block) *Block {
	hash, nonce := GenerateHash(prevBlock)

	return &Block{
		Id:           prevBlock.Id + 1,
		Hash:         hash,
		PrevHash:     prevBlock.Hash,
		Nonce:        nonce,
		Timestamp:    time.Now(),
		Transactions: []Transaction{},
	}
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.Blockchain = append(bc.Blockchain, NewGenesisBlock())
	return bc
}

func GenerateHash(block *Block) (string, int) {
	nonce := 0
	want := strings.Repeat("0", Difficulty)

	for {
		data := []byte(fmt.Sprintf("%v%v%v%v", block.Transactions, block.PrevHash, block.Timestamp.String(), nonce))

		hash := sha256.Sum256(data)
		encodedHash := hex.EncodeToString(hash[:])
		if encodedHash[:Difficulty] == want {
			return encodedHash, nonce
		}

		nonce++
	}
}

func main() {
	blockchain := NewBlockchain()

	tx1 := NewTransaction("Alice", "Bob", 100)
	tx2 := NewTransaction("Mert", "Ali", 3)
	tx3 := NewTransaction("Anıl", "Cenk", 10)
	tx4 := NewTransaction("Anıl", "Anıl", 203)

	block := NewBlock(blockchain.Blockchain[len(blockchain.Blockchain)-1])
	block.AddTransaction(tx1)
	block.AddTransaction(tx2)
	block.AddTransaction(tx3)
	block.AddTransaction(tx4)
	blockchain.Blockchain = append(blockchain.Blockchain, block)

	tx222 := NewTransaction("Mert", "Onur", 3)
	block2 := NewBlock(blockchain.Blockchain[len(blockchain.Blockchain)-1])
	block2.AddTransaction(tx222)
	blockchain.Blockchain = append(blockchain.Blockchain, block2)

	for _, block := range blockchain.Blockchain {
		fmt.Printf("{\n ID: %v,\n PrevHash: %v,\n Hash: %v,\n Nonce: %v,\n Transactions: %v \n}, \n", block.Id, block.PrevHash, block.Hash, block.Nonce, block.Transactions)
	}
}
