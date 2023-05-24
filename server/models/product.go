package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primaryKey" `
  Name string 
  Price float64 
  Description string 
  Options []ProductOption `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
    p.ID = uuid.New()
    return
}


