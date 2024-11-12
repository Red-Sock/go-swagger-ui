package main

import (
	"fmt"
	"log"
	"net/http"

	swaggerui "github.com/Red-Sock/go-swagger-ui"
	"github.com/Red-Sock/go-swagger-ui/examples/helpers/swaggers"
)

func main() {

	port := ":80"

	mux := http.NewServeMux()

	const swaggerPath = "/swagger/"

	mux.HandleFunc("/", swaggerui.NewHandler(
		swaggerui.WithHTMLTitle("One swaggers by url"),
		swaggerui.WithSpecURL("http://localhost"+port+swaggerPath+"hello_world_api.swagger.json"),
	))

	mux.Handle(swaggerPath, http.StripPrefix(swaggerPath, swaggers.HostSwaggers()))

	fmt.Println("Starting Swagger UI on http://localhost" + port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
