package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	Difficulty          = 3
	MaxTransactionCount = 3
)

type Mempool struct {
	mu           sync.Mutex
	Transactions []Transaction
}

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

func (m *Mempool) GetTransaction() *Transaction {
	m.mu.Lock()
	defer m.mu.Unlock()

	tx := m.Transactions[0]
	m.Transactions = m.Transactions[1:]

	return &tx
}

func Mine(blockchain *Blockchain, mem *Mempool) {
	var wg sync.WaitGroup
	txList := []Transaction{}

	for i := range len(mem.Transactions) {
		if i > MaxTransactionCount-1 {
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			txList = append(txList, *mem.GetTransaction())
		}()
	}
	wg.Wait()

	prevBlock := blockchain.Blockchain[len(blockchain.Blockchain)-1]
	hash, nonce := GenerateHash(txList, prevBlock.Hash, time.Now())

	minedBlock := &Block{
		Id:           prevBlock.Id + 1,
		Hash:         hash,
		PrevHash:     prevBlock.Hash,
		Nonce:        nonce,
		Timestamp:    time.Now(),
		Transactions: txList,
	}

	blockchain.Blockchain = append(blockchain.Blockchain, minedBlock)
}

func MineGenesisBlock() *Block {
	initHash := "0000000000000000000000000000000000000000000000000000000000000000"
	hash, nonce := GenerateHash([]Transaction{}, initHash, time.Now())

	return &Block{
		Id:           0,
		Hash:         hash,
		PrevHash:     initHash,
		Nonce:        nonce,
		Timestamp:    time.Now(),
		Transactions: []Transaction{},
	}
}

func NewMempool() *Mempool {
	return &Mempool{}
}

func NewTransaction(mem *Mempool, from, to Address, amount int) {
	tx := Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}

	mem.Transactions = append(mem.Transactions, tx)
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	genesis := MineGenesisBlock()

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
