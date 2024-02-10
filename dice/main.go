package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	dice := newDice("dice.txt")
	dice = shuffle(dice)
	print(dice)
}

func newDice(filename string) int {

	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	no, _ := strconv.Atoi(string(bs))

	return no
}

func shuffle(loop int) int {

	var nums []int

	for i := 1; i <= loop; i++ {
		nums = append(nums, i)
	}

	rand.Seed(time.Now().Unix())
	n := (rand.Int()%len(nums) + 1)

	return n
}

func print(d int) {
	fmt.Println(d)
}
