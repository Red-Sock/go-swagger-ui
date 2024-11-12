package swaggers

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var swaggers embed.FS

//go:embed dist/hello_world_api.swagger.json
var SwaggerV1 []byte

func HostSwaggers() http.HandlerFunc {
	stripped, err := fs.Sub(swaggers, "dist")
	if err != nil {
		log.Fatal(err)
	}

	ffs := http.FileServer(http.FS(stripped))

	return ffs.ServeHTTP
}
