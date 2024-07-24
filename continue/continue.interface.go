package continueplay

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type ContinuePlay struct {
	Id           uint       `json:"id" gorm:"primaryKey;autoIncrement" `
	ContinueUserID       uint       `json:"continue_user_id" gorm:"not null"`
	MovieGroupID uint       `json:"movie_group_id" gorm:"not null"`
	MovieID      uint       `json:"movie_id" gorm:"not null"`
	Season    float64 		`json:"season" `
	StampTime    string    `json:"stamp_time" gorm:"not null;type:time"`
		EndTime    string  `json:"end_time" gorm:"not null;type:time"`
	CreatedAt    *time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `json:"-" gorm:"autoUpdateTime" `
}

type ResponseContinuePlay struct {
	Id           uint       `json:"id" `
	MovieGroupID uint       `json:"movie_group_id"`
	MovieID      uint       `json:"movie_id"`
	StampTime    string    `json:"stamp_time"`
	EndTime    string 		`json:"end_time" `
	Season    float64 		`json:"season" `
}
