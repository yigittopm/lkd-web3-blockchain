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

	// New Mempool instance
	mempool := NewMempool()

	// Create a new transaction
	NewTransaction(mempool, "Alice", "Bob", 100)
	NewTransaction(mempool, "Mert", "Ali", 3)
	NewTransaction(mempool, "Anıl", "Cenk", 10)
	NewTransaction(mempool, "Mert", "Mert", 2013)
	NewTransaction(mempool, "Cenk", "Alp", 2014)

	// Mine the block
	for len(mempool.Transactions) > 0 {
		Mine(blockchain, mempool)
	}

	// Print the blockchain
	Log(blockchain)
}
