package threeqgo

type ThreeQGo interface {
	// Utility

	// Welcome will test authentication
	Welcome() error
	SetAPIKey(apiKey string)

	// ApiKey

	// GetAPIKeyByUser use your username and password to get the apiKey
	GetAPIKeyByUser(username, password string) (string, error)

	// Projects

	GetProjects() ([]Project, error)
}
