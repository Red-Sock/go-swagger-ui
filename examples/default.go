package main

import (
	"fmt"
	"log"
	"net/http"

	swaggerui "github.com/alexliesenfeld/go-swagger-ui"
)

func main() {
	http.HandleFunc("/", swaggerui.NewHandler(
		swaggerui.WithHTMLTitle("My Example Petstore API"),
		swaggerui.WithSpecURLs("v1",
			[]swaggerui.SpecURL{
				{
					Name: "v1",
					URL:  "http://[::]/docs/swagger/hello_world_api.swagger.json",
				},
				{
					Name: "v2",
					URL:  "http://[::]/docs/swagger/hello_world_api_v2.swagger.json",
				},
			},
		),
		swaggerui.WithDocExpansion(swaggerui.DocExpansionFull),
	))

	fmt.Println("Starting Swagger UI on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
