package main

import "fmt"

func main() {
	fmt.Println("Yellow...")

	checkout := NewCheckout()
	checkout.Scan("A")
	checkout.Scan("B")
	checkout.Scan("A")
	totalPrice := checkout.GetTotalPrice()

	fmt.Println("Total price: ", totalPrice)
}
