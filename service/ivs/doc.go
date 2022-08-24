// Code generated by smithy-go-codegen DO NOT EDIT.

// Package ivs provides the API client, operations, and parameter types for Amazon
// Interactive Video Service.
//
// Introduction The Amazon Interactive Video Service (IVS) API is REST compatible,
// using a standard HTTP API and an Amazon Web Services EventBridge event stream
// for responses. JSON is used for both requests and responses, including errors.
// The API is an Amazon Web Services regional service. For a list of supported
// regions and Amazon IVS HTTPS service endpoints, see the Amazon IVS page
// (https://docs.aws.amazon.com/general/latest/gr/ivs.html) in the Amazon Web
// Services General Reference. All API request parameters and URLs are case
// sensitive. For a summary of notable documentation changes in each release, see
// Document History
// (https://docs.aws.amazon.com/ivs/latest/userguide/doc-history.html). Allowed
// Header Values
//
// * Accept:  application/json
//
// * Accept-Encoding:  gzip, deflate
//
// *
// Content-Type: application/json
//
// Resources The following resources contain
// information about your IVS live stream (see  Getting Started with Amazon IVS
// (https://docs.aws.amazon.com/ivs/latest/userguide/getting-started.html)):
//
// *
// Channel — Stores configuration data related to your live stream. You first
// create a channel and then use the channel’s stream key to start your live
// stream. See the Channel endpoints for more information.
//
// * Stream key — An
// identifier assigned by Amazon IVS when you create a channel, which is then used
// to authorize streaming. See the StreamKey endpoints for more information. Treat
// the stream key like a secret, since it allows anyone to stream to the
// channel.
//
// * Playback key pair — Video playback may be restricted using
// playback-authorization tokens, which use public-key encryption. A playback key
// pair is the public-private pair of keys used to sign and validate the
// playback-authorization token. See the PlaybackKeyPair endpoints for more
// information.
//
// * Recording configuration — Stores configuration related to
// recording a live stream and where to store the recorded content. Multiple
// channels can reference the same recording configuration. See the Recording
// Configuration endpoints for more information.
//
// Tagging A tag is a metadata label
// that you assign to an Amazon Web Services resource. A tag comprises a key and a
// value, both set by you. For example, you might set a tag as topic:nature to
// label a particular video category. See Tagging Amazon Web Services Resources
// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) for more
// information, including restrictions that apply to tags and "Tag naming limits
// and requirements"; Amazon IVS has no service-specific constraints beyond what is
// documented there. Tags can help you identify and organize your Amazon Web
// Services resources. For example, you can use the same tag for different
// resources to indicate that they are related. You can also use tags to manage
// access (see  Access Tags
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html)). The Amazon
// IVS API has these tag-related endpoints: TagResource, UntagResource, and
// ListTagsForResource. The following resources support tagging: Channels, Stream
// Keys, Playback Key Pairs, and Recording Configurations. At most 50 tags can be
// applied to a resource. Authentication versus Authorization Note the differences
// between these concepts:
//
// * Authentication is about verifying identity. You need
// to be authenticated to sign Amazon IVS API requests.
//
// * Authorization is about
// granting permissions. Your IAM roles need to have permissions for Amazon IVS API
// requests. In addition, authorization is needed to view Amazon IVS private
// channels
// (https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html).
// (Private channels are channels that are enabled for "playback
// authorization.")
//
// Authentication All Amazon IVS API requests must be
// authenticated with a signature. The Amazon Web Services Command-Line Interface
// (CLI) and Amazon IVS Player SDKs take care of signing the underlying API calls
// for you. However, if your application calls the Amazon IVS API directly, it’s
// your responsibility to sign the requests. You generate a signature using valid
// Amazon Web Services credentials that have permission to perform the requested
// action. For example, you must sign PutMetadata requests with a signature
// generated from an IAM user account that has the ivs:PutMetadata permission. For
// more information:
//
// * Authentication and generating signatures — See
// Authenticating Requests (Amazon Web Services Signature Version 4)
// (https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon Web Services General Reference.
//
// * Managing Amazon IVS permissions
// — See Identity and Access Management
// (https://docs.aws.amazon.com/ivs/latest/userguide/security-iam.html) on the
// Security page of the Amazon IVS User Guide.
//
// Amazon Resource Names (ARNs) ARNs
// uniquely identify AWS resources. An ARN is required when you need to specify a
// resource unambiguously across all of AWS, such as in IAM policies and API calls.
// For more information, see Amazon Resource Names
// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
// the AWS General Reference. Channel Endpoints
//
// * CreateChannel — Creates a new
// channel and an associated stream key to start streaming.
//
// * GetChannel — Gets
// the channel configuration for the specified channel ARN.
//
// * BatchGetChannel —
// Performs GetChannel on multiple ARNs simultaneously.
//
// * ListChannels — Gets
// summary information about all channels in your account, in the Amazon Web
// Services region where the API request is processed. This list can be filtered to
// match a specified name or recording-configuration ARN. Filters are mutually
// exclusive and cannot be used together. If you try to use both filters, you will
// get an error (409 Conflict Exception).
//
// * UpdateChannel — Updates a channel's
// configuration. This does not affect an ongoing stream of this channel. You must
// stop and restart the stream for the changes to take effect.
//
// * DeleteChannel —
// Deletes the specified channel.
//
// # StreamKey Endpoints
//
// * CreateStreamKey — Creates
// a stream key, used to initiate a stream, for the specified channel ARN.
//
// *
// GetStreamKey — Gets stream key information for the specified ARN.
//
// *
// BatchGetStreamKey — Performs GetStreamKey on multiple ARNs simultaneously.
//
// *
// ListStreamKeys — Gets summary information about stream keys for the specified
// channel.
//
// * DeleteStreamKey — Deletes the stream key for the specified ARN, so
// it can no longer be used to stream.
//
// # Stream Endpoints
//
// * GetStream — Gets
// information about the active (live) stream on a specified channel.
//
// *
// GetStreamSession — Gets metadata on a specified stream.
//
// * ListStreams — Gets
// summary information about live streams in your account, in the Amazon Web
// Services region where the API request is processed.
//
// * ListStreamSessions — Gets
// a summary of current and previous streams for a specified channel in your
// account, in the AWS region where the API request is processed.
//
// * StopStream —
// Disconnects the incoming RTMPS stream for the specified channel. Can be used in
// conjunction with DeleteStreamKey to prevent further streaming to a channel.
//
// *
// PutMetadata — Inserts metadata into the active stream of the specified channel.
// At most 5 requests per second per channel are allowed, each with a maximum 1 KB
// payload. (If 5 TPS is not sufficient for your needs, we recommend batching your
// data into a single PutMetadata call.) At most 155 requests per second per
// account are allowed.
//
// PlaybackKeyPair Endpoints For more information, see
// Setting Up Private Channels
// (https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html) in the
// Amazon IVS User Guide.
//
// * ImportPlaybackKeyPair — Imports the public portion of
// a new key pair and returns its arn and fingerprint. The privateKey can then be
// used to generate viewer authorization tokens, to grant viewers access to private
// channels (channels enabled for playback authorization).
//
// * GetPlaybackKeyPair —
// Gets a specified playback authorization key pair and returns the arn and
// fingerprint. The privateKey held by the caller can be used to generate viewer
// authorization tokens, to grant viewers access to private channels.
//
// *
// ListPlaybackKeyPairs — Gets summary information about playback key pairs.
//
// *
// DeletePlaybackKeyPair — Deletes a specified authorization key pair. This
// invalidates future viewer tokens generated using the key pair’s
// privateKey.
//
// # RecordingConfiguration Endpoints
//
// * CreateRecordingConfiguration —
// Creates a new recording configuration, used to enable recording to Amazon S3.
//
// *
// GetRecordingConfiguration — Gets the recording-configuration metadata for the
// specified ARN.
//
// * ListRecordingConfigurations — Gets summary information about
// all recording configurations in your account, in the Amazon Web Services region
// where the API request is processed.
//
// * DeleteRecordingConfiguration — Deletes
// the recording configuration for the specified ARN.
//
// Amazon Web Services Tags
// Endpoints
//
// * TagResource — Adds or updates tags for the Amazon Web Services
// resource with the specified ARN.
//
// * UntagResource — Removes tags from the
// resource with the specified ARN.
//
// * ListTagsForResource — Gets information about
// Amazon Web Services tags for the specified ARN.
package ivs
