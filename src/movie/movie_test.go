package movie

import (
	"errors"
	"testing"
)

// MockMovieRepository is a mock implementation of the Repository interface.
type MockMovieRepository struct{}

func (mr *MockMovieRepository) GetMovies() ([]Movie, error) {
    // Implement the mock GetMovies method for testing.
    // Return a slice of Movie objects or an error.
    return []Movie{}, nil // Implement with appropriate data or error handling.
}

func (mr *MockMovieRepository) GetMovieByID(id string) (Movie, error) {
    // Implement the mock GetMovieByID method for testing.
    // Return a Movie object with the expected name or an error.
    if id == "5cd95395de30eff6ebccde5b" {
        return Movie{
            ID:   id,
            Name: "The Two Towers",
            // Add other fields as needed.
        }, nil
    }
    // Return an error for unknown movie IDs.
    return Movie{}, errors.New("movie not found")
}

func (mr *MockMovieRepository) GetMovieByName(movieName string) (Movie, error) {
    // Implement the mock GetMovieByName method for testing.
    // Return a Movie object or an error.
    return Movie{}, nil // Implement with appropriate data or error handling.
}

func (mr *MockMovieRepository) GetMovieNames() ([]string, error) {
    // Implement the mock GetMovieNames method for testing.
    // Return a slice of movie names or an error.
    return []string{}, nil // Implement with appropriate data or error handling.
}

func TestGetMovieByID(t *testing.T) {
    // Create a new instance of the mock repository for testing.
    repo := &MockMovieRepository{}

    movieID := "5cd95395de30eff6ebccde5b"
    movieName := "The Two Towers"

    // Create a new instance of the MovieService using the mock repository.
    service := NewMovieService(repo)

    got, err := service.GetMovieByID(movieID)
    want := movieName

    if err != nil {
        t.Errorf("encountered error %v", err)
    }
    if got.Name != want {
        t.Errorf("got %q, wanted %q", got.Name, want)
    }
}

// Add more test cases for other methods as needed...
