package kalshigo

import (
	"context"
	"fmt"
)

// GetMultivariateEventCollection retrieves a multivariate event collection by ticker
func (c *Client) GetMultivariateEventCollection(ctx context.Context, collectionTicker string) (*GetMultivariateEventCollectionResponse, error) {
	var result GetMultivariateEventCollectionResponse
	path := fmt.Sprintf("/multivariate_event_collections/%s", collectionTicker)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMultivariateEventCollections retrieves multivariate event collections with optional filters
func (c *Client) GetMultivariateEventCollections(ctx context.Context, params *GetMveCollectionsParams) (*GetMultivariateEventCollectionsResponse, error) {
	var result GetMultivariateEventCollectionsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/multivariate_event_collections", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateMarketInMultivariateEventCollection creates a market in a multivariate event collection
func (c *Client) CreateMarketInMultivariateEventCollection(ctx context.Context, collectionTicker string, req *CreateMarketInMveCollectionRequest) (*CreateMarketInMveCollectionResponse, error) {
	var result CreateMarketInMveCollectionResponse
	path := fmt.Sprintf("/multivariate_event_collections/%s", collectionTicker)
	if err := c.post(ctx, path, req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// LookupTickersInMultivariateEventCollection looks up tickers for a market in a multivariate event collection
func (c *Client) LookupTickersInMultivariateEventCollection(ctx context.Context, collectionTicker string, req *LookupTickersInMveCollectionRequest) (*LookupTickersInMveCollectionResponse, error) {
	var result LookupTickersInMveCollectionResponse
	path := fmt.Sprintf("/multivariate_event_collections/%s/lookup", collectionTicker)
	if err := c.put(ctx, path, req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMultivariateEventCollectionLookupHistory retrieves lookup history for a multivariate event collection
func (c *Client) GetMultivariateEventCollectionLookupHistory(ctx context.Context, collectionTicker string, params *GetMveLookupHistoryParams) (*GetMveLookupHistoryResponse, error) {
	var result GetMveLookupHistoryResponse
	path := fmt.Sprintf("/multivariate_event_collections/%s/lookup", collectionTicker)
	query := buildQuery(params)
	if err := c.get(ctx, path, query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}
