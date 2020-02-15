package offer

import (
	"testing"

	"github.com/ferdinandjason/codekata/kata01/domain"
	"github.com/martinlindhe/unit"
)

func TestBuyXForYOffer_CalculateDiscount(t *testing.T) {
	var tcs = []struct {
		item         domain.Item
		itemQuantity float64
		targetUnit   unit.Mass
		expected     float64
	}{
		{
			item: domain.Item{
				Name:  "Metal Straw 1",
				Price: 10000.0,
				Unit:  unit.Kilogram,
				Offer: BuyXForYOffer{
					BuyUnit:      3,
					SpecialPrice: 25000.0,
				},
			},
			itemQuantity: 3.0,
			targetUnit:   unit.Kilogram,
			expected:     5000.0,
		}, {
			item: domain.Item{
				Name:  "Metal Straw 2",
				Price: 10000.0,
				Unit:  unit.Kilogram,
				Offer: BuyXForYOffer{
					BuyUnit:      3,
					SpecialPrice: 25000.0,
				},
			},
			itemQuantity: 4.0,
			targetUnit:   unit.Kilogram,
			expected:     5000.0,
		}, {
			item: domain.Item{
				Name:  "Metal Straw 3",
				Price: 10000.0,
				Unit:  unit.Kilogram,
				Offer: BuyXForYOffer{
					BuyUnit:      3,
					SpecialPrice: 25000.0,
				},
			},
			itemQuantity: 5.0,
			targetUnit:   unit.Kilogram,
			expected:     5000.0,
		}, {
			item: domain.Item{
				Name:  "Metal Straw 4",
				Price: 10000.0,
				Unit:  unit.Kilogram,
				Offer: BuyXForYOffer{
					BuyUnit:      3,
					SpecialPrice: 25000.0,
				},
			},
			itemQuantity: 6.0,
			targetUnit:   unit.Kilogram,
			expected:     10000.0,
		},
	}

	for _, tc := range tcs {
		actual := tc.item.Offer.CalculateDiscount(tc.item.Price, tc.item.Unit, tc.itemQuantity, tc.targetUnit)
		if actual != tc.expected {
			t.Errorf("BuyXForYOffer.CalculateDiscount(): expected %f, got %f", tc.expected, actual)
		}
	}
}
