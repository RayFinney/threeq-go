package threeqgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var basePath = "https://sdn.3qsdn.com/api/v2"

type threeQGo struct {
	httpClient *http.Client
	apiKey     string
}

func (t *threeQGo) checkForErrorsInResponse(response *http.Response) error {
	if response == nil {
		return errors.New("no response")
	}
	if response.StatusCode >= 400 || response.StatusCode >= 500 {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("%s - %s", response.Status, string(body))
	}
	return nil
}

func (t *threeQGo) setRequestHeaders(req *http.Request) {
	req.Header.Set("X-AUTH-APIKEY", t.apiKey)
	req.Header.Set("accept", "application/json")
}

func (t *threeQGo) SetAPIKey(apiKey string) {
	t.apiKey = apiKey
}

func (t *threeQGo) GetAPIKeyByUser(username, password string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/apikey", basePath), bytes.NewBufferString(""))
	req.Header.Set("X-AUTH-USERNAME", username)
	req.Header.Set("X-AUTH-PASSWD", password)
	req.Header.Set("accept", "application/json")
	if err != nil {
		return "", err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return "", err
	}

	apiKeyStruct := struct {
		ApiKey string `json:"APIKEY"`
	}{}
	err = json.Unmarshal(body, &apiKeyStruct)
	return apiKeyStruct.ApiKey, err
}

func (t *threeQGo) Welcome() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/welcome", basePath), bytes.NewBufferString(""))
	t.setRequestHeaders(req)
	if err != nil {
		return err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return t.checkForErrorsInResponse(response)
}

func (t *threeQGo) GetProjects() ([]Project, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects", basePath), bytes.NewBufferString(""))
	t.setRequestHeaders(req)
	if err != nil {
		return nil, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return nil, err
	}
	var pResult ProjectsResult
	err = json.Unmarshal(body, &pResult)
	return pResult.Projects, err
}

// NewClient creates a new Client, apiKey and httpClient are optional.
func NewClient(httpClient *http.Client) ThreeQGo {
	c := threeQGo{
		httpClient: httpClient,
	}
	if c.httpClient == nil {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = 100
		t.MaxConnsPerHost = 100
		t.MaxIdleConnsPerHost = 100

		c.httpClient = &http.Client{
			Timeout:   10 * time.Second,
			Transport: t,
		}
	}
	return &c
}
