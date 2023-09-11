package schema

/*
   Form:
     description: Describes a form
     type: object
     properties:
       url:
         description: 'The URL from where the form can be fetched. The content fetched from the url must be processed as per the mime_type specified in this object. Once fetched, the rendering platform can choosed to render the form as-is as an embeddable element; or process it further to blend with the theme of the application. In case the interface is non-visual, the the render can process the form data and reproduce it as per the standard specified in the form.'
         type: string
         format: uri
       data:
         description: The form content string. This content will again follow the mime_type field for processing. Typically forms should be sent as an html string starting with <form></form> tags. The application must render this form after removing any css or javascript code if necessary. The `action` attribute in the form should have a url where the form needs to be submitted.
         type: string
       mime_type:
         description: This field indicates the nature and format of the form received by querying the url. MIME types are defined and standardized in IETF's RFC 6838.
         type: string
*/

type Form struct {
	Url      Uri    `json:"url,omitempty"`       // The URL from where the form can be fetched. The content fetched from the url must be processed as per the mime_type specified in this object. Once fetched, the rendering platform can choosed to render the form as-is as an embeddable element; or process it further to blend with the theme of the application. In case the interface is non-visual, the the render can process the form data and reproduce it as per the standard specified in the form.
	Data     string `json:"data,omitempty"`      // The form content string. This content will again follow the mime_type field for processing. Typically forms should be sent as an html string starting with <form></form> tags. The application must render this form after removing any css or javascript code if necessary. The `action` attribute in the form should have a url where the form needs to be submitted.
	MimeType string `json:"mime_type,omitempty"` // This field indicates the nature and format of the form received by querying the url. MIME types are defined and standardized in IETF's RFC 6838.
}
