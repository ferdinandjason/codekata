package domain

import (
	"testing"

	"github.com/ferdinandjason/codekata/kata01/offer"
	"github.com/martinlindhe/unit"
)

func TestItemGroup_PriceWithoutOffer(t *testing.T) {
	var tcs = []struct {
		itemGroup ItemGroup
		expected  float64
	}{
		{
			itemGroup: ItemGroup{
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
			},
			expected: 120000.0,
		},
	}

	for _, tc := range tcs {
		actual := tc.itemGroup.PriceWithoutOffer()
		if actual != tc.expected {
			t.Errorf("ItemGroup.PriceWithoutOffer(): expected %f, got %f", tc.expected, actual)
		}
	}
}

func TestItemGroup_CalculatePrice(t *testing.T) {
	var tcs = []struct {
		itemGroup ItemGroup
		expected  float64
	}{
		{
			itemGroup: ItemGroup{
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
			},
			expected: 90000.0,
		}, {
			itemGroup: ItemGroup{
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
			},
			expected: 25000.0,
		}, {
			itemGroup: ItemGroup{
				Item: Item{
					Name:  "Vegetables",
					Price: 20000.0,
					Unit:  unit.AvoirdupoisPound,
					Offer: offer.NoOffer{},
				},
				Quantity: 4,
				Unit:     unit.AvoirdupoisOunce,
			},
			expected: 5000.0,
		},
	}

	for _, tc := range tcs {
		actual := tc.itemGroup.CalculatePrice()
		if actual != tc.expected {
			t.Errorf("ItemGroup.CalculatePrice(): expected %f, got %f", tc.expected, actual)
		}
	}
}
