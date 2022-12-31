package read

import (
	"fmt"
	"bufio"
	"io"
	"log"

	str "guess_the_number/str"
)

type Reader struct {
	*bufio.Reader
}

func New(rd io.Reader) *Reader {
	return &Reader{
		bufio.NewReader(rd),
	}
}

func (r *Reader) ReadFromRd(msg string) int {
	fmt.Print(msg)
	a, _ := r.ReadString('\n')
	n, err := str.ToIntWithFL(a)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return n
}