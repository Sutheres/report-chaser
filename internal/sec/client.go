//go:generate mockgen -package=mocks -destination=../../mocks/sec.go github.com/Sutheres/report-chaser/internal/sec SEC

package sec

import (
	"net/http"
)

type SEC interface {

}

type client struct {
	host string
	c    http.Client
}

func NewClient(host string) SEC {
	return &client{
		host: host,
		c:    http.Client{},
	}
}
