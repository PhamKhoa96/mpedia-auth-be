package dsoauth

import (
	"log"

	"github.com/casbin/casbin/v2"
)

type AccessList struct {
	Enforcer *casbin.Enforcer
}

func NewAccessList(conf, csv string) *AccessList {
	ce, err := casbin.NewEnforcer(conf, csv)
	if err != nil {
		log.Fatalln("Error while loading access list")
	}

	return &AccessList{
		Enforcer: ce,
	}
}
