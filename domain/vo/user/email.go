package user

import (
	"errors"
	"regexp"
)

type Email string

// NewEmail Because it is valueObject, it is returned as a value type.
func NewEmail(email string) (Email, error) {
	if !isValidEmail(email) {
		return "", errors.New("invalid Email")
	}

	return Email(email), nil
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return regex.MatchString(email)
}

func (e Email) Value() string {
	return string(e)
}
