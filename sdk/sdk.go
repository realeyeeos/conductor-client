package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Client struct {
	BaseURL string
}

type CreateWorkflowResponse struct {
	Status           int               `json:"status"`
	Message          string            `json:"message"`
	Instance         string            `json:"instance,omitempty"`
	Retryable        bool              `json:"retryable"`
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
}

type ValidationError struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

type RunWorkflowRequest struct {
	Name    string      `json:"name"`
	Version string      `json:"version"`
	Input   interface{} `json:"input"`
}

type RunWorkflowResponse struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Instance  string `json:"instance,omitempty"`
	Retryable bool   `json:"retryable"`
}

// NewClient  Example "http://192.168.4.245:31500"
func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) CreateWorkFlow(data string) (CreateWorkflowResponse, error) {
	var apiResponse CreateWorkflowResponse
	url := c.BaseURL + "/api/metadata/workflow"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		return apiResponse, err
	}
	if resp.StatusCode != http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&apiResponse)
		if err != nil {
			return apiResponse, err
		}
		return apiResponse, errors.New(apiResponse.Message)
	}
	//表示创建成功
	return apiResponse, nil
}

func (c *Client) RunWorkFlow(workFlowName string, workFlowVersion string, input map[string]interface{}) (RunWorkflowResponse, error) {

	var data RunWorkflowRequest
	var apiResponse RunWorkflowResponse
	data.Name = workFlowName
	data.Version = workFlowVersion
	data.Input = input

	jsonData, err := json.Marshal(data)
	if err != nil {
		return apiResponse, err
	}

	url := c.BaseURL + "/api/workflow"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return apiResponse, err
	}
	if resp.StatusCode != http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&apiResponse)
		if err != nil {
			return apiResponse, err
		}
		return apiResponse, errors.New(apiResponse.Message)
	}
	return apiResponse, nil
}
