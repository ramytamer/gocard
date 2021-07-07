package models

import "fmt"

func (ms *ModelSuite) Test_IsValidCode() {
	suit := []string{"C", "D", "H", "S"}
	vals := []string{"K", "Q", "J", "A", "2", "3", "5", "1"}

	for _, val := range vals {
		for _, suit := range suit {
			ms.True(isValidCode(fmt.Sprintf("%s%s", val, suit)))
		}
	}

	ms.False(isValidCode("foo"))
	ms.False(isValidCode(""))
	ms.False(isValidCode("13H"))
}

func (ms *ModelSuite) Test_GetCardCode() {
	ms.Equal("AH", GetCardCode(1, HEARTS))
	ms.Equal("KH", GetCardCode(13, HEARTS))
	ms.Equal("QH", GetCardCode(12, HEARTS))
	ms.Equal("JH", GetCardCode(11, HEARTS))
	ms.Equal("2H", GetCardCode(2, HEARTS))
}

func (ms *ModelSuite) Test_MakeCardFromCode() {
	ah, _ := MakeCardFromCode("AH")
	ms.Equal(1, ah.Value)
	ms.Equal(HEARTS, ah.Suit)

	kh, _ := MakeCardFromCode("KH")
	ms.Equal(13, kh.Value)
	ms.Equal(HEARTS, kh.Suit)

	qh, _ := MakeCardFromCode("QH")
	ms.Equal(12, qh.Value)
	ms.Equal(HEARTS, qh.Suit)

	jh, _ := MakeCardFromCode("JH")
	ms.Equal(11, jh.Value)
	ms.Equal(HEARTS, jh.Suit)

	h3, _ := MakeCardFromCode("3H")
	ms.Equal(3, h3.Value)
	ms.Equal(HEARTS, h3.Suit)
}
