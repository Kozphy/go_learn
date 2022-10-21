package slidingwindows

import "fmt"

func maxProfit(prices []int) int {
	var max_profit int = 0
	smaller_value := 100000

	for _, price := range prices {
		if smaller_value > price {
			smaller_value = price
		}

		if price < smaller_value {
			continue
		}

		if (price - smaller_value) > max_profit {
			max_profit = price - smaller_value
		}

	}
	return max_profit
}

func Execute_maxProfit() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
