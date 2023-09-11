package schema

/*
   Tracking:
     description: Contains tracking information that can be used by the BAP to track the fulfillment of an order in real-time. which is useful for knowing the location of time sensitive deliveries.
     type: object
     properties:
       id:
         description: A unique tracking reference number
         type: string
       url:
         description: 'A URL to the tracking endpoint. This can be a link to a tracking webpage, a webhook URL created by the BAP where BPP can push the tracking data, or a GET url creaed by the BPP which the BAP can poll to get the tracking data. It can also be a websocket URL where the BPP can push real-time tracking data.'
         type: string
         format: uri
       location:
         description: 'In case there is no real-time tracking endpoint available, this field will contain the latest location of the entity being tracked. The BPP will update this value everytime the BAP calls the track API.'
         allOf:
           - $ref: '#/components/schemas/Location'
       status:
         description: 'This value indicates if the tracking is currently active or not. If this value is `active`, then the BAP can begin tracking the order. If this value is `inactive`, the tracking URL is considered to be expired and the BAP should stop tracking the order.'
         type: string
         enum:
           - active
           - inactive
       tags:
         $ref: '#/components/schemas/TagGroup'

*/

type TrackingStatus string

const (
	ActiveTracking   TrackingStatus = "active"
	InactiveTracking TrackingStatus = "inactive"
)

type Tracking struct {
	Id       string         `json:"id,omitempty"`       // A unique tracking reference number
	Url      Uri            `json:"url,omitempty"`      // A URL to the tracking endpoint. This can be a link to a tracking webpage, a webhook URL created by the BAP where BPP can push the tracking data, or a GET url creaed by the BPP which the BAP can poll to get the tracking data. It can also be a websocket URL where the BPP can push real-time tracking data.
	Location Location       `json:"location,omitempty"` // In case there is no real-time tracking endpoint available, this field will contain the latest location of the entity being tracked. The BPP will update this value everytime the BAP calls the track API.
	Status   TrackingStatus `json:"status,omitempty"`   // This value indicates if the tracking is currently active or not. If this value is `active`, then the BAP can begin tracking the order. If this value is `inactive`, the tracking URL is considered to be expired and the BAP should stop tracking the order.
	Tags     TagGroup       `json:"tags,omitempty"`
}
