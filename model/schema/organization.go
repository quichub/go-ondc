package schema

/*
Organization:
     description: Describes an organization
     type: object
     properties:
       name:
         type: string
       cred:
         type: string
*/

type Organization struct {
	Name string `json:"name"`
	Cred string `json:"cred"`
}
