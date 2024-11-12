package go_swagger_ui

import (
	"strings"

	"github.com/alexliesenfeld/go-swagger-ui/internal/config"
)

// Option is a function that takes a pointer to config.UiConfig and modifies it.
type Option func(*config.UiConfig)

var (
	DocExpansionList config.DocExpansion = "list"
	DocExpansionFull config.DocExpansion = "full"
	DocExpansionNone config.DocExpansion = "none"
)

// WithDocExpansion controls the default expansion setting for the operations and tags.
func WithDocExpansion(value config.DocExpansion) Option {
	return func(cfg *config.UiConfig) {
		cfg.DocExpansion = config.Value[config.DocExpansion]{
			IsSet: true,
			Value: value,
		}
	}
}

// WithDefaultModelExpandDepth sets the default expansion depth for the model on the model-example section.
func WithDefaultModelExpandDepth(defaultModelExpandDepth int) Option {
	return func(cfg *config.UiConfig) {
		cfg.DefaultModelExpandDepth = config.Value[int]{Value: defaultModelExpandDepth, IsSet: true}
	}
}

// WithDefaultModelsExpandDepth sets the default expansion depth for models
// (set to -1 completely hide the models).
func WithDefaultModelsExpandDepth(defaultModelsExpandDepth int) Option {
	return func(cfg *config.UiConfig) {
		cfg.DefaultModelsExpandDepth = config.Value[int]{Value: defaultModelsExpandDepth, IsSet: true}
	}
}

var (
	ModelRenderingExample config.ModelRendering = "example"
	ModelRenderingModel   config.ModelRendering = "model"
)

// WithDefaultModelRendering controls how the model is shown when the API is first rendered.
// The user can always switch the rendering for a given model by clicking the 'Model' and 'Example Value' links.
func WithDefaultModelRendering(defaultModelRendering config.ModelRendering) Option {
	return func(cfg *config.UiConfig) {
		cfg.DefaultModelRendering = config.Value[config.ModelRendering]{Value: defaultModelRendering, IsSet: true}
	}
}

// WithQueryConfigEnabled enables overriding configuration parameters via URL search params.
func WithQueryConfigEnabled(queryConfigEnabled bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.QueryConfigEnabled = config.Value[bool]{Value: queryConfigEnabled, IsSet: true}
	}
}

// WithSupportedSubmitMethods sets a list of HTTP methods that have the "Try it out" feature enabled.
// An empty array disables "Try it out" for all operations. This does not filter the operations from the display.
// Default is: ["get", "put", "post", "delete", "options", "head", "patch", "trace"].
func WithSupportedSubmitMethods(supportedSubmitMethods ...string) Option {
	return func(cfg *config.UiConfig) {
		cfg.SupportedSubmitMethods = append(cfg.SupportedSubmitMethods, supportedSubmitMethods...)
	}
}

// WithDeepLinking enables deep linking. See documentation at
// https://swagger.io/docs/open-source-tools/swagger-ui/usage/deep-linking/
// for more information.
func WithDeepLinking(deepLinking bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.DeepLinking = config.Value[bool]{Value: deepLinking, IsSet: true}
	}
}

// WithShowExtensions controls the display of vendor extension (x-) fields and values for
// Operations, Parameters, Responses, and Schema.
func WithShowExtensions(showExtensions bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.ShowExtensions = config.Value[bool]{Value: showExtensions, IsSet: true}
	}
}

// WithShowCommonExtensions controls the display of extensions (pattern, maxLength, minLength, maximum, minimum)
// fields and values for Parameters.
func WithShowCommonExtensions(showCommonExtensions bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.ShowCommonExtensions = config.Value[bool]{Value: showCommonExtensions, IsSet: true}
	}
}

// WithFilter enables filtering. The top bar will show an edit box that you can use to filter the tagged
// operations that are shown. If enabled and a non-empty expression string is passed, then filtering
// will be enabled using that string as the filter expression. Filtering is case-sensitive matching
// the filter expression anywhere inside the tag. Leave the expression empty, if you only want to
// enable filtering but do not need a filter expression.
func WithFilter(enabled bool, expression string) Option {
	return func(cfg *config.UiConfig) {
		cfg.Filter = config.Value[bool]{Value: enabled, IsSet: true}
		if enabled && len(expression) > 0 {
			cfg.FilterString = config.Value[string]{Value: expression, IsSet: true}
		}
	}
}

// WithDisplayOperation controls the display of operationId in operations list. The default is false.
func WithDisplayOperation(displayOperationID bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.DisplayOperationID = config.Value[bool]{Value: displayOperationID, IsSet: true}

	}
}

// WithTryItOutEnabled controls whether the "Try it out" section should be enabled by default.
func WithTryItOutEnabled(tryItOutEnabled bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.TryItOutEnabled = config.Value[bool]{Value: tryItOutEnabled, IsSet: true}

	}
}

// WithDisplayRequestDuration controls the display of the request duration (in milliseconds) for "Try it out" requests.
func WithDisplayRequestDuration(displayRequestDuration bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.DisplayRequestDuration = config.Value[bool]{Value: displayRequestDuration, IsSet: true}
	}
}

