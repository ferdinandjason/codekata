package domain

import (
	"github.com/martinlindhe/unit"
)

type Item struct {
	Name  string
	Price float64
	Unit  unit.Mass
	Offer Offer
}
