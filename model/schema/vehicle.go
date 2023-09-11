package schema

/*
   Vehicle:
     description: Describes the properties of a vehicle used in a mobility service
     type: object
     properties:
       category:
         type: string
       capacity:
         type: integer
       make:
         type: string
       model:
         type: string
       size:
         type: string
       variant:
         type: string
       color:
         type: string
       energy_type:
         type: string
       registration:
         type: string

*/

type Vehicle struct {
	Category     string `json:"category,omitempty"`     // Category of the vehicle
	Capacity     int    `json:"capacity,omitempty"`     // Capacity of the vehicle
	Make         string `json:"make,omitempty"`         // Make of the vehicle
	Model        string `json:"model,omitempty"`        // Model of the vehicle
	Size         string `json:"size,omitempty"`         // Size of the vehicle
	Variant      string `json:"variant,omitempty"`      // Variant of the vehicle
	Color        string `json:"color,omitempty"`        // Color of the vehicle
	EnergyType   string `json:"energy_type,omitempty"`  // Energy type of the vehicle
	Registration string `json:"registration,omitempty"` // Registration number of the vehicle
}
