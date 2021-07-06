package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	CLUBS int = iota
	DIAMOND
	HEARTS
	SPADES
)

type Card struct {
	Value int `json:"value"`
	Suit  int `json:"suit"`
}

func (c Card) ToMap() CardMap {
	return CardMap{
		"value": strconv.Itoa(c.Value),
		"suit":  strconv.Itoa(c.Suit),
		"code":  GetCardCode(c.Value, c.Suit),
	}
}

func MakeCard(value, suit int) *Card {
	return &Card{
		Value: value,
		Suit:  suit,
	}
}

func MakeCardFromCode(code string) (*Card, error) {
	if !isValidCode(code) {
		return nil, errors.New("Invalid Code!")
	}

	splittedCode := strings.Split(code, "")

	var s, v int

	switch splittedCode[0] {
	case "J":
		v = 11
	case "Q":
		v = 12
	case "K":
		v = 13
	case "A":
		v = 1
	default:
		v, _ = strconv.Atoi(splittedCode[0])
	}

	switch splittedCode[1] {
	case "C":
		s = CLUBS
	case "D":
		s = DIAMOND
	case "H":
		s = HEARTS
	case "S":
		s = SPADES
	}

	return &Card{
		Value: v,
		Suit:  s,
	}, nil
}

func GetCardCode(value, suit int) string {
	v := strconv.Itoa(value)

	switch value {
	case 1:
		v = "A"
	case 11:
		v = "J"
	case 12:
		v = "Q"
	case 13:
		v = "K"
	}

	var s string

	switch suit {
	case CLUBS:
		s = "C"
	case DIAMOND:
		s = "D"
	case HEARTS:
		s = "H"
	case SPADES:
		s = "S"
	}

	return fmt.Sprintf("%s%s", v, s)
}

func isValidCode(code string) bool {
	if len(code) != 2 {
		log.Println("invalid length", code, len(code))
		return false
	}

	splittedCode := strings.Split(code, "")
	return isValidValue(splittedCode[0]) && isValidSuit(splittedCode[1])
}

func isValidSuit(s string) bool {
	switch s {
	case "C", "D", "H", "S":
		return true
	}

	log.Println("invalid suit", s)
	return false
}

func isValidValue(v string) bool {
	switch v {
	case "J", "Q", "K", "A":
		return true
	default:
		val, _ := strconv.Atoi(v)

		return val >= 1 && val <= 10
	}
}
