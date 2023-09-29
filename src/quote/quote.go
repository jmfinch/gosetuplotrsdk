package quote

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/jmfinch/lotrsdk/src/request"
)

type QuoteResponse struct {
	Quotes []Quote `json:"docs"`
	Total  int     `json:"total"`
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Page   int     `json:"page"`
	Pages  int     `json:"pages"`
}

type Quote struct {
	ID          string `json:"_id"`
	Dialog      string `json:"dialog"`
	MovieID     string `json:"movie"`
	CharacterID string `json:"character"`
}

func GetAllQuotes() ([]Quote, error) {
	response, err := request.CallApi("/quote")
	if err != nil {
		log.Printf("error calling API %v\n", err)
		return []Quote{}, err
	}
	quotes, err := deserializeQuotes(response)
	if err != nil {
		return []Quote{}, err
	}

	return quotes, nil
}

func GetQuotes(opts ...request.Option) ([]Quote, error) {

	response, err := request.CallApi("/quote", opts...)
	if err != nil {
		log.Printf("error calling API %v\n", err)
		return []Quote{}, err
	}
	quotes, err := deserializeQuotes(response)
	if err != nil {
		return []Quote{}, err
	}
	return quotes, nil
}

func GetQuoteByID(id string) (Quote, error) {
	response, err := request.CallApi("/quote/" + id)
	if err != nil {
		log.Printf("error calling API %v\n", err)
		return Quote{}, err
	}
	quotes, err := deserializeQuotes(response)
	if err != nil {
		return Quote{}, err
	}
	if len(quotes) > 0 {
		return quotes[0], nil
	}

	return Quote{}, errors.New("malformed API response")
}

func GetQuotesFromMovie(movieID string) ([]Quote, error) {
	response, err := request.CallApi("/movie/" + movieID + "/quote")
	if err != nil {
		log.Printf("error calling API %v\n", err)
		return []Quote{}, err
	}
	quotes, err := deserializeQuotes(response)
	if err != nil {
		return []Quote{}, err
	}

	return quotes, nil
}

func GetQuoteByDialog(quoteDialog string) (Quote, error) {
	quotes, err := GetQuotes()
	if err != nil {
		log.Printf("error calling API %v\n", err)
		return Quote{}, err
	}
	for _, quote := range quotes {
		if quote.Dialog == quoteDialog {
			return quote, nil
		}
	}
	errorMessage := fmt.Sprintf("Cannont find matching quote for %s in Lord of the Rings movies", quoteDialog)
	return Quote{}, errors.New(errorMessage)
}

func deserializeQuotes(responceBytes []byte) ([]Quote, error) {
	var quotesResponse QuoteResponse
	if err := json.Unmarshal(responceBytes, &quotesResponse); err != nil {
		log.Printf("error deserializing json data %v\n", err)
		return []Quote{}, err
	}
	return quotesResponse.Quotes, nil
}