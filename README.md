# Blockchain
Gabriella Munger, Justin Gomez, Sebastian Fernandez

Description
----
This project is meant to simulate a basic blockchain structured similarly to bitcoin. There is one Logger keeping track of a tamper-resistant log and several Miners (number is specified by user input) that work to create the next block. When one of the Miners solves the hash puzzle (of a difficulty set by user input), the Miner sends their nonce (solution) and the new block to the Logger for verification. When the Logger has verified the validity of the block, it logs the block and sends the new hash of the new block to each of the Miners to serve as the next puzzle. This continues until the chain reaches the number of blocks desired (as specificed by user input). The user can also specify how many concurrent threads are used to run the program, which is then controlled using GOMAXPROCS().

To make this as similar to the implementation of bitcoin while supporting our needs, we kept the same data fields for the block as are used for a block in bitcoin. We looked at the source code for bitcoin () and also other resources, including http://www.herongyang.com/Bitcoin/Bitcoin-Data-Block-Data-Field.html. For fields that are meaningless in this project, we used "dummy" data.

How To Run
----
### 1.  Clone Git Repository
### 2.  Run using:
`go run main.go block.go miner.go logger.go`

### 3.  Input parameters as requested
The program will request integer inputs for the difficulty level, number of miners, number of rounds (blocks in the chain), and number of concurrent threads. If invalid input is received, the program will ask the user to try again.

-- will add a screenshot of input screen--
### 4. Example of expected output

-- will add screenshots and descriptions of sample outputs--

Workflow Diagram
----
-- will add workflow diagram --
