package kalshigo

// CreateOrderRequest represents a request to create an order
type CreateOrderRequest struct {
	Ticker                  string                   `json:"ticker"`
	ClientOrderID           *string                  `json:"client_order_id,omitempty"`
	Side                    OrderSide                `json:"side"`
	Action                  OrderAction              `json:"action"`
	Count                   *int                     `json:"count,omitempty"`
	CountFp                 *FixedPointCount         `json:"count_fp,omitempty"`
	Type                    *OrderType               `json:"type,omitempty"`
	YesPrice                *int                     `json:"yes_price,omitempty"`
	NoPrice                 *int                     `json:"no_price,omitempty"`
	YesPriceDollars         *FixedPointDollars       `json:"yes_price_dollars,omitempty"`
	NoPriceDollars          *FixedPointDollars       `json:"no_price_dollars,omitempty"`
	ExpirationTs            *int64                   `json:"expiration_ts,omitempty"`
	TimeInForce             *TimeInForce             `json:"time_in_force,omitempty"`
	BuyMaxCost              *int                     `json:"buy_max_cost,omitempty"`
	PostOnly                *bool                    `json:"post_only,omitempty"`
	ReduceOnly              *bool                    `json:"reduce_only,omitempty"`
	SellPositionFloor       *int                     `json:"sell_position_floor,omitempty"`
	SelfTradePreventionType *SelfTradePreventionType `json:"self_trade_prevention_type,omitempty"`
	OrderGroupID            *string                  `json:"order_group_id,omitempty"`
	CancelOrderOnPause      *bool                    `json:"cancel_order_on_pause,omitempty"`
	Subaccount              *int                     `json:"subaccount,omitempty"`
}

// BatchCreateOrdersRequest represents a batch order creation request
type BatchCreateOrdersRequest struct {
	Orders []CreateOrderRequest `json:"orders"`
}

// BatchCancelOrdersRequest represents a batch order cancellation request
type BatchCancelOrdersRequest struct {
	IDs    []string                        `json:"ids,omitempty"`
	Orders []BatchCancelOrdersRequestOrder `json:"orders,omitempty"`
}

// BatchCancelOrdersRequestOrder represents an individual order to cancel in a batch request
type BatchCancelOrdersRequestOrder struct {
	OrderID    string `json:"order_id"`
	Subaccount *int   `json:"subaccount,omitempty"`
}

// AmendOrderRequest represents a request to amend an order
type AmendOrderRequest struct {
	Ticker               string             `json:"ticker"`
	Side                 OrderSide          `json:"side"`
	Action               OrderAction        `json:"action"`
	ClientOrderID        string             `json:"client_order_id,omitempty"`
	UpdatedClientOrderID string             `json:"updated_client_order_id,omitempty"`
	YesPrice             *int               `json:"yes_price,omitempty"`
	NoPrice              *int               `json:"no_price,omitempty"`
	YesPriceDollars      *FixedPointDollars `json:"yes_price_dollars,omitempty"`
	NoPriceDollars       *FixedPointDollars `json:"no_price_dollars,omitempty"`
	Count                *int               `json:"count,omitempty"`
	CountFp              *FixedPointCount   `json:"count_fp,omitempty"`
	Subaccount           *int               `json:"subaccount,omitempty"`
}

// DecreaseOrderRequest represents a request to decrease an order
type DecreaseOrderRequest struct {
	ReduceBy   *int             `json:"reduce_by,omitempty"`
	ReduceByFp *FixedPointCount `json:"reduce_by_fp,omitempty"`
	ReduceTo   *int             `json:"reduce_to,omitempty"`
	ReduceToFp *FixedPointCount `json:"reduce_to_fp,omitempty"`
	Subaccount *int             `json:"subaccount,omitempty"`
}

// CreateOrderGroupRequest represents a request to create an order group
type CreateOrderGroupRequest struct {
	ContractsLimit   *int64           `json:"contracts_limit,omitempty"`
	ContractsLimitFp *FixedPointCount `json:"contracts_limit_fp,omitempty"`
	Subaccount       *int             `json:"subaccount,omitempty"`
}

