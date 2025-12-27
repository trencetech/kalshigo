# Kalshigo

A Go client library for the [Kalshi](https://kalshi.com) prediction markets trading API.

## Installation

```bash
go get github.com/trencetech/kalshigo
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/trencetech/kalshigo"
)

func main() {
    // Create a client with API credentials
    client := kalshigo.NewClient(kalshigo.ClientConfig{
        APIKeyID:   "your-api-key-id",
        PrivateKey: "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----",
    })

    ctx := context.Background()

    // Get your portfolio balance
    balance, err := client.GetBalance(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Balance: %s\n", balance.Balance)

    // Fetch available markets
    markets, err := client.GetMarkets(ctx, kalshigo.GetMarketsParams{
        Status: kalshigo.MarketStatusOpen,
        Limit:  10,
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, market := range markets.Markets {
        fmt.Printf("Market: %s - %s\n", market.Ticker, market.Title)
    }
}
```

## Features

- Complete coverage of the Kalshi Trade API v2
- RSA-based request signing for authentication
- Support for both production and demo environments
- Context-aware requests with cancellation support
- Structured logging with zap
- Fixed-point dollar representation for financial precision

## Configuration

### Client Options

```go
client := kalshigo.NewClient(kalshigo.ClientConfig{
    BaseURL:    kalshigo.DefaultBaseURL,  // optional, defaults to production API
    APIKeyID:   "your-api-key-id",
    PrivateKey: "your-pem-encoded-private-key",
    HTTPClient: &http.Client{Timeout: 30 * time.Second},
    Logger:     *zap.NewProduction(),
    UserAgent:  "my-trading-bot/1.0",
})
```

### API URL Configuration

The `BaseURL` is configurable and defaults to the production API if not specified:

```go
// Production (default) - no need to specify BaseURL
client := kalshigo.NewClient(kalshigo.ClientConfig{
    APIKeyID:   "your-api-key-id",
    PrivateKey: "your-pem-encoded-private-key",
})

// Demo environment
client := kalshigo.NewClient(kalshigo.ClientConfig{
    BaseURL:    kalshigo.DemoBaseURL,
    APIKeyID:   "your-api-key-id",
    PrivateKey: "your-pem-encoded-private-key",
})

// Custom API URL
client := kalshigo.NewClient(kalshigo.ClientConfig{
    BaseURL:    "https://custom-api.example.com/trade-api/v2",
    APIKeyID:   "your-api-key-id",
    PrivateKey: "your-pem-encoded-private-key",
})
```

### Environment URLs

| Constant | Environment | URL |
|----------|-------------|-----|
| `DefaultBaseURL` | Production | `https://api.elections.kalshi.com/trade-api/v2` |
| `DemoBaseURL` | Demo | `https://demo-api.kalshi.co/trade-api/v2` |

## API Coverage

### Portfolio & Account

```go
GetBalance(ctx)
GetPositions(ctx, params)
GetSettlements(ctx, params)
GetFills(ctx, params)
GetPortfolioRestingOrderTotalValue(ctx)
```

### Orders

```go
GetOrders(ctx, params)
GetOrder(ctx, orderID)
CreateOrder(ctx, request)
CancelOrder(ctx, orderID)
AmendOrder(ctx, orderID, request)
DecreaseOrder(ctx, orderID, request)
BatchCreateOrders(ctx, request)
BatchCancelOrders(ctx, request)
GetOrderQueuePositions(ctx, ticker)
```

### Markets

```go
GetMarkets(ctx, params)
GetMarket(ctx, ticker)
GetMarketOrderbook(ctx, ticker, depth)
GetMarketCandlesticks(ctx, seriesTicker, ticker, params)
BatchGetMarketCandlesticks(ctx, seriesTicker, params)
GetTrades(ctx, params)
```

### Events

```go
GetEvents(ctx, params)
GetEvent(ctx, eventTicker)
GetEventMetadata(ctx, eventTicker)
GetEventCandlesticks(ctx, seriesTicker, eventTicker, params)
GetEventForecastPercentilesHistory(ctx, seriesTicker, eventTicker, params)
GetMultivariateEvents(ctx, params)
```

### Exchange

```go
GetExchangeStatus(ctx)
GetExchangeAnnouncements(ctx)
GetExchangeSchedule(ctx)
GetSeriesFeeChanges(ctx, params)
```

### RFQ & Quotes

```go
CreateRFQ(ctx, request)
GetRFQs(ctx, params)
GetRFQ(ctx, rfqID)
DeleteRFQ(ctx, rfqID)
CreateQuote(ctx, request)
GetQuotes(ctx, params)
GetQuote(ctx, quoteID)
DeleteQuote(ctx, quoteID)
AcceptQuote(ctx, quoteID)
ConfirmQuote(ctx, quoteID, request)
```

### Multivariate Events

```go
GetMultivariateEventCollections(ctx, params)
GetMultivariateEventCollection(ctx, mveCollectionTicker)
CreateMarketInMultivariateEventCollection(ctx, ticker, request)
LookupTickersInMultivariateEventCollection(ctx, ticker, request)
GetMultivariateEventCollectionLookupHistory(ctx, ticker, params)
```

### API Keys

```go
GetApiKeys(ctx)
CreateApiKey(ctx, request)
GenerateApiKey(ctx, request)
DeleteApiKey(ctx, keyID)
```

## Order Types

```go
// Create a limit order
order, err := client.CreateOrder(ctx, kalshigo.CreateOrderRequest{
    Ticker:     "INXD-25MAR21-B4300",
    Action:     kalshigo.OrderActionBuy,
    Side:       kalshigo.OrderSideYes,
    Type:       kalshigo.OrderTypeLimit,
    Count:      10,
    YesPrice:   50,  // Price in cents
    Expiration: time.Now().Add(24 * time.Hour),
})

// Create a market order
order, err := client.CreateOrder(ctx, kalshigo.CreateOrderRequest{
    Ticker: "INXD-25MAR21-B4300",
    Action: kalshigo.OrderActionBuy,
    Side:   kalshigo.OrderSideYes,
    Type:   kalshigo.OrderTypeMarket,
    Count:  5,
})
```

## Authentication

Kalshi uses RSA-based request signing. You need:

1. An API Key ID from your Kalshi account settings
2. An RSA private key (PEM format) associated with that API key

The client automatically signs all authenticated requests with:
- `KALSHI-ACCESS-KEY`: Your API key ID
- `KALSHI-ACCESS-SIGNATURE`: RSA-PSS signature of the request
- `KALSHI-ACCESS-TIMESTAMP`: Current timestamp in milliseconds

## Constants

The library provides typed constants for common values:

```go
// Order status
kalshigo.OrderStatusResting
kalshigo.OrderStatusCanceled
kalshigo.OrderStatusExecuted

// Order sides
kalshigo.OrderSideYes
kalshigo.OrderSideNo

// Order actions
kalshigo.OrderActionBuy
kalshigo.OrderActionSell

// Order types
kalshigo.OrderTypeLimit
kalshigo.OrderTypeMarket

// Market status
kalshigo.MarketStatusOpen
kalshigo.MarketStatusClosed
kalshigo.MarketStatusSettled

// Market types
kalshigo.MarketTypeBinary
kalshigo.MarketTypeScalar
```

## Error Handling

API errors are returned as `*kalshigo.ErrorResponse`:

```go
order, err := client.CreateOrder(ctx, request)
if err != nil {
    if apiErr, ok := err.(*kalshigo.ErrorResponse); ok {
        fmt.Printf("API Error: %s (code: %s)\n", apiErr.Message, apiErr.Code)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Dependencies

- [go.uber.org/zap](https://github.com/uber-go/zap) - Structured logging

## Requirements

- Go 1.21 or later

## License

See [LICENSE](LICENSE) for details.