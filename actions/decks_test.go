package actions

import (
	"encoding/json"
	"fmt"
	"gocard/models"
	"net/http"
)

func (as *ActionSuite) Test_DecksResource_Show_full_deck_normal() {
	as.LoadFixture("full_deck_normal")
	deck := models.Deck{}
	as.DB.First(&deck)

	response := as.JSON(fmt.Sprintf("/decks/%s", deck.ID)).Get()
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusOK, response.Code)
	as.Equal(false, responseData["shuffled"])
	as.Equal(52.0, responseData["remaining"])
	as.Equal(52, len(responseData["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Show_full_deck_shuffled() {
	as.LoadFixture("full_deck_shuffled")
	deck := models.Deck{}
	as.DB.First(&deck)

	response := as.JSON(fmt.Sprintf("/decks/%s", deck.ID)).Get()
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusOK, response.Code)
	as.Equal(true, responseData["shuffled"])
	as.Equal(52.0, responseData["remaining"])
	as.Equal(52, len(responseData["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Show_partial_deck_normal() {
	as.LoadFixture("partial_deck_normal")
	deck := models.Deck{}
	as.DB.First(&deck)

	response := as.JSON(fmt.Sprintf("/decks/%s", deck.ID)).Get()
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusOK, response.Code)
	as.Equal(false, responseData["shuffled"])
	as.Equal(2.0, responseData["remaining"])
	as.Equal(2, len(responseData["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Show_partial_deck_shuffled() {
	as.LoadFixture("partial_deck_shuffled")
	deck := models.Deck{}
	as.DB.First(&deck)

	response := as.JSON(fmt.Sprintf("/decks/%s", deck.ID)).Get()
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusOK, response.Code)
	as.Equal(true, responseData["shuffled"])
	as.Equal(2.0, responseData["remaining"])
	as.Equal(2, len(responseData["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Show_not_found() {
	response := as.JSON(fmt.Sprintf("/decks/%s", "404")).Get()
	as.Equal(http.StatusNotFound, response.Code)
}

func (as *ActionSuite) Test_DecksResource_Create_full_deck_normal_by_default() {
	response := as.JSON(fmt.Sprintf("/decks")).Post(
		map[string]interface{}{},
	)
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusCreated, response.Code)

	deck := models.Deck{}
	as.DB.First(&deck)

	as.False(deck.Shuffled)
	as.Equal(52, len(deck.Data["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Create_full_deck_shuffled() {
	response := as.JSON(fmt.Sprintf("/decks")).Post(
		map[string]interface{}{"shuffled": true},
	)
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusCreated, response.Code)

	deck := models.Deck{}
	as.DB.First(&deck)

	as.True(deck.Shuffled)
	as.Equal(52, len(deck.Data["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Create_full_deck_if_partial_cards_invalid() {
	response := as.JSON(fmt.Sprintf("/decks")).Post(
		map[string]interface{}{"partial_cards": "foo,bar"},
	)
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusCreated, response.Code)

	deck := models.Deck{}
	as.DB.First(&deck)

	as.False(deck.Shuffled)
	as.Equal(52, len(deck.Data["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Create_partial_cards() {
	response := as.JSON(fmt.Sprintf("/decks")).Post(
		map[string]interface{}{"partial_cards": "AH,KC,3S", "shuffled": true},
	)
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusCreated, response.Code)

	deck := models.Deck{}
	as.DB.First(&deck)

	as.True(deck.Shuffled)
	as.Equal(3, len(deck.Data["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Draw_full_deck_normal() {
	as.LoadFixture("full_deck_normal")
	deck := models.Deck{}
	as.DB.First(&deck)

	response := as.JSON(fmt.Sprintf("/decks/%s/draw", deck.ID)).Put(
		map[string]interface{}{"count": 1},
	)
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusOK, response.Code)
	as.Equal(1, len(responseData["cards"].([]interface{})))

	deck = models.Deck{}
	as.DB.First(&deck)

	as.Equal(51, len(deck.Data["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Draw_partial_deck_normal() {
	as.LoadFixture("partial_deck_normal")
	deck := models.Deck{}
	as.DB.First(&deck)

	response := as.JSON(fmt.Sprintf("/decks/%s/draw", deck.ID)).Put(
		map[string]interface{}{"count": 2},
	)
	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	as.Equal(http.StatusOK, response.Code)
	as.Equal(2, len(responseData["cards"].([]interface{})))

	deck = models.Deck{}
	as.DB.First(&deck)

	as.Equal(0, len(deck.Data["cards"].([]interface{})))
}

func (as *ActionSuite) Test_DecksResource_Draw_handle_errors() {
	as.LoadFixture("partial_deck_normal")
	deck := models.Deck{}
	as.DB.First(&deck)

	response := as.JSON(fmt.Sprintf("/decks/%s/draw", deck.ID)).Put(map[string]interface{}{"count": 0})
	as.Equal(http.StatusUnprocessableEntity, response.Code)

	response = as.JSON(fmt.Sprintf("/decks/%s/draw", deck.ID)).Put(map[string]interface{}{"count": -1})
	as.Equal(http.StatusUnprocessableEntity, response.Code)

	response = as.JSON(fmt.Sprintf("/decks/%s/draw", deck.ID)).Put(map[string]interface{}{"count": -2})
	as.Equal(http.StatusUnprocessableEntity, response.Code)

	response = as.JSON(fmt.Sprintf("/decks/%s/draw", deck.ID)).Put(map[string]interface{}{"count": 53})
	as.Equal(http.StatusUnprocessableEntity, response.Code)

	cardsCount := len(deck.Data["cards"].([]interface{}))
	response = as.JSON(fmt.Sprintf("/decks/%s/draw", deck.ID)).Put(map[string]interface{}{"count": cardsCount + 1})
	as.Equal(http.StatusUnprocessableEntity, response.Code)

	response = as.JSON(fmt.Sprintf("/decks/%s/draw", "404")).Put(map[string]interface{}{"count": 1})
	as.Equal(http.StatusNotFound, response.Code)
}
