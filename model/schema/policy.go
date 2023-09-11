package schema

/*
   Policy:
     description: Describes a policy. Allows for domain extension.
     type: object
     properties:
       id:
         type: string
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       parent_policy_id:
         $ref: '#/components/schemas/Policy/properties/id'
       time:
         $ref: '#/components/schemas/Time'
*/

type PolicyId string

type Policy struct {
	Id             string     `json:"id,omitempty"`               // ID of the policy. This follows the syntax {item.id}/policy/{policy unique id} for item specific policy OR
	Descriptor     Descriptor `json:"descriptor,omitempty"`       // Describes the description of a real-world object.
	ParentPolicyId PolicyId   `json:"parent_policy_id,omitempty"` // ID of the parent policy. This follows the syntax {item.id}/policy/{policy unique id} for item specific policy OR
	Time           Time       `json:"time,omitempty"`             // Describes time
}
