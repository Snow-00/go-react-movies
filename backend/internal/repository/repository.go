package repository

import (
	"database/sql"

	"github.com/Snow-00/go-react-movies-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies(genre ...int) ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)

	OneMovie(id int) (*models.Movie, error)                         // for the public to see
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error) // for edit purpose / admin only
	AllGenres() ([]*models.Genre, error)
	InsertMovie(movie models.Movie) (int, error)
	UpdateMovieGenres(id int, genreIDs []int) error
	UpdateMovie(movie models.Movie) error
	DeleteMovie(id int) error
}
