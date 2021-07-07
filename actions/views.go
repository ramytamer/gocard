package actions

import (
	"gocard/models"
	"math/rand"
	"time"
)

func renderCreateDeck(deck models.Deck) map[string]interface{} {
	cards := deck.Data["cards"].([]interface{})
	return map[string]interface{}{
		"deck_id":   deck.ID.String(),
		"shuffled":  deck.Shuffled,
		"remaining": len(cards),
	}
}

func renderShowDeck(deck models.Deck) map[string]interface{} {
	cards := deck.Data["cards"].([]interface{})
	if deck.Shuffled {
		cards = shuffleCards(cards)
	}
	return map[string]interface{}{
		"deck_id":   deck.ID.String(),
		"shuffled":  deck.Shuffled,
		"remaining": len(cards),
		"cards":     cards,
	}
}

func renderDrawnCards(cards []interface{}) map[string]interface{} {
	return map[string]interface{}{"cards": cards}
}

func shuffleCards(cards []interface{}) []interface{} {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

	return cards
}
