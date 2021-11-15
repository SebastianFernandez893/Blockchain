
package main

import (
	"bytes"
	"fmt"
	//	"sync"
	//	"time"
	"crypto/sha256"
	"strconv"

)

type Block struct {
	previousBlock *Block
	nonce         int
	tx            string //Do we need a tx at all? I guess it would then need to be its own struct, so maybe unessecary
}

func main() {
	//var miners, diff,rounds int
//	diff, miners, rounds := askInput()
	nonce, hash := findNonce("abcd",1)
	fmt.Printf("%x %x \n", nonce, hash)

}

func askInput() (int, int,int){
	fmt.Println("This program will simulate a blockchain by initializing several miners and a single logger. \n")
	fmt.Println("The miners will attempt to solve cryptographic puzzles according to a difficulty you set. \n")
	fmt.Println("The difficulty is determined by comparing the most significant bits of the two hashes. \n")


	fmt.Println("Please choose how many leading bits you would like to be compared. \n")
	var difficulty int
	fmt.Scanln(&difficulty)

	fmt.Println("Please input the number of miners you would like to simulate. \n")
	var numOfMiners int
	fmt.Scanln(&numOfMiners)

	fmt.Println("Please input the number of blocks you would like to add to the blockchain \n")
	var numOfRounds int
	fmt.Scanln(&numOfRounds)

	return difficulty, numOfMiners,numOfRounds

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

