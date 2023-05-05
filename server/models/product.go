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
  Options []ProductOption `json:"options,omitempty" gorm:"foreignKey:ProductID"`
}

type ProductOption struct {
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
  Name string `json:"name"`
  Price float64 `json:"price"`
  ProductID uuid.UUID `json:"product_id"`
}

