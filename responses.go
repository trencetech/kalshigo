package kalshigo

import "time"

// GetExchangeAnnouncementsResponse represents the response for exchange announcements
type GetExchangeAnnouncementsResponse struct {
	Announcements []Announcement `json:"announcements"`
}

// GetExchangeScheduleResponse represents the response for exchange schedule
type GetExchangeScheduleResponse struct {
	Schedule Schedule `json:"schedule"`
}

// GetUserDataTimestampResponse represents the response for user data timestamp
type GetUserDataTimestampResponse struct {
	AsOfTime time.Time `json:"as_of_time"`
}

// GetSeriesFeeChangesResponse represents the response for series fee changes
type GetSeriesFeeChangesResponse struct {
	SeriesFeeChangeArr []SeriesFeeChange `json:"series_fee_change_arr"`
}

// GetApiKeysResponse represents the response for API keys
type GetApiKeysResponse struct {
	ApiKeys []ApiKey `json:"api_keys"`
}

// CreateApiKeyResponse represents the response for creating an API key
type CreateApiKeyResponse struct {
	ApiKeyID string `json:"api_key_id"`
}

// GenerateApiKeyResponse represents the response for generating an API key
type GenerateApiKeyResponse struct {
	ApiKeyID   string `json:"api_key_id"`
	PrivateKey string `json:"private_key"`
}

// GetOrdersResponse represents the response for getting orders
type GetOrdersResponse struct {
	Orders []Order `json:"orders"`
	Cursor string  `json:"cursor"`
}

// GetOrderResponse represents the response for getting a single order
type GetOrderResponse struct {
	Order Order `json:"order"`
}

// CreateOrderResponse represents the response for creating an order
type CreateOrderResponse struct {
	Order Order `json:"order"`
}

// BatchCreateOrdersIndividualResponse represents individual response in batch order creation
type BatchCreateOrdersIndividualResponse struct {
	ClientOrderID *string        `json:"client_order_id,omitempty"`
	Order         *Order         `json:"order,omitempty"`
	Error         *ErrorResponse `json:"error,omitempty"`
}

// BatchCreateOrdersResponse represents the response for batch order creation
type BatchCreateOrdersResponse struct {
	Orders []BatchCreateOrdersIndividualResponse `json:"orders"`
}

// BatchCancelOrdersIndividualResponse represents individual response in batch order cancellation
type BatchCancelOrdersIndividualResponse struct {
	OrderID     string          `json:"order_id"`
	Order       *Order          `json:"order,omitempty"`
	ReducedBy   int             `json:"reduced_by"`
	ReducedByFp FixedPointCount `json:"reduced_by_fp"`
	Error       *ErrorResponse  `json:"error,omitempty"`
}

// BatchCancelOrdersResponse represents the response for batch order cancellation
type BatchCancelOrdersResponse struct {
	Orders []BatchCancelOrdersIndividualResponse `json:"orders"`
}

// CancelOrderResponse represents the response for canceling an order
type CancelOrderResponse struct {
	Order       Order           `json:"order"`
	ReducedBy   int             `json:"reduced_by"`
	ReducedByFp FixedPointCount `json:"reduced_by_fp"`
}

// AmendOrderResponse represents the response for amending an order
type AmendOrderResponse struct {
	OldOrder Order `json:"old_order"`
	Order    Order `json:"order"`
}

// DecreaseOrderResponse represents the response for decreasing an order
type DecreaseOrderResponse struct {
	Order Order `json:"order"`
}

// GetOrderQueuePositionResponse represents the response for getting order queue position
type GetOrderQueuePositionResponse struct {
	QueuePositionFp FixedPointCount `json:"queue_position_fp"`
}

// GetOrderQueuePositionsResponse represents the response for getting order queue positions
type GetOrderQueuePositionsResponse struct {
	QueuePositions []OrderQueuePosition `json:"queue_positions"`
}

// GetOrderGroupsResponse represents the response for getting order groups
type GetOrderGroupsResponse struct {
	OrderGroups []OrderGroup `json:"order_groups,omitempty"`
}

// GetOrderGroupResponse represents the response for getting a single order group
type GetOrderGroupResponse struct {
	IsAutoCancelEnabled bool            `json:"is_auto_cancel_enabled"`
	Orders              []string        `json:"orders,omitempty"`
	ContractsLimitFp    FixedPointCount `json:"contracts_limit_fp,omitempty"`
}

// CreateOrderGroupResponse represents the response for creating an order group
type CreateOrderGroupResponse struct {
	OrderGroupID string `json:"order_group_id"`
}

// GetBalanceResponse represents the response for getting balance
type GetBalanceResponse struct {
	Balance        int64 `json:"balance"`
	PortfolioValue int64 `json:"portfolio_value"`
	UpdatedTs      int64 `json:"updated_ts"`
}

