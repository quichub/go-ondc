/*
   State:
     description: Describes a state
     type: object
     properties:
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       updated_at:
         type: string
         format: date-time
       updated_by:
         type: string
         description: ID of entity which changed the state
*/

package schema

type State struct {
	Descriptor Descriptor `json:"descriptor,omitempty"` // Describes a descriptor
	UpdatedAt  DateTime   `json:"updated_at,omitempty"` // Timestamp in RFC3339 format
	UpdatedBy  string     `json:"updated_by,omitempty"` // ID of entity which changed the state
}
