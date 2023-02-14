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

func TestGetFileEncoderSettings(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.GetFileEncoderSettings(49763)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateFileEncoderSettings(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.UpdateFileEncoderSettings(49763, FileEncoderSettingsUpdate{
		UseEncoding:             true,
		UsePremiumEncoding:      false,
		UsePreProcessing:        false,
		UseAIServices:           false,
		SourceLanguage:          "eng",
		GenerateSubtitles:       false,
		UseDeinterlace:          false,
		UseTwoPass:              false,
		UseCMAF:                 false,
		UseBPFrames:             true,
		PackageAudioOnlyVariant: false,
		PackageForDRM:           false,
		UseWatermark:            false,
		WatermarkPosition:       "bottom-right",
		UseCropping:             false,
		CroppingParameters:      "",
		NormalizeAudio:          false,
		UseFirstAudioTrack:      false,
		UseAllAudioTracks:       false,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFileFormatSettings(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.GetFileFormatSettings(49763)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFileFormat(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.GetFileFormat(49763, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateFileFormat(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.UpdateFileFormat(49763, 1, FileFormatUpdate{
		VideoBitRate:    4000,
		VideoProfile:    "main",
		VideoFPS:        25,
		AudioBitRate:    152,
		AudioSampleRate: 48000,
		AudioChannels:   2,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddFileFormat(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.AddFileFormat(49763, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveFileFormat(t *testing.T) {
	client := NewClient(nil)
	client.SetAPIKey(os.Getenv("API_KEY"))
	_, err := client.RemoveFileFormat(49763, 1)
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
		IncludeDeleted:    true,
		IncludeMetadata:   true,
		IncludePlayouts:   true,
		IncludeProperties: true,
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
		IncludeDeleted:    true,
		IncludeMetadata:   true,
		IncludePlayouts:   true,
		IncludeProperties: true,
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
