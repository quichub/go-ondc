package schema

/*
   Ack:
     description: Describes the ACK response
     type: object
     properties:
       status:
         type: string
         description: 'Describe the status of the ACK response. If schema validation passes, status is ACK else it is NACK'
         enum:
           - ACK
           - NACK
       tags:
         description: A list of tags containing any additional information sent along with the Acknowledgement.
         type: array
         items:
           $ref: '#/components/schemas/TagGroup'
     required:
       - status
*/

type AckStatus string // Describe the status of the ACK response. If schema validation passes, status is ACK else it is NACK

const (
	Ack  AckStatus = "ACK"
	Nack AckStatus = "NACK"
)

type AckType struct {
	Status AckStatus  `json:"status"` // Describe the status of the ACK response. If schema validation passes, status is ACK else it is NACK
	Tags   []TagGroup `json:"tags,omitempty"`
}

//validate the Ack struct

func (a *AckType) Validate() error {

	if a.Status == "" {
		return NewValidationError("status is required")
	}

	return nil
}
