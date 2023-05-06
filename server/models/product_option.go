package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductOption struct {
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
  Name string `json:"name"`
  Price float64 `json:"price"`
  ProductID uuid.UUID `json:"product_id"`
}

func (po *ProductOption) BeforeCreate(tx *gorm.DB) (err error) {
    po.ID = uuid.New()
    return
}

