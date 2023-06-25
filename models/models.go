package models

type Stock struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageID     string `json:"image_id"`
}

type Image struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	ImageURL     string `json:"image_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}
