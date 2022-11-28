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
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "threeqgo/1.0")
}

func (t *threeQGo) SetAPIKey(apiKey string) {
	t.apiKey = apiKey
}

func (t *threeQGo) GetAPIKeyByUser(username, password string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/apikey", basePath), nil)
	req.Header.Set("X-AUTH-USERNAME", username)
	req.Header.Set("X-AUTH-PASSWD", password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "threeqgo/1.0")
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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/welcome", basePath), nil)
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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects", basePath), nil)
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
	var pResult ProjectsResponse
	err = json.Unmarshal(body, &pResult)
	return pResult.Projects, err
}

func (t *threeQGo) CreateProject(project ProjectCreate) (ProjectCreateResponse, error) {
	payloadJson, err := json.Marshal(project)
	if err != nil {
		return ProjectCreateResponse{}, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/projects", basePath), bytes.NewBuffer(payloadJson))
	t.setRequestHeaders(req)
	if err != nil {
		return ProjectCreateResponse{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return ProjectCreateResponse{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ProjectCreateResponse{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return ProjectCreateResponse{}, err
	}
	var pResult ProjectCreateResponse
	err = json.Unmarshal(body, &pResult)
	return pResult, nil
}

func (t *threeQGo) GetProject(id int64) (Project, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d", basePath, id), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return Project{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Project{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Project{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Project{}, err
	}
	var project Project
	err = json.Unmarshal(body, &project)
	return project, err
}

func (t *threeQGo) UpdateProject(id int64, project ProjectUpdate) (Project, error) {
	payloadJson, err := json.Marshal(project)
	if err != nil {
		return Project{}, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%d", basePath, id), bytes.NewBuffer(payloadJson))
	t.setRequestHeaders(req)
	if err != nil {
		return Project{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Project{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Project{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Project{}, err
	}
	var pResult Project
	err = json.Unmarshal(body, &pResult)
	return pResult, nil
}

func (t *threeQGo) GetProjectChannels(id int64) ([]Channel, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d/channels", basePath, id), nil)
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
	var cResponse ChannelsResponse
	err = json.Unmarshal(body, &cResponse)
	return cResponse.Channels, err
}

func (t *threeQGo) DeleteProject(id int64) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/projects/%d", basePath, id), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return err
	}
	return nil
}

func (t *threeQGo) GetChannels() ([]Channel, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/channels", basePath), nil)
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
	var cResponse ChannelsResponse
	err = json.Unmarshal(body, &cResponse)
	return cResponse.Channels, err
}

func (t *threeQGo) GetChannel(id int64) (Channel, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/channels/%d", basePath, id), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return Channel{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Channel{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Channel{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Channel{}, err
	}
	var cResponse Channel
	err = json.Unmarshal(body, &cResponse)
	return cResponse, err
}

func (t *threeQGo) GetChannelRecorders() ([]Recorder, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/channels/recorders", basePath), nil)
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
	var cResponse RecordersResponse
	err = json.Unmarshal(body, &cResponse)
	return cResponse.ChannelRecorders, err
}

func (t *threeQGo) GetChannelRecorder(channelID, recorderID int64) (Recorder, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/channels/%d/recorders/%d", basePath, channelID, recorderID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return Recorder{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Recorder{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Recorder{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Recorder{}, err
	}
	var cResponse Recorder
	err = json.Unmarshal(body, &cResponse)
	return cResponse, err
}

func (t *threeQGo) UpdateChannelRecorder(channelID, recorderID int64, recorder RecorderUpdate) (Recorder, error) {
	payloadJson, err := json.Marshal(recorder)
	if err != nil {
		return Recorder{}, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/channels/%d/recorders/%d", basePath, channelID, recorderID), bytes.NewBuffer(payloadJson))
	t.setRequestHeaders(req)
	if err != nil {
		return Recorder{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Recorder{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Recorder{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Recorder{}, err
	}
	var pResult Recorder
	err = json.Unmarshal(body, &pResult)
	return pResult, nil
}

func (t *threeQGo) DeleteChannelRecorder(channelID, recorderID int64) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%d/recorders/%d", basePath, channelID, recorderID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return err
	}
	return nil
}

func (t *threeQGo) CreateChannelRecorder(channelID, dstProjectID int64, recorder RecorderCreate) (Recorder, error) {
	payloadJson, err := json.Marshal(recorder)
	if err != nil {
		return Recorder{}, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/channels/%d/recorders/%d", basePath, channelID, dstProjectID), bytes.NewBuffer(payloadJson))
	t.setRequestHeaders(req)
	if err != nil {
		return Recorder{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Recorder{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Recorder{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Recorder{}, err
	}
	var pResult Recorder
	err = json.Unmarshal(body, &pResult)
	return pResult, nil
}

func (t *threeQGo) ChannelRecorderAddCategory(channelID, recorderID, categoryID int64) (Recorder, error) {
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/channels/%d/recorders/%d/categories/%d", basePath, channelID, recorderID, categoryID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return Recorder{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Recorder{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Recorder{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Recorder{}, err
	}
	var cResponse Recorder
	err = json.Unmarshal(body, &cResponse)
	return cResponse, err
}

func (t *threeQGo) ChannelRecorderRemoveCategory(channelID, recorderID, categoryID int64) (Recorder, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/channels/%d/recorders/%d/categories/%d", basePath, channelID, recorderID, categoryID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return Recorder{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return Recorder{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Recorder{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return Recorder{}, err
	}
	var cResponse Recorder
	err = json.Unmarshal(body, &cResponse)
	return cResponse, err
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
