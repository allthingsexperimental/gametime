package character

import "github.com/allthingsexperimental/gametime/ability"

type Character struct {
	Name      string
	Abilities []*ability.Ability
}
