package v1

import (
	"go-webapp-template/config"
)

func GetUsers() ([]config.User, error) {
	conf := config.Get()

	return conf.Users, nil
}
