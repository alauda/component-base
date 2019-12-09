package resource

import (
	"bytes"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"
)

// ParseFromManifest parse a string to resource list, the string is a multi-part yaml
func ParseFromManifest(manifest string) ([]unstructured.Unstructured, error) {
	var rs []unstructured.Unstructured

	ss := strings.Split(manifest, "---")
	for _, item := range ss {
		if item == "" {
			continue
		}
		var res unstructured.Unstructured
		d := yaml.NewYAMLOrJSONDecoder(bytes.NewBufferString(item), len([]byte(item)))
		err := d.Decode(&res)

		// err := yaml.Unmarshal([]byte(item), &res)
		if err != nil {
			return nil, err
		}
		rs = append(rs, res)
	}

	return rs, nil

}
