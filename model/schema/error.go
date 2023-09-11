package schema

/*
   Error:
     description: Describes an error object
     type: object
     properties:
       type:
         type: string
         enum:
           - CONTEXT-ERROR
           - CORE-ERROR
           - DOMAIN-ERROR
           - POLICY-ERROR
           - JSON-SCHEMA-ERROR
       code:
         type: string
         description: ONDC specific error code. For full list of error codes, refer to docs/drafts/Error Codes.md of this repo
       path:
         type: string
         description: Path to json schema generating the error. Used only during json schema validation errors
       message:
         type: string
         description: Human readable message describing the error
     required:
       - type
       - code
*/

type ErrorType string

const (
	ContextError    ErrorType = "CONTEXT-ERROR"
	CoreError       ErrorType = "CORE-ERROR"
	DomainError     ErrorType = "DOMAIN-ERROR"
	PolicyError     ErrorType = "POLICY-ERROR"
	JsonSchemaError ErrorType = "JSON-SCHEMA-ERROR"
)

type Error struct {
	Type    ErrorType `json:"type"`
	Code    string    `json:"code"`
	Path    string    `json:"path,omitempty"`
	Message string    `json:"message"`
}

func (e Error) Validate() error {
	if e.Type == "" {
		return NewValidationError("Error.Type is required")
	}

	if e.Code == "" {
		return NewValidationError("Error.Code is required")
	}

	return nil
}
