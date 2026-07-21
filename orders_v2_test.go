package kalshigo

import (
	"encoding/json"
	"testing"
)

func TestCreateEventOrderRequest_Marshal(t *testing.T) {
	req := CreateEventOrderRequest{
		Ticker:                  "KXTEST-25JAN01",
		Side:                    EventOrderSideBid,
		Count:                   "10.00",
		Price:                   "0.5600",
		TimeInForce:             TimeInForceFillOrKill,
		SelfTradePreventionType: SelfTradePreventionTakerAtCross,
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	// Required fields must be present with the exact V2 wire names.
	wantRequired := map[string]any{
		"ticker":                     "KXTEST-25JAN01",
		"side":                       "bid",
		"count":                      "10.00",
		"price":                      "0.5600",
		"time_in_force":              "fill_or_kill",
		"self_trade_prevention_type": "taker_at_cross",
	}
	for k, want := range wantRequired {
		v, ok := got[k]
		if !ok {
			t.Errorf("missing required field %q in %s", k, b)
			continue
		}
		if v != want {
			t.Errorf("field %q = %v, want %v", k, v, want)
		}
	}

	// Legacy fields must NOT appear on the V2 request.
	for _, k := range []string{"action", "type", "count_fp", "yes_price", "no_price", "yes_price_dollars", "no_price_dollars"} {
		if _, ok := got[k]; ok {
			t.Errorf("unexpected legacy field %q on V2 request: %s", k, b)
		}
	}

	// Unset optional fields must be omitted.
	for _, k := range []string{"client_order_id", "expiration_time", "post_only", "cancel_order_on_pause", "reduce_only", "subaccount", "order_group_id", "exchange_index"} {
		if _, ok := got[k]; ok {
			t.Errorf("optional field %q should be omitted when unset: %s", k, b)
		}
	}
}

func TestAmendEventOrderRequest_Marshal(t *testing.T) {
	req := AmendEventOrderRequest{
		Ticker: "KXTEST-25JAN01",
		Side:   EventOrderSideAsk,
		Price:  "0.4200",
		Count:  "8.00",
	}

	got := marshalToMap(t, req)

	for k, want := range map[string]any{
		"ticker": "KXTEST-25JAN01",
		"side":   "ask",
		"price":  "0.4200",
		"count":  "8.00",
	} {
		if got[k] != want {
			t.Errorf("field %q = %v, want %v", k, got[k], want)
		}
	}
	for _, k := range []string{"action", "yes_price", "no_price", "count_fp", "subaccount"} {
		if _, ok := got[k]; ok {
			t.Errorf("unexpected field %q on V2 amend request: %v", k, got)
		}
	}
}

func TestDecreaseEventOrderRequest_ExactlyOne(t *testing.T) {
	reduceBy := FixedPointCount("5.00")
	req := DecreaseEventOrderRequest{ReduceBy: &reduceBy}

	got := marshalToMap(t, req)

	if got["reduce_by"] != "5.00" {
		t.Errorf("reduce_by = %v, want 5.00", got["reduce_by"])
	}
	// The unset alternative must be omitted so the "exactly one" constraint holds.
	if _, ok := got["reduce_to"]; ok {
		t.Errorf("reduce_to should be omitted when unset: %v", got)
	}
	// Ensure the legacy _fp names are not emitted.
	for _, k := range []string{"reduce_by_fp", "reduce_to_fp"} {
		if _, ok := got[k]; ok {
			t.Errorf("unexpected legacy field %q: %v", k, got)
		}
	}
}

func TestBatchCreateEventOrdersRequest_Marshal(t *testing.T) {
	req := BatchCreateEventOrdersRequest{
		Orders: []CreateEventOrderRequest{{
			Ticker:                  "KXTEST-25JAN01",
			Side:                    EventOrderSideBid,
			Count:                   "1.00",
			Price:                   "0.5000",
			TimeInForce:             TimeInForceImmediateOrCancel,
			SelfTradePreventionType: SelfTradePreventionMaker,
		}},
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got struct {
		Orders []map[string]any `json:"orders"`
	}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(got.Orders) != 1 {
		t.Fatalf("orders len = %d, want 1", len(got.Orders))
	}
	if got.Orders[0]["side"] != "bid" || got.Orders[0]["count"] != "1.00" {
		t.Errorf("batch order element wrong: %v", got.Orders[0])
	}
}

func TestBatchCancelEventOrdersResponse_Unmarshal(t *testing.T) {
	raw := `{"orders":[
		{"order_id":"ord_1","reduced_by":"3.00","ts_ms":1730000000000},
		{"order_id":"ord_2","error":{"code":"not_found","message":"no such order"}}
	]}`

	var resp BatchCancelEventOrdersResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Orders) != 2 {
		t.Fatalf("orders len = %d, want 2", len(resp.Orders))
	}
	if resp.Orders[0].ReducedBy == nil || *resp.Orders[0].ReducedBy != "3.00" {
		t.Errorf("orders[0].ReducedBy = %v, want 3.00", resp.Orders[0].ReducedBy)
	}
	if resp.Orders[1].Error == nil || resp.Orders[1].Error.Code != "not_found" {
		t.Errorf("orders[1].Error = %v, want code not_found", resp.Orders[1].Error)
	}
}

func marshalToMap(t *testing.T, v any) map[string]any {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	return m
}

func TestCreateEventOrderResponse_Unmarshal(t *testing.T) {
	// Flat shape per the documented V2 create-order response schema.
	raw := `{
		"order_id": "ord_123",
		"client_order_id": "cli_456",
		"fill_count": "7.00",
		"remaining_count": "3.00",
		"average_fill_price": "0.5600",
		"average_fee_paid": "0.0100",
		"ts_ms": 1730000000000
	}`

	var resp CreateEventOrderResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if resp.OrderID != "ord_123" {
		t.Errorf("OrderID = %q, want ord_123", resp.OrderID)
	}
	if resp.FillCount != "7.00" {
		t.Errorf("FillCount = %q, want 7.00", resp.FillCount)
	}
	if resp.RemainingCount != "3.00" {
		t.Errorf("RemainingCount = %q, want 3.00", resp.RemainingCount)
	}
	if resp.AverageFillPrice == nil || *resp.AverageFillPrice != "0.5600" {
		t.Errorf("AverageFillPrice = %v, want 0.5600", resp.AverageFillPrice)
	}
	if resp.TsMs != 1730000000000 {
		t.Errorf("TsMs = %d, want 1730000000000", resp.TsMs)
	}
}
