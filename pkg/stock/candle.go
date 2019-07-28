package stock

import "time"

type Candle struct {
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
	Time   time.Time
}
