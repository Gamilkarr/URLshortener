package main

import (
	"net/http"

	"github.com/Gamilkarr/URLshortener/config"
	"github.com/Gamilkarr/URLshortener/internal/app/handlers"
	"github.com/Gamilkarr/URLshortener/internal/app/router"
)

func main() {
	cfg := config.New()
	h := handlers.New()
	r := router.New(h)

	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		panic(err)
	}
}
