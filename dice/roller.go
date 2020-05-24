package dice

import (
	"math/rand"
	"time"
)

func Roller(die uint8) uint8 {
	t := time.Now().Nanosecond()
	rand.Seed(int64(t))
	i := uint8(rand.Intn(int(die)))
	return i
}
