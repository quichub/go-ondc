package schema

/*
   Dimensions:
     description: Describes the dimensions of a real-world object
     type: object
     properties:
       length:
         $ref: '#/components/schemas/Scalar'
       breadth:
         $ref: '#/components/schemas/Scalar'
       height:
         $ref: '#/components/schemas/Scalar'
*/

type Dimensions struct {
	Length  Scalar `json:"length,omitempty"`
	Breadth Scalar `json:"breadth,omitempty"`
	Height  Scalar `json:"height,omitempty"`
}
