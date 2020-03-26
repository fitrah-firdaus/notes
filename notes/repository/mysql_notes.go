package repository

import (
	"github.com/fitrah-firdaus/notes/models"
	"github.com/fitrah-firdaus/notes/notes"
	"github.com/jinzhu/gorm"
)

type mysqlNotesRepository struct {
	DB *gorm.DB
}

func NewMysqlNotesRepository(db *gorm.DB) notes.Repository {
	return &mysqlNotesRepository{db}
}

func (m mysqlNotesRepository) GetByID(id int64) (*models.Notes, error) {
	var n *models.Notes

	if err := m.DB.First(&n, id).Error; err != nil {
		return nil, err
	}

	return n, nil
}

func (m mysqlNotesRepository) GetByTitle(title string) (*models.Notes, error) {
	var n *models.Notes

	if err := m.DB.First(&n, "title = ?", title).Error; err != nil {
		return &models.Notes{}, err
	}

	return n, nil
}

func (m mysqlNotesRepository) Save(notes *models.Notes) (*models.Notes, error) {
	if err := m.DB.Create(notes).Error; err != nil {
		return notes, err
	}
	return notes, nil
}

func (m mysqlNotesRepository) Update(notes *models.Notes) (*models.Notes, error) {
	if err := m.DB.Update(notes).Error; err != nil {
		return notes, err
	}
	return notes, nil
}

func (m mysqlNotesRepository) Delete(id int64) (int64, error) {
	var note *models.Notes

	db := m.DB.Delete(&note, "id = ?", id)

	if err := db.Error; err != nil {
		return 0, err
	}

	return db.RowsAffected, nil
}
