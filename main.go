package main

import (
	"bytes"
	"fmt"
	//	"sync"
	//	"time"
	"crypto/sha256"
	"strconv"
	//"strings"
)

type Block struct {
	previousBlock *Block
	nonce         int
	tx            string //Do we need a tx at all? I guess it would then need to be its own struct, so maybe unessecary

}
type chain struct{
	head Block

}

func findNonce(hash string) int {
	//Will need a way to change the difficulty of diffSlice
	diffSlice := []byte{0}  //Slice which is used to compare the found hash
	fmt.Printf("The difficulty is: %x \n", diffSlice)
	nonceFound := false
	nonce := -1
	var newHash []byte

	//This for loop continuely creates new hashes by concatenating a new nonce and the previous blocks hash each iteration.
	//It then compares it to the difficulty
	for !nonceFound {
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash := sha256.Sum256([]byte(strNonce + hash))
		if nonce%10000 == 0 {
			fmt.Printf("Here %x \n", newHash)
		}
		x := newHash[:1]
		if bytes.Equal(x, diffSlice) {
			fmt.Printf("In if statement %x \n", newHash)
			nonceFound = true
		}
	}
	fmt.Printf("%x \n", newHash) //So there is a problem with how outside of the for loop, newHash won't print out.
	return nonce

}

func main() {
	//var miners, diff,rounds int
//	diff, miners, rounds := askInput()
	nonce, hash := findNonce("hteh",1)
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
