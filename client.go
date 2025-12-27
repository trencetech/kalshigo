package kalshigo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	DefaultBaseURL   = "https://api.elections.kalshi.com/trade-api/v2"
	DemoBaseURL      = "https://demo-api.kalshi.co/trade-api/v2"
	DefaultTimeout   = 2 * time.Minute
	DefaultUserAgent = "kalshi-go-client/1.0"
)

// ClientConfig holds configuration for creating a Client
type ClientConfig struct {
	BaseURL    string
	APIKeyID   string
	PrivateKey string
	HTTPClient *http.Client
	Logger     zap.Logger
	UserAgent  string
}

// Client is the Kalshi API client
type Client struct {
	baseURL     string
	credentials *Credentials
	httpClient  *http.Client
	logger      zap.Logger
	userAgent   string
}

// NewClient creates a new Kalshi API client
func NewClient(config ClientConfig) *Client {
	if config.BaseURL == "" {
		config.BaseURL = DefaultBaseURL
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: DefaultTimeout,
		}
	}

	if config.UserAgent == "" {
		config.UserAgent = DefaultUserAgent
	}

	var creds *Credentials
	if config.APIKeyID != "" && config.PrivateKey != "" {
		var err error
		creds, err = NewCredentials(config.APIKeyID, config.PrivateKey)
		if err != nil {
			return nil
		}
	}

	return &Client{
		baseURL:     strings.TrimSuffix(config.BaseURL, "/"),
		credentials: creds,
		httpClient:  config.HTTPClient,
		logger:      config.Logger,
		userAgent:   config.UserAgent,
	}
}

// NewClientWithCredentials creates a new client with pre-parsed credentials
func NewClientWithCredentials(baseURL string, creds *Credentials, httpClient *http.Client, logger zap.Logger) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	if httpClient == nil {
		httpClient = &http.Client{Timeout: DefaultTimeout}
	}

	return &Client{
		baseURL:     strings.TrimSuffix(baseURL, "/"),
		credentials: creds,
		httpClient:  httpClient,
		logger:      logger,
		userAgent:   DefaultUserAgent,
	}
}

// request performs an HTTP request to the Kalshi API
func (c *Client) request(ctx context.Context, method, path string, query url.Values, body any, result any, authenticated bool) error {
	fullURL := c.baseURL + path
	if len(query) > 0 {
		fullURL += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	if authenticated {
		if c.credentials == nil {
			return fmt.Errorf("authentication required but no credentials configured")
		}

		// Kalshi requires the full path with /trade-api/v2 prefix for signing
		// Query parameters must NOT be included in the signature
		parsedURL, err := url.Parse(c.baseURL)
		if err != nil {
			return fmt.Errorf("failed to parse base URL: %w", err)
		}
		authPath := parsedURL.Path + path

		headers, err := c.credentials.GetAuthHeaders(method, authPath)
		if err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	c.logger.Debug("kalshi request", zap.String("url", fullURL), zap.String("method", method))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	c.logger.Debug("kalshi response", zap.String("url", fullURL), zap.Int("status", resp.StatusCode), zap.Any("body", respBody))

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		if json.Unmarshal(respBody, &errResp) == nil && (errResp.Message != "" || errResp.Code != "") {
			return &errResp
		}
		return fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// get performs a GET request
func (c *Client) get(ctx context.Context, path string, query url.Values, result any, authenticated bool) error {
	return c.request(ctx, http.MethodGet, path, query, nil, result, authenticated)
}

// post performs a POST request
func (c *Client) post(ctx context.Context, path string, body any, result any, authenticated bool) error {
	return c.request(ctx, http.MethodPost, path, nil, body, result, authenticated)
}

// put performs a PUT request
func (c *Client) put(ctx context.Context, path string, body any, result any, authenticated bool) error {
	return c.request(ctx, http.MethodPut, path, nil, body, result, authenticated)
}

// delete performs a DELETE request
func (c *Client) delete(ctx context.Context, path string, body any, result any, authenticated bool) error {
	return c.request(ctx, http.MethodDelete, path, nil, body, result, authenticated)
}

// buildQuery builds a query string from a params struct using the url tag
func buildQuery(params any) url.Values {
	query := url.Values{}
	if params == nil {
		return query
	}

	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return query
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return query
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("url")
		if tag == "" || tag == "-" {
			continue
		}

		parts := strings.Split(tag, ",")
		name := parts[0]
		omitEmpty := len(parts) > 1 && parts[1] == "omitempty"

		if omitEmpty && isEmptyValue(field) {
			continue
		}

		var value string
		switch field.Kind() {
		case reflect.String:
			value = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() != 0 || !omitEmpty {
				value = strconv.FormatInt(field.Int(), 10)
			}
		case reflect.Bool:
			if field.Bool() || !omitEmpty {
				value = strconv.FormatBool(field.Bool())
			}
		case reflect.Slice:
			if field.Len() > 0 {
				for j := 0; j < field.Len(); j++ {
					elem := field.Index(j)
					switch elem.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						query.Add(name, strconv.FormatInt(elem.Int(), 10))
					case reflect.String:
						query.Add(name, elem.String())
					}
				}
				continue
			}
		}

		if value != "" {
			query.Set(name, value)
		}
	}

	return query
}

// isEmptyValue checks if a reflect.Value is empty
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
