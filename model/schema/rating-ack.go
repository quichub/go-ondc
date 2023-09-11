package schema

/*
RatingAck:
     type: object
     properties:
       feedback_ack:
         description: If feedback has been recorded or not
         type: boolean
       rating_ack:
         description: If rating has been recorded or not
         type: boolean
*/

type RatingAck struct {
	FeedbackAck bool `json:"feedback_ack"` // If feedback has been recorded or not
	RatingAck   bool `json:"rating_ack"`   // If rating has been recorded or not
}
