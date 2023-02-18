package threeqgo

import "io"

type ThreeQGo interface {
	// Utility

	// Welcome will test authentication
	Welcome() error
	SetAPIKey(apiKey string)

	// ApiKey

	// GetAPIKeyByUser use your username and password to get the apiKey
	GetAPIKeyByUser(username, password string) (string, error)

	// Projects

	// GetProjects returns a collection of available projects
	GetProjects() ([]Project, error)
	// CreateProject creates a new project
	CreateProject(project ProjectCreate) (ProjectCreateResponse, error)
	// GetProject return Project
	GetProject(id int64) (Project, error)
	// UpdateProject changes project properties
	UpdateProject(id int64, project ProjectUpdate) (Project, error)
	// DeleteProject delete a Project by id
	DeleteProject(id int64) error
	// GetProjectChannels return Channel(s) of Project if livestream project
	GetProjectChannels(id int64) ([]Channel, error)

	// Project - Custom Meta Definitions
	// TODO

	// Project - FileEncoders

	// GetFileEncoderSettings return the global settings of FileEncoding in a Video on Demand project
	GetFileEncoderSettings(projectID int64) (FileEncodingSetting, error)
	// UpdateFileEncoderSettings set the global settings for FileEncoding in a Video on Demand project
	UpdateFileEncoderSettings(projectID int64, settings FileEncoderSettingsUpdate) (FileEncodingSetting, error)
	// SetWatermarkPicture set a image as watermark
	SetWatermarkPicture(projectID int64, filename, contentType string, watermark io.Reader) error
	// GetFileFormatSettings return all global FileFormatSettings of a Video on Demand project
	GetFileFormatSettings(projectID int64) (FileFormatSettings, error)
	// GetFileFormat return a FileFormat of a Video on Demand project
	GetFileFormat(projectID int64, fileFormatID int64) (FileFormat, error)
	// UpdateFileFormat Set a FileFormat setting of a Video on Demand project
	UpdateFileFormat(projectID int64, fileFormatID int64, fileFormat FileFormatUpdate) (FileFormat, error)
	// AddEncoderSettingsFileFormat add (link) FileFormat to the FileEncoderSettings
	AddEncoderSettingsFileFormat(projectID int64, fileFormatID int64) (FileEncodingSetting, error)
	// RemoveEncoderSettingsFileFormat removes (unlink) FileFormat from the FileEncoderSettings
	RemoveEncoderSettingsFileFormat(projectID int64, fileFormatID int64) (FileEncodingSetting, error)

	// Project - FileEncoderPipeline

	// GetEncodingPipeline return the assets in the encoding pipeline of a video on demand project
	GetEncodingPipeline(projectID int64) ([]FileEncoderPipeline, error)
	// GetEncodingPipelineFile return the pipeline asset of the specified file of a video in demand project
	GetEncodingPipelineFile(projectID int64, fileID int64) (FileEncoderPipeline, error)
	// UpdateEncodingPipelineFile set the pipeline asset settings if the specified file of a video in demand project
	UpdateEncodingPipelineFile(projectID int64, fileID int64, settings FileEncoderSettingsUpdate) (FileEncoderPipeline, error)
	// AddEncodingPipelineFileFileFormat add (link) FileFormat to the FileEncoderSettings
	AddEncodingPipelineFileFileFormat(projectID int64, fileID int64, fileFormatID int64) (FileEncodingSetting, error)
	// RemoveEncodingPipelineFileFileFormat removes (unlink) FileFormat from the FileEncoderSettings
	RemoveEncodingPipelineFileFileFormat(projectID int64, fileID int64, fileFormatID int64) (FileEncodingSetting, error)

	// Channels

	// GetChannels return a collection of available channels
	GetChannels() ([]Channel, error)
	// GetChannel return a Channel by id
	GetChannel(id int64) (Channel, error)

	// Channel - Picture
	// TODO

	// Channel - Metadata
	// TODO

	// Channel - Ingest
	// TODO

	// Channel - Output
	// TODO

	// Channel - Embed
	// TODO

	// Channel - Transcoder
	// TODO

	// Channel - Distributions
	// TODO

	// Channel - Timeshift2VoD
	// TODO

	// Channel - Recorder

	// GetChannelRecorders return enabled channel recorder
	GetChannelRecorders() ([]Recorder, error)
	// GetChannelRecorder return a Recorder by id
	GetChannelRecorder(channelID, recorderID int64) (Recorder, error)
	// UpdateChannelRecorder Edit a Recorder
	UpdateChannelRecorder(channelID, recorderID int64, recorder RecorderUpdate) (Recorder, error)
	// DeleteChannelRecorder delete a Recorder by id
	DeleteChannelRecorder(channelID, recorderID int64) error
	// CreateChannelRecorder create a new Recorder
	CreateChannelRecorder(channelID, dstProjectID int64, recorder RecorderCreate) (Recorder, error)
	// ChannelRecorderAddCategory add(link) a Category to the Recorder
	ChannelRecorderAddCategory(channelID, recorderID, categoryID int64) (Recorder, error)
	// ChannelRecorderRemoveCategory remove(unlink) a Category from the Recorder
	ChannelRecorderRemoveCategory(channelID, recorderID, categoryID int64) (Recorder, error)

	// Channel - Purge Timeshift
	// TODO

	// Files

	// GetFiles returns a collection of File in Project
	GetFiles(projectID int64, queryParams FileSearchOptions) ([]File, error)
	// GetFile return a File by id in Project
	GetFile(projectID, fileID int64) (File, error)

	// File - Metadata
	// TODO

	// File - Metadata - Categories
	// TODO

	// File - Metadata - Videotype
	// TODO

	// File - Playout
	// TODO

	// File - Output
	// TODO

	// File - Picture
	// TODO

	// File - SubTitle
	// TODO

	// File - Most viewed
	// TODO

	// File - Merge files
	// TODO

	// File - Video intelligence output
	// TODO

	// File - Encoding Progress

	// GetEncodingProgress return the EncodingProgress of file processing
	GetEncodingProgress(projectID, fileID int64) (EncodingProgress, error)

	// File - Metadata by ProgrammId
	// TODO

	// File - Text2Speech
	// TODO

	// File - Release Status
	// TODO

	// Categories
	// TODO

	// Podcasts
	// TODO

	// Podcast-Episodes
	// TODO

	// Reporting - Project
	// TODO

	// Reporting - File
	// TODO

	// Reporting - User
	// TODO

	// Reporting - Country, Device
	// TODO

	// Playlists
	// TODO

	// Analytics - Dashboard
	// TODO

	// Analytics - Audience
	// TODO

	// Analytics - Locations
	// TODO

	// Analytics - Ads
	// TODO

	// Analytics - UserToken overview
	// TODO

	// Services - Configuration
	// TODO

	// Services - User-Playlists
	// TODO
}
