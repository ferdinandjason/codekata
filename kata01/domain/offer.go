package domain

import (
	"github.com/martinlindhe/unit"
)

type Offer interface {
	CalculateDiscount(itemPrice float64, sourceUnit unit.Mass, itemQuantity float64, targetUnit unit.Mass) float64
}
