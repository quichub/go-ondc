/*
   Feedback:
     description: Feedback for a service
     type: object
     properties:
       feedback_form:
         $ref: '#/components/schemas/FeedbackForm'
       feedback_url:
         $ref: '#/components/schemas/FeedbackUrl'
*/

package schema

type Feedback struct {
	FeedbackForm FeedbackForm `json:"feedback_form,omitempty"` // Feedback for a service
	FeedbackUrl  FeedbackUrl  `json:"feedback_url,omitempty"`  // Feedback for a service
}
