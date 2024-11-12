package go_swagger_ui

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alexliesenfeld/go-swagger-ui/internal/config"
	"github.com/alexliesenfeld/go-swagger-ui/swagger-ui/templates"
)

//go:embed swagger-ui/dist/*
var swaggerUIFS embed.FS

//go:embed swagger-ui/templates/*
var templatesFS embed.FS

//var tplOverrides = map[string]*template.Template{
//	"index.html":             template.Must(template.ParseFS(templatesFS, "swagger-ui/templates/index.html")),
//	"swagger-initializer.js": template.Must(template.ParseFS(templatesFS, "swagger-ui/templates/swagger-initializer.js")),
//}

var tplGenerators = map[string]templates.Generator{
	"index.html":             templates.Index{},
	"swagger-initializer.js": templates.SwaggerInitializer{},
}

var allFilePaths = Must(walkFS("swagger-ui/dist/", &swaggerUIFS, "."))

func NewHandler(opts ...Option) http.HandlerFunc {
	cfg := config.UiConfig{
		TtmlTitle: "Swagger UI",
		Plugins:   map[config.Plugin]struct{}{},
		DocExpansion: config.Value[config.DocExpansion]{
			IsSet: true,
			Value: DocExpansionList,
		},
	}

	for idx := range opts {
		opts[idx](&cfg)
	}

	if len(cfg.Spec) > 0 {
		cfg.Spec = Must(yamlOrJSONToJSON(cfg.Spec))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fileName := strings.TrimPrefix(strings.TrimSpace(path.Base(r.URL.Path)), "/")
		if fileName == "" {
			fileName = "index.html"
		}

		// Always serve "index.html" if a file is being asked for does not exist.
		// These cases are usually caused by http.Handler instances that are mounted on URL paths
		// that do not end with a slash (e.g., https://example.com/hello, in which case the
		// file name would be "hello", although "index.html" is what is expected to be returned).
		if _, exists := allFilePaths[fileName]; !exists {
			fileName = "index.html"
		}

		// We reload the Spec file only for the CLI. In a normal production HTTP mode
		// "SpecFilePath" should be unset and not used at all. See WithSpecFilePath.
		if cfg.SpecFilePath != "" && fileName == "index.html" {
			newSpecContent, err := readSpecFile(cfg.SpecFilePath)
			if err != nil {
				slog.Error("error reading Swagger UI file", "err", err.Error())
				sendError(w, err)
				return
			}

			cfg.Spec = newSpecContent
		}

		// We either load the requested file from the embed filesystem directly or rendering
		// a template instead.
		var responseBody []byte
		if generator, ok := tplGenerators[fileName]; ok {
			var buf bytes.Buffer

			err := generator.Generate(&buf, cfg)
			if err != nil {
				slog.Error("failed to use Swagger UI template", "err", err.Error())
				sendError(w, err)
				return
			}
			responseBody = buf.Bytes()
		} else {
			var err error
			responseBody, err = fs.ReadFile(swaggerUIFS, "swagger-ui/dist/"+fileName)
			if err != nil {
				slog.Error("error reading file", "err", err.Error())
				sendError(w, err)
				return
			}
		}

		w.Header().Set("Content-Type", getContentType(fileName, responseBody))
		w.Write(responseBody)
	}
}

func sendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if errors.Is(err, fs.ErrNotExist) {
		w.WriteHeader(404)
	} else {
		w.WriteHeader(500)
	}
}

func readSpecFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading Spec file: %w", err)
	}

	return data, nil
}

func getContentType(fileName string, content []byte) string {
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		// read a chunk to decide between utf-8 text and binary
		var buf [512]byte
		n, _ := io.ReadFull(bytes.NewReader(content), buf[:])
		contentType = http.DetectContentType(buf[:n])
	}

	return contentType
}
