package controllers

import (
	"time"

	"github.com/device-fish-feeder/models"
)

var Device *models.Device

func FeedFish() (*models.FeedReport, error) {
	time.Sleep(4 * time.Second)
	success := true // place holders
	// err := nil
	feed_report := models.NewFeedReport("1324", success)

	return feed_report, nil
}
