// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package users

import (
	"errors"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/platrum"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserPassword(user *model.User, password string) error {
	if strings.Contains(user.Roles, model.SystemAdminRoleId) || user.IsBot {
		if err := ComparePassword(user.Password, password); err != nil {
			return NewErrInvalidPassword("")
		}

		return nil
	}

	if res := platrum.CheckUserPassword(user.Email, password); !res {
		return NewErrInvalidPassword("")
	}

	return nil
}

// HashPassword generates a hash using the bcrypt.GenerateFromPassword
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func ComparePassword(hash string, password string) error {
	if password == "" || hash == "" {
		return errors.New("empty password or hash")
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (us *UserService) isPasswordValid(password string) error {

	if *us.config().ServiceSettings.EnableDeveloper {
		return nil
	}

	return IsPasswordValidWithSettings(password, &us.config().PasswordSettings)
}

// IsPasswordValidWithSettings is a utility functions that checks if the given password
// conforms to the password settings. It returns the error id as error value.
func IsPasswordValidWithSettings(password string, settings *model.PasswordSettings) error {
	id := "model.user.is_valid.pwd"
	isError := false

	if len(password) < *settings.MinimumLength || len(password) > model.PasswordMaximumLength {
		isError = true
	}

	if *settings.Lowercase {
		if !strings.ContainsAny(password, model.LowercaseLetters) {
			isError = true
		}

		id = id + "_lowercase"
	}

	if *settings.Uppercase {
		if !strings.ContainsAny(password, model.UppercaseLetters) {
			isError = true
		}

		id = id + "_uppercase"
	}

	if *settings.Number {
		if !strings.ContainsAny(password, model.NUMBERS) {
			isError = true
		}

		id = id + "_number"
	}

	if *settings.Symbol {
		if !strings.ContainsAny(password, model.SYMBOLS) {
			isError = true
		}

		id = id + "_symbol"
	}

	if isError {
		return NewErrInvalidPassword(id + ".app_error")
	}

	return nil
}
