package domain

import (
	"github.com/martinlindhe/unit"
)

type Item struct {
	Name     string
	Price    float64
	Unit     unit.Mass
	Offer    Offer
	Quantity float64
	BuyUnit  unit.Mass
}

func (i Item) PriceWithoutOffer() float64 {
	multiplierUnit := float64(i.BuyUnit / i.Unit)
	price := i.Price * multiplierUnit * i.Quantity

	return price
}

func (i Item) CalculatePrice() float64 {
	totalPrice := i.PriceWithoutOffer()
	discountPrice := i.Offer.CalculateDiscount(i.Price, i.Unit, i.Quantity, i.Unit)
	grandTotal := totalPrice - discountPrice

	return grandTotal
}
