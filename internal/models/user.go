package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
}
