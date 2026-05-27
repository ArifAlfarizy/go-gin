package models

type Todo struct { // JSON tag				// GORM tag
	ID        uint   `json:"id"` // field named `ID` will be used as a primary field by default
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
