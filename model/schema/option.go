package schema

/*
Option:
     description: Describes a selectable option
     type: object
     properties:
       id:
         type: string
       descriptor:
         $ref: '#/components/schemas/Descriptor'
*/

type Option struct {
	Id         string     `json:"id,omitempty"`
	Descriptor Descriptor `json:"descriptor,omitempty"`
}

type OptionId string
