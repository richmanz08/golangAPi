package movies

import "time"

type IMovie struct {
	ID             int32   `json:"id"`
	MovieNameLocal string  `json:"movie_name_local" binding:"required"`
	MovieNameEng   string  `json:"movie_name_eng" binding:"required"`
	Episodes       int32   `json:"episodes" binding:"required"`
	Type           string  `json:"type" binding:"required"`
	Rating         float64 `json:"rating" binding:"required"`
	Duration       float64 `json:"duration" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	QualityType    string  `json:"quality_type" binding:"required"`
	Directors_id   string  `json:"directors_id" binding:"required"`
	Casters_id     string  `json:"casters_id" binding:"required"`
	PosterURL      string  `json:"poster_url" binding:"required"`
	DirectoryName  string  `json:"directory_name" binding:"required"`
	Year           int32   `json:"year" binding:"required"`
	CreateAt       *time.Time  `json:"create_at" `
	UpdateAt       *time.Time  `json:"update_at" `
	DeleteAt       *time.Time     `json:"delete_at" `
}