package schema

/*
   Credential:
     description: Describes a credential of an entity - Person or Organization
     type: object
     properties:
       id:
         type: string
       type:
         type: string
         default: VerifiableCredential
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       url:
         description: URL of the credential
         type: string
         format: uri
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type Credential struct {
	Id         string     `json:"id,omitempty"`
	Type       string     `json:"type,omitempty" default:"VerifiableCredential"`
	Descriptor Descriptor `json:"descriptor,omitempty"`
	Url        string     `json:"url,omitempty"`
	Tags       []TagGroup `json:"tags,omitempty"`
}
