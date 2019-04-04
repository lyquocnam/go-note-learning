package model

import (
	validator "github.com/asaskevich/govalidator"
	"time"
)

type Note struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Title       string     `gorm:"unique" json:"title" valid:"required~Tiêu đề không được trống,runelength(1|80)~Tiêu đề phải từ 1 - 80 ký tự"`
	IsCompleted bool       `json:"is_completed" valid:"required~Trạng thái hoàn tất không được rỗng"`
}

func (n *Note) Validate() (bool, error) {
	return validator.ValidateStruct(n)
}
