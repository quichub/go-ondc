package schema

/*
   Price:
     description: Describes the price of an item. Allows for domain extension.
     type: object
     properties:
       currency:
         type: string
         description: ISO 4217 alphabetic currency code e.g. 'INR'
       value:
         $ref: '#/components/schemas/DecimalValue'
       estimated_value:
         $ref: '#/components/schemas/DecimalValue'
       computed_value:
         $ref: '#/components/schemas/DecimalValue'
       listed_value:
         $ref: '#/components/schemas/DecimalValue'
       offered_value:
         $ref: '#/components/schemas/DecimalValue'
       minimum_value:
         $ref: '#/components/schemas/DecimalValue'
       maximum_value:
         $ref: '#/components/schemas/DecimalValue'
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type PriceCurrency string

type PriceValue DecimalValue

type Price struct {
	Currency       PriceCurrency `json:"currency,omitempty"`        // ISO 4217 alphabetic currency code e.g. 'INR'
	Value          PriceValue    `json:"value,omitempty"`           // Value of the price
	EstimatedValue DecimalValue  `json:"estimated_value,omitempty"` // Estimated value of the price
	ComputedValue  DecimalValue  `json:"computed_value,omitempty"`  // Computed value of the price
	ListedValue    DecimalValue  `json:"listed_value,omitempty"`    // Listed value of the price
	OfferedValue   DecimalValue  `json:"offered_value,omitempty"`   // Offered value of the price
	MinimumValue   DecimalValue  `json:"minimum_value,omitempty"`   // Minimum value of the price
	MaximumValue   DecimalValue  `json:"maximum_value,omitempty"`   // Maximum value of the price
	Tags           []TagGroup    `json:"tags,omitempty"`
}
