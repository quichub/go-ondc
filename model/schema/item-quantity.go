/*
   ItemQuantity:
     description: Describes count or amount of an item
     type: object
     properties:
       allocated:
         type: object
         properties:
           count:
             type: integer
             minimum: 0
           measure:
             $ref: '#/components/schemas/Scalar'
       available:
         type: object
         properties:
           count:
             type: integer
             minimum: 0
           measure:
             $ref: '#/components/schemas/Scalar'
       maximum:
         type: object
         properties:
           count:
             type: integer
             minimum: 1
           measure:
             $ref: '#/components/schemas/Scalar'
       minimum:
         type: object
         properties:
           count:
             type: integer
             minimum: 0
           measure:
             $ref: '#/components/schemas/Scalar'
       selected:
         type: object
         properties:
           count:
             type: integer
             minimum: 0
           measure:
             $ref: '#/components/schemas/Scalar'
       unitized:
         description: This represents the quantity available in a single unit of the item
         type: object
         properties:
           count:
             type: integer
             minimum: 1
             maximum: 1
           measure:
             $ref: '#/components/schemas/Scalar'
*/

package schema

type Quantity struct {
	Count   int     `json:"count,omitempty"`
	Measure *Scalar `json:"measure,omitempty"`
}

type ItemQuantitySelected Quantity

type ItemQuantity struct {
	Allocated Quantity `json:"allocated,omitempty"`
	Available Quantity `json:"available,omitempty"`
	Maximum   Quantity `json:"maximum,omitempty"`
	Minimum   Quantity `json:"minimum,omitempty"`
	Selected  Quantity `json:"selected,omitempty"`
	Unitized  Quantity `json:"unitized,omitempty"`
}

func (iq *ItemQuantity) Validate() error {

	//validate Allocated
	if iq.Allocated.Count < 0 {
		return NewValidationError("Allocated count should be greater than or equal to 0")
	}

	//validate Available
	if iq.Available.Count < 0 {
		return NewValidationError("Available count should be greater than or equal to 0")
	}

	//validate Maximum
	if iq.Maximum.Count < 1 {
		return NewValidationError("Maximum count should be greater than or equal to 1")
	}

	//validate Minimum
	if iq.Minimum.Count < 0 {
		return NewValidationError("Minimum count should be greater than or equal to 0")
	}

	//validate Selected
	if iq.Selected.Count < 0 {
		return NewValidationError("Selected count should be greater than or equal to 0")
	}

	//validate Unitized
	if iq.Unitized.Count != 1 {
		return NewValidationError("Unitized count should be equal to 1")
	}

	return nil
}
