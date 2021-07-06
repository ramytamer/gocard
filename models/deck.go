package models

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type CardMap map[string]string

// Deck is used by pop to map your decks database table to your go code.
type Deck struct {
	ID           uuid.UUID `json:"deck_id" db:"id"`
	Shuffled     bool      `json:"shuffled" db:"shuffled"`
	Cards        []CardMap `json:"-" db:"cards"`
	PartialCards string    `json:"partial_cards" db:"-"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d Deck) String() string {
	cards := make([]CardMap, len(d.Cards))

	for i, card := range d.Cards {
		v, _ := strconv.Atoi(card["value"])
		s, _ := strconv.Atoi(card["suite"])

		c := MakeCard(v, s)

		cards[i] = c.ToMap()
	}

	if d.Shuffled {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	}

	d.Cards = cards

	jd, _ := json.Marshal(d)
	return string(jd)
}

// Decks is not required by pop and may be deleted
type Decks []Deck

// String is not required by pop and may be deleted
func (d Decks) String() string {
	// log.Println("+++++++++++++++++++++++++++++")
	// log.Println("+++++++++++++++++++++++++++++")
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Deck) Validate(tx *pop.Connection) (*validate.Errors, error) {
	// return validate.Validate(
	// 	&validators.FuncValidator{
	// 		Fn: func() bool {

	// 		},
	// 	}
	// ), nil
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Deck) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Deck) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
