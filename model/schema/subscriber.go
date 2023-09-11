/*
   Subscriber:
     type: object
     description: 'Any entity which wants to authenticate itself on a network. This can be a Buyer App, Seller App or Gateway.'
     properties:
       subscriber_id:
         type: string
         description: Registered domain name of the subscriber. Must have a valid SSL certificate issued by a Certificate Authority of the operating region
       type:
         type: string
         enum:
           - bap
           - bpp
           - bg
       cb_url:
         type: string
         description: Callback URL of the subscriber. The Registry will call this URL's on_subscribe API to validate the subscriber\'s credentials
       domain:
         $ref: '#/components/schemas/Domain'
       city:
         $ref: '#/components/schemas/City/properties/code'
       country:
         $ref: '#/components/schemas/Country/properties/code'
       signing_public_key:
         type: string
         description: 'Signing Public key of the subscriber. <br/><br/>Any subscriber platform (Buyer App, Seller App, Gateway) who wants to transact on the network must digitally sign the ```requestBody``` using the corresponding private key of this public key and send it in the transport layer header. In case of ```HTTP``` it is the ```Authorization``` header. <br><br/>The ```Authorization``` will be used to validate the signature of a Buyer App or Seller App.<br/><br/>Furthermore, if an API call is being proxied or multicast by a ONDC Gateway, the Gateway must use it\''s signing key to digitally sign the ```requestBody``` using the corresponding private key of this public key and send it in the ```X-Gateway-Authorization``` header.'
       encryption_public_key:
         type: string
         description: Encryption public key of the Buyer App. Any Seller App must encrypt the ```requestBody.message``` value of the ```on_search``` API using this public key.
       status:
         type: string
         enum:
           - INITIATED
           - UNDER_SUBSCRIPTION
           - SUBSCRIBED
           - INVALID_SSL
           - UNSUBSCRIBED
       created:
         type: string
         description: Timestamp when a subscriber was added to the registry with status = INITIATED
         format: date-time
       updated:
         type: string
         format: date-time
       expires:
         type: string
         description: Expiry timestamp in UTC derived from the ```lease_time``` of the subscriber
         format: date-time
*/

package schema

//define all enum types in above open api schema

type StatusType string

const (
	StatusTypeInitiated         StatusType = "INITIATED"
	StatusTypeUnderSubscription StatusType = "UNDER_SUBSCRIPTION"
	StatusTypeSubscribed        StatusType = "SUBSCRIBED"
	StatusTypeInvalidSSL        StatusType = "INVALID_SSL"
	StatusTypeUnsubscribed      StatusType = "UNSUBSCRIBED"
)

type SubscriberType string

const (
	SubscriberTypeBap SubscriberType = "bap"
	SubscriberTypeBpp SubscriberType = "bpp"
	SubscriberTypeBg  SubscriberType = "bg"
)

type Subscriber struct {
	SubscriberID        string         `json:"subscriber_id,omitempty"` // Registered domain name of the subscriber. Must have a valid SSL certificate issued by a Certificate Authority of the operating region
	Type                SubscriberType `json:"type,omitempty"`          // Type of subscriber
	CbURL               string         `json:"cb_url,omitempty"`        // Callback URL of the subscriber. The Registry will call this URL's on_subscribe API to validate the subscriber's credentials
	Domain              Domain         `json:"domain,omitempty"`        // Describes a domain
	City                CityCode       `json:"city,omitempty"`          // Describes a city
	Country             CountryCode    `json:"country,omitempty"`       // Describes a country
	SigningPublicKey    string         `json:"signing_public_key"`
	EncryptionPublicKey string         `json:"encryption_public_key"`
	Status              StatusType     `json:"status,omitempty"`  // Status of the subscriber
	Created             DateTime       `json:"created,omitempty"` // Timestamp when a subscriber was added to the registry with status = INITIATED
	Updated             DateTime       `json:"updated,omitempty"`
	Expires             DateTime       `json:"expires,omitempty"` // Expiry timestamp in UTC derived from the lease_time of the subscriber
}
