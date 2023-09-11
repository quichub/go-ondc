package schema

import "fmt"

type FulfillmentId string

/*
   Fulfillment:
     description: Describes how a single product/service will be rendered/fulfilled to the end customer
     type: object
     properties:
       id:
         type: string
         description: Unique reference ID to the fulfillment of an order
       type:
         type: string
         description: This describes the type of fulfillment ("Pickup" - Buyer picks up from store by themselves or through their logistics provider; "Delivery" - seller delivers to buyer)
         enum:
           - Delivery
           - Pickup
           - Delivery and Pickup
           - Reverse QC
       "@ondc/org/category":
         type: string
         description: Fulfillment Category
       "@ondc/org/TAT":
         description: fulfillment turnaround time in ISO8601 durations format e.g. 'PT24H' indicates 24 hour TAT
         type: string
       provider_id:
         $ref: '#/components/schemas/Provider/properties/id'
       "@ondc/org/provider_name":
         type: string
       rating:
         $ref: '#/components/schemas/Rating/properties/value'
       state:
         $ref: '#/components/schemas/State'
       tracking:
         type: boolean
         description: Indicates whether the fulfillment allows tracking
         default: false
       customer:
         type: object
         properties:
           person:
             $ref: '#/components/schemas/Person'
           contact:
             $ref: '#/components/schemas/Contact'
       agent:
         $ref: '#/components/schemas/Agent'
       person:
         $ref: '#/components/schemas/Person'
       contact:
         $ref: '#/components/schemas/Contact'
       vehicle:
         $ref: '#/components/schemas/Vehicle'
       start:
         description: Details on the start of fulfillment
         type: object
         properties:
           location:
             $ref: '#/components/schemas/Location'
           time:
             $ref: '#/components/schemas/Time'
           instructions:
             $ref: '#/components/schemas/Descriptor'
           contact:
             $ref: '#/components/schemas/Contact'
           person:
             $ref: '#/components/schemas/Person'
           authorization:
             $ref: '#/components/schemas/Authorization'
       end:
         description: Details on the end of fulfillment
         type: object
         properties:
           location:
             $ref: '#/components/schemas/Location'
           time:
             $ref: '#/components/schemas/Time'
           instructions:
             $ref: '#/components/schemas/Descriptor'
           contact:
             $ref: '#/components/schemas/Contact'
           person:
             $ref: '#/components/schemas/Person'
           authorization:
             $ref: '#/components/schemas/Authorization'
       rateable:
         $ref: '#/components/schemas/Rateable'
       tags:
         $ref: '#/components/schemas/TagGroup'
     required:
       - id
       - type
*/

type FulfillmentType string

const (
	Delivery          FulfillmentType = "Delivery"
	Pickup            FulfillmentType = "Pickup"
	DeliveryAndPickup FulfillmentType = "Delivery and Pickup"
	ReverseQC         FulfillmentType = "Reverse QC"
)

type Customer struct {
	Person  Person  `json:"person,omitempty"`
	Contact Contact `json:"contact,omitempty"`
}

type FulfillmentStartEnd struct {
	Location      Location      `json:"location,omitempty"`
	Time          Time          `json:"time,omitempty"`
	Instructions  Descriptor    `json:"instructions,omitempty"`
	Contact       Contact       `json:"contact,omitempty"`
	Person        Person        `json:"person,omitempty"`
	Authorization Authorization `json:"authorization,omitempty"`
}

type Fulfillment struct {
	Id                  string               `json:"id,omitempty"`                 // Unique reference ID to the fulfillment of an order
	Type                FulfillmentType      `json:"type,omitempty"`               // This describes the type of fulfillment ("Pickup" - Buyer picks up from store by themselves or through their logistics provider; "Delivery" - seller delivers to buyer)
	OndcOrgCategory     *string              `json:"@ondc/org/category,omitempty"` // Fulfillment Category
	OndcOrgTAT          *string              `json:"@ondc/org/TAT,omitempty"`      // fulfillment turnaround time in ISO8601 durations format e.g. 'PT24H' indicates 24 hour TAT
	ProviderId          *ProviderId          `json:"provider_id,omitempty"`
	OndcOrgProviderName *string              `json:"@ondc/org/provider_name,omitempty"`
	Rating              *RatingValue         `json:"rating,omitempty"`
	State               *State               `json:"state,omitempty"`
	Tracking            *bool                `json:"tracking,omitempty"` // Indicates whether the fulfillment allows tracking
	Customer            *Customer            `json:"customer,omitempty"`
	Agent               *Agent               `json:"agent,omitempty"`
	Person              *Person              `json:"person,omitempty"`
	Contact             *Contact             `json:"contact,omitempty"`
	Vehicle             *Vehicle             `json:"vehicle,omitempty"`
	Start               *FulfillmentStartEnd `json:"start,omitempty"` // Details on the start of fulfillment
	End                 *FulfillmentStartEnd `json:"end,omitempty"`   // Details on the end of fulfillment
	Rateable            *Rateable            `json:"rateable,omitempty"`
	Tags                *TagGroup            `json:"tags,omitempty"`
}

// validate validates the Fulfillment object
func (f *Fulfillment) Validate() error {
	if f.Type == "" {
		return fmt.Errorf("Fulfillment.Type is required")
	}
	if f.Id == "" {
		return fmt.Errorf("Fulfillment.Id is required")
	}
	return nil
}
