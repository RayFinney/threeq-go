package threeqgo

import "time"

type ProjectsResult struct {
	Projects []Project `json:"Projects"`
}

type Project struct {
	ID            int64             `json:"Id"`
	StreamType    StreamType        `json:"StreamType"`
	Cluster       string            `json:"Cluster"`
	Category      []ProjectCategory `json:"Category"`
	Label         string            `json:"Label"`
	SecurityKey   string            `json:"SecurityKey"`
	TokenSecurity bool              `json:"TokenSecurity"`
	ThumbURI      string            `json:"ThumbURI"`
	CreatedAt     time.Time         `json:"CreatedAt"`
	LastUpdatedAt time.Time         `json:"LastUpdatedAt"`
	Expires       bool              `json:"Expires"`
	ExpiresAt     time.Time         `json:"ExpiresAt"`
	UsePlayerV5   bool              `json:"UsePlayerV5"`
	BottalkApiKey string            `json:"BottalkApiKey"`
	UsePubdate    bool              `json:"UsePubdate"`
}

type StreamType struct {
	ID int64 `json:"Id"`
}

type ProjectCategory struct {
	ID       int64  `json:"Id"`
	Label    string `json:"Label"`
	CustomId string `json:"CustomId"`
	ThumbURI string `json:"ThumbURI"`
}
