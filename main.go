package main

import (
	"fmt"

	"okex-websocket/eqsModels"
	Okex "okex-websocket/exchange/okex"
)

func main() {
	symbol := eqsModels.EqsSymbol{
		Base:  "BTC",
		Quote: "USDT",
		Sep:   "/",
	}

	okexConn := Okex.NewOkexWsConnWithHost("wss", "real.okex.com:8443", symbol)
	go okexConn.StartspotTrade()

	// 获取行情数据
	go func() {
		for {
			outData := <-okexConn.OutChan
			data := outData.(Okex.OkexSpotTrade)
			fmt.Printf("%+v\n", data)
		}
	}()

	// 获取报错信息
	go func() {
		for {
			fmt.Println(<-okexConn.ErrChan)
		}
	}()

	select {}
}
