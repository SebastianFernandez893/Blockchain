package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"sync"
)

type Miner struct{
	id int
	pullChan chan bool
	pushChan chan Block
}



func createMiner(id int)Miner{
	pull:= make(chan bool, 10)
	push:= make(chan Block,10)
	miner:= Miner{id,pull,push}
	return miner
}

func run(currMiner *Miner, wg *sync.WaitGroup,){
	//Miner will need to know the difficulty, and the hash of the previousblock
	//Structure will sort of be
	//While there is no block coming in from the logger, findNonce, once nonce found, create and then send block
	
}

func findNonce(hash string, diff int) (int, [32]byte) {

	diffSlice := make([]byte,diff)//Slice which is used to compare the found hash
	nonceFound := false
	nonce := -1

	//This for loop continuely creates new hashes by concatenating a new nonce and the previous blocks hash each iteration.
	//It then compares it to the difficulty
	var newHash [32]byte
	for !nonceFound {
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash = sha256.Sum256([]byte(strNonce + hash))
		x := newHash[:diff]
		if bytes.Equal(x, diffSlice) {
			nonceFound = true
		}
	}
	return nonce, newHash
}