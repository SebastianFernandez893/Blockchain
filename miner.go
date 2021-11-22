package main

import (
	"crypto/sha256"
	"strconv"

	"bytes"
)

type Miner struct {
	id       int
	pullChan *chan Block
	pushChan chan Block
}

func createMiner(id int) Miner {
	pull := make(chan Block, 10)
	push := make(chan Block, 10)
	miner := Miner{id, &pull, push}
	return miner
}

func run(currMiner *Miner, diff int) {
	//Miner will need to know the difficulty, and the hash of the previousblock
	//Structure will sort of be
	//While there is no block coming in from the logger, findNonce, once nonce found, create and then send block
	//The determinant for if the block is coming will need to be a channel pulling from the logger, for now I will just make it a boolean
	var prevBlock Block
	select {
	case newBlock, ok := <-*currMiner.pullChan:
		if ok {
			prevBlock = newBlock
		}
	default:
		prevBlockHash := prevBlock.hash
		nonce, hash := findNonce(prevBlockHash, 1)
		block := createBlock(nonce, hash, diff, &prevBlock)
		currMiner.pushChan <- block

	}
}

func findNonce(seed [32]byte, diff int) (int, [32]byte) {
	hashSeed := bytes.NewBuffer(seed[:]).String()
	diffSlice := make([]byte, diff) //Slice which is used to compare the found hash
	nonceFound := false
	nonce := -1

	//This for loop continuely creates new hashes by concatenating a new nonce and the previous blocks hash each iteration.
	//It then compares it to the difficulty
	var newHash [32]byte
	for !nonceFound {
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash = sha256.Sum256([]byte(strNonce + hashSeed))
		x := newHash[:diff]
		if bytes.Equal(x, diffSlice) {
			nonceFound = true
		}
	}
	return nonce, newHash
}
