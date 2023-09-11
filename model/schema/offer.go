package schema

/*
   Offer:
     description: Describes an offer
     type: object
     properties:
       id:
         type: string
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       location_ids:
         type: array
         items:
           $ref: '#/components/schemas/Location/properties/id'
       category_ids:
         type: array
         items:
           $ref: '#/components/schemas/Category/properties/id'
       item_ids:
         type: array
         items:
           $ref: '#/components/schemas/Item/properties/id'
       time:
         $ref: '#/components/schemas/Time'
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type Offer struct {
	Id          string     `json:"id,omitempty"`
	Descriptor  Descriptor `json:"descriptor,omitempty"`
	LocationIds []string   `json:"location_ids,omitempty"`
	CategoryIds []string   `json:"category_ids,omitempty"`
	ItemIds     []string   `json:"item_ids,omitempty"`
	Time        Time       `json:"time,omitempty"`
	Tags        []TagGroup `json:"tags,omitempty"`
}
