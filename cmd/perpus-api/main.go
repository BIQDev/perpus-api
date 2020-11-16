package main

import (
	"fmt"
	"github.com/BIQDev/perpus-api/internal/db"
	_ "github.com/BIQDev/perpus-api/internal/modules"
	"github.com/BIQDev/perpus-api/internal/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Init() {
	db.Mongo.Init()
}

func Exit() {
	db.Mongo.Conn().Disconnect(db.Mongo.GetCtx())
}

func main() {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("GO_ENV") == "development" || os.Getenv("BIQ_IS_DEBUG") == "true" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	Init()


	svcPort := os.Getenv("SERVICE_PORT")
	log.Println("Listening at port:", svcPort)
	err = http.ListenAndServe(
		fmt.Sprintf(":%s", svcPort),
		router.BIQRouter,
	)

	Exit()

}
