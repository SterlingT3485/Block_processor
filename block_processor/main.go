package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Block struct {
	ID   string `json:"id"`   // unique identifier
	View int    `json:"view"` // ordering
}

type Vote struct {
	BlockID string `json:"block_id"` // the block id to vote
}

var (
	blockLock     sync.Mutex            // lock for block
	voteLock      sync.Mutex            // lock for vote
	blocks        = make(map[string]Block) // storing all blocks
	votes         = make(map[string]bool)  // storing all vote
	acceptedViews = make(map[int]bool)     // the views done
	pendingBlocks = make(map[int]Block)    // the views to be processed
)

// start http server
func main() {
	http.HandleFunc("/block", handleBlock) // http for blocks
	http.HandleFunc("/vote", handleVote)   // http for votes

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // start server with potential error report
}

// process block request
func handleBlock(w http.ResponseWriter, r *http.Request) {
	var block Block
	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // if the request cannot be read properly
		return
	}

	log.Printf("Received block: ID=%s, View=%d\n", block.ID, block.View)

	blockLock.Lock()
	blocks[block.ID] = block
	blockLock.Unlock()

	checkAndAcceptBlocks() // call helper function to check and accept the received block
}

// process vote request
func handleVote(w http.ResponseWriter, r *http.Request) {
	var vote Vote
	if err := json.NewDecoder(r.Body).Decode(&vote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // if the request cannot be read properly
		return
	}

	log.Printf("Received vote for block ID=%s\n", vote.BlockID)

	voteLock.Lock()
	votes[vote.BlockID] = true
	voteLock.Unlock()

	checkAndAcceptBlocks() // call helper function to check and accept the received block
}

// helper function that checks and accepts the received block
func checkAndAcceptBlocks() {
	// protect the whole function with lock
	blockLock.Lock()
	defer blockLock.Unlock()

	// check for vote first
	for _, block := range blocks {
		if votes[block.ID] { // if the block is voted
			pendingBlocks[block.View] = block
		}
	}

	// check for smaller blocks next
	for {
		if block, ok := pendingBlocks[len(acceptedViews)]; ok { // checks if all blocks with views smaller than this block has been accepted
			fmt.Printf("Accepted Block: ID=%s, View=%d\n", block.ID, block.View)
			acceptedViews[block.View] = true
			delete(pendingBlocks, block.View)
		} else {
			break
		}
	}
}
