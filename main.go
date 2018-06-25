package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ant0ine/go-json-rest/rest"
	"fmt"
	"net/http"
	"log"
	"github.com/xpp666/go.orm/service"
)


func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router,err :=rest.MakeRouter(
		rest.Post("/register",service.Register),
		rest.Get("/query/:id", service.QueryUserById),
		rest.Delete("/delete/:id", service.DeleteUser),
		rest.Put("/update",service.Update),
	)
	if err != nil {
		fmt.Println("err=>", err)
	}
	api.SetApp(router)
	http.Handle("/user/",http.StripPrefix("/user", api.MakeHandler()))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
