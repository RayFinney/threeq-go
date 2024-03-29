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
	CreatedAt     time.Time       `json:"CreatedAt"`
	LastUpdatedAt time.Time       `json:"LastUpdatedAt"`
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
	Playouts         []struct {
		Playout Playout `json:"Playout"`
	} `json:"Playouts"`
	CreatedAt    time.Time `json:"CreatedAt"`
	LastUpdateAt time.Time `json:"LastUpdateAt"`
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

type FileEncodingSetting struct {
	UseEncoding             bool             `json:"UseEncoding"`
	UsePremiumEncoding      bool             `json:"UsePremiumEncoding"`
	UsePreProcessing        bool             `json:"UsePreProcessing"`
	UseAIServices           bool             `json:"UseAIServices"`
	ISOLanguageCode         string           `json:"ISOLanguageCode"`
	GenerateSubtitles       bool             `json:"GenerateSubtitles"`
	UseDeinterlace          bool             `json:"UseDeinterlace"`
	UseTwoPass              bool             `json:"UseTwoPass"`
	UseCMAF                 bool             `json:"UseCMAF"`
	UseBPFrames             bool             `json:"UseBPFrames"`
	PackageAudioOnlyVariant bool             `json:"PackageAudioOnlyVariant"`
	PackageForDRM           bool             `json:"PackageForDRM"`
	UseWatermark            bool             `json:"UseWatermark"`
	WatermarkURI            string           `json:"WatermarkURI"`
	WatermarkPosition       string           `json:"WatermarkPosition"` // top-left, top-right, bottom-left, bottom-right
	UseCropping             bool             `json:"UseCropping"`
	CroppingParameters      string           `json:"CroppingParameters"`
	NormalizeAudio          bool             `json:"NormalizeAudio"`
	UseFirstAudioTrack      bool             `json:"UseFirstAudioTrack"`
	UseAllAudioTracks       bool             `json:"UseAllAudioTracks"`
	FileFormats             []FileFormat     `json:"FileFormats"`
	AIFunctions             []AiconixJobType `json:"AIFunctions"`
}

type FileEncoderSettingsUpdate struct {
	UseEncoding             bool     `json:"UseEncoding"`
	UseCMAF                 bool     `json:"UseCMAF"`
	UsePremiumEncoding      bool     `json:"UsePremiumEncoding"`
	UsePreProcessing        bool     `json:"UsePreProcessing"`
	UseAIServices           bool     `json:"UseAIServices"`
	GenerateSubtitles       bool     `json:"GenerateSubtitles"`
	SourceLanguage          string   `json:"SourceLanguage,omitempty"`
	SubtitleTranslation     []string `json:"SubtitleTranslation,omitempty"`
	UseDeinterlace          bool     `json:"UseDeinterlace"`
	UseTwoPass              bool     `json:"UseTwoPass"`
	UseBPFrames             bool     `json:"UseBPFrames"`
	UseCropping             bool     `json:"UseCropping"`
	CroppingParameters      string   `json:"CroppingParameters,omitempty"`
	NormalizeAudio          bool     `json:"NormalizeAudio"`
	UseFirstAudioTrack      bool     `json:"UseFirstAudioTrack"`
	UseAllAudioTracks       bool     `json:"UseAllAudioTracks"`
	PackageAudioOnlyVariant bool     `json:"PackageAudioOnlyVariant"`
	PackageForDRM           bool     `json:"PackageForDRM"`
	WatermarkPosition       string   `json:"WatermarkPosition"` // top-left, top-right, bottom-left, bottom-right
	UseWatermark            bool     `json:"UseWatermark"`
}

type FileFormat struct {
	FileFormatID    int64   `json:"FileFormatId"`
	Label           string  `json:"Label"`
	MIMEType        string  `json:"MIME-Type"`
	VideoCodec      string  `json:"VideoCodec"`
	VideoHeight     int64   `json:"VideoHeight"`
	VideoBitRate    int64   `json:"VideoBitRate"`
	VideoProfile    string  `json:"VideoProfile"`
	VideoFPS        float64 `json:"VideoFPS"`
	AudioBitRate    int64   `json:"AudioBitRate"`
	AudioSampleRate int64   `json:"AudioSampleRate"`
	AudioChannels   int64   `json:"AudioChannels"`
}

type FileFormatUpdate struct {
	VideoBitRate    int64   `json:"VideoBitRate"`
	VideoProfile    string  `json:"VideoProfile"`
	VideoFPS        float64 `json:"VideoFPS"`
	AudioBitRate    int64   `json:"AudioBitRate"`
	AudioSampleRate int64   `json:"AudioSampleRate"`
	AudioChannels   int64   `json:"AudioChannels"`
}

type AiconixJobType struct {
	ID             int64  `json:"Id"`
	AIFunction     string `json:"AIFunction"`     // blockdetection, keywords, labeldetection, objecttracking, shotdetection, speechtotext, contentmoderation, emotion, facedetection, facialattributes, iabcategories, landmarks, ocr, persontracking
	AIResultType   string `json:"AIResultType"`   // rawresult, bestresult
	SourceLanguage string `json:"SourceLanguage"` // ISO-639-2/b (three-letter codes) Language Code
	TranslateTo    string `json:"TranslateTo"`    // ISO-639-2/b (three-letter codes) Language Code
}

type FileFormatSettings struct {
	FileFormatSettings []FileFormat `json:"FileFormatSettings"`
}

type FileEncoderPipeline struct {
	PipelineID            int64 `json:"PipelineId"`
	FileID                int64 `json:"FileId"`
	TimeFromStart         int64 `json:"TimeFromStart"`
	Duration              int64 `json:"Duration"`
	RemoveAudio           bool  `json:"RemoveAudio"`
	ExtractAudioChannel   bool  `json:"ExtractAudioChannel"`
	AudioChannelToExtract int64 `json:"AudioChannelToExtract"`
	FileEncodingSetting
}
