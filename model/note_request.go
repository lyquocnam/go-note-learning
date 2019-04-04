package model

import validator "github.com/asaskevich/govalidator"

type NoteRequest struct {
	Title       *string `gorm:"unique" json:"title" valid:"required~Tiêu đề không được trống,runelength(1|80)~Tiêu đề phải từ 1 - 80 ký tự"`
	IsCompleted *bool   `json:"is_completed" valid:"required~Trạng thái hoàn tất không được rỗng"`
}

func (n *NoteRequest) Validate() (bool, error) {
	return validator.ValidateStruct(n)
}
