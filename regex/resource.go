package regex

import (
	"fmt"
	"regexp"
)

var (
	// resourceNameRegex is resource name regex for most kubernetes resources
	resourceNameRegex = `^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
)

//IsValidResourceName validate if a resource name is valid
func IsValidResourceName(name string) bool {
	matched, _ := regexp.MatchString(resourceNameRegex, name)
	return matched
}

// CreateResourceRegexError create kubernetes like error message in kubectl about invalid HelmRequest
// resource
func CreateResourceRegexError(regex, kind, name, key, value string) string {
	template := `The %s "%s" is invalid: %s: Invalid value: "%s": ` +
		`a DNS-1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', ` +
		`and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '%s')`
	return fmt.Sprintf(template, kind, name, key, value, regex)
}
