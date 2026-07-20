package kalshigo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// EventOrderSide represents the side of a V2 event order. Unlike the legacy
// OrderSide (yes/no), the V2 event-order endpoints use a single-book bid/ask
// model quoted from the YES leg: bid = buy YES, ask = sell YES. Buying NO is
// expressed as an ask at (1 - noPrice); selling NO as a bid at (1 - noPrice).
type EventOrderSide string

const (
	EventOrderSideBid EventOrderSide = "bid"
	EventOrderSideAsk EventOrderSide = "ask"
)

// CreateEventOrderRequest is the body for the V2 create-order endpoint
// (POST /portfolio/events/orders). It replaces CreateOrderRequest: there is no
// action or type field, a single fixed-point Price replaces yes/no price fields,
// Count is a fixed-point string named "count" (not "count_fp"), and
// SelfTradePreventionType is required.
type CreateEventOrderRequest struct {
	Ticker                  string                  `json:"ticker"`
	Side                    EventOrderSide          `json:"side"`
	Count                   FixedPointCount         `json:"count"`
	Price                   FixedPointDollars       `json:"price"`
	TimeInForce             TimeInForce             `json:"time_in_force"`
	SelfTradePreventionType SelfTradePreventionType `json:"self_trade_prevention_type"`
	ClientOrderID           *string                 `json:"client_order_id,omitempty"`
	ExpirationTime          *int64                  `json:"expiration_time,omitempty"`
	PostOnly                *bool                   `json:"post_only,omitempty"`
	CancelOrderOnPause      *bool                   `json:"cancel_order_on_pause,omitempty"`
	ReduceOnly              *bool                   `json:"reduce_only,omitempty"`
	Subaccount              *int                    `json:"subaccount,omitempty"`
	OrderGroupID            *string                 `json:"order_group_id,omitempty"`
	ExchangeIndex           *int                    `json:"exchange_index,omitempty"`
}

// EventOrderResult is a single element of a batch-create response, and also the
// shape of a single create response. Fills are reported via fixed-point count
// and price fields plus a matching-engine timestamp. Error is populated (and the
// fill fields nil) when that individual order in a batch failed.
type EventOrderResult struct {
	OrderID          string             `json:"order_id"`
	ClientOrderID    string             `json:"client_order_id,omitempty"`
	FillCount        *FixedPointCount   `json:"fill_count,omitempty"`
	RemainingCount   *FixedPointCount   `json:"remaining_count,omitempty"`
	AverageFillPrice *FixedPointDollars `json:"average_fill_price,omitempty"`
	AverageFeePaid   *FixedPointDollars `json:"average_fee_paid,omitempty"`
	TsMs             *int64             `json:"ts_ms,omitempty"`
	Error            *ErrorResponse     `json:"error,omitempty"`
}

// CreateEventOrderResponse is the V2 create-order response. It is flat (no
// nested "order" object). AverageFillPrice/AverageFeePaid are only populated
// when FillCount > 0.
type CreateEventOrderResponse struct {
	OrderID          string             `json:"order_id"`
	ClientOrderID    string             `json:"client_order_id,omitempty"`
	FillCount        FixedPointCount    `json:"fill_count"`
	RemainingCount   FixedPointCount    `json:"remaining_count"`
	AverageFillPrice *FixedPointDollars `json:"average_fill_price,omitempty"`
	AverageFeePaid   *FixedPointDollars `json:"average_fee_paid,omitempty"`
	TsMs             int64              `json:"ts_ms"`
}

// BatchCreateEventOrdersRequest is the body for POST /portfolio/events/orders/batched.
type BatchCreateEventOrdersRequest struct {
	Orders []CreateEventOrderRequest `json:"orders"`
}

// BatchCreateEventOrdersResponse holds one result per submitted order, in order.
type BatchCreateEventOrdersResponse struct {
	Orders []EventOrderResult `json:"orders"`
}

