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
	StampTime    float64    `json:"stamp_time" gorm:"not null" `
	CreatedAt    *time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `json:"-" gorm:"autoUpdateTime" `
}
