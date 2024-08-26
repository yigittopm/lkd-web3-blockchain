package main

import (
	"fmt"
)

func Log(b *Blockchain) {
	for _, block := range b.Blockchain {
		fmt.Printf("{\n ID: %v,\n PrevHash: %v,\n Hash: %v,\n Nonce: %v,\n Transactions: %v \n}, \n", block.Id, block.PrevHash, block.Hash, block.Nonce, block.Transactions)
	}
}

func main() {
	// New Blockchain instance
	blockchain := NewBlockchain()

	// Create a new transaction
	tx1 := NewTransaction("Alice", "Bob", 100)
	tx2 := NewTransaction("Mert", "Ali", 3)
	tx3 := NewTransaction("Anıl", "Cenk", 10)
	tx4 := NewTransaction("Mert", "Mert", 203)

	// Create a new block
	block := NewBlock()

	// Add transactions to the block
	block.AddTransaction(tx1)
	block.AddTransaction(tx2)
	block.AddTransaction(tx3)
	block.AddTransaction(tx4)

	// Mine the block
	minedBlock := block.MineBlock(blockchain)

	// Append the block to the blockchain
	blockchain.Blockchain = append(blockchain.Blockchain, minedBlock)

	// Print the blockchain
	Log(blockchain)
}
