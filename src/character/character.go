package character

import (
	"encoding/json"
	"errors"

	"github.com/jmfinch/lotrsdk/src/request"
)

type characterResponse struct {
	Characters []Character `json:"docs"`
	Total      int         `json:"total"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	Page       int         `json:"page"`
	Pages      int         `json:"pages"`
}

type Character struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	WikiUrl string `json:"wikiUrl"`
	Race    string `json:"race"`
}

func GetCharacters(opts ...request.Option) ([]Character, error) {

	response, err := request.CallApi("/character", opts...)
	if err != nil {
		return []Character{}, err
	}
	characters, err := deserializeCharacters(response)
	if err != nil {
		return []Character{}, err
	}
	return characters, nil
}

func GetCharacterByID(id string) (Character, error) {
	response, err := request.CallApi("/character/" + id)
	if err != nil {
		return Character{}, err
	}
	characters, err := deserializeCharacters(response)
	if err != nil {
		return Character{}, err
	}
	if len(characters) > 0 {
		return characters[0], nil
	}

	return Character{}, errors.New("empty character returned")
}
func deserializeCharacters(resp []byte) ([]Character, error) {
	var charResp characterResponse
	if err := json.Unmarshal(resp, &charResp); err != nil {
		return []Character{}, err
	}
	return charResp.Characters, nil
}