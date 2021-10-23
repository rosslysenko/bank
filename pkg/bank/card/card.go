package card

import (
	"github.com/rosslysenko/bank/pkg/bank/types"
	// "fmt"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	newCard := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return newCard
}
