package schema

/*
   Country:
     description: Describes a country.
     type: object
     properties:
       name:
         type: string
         description: Name of the country
       code:
         type: string
         description: Country code as per ISO 3166 Alpha-3 code format
*/

type CountryCode string

type Country struct {
	Name string      `json:"name"` // Name of the country
	Code CountryCode `json:"code"` // Country code as per ISO 3166 Alpha-3 code format
}
