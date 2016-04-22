package main

import (
	"fmt"
	"log"
	"myapi"
	"net/http"
	_ "postgresql"
	_ "sync"
	_ "util"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	fmt.Println("server start")

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/users", myapi.GetAllUsers),
		rest.Post("/users", myapi.PostUser),
		rest.Get("/users/:userid/relationships", myapi.GetUserRelationships),
		rest.Put("/users/:userid/relationships/:other_user", myapi.PutUserRelationships),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9090", api.MakeHandler()))

	fmt.Println("server end")

}
