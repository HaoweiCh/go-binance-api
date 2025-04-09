package main

import (
	"context"
	"fmt"

	"github.com/HaoweiCh/go-binance-api"
	"github.com/HaoweiCh/go-binance-api/futures"
)

func SpotOrder() {
	binance.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance_api.NewClient(apiKey, secret)

	symbol := "BTCUSDT"
	side := binance.SideTypeSell
	orderType := binance.OrderTypeMarket
	quantity := "0.0001"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func FuturesOrder() {
	futures.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance_api.NewFuturesClient(apiKey, secret)

	symbol := "LTCUSDT"
	side := futures.SideTypeSell
	orderType := futures.OrderTypeMarket
	quantity := "0.1"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).PositionSide(futures.PositionSideTypeLong).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

}
