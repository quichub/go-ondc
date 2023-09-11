package schema

/*
   XInputResponse:
     description: The response to the form fetched via the XInput URL
     type: array
     items:
       type: object
       properties:
         input:
           description: The _name_ attribute of the input tag in the XInput form
           type: string
         value:
           description: 'The value of the input field. Files must be sent as data URLs. For more information on Data URLs visit https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs'
           type: string
*/

type XInputResponseItem struct {
	Input string `json:"input"` // The _name_ attribute of the input tag in the XInput form
	Value string `json:"value"` // The value of the input field. Files must be sent as data URLs. For more information on Data URLs visit https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs
}

type XInputResponse []XInputResponseItem
