package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
  Username string `gorm:"unique" json:"username"`
  Password string `json:"password"`
  Email string `json:"email"`
  StudentId string `json:"student_id"`
  FirstName string `json:"f_name"`
  LastName string `json:"l_name"`
}
