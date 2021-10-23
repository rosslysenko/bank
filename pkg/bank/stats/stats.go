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
