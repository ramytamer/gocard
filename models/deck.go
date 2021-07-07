package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/pop/v5/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type CardMap map[string]interface{}

// Deck is used by pop to map your decks database table to your go code.
type Deck struct {
	ID           uuid.UUID  `json:"deck_id" db:"id"`
	Shuffled     bool       `json:"shuffled" db:"shuffled"`
	Data         slices.Map `json:"-" db:"data"`
	Count        int        `json:"count" db:"-"`
	PartialCards string     `json:"partial_cards" db:"-"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d Deck) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Decks is not required by pop and may be deleted
type Decks []Deck

// String is not required by pop and may be deleted
func (d Decks) String() string {
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
