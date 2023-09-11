package schema

/*
   Authorization:
     description: Describes an authorization mechanism
     type: object
     properties:
       type:
         type: string
         description: Type of authorization mechanism used
       token:
         type: string
         description: Token used for authorization
       valid_from:
         type: string
         format: date-time
         description: Timestamp in RFC3339 format from which token is valid
       valid_to:
         type: string
         format: date-time
         description: Timestamp in RFC3339 format until which token is valid
       status:
         type: string
         description: Status of the token
*/

type Authorization struct {
	Type      string    `json:"type,omitempty"`       // Type of authorization mechanism used
	Token     string    `json:"token,omitempty"`      // Token used for authorization
	ValidFrom Timestamp `json:"valid_from,omitempty"` // Timestamp in RFC3339 format from which token is valid
	ValidTo   Timestamp `json:"valid_to,omitempty"`   // Timestamp in RFC3339 format until which token is valid
	Status    string    `json:"status,omitempty"`     // Status of the token
}
