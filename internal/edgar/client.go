package edgar

import (
	"github.com/Sutheres/report-chaser/internal/edgar/models"
	"net/http"
)

type EdgarIndex string

var (
	JSONExtension EdgarIndex = "index.json"
)

type Edgar interface {
	GetDailyReports() ([]models.Item, error)
	GetFileExtension(file string) string
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

func (c *client) GetFileExtension(file string) string {
	return file[len(file)-3:]
}
