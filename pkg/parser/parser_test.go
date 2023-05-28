package parser

import "testing"

func TestGetExchangeRate(t *testing.T) {
	binanceParser := NewBinanceCryptoParser()
	rate, err := binanceParser.GetExchangeRate("BTC", "UAH")
	if err != nil {
		t.Errorf("failure occurs while parsing exhange rate: %v", err)
	}

	if rate < 0 {
		t.Errorf("echange rate is negative value: %v", rate)
	}
}
