package schema

/*
   Agent:
     description: Describes an order executor
     allOf:
       - $ref: '#/components/schemas/Person'
       - $ref: '#/components/schemas/Contact'
       - type: object
         properties:
           rateable:
             $ref: '#/components/schemas/Rateable'
*/

type Agent struct {
	Person,
	Contact,
	Rateable Rateable `json:"rateable,omitempty"`
}
