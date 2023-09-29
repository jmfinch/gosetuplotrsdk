package main

import (
	"log"

	"github.com/jmfinch/lotrsdk/src/character"
	"github.com/jmfinch/lotrsdk/src/movie"
	"github.com/jmfinch/lotrsdk/src/quote"
	"github.com/jmfinch/lotrsdk/src/request"
)

func main() {
	testMovies()
}

func testMovies() {
	movies, err := movie.GetMovies()
	if err != nil {
		log.Println(err)
	}
	for _, movie := range movies {
		log.Println(movie.ID, movie.Name)
	}

	TwoTowers, err := movie.GetMovieByID("5cd95395de30eff6ebccde5b")
	if err != nil {
		log.Println(err)
	}
	log.Println(TwoTowers.Name, TwoTowers.BudgetInMillions)
	ReturnOftheKing, err := movie.GetMovieByName("The Return of the King")
	if err != nil {
		log.Println(err)
	}
	log.Println(ReturnOftheKing.Name, ReturnOftheKing.AcademyAwardWins)
}

func testQuotes() {
	//Getting all quotes
	quotes, err := quote.GetQuotes(
		request.WithSort("character", "asc"),
		request.WithLimit(10),
		request.WithPage(3),
	)
	if err != nil {
		log.Println(err)
	}
	for _, quote := range quotes {
		char, err := character.GetCharacterByID(quote.CharacterID)
		if err != nil {
			log.Println("ooof")
		}
		log.Println(char.Name, quote.Dialog)
	}

}

func quoteHolder() {
	movieId := "5cd95395de30eff6ebccde5b"
	testQuotes, err := quote.GetQuotesFromMovie(movieId)
	if err != nil {
		log.Println(err)
	}
	for _, quote := range testQuotes {
		log.Println(quote.ID, quote.Dialog)
	}

	//Get Quote by ID
	quoteId := "5cd96e05de30eff6ebccebb1"
	testQuote, err := quote.GetQuoteByID(quoteId)
	if err != nil {
		log.Println(err)
	}
	log.Println(testQuote.Dialog, testQuote.ID)
	//quoteDialog := "Hey, stinker, don't go gettingtoo far ahead."
}