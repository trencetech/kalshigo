package kalshigo

import (
	"context"
	"fmt"
	"net/url"
)

// GetEvents retrieves events with optional filters
func (c *Client) GetEvents(ctx context.Context, params *GetEventsParams) (*GetEventsResponse, error) {
	var result GetEventsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/events", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMultivariateEvents retrieves multivariate events with optional filters
func (c *Client) GetMultivariateEvents(ctx context.Context, params *GetMultivariateEventsParams) (*GetMultivariateEventsResponse, error) {
	var result GetMultivariateEventsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/events/multivariate", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEvent retrieves a single event by ticker
func (c *Client) GetEvent(ctx context.Context, eventTicker string, withNestedMarkets bool) (*GetEventResponse, error) {
	var result GetEventResponse
	path := fmt.Sprintf("/events/%s", eventTicker)

	var query url.Values
	if withNestedMarkets {
		query = url.Values{}
		query.Set("with_nested_markets", "true")
	}

	if err := c.get(ctx, path, query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEventMetadata retrieves metadata for an event
func (c *Client) GetEventMetadata(ctx context.Context, eventTicker string) (*GetEventMetadataResponse, error) {
	var result GetEventMetadataResponse
	path := fmt.Sprintf("/events/%s/metadata", eventTicker)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEventCandlesticks retrieves candlestick data for an event
func (c *Client) GetEventCandlesticks(ctx context.Context, seriesTicker, eventTicker string, params *GetCandlesticksParams) (*GetEventCandlesticksResponse, error) {
	var result GetEventCandlesticksResponse
	path := fmt.Sprintf("/series/%s/events/%s/candlesticks", seriesTicker, eventTicker)
	query := buildQuery(params)
	if err := c.get(ctx, path, query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEventForecastPercentilesHistory retrieves forecast percentile history for an event
func (c *Client) GetEventForecastPercentilesHistory(ctx context.Context, seriesTicker, eventTicker string, params *GetForecastHistoryParams) (*GetEventForecastPercentilesHistoryResponse, error) {
	var result GetEventForecastPercentilesHistoryResponse
	path := fmt.Sprintf("/series/%s/events/%s/forecast_percentile_history", seriesTicker, eventTicker)
	query := buildQuery(params)
	if err := c.get(ctx, path, query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}
