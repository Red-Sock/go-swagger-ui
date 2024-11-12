package templates

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/Red-Sock/go-swagger-ui/config"
)

var (
	//go:embed swagger-initializer.js.pattern
	swaggerInitializerTemplatesFS string

	swaggerInitializerTemplate = template.Must(template.New("swagger-initializer.js").Parse(swaggerInitializerTemplatesFS))
)

type SwaggerInitializer struct {
	Plugins    []string
	ConfigURL  string
	Spec       string
	URL        string
	URLs       string
	PrimaryURL string

	DocExpansion             string
	DefaultModelExpandDepth  string
	DefaultModelsExpandDepth string
	DefaultModelRendering    string
	QueryConfigEnabled       string
	SupportedSubmitMethods   string
	DeepLinking              string
	ShowMutatedRequest       string
	ShowExtensions           string
	ShowCommonExtensions     string
	Filter                   string
	FilterString             string
	DisplayOperationId       string
	TryItOutEnabled          string
	DisplayRequestDuration   string
	PersistAuthorization     string
	WithCredentials          string
	OAuth2RedirectUrl        string
	Layout                   string
	ValidatorURL             string
	MaxDisplayedTags         string
}

func (SwaggerInitializer) Generate(w io.Writer, cfg config.UiConfig) error {
	urlsAsBase64EncodedJSON, err := marshalObject(cfg.Urls)
	if err != nil {
		return fmt.Errorf("cannot marshal URLs: %w", err)
	}

	args := SwaggerInitializer{
		ConfigURL:                fromStringConfigValue(cfg.ConfigURL),
		Spec:                     strings.TrimSpace(base64.StdEncoding.EncodeToString(cfg.Spec)),
		URL:                      fromStringConfigValue(cfg.Url),
		DocExpansion:             fromDocExpansionConfigValue(cfg.DocExpansion),
		DefaultModelExpandDepth:  fromIntConfigValue(cfg.DefaultModelExpandDepth),
		DefaultModelsExpandDepth: fromIntConfigValue(cfg.DefaultModelsExpandDepth),
		DefaultModelRendering:    fromModelRenderingConfigValue(cfg.DefaultModelRendering),
		QueryConfigEnabled:       fromBoolConfigValue(cfg.QueryConfigEnabled),
		SupportedSubmitMethods:   strings.TrimSpace(strings.Join(cfg.SupportedSubmitMethods, ",")),
		DeepLinking:              fromBoolConfigValue(cfg.DeepLinking),
		ShowMutatedRequest:       fromBoolConfigValue(cfg.ShowMutatedRequest),
		ShowExtensions:           fromBoolConfigValue(cfg.ShowExtensions),
		ShowCommonExtensions:     fromBoolConfigValue(cfg.ShowCommonExtensions),
		Filter:                   fromBoolConfigValue(cfg.Filter),
		FilterString:             fromStringConfigValue(cfg.FilterString),
		DisplayOperationId:       fromBoolConfigValue(cfg.DisplayOperationID),
		TryItOutEnabled:          fromBoolConfigValue(cfg.TryItOutEnabled),
		DisplayRequestDuration:   fromBoolConfigValue(cfg.DisplayRequestDuration),
		PersistAuthorization:     fromBoolConfigValue(cfg.PersistAuthorization),
		WithCredentials:          fromBoolConfigValue(cfg.WithCredentials),
		OAuth2RedirectUrl:        fromStringConfigValue(cfg.Oauth2RedirectUrl),
		Layout:                   fromStringConfigValue(cfg.Layout),
		ValidatorURL:             fromStringConfigValue(cfg.ValidatorUrl),
		MaxDisplayedTags:         fromIntConfigValue(cfg.MaxDisplayedTags),
		PrimaryURL:               fromStringConfigValue(cfg.UrlsPrimary),
		URLs:                     urlsAsBase64EncodedJSON,
		Plugins:                  fromMapToArray(cfg.Plugins),
	}

	return swaggerInitializerTemplate.Execute(w, args)

}
