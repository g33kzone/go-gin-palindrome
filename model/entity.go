package model

// Note - struct symbolises the Message DB entity
type Note struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Message      string `json:"message"`
	IsPalindrome bool   `json:"ispalindrome"`
}
