package offer

import (
	"math"

	"github.com/martinlindhe/unit"
)

type BuyXGetYOffer struct {
	BuyUnit  int
	FreeUnit int
}

func (bxgy BuyXGetYOffer) CalculateDiscount(itemPrice float64, sourceUnit unit.Mass, itemQuantity float64, targetUnit unit.Mass) float64 {
	multiplierUnit := float64(targetUnit / sourceUnit)
	groupedUnit := float64(bxgy.BuyUnit + bxgy.FreeUnit)

	itemQuantityDiscounted := math.Floor((itemQuantity * multiplierUnit * float64(bxgy.FreeUnit)) / groupedUnit)
	totalDiscount := itemQuantityDiscounted * itemPrice

	return totalDiscount
}
