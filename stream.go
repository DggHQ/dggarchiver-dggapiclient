package dggapiclient

import (
	"context"
	"fmt"
	"net/http"
)

/*
The API data struct where the json will be parsed into
*/
type API struct {
	Data struct {
		Streams struct {
			Youtube struct {
				Live       bool   `json:"live"`
				Preview    string `json:"preview"`
				StatusText string `json:"status_text"`
				StartedAt  string `json:"started_at"`
				EndedAt    string `json:"ended_at"`
				Duration   int    `json:"duration"`
				ID         any    `json:"id"`
				Platform   string `json:"platform"`
				Type       string `json:"type"`
			} `json:"youtube"`
		} `json:"streams"`
	} `json:"data"`
}

/*
Get stream info on /info/stream
*/
func (c *DGGApiClient) GetStreamInfo(ctx context.Context) (*API, error) {
	// Get API Endpoint for all Jobs
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/info/stream", c.URL), nil)
	if err != nil {
		return nil, err
	}
	// Execute the request
	req = req.WithContext(ctx)

	// Send the request to the endpoint
	var res API
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
