package kalshigo

import (
	"context"
	"fmt"
)

// GetOrders retrieves orders with optional filters
func (c *Client) GetOrders(ctx context.Context, params *GetOrdersParams) (*GetOrdersResponse, error) {
	var result GetOrdersResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/portfolio/orders", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrder retrieves a single order by ID
func (c *Client) GetOrder(ctx context.Context, orderID string) (*GetOrderResponse, error) {
	var result GetOrderResponse
	path := fmt.Sprintf("/portfolio/orders/%s", orderID)
	if err := c.get(ctx, path, nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateOrder creates a new order
func (c *Client) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	var result CreateOrderResponse
	if err := c.post(ctx, "/portfolio/orders", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelOrder cancels an existing order
func (c *Client) CancelOrder(ctx context.Context, orderID string) (*CancelOrderResponse, error) {
	var result CancelOrderResponse
	path := fmt.Sprintf("/portfolio/orders/%s", orderID)
	if err := c.delete(ctx, path, nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchCreateOrders creates multiple orders at once
func (c *Client) BatchCreateOrders(ctx context.Context, req *BatchCreateOrdersRequest) (*BatchCreateOrdersResponse, error) {
	var result BatchCreateOrdersResponse
	if err := c.post(ctx, "/portfolio/orders/batched", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchCancelOrders cancels multiple orders at once
func (c *Client) BatchCancelOrders(ctx context.Context, req *BatchCancelOrdersRequest) (*BatchCancelOrdersResponse, error) {
	var result BatchCancelOrdersResponse
	if err := c.delete(ctx, "/portfolio/orders/batched", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// AmendOrder amends an existing order
func (c *Client) AmendOrder(ctx context.Context, orderID string, req *AmendOrderRequest) (*AmendOrderResponse, error) {
	var result AmendOrderResponse
	path := fmt.Sprintf("/portfolio/orders/%s/amend", orderID)
	if err := c.post(ctx, path, req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// DecreaseOrder decreases the number of contracts in an existing order
func (c *Client) DecreaseOrder(ctx context.Context, orderID string, req *DecreaseOrderRequest) (*DecreaseOrderResponse, error) {
	var result DecreaseOrderResponse
	path := fmt.Sprintf("/portfolio/orders/%s/decrease", orderID)
	if err := c.post(ctx, path, req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrderQueuePosition retrieves the queue position of an order
func (c *Client) GetOrderQueuePosition(ctx context.Context, orderID string) (*GetOrderQueuePositionResponse, error) {
	var result GetOrderQueuePositionResponse
	path := fmt.Sprintf("/portfolio/orders/%s/queue_position", orderID)
	if err := c.get(ctx, path, nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrderQueuePositions retrieves queue positions for multiple orders
func (c *Client) GetOrderQueuePositions(ctx context.Context, params *GetOrderQueuePositionsParams) (*GetOrderQueuePositionsResponse, error) {
	var result GetOrderQueuePositionsResponse
	query := buildQuery(params)
	if err := c.get(ctx, "/portfolio/orders/queue_positions", query, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrderGroups retrieves all order groups
func (c *Client) GetOrderGroups(ctx context.Context) (*GetOrderGroupsResponse, error) {
	var result GetOrderGroupsResponse
	if err := c.get(ctx, "/portfolio/order_groups", nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrderGroup retrieves a single order group
func (c *Client) GetOrderGroup(ctx context.Context, orderGroupID string) (*GetOrderGroupResponse, error) {
	var result GetOrderGroupResponse
	path := fmt.Sprintf("/portfolio/order_groups/%s", orderGroupID)
	if err := c.get(ctx, path, nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateOrderGroup creates a new order group
func (c *Client) CreateOrderGroup(ctx context.Context, req *CreateOrderGroupRequest) (*CreateOrderGroupResponse, error) {
	var result CreateOrderGroupResponse
	if err := c.post(ctx, "/portfolio/order_groups/create", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteOrderGroup deletes an order group
func (c *Client) DeleteOrderGroup(ctx context.Context, orderGroupID string) error {
	path := fmt.Sprintf("/portfolio/order_groups/%s", orderGroupID)
	if err := c.delete(ctx, path, nil, nil, true); err != nil {
		return err
	}
	return nil
}

// ResetOrderGroup resets an order group's matched contracts counter
func (c *Client) ResetOrderGroup(ctx context.Context, orderGroupID string) error {
	path := fmt.Sprintf("/portfolio/order_groups/%s/reset", orderGroupID)
	if err := c.put(ctx, path, nil, nil, true); err != nil {
		return err
	}
	return nil
}

// TriggerOrderGroup triggers the order group, canceling all orders in the group.
func (c *Client) TriggerOrderGroup(ctx context.Context, orderGroupID string) error {
	path := fmt.Sprintf("/portfolio/order_groups/%s/trigger", orderGroupID)
	if err := c.put(ctx, path, nil, nil, true); err != nil {
		return err
	}
	return nil
}

// UpdateOrderGroupLimit updates the order group's contracts limit.
func (c *Client) UpdateOrderGroupLimit(ctx context.Context, orderGroupID string, req *UpdateOrderGroupLimitRequest) error {
	path := fmt.Sprintf("/portfolio/order_groups/%s/limit", orderGroupID)
	if err := c.put(ctx, path, req, nil, true); err != nil {
		return err
	}
	return nil
}
