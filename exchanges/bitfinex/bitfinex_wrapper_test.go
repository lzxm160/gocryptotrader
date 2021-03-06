package bitfinex

import (
	"testing"

	"github.com/lzxm160/gocryptotrader/currency/pair"
	"github.com/lzxm160/gocryptotrader/exchanges/ticker"
)

func TestStart(t *testing.T) {
	start := Bitfinex{}
	start.Start()
}

func TestRun(t *testing.T) {
	run := Bitfinex{}
	run.Run()
}

func TestGetTickerPrice(t *testing.T) {
	getTickerPrice := Bitfinex{}
	_, err := getTickerPrice.GetTickerPrice(pair.NewCurrencyPair("BTC", "USD"),
		ticker.Spot)
	if err != nil {
		t.Errorf("Test Failed - Bitfinex GetTickerPrice() error: %s", err)
	}
}

func TestGetOrderbookEx(t *testing.T) {
	getOrderBookEx := Bitfinex{}
	_, err := getOrderBookEx.GetOrderbookEx(pair.NewCurrencyPair("BTC", "USD"),
		ticker.Spot)
	if err != nil {
		t.Errorf("Test Failed - Bitfinex GetOrderbookEx() error: %s", err)
	}
}
