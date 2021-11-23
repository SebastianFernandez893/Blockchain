// logger funtions and struct

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"sync"
)

func runLogger(wg *sync.WaitGroup, minerArray []Miner, toLoggerChan chan toLoggerData, chainLength int, diff int) {
	defer wg.Done()
	fmt.Println("Started Logger")
	i := 1
	oldBlock := createFirstBlock(diff)
	blockToString(&oldBlock)
	// send first block to all miners
	for i := 0; i < len(minerArray); i++ {
		minerArray[i].toMinerChan <- oldBlock
	}
	fmt.Println("sent initial block to miners")
	for i <= chainLength {
		newmsg := <-toLoggerChan
		fmt.Println("received block proposal from miner", newmsg.miner.id)
		isValid, currBlock := notifyMiner(minerArray, newmsg, oldBlock)
		if isValid {
			i++
			fmt.Println("incremented chain height")
		}
		fmt.Println("i value is", i)
		oldBlock = currBlock
	}
	fmt.Println("Ending Logger")
}

func loggerVerify(b *Block) bool {
	prevBlockHash := b.prevBlockHash
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

func notifyMiner(minerArray []Miner, newmsg toLoggerData, oldBlock Block) (bool, Block) {
	// close current channels with miner
	// select case
	block := newmsg.block
	if loggerVerify(&block) == true {
		fmt.Println("Block was verified!! from", newmsg.miner.id)
		for i := 0; i < len(minerArray); i++ {
			//minerArray[i].notifyChan <- true
			minerArray[i].toMinerChan <- block
			fmt.Println("Block was sent to miner", minerArray[i].id)
		}
		return true, block // if block was valid and added to chain
	} else {
		fmt.Println("Block was invalid, notifying miner", newmsg.miner.id)
		newmsg.miner.toMinerChan <- oldBlock
		return false, oldBlock // block was invalid and not added to chain
	}
}
