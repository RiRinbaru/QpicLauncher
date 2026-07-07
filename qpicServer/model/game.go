package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	CreatorID   string         `gorm:"type:varchar(255);not null" json:"creatorId"`
	Genre       string         `gorm:"type:varchar(100)" json:"genre"`
	Tags        string         `gorm:"type:text" json:"tags"`
	Explanation string         `gorm:"type:text" json:"explanation"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}