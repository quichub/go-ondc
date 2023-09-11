package operations

import "ondc/model/schema"

/*
OpenAPI schema

	schema:
	  type: object
	  properties:
	    context:
	      $ref: '#/components/schemas/Context'
	    message:
	      type: object
	      properties:
	        intent:
	          $ref: '#/components/schemas/Intent'
	  required:
	    - context
	    - message
*/

type SearchRequestMessage struct {
	Intent schema.Intent `json:"intent"`
}

type SearchRequest struct {
	Context schema.Context       `json:"context"`
	Message SearchRequestMessage `json:"message"`
}

type SearchResponse struct {
}
