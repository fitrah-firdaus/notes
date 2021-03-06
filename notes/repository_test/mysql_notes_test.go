package repository_test

import (
	"github.com/fitrah-firdaus/notes/models"
	"github.com/fitrah-firdaus/notes/notes"
	"github.com/fitrah-firdaus/notes/notes/repository"
	"github.com/fitrah-firdaus/notes/notes/testhelper"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

var notesRepository notes.Repository

func init() {
	db, _ := testhelper.OpenDB()
	db.DropTable(&models.Notes{})
	db.AutoMigrate(&models.Notes{})
	notesRepository = repository.NewMysqlNotesRepository(db)
}

func TestSaveNotes(t *testing.T) {
	newNotes := &models.Notes{
		Title:   "Title_test",
		Content: "Content_test",
		Author:  "Author_tes",
	}

	savedNotes, err := notesRepository.Save(newNotes)

	if err != nil {
		t.Errorf("this is error saving notes: %v\n", err)
		return
	}

	assert.Equal(t, newNotes.Title, savedNotes.Title)
}