// GetPositionsResponse represents the response for getting positions
type GetPositionsResponse struct {
	Cursor          string           `json:"cursor,omitempty"`
	MarketPositions []MarketPosition `json:"market_positions"`
	EventPositions  []EventPosition  `json:"event_positions"`
}

// GetSettlementsResponse represents the response for getting settlements
type GetSettlementsResponse struct {
	Settlements []Settlement `json:"settlements"`
	Cursor      string       `json:"cursor,omitempty"`
}

// GetPortfolioRestingOrderTotalValueResponse represents the response for total resting order value
type GetPortfolioRestingOrderTotalValueResponse struct {
	TotalRestingOrderValue int `json:"total_resting_order_value"`
}

// GetFillsResponse represents the response for getting fills
type GetFillsResponse struct {
	Fills  []Fill `json:"fills"`
	Cursor string `json:"cursor"`
}

// GetTradesResponse represents the response for getting trades
type GetTradesResponse struct {
	Trades []Trade `json:"trades"`
	Cursor string  `json:"cursor"`
}

// GetMarketsResponse represents the response for getting markets
type GetMarketsResponse struct {
	Markets []Market `json:"markets"`
	Cursor  string   `json:"cursor"`
}

// GetMarketResponse represents the response for getting a single market
type GetMarketResponse struct {
	Market Market `json:"market"`
}

// GetMarketOrderbookResponse represents the response for getting market orderbook
type GetMarketOrderbookResponse struct {
	Orderbook   Orderbook        `json:"orderbook"`
	OrderbookFp OrderbookCountFp `json:"orderbook_fp"`
}

// GetMarketCandlesticksResponse represents the response for getting market candlesticks
type GetMarketCandlesticksResponse struct {
	Ticker       string              `json:"ticker"`
	Candlesticks []MarketCandlestick `json:"candlesticks"`
}

// GetEventCandlesticksResponse represents the response for getting event candlesticks
type GetEventCandlesticksResponse struct {
	MarketTickers      []string              `json:"market_tickers"`
	MarketCandlesticks [][]MarketCandlestick `json:"market_candlesticks"`
	AdjustedEndTs      int64                 `json:"adjusted_end_ts"`
}

// BatchGetMarketCandlesticksResponse represents the response for batch getting candlesticks
type BatchGetMarketCandlesticksResponse struct {
	Markets []MarketCandlesticksResponse `json:"markets"`
}

// GetSeriesResponse represents the response for getting a series
type GetSeriesResponse struct {
	Series Series `json:"series"`
}

// GetSeriesListResponse represents the response for getting series list
type GetSeriesListResponse struct {
	Series []Series `json:"series"`
}

// GetEventsResponse represents the response for getting events
type GetEventsResponse struct {
	Events     []EventData `json:"events"`
	Milestones []Milestone `json:"milestones,omitempty"`
	Cursor     string      `json:"cursor"`
}

// GetMultivariateEventsResponse represents the response for getting multivariate events
type GetMultivariateEventsResponse struct {
	Events []EventData `json:"events"`
	Cursor string      `json:"cursor"`
}

// GetEventResponse represents the response for getting a single event
type GetEventResponse struct {
	Event   EventData `json:"event"`
	Markets []Market  `json:"markets"`
}

// GetEventMetadataResponse represents the response for getting event metadata
type GetEventMetadataResponse struct {
	ImageURL          string             `json:"image_url"`
	FeaturedImageURL  string             `json:"featured_image_url,omitempty"`
	MarketDetails     []MarketMetadata   `json:"market_details"`
	SettlementSources []SettlementSource `json:"settlement_sources"`
	Competition       string             `json:"competition,omitempty"`
	CompetitionScope  string             `json:"competition_scope,omitempty"`
}

// GetEventForecastPercentilesHistoryResponse represents the response for forecast history
type GetEventForecastPercentilesHistoryResponse struct {
	ForecastHistory []ForecastPercentilesPoint `json:"forecast_history"`
}

// GetMilestoneResponse represents the response for getting a milestone
type GetMilestoneResponse struct {
	Milestone Milestone `json:"milestone"`
}

// GetMilestonesResponse represents the response for getting milestones
type GetMilestonesResponse struct {
	Milestones []Milestone `json:"milestones"`
	Cursor     string      `json:"cursor,omitempty"`
}

// GetLiveDataResponse represents the response for getting live data
type GetLiveDataResponse struct {
	LiveData LiveData `json:"live_data"`
}

// GetLiveDatasResponse represents the response for getting multiple live data
type GetLiveDatasResponse struct {
	LiveDatas []LiveData `json:"live_datas"`
}

// GetIncentiveProgramsResponse represents the response for getting incentive programs
type GetIncentiveProgramsResponse struct {
	IncentivePrograms []IncentiveProgram `json:"incentive_programs"`
	NextCursor        string             `json:"next_cursor,omitempty"`
}

// GetStructuredTargetsResponse represents the response for getting structured targets
type GetStructuredTargetsResponse struct {
	StructuredTargets []StructuredTarget `json:"structured_targets,omitempty"`
	Cursor            string             `json:"cursor,omitempty"`
}

