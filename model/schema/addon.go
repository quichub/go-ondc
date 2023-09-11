package schema

/*
   AddOn:
     description: Describes an add-on
     type: object
     properties:
       id:
         type: string
         description: 'ID of the add-on. This follows the syntax {item.id}/add-on/{add-on unique id} for item specific add-on OR '
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       price:
         $ref: '#/components/schemas/Price'
*/

type AddOnId string

type AddOn struct {
	Id         AddOnId    `json:"id,omitempty"`         // ID of the add-on. This follows the syntax {item.id}/add-on/{add-on unique id} for item specific add-on OR
	Descriptor Descriptor `json:"descriptor,omitempty"` // Describes the description of a real-world object.
	Price      Price      `json:"price,omitempty"`      // Describes the price of a real-world object
}
