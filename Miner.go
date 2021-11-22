package main

import (
	"crypto/sha256"
	"strconv"
	"math/rand"
	"bytes"
)

type Miner struct{
	id int
	pullChan chan Block
	pushChan chan pushChanData
	notifyChan chan bool
}
type pushChanData struct{
	miner *Miner
	block Block
}


func createMiner(id int,push chan pushChanData)Miner{
	pull:= make(chan Block, 10)
	notifyChan := make(chan bool)
	miner:= Miner{id,pull,push,notifyChan}
	return miner
}

func run(currMiner *Miner,diff int, blockcount int){
	//Miner will need to know the difficulty, and the hash of the previousblock
	//Structure will sort of be
	//While there is no block coming in from the logger, findNonce, once nonce found, create and then send block
	//The determinant for if the block is coming will need to be a channel pulling from the logger, for now I will just make it a boolean
	var prevBlock Block
	for true{
		select{
		case newBlock := <- currMiner.pullChan:
			prevBlock = newBlock
			prevBlockHash := prevBlock.hash
			nonce,hash := findNonce(prevBlockHash,diff,currMiner)
			//Creating a bad nonce
			randNum := rand.Intn(100)
			if randNum<10{
				nonce,hash = findBadNonce(prevBlockHash,diff)
			}
			newBlockHeight := prevBlock.height+1
			block:=createBlock(nonce,hash,diff,&prevBlock,newBlockHeight)
			data := pushChanData{currMiner,block}
			currMiner.pushChan<-data
		default:
			//Do Nothing
		}
	}
}
func findBadNonce(seed [32]byte,diff int) (int, [32]byte){
	hashSeed:= bytes.NewBuffer(seed[:]).String()
	diffSlice := make([]byte,diff)//Slice which is used to compare the found hash
	nonceFound := false
	nonce := -1
	var newHash [32]byte
	for !nonceFound {
			nonce++
			strNonce := strconv.Itoa(nonce)
			newHash = sha256.Sum256([]byte(strNonce + hashSeed))
			x := newHash[:diff]
			if !bytes.Equal(x, diffSlice) {
				nonceFound = true
			}
		}
	return nonce, newHash
}

func findNonce(seed [32]byte, diff int, miner *Miner) (int, [32]byte) {

	hashSeed:= bytes.NewBuffer(seed[:]).String()
	diffSlice := make([]byte,diff)//Slice which is used to compare the found hash
	nonceFound := false
	nonce := -1

	//This for loop continuely creates new hashes by concatenating a new nonce and the previous blocks hash each iteration.
	//It then compares it to the difficulty
	var newHash [32]byte
	for !nonceFound {
		select{
			case here := <-miner.notifyChan:
				if here==true{
					break
				}
		default:
			nonce++
			strNonce := strconv.Itoa(nonce)
			newHash = sha256.Sum256([]byte(strNonce + hashSeed))
			x := newHash[:diff]
			if bytes.Equal(x, diffSlice) {
				nonceFound = true
			}
		}
	}
	return nonce, newHash
}