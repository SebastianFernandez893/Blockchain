// logger funtions and struct

package main

import (
	"bytes"
	"crypto/sha256"
)

type Logger struct {
	notify chan *bool
	verify bool
}

func loggerVerify(Block) {
	// receive nonce + block from miner

	// verify nonce value given by hashing nonce received
	prevHash = block.previousBlock
	nonceFound = block.nonce
	diff = block.difficulty
	diffSlice := []byte{diff}
	strHash := strcon.Itoa(nonceFound)
	verifyHash := sha256.Sum256([]byte(nonceFound + prevHash))
	//verify difficulty condition is met
	x := verifyHash[:diff+1]
	if bytes.Equal(x, diffSlice) {
		verify = True
	}
	// append block to list of blocks
	//end routine and start logger notify
}

func loggerNotify(notify) {
	// close current channels with miner
	//initialize channels to send the block to each miner

}
