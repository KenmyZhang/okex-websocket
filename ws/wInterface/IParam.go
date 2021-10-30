package wInterface

import . "okex-websocket/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
