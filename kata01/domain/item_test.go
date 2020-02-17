package domain

import (
	"testing"

	"github.com/ferdinandjason/codekata/kata01/offer"
	"github.com/martinlindhe/unit"
)

func TestItem_PriceWithoutOffer(t *testing.T) {
	var tcs = []struct {
		name     string
		input    Item
		expected float64
	}{
		{
			name: "calculate price without offer",
			input: Item{
				Name:  "Mouse Pad 1",
				Price: 30000.0,
				Unit:  unit.Kilogram,
				Offer: offer.BuyXGetYOffer{
					BuyUnit:  3,
					FreeUnit: 1,
				},
				Quantity: 4.0,
				BuyUnit:  unit.Kilogram,
			},
			expected: 120000.0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.PriceWithoutOffer()
			if actual != tc.expected {
				t.Errorf("Item.PriceWithoutOffer(): expected %f, got %f", tc.expected, actual)
			}
		})
	}
}

func TestItem_CalculatePrice(t *testing.T) {
	var tcs = []struct {
		name     string
		input    Item
		expected float64
	}{
		{
			name: "calculate price for buy x get y offer",
			input: Item{
				Name:  "Mouse Pad 1",
				Price: 30000.0,
				Unit:  unit.Kilogram,
				Offer: offer.BuyXGetYOffer{
					BuyUnit:  3,
					FreeUnit: 1,
				},
				Quantity: 4.0,
				BuyUnit:  unit.Kilogram,
			},
			expected: 90000.0,
		}, {
			name: "calculate price for buy x for y offer",
			input: Item{
				Name:  "Metal Straw 1",
				Price: 10000.0,
				Unit:  unit.Kilogram,
				Offer: offer.BuyXForYOffer{
					BuyUnit:      3,
					SpecialPrice: 25000.0,
				},
				Quantity: 3.0,
				BuyUnit:  unit.Kilogram,
			},
			expected: 25000.0,
		}, {
			name: "calculate price for item that need unit conversion",
			input: Item{
				Name:     "Vegetables",
				Price:    20000.0,
				Unit:     unit.AvoirdupoisPound,
				Offer:    offer.NoOffer{},
				Quantity: 4.0,
				BuyUnit:  unit.AvoirdupoisOunce,
			},
			expected: 5000.0,
		}, {
			name: "calculate price for buy x get y offer but the quantity is less than x+y",
			input: Item{
				Name:  "Mouse Pad 2",
				Price: 24000.0,
				Unit:  unit.Kilogram,
				Offer: offer.BuyXGetYOffer{
					BuyUnit:  3,
					FreeUnit: 1,
				},
				Quantity: 3.0,
				BuyUnit:  unit.Kilogram,
			},
			expected: 72000.0,
		}, {
			name: "calculate price for buy x for y offer but the quantity is less than x",
			input: Item{
				Name:  "Metal Straw 1",
				Price: 7000.0,
				Unit:  unit.Kilogram,
				Offer: offer.BuyXForYOffer{
					BuyUnit:      3,
					SpecialPrice: 15000.0,
				},
				Quantity: 2.0,
				BuyUnit:  unit.Kilogram,
			},
			expected: 14000.0,
		},
	}

	for _, tc := range tcs {
		actual := tc.input.CalculatePrice()
		if actual != tc.expected {
			t.Errorf("Item.CalculatePrice(): expected %f, got %f", tc.expected, actual)
		}
	}
}
