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
		swaggerui.WithSpecURLs("One", []swaggerui.SpecURL{
			{
				Name: "One",
				URL:  "https://petstore.swagger.io/v2/swagger.json",
			},
			{
				Name: "Two",
				URL:  "https://petstore.swagger.io/v2/swagger.json",
			},
		}),
		swaggerui.WithLayout(swaggerui.LayoutStandaloneLayout),
		swaggerui.WithDocExpansion(swaggerui.DocExpansionFull),
		swaggerui.WithPlugins(swaggerui.PluginTopBar),
	))

	fmt.Println("Starting Swagger UI on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