// AmendEventOrderRequest is the body for POST /portfolio/events/orders/{order_id}/amend.
// Count is the new total/max fillable count. Decreasing size preserves queue
// position; any other amendment forfeits it. The subaccount is a query param
// (see AmendEventOrder), not a body field.
type AmendEventOrderRequest struct {
	Ticker               string            `json:"ticker"`
	Side                 EventOrderSide    `json:"side"`
	Price                FixedPointDollars `json:"price"`
	Count                FixedPointCount   `json:"count"`
	ClientOrderID        *string           `json:"client_order_id,omitempty"`
	UpdatedClientOrderID *string           `json:"updated_client_order_id,omitempty"`
	ExchangeIndex        *int              `json:"exchange_index,omitempty"`
}

// AmendEventOrderResponse is the V2 amend response. RemainingCount/FillCount and
// the average fields are nullable.
type AmendEventOrderResponse struct {
	OrderID          string             `json:"order_id"`
	ClientOrderID    string             `json:"client_order_id,omitempty"`
	RemainingCount   *FixedPointCount   `json:"remaining_count,omitempty"`
	FillCount        *FixedPointCount   `json:"fill_count,omitempty"`
	AverageFillPrice *FixedPointDollars `json:"average_fill_price,omitempty"`
	AverageFeePaid   *FixedPointDollars `json:"average_fee_paid,omitempty"`
	TsMs             int64              `json:"ts_ms"`
}

// DecreaseEventOrderRequest is the body for POST /portfolio/events/orders/{order_id}/decrease.
// Exactly one of ReduceBy or ReduceTo must be set. The subaccount is a query
// param (see DecreaseEventOrder), not a body field.
type DecreaseEventOrderRequest struct {
	ReduceBy      *FixedPointCount `json:"reduce_by,omitempty"`
	ReduceTo      *FixedPointCount `json:"reduce_to,omitempty"`
	ExchangeIndex *int             `json:"exchange_index,omitempty"`
}

// DecreaseEventOrderResponse is the V2 decrease response.
type DecreaseEventOrderResponse struct {
	OrderID        string          `json:"order_id"`
	ClientOrderID  string          `json:"client_order_id,omitempty"`
	RemainingCount FixedPointCount `json:"remaining_count"`
	TsMs           int64           `json:"ts_ms"`
}

// CancelEventOrderParams holds the optional query parameters for cancelling a V2
// event order. MarketTicker is required when ExchangeIndex is -1 (auto-routing).
type CancelEventOrderParams struct {
	Subaccount    *int
	ExchangeIndex *int
	MarketTicker  *string
}

// CancelEventOrderResponse is the V2 single-cancel response. ReducedBy is the
// remaining count at the time of cancellation (i.e. the number of contracts
// removed from the book).
type CancelEventOrderResponse struct {
	OrderID       string          `json:"order_id"`
	ClientOrderID string          `json:"client_order_id,omitempty"`
	ReducedBy     FixedPointCount `json:"reduced_by"`
	TsMs          int64           `json:"ts_ms"`
}

// BatchCancelEventOrder identifies one order to cancel in a V2 batch cancel.
// MarketTicker is required when ExchangeIndex is -1 (auto-routing).
type BatchCancelEventOrder struct {
	OrderID       string  `json:"order_id"`
	Subaccount    *int    `json:"subaccount,omitempty"`
	ExchangeIndex *int    `json:"exchange_index,omitempty"`
	MarketTicker  *string `json:"market_ticker,omitempty"`
}

// BatchCancelEventOrdersRequest is the body for DELETE /portfolio/events/orders/batched.
type BatchCancelEventOrdersRequest struct {
	Orders []BatchCancelEventOrder `json:"orders"`
}

// EventOrderCancelResult is a single element of a batch-cancel response. Error is
// populated (and ReducedBy/TsMs nil) when that individual cancellation failed.
type EventOrderCancelResult struct {
	OrderID       string           `json:"order_id"`
	ClientOrderID *string          `json:"client_order_id,omitempty"`
	ReducedBy     *FixedPointCount `json:"reduced_by,omitempty"`
	TsMs          *int64           `json:"ts_ms,omitempty"`
	Error         *ErrorResponse   `json:"error,omitempty"`
}

// BatchCancelEventOrdersResponse holds one result per requested cancellation.
type BatchCancelEventOrdersResponse struct {
	Orders []EventOrderCancelResult `json:"orders"`
}

