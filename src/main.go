package main

import (
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/frontend/handler"
	"HomegrownDB/frontend/server"
	"HomegrownDB/starter"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	switch {
	case len(os.Args) < 2:
		println("please enter command: install or start")
		os.Exit(-1)
	case os.Args[1] == "install":
		install()
	case os.Args[1] == "uninstall":
		uninstall()
	case os.Args[1] == "start":
		start()
	default:
		println("not supported command: " + os.Args[0])
		os.Exit(-1)
	}
}

func start() {
	time0 := time.Now()

	db, err := hg.Load(nil)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	setUpShutdownFunction(db)

	frontendServer := server.CreateDefaultServer(
		"0.0.0.0",
		"8080",
		handler.Handlers{
			SqlHandler: handler.NewSqlHandler(db.FrontendContainer()),
		},
	)

	err = frontendServer.Start()
	if err != nil {
		log.Printf("can start frontend frontendServer: %s", err.Error())
		os.Exit(1)
	}

	startTime := time.Now().Sub(time0).Milliseconds()
	startTimeInSec := startTime / 1_000
	log.Printf("Successfully started database in %d.%03d ms", startTimeInSec, startTime-startTimeInSec*1_000)
	time.Sleep(time.Duration(1<<63 - 1))
}

func setUpShutdownFunction(db hg.DB) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		err := db.Shutdown()
		if err != nil {
			log.Printf("Could not shutdown database. Reason:\n%s", err.Error())
		}
	}()
}

func install() {
	err := starter.InstallDefault()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func uninstall() {
	err := starter.UninstallDefault()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
