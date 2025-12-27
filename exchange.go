package kalshigo

import (
	"context"
)

// GetExchangeStatus gets the current exchange status
func (c *Client) GetExchangeStatus(ctx context.Context) (*ExchangeStatus, error) {
	var result ExchangeStatus
	if err := c.get(ctx, "/exchange/status", nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetExchangeAnnouncements gets all exchange-wide announcements
func (c *Client) GetExchangeAnnouncements(ctx context.Context) (*GetExchangeAnnouncementsResponse, error) {
	var result GetExchangeAnnouncementsResponse
	if err := c.get(ctx, "/exchange/announcements", nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetExchangeSchedule gets the exchange schedule
func (c *Client) GetExchangeSchedule(ctx context.Context) (*GetExchangeScheduleResponse, error) {
	var result GetExchangeScheduleResponse
	if err := c.get(ctx, "/exchange/schedule", nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserDataTimestamp gets the timestamp when user data was last validated
func (c *Client) GetUserDataTimestamp(ctx context.Context) (*GetUserDataTimestampResponse, error) {
	var result GetUserDataTimestampResponse
	if err := c.get(ctx, "/exchange/user_data_timestamp", nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSeriesFeeChanges gets series fee changes
func (c *Client) GetSeriesFeeChanges(ctx context.Context, params *GetSeriesFeeChangesParams) (*GetSeriesFeeChangesResponse, error) {
	var result GetSeriesFeeChangesResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/series/fee_changes", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}
