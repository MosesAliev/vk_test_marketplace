package model

import (
	"errors"
	"strings"
	"unicode"
)

type User struct {
	Login    string `gorm:"primaryKey"`
	Password string
	Ads      []Ad `gorm:"foreignKey:UserLogin"`
}

func (user User) hasDigits() bool {
	for _, r := range user.Password {
		if unicode.IsDigit(r) {
			return true
		}
	}

	return false
}

func (user User) hasUppers() bool {
	for _, r := range user.Password {
		if unicode.IsUpper(r) {
			return true
		}
	}

	return false
}

func (user User) hasLowers() bool {
	for _, r := range user.Password {
		if unicode.IsLower(r) {
			return true
		}
	}

	return false
}

func (user User) IsValidPassword() error {
	if len(user.Password) < 8 {
		return errors.New("password length must be more than 8")
	}

	if !user.hasDigits() {
		return errors.New("the password must contain numbers")
	}

	if !strings.ContainsAny(user.Password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return errors.New("the password must contain letters")
	}

	if !user.hasUppers() {
		return errors.New("password must contain uppercase characters")
	}

	if !user.hasLowers() {
		return errors.New("password must contain lowercase characters")
	}

	return nil
}
