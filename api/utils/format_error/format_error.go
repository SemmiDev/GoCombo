package format_error

import (
	"github.com/SemmiDev/go-combo/api/errors_messages"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors_messages.ErrUsernameAlreadyTaken
	}

	if strings.Contains(err, "email") {
		return errors_messages.ErrEmailAlreadyTaken
	}

	if strings.Contains(err, "hashedPassword") {
		return errors_messages.ErrIncorrectPassword
	}

	return errors_messages.ErrIncorrectDetails
}
