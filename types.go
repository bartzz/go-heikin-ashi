package go_heikin_ashi

type Candle struct {
	Open  float64
	Close float64
	High  float64
	Low   float64
}

type HeikinAshiCandle struct {
	Open      float64
	Close     float64
	High      float64
	Low       float64
	IsBullish bool
}
