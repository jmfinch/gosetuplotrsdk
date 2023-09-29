package main

import (
	"log"

	"github.com/jmfinch/lotrsdk/src/movie"

)

func main() {
	testMovies()
}

func testMovies() {
	// Create a repository instance using NewMovieRepository()
	movieRepo := movie.NewMovieRepository()

	movies, err := movieRepo.GetMovies()
	if err != nil {
		log.Println(err)
	}
	for _, movie := range movies {
		log.Println(movie.ID, movie.Name)
	}

	TwoTowers, err := movieRepo.GetMovieByID("5cd95395de30eff6ebccde5b")
	if err != nil {
		log.Println(err)
	}
	log.Println(TwoTowers.Name, TwoTowers.BudgetInMillions)
	ReturnOftheKing, err := movieRepo.GetMovieByName("The Return of the King")
	if err != nil {
		log.Println(err)
	}
	log.Println(ReturnOftheKing.Name, ReturnOftheKing.AcademyAwardWins)
}