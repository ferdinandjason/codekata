package domain

type Cart struct {
	Items []Item
}

func (c Cart) GrandTotal() float64 {
	var totalPrice = 0.0

	for _, item := range c.Items {
		totalPrice += item.CalculatePrice()
	}

	return totalPrice
}