// CreateApiKeyRequest represents a request to create an API key
type CreateApiKeyRequest struct {
	Name      string   `json:"name"`
	PublicKey string   `json:"public_key"`
	Scopes    []string `json:"scopes,omitempty"`
}

// GenerateApiKeyRequest represents a request to generate an API key
type GenerateApiKeyRequest struct {
	Name   string   `json:"name"`
	Scopes []string `json:"scopes,omitempty"`
}

// CreateRFQRequest represents a request to create an RFQ
type CreateRFQRequest struct {
	MarketTicker         string             `json:"market_ticker"`
	Contracts            *int               `json:"contracts,omitempty"`
	ContractsFp          *FixedPointCount   `json:"contracts_fp,omitempty"`
	TargetCostCentiCents *int64             `json:"target_cost_centi_cents,omitempty"`
	TargetCostDollars    *FixedPointDollars `json:"target_cost_dollars,omitempty"`
	RestRemainder        bool               `json:"rest_remainder"`
	ReplaceExisting      *bool              `json:"replace_existing,omitempty"`
	SubtraderID          *string            `json:"subtrader_id,omitempty"`
	Subaccount           *int               `json:"subaccount,omitempty"`
}

// CreateQuoteRequest represents a request to create a quote
type CreateQuoteRequest struct {
	RfqID         string            `json:"rfq_id"`
	YesBid        FixedPointDollars `json:"yes_bid"`
	NoBid         FixedPointDollars `json:"no_bid"`
	RestRemainder bool              `json:"rest_remainder"`
	Subaccount    *int              `json:"subaccount,omitempty"`
}

// AcceptQuoteRequest represents a request to accept a quote
type AcceptQuoteRequest struct {
	AcceptedSide OrderSide `json:"accepted_side"`
}

// CreateMarketInMveCollectionRequest represents a request to create a market in MVE collection
type CreateMarketInMveCollectionRequest struct {
	SelectedMarkets   []TickerPair `json:"selected_markets"`
	WithMarketPayload *bool        `json:"with_market_payload,omitempty"`
}

// LookupTickersInMveCollectionRequest represents a request to lookup tickers in MVE collection
type LookupTickersInMveCollectionRequest struct {
	SelectedMarkets []TickerPair `json:"selected_markets"`
}

// GetOrdersParams represents query parameters for getting orders
type GetOrdersParams struct {
	Ticker      string      `url:"ticker,omitempty"`
	EventTicker string      `url:"event_ticker,omitempty"`
	MinTs       int64       `url:"min_ts,omitempty"`
	MaxTs       int64       `url:"max_ts,omitempty"`
	Status      OrderStatus `url:"status,omitempty"`
	Limit       int         `url:"limit,omitempty"`
	Cursor      string      `url:"cursor,omitempty"`
	Subaccount  *int        `url:"subaccount,omitempty"`
}

// GetFillsParams represents query parameters for getting fills
type GetFillsParams struct {
	Ticker     string `url:"ticker,omitempty"`
	OrderID    string `url:"order_id,omitempty"`
	MinTs      int64  `url:"min_ts,omitempty"`
	MaxTs      int64  `url:"max_ts,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Cursor     string `url:"cursor,omitempty"`
	Subaccount *int   `url:"subaccount,omitempty"`
}

// GetPositionsParams represents query parameters for getting positions
type GetPositionsParams struct {
	Cursor      string `url:"cursor,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	CountFilter string `url:"count_filter,omitempty"`
	Ticker      string `url:"ticker,omitempty"`
	EventTicker string `url:"event_ticker,omitempty"`
	Subaccount  *int   `url:"subaccount,omitempty"`
}

