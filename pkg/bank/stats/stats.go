package stats

import (
	"github.com/rosslysenko/bank/pkg/bank/types"
)

// Avg calculates the average payment amount
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	var numberPayments types.Money
	for _, payment := range payments {
		sum += payment.Amount
		numberPayments++
	}
	return sum / numberPayments
}

// TotalInCategory counts the amount of payments by category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sumPerCategory types.Money
	for _, payment := range payments {
		if payment.Category == category {
			sumPerCategory += payment.Amount
		}
	}
	return sumPerCategory
}
