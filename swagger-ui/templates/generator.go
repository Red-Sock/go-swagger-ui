package templates

import (
	"io"

	"github.com/alexliesenfeld/go-swagger-ui/internal/config"
)

type Generator interface {
	Generate(writer io.Writer, config config.UiConfig) error
}
