package schema

/*
   Person:
     description: Describes a person.
     type: object
     properties:
       name:
         $ref: '#/components/schemas/Name'
       image:
         $ref: '#/components/schemas/Image'
       dob:
         type: string
         format: date
       gender:
         type: string
         description: 'Gender of something, typically a Person, but possibly also fictional characters, animals, etc. While Male and Female may be used, text strings are also acceptable for people who do not identify as a binary gender'
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type Person struct {
	Name   Name     `json:"name,omitempty"`   // Describes a name
	Image  Image    `json:"image,omitempty"`  // Describes an image
	Dob    Date     `json:"dob,omitempty"`    // Date of birth of the person
	Gender string   `json:"gender,omitempty"` // Gender of something, typically a Person, but possibly also fictional characters, animals, etc. While Male and Female may be used, text strings are also acceptable for people who do not identify as a binary gender
	Tags   TagGroup `json:"tags,omitempty"`   // Describes a tag group
}
