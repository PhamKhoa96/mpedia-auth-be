package main

import (
	"log"
	"os"

	"auth_be/config"
)

func main() {
	args := config.LoadArgs()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	conf := config.LoadConfigFromToml(args.ConfigFile)

	logger.Println("Start server")
	server := webserver.NewServer(conf, logger, nil, accesslist, auth, secretkey)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}
