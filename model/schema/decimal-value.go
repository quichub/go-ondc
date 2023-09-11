package schema

import (
	"fmt"
	"strconv"
)

/*
   DecimalValue:
     description: Describes a decimal value
     type: string
     pattern: '[+-]?([0-9]*[.])?[0-9]+'
*/

type DecimalValue float64

func NewDecimalValue(value float64) DecimalValue {
	return DecimalValue(value)
}

func (d DecimalValue) Value() float64 {
	return float64(d)
}

func (d DecimalValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%f", d.Value())), nil
}

func (d *DecimalValue) UnmarshalJSON(data []byte) error {
	//parse from string to float64
	val, err := strconv.ParseFloat(string(data), 64)
	if err == nil {
		*d = DecimalValue(val)
	}

	return err
}
