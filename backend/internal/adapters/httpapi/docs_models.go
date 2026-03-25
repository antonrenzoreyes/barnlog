package httpapi

import (
	"fmt"
	"time"

	openapicontract "barnlog/backend/internal/contracts/openapi"
)

func newStatusResponse(status string) openapicontract.HttpapiStatusResponse {
	return openapicontract.HttpapiStatusResponse{Status: &status}
}

func newReadyResponse(status string, timestamp time.Time) openapicontract.HttpapiReadyResponse {
	return openapicontract.HttpapiReadyResponse{
		Status:    &status,
		Timestamp: &timestamp,
	}
}

func newErrorResponse(code string) openapicontract.HttpapiErrorResponse {
	return openapicontract.HttpapiErrorResponse{Error: &code}
}

func newUploadFileResponse(
	fileID, fileName, contentType string,
	sizeBytes int64,
) (openapicontract.HttpapiUploadFileResponse, error) {
	maxInt := int64(^uint(0) >> 1)
	if sizeBytes < 0 || sizeBytes > maxInt {
		return openapicontract.HttpapiUploadFileResponse{}, fmt.Errorf("size_bytes out of range: %d", sizeBytes)
	}

	return openapicontract.HttpapiUploadFileResponse{
		FileId:      fileID,
		FileName:    fileName,
		ContentType: contentType,
		SizeBytes:   int(sizeBytes),
	}, nil
}
