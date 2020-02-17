package domain

import (
	"github.com/ferdinandjason/codekata/kata01/offer"
	"github.com/martinlindhe/unit"
)

var FourMousePadWithBuy3Get1Offer = Item{
	Name:  "Mouse Pad 1",
	Price: 30000.0,
	Unit:  unit.Kilogram,
	Offer: offer.BuyXGetYOffer{
		BuyUnit:  3,
		FreeUnit: 1,
	},
	Quantity: 4.0,
	BuyUnit:  unit.Kilogram,
}

var ThreeMetalStrawWithBuy3For25000Offer = Item{
	Name:  "Metal Straw 1",
	Price: 10000.0,
	Unit:  unit.Kilogram,
	Offer: offer.BuyXForYOffer{
		BuyUnit:      3,
		SpecialPrice: 25000.0,
	},
	Quantity: 3.0,
	BuyUnit:  unit.Kilogram,
}

var FourOuncesVegetablesWith20000PerPound = Item{
	Name:     "Vegetables",
	Price:    20000.0,
	Unit:     unit.AvoirdupoisPound,
	Offer:    offer.NoOffer{},
	Quantity: 4.0,
	BuyUnit:  unit.AvoirdupoisOunce,
}
