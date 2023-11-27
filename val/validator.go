package val

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/fauzanfebrian/simplebank/util"
)

var (
	isValidUsername = regexp.MustCompile("^[a-z0-9_]+$").MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidataString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidataString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must containe only aphanumeric")
	}
	return nil
}

func ValidateRole(value string) error {
	if err := ValidataString(value, 3, 100); err != nil {
		return err
	}
	if value != util.BankerRole && value != util.DepositorRole {
		return fmt.Errorf("incorrect value")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidataString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must containe only letter & space")
	}
	return nil
}

func ValidatePassword(password string) error {
	return ValidataString(password, 6, 100)
}

func ValidateEmail(email string) error {
	if err := ValidataString(email, 3, 100); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("email invalid")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be positive integer")
	}
	return nil
}

func ValidateSecret(value string) error {
	return ValidataString(value, 32, 128)
}
