package main

import (
	"log"

	"github.com/devasherr/socket-rooms/db"
	"github.com/devasherr/socket-rooms/internal/user"
	"github.com/devasherr/socket-rooms/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize db connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:7777")
}
