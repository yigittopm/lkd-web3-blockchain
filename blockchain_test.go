package main

import (
	"strings"
	"testing"
	"time"
)

const errorMessage = "Expected %v but got %v"

func TestCreateTransaction(t *testing.T) {
	mempool := NewMempool()
	NewTransaction(mempool, "Alice", "Bob", 100)

	tx := mempool.Transactions[0]
	if tx.From != "Alice" {
		t.Errorf(errorMessage, "Alice", tx.From)
	}
	if tx.To != "Bob" {
		t.Errorf(errorMessage, "Bob", tx.To)
	}
	if tx.Amount != 100 {
		t.Errorf(errorMessage, 100, tx.Amount)
	}
}

func TestMineBlock(t *testing.T) {
	blockchain := NewBlockchain()
	mempool := NewMempool()

	NewTransaction(mempool, "Alice", "Bob", 100)
	NewTransaction(mempool, "Mert", "Ali", 3)
	NewTransaction(mempool, "Anıl", "Cenk", 10)
	NewTransaction(mempool, "Mert", "Mert", 203)

	if len(mempool.Transactions) != 4 {
		t.Errorf(errorMessage, 4, len(mempool.Transactions))
	}

	if len(blockchain.Blockchain) != 1 {
		t.Errorf(errorMessage, 1, len(blockchain.Blockchain))
	}

	// Mine block
	Mine(blockchain, mempool)

	if len(mempool.Transactions) != 1 {
		t.Errorf(errorMessage, 1, len(mempool.Transactions))
	}

	if len(blockchain.Blockchain) != 2 {
		t.Errorf(errorMessage, 2, len(blockchain.Blockchain))
	}
}

func TestCreateBlockchain(t *testing.T) {
	blockchain := NewBlockchain()
	if len(blockchain.Blockchain) != 1 {
		t.Errorf(errorMessage, 1, len(blockchain.Blockchain))
	}

	if blockchain.Blockchain[0].PrevHash != "0000000000000000000000000000000000000000000000000000000000000000" {
		t.Errorf(errorMessage, "00..00", blockchain.Blockchain[0].PrevHash)
	}
}

func TestGenerateHash(t *testing.T) {
	tx := []Transaction{}

	hash, _ := GenerateHash(tx, "", time.Now())
	difficulty := strings.Repeat("0", Difficulty)

	if strings.HasPrefix(hash, difficulty) == false {
		t.Errorf("Expected hash to start with %v but got %v", difficulty, hash)
	}
}
