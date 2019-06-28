package regex

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestCreateResourceRegexError(t *testing.T) {
	exp := `The ConfigMap "ff--" is invalid: metadata.name: Invalid value: "ff--": a DNS-1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$')`
	assert.Equal(t, exp, DefaultResourceNameRegexError(
		"ConfigMap",
		"ff--",
		"metadata.name",
		"ff--",
	))
}

func TestIsValidResourceName(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		assert.Equal(t, false, IsValidResourceName("f-0---"))
	})

	t.Run("invalid", func(t *testing.T) {
		assert.Equal(t, true, IsValidResourceName("abc-b"))
	})
}
