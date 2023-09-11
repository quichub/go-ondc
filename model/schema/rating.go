package schema

/*
	Rating:
     description: Describes the rating of a person or an object.
     type: object
     properties:
       rating_category:
         description: Category of the object being rated
         type: string
       id:
         description: Id of the object being rated
         type: string
       value:
         description: Rating value given to the object (1 - Poor; 2 - Needs improvement; 3 - Satisfactory; 4 - Good; 5 - Excellent)
         type: number
         minimum: 1
         maximum: 5
       feedback_form:
         $ref: '#/components/schemas/FeedbackForm'
       feedback_id:
         $ref: '#/components/schemas/FeedbackUrl/properties/params/properties/feedback_id'
*/

type RatingValue int

type RatingId string

type Rating struct {
	RatingCategory string         `json:"rating_category,omitempty"` // Category of the object being rated
	Id             RatingId       `json:"id,omitempty"`              // Id of the object being rated
	Value          RatingValue    `json:"value,omitempty"`           // Rating value given to the object (1 - Poor; 2 - Needs improvement; 3 - Satisfactory; 4 - Good; 5 - Excellent)
	FeedbackForm   FeedbackForm   `json:"feedback_form,omitempty"`
	FeedbackId     FeedbackIdType `json:"feedback_id,omitempty"`
}

func (r *Rating) Validate() error {

	if r.Value < 1 || r.Value > 5 {
		return NewValidationError("Rating value should be between 1 and 5")
	}

	return nil
}
