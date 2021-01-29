package model

type Article struct {
	Model

	TagId int `json:"tag_id" gorm:"index"`
}
