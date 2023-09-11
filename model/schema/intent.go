package schema

/*
   Intent:
     description: Intent of a user. Used for searching for services. Buyer App can set finder fee type in payment."@ondc/org/buyer_app_finder_fee_type" and amount in "@ondc/org/buyer_app_finder_fee_amount"
     type: object
     properties:
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       provider:
         $ref: '#/components/schemas/Provider'
       fulfillment:
         $ref: '#/components/schemas/Fulfillment'
       payment:
         $ref: '#/components/schemas/Payment'
       category:
         $ref: '#/components/schemas/Category'
       offer:
         $ref: '#/components/schemas/Offer'
       item:
         $ref: '#/components/schemas/Item'
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type Intent struct {
	Descriptor  Descriptor  `json:"descriptor,omitempty"`
	Provider    Provider    `json:"provider,omitempty"`
	Fulfillment Fulfillment `json:"fulfillment,omitempty"`
	Payment     Payment     `json:"payment,omitempty"`
	Category    Category    `json:"category,omitempty"`
	Offer       Offer       `json:"offer,omitempty"`
	Item        Item        `json:"item,omitempty"`
	Tags        []TagGroup  `json:"tags,omitempty"`
}
