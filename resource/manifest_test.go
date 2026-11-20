package resource

import (
	"testing"
)

func TestParseFromManifest(t *testing.T) {
	t.Run("When manifest contains multiple YAML documents, it should return each resource", func(t *testing.T) {
		manifest := `apiVersion: v1
kind: ConfigMap
metadata:
  name: first
---
apiVersion: v1
kind: Secret
metadata:
  name: second
`

		got, err := ParseFromManifest(manifest)
		if err != nil {
			t.Fatalf("ParseFromManifest() error = %v", err)
		}
		if len(got) != 2 {
			t.Fatalf("ParseFromManifest() returned %d resources, want 2", len(got))
		}
		if got[0].GetKind() != "ConfigMap" || got[0].GetName() != "first" {
			t.Fatalf("first resource = %s/%s, want ConfigMap/first", got[0].GetKind(), got[0].GetName())
		}
		if got[1].GetKind() != "Secret" || got[1].GetName() != "second" {
			t.Fatalf("second resource = %s/%s, want Secret/second", got[1].GetKind(), got[1].GetName())
		}
	})

	t.Run("When manifest contains invalid YAML, it should return an error", func(t *testing.T) {
		_, err := ParseFromManifest("apiVersion: v1\nmetadata: [")
		if err == nil {
			t.Fatal("ParseFromManifest() error = nil, want an error")
		}
	})
}
