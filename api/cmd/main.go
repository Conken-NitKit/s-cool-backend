package main

import (
	"condog/internal/config"
	"condog/internal/di"
	"log"

	"github.com/srvc/appctx"
)

func main() {
	log.Println("start application")
	ctx := appctx.Global()
	app, err := di.InitializeApp(ctx, config.GetAppConfig())
	if err != nil {
		log.Fatalln("Failed to dependency injection:", err)
	}
	log.Fatalln(app.Run())
}