package main

import (
	"fmt"
	"net/http"

	"github.com/JMustang/go-crud/configs"
	"github.com/go-chi/chi"
)

func Main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
