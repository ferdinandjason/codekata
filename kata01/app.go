package main

import (
	"fmt"

	"github.com/ferdinandjason/codekata/kata01/domain"
	"github.com/ferdinandjason/codekata/kata01/offer"
	"github.com/martinlindhe/unit"
)

func main() {
	SupermarketCart := domain.Cart{
		Items: []domain.ItemGroup{
			{
				Item: domain.Item{
					Name:  "Mouse Pad 1",
					Price: 30000.0,
					Unit:  unit.Kilogram,
					Offer: offer.BuyXGetYOffer{
						BuyUnit:  3,
						FreeUnit: 1,
					},
				},
				Quantity: 4.0,
				Unit:     unit.Kilogram,
			}, {
				Item: domain.Item{
					Name:  "Metal Straw 1",
					Price: 10000.0,
					Unit:  unit.Kilogram,
					Offer: offer.BuyXForYOffer{
						BuyUnit:      3,
						SpecialPrice: 25000.0,
					},
				},
				Quantity: 3.0,
				Unit:     unit.Kilogram,
			}, {
				Item: domain.Item{
					Name:  "Vegetables",
					Price: 20000.0,
					Unit:  unit.AvoirdupoisPound,
					Offer: offer.NoOffer{},
				},
				Quantity: 4,
				Unit:     unit.AvoirdupoisOunce,
			},
		},
	}

	fmt.Println("Grand Total: Rp.", SupermarketCart.GrandTotal())
}
