package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	diff, minerCount, blockCount, threadCount := loopInput()

	runtime.GOMAXPROCS(threadCount)
	toLoggerChan := make(chan toLoggerData, minerCount)
	minerArray := initMiners(minerCount, toLoggerChan, blockCount)

	//Send it through channel
	time1 := time.Now()
	for i := 0; i < minerCount; i++ {

		go run(&minerArray[i], diff) //Does this need a waitgroup? Probably not, why would the miners need to wait for other miners to finsih?
	}
	go runLogger(&wg, minerArray, toLoggerChan, blockCount, diff)
	wg.Wait()

	time2 := time.Now()
	timeDiff := time2.Sub(time1)
	fmt.Println("the total puzzle time is : ", &timeDiff)

}

/*
	function loops askInput() until correct input is submitted
	@output 4 integers corresponding to user input for difficulty level, miner count, block count, and thread count
*/
func initMiners(minerCount int, toLoggerChan chan toLoggerData, blockCount int) []Miner {
	minerArray := make([]Miner, minerCount)
	for i := 0; i < minerCount; i++ {
		minerArray[i] = createMiner(i, toLoggerChan, blockCount)
	}
	return minerArray
}

func loopInput() (int, int, int, int) {
	needInput := true
	input := []int{0, 0, 0, 0}
	for needInput {
		diff, miners, rounds, procsNum, updateNeedInput := askInput()
		input[0] = diff
		input[1] = miners
		input[2] = rounds
		input[3] = procsNum
		needInput = updateNeedInput
	}
	return input[0], input[1], input[2], input[3]
}

/*
	@output 4 integers corresponding to input accepted from user
			1 boolean representing whether the askInput() needs to be called again
*/
func askInput() (int, int, int, int, bool) {
	fmt.Println("This program will simulate a blockchain by initializing several miners and a single logger.")
	fmt.Println("The miners will attempt to solve cryptographic puzzles according to a difficulty you set.")
	fmt.Println("The difficulty is determined by comparing the most significant bits of the two hashes.")
	fmt.Println("-------------------------------------------------------------------------------")

	fmt.Println("Please choose how many leading bits you would like to be compared.")
	var difficulty int
	_, errD := fmt.Scanln(&difficulty)
	if errD != nil {
		fmt.Println("Invalid difficulty level! Try again using an integer.")
		return 0, 0, 0, 0, true
	}

	fmt.Println("Please input the number of miners you would like to simulate.")
	var numOfMiners int
	_, errM := fmt.Scanln(&numOfMiners)
	if errM != nil {
		fmt.Println("Invalid number of miners! Try again using an integer.")
		return 0, 0, 0, 0, true
	}

	fmt.Println("Please input the number of blocks you would like to add to the blockchain.")
	var numOfRounds int
	_, errR := fmt.Scanln(&numOfRounds)
	if errR != nil {
		fmt.Println("Invalid number of miners! Try again using an integer.")
		return 0, 0, 0, 0, true
	}
	fmt.Println("Please input the number of concurrent threads you would like to use.")
	var numOfProcs int
	_, errP := fmt.Scanln(&numOfProcs)
	if errP != nil {
		fmt.Println("Invalid number of threads! Try again using an integer.")
		return 0, 0, 0, 0, true
	}

	fmt.Println("Thanks! We will start the simulation with", numOfMiners, "miners on difficulty level", difficulty,
		"for", numOfRounds, "rounds using a GOMAXPROCS number of", numOfProcs, ".")
	fmt.Println("-------------------------------------------------------------------------------")
	return difficulty, numOfMiners, numOfRounds, numOfProcs, false
}