// WithPersistAuthorization configures Swagger UI to persist authorization data, so that it is not lost
// on browser close/refresh.
func WithPersistAuthorization(persistAuthorization bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.PersistAuthorization = config.Value[bool]{Value: persistAuthorization, IsSet: true}
	}
}

// WithCredentials enables passing credentials, as defined in the Fetch standard, in CORS requests that are
// sent by the browser. Note that Swagger UI cannot currently set cookies cross-domain (see swagger-js#1163) -
// as a result, you will have to rely on browser-supplied cookies (which this setting enables sending)
// that Swagger UI cannot control.
func WithCredentials(withCredentials bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.WithCredentials = config.Value[bool]{Value: withCredentials, IsSet: true}
	}
}

// WithOauth2RedirectUrl sets the OAuth redirect URL.
func WithOauth2RedirectUrl(oauth2RedirectUrl string) Option {
	return func(cfg *config.UiConfig) {
		cfg.Oauth2RedirectUrl = config.Value[string]{Value: oauth2RedirectUrl, IsSet: true}
	}
}

// WithHTMLTitle sets the index HTML page TtmlTitle.
func WithHTMLTitle(title string) Option {
	return func(cfg *config.UiConfig) {
		cfg.TtmlTitle = title
	}
}

var (
	LayoutBaseLayout       config.Layout = "BaseLayout"
	LayoutStandaloneLayout config.Layout = "StandaloneLayout"
)

// WithLayout sets the name of a component available via the plugin system to use as the top-level
// layout for Swagger UI.
// Possible values are "BaseLayout" and "StandaloneLayout".
// Default is "BaseLayout".
func WithLayout(layout config.Layout) Option {
	return func(cfg *config.UiConfig) {
		cfg.Layout = config.Value[string]{Value: string(layout), IsSet: true}
	}
}

var (
	PresetAPIPreset config.Preset = "ApiPreset"
)

// WithPresets sets the list of presets to use in Swagger UI. Usually, you'll want to
// include PresetAPIPreset if you use this option.
func WithPresets(preset config.Preset) Option {
	return func(cfg *config.UiConfig) {
		cfg.Layout = config.Value[string]{Value: string(preset), IsSet: true}
	}
}

// WithMaxDisplayedTags limits the number of tagged operations displayed to at most this many.
// The default is to show all operations.
func WithMaxDisplayedTags(maxTags int) Option {
	return func(cfg *config.UiConfig) {
		cfg.MaxDisplayedTags = config.Value[int]{Value: maxTags, IsSet: true}
	}
}

// WithValidatorURL sets the validator URL to use to validate specification files. By default, Swagger UI
// attempts to validate specs against swagger.io's online validator. You can use this parameter to set a
// different validator URL, for example for locally deployed validators (e.g., Validator Badge,
// see https://github.com/swagger-api/validator-badge). Disabling it or setting the URL to 127.0.0.1 or localhost
// will disable validation.
func WithValidatorURL(enabled bool, validatorUrl string) Option {
	return func(cfg *config.UiConfig) {
		var valURL string
		if enabled {
			valURL = validatorUrl
		}

		cfg.ValidatorUrl = config.Value[string]{Value: valURL, IsSet: enabled}
	}
}

// WithShowMutatedRequest configures Swagger UI to use the mutated request returned from a
// requestInterceptor to produce the curl command in the UI, otherwise the request before
// the requestInterceptor was applied is used.
// Refer to https://swagger.io/docs/open-source-tools/swagger-ui/usage/configuration/ for more information.
func WithShowMutatedRequest(showMutatedRequest bool) Option {
	return func(cfg *config.UiConfig) {
		cfg.ShowMutatedRequest = config.Value[bool]{Value: showMutatedRequest, IsSet: true}
	}
}

// WithConfigURL sets the URL to fetch external configuration document from.
func WithConfigURL(configURL string) Option {
	return func(cfg *config.UiConfig) {
		cfg.ConfigURL = config.Value[string]{Value: configURL, IsSet: true}
	}
}

// WithBasePath sets the path prefix Swagger UI is provided on the server
// For example, if Swagger UI is provided under https://example.com/my-service/swagger-ui,
// the base path would be "/my-service/swagger-ui"). This will allow the handler to receive
// requests on path "my-service/swagger-ui" without a trailing slash
// (i.e., "/my-service/swagger-ui/"). Internally, the base path will be used to set
// a prefix for Swagger UI asset files (CSS, JavaScript, etc.).
func WithBasePath(basePath string) Option {
	return func(cfg *config.UiConfig) {
		cfg.BasePath = strings.TrimSuffix(basePath, "/") + "/"
	}
}

const (
	TopBarPlugin config.Plugin = "SwaggerUIBundle.plugins.TopBar"
)

// WithPlugins adds swagger plugins to final build
func WithPlugins(pluginsNames ...config.Plugin) Option {
	return func(cfg *config.UiConfig) {
		for _, name := range pluginsNames {
			cfg.Plugins[name] = struct{}{}
		}
	}
}

// isSpecSet - checks whether specifications is already set
func isSpecSet(cfg *config.UiConfig) bool {
	return len(cfg.Spec) != 0 ||
		cfg.ConfigURL.IsSet ||
		cfg.SpecFilePath != "" ||
		cfg.Url.IsSet ||
		len(cfg.Urls) != 0
}
