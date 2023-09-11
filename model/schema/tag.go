package schema

/*
   Tag:
     description: 'Describes a tag. This is a simple key-value store which is used to contain extended metadata. This object can be added as a property to any schema to describe extended attributes. For BAPs, tags can be sent during search to optimize and filter search results. BPPs can use tags to index their catalog to allow better search functionality. Tags are sent by the BPP as part of the catalog response in the `on_search` callback. Tags are also meant for display purposes. Upon receiving a tag, BAPs are meant to render them as name-value pairs. This is particularly useful when rendering tabular information about a product or service.'
     type: object
     properties:
       code:
         description: 'The machine-readable name of the tag. The allowed values of this property can be published at three levels namely, a) Core specification, b) industry sector-specific adaptations, and c) Network-specific adaptations. Except core, each adaptation (sector or network) should prefix a unique namespace with the allowed value.'
         type: string
       name:
         description: 'The human-readable name of the tag. This set by the BPP and rendered as-is by the BAP. Sometimes, the network policy may reserve some names for this property. Values outside the reserved values can be set by the BPP. However,the BAP may choose to rename or even ignore this value and render the output purely using the `code` property, but it is recommended for BAPs to keep the name same to avoid confusion and provide consistency.'
         type: string
       value:
         description: The value of the tag. This set by the BPP and rendered as-is by the BAP.
         type: string
       display:
         description: 'This value indicates if the tag is intended for display purposes. If set to `true`, then this tag must be displayed. If it is set to `false`, it should not be displayed. This value can override the group display value.'
         type: boolean
*/

type Tag struct {
	Code    string `json:"code,omitempty"`    // The machine-readable name of the tag. The allowed values of this property can be published at three levels namely, a) Core specification, b) industry sector-specific adaptations, and c) Network-specific adaptations. Except core, each adaptation (sector or network) should prefix a unique namespace with the allowed value.
	Name    string `json:"name,omitempty"`    // The human-readable name of the tag. This set by the BPP and rendered as-is by the BAP. Sometimes, the network policy may reserve some names for this property. Values outside the reserved values can be set by the BPP. However,the BAP may choose to rename or even ignore this value and render the output purely using the `code` property, but it is recommended for BAPs to keep the name same to avoid confusion and provide consistency.
	Value   string `json:"value,omitempty"`   // The value of the tag. This set by the BPP and rendered as-is by the BAP.
	Display bool   `json:"display,omitempty"` // This value indicates if the tag is intended for display purposes. If set to `true`, then this tag must be displayed. If it is set to `false`, it should not be displayed. This value can override the group display value.
}
