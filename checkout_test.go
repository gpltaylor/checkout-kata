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
		{"C", 20},
		{"D", 15},
		{"ABCD", 115},
		{"ABCDABCD", 215},
		{"ABCDABCDABCD", 310},
		{"ABCDABCDABCDABCD", 410},
		{"ABCDABCDABCDABCDABCD", 525},
		{"ABCDABCDABCDABCDABCDABCD", 605},
		{"ABCDABCDABCDABCDABCDABCDABCD", 720},
		{"ABCDABCDABCDABCDABCDABCDABCDABCD", 820},
		{"ABCDABCDABCDABCDABCDABCDABCDABCDABCD", 915},
		{"ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD", 1015},
		{"ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD", 1130},
		{"ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD", 1210},
		{"ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD", 1325},
		{"ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD", 1425},
		{"ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD", 1520},
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

func BenchmarkCheckout(b *testing.B) {
	checkout := NewCheckout()
	for i := 0; i < b.N; i++ {
		checkout.Scan("ABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCDABCD")
		checkout.GetTotalPrice()
	}
}
