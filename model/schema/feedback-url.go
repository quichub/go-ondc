/*

Create code for the following open api scehma

FeedbackUrl:
     description: Describes how a feedback URL will be sent by the Seller App
     type: object
     properties:
       url:
         description: feedback URL sent by the Seller App
         type: string
         format: uri
       tl_method:
         type: string
         enum:
           - http/get
           - http/post
       params:
         type: object
         properties:
           feedback_id:
             type: string
             description: This value will be placed in the the $feedback_id url param in case of http/get and in the requestBody http/post requests
         additionalProperties:
           type: string
         required:
           - feedback_id
*/

package schema

import "encoding/json"

type TlMethodType string

type FeedbackIdType string

const (
	HttpGetTlMethod  TlMethodType = "http/get"
	HttpPostTlMethod TlMethodType = "http/post"
)

type FeedbackUrl struct {
	Url      Uri          `json:"url,omitempty"` // feedback URL sent by the Seller App
	TlMethod TlMethodType `json:"tl_method,omitempty"`
	Params   ParamsType   `json:"params,omitempty"`
}

type ParamsType struct {
	FeedbackId           FeedbackIdType
	AdditionalProperties map[string]string
}

//custom marshal unmarshall ParamsType as json with additionalProperiies as key value pairs

func (p ParamsType) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	for k, v := range p.AdditionalProperties {
		m[k] = v
	}
	m["feedback_id"] = string(p.FeedbackId)
	return json.Marshal(m)
}

func (p *ParamsType) UnmarshalJSON(data []byte) error {
	m := make(map[string]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	p.AdditionalProperties = make(map[string]string)
	for k, v := range m {
		if k == "feedback_id" {
			p.FeedbackId = FeedbackIdType(v)
		} else {
			p.AdditionalProperties[k] = v
		}
	}
	return nil
}
