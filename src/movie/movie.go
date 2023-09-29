package movie

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"

	"github.com/jmfinch/lotrsdk/src/request"
)

// Define the Repository interface and MovieRepository struct.
type Repository interface {
    GetMovies() ([]Movie, error)
    GetMovieByID(id string) (Movie, error)
    GetMovieByName(movieName string) (Movie, error)
    GetMovieNames() ([]string, error)
}

type MovieRepository struct{}

func NewMovieRepository() Repository {
    return &MovieRepository{}
}

type movieResponse struct {
	Movies []Movie `json:"docs"`
	Total  int     `json:"total"`
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Page   int     `json:"page"`
	Pages  int     `json:"pages"`
}

type Movie struct {
	ID                         string  `json:"_id"`
	Name                       string  `json:"name"`
	RuntimeInMinutes           float32 `json:"runtimeInMinutes"`
	BudgetInMillions           float32 `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions float32 `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int     `json:"academyAwardNominations"`
	AcademyAwardWins           int     `json:"academyAwardWins"`
	RottenTomatoesScore        float32 `json:"rottenTomatoesScore"`
}

func (mr *MovieRepository) GetMovies() ([]Movie, error) {
    response, err := request.CallApi("/movie")
    if err != nil {
        log.Printf("error calling API %v\n", err)
        return nil, err
    }
    movies, err := deserializeMovies(response)
    if err != nil {
        return nil, err
    }
    return movies, nil
}

func (mr *MovieRepository) GetMovieByID(id string) (Movie, error) {
    response, err := request.CallApi("/movie/" + id)
    if err != nil {
        log.Printf("error calling API %v\n", err)
        return Movie{}, err
    }
    movies, err := deserializeMovies(response)
    if err != nil {
        return Movie{}, err
    }
    if len(movies) > 0 {
        return movies[0], nil
    }

    return Movie{}, errors.New("malformed API response")
}

func (mr *MovieRepository) GetMovieByName(movieName string) (Movie, error) {
    movies, err := mr.GetMovies()
    if err != nil {
        return Movie{}, err
    }
    for _, movie := range movies {
        if movie.Name == movieName {
            return movie, nil
        }
    }
    errorMessage := fmt.Sprintf("Cannont find matching name for %s in Lord of the Rings movies", movieName)
    return Movie{}, errors.New(errorMessage)
}

func (mr *MovieRepository) GetMovieNames() ([]string, error) {
    var result []string
    movies, err := mr.GetMovies()
    if err != nil {
        return nil, err
    }
    for _, movie := range movies {
        result = append(result, movie.Name)
    }
    return result, nil
}

func deserializeMovies(responseBytes []byte) ([]Movie, error) {
    var movieResp movieResponse
    if err := json.Unmarshal(responseBytes, &movieResp); err != nil {
        log.Printf("error deserializing json data %v\n", err)
        return nil, err
    }
    return movieResp.Movies, nil
}