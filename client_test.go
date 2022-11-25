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
