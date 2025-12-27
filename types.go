package kalshigo

import "time"

// FixedPointDollars represents a US dollar amount as a fixed-point decimal string
type FixedPointDollars string

// ErrorResponse represents an API error
type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
	Service string `json:"service,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Code != "" {
		return e.Code
	}
	return "unknown error"
}

// SelfTradePreventionType represents the self-trade prevention type for orders
type SelfTradePreventionType string

const (
	SelfTradePreventionTakerAtCross SelfTradePreventionType = "taker_at_cross"
	SelfTradePreventionMaker        SelfTradePreventionType = "maker"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusResting  OrderStatus = "resting"
	OrderStatusCanceled OrderStatus = "canceled"
	OrderStatusExecuted OrderStatus = "executed"
)

// OrderSide represents the side of an order
type OrderSide string

const (
	OrderSideYes OrderSide = "yes"
	OrderSideNo  OrderSide = "no"
)

// OrderAction represents the action of an order
type OrderAction string

const (
	OrderActionBuy  OrderAction = "buy"
	OrderActionSell OrderAction = "sell"
)

// OrderType represents the type of an order
type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

// TimeInForce represents the time in force of an order
type TimeInForce string

const (
	TimeInForceFillOrKill        TimeInForce = "fill_or_kill"
	TimeInForceGoodTillCanceled  TimeInForce = "good_till_canceled"
	TimeInForceImmediateOrCancel TimeInForce = "immediate_or_cancel"
)

// MarketStatus represents the status of a market
type MarketStatus string

const (
	MarketStatusUnopened    MarketStatus = "unopened"
	MarketStatusOpen        MarketStatus = "open"
	MarketStatusPaused      MarketStatus = "paused"
	MarketStatusClosed      MarketStatus = "closed"
	MarketStatusSettled     MarketStatus = "settled"
	MarketStatusInitialized MarketStatus = "initialized"
	MarketStatusInactive    MarketStatus = "inactive"
	MarketStatusActive      MarketStatus = "active"
	MarketStatusDetermined  MarketStatus = "determined"
	MarketStatusDisputed    MarketStatus = "disputed"
	MarketStatusAmended     MarketStatus = "amended"
	MarketStatusFinalized   MarketStatus = "finalized"
)

// MarketType represents the type of a market
type MarketType string

const (
	MarketTypeBinary MarketType = "binary"
	MarketTypeScalar MarketType = "scalar"
)

// MarketResult represents the result of a market
type MarketResult string

const (
	MarketResultYes MarketResult = "yes"
	MarketResultNo  MarketResult = "no"
)

// EventStatus represents the status of an event
type EventStatus string

const (
	EventStatusOpen    EventStatus = "open"
	EventStatusClosed  EventStatus = "closed"
	EventStatusSettled EventStatus = "settled"
)

// AnnouncementType represents the type of an announcement
type AnnouncementType string

const (
	AnnouncementTypeInfo    AnnouncementType = "info"
	AnnouncementTypeWarning AnnouncementType = "warning"
	AnnouncementTypeError   AnnouncementType = "error"
)

// AnnouncementStatus represents the status of an announcement
type AnnouncementStatus string

const (
	AnnouncementStatusActive   AnnouncementStatus = "active"
	AnnouncementStatusInactive AnnouncementStatus = "inactive"
)

// FeeType represents the fee structure type
type FeeType string

const (
	FeeTypeQuadratic              FeeType = "quadratic"
	FeeTypeQuadraticWithMakerFees FeeType = "quadratic_with_maker_fees"
	FeeTypeFlat                   FeeType = "flat"
)

// StrikeType represents how the market strike is defined and evaluated
type StrikeType string

const (
	StrikeTypeGreater        StrikeType = "greater"
	StrikeTypeGreaterOrEqual StrikeType = "greater_or_equal"
	StrikeTypeLess           StrikeType = "less"
	StrikeTypeLessOrEqual    StrikeType = "less_or_equal"
	StrikeTypeBetween        StrikeType = "between"
	StrikeTypeFunctional     StrikeType = "functional"
	StrikeTypeCustom         StrikeType = "custom"
	StrikeTypeStructured     StrikeType = "structured"
)

// TradeTakerSide represents the taker side of a trade
type TradeTakerSide string

const (
	TradeTakerSideYes TradeTakerSide = "yes"
	TradeTakerSideNo  TradeTakerSide = "no"
)

// SettlementResult represents the market settlement result
type SettlementResult string

const (
	SettlementResultYes    SettlementResult = "yes"
	SettlementResultNo     SettlementResult = "no"
	SettlementResultScalar SettlementResult = "scalar"
	SettlementResultVoid   SettlementResult = "void"
)

// IncentiveType represents the type of incentive program
type IncentiveType string

const (
	IncentiveTypeLiquidity IncentiveType = "liquidity"
	IncentiveTypeVolume    IncentiveType = "volume"
)

// QuoteStatus represents the status of a quote
type QuoteStatus string

const (
	QuoteStatusOpen      QuoteStatus = "open"
	QuoteStatusAccepted  QuoteStatus = "accepted"
	QuoteStatusConfirmed QuoteStatus = "confirmed"
	QuoteStatusExecuted  QuoteStatus = "executed"
	QuoteStatusCancelled QuoteStatus = "cancelled"
)

// RFQStatus represents the status of an RFQ
type RFQStatus string

const (
	RFQStatusOpen   RFQStatus = "open"
	RFQStatusClosed RFQStatus = "closed"
)

// ApiKey represents an API key
type ApiKey struct {
	ApiKeyID string   `json:"api_key_id"`
	Name     string   `json:"name"`
	Scopes   []string `json:"scopes"`
}

// ExchangeStatus represents the exchange status
type ExchangeStatus struct {
	ExchangeActive              bool       `json:"exchange_active"`
	TradingActive               bool       `json:"trading_active"`
	ExchangeEstimatedResumeTime *time.Time `json:"exchange_estimated_resume_time,omitempty"`
}

// Announcement represents an exchange announcement
type Announcement struct {
	Type         AnnouncementType   `json:"type"`
	Message      string             `json:"message"`
	DeliveryTime time.Time          `json:"delivery_time"`
	Status       AnnouncementStatus `json:"status"`
}

// DailySchedule represents trading hours for a day
type DailySchedule struct {
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

// MaintenanceWindow represents a scheduled maintenance window
type MaintenanceWindow struct {
	StartDatetime time.Time `json:"start_datetime"`
	EndDatetime   time.Time `json:"end_datetime"`
}

// WeeklySchedule represents the weekly trading schedule
type WeeklySchedule struct {
	StartTime time.Time       `json:"start_time"`
	EndTime   time.Time       `json:"end_time"`
	Monday    []DailySchedule `json:"monday"`
	Tuesday   []DailySchedule `json:"tuesday"`
	Wednesday []DailySchedule `json:"wednesday"`
	Thursday  []DailySchedule `json:"thursday"`
	Friday    []DailySchedule `json:"friday"`
	Saturday  []DailySchedule `json:"saturday"`
	Sunday    []DailySchedule `json:"sunday"`
}

// Schedule represents the exchange schedule
type Schedule struct {
	StandardHours      []WeeklySchedule    `json:"standard_hours"`
	MaintenanceWindows []MaintenanceWindow `json:"maintenance_windows"`
}

// Order represents a trading order
type Order struct {
	OrderID                 string                   `json:"order_id"`
	UserID                  string                   `json:"user_id"`
	ClientOrderID           string                   `json:"client_order_id"`
	Ticker                  string                   `json:"ticker"`
	Side                    OrderSide                `json:"side"`
	Action                  OrderAction              `json:"action"`
	Type                    OrderType                `json:"type"`
	Status                  OrderStatus              `json:"status"`
	YesPrice                int                      `json:"yes_price"`
	NoPrice                 int                      `json:"no_price"`
	YesPriceDollars         FixedPointDollars        `json:"yes_price_dollars"`
	NoPriceDollars          FixedPointDollars        `json:"no_price_dollars"`
	FillCount               int                      `json:"fill_count"`
	RemainingCount          int                      `json:"remaining_count"`
	InitialCount            int                      `json:"initial_count"`
	TakerFees               int                      `json:"taker_fees"`
	MakerFees               int                      `json:"maker_fees"`
	TakerFillCost           int                      `json:"taker_fill_cost"`
	MakerFillCost           int                      `json:"maker_fill_cost"`
	TakerFillCostDollars    FixedPointDollars        `json:"taker_fill_cost_dollars"`
	MakerFillCostDollars    FixedPointDollars        `json:"maker_fill_cost_dollars"`
	QueuePosition           int                      `json:"queue_position"`
	TakerFeesDollars        *FixedPointDollars       `json:"taker_fees_dollars,omitempty"`
	MakerFeesDollars        *FixedPointDollars       `json:"maker_fees_dollars,omitempty"`
	ExpirationTime          *time.Time               `json:"expiration_time,omitempty"`
	CreatedTime             *time.Time               `json:"created_time,omitempty"`
	LastUpdateTime          *time.Time               `json:"last_update_time,omitempty"`
	SelfTradePreventionType *SelfTradePreventionType `json:"self_trade_prevention_type,omitempty"`
	OrderGroupID            *string                  `json:"order_group_id,omitempty"`
	CancelOrderOnPause      bool                     `json:"cancel_order_on_pause,omitempty"`
}

// OrderGroup represents an order group
type OrderGroup struct {
	ID                  string `json:"id"`
	IsAutoCancelEnabled bool   `json:"is_auto_cancel_enabled"`
}

// Trade represents a market trade
type Trade struct {
	TradeID         string            `json:"trade_id"`
	Ticker          string            `json:"ticker"`
	Price           float64           `json:"price"`
	Count           int               `json:"count"`
	YesPrice        int               `json:"yes_price"`
	NoPrice         int               `json:"no_price"`
	YesPriceDollars FixedPointDollars `json:"yes_price_dollars"`
	NoPriceDollars  FixedPointDollars `json:"no_price_dollars"`
	TakerSide       TradeTakerSide    `json:"taker_side"`
	CreatedTime     *time.Time        `json:"created_time,omitempty"`
}

// Fill represents a fill (matched trade)
type Fill struct {
	FillID        string      `json:"fill_id"`
	TradeID       string      `json:"trade_id"`
	OrderID       string      `json:"order_id"`
	ClientOrderID string      `json:"client_order_id,omitempty"`
	Ticker        string      `json:"ticker"`
	MarketTicker  string      `json:"market_ticker"`
	Side          OrderSide   `json:"side"`
	Action        OrderAction `json:"action"`
	Count         int         `json:"count"`
	Price         float64     `json:"price"`
	YesPrice      int         `json:"yes_price"`
	NoPrice       int         `json:"no_price"`
	YesPriceFixed string      `json:"yes_price_fixed"`
	NoPriceFixed  string      `json:"no_price_fixed"`
	IsTaker       bool        `json:"is_taker"`
	CreatedTime   *time.Time  `json:"created_time,omitempty"`
	Ts            int64       `json:"ts,omitempty"`
}

// MarketPosition represents a user's position in a market
type MarketPosition struct {
	Ticker                string            `json:"ticker"`
	TotalTraded           int               `json:"total_traded"`
	TotalTradedDollars    FixedPointDollars `json:"total_traded_dollars"`
	Position              int               `json:"position"`
	MarketExposure        int               `json:"market_exposure"`
	MarketExposureDollars FixedPointDollars `json:"market_exposure_dollars"`
	RealizedPnl           int               `json:"realized_pnl"`
	RealizedPnlDollars    FixedPointDollars `json:"realized_pnl_dollars"`
	RestingOrdersCount    int               `json:"resting_orders_count"`
	FeesPaid              int               `json:"fees_paid"`
	FeesPaidDollars       FixedPointDollars `json:"fees_paid_dollars"`
	LastUpdatedTs         *time.Time        `json:"last_updated_ts,omitempty"`
}

// EventPosition represents a user's position in an event
type EventPosition struct {
	EventTicker          string            `json:"event_ticker"`
	TotalCost            int               `json:"total_cost"`
	TotalCostDollars     FixedPointDollars `json:"total_cost_dollars"`
	TotalCostShares      int64             `json:"total_cost_shares"`
	EventExposure        int               `json:"event_exposure"`
	EventExposureDollars FixedPointDollars `json:"event_exposure_dollars"`
	RealizedPnl          int               `json:"realized_pnl"`
	RealizedPnlDollars   FixedPointDollars `json:"realized_pnl_dollars"`
	FeesPaid             int               `json:"fees_paid"`
	FeesPaidDollars      FixedPointDollars `json:"fees_paid_dollars"`
}

// Settlement represents a market settlement
type Settlement struct {
	Ticker       string           `json:"ticker"`
	EventTicker  string           `json:"event_ticker"`
	MarketResult SettlementResult `json:"market_result"`
	YesCount     int64            `json:"yes_count"`
	YesTotalCost int              `json:"yes_total_cost"`
	NoCount      int64            `json:"no_count"`
	NoTotalCost  int              `json:"no_total_cost"`
	Revenue      int              `json:"revenue"`
	SettledTime  time.Time        `json:"settled_time"`
	FeeCost      string           `json:"fee_cost"`
	Value        *int             `json:"value,omitempty"`
}

// BidAskDistribution represents OHLC data for bid/ask
type BidAskDistribution struct {
	Open         int               `json:"open"`
	OpenDollars  FixedPointDollars `json:"open_dollars"`
	Low          int               `json:"low"`
	LowDollars   FixedPointDollars `json:"low_dollars"`
	High         int               `json:"high"`
	HighDollars  FixedPointDollars `json:"high_dollars"`
	Close        int               `json:"close"`
	CloseDollars FixedPointDollars `json:"close_dollars"`
}

// PriceDistribution represents OHLC data for prices
type PriceDistribution struct {
	Open            *int               `json:"open,omitempty"`
	OpenDollars     *FixedPointDollars `json:"open_dollars,omitempty"`
	Low             *int               `json:"low,omitempty"`
	LowDollars      *FixedPointDollars `json:"low_dollars,omitempty"`
	High            *int               `json:"high,omitempty"`
	HighDollars     *FixedPointDollars `json:"high_dollars,omitempty"`
	Close           *int               `json:"close,omitempty"`
	CloseDollars    *FixedPointDollars `json:"close_dollars,omitempty"`
	Mean            *int               `json:"mean,omitempty"`
	MeanDollars     *FixedPointDollars `json:"mean_dollars,omitempty"`
	Previous        *int               `json:"previous,omitempty"`
	PreviousDollars *FixedPointDollars `json:"previous_dollars,omitempty"`
	Min             *int               `json:"min,omitempty"`
	MinDollars      *FixedPointDollars `json:"min_dollars,omitempty"`
	Max             *int               `json:"max,omitempty"`
	MaxDollars      *FixedPointDollars `json:"max_dollars,omitempty"`
}

// MarketCandlestick represents candlestick data for a market
type MarketCandlestick struct {
	EndPeriodTs  int64              `json:"end_period_ts"`
	YesBid       BidAskDistribution `json:"yes_bid"`
	YesAsk       BidAskDistribution `json:"yes_ask"`
	Price        PriceDistribution  `json:"price"`
	Volume       int64              `json:"volume"`
	OpenInterest int64              `json:"open_interest"`
}

// PriceLevel represents a price level in the order book
type PriceLevel [2]int

// PriceLevelDollars represents a price level in dollars
type PriceLevelDollars [2]any

// Orderbook represents the order book
type Orderbook struct {
	Yes        []PriceLevel        `json:"yes"`
	No         []PriceLevel        `json:"no"`
	YesDollars []PriceLevelDollars `json:"yes_dollars"`
	NoDollars  []PriceLevelDollars `json:"no_dollars"`
}

// PriceRange represents a valid price range for orders
type PriceRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Step  string `json:"step"`
}

// MveSelectedLeg represents a selected leg for MVE
type MveSelectedLeg struct {
	EventTicker  string `json:"event_ticker,omitempty"`
	MarketTicker string `json:"market_ticker,omitempty"`
	Side         string `json:"side,omitempty"`
}

// Market represents a market
type Market struct {
	Ticker                  string             `json:"ticker"`
	EventTicker             string             `json:"event_ticker"`
	MarketType              MarketType         `json:"market_type"`
	Title                   string             `json:"title"`
	Subtitle                string             `json:"subtitle"`
	YesSubTitle             string             `json:"yes_sub_title"`
	NoSubTitle              string             `json:"no_sub_title"`
	CreatedTime             time.Time          `json:"created_time"`
	OpenTime                time.Time          `json:"open_time"`
	CloseTime               time.Time          `json:"close_time"`
	ExpectedExpirationTime  *time.Time         `json:"expected_expiration_time,omitempty"`
	ExpirationTime          time.Time          `json:"expiration_time"`
	LatestExpirationTime    time.Time          `json:"latest_expiration_time"`
	SettlementTimerSeconds  int                `json:"settlement_timer_seconds"`
	Status                  MarketStatus       `json:"status"`
	ResponsePriceUnits      string             `json:"response_price_units"`
	YesBid                  float64            `json:"yes_bid"`
	YesBidDollars           FixedPointDollars  `json:"yes_bid_dollars"`
	YesAsk                  float64            `json:"yes_ask"`
	YesAskDollars           FixedPointDollars  `json:"yes_ask_dollars"`
	NoBid                   float64            `json:"no_bid"`
	NoBidDollars            FixedPointDollars  `json:"no_bid_dollars"`
	NoAsk                   float64            `json:"no_ask"`
	NoAskDollars            FixedPointDollars  `json:"no_ask_dollars"`
	LastPrice               float64            `json:"last_price"`
	LastPriceDollars        FixedPointDollars  `json:"last_price_dollars"`
	Volume                  int                `json:"volume"`
	Volume24h               int                `json:"volume_24h"`
	Result                  MarketResult       `json:"result,omitempty"`
	CanCloseEarly           bool               `json:"can_close_early"`
	OpenInterest            int                `json:"open_interest"`
	NotionalValue           int                `json:"notional_value"`
	NotionalValueDollars    FixedPointDollars  `json:"notional_value_dollars"`
	PreviousYesBid          int                `json:"previous_yes_bid"`
	PreviousYesBidDollars   FixedPointDollars  `json:"previous_yes_bid_dollars"`
	PreviousYesAsk          int                `json:"previous_yes_ask"`
	PreviousYesAskDollars   FixedPointDollars  `json:"previous_yes_ask_dollars"`
	PreviousPrice           int                `json:"previous_price"`
	PreviousPriceDollars    FixedPointDollars  `json:"previous_price_dollars"`
	Liquidity               int                `json:"liquidity"`
	LiquidityDollars        FixedPointDollars  `json:"liquidity_dollars"`
	SettlementValue         *int               `json:"settlement_value,omitempty"`
	SettlementValueDollars  *FixedPointDollars `json:"settlement_value_dollars,omitempty"`
	SettlementTs            *time.Time         `json:"settlement_ts,omitempty"`
	ExpirationValue         string             `json:"expiration_value"`
	Category                string             `json:"category"`
	RiskLimitCents          int                `json:"risk_limit_cents"`
	FeeWaiverExpirationTime *time.Time         `json:"fee_waiver_expiration_time,omitempty"`
	EarlyCloseCondition     *string            `json:"early_close_condition,omitempty"`
	TickSize                int                `json:"tick_size"`
	StrikeType              StrikeType         `json:"strike_type,omitempty"`
	FloorStrike             *float64           `json:"floor_strike,omitempty"`
	CapStrike               *float64           `json:"cap_strike,omitempty"`
	FunctionalStrike        *string            `json:"functional_strike,omitempty"`
	CustomStrike            map[string]any     `json:"custom_strike,omitempty"`
	RulesPrimary            string             `json:"rules_primary"`
	RulesSecondary          string             `json:"rules_secondary"`
	MveCollectionTicker     string             `json:"mve_collection_ticker,omitempty"`
	MveSelectedLegs         []MveSelectedLeg   `json:"mve_selected_legs,omitempty"`
	PrimaryParticipantKey   *string            `json:"primary_participant_key,omitempty"`
	PriceLevelStructure     string             `json:"price_level_structure"`
	PriceRanges             []PriceRange       `json:"price_ranges"`
}

// SettlementSource represents a settlement source
type SettlementSource struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Series represents a series
type Series struct {
	Ticker                 string             `json:"ticker"`
	Frequency              string             `json:"frequency"`
	Title                  string             `json:"title"`
	Category               string             `json:"category"`
	Tags                   []string           `json:"tags"`
	SettlementSources      []SettlementSource `json:"settlement_sources"`
	ContractURL            string             `json:"contract_url"`
	ContractTermsURL       string             `json:"contract_terms_url"`
	ProductMetadata        map[string]any     `json:"product_metadata,omitempty"`
	FeeType                FeeType            `json:"fee_type"`
	FeeMultiplier          float64            `json:"fee_multiplier"`
	AdditionalProhibitions []string           `json:"additional_prohibitions"`
}

// SeriesFeeChange represents a fee change for a series
type SeriesFeeChange struct {
	ID            string    `json:"id"`
	SeriesTicker  string    `json:"series_ticker"`
	FeeType       FeeType   `json:"fee_type"`
	FeeMultiplier float64   `json:"fee_multiplier"`
	ScheduledTs   time.Time `json:"scheduled_ts"`
}

// EventData represents an event
type EventData struct {
	EventTicker          string         `json:"event_ticker"`
	SeriesTicker         string         `json:"series_ticker"`
	SubTitle             string         `json:"sub_title"`
	Title                string         `json:"title"`
	CollateralReturnType string         `json:"collateral_return_type"`
	MutuallyExclusive    bool           `json:"mutually_exclusive"`
	Category             string         `json:"category"`
	StrikeDate           *time.Time     `json:"strike_date,omitempty"`
	StrikePeriod         *string        `json:"strike_period,omitempty"`
	Markets              []Market       `json:"markets,omitempty"`
	AvailableOnBrokers   bool           `json:"available_on_brokers"`
	ProductMetadata      map[string]any `json:"product_metadata,omitempty"`
}

// MarketMetadata represents market metadata
type MarketMetadata struct {
	MarketTicker string `json:"market_ticker"`
	ImageURL     string `json:"image_url"`
	ColorCode    string `json:"color_code"`
}

// Milestone represents a milestone
type Milestone struct {
	ID                  string         `json:"id"`
	Category            string         `json:"category"`
	Type                string         `json:"type"`
	StartDate           time.Time      `json:"start_date"`
	EndDate             *time.Time     `json:"end_date,omitempty"`
	RelatedEventTickers []string       `json:"related_event_tickers"`
	Title               string         `json:"title"`
	NotificationMessage string         `json:"notification_message"`
	SourceID            *string        `json:"source_id,omitempty"`
	Details             map[string]any `json:"details"`
	PrimaryEventTickers []string       `json:"primary_event_tickers"`
	LastUpdatedTs       time.Time      `json:"last_updated_ts"`
}

// LiveData represents live data
type LiveData struct {
	Type        string         `json:"type"`
	Details     map[string]any `json:"details"`
	MilestoneID string         `json:"milestone_id"`
}

// IncentiveProgram represents an incentive program
type IncentiveProgram struct {
	ID                string        `json:"id"`
	MarketTicker      string        `json:"market_ticker"`
	IncentiveType     IncentiveType `json:"incentive_type"`
	StartDate         time.Time     `json:"start_date"`
	EndDate           time.Time     `json:"end_date"`
	PeriodReward      int64         `json:"period_reward"`
	PaidOut           bool          `json:"paid_out"`
	DiscountFactorBps *int          `json:"discount_factor_bps,omitempty"`
	TargetSize        *int          `json:"target_size,omitempty"`
}

// StructuredTarget represents a structured target
type StructuredTarget struct {
	ID            string         `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Type          string         `json:"type,omitempty"`
	Details       map[string]any `json:"details,omitempty"`
	SourceID      string         `json:"source_id,omitempty"`
	LastUpdatedTs *time.Time     `json:"last_updated_ts,omitempty"`
}

