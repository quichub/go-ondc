package schema

/*

 Document:
     description: Describes a document which can be sent as a url
     type: object
     properties:
       url:
         type: string
         format: uri
       label:
         type: string

*/

type Document struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}
