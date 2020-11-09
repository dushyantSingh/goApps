package main

import "fmt"

//Trade struct to store info
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) value() float64 {
	return float64(t.Volume) * t.Price
}

func newTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol cant be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("Volume cant be zero or less")
	}
	if price <= 0 {
		return nil, fmt.Errorf("Price cant be zero or less")
	}

	trade := &Trade{Symbol: symbol, Volume: volume, Price: price, Buy: buy}
	return trade, nil
}
func main() {
	t := Trade{"MSFT", 30, 202, true}

	fmt.Println(t)
	fmt.Println(t.value())

	t2, err := newTrade("MSFT", 30, 400, true)

	fmt.Println(t2, err)
}
