package movies

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

// ----------------- Table
type Movie struct {
	ID            uint            `gorm:"primaryKey;autoIncrement" `
	MovieGroupID  uint            `gorm:"not null" `
	NameLocal     string          `gorm:"unique"`
	NameEng       string          `gorm:"unique" `
	Type          string          `gorm:"not null" ` // MOVIE OR SERIES
	Status        string          `gorm:"not null" ` // ACTIVE OR INACTIVE
	Duration      string         `gorm:"type:time" `
	Description   string          `gorm:"" `
	QualityType   string          `gorm:""`
	Season        int32           `gorm:"" ` // Required when type SERIES
	Episode       int32           `gorm:"" ` // Required when type SERIES
	DirectoryName string          `gorm:"unique" `
	Year          int32           `gorm:"not null" `
	Casters       json.RawMessage `json:"Casters" gorm:"type:json"`
	Directors     json.RawMessage `json:"Directors" gorm:"type:json"`
	CreatedAt     *time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     *time.Time      `gorm:"autoUpdateTime" `
	DeletedAt     gorm.DeletedAt  `gorm:"index"`
}

type MovieGroup struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" `
	NameLocal   string         `gorm:"not null;unique" `
	NameEng     string         `gorm:"not null;unique" `
	Type        string         `gorm:"not null" ` // MOVIE OR SERIES
	Status      string         `gorm:"not null" ` // ACTIVE OR INACTIVE
	Description string         `gorm:"" `
	PosterPath  string         `gorm:"not null" `
	CreatedAt   *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" `
}

// ----------------- type

type ParamsMovies struct {
	MovieGroupID int
	PageSize     int `json:"pageSize" binding:"required"`
	Current      int `json:"current" binding:"required"`
	Season       int `json:"season"`
}

type FilterBy string
const (
    Trending  FilterBy = "TRENDING"
    ComingSoon FilterBy = "COMING_SOON"
)
type ParamsMovieGroup struct {
	Name                  string `json:"Name"`
	Status                string `json:"Status"`
	PageSize              int    `json:"pageSize" binding:"required"`
	Current               *int    `json:"current" binding:"required"`
	InCludesMovieGroupIds string `json:"IncludesMovieGroupIds"`
	ExCludesMovieGroupIds string `json:"ExcludesMovieGroupIds"`
	FilterBy              FilterBy `json:"filter_by"`
}
type ResponseMovieGroup struct {
	ID          uint
	NameLocal   string
	NameEng     string
	Type        string
	Status      string
	PosterPath  string
	Description string
	MovieTime    string  `json:"movie_time" gorm:"not null;type:time"`
	MovieYear   int  `json:"movie_year" gorm:"not null"`
	MovieSeason int     `json:"movie_season"`
	MovieQuality string `json:"movie_quality"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type ResponseEpisodesOption struct {
	Season       int32 `json:"season"`
	EpisodeTotal int   `json:"episode_total"`
}




type ResponseMovieRecommendNow struct {
	MovieGroupID  int            `json:"movie_group_id"`
	NameLocal   string `json:"movie_name_local"`
	NameEng     string `json:"movie_name_eng"`
	Description     string `json:"description"`
	Type        string `json:"movie_type"`
	MovieQuality string `json:"movie_quality"`
	MovieTime    string  `json:"movie_time"`
	PosterURL string `json:"poster_url"`
	UpdatedAt   *time.Time `json:"updated_at"`
}