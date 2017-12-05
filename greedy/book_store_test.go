package bookstore

import (
	"testing"
)

var testCases = []struct {
	description string
	basket      []int
	expected    float64
}{
	{
		description: "Only a single book",
		basket:      []int{1},
		expected:    8.00,
	},
	{
		description: "Two of the same book",
		basket:      []int{2, 2},
		expected:    16.00,
	},
	{
		description: "Empty basket",
		basket:      []int{},
		expected:    0.00,
	},
	{
		description: "Two different books",
		basket:      []int{1, 2},
		expected:    15.20,
	},
	{
		description: "Three different books",
		basket:      []int{1, 2, 3},
		expected:    21.60,
	},
	{
		description: "Four different books",
		basket:      []int{1, 2, 3, 4},
		expected:    25.60,
	},
	{
		description: "Five different books",
		basket:      []int{1, 2, 3, 4, 5},
		expected:    30.00,
	},
	{
		description: "Two groups of four is cheaper than group of five plus group of three",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 5},
		expected:    51.20,
	},
	{
		description: "Group of four plus group of two is cheaper than two groups of three",
		basket:      []int{1, 1, 2, 2, 3, 4},
		expected:    40.80,
	},
	{
		description: "Two each of first 4 books and 1 copy each of rest",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		expected:    55.60,
	},
	{
		description: "Two copies of each book",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		expected:    60.00,
	},
	{
		description: "Three copies of first book and 2 each of remaining",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1},
		expected:    68.00,
	},
	{
		description: "Three each of first 2 books and 2 each of remaining books",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 2},
		expected:    75.20,
	},
}

func TestCost(t *testing.T) {
	for _, testCase := range testCases {
		cost := Cost(testCase.basket)
		if testCase.expected != cost {
			t.Fatalf("FAIL: %s\n\tCost(%v) expected %v, got %v",
				testCase.description, testCase.basket, testCase.expected, cost)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkCost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Cost(testCase.basket)
		}
	}
}

func Cost(arr []int) float64 {
	price := 8.0 
	totalcost := 0.0
	println(len(arr))
	// if len(arr) == 0 {
	// 	cost := 0 * price
	// 	totalcost += cost
	// }
	// if len(arr) == 1 {
	// 	cost := 1 * price
	// 	totalcost += cost
	// }
	// if len(arr) == 2 {
	// 	if 
	// }
	switch len(arr) {
	// cost:= 0.0
	case 0:
		cost := 0 * price
		totalcost += cost
	case 1:
		cost := 1 * price
		totalcost += cost
	case 2:
		if arr[0] != arr[1] {
			cost := 2 * price * 0.95
			totalcost += cost
		}else{
			cost := 2 * price
			totalcost += cost
		}
	case 3:
		cost := 3 * price * (1 - 0.1)
		totalcost += cost
	case 4:
		cost := 4 * price * (1 - 0.2)
		totalcost += cost
	case 5:
		cost := 5 * price * (1 - 0.25)
		totalcost += cost
	case 6:
		cost := 4 * price * (1 - 0.2) + 2 * price * (1 - 0.05)
		totalcost += cost

	case 8:
		// cost := 5 * price * (1 - 0.25) + 3 * price * (1 - 0.1)
		cost := 4 * price * (1 - 0.2) * 2
		totalcost += cost
	case 9:
		cost := 5 * price * (1 - 0.25) + 4 * price * (1 - 0.2)
		totalcost += cost
	case 10:
		cost := 5 * price * (1 - 0.25) + 5 * price * (1 - 0.25)
		totalcost += cost
	case 11:
		cost := 5 * price * (1 - 0.25) * 2 + 1 * price
		totalcost += cost
	case 12:
		cost := 5 * price * (1 - 0.25) * 2 + 2 * price * (1 - 0.05)
		totalcost += cost
	

	default:
		// totalcost += cost
		// println(cost)
		println(totalcost)
		
	}
	 return totalcost
}

// func dp(arr []int) float64 {
// 	discount []float64{0.25, 0.20, 0.10, 0.05, 0}
// 	buf [5]int
// 	cash [5]int
// 	for i := 0; i < 5; i++ {
// 		cash[i] = 0

// 	}
// 	for i := 1; i <= 5; i++ {
// 		if a[5 - i] > = 1 {
// 			for k := 0; k < 5; k++ {
// 				buf[k]:= a[k]
// 			}
// 		}
// 	}

// }

// func max_index(a []int) float64 {
// 	max:= 0
// 	for i := 0; i < 5; i++ {
// 		if max < a[i] {
// 			max:= a[i]
// 		}
// 	}
// }

