package sdk

import (
	"bytes"
	"net/http"
)

type Client struct {
	BaseURL string
}

// NewClient  Example "http://192.168.4.245:31500"
func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) CreateWorkFlow(data string) (*http.Response, error) {
	url := c.BaseURL + "/api/metadata/workflow"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RunWorkFlow(data string) (*http.Response, error) {
	url := c.BaseURL + "/api/workflow"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
