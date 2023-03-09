package main

import (
	"fmt"
	"log"
	"net/http"
	"platform/pkg/config"
	"platform/pkg/handlers"
)

const portNumber = ":8084"

func main() {
	var app config.AppConfig
	app.Viewsmap = make(map[string]*config.Video)

	repo := handlers.NewRepo(&app)
	handlers.Newhandlers(repo)
	//render.NewVideos(&app)

	//fmt.Println(fmt.Sprintf("starting at port %s", portNumber))
	fmt.Printf("starting at port %s \n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
