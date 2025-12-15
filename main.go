package main

import "fmt"

func main() {
	cal := NewCalculator()

	// Example 1: Red + Green (No member card)
	order1 := map[string]int{"Red": 1, "Green": 1}
	price1 := cal.CalculatePrice(order1, false)
	fmt.Printf("Example 1 Price: %.2f บาท\n\n", price1)

	// Example 2: Red + Green with Member Card
	order2 := map[string]int{"Red": 1, "Green": 1}
	price2 := cal.CalculatePrice(order2, true)
	fmt.Printf("Example 2 Price: %.2f บาท\n\n", price2)

	// Example 3: Orange 5 (No member card)
	order3 := map[string]int{"Orange": 5}
	price3 := cal.CalculatePrice(order3, false)
	fmt.Printf("Example 3 Price: %.2f บาท\n\n", price3)

	// Example 4: Orange 5 sets Member Card
	order4 := map[string]int{"Orange": 5}
	price4 := cal.CalculatePrice(order4, true)
	fmt.Printf("Example 4 Price: %.2f บาท\n\n", price4)

}
