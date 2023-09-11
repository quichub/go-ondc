package schema

/*
Contact:
     type: object
     properties:
       phone:
         type: string
       email:
         type: string
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type Contact struct {
	Phone string     `json:"phone,omitempty"`
	Email string     `json:"email,omitempty"`
	Tags  []TagGroup `json:"tags,omitempty"`
}
