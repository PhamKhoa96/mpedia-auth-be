package webserver

import (
	"log"
	"net/smtp"

	"auth_be/config"
)

type Webserver struct {
	debug             bool
	config            *config.Config
	logger            *log.Logger
	e                 *echo.Echo
	secret            []byte
	acl               *dsoauth.AccessList
	auth              *dsoauth.DsoAuth
	smtp              *smtp.Auth
	keysUpdateTimeOut int
}

func NewServer(config *config.Config, logger *log.Logger, smtp *smtp.Auth, accesslist *dsoauth.AccessList, auth *dsoauth.DsoAuth, secret string) Webserver {

	s := Webserver{
		debug:  config.Web.Debug,
		logger: logger,
		config: config,
		smtp:   smtp,
		acl:    accesslist,
		auth:   auth,
		secret: []byte(secret),
	}

	s.e = s.loadRouter()

	go s.backgroundUpdateKeys()

	return s
}
