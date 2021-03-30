package routes

import (
	"net/http"

	"github.com/juliocmalvares/loja/controllers"
)

func Router() {
	http.HandleFunc("/", controllers.Index)
}
