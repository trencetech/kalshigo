package kalshigo

import (
	"context"
	"fmt"
)

// GetCommunicationsID retrieves the communications ID for the authenticated user
func (c *Client) GetCommunicationsID(ctx context.Context) (*GetCommunicationsIDResponse, error) {
	var result GetCommunicationsIDResponse
	if err := c.get(ctx, "/communications/id", nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRFQs retrieves RFQs with optional filters
func (c *Client) GetRFQs(ctx context.Context, params *GetRFQsParams) (*GetRFQsResponse, error) {
	var result GetRFQsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/communications/rfqs", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRFQ retrieves a single RFQ by ID
func (c *Client) GetRFQ(ctx context.Context, rfqID string) (*GetRFQResponse, error) {
	var result GetRFQResponse
	path := fmt.Sprintf("/communications/rfqs/%s", rfqID)
	if err := c.get(ctx, path, nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateRFQ creates a new RFQ
func (c *Client) CreateRFQ(ctx context.Context, req *CreateRFQRequest) (*CreateRFQResponse, error) {
	var result CreateRFQResponse
	if err := c.post(ctx, "/communications/rfqs", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteRFQ deletes an RFQ by ID
func (c *Client) DeleteRFQ(ctx context.Context, rfqID string) error {
	path := fmt.Sprintf("/communications/rfqs/%s", rfqID)
	if err := c.delete(ctx, path, nil, nil, true); err != nil {
		return err
	}
	return nil
}

// GetQuotes retrieves quotes with optional filters
func (c *Client) GetQuotes(ctx context.Context, params *GetQuotesParams) (*GetQuotesResponse, error) {
	var result GetQuotesResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/communications/quotes", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetQuote retrieves a single quote by ID
func (c *Client) GetQuote(ctx context.Context, quoteID string) (*GetQuoteResponse, error) {
	var result GetQuoteResponse
	path := fmt.Sprintf("/communications/quotes/%s", quoteID)
	if err := c.get(ctx, path, nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateQuote creates a quote in response to an RFQ
func (c *Client) CreateQuote(ctx context.Context, req *CreateQuoteRequest) (*CreateQuoteResponse, error) {
	var result CreateQuoteResponse
	if err := c.post(ctx, "/communications/quotes", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteQuote deletes a quote by ID
func (c *Client) DeleteQuote(ctx context.Context, quoteID string) error {
	path := fmt.Sprintf("/communications/quotes/%s", quoteID)
	if err := c.delete(ctx, path, nil, nil, true); err != nil {
		return err
	}
	return nil
}

// AcceptQuote accepts a quote
func (c *Client) AcceptQuote(ctx context.Context, quoteID string, req *AcceptQuoteRequest) error {
	path := fmt.Sprintf("/communications/quotes/%s/accept", quoteID)
	if err := c.put(ctx, path, req, nil, true); err != nil {
		return err
	}
	return nil
}

// ConfirmQuote confirms a quote
func (c *Client) ConfirmQuote(ctx context.Context, quoteID string) error {
	path := fmt.Sprintf("/communications/quotes/%s/confirm", quoteID)
	if err := c.put(ctx, path, nil, nil, true); err != nil {
		return err
	}
	return nil
}
