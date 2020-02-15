package offer

import (
	"github.com/martinlindhe/unit"
)

type NoOffer struct{}

func (no NoOffer) CalculateDiscount(itemPrice float64, sourceUnit unit.Mass, itemQuantity float64, targetUnit unit.Mass) float64 {
	return 0.0
}
