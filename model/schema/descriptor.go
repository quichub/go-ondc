package schema

/*
   Descriptor:
     description: Describes the description of a real-world object.
     type: object
     properties:
       name:
         type: string
       code:
         type: string
       symbol:
         type: string
       short_desc:
         type: string
       long_desc:
         type: string
       images:
         type: array
         items:
           $ref: '#/components/schemas/Image'
       audio:
         type: string
         format: uri
       3d_render:
         type: string
         format: uri
*/

type Descriptor struct {
	Name      string  `json:"name,omitempty"`
	Code      string  `json:"code,omitempty"`
	Symbol    string  `json:"symbol,omitempty"`
	ShortDesc string  `json:"short_desc,omitempty"`
	LongDesc  string  `json:"long_desc,omitempty"`
	Images    []Image `json:"images,omitempty"`
	Audio     Uri     `json:"audio,omitempty"`
	Render3d  Uri     `json:"3d_render,omitempty"`
}
