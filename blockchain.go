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
