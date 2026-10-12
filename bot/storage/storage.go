package storage

import "OCM/bot/models"

type Storage interface {
	Save(p *models.Player, platform models.Platform, externalID string) (ID int64, err error)
	Remove(id int64) error
	Update(p *models.Player) error
	Find(platform models.Platform, externalID string) (*models.Player, error)
	IsPhoneExist(phoneNumber string) (bool, error)
	IsEmailExist(email string) (bool, error)
	Join(teamID int64, playerID int64) error
	Request(teamID int64, playerID int64) (ID int64, err error)
	Reply(requestID int64, approved bool) (err error)
}
