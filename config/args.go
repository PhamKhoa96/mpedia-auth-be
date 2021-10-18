package config

import (
	"flag"
)

type Args struct {
	ConfigFile     string
	Verbose        bool
	Debug          bool
	Newuser        string
	ChangePassword string
	DB             string
	NoAcl          bool
	AclModel       string
	AclPolicy      string
	RoleListFile   string
}

func LoadArgs() Args {
	configFile := flag.String("config", "auth.conf", "Path to config file")
	noAcl := flag.Bool("noacl", false, "No access list")
	aclModel := flag.String("aclmo", "./acl/rbac_model.conf", "Path to acl model file")
	aclPolicy := flag.String("aclpo", "./acl/rbac_policy.csv", "Path to acl policy file")
	roleListFile := flag.String("role", "./acl/role_list.txt", "Path to role list file")
	verbose := flag.Bool("verb", false, "Print all to stdout")
	debug := flag.Bool("debug", false, "Print all to stdout")
	newuser := flag.String("newuser", "", "Create new user")
	passwd := flag.String("passwd", "", "Create new user")
	db := flag.String("db", "", "Specify which database to use")
	flag.Parse()
	args := Args{
		ConfigFile:     *configFile,
		Verbose:        *verbose,
		Debug:          *debug,
		Newuser:        *newuser,
		ChangePassword: *passwd,
		DB:             *db,
		NoAcl:          *noAcl,
		AclModel:       *aclModel,
		AclPolicy:      *aclPolicy,
		RoleListFile:   *roleListFile,
	}
	return args
}
