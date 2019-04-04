package storage

import (
	"github.com/lyquocnam/go-note-learning/model"
	"gopkg.in/mgo.v2"
)

type noteMongo struct {
	noteCollection *mgo.Collection
}

func (m *noteMongo) Get(id uint) (*model.Note, error) {
	var result model.Note
	err := m.noteCollection.FindId(id).One(&result)
	return &result, err
}

func (m *noteMongo) GetByTitle(title string) (*model.Note, error) {
	panic("implement me")
}

func (m *noteMongo) GetList() ([]*model.Note, error) {
	panic("implement me")
}

func (m *noteMongo) Insert(note *model.Note) (*model.Note, error) {
	panic("implement me")
}

func (m *noteMongo) Update(id uint, note *model.Note) (*model.Note, error) {
	panic("implement me")
}

func (m *noteMongo) Delete(note *model.Note) error {
	panic("implement me")
}

func (m *noteMongo) Count(where interface{}, args ...interface{}) (int, error) {
	panic("implement me")
}
