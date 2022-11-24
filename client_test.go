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
