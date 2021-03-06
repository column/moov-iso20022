/*
 * ISO20022 API
 *
 * ISO 20022 is an ISO standard for electronic data interchange between financial institutions. It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The metadata is stored in UML models with a special ISO 20022 UML Profile. The metadata is transformed into the syntax of messages used in financial networks. The first syntax supported for messages was XML Schema. Package ISO20022 implements a message reader and writer written in Go decorated with a HTTP API for creating, parsing, and validating meta data messages. Package ISO20022 supported xml and json format for message  | Input      | Output     |  |------------|------------|  | JSON       | JSON       |  | XML        | XML        |
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"net/http"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

// NewAPIResponse returns a new APIResonse object.
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError returns a new APIResponse object with the provided error message.
func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
