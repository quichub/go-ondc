/*
   Provider:
     description: Describes a service provider. This can be a restaurant, a hospital, a Store etc
     type: object
     properties:
       id:
         type: string
         description: Id of the provider
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       category_id:
         type: string
         description: Category Id of the provider
       "@ondc/org/fssai_license_no":
         type: string
         description: FSSAI license no. Mandatory for category_id "F&B"
       rating:
         $ref: '#/components/schemas/Rating/properties/value'
       time:
         $ref: '#/components/schemas/Time'
       categories:
         type: array
         items:
           $ref: '#/components/schemas/Category'
       creds:
         type: array
         items:
           $ref: '#/components/schemas/Credential'
       fulfillments:
         type: array
         items:
           $ref: '#/components/schemas/Fulfillment'
       payments:
         type: array
         items:
           $ref: '#/components/schemas/Payment'
       locations:
         type: array
         items:
           allOf:
           - $ref: '#/components/schemas/Location'
           - type: object
             properties:
               rateable:
                 $ref: '#/components/schemas/Rateable'
       offers:
         type: array
         items:
           $ref: '#/components/schemas/Offer'
       items:
         type: array
         items:
           $ref: '#/components/schemas/Item'
       ttl:
         type: string
         description: Validity of catalog in ISO8601 durations format after which it has to be refreshed e.g. 'P7D' indicates validity of 7 days; value of 0 indicates catalog is not cacheable
       exp:
         type: string
         description: Time after which catalog has to be refreshed
         format: date-time
       rateable:
         $ref: '#/components/schemas/Rateable'
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

package schema

type ProviderId string

type RateableLocation struct {
	Location
	Rateable
}

type Provider struct {
	ID             ProviderId         `json:"id,omitempty"` // Id of the provider
	Descriptor     Descriptor         `json:"descriptor,omitempty"`
	CategoryId     string             `json:"category_id,omitempty"`                // Category Id of the provider
	FssaiLicenseNo string             `json:"@ondc/org/fssai_license_no,omitempty"` // FSSAI license no. Mandatory for category_id "F&B"
	Rating         *RatingValue       `json:"rating,omitempty"`
	Time           Time               `json:"time,omitempty"`
	Categories     []Category         `json:"categories,omitempty"`
	Creds          []Credential       `json:"creds,omitempty"`
	Fulfillments   []Fulfillment      `json:"fulfillments,omitempty"`
	Payments       []Payment          `json:"payments,omitempty"`
	Locations      []RateableLocation `json:"locations,omitempty"`
	Offers         []Offer            `json:"offers,omitempty"`
	Items          []Item             `json:"items,omitempty"`
	Ttl            string             `json:"ttl,omitempty"` // Validity of catalog in ISO8601 durations format after which it has to be refreshed e.g. 'P7D' indicates validity of 7 days; value of 0 indicates catalog is not cacheable
	Exp            DateTime           `json:"exp,omitempty"` // Time after which catalog has to be refreshed
	Rateable       *Rateable          `json:"rateable,omitempty"`
	Tags           *TagGroup          `json:"tags,omitempty"`
}
