package hash

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

type testSpec struct {
	ClusterName string `json:"clusterName"`
	Chart       string `json:"chart"`
	Version     string `json:"version"`
	Namespace   string `json:"namespace"`
}

func TestHash(t *testing.T) {
	data := testSpec{
		ClusterName: "a",
		Chart:       "ba/b",
		Version:     "1.2",
	}
	assert.Equal(t, "13162844096571394621", GenHashStr(data))
}
