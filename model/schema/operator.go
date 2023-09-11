package schema

/*
Operator:
     description: Describes the agent of a service
     allOf:
       - $ref: "#/components/schemas/Person"
       - properties:
           experience:
             type: object
             properties:
               label:
                 type: string
               value:
                 type: string
               unit:
                 type: string
*/

type Operator struct {
	Person
	Experience struct {
		Label string `json:"label,omitempty"`
		Value string `json:"value,omitempty"`
		Unit  string `json:"unit,omitempty"`
	} `json:"experience,omitempty"`
}
