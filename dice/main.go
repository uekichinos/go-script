package main

func main() {
	dice := newDice("dice.txt")
	dice = shuffle(dice)
	print(dice)
}
