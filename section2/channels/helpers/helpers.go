package helpers

import (
	"math/rand"
	"time"
)

func GenerateRamdomNumber(n int) int {
	r := rand.New(rand.NewSource((time.Now().UnixNano())))
	value := r.Intn(n)
	return value

}
