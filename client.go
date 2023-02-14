package threeqgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var basePath = "https://sdn.3qsdn.com/api/v2"
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type threeQGo struct {
	httpClient *http.Client
	apiKey     string
}

func (t *threeQGo) checkForErrorsInResponse(response *http.Response) error {
	if response == nil {
		return errors.New("no response")
	}
	if response.StatusCode >= 400 || response.StatusCode >= 500 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
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

func (t *threeQGo) GetFileEncoderSettings(projectID int64) (FileEncodingSetting, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d/fileencodersettings", basePath, projectID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return FileEncodingSetting{}, err
	}
	var settings FileEncodingSetting
	err = json.Unmarshal(body, &settings)
	return settings, err
}

func (t *threeQGo) UpdateFileEncoderSettings(projectID int64, settings FileEncoderSettingsUpdate) (FileEncodingSetting, error) {
	payloadJson, err := json.Marshal(settings)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%d/fileencodersettings", basePath, projectID), bytes.NewBuffer(payloadJson))
	t.setRequestHeaders(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return FileEncodingSetting{}, err
	}
	var sResult FileEncodingSetting
	err = json.Unmarshal(body, &sResult)
	return sResult, nil
}

func (t *threeQGo) SetWatermarkPicture(projectID int64, filename, contentType string, watermark io.Reader) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		EscapeQuotes("file"), EscapeQuotes(filepath.Base(filename))))
	h.Set("Content-Type", contentType)
	part, _ := writer.CreatePart(h)
	_, err := io.Copy(part, watermark)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/projects/%d/fileencodersettings/watermarkpicture", basePath, projectID), body)
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

func (t *threeQGo) GetFileFormatSettings(projectID int64) (FileFormatSettings, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d/fileformatsettings", basePath, projectID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return FileFormatSettings{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return FileFormatSettings{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FileFormatSettings{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return FileFormatSettings{}, err
	}
	var settings FileFormatSettings
	err = json.Unmarshal(body, &settings)
	return settings, err
}

func (t *threeQGo) GetFileFormat(projectID int64, fileFormatID int64) (FileFormat, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d/fileformatsettings/%d", basePath, projectID, fileFormatID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return FileFormat{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return FileFormat{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FileFormat{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return FileFormat{}, err
	}
	var settings FileFormat
	err = json.Unmarshal(body, &settings)
	return settings, err
}

func (t *threeQGo) AddFileFormat(projectID int64, fileFormatID int64) (FileEncodingSetting, error) {
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/projects/%d/fileencodersettings/fileformat/%d", basePath, projectID, fileFormatID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return FileEncodingSetting{}, err
	}
	var settings FileEncodingSetting
	err = json.Unmarshal(body, &settings)
	return settings, err
}

func (t *threeQGo) RemoveFileFormat(projectID int64, fileFormatID int64) (FileEncodingSetting, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/projects/%d/fileencodersettings/fileformat/%d", basePath, projectID, fileFormatID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FileEncodingSetting{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return FileEncodingSetting{}, err
	}
	var settings FileEncodingSetting
	err = json.Unmarshal(body, &settings)
	return settings, err
}

func (t *threeQGo) UpdateFileFormat(projectID int64, fileFormatID int64, fileFormat FileFormatUpdate) (FileFormat, error) {
	payloadJson, err := json.Marshal(fileFormat)
	if err != nil {
		return FileFormat{}, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/projects/%d/fileformatsettings/%d", basePath, projectID, fileFormatID), bytes.NewBuffer(payloadJson))
	t.setRequestHeaders(req)
	if err != nil {
		return FileFormat{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return FileFormat{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return FileFormat{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return FileFormat{}, err
	}
	var sResult FileFormat
	err = json.Unmarshal(body, &sResult)
	return sResult, nil
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

func encodeQuery(s interface{}) string {
	query := ""
	v := reflect.ValueOf(s)
	fields := reflect.VisibleFields(reflect.TypeOf(s))
	for i, field := range fields {
		value := ""
		switch field.Type.Kind() {
		case reflect.String:
			value = v.Field(i).String()
		case reflect.Int64:
			value = strconv.FormatInt(v.Field(i).Int(), 10)
			break
		case reflect.Bool:
			value = strconv.FormatBool(v.Field(i).Bool())
			break
		}
		if value != "" && value != "0" {
			if i > 0 {
				query += "&"
			}
			query += fmt.Sprintf("%s=%s", field.Name, value)
		}
	}
	if query != "" {
		query = "?" + query
	}
	return query
}

func (t *threeQGo) GetFiles(projectID int64, queryParams FileSearchOptions) ([]File, error) {
	if queryParams.Limit == 0 {
		queryParams.Limit = 100
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d/files%s", basePath, projectID, encodeQuery(queryParams)), nil)
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
	var cResponse FilesResponse
	err = json.Unmarshal(body, &cResponse)
	return cResponse.Files, err
}

func (t *threeQGo) GetFile(projectID, fileID int64) (File, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d/files/%d", basePath, projectID, fileID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return File{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return File{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return File{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return File{}, err
	}
	var cResponse File
	err = json.Unmarshal(body, &cResponse)
	return cResponse, err
}

func (t *threeQGo) GetEncodingProgress(projectID, fileID int64) (EncodingProgress, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects/%d/files/%d/progress", basePath, projectID, fileID), nil)
	t.setRequestHeaders(req)
	if err != nil {
		return EncodingProgress{}, err
	}
	response, err := t.httpClient.Do(req)
	if err != nil {
		return EncodingProgress{}, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return EncodingProgress{}, err
	}
	defer response.Body.Close()
	if err := t.checkForErrorsInResponse(response); err != nil {
		return EncodingProgress{}, err
	}
	var cResponse EncodingProgress
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
