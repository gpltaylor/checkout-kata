package main

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct {
	totalPrice int
	basket     map[string]int
	ItemPrices map[string]ItemPrice
}

type ItemPrice struct {
	sku           string
	price         int
	multiPrice    int
	multiPriceQty int
}

func (c *Checkout) Scan(item string) {
	// Add scanned item to a Key Value store (Map)
	if c.basket == nil {
		c.basket = make(map[string]int)
	}
	c.basket[item]++
}

func (c *Checkout) GetTotalPrice() int {
	// Return the total price including the discount
	return c.totalPrice
}

func NewCheckout() ICheckout {
	return &Checkout{
		// TODO: Load from Redis or something (Dapr state store?)
		ItemPrices: map[string]ItemPrice{
			"A": {"A", 50, 130, 3},
			"B": {"B", 30, 45, 2},
			"C": {"C", 20, 0, 0},
			"D": {"D", 15, 0, 0},
		},
	}
}
