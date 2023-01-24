package main

import (
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/frontend/server"
	"HomegrownDB/starter"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	switch {
	case len(os.Args) < 2:
		println("please enter command: install or start")
		os.Exit(-1)
	case os.Args[1] == "install":
		err := starter.InstallDefault()
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	case os.Args[1] == "start":
		db, err := hg.Load(nil)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		_ = db
		frontendServer := server.CreateDefaultServer(
			"0.0.0.0",
			"8080",
			nil,
		)
		err = frontendServer.Start()
		if err != nil {
			log.Printf("can start frontend frontendServer: %s", err.Error())
			os.Exit(1)
		}
	default:
		println("not supported command: " + os.Args[0])
		os.Exit(-1)
	}

	log.Print("Successfully started database")
}
