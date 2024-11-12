package templates

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexliesenfeld/go-swagger-ui/internal/config"
)

func fromStringConfigValue(v config.Value[string]) string {
	if v.IsSet {
		return strings.ReplaceAll(v.Value, "\n", "\\n")
	}

	return ""
}

func fromDocExpansionConfigValue(v config.Value[config.DocExpansion]) string {
	if v.IsSet {
		return string(v.Value)
	}

	return ""
}

func fromModelRenderingConfigValue(v config.Value[config.ModelRendering]) string {
	if v.IsSet {
		return string(v.Value)
	}

	return ""
}

func fromIntConfigValue(v config.Value[int]) string {
	if v.IsSet {
		return fmt.Sprintf("%d", v.Value)
	}

	return ""
}

func fromBoolConfigValue(v config.Value[bool]) string {
	if v.IsSet {
		return strconv.FormatBool(v.Value)
	}

	return ""
}

func marshalObject(v any) (string, error) {
	if v == nil {
		return "", nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("cannot marshal object: %w", err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func fromMapToArray(in map[config.Plugin]struct{}) []string {
	out := make([]string, 0, len(in))

	for key := range in {
		out = append(out, string(key))
	}

	return out
}
