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

// func (as *ActionSuite) Test_DecksResource_Create() {
//   as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_DecksResource_Update() {
//   as.Fail("Not Implemented!")
// }
