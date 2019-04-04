package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/lyquocnam/go-note-learning/model"
)

type notePostgresStorage struct {
	db *gorm.DB
}

func NewNotePostgresStorage(db *gorm.DB) *notePostgresStorage {
	return &notePostgresStorage{db: db}
}

func (n *notePostgresStorage) Get(id uint) (*model.Note, error) {
	var note model.Note
	err := n.db.New().First(&note, "id = ?", id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
	}
	return &note, err
}

func (n *notePostgresStorage) GetByTitle(title string) (*model.Note, error) {
	var note model.Note
	err := n.db.New().First(&note, "title = ?", title).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
	}
	return &note, err
}

func (n *notePostgresStorage) GetList() ([]*model.Note, error) {
	var notes []*model.Note
	err := n.db.New().Find(&notes).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
	}
	return notes, err
}

func (n *notePostgresStorage) Insert(note *model.Note) (*model.Note, error) {
	err := n.db.New().Create(&note).Error
	return note, err
}

func (n *notePostgresStorage) Update(id uint, note *model.Note) (*model.Note, error) {
	err := n.db.New().Save(&note).Error
	return note, err
}

func (n *notePostgresStorage) Delete(note *model.Note) error {
	return n.db.New().Unscoped().Delete(note).Error
}

func (n *notePostgresStorage) Count(where interface{}, args ...interface{}) (int, error) {
	count := 0
	err := n.db.New().Model(model.Note{}).Where(where, args...).Count(&count).Error
	return count, err
}
