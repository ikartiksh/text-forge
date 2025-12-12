package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/png"
	"strings"
)

func FormatJSON(input string) (string, error) {
	var jsonObj interface{}

	if err := json.Unmarshal([]byte(input), &jsonObj); err != nil {
		return "", fmt.Errorf("invalid JSON: %v", err)
	}
	pretty, err := json.MarshalIndent(jsonObj, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting JSON: %v", err)
	}
	return string(pretty), nil
}

func DecodeJWT(token string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid JWT format")
	}

	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode header: %v", err)
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %v", err)
	}

	var header, payload map[string]interface{}
	json.Unmarshal(headerJSON, &header)
	json.Unmarshal(payloadJSON, &payload)

	result := map[string]interface{}{
		"header":  header,
		"payload": payload,
	}

	return result, nil
}

func GenerateQRCode(content string, size int) (string, error) {
	if size == 0 {
		size = 256
	}

	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return "", err
	}

	qr.DisableBorder = false

	img := qr.Image(size)

	var buf strings.Builder
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err = png.Encode(encoder, img)
	if err != nil {
		return "", err
	}
	encoder.Close()

	return "data:image/png;base64," + buf.String(), nil
}
