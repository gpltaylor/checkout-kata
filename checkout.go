package main

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct {
	totalPrice int
}

func (c *Checkout) Scan(item string) {
	// Add scanned item to a Key Value store (Map)
}

func (c *Checkout) GetTotalPrice() int {
	// Return the total price including the discount
	return c.totalPrice
}

func NewCheckout() ICheckout {
	return &Checkout{}
}
