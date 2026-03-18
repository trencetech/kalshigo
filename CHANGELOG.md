# Changelog

## v1.0.2 - 2026-03-18

### New Endpoints

- **Historical Data** (`historical.go`)
  - `GetHistoricalCutoff` - Get cutoff timestamps between live and historical data
  - `GetHistoricalMarkets` - Retrieve archived markets
  - `GetHistoricalMarket` - Retrieve a specific archived market by ticker
  - `GetHistoricalMarketCandlesticks` - Get candlestick data for archived markets
  - `GetHistoricalFills` - Get historical fills for authenticated user
  - `GetHistoricalOrders` - Get archived orders
  - `GetHistoricalTrades` - Get historical trades

- **Account** (`account.go`)
  - `GetAccountApiLimits` - Retrieve API tier limits for authenticated user

- **Order Groups** (`orders.go`)
  - `TriggerOrderGroup` - Cancel all orders in a group
  - `UpdateOrderGroupLimit` - Update rolling 15-second contracts limit

- **Subaccount Netting** (`portfolio.go`)
  - `GetSubaccountNetting` - Get netting settings for all subaccounts
  - `UpdateSubaccountNetting` - Update netting setting for a subaccount

### Updated Method Signatures

- `GetSeries` now accepts `*GetSeriesParams` to support `include_volume` query parameter

### New Types

- `ExchangeInstance` - Exchange instance type enum (`event_contract`, `margined`)
- `BidAskDistributionHistorical` - Historical OHLC bid/ask data using dollar strings
- `PriceDistributionHistorical` - Historical OHLC price data using dollar strings
- `MarketCandlestickHistorical` - Historical candlestick data
- `SubaccountNettingConfig` - Subaccount netting configuration
- `BatchCancelOrdersRequestOrder` - Per-order subaccount support in batch cancel
- `GetHistoricalMarketsParams`, `GetHistoricalCandlesticksParams`, `GetHistoricalFillsParams`, `GetHistoricalOrdersParams`, `GetHistoricalTradesParams` - Query params for historical endpoints
- `GetSeriesParams` - Query params for single series retrieval
- `GetOrderGroupsParams` - Query params for order groups
- `UpdateOrderGroupLimitRequest` - Request body for updating order group limit
- `UpdateSubaccountNettingRequest` - Request body for updating netting settings

### New Response Types

- `GetHistoricalCutoffResponse`
- `GetMarketCandlesticksHistoricalResponse`
- `GetAccountApiLimitsResponse`
- `GetSubaccountNettingResponse`

### Updated Types

- **`Market`** - Added `UpdatedTime`, `YesBidSizeFp`, `YesAskSizeFp`, `FractionalTradingEnabled`
- **`Order`** - Added `SubaccountNumber`
- **`OrderGroup`** - Added `ContractsLimitFp`
- **`Fill`** - Added `YesPriceDollars`, `NoPriceDollars`, `FeeCost` (FixedPointDollars), `SubaccountNumber`
- **`Settlement`** - Added `YesTotalCostDollars`, `NoTotalCostDollars`; changed `FeeCost` to `FixedPointDollars`
- **`EventData`** - Added `LastUpdatedTs`
- **`Series`** - Added `LastUpdatedTs`
- **`Milestone`** - Added `SourceIDs`
- **`StructuredTarget`** - Added `SourceIDs`
- **`IncentiveProgram`** - Added `MarketID`
- **`RFQ`** - Added `TargetCostDollars`
- **`Quote`** - Added `RfqTargetCostDollars`, `YesContractsFp`, `NoContractsFp`
- **`OrderQueuePosition`** - Changed `QueuePosition` to `QueuePositionFp` (FixedPointCount)
- **`SubaccountBalance`** - Changed `Balance` to `FixedPointDollars`
- **`EventStatus`** - Added `EventStatusUnopened`

### Updated Request Types

- **`BatchCancelOrdersRequest`** - Added `Orders` field with per-order subaccount support
- **`CreateOrderGroupRequest`** - Added `Subaccount`
- **`CreateRFQRequest`** - Added `ContractsFp`, `TargetCostDollars`, `ReplaceExisting`, `SubtraderID`, `Subaccount`
- **`CreateQuoteRequest`** - Added `Subaccount`
- **`AmendOrderRequest`** - Added `Ticker`, `Side`, `Action`, `ClientOrderID`, `UpdatedClientOrderID`, `YesPriceDollars`, `NoPriceDollars`, `CountFp`, `Subaccount`
- **`DecreaseOrderRequest`** - Added `ReduceByFp`, `ReduceToFp`, `Subaccount`
- **`GetMarketsParams`** - Added `MinUpdatedTs`, `MinSettledTs`, `MaxSettledTs`
- **`GetEventsParams`** - Added `MinUpdatedTs`
- **`GetMilestonesParams`** - Added `MinUpdatedTs`
- **`GetSeriesListParams`** - Added `IncludeVolume`, `MinUpdatedTs`
- **`GetCandlesticksParams`** - Added `IncludeLatestBeforeStart`
- **`GetRFQsParams`** - Added `Subaccount`
- **`GetOrderQueuePositionsParams`** - Added `Subaccount`

### Updated Response Types

- **`GetOrderGroupResponse`** - Added `ContractsLimitFp`
- **`GetOrderQueuePositionResponse`** - Changed to `QueuePositionFp` (FixedPointCount)