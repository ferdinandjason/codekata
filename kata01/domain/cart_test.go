package domain

import (
	"testing"
)

func TestCart_GrandTotal(t *testing.T) {
	var tcs = []struct {
		name     string
		input    Cart
		expected float64
	}{
		{
			name: "calculate total price in the cart",
			input: Cart{
				Items: []Item{
					FourMousePadWithBuy3Get1Offer,
					ThreeMetalStrawWithBuy3For25000Offer,
					FourOuncesVegetablesWith20000PerPound,
				},
			},
			expected: 120000.0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.GrandTotal()
			if actual != tc.expected {
				t.Errorf("ItemGroup.CalculatePrice(): expected %f, got %f", tc.expected, actual)
			}
		})
	}
}
