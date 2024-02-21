module github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson

go 1.20

require (
	github.com/aws/aws-sdk-go-v2 v1.25.1
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.1
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.1
	github.com/aws/smithy-go v1.20.1
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/
