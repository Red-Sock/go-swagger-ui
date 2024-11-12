package templates

import (
	"io"

	"github.com/Red-Sock/go-swagger-ui/config"
)

type Generator interface {
	Generate(writer io.Writer, config config.UiConfig) error
}
