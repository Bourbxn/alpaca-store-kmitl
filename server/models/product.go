package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
  Name string `json:"name"`
  Price float64 `json:"price"`
  Description string `json:"description"`
  Options []ProductOption `json:"options,omitempty" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
    p.ID = uuid.New()
    return
}


