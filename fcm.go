package kalshigo

import (
	"context"
)

// GetFCMOrders retrieves orders for FCM members filtered by subtrader ID
func (c *Client) GetFCMOrders(ctx context.Context, params *GetFCMOrdersParams) (*GetOrdersResponse, error) {
	var result GetOrdersResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/fcm/orders", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFCMPositions retrieves positions for FCM members filtered by subtrader ID
func (c *Client) GetFCMPositions(ctx context.Context, params *GetFCMPositionsParams) (*GetPositionsResponse, error) {
	var result GetPositionsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/fcm/positions", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}
