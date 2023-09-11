package schema

/*
   Domain:
     description: Codification of domain for ONDC
     type: string
     enum:
       - nic2004:52110
*/

type Domain string // Codification of domain for ONDC

const (
	Nic2004_52110 Domain = "nic2004:52110"
)
