package main

import (
	"fmt"
	"net/http"

	"github.com/JMustang/go-crud/configs"
	"github.com/JMustang/go-crud/handles"
	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handles.Create)
	r.Put("/{id}", handles.Update)
	r.Delete("/{id}", handles.Delete)
	r.Get("/", handles.List)
	r.Get("/{id}", handles.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
