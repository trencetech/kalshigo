package kalshigo

import (
	"context"
	"fmt"
	"net/url"
)

// GetMilestone retrieves a single milestone by ID
func (c *Client) GetMilestone(ctx context.Context, milestoneID string) (*GetMilestoneResponse, error) {
	var result GetMilestoneResponse
	path := fmt.Sprintf("/milestones/%s", milestoneID)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMilestones retrieves milestones with optional filters
func (c *Client) GetMilestones(ctx context.Context, params *GetMilestonesParams) (*GetMilestonesResponse, error) {
	var result GetMilestonesResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/milestones", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLiveData retrieves live data for a specific milestone
func (c *Client) GetLiveData(ctx context.Context, dataType, milestoneID string) (*GetLiveDataResponse, error) {
	var result GetLiveDataResponse
	path := fmt.Sprintf("/live_data/%s/milestone/%s", dataType, milestoneID)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLiveDatas retrieves live data for multiple milestones
func (c *Client) GetLiveDatas(ctx context.Context, milestoneIDs []string) (*GetLiveDatasResponse, error) {
	var result GetLiveDatasResponse
	query := url.Values{}
	for _, id := range milestoneIDs {
		query.Add("milestone_ids", id)
	}
	if err := c.get(ctx, "/live_data/batch", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetIncentivePrograms retrieves incentive programs with optional filters
func (c *Client) GetIncentivePrograms(ctx context.Context, params *GetIncentiveProgramsParams) (*GetIncentiveProgramsResponse, error) {
	var result GetIncentiveProgramsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/incentive_programs", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetStructuredTargets retrieves structured targets with optional filters
func (c *Client) GetStructuredTargets(ctx context.Context, params *GetStructuredTargetsParams) (*GetStructuredTargetsResponse, error) {
	var result GetStructuredTargetsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/structured_targets", query, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetStructuredTarget retrieves a single structured target by ID
func (c *Client) GetStructuredTarget(ctx context.Context, structuredTargetID string) (*GetStructuredTargetResponse, error) {
	var result GetStructuredTargetResponse
	path := fmt.Sprintf("/structured_targets/%s", structuredTargetID)
	if err := c.get(ctx, path, nil, &result, false); err != nil {
		return nil, err
	}
	return &result, nil
}
