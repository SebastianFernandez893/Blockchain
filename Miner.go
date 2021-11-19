package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

//type Miner struct {
//	id       int
//	pullChan chan bool
//	pushChan chan Block
//}
//
//func createMiner(id int) Miner {
//	pull := make(chan bool, 10)
//	push := make(chan Block, 10)
//	miner := Miner{id, pull, push}
//	return miner
//}

//func run(currMiner *Miner, wg *sync.WaitGroup) {
//	//Miner will need to know the difficulty, and the hash of the previousblock
//	//Structure will sort of be
//	//While there is no block coming in from the logger, findNonce, once nonce found, create and then send block
//	//The determinant for if the block is coming will need to be a channel pulling from the logger, for now I will just make it a boolean
//	noNewBlock := true
//	//Difficulty and previous hash will need to be pulled from the channel as well. Or its accessed from the blockchain?
//	for noNewBlock {
//		nonce, hash := findNonce("abc", 1)
//		//create Block here
//		//Send block here
//	}
//
//}

/*
	@output nonce 	integer used to solve puzzle
			newHash string of
*/
func findNonce(hash [32]byte, diff int) (int, [32]byte) {
	diffSlice := make([]byte, diff) // Slice used to compare the found hash
	nonce := -1                     // starting number for nonce guess
	hashString := string(hash[:32])

	//creates new hashes by concatenating a new nonce and the previous blocks hash each iteration.
	//then compares it to the difficulty
	var newHash [32]byte
	for true {
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash := sha256.Sum256([]byte(strNonce + hashString))
		compareSlice := newHash[:diff]
		if bytes.Equal(compareSlice, diffSlice) {
			return nonce, newHash
		}
	}
	return nonce, newHash
}
