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
		swaggerui.WithHTMLTitle("Multiple swaggers"),
		swaggerui.WithSpecURLs("HelloWorld_v1",
			[]swaggerui.SpecURL{
				{
					Name: "HelloWorld_v1",
					URL:  "http://localhost" + port + swaggerPath + "hello_world_api.swagger.json",
				},
				{
					Name: "HelloWorld_v2",
					URL:  "http://localhost" + port + swaggerPath + "/hello_world_api_v2.swagger.json",
				},
			},
		),
	))

	mux.Handle(swaggerPath, http.StripPrefix(swaggerPath, swaggers.HostSwaggers()))

	fmt.Println("Starting Swagger UI on http://localhost" + port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
