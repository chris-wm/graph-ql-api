package entity

import (
	"time"

	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("entity")
}

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type BaseEntity struct {
	ID        string         `gorm:"type:string;size:16;primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (e *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = hash.GenerateHash(12)

	return
}

type UuidEntity struct {
	ID        uuid.UUID      `gorm:"type:char(36);primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (e *UuidEntity) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID, err = uuid.NewRandom()

	if err != nil {
		return err
	}

	return
}
