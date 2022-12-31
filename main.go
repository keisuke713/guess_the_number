package main

import (
	"fmt"
	"os"
	"log"
	"errors"

	read "guess_the_number/read"
	rand "guess_the_number/rand"
	guess "guess_the_number/guess"
)

const (
	instruction = `This is the Guess the Number game.
You type min number and max number.
Then, we generate a number between two number you just typed
And you guess what it is!	
`
)

var (
	lessThanOrEqualZeroErr = errors.New("You have to choose a number which is more than zero")
	secondNumberSmallerThanFirstOneErr = errors.New("Second number have to be bigger than first one")
)

var (
	reader = read.New(os.Stdin)
)

func main() {
	fmt.Println(instruction)

	min := reader.ReadFromRd("Type min number: ")
	if min <= 0 {
		log.Fatalf(lessThanOrEqualZeroErr.Error())
	}

	max := reader.ReadFromRd("Type max number: ")
	if max <= 0 {
		log.Fatalf(lessThanOrEqualZeroErr.Error())
	}

	if min >= max {
		log.Fatalf(secondNumberSmallerThanFirstOneErr.Error())
	}
	
	// ランダムな値生成
	res := rand.NewInt(min, max)
	// 数字を推測させる
	guess.Guess(res)
}