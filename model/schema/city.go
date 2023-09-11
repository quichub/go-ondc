package schema

/*
   City:
     description: Describes a city
     type: object
     properties:
       name:
         type: string
         description: Name of the city
       code:
         type: string
         description: Codification of city code will be using the std code of the city e.g. for Bengaluru, city code is 'std:080'
*/

type CityCode string

type City struct {
	Name string   `json:"name"` // Name of the city
	Code CityCode `json:"code"` // Codification of city code will be using the std code of the city e.g. for Bengaluru, city code is 'std:080'
}
