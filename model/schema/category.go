package schema

/*
Category:
        description: Describes a category
     type: object
     properties:
       id:
         type: string
         description: Unique id of the category
       parent_category_id:
         type: string
         description: Unique id of the parent category
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       time:
         $ref: '#/components/schemas/Time'
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type CategoryId string

type Category struct {
	Id               CategoryId `json:"id,omitempty"`                 // Unique id of the category
	ParentCategoryId string     `json:"parent_category_id,omitempty"` // Unique id of the parent category
	Descriptor       Descriptor `json:"descriptor,omitempty"`         // Describes the description of a real-world object.
	Time             Time       `json:"time,omitempty"`
	Tags             []TagGroup `json:"tags,omitempty"`
}
