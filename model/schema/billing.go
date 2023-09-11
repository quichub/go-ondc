package schema

/*
   Billing:
     description: Describes a billing event
     type: object
     properties:
       name:
         description: Personal details of the customer needed for billing.
         type: string
       organization:
         $ref: '#/components/schemas/Organization'
       address:
         $ref: '#/components/schemas/Address'
       email:
         type: string
         format: email
       phone:
         type: string
       time:
         $ref: '#/components/schemas/Time'
       tax_number:
         description: GST number
         type: string
       created_at:
         type: string
         format: date-time
       updated_at:
         type: string
         format: date-time
*/

type Billing struct {
	Name         string       `json:"name,omitempty"`         // Personal details of the customer needed for billing.
	Organization Organization `json:"organization,omitempty"` // Describes an organization
	Address      Address      `json:"address,omitempty"`      // Describes an address
	Email        string       `json:"email,omitempty"`        // Email address of the customer
	Phone        string       `json:"phone,omitempty"`        // Phone number of the customer
	Time         Time         `json:"time,omitempty"`         // Describes time
	TaxNumber    string       `json:"tax_number,omitempty"`   // GST number
	CreatedAt    DateTime     `json:"created_at,omitempty"`   // Timestamp in RFC3339 format
	UpdatedAt    DateTime     `json:"updated_at,omitempty"`   // Timestamp in RFC3339 format
}
