package schema

/*
   Catalog:
     description: Describes a Seller App catalog
     type: object
     properties:
       bpp/descriptor:
         $ref: '#/components/schemas/Descriptor'
       bpp/categories:
         type: array
         items:
           $ref: '#/components/schemas/Category'
       bpp/fulfillments:
         type: array
         items:
           $ref: '#/components/schemas/Fulfillment'
       bpp/payments:
         type: array
         items:
           $ref: '#/components/schemas/Payment'
       bpp/offers:
         type: array
         items:
           $ref: '#/components/schemas/Offer'
       bpp/providers:
         type: array
         items:
           $ref: '#/components/schemas/Provider'
       exp:
         type: string
         description: Time after which catalog has to be refreshed
         format: date-time
*/

type Catalog struct {
	Descriptor   Descriptor    `json:"bpp/descriptor,omitempty"`
	Categories   []Category    `json:"bpp/categories,omitempty"`
	Fulfillments []Fulfillment `json:"bpp/fulfillments,omitempty"`
	Payments     []Payment     `json:"bpp/payments,omitempty"`
	Offers       []Offer       `json:"bpp/offers,omitempty"`
	Providers    []Provider    `json:"bpp/providers,omitempty"`
	Exp          DateTime      `json:"exp,omitempty"` // Time after which catalog has to be refreshed
}
