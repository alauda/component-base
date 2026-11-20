package compress

import (
	"bytes"
	stdcompress "compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
)

var magicGzip = []byte{0x1f, 0x8b, 0x08}

// DecodeToString decodes a base64-encoded, optionally gzip-compressed JSON string.
func DecodeToString(data string) (string, error) {
	decoded, err := Decode(data)
	if err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(decoded, &result); err != nil {
		return "", err
	}
	return result, nil
}

// Decode decodes base64 data and transparently decompresses gzip payloads.
// Non-gzip payloads are returned unchanged for compatibility with older data.
func Decode(data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	if len(decoded) < len(magicGzip) || !bytes.Equal(decoded[:len(magicGzip)], magicGzip) {
		return decoded, nil
	}

	reader, err := stdcompress.NewReader(bytes.NewReader(decoded))
	if err != nil {
		return nil, err
	}
	decompressed, readErr := io.ReadAll(reader)
	closeErr := reader.Close()
	if readErr != nil {
		return nil, readErr
	}
	if closeErr != nil {
		return nil, closeErr
	}
	return decompressed, nil
}
