package go_swagger_ui

import (
	"github.com/alexliesenfeld/go-swagger-ui/internal/config"
)

// WithSpec sets the Spec field of https://github.com/swagger-api/swagger-ui/blob/HEAD/docs/usage/configuration.md..

// WithSpec sets an OpenAPI specification document content. When used, the URL configuration setting will not be used.
// This is useful for testing manually-generated definitions without hosting them.
func WithSpec(value []byte) Option {
	return func(cfg *config.UiConfig) {
		if isSpecSet(cfg) {
			panic("Spec already set")
		}

		cfg.Spec = value
	}
}

// WithSpecURL sets the URL pointing to API definition (normally swagger.json or swagger.yaml).
// Will be ignored if WithSpecURLs or WithSpec is used.
func WithSpecURL(value string) Option {
	return func(cfg *config.UiConfig) {
		if isSpecSet(cfg) {
			panic("Spec already set")
		}

		cfg.Url = config.Value[string]{
			IsSet: true,
			Value: value,
		}
	}
}

type SpecURL config.SpecURL

// WithSpecURLs sets the URLs array to multiple API definitions that are used by Topbar plugin.
// Panics if used along with WithSpec, WithSpecURL, WithSpecFilePath
func WithSpecURLs(primary string, urls []SpecURL) Option {
	return func(cfg *config.UiConfig) {
		if isSpecSet(cfg) {
			panic("Spec already set")
		}

		cfg.Plugins[TopBarPlugin] = struct{}{}

		cfg.Urls = make([]config.SpecURL, len(urls))
		for i, url := range urls {
			cfg.Urls[i] = config.SpecURL(url)
		}
		if len(primary) > 0 {
			cfg.UrlsPrimary = config.Value[string]{
				IsSet: true,
				Value: primary,
			}
		}
	}
}

// WithSpecFilePath sets a file path to read from the OS file system.
// THIS OPTION IS NOT RECOMMENDED FOR PRODUCTION USE, because it reloads the file on every request.
// This option only exist to for testing purposes. Once file content is read, it will be used to set the Spec field of
// https://github.com/swagger-api/swagger-ui/blob/HEAD/docs/usage/configuration.md and is equivalent to the
// WithSpec function.
func WithSpecFilePath(path string) Option {
	return func(cfg *config.UiConfig) {
		if isSpecSet(cfg) {
			panic("Spec already set")
		}

		cfg.SpecFilePath = path
	}
}
