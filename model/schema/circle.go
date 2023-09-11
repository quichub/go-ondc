package schema

/*
Circle:
     description: Describes a circular area on the map
     type: object
     properties:
       gps:
         $ref: '#/components/schemas/Gps'
       radius:
         $ref: '#/components/schemas/Scalar'
     required:
       - gps
       - radius
*/

type Circle struct {
	Gps    Gps    `json:"gps"`
	Radius Scalar `json:"radius"`
}
