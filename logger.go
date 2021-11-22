// logger funtions and struct

package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

func loggerVerify(b *Block) {

	prevBlockHash := b.hash
	nonce := b.nonce
	diff := b.difficulty

	hashSeed := bytes.NewBuffer(prevBlockHash[:]).String()
	diffSlice := make([]byte, diff)
	strNonce := strconv.Itoa(nonce)
	verifyHash := sha256.Sum256([]byte(strNonce + hashSeed))
	x := verifyHash[:diff]
	// verify nonce
	if bytes.Equal(x, diffSlice) {
		return true
	}

	return false
	// append block to list of blocks
	//end routine and start logger notify
}

func runLogger(minerArray, pushChan) {
// close current channels with miner
// select case

loggerVerify()

//initialize channels to send the block to each miner
Notify chan <- true 
 }
