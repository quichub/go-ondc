/*
   Schedule:
     description: Describes a schedule
     type: object
     properties:
       frequency:
         $ref: '#/components/schemas/Duration'
       holidays:
         type: array
         items:
           type: string
           format: date-time
       times:
         type: array
         items:
           type: string
           format: date-time

*/

package schema

type Schedule struct {
	Frequency Duration   `json:"frequency,omitempty"` // Describes a duration
	Holidays  []DateTime `json:"holidays,omitempty"`  // Timestamp in RFC3339 format
	Times     []DateTime `json:"times,omitempty"`     // Timestamp in RFC3339 format
}
