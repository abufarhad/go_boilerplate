package model

import "time"

type AboutUs struct {
	Id             uint       `json:"id"`
	Title          string     `json:"title"`
	SubTitle       string     `json:"sub_title"`
	Description    string     `json:"description"`
	TitleImagePath string     `json:"title_image_path"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
