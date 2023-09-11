package schema

/*
   Cancellation:
     description: Describes a cancellation event
     type: object
     properties:
       type:
         type: string
         enum:
           - full
           - partial
       ref_id:
         type: string
       policies:
         type: array
         items:
           $ref: '#/components/schemas/Policy'
       time:
         type: string
         format: date-time
       cancelled_by:
         type: string
       reasons:
         $ref: '#/components/schemas/Option'
       selected_reason:
         type: object
         properties:
           id:
             $ref: '#/components/schemas/Option/properties/id'
       additional_description:
         $ref: '#/components/schemas/Descriptor'
*/

type CancellationType string

const (
	FullCancellation    CancellationType = "full"
	PartialCancellation CancellationType = "partial"
)

type SelectedReasonType struct {
	Id OptionId `json:"id,omitempty"`
}

type Cancellation struct {
	Type                  CancellationType   `json:"type,omitempty"`
	RefId                 string             `json:"ref_id,omitempty"`
	Policies              []Policy           `json:"policies,omitempty"`
	Time                  DateTime           `json:"time,omitempty"`
	CancelledBy           string             `json:"cancelled_by,omitempty"`
	Reasons               Option             `json:"reasons,omitempty"`
	SelectedReason        SelectedReasonType `json:"selected_reason,omitempty"`
	AdditionalDescription Descriptor         `json:"additional_description,omitempty"`
}
