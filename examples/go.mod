module main

go 1.23.4

// require  v2.8.0

replace github.com/HaoweiCh/go-binance-api => ../

replace github.com/HaoweiCh/go-binance-api/futures => ../futures

require github.com/HaoweiCh/go-binance-api v1.0.0-00010101000000-000000000000

require (
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
)
