package kalshigo

import (
	"context"
	"fmt"
	"net/url"
)

// GetHistoricalCutoff returns the cutoff timestamps that define the boundary between live and historical data.
func (c *Client) GetHistoricalCutoff(ctx context.Context) (*GetHistoricalCutoffResponse, error) {
	var result GetHistoricalCutoffResponse
	if err := c.get(ctx, "/historical/cutoff", nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHistoricalMarkets retrieves markets that have been archived to the historical database.
func (c *Client) GetHistoricalMarkets(ctx context.Context, params *GetHistoricalMarketsParams) (*GetMarketsResponse, error) {
	var query url.Values
	if params != nil {
		query = buildQuery(params)
	}
	var result GetMarketsResponse
	if err := c.get(ctx, "/historical/markets", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHistoricalMarket retrieves a specific historical market by its ticker.
func (c *Client) GetHistoricalMarket(ctx context.Context, ticker string) (*GetMarketResponse, error) {
	var result GetMarketResponse
	path := fmt.Sprintf("/historical/markets/%s", ticker)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHistoricalMarketCandlesticks retrieves candlestick data for archived markets.
func (c *Client) GetHistoricalMarketCandlesticks(ctx context.Context, ticker string, params *GetHistoricalCandlesticksParams) (*GetMarketCandlesticksHistoricalResponse, error) {
	var query url.Values
	if params != nil {
		query = buildQuery(params)
	}
	var result GetMarketCandlesticksHistoricalResponse
	path := fmt.Sprintf("/historical/markets/%s/candlesticks", ticker)
	if err := c.get(ctx, path, query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHistoricalFills retrieves historical fills for the authenticated user.
func (c *Client) GetHistoricalFills(ctx context.Context, params *GetHistoricalFillsParams) (*GetFillsResponse, error) {
	var query url.Values
	if params != nil {
		query = buildQuery(params)
	}
	var result GetFillsResponse
	if err := c.get(ctx, "/historical/fills", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHistoricalOrders retrieves orders that have been archived to the historical database.
func (c *Client) GetHistoricalOrders(ctx context.Context, params *GetHistoricalOrdersParams) (*GetOrdersResponse, error) {
	var query url.Values
	if params != nil {
		query = buildQuery(params)
	}
	var result GetOrdersResponse
	if err := c.get(ctx, "/historical/orders", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHistoricalTrades retrieves all historical trades.
func (c *Client) GetHistoricalTrades(ctx context.Context, params *GetHistoricalTradesParams) (*GetTradesResponse, error) {
	var query url.Values
	if params != nil {
		query = buildQuery(params)
	}
	var result GetTradesResponse
	if err := c.get(ctx, "/historical/trades", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}
