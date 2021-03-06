/*
获取市场交易
api文档: https://www.okex.me/docs/zh/#spot_ws-trade
*/
package Okex

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
字段		数据类型	描述
id		integer	unix时间，同时作为消息ID
amount	float	24小时成交量
count	integer	24小时成交笔数
open	float	24小时开盘价
close	float	最新价
low		float	24小时最低价
high	float	24小时最高价
vol		float	24小时成交额
*/
type OkexSpotTrade struct {
	Table string
	Data  []struct {
		InstrumentId string
		Price        string
		Side         string
		Size         string
		Timestamp    string
		TradeId      string
	}
}

func (h *OkexWsConn) StartspotTrade() {
	// 创建ws
	for {
		fmt.Println("okex create connection begin")
		if h.createConnection() {
			fmt.Println("okex create connection success")
			break
		}
		fmt.Println("okex create connectting")
	}
	defer h.wsConn.Close()

	// 订阅
	for {
		if h.subscribespotTrade() {
			break
		}
	}

	// ping
	go h.ping()

	go h.readspotTrade()

	select {}
}

/*
订阅频道
*/
func (h *OkexWsConn) subscribespotTrade() bool {
	//symbol := h.symbol.ToUpperWithSep("-")
	args := []string{
		//fmt.Sprintf("spot/trade:%s", symbol),
		`{"channel": "candle15m", "instId": "BTC-USDT"}`,
	}
	message, _ := json.Marshal(map[string]interface{}{
		"op":   "subscribe",
		"args": args,
	})
	fmt.Println("message:", string(message))
	if !h.writeMessage(message) {
		return false
	}

	return true
}

/*
获取市场成交
*/
func (h *OkexWsConn) readspotTrade() {
	defer func() {
		go h.StartspotTrade()
	}()

	for {
		unzipData, err := h.readMessage()
		if err != nil {
			h.ErrChan <- err
		}

		dataStr := string(unzipData)
		if strings.Contains(dataStr, "event") { // 订阅成功消息
			continue
		} else {
			rt := OkexSpotTrade{
				Table: "",
				Data:  nil,
			}
			err := json.Unmarshal(unzipData, &rt)
			if err != nil {
				h.ErrChan <- err
			} else {
				h.OutChan <- rt
			}
		}
	}
}
