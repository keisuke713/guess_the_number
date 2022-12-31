package guess

import (
	"fmt"
	"os"

	read "guess_the_number/read"
)

type GuessImpl struct {
	res int
	*read.Reader
}

const (
	correctMsg = `You did it!!!
%d is the one we selected`

	incorrectMsg = `Oops!!!
Not the one!!
You can try as much as you want`
)

func Guess(n int) {
	New(n).Guess()
}

func New(n int) *GuessImpl {
	return &GuessImpl{
		n,
		read.New(os.Stdin),
	}
}

func (g *GuessImpl) Guess() {
	for {
		n := g.ReadFromRd("Guess the number: ")

		if g.res == n {
			msg := fmt.Sprintf(correctMsg, g.res)
			fmt.Println(msg)
			break
		}
		fmt.Println(incorrectMsg)
	}
}