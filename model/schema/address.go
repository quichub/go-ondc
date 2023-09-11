package schema

/*
   Address:
     description: Describes an address
     type: object
     properties:
       door:
         type: string
         description: Door / Shop number of the address
       name:
         type: string
         description: 'Name of address if applicable. Example, shop name'
       building:
         type: string
         description: Name of the building or block
       street:
         type: string
         description: Street name or number
       locality:
         type: string
         description: 'Name of the locality, apartments'
       ward:
         type: string
         description: Name or number of the ward if applicable
       city:
         type: string
         description: City name
       state:
         type: string
         description: State name
       country:
         type: string
         description: Country name
       area_code:
         type: string
         description: 'Area code. This can be Pincode, ZIP code or any equivalent'
*/

type Address struct {
	Door     string `json:"door,omitempty"`      // Door / Shop number of the address
	Name     string `json:"name,omitempty"`      // Name of address if applicable. Example, shop name
	Building string `json:"building,omitempty"`  // Name of the building or block
	Street   string `json:"street,omitempty"`    // Street name or number
	Locality string `json:"locality,omitempty"`  // Name of the locality, apartments
	Ward     string `json:"ward,omitempty"`      // Name or number of the ward if applicable
	City     string `json:"city,omitempty"`      // City name
	State    string `json:"state,omitempty"`     // State name
	Country  string `json:"country,omitempty"`   // Country name
	AreaCode string `json:"area_code,omitempty"` // Area code. This can be Pincode, ZIP code or any equivalent
}
