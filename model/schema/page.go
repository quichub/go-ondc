package schema

/*
   Page:
     description: Describes a page in a search result
     type: object
     properties:
       id:
         type: string
       next_id:
         type: string
*/

type Page struct {
	Id     string `json:"id"`
	NextId string `json:"next_id"`
}
