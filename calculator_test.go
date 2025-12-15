package main

import (
	"testing"
)

// Test NewCalculator for check bundlediscount and memberdiscount
func TestNewCalculator(t *testing.T) {
	calc := NewCalculator()

	if calc.BundleDiscount != 0.05 {
		t.Errorf("Expected BundleDiscount to be 0.05, got %f", calc.BundleDiscount)
	}

	if calc.MemberDiscount != 0.10 {
		t.Errorf("Expected MemberDiscount to be 0.10, got %f", calc.MemberDiscount)
	}
}

// Test calculateItemPrice for regular items
func TestCalculateItemPrice_RegularItem(t *testing.T) {
	calc := NewCalculator()

	price := calc.calculateItemPrice("Red", 3)
	expected := 150.0
	//t.Logf("Expected: %.2f, Got: %.2f", expected, price)

	if price != expected {
		t.Errorf("Expected %f, got %f", expected, price)
	}
}

// Test calculateItemPrice for bundle items
func TestCalculateItemPrice_BundleItem(t *testing.T) {
	calc := NewCalculator()

	price := calc.calculateItemPrice("Orange", 3)
	expected := 348.0
	//t.Logf("Expected: %.2f, Got: %.2f", expected, price)

	if price != expected {
		t.Errorf("Expected %f, got %f", expected, price)
	}
}

// Test calculateItemPrice for non-existent item
func TestCalculateItemPrice_NonExistentItem(t *testing.T) {
	calc := NewCalculator()

	price := calc.calculateItemPrice("White", 2)
	expected := 0.0
	//t.Logf("Expected: %.2f, Got: %.2f", expected, price)

	if price != expected {
		t.Errorf("Expected %f, got %f", expected, price)
	}
}

// Test bundlePrice
func TestBundlePrice(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name     string
		price    float64
		qty      int
		expected float64
	}{
		{"2 items", 100.0, 2, 190.0},
		{"3 items", 100.0, 3, 290.0},
		{"4 items", 100.0, 4, 380.0},
		{"1 item", 100.0, 1, 100.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.bundlePrice(tt.price, tt.qty)
			//t.Logf("result: %.2f", result)
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

// Test memberCardDiscount
func TestMemberCardDiscount(t *testing.T) {
	calc := NewCalculator()

	amount := 1000.0
	expected := 900.0

	result := calc.memberCardDiscount(amount)
	//t.Logf("result: %.2f", result)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

// Test CalculatePrice without member card
func TestCalculatePrice_NoMemberCard(t *testing.T) {
	calc := NewCalculator()

	items := map[string]int{
		"Red":   2,
		"Green": 2,
	}

	total := calc.CalculatePrice(items, false)
	expected := 176.0
	//t.Logf("total: %.2f, expected: %.2f", total, expected)

	if total != expected {
		t.Errorf("Expected %f, got %f", expected, total)
	}
}

// Test CalculatePrice with member card
func TestCalculatePrice_WithMemberCard(t *testing.T) {
	calc := NewCalculator()

	items := map[string]int{
		"Red":   2,
		"Green": 2,
	}
	// Total: 176
	// With member discount: 176 * 0.90 = 158.4

	total := calc.CalculatePrice(items, true)
	expected := 158.4
	//t.Logf("total: %.2f, expected: %.2f", total, expected)

	if total != expected {
		t.Errorf("Expected %f, got %f", expected, total)
	}
}

// Test CalculatePrice with mixed items
func TestCalculatePrice_ComplexScenario(t *testing.T) {
	calc := NewCalculator()

	items := map[string]int{
		"Orange": 3, // (120 * 2 * 0.95) + 120 = 348
		"Pink":   2, // 80 * 2 * 0.95 = 152
		"Red":    1, // 50 * 1 = 50
	}
	// Total: 348 + 152 + 50 = 550
	// With member card: 550 * 0.90 = 495

	total := calc.CalculatePrice(items, true)
	expected := 495.0

	if total != expected {
		t.Errorf("Expected %f, got %f", expected, total)
	}
}
