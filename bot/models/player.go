package models

type Player struct {
	FirstName   string
	LastName    string
	Patronymic  *string
	ID          int64
	PhoneNumber *string
	Email       *string
}
