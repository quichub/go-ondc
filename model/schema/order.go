/*
Order:
     description: Describes the details of an order
     type: object
     properties:
       id:
         type: string
         description: Unique identifier for Order across the network
       state:
         type: string
       provider:
         type: object
         properties:
           id:
             $ref: '#/components/schemas/Provider/properties/id'
           locations:
             type: array
             maxItems: 1
             items:
               type: object
               properties:
                 id:
                   $ref: '#/components/schemas/Location/properties/id'
       items:
         type: array
         items:
           $ref: '#/components/schemas/Item'
       add_ons:
         type: array
         items:
           type: object
           properties:
             id:
               $ref: '#/components/schemas/AddOn/properties/id'
       offers:
         type: array
         items:
           type: object
           properties:
             id:
               $ref: '#/components/schemas/Offer/properties/id'
       documents:
         type: array
         items:
           $ref: '#/components/schemas/Document'
       billing:
         $ref: '#/components/schemas/Billing'
       fulfillments:
         type: array
         items:
           $ref: "#/components/schemas/Fulfillment"
       cancellation_terms:
         description: The cancellation terms of this order. This can be overriden at the item level cancellation terms.
         type: array
         items:
           $ref: '#/components/schemas/CancellationTerm'
       quote:
         $ref: '#/components/schemas/Quotation'
       payment:
         $ref: '#/components/schemas/Payment'
       created_at:
         type: string
         format: date-time
       updated_at:
         type: string
         format: date-time
*/

package schema

type OrderId string

type OrderProviderLocation struct {
	Id LocationId `json:"id,omitempty"`
}

type OrderProvider struct {
	Id        ProviderId `json:"id,omitempty"`        // Unique identifier for Provider across the network
	Locations []Location `json:"locations,omitempty"` // Describes a location
}

type Order struct {
	Id                OrderId            `json:"id,omitempty"` // Unique identifier for Order across the network
	State             string             `json:"state,omitempty"`
	Provider          Provider           `json:"provider,omitempty"`
	Items             []Item             `json:"items,omitempty"`
	AddOns            []AddOn            `json:"add_ons,omitempty"`
	Offers            []Offer            `json:"offers,omitempty"`
	Documents         []Document         `json:"documents,omitempty"`
	Billing           Billing            `json:"billing,omitempty"`
	Fulfillments      []Fulfillment      `json:"fulfillments,omitempty"`
	CancellationTerms []CancellationTerm `json:"cancellation_terms,omitempty"`
	Quote             Quotation          `json:"quote,omitempty"`
	Payment           Payment            `json:"payment,omitempty"`
	CreatedAt         DateTime           `json:"created_at,omitempty"` // Timestamp in RFC3339 format
	UpdatedAt         DateTime           `json:"updated_at,omitempty"` // Timestamp in RFC3339 format
}
