<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
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
<<<<<<< Updated upstream

}

func main() {
	nonce := findNonce("abcd")
	fmt.Printf("%d \n", nonce)
=======
}

type chain struct{
	head Block
}

func main() {
	//var miners, diff,rounds int
	diff, miners, rounds := askInput()
	nonce := findNonce("abcd",diff)
	fmt.Printf("%x \n", nonce)
>>>>>>> Stashed changes
}

func askInput() (int, int,int){
	fmt.Println("This program will simulate a blockchain by initializing several miners and a single logger. \n")
	fmt.Println("The miners will attempt to solve cryptographic puzzles according to a difficulty you set. \n")
	fmt.Println("The difficulty is determined by comparing the most significant bits of the two hashes. \n")

<<<<<<< Updated upstream
=======
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
>>>>>>> Stashed changes
}


func findNonce(hash string, diff int) int {

	diffSlice := make([]byte,diff)//Slice which is used to compare the found hash
	nonceFound := false
	nonce := -1
<<<<<<< Updated upstream
	var newHash []byte
=======

>>>>>>> Stashed changes
	//This for loop continuely creates new hashes by concatenating a new nonce and the previous blocks hash each iteration.
	//It then compares it to the difficulty
	for !nonceFound {
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash := sha256.Sum256([]byte(strNonce + hash))
<<<<<<< Updated upstream
		if nonce%10000 == 0 {
			fmt.Printf("Here %x \n", newHash)
		}

		x := newHash[:1]
=======
		x := newHash[:diff]
>>>>>>> Stashed changes
		if bytes.Equal(x, diffSlice) {
			nonceFound = true
		}
	}
	return nonce
}
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
