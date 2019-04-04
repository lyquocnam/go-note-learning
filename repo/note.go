package repo

import (
	"errors"
	"github.com/lyquocnam/go-note-learning/lib"
	"github.com/lyquocnam/go-note-learning/model"
	"github.com/lyquocnam/go-note-learning/storage"
	"net/http"
)

type noteRepo struct {
	noteStorage storage.NoteStorage
}

func NewNoteRepo(noteStorage storage.NoteStorage) *noteRepo {
	return &noteRepo{noteStorage: noteStorage}
}

type NoteRepo interface {
	Get(id uint) (*model.Note, error)
	GetList() ([]*model.Note, error)
	ExistByTitle(title string) (bool, error)
	Exist(id uint) (bool, error)
	Insert(note *model.NoteRequest) (*model.Note, int, error)
	Update(id uint, request *model.NoteRequest) (*model.Note, int, error)
	Delete(id uint) (uint, int, error)
}

func (r *noteRepo) Get(id uint) (*model.Note, error) {
	return r.noteStorage.Get(id)
}

func (r *noteRepo) GetList() ([]*model.Note, error) {
	return r.noteStorage.GetList()
}

func (r *noteRepo) ExistByTitle(title string) (bool, error) {
	count, err := r.noteStorage.Count("title = ?", title)
	return count > 0, err
}

func (r *noteRepo) Exist(id uint) (bool, error) {
	count, err := r.noteStorage.Count("id = ?", id)
	return count > 0, err
}

func (r *noteRepo) Insert(request *model.NoteRequest) (*model.Note, int, error) {
	if request.Title == nil {
		return nil, http.StatusBadRequest, errors.New(lib.NoteTitleRequired)
	}

	note, err := r.noteStorage.GetByTitle(*request.Title)
	if err != nil {
		return nil, 500, err
	}
	if note != nil {
		return nil, http.StatusConflict, errors.New(lib.NoteTitleAlreadyExistError)
	}

	note = &model.Note{
		Title: *request.Title,
	}
	if request.IsCompleted != nil {
		note.IsCompleted = *request.IsCompleted
	}

	result, err := r.noteStorage.Insert(note)
	if err != nil {
		return nil, 500, err
	}
	return result, 200, nil
}

func (r *noteRepo) Update(id uint, request *model.NoteRequest) (*model.Note, int, error) {
	note, err := r.noteStorage.Get(id)
	// exist, err := u.noteRepo.Exist(id)
	if err != nil {
		return nil, 500, err
	}
	if note == nil {
		return nil, http.StatusNotFound, errors.New(lib.NoteNotExistError)
	}

	if request.Title != nil {
		note.Title = *request.Title
	}
	if request.IsCompleted != nil {
		note.IsCompleted = *request.IsCompleted
	}

	result, err := r.noteStorage.Update(id, note)
	if err != nil {
		return nil, 500, err
	}
	return result, 200, nil
}

func (r *noteRepo) Delete(id uint) (uint, int, error) {
	note, err := r.noteStorage.Get(id)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	if note == nil {
		return 0, http.StatusNotFound, errors.New(lib.NoteNotExistError)
	}
	err = r.noteStorage.Delete(note)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return note.ID, 200, nil
}
