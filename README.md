# Blockchain
Gabriella Munger, Justin Gomez, Sebastian Fernandez

Description
----
This project simulates a blockchain based on Bitcoin's cryptographic puzzle using Golang's native channels and routines in a synchronous network. There is one Logger keeping track of a tamper-resistant log and several Miners (number is specified by user input) that work to create the next block. When one of the Miners solves the hash puzzle (of a difficulty set by user input), the Miner sends their nonce (solution) and the new block to the Logger for verification. When the Logger has verified the validity of the block, it logs the block and sends the new hash of the new block to each of the Miners to serve as the next puzzle. This continues until the chain reaches the number of blocks desired (as specificed by user input). The user can also specify how many concurrent threads are used to run the program, which is then controlled using GOMAXPROCS().

To make this as similar to the implementation of bitcoin while supporting our needs, we kept the same data fields for the block as are used for a block in bitcoin. We looked at the source code for bitcoin (https://github.com/bitcoin/bitcoin/tree/master/src) and also other resources, including http://www.herongyang.com/Bitcoin/Bitcoin-Data-Block-Data-Field.html. For fields that are meaningless in this project, we used "dummy" data.

Input and Output
----
The user will input the number of miners, blocks, and processors being ran, as well as the difficulty of the cryptographic puzzle. The difficulty is defined by how many leading zeros need to be present in the proposed hash. A higher difficulty level means that more leading zeros are required.

How to Run:
----

### 1: Clone Git Repository

Clone the repository using `git clone https://github.com/SebastianFernandez1999/Blockchain`

#### 2: Run using 
`go run *.go`

### 3.  Input parameters as requested
The program will request integer inputs for the difficulty level, number of miners, number of blocks in the chain, and number of concurrent threads. If invalid input is received, the program will ask the user to try again.

-- will add a screenshot of input screen--
### 4. Example of expected output

-- will add screenshots and descriptions of sample outputs--

Workflow Diagram
----
![MP2 - Algorithm flowchart example](https://user-images.githubusercontent.com/90423480/142967631-9070163d-e2b7-4b43-929b-eb4be783675b.jpeg)


Discussion of Technologies Used
---
Hash Puzzle: Each miner is attempting to solve a cryptographic puzzle.
When presented with the hash (h) of the previous block, the miner must find a nonce(some number) that when hashed, using the SHA-256 protocol, with h, will produce a new hash with a certain amount of leading zeroes.

Blockchain: The blockchain in this program is not held by any data structure. Instead to access it, the current block must be used, and can only be searched through backwards to access the block you want.  

