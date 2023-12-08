package constant

import (
	"math/rand"
	"time"
)

func RandomSixDigitNumber() int {
	min := 100000
	max := 999999
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
