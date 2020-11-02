package sec

import (
	"github.com/Sutheres/report-chaser/internal/sec/models"
	"net/http"
)

type EdgarIndex string

var (
	JSONExtension EdgarIndex = "index.json"
)

type SEC interface {
	GetDailyReports() ([]models.Item, error)
	GetFileExtension(file string) string
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

func (c *client) GetFileExtension(file string) string {
	return file[len(file)-3:]
}
