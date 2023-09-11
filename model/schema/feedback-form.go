/*
   FeedbackForm:
     description: Describes a feedback form that a Seller App can send to get feedback from the Buyer App
     type: array
     items:
       $ref: '#/components/schemas/FeedbackFormElement'
*/

package schema

type FeedbackFormId string

type FeedbackForm []FeedbackFormElement
