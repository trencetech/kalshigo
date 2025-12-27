package kalshigo

import (
	"context"
)

// GetBalance retrieves the user's balance and portfolio value
func (c *Client) GetBalance(ctx context.Context) (*GetBalanceResponse, error) {
	var result GetBalanceResponse
	if err := c.get(ctx, "/portfolio/balance", nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPositions retrieves the user's positions
func (c *Client) GetPositions(ctx context.Context, params *GetPositionsParams) (*GetPositionsResponse, error) {
	var result GetPositionsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/portfolio/positions", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSettlements retrieves the user's settlement history
func (c *Client) GetSettlements(ctx context.Context, params *GetSettlementsParams) (*GetSettlementsResponse, error) {
	var result GetSettlementsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/portfolio/settlements", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPortfolioRestingOrderTotalValue retrieves the total value of resting orders
func (c *Client) GetPortfolioRestingOrderTotalValue(ctx context.Context) (*GetPortfolioRestingOrderTotalValueResponse, error) {
	var result GetPortfolioRestingOrderTotalValueResponse
	if err := c.get(ctx, "/portfolio/summary/total_resting_order_value", nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFills retrieves the user's fills
func (c *Client) GetFills(ctx context.Context, params *GetFillsParams) (*GetFillsResponse, error) {
	var result GetFillsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/portfolio/fills", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}
