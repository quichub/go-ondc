package schema

import (
	"errors"
	"fmt"
	"strings"
)

/*

   Name:
     type: string
     description: 'Describes the name of a person in format: ./{given_name}/{honorific_prefix}/{first_name}/{middle_name}/{last_name}/{honorific_suffix}'
*/
//pattern: '^\./[^/]+/[^/]*/[^/]*/[^/]*/[^/]*/[^/]*$'

type Name struct {
	GivenName       string
	HonorificPrefix string
	FirstName       string
	MiddleName      string
	LastName        string
	HonorificSuffix string
}

func NewName(givenName string, honorificPrefix string, firstName string, middleName string, lastName string, honorificSuffix string) Name {
	return Name{
		GivenName:       givenName,
		HonorificPrefix: honorificPrefix,
		FirstName:       firstName,
		MiddleName:      middleName,
		LastName:        lastName,
		HonorificSuffix: honorificSuffix,
	}
}

func (n Name) Value() string {
	return fmt.Sprintf("./%s/%s/%s/%s/%s/%s", n.GivenName, n.HonorificPrefix, n.FirstName, n.MiddleName, n.LastName, n.HonorificSuffix)
}

func (n Name) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", n.Value())), nil
}

func (n *Name) UnmarshalJSON(data []byte) error {
	// split the string into parts
	parts := strings.Split(string(data), "/")
	if len(parts) != 7 {
		return errors.New("invalid name format")
	}

	// assign the parts
	n.GivenName = parts[1]
	n.HonorificPrefix = parts[2]
	n.FirstName = parts[3]
	n.MiddleName = parts[4]
	n.LastName = parts[5]
	n.HonorificSuffix = parts[6]

	return nil
}
