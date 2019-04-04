package storage

import "github.com/lyquocnam/go-note-learning/model"

type NoteStorage interface {
	Get(id uint) (*model.Note, error)
	GetByTitle(title string) (*model.Note, error)
	GetList() ([]*model.Note, error)
	Insert(note *model.Note) (*model.Note, error)
	Update(id uint, note *model.Note) (*model.Note, error)
	Delete(note *model.Note) error
	Count(where interface{}, args ...interface{}) (int, error)
}
