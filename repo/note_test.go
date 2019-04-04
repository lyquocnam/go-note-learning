package repo

import (
	"errors"
	"github.com/lyquocnam/go-note-learning/lib"
	"github.com/lyquocnam/go-note-learning/mocks"
	"github.com/lyquocnam/go-note-learning/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNoteRepo_Get(t *testing.T) {
	note := model.Note{
		ID:          1,
		Title:       "Hello",
		IsCompleted: true,
	}
	cases := []struct {
		name   string
		expect *model.Note
		err    error
	}{
		{
			name:   "case get note ok",
			expect: &note,
			err:    nil,
		},
		{
			name:   "case get note not found",
			expect: nil,
			err:    errors.New(lib.NoteNotExistError),
		},
		{
			name:   "case can not get",
			expect: nil,
			err:    errors.New("can not get data"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockStorage := &mocks.NoteStorage{}
			repo := NewNoteRepo(mockStorage)
			mockStorage.On("Get", note.ID).Return(c.expect, c.err)
			actual, err := repo.Get(note.ID)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expect, actual)
		})
	}

}

func TestNoteRepo_GetList(t *testing.T) {
	notes := make([]*model.Note, 0)
	notes = append(notes, &model.Note{
		ID:          1,
		Title:       "Hello",
		IsCompleted: true,
	})
	notes = append(notes, &model.Note{
		ID:          2,
		Title:       "Hello 2",
		IsCompleted: false,
	})
	cases := []struct {
		name   string
		expect []*model.Note
		err    error
	}{
		{
			name:   "case get note ok",
			expect: notes,
			err:    nil,
		},
		{
			name:   "case can not get data",
			expect: nil,
			err:    errors.New("can not get data"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockStorage := &mocks.NoteStorage{}
			repo := NewNoteRepo(mockStorage)
			mockStorage.On("GetList").Return(c.expect, c.err)
			actual, err := repo.GetList()
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expect, actual)
		})
	}
}

func TestNoteRepo_ExistByTitle(t *testing.T) {
	note := model.Note{
		ID:          1,
		Title:       "Hello",
		IsCompleted: true,
	}
	cases := []struct {
		name   string
		expect bool
		count  int
		err    error
	}{
		{
			name:   "case get note ok",
			expect: true,
			count:  1,
			err:    nil,
		},
		{
			name:   "case can not get",
			expect: false,
			count:  0,
			err:    errors.New("can not get data"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockStorage := &mocks.NoteStorage{}
			repo := NewNoteRepo(mockStorage)
			mockStorage.On("Count", "title = ?", note.Title).Return(c.count, c.err)
			actual, err := repo.ExistByTitle(note.Title)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expect, actual)
		})
	}
}

func TestNoteRepo_Exist(t *testing.T) {
	note := model.Note{
		ID:          1,
		Title:       "Hello",
		IsCompleted: true,
	}
	cases := []struct {
		name   string
		expect bool
		count  int
		err    error
	}{
		{
			name:   "case get note ok",
			expect: true,
			count:  1,
			err:    nil,
		},
		{
			name:   "case can not get",
			expect: false,
			count:  0,
			err:    errors.New("can not get data"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockStorage := &mocks.NoteStorage{}
			repo := NewNoteRepo(mockStorage)
			mockStorage.On("Count", "id = ?", note.ID).Return(c.count, c.err)
			actual, err := repo.Exist(note.ID)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expect, actual)
		})
	}
}

func TestNoteRepo_Insert(t *testing.T) {
	isCompleted := false
	note := model.Note{
		ID:          1,
		Title:       "Hello",
		IsCompleted: true,
	}
	before := model.Note{
		Title: "Hello",
	}
	request := model.NoteRequest{
		Title:       &note.Title,
		IsCompleted: &isCompleted,
	}
	cases := []struct {
		name             string
		expect           *model.Note
		request          model.NoteRequest
		beforeInsert     *model.Note
		afterInsert      *model.Note
		code             int
		err              error
		getByTitleResult *model.Note
		getByTitleErr    error
	}{
		{
			name:             "case 1: get note ok",
			expect:           &note,
			beforeInsert:     &before,
			afterInsert:      &note,
			request:          request,
			code:             200,
			err:              nil,
			getByTitleResult: nil,
			getByTitleErr:    nil,
		},
		{
			name:             "case 2: title has already exist",
			expect:           nil,
			request:          request,
			code:             http.StatusConflict,
			err:              errors.New(lib.NoteTitleAlreadyExistError),
			getByTitleResult: &note,
			getByTitleErr:    nil,
		},
		{
			name:   "case 3: title is not valid",
			expect: nil,
			request: model.NoteRequest{
				Title: nil,
			},
			code:             http.StatusBadRequest,
			err:              errors.New(lib.NoteTitleRequired),
			getByTitleResult: nil,
			getByTitleErr:    nil,
		},
		{
			name:             "case 4: getByTitle has error",
			expect:           nil,
			request:          request,
			code:             http.StatusInternalServerError,
			err:              errors.New(""),
			getByTitleResult: nil,
			getByTitleErr:    errors.New(""),
		},
		{
			name:             "case 5: Can not insert to db",
			expect:           nil,
			request:          request,
			beforeInsert:     &before,
			afterInsert:      &note,
			code:             http.StatusInternalServerError,
			err:              errors.New(""),
			getByTitleResult: nil,
			getByTitleErr:    nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockStorage := &mocks.NoteStorage{}
			repo := NewNoteRepo(mockStorage)
			mockStorage.On("GetByTitle", note.Title).Return(c.getByTitleResult, c.getByTitleErr)
			mockStorage.On("Insert", c.beforeInsert).Return(c.afterInsert, c.err)
			actual, code, err := repo.Insert(&c.request)
			assert.Equal(t, c.code, code)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expect, actual)
		})
	}
}

