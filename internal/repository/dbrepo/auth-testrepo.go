package dbrepo

import (
	"errors"

	"github.com/Noblefel/Rest-Api-Managemen-Kontak/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (ar *testAuthRepo) Register(u models.User) (int, error) {
	if u.Email == "alreadyexists@error.com" {
		return 0, errors.New("duplicate key value")
	}

	if u.Email == "unexpected@error.com" {
		return 0, errors.New("unexpected error")
	}

	return 1, nil
}

func (ar *testAuthRepo) Authenticate(u models.User) (int, int, error) {
	if u.Password == "unexpected error" {
		return 0, 0, errors.New("unexpected error")
	}

	if u.Password == "jwt error" {
		return -1, 0, nil
	}

	correctPassword := "password"

	if u.Password != correctPassword {
		return 0, 0, bcrypt.ErrMismatchedHashAndPassword
	}

	return 1, 0, nil
}