// CreateEventOrder creates an order via the V2 event-order endpoint
// (POST /portfolio/events/orders). This is the replacement for the deprecated
// CreateOrder (POST /portfolio/orders); callers translate their yes/no +
// buy/sell intent into the bid/ask single-book model (see EventOrderSide).
func (c *Client) CreateEventOrder(ctx context.Context, req *CreateEventOrderRequest) (*CreateEventOrderResponse, error) {
	var result CreateEventOrderResponse
	if err := c.post(ctx, "/portfolio/events/orders", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchCreateEventOrders creates multiple orders in one request via the V2
// endpoint (POST /portfolio/events/orders/batched). Each order is billed
// independently; a per-order failure is reported via the result's Error field
// rather than failing the whole batch.
func (c *Client) BatchCreateEventOrders(ctx context.Context, req *BatchCreateEventOrdersRequest) (*BatchCreateEventOrdersResponse, error) {
	var result BatchCreateEventOrdersResponse
	if err := c.post(ctx, "/portfolio/events/orders/batched", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelEventOrder cancels a single order via the V2 endpoint
// (DELETE /portfolio/events/orders/{order_id}). params may be nil to accept the
// defaults (primary subaccount, exchange_index 0).
func (c *Client) CancelEventOrder(ctx context.Context, orderID string, params *CancelEventOrderParams) (*CancelEventOrderResponse, error) {
	path := fmt.Sprintf("/portfolio/events/orders/%s", orderID)
	var query url.Values
	if params != nil {
		query = eventOrderQuery(params.Subaccount, params.ExchangeIndex, params.MarketTicker)
	}
	var result CancelEventOrderResponse
	if err := c.request(ctx, http.MethodDelete, path, query, nil, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchCancelEventOrders cancels multiple orders in one request via the V2
// endpoint (DELETE /portfolio/events/orders/batched).
func (c *Client) BatchCancelEventOrders(ctx context.Context, req *BatchCancelEventOrdersRequest) (*BatchCancelEventOrdersResponse, error) {
	var result BatchCancelEventOrdersResponse
	if err := c.delete(ctx, "/portfolio/events/orders/batched", req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// AmendEventOrder amends an order's price and/or count via the V2 endpoint
// (POST /portfolio/events/orders/{order_id}/amend). subaccount is an optional
// query parameter (nil = primary).
func (c *Client) AmendEventOrder(ctx context.Context, orderID string, req *AmendEventOrderRequest, subaccount *int) (*AmendEventOrderResponse, error) {
	path := fmt.Sprintf("/portfolio/events/orders/%s/amend", orderID)
	query := eventOrderQuery(subaccount, nil, nil)
	var result AmendEventOrderResponse
	if err := c.request(ctx, http.MethodPost, path, query, req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// DecreaseEventOrder decreases an order's resting count via the V2 endpoint
// (POST /portfolio/events/orders/{order_id}/decrease). Exactly one of
// req.ReduceBy or req.ReduceTo must be set. subaccount is an optional query
// parameter (nil = primary).
func (c *Client) DecreaseEventOrder(ctx context.Context, orderID string, req *DecreaseEventOrderRequest, subaccount *int) (*DecreaseEventOrderResponse, error) {
	path := fmt.Sprintf("/portfolio/events/orders/%s/decrease", orderID)
	query := eventOrderQuery(subaccount, nil, nil)
	var result DecreaseEventOrderResponse
	if err := c.request(ctx, http.MethodPost, path, query, req, &result, true); err != nil {
		return nil, err
	}
	return &result, nil
}

// eventOrderQuery builds the optional query string shared by the V2 mutation
// endpoints. Nil values are omitted so the server applies its defaults.
func eventOrderQuery(subaccount, exchangeIndex *int, marketTicker *string) url.Values {
	q := url.Values{}
	if subaccount != nil {
		q.Set("subaccount", strconv.Itoa(*subaccount))
	}
	if exchangeIndex != nil {
		q.Set("exchange_index", strconv.Itoa(*exchangeIndex))
	}
	if marketTicker != nil {
		q.Set("market_ticker", *marketTicker)
	}
	return q
}