// GetSettlementsParams represents query parameters for getting settlements
type GetSettlementsParams struct {
	Limit       int    `url:"limit,omitempty"`
	Cursor      string `url:"cursor,omitempty"`
	Ticker      string `url:"ticker,omitempty"`
	EventTicker string `url:"event_ticker,omitempty"`
	MinTs       int64  `url:"min_ts,omitempty"`
	MaxTs       int64  `url:"max_ts,omitempty"`
	Subaccount  *int   `url:"subaccount,omitempty"`
}

// GetTradesParams represents query parameters for getting trades
type GetTradesParams struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
	Ticker string `url:"ticker,omitempty"`
	MinTs  int64  `url:"min_ts,omitempty"`
	MaxTs  int64  `url:"max_ts,omitempty"`
}

// GetMarketsParams represents query parameters for getting markets
type GetMarketsParams struct {
	Limit        int          `url:"limit,omitempty"`
	Cursor       string       `url:"cursor,omitempty"`
	EventTicker  string       `url:"event_ticker,omitempty"`
	SeriesTicker string       `url:"series_ticker,omitempty"`
	MinCreatedTs int64        `url:"min_created_ts,omitempty"`
	MaxCreatedTs int64        `url:"max_created_ts,omitempty"`
	MinCloseTs   int64        `url:"min_close_ts,omitempty"`
	MaxCloseTs   int64        `url:"max_close_ts,omitempty"`
	MinSettledTs int64        `url:"min_settled_ts,omitempty"`
	MaxSettledTs int64        `url:"max_settled_ts,omitempty"`
	Status       MarketStatus `url:"status,omitempty"`
	Tickers      string       `url:"tickers,omitempty"`
	MveFilter    string       `url:"mve_filter,omitempty"`
	MinUpdatedTs *int64       `url:"min_updated_ts,omitempty"`
}

// GetEventsParams represents query parameters for getting events
type GetEventsParams struct {
	Limit             int         `url:"limit,omitempty"`
	Cursor            string      `url:"cursor,omitempty"`
	WithNestedMarkets bool        `url:"with_nested_markets,omitempty"`
	WithMilestones    bool        `url:"with_milestones,omitempty"`
	Status            EventStatus `url:"status,omitempty"`
	SeriesTicker      string      `url:"series_ticker,omitempty"`
	MinCloseTs        int64       `url:"min_close_ts,omitempty"`
	MinUpdatedTs      *int64      `url:"min_updated_ts,omitempty"`
}

// GetMultivariateEventsParams represents query parameters for getting multivariate events
type GetMultivariateEventsParams struct {
	Limit             int    `url:"limit,omitempty"`
	Cursor            string `url:"cursor,omitempty"`
	SeriesTicker      string `url:"series_ticker,omitempty"`
	CollectionTicker  string `url:"collection_ticker,omitempty"`
	WithNestedMarkets bool   `url:"with_nested_markets,omitempty"`
}

// GetCandlesticksParams represents query parameters for getting candlesticks
type GetCandlesticksParams struct {
	StartTs                  int64 `url:"start_ts"`
	EndTs                    int64 `url:"end_ts"`
	PeriodInterval           int   `url:"period_interval"`
	IncludeLatestBeforeStart bool  `url:"include_latest_before_start,omitempty"`
}

// BatchGetCandlesticksParams represents query parameters for batch getting candlesticks
type BatchGetCandlesticksParams struct {
	MarketTickers            string `url:"market_tickers"`
	StartTs                  int64  `url:"start_ts"`
	EndTs                    int64  `url:"end_ts"`
	PeriodInterval           int    `url:"period_interval"`
	IncludeLatestBeforeStart bool   `url:"include_latest_before_start,omitempty"`
}

// GetSeriesFeeChangesParams represents query parameters for getting series fee changes
type GetSeriesFeeChangesParams struct {
	SeriesTicker   string `url:"series_ticker,omitempty"`
	ShowHistorical bool   `url:"show_historical,omitempty"`
}

