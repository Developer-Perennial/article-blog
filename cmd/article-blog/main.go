package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/DevPer/article-blog/app"
	"github.com/DevPer/article-blog/internal/repository/db"
)

var (
	configFilePath = flag.String("c", "deployment/config.yaml", "Server Config File")
	autoMigrate    = flag.Bool("m", false, "Auto-migrate DB schema")
)

func main() {
	flag.Parse()

	a, err := app.Init(*configFilePath)
	if err != nil {
		panic(err)
	}

	// migrate entities to DB
	if *autoMigrate {
		err = db.AutoMigrate(a.DbConfig, a.Ds)
		if err != nil {
			panic(err)
		}
	}

	// create channel to gracefully stop server
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	srvCtx := context.Background()

	// start server in separate goroutine to setup graceful server termination
	go func() {
		if err := a.Run(srvCtx); err != nil {
			log.Fatalln(err)
		}
	}()

	// wait for termination signal
	<-sc

	log.Println("Closing Server")
	err = a.ShutDown(srvCtx)
	if err != nil {
		panic(err)
	}
	log.Println("Server Closed!!")
}
