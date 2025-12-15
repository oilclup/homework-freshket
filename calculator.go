package main

// menu lists
var prices = map[string]float64{
	"Red":    50,
	"Green":  40,
	"Blue":   30,
	"Yellow": 50,
	"Pink":   80,
	"Purple": 90,
	"Orange": 120,
}

// bundle lists
var bundleItems = map[string]bool{
	"Orange": true,
	"Pink":   true,
	"Green":  true,
}

type Calculator struct {
	BundleDiscount float64 // Bundle Discount (0.05 = 5%)
	MemberDiscount float64 // Member card discount (0.10 = 10%)
}

func NewCalculator() *Calculator {
	return &Calculator{
		BundleDiscount: 0.05,
		MemberDiscount: 0.10,
	}
}

// CalculatePrice
func (c *Calculator) CalculatePrice(items map[string]int, hasMemberCard bool) float64 {
	total := 0.0

	for name, qty := range items {
		if qty <= 0 {
			continue
		}

		itemPrice := c.calculateItemPrice(name, qty)
		total += itemPrice
	}
	//case hasMemberCard
	if hasMemberCard {
		total = c.memberCardDiscount(total)

	}

	return total
}

// calculateItemPrice calculates price for a single item type
func (c *Calculator) calculateItemPrice(name string, qty int) float64 {
	price, exists := prices[name]
	if !exists {
		return 0
	}

	if bundleItems[name] {
		return c.bundlePrice(price, qty)
	}

	return price * float64(qty)
}

// bundlePrice applies 5% discount for every 2 items
func (c *Calculator) bundlePrice(price float64, qty int) float64 {
	pairs := qty / 2
	single := qty % 2

	pairQty := pairs * 2
	pairPrice := price * float64(pairQty) * (1 - c.BundleDiscount)

	singlePrice := price * float64(single)

	return pairPrice + singlePrice
}

// memberCardDiscount applies 10% discount on total amount
func (c *Calculator) memberCardDiscount(amount float64) float64 {
	return amount * (1 - c.MemberDiscount)
}
