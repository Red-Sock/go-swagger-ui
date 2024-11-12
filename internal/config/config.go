package config

type Value[T any] struct {
	IsSet bool
	Value T
}

type UiConfig struct {
	TtmlTitle string
	BasePath  string

	Spec         []byte
	ConfigURL    Value[string]
	SpecFilePath string
	Url          Value[string]
	Urls         []SpecURL
	UrlsPrimary  Value[string]

	Layout                   Value[string]
	DocExpansion             Value[DocExpansion]
	DefaultModelExpandDepth  Value[int]
	DefaultModelsExpandDepth Value[int]
	DefaultModelRendering    Value[ModelRendering]
	QueryConfigEnabled       Value[bool]
	SupportedSubmitMethods   []string
	ShowMutatedRequest       Value[bool]
	DeepLinking              Value[bool]
	ShowExtensions           Value[bool]
	ShowCommonExtensions     Value[bool]
	Filter                   Value[bool]
	FilterString             Value[string]
	DisplayOperationID       Value[bool]
	TryItOutEnabled          Value[bool]
	DisplayRequestDuration   Value[bool]
	PersistAuthorization     Value[bool]
	WithCredentials          Value[bool]
	Oauth2RedirectUrl        Value[string]
	MaxDisplayedTags         Value[int]
	ValidatorUrl             Value[string]
	Plugins                  map[Plugin]struct{}
}

type DocExpansion string

type ModelRendering string

type Layout string

type Preset string

type SpecURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Plugin string
