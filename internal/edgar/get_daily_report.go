package edgar

import (
	"encoding/json"
	"fmt"
	"github.com/Sutheres/report-chaser/internal/edgar/models"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"time"
)

func (c *client) GetDailyReports() ([]models.Item, error) {
	quarter := string(models.Q1)
	var masterIndexes []models.Item
	index, err := c.getDailyMasterReports(quarter)
	if err != nil {
		return masterIndexes, errors.Wrap(err, "getDailyMasterReport")
	}

	for _, item := range index.Items {
		if item.Name[:6] == "master" {
			today := time.Now().AddDate(0, 0, -5)
			parsedLastModified := item.LastModified[:len(item.LastModified)-3]
			lastUpdated, err := time.Parse("01/02/2006 15:04:05", parsedLastModified)
			if err != nil {
				return masterIndexes, errors.Wrap(err, "Parse")
			}
			if today.Equal(lastUpdated) ||
				today.Before(lastUpdated) {
				if c.GetFileExtension(item.Name) == "idx" {
					masterIndexes = append(masterIndexes, item)
				}
			}
		}
	}

	if masterIndexes != nil {
		for _, r := range masterIndexes {
			req, err := http.NewRequest(
				"GET",
				fmt.Sprintf("%s/%d/%s/%s", c.host, time.Now().Year(), quarter, r.Href),
				nil,
			)
			resp, err := c.c.Do(req)
			if err != nil {
				return masterIndexes, errors.Wrap(err, "Do")
			}
			defer resp.Body.Close()
			txtFile, err := os.Create(r.Href)
			if err != nil {
				return masterIndexes, errors.Wrap(err, "Create")
			}
			defer txtFile.Close()

			_, err = io.Copy(txtFile, resp.Body)
			if err != nil {
				return masterIndexes, errors.Wrap(err, "Copy")
			}
		}
	}
	return masterIndexes, nil
}

func (c *client) getDailyMasterReports(quarter string) (models.DailyIndex, error) {
	var d models.DailyIndex
	year := time.Now().Year()
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%d/%s/%s", c.host, year, quarter, JSONExtension),
		nil,
	)
	if err != nil {
		return d, errors.Wrap(err, "NewRequest")
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return d, errors.Wrap(err, "Do")
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return d, errors.Wrap(err, "Decode")
	}

	return d, nil
}