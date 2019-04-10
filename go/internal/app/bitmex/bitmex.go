package bitmex

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ccxt/ccxt/go/pkg/ccxt"
)

// BitmexExchange struct
type BitmexExchange struct {
}

// LoadMarkets provides the available markets
// func (x *BitmexExchange) LoadMarkets() (map[string]ccxt.Market, error) {
// 	markets, err := x.FetchMarkets()
// 	if err != nil {
// 		return nil, err
// 	}
// 	mmap := make(map[string]Market, len(markets))
// 	for _, v := range markets {
// 		v.Exchange = x
// 		mmap[v.Symbol] = v
// 	}
// 	return mmap, nil
// }

// func (x *BitmexExchange) FetchTickers(symbols []string, params map[string]interface{}) (map[string]Ticker, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchTicker(symbol string, params map[string]interface{}) (Ticker, error) {
// 	ticker := Ticker{}
// 	return ticker, nil
// }
// func (x *BitmexExchange) FetchOHLCV(symbol, tf string, since *util.JSONTime, limit *int, params map[string]interface{}) ([]OHLCV, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchOrderBook(symbol string, limit *int, params map[string]interface{}) (OrderBook, error) {
// 	orderbook := OrderBook{}
// 	return orderbook, nil
// }
// func (x *BitmexExchange) FetchL2OrderBook(symbol string, limit *int, params map[string]interface{}) (OrderBook, error) {
// 	orderbook := OrderBook{}
// 	return orderbook, nil
// }
// func (x *BitmexExchange) FetchTrades(symbol string, since *util.JSONTime, params map[string]interface{}) ([]Trade, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchOrder(id string, symbol *string, params map[string]interface{}) (Order, error) {
// 	order := Order{}
// 	return order, nil
// }
// func (x *BitmexExchange) FetchOrders(symbol *string, since *util.JSONTime, limit *int, params map[string]interface{}) ([]Order, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchOpenOrders(symbol *string, since *util.JSONTime, limit *int, params map[string]interface{}) ([]Order, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchClosedOrders(symbol *string, since *util.JSONTime, limit *int, params map[string]interface{}) ([]Order, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchMyTrades(symbol *string, since *util.JSONTime, limit *int, params map[string]interface{}) ([]Trade, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchBalance(params map[string]interface{}) (Balances, error) {
// 	balances := Balances{}
// 	return balances, nil
// }
// func (x *BitmexExchange) FetchCurrencies() ([]Currency, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) FetchMarkets() ([]Market, error) {
// 	return nil, nil
// }
// func (x *BitmexExchange) CreateOrder(symbol, t, side string, amount float64, price *float64, params map[string]interface{}) (Order, error) {
// 	order := Order{}
// 	return order, nil
// }
// func (x *BitmexExchange) CancelOrder(id string, symbol *string, params map[string]interface{}) error {
// 	return nil
// }

