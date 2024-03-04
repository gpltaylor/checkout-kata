# Checkout Kata

Implement the code for a checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."

Implement the code for a supermarket checkout that calculates the total price of a number of items. In a normal supermarket, things are identified using Stock Keeping Units, or SKUs. In our store, we’ll use individual letters of the alphabet (A, B, C, and so on). Our goods are priced individually. In addition, some items are multi-priced: buy n of them, and they’ll cost you y pence. For example, item A might cost 50 individually, but this week we have a special offer: buy three As and they’ll cost you 130. In fact the prices are:

| SKU  | Unit Price | Special Price |
| ---- | ---------- | ------------- |
| A    | 50         | 3 for 130     |
| B    | 30         | 2 for 45      |
| C    | 20         |               |
| D    | 15         |               |

The checkout accepts items in any order, so that if we scan a B, an A, and another B, we’ll recognize the two Bs and price them at 45 (for a total price so far of 95). **The pricing changes frequently, so pricing should be independent of the checkout.**

The interface to the checkout could look like:

```cs
interface ICheckout
{
    void Scan(string item);
    int GetTotalPrice();
}
```

## Implementation
A Checkout implementing the above Interface has been created. This will allow DI to load into any application and use the Checkout.

## Arch
This is a simple online test so it's fine to place all items in root.

Based on the Product Owner requirements, business requirements and solutions architecture, the product structure would change.

- pkg folder
- cmd folder
- internal folder
- api folder
- web folder

### Data
The data would not be stored within a File :) 
This would be load using Redis Cache with a set TTL. Using something like Dapr to load this into the microservice arch.

I hope this is enough to highlist the basics... there is a LOT mroe I would do here but I was told to only put a few hours into this.

