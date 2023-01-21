package goutil

import (
	"math/rand"
	"strconv"
	"time"
)

// GenerateId generates an string of numbers
func GenerateId(amount int) string {
	returned := ""
	for amount >= 0 {
		rand.Seed(time.Now().UnixNano())
		returned += strconv.Itoa(rand.Intn(9))
		amount--
	}
	return returned
}