// Info on the BitmexExchange
func (x *BitmexExchange) Info() (ccxt.ExchangeInfo, error) {
	var info ccxt.ExchangeInfo
	configFile := "bitmex.json"
	f, err := os.Open(configFile)
	defer f.Close()
	if err != nil {
		return info, fmt.Errorf("Config %s missing or errored: %v", configFile, err)
	}
	json.NewDecoder(f).Decode(&info)
	// info := ccxt.ExchangeInfo{
	// 	ID:        "bitmex",
	// 	Name:      "bitMex",
	// 	Countries: []string{"SC"},
	// 	Version:   "v1",
	// 	// UserAgent: nil,
	// 	RateLimit: 2000,
	// 	Has: ccxt.HasDescription{
	// 		CORS:              false,
	// 		FetchOHLCV:        "true",
	// 		Withdraw:          true,
	// 		EditOrder:         "true",
	// 		FetchOrder:        true,
	// 		FetchOrders:       true,
	// 		FetchOpenOrders:   true,
	// 		FetchClosedOrders: true,
	// 		FetchMyTrades:     true,
	// 	},
	// 	Timeframes: map[string]string{
	// 		"1m": "1m",
	// 		"5m": "5m",
	// 		"1h": "1h",
	// 		"1d": "1d",
	// 	},
	// 	URLs: ccxt.URLs{
	// 		Logo: ccxt.StringSlice{"https://user-images.githubusercontent.com/1294454/27766319-f653c6e6-5ed4-11e7-933d-f0bc3699ae8f.jpg"},
	// 		WWW:  ccxt.StringSlice{"https://www.bitmex.com"},
	// 		Doc:  ccxt.StringSlice{"https://www.bitmex.com/app/apiOverview", "https://github.com/BitMEX/api-connectors/tree/master/official-http"},
	// 		Fees: ccxt.StringSlice{"https://www.bitmex.com/app/fees"},
	// 	},
	// 	API: ccxt.APIs{
	// 		Public: ccxt.APIMethods{
	// 			Get: []string{
	// 				"announcement",
	// 				"announcement/urgent",
	// 				"funding",
	// 				"instrument",
	// 				"instrument/active",
	// 				"instrument/activeAndIndices",
	// 				"instrument/activeIntervals",
	// 				"instrument/compositeIndex",
	// 				"instrument/indices",
	// 				"insurance",
	// 				"leaderboard",
	// 				"liquidation",
	// 				"orderBook",
	// 				"orderBook/L2",
	// 				"quote",
	// 				"quote/bucketed",
	// 				"schema",
	// 				"schema/websocketHelp",
	// 				"settlement",
	// 				"stats",
	// 				"stats/history",
	// 				"trade",
	// 				"trade/bucketed",
	// 			},
	// 		},
	// 		Private: ccxt.APIMethods{
	// 			Get: []string{
	// 				"apiKey",
	// 				"chat",
	// 				"chat/channels",
	// 				"chat/connected",
	// 				"execution",
	// 				"execution/tradeHistory",
	// 				"notification",
	// 				"order",
	// 				"position",
	// 				"user",
	// 				"user/affiliateStatus",
	// 				"user/checkReferralCode",
	// 				"user/commission",
	// 				"user/depositAddress",
	// 				"user/margin",
	// 				"user/minWithdrawalFee",
	// 				"user/wallet",
	// 				"user/walletHistory",
	// 				"user/walletSummary",
	// 			},
	// 			Post: []string{
	// 				"apiKey",
	// 				"apiKey/disable",
	// 				"apiKey/enable",
	// 				"chat",
	// 				"order",
	// 				"order/bulk",
	// 				"order/cancelAllAfter",
	// 				"order/closePosition",
	// 				"position/isolate",
	// 				"position/leverage",
	// 				"position/riskLimit",
	// 				"position/transferMargin",
	// 				"user/cancelWithdrawal",
	// 				"user/confirmEmail",
	// 				"user/confirmEnableTFA",
	// 				"user/confirmWithdrawal",
	// 				"user/disableTFA",
	// 				"user/logout",
	// 				"user/logoutAll",
	// 				"user/preferences",
	// 				"user/requestEnableTFA",
	// 				"user/requestWithdrawal",
	// 			},
	// 			Put: []string{
	// 				"order",
	// 				"order/bulk",
	// 				"user",
	// 			},
	// 			Delete: []string{
	// 				"apiKey",
	// 				"order",
	// 				"order/all",
	// 			},
	// 		},
	// 	},
	// 	Exceptions: ccxt.Exceptions{
	// 		Exact: map[string]string{
	// 			"Invalid API Key.":    "AuthenticationError",
	// 			"Access Denied":       "PermissionDenied",
	// 			"Duplicate clOrdID":   "InvalidOrder",
	// 			"Signature not valid": "AuthenticationError",
	// 		},
	// 		Broad: map[string]string{
	// 			"overloaded": "ExchangeNotAvailable",
	// 			"Account has insufficient Available Balance": "InsufficientFunds",
	// 		},
	// 	},
	// }
	return info, nil
}

// func (x *BitmexExchange) GetMarket(symbol string) (Market, error) {
// 	market := Market{}
// 	return market, nil
// }
// func (x *BitmexExchange) CreateLimitBuyOrder(symbol string, amount float64, price *float64, params map[string]interface{}) (Order, error) {
// 	order := Order{}
// 	return order, nil
// }
// func (x *BitmexExchange) CreateLimitSellOrder(symbol string, amount float64, price *float64, params map[string]interface{}) (Order, error) {
// 	order := Order{}
// 	return order, nil
// }
// func (x *BitmexExchange) CreateMarketBuyOrder(symbol string, amount float64, params map[string]interface{}) (Order, error) {
// 	order := Order{}
// 	return order, nil
// }
// func (x *BitmexExchange) CreateMarketSellOrder(symbol string, amount float64, params map[string]interface{}) (Order, error) {
// 	order := Order{}
// 	return order, nil
// }
