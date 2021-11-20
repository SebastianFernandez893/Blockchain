package main

import (
"crypto/sha256"
"fmt"
"strconv"
"time"
)

type Block struct {
	hash            [32]byte
	confirmations   int      // not utilized
	size            int      // not utilized
	strippedSize    int      // excluding witness data; not utilized
	weight          int      // as defined in BIP 141; not utilized
	height          int      // not utilized
	version         int      // not utilized
	versionHex      string   // not utilized
	merkleRoot      string   // not utilized
	tx              []string // only use first index
	timestamp       int64    // not utilized
	medianTimestamp int64    // not utilized
	nonce           int
	bits            string // not utilized
	difficulty      int
	chainWork       string // not utilized
	prevBlockHash   [32]byte
	previousBlock   *Block // might not end up keeping
}

func blockToString(b *Block) {
	fmt.Println("Printing Block information:")
	fmt.Println("\thash:", b.hash)
	fmt.Println("\tconfirmations:", b.confirmations)
	fmt.Println("\tsize:", b.size)
	fmt.Println("\tstripped size:", b.strippedSize)
	fmt.Println("\tweight:", b.weight)
	fmt.Println("\tversion:", b.version)
	fmt.Println("\tversion hex:", b.versionHex)
	fmt.Printf("\tmerkle root: %x\n", b.merkleRoot)
	fmt.Printf("\ttransaction: %x\n", b.tx[0])
	fmt.Println("\ttime:", b.timestamp)
	fmt.Println("\tmedian time:", b.medianTimestamp)
	fmt.Println("\tnonce:", b.nonce)
	fmt.Println("\tbits:", b.bits)
	fmt.Println("\tdifficulty:", b.difficulty)
	fmt.Println("\tchain work:", b.chainWork)
	fmt.Println("\tprevBlockHash:", b.prevBlockHash)
}

/*
	@param 	Block b
	@output hash of all block data
*/
func setBlockHash(b *Block) {
	// concatenate string of all data fields except for hash
	hashIn := strconv.Itoa(b.confirmations) + strconv.Itoa(b.size) + strconv.Itoa(b.strippedSize) +
		strconv.Itoa(b.weight) + strconv.Itoa(b.height) + strconv.Itoa(b.version) + b.versionHex + b.merkleRoot +
		b.tx[0] + strconv.FormatInt(b.timestamp, 10) + strconv.FormatInt(b.medianTimestamp, 10) + strconv.Itoa(b.nonce) +
		b.bits + strconv.Itoa(b.difficulty) + b.chainWork + string(b.prevBlockHash[:32])
	hashOut := sha256.Sum256([]byte(hashIn))
	//output := string(hashOut[:32]) // there might be an issue with the length of the array
	b.hash = hashOut
}

/*
	creates a new block given the previous block, new nonce, and difficulty level
	generates fake data for non-utilized data fields
	@output new block
*/
func createBlock(nonce int, hash [32]byte, difficulty int, prevBlock *Block) Block {
	merkleRootByteArray := sha256.Sum256([]byte("transaction"))
	merkleRoot := string(merkleRootByteArray[:32])
	tx0ByteArray := sha256.Sum256([]byte("transaction1"))
	tx0String := string(tx0ByteArray[:32])
	tx := []string{tx0String}
	timestamp := time.Now().Unix()
	medianTimestamp := time.Now().Unix()
	chainWork := "000000000000000000000000000000000000000000000000000001f501f501f5"
	prevBlockHash := prevBlock.hash
	block := Block{hash, 1, 1, 1, 1,
		1, 1, "00000001", merkleRoot, tx,
		timestamp, medianTimestamp, nonce,
		"1d00ffff", difficulty, chainWork, prevBlockHash, prevBlock}
	return block
}

/*
	creates the first block given difficulty level
	generates fake data for non-utilized data fields
	@output new block
*/
func createFirstBlock(difficulty int) Block {
	hash := sha256.Sum256([]byte("hash"))
	merkleRootByteArray := sha256.Sum256([]byte("transaction"))
	merkleRoot := string(merkleRootByteArray[:32])
	tx0ByteArray := sha256.Sum256([]byte("transaction1"))
	tx0String := string(tx0ByteArray[:32])
	tx := []string{tx0String}
	timestamp := time.Now().Unix()
	medianTimestamp := time.Now().Unix()
	chainWork := "000000000000000000000000000000000000000000000000000001f501f501f5"
	prevBlockHash := sha256.Sum256([]byte(""))
	nonce := 0
	block := Block{hash, 1, 1, 1, 1,
		1, 1, "00000001", merkleRoot, tx,
		timestamp, medianTimestamp, nonce,
		"1d00ffff", difficulty, chainWork, prevBlockHash, nil}
	return block
}
