# Blockchain
This project simulates a blockchain based on Bitcoin's cryptographic puzzle using Golang's native channels and routines in a synchronous netowrk.

##Input and Output
The user will input the number of miners, blocks, and processes being ran. As well as the difficulty of the cryptographic puzzle. The difficulty is defined as how many leading zeroes need to be present in the proposed hash.

How to Run:
----

### 1: Clone Git Repository

Clone the repository with `git clone https://github.com/SebastianFernandez1999/Blockchain`

#### 2: Run using 
`go run *.go`

#### 3: Input number of miners, blocks, processes, and difficulty.

Discussion of Technologies Used
---
Hash Puzzle: Each miner is attempting to solve a cryptographic puzzle.
When presented with the hash (h) of the previous block, the miner must find a nonce(some number) that when hashed, using the SHA-256 protocol, with h, will produce a new hash with a certain amount of leading zeroes.

Blockchain: The blockchain in this program is not held by any data structure. Instead to access it, the current block must be used, and can only be searched through backwards to access the block you want.  

