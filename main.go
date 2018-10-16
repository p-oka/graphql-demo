package main

import (
	"net/http"
	"os"
	"github.com/p-oka/graphql-demo/server"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/graphql", server.NewHandler())
	mux.Handle("/", http.FileServer(Assets))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		panic(err)
	}
}
