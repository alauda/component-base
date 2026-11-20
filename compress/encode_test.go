package compress

import (
	"bytes"
	stdcompress "compress/gzip"
	"encoding/base64"
	"encoding/json"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("When data is base64 encoded without compression, it should return the original bytes", func(t *testing.T) {
		want := []byte(`{"apiVersion":"v1"}`)
		got, err := Decode(base64.StdEncoding.EncodeToString(want))
		if err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("Decode() = %q, want %q", got, want)
		}
	})

	t.Run("When data is gzip compressed and base64 encoded, it should return the decompressed bytes", func(t *testing.T) {
		want := []byte(`{"kind":"ConfigMap"}`)
		var compressed bytes.Buffer
		writer := stdcompress.NewWriter(&compressed)
		if _, err := writer.Write(want); err != nil {
			t.Fatalf("gzip.Write() error = %v", err)
		}
		if err := writer.Close(); err != nil {
			t.Fatalf("gzip.Close() error = %v", err)
		}

		got, err := Decode(base64.StdEncoding.EncodeToString(compressed.Bytes()))
		if err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("Decode() = %q, want %q", got, want)
		}
	})

	t.Run("When decoded data is shorter than the gzip header, it should return the data without panicking", func(t *testing.T) {
		want := []byte{0x1f}
		got, err := Decode(base64.StdEncoding.EncodeToString(want))
		if err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("Decode() = %x, want %x", got, want)
		}
	})
}

func TestDecodeToString(t *testing.T) {
	t.Run("When compressed data contains a JSON string, it should return the string", func(t *testing.T) {
		want := "apiVersion: v1\nkind: ConfigMap\n"
		payload, err := json.Marshal(want)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}
		var compressed bytes.Buffer
		writer := stdcompress.NewWriter(&compressed)
		if _, err := writer.Write(payload); err != nil {
			t.Fatalf("gzip.Write() error = %v", err)
		}
		if err := writer.Close(); err != nil {
			t.Fatalf("gzip.Close() error = %v", err)
		}

		got, err := DecodeToString(base64.StdEncoding.EncodeToString(compressed.Bytes()))
		if err != nil {
			t.Fatalf("DecodeToString() error = %v", err)
		}
		if got != want {
			t.Fatalf("DecodeToString() = %q, want %q", got, want)
		}
	})
}
