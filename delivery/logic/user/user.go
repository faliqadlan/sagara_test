package user

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type Logic struct{}

func New() *Logic {
	return &Logic{}
}

func (l *Logic) ValidationStruct(req Req) error {
	var v = validator.New()
	if err := v.Struct(req); err != nil {
		log.Warn(err)
		switch {
		case strings.Contains(err.Error(), "UserName"):
			err = errors.New("invalid userName")
		case strings.Contains(err.Error(), "Email"):
			err = errors.New("invalid email")
		case strings.Contains(err.Error(), "Password"):
			err = errors.New("invalid password")
		case strings.Contains(err.Error(), "Name"):
			err = errors.New("invalid name")
		default:
			err = errors.New("invalid input")
		}
		return err
	}
	return nil
}
