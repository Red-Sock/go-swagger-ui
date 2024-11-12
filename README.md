<img src="./assets/logo.png" align="right" height="200"/>

# go-swagger-ui

`go-swagger-ui` provides a simple way to serve Swagger UI in your Go web application. 
It embeds Swagger UI and provides a customizable [http.Handler](https://pkg.go.dev/net/http#Handler) 
to serve it.

## Features

* Provides a customizable HTTP Handler to serve Swagger UI.
* Supports many of Swagger UI's configuration options.
* Supports dynamic UI configuration in your Go application.
* Provides a CLI application to open OpenAPI specification files in a Swagger UI instance (browser window).

## Installation

To install the `go-swagger-ui` package, use the following command:

```bash
go get github.com/Red-Sock/go-swagger-ui
```

## Usage
```go
package main

import (
	"fmt"
	swaggerui "github.com/Red-Sock/go-swagger-ui"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", swaggerui.NewHandler(
		swaggerui.WithHTMLTitle("My Example Petstore API"),
		swaggerui.WithSpecURL("https://petstore.swagger.io/v2/swagger.json"),
		swaggerui.WithTryItOutEnabled(true),
		swaggerui.WithPersistAuthorization(true),
	))

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

## CLI Usage

Install the CLI application:

```bash
go install github.com/Red-Sock/go-swagger-ui/cmd/swui@latest
```

You can then use the CLI tool to serve OpenAPI spec files in a separate Swagger UI instance in a browser window.
```bash
swui /path/to/openapi-spec.yaml
```

`swui` also allows you to reload the browser window to see changes made to the spec file.

## Roadmap

- [x] Embed Swagger UI
- [x] Provide simple but powerful `http.Handler` that is compatible with major web libraries
- [x] Allow to change the majority of configuration parameters from within a Go application
- [x] Make it possible to configure multiple spec file urls
- [x] Provide a CLI tool to view OpenAPI spec files locally in a browser
- [ ] Add OAuth2 configuration possibilities (https://github.com/swagger-api/swagger-ui/blob/master/docs/usage/oauth2.md)
- [ ] Make plugins configurable
- [ ] Make presets configurable
- [ ] Allow using CDN instead of embedding Swagger UI

## License

`go-swagger-ui` is free software: you can use it under the terms of the 
[Apache License 2.0](LICENSE) license.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; 
without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. 

This library includes Swagger UI as static resources. Swagger UI is a product of 
SmartBear Software Inc. and is made available under the Apache License 2.0. 
For complete details of the license, please refer to the official license text available at 
[this link](https://github.com/swagger-api/swagger-ui/blob/master/LICENSE). 
It's important to note that as per the requirements of the Apache-2.0 license, this documentation 
serves as an attribution to SmartBear Software Inc., acknowledging their development of Swagger UI. 
Users of this library are advised to ensure they are in compliance with the Apache-2.0 license terms 
when utilizing Swagger UI as part of this library.