// RFQ represents a request for quote
type RFQ struct {
	ID                   string           `json:"id"`
	CreatorID            string           `json:"creator_id"`
	MarketTicker         string           `json:"market_ticker"`
	Contracts            int              `json:"contracts"`
	TargetCostCentiCents int64            `json:"target_cost_centi_cents,omitempty"`
	Status               RFQStatus        `json:"status"`
	CreatedTs            time.Time        `json:"created_ts"`
	MveCollectionTicker  string           `json:"mve_collection_ticker,omitempty"`
	MveSelectedLegs      []MveSelectedLeg `json:"mve_selected_legs,omitempty"`
	RestRemainder        bool             `json:"rest_remainder,omitempty"`
	CancellationReason   string           `json:"cancellation_reason,omitempty"`
	CreatorUserID        string           `json:"creator_user_id,omitempty"`
	CancelledTs          *time.Time       `json:"cancelled_ts,omitempty"`
	UpdatedTs            *time.Time       `json:"updated_ts,omitempty"`
}

// Quote represents a quote response to an RFQ
type Quote struct {
	ID                      string            `json:"id"`
	RfqID                   string            `json:"rfq_id"`
	CreatorID               string            `json:"creator_id"`
	RfqCreatorID            string            `json:"rfq_creator_id,omitempty"`
	MarketTicker            string            `json:"market_ticker"`
	Contracts               int               `json:"contracts"`
	YesBid                  int               `json:"yes_bid"`
	NoBid                   int               `json:"no_bid"`
	YesBidDollars           FixedPointDollars `json:"yes_bid_dollars"`
	NoBidDollars            FixedPointDollars `json:"no_bid_dollars"`
	CreatedTs               time.Time         `json:"created_ts"`
	UpdatedTs               time.Time         `json:"updated_ts"`
	Status                  QuoteStatus       `json:"status"`
	AcceptedSide            OrderSide         `json:"accepted_side,omitempty"`
	AcceptedTs              *time.Time        `json:"accepted_ts,omitempty"`
	ConfirmedTs             *time.Time        `json:"confirmed_ts,omitempty"`
	ExecutedTs              *time.Time        `json:"executed_ts,omitempty"`
	CancelledTs             *time.Time        `json:"cancelled_ts,omitempty"`
	RestRemainder           bool              `json:"rest_remainder,omitempty"`
	CancellationReason      string            `json:"cancellation_reason,omitempty"`
	CreatorUserID           string            `json:"creator_user_id,omitempty"`
	RfqCreatorUserID        string            `json:"rfq_creator_user_id,omitempty"`
	RfqTargetCostCentiCents int64             `json:"rfq_target_cost_centi_cents,omitempty"`
	RfqCreatorOrderID       string            `json:"rfq_creator_order_id,omitempty"`
	CreatorOrderID          string            `json:"creator_order_id,omitempty"`
}

