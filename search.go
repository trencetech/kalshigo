package kalshigo

import (
	"context"
)

// GetTagsForSeriesCategories retrieves tags organized by series categories
func (c *Client) GetTagsForSeriesCategories(ctx context.Context) (*GetTagsForSeriesCategoriesResponse, error) {
	var result GetTagsForSeriesCategoriesResponse
	if err := c.get(ctx, "/search/tags_by_categories", nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFiltersForSports retrieves available filters organized by sport
func (c *Client) GetFiltersForSports(ctx context.Context) (*GetFiltersBySportsResponse, error) {
	var result GetFiltersBySportsResponse
	if err := c.get(ctx, "/search/filters_by_sport", nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}
