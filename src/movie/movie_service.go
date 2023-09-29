package movie

// MovieService implements the Service interface for movie-related business logic.
type MovieService struct {
    repo Repository
}

// NewMovieService creates a new MovieService instance.
func NewMovieService(repo Repository) Service {
    return &MovieService{
        repo: repo,
    }
}

// Service defines the methods that the MovieService should implement.
type Service interface {
    GetMovies() ([]Movie, error)
    GetMovieByID(id string) (Movie, error)
    GetMovieByName(movieName string) (Movie, error)
    GetMovieNames() ([]string, error)
}

func (ms *MovieService) GetMovies() ([]Movie, error) {
    return ms.repo.GetMovies()
}

func (ms *MovieService) GetMovieByID(id string) (Movie, error) {
    return ms.repo.GetMovieByID(id)
}

func (ms *MovieService) GetMovieByName(movieName string) (Movie, error) {
    return ms.repo.GetMovieByName(movieName)
}

func (ms *MovieService) GetMovieNames() ([]string, error) {
    return ms.repo.GetMovieNames()
}
