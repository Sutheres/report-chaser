package edgar

import (
	"net/http"
)

type Edgar interface {

}

type client struct {
	host string
	c    http.Client
}

func NewClient(host string) Edgar {
	return &client{
		host: host,
		c:    http.Client{},
	}
}
