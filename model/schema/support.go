/*
   Support:
     description: Customer support
     type: object
     properties:
       type:
         type: string
         enum:
           - order
           - billing
           - fulfillment
       ref_id:
         type: string
       channels:
         $ref: '#/components/schemas/TagGroup'
*/

package schema

type SupportType string

const (
	OrderSupport       SupportType = "order"
	BillingSupport     SupportType = "billing"
	FulfillmentSupport SupportType = "fulfillment"
)

type Support struct {
	Type     SupportType `json:"type,omitempty"`
	RefId    string      `json:"ref_id,omitempty"`
	Channels TagGroup    `json:"channels,omitempty"`
}
