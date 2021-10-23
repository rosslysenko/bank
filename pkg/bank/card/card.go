package card

import (
	"github.com/rosslysenko/bank/pkg/bank/types"
)

// Withdraws money from the card
func Withdraw(card types.Card, amount types.Money) types.Card {
	if amount > card.Balance || amount < 0 || amount > 2000000 || !card.Active {
		return card
	}
	card.Balance -= amount
	return card
}

// IssueCard creates a card instance with certain fields
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
