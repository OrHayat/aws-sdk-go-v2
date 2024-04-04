// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Object specifying a stream’s audio configuration, as set up by the broadcaster
// (usually in an encoder). This is part of the IngestConfiguration object and
// used for monitoring stream health.
type AudioConfiguration struct {

	// Number of audio channels.
	Channels int64

	// Codec used for the audio encoding.
	Codec *string

	// Number of audio samples recorded per second.
	SampleRate int64

	// The expected ingest bitrate (bits per second). This is configured in the
	// encoder.
	TargetBitrate int64

	noSmithyDocumentSerde
}

// Error related to a specific channel, specified by its ARN.
type BatchError struct {

	// Channel ARN.
	Arn *string

	// Error code.
	Code *string

	// Error message, determined by the application.
	Message *string

	noSmithyDocumentSerde
}

// Error for a request in the batch for BatchStartViewerSessionRevocation. Each
// error is related to a specific channel-ARN and viewer-ID pair.
type BatchStartViewerSessionRevocationError struct {

	// Channel ARN.
	//
	// This member is required.
	ChannelArn *string

	// The ID of the viewer session to revoke.
	//
	// This member is required.
	ViewerId *string

	// Error code.
	Code *string

	// Error message, determined by the application.
	Message *string

	noSmithyDocumentSerde
}

// A viewer session to revoke in the call to BatchStartViewerSessionRevocation .
type BatchStartViewerSessionRevocationViewerSession struct {

	// The ARN of the channel associated with the viewer session to revoke.
	//
	// This member is required.
	ChannelArn *string

	// The ID of the viewer associated with the viewer session to revoke. Do not use
	// this field for personally identifying, confidential, or sensitive information.
	//
	// This member is required.
	ViewerId *string

	// An optional filter on which versions of the viewer session to revoke. All
	// versions less than or equal to the specified version will be revoked. Default:
	// 0.
	ViewerSessionVersionsLessThanOrEqualTo int32

	noSmithyDocumentSerde
}

// Object specifying a channel.
type Channel struct {

	// Channel ARN.
	Arn *string

	// Whether the channel is private (enabled for playback authorization). Default:
	// false .
	Authorized bool

	// Channel ingest endpoint, part of the definition of an ingest server, used when
	// you set up streaming software.
	IngestEndpoint *string

	// Whether the channel allows insecure RTMP ingest. Default: false .
	InsecureIngest bool

	// Channel latency mode. Use NORMAL to broadcast and deliver live video up to Full
	// HD. Use LOW for near-real-time interaction with viewers. Default: LOW .
	LatencyMode ChannelLatencyMode

	// Channel name.
	Name *string

	// Playback-restriction-policy ARN. A valid ARN value here both specifies the ARN
	// and enables playback restriction. Default: "" (empty string, no playback
	// restriction policy is applied).
	PlaybackRestrictionPolicyArn *string

	// Channel playback URL.
	PlaybackUrl *string

	// Optional transcode preset for the channel. This is selectable only for
	// ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default
	// preset is HIGHER_BANDWIDTH_DELIVERY . For other channel types ( BASIC and
	// STANDARD ), preset is the empty string ( "" ).
	Preset TranscodePreset

	// Recording-configuration ARN. A valid ARN value here both specifies the ARN and
	// enables recording. Default: "" (empty string, recording is disabled).
	RecordingConfigurationArn *string

	// Specifies the endpoint and optional passphrase for streaming with the SRT
	// protocol.
	Srt *Srt

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	// Channel type, which determines the allowable resolution and bitrate. If you
	// exceed the allowable input resolution or bitrate, the stream probably will
	// disconnect immediately. Default: STANDARD . For details, see Channel Types (https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/channel-types.html)
	// .
	Type ChannelType

	noSmithyDocumentSerde
}

// Summary information about a channel.
type ChannelSummary struct {

	// Channel ARN.
	Arn *string

	// Whether the channel is private (enabled for playback authorization). Default:
	// false .
	Authorized bool

	// Whether the channel allows insecure RTMP ingest. Default: false .
	InsecureIngest bool

	// Channel latency mode. Use NORMAL to broadcast and deliver live video up to Full
	// HD. Use LOW for near-real-time interaction with viewers. Default: LOW .
	LatencyMode ChannelLatencyMode

	// Channel name.
	Name *string

	// Playback-restriction-policy ARN. A valid ARN value here both specifies the ARN
	// and enables playback restriction. Default: "" (empty string, no playback
	// restriction policy is applied).
	PlaybackRestrictionPolicyArn *string

	// Optional transcode preset for the channel. This is selectable only for
	// ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default
	// preset is HIGHER_BANDWIDTH_DELIVERY . For other channel types ( BASIC and
	// STANDARD ), preset is the empty string ( "" ).
	Preset TranscodePreset

	// Recording-configuration ARN. A valid ARN value here both specifies the ARN and
	// enables recording. Default: "" (empty string, recording is disabled).
	RecordingConfigurationArn *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	// Channel type, which determines the allowable resolution and bitrate. If you
	// exceed the allowable input resolution or bitrate, the stream probably will
	// disconnect immediately. Default: STANDARD . For details, see Channel Types (https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/channel-types.html)
	// .
	Type ChannelType

	noSmithyDocumentSerde
}

