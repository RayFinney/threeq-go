package threeqgo

import "time"

var StreamTypeVideoOnDemand, StreamTypeLivestream int64 = 1, 2

type ProjectsResponse struct {
	Projects []Project `json:"Projects"`
}

type ProjectCreate struct {
	Label        string `json:"Label"`
	StreamTypeId int64  `json:"StreamTypeId"` // 1 = Video Platform	| 2 = Video Livestream
}

type ProjectUpdate struct {
	Label         string `json:"Label"`
	TokenSecurity bool   `json:"TokenSecurity"`
	UsePlayerV5   bool   `json:"UsePlayerV5"`
}

type ProjectCreateResponse struct {
	ProjectId int64 `json:"ProjectId"`
	ChannelId int64 `json:"ChannelId"`
}

type Project struct {
	ID            int64      `json:"Id"`
	StreamType    StreamType `json:"StreamType"`
	Cluster       string     `json:"Cluster"`
	Category      []Category `json:"Category"`
	Label         string     `json:"Label"`
	SecurityKey   string     `json:"SecurityKey"`
	TokenSecurity bool       `json:"TokenSecurity"`
	ThumbURI      string     `json:"ThumbURI"`
	CreatedAt     time.Time  `json:"CreatedAt"`
	LastUpdatedAt time.Time  `json:"LastUpdatedAt"`
	Expires       bool       `json:"Expires"`
	ExpiresAt     time.Time  `json:"ExpiresAt"`
	UsePlayerV5   bool       `json:"UsePlayerV5"`
	BottalkApiKey string     `json:"BottalkApiKey"`
	UsePubdate    bool       `json:"UsePubdate"`
}

type StreamType struct {
	ID    int64  `json:"Id"`
	Label string `json:"Label"`
}

type Category struct {
	ID       int64  `json:"Id"`
	Label    string `json:"Label"`
	CustomId string `json:"CustomId"`
	ThumbURI string `json:"ThumbURI"`
}

type ChannelsResponse struct {
	Channels []Channel `json:"Channels"`
}

type Channel struct {
	ID            int64           `json:"Id"`
	ChannelStatus []ChannelStatus `json:"ChannelStatus"`
	Project       Project         `json:"Project"`
}

type ChannelStatus struct {
	IsOnline     bool   `json:"IsOnline"`
	IsPrimary    bool   `json:"IsPrimary"`
	Origin       string `json:"Origin"`
	VideoFormat  string `json:"VideoFormat"`
	VideoBitRate int64  `json:"VideoBitRate"`
	VideoWidth   int64  `json:"VideoWidth"`
	VideoHeight  int64  `json:"VideoHeight"`
	AudioFormat  string `json:"AudioFormat"`
	AudioBitRate int64  `json:"AudioBitRate"`
}

type RecorderUpdate struct {
	Title                string `json:"Title"`
	Description          string `json:"Description"`
	AutoRecording        bool   `json:"AutoRecording"`
	UseRecordingInterval bool   `json:"UseRecordingInterval"`
	RecordingInterval    int64  `json:"RecordingInterval"`
}

type RecorderCreate struct {
	Title                string `json:"Title"`
	Description          string `json:"Description"`
	AutoRecording        bool   `json:"AutoRecording"`
	UseRecordingInterval bool   `json:"UseRecordingInterval"`
	RecordingInterval    int64  `json:"RecordingInterval"`
}

type RecordersResponse struct {
	ChannelRecorders []Recorder `json:"ChannelRecorders"`
}

type Recorder struct {
	ID                   int64      `json:"Id"`
	Title                string     `json:"Title"`
	Description          string     `json:"Description"`
	Channel              Channel    `json:"Channel"`
	Project              Project    `json:"Project"`
	DstProject           Project    `json:"DstProject"`
	IsRecording          bool       `json:"IsRecording"`
	AutoRecording        bool       `json:"AutoRecording"`
	UseRecordingInterval bool       `json:"UseRecordingInterval"`
	RecordingInterval    int64      `json:"RecordingInterval"`
	RecordStartTime      time.Time  `json:"RecordStartTime"`
	SecondsRecorded      int64      `json:"SecondsRecorded"`
	Category             []Category `json:"Category"`
	IsStarting           bool       `json:"IsStarting"`
	IsStopping           bool       `json:"IsStopping"`
	CreatedAt            time.Time  `json:"CreatedAt"`
	LastUpdatedAt        time.Time  `json:"LastUpdatedAt"`
}
