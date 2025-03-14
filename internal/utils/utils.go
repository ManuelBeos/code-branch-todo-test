package utils

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomNumber(min, max int) int {
	return rng.Intn(max-min+1) + min
}
