package schema

import "ondc/model"

/*
   Context:
     description: Describes a ONDC message context
     type: object
     properties:
       domain:
         $ref: '#/components/schemas/Domain'
       country:
         $ref: '#/components/schemas/Country/properties/code'
       city:
         $ref: '#/components/schemas/City/properties/code'
       action:
         type: string
         description: Defines the ONDC API call. Any actions other than the enumerated actions are not supported by ONDC Protocol
         enum:
           - search
           - select
           - init
           - confirm
           - update
           - status
           - track
           - cancel
           - rating
           - support
           - on_search
           - on_select
           - on_init
           - on_confirm
           - on_update
           - on_status
           - on_track
           - on_cancel
           - on_rating
           - on_support
       core_version:
         type: string
         description: Version of ONDC core API specification being used
       bap_id:
         type: string
         description: Unique id of the Buyer App. By default it is the fully qualified domain name of the Buyer App
       bap_uri:
         type: string
         format: uri
         description: URI of the Buyer App for accepting callbacks. Must have the same domain name as the bap_id
       bpp_id:
         type: string
         description: Unique id of the Seller App. By default it is the fully qualified domain name of the Seller App,
                      mandatory for all peer-to-peer API requests, i.e. except search and on_search
       bpp_uri:
         type: string
         format: uri
         description: URI of the Seller App. Must have the same domain name as the bap_id,
                      mandatory for all peer-to-peer API requests, i.e. except search and on_search
       transaction_id:
         type: string
         description: This is a unique value which persists across all API calls from search through confirm
       message_id:
         type: string
         description: This is a unique value which persists during a request / callback cycle
       timestamp:
         type: string
         format: date-time
         description: Time of request generation in RFC3339 format
       key:
         type: string
         description: The encryption public key of the sender
       ttl:
         type: string
         description: Timestamp for which this message holds valid in ISO8601 durations format -
                      Outer limit for ttl for search, select, init, confirm, status, track, cancel, update, rating, support is 'PT30S' which is 30 seconds,
                      different buyer apps can change this to meet their UX requirements, but it shouldn't exceed this outer limit
     required:
       - domain
       - action
       - core_version
       - bap_id
       - bap_uri
       - transaction_id
       - message_id
       - city
       - country
       - timestamp
*/

type Context struct {
	Domain        Domain      `json:"domain"`
	Country       CountryCode `json:"country"`
	City          CityCode    `json:"city"`
	Action        ActionType  `json:"action"`
	CoreVersion   string      `json:"core_version"`
	BapId         string      `json:"bap_id"`
	BapUri        string      `json:"bap_uri"`
	BppId         *string     `json:"bpp_id"`
	BppUri        *string     `json:"bpp_uri"`
	TransactionId string      `json:"transaction_id"`
	MessageId     string      `json:"message_id"`
	Timestamp     Timestamp   `json:"timestamp"`
	Key           *string     `json:"key"`
	Ttl           *string     `json:"ttl"`
}

func (c Context) Validate() error {
	if c.Domain == "" {
		return model.NewApiError(400, "domain is required")
	}

	if c.Action == "" {
		return model.NewApiError(400, "action is required")
	}

	if c.CoreVersion == "" {
		return model.NewApiError(400, "core_version is required")
	}

	if c.BapId == "" {
		return model.NewApiError(400, "bap_id is required")
	}

	if c.BapUri == "" {
		return model.NewApiError(400, "bap_uri is required")
	}

	if c.TransactionId == "" {
		return model.NewApiError(400, "transaction_id is required")
	}

	if c.MessageId == "" {
		return model.NewApiError(400, "message_id is required")
	}

	return nil
}