// OrderQueuePosition represents the queue position of an order
type OrderQueuePosition struct {
	OrderID       string `json:"order_id"`
	MarketTicker  string `json:"market_ticker"`
	QueuePosition int    `json:"queue_position"`
}

// AssociatedEvent represents an event associated with a collection
type AssociatedEvent struct {
	Ticker        string   `json:"ticker"`
	IsYesOnly     bool     `json:"is_yes_only"`
	SizeMax       *int     `json:"size_max,omitempty"`
	SizeMin       *int     `json:"size_min,omitempty"`
	ActiveQuoters []string `json:"active_quoters"`
}

// MultivariateEventCollection represents a multivariate event collection
type MultivariateEventCollection struct {
	CollectionTicker       string            `json:"collection_ticker"`
	SeriesTicker           string            `json:"series_ticker"`
	Title                  string            `json:"title"`
	Description            string            `json:"description"`
	OpenDate               time.Time         `json:"open_date"`
	CloseDate              time.Time         `json:"close_date"`
	AssociatedEvents       []AssociatedEvent `json:"associated_events"`
	AssociatedEventTickers []string          `json:"associated_event_tickers"`
	IsOrdered              bool              `json:"is_ordered"`
	IsSingleMarketPerEvent bool              `json:"is_single_market_per_event"`
	IsAllYes               bool              `json:"is_all_yes"`
	SizeMin                int               `json:"size_min"`
	SizeMax                int               `json:"size_max"`
	FunctionalDescription  string            `json:"functional_description"`
}

