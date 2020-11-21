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

	if os.Getenv("SSL_MODE") == "self-signed" {
		log.Println(os.Getenv("SSL_SELF_SIGNED_PATH") + "/" + os.Getenv("SSL_SELF_SIGNED_NAME") + ".crt")
		err = http.ListenAndServeTLS(
			fmt.Sprintf(":%s", svcPort),
			os.Getenv("SSL_SELF_SIGNED_PATH") + "/" + os.Getenv("SSL_SELF_SIGNED_NAME") + ".crt",
			os.Getenv("SSL_SELF_SIGNED_PATH") + "/" + os.Getenv("SSL_SELF_SIGNED_NAME") + ".key",
			nil,
		)
	} else {
		err = http.ListenAndServe(
			fmt.Sprintf(":%s", svcPort),
			router.BIQRouter,
		)
	}

	if err != nil {
		log.Println(err)
	}

	Exit()

}
