package GoUtils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateId(amount int) string {
	returned := ""
	for amount >= 0 {
		rand.Seed(time.Now().UnixNano())
		returned += strconv.Itoa(rand.Intn(9))
		amount--
	}
	return returned
}
