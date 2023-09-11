package schema

/*
	Language:
     description: indicates language code. ONDC supports language codes as per ISO 639.2 standard
     type: object
     properties:
       code:
         type: string
*/

type Language struct {
	Code string `json:"code"`
}
