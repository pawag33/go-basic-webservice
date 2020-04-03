package main

import (
	"net/http"

	"github.com/pawag33/go-basic-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
