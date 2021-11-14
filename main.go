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

type Block struct{
	previousBlock *Block
	nonce int
	tx string //Do we need a tx at all? I guess it would then need to be its own struct, so maybe unessecary

}
func main() {
	 str :=findNonce("abcd")
	 fmt.Printf("%d \n", str)
}

func askInput(){

}

func findNonce(hash string) []byte{

	diffSlice := []byte{0}
	fmt.Printf("The difficulty is: %x \n", diffSlice)
	nonceFound := false
	nonce:= -1
	var newHash []byte
	for !nonceFound{
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash := sha256.Sum256([]byte(strNonce+hash))
		if nonce%1000==0 {
			fmt.Printf("%x \n", newHash)
		}
		x :=  newHash[:1]
		if bytes.Equal(x,diffSlice) {
			nonceFound = true
		}
	}
	return newHash

}
