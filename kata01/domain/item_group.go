package domain

import (
	"github.com/martinlindhe/unit"
)

type ItemGroup struct {
	Item     Item
	Quantity float64
	Unit     unit.Mass
}

func (ig ItemGroup) PriceWithoutOffer() float64 {
	multiplierUnit := float64(ig.Unit / ig.Item.Unit)
	price := ig.Item.Price * multiplierUnit * ig.Quantity

	return price
}

func (ig ItemGroup) CalculatePrice() float64 {
	totalPrice := ig.PriceWithoutOffer()
	discountPrice := ig.Item.Offer.CalculateDiscount(ig.Item.Price, ig.Item.Unit, ig.Quantity, ig.Unit)
	grandTotal := totalPrice - discountPrice

	return grandTotal
}
