package kalshigo

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	headerAccessKey       = "KALSHI-ACCESS-KEY"
	headerAccessSignature = "KALSHI-ACCESS-SIGNATURE"
	headerAccessTimestamp = "KALSHI-ACCESS-TIMESTAMP"
)

// Credentials holds API credentials for authentication
type Credentials struct {
	APIKeyID   string
	PrivateKey *rsa.PrivateKey
}

// NewCredentials creates new credentials from an API key ID and private key PEM
func NewCredentials(apiKeyID string, privateKeyPEM string) (*Credentials, error) {
	privateKey, err := parsePrivateKey(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &Credentials{
		APIKeyID:   apiKeyID,
		PrivateKey: privateKey,
	}, nil
}

// parsePrivateKey parses a PEM-encoded RSA private key
// Supports both actual newlines and escaped \n characters (common in env vars)
func parsePrivateKey(pemData string) (*rsa.PrivateKey, error) {
	// Trim any surrounding whitespace or quotes
	pemData = strings.TrimSpace(pemData)
	pemData = strings.Trim(pemData, "\"'")

	// Convert literal \n (backslash followed by n) to actual newlines
	pemData = strings.ReplaceAll(pemData, `\n`, "\n")

	// Also handle \r\n and lone \r
	pemData = strings.ReplaceAll(pemData, "\r\n", "\n")
	pemData = strings.ReplaceAll(pemData, "\r", "\n")

	// Fix malformed PEM headers (missing spaces)
	pemData = strings.ReplaceAll(pemData, "-----BEGINRSAPRIVATEKEY-----", "-----BEGIN RSA PRIVATE KEY-----")
	pemData = strings.ReplaceAll(pemData, "-----ENDRSAPRIVATEKEY-----", "-----END RSA PRIVATE KEY-----")
	pemData = strings.ReplaceAll(pemData, "-----BEGINPRIVATEKEY-----", "-----BEGIN PRIVATE KEY-----")
	pemData = strings.ReplaceAll(pemData, "-----ENDPRIVATEKEY-----", "-----END PRIVATE KEY-----")

	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		// Provide more context for debugging
		preview := pemData
		if len(preview) > 100 {
			preview = preview[:100] + "..."
		}
		return nil, fmt.Errorf("failed to decode PEM block, data starts with: %q", preview)
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, ok := key.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("key is not an RSA private key")
		}
		return rsaKey, nil
	default:
		return nil, fmt.Errorf("unsupported key type: %s", block.Type)
	}
}

// signRequest creates authentication headers for a request
func (c *Credentials) signRequest(timestamp int64, method, path string) (map[string]string, error) {
	timestampStr := strconv.FormatInt(timestamp, 10)

	msgToSign := timestampStr + method + path

	hash := sha256.Sum256([]byte(msgToSign))

	signature, err := rsa.SignPSS(
		rand.Reader,
		c.PrivateKey,
		crypto.SHA256,
		hash[:],
		&rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
			Hash:       crypto.SHA256,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	return map[string]string{
		headerAccessKey:       c.APIKeyID,
		headerAccessSignature: signatureB64,
		headerAccessTimestamp: timestampStr,
	}, nil
}

// GetAuthHeaders returns authentication headers for a request
func (c *Credentials) GetAuthHeaders(method, path string) (map[string]string, error) {
	timestamp := time.Now().UnixMilli()
	return c.signRequest(timestamp, method, path)
}

// GetAuthHeadersWithTimestamp returns authentication headers with a specific timestamp
func (c *Credentials) GetAuthHeadersWithTimestamp(timestamp int64, method, path string) (map[string]string, error) {
	return c.signRequest(timestamp, method, path)
}
