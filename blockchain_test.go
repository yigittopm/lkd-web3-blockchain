package main

import (
	"strings"
	"testing"
	"time"
)

const errorMessage = "Expected %v but got %v"

func TestCreateTransaction(t *testing.T) {
	tx := NewTransaction("Alice", "Bob", 100)
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
	tx1 := NewTransaction("Alice", "Bob", 100)
	tx2 := NewTransaction("Mert", "Ali", 3)
	tx3 := NewTransaction("Anıl", "Cenk", 10)
	tx4 := NewTransaction("Mert", "Mert", 203)

	block := NewBlock()
	block.AddTransaction(tx1)
	block.AddTransaction(tx2)
	block.AddTransaction(tx3)
	block.AddTransaction(tx4)
	minedBlock := block.Mine(blockchain)
	blockchain.Blockchain = append(blockchain.Blockchain, minedBlock)

	if len(blockchain.Blockchain) != 2 {
		t.Errorf(errorMessage, 2, len(blockchain.Blockchain))
	}

	if blockchain.Blockchain[0].Hash != blockchain.Blockchain[1].PrevHash {
		t.Errorf(errorMessage, blockchain.Blockchain[0].Hash, blockchain.Blockchain[1].PrevHash)
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
