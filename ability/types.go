package ability

import (
	"sort"

	"github.com/allthingsexperimental/gametime/dice"
)

type Ability struct {
	Name     string
	Base     uint8
	Modifier int8
}

const (
	DefaultDie       = 20
	DefaultRollCount = 5
)

var DefaultAbilities = []string{
	"Strength",
	"Dexterity",
	"Constitution",
	"Intelligence",
	"Wisdom",
	"Charisma",
}

func New(name string) *Ability {
	return &Ability{
		Name: name,
		Base: rollNTimes(DefaultRollCount),
	}
}

func rollNTimes(n uint8) uint8 {
	var rolls []int
	for len(rolls) < int(n) {
		rolls = append(rolls, int(dice.Roller(DefaultDie)))
	}
	sort.Ints(rolls)
	var s int
	for _, v := range rolls[1:] {
		s += v
	}
	return uint8(s / len(rolls[1:]))
}
