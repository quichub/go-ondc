/*
FeedbackFormElement:
     description: An element in the feedback form. It can be question or an answer to the question.
     type: object
     properties:
       id:
         type: string
       parent_id:
         $ref: '#/components/schemas/FeedbackFormElement/properties/id'
       question:
         description: Specifies the question to which the answer options will be contained in the child FeedbackFormElements
         type: string
       answer:
         description: Specifies an answer option to which the question will be in the FeedbackFormElement specified in parent_id
         type: string
       answer_type:
         description: Specifies how the answer option should be rendered.
         type: string
         enum :
           - radio
           - checkbox
           - text
*/

package schema

type AnswerType string

const (
	RadioAnswerType    AnswerType = "radio"
	CheckboxAnswerType AnswerType = "checkbox"
	TextAnswerType     AnswerType = "text"
)

type FeedbackFormElement struct {
	Id         string         `json:"id,omitempty"`
	ParentId   FeedbackFormId `json:"parent_id,omitempty"`
	Question   string         `json:"question,omitempty"`    // Specifies the question to which the answer options will be contained in the child FeedbackFormElements
	Answer     string         `json:"answer,omitempty"`      // Specifies an answer option to which the question will be in the FeedbackFormElement specified in parent_id
	AnswerType AnswerType     `json:"answer_type,omitempty"` // Specifies how the answer option should be rendered.
}