// GetStructuredTargetResponse represents the response for getting a structured target
type GetStructuredTargetResponse struct {
	StructuredTarget StructuredTarget `json:"structured_target,omitempty"`
}

// GetTagsForSeriesCategoriesResponse represents the response for tags by categories
type GetTagsForSeriesCategoriesResponse struct {
	TagsByCategories map[string][]string `json:"tags_by_categories"`
}

// GetFiltersBySportsResponse represents the response for filters by sports
type GetFiltersBySportsResponse struct {
	FiltersBySports map[string]SportFilterDetails `json:"filters_by_sports"`
	SportOrdering   []string                      `json:"sport_ordering"`
}

// GetCommunicationsIDResponse represents the response for communications ID
type GetCommunicationsIDResponse struct {
	CommunicationsID string `json:"communications_id"`
}

// GetRFQsResponse represents the response for getting RFQs
type GetRFQsResponse struct {
	RFQs   []RFQ  `json:"rfqs"`
	Cursor string `json:"cursor,omitempty"`
}

// GetRFQResponse represents the response for getting a single RFQ
type GetRFQResponse struct {
	RFQ RFQ `json:"rfq"`
}

// CreateRFQResponse represents the response for creating an RFQ
type CreateRFQResponse struct {
	ID string `json:"id"`
}

// GetQuotesResponse represents the response for getting quotes
type GetQuotesResponse struct {
	Quotes []Quote `json:"quotes"`
	Cursor string  `json:"cursor,omitempty"`
}

// GetQuoteResponse represents the response for getting a single quote
type GetQuoteResponse struct {
	Quote Quote `json:"quote"`
}

// CreateQuoteResponse represents the response for creating a quote
type CreateQuoteResponse struct {
	ID string `json:"id"`
}

// GetMultivariateEventCollectionResponse represents the response for getting MVE collection
type GetMultivariateEventCollectionResponse struct {
	MultivariateContract MultivariateEventCollection `json:"multivariate_contract"`
}

// GetMultivariateEventCollectionsResponse represents the response for getting MVE collections
type GetMultivariateEventCollectionsResponse struct {
	MultivariateContracts []MultivariateEventCollection `json:"multivariate_contracts"`
	Cursor                string                        `json:"cursor,omitempty"`
}

// LookupTickersInMveCollectionResponse represents the response for lookup in MVE collection
type LookupTickersInMveCollectionResponse struct {
	EventTicker  string `json:"event_ticker"`
	MarketTicker string `json:"market_ticker"`
}

// CreateMarketInMveCollectionResponse represents the response for creating market in MVE collection
type CreateMarketInMveCollectionResponse struct {
	EventTicker  string  `json:"event_ticker"`
	MarketTicker string  `json:"market_ticker"`
	Market       *Market `json:"market,omitempty"`
}

// GetMveLookupHistoryResponse represents the response for MVE lookup history
type GetMveLookupHistoryResponse struct {
	LookupPoints []LookupPoint `json:"lookup_points"`
}

// CreateSubaccountResponse represents the response for creating a subaccount
type CreateSubaccountResponse struct {
	SubaccountNumber int `json:"subaccount_number"`
}

// ApplySubaccountTransferResponse represents the response for transferring between subaccounts
type ApplySubaccountTransferResponse struct {
	// Empty response indicating successful transfer
}

// GetSubaccountBalancesResponse represents the response for getting subaccount balances
type GetSubaccountBalancesResponse struct {
	SubaccountBalances []SubaccountBalance `json:"subaccount_balances"`
}

// GetSubaccountTransfersResponse represents the response for getting subaccount transfers
type GetSubaccountTransfersResponse struct {
	Transfers []SubaccountTransfer `json:"transfers"`
	Cursor    *string              `json:"cursor,omitempty"`
}

// GetHistoricalCutoffResponse represents the response for historical cutoff timestamps
type GetHistoricalCutoffResponse struct {
	MarketSettledTs time.Time `json:"market_settled_ts"`
	TradesCreatedTs time.Time `json:"trades_created_ts"`
	OrdersUpdatedTs time.Time `json:"orders_updated_ts"`
}

// GetMarketCandlesticksHistoricalResponse represents historical candlestick response
type GetMarketCandlesticksHistoricalResponse struct {
	Ticker       string                        `json:"ticker"`
	Candlesticks []MarketCandlestickHistorical `json:"candlesticks"`
}

// GetAccountApiLimitsResponse represents the response for account API limits
type GetAccountApiLimitsResponse struct {
	UsageTier  string `json:"usage_tier"`
	ReadLimit  int    `json:"read_limit"`
	WriteLimit int    `json:"write_limit"`
}

// GetSubaccountNettingResponse represents the response for subaccount netting settings
type GetSubaccountNettingResponse struct {
	NettingConfigs []SubaccountNettingConfig `json:"netting_configs"`
}
