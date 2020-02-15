package offer

import (
	"math"

	"github.com/martinlindhe/unit"
)

type BuyXForYOffer struct {
	BuyUnit      int
	SpecialPrice float64
}

func (bxfy BuyXForYOffer) CalculateDiscount(itemPrice float64, sourceUnit unit.Mass, itemQuantity float64, targetUnit unit.Mass) float64 {
	multiplierUnit := float64(targetUnit / sourceUnit)
	discountedGroup := math.Floor((itemQuantity * multiplierUnit) / float64(bxfy.BuyUnit))

	discountPerGroupPrice := float64(bxfy.BuyUnit)*itemPrice - bxfy.SpecialPrice
	totalDiscount := discountedGroup * discountPerGroupPrice

	return totalDiscount
}
