// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AggregationFunction string

// Enum values for AggregationFunction
const (
	AggregationFunctionAvg AggregationFunction = "AVG"
	AggregationFunctionSum AggregationFunction = "SUM"
)

// Values returns all known values for AggregationFunction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AggregationFunction) Values() []AggregationFunction {
	return []AggregationFunction{
		"AVG",
		"SUM",
	}
}

type AlertStatus string

// Enum values for AlertStatus
const (
	AlertStatusActive   AlertStatus = "ACTIVE"
	AlertStatusInactive AlertStatus = "INACTIVE"
)

// Values returns all known values for AlertStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AlertStatus) Values() []AlertStatus {
	return []AlertStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type AlertType string

// Enum values for AlertType
const (
	AlertTypeSns    AlertType = "SNS"
	AlertTypeLambda AlertType = "LAMBDA"
)

// Values returns all known values for AlertType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AlertType) Values() []AlertType {
	return []AlertType{
		"SNS",
		"LAMBDA",
	}
}

type AnomalyDetectionTaskStatus string

// Enum values for AnomalyDetectionTaskStatus
const (
	AnomalyDetectionTaskStatusPending          AnomalyDetectionTaskStatus = "PENDING"
	AnomalyDetectionTaskStatusInProgress       AnomalyDetectionTaskStatus = "IN_PROGRESS"
	AnomalyDetectionTaskStatusCompleted        AnomalyDetectionTaskStatus = "COMPLETED"
	AnomalyDetectionTaskStatusFailed           AnomalyDetectionTaskStatus = "FAILED"
	AnomalyDetectionTaskStatusFailedToSchedule AnomalyDetectionTaskStatus = "FAILED_TO_SCHEDULE"
)

// Values returns all known values for AnomalyDetectionTaskStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyDetectionTaskStatus) Values() []AnomalyDetectionTaskStatus {
	return []AnomalyDetectionTaskStatus{
		"PENDING",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
		"FAILED_TO_SCHEDULE",
	}
}

type AnomalyDetectorFailureType string

// Enum values for AnomalyDetectorFailureType
const (
	AnomalyDetectorFailureTypeActivationFailure         AnomalyDetectorFailureType = "ACTIVATION_FAILURE"
	AnomalyDetectorFailureTypeBackTestActivationFailure AnomalyDetectorFailureType = "BACK_TEST_ACTIVATION_FAILURE"
	AnomalyDetectorFailureTypeDeletionFailure           AnomalyDetectorFailureType = "DELETION_FAILURE"
	AnomalyDetectorFailureTypeDeactivationFailure       AnomalyDetectorFailureType = "DEACTIVATION_FAILURE"
)

// Values returns all known values for AnomalyDetectorFailureType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyDetectorFailureType) Values() []AnomalyDetectorFailureType {
	return []AnomalyDetectorFailureType{
		"ACTIVATION_FAILURE",
		"BACK_TEST_ACTIVATION_FAILURE",
		"DELETION_FAILURE",
		"DEACTIVATION_FAILURE",
	}
}

type AnomalyDetectorStatus string

// Enum values for AnomalyDetectorStatus
const (
	AnomalyDetectorStatusActive             AnomalyDetectorStatus = "ACTIVE"
	AnomalyDetectorStatusActivating         AnomalyDetectorStatus = "ACTIVATING"
	AnomalyDetectorStatusDeleting           AnomalyDetectorStatus = "DELETING"
	AnomalyDetectorStatusFailed             AnomalyDetectorStatus = "FAILED"
	AnomalyDetectorStatusInactive           AnomalyDetectorStatus = "INACTIVE"
	AnomalyDetectorStatusLearning           AnomalyDetectorStatus = "LEARNING"
	AnomalyDetectorStatusBackTestActivating AnomalyDetectorStatus = "BACK_TEST_ACTIVATING"
	AnomalyDetectorStatusBackTestActive     AnomalyDetectorStatus = "BACK_TEST_ACTIVE"
	AnomalyDetectorStatusBackTestComplete   AnomalyDetectorStatus = "BACK_TEST_COMPLETE"
	AnomalyDetectorStatusDeactivated        AnomalyDetectorStatus = "DEACTIVATED"
	AnomalyDetectorStatusDeactivating       AnomalyDetectorStatus = "DEACTIVATING"
)

// Values returns all known values for AnomalyDetectorStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnomalyDetectorStatus) Values() []AnomalyDetectorStatus {
	return []AnomalyDetectorStatus{
		"ACTIVE",
		"ACTIVATING",
		"DELETING",
		"FAILED",
		"INACTIVE",
		"LEARNING",
		"BACK_TEST_ACTIVATING",
		"BACK_TEST_ACTIVE",
		"BACK_TEST_COMPLETE",
		"DEACTIVATED",
		"DEACTIVATING",
	}
}

type CSVFileCompression string

// Enum values for CSVFileCompression
const (
	CSVFileCompressionNone CSVFileCompression = "NONE"
	CSVFileCompressionGzip CSVFileCompression = "GZIP"
)

// Values returns all known values for CSVFileCompression. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CSVFileCompression) Values() []CSVFileCompression {
	return []CSVFileCompression{
		"NONE",
		"GZIP",
	}
}

type Frequency string

// Enum values for Frequency
const (
	FrequencyP1d   Frequency = "P1D"
	FrequencyPt1h  Frequency = "PT1H"
	FrequencyPt10m Frequency = "PT10M"
	FrequencyPt5m  Frequency = "PT5M"
)

// Values returns all known values for Frequency. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Frequency) Values() []Frequency {
	return []Frequency{
		"P1D",
		"PT1H",
		"PT10M",
		"PT5M",
	}
}

type JsonFileCompression string

// Enum values for JsonFileCompression
const (
	JsonFileCompressionNone JsonFileCompression = "NONE"
	JsonFileCompressionGzip JsonFileCompression = "GZIP"
)

// Values returns all known values for JsonFileCompression. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JsonFileCompression) Values() []JsonFileCompression {
	return []JsonFileCompression{
		"NONE",
		"GZIP",
	}
}

type RelationshipType string

// Enum values for RelationshipType
const (
	RelationshipTypeCauseOfInputAnomalyGroup  RelationshipType = "CAUSE_OF_INPUT_ANOMALY_GROUP"
	RelationshipTypeEffectOfInputAnomalyGroup RelationshipType = "EFFECT_OF_INPUT_ANOMALY_GROUP"
)

// Values returns all known values for RelationshipType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelationshipType) Values() []RelationshipType {
	return []RelationshipType{
		"CAUSE_OF_INPUT_ANOMALY_GROUP",
		"EFFECT_OF_INPUT_ANOMALY_GROUP",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}
