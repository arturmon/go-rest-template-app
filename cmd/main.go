package main

import (
	_ "go-rest-template-app/cmd/docs"
	"go-rest-template-app/config"
	"go-rest-template-app/model"
	"go-rest-template-app/web"

	"log"
	"os"
	"os/signal"
)

//	@title			Rest Template App
//	@version		0.0.1
//	@description	This is a Rest Template App.
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes					http https
// @host						localhost:8080
// @BasePath					/v1
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to reading config file, %s\n", err)
	}

	service, err := model.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize model for operating all service, %s\n", err)
	}

	server := web.NewServer(cfg, service)
	go func() {
		if err = server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen for http server, %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log.Println("app is running")
	<-quit
	log.Println("app is stopped")
}