// GetSeriesListParams represents query parameters for getting series list
type GetSeriesListParams struct {
	Category               string `url:"category,omitempty"`
	Tags                   string `url:"tags,omitempty"`
	IncludeProductMetadata bool   `url:"include_product_metadata,omitempty"`
	IncludeVolume          bool   `url:"include_volume,omitempty"`
	MinUpdatedTs           *int64 `url:"min_updated_ts,omitempty"`
}

// GetMilestonesParams represents query parameters for getting milestones
type GetMilestonesParams struct {
	Limit              int    `url:"limit"`
	MinimumStartDate   string `url:"minimum_start_date,omitempty"`
	Category           string `url:"category,omitempty"`
	Competition        string `url:"competition,omitempty"`
	SourceID           string `url:"source_id,omitempty"`
	Type               string `url:"type,omitempty"`
	RelatedEventTicker string `url:"related_event_ticker,omitempty"`
	Cursor             string `url:"cursor,omitempty"`
	MinUpdatedTs       *int64 `url:"min_updated_ts,omitempty"`
}

// GetIncentiveProgramsParams represents query parameters for getting incentive programs
type GetIncentiveProgramsParams struct {
	Status string `url:"status,omitempty"`
	Type   string `url:"type,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetStructuredTargetsParams represents query parameters for getting structured targets
type GetStructuredTargetsParams struct {
	Type        string `url:"type,omitempty"`
	Competition string `url:"competition,omitempty"`
	PageSize    int    `url:"page_size,omitempty"`
	Cursor      string `url:"cursor,omitempty"`
}

// GetRFQsParams represents query parameters for getting RFQs
type GetRFQsParams struct {
	Cursor        string `url:"cursor,omitempty"`
	EventTicker   string `url:"event_ticker,omitempty"`
	MarketTicker  string `url:"market_ticker,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Status        string `url:"status,omitempty"`
	CreatorUserID string `url:"creator_user_id,omitempty"`
	Subaccount    *int   `url:"subaccount,omitempty"`
}

// GetQuotesParams represents query parameters for getting quotes
type GetQuotesParams struct {
	Cursor                string `url:"cursor,omitempty"`
	EventTicker           string `url:"event_ticker,omitempty"`
	MarketTicker          string `url:"market_ticker,omitempty"`
	Limit                 int    `url:"limit,omitempty"`
	Status                string `url:"status,omitempty"`
	QuoteCreatorUserID    string `url:"quote_creator_user_id,omitempty"`
	RfqCreatorUserID      string `url:"rfq_creator_user_id,omitempty"`
	RfqCreatorSubtraderID string `url:"rfq_creator_subtrader_id,omitempty"`
	RfqID                 string `url:"rfq_id,omitempty"`
}

// GetOrderQueuePositionsParams represents query parameters for getting order queue positions
type GetOrderQueuePositionsParams struct {
	MarketTickers string `url:"market_tickers,omitempty"`
	EventTicker   string `url:"event_ticker,omitempty"`
	Subaccount    *int   `url:"subaccount,omitempty"`
}

// GetMveCollectionsParams represents query parameters for getting MVE collections
type GetMveCollectionsParams struct {
	Status                string `url:"status,omitempty"`
	AssociatedEventTicker string `url:"associated_event_ticker,omitempty"`
	SeriesTicker          string `url:"series_ticker,omitempty"`
	Limit                 int    `url:"limit,omitempty"`
	Cursor                string `url:"cursor,omitempty"`
}

// GetMveLookupHistoryParams represents query parameters for getting MVE lookup history
type GetMveLookupHistoryParams struct {
	LookbackSeconds int `url:"lookback_seconds"`
}

// GetFCMOrdersParams represents query parameters for getting FCM orders
type GetFCMOrdersParams struct {
	SubtraderID string      `url:"subtrader_id"`
	Cursor      string      `url:"cursor,omitempty"`
	EventTicker string      `url:"event_ticker,omitempty"`
	Ticker      string      `url:"ticker,omitempty"`
	MinTs       int64       `url:"min_ts,omitempty"`
	MaxTs       int64       `url:"max_ts,omitempty"`
	Status      OrderStatus `url:"status,omitempty"`
	Limit       int         `url:"limit,omitempty"`
}

