# go-heikin-ashi

Convert regular OCHL candles to Heikin Ashi type candles.

<img width="752" alt="Screenshot 2021-07-15 at 18 41 49" src="https://user-images.githubusercontent.com/7526930/125825592-e354fa1a-5e1e-4489-ba9f-d1d3bfb0f3a5.png">

### Usage Example
```golang
ha := NewHeikinAshi()

candles := []Candle{
    {Open: 4261.48, Close: 4086.29, High: 4485.39, Low: 3850.00},
    {Open: 4069.13, Close: 4310.01, High: 4453.91, Low: 3400.00},
    {Open: 4310.01, Close: 4509.08, High: 4939.19, Low: 4124.54},
    {Open: 4505.00, Close: 4130.37, High: 4788.59, Low: 3603.00},
    {Open: 4153.62, Close: 3699.99, High: 4394.59, Low: 2817.00},
    {Open: 3690.00, Close: 3660.02, High: 4123.20, Low: 3505.55},
    {Open: 3660.02, Close: 4378.48, High: 4406.52, Low: 3653.69},
    {Open: 4400.00, Close: 4640.00, High: 4658.00, Low: 4110.00},
    {Open: 4640.00, Close: 5709.99, High: 5922.30, Low: 4550.00},
}

var results []HeikinAshiCandle

for _, candle := range candles {
    haCandle := ha.CalculateHeikinAshi(candle)
    results = append(results, haCandle)
}
```


### Features
- Fully compliant with TradingView
- Low memory usage
- Straightforward usage

### Planned features
- Multi thread safety
- More usage examples
- Maybe some extra caching 
