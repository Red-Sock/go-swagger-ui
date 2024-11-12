package main

import (
	"fmt"
	"log"
	"net/http"

	swaggerui "github.com/alexliesenfeld/go-swagger-ui"
	"github.com/alexliesenfeld/go-swagger-ui/examples/helpers/swaggers"
)

func main() {

	port := ":80"

	mux := http.NewServeMux()

	const swaggerPath = "/swagger/"

	mux.HandleFunc("/", swaggerui.NewHandler(
		swaggerui.WithHTMLTitle("Bytes swaggers"),
		swaggerui.WithSpec(swaggers.SwaggerV1),
	))

	mux.Handle(swaggerPath, http.StripPrefix(swaggerPath, swaggers.HostSwaggers()))

	fmt.Println("Starting Swagger UI on http://localhost" + port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