// GetFCMPositionsParams represents query parameters for getting FCM positions
type GetFCMPositionsParams struct {
	SubtraderID      string `url:"subtrader_id"`
	Ticker           string `url:"ticker,omitempty"`
	EventTicker      string `url:"event_ticker,omitempty"`
	CountFilter      string `url:"count_filter,omitempty"`
	SettlementStatus string `url:"settlement_status,omitempty"`
	Limit            int    `url:"limit,omitempty"`
	Cursor           string `url:"cursor,omitempty"`
}

// GetForecastHistoryParams represents query parameters for getting forecast history
type GetForecastHistoryParams struct {
	Percentiles    []int `url:"percentiles"`
	StartTs        int64 `url:"start_ts"`
	EndTs          int64 `url:"end_ts"`
	PeriodInterval int   `url:"period_interval"`
}

// ApplySubaccountTransferRequest represents a request to transfer between subaccounts
type ApplySubaccountTransferRequest struct {
	ClientTransferID string `json:"client_transfer_id"`
	FromSubaccount   int    `json:"from_subaccount"`
	ToSubaccount     int    `json:"to_subaccount"`
	AmountCents      int64  `json:"amount_cents"`
}

// GetSubaccountTransfersParams represents query parameters for getting subaccount transfers
type GetSubaccountTransfersParams struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// GetSeriesParams represents query parameters for getting a single series
type GetSeriesParams struct {
	IncludeVolume bool `url:"include_volume,omitempty"`
}

// GetOrderGroupsParams represents query parameters for getting order groups
type GetOrderGroupsParams struct {
	Subaccount *int `url:"subaccount,omitempty"`
}

// UpdateOrderGroupLimitRequest represents a request to update an order group's contracts limit
type UpdateOrderGroupLimitRequest struct {
	ContractsLimit   *int             `json:"contracts_limit,omitempty"`
	ContractsLimitFp *FixedPointCount `json:"contracts_limit_fp,omitempty"`
}

// UpdateSubaccountNettingRequest represents a request to update subaccount netting settings
type UpdateSubaccountNettingRequest struct {
	SubaccountNumber int  `json:"subaccount_number"`
	Enabled          bool `json:"enabled"`
}

// GetHistoricalMarketsParams represents query parameters for getting historical markets
type GetHistoricalMarketsParams struct {
	Limit               *int   `url:"limit,omitempty"`
	Cursor              string `url:"cursor,omitempty"`
	Tickers             string `url:"tickers,omitempty"`
	EventTicker         string `url:"event_ticker,omitempty"`
	MveHistoricalFilter string `url:"mve_filter,omitempty"`
}

// GetHistoricalCandlesticksParams represents query parameters for getting historical candlesticks
type GetHistoricalCandlesticksParams struct {
	StartTs        int64 `url:"start_ts"`
	EndTs          int64 `url:"end_ts"`
	PeriodInterval int   `url:"period_interval"`
}

// GetHistoricalFillsParams represents query parameters for getting historical fills
type GetHistoricalFillsParams struct {
	Ticker string `url:"ticker,omitempty"`
	MaxTs  *int64 `url:"max_ts,omitempty"`
	Limit  *int   `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetHistoricalOrdersParams represents query parameters for getting historical orders
type GetHistoricalOrdersParams struct {
	Ticker string `url:"ticker,omitempty"`
	MaxTs  *int64 `url:"max_ts,omitempty"`
	Limit  *int   `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetHistoricalTradesParams represents query parameters for getting historical trades
type GetHistoricalTradesParams struct {
	Ticker string `url:"ticker,omitempty"`
	MinTs  *int64 `url:"min_ts,omitempty"`
	MaxTs  *int64 `url:"max_ts,omitempty"`
	Limit  *int   `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}
