
package main

import (
	"fmt"
)

type Block struct {
	previousBlock *Block
	nonce         int
	//tx            string //Do we need a tx at all? I guess it would then need to be its own struct, so maybe unessecary
	hash [32]byte
}
func createBlock(previousBlock *Block, nonce int,hash [32]byte) Block{
	block := Block{previousBlock, nonce, hash}
	return block
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




