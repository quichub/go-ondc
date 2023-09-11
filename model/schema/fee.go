package schema

/*
  Fee:
     description: A fee applied on a particular entity
     type: object
     properties:
       percentage:
         description: Percentage of a value
         allOf:
           - $ref: '#/components/schemas/DecimalValue'
       amount:
         description: A fixed value
         allOf:
           - $ref: '#/components/schemas/Price'
*/

type Fee struct {
	Percentage DecimalValue `json:"percentage,omitempty"` // Percentage of a value
	Amount     Price        `json:"amount,omitempty"`     // A fixed value
}
