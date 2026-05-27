package models

import "gorm.io/gorm"

/* gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
*/

type Todo struct { // JSON tag				// GORM tag
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
