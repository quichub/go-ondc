package schema

/*
	Scalar:
     description: An object representing a scalar quantity.
     type: object
     properties:
       type:
         type: string
         enum:
           - CONSTANT
           - VARIABLE
       value:
         type: number
       estimated_value:
         type: number
       computed_value:
         type: number
       range:
         type: object
         properties:
           min:
             type: number
           max:
             type: number
       unit:
         type: string
     required:
       - value
       - unit
*/

type ScalarType string

const (
	ConstScalar    ScalarType = "CONSTANT"
	VariableScalar ScalarType = "VARIABLE"
)

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Scalar struct {
	Type           *ScalarType `json:"type"`
	Value          float64     `json:"value"`
	EstimatedValue *float64    `json:"estimated_value"`
	ComputedValue  *float64    `json:"computed_value"`
	Range          *Range      `json:"range"`
	Unit           string      `json:"unit"`
}
