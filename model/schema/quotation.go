package schema

/*
   Quotation:
     description: Describes a quote
     type: object
     properties:
       price:
         $ref: '#/components/schemas/Price'
       breakup:
         type: array
         items:
           type: object
           properties:
             "@ondc/org/item_id":
               $ref: '#/components/schemas/Item/properties/id'
             "@ondc/org/item_quantity":
               $ref: '#/components/schemas/ItemQuantity/properties/selected'
             "@ondc/org/title_type":
               type: string
               enum:
                 - item
                 - delivery
                 - packing
                 - tax
                 - misc
                 - discount
             item:
               $ref: '#/components/schemas/Item'
             title:
               type: string
             price:
               $ref: '#/components/schemas/Price'
       ttl:
         $ref: '#/components/schemas/Duration'
*/

type TitleType string

const (
	ItemTitle     TitleType = "item"
	DeliveryTitle TitleType = "delivery"
	PackingTitle  TitleType = "packing"
	TaxTitle      TitleType = "tax"
	MiscTitle     TitleType = "misc"
	DiscountTitle TitleType = "discount"
)

type BreakupItem struct {
	ItemId    string               `json:"@ondc/org/item_id,omitempty"`
	Quantity  ItemQuantitySelected `json:"@ondc/org/item_quantity,omitempty"`
	TitleType TitleType            `json:"@ondc/org/title_type,omitempty"`
	Item      Item                 `json:"item,omitempty"`
	Title     string               `json:"title,omitempty"`
	Price     Price                `json:"price,omitempty"`
}

type Quotation struct {
	Price   Price         `json:"price,omitempty"`
	Breakup []BreakupItem `json:"breakup,omitempty"`
	Ttl     Duration      `json:"ttl,omitempty"`
}
