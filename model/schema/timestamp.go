package schema

import "time"

// RFC3339 time stamp
type Timestamp struct {
	value time.Time
}

func NewTimestamp(value time.Time) Timestamp {
	return Timestamp{value}
}

func (t Timestamp) Value() time.Time {
	return t.value
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	//to RFC3339 format
	return []byte(t.value.Format(time.RFC3339)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	//from RFC3339 format
	val, err := time.Parse(time.RFC3339, string(data))
	if err == nil {
		t.value = val
	}

	return err
}
