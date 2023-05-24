package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type ProductOption struct {
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primaryKey"`
  Name string 
  Price float64 
  ProductID uuid.UUID `gorm:"type:char(36);"`
}

func (po *ProductOption) BeforeCreate(tx *gorm.DB) (err error) {
    po.ID = uuid.New()
    return
}


