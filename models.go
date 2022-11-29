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

type FileCreate struct {
}

type FileUpdate struct {
}

type FilesResponse struct {
	Files []File `json:"Files"`
}

type FileSearchOptions struct {
	IncludeDeleted    bool
	IncludeMetadata   bool
	IncludePlayouts   bool
	IncludeProperties bool
	ReleaseStatus     string
	OrderBy           string
	Sort              string
	Period            string
	CategoryId        int64
	VideoTypeId       int64
	Offset            int64
	Limit             int64
}

type File struct {
	ID               int64          `json:"Id"`
	Name             string         `json:"Name"`
	IsFinished       bool           `json:"IsFinished"`
	IsEncoding       bool           `json:"IsEncoding"`
	UseEncoding      bool           `json:"UseEncoding"`
	EncodingPriority int64          `json:"EncodingPriority"`
	HasErrors        bool           `json:"HasErrors"`
	ErrorMessage     string         `json:"ErrorMessage"`
	Properties       FileProperties `json:"Properties"`
	MetaData         MetaData       `json:"MetaData"`
	Playouts         []Playout      `json:"Playouts"`
	CreatedAt        time.Time      `json:"CreatedAt"`
	LastUpdateAt     time.Time      `json:"LastUpdateAt"`
}

type FileProperties struct {
	Length          float64 `json:"Length"`
	Size            string  `json:"Size"`
	VideoFormat     string  `json:"VideoFormat"`
	VideoBitRate    int64   `json:"VideoBitRate"`
	VideoWidth      int64   `json:"VideoWidth"`
	VideoHeight     int64   `json:"VideoHeight"`
	VideoFPS        float64 `json:"VideoFPS"`
	AudioFormat     string  `json:"AudioFormat"`
	AudioBitRate    int64   `json:"AudioBitRate"`
	AudioSampleRate int64   `json:"AudioSampleRate"`
	AudioChannels   int64   `json:"AudioChannels"`
}

type MetaData struct {
	StandardFilePicture    StandardFilePicture `json:"StandardFilePicture"`
	Title                  string              `json:"Title"`
	Description            string              `json:"Description"`
	DisplayTitle           string              `json:"DisplayTitle"`
	DisplayTitleSecondLine string              `json:"DisplayTitleSecondLine"`
	Tags                   string              `json:"Tags"`
	Genre                  string              `json:"Genre"`
	Studio                 string              `json:"Studio"`
	IsPublicAt             time.Time           `json:"IsPublicAt"`
	IsPublicUntil          time.Time           `json:"IsPublicUntil"`
	Category               []Category          `json:"Category"`
	IABCategory            []IABCategory       `json:"IABCategory"`
	Share                  []Share             `json:"Share"`
	Series                 string              `json:"Series"`
	ProductionCountry      string              `json:"ProductionCountry"`
	NativeLanguage         string              `json:"NativeLanguage"`
	ProgramID              string              `json:"ProgramId"`
	Source                 string              `json:"Source"`
	Licensor               string              `json:"Licensor"`
	LicenseArea            string              `json:"LicenseArea"`
	RelationShip           string              `json:"RelationShip"`
	Rating                 float64             `json:"Rating"`
	Latitude               float64             `json:"Latitude"`
	Longitude              float64             `json:"Longitude"`
	VideoType              []VideoType         `json:"VideoType"`
	OriginalFileName       string              `json:"OriginalFileName"`
	Deeplink               string              `json:"Deeplink"`
	MetaPictureURI         string              `json:"MetaPictureURI"`
	CustomMetadata         map[string]string
}

type StandardFilePicture struct {
	FilePictureId int64  `json:"FilePictureId"`
	URI           string `json:"URI"`
	ThumbURI      string `json:"ThumbURI"`
	SrcPictureURI string `json:"SrcPictureURI"`
	IsStandard    bool   `json:"IsStandard"`
}

type IABCategory struct {
	ID    int64  `json:"Id"`
	Label string `json:"Label"`
	TABId string `json:"TABId"`
}

type Share struct {
	ID          int64  `json:"Id"`
	Label       string `json:"Label"`
	Description string `json:"Description"`
	Author      string `json:"Author"`
	AccessKey   string `json:"AccessKey"`
	ThumbURI    string `json:"ThumbURI"`
}

type VideoType struct {
	ID    int64  `json:"Id"`
	Label string `json:"Label"`
}

type EncodingProgress struct {
	ID               int64   `json:"Id"`
	Format           string  `json:"Format"`
	FPS              float64 `json:"FPS"`
	RealTime         int64   `json:"RealTime"`
	EncodingProgress int64   `json:"EncodingProgress"`
	UploadProgress   int64   `json:"UploadProgress"`
	HasErrors        bool    `json:"HasErrors"`
	IsFinished       bool    `json:"IsFinished"`
	IsStored         bool    `json:"IsStored"`
}

type Playout struct {
	ID    string `json:"Id"`
	Label string `json:"Label"`
}
