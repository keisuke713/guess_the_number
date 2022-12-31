package rand

import (
	"math/rand"
	"time"
)

func NewInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}