// A complex type that describes a location where recorded videos will be stored.
// Each member represents a type of destination configuration. For recording, you
// define one and only one type of destination configuration.
type DestinationConfiguration struct {

	// An S3 destination configuration where recorded videos will be stored.
	S3 *S3DestinationConfiguration

	noSmithyDocumentSerde
}

// Object specifying the ingest configuration set up by the broadcaster, usually
// in an encoder.
type IngestConfiguration struct {

	// Encoder settings for audio.
	Audio *AudioConfiguration

	// Encoder settings for video.
	Video *VideoConfiguration

	noSmithyDocumentSerde
}

// A key pair used to sign and validate a playback authorization token.
type PlaybackKeyPair struct {

	// Key-pair ARN.
	Arn *string

	// Key-pair identifier.
	Fingerprint *string

	// Playback-key-pair name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about a playback key pair.
type PlaybackKeyPairSummary struct {

	// Key-pair ARN.
	Arn *string

	// Playback-key-pair name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object representing a policy to constrain playback by country and/or origin
// sites.
type PlaybackRestrictionPolicy struct {

	// A list of country codes that control geoblocking restriction. Allowed values
	// are the officially assigned ISO 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	// codes. Default: All countries (an empty array).
	//
	// This member is required.
	AllowedCountries []string

	// A list of origin sites that control CORS restriction. Allowed values are the
	// same as valid values of the Origin header defined at
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin)
	// . Default: All origins (an empty array).
	//
	// This member is required.
	AllowedOrigins []string

	// Playback-restriction-policy ARN
	//
	// This member is required.
	Arn *string

	// Whether channel playback is constrained by origin site. Default: false .
	EnableStrictOriginEnforcement *bool

	// Playback-restriction-policy name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about a PlaybackRestrictionPolicy.
type PlaybackRestrictionPolicySummary struct {

	// A list of country codes that control geoblocking restriction. Allowed values
	// are the officially assigned ISO 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	// codes. Default: All countries (an empty array).
	//
	// This member is required.
	AllowedCountries []string

	// A list of origin sites that control CORS restriction. Allowed values are the
	// same as valid values of the Origin header defined at
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin)
	// . Default: All origins (an empty array).
	//
	// This member is required.
	AllowedOrigins []string

	// Playback-restriction-policy ARN
	//
	// This member is required.
	Arn *string

	// Whether channel playback is constrained by origin site. Default: false .
	EnableStrictOriginEnforcement *bool

	// Playback-restriction-policy name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object representing a configuration to record a channel stream.
type RecordingConfiguration struct {

	// Recording-configuration ARN.
	//
	// This member is required.
	Arn *string

	// A complex type that contains information about where recorded video will be
	// stored.
	//
	// This member is required.
	DestinationConfiguration *DestinationConfiguration

	// Indicates the current state of the recording configuration. When the state is
	// ACTIVE , the configuration is ready for recording a channel stream.
	//
	// This member is required.
	State RecordingConfigurationState

	// Recording-configuration name. The value does not need to be unique.
	Name *string

	// If a broadcast disconnects and then reconnects within the specified interval,
	// the multiple streams will be considered a single broadcast and merged together.
	// Default: 0.
	RecordingReconnectWindowSeconds int32

	// Object that describes which renditions should be recorded for a stream.
	RenditionConfiguration *RenditionConfiguration

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	// A complex type that allows you to enable/disable the recording of thumbnails
	// for a live session and modify the interval at which thumbnails are generated for
	// the live session.
	ThumbnailConfiguration *ThumbnailConfiguration

	noSmithyDocumentSerde
}

// Summary information about a RecordingConfiguration.
type RecordingConfigurationSummary struct {

	// Recording-configuration ARN.
	//
	// This member is required.
	Arn *string

	// A complex type that contains information about where recorded video will be
	// stored.
	//
	// This member is required.
	DestinationConfiguration *DestinationConfiguration

	// Indicates the current state of the recording configuration. When the state is
	// ACTIVE , the configuration is ready for recording a channel stream.
	//
	// This member is required.
	State RecordingConfigurationState

	// Recording-configuration name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Object that describes which renditions should be recorded for a stream.
type RenditionConfiguration struct {

	// Indicates which set of renditions are recorded for a stream. For BASIC
	// channels, the CUSTOM value has no effect. If CUSTOM is specified, a set of
	// renditions must be specified in the renditions field. Default: ALL .
	RenditionSelection RenditionConfigurationRenditionSelection

	// Indicates which renditions are recorded for a stream, if renditionSelection is
	// CUSTOM ; otherwise, this field is irrelevant. The selected renditions are
	// recorded if they are available during the stream. If a selected rendition is
	// unavailable, the best available rendition is recorded. For details on the
	// resolution dimensions of each rendition, see Auto-Record to Amazon S3 (https://docs.aws.amazon.com/ivs/latest/userguide/record-to-s3.html)
	// .
	Renditions []RenditionConfigurationRendition

	noSmithyDocumentSerde
}

// A complex type that describes an S3 location where recorded videos will be
// stored.
type S3DestinationConfiguration struct {

	// Location (S3 bucket name) where recorded videos will be stored.
	//
	// This member is required.
	BucketName *string

	noSmithyDocumentSerde
}

// Specifies information needed to stream using the SRT protocol.
type Srt struct {

	// The endpoint to be used when streaming with IVS using the SRT protocol.
	Endpoint *string

	// Auto-generated passphrase to enable encryption. This field is applicable only
	// if the end user has not enabled the insecureIngest option for the channel.
	Passphrase *string

	noSmithyDocumentSerde
}

// Specifies a live video stream that has been ingested and distributed.
type Stream struct {

	// Channel ARN for the stream.
	ChannelArn *string

	// The stream’s health.
	Health StreamHealth

	// URL of the master playlist, required by the video player to play the HLS stream.
	PlaybackUrl *string

	// Time of the stream’s start. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	StartTime *time.Time

	// The stream’s state. Do not rely on the OFFLINE state, as the API may not return
	// it; instead, a "NotBroadcasting" error will indicate that the stream is not
	// live.
	State StreamState

	// Unique identifier for a live or previously live stream in the specified channel.
	StreamId *string

	// A count of concurrent views of the stream. Typically, a new view appears in
	// viewerCount within 15 seconds of when video playback starts and a view is
	// removed from viewerCount within 1 minute of when video playback ends. A value
	// of -1 indicates that the request timed out; in this case, retry.
	ViewerCount int64

	noSmithyDocumentSerde
}

// Object specifying a stream’s events. For a list of events, see Using Amazon
// EventBridge with Amazon IVS (https://docs.aws.amazon.com/ivs/latest/userguide/eventbridge.html)
// .
type StreamEvent struct {

	// Time when the event occurred. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	EventTime *time.Time

	// Name that identifies the stream event within a type .
	Name *string

	// Logical group for certain events.
	Type *string

	noSmithyDocumentSerde
}

// Object specifying the stream attribute on which to filter.
type StreamFilters struct {

	// The stream’s health.
	Health StreamHealth

	noSmithyDocumentSerde
}

// Object specifying a stream key.
type StreamKey struct {

	// Stream-key ARN.
	Arn *string

	// Channel ARN for the stream.
	ChannelArn *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	// Stream-key value.
	Value *string

	noSmithyDocumentSerde
}

// Summary information about a stream key.
type StreamKeySummary struct {

	// Stream-key ARN.
	Arn *string

	// Channel ARN for the stream.
	ChannelArn *string

	// Tags attached to the resource. Array of 1-50 maps, each of the form
	// string:string (key:value) . See Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Object that captures the Amazon IVS configuration that the customer
// provisioned, the ingest configurations that the broadcaster used, and the most
// recent Amazon IVS stream events it encountered.
type StreamSession struct {

	// The properties of the channel at the time of going live.
	Channel *Channel

	// Time when the channel went offline. This is an ISO 8601 timestamp; note that
	// this is returned as a string. For live streams, this is NULL .
	EndTime *time.Time

	// The properties of the incoming RTMP stream for the stream.
	IngestConfiguration *IngestConfiguration

	// The properties of recording the live stream.
	RecordingConfiguration *RecordingConfiguration

	// Time when the channel went live. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	StartTime *time.Time

	// Unique identifier for a live or previously live stream in the specified channel.
	StreamId *string

	// List of Amazon IVS events that the stream encountered. The list is sorted by
	// most recent events and contains up to 500 events. For Amazon IVS events, see
	// Using Amazon EventBridge with Amazon IVS (https://docs.aws.amazon.com/ivs/latest/userguide/eventbridge.html)
	// .
	TruncatedEvents []StreamEvent

	noSmithyDocumentSerde
}

// Summary information about a stream session.
type StreamSessionSummary struct {

	// Time when the channel went offline. This is an ISO 8601 timestamp; note that
	// this is returned as a string. For live streams, this is NULL .
	EndTime *time.Time

	// If true , this stream encountered a quota breach or failure.
	HasErrorEvent bool

	// Time when the channel went live. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	StartTime *time.Time

	// Unique identifier for a live or previously live stream in the specified channel.
	StreamId *string

	noSmithyDocumentSerde
}

// Summary information about a stream.
type StreamSummary struct {

	// Channel ARN for the stream.
	ChannelArn *string

	// The stream’s health.
	Health StreamHealth

	// Time of the stream’s start. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	StartTime *time.Time

	// The stream’s state. Do not rely on the OFFLINE state, as the API may not return
	// it; instead, a "NotBroadcasting" error will indicate that the stream is not
	// live.
	State StreamState

	// Unique identifier for a live or previously live stream in the specified channel.
	StreamId *string

	// A count of concurrent views of the stream. Typically, a new view appears in
	// viewerCount within 15 seconds of when video playback starts and a view is
	// removed from viewerCount within 1 minute of when video playback ends. A value
	// of -1 indicates that the request timed out; in this case, retry.
	ViewerCount int64

	noSmithyDocumentSerde
}

// An object representing a configuration of thumbnails for recorded video.
type ThumbnailConfiguration struct {

	// Thumbnail recording mode. Default: INTERVAL .
	RecordingMode RecordingMode

	// Indicates the desired resolution of recorded thumbnails. Thumbnails are
	// recorded at the selected resolution if the corresponding rendition is available
	// during the stream; otherwise, they are recorded at source resolution. For more
	// information about resolution values and their corresponding height and width
	// dimensions, see Auto-Record to Amazon S3 (https://docs.aws.amazon.com/ivs/latest/userguide/record-to-s3.html)
	// . Default: Null (source resolution is returned).
	Resolution ThumbnailConfigurationResolution

	// Indicates the format in which thumbnails are recorded. SEQUENTIAL records all
	// generated thumbnails in a serial manner, to the media/thumbnails directory.
	// LATEST saves the latest thumbnail in media/latest_thumbnail/thumb.jpg and
	// overwrites it at the interval specified by targetIntervalSeconds . You can
	// enable both SEQUENTIAL and LATEST . Default: SEQUENTIAL .
	Storage []ThumbnailConfigurationStorage

	// The targeted thumbnail-generation interval in seconds. This is configurable
	// (and required) only if recordingMode is INTERVAL . Default: 60. Important: For
	// the BASIC channel type, setting a value for targetIntervalSeconds does not
	// guarantee that thumbnails are generated at the specified interval. For
	// thumbnails to be generated at the targetIntervalSeconds interval, the
	// IDR/Keyframe value for the input video must be less than the
	// targetIntervalSeconds value. See  Amazon IVS Streaming Configuration (https://docs.aws.amazon.com/ivs/latest/userguide/streaming-config.html)
	// for information on setting IDR/Keyframe to the recommended value in
	// video-encoder settings.
	TargetIntervalSeconds *int64

	noSmithyDocumentSerde
}

// Object specifying a stream’s video configuration, as set up by the broadcaster
// (usually in an encoder). This is part of the IngestConfiguration object and
// used for monitoring stream health.
type VideoConfiguration struct {

	// Indicates the degree of required decoder performance for a profile. Normally
	// this is set automatically by the encoder. For details, see the H.264
	// specification.
	AvcLevel *string

	// Indicates to the decoder the requirements for decoding the stream. For
	// definitions of the valid values, see the H.264 specification.
	AvcProfile *string

	// Codec used for the video encoding.
	Codec *string

	// Software or hardware used to encode the video.
	Encoder *string

	// The expected ingest bitrate (bits per second). This is configured in the
	// encoder.
	TargetBitrate int64

	// The expected ingest framerate. This is configured in the encoder.
	TargetFramerate int64

	// Video-resolution height in pixels.
	VideoHeight int64

	// Video-resolution width in pixels.
	VideoWidth int64

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
