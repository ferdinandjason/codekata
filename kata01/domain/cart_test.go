package domain

import (
	"testing"

	"github.com/ferdinandjason/codekata/kata01/offer"
	"github.com/martinlindhe/unit"
)

func TestCart_GrandTotal(t *testing.T) {
	var tcs = []struct {
		cart     Cart
		expected float64
	}{
		{
			cart: Cart{
				Items: []ItemGroup{
					{
						Item: Item{
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
						Item: Item{
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
						Item: Item{
							Name:  "Vegetables",
							Price: 20000.0,
							Unit:  unit.AvoirdupoisPound,
							Offer: offer.NoOffer{},
						},
						Quantity: 4,
						Unit:     unit.AvoirdupoisOunce,
					},
				},
			},
			expected: 120000.0,
		},
	}

	for _, tc := range tcs {
		actual := tc.cart.GrandTotal()
		if actual != tc.expected {
			t.Errorf("ItemGroup.CalculatePrice(): expected %f, got %f", tc.expected, actual)
		}
	}
}
