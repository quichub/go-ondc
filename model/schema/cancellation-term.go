package schema

/*
   CancellationTerm:
     description: Describes the cancellation terms of an item or an order. This can be referenced at an item or order level. Item-level cancellation terms can override the terms at the order level.
     type: object
     properties:
       reason_required:
         description: Indicates whether a reason is required to cancel the order
         type: boolean
       refund_eligible:
         description: Indicates if cancellation will result in a refund
         type: boolean
       return_eligible:
         description: Indicates if cancellation will result in a return to origin
         type: boolean
       fulfillment_state:
         description: The state of fulfillment during which these terms are applicable.
         allOf:
           - $ref: '#/components/schemas/State'
       return_policy:
         description: Describes the return policy of an item or an order
         type: object
         properties:
           return_eligible:
             description: Indicates whether the item is eligible for return
             type: boolean
           return_within:
             description: 'Applicable only for buyer managed returns where the buyer has to return the item to the origin before a certain date-time, failing which they will not be eligible for refund.'
             allOf:
               - $ref: '#/components/schemas/Time'
           return_location:
             $ref: '#/components/schemas/Location'
           fulfillment_managed_by:
             enum:
               - customer
               - provider
       refund_policy:
         type: object
         properties:
           refund_eligible:
             description: Indicates if cancellation will result in a refund
             type: boolean
           refund_within:
             description: Time within which refund will be processed after successful cancellation.
             allOf:
               - $ref: '#/components/schemas/Time'
           refund_amount:
             $ref: '#/components/schemas/Price'
       cancel_by:
         description: Information related to the time of cancellation.
         allOf:
           - $ref: '#/components/schemas/Time'
       cancellation_fee:
         $ref: '#/components/schemas/Fee'
       xinput_required:
         $ref: '#/components/schemas/XInput'
       xinput_response:
         $ref: '#/components/schemas/XInputResponse'
       external_ref:
         $ref: '#/components/schemas/MediaFile'*/

type FulfillmentManagedByType string

const (
	CustomerFulfilled FulfillmentManagedByType = "customer"
	ProviderFulfilled FulfillmentManagedByType = "provider"
)

type ReturnPolicy struct {
	ReturnEligible       bool                     `json:"return_eligible,omitempty"` // Indicates whether the item is eligible for return
	ReturnWithin         Time                     `json:"return_within,omitempty"`   // Applicable only for buyer managed returns where the buyer has to return the item to the origin before a certain date-time, failing which they will not be eligible for refund.
	ReturnLocation       Location                 `json:"return_location,omitempty"`
	FulfillmentManagedBy FulfillmentManagedByType `json:"fulfillment_managed_by,omitempty"`
}

type RefundPolicy struct {
	RefundEligible bool  `json:"refund_eligible,omitempty"` // Indicates if cancellation will result in a refund
	RefundWithin   Time  `json:"refund_within,omitempty"`   // Time within which refund will be processed after successful cancellation.
	RefundAmount   Price `json:"refund_amount,omitempty"`
}

type CancellationTerm struct {
	ReasonRequired   bool           `json:"reason_required,omitempty"` // Indicates whether a reason is required to cancel the order
	RefundEligible   bool           `json:"refund_eligible,omitempty"` // Indicates if cancellation will result in a refund
	ReturnEligible   bool           `json:"return_eligible,omitempty"` // Indicates if cancellation will result in a return to origin
	FulfillmentState State          `json:"fulfillment_state,omitempty"`
	ReturnPolicy     ReturnPolicy   `json:"return_policy,omitempty"` // Describes the return policy of an item or an order
	RefundPolicy     RefundPolicy   `json:"refund_policy,omitempty"`
	CancelBy         Time           `json:"cancel_by,omitempty"` // Information related to the time of cancellation.
	CancellationFee  Fee            `json:"cancellation_fee,omitempty"`
	XinputRequired   XInput         `json:"xinput_required,omitempty"`
	XinputResponse   XInputResponse `json:"xinput_response,omitempty"`
	ExternalRef      MediaFile      `json:"external_ref,omitempty"`
}
