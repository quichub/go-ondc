package schema

/*
   Location:
     description: Describes the location of a runtime object.
     type: object
     properties:
       id:
         type: string
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       gps:
         $ref: '#/components/schemas/Gps'
       address:
         $ref: '#/components/schemas/Address'
       station_code:
         type: string
       city:
         $ref: '#/components/schemas/City'
       country:
         $ref: '#/components/schemas/Country'
       circle:
         $ref: '#/components/schemas/Circle'
       polygon:
         type: string
       3dspace:
         type: string
       time:
         $ref: '#/components/schemas/Time'
*/

type LocationId string

type Location struct {
	Id          LocationId `json:"id,omitempty"`
	Descriptor  Descriptor `json:"descriptor,omitempty"` // Describes the description of a real-world object.
	Gps         Gps        `json:"gps,omitempty"`
	Address     Address    `json:"address,omitempty"`
	StationCode string     `json:"station_code,omitempty"`
	City        City       `json:"city,omitempty"`
	Country     Country    `json:"country,omitempty"`
	Circle      Circle     `json:"circle,omitempty"`
	Polygon     string     `json:"polygon,omitempty"`
	Thee3dspace string     `json:"3dspace,omitempty"`
	Time        Time       `json:"time,omitempty"`
}
