package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
  Id string `gorm:"primaryKey" json:"id"` 
  Username string `gorm:"unique" json:"username"`
  Password string `json:"password"`
  Email string `json:"email"`
  StudentId string `json:"student_id"`
  FirstName string `json:"f_name"`
  LastName string `json:"l_name"`
}
