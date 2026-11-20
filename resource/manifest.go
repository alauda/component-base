package resource

import (
	"io"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// ParseFromManifest parses a multi-document YAML or JSON manifest into resources.
func ParseFromManifest(manifest string) ([]unstructured.Unstructured, error) {
	decoder := yaml.NewYAMLOrJSONDecoder(strings.NewReader(manifest), 4096)
	resources := make([]unstructured.Unstructured, 0)

	for {
		var resource unstructured.Unstructured
		err := decoder.Decode(&resource)
		if err == io.EOF {
			return resources, nil
		}
		if err != nil {
			return nil, err
		}
		if len(resource.Object) == 0 {
			continue
		}
		resources = append(resources, resource)
	}
}
