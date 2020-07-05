package td

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
func GenerateRangeNum(min, max int) int {
	randNum := r.Intn(max - min) + min
	return randNum
}