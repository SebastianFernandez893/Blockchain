// logger funtions and struct

package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

func loggerVerify(b *Block) {

	verify := false

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
		verify = true
	}

	_ = verify
	// append block to list of blocks
	//end routine and start logger notify
}

func loggerNotify(notify) {
	// close current channels with miner
	close(pull)
	//initialize channels to send the block to each miner

}