package notes

import (
	"github.com/fitrah-firdaus/notes/models"
)

type Repository interface {
	GetByID(id int64) (*models.Notes, error)
	GetByTitle(title string) (*models.Notes, error)
	Save(notes *models.Notes) (*models.Notes, error)
	Update(notes *models.Notes) (*models.Notes, error)
	Delete(id int64) (int64, error)
}
