package main

import "testing"

func TestCalculateBasket(t *testing.T) {
	tests := []struct {
		inputSKUs     string
		expectedPrice int
	}{
		{"AA", 100},
		{"AAA", 130},
		{"AAAA", 180},
		{"B", 30},
		{"BB", 45},
		{"BBB", 75},
		{"BBBB", 90},
	}

	for _, tt := range tests {
		checkout := NewCheckout()
		for _, sku := range tt.inputSKUs {
			checkout.Scan(string(sku))
		}
		totalPrice := checkout.GetTotalPrice()
		if totalPrice != tt.expectedPrice {
			t.Errorf("Expected %d, got %d", tt.expectedPrice, totalPrice)
		}
	}
}
