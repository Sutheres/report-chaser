package sec

type Ticker struct {
	CIK    int64  `json:"cik_str"`
	Ticker string `json:"ticker"`
	Title  string `json:"title"`
}
