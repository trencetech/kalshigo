package kalshigo

import (
	"context"
	"fmt"
	"net/url"
)

// GetMarkets retrieves markets with optional filters
func (c *Client) GetMarkets(ctx context.Context, params *GetMarketsParams) (*GetMarketsResponse, error) {
	var result GetMarketsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/markets", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMarket retrieves a single market by ticker
func (c *Client) GetMarket(ctx context.Context, ticker string) (*GetMarketResponse, error) {
	var result GetMarketResponse
	path := fmt.Sprintf("/markets/%s", ticker)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMarketOrderbook retrieves the order book for a market
func (c *Client) GetMarketOrderbook(ctx context.Context, ticker string, depth int) (*GetMarketOrderbookResponse, error) {
	var result GetMarketOrderbookResponse
	path := fmt.Sprintf("/markets/%s/orderbook", ticker)

	var query url.Values
	if depth > 0 {
		query = url.Values{}
		query.Set("depth", fmt.Sprintf("%d", depth))
	}

	if err := c.get(ctx, path, query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMarketCandlesticks retrieves candlestick data for a market
func (c *Client) GetMarketCandlesticks(ctx context.Context, seriesTicker, ticker string, params *GetCandlesticksParams) (*GetMarketCandlesticksResponse, error) {
	var result GetMarketCandlesticksResponse
	path := fmt.Sprintf("/series/%s/markets/%s/candlesticks", seriesTicker, ticker)
	query := buildQuery(params)
	if err := c.get(ctx, path, query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchGetMarketCandlesticks retrieves candlestick data for multiple markets
func (c *Client) BatchGetMarketCandlesticks(ctx context.Context, params *BatchGetCandlesticksParams) (*BatchGetMarketCandlesticksResponse, error) {
	var result BatchGetMarketCandlesticksResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/markets/candlesticks", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTrades retrieves trades for all markets
func (c *Client) GetTrades(ctx context.Context, params *GetTradesParams) (*GetTradesResponse, error) {
	var result GetTradesResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/markets/trades", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSeries retrieves a single series by ticker
func (c *Client) GetSeries(ctx context.Context, seriesTicker string) (*GetSeriesResponse, error) {
	var result GetSeriesResponse
	path := fmt.Sprintf("/series/%s", seriesTicker)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSeriesList retrieves a list of series
func (c *Client) GetSeriesList(ctx context.Context, params *GetSeriesListParams) (*GetSeriesListResponse, error) {
	var result GetSeriesListResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/series", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}
