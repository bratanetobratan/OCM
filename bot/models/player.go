package models

type Player struct {
	FirstName   string
	SecondName  string
	Patronymic  string
	ID          int64
	PhoneNumber string
	Email       string
	TeamIDs     []int64
}
