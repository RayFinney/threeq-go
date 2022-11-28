package threeqgo

import (
	"os"
	"testing"
)

func TestWelcome(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	err := client.Welcome()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAPIKeyByUser(t *testing.T) {
	client := NewClient(nil)
	apiKey, err := client.GetAPIKeyByUser(os.Getenv("USERNAME"), os.Getenv("PASSWD"))
	if err != nil || apiKey == "" {
		t.Fatal(err)
	}
}

func TestGetProjects(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.GetProjects()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateProject(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.CreateProject(ProjectCreate{
		Label:        "Test",
		StreamTypeId: StreamTypeVideoOnDemand,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProject(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	project, err := client.GetProject(49763)
	if err != nil || project.ID != 49763 {
		t.Fatal(err)
	}
}

func TestUpdateProject(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	project, err := client.UpdateProject(49763, ProjectUpdate{
		Label:         "Virtual Events DEV VoD",
		TokenSecurity: false,
		UsePlayerV5:   true,
	})
	if err != nil || project.ID != 49763 {
		t.Fatal(err)
	}
}

func TestDeleteProject(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	err := client.DeleteProject(62577)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProjectChannels(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	channels, err := client.GetProjectChannels(58931)
	if err != nil || channels[0].ID == 0 {
		t.Fatal(err)
	}
}

func TestGetChannels(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	channels, err := client.GetChannels()
	if err != nil || channels[0].ID == 0 {
		t.Fatal(err)
	}
}

func TestGetChannel(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	channels, err := client.GetChannel(50559)
	if err != nil || channels.ID != 50559 {
		t.Fatal(err)
	}
}

func TestGetChannelRecorders(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.GetChannelRecorders()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetChannelRecorder(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.GetChannelRecorder(50559, 7863)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateChannelRecorder(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.UpdateChannelRecorder(50559, 7863, RecorderUpdate{
		Title:                "Auto Test",
		Description:          "Auto Test",
		AutoRecording:        true,
		UseRecordingInterval: false,
		RecordingInterval:    120,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteChannelRecorder(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	err := client.DeleteChannelRecorder(50559, 7863)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateChannelRecorder(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.CreateChannelRecorder(50559, 49763, RecorderCreate{
		Title:                "Auto Test",
		Description:          "Auto Test",
		AutoRecording:        true,
		UseRecordingInterval: false,
		RecordingInterval:    120,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncodeQuery(t *testing.T) {
	query := encodeQuery(FileSearchOptions{
		IncludeDeleted:    false,
		IncludeMetadata:   false,
		IncludePlayouts:   false,
		IncludeProperties: false,
		Offset:            0,
		Limit:             5,
	})
	expect := "?IncludeDeleted=false&IncludeMetadata=false&IncludePlayouts=false&IncludeProperties=false&Limit=5"
	if query != expect {
		t.Errorf("expected: %s \n got: %s", expect, query)
	}
}

func TestGetFiles(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	files, err := client.GetFiles(49763, FileSearchOptions{
		IncludeDeleted:    false,
		IncludeMetadata:   false,
		IncludePlayouts:   false,
		IncludeProperties: false,
		Offset:            0,
		Limit:             1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Errorf("size should be 1 but is %d", len(files))
	}
	files, err = client.GetFiles(49763, FileSearchOptions{
		IncludeDeleted:    false,
		IncludeMetadata:   false,
		IncludePlayouts:   false,
		IncludeProperties: false,
		Offset:            0,
		Limit:             5,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 5 {
		t.Errorf("size should be 5 but is %d", len(files))
	}
}

func TestGetFile(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	file, err := client.GetFile(49763, 6794663)
	if err != nil && file.ID == 6794663 {
		t.Fatal(err)
	}
}
