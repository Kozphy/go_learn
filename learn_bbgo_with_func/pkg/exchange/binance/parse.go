package binance

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

// IsBookTicker document ref :https://binance-docs.github.io/apidocs/spot/en/#individual-symbol-book-ticker-streams
// use key recognition because there's no identify in the content.
func IsBookTicker(val *fastjson.Value) bool {
	return !val.Exists("e") && val.Exists("u") &&
		val.Exists("s") && val.Exists("b") &&
		val.Exists("B") && val.Exists("a") && val.Exists("A")
}

func parseWebSocketEvent(message []byte) (interface{}, error) {
	val, err := fastjson.ParseBytes(message)

	if err != nil {
		return nil, err
	}

	// res, err := json.MarshalIndent(message, "", "  ")
	// if err != nil {
	//	log.Fatal(err)
	// }
	// str := strings.ReplaceAll(string(res), "\\", "")
	// fmt.Println(str)
	eventType := string(val.GetStringBytes("e"))
	if eventType == "" && IsBookTicker(val) {
		eventType = "bookTicker"
	}

	switch eventType {
	case "kline":
		var event KLineEvent
		err := json.Unmarshal([]byte(message), &event)
		return &event, err
	case "bookTicker":
		var event BookTickerEvent
		err := json.Unmarshal([]byte(message), &event)
		event.Event = eventType
		return &event, err

	case "outboundAccountPosition":
		var event OutboundAccountPositionEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	case "outboundAccountInfo":
		var event OutboundAccountInfoEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	case "balanceUpdate":
		var event BalanceUpdateEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	case "executionReport":
		var event ExecutionReportEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	case "depthUpdate":
		return parseDepthEvent(val)

	case "markPriceUpdate":
		var event MarkPriceUpdateEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	case "listenKeyExpired":
		var event ListenKeyExpired
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	// Binance futures data --------------
	case "continuousKline":
		var event ContinuousKLineEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	case "ORDER_TRADE_UPDATE":
		var event OrderTradeUpdateEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	// Event: Balance and Position Update
	case "ACCOUNT_UPDATE":
		var event AccountUpdateEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	// Event: Order Update
	case "ACCOUNT_CONFIG_UPDATE":
		var event AccountConfigUpdateEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	case "trade":
		var event MarketTradeEvent
		err = json.Unmarshal([]byte(message), &event)
		return &event, err

	default:
		id := val.GetInt("id")
		if id > 0 {
			return &ResultEvent{ID: id}, nil
		}
	}

	return nil, fmt.Errorf("unsupported message: %s", message)
}
