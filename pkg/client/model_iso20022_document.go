/*
 * ISO20022 API
 *
 * ISO 20022 is an ISO standard for electronic data interchange between financial institutions. It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The metadata is stored in UML models with a special ISO 20022 UML Profile. The metadata is transformed into the syntax of messages used in financial networks. The first syntax supported for messages was XML Schema. Package ISO20022 implements a message reader and writer written in Go decorated with a HTTP API for creating, parsing, and validating meta data messages. Package ISO20022 supported xml and json format for message  | Input      | Output     |  |------------|------------|  | JSON       | JSON       |  | XML        | XML        |
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Iso20022Document https://www.iso20022.org/iso-20022-message-definitions?business-domain=1
type Iso20022Document struct {
	Xmlns         string                 `json:"Xmlns"`
	RequestObject map[string]interface{} `json:"RequestObject"`
}
