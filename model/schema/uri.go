package schema

type Uri string // URI

func NewUri(value string) Uri {
	return Uri(value)
}

func (u Uri) Value() string {
	return string(u)
}
