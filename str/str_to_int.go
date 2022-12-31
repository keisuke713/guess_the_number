package str

import (
	"strconv"
)

func ToIntWithFL(a string) (int, error) {
	len := len(a)
	aWithoutLF := a[:len-1]
	return strconv.Atoi(aWithoutLF)	
}