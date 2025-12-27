package kalshigo

import (
	"context"
	"fmt"
)

// GetApiKeys retrieves all API keys for the authenticated user
func (c *Client) GetApiKeys(ctx context.Context) (*GetApiKeysResponse, error) {
	var result GetApiKeysResponse
	if err := c.get(ctx, "/api_keys", nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateApiKey creates a new API key with a user-provided public key
func (c *Client) CreateApiKey(ctx context.Context, req *CreateApiKeyRequest) (*CreateApiKeyResponse, error) {
	var result CreateApiKeyResponse
	if err := c.post(ctx, "/api_keys", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GenerateApiKey generates a new API key with an automatically created key pair
func (c *Client) GenerateApiKey(ctx context.Context, req *GenerateApiKeyRequest) (*GenerateApiKeyResponse, error) {
	var result GenerateApiKeyResponse
	if err := c.post(ctx, "/api_keys/generate", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteApiKey deletes an existing API key
func (c *Client) DeleteApiKey(ctx context.Context, apiKey string) error {
	path := fmt.Sprintf("/api_keys/%s", apiKey)
	if err := c.delete(ctx, path, nil, nil, true); err != nil {
		return err
	}
	return nil
}
