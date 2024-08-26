package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

var (
	Difficulty = 4
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

type GenerateHashPayload struct {
	PrevHash     string
	Timestamp    time.Time
	Transactions []Transaction
}

func (b *Block) AddTransaction(tx Transaction) {
	b.Transactions = append(b.Transactions, tx)
}

func (b *Block) MineBlock(blockchain *Blockchain) *Block {
	prevBlock := blockchain.Blockchain[len(blockchain.Blockchain)-1]
	hash, nonce := GenerateHash(b.Transactions, prevBlock.Hash, time.Now())

	return &Block{
		Id:           prevBlock.Id + 1,
		Hash:         hash,
		PrevHash:     prevBlock.Hash,
		Nonce:        nonce,
		Timestamp:    time.Now(),
		Transactions: b.Transactions,
	}
}

func (b *Block) MineGenesisBlock() *Block {
	hash, nonce := GenerateHash([]Transaction{}, "0000000000000000000000000000000000000000000000000000000000000000", time.Now())

	return &Block{
		Id:           0,
		Hash:         hash,
		PrevHash:     "0000000000000000000000000000000000000000000000000000000000000000",
		Nonce:        nonce,
		Timestamp:    time.Now(),
		Transactions: []Transaction{},
	}
}

func NewTransaction(from, to Address, amount int) Transaction {
	return Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
}

func NewBlock() *Block {
	return &Block{}
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	block := NewBlock()
	genesis := block.MineGenesisBlock()

	bc.Blockchain = append(bc.Blockchain, genesis)
	return bc
}

func GenerateHash(tx []Transaction, prevHash string, timestamp time.Time) (string, int) {
	nonce := 0
	want := strings.Repeat("0", Difficulty)

	for {
		data := []byte(fmt.Sprintf("%v%v%v%v", tx, prevHash, timestamp.String(), nonce))

		hash := sha256.Sum256(data)
		encodedHash := hex.EncodeToString(hash[:])
		if encodedHash[:Difficulty] == want {
			return encodedHash, nonce
		}

		nonce++
	}
}
