package repository

import (
	"github.com/Noblefel/Rest-Api-Managemen-Kontak/internal/models"
)

type AuthRepo interface {
	Register(u models.User) (int, error)
	Authenticate(u models.User) (int, int, error)
}

type UserRepo interface {
	GetUser(id int) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUser(u models.User) error
	DeleteUser(id int) error
}

type ContactRepo interface {
	GetContact(id int) (models.Contact, error)
	GetContactWithUser(id int) (models.Contact, error)
	GetAllContact() ([]models.Contact, error)
	GetUserContact(user_id int) ([]models.Contact, error)
	CreateContact(c models.Contact) (int, error)
	UpdateContact(c models.Contact) error
	DeleteContact(id int) error
}
