package main

import (
	"github.com/bariasabda/monorepo/packages/fetch/cmd/http"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	http.Execute()
}
