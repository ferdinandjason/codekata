package offer

import (
	"testing"

	"github.com/ferdinandjason/codekata/kata01/domain"
	"github.com/martinlindhe/unit"
)

func TestNoOffer_CalculateDiscount(t *testing.T) {
	var tcs = []struct {
		item         domain.Item
		itemQuantity float64
		targetUnit   unit.Mass
		expected     float64
	}{
		{
			item: domain.Item{
				Name:  "Aqua",
				Price: 5000.0,
				Unit:  unit.Kilogram,
				Offer: NoOffer{},
			},
			itemQuantity: 3.0,
			targetUnit:   unit.Kilogram,
			expected:     0.0,
		},
	}

	for _, tc := range tcs {
		actual := tc.item.Offer.CalculateDiscount(tc.item.Price, tc.item.Unit, tc.itemQuantity, tc.targetUnit)
		if actual != tc.expected {
			t.Errorf("NoOffer.CalculateDiscount(): expected %f, got %f", tc.expected, actual)
		}
	}
}
