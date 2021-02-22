package sec

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func (s *client) GetTickerValues() ([]Ticker, error) {
	var response map[string]Ticker

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/files/company_tickers.json", s.host),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest")
	}
	resp, err := s.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Do")
	}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "Decode")
	}
	defer resp.Body.Close()

	var tickers []Ticker
	for _, v := range response {
		tickers = append(tickers, v)
	}

	return tickers, nil
}
