package movies

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB
type Movie struct {
	ID            uint          `gorm:"primaryKey;autoIncrement" `
	MovieGroupID  uint          `gorm:"not null" `
	NameLocal     string         `gorm:"unique"`
	NameEng       string         `gorm:"unique" `
	Type          string         `gorm:"not null" ` // MOVIE OR SERIES
	Status     string         	`gorm:"not null" ` // ACTIVE OR INACTIVE
	Duration      float64        `gorm:"" `
	Description   string         `gorm:"" `
	QualityType   string         `gorm:""`
	DirectorsID   string         `gorm:"" `
	CastersID     string         `gorm:"" `
	Season        int32          `gorm:"" ` // Required when type SERIES
	Episode       int32          `gorm:"" ` // Required when type SERIES
	DirectoryName string         `gorm:"unique" `
	Year          int32          `gorm:"not null" `
	CreatedAt     *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt     *time.Time     `gorm:"autoUpdateTime" `
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type MovieGroup struct {
	ID         uint          `gorm:"primaryKey;autoIncrement" `
	NameLocal  string         `gorm:"not null;unique" `
	NameEng    string         `gorm:"not null;unique" `
	Type       string         `gorm:"not null" ` // MOVIE OR SERIES
	Status     string         `gorm:"not null" ` // ACTIVE OR INACTIVE
	PosterPath string         `gorm:"not null" `
	CreatedAt  *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index" `
}