// TickerPair represents a ticker pair for MVE operations
type TickerPair struct {
	MarketTicker string    `json:"market_ticker"`
	EventTicker  string    `json:"event_ticker"`
	Side         OrderSide `json:"side"`
}

// LookupPoint represents a lookup point in MVE collection
type LookupPoint struct {
	EventTicker     string       `json:"event_ticker"`
	MarketTicker    string       `json:"market_ticker"`
	SelectedMarkets []TickerPair `json:"selected_markets"`
	LastQueriedTs   time.Time    `json:"last_queried_ts"`
}

// SportFilterDetails represents filter details for a sport
type SportFilterDetails struct {
	Scopes       []string             `json:"scopes"`
	Competitions map[string]ScopeList `json:"competitions"`
}

// ScopeList represents a list of scopes
type ScopeList struct {
	Scopes []string `json:"scopes"`
}

// PercentilePoint represents a percentile point for forecasts
type PercentilePoint struct {
	Percentile           int     `json:"percentile"`
	RawNumericalForecast float64 `json:"raw_numerical_forecast"`
	NumericalForecast    float64 `json:"numerical_forecast"`
	FormattedForecast    string  `json:"formatted_forecast"`
}

// ForecastPercentilesPoint represents forecast percentile data over time
type ForecastPercentilesPoint struct {
	EventTicker      string            `json:"event_ticker"`
	EndPeriodTs      int64             `json:"end_period_ts"`
	PeriodInterval   int               `json:"period_interval"`
	PercentilePoints []PercentilePoint `json:"percentile_points"`
}

// MarketCandlesticksResponse represents candlestick data for a market
type MarketCandlesticksResponse struct {
	MarketTicker string              `json:"market_ticker"`
	Candlesticks []MarketCandlestick `json:"candlesticks"`
}
