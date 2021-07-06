package actions

import "gocard/models"

func renderCreateDeck(deck models.Deck) map[string]interface{} {
	return map[string]interface{}{
		"deck_id":   deck.ID.String(),
		"shuffled":  deck.Shuffled,
		"remaining": len(deck.Cards),
	}
}
