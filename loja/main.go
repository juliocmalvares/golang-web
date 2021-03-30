package main

import (
	"net/http"

	"github.com/juliocmalvares/loja/routes"
)

func main() {
	routes.Router()
	http.ListenAndServe(":8000", nil)
}
