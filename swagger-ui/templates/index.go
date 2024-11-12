package templates

import (
	_ "embed"
	"io"
	"text/template"

	"github.com/Red-Sock/go-swagger-ui/config"
)

type Index struct {
	BasePath  string
	HTMLTitle string
}

var (
	//go:embed index.html
	indexTemplatesFS string

	indexTemplate = template.Must(template.New("index.html").Parse(indexTemplatesFS))
)

func (i Index) Generate(w io.Writer, cfg config.UiConfig) error {
	args := Index{
		BasePath:  cfg.BasePath,
		HTMLTitle: cfg.TtmlTitle,
	}

	return indexTemplate.Execute(w, args)
}
