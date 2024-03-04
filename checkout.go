package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

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
	Sku           string `json:"sku"`
	Price         int    `json:"price"`
	MultiPrice    int    `json:"multiPrice"`
	MultiPriceQty int    `json:"multiPriceQty"`
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
	for sku, qty := range c.basket {
		itemPrice := c.ItemPrices[sku]
		if itemPrice.MultiPrice > 0 && qty >= itemPrice.MultiPriceQty {
			c.totalPrice += (qty/itemPrice.MultiPriceQty)*itemPrice.MultiPrice + (qty%itemPrice.MultiPriceQty)*itemPrice.Price
		} else {
			c.totalPrice += qty * itemPrice.Price
		}
	}
	return c.totalPrice
}

func NewCheckout() ICheckout {
	ItemPrices := make(map[string]ItemPrice)

	// Read JSON file and populate the ItemPrices map
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	file, err := os.Open(filepath.Join(cwd, "item-prices.json"))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ItemPrices); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return &Checkout{
		ItemPrices: ItemPrices,
	}
}
