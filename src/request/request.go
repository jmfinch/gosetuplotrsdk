package request

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const LOTR_ENDPOINT string = "https://the-one-api.dev/v2"
const ERR_NOTOKEN = "LOTR_ACESS_TOKEN not set."

type Sorting struct {
	Field string
	Order string
}

type Filtering struct {
}

type Option func(*string)

func WithLimit(limit int) Option {
	return func(str *string) {
		*str += "limit=" + strconv.Itoa(limit)
	}
}

func WithOffset(offset int) Option {
	return func(str *string) {
		*str += "offset=" + strconv.Itoa(offset)
	}
}

func WithPage(page int) Option {
	return func(str *string) {
		*str += "page=" + strconv.Itoa(page)
	}
}

func WithSort(field, order string) Option {
	return func(str *string) {
		*str += "sort=" + field + ":" + order
	}
}

func CallApi(uri string, opts ...Option) ([]byte, error) {
	if len(opts) > 1 {
		updateUri(&uri, opts...)
	}

	req, err := buildRequest(uri)
	if err != nil {
		log.Printf("error building HTTP request: %v\n", err)
		return []byte{}, err
	}

	responseBytes, err := makeRequest(req)
	if err != nil {
		log.Printf("error with HTTP response: %v\n", err)
		return []byte{}, err
	}

	return responseBytes, nil

}

func buildRequest(uri string) (*http.Request, error) {

	request, err := http.NewRequest(
		http.MethodGet,
		LOTR_ENDPOINT+uri,
		nil,
	)
	if err != nil {
		log.Printf("error creating HTTP request: %v", err)
		return request, err
	}

	if os.Getenv("LOTR_ACCESS_TOKEN") == "" {
		return request, errors.New(ERR_NOTOKEN)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("LOTR_ACCESS_TOKEN"))
	return request, nil
}

func makeRequest(request *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("error sending HTTP request: %v", err)
		return []byte{}, err
	}
	if res.StatusCode == http.StatusUnauthorized {
		return []byte{}, errors.New("unauthorized request, check your API token")
	}
	if res.Header.Get("X-Ratelimit-Remaining") == "0" {
		return []byte{}, errors.New("api rate limit reached, rate limit will reset within 10 minutes")
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading HTTP response body: %v\n", err)
		return responseBytes, err
	}

	return responseBytes, nil
}

func updateUri(uri *string, opts ...Option) {
	*uri += "?"
	for i, opt := range opts {
		opt(uri)
		if i != len(opts)-1 {
			*uri = *uri + "&"
		}
	}

}