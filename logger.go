// logger funtions and struct

package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Logger struct {
	notify chan *bool
	verify bool
}

func loggerVerify(Block) {

	verify := false

	prevBlockHash := Block.hash
	nonce := Block.nonce
	diff := Block.difficulty
	strHash := strcon.Itoa(nonceFound)

	hashSeed := bytes.NewBuffer(prevBlockHash[:]).String()
	diffSlice := make([]byte, diff)
	strNonce := strconv.Itoa(nonce)
	verifyHash := sha256.Sum256([]byte(strNonce + hashSeed))
	x := verifyHash[:diff]
	// verify nonce
	if bytes.Equal(x, diffSlice) {
		verify := true
	}

	// append block to list of blocks
	//end routine and start logger notify
}

func loggerNotify(notify) {
	// close current channels with miner
	//initialize channels to send the block to each miner

}
