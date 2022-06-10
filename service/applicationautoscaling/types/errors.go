// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Concurrent updates caused an exception, for example, if you request an update to
// an Application Auto Scaling resource that already has a pending update.
type ConcurrentUpdateException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentUpdateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentUpdateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentUpdateException) ErrorCode() string             { return "ConcurrentUpdateException" }
func (e *ConcurrentUpdateException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Failed access to resources caused an exception. This exception is thrown when
// Application Auto Scaling is unable to retrieve the alarms associated with a
// scaling policy due to a client error, for example, if the role ARN specified for
// a scalable target does not have permission to call the CloudWatch DescribeAlarms
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarms.html)
// on your behalf.
type FailedResourceAccessException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *FailedResourceAccessException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FailedResourceAccessException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FailedResourceAccessException) ErrorCode() string             { return "FailedResourceAccessException" }
func (e *FailedResourceAccessException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service encountered an internal error.
type InternalServiceException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The next token supplied was invalid.
type InvalidNextTokenException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A per-account resource limit is exceeded. For more information, see Application
// Auto Scaling service quotas
// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-limits.html).
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified object could not be found. For any operation that depends on the
// existence of a scalable target, this exception is thrown if the scalable target
// with the specified service namespace, resource ID, and scalable dimension does
// not exist. For any operation that deletes or deregisters a resource, this
// exception is thrown if the resource cannot be found.
type ObjectNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ObjectNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectNotFoundException) ErrorCode() string             { return "ObjectNotFoundException" }
func (e *ObjectNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception was thrown for a validation issue. Review the available parameters
// for the API request.
type ValidationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }