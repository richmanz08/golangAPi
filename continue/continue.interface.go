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
	Season    int 		`json:"season" `
	StampTime    string    `json:"stamp_time" gorm:"not null;type:time"`
		EndTime    string  `json:"end_time" gorm:"not null;type:time"`
	CreatedAt    *time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `json:"-" gorm:"autoUpdateTime" `
}

type ContinuePlayParams struct {
	MovieGroupID uint       `json:"movie_group_id" binding:"required"`
	MovieID      uint       `json:"movie_id" binding:"required"`
	Season    int 		`json:"season"`
	StampTime    string    `json:"stamp_time" binding:"required"`
		EndTime    string  `json:"end_time" binding:"required"`
}

type ResponseContinuePlay struct {
	Id           uint       `json:"id" `
	MovieGroupID uint       `json:"movie_group_id"`
	MovieID      uint       `json:"movie_id"`
	StampTime    string    `json:"stamp_time"`
	EndTime    string 		`json:"end_time" `
	Season    int 		`json:"season" `
}
