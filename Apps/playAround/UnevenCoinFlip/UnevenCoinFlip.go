package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//IterationVar is a constant for number of flips. Can be user defined
var IterationVar int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	setIterVar()

	sum := 0
	for i := 0; i < IterationVar; i++ {
		sum += simulateFlip()
	}

	fmt.Println("***********************************")
	fmt.Printf("TOTAL SUM: %d", sum)
	fmt.Printf("\nNUMBER OF FLIPS: %d", IterationVar)
	avg := float64(sum) / float64(IterationVar)
	fmt.Printf("\nCALCULATED COIN PROBABILITY: %f", avg)
	fmt.Println("\n***********************************")

}

func setIterVar() {
	if len(os.Args) > 1 {
		IterationVar, _ = strconv.Atoi(os.Args[1])
	} else {
		//Default flip count
		IterationVar = 10000
	}
}

func simulateFlip() int {
	coinFlips := 0
	for 1 != 0 {
		flip1 := flipCoin()
		flip2 := flipCoin()
		coinFlips += 2
		if flip1 != flip2 {
			return flip1
		}
	}
	return -1
}

func flipCoin() int {
	if rand.Float64() < .5 {
		return 0
	}
	return 1
}