func TestNoteRepo_Update(t *testing.T) {
	isCompleted := false
	var noteId uint = 1
	note := model.Note{
		ID:          noteId,
		Title:       "Hello",
		IsCompleted: true,
	}
	before := model.Note{
		ID:    noteId,
		Title: "Hello",
	}
	request := model.NoteRequest{
		Title:       &note.Title,
		IsCompleted: &isCompleted,
	}
	cases := []struct {
		name         string
		expect       *model.Note
		request      model.NoteRequest
		beforeUpdate *model.Note
		afterUpdate  *model.Note
		code         int
		err          error
		getResult    *model.Note
		getErr       error
	}{
		{
			name:         "case 1: get note ok",
			expect:       &note,
			beforeUpdate: &before,
			afterUpdate:  &note,
			request:      request,
			code:         200,
			err:          nil,
			getResult:    &note,
			getErr:       nil,
		},
		{
			name:         "case 3: can not get note",
			expect:       nil,
			beforeUpdate: &before,
			afterUpdate:  &note,
			request:      request,
			code:         500,
			err:          errors.New("can not get note"),
			getResult:    nil,
			getErr:       errors.New("can not get note"),
		},
		{
			name:         "case 2: note not exist",
			expect:       nil,
			beforeUpdate: &before,
			afterUpdate:  &note,
			request:      request,
			code:         404,
			err:          errors.New(lib.NoteNotExistError),
			getResult:    nil,
			getErr:       nil,
		},
		{
			name:         "case 3: can not update note",
			expect:       nil,
			beforeUpdate: &before,
			afterUpdate:  &note,
			request:      request,
			code:         500,
			err:          errors.New("can not update note"),
			getResult:    &note,
			getErr:       nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockStorage := &mocks.NoteStorage{}
			repo := NewNoteRepo(mockStorage)
			mockStorage.On("Get", noteId).Return(c.getResult, c.getErr)
			mockStorage.On("Update", noteId, c.beforeUpdate).Return(c.afterUpdate, c.err)
			actual, code, err := repo.Update(noteId, &c.request)
			assert.Equal(t, c.code, code)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expect, actual)
		})
	}
}

func TestNoteRepo_Delete(t *testing.T) {
	isCompleted := false
	var noteId uint = 1
	note := model.Note{
		ID:          noteId,
		Title:       "Hello",
		IsCompleted: true,
	}
	var noteIdZero uint = 0
	request := model.NoteRequest{
		Title:       &note.Title,
		IsCompleted: &isCompleted,
	}
	cases := []struct {
		name         string
		expect       *uint
		request      model.NoteRequest
		beforeDelete *model.Note
		code         int
		err          error
		getResult    *model.Note
		getErr       error
	}{
		{
			name:         "case 1: get note ok",
			expect:       &noteId,
			beforeDelete: &note,
			request:      request,
			code:         200,
			err:          nil,
			getResult:    &note,
			getErr:       nil,
		},
		{
			name:      "case 2: can not get note",
			expect:    &noteIdZero,
			request:   request,
			code:      500,
			err:       errors.New("can not get note"),
			getResult: nil,
			getErr:    errors.New("can not get note"),
		},
		{
			name:      "case 3: note not exist",
			expect:    &noteIdZero,
			request:   request,
			code:      404,
			err:       errors.New(lib.NoteNotExistError),
			getResult: nil,
			getErr:    nil,
		},
		{
			name:         "case 4: can not update note",
			expect:       &noteIdZero,
			beforeDelete: &note,
			request:      request,
			code:         500,
			err:          errors.New("can not update note"),
			getResult:    &note,
			getErr:       nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockStorage := &mocks.NoteStorage{}
			repo := NewNoteRepo(mockStorage)
			mockStorage.On("Get", noteId).Return(c.getResult, c.getErr)
			mockStorage.On("Delete", c.beforeDelete).Return(c.err)
			actual, code, err := repo.Delete(noteId)
			assert.Equal(t, c.code, code)
			assert.Equal(t, c.err, err)
			assert.Equal(t, *c.expect, actual)
		})
	}
}
