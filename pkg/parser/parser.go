package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BinanceCryptoParser struct {
	baseURL string
}

func NewBinanceCryptoParser() *BinanceCryptoParser {
	return &BinanceCryptoParser{
		baseURL: "https://api.binance.com/api/v3/ticker",
	}
}

func (p *BinanceCryptoParser) GetExchangeRate(baseCurrency string, quoteCurrency string) (float64, error) {
	requestURL := fmt.Sprintf("%s/price?symbol=%s%s", p.baseURL, baseCurrency, quoteCurrency)

	response, err := http.Get(requestURL)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	var result struct {
		Price string `json:"price"`
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return 0, err
	}

	rate, err := strconv.ParseFloat(result.Price, 64)
	if err != nil {
		return 0, err
	}

	return rate, nil
}
