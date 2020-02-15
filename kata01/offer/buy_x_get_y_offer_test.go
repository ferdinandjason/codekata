package offer

import (
	"testing"

	"github.com/ferdinandjason/codekata/kata01/domain"
	"github.com/martinlindhe/unit"
)

func TestBuyXGetYOffer_CalculateDiscount(t *testing.T) {
	var tcs = []struct {
		item         domain.Item
		itemQuantity float64
		targetUnit   unit.Mass
		expected     float64
	}{
		{
			item: domain.Item{
				Name:  "Mouse Pad 1",
				Price: 30000.0,
				Unit:  unit.Kilogram,
				Offer: BuyXGetYOffer{
					BuyUnit:  3,
					FreeUnit: 1,
				},
			},
			itemQuantity: 3.0,
			targetUnit:   unit.Kilogram,
			expected:     0.0,
		}, {
			item: domain.Item{
				Name:  "Mouse Pad 2",
				Price: 20000.0,
				Unit:  unit.Kilogram,
				Offer: BuyXGetYOffer{
					BuyUnit:  1,
					FreeUnit: 1,
				},
			},
			itemQuantity: 3.0,
			targetUnit:   unit.Kilogram,
			expected:     20000.0,
		}, {
			item: domain.Item{
				Name:  "Mouse Pad 3",
				Price: 30000.0,
				Unit:  unit.Kilogram,
				Offer: BuyXGetYOffer{
					BuyUnit:  3,
					FreeUnit: 1,
				},
			},
			itemQuantity: 4.0,
			targetUnit:   unit.Kilogram,
			expected:     30000.0,
		},
	}

	for _, tc := range tcs {
		actual := tc.item.Offer.CalculateDiscount(tc.item.Price, tc.item.Unit, tc.itemQuantity, tc.targetUnit)
		if actual != tc.expected {
			t.Errorf("BuyXGetYOffer.CalculateDiscount(): expected %f, got %f", tc.expected, actual)
		}
	}
}
