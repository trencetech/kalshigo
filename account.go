package kalshigo

import "context"

// GetAccountApiLimits retrieves the API tier limits for the authenticated user.
func (c *Client) GetAccountApiLimits(ctx context.Context) (*GetAccountApiLimitsResponse, error) {
	var result GetAccountApiLimitsResponse
	if err := c.get(ctx, "/account/limits", nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}